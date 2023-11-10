# MsgVpnReplayLog

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EgressEnabled** | [****bool**](*bool.md) | Indicates whether the transmission of messages from the Replay Log is enabled. | [optional] [default to null]
**IngressEnabled** | [****bool**](*bool.md) | Indicates whether the reception of messages to the Replay Log is enabled. | [optional] [default to null]
**MaxSpoolUsage** | **int64** | The maximum spool usage allowed by the Replay Log, in megabytes (MB). If this limit is exceeded, old messages will be trimmed. | [optional] [default to null]
**MsgSpoolUsage** | **int64** | The spool usage of the Replay Log, in bytes (B). | [optional] [default to null]
**MsgVpnName** | **string** | The name of the Message VPN. | [optional] [default to null]
**ReplayLogName** | **string** | The name of the Replay Log. | [optional] [default to null]
**TopicFilterEnabled** | [****bool**](*bool.md) | Enable or disable topic filtering for the Replay Log. Available since 2.27. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

