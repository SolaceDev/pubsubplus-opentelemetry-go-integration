# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/action*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DoMsgVpnTopicEndpointCancelReplay**](TopicEndpointApi.md#DoMsgVpnTopicEndpointCancelReplay) | **Put** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/cancelReplay | Cancel the replay of messages to the Topic Endpoint.
[**DoMsgVpnTopicEndpointClearStats**](TopicEndpointApi.md#DoMsgVpnTopicEndpointClearStats) | **Put** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/clearStats | Clear the statistics for the Topic Endpoint.
[**DoMsgVpnTopicEndpointCopyMsgFromQueue**](TopicEndpointApi.md#DoMsgVpnTopicEndpointCopyMsgFromQueue) | **Put** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/copyMsgFromQueue | Copy a message from a Queue to this Topic Endpoint.
[**DoMsgVpnTopicEndpointCopyMsgFromReplayLog**](TopicEndpointApi.md#DoMsgVpnTopicEndpointCopyMsgFromReplayLog) | **Put** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/copyMsgFromReplayLog | Copy a message from a Replay Log to this Topic Endpoint.
[**DoMsgVpnTopicEndpointCopyMsgFromTopicEndpoint**](TopicEndpointApi.md#DoMsgVpnTopicEndpointCopyMsgFromTopicEndpoint) | **Put** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/copyMsgFromTopicEndpoint | Copy a message from another Topic Endpoint to this Topic Endpoint.
[**DoMsgVpnTopicEndpointDeleteMsgs**](TopicEndpointApi.md#DoMsgVpnTopicEndpointDeleteMsgs) | **Put** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/deleteMsgs | Delete all spooled messages from the Topic Endpoint.
[**DoMsgVpnTopicEndpointMsgDelete**](TopicEndpointApi.md#DoMsgVpnTopicEndpointMsgDelete) | **Put** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/msgs/{msgId}/delete | Delete the Message from the Topic Endpoint.
[**DoMsgVpnTopicEndpointStartReplay**](TopicEndpointApi.md#DoMsgVpnTopicEndpointStartReplay) | **Put** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/startReplay | Start the replay of messages to the Topic Endpoint.
[**GetMsgVpnTopicEndpoint**](TopicEndpointApi.md#GetMsgVpnTopicEndpoint) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName} | Get a Topic Endpoint object.
[**GetMsgVpnTopicEndpointMsg**](TopicEndpointApi.md#GetMsgVpnTopicEndpointMsg) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/msgs/{msgId} | Get a Topic Endpoint Message object.
[**GetMsgVpnTopicEndpointMsgs**](TopicEndpointApi.md#GetMsgVpnTopicEndpointMsgs) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/msgs | Get a list of Topic Endpoint Message objects.
[**GetMsgVpnTopicEndpoints**](TopicEndpointApi.md#GetMsgVpnTopicEndpoints) | **Get** /msgVpns/{msgVpnName}/topicEndpoints | Get a list of Topic Endpoint objects.

# **DoMsgVpnTopicEndpointCancelReplay**
> SempMetaOnlyResponse DoMsgVpnTopicEndpointCancelReplay(ctx, body, msgVpnName, topicEndpointName)
Cancel the replay of messages to the Topic Endpoint.

Cancel the replay of messages to the Topic Endpoint.    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnTopicEndpointCancelReplay**](MsgVpnTopicEndpointCancelReplay.md)| The Cancel Replay action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointName** | **string**| The name of the Topic Endpoint. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnTopicEndpointClearStats**
> SempMetaOnlyResponse DoMsgVpnTopicEndpointClearStats(ctx, body, msgVpnName, topicEndpointName)
Clear the statistics for the Topic Endpoint.

Clear the statistics for the Topic Endpoint.    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnTopicEndpointClearStats**](MsgVpnTopicEndpointClearStats.md)| The Clear Stats action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointName** | **string**| The name of the Topic Endpoint. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnTopicEndpointCopyMsgFromQueue**
> SempMetaOnlyResponse DoMsgVpnTopicEndpointCopyMsgFromQueue(ctx, body, msgVpnName, topicEndpointName)
Copy a message from a Queue to this Topic Endpoint.

Copy a message from a Queue to this Topic Endpoint.   Attribute|Required|Deprecated :---|:---:|:---: replicationGroupMsgId|x| sourceQueueName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.29.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnTopicEndpointCopyMsgFromQueue**](MsgVpnTopicEndpointCopyMsgFromQueue.md)| The Copy Message From Queue action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointName** | **string**| The name of the Topic Endpoint. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnTopicEndpointCopyMsgFromReplayLog**
> SempMetaOnlyResponse DoMsgVpnTopicEndpointCopyMsgFromReplayLog(ctx, body, msgVpnName, topicEndpointName)
Copy a message from a Replay Log to this Topic Endpoint.

Copy a message from a Replay Log to this Topic Endpoint.   Attribute|Required|Deprecated :---|:---:|:---: replicationGroupMsgId|x| sourceReplayLogName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.29.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnTopicEndpointCopyMsgFromReplayLog**](MsgVpnTopicEndpointCopyMsgFromReplayLog.md)| The Copy Message From Replay Log action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointName** | **string**| The name of the Topic Endpoint. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnTopicEndpointCopyMsgFromTopicEndpoint**
> SempMetaOnlyResponse DoMsgVpnTopicEndpointCopyMsgFromTopicEndpoint(ctx, body, msgVpnName, topicEndpointName)
Copy a message from another Topic Endpoint to this Topic Endpoint.

Copy a message from another Topic Endpoint to this Topic Endpoint.   Attribute|Required|Deprecated :---|:---:|:---: replicationGroupMsgId|x| sourceTopicEndpointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.29.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnTopicEndpointCopyMsgFromTopicEndpoint**](MsgVpnTopicEndpointCopyMsgFromTopicEndpoint.md)| The Copy Message From Topic Endpoint action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointName** | **string**| The name of the Topic Endpoint. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnTopicEndpointDeleteMsgs**
> SempMetaOnlyResponse DoMsgVpnTopicEndpointDeleteMsgs(ctx, body, msgVpnName, topicEndpointName)
Delete all spooled messages from the Topic Endpoint.

Delete all spooled messages from the Topic Endpoint.    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.28.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnTopicEndpointDeleteMsgs**](MsgVpnTopicEndpointDeleteMsgs.md)| The Delete All Messages action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointName** | **string**| The name of the Topic Endpoint. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnTopicEndpointMsgDelete**
> SempMetaOnlyResponse DoMsgVpnTopicEndpointMsgDelete(ctx, body, msgVpnName, topicEndpointName, msgId)
Delete the Message from the Topic Endpoint.

Delete the Message from the Topic Endpoint.    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnTopicEndpointMsgDelete**](MsgVpnTopicEndpointMsgDelete.md)| The Delete action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointName** | **string**| The name of the Topic Endpoint. | 
  **msgId** | **string**| The identifier (ID) of the Message. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnTopicEndpointStartReplay**
> SempMetaOnlyResponse DoMsgVpnTopicEndpointStartReplay(ctx, body, msgVpnName, topicEndpointName)
Start the replay of messages to the Topic Endpoint.

Start the replay of messages to the Topic Endpoint.    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnTopicEndpointStartReplay**](MsgVpnTopicEndpointStartReplay.md)| The Start Replay action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointName** | **string**| The name of the Topic Endpoint. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnTopicEndpoint**
> MsgVpnTopicEndpointResponse GetMsgVpnTopicEndpoint(ctx, msgVpnName, topicEndpointName, optional)
Get a Topic Endpoint object.

Get a Topic Endpoint object.  A Topic Endpoint attracts messages published to a topic for which the Topic Endpoint has a matching topic subscription. The topic subscription for the Topic Endpoint is specified in the client request to bind a Flow to that Topic Endpoint. Queues are significantly more flexible than Topic Endpoints and are the recommended approach for most applications. The use of Topic Endpoints should be restricted to JMS applications.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| topicEndpointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointName** | **string**| The name of the Topic Endpoint. | 
 **optional** | ***TopicEndpointApiGetMsgVpnTopicEndpointOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TopicEndpointApiGetMsgVpnTopicEndpointOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointResponse**](MsgVpnTopicEndpointResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnTopicEndpointMsg**
> MsgVpnTopicEndpointMsgResponse GetMsgVpnTopicEndpointMsg(ctx, msgVpnName, topicEndpointName, msgId, optional)
Get a Topic Endpoint Message object.

Get a Topic Endpoint Message object.  A Topic Endpoint Message is a packet of information sent from producers to consumers using the Topic Endpoint.   Attribute|Identifying|Deprecated :---|:---:|:---: msgId|x| msgVpnName|x| topicEndpointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointName** | **string**| The name of the Topic Endpoint. | 
  **msgId** | **string**| The identifier (ID) of the Message. | 
 **optional** | ***TopicEndpointApiGetMsgVpnTopicEndpointMsgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TopicEndpointApiGetMsgVpnTopicEndpointMsgOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointMsgResponse**](MsgVpnTopicEndpointMsgResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnTopicEndpointMsgs**
> MsgVpnTopicEndpointMsgsResponse GetMsgVpnTopicEndpointMsgs(ctx, msgVpnName, topicEndpointName, optional)
Get a list of Topic Endpoint Message objects.

Get a list of Topic Endpoint Message objects.  A Topic Endpoint Message is a packet of information sent from producers to consumers using the Topic Endpoint.   Attribute|Identifying|Deprecated :---|:---:|:---: msgId|x| msgVpnName|x| topicEndpointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointName** | **string**| The name of the Topic Endpoint. | 
 **optional** | ***TopicEndpointApiGetMsgVpnTopicEndpointMsgsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TopicEndpointApiGetMsgVpnTopicEndpointMsgsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointMsgsResponse**](MsgVpnTopicEndpointMsgsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnTopicEndpoints**
> MsgVpnTopicEndpointsResponse GetMsgVpnTopicEndpoints(ctx, msgVpnName, optional)
Get a list of Topic Endpoint objects.

Get a list of Topic Endpoint objects.  A Topic Endpoint attracts messages published to a topic for which the Topic Endpoint has a matching topic subscription. The topic subscription for the Topic Endpoint is specified in the client request to bind a Flow to that Topic Endpoint. Queues are significantly more flexible than Topic Endpoints and are the recommended approach for most applications. The use of Topic Endpoints should be restricted to JMS applications.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| topicEndpointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***TopicEndpointApiGetMsgVpnTopicEndpointsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TopicEndpointApiGetMsgVpnTopicEndpointsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointsResponse**](MsgVpnTopicEndpointsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

