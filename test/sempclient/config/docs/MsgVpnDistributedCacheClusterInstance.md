# MsgVpnDistributedCacheClusterInstance

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoStartEnabled** | [****bool**](*bool.md) | Enable or disable auto-start for the Cache Instance. When enabled, the Cache Instance will automatically attempt to transition from the Stopped operational state to Up whenever it restarts or reconnects to the message broker. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is &#x60;false&#x60;. | [optional] [default to null]
**CacheName** | **string** | The name of the Distributed Cache. | [optional] [default to null]
**ClusterName** | **string** | The name of the Cache Cluster. | [optional] [default to null]
**Enabled** | [****bool**](*bool.md) | Enable or disable the Cache Instance. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is &#x60;false&#x60;. | [optional] [default to null]
**InstanceName** | **string** | The name of the Cache Instance. | [optional] [default to null]
**MsgVpnName** | **string** | The name of the Message VPN. | [optional] [default to null]
**StopOnLostMsgEnabled** | [****bool**](*bool.md) | Enable or disable stop-on-lost-message for the Cache Instance. When enabled, the Cache Instance will transition to the stopped operational state upon losing a message. When stopped, it cannot accept or respond to cache requests, but continues to cache messages. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is &#x60;true&#x60;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

