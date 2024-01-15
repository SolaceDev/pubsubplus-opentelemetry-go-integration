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
	"encoding/hex"
	"time"

	"solace.dev/go/messaging"
	"solace.dev/go/messaging/pkg/solace"
	"solace.dev/go/messaging/pkg/solace/config"
	"solace.dev/go/messaging/pkg/solace/message"
	"solace.dev/go/messaging/pkg/solace/resource"
	"solace.dev/go/messaging/trace/propagation"
	"solace.dev/go/messaging/trace/propagation/logging"
	"solace.dev/go/messaging/trace/propagation/test/helpers"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
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

// OutboundMessageWithTracingSupport represents a message received by a consumer.
type OutboundMessageWithTracingSupport interface {
	// Extend the OutboundMessage interface.
	message.OutboundMessage

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

var _ = Describe("Local MessageBuilder Tests", func() {
	var messageBuilder solace.OutboundMessageBuilder

	// set the log level to Debug
	logging.SetLogLevel(logging.LogLevelDebug)

	// We are intentionally doing this outside of the BeforeEach block in order to not allocate hundreds of messaging services
	// as they are expensive to create and can cause issues with number of open files on some test systems
	messagingService, _ := messaging.NewMessagingServiceBuilder().FromConfigurationProvider(config.ServicePropertyMap{
		config.AuthenticationPropertySchemeBasicUserName: "default",
		config.AuthenticationPropertySchemeBasicPassword: "default",
		config.TransportLayerPropertyHost:                "localhost",
		config.ServicePropertyVPNName:                    "default",
	}).Build()

	BeforeEach(func() {
		messageBuilder = messagingService.MessageBuilder()
	})

	Describe("OutboundMessage Carrier setters and getters", func() {

		It("does not panic when calling Set(TraceParent) on carrier message", func() {
			msg, err := messageBuilder.Build()
			Expect(err).ToNot(HaveOccurred())

			var carrier = propagation.NewOutboundMessageCarrier(msg)
			traceParent1 := "00-79f90916c9a3dad1eb4b328e00469e45-3b364712c4e1f17f-00" // creation context
			carrier.Set(propagation.TracingPropertyName.TraceParent, traceParent1)
			Expect(carrier.Get(propagation.TracingPropertyName.TraceParent)).ToNot(BeEmpty())
		})

		It("does not panic when calling Set(TraceState) on carrier message", func() {
			msg, err := messageBuilder.Build()
			Expect(err).ToNot(HaveOccurred())

			var carrier = propagation.NewOutboundMessageCarrier(msg)
			traceState1 := "trace1=value1;trace2=value2;trace322=ewrHB554CGF" // creation trace state
			carrier.Set(propagation.TracingPropertyName.TraceState, traceState1)
			Expect(carrier.Get(propagation.TracingPropertyName.TraceState)).ToNot(BeEmpty())
		})

		It("does not panic when calling Set(Baggage) on carrier message", func() {
			msg, err := messageBuilder.Build()
			Expect(err).ToNot(HaveOccurred())

			var carrier = propagation.NewOutboundMessageCarrier(msg)
			baggageStr := "newbaggage=Oseme,example=yammer,foo=bar"
			carrier.Set(propagation.TracingPropertyName.Baggage, baggageStr)
			Expect(carrier.Get(propagation.TracingPropertyName.Baggage)).ToNot(BeEmpty())
		})

		It("does not panic when calling Get(TraceParent) on carrier message", func() {
			msg, err := messageBuilder.Build()
			Expect(err).ToNot(HaveOccurred())

			var carrier = propagation.NewOutboundMessageCarrier(msg)
			Expect(func() { carrier.Get(propagation.TracingPropertyName.TraceParent) }).ToNot(Panic())
		})

		It("does not panic when calling Get(TraceState) on carrier message", func() {
			msg, err := messageBuilder.Build()
			Expect(err).ToNot(HaveOccurred())

			var carrier = propagation.NewOutboundMessageCarrier(msg)
			Expect(func() { carrier.Get(propagation.TracingPropertyName.TraceState) }).ToNot(Panic())
		})

		It("does not panic when calling Get(Baggage) on carrier message", func() {
			msg, err := messageBuilder.Build()
			Expect(err).ToNot(HaveOccurred())

			var carrier = propagation.NewOutboundMessageCarrier(msg)
			Expect(func() { carrier.Get(propagation.TracingPropertyName.Baggage) }).ToNot(Panic())
		})

	})

})

var _ = Describe("Remote Message Tests", func() {

	const topic = "remote-message-tests"

	var messagingService solace.MessagingService
	var messageBuilder solace.OutboundMessageBuilder

	// set the log level to Debug
	logging.SetLogLevel(logging.LogLevelDebug)

	BeforeEach(func() {
		builder := messaging.NewMessagingServiceBuilder().
			FromConfigurationProvider(helpers.DefaultConfiguration()).
			FromConfigurationProvider(config.ServicePropertyMap{
				config.ServicePropertyGenerateSendTimestamps:    true,
				config.ServicePropertyGenerateReceiveTimestamps: true,
				config.ServicePropertyGenerateSenderID:          true,
			})

		var err error
		messagingService, err = builder.Build()
		Expect(err).ToNot(HaveOccurred())
		messageBuilder = messagingService.MessageBuilder()
	})

	Describe("Received InboundMessage Carrier setters and getters", func() {
		var publisher solace.DirectMessagePublisher
		var receiver solace.DirectMessageReceiver
		var inboundMessageChannel chan message.InboundMessage

		BeforeEach(func() {
			var err error
			err = messagingService.Connect()
			Expect(err).ToNot(HaveOccurred())

			publisher, err = messagingService.CreateDirectMessagePublisherBuilder().Build()
			Expect(err).ToNot(HaveOccurred())
			receiver, err = messagingService.CreateDirectMessageReceiverBuilder().WithSubscriptions(resource.TopicSubscriptionOf(topic)).Build()
			Expect(err).ToNot(HaveOccurred())

			err = publisher.Start()
			Expect(err).ToNot(HaveOccurred())

			inboundMessageChannel = make(chan message.InboundMessage)
			receiver.ReceiveAsync(func(inboundMessage message.InboundMessage) {
				inboundMessageChannel <- inboundMessage
			})

			err = receiver.Start()
			Expect(err).ToNot(HaveOccurred())
		})

		AfterEach(func() {
			var err error
			err = publisher.Terminate(10 * time.Second)
			Expect(err).ToNot(HaveOccurred())
			err = receiver.Terminate(10 * time.Second)
			Expect(err).ToNot(HaveOccurred())

			err = messagingService.Disconnect()
			Expect(err).ToNot(HaveOccurred())
		})

		It("does not panic when calling Set(TraceParent) on carrier inbound message", func() {
			message, err := messageBuilder.Build()
			Expect(err).ToNot(HaveOccurred())

			publisher.Publish(message, resource.TopicOf(topic))

			select {
			case msg := <-inboundMessageChannel:
				var carrier = propagation.NewInboundMessageCarrier(msg)
				traceParent1 := "00-79f90916c9a3dad1eb4b328e00469e45-3b364712c4e1f17f-01" // creation context
				carrier.Set(propagation.TracingPropertyName.TraceParent, traceParent1)
				Expect(carrier.Get(propagation.TracingPropertyName.TraceParent)).ToNot(BeEmpty())

				msgWithTracingSupport := msg.(InboundMessageWithTracingSupport)
				traceID, spanID, sampled, traceState, ok := msgWithTracingSupport.GetCreationTraceContext()
				Expect(ok).To(BeTrue())
				Expect(traceID).ToNot(BeEmpty()) // not be empty
				Expect(spanID).ToNot(BeEmpty())  // not be empty
				Expect(sampled).To(BeTrue())
				Expect(traceState).To(Equal(""))
			case <-time.After(1 * time.Second):
				Fail("timed out waiting for message to be delivered")
			}
		})

		It("does not panic when calling Set(TraceState) on carrier inbound message", func() {
			message, err := messageBuilder.Build()
			Expect(err).ToNot(HaveOccurred())

			publisher.Publish(message, resource.TopicOf(topic))

			select {
			case msg := <-inboundMessageChannel:
				var carrier = propagation.NewInboundMessageCarrier(msg)
				traceParent1 := "00-79f90916c9a3dad1eb4b328e00469e45-3b364712c4e1f17f-01" // creation context
				carrier.Set(propagation.TracingPropertyName.TraceParent, traceParent1)

				traceState1 := "trace1=value1;trace2=value2;trace322=ewrHB554CGF" // creation trace state
				carrier.Set(propagation.TracingPropertyName.TraceState, traceState1)
				Expect(carrier.Get(propagation.TracingPropertyName.TraceState)).ToNot(BeEmpty())

				msgWithTracingSupport := msg.(InboundMessageWithTracingSupport)
				traceID, spanID, sampled, _, ok := msgWithTracingSupport.GetCreationTraceContext()
				Expect(ok).To(BeTrue())
				Expect(traceID).ToNot(BeEmpty()) // not be empty
				Expect(spanID).ToNot(BeEmpty())  // not be empty
				Expect(sampled).To(BeTrue())
				// Expect(traceState).To(Equal(traceState1))
			case <-time.After(1 * time.Second):
				Fail("timed out waiting for message to be delivered")
			}
		})

		It("does not panic when calling Set(Baggage) on carrier inbound message", func() {
			message, err := messageBuilder.Build()
			Expect(err).ToNot(HaveOccurred())

			publisher.Publish(message, resource.TopicOf(topic))

			select {
			case msg := <-inboundMessageChannel:
				var carrier = propagation.NewInboundMessageCarrier(msg)
				baggageStr := "newbaggage=Oseme,example=yammer,foo=bar"
				carrier.Set(propagation.TracingPropertyName.Baggage, baggageStr)
				Expect(carrier.Get(propagation.TracingPropertyName.Baggage)).ToNot(BeEmpty())

				msgWithTracingSupport := msg.(InboundMessageWithTracingSupport)
				baggage, ok := msgWithTracingSupport.GetBaggage()
				Expect(ok).To(BeTrue())
				Expect(baggage).To(Equal(baggageStr))
			case <-time.After(1 * time.Second):
				Fail("timed out waiting for message to be delivered")
			}
		})

		It("does not panic when calling Get(TraceParent) on carrier inbound message", func() {
			message, err := messageBuilder.Build()
			Expect(err).ToNot(HaveOccurred())

			publisher.Publish(message, resource.TopicOf(topic))

			select {
			case msg := <-inboundMessageChannel:
				var carrier = propagation.NewInboundMessageCarrier(msg)
				Expect(func() { carrier.Get(propagation.TracingPropertyName.TraceParent) }).ToNot(Panic())
			case <-time.After(1 * time.Second):
				Fail("timed out waiting for message to be delivered")
			}
		})

		It("does not panic when calling Get(TraceState) on carrier inbound message", func() {
			message, err := messageBuilder.Build()
			Expect(err).ToNot(HaveOccurred())

			publisher.Publish(message, resource.TopicOf(topic))

			select {
			case msg := <-inboundMessageChannel:
				var carrier = propagation.NewInboundMessageCarrier(msg)
				Expect(func() { carrier.Get(propagation.TracingPropertyName.TraceState) }).ToNot(Panic())

				msgWithTracingSupport := msg.(InboundMessageWithTracingSupport)
				_, _, _, traceState, _ := msgWithTracingSupport.GetCreationTraceContext()
				Expect(traceState).To(BeEmpty())
			case <-time.After(1 * time.Second):
				Fail("timed out waiting for message to be delivered")
			}
		})

		It("does not panic when calling Get(Baggage) on carrier inbound message", func() {
			message, err := messageBuilder.Build()
			Expect(err).ToNot(HaveOccurred())

			publisher.Publish(message, resource.TopicOf(topic))

			select {
			case msg := <-inboundMessageChannel:
				var carrier = propagation.NewInboundMessageCarrier(msg)
				Expect(func() { carrier.Get(propagation.TracingPropertyName.Baggage) }).ToNot(Panic())

				msgWithTracingSupport := msg.(InboundMessageWithTracingSupport)
				baggage, _ := msgWithTracingSupport.GetBaggage()
				Expect(baggage).To(Equal(""))
			case <-time.After(1 * time.Second):
				Fail("timed out waiting for message to be delivered")
			}
		})

	})

	Describe("Published and received message with Distributed Tracing support", func() {
		var publisher solace.DirectMessagePublisher
		var receiver solace.DirectMessageReceiver
		var inboundMessageChannel chan InboundMessageWithTracingSupport

		BeforeEach(func() {
			var err error
			err = messagingService.Connect()
			Expect(err).ToNot(HaveOccurred())

			publisher, err = messagingService.CreateDirectMessagePublisherBuilder().Build()
			Expect(err).ToNot(HaveOccurred())
			receiver, err = messagingService.CreateDirectMessageReceiverBuilder().WithSubscriptions(resource.TopicSubscriptionOf(topic)).Build()
			Expect(err).ToNot(HaveOccurred())

			err = publisher.Start()
			Expect(err).ToNot(HaveOccurred())

			inboundMessageChannel = make(chan InboundMessageWithTracingSupport)
			receiver.ReceiveAsync(func(inboundMessage message.InboundMessage) {
				inboundMessageChannel <- inboundMessage.(InboundMessageWithTracingSupport)
			})

			err = receiver.Start()
			Expect(err).ToNot(HaveOccurred())
		})

		AfterEach(func() {
			var err error
			err = publisher.Terminate(10 * time.Second)
			Expect(err).ToNot(HaveOccurred())
			err = receiver.Terminate(10 * time.Second)
			Expect(err).ToNot(HaveOccurred())

			err = messagingService.Disconnect()
			Expect(err).ToNot(HaveOccurred())
		})

		It("should be able to publish/receive a message with no creation context", func() {
			message, err := messageBuilder.Build() // no creation context is set on message
			Expect(err).ToNot(HaveOccurred())

			publisher.Publish(message, resource.TopicOf(topic))

			select {
			case message := <-inboundMessageChannel:
				traceID, spanID, sampled, traceState, ok := message.GetCreationTraceContext()
				Expect(ok).To(BeFalse())
				Expect(traceID).To(Equal([16]byte{})) // empty
				Expect(spanID).To(Equal([8]byte{}))   // empty
				Expect(sampled).To(BeFalse())
				Expect(traceState).To(Equal(""))
			case <-time.After(1 * time.Second):
				Fail("timed out waiting for message to be delivered")
			}
		})

		It("should be able to publish/receive a message with a valid creation context", func() {
			message, err := messageBuilder.Build()
			Expect(err).ToNot(HaveOccurred())

			// cast the message to the extended interface that has message tracing support
			messageWithDT := message.(OutboundMessageWithTracingSupport)

			// set creation context on message
			creationCtxTraceID, _ := hex.DecodeString("79f90916c9a3dad1eb4b328e00469e45")
			creationCtxSpanID, _ := hex.DecodeString("3b364712c4e1f17f")
			sampledValue := true
			traceStateValue := "sometrace=Example"

			var creationCtxTraceID16 [16]byte
			var creationCtxSpanID8 [8]byte
			copy(creationCtxTraceID16[:], creationCtxTraceID)
			copy(creationCtxSpanID8[:], creationCtxSpanID)
			ok := messageWithDT.SetCreationTraceContext(creationCtxTraceID16, creationCtxSpanID8, sampledValue, &traceStateValue)
			Expect(ok).To(BeTrue())

			publisher.Publish(messageWithDT, resource.TopicOf(topic))

			select {
			case message := <-inboundMessageChannel:
				traceID, spanID, sampled, traceState, ok := message.GetCreationTraceContext()
				Expect(ok).To(BeTrue())
				Expect(traceID).To(Equal(creationCtxTraceID16)) // should be equal
				Expect(spanID).To(Equal(creationCtxSpanID8))    // should be equal
				Expect(sampled).To(Equal(sampledValue))
				Expect(traceState).To(Equal(traceStateValue))
			case <-time.After(1 * time.Second):
				Fail("timed out waiting for message to be delivered")
			}
		})

		It("should be able to publish/receive a message with a valid creation context without trace state", func() {
			message, err := messageBuilder.Build()
			Expect(err).ToNot(HaveOccurred())

			// cast the message to the extended interface that has message tracing support
			messageWithDT := message.(OutboundMessageWithTracingSupport)

			// set creation context on message
			creationCtxTraceID, _ := hex.DecodeString("79f90916c9a3dad1eb4b328e00469e45")
			creationCtxSpanID, _ := hex.DecodeString("3b364712c4e1f17f")
			sampledValue := true

			var creationCtxTraceID16 [16]byte
			var creationCtxSpanID8 [8]byte
			copy(creationCtxTraceID16[:], creationCtxTraceID)
			copy(creationCtxSpanID8[:], creationCtxSpanID)
			ok := messageWithDT.SetCreationTraceContext(creationCtxTraceID16, creationCtxSpanID8, sampledValue, nil) // no trace state
			Expect(ok).To(BeTrue())

			publisher.Publish(messageWithDT, resource.TopicOf(topic))

			select {
			case message := <-inboundMessageChannel:
				traceID, spanID, sampled, traceState, ok := message.GetCreationTraceContext()
				Expect(ok).To(BeTrue())
				Expect(traceID).To(Equal(creationCtxTraceID16)) // should be equal
				Expect(spanID).To(Equal(creationCtxSpanID8))    // should be equal
				Expect(sampled).To(Equal(sampledValue))
				Expect(traceState).To(Equal("")) // should be empty
			case <-time.After(1 * time.Second):
				Fail("timed out waiting for message to be delivered")
			}
		})

		It("should be able to publish/receive a message with no tranport context", func() {
			message, err := messageBuilder.Build() // no creation context is set on message
			Expect(err).ToNot(HaveOccurred())

			publisher.Publish(message, resource.TopicOf(topic))

			select {
			case message := <-inboundMessageChannel:
				traceID, spanID, sampled, traceState, ok := message.GetTransportTraceContext()
				Expect(ok).To(BeFalse())
				Expect(traceID).To(Equal([16]byte{})) // empty
				Expect(spanID).To(Equal([8]byte{}))   // empty
				Expect(sampled).To(BeFalse())
				Expect(traceState).To(Equal(""))
			case <-time.After(1 * time.Second):
				Fail("timed out waiting for message to be delivered")
			}
		})

		It("should be able to publish/receive a message with a valid transport context", func() {
			message, err := messageBuilder.Build()
			Expect(err).ToNot(HaveOccurred())

			// cast the message to the extended interface that has message tracing support
			messageWithDT := message.(OutboundMessageWithTracingSupport)

			// set transport context on message
			transportCtxTraceID, _ := hex.DecodeString("55d30916c9a3dad1eb4b328e00469e45")
			transportCtxSpanID, _ := hex.DecodeString("a7164712c4e1f17f")

			sampledValue := true
			traceStateValue := "trace=Sample"

			var transportCtxTraceID16 [16]byte
			var transportCtxSpanID8 [8]byte
			copy(transportCtxTraceID16[:], transportCtxTraceID)
			copy(transportCtxSpanID8[:], transportCtxSpanID)
			ok := messageWithDT.SetTransportTraceContext(transportCtxTraceID16, transportCtxSpanID8, sampledValue, &traceStateValue)
			Expect(ok).To(BeTrue())

			publisher.Publish(messageWithDT, resource.TopicOf(topic))

			select {
			case message := <-inboundMessageChannel:
				traceID, spanID, sampled, traceState, ok := message.GetTransportTraceContext()
				Expect(ok).To(BeTrue())
				Expect(traceID).To(Equal(transportCtxTraceID16)) // should be equal
				Expect(spanID).To(Equal(transportCtxSpanID8))    // should be equal
				Expect(sampled).To(BeTrue())
				Expect(traceState).To(Equal(traceStateValue))
			case <-time.After(1 * time.Second):
				Fail("timed out waiting for message to be delivered")
			}
		})

		It("should be able to publish/receive a message with a valid transport context without trace state", func() {
			message, err := messageBuilder.Build()
			Expect(err).ToNot(HaveOccurred())

			// cast the message to the extended interface that has message tracing support
			messageWithDT := message.(OutboundMessageWithTracingSupport)

			// set transport context on message
			transportCtxTraceID, _ := hex.DecodeString("55d30916c9a3dad1eb4b328e00469e45")
			transportCtxSpanID, _ := hex.DecodeString("a7164712c4e1f17f")
			sampledValue := true

			var transportCtxTraceID16 [16]byte
			var transportCtxSpanID8 [8]byte
			copy(transportCtxTraceID16[:], transportCtxTraceID)
			copy(transportCtxSpanID8[:], transportCtxSpanID)
			ok := messageWithDT.SetTransportTraceContext(transportCtxTraceID16, transportCtxSpanID8, sampledValue, nil)
			Expect(ok).To(BeTrue())

			publisher.Publish(messageWithDT, resource.TopicOf(topic))

			select {
			case message := <-inboundMessageChannel:
				traceID, spanID, sampled, traceState, ok := message.GetTransportTraceContext()
				Expect(ok).To(BeTrue())
				Expect(traceID).To(Equal(transportCtxTraceID16)) // should be equal
				Expect(spanID).To(Equal(transportCtxSpanID8))    // should be equal
				Expect(sampled).To(BeTrue())
				Expect(traceState).To(Equal("")) // should be empty
			case <-time.After(1 * time.Second):
				Fail("timed out waiting for message to be delivered")
			}
		})

		It("should be able to publish/receive a message with a valid creation context and no transport context", func() {
			message, err := messageBuilder.Build()
			Expect(err).ToNot(HaveOccurred())

			// cast the message to the extended interface that has message tracing support
			messageWithDT := message.(OutboundMessageWithTracingSupport)

			// set creation context on message
			creationCtxTraceID, _ := hex.DecodeString("79f90916c9a3dad1eb4b328e00469e45")
			creationCtxSpanID, _ := hex.DecodeString("3b364712c4e1f17f")
			sampledValue := true
			traceStateValue := "sometrace=Example"

			var creationCtxTraceID16 [16]byte
			var creationCtxSpanID8 [8]byte
			copy(creationCtxTraceID16[:], creationCtxTraceID)
			copy(creationCtxSpanID8[:], creationCtxSpanID)
			ok := messageWithDT.SetCreationTraceContext(creationCtxTraceID16, creationCtxSpanID8, sampledValue, &traceStateValue)
			Expect(ok).To(BeTrue())

			publisher.Publish(messageWithDT, resource.TopicOf(topic))

			select {
			case message := <-inboundMessageChannel:
				creationTraceID, creationSpanID, creationSampled, creationTraceState, creationOk := message.GetCreationTraceContext()
				transportTraceID, transportSpanID, transportSampled, transportTraceState, transportOk := message.GetTransportTraceContext()

				Expect(creationOk).To(BeTrue())
				Expect(creationTraceID).To(Equal(creationCtxTraceID16)) // should be equal
				Expect(creationSpanID).To(Equal(creationCtxSpanID8))    // should be equal
				Expect(creationSampled).To(BeTrue())
				Expect(creationTraceState).To(Equal(traceStateValue))

				Expect(transportOk).To(BeFalse())
				Expect(transportTraceID).To(Equal([16]byte{})) // empty
				Expect(transportSpanID).To(Equal([8]byte{}))   // empty
				Expect(transportSampled).To(BeFalse())
				Expect(transportTraceState).To(Equal(""))

			case <-time.After(1 * time.Second):
				Fail("timed out waiting for message to be delivered")
			}
		})

		It("should be able to publish/receive a message with different creation context and transport context", func() {
			message, err := messageBuilder.Build()
			Expect(err).ToNot(HaveOccurred())

			// cast the message to the extended interface that has message tracing support
			messageWithDT := message.(OutboundMessageWithTracingSupport)

			// set creation context on message
			creationCtxTraceID, _ := hex.DecodeString("79f90916c9a3dad1eb4b328e00469e45")
			creationCtxSpanID, _ := hex.DecodeString("3b364712c4e1f17f")
			creationCtxTraceState := "sometrace1=example1"

			// set transport context on message
			transportCtxTraceID, _ := hex.DecodeString("55d30916c9a3dad1eb4b328e00469e45")
			transportCtxSpanID, _ := hex.DecodeString("a7164712c4e1f17f")
			transportCtxTraceState := "sometrace2=example2"

			var creationCtxTraceID16, transportCtxTraceID16 [16]byte
			var creationCtxSpanID8, transportCtxSpanID8 [8]byte
			copy(creationCtxTraceID16[:], creationCtxTraceID)
			copy(creationCtxSpanID8[:], creationCtxSpanID)
			setCreationCtxOk := messageWithDT.SetCreationTraceContext(creationCtxTraceID16, creationCtxSpanID8, true, &creationCtxTraceState)
			Expect(setCreationCtxOk).To(BeTrue())

			copy(transportCtxTraceID16[:], transportCtxTraceID)
			copy(transportCtxSpanID8[:], transportCtxSpanID)
			setTransportCtxOk := messageWithDT.SetTransportTraceContext(transportCtxTraceID16, transportCtxSpanID8, true, &transportCtxTraceState)
			Expect(setTransportCtxOk).To(BeTrue())

			publisher.Publish(messageWithDT, resource.TopicOf(topic))

			select {
			case message := <-inboundMessageChannel:
				creationTraceID, creationSpanID, creationSampled, creationTraceState, creationOk := message.GetCreationTraceContext()
				transportTraceID, transportSpanID, transportSampled, transportTraceState, transportOk := message.GetTransportTraceContext()

				Expect(creationOk).To(BeTrue())
				Expect(creationTraceID).To(Equal(creationCtxTraceID16)) // should be equal
				Expect(creationSpanID).To(Equal(creationCtxSpanID8))    // should be equal
				Expect(creationSampled).To(BeTrue())
				Expect(creationTraceState).To(Equal(creationCtxTraceState))

				Expect(transportOk).To(BeTrue())
				Expect(transportTraceID).ToNot(Equal([16]byte{}))         // not empty
				Expect(transportTraceID).To(Equal(transportCtxTraceID16)) // should be equal

				Expect(transportSpanID).ToNot(Equal([8]byte{}))        // not empty
				Expect(transportSpanID).To(Equal(transportCtxSpanID8)) // should be equal

				Expect(transportSampled).To(BeTrue())
				Expect(transportTraceState).To(Equal(transportCtxTraceState))

				Expect(creationTraceID).ToNot(Equal(transportTraceID))       // should be not be equal
				Expect(creationSpanID).ToNot(Equal(transportSpanID))         // should be not be equal
				Expect(creationTraceState).ToNot(Equal(transportTraceState)) // should be not be equal

			case <-time.After(1 * time.Second):
				Fail("timed out waiting for message to be delivered")
			}
		})

		It("should be able to publish/receive a message with no baggage", func() {
			message, err := messageBuilder.Build()
			Expect(err).ToNot(HaveOccurred())

			// cast the message to the extended interface that has message tracing support
			messageWithDT := message.(OutboundMessageWithTracingSupport)

			baggageErr := messageWithDT.SetBaggage("") // set empty baggage
			Expect(baggageErr).To(BeNil())

			publisher.Publish(messageWithDT, resource.TopicOf(topic))

			select {
			case message := <-inboundMessageChannel:
				baggage, ok := message.GetBaggage()
				Expect(ok).To(BeTrue())
				Expect(baggage).To(Equal(""))
			case <-time.After(1 * time.Second):
				Fail("timed out waiting for message to be delivered")
			}
		})

		It("should be able to publish/receive a message with a valid baggage", func() {
			baggage := "baggage1=baggage1"
			message, err := messageBuilder.Build()
			Expect(err).ToNot(HaveOccurred())

			// cast the message to the extended interface that has message tracing support
			messageWithDT := message.(OutboundMessageWithTracingSupport)

			baggageErr := messageWithDT.SetBaggage(baggage) // set a valid baggage string
			Expect(baggageErr).To(BeNil())

			publisher.Publish(messageWithDT, resource.TopicOf(topic))

			select {
			case message := <-inboundMessageChannel:
				receivedBaggage, ok := message.GetBaggage()
				Expect(ok).To(BeTrue())
				Expect(receivedBaggage).To(Equal(baggage))
			case <-time.After(1 * time.Second):
				Fail("timed out waiting for message to be delivered")
			}
		})

	})

})
