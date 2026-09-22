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

// Package sempclient contains generated SEMPv2 code
package sempclient

// The SEMPv2 clients are generated with swagger-codegen running in Docker.
// makeclean removes any previously generated code on the host, then
// generate-semp.sh builds the codegen image and produces the action, config
// and monitor clients. Files move in and out of the containers through a
// named volume and `docker cp`, so this works unchanged against a remote
// DOCKER_HOST as well as a local daemon.
//go:generate ./makeclean
//go:generate ./generate-semp.sh
