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

package test

import (
	"solace.dev/go/messaging/trace/propagation"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("API Version Info", func() {
	It("can retrieve the version from the API", func() {
		apiVersionInfo := propagation.Version()
		Expect(apiVersionInfo).ToNot(BeNil())
		Expect(apiVersionInfo).ToNot(BeEmpty())
	})
})
