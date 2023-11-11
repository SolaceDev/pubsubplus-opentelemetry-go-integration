// pubsubplus-opentelemetry-go-integration
//
// Copyright 2023-2023 Solace Corporation. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package propagation

import (
	"solace.dev/go/messaging/pkg/solace/message"
	impl "solace.dev/go/trace/propagation/internal/impl"
	"solace.dev/go/trace/propagation/propagation"
)

// NewInboundMessageCarrier returns an instance of propagation.InboundMessageCarrier that can be used
// with otel integration for extracting tracing information from a received Inbound message.
func NewInboundMessageCarrier(carrier message.InboundMessage) propagation.InboundMessageCarrier {
	return impl.NewInboundMessageCarrier(carrier)
}

// NewOutboundMessageCarrier returns an instance of propagation.OutboundMessageCarrier that can be used
// with otel integration for injecting tracing information before publishing an Outbound message.
func NewOutboundMessageCarrier(carrier message.OutboundMessage) propagation.OutboundMessageCarrier {
	return impl.NewOutboundMessageCarrier(carrier)
}
