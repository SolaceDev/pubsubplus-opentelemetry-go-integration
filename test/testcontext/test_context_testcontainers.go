// pubsubplus-opentelemetry-go-integration
//
// Copyright 2024 Solace Corporation. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build !remote
// +build !remote

package testcontext

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/google/uuid"
	"github.com/testcontainers/testcontainers-go"
)

// the default config file relative to the location the tests are run from
// this assumes that the tests are run from the root of the `test` directory
const testContainersConfig = "./data/config/config_testcontainers.json"
const defaultComposeFilePath = "data/compose/docker-compose.yml"

type testContainersTestContext struct {
	testContextCommon
	compose *testcontainers.LocalDockerCompose
}

func getTestContext() testContext {
	context := testContainersTestContext{}
	return &context
}

// SetupTestContext sets up new context
func (context *testContainersTestContext) Setup() error {
	fmt.Println()
	fmt.Println("-- TESTCONTAINERS SETUP --")

	context.setupCommon(testContainersConfig)

	identifier := strings.ToLower(uuid.New().String())
	composeFilePaths := []string{defaultComposeFilePath}
	compose := testcontainers.NewLocalDockerCompose(composeFilePaths, identifier)

	context.compose = compose

	context.semp = newSempV2(context.config.SEMP)
	context.toxi = newToxiProxy(context.config.ToxiProxy)

	fmt.Println("Starting dockerized broker via docker-compose")

	execError := context.compose.
		WithEnv(context.config.ToEnvironment()).
		WithCommand([]string{"up", "-d"}).
		Invoke()

	err := execError.Error
	if err != nil {
		return err
	}

	fmt.Println("Docker-compose command succeeded")
	fmt.Println("Waiting for SEMPv2 to be up")

	err = context.waitForSEMP()
	if err != nil {
		return err
	}

	fmt.Println("Waiting for Messaging to be up")
	err = context.waitForMessaging()
	if err != nil {
		return err
	}

	err = context.semp.setup()
	if err != nil {
		return err
	}

	fmt.Println("Waiting for ToxiProxy to be up")
	err = context.waitForToxiProxy()
	if err != nil {
		return err
	}

	fmt.Println("Setting up ToxiProxy proxies")
	err = context.toxi.setup()
	if err != nil {
		return err
	}

	fmt.Println("Waiting for MsgVPN " + context.config.Messaging.VPN + " to be up")
	err = context.waitForVPNState("up")
	if err != nil {
		return err
	}

	fmt.Println("-- TESTCONTAINERS SETUP COMPLETE --")
	fmt.Println()
	return nil
}

// TeardownTestContext impl of test context
func (context *testContainersTestContext) Teardown() error {
	fmt.Println()
	fmt.Println("-- TESTCONTAINERS TEARDOWN --")
	pubsubHostname := context.config.TestContainers.BrokerHostname
	err := context.toxi.teardown()
	if err != nil {
		fmt.Println("Encountered error tearing down toxiproxy: " + err.Error())
	}
	rc, output, err := context.dockerLogs(pubsubHostname)
	if err == nil {
		fmt.Println("Broker logs for " + pubsubHostname + ":")
		fmt.Println(output)
	} else {
		fmt.Println("Encountered error getting docker logs from " + pubsubHostname + ": rc=" + string(rc) + " err:" + err.Error())
	}
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Encountered error getting working directory for " + pubsubHostname + " diagnostics err:" + err.Error())
	}

	err = context.gatherBrokerDiagnostics(path.Join(wd, "diagnostics.tgz"))
	if err != nil {
		fmt.Println("Encountered error getting " + pubsubHostname + " diagnostics err:" + err.Error())
	}

	composeResult := context.compose.Down()
	if composeResult.Error != nil {
		err = composeResult.Error
		fmt.Println("Encountered error tearing down docker compose: " + composeResult.Error.Error())
	}
	fmt.Println("-- TESTCONTAINERS TEARDOWN COMPLETE --")
	return err
}

func (context *testContainersTestContext) gatherBrokerDiagnostics(destinationPath string) error {
	fmt.Println()
	pubsubHostname := context.config.TestContainers.BrokerHostname
	// gather all important infomation and logs from pubsubHostname container
	fmt.Println("Run gather-diagnostics for " + pubsubHostname + "...")
	resp, output, err := context.dockerExec(pubsubHostname, []string{"/bin/bash", "-l", "-c", "\"gather-diagnostics\""})
	if err != nil {
		return err
	}
	fmt.Println(output)
	if resp != 0 {
		return fmt.Errorf("failed to gather %s diagnostics", pubsubHostname)
	}
	fmt.Println("Gathered gather-diagnostics for " + pubsubHostname)
	// extract diagnostic to host
	// first get absolute path from container
	resp, diagnosticPath, err := context.dockerExec(pubsubHostname, []string{"/bin/bash", "-l", "-c", "ls -rt /usr/sw/jail/logs/gather-diagnostics*.tgz | tail -n 1"})
	//resp, diagnosticPath, err := context.dockerExec(pubsubHostname, []string{"/bin/bash", "-l", "-c", " realpath $(ls -rt /usr/sw/jail/logs/gather-diagnostics*.tgz | tail -n 1)"})
	if err != nil {
		return err
	}
	if resp != 0 {
		return fmt.Errorf("Failed to locate %s diagnostics", pubsubHostname)
	}
	fmt.Println("Exacting gather-diagnostics " + diagnosticPath + " for " + pubsubHostname + " to " + destinationPath + "...")
	err = context.dockerCpToHost(pubsubHostname, strings.TrimSpace(diagnosticPath), destinationPath)
	return err
}

const dockerCmd = "docker"

// dockerCp copies a file from one container to another container
func (ctx *testContainersTestContext) dockerCp(sourceContainer, sourcePath, destContainer, destPath string) error {
	tmpFileName := path.Join(os.TempDir(), "tmp-go-docker-cp")
	defer os.Remove(tmpFileName)
	cmd := exec.Command(dockerCmd, "cp", sourceContainer+":"+sourcePath, tmpFileName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(output))
		return fmt.Errorf("encountered error running docker cp to local from %s: %w", sourceContainer, err)
	}
	err = os.Chmod(tmpFileName, 0655)
	if err != nil {
		return fmt.Errorf("encountered error running chmod on %s: %w", tmpFileName, err)
	}
	cmd = exec.Command(dockerCmd, "cp", tmpFileName, destContainer+":"+destPath)
	output, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(output))
		return fmt.Errorf("encountered error running docker cp from local to %s: %w", destContainer, err)
	}
	return nil
}

func (ctx *testContainersTestContext) dockerCpToHost(sourceContainer, sourcePath, destPath string) error {
	hostFileName := destPath
	fmt.Println("Copying source file " + sourceContainer + ":" + sourcePath + " to " + hostFileName)
	cmd := exec.Command(dockerCmd, "cp", "-L", sourceContainer+":"+sourcePath, hostFileName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(output))
		return fmt.Errorf("encountered error running docker cp to local from %s: %w", sourceContainer, err)
	}
	err = os.Chmod(hostFileName, 0655)
	if err != nil {
		return fmt.Errorf("encountered error running chmod on %s: %w", hostFileName, err)
	}
	return nil
}

// dockerExec runs the given command on the given container using the given user if not empty
func (ctx *testContainersTestContext) dockerExec(container string, command []string) (int, string, error) {
	args := append([]string{"exec", container}, command...)
	cmd := exec.Command(dockerCmd, args...)
	output, err := cmd.CombinedOutput()
	return cmd.ProcessState.ExitCode(), string(output), err
}

// dockerLogs runs docker log on given container
func (ctx *testContainersTestContext) dockerLogs(container string) (int, string, error) {
	args := []string{"logs", container}
	cmd := exec.Command(dockerCmd, args...)
	output, err := cmd.CombinedOutput()
	return cmd.ProcessState.ExitCode(), string(output), err
}
