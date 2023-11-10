# MsgVpnReplayLog

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EgressEnabled** | [****bool**](*bool.md) | Enable or disable the transmission of messages from the Replay Log. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is &#x60;false&#x60;. | [optional] [default to null]
**IngressEnabled** | [****bool**](*bool.md) | Enable or disable the reception of messages to the Replay Log. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is &#x60;false&#x60;. | [optional] [default to null]
**MaxSpoolUsage** | **int64** | The maximum spool usage allowed by the Replay Log, in megabytes (MB). If this limit is exceeded, old messages will be trimmed. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is &#x60;0&#x60;. | [optional] [default to null]
**MsgVpnName** | **string** | The name of the Message VPN. | [optional] [default to null]
**ReplayLogName** | **string** | The name of the Replay Log. | [optional] [default to null]
**TopicFilterEnabled** | [****bool**](*bool.md) | Enable or disable topic filtering for the Replay Log. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is &#x60;false&#x60;. Available since 2.27. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

