# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/config*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMsgVpnReplayLog**](ReplayLogApi.md#CreateMsgVpnReplayLog) | **Post** /msgVpns/{msgVpnName}/replayLogs | Create a Replay Log object.
[**CreateMsgVpnReplayLogTopicFilterSubscription**](ReplayLogApi.md#CreateMsgVpnReplayLogTopicFilterSubscription) | **Post** /msgVpns/{msgVpnName}/replayLogs/{replayLogName}/topicFilterSubscriptions | Create a Topic Filter Subscription object.
[**DeleteMsgVpnReplayLog**](ReplayLogApi.md#DeleteMsgVpnReplayLog) | **Delete** /msgVpns/{msgVpnName}/replayLogs/{replayLogName} | Delete a Replay Log object.
[**DeleteMsgVpnReplayLogTopicFilterSubscription**](ReplayLogApi.md#DeleteMsgVpnReplayLogTopicFilterSubscription) | **Delete** /msgVpns/{msgVpnName}/replayLogs/{replayLogName}/topicFilterSubscriptions/{topicFilterSubscription} | Delete a Topic Filter Subscription object.
[**GetMsgVpnReplayLog**](ReplayLogApi.md#GetMsgVpnReplayLog) | **Get** /msgVpns/{msgVpnName}/replayLogs/{replayLogName} | Get a Replay Log object.
[**GetMsgVpnReplayLogTopicFilterSubscription**](ReplayLogApi.md#GetMsgVpnReplayLogTopicFilterSubscription) | **Get** /msgVpns/{msgVpnName}/replayLogs/{replayLogName}/topicFilterSubscriptions/{topicFilterSubscription} | Get a Topic Filter Subscription object.
[**GetMsgVpnReplayLogTopicFilterSubscriptions**](ReplayLogApi.md#GetMsgVpnReplayLogTopicFilterSubscriptions) | **Get** /msgVpns/{msgVpnName}/replayLogs/{replayLogName}/topicFilterSubscriptions | Get a list of Topic Filter Subscription objects.
[**GetMsgVpnReplayLogs**](ReplayLogApi.md#GetMsgVpnReplayLogs) | **Get** /msgVpns/{msgVpnName}/replayLogs | Get a list of Replay Log objects.
[**ReplaceMsgVpnReplayLog**](ReplayLogApi.md#ReplaceMsgVpnReplayLog) | **Put** /msgVpns/{msgVpnName}/replayLogs/{replayLogName} | Replace a Replay Log object.
[**UpdateMsgVpnReplayLog**](ReplayLogApi.md#UpdateMsgVpnReplayLog) | **Patch** /msgVpns/{msgVpnName}/replayLogs/{replayLogName} | Update a Replay Log object.

# **CreateMsgVpnReplayLog**
> MsgVpnReplayLogResponse CreateMsgVpnReplayLog(ctx, body, msgVpnName, optional)
Create a Replay Log object.

Create a Replay Log object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  When the Message Replay feature is enabled, message brokers store persistent messages in a Replay Log. These messages are kept until the log is full, after which the oldest messages are removed to free up space for new messages.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x||| replayLogName|x|x||||    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.10.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnReplayLog**](MsgVpnReplayLog.md)| The Replay Log object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***ReplayLogApiCreateMsgVpnReplayLogOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplayLogApiCreateMsgVpnReplayLogOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnReplayLogResponse**](MsgVpnReplayLogResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnReplayLogTopicFilterSubscription**
> MsgVpnReplayLogTopicFilterSubscriptionResponse CreateMsgVpnReplayLogTopicFilterSubscription(ctx, body, msgVpnName, replayLogName, optional)
Create a Topic Filter Subscription object.

Create a Topic Filter Subscription object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  One or more Subscriptions can be added to a replay-log so that only guaranteed messages published to matching topics are stored in the Replay Log.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x||| replayLogName|x||x||| topicFilterSubscription|x|x||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.27.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnReplayLogTopicFilterSubscription**](MsgVpnReplayLogTopicFilterSubscription.md)| The Topic Filter Subscription object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **replayLogName** | **string**| The name of the Replay Log. | 
 **optional** | ***ReplayLogApiCreateMsgVpnReplayLogTopicFilterSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplayLogApiCreateMsgVpnReplayLogTopicFilterSubscriptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnReplayLogTopicFilterSubscriptionResponse**](MsgVpnReplayLogTopicFilterSubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnReplayLog**
> SempMetaOnlyResponse DeleteMsgVpnReplayLog(ctx, msgVpnName, replayLogName)
Delete a Replay Log object.

Delete a Replay Log object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  When the Message Replay feature is enabled, message brokers store persistent messages in a Replay Log. These messages are kept until the log is full, after which the oldest messages are removed to free up space for new messages.  A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.10.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **replayLogName** | **string**| The name of the Replay Log. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnReplayLogTopicFilterSubscription**
> SempMetaOnlyResponse DeleteMsgVpnReplayLogTopicFilterSubscription(ctx, msgVpnName, replayLogName, topicFilterSubscription)
Delete a Topic Filter Subscription object.

Delete a Topic Filter Subscription object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  One or more Subscriptions can be added to a replay-log so that only guaranteed messages published to matching topics are stored in the Replay Log.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.27.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **replayLogName** | **string**| The name of the Replay Log. | 
  **topicFilterSubscription** | **string**| The topic of the Subscription. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnReplayLog**
> MsgVpnReplayLogResponse GetMsgVpnReplayLog(ctx, msgVpnName, replayLogName, optional)
Get a Replay Log object.

Get a Replay Log object.  When the Message Replay feature is enabled, message brokers store persistent messages in a Replay Log. These messages are kept until the log is full, after which the oldest messages are removed to free up space for new messages.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| replayLogName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.10.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **replayLogName** | **string**| The name of the Replay Log. | 
 **optional** | ***ReplayLogApiGetMsgVpnReplayLogOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplayLogApiGetMsgVpnReplayLogOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnReplayLogResponse**](MsgVpnReplayLogResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnReplayLogTopicFilterSubscription**
> MsgVpnReplayLogTopicFilterSubscriptionResponse GetMsgVpnReplayLogTopicFilterSubscription(ctx, msgVpnName, replayLogName, topicFilterSubscription, optional)
Get a Topic Filter Subscription object.

Get a Topic Filter Subscription object.  One or more Subscriptions can be added to a replay-log so that only guaranteed messages published to matching topics are stored in the Replay Log.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| replayLogName|x||| topicFilterSubscription|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.27.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **replayLogName** | **string**| The name of the Replay Log. | 
  **topicFilterSubscription** | **string**| The topic of the Subscription. | 
 **optional** | ***ReplayLogApiGetMsgVpnReplayLogTopicFilterSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplayLogApiGetMsgVpnReplayLogTopicFilterSubscriptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnReplayLogTopicFilterSubscriptionResponse**](MsgVpnReplayLogTopicFilterSubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnReplayLogTopicFilterSubscriptions**
> MsgVpnReplayLogTopicFilterSubscriptionsResponse GetMsgVpnReplayLogTopicFilterSubscriptions(ctx, msgVpnName, replayLogName, optional)
Get a list of Topic Filter Subscription objects.

Get a list of Topic Filter Subscription objects.  One or more Subscriptions can be added to a replay-log so that only guaranteed messages published to matching topics are stored in the Replay Log.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| replayLogName|x||| topicFilterSubscription|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.27.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **replayLogName** | **string**| The name of the Replay Log. | 
 **optional** | ***ReplayLogApiGetMsgVpnReplayLogTopicFilterSubscriptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplayLogApiGetMsgVpnReplayLogTopicFilterSubscriptionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnReplayLogTopicFilterSubscriptionsResponse**](MsgVpnReplayLogTopicFilterSubscriptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnReplayLogs**
> MsgVpnReplayLogsResponse GetMsgVpnReplayLogs(ctx, msgVpnName, optional)
Get a list of Replay Log objects.

Get a list of Replay Log objects.  When the Message Replay feature is enabled, message brokers store persistent messages in a Replay Log. These messages are kept until the log is full, after which the oldest messages are removed to free up space for new messages.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| replayLogName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.10.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***ReplayLogApiGetMsgVpnReplayLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplayLogApiGetMsgVpnReplayLogsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnReplayLogsResponse**](MsgVpnReplayLogsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnReplayLog**
> MsgVpnReplayLogResponse ReplaceMsgVpnReplayLog(ctx, body, msgVpnName, replayLogName, optional)
Replace a Replay Log object.

Replace a Replay Log object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  When the Message Replay feature is enabled, message brokers store persistent messages in a Replay Log. These messages are kept until the log is full, after which the oldest messages are removed to free up space for new messages.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- msgVpnName|x||x||||| replayLogName|x||x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.10.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnReplayLog**](MsgVpnReplayLog.md)| The Replay Log object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **replayLogName** | **string**| The name of the Replay Log. | 
 **optional** | ***ReplayLogApiReplaceMsgVpnReplayLogOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplayLogApiReplaceMsgVpnReplayLogOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnReplayLogResponse**](MsgVpnReplayLogResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnReplayLog**
> MsgVpnReplayLogResponse UpdateMsgVpnReplayLog(ctx, body, msgVpnName, replayLogName, optional)
Update a Replay Log object.

Update a Replay Log object. Any attribute missing from the request will be left unchanged.  When the Message Replay feature is enabled, message brokers store persistent messages in a Replay Log. These messages are kept until the log is full, after which the oldest messages are removed to free up space for new messages.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- msgVpnName|x|x||||| replayLogName|x|x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.10.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnReplayLog**](MsgVpnReplayLog.md)| The Replay Log object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **replayLogName** | **string**| The name of the Replay Log. | 
 **optional** | ***ReplayLogApiUpdateMsgVpnReplayLogOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplayLogApiUpdateMsgVpnReplayLogOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnReplayLogResponse**](MsgVpnReplayLogResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

