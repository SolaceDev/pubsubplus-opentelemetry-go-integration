# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/monitor*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMsgVpnReplayLog**](ReplayLogApi.md#GetMsgVpnReplayLog) | **Get** /msgVpns/{msgVpnName}/replayLogs/{replayLogName} | Get a Replay Log object.
[**GetMsgVpnReplayLogMsg**](ReplayLogApi.md#GetMsgVpnReplayLogMsg) | **Get** /msgVpns/{msgVpnName}/replayLogs/{replayLogName}/msgs/{msgId} | Get a Message object.
[**GetMsgVpnReplayLogMsgs**](ReplayLogApi.md#GetMsgVpnReplayLogMsgs) | **Get** /msgVpns/{msgVpnName}/replayLogs/{replayLogName}/msgs | Get a list of Message objects.
[**GetMsgVpnReplayLogTopicFilterSubscription**](ReplayLogApi.md#GetMsgVpnReplayLogTopicFilterSubscription) | **Get** /msgVpns/{msgVpnName}/replayLogs/{replayLogName}/topicFilterSubscriptions/{topicFilterSubscription} | Get a Topic Filter Subscription object.
[**GetMsgVpnReplayLogTopicFilterSubscriptions**](ReplayLogApi.md#GetMsgVpnReplayLogTopicFilterSubscriptions) | **Get** /msgVpns/{msgVpnName}/replayLogs/{replayLogName}/topicFilterSubscriptions | Get a list of Topic Filter Subscription objects.
[**GetMsgVpnReplayLogs**](ReplayLogApi.md#GetMsgVpnReplayLogs) | **Get** /msgVpns/{msgVpnName}/replayLogs | Get a list of Replay Log objects.

# **GetMsgVpnReplayLog**
> MsgVpnReplayLogResponse GetMsgVpnReplayLog(ctx, msgVpnName, replayLogName, optional)
Get a Replay Log object.

Get a Replay Log object.  When the Message Replay feature is enabled, message brokers store persistent messages in a Replay Log. These messages are kept until the log is full, after which the oldest messages are removed to free up space for new messages.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| replayLogName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

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


 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnReplayLogResponse**](MsgVpnReplayLogResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnReplayLogMsg**
> MsgVpnReplayLogMsgResponse GetMsgVpnReplayLogMsg(ctx, msgVpnName, replayLogName, msgId, optional)
Get a Message object.

Get a Message object.  A Message is a packet of information sent from producers to consumers. Messages are the central units of information that clients exchange using the message broker and which are cached in the Replay Log.   Attribute|Identifying|Deprecated :---|:---:|:---: msgId|x| msgVpnName|x| replayLogName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **replayLogName** | **string**| The name of the Replay Log. | 
  **msgId** | **string**| The identifier (ID) of the message. | 
 **optional** | ***ReplayLogApiGetMsgVpnReplayLogMsgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplayLogApiGetMsgVpnReplayLogMsgOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnReplayLogMsgResponse**](MsgVpnReplayLogMsgResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnReplayLogMsgs**
> MsgVpnReplayLogMsgsResponse GetMsgVpnReplayLogMsgs(ctx, msgVpnName, replayLogName, optional)
Get a list of Message objects.

Get a list of Message objects.  A Message is a packet of information sent from producers to consumers. Messages are the central units of information that clients exchange using the message broker and which are cached in the Replay Log.   Attribute|Identifying|Deprecated :---|:---:|:---: msgId|x| msgVpnName|x| replayLogName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **replayLogName** | **string**| The name of the Replay Log. | 
 **optional** | ***ReplayLogApiGetMsgVpnReplayLogMsgsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplayLogApiGetMsgVpnReplayLogMsgsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnReplayLogMsgsResponse**](MsgVpnReplayLogMsgsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnReplayLogTopicFilterSubscription**
> MsgVpnReplayLogTopicFilterSubscriptionResponse GetMsgVpnReplayLogTopicFilterSubscription(ctx, msgVpnName, replayLogName, topicFilterSubscription, optional)
Get a Topic Filter Subscription object.

Get a Topic Filter Subscription object.  One or more Subscriptions can be added to a replay-log so that only guaranteed messages published to matching topics are stored in the Replay Log.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| replayLogName|x| topicFilterSubscription|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.27.

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

Get a list of Topic Filter Subscription objects.  One or more Subscriptions can be added to a replay-log so that only guaranteed messages published to matching topics are stored in the Replay Log.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| replayLogName|x| topicFilterSubscription|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.27.

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

Get a list of Replay Log objects.  When the Message Replay feature is enabled, message brokers store persistent messages in a Replay Log. These messages are kept until the log is full, after which the oldest messages are removed to free up space for new messages.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| replayLogName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

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

