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
	"solace.dev/go/messaging/pkg/solace/message"
	"solace.dev/go/messaging/trace/propagation/internal/impl/logging"
)

// InboundMessageCarrier injects and extracts traces from a message.InboundMessage at 'solace.dev/go/messaging/pkg/solace/message'.
// it implements go.opentelemetry.io/otel/propagation.TextMapCarrier interface.
type InboundMessageCarrier struct {
	messagePointer message.InboundMessage
}

// NewInboundMessageCarrier creates a new InboundMessageCarrierImpl to be used with OpenTelemetry propagator API.
func NewInboundMessageCarrier(message message.InboundMessage) *InboundMessageCarrier {
	// perform null check on a pointer to the underlying message
	if message == nil {
		return nil
	}
	return &InboundMessageCarrier{messagePointer: message}
}

// GetInboundMessageCarrierPointer function
func GetInboundMessageCarrierPointer(message *InboundMessageCarrier) message.InboundMessage {
	if message == nil {
		return nil
	}
	return message.messagePointer
}

// Get returns the value associated with the passed key.
func (carrier InboundMessageCarrier) Get(key string) string {
	// 1. when key is 'traceparent', retrieve all parts of the traceparent from the underlying message pointer
	// 2. when key is 'tracestate', retrieve all parts of the tracestate from the underlying message pointer
	// 3. when key is 'baggage', get baggage from the underlying message pointer
	// when key is non of above (log as it is not supported format), try to get from message user property using given key as is, when found return value as is, otherwise return empty string.

	if key == "" {
		return "" // invalid key argument passed in
	}

	// try to get from message user property using given key as is, when found return value as is, otherwise return empty string.
	if retrievedProperty, found := carrier.messagePointer.GetProperty(key); !IsTracingKeySupported(key) && found {
		logging.Default.Warning("key does not match supported OTEL tracing property name, returning from message's User Property Map")
		return retrievedProperty.(string) // return the retrieved property from the message
	}

	return GetTracingField(carrier.messagePointer, key) // return the retrieved tracing property value
}

// Set stores the key-value pair.
func (carrier InboundMessageCarrier) Set(key, val string) {
	// Use an internal MessageImpl data storage to insert otel data from propagator:
	// 1. when key is 'traceparent' decompose value and store all parts of the traceparent into the underlying message pointer
	// 2. when key is 'tracestate' store all parts of the tracestate from the underlying message pointer (determine if transport or creation context to set)
	// 3. when key is 'baggage' set baggage on the underlying message pointer
	// when key is non of above ignore (log as it is not supported format with warn level)

	SetTracingField(carrier.messagePointer, key, val) // set the tracing field value
}

// Keys lists the keys stored in this carrier.
func (carrier InboundMessageCarrier) Keys() []string {
	// return these 'traceparent', 'tracestate', 'baggage'
	return GetSupportedTracingKeys()
}
