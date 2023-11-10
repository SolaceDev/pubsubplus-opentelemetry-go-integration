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

package test

import (
	// "solace.dev/go/trace/propagation"
	"solace.dev/go/messaging"
	"solace.dev/go/messaging/pkg/solace"

	"solace.dev/go/trace/test/helpers"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("API Version Info", func() {
	var builder solace.MessagingServiceBuilder
	BeforeEach(func() {
		builder = messaging.NewMessagingServiceBuilder().FromConfigurationProvider(helpers.DefaultConfiguration())
	})
	It("can retrieve the version from the API", func() {
		messagingService := helpers.BuildMessagingService(builder)
		client := helpers.ConnectMessagingService(messagingService)
		apiInfo := messagingService.Info()
		Expect(apiInfo).ToNot(BeNil())
		Expect(apiInfo.GetAPIVersion()).ToNot(BeEmpty())
		Expect(client.SoftwareVersion).To(Equal(apiInfo.GetAPIVersion()))
		helpers.DisconnectMessagingService(messagingService)
	})
})
