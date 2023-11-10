# MsgVpnClientConnection

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientAddress** | **string** | The IP address and TCP port on the Client side of the Client Connection. | [optional] [default to null]
**ClientName** | **string** | The name of the Client. | [optional] [default to null]
**Compression** | [****bool**](*bool.md) | Indicates whether compression is enabled for the Client Connection. | [optional] [default to null]
**Encryption** | [****bool**](*bool.md) | Indicates whether encryption (TLS) is enabled for the Client Connection. | [optional] [default to null]
**FastRetransmitCount** | **int32** | The number of TCP fast retransmits due to duplicate acknowledgments (ACKs). See RFC 5681 for further details. | [optional] [default to null]
**MsgVpnName** | **string** | The name of the Message VPN. | [optional] [default to null]
**RxQueueByteCount** | **int32** | The number of bytes currently in the receive queue for the Client Connection. | [optional] [default to null]
**SegmentReceivedOutOfOrderCount** | **int32** | The number of TCP segments received from the Client Connection out of order. | [optional] [default to null]
**SmoothedRoundTripTime** | **int64** | The TCP smoothed round-trip time (SRTT) for the Client Connection, in nanoseconds. See RFC 2988 for further details. | [optional] [default to null]
**TcpState** | **string** | The TCP state of the Client Connection. When fully operational, should be: established. See RFC 793 for further details. The allowed values and their meaning are:  &lt;pre&gt; \&quot;closed\&quot; - No connection state at all. \&quot;listen\&quot; - Waiting for a connection request from any remote TCP and port. \&quot;syn-sent\&quot; - Waiting for a matching connection request after having sent a connection request. \&quot;syn-received\&quot; - Waiting for a confirming connection request acknowledgment after having both received and sent a connection request. \&quot;established\&quot; - An open connection, data received can be delivered to the user. \&quot;close-wait\&quot; - Waiting for a connection termination request from the local user. \&quot;fin-wait-1\&quot; - Waiting for a connection termination request from the remote TCP, or an acknowledgment of the connection termination request previously sent. \&quot;closing\&quot; - Waiting for a connection termination request acknowledgment from the remote TCP. \&quot;last-ack\&quot; - Waiting for an acknowledgment of the connection termination request previously sent to the remote TCP. \&quot;fin-wait-2\&quot; - Waiting for a connection termination request from the remote TCP. \&quot;time-wait\&quot; - Waiting for enough time to pass to be sure the remote TCP received the acknowledgment of its connection termination request. &lt;/pre&gt;  | [optional] [default to null]
**TimedRetransmitCount** | **int32** | The number of TCP segments retransmitted due to timeout awaiting an acknowledgement (ACK). See RFC 793 for further details. | [optional] [default to null]
**TxQueueByteCount** | **int32** | The number of bytes currently in the transmit queue for the Client Connection. | [optional] [default to null]
**Uptime** | **int64** | The amount of time in seconds since the Client Connection was established. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

