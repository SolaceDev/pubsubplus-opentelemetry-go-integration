# MsgVpnDistributedCacheCluster

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheName** | **string** | The name of the Distributed Cache. | [optional] [default to null]
**ClusterName** | **string** | The name of the Cache Cluster. | [optional] [default to null]
**DeliverToOneOverrideEnabled** | [****bool**](*bool.md) | Enable or disable deliver-to-one override for the Cache Cluster. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is &#x60;true&#x60;. | [optional] [default to null]
**Enabled** | [****bool**](*bool.md) | Enable or disable the Cache Cluster. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is &#x60;false&#x60;. | [optional] [default to null]
**EventDataByteRateThreshold** | [***EventThresholdByValue**](EventThresholdByValue.md) |  | [optional] [default to null]
**EventDataMsgRateThreshold** | [***EventThresholdByValue**](EventThresholdByValue.md) |  | [optional] [default to null]
**EventMaxMemoryThreshold** | [***EventThresholdByPercent**](EventThresholdByPercent.md) |  | [optional] [default to null]
**EventMaxTopicsThreshold** | [***EventThresholdByPercent**](EventThresholdByPercent.md) |  | [optional] [default to null]
**EventRequestQueueDepthThreshold** | [***EventThresholdByPercent**](EventThresholdByPercent.md) |  | [optional] [default to null]
**EventRequestRateThreshold** | [***EventThresholdByValue**](EventThresholdByValue.md) |  | [optional] [default to null]
**EventResponseRateThreshold** | [***EventThresholdByValue**](EventThresholdByValue.md) |  | [optional] [default to null]
**GlobalCachingEnabled** | [****bool**](*bool.md) | Enable or disable global caching for the Cache Cluster. When enabled, the Cache Instances will fetch topics from remote Home Cache Clusters when requested, and subscribe to those topics to cache them locally. When disabled, the Cache Instances will remove all subscriptions and cached messages for topics from remote Home Cache Clusters. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is &#x60;false&#x60;. | [optional] [default to null]
**GlobalCachingHeartbeat** | **int64** | The heartbeat interval, in seconds, used by the Cache Instances to monitor connectivity with the remote Home Cache Clusters. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is &#x60;3&#x60;. | [optional] [default to null]
**GlobalCachingTopicLifetime** | **int64** | The topic lifetime, in seconds. If no client requests are received for a given global topic over the duration of the topic lifetime, then the Cache Instance will remove the subscription and cached messages for that topic. A value of 0 disables aging. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is &#x60;3600&#x60;. | [optional] [default to null]
**MaxMemory** | **int64** | The maximum memory usage, in megabytes (MB), for each Cache Instance in the Cache Cluster. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is &#x60;2048&#x60;. | [optional] [default to null]
**MaxMsgsPerTopic** | **int64** | The maximum number of messages per topic for each Cache Instance in the Cache Cluster. When at the maximum, old messages are removed as new messages arrive. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is &#x60;1&#x60;. | [optional] [default to null]
**MaxRequestQueueDepth** | **int64** | The maximum queue depth for cache requests received by the Cache Cluster. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is &#x60;100000&#x60;. | [optional] [default to null]
**MaxTopicCount** | **int64** | The maximum number of topics for each Cache Instance in the Cache Cluster. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is &#x60;2000000&#x60;. | [optional] [default to null]
**MsgLifetime** | **int64** | The message lifetime, in seconds. If a message remains cached for the duration of its lifetime, the Cache Instance will remove the message. A lifetime of 0 results in the message being retained indefinitely. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is &#x60;0&#x60;. | [optional] [default to null]
**MsgVpnName** | **string** | The name of the Message VPN. | [optional] [default to null]
**NewTopicAdvertisementEnabled** | [****bool**](*bool.md) | Enable or disable the advertising, onto the message bus, of new topics learned by each Cache Instance in the Cache Cluster. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is &#x60;false&#x60;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

