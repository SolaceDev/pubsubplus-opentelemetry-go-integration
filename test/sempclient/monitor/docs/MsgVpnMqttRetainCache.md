# MsgVpnMqttRetainCache

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupCacheInstance** | **string** | The name of the backup Cache Instance associated with this MQTT Retain Cache. | [optional] [default to null]
**BackupFailureReason** | **string** | The reason why the backup cache associated with this MQTT Retain Cache is operationally down, if any. | [optional] [default to null]
**BackupUp** | [****bool**](*bool.md) | Indicates whether the backup cache associated with this MQTT Retain Cache is operationally up. | [optional] [default to null]
**BackupUptime** | **int32** | The number of seconds that the backup cache associated with this MQTT Retain Cache has been operationally up. | [optional] [default to null]
**CacheCluster** | **string** | The name of the Cache Cluster associated with this MQTT Retain Cache. | [optional] [default to null]
**CacheName** | **string** | The name of the MQTT Retain Cache. | [optional] [default to null]
**DistributedCache** | **string** | The name of the Distributed Cache associated with this MQTT Retain Cache. | [optional] [default to null]
**Enabled** | [****bool**](*bool.md) | Indicates whether this MQTT Retain Cache is enabled. When the cache is disabled, neither retain messages nor retain requests will be delivered by the cache. However, live retain messages will continue to be delivered to currently connected MQTT clients. | [optional] [default to null]
**FailureReason** | **string** | The reason why this MQTT Retain Cache is operationally down, if any. | [optional] [default to null]
**MsgLifetime** | **int64** | The message lifetime, in seconds. If a message remains cached for the duration of its lifetime, the cache will remove the message. A lifetime of 0 results in the message being retained indefinitely, otherwise it must be 3 seconds or more. | [optional] [default to null]
**MsgVpnName** | **string** | The name of the Message VPN. | [optional] [default to null]
**PrimaryCacheInstance** | **string** | The name of the primary Cache Instance associated with this MQTT Retain Cache. | [optional] [default to null]
**PrimaryFailureReason** | **string** | The reason why the primary cache associated with this MQTT Retain Cache is operationally down, if any. | [optional] [default to null]
**PrimaryUp** | [****bool**](*bool.md) | Indicates whether the primary cache associated with this MQTT Retain Cache is operationally up. | [optional] [default to null]
**PrimaryUptime** | **int32** | The number of seconds that the primary cache associated with this MQTT Retain Cache has been operationally up. | [optional] [default to null]
**Up** | [****bool**](*bool.md) | Indicates whether this MQTT Retain Cache is operationally up. | [optional] [default to null]
**Uptime** | **int32** | The number of seconds that the MQTT Retain Cache has been operationally up. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

