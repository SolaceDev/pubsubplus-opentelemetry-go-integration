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

package internal

import (
	"encoding/hex"
	"fmt"
	"strings"

	"solace.dev/go/messaging/trace/propagation/internal/impl/logging"
)

const Version string = "00"
const TraceParentDelimiter string = "-"

const BaggageListMembersSeparator string = ","
const BaggageListMemberKeyValueSplitter string = "="
const BaggageListMemberMetadataValueSplitter string = ";"

const TraceIDLength int = 32
const SpanIDLength int = 16

// GetTraceParentFromTraceContextProperties serializes a tracing context object into a string.
//
// @param traceID - the trace ID [16]byte property
// @param spanID - the span ID [8]byte property
// @param sampled - the sampled boolean property
// returns - the formated string representation of the context object
func GetTraceParentFromTraceContextProperties(traceID [16]byte, spanID [8]byte, sampled bool) string {
	// if the traceID and spanID is not available for this context
	if traceID == [16]byte{} || spanID == [8]byte{} {
		logging.Default.Info("[formatContextAsString]: Failed to retrieve trace parent as string from Solace message with tracing support")
		return ""
	}

	traceFlagsHex := "00"
	if sampled {
		traceFlagsHex = "01" // since the flag should be < 2 (binary: 10)
	}

	encodedTraceParent := fmt.Sprintf(
		"%s%s%x%s%x%s%s",
		Version,
		TraceParentDelimiter,
		traceID,
		TraceParentDelimiter,
		spanID,
		TraceParentDelimiter,
		traceFlagsHex,
	)

	return encodedTraceParent
}

// GetTraceContextPropertiesFromTraceParent deserializes a trace parent string into its tracing context object properties.
// @param traceParent the trace parent string to deserialize
// returns - the tracing context object properties
func GetTraceContextPropertiesFromTraceParent(traceParent string) (traceID *[16]byte, spanID *[8]byte, sampled bool, ok bool) {
	// valid value "00-75e792db89dec2cf3b3333a2f71869e4-982f925c36fb8a1c-01"
	traceParentParts := strings.Split(traceParent, TraceParentDelimiter)

	if len(traceParentParts) == 4 && traceParentParts[0] == Version &&
		len(traceParentParts[1]) == TraceIDLength && len(traceParentParts[2]) == SpanIDLength {
		// ignore the traceParentParts[0] which contains the version, we don't care about it
		var traceIDHex = traceParentParts[1]
		var spanIDHex = traceParentParts[2]
		var isSampledFlag = traceParentParts[3]

		sampled = (isSampledFlag != "00")

		traceIDBytes, errTraceID := hex.DecodeString(traceIDHex)
		var traceIDVal [16]byte
		copy(traceIDVal[:], traceIDBytes)

		spanIDBytes, errSpanID := hex.DecodeString(spanIDHex)
		var spanIDVal [8]byte
		copy(spanIDVal[:], spanIDBytes)

		if errTraceID == nil && errSpanID == nil {
			return &traceIDVal, &spanIDVal, sampled, true
		}
	}

	return nil, nil, false, false
}

// IsValidBaggageKey determines whether the given string is a valid entry key.
// @param key - the entry key to be validated.
// returns boolean of whether the key is valid.
func IsValidBaggageKey(key string) bool {
	return key != ""
}

// IsValidBaggageValue determines whether the given string is a valid
// baggage entry value.
// @param value - the entry value to be validated.
// returns boolean of whether the value is valid.
func IsValidBaggageValue(value string) bool {
	return value != ""
}
