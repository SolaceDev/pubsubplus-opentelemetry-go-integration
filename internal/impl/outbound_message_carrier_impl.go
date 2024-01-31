// pubsubplus-opentelemetry-go-integration
//
// Copyright 2024 Solace Corporation. All rights reserved.
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

// Package impl contains the main type definitions
// for the Getter and Setter for handling OTEL tracing.
package impl

import (
	"solace.dev/go/messaging-trace/opentelemetry/internal/impl/logging"
	"solace.dev/go/messaging/pkg/solace/message"
)

// OutboundMessageCarrier injects and extracts traces from a message.OutboundMessage at 'solace.dev/go/messaging/pkg/solace/message',
// it implements go.opentelemetry.io/otel/propagation.TextMapCarrier interface.
type OutboundMessageCarrier struct {
	messagePointer message.OutboundMessage
}

// NewOutboundMessageCarrier creates a new OutboundIMessageCarrier to be used with OpenTelemetry propagator API.
func NewOutboundMessageCarrier(message message.OutboundMessage) *OutboundMessageCarrier {
	// perform null check on a pointer to the underlying message
	if message == nil {
		return nil
	}
	return &OutboundMessageCarrier{messagePointer: message}
}

// GetOutboundMessageCarrierPointer function
func GetOutboundMessageCarrierPointer(message *OutboundMessageCarrier) message.OutboundMessage {
	if message == nil {
		return nil
	}
	return message.messagePointer
}

// Get returns the value associated with the passed key.
func (carrier OutboundMessageCarrier) Get(key string) string {
	// 1. when key is 'traceparent', get all parts of the traceparent from the underlying message pointer
	// 2. when key is 'tracestate', get all parts of the tracestate from the underlying message pointer
	// 3. when key is 'baggage', get baggage from the underlying message pointer
	// when key is non of above ignore (log as it is not supported format)

	if key == "" {
		return "" // invalid key argument passed in
	}

	return GetTracingField(carrier.messagePointer, key) // return the retrieved tracing property value
}

// Set stores the key-value pair.
func (carrier OutboundMessageCarrier) Set(key, val string) {
	// Use an internal MessageImpl data storage to insert otel data from propagator:
	// 1. when key is 'traceparent' decompose value and store all parts of the traceparent into the underlying message pointer
	// 2. when key is 'tracestate' all parts of the tracestate from the underlying message pointer (determine if transport or creation context to return)
	// 3. when key is 'baggage' set baggage on the underlying message pointer
	// when key is non of above (log as it is not supported format) insert into the message as user property, key as is & value string as is.

	// if key is not empty but not supported, set it in the User Property Map of the OutboundMessage
	if key != "" && !IsTracingKeySupported(key) {
		// TODO: Set into message as user properties in UserProperty Map
		logging.Default.Warning("Ignoring any other OTEL third party tracing properties")
		return // return from function
	}

	SetTracingField(carrier.messagePointer, key, val) // set the tracing field value
}

// Keys lists the keys stored in this carrier.
func (carrier OutboundMessageCarrier) Keys() []string {
	// return these 'traceparent', 'tracestate', 'baggage'
	return GetSupportedTracingKeys()
}
