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
	"strings"

	"solace.dev/go/messaging/pkg/solace/message"
	"solace.dev/go/messaging/trace/propagation/internal"
	"solace.dev/go/messaging/trace/propagation/internal/impl/logging"
)

// InboundMessageWithTracingSupport represents a message received by a consumer.
type InboundMessageWithTracingSupport interface {
	// Extend the InboundMessage interface.
	message.InboundMessage

	// GetCreationTraceContext will return the trace context metadata used for distributed message tracing message
	GetCreationTraceContext() (traceID [16]byte, spanID [8]byte, sampled bool, traceState string, ok bool)

	// SetTraceContext will set creation trace context metadata used for distributed message tracing.
	SetCreationTraceContext(traceID [16]byte, spanID [8]byte, sampled bool, traceState *string) (ok bool)

	// GetTransportTraceContext will return the trace context metadata used for distributed message tracing
	GetTransportTraceContext() (traceID [16]byte, spanID [8]byte, sampled bool, traceState string, ok bool)

	// SetTraceContext will set transport trace context metadata used for distributed message tracing.
	SetTransportTraceContext(traceID [16]byte, spanID [8]byte, sampled bool, traceState *string) (ok bool)

	// GetBaggage will return the baggage string associated with the message
	GetBaggage() (baggage string, ok bool)

	// SetBaggage will set the baggage string associated with the message
	SetBaggage(baggage string) error
}

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
	// TODO: ***Note use appropriate otel decoding functions for to string conversion
	// TODO: determine if transport or creation context to return, use GetCreationTraceContext() or GetTransportTraceContext() functions on the underlying message pointer
	// TODO: use an internal MessageImpl data storage to retrieve otel data to propagator:
	// 1. when key is 'traceparent' all parts of the traceparent from the underlying message pointer
	// 2. when key is 'tracestate' all parts of the traceparent from the underlying message pointer
	// 3. when key is 'baggage' GetBaggage() function on the underlying message pointer
	// when key is non of above (log as it is not supported format), try to get from message user property using given key as is, when found return value as is, otherwise return empty string.

	if key == "" {
		return "" // invalid key argument passed in
	}

	// cast the message to the extended interface that has message tracing support
	messageWithDT := carrier.messagePointer.(InboundMessageWithTracingSupport)

	// the tracing property names
	TracingPropertyName := internal.NewTracingPropertyNames()

	switch key {

	case TracingPropertyName.TraceParent:
		logging.Default.Info("Returning the TraceParent for this message")
		// return the trasport context as the Trace Parent
		if traceID, spanID, sampled, _, ok := messageWithDT.GetTransportTraceContext(); ok {
			// return the TraceParent from transport context since found
			logging.Default.Info("[GetTraceParentAsString]: Transport context was found, retreiving from Transport context")
			return internal.GetTraceParentFromTraceContextProperties(traceID, spanID, sampled)
		}
		// Else if not present, then return the creation context as the Trace Parent
		if traceID, spanID, sampled, _, ok := messageWithDT.GetCreationTraceContext(); ok {
			// return the TraceParent from transport context since found
			logging.Default.Info("[GetTraceParentAsString]: Creation context was found, retreiving from Creation context")
			return internal.GetTraceParentFromTraceContextProperties(traceID, spanID, sampled)
		}

		logging.Default.Debug("[GetTraceParentAsString]: No context information found on this Solace message")
		return "" // no value found

	case TracingPropertyName.TraceState:
		logging.Default.Info("Returning the TraceState for this message")
		// return the TraceState from the transport context
		_, _, _, transportTraceState, _ := messageWithDT.GetTransportTraceContext()
		_, _, _, creationTraceState, _ := messageWithDT.GetCreationTraceContext()

		if transportTraceState == "" && creationTraceState == "" {
			logging.Default.Debug("[getTraceStateAsString]: Transport and creation context not present")
			return ""
		}

		var traceState = "" // to hold the trace state from the message
		if transportTraceState != "" {
			logging.Default.Info("[getTraceStateAsString]: Found trace state for transport context")
			logging.Default.Debug("[getTraceStateAsString]: Transport context trace state: " + transportTraceState)
			traceState = transportTraceState
		}

		if creationTraceState != "" {
			logging.Default.Info("[getTraceStateAsString]: Found trace state for creation context")
			logging.Default.Debug("[getTraceStateAsString]: Creation context trace state: " + creationTraceState)

			if traceState != "" {
				traceState = traceState + "," + creationTraceState
			} else {
				traceState = creationTraceState
			}
		}
		return traceState // the TraceState value from the context(s)

	case TracingPropertyName.Baggage:
		if baggage, ok := messageWithDT.GetBaggage(); ok {
			logging.Default.Debug("Baggage string found for Solace message: " + baggage)
			return baggage
		}
		logging.Default.Debug("[getTraceStateAsString]: Baggage not present on this Solace message")
		return "" // empty baggage
	}

	// try to get from message user property using given key as is, when found return value as is, otherwise return empty string.
	if retrievedProperty, found := carrier.messagePointer.GetProperty(key); found {
		logging.Default.Warning("key does not match supported OTEL tracing property name, returning from message's User Property Map")
		return retrievedProperty.(string) // return the retrievd=ed property from the message
	}

	// return an empty string if we could not retrieve any of the tracing properties from the message
	logging.Default.Warning("key does not match any known OTEL tracing property name, returning an empty string")
	return ""
}

// Set stores the key-value pair.
func (carrier InboundMessageCarrier) Set(key, val string) {
	// TODO: ***Note use appropriate otel decoding functions for from string conversion to binary format
	// TODO:  use an internal MessageImpl data storage to insert otel data from propagator:
	// 1. when key is 'traceparent' decompose value and store all parts of the traceparent into the underlying message pointer, use method 'SetTraceContext'
	// 2. when key is 'tracestate' all parts of the traceparent from the underlying message pointer (determine if transport or creation context to return)
	// 3. when key is 'baggage' SetBaggage() function on the underlying message pointer
	// when key is non of above ignore (log as it is not supported format with warn level)

	if key != "" {
		if val != "" {
			// cast the message to the extended interface that has message tracing support
			message := carrier.messagePointer
			messageWithDT := message.(InboundMessageWithTracingSupport)

			// the tracing property names
			TracingPropertyName := internal.NewTracingPropertyNames()

			switch key {

			case TracingPropertyName.TraceParent:
				logging.Default.Info("Injecting trace parent into the Solace message")
				// return the trasport context as the Trace Parent
				if traceID, spanID, sampled, ok := internal.GetTraceContextPropertiesFromTraceParent(val); ok {
					// Set the Transport context if we find a creation context
					if creationTraceID, _, _, _, _ok := messageWithDT.GetCreationTraceContext(); _ok && creationTraceID != [16]byte{} {
						// else set them as the Transport context and clear the trace state of the transport context, if available
						emptyStr := ""
						messageWithDT.SetTransportTraceContext(*traceID, *spanID, sampled, &emptyStr)
						logging.Default.Info("Set[TraceParent]: Creation context was found, Injecting as transport context")
					} else {
						// else set them as the Creation context
						messageWithDT.SetCreationTraceContext(*traceID, *spanID, sampled, nil)
						logging.Default.Info("Set[TraceParent]: Injecting as creation context")
					}
				}

			case TracingPropertyName.TraceState:
				logging.Default.Info("Injecting trace state into the Solace message")
				newTraceState := string(val) // new string since we are passing a reference into function
				// inject traceState into creation context since it is not null
				if traceID, spanID, sampled, traceState, _ok := messageWithDT.GetCreationTraceContext(); !_ok || traceState == "" {
					// set it as the Creation context
					logging.Default.Info("Set[TraceState]: Either creation context was not found or was found without trace state, Injecting into creation context")
					messageWithDT.SetCreationTraceContext(traceID, spanID, sampled, &newTraceState)
				} else {
					// set it as the Transport context
					if traceID, spanID, sampled, _, ok := messageWithDT.GetTransportTraceContext(); ok {
						// use the tracing properties already in transport context as args
						logging.Default.Info("Set[TraceState]: Trace state found in creation context, Injecting into transport context")
						messageWithDT.SetTransportTraceContext(traceID, spanID, sampled, &newTraceState)
					}
				}

			case TracingPropertyName.Baggage:
				logging.Default.Info("'Injecting baggage information into the Solace message")
				// check that baggage is valid before injection
				if internal.IsValidBaggageValue(val) {
					logging.Default.Debug("Injected Baggage string: " + val)
					_error := messageWithDT.SetBaggage(val)
					if _error != nil {
						logging.Default.Error("Baggage injection failed: " + _error.Error())
					}
				} else {
					logging.Default.Warning("Baggage injection failed: Invalid Baggage value")
				}
			default:
				// when key is non of above ignore (log as it is not supported format with warn level)
				logging.Default.Warning("Ignoring any other OTEL third party tracing properties")
			}
		} else {
			// invalid val argument passed in
			logging.Default.Warning("val cannot be null or an empty string")
		}

	} else {
		// invalid key argument passed in
		logging.Default.Warning("key cannot be null or an empty string")
	}
}

// Keys lists the keys stored in this carrier.
func (carrier InboundMessageCarrier) Keys() []string {
	// it can return 'traceparent', 'tracestate', 'baggage' when they are supported by the underlying message pointer
	// the tracing property names
	TracingPropertyName := internal.NewTracingPropertyNames()
	propertyNames := TracingPropertyName.List()
	logging.Default.Debug("Retreived support tracing property names: [ " + strings.Join(propertyNames, ", ") + " ]")
	return propertyNames
}
