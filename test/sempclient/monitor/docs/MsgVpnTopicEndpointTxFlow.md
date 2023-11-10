# MsgVpnTopicEndpointTxFlow

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AckedMsgCount** | **int64** | The number of guaranteed messages delivered and acknowledged by the consumer. | [optional] [default to null]
**ActivityState** | **string** | The activity state of the Flow. The allowed values and their meaning are:  &lt;pre&gt; \&quot;active-browser\&quot; - The Flow is active as a browser. \&quot;active-consumer\&quot; - The Flow is active as a consumer. \&quot;inactive\&quot; - The Flow is inactive. &lt;/pre&gt;  | [optional] [default to null]
**BindTime** | **int32** | The timestamp of when the Flow bound to the Topic Endpoint. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). | [optional] [default to null]
**ClientName** | **string** | The name of the Client. | [optional] [default to null]
**ConsumerRedeliveryRequestAllowed** | [****bool**](*bool.md) | Indicates whether redelivery requests can be received as negative acknowledgements (NACKs) from the consumer. Applicable only to REST consumers. | [optional] [default to null]
**CutThroughAckedMsgCount** | **int64** | The number of guaranteed messages that used cut-through delivery and are acknowledged by the consumer. | [optional] [default to null]
**DeliveryState** | **string** | The delivery state of the Flow. The allowed values and their meaning are:  &lt;pre&gt; \&quot;closed\&quot; - The Flow is unbound. \&quot;opened\&quot; - The Flow is bound but inactive. \&quot;unbinding\&quot; - The Flow received an unbind request. \&quot;handshaking\&quot; - The Flow is handshaking to become active. \&quot;deliver-cut-through\&quot; - The Flow is streaming messages using direct+guaranteed delivery. \&quot;deliver-from-input-stream\&quot; - The Flow is streaming messages using guaranteed delivery. \&quot;deliver-from-memory\&quot; - The Flow throttled causing message delivery from memory (RAM). \&quot;deliver-from-spool\&quot; - The Flow stalled causing message delivery from spool (ADB or disk). &lt;/pre&gt;  | [optional] [default to null]
**FlowId** | **int64** | The identifier (ID) of the Flow. | [optional] [default to null]
**HighestAckPendingMsgId** | **int64** | The highest identifier (ID) of message transmitted and waiting for acknowledgement. | [optional] [default to null]
**LastAckedMsgId** | **int64** | The identifier (ID) of the last message transmitted and acknowledged by the consumer. | [optional] [default to null]
**LowestAckPendingMsgId** | **int64** | The lowest identifier (ID) of message transmitted and waiting for acknowledgement. | [optional] [default to null]
**MaxUnackedMsgsExceededMsgCount** | **int64** | The number of guaranteed messages that exceeded the maximum number of delivered unacknowledged messages. | [optional] [default to null]
**MsgVpnName** | **string** | The name of the Message VPN. | [optional] [default to null]
**NoLocalDelivery** | [****bool**](*bool.md) | Indicates whether not to deliver messages to a consumer that published them. | [optional] [default to null]
**RedeliveredMsgCount** | **int64** | The number of guaranteed messages that were redelivered. | [optional] [default to null]
**RedeliveryRequestCount** | **int64** | The number of consumer requests via negative acknowledgements (NACKs) to redeliver guaranteed messages. | [optional] [default to null]
**SessionName** | **string** | The name of the Transacted Session for the Flow. | [optional] [default to null]
**StoreAndForwardAckedMsgCount** | **int64** | The number of guaranteed messages that used store and forward delivery and are acknowledged by the consumer. | [optional] [default to null]
**TopicEndpointName** | **string** | The name of the Topic Endpoint. | [optional] [default to null]
**TransportRetransmitMsgCount** | **int64** | The number of guaranteed messages that were retransmitted at the transport layer as part of a single delivery attempt. Available since 2.18. | [optional] [default to null]
**UnackedMsgCount** | **int64** | The number of guaranteed messages delivered but not yet acknowledged by the consumer. | [optional] [default to null]
**UsedWindowSize** | **int64** | The number of guaranteed messages using the available window size. | [optional] [default to null]
**WindowClosedCount** | **int64** | The number of times the window for guaranteed messages was filled and closed before an acknowledgement was received. | [optional] [default to null]
**WindowSize** | **int64** | The number of outstanding guaranteed messages that can be transmitted over the Flow before an acknowledgement is received. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

