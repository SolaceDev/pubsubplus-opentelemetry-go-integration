# MsgVpnQueueStartReplay

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AfterMsg** | **string** | The Message after which to begin replay, identified by its Replication Group Message ID. Available since 2.21. | [optional] [default to null]
**FromTime** | **int32** | The time to begin replaying messages from. The value must be no less than the time that the replay log was created. To play back the whole log, this parameter must be omitted. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). | [optional] [default to null]
**ReplayLogName** | **string** | The name of the Replay Log. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

