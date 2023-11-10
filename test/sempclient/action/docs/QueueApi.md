# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/action*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DoMsgVpnQueueCancelReplay**](QueueApi.md#DoMsgVpnQueueCancelReplay) | **Put** /msgVpns/{msgVpnName}/queues/{queueName}/cancelReplay | Cancel the replay of messages to the Queue.
[**DoMsgVpnQueueClearStats**](QueueApi.md#DoMsgVpnQueueClearStats) | **Put** /msgVpns/{msgVpnName}/queues/{queueName}/clearStats | Clear the statistics for the Queue.
[**DoMsgVpnQueueCopyMsgFromQueue**](QueueApi.md#DoMsgVpnQueueCopyMsgFromQueue) | **Put** /msgVpns/{msgVpnName}/queues/{queueName}/copyMsgFromQueue | Copy a message from another Queue to this Queue.
[**DoMsgVpnQueueCopyMsgFromReplayLog**](QueueApi.md#DoMsgVpnQueueCopyMsgFromReplayLog) | **Put** /msgVpns/{msgVpnName}/queues/{queueName}/copyMsgFromReplayLog | Copy a message from a Replay Log to this Queue.
[**DoMsgVpnQueueCopyMsgFromTopicEndpoint**](QueueApi.md#DoMsgVpnQueueCopyMsgFromTopicEndpoint) | **Put** /msgVpns/{msgVpnName}/queues/{queueName}/copyMsgFromTopicEndpoint | Copy a message from a Topic Endpoint to this Queue.
[**DoMsgVpnQueueDeleteMsgs**](QueueApi.md#DoMsgVpnQueueDeleteMsgs) | **Put** /msgVpns/{msgVpnName}/queues/{queueName}/deleteMsgs | Delete all spooled messages from the Queue.
[**DoMsgVpnQueueMsgDelete**](QueueApi.md#DoMsgVpnQueueMsgDelete) | **Put** /msgVpns/{msgVpnName}/queues/{queueName}/msgs/{msgId}/delete | Delete the Message from the Queue.
[**DoMsgVpnQueueStartReplay**](QueueApi.md#DoMsgVpnQueueStartReplay) | **Put** /msgVpns/{msgVpnName}/queues/{queueName}/startReplay | Start the replay of messages to the Queue.
[**GetMsgVpnQueue**](QueueApi.md#GetMsgVpnQueue) | **Get** /msgVpns/{msgVpnName}/queues/{queueName} | Get a Queue object.
[**GetMsgVpnQueueMsg**](QueueApi.md#GetMsgVpnQueueMsg) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/msgs/{msgId} | Get a Queue Message object.
[**GetMsgVpnQueueMsgs**](QueueApi.md#GetMsgVpnQueueMsgs) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/msgs | Get a list of Queue Message objects.
[**GetMsgVpnQueues**](QueueApi.md#GetMsgVpnQueues) | **Get** /msgVpns/{msgVpnName}/queues | Get a list of Queue objects.

# **DoMsgVpnQueueCancelReplay**
> SempMetaOnlyResponse DoMsgVpnQueueCancelReplay(ctx, body, msgVpnName, queueName)
Cancel the replay of messages to the Queue.

Cancel the replay of messages to the Queue.    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnQueueCancelReplay**](MsgVpnQueueCancelReplay.md)| The Cancel Replay action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The name of the Queue. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnQueueClearStats**
> SempMetaOnlyResponse DoMsgVpnQueueClearStats(ctx, body, msgVpnName, queueName)
Clear the statistics for the Queue.

Clear the statistics for the Queue.    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnQueueClearStats**](MsgVpnQueueClearStats.md)| The Clear Stats action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The name of the Queue. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnQueueCopyMsgFromQueue**
> SempMetaOnlyResponse DoMsgVpnQueueCopyMsgFromQueue(ctx, body, msgVpnName, queueName)
Copy a message from another Queue to this Queue.

Copy a message from another Queue to this Queue.   Attribute|Required|Deprecated :---|:---:|:---: replicationGroupMsgId|x| sourceQueueName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.29.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnQueueCopyMsgFromQueue**](MsgVpnQueueCopyMsgFromQueue.md)| The Copy Message From Queue action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The name of the Queue. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnQueueCopyMsgFromReplayLog**
> SempMetaOnlyResponse DoMsgVpnQueueCopyMsgFromReplayLog(ctx, body, msgVpnName, queueName)
Copy a message from a Replay Log to this Queue.

Copy a message from a Replay Log to this Queue.   Attribute|Required|Deprecated :---|:---:|:---: replicationGroupMsgId|x| sourceReplayLogName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.29.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnQueueCopyMsgFromReplayLog**](MsgVpnQueueCopyMsgFromReplayLog.md)| The Copy Message From Replay Log action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The name of the Queue. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnQueueCopyMsgFromTopicEndpoint**
> SempMetaOnlyResponse DoMsgVpnQueueCopyMsgFromTopicEndpoint(ctx, body, msgVpnName, queueName)
Copy a message from a Topic Endpoint to this Queue.

Copy a message from a Topic Endpoint to this Queue.   Attribute|Required|Deprecated :---|:---:|:---: replicationGroupMsgId|x| sourceTopicEndpointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.29.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnQueueCopyMsgFromTopicEndpoint**](MsgVpnQueueCopyMsgFromTopicEndpoint.md)| The Copy Message From Topic Endpoint action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The name of the Queue. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnQueueDeleteMsgs**
> SempMetaOnlyResponse DoMsgVpnQueueDeleteMsgs(ctx, body, msgVpnName, queueName)
Delete all spooled messages from the Queue.

Delete all spooled messages from the Queue.    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.28.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnQueueDeleteMsgs**](MsgVpnQueueDeleteMsgs.md)| The Delete All Messages action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The name of the Queue. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnQueueMsgDelete**
> SempMetaOnlyResponse DoMsgVpnQueueMsgDelete(ctx, body, msgVpnName, queueName, msgId)
Delete the Message from the Queue.

Delete the Message from the Queue.    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnQueueMsgDelete**](MsgVpnQueueMsgDelete.md)| The Delete action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The name of the Queue. | 
  **msgId** | **string**| The identifier (ID) of the Message. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnQueueStartReplay**
> SempMetaOnlyResponse DoMsgVpnQueueStartReplay(ctx, body, msgVpnName, queueName)
Start the replay of messages to the Queue.

Start the replay of messages to the Queue.    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnQueueStartReplay**](MsgVpnQueueStartReplay.md)| The Start Replay action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The name of the Queue. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnQueue**
> MsgVpnQueueResponse GetMsgVpnQueue(ctx, msgVpnName, queueName, optional)
Get a Queue object.

Get a Queue object.  A Queue acts as both a destination that clients can publish messages to, and as an endpoint that clients can bind consumers to and consume messages from.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| queueName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The name of the Queue. | 
 **optional** | ***QueueApiGetMsgVpnQueueOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QueueApiGetMsgVpnQueueOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueueResponse**](MsgVpnQueueResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnQueueMsg**
> MsgVpnQueueMsgResponse GetMsgVpnQueueMsg(ctx, msgVpnName, queueName, msgId, optional)
Get a Queue Message object.

Get a Queue Message object.  A Queue Message is a packet of information sent from producers to consumers using the Queue.   Attribute|Identifying|Deprecated :---|:---:|:---: msgId|x| msgVpnName|x| queueName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The name of the Queue. | 
  **msgId** | **string**| The identifier (ID) of the Message. | 
 **optional** | ***QueueApiGetMsgVpnQueueMsgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QueueApiGetMsgVpnQueueMsgOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueueMsgResponse**](MsgVpnQueueMsgResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnQueueMsgs**
> MsgVpnQueueMsgsResponse GetMsgVpnQueueMsgs(ctx, msgVpnName, queueName, optional)
Get a list of Queue Message objects.

Get a list of Queue Message objects.  A Queue Message is a packet of information sent from producers to consumers using the Queue.   Attribute|Identifying|Deprecated :---|:---:|:---: msgId|x| msgVpnName|x| queueName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The name of the Queue. | 
 **optional** | ***QueueApiGetMsgVpnQueueMsgsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QueueApiGetMsgVpnQueueMsgsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueueMsgsResponse**](MsgVpnQueueMsgsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnQueues**
> MsgVpnQueuesResponse GetMsgVpnQueues(ctx, msgVpnName, optional)
Get a list of Queue objects.

Get a list of Queue objects.  A Queue acts as both a destination that clients can publish messages to, and as an endpoint that clients can bind consumers to and consume messages from.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| queueName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***QueueApiGetMsgVpnQueuesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QueueApiGetMsgVpnQueuesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueuesResponse**](MsgVpnQueuesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

