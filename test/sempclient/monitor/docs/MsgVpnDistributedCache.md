# MsgVpnDistributedCache

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheManagementUp** | [****bool**](*bool.md) | Indicates whether managing of the distributed cache over the  message bus is operationally up in the Message VPN. Available since 2.28. | [optional] [default to null]
**CacheName** | **string** | The name of the Distributed Cache. | [optional] [default to null]
**CacheVirtualRouter** | **string** | The virtual router of the Distributed Cache. The allowed values and their meaning are:  &lt;pre&gt; \&quot;auto\&quot; - The Distributed Cache is automatically assigned a virtual router at creation, depending on the broker&#x27;s active-standby role. &lt;/pre&gt;  Available since 2.28. | [optional] [default to null]
**Enabled** | [****bool**](*bool.md) | Indicates whether the Distributed Cache is enabled. | [optional] [default to null]
**Heartbeat** | **int64** | The heartbeat interval, in seconds, used by the Cache Instances to monitor connectivity with the message broker. | [optional] [default to null]
**LastFailureReason** | **string** | The reason for the last distributed cache management failure. Available since 2.28. | [optional] [default to null]
**LastFailureTime** | **int32** | The timestamp of the last distributed cache management failure. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). Available since 2.28. | [optional] [default to null]
**MsgVpnName** | **string** | The name of the Message VPN. | [optional] [default to null]
**MsgsLost** | [****bool**](*bool.md) | Indicates whether one or more messages were lost by any Cache Instance in the Distributed Cache. | [optional] [default to null]
**ScheduledDeleteMsgDayList** | **string** | The scheduled delete message day(s), specified as \&quot;daily\&quot; or a comma-separated list of days. Days must be specified as \&quot;Sun\&quot;, \&quot;Mon\&quot;, \&quot;Tue\&quot;, \&quot;Wed\&quot;, \&quot;Thu\&quot;, \&quot;Fri\&quot;, or \&quot;Sat\&quot;, with no spaces, and in sorted order from Sunday to Saturday. The empty-string (\&quot;\&quot;) can also be specified, indicating no schedule is configured (\&quot;scheduledDeleteMsgTimeList\&quot; must also be configured to the empty-string). | [optional] [default to null]
**ScheduledDeleteMsgTimeList** | **string** | The scheduled delete message time(s), specified as \&quot;hourly\&quot; or a comma-separated list of 24-hour times in the form hh:mm, or h:mm. There must be no spaces, and times (up to 4) must be in sorted order from 0:00 to 23:59. The empty-string (\&quot;\&quot;) can also be specified, indicating no schedule is configured (\&quot;scheduledDeleteMsgDayList\&quot; must also be configured to the empty-string). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

