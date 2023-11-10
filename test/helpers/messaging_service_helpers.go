// pubsubplus-opentelemetry-go-integration
//
// Copyright 2023-2023 Solace Corporation. All rights reserved.
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

// Package helpers contains various test helpers for assertions, for example it contains a helper to connect and disconnect a messaging service
// based on a messaging service builder.
package helpers

import (
	"fmt"
	"net/url"
	"time"

	"solace.dev/go/messaging/pkg/solace"
	"solace.dev/go/messaging/pkg/solace/config"
	"solace.dev/go/trace/test/sempclient/monitor"
	"solace.dev/go/trace/test/testcontext"

	//lint:ignore ST1001 dot import is fine for tests
	. "github.com/onsi/ginkgo/v2"
	//lint:ignore ST1001 dot import is fine for tests
	. "github.com/onsi/gomega"
)

// Helper functions for various connect/disconnect functionality, paramtereizing connects and disconnects

const defaultTimeout = 10 * time.Second

// ConnectFunction type defined
type ConnectFunction func(service solace.MessagingService) error

// SynchronousConnectFunction defined
var SynchronousConnectFunction = func(service solace.MessagingService) error {
	return service.Connect()
}

// AsynchronousConnectFunction defined
var AsynchronousConnectFunction = func(service solace.MessagingService) error {
	select {
	case err := <-service.ConnectAsync():
		return err
	case <-time.After(defaultTimeout):
		Fail("timed out waiting for async connect to return")
		return nil
	}
}

// CallbackConnectFunction defined
var CallbackConnectFunction = func(service solace.MessagingService) error {
	var passedService solace.MessagingService
	cErr := make(chan error)
	service.ConnectAsyncWithCallback(func(newService solace.MessagingService, err error) {
		passedService = newService
		cErr <- err
	})
	select {
	case err := <-cErr:
		ExpectWithOffset(1, passedService).To(Equal(service), "Messaging service passed to callback did not equal calling service")
		return err
	case <-time.After(defaultTimeout):
		Fail("timed out waiting for callback connect to return")
		return nil
	}
}

// ConnectFunctions defined
var ConnectFunctions = map[string]ConnectFunction{
	"synchronous connect":  SynchronousConnectFunction,
	"asynchronous connect": AsynchronousConnectFunction,
	"callback connect":     CallbackConnectFunction,
}

// DisconnectFunction defined
type DisconnectFunction func(service solace.MessagingService) error

// SynchronousDisconnectFunction defined
var SynchronousDisconnectFunction = func(service solace.MessagingService) error {
	return service.Disconnect()
}

// AsynchronousDisconnectFunction defined
var AsynchronousDisconnectFunction = func(service solace.MessagingService) error {
	select {
	case err := <-service.DisconnectAsync():
		return err
	case <-time.After(defaultTimeout):
		Fail("timed out waiting for async disconnect to return")
		return nil
	}
}

// CallbackDisconnectFunction defined
var CallbackDisconnectFunction = func(service solace.MessagingService) error {
	cErr := make(chan error)
	service.DisconnectAsyncWithCallback(func(err error) {
		cErr <- err
	})
	select {
	case err := <-cErr:
		return err
	case <-time.After(defaultTimeout):
		Fail("timed out waiting for callback disconnect to return")
		return nil
	}
}

// DisconnectFunctions defined
var DisconnectFunctions = map[string]DisconnectFunction{
	"synchronous disconnect":  SynchronousDisconnectFunction,
	"asynchronous disconnect": AsynchronousDisconnectFunction,
	"callback disconnect":     CallbackDisconnectFunction,
}

// DefaultConfiguration defined
func DefaultConfiguration() config.ServicePropertyMap {
	connectionDetails := testcontext.Messaging()
	url := fmt.Sprintf("%s:%d", connectionDetails.Host, connectionDetails.MessagingPorts.PlaintextPort)
	config := config.ServicePropertyMap{
		config.ServicePropertyVPNName:                     connectionDetails.VPN,
		config.TransportLayerPropertyHost:                 url,
		config.AuthenticationPropertySchemeBasicUserName:  connectionDetails.Authentication.BasicUsername,
		config.AuthenticationPropertySchemeBasicPassword:  connectionDetails.Authentication.BasicPassword,
		config.TransportLayerPropertyReconnectionAttempts: 0,
	}
	return config
}

// helper to connect and disconnect the messaging service

// BuildMessagingService function
func BuildMessagingService(builder solace.MessagingServiceBuilder) solace.MessagingService {
	return buildMessagingService(builder, 2)
}

// ConnectMessagingService function
func ConnectMessagingService(messagingService solace.MessagingService) *monitor.MsgVpnClient {
	return connectMessagingServiceWithFunction(messagingService, SynchronousConnectFunction, 2)
}

// DisconnectMessagingService function
func DisconnectMessagingService(messagingService solace.MessagingService) {
	disconnectMessagingServiceWithFunction(messagingService, SynchronousDisconnectFunction, 2)
}

func buildMessagingService(builder solace.MessagingServiceBuilder, callDepth int) solace.MessagingService {
	messagingService, err := builder.Build()
	ExpectWithOffset(callDepth, err).To(BeNil(), "Got unexpected error while building messaging service")
	return messagingService
}

func connectMessagingServiceWithFunction(messagingService solace.MessagingService, connectFunction ConnectFunction, callDepth int) *monitor.MsgVpnClient {
	success := false
	// make sure that when we connect we defer the disconnect
	defer func() {
		if !success {
			// make sure that if we fail, we disconnect the messaging service
			messagingService.Disconnect()
		}
	}()
	err := connectFunction(messagingService)
	ExpectWithOffset(callDepth, err).To(BeNil(), "Got unexpected error when connecting messaging service")

	client := getClient(messagingService, callDepth+1)
	ExpectWithOffset(callDepth, client.ClientName).To(Equal(messagingService.GetApplicationID()), "Expected ClientName to equal application ID when checking for connected client")
	success = true
	return client
}

func disconnectMessagingServiceWithFunction(messagingService solace.MessagingService, disconnectFunction func(solace.MessagingService) error, callDepth int) {
	err := disconnectFunction(messagingService)
	ExpectWithOffset(callDepth, err).To(BeNil(), "Got unexpected error when disconnecting messaging service")

	resp, _, err := testcontext.SEMP().Monitor().MsgVpnApi.
		GetMsgVpnClient(testcontext.SEMP().MonitorCtx(), testcontext.Messaging().VPN, "notexist", nil)
	ExpectWithOffset(callDepth, err).To(HaveOccurred(), "Expected SEMP to reject call checking for connected client after disconnecting client")
	decodeMonitorSwaggerError(err, &resp, callDepth+1)
	ExpectWithOffset(callDepth, resp.Meta).ToNot(BeNil(), "Expected response from SEMP to contain meta")
	ExpectWithOffset(callDepth, resp.Meta.Error_).ToNot(BeNil(), "Expected response from SEMP to contain an error")
	ExpectWithOffset(callDepth, resp.Meta.Error_.Status).To(Equal("NOT_FOUND"), "Expected response from SEMP to have error status NOT_FOUND")
}

// GetClient function
func GetClient(messagingService solace.MessagingService) *monitor.MsgVpnClient {
	return getClient(messagingService, 2)
}

func getClient(messagingService solace.MessagingService, callDepth int) *monitor.MsgVpnClient {
	clientResponse, _, err := testcontext.SEMP().Monitor().MsgVpnApi.
		GetMsgVpnClient(testcontext.SEMP().MonitorCtx(), testcontext.Messaging().VPN, url.QueryEscape(messagingService.GetApplicationID()), nil)
	ExpectWithOffset(callDepth, err).To(BeNil(), "Got unexpected error checking SEMP for connected client")
	ExpectWithOffset(callDepth, clientResponse.Data).To(Not(BeNil()), "Expected SEMP response data to not be nil when checking for connected client")
	return clientResponse.Data
}
