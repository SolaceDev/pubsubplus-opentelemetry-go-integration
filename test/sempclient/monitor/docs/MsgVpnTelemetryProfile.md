# MsgVpnTelemetryProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AclProfileName** | **string** | The name of the Telemetry Profile ACL Profile. | [optional] [default to null]
**ClientProfileName** | **string** | The name of the Telemetry Profile Client Profile. | [optional] [default to null]
**MsgVpnName** | **string** | The name of the Message VPN. | [optional] [default to null]
**QueueEventBindCountThreshold** | [***EventThreshold**](EventThreshold.md) |  | [optional] [default to null]
**QueueEventMsgSpoolUsageThreshold** | [***EventThreshold**](EventThreshold.md) |  | [optional] [default to null]
**QueueMaxBindCount** | **int64** | The maximum number of consumer flows that can bind to the Queue. | [optional] [default to null]
**QueueMaxMsgSpoolUsage** | **int64** | The maximum message spool usage allowed by the Queue, in megabytes (MB). | [optional] [default to null]
**QueueName** | **string** | The name of the Telemetry Profile Queue. | [optional] [default to null]
**ReceiverAclConnectDefaultAction** | **string** | The default action to take when a receiver client connects to the broker. The allowed values and their meaning are:  &lt;pre&gt; \&quot;allow\&quot; - Allow client connection unless an exception is found for it. \&quot;disallow\&quot; - Disallow client connection unless an exception is found for it. &lt;/pre&gt;  | [optional] [default to null]
**ReceiverEnabled** | [****bool**](*bool.md) | Enable or disable the ability for receiver clients to consume from the #telemetry queue. | [optional] [default to null]
**ReceiverEventConnectionCountPerClientUsernameThreshold** | [***EventThreshold**](EventThreshold.md) |  | [optional] [default to null]
**ReceiverMaxConnectionCountPerClientUsername** | **int64** | The maximum number of receiver connections per Client Username. | [optional] [default to null]
**ReceiverTcpCongestionWindowSize** | **int64** | The TCP initial congestion window size for clients using the Client Profile, in multiples of the TCP Maximum Segment Size (MSS). Changing the value from its default of 2 results in non-compliance with RFC 2581. Contact support before changing this value. | [optional] [default to null]
**ReceiverTcpKeepaliveCount** | **int64** | The number of TCP keepalive retransmissions to a client using the Client Profile before declaring that it is not available. | [optional] [default to null]
**ReceiverTcpKeepaliveIdleTime** | **int64** | The amount of time a client connection using the Client Profile must remain idle before TCP begins sending keepalive probes, in seconds. | [optional] [default to null]
**ReceiverTcpKeepaliveInterval** | **int64** | The amount of time between TCP keepalive retransmissions to a client using the Client Profile when no acknowledgement is received, in seconds. | [optional] [default to null]
**ReceiverTcpMaxSegmentSize** | **int64** | The TCP maximum segment size for clients using the Client Profile, in bytes. Changes are applied to all existing connections. | [optional] [default to null]
**ReceiverTcpMaxWindowSize** | **int64** | The TCP maximum window size for clients using the Client Profile, in kilobytes. Changes are applied to all existing connections. | [optional] [default to null]
**TelemetryProfileName** | **string** | The name of the Telemetry Profile. | [optional] [default to null]
**TraceEnabled** | [****bool**](*bool.md) | Enable or disable generation of all trace span data messages. When enabled, the state of configured trace filters control which messages get traced. When disabled, trace span data messages are never generated, regardless of the state of trace filters. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

