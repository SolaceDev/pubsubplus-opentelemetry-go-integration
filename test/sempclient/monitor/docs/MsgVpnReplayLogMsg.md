# MsgVpnReplayLogMsg

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachmentSize** | **int64** | The size of the message attachment, in bytes (B). | [optional] [default to null]
**ContentSize** | **int64** | The size of the message content, in bytes (B). | [optional] [default to null]
**DmqEligible** | [****bool**](*bool.md) | Indicates whether the message is eligible for the Dead Message Queue (DMQ). | [optional] [default to null]
**MsgId** | **int64** | The identifier (ID) of the message. | [optional] [default to null]
**MsgVpnName** | **string** | The name of the Message VPN. | [optional] [default to null]
**Priority** | **int32** | The priority level of the message. | [optional] [default to null]
**PublisherId** | **int64** | The identifier (ID) of the message publisher. | [optional] [default to null]
**ReplayLogName** | **string** | The name of the Replay Log. | [optional] [default to null]
**ReplicationGroupMsgId** | **string** | An ID that uniquely identifies this Message within this replication group. Available since 2.21. | [optional] [default to null]
**SequenceNumber** | **int64** | The sequence number assigned to the message. Applicable only to messages received on sequenced topics. | [optional] [default to null]
**SpooledTime** | **int32** | The timestamp of when the message was spooled in the Replay Log. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

