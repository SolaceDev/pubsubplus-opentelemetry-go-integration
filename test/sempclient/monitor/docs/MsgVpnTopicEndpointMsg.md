# MsgVpnTopicEndpointMsg

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachmentSize** | **int64** | The size of the Message attachment, in bytes (B). | [optional] [default to null]
**ContentSize** | **int64** | The size of the Message content, in bytes (B). | [optional] [default to null]
**DeliveryEligibleTime** | **int32** | The timestamp of when the Message is eligible for delivery. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). Available since 2.22. | [optional] [default to null]
**DmqEligible** | [****bool**](*bool.md) | Indicates whether the Message is eligible for the Dead Message Queue (DMQ). | [optional] [default to null]
**ExpiryTime** | **int32** | The timestamp of when the Message expires. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). | [optional] [default to null]
**MsgId** | **int64** | The identifier (ID) of the Message. | [optional] [default to null]
**MsgVpnName** | **string** | The name of the Message VPN. | [optional] [default to null]
**Priority** | **int32** | The priority level of the Message, from 9 (highest) to 0 (lowest). | [optional] [default to null]
**PublisherId** | **int64** | The identifier (ID) of the Message publisher. | [optional] [default to null]
**RedeliveryCount** | **int32** | The number of times the Message has been redelivered. | [optional] [default to null]
**ReplicatedMateMsgId** | **int64** | The Message identifier (ID) on the replication mate. Applicable only to replicated messages. | [optional] [default to null]
**ReplicationGroupMsgId** | **string** | An ID that uniquely identifies this Message within this replication group. Available since 2.21. | [optional] [default to null]
**ReplicationState** | **string** | The replication state of the Message. The allowed values and their meaning are:  &lt;pre&gt; \&quot;replicated\&quot; - The Message is replicated to the remote Message VPN. \&quot;not-replicated\&quot; - The Message is not being replicated to the remote Message VPN. \&quot;pending-replication\&quot; - The Message is queued for replication to the remote Message VPN. &lt;/pre&gt;  | [optional] [default to null]
**SpooledTime** | **int32** | The timestamp of when the Message was spooled in the Topic Endpoint. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). | [optional] [default to null]
**TopicEndpointName** | **string** | The name of the Topic Endpoint. | [optional] [default to null]
**Undelivered** | [****bool**](*bool.md) | Indicates whether delivery of the Message has never been attempted. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

