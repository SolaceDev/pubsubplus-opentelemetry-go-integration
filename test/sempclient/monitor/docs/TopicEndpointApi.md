# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/monitor*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMsgVpnTopicEndpoint**](TopicEndpointApi.md#GetMsgVpnTopicEndpoint) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName} | Get a Topic Endpoint object.
[**GetMsgVpnTopicEndpointMsg**](TopicEndpointApi.md#GetMsgVpnTopicEndpointMsg) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/msgs/{msgId} | Get a Topic Endpoint Message object.
[**GetMsgVpnTopicEndpointMsgs**](TopicEndpointApi.md#GetMsgVpnTopicEndpointMsgs) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/msgs | Get a list of Topic Endpoint Message objects.
[**GetMsgVpnTopicEndpointPriorities**](TopicEndpointApi.md#GetMsgVpnTopicEndpointPriorities) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/priorities | Get a list of Topic Endpoint Priority objects.
[**GetMsgVpnTopicEndpointPriority**](TopicEndpointApi.md#GetMsgVpnTopicEndpointPriority) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/priorities/{priority} | Get a Topic Endpoint Priority object.
[**GetMsgVpnTopicEndpointTxFlow**](TopicEndpointApi.md#GetMsgVpnTopicEndpointTxFlow) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/txFlows/{flowId} | Get a Topic Endpoint Transmit Flow object.
[**GetMsgVpnTopicEndpointTxFlows**](TopicEndpointApi.md#GetMsgVpnTopicEndpointTxFlows) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/txFlows | Get a list of Topic Endpoint Transmit Flow objects.
[**GetMsgVpnTopicEndpoints**](TopicEndpointApi.md#GetMsgVpnTopicEndpoints) | **Get** /msgVpns/{msgVpnName}/topicEndpoints | Get a list of Topic Endpoint objects.

# **GetMsgVpnTopicEndpoint**
> MsgVpnTopicEndpointResponse GetMsgVpnTopicEndpoint(ctx, msgVpnName, topicEndpointName, optional)
Get a Topic Endpoint object.

Get a Topic Endpoint object.  A Topic Endpoint attracts messages published to a topic for which the Topic Endpoint has a matching topic subscription. The topic subscription for the Topic Endpoint is specified in the client request to bind a Flow to that Topic Endpoint. Queues are significantly more flexible than Topic Endpoints and are the recommended approach for most applications. The use of Topic Endpoints should be restricted to JMS applications.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| topicEndpointName|x| virtualRouter||x    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

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

Get a Topic Endpoint Message object.  A Topic Endpoint Message is a packet of information sent from producers to consumers using the Topic Endpoint.   Attribute|Identifying|Deprecated :---|:---:|:---: msgId|x| msgVpnName|x| topicEndpointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

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

Get a list of Topic Endpoint Message objects.  A Topic Endpoint Message is a packet of information sent from producers to consumers using the Topic Endpoint.   Attribute|Identifying|Deprecated :---|:---:|:---: msgId|x| msgVpnName|x| topicEndpointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

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

# **GetMsgVpnTopicEndpointPriorities**
> MsgVpnTopicEndpointPrioritiesResponse GetMsgVpnTopicEndpointPriorities(ctx, msgVpnName, topicEndpointName, optional)
Get a list of Topic Endpoint Priority objects.

Get a list of Topic Endpoint Priority objects.  Topic Endpoints can optionally support priority message delivery; all messages of a higher priority are delivered before any messages of a lower priority. A Priority object contains information about the number and size of the messages with a particular priority in the Topic Endpoint.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| priority|x| topicEndpointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointName** | **string**| The name of the Topic Endpoint. | 
 **optional** | ***TopicEndpointApiGetMsgVpnTopicEndpointPrioritiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TopicEndpointApiGetMsgVpnTopicEndpointPrioritiesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointPrioritiesResponse**](MsgVpnTopicEndpointPrioritiesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnTopicEndpointPriority**
> MsgVpnTopicEndpointPriorityResponse GetMsgVpnTopicEndpointPriority(ctx, msgVpnName, topicEndpointName, priority, optional)
Get a Topic Endpoint Priority object.

Get a Topic Endpoint Priority object.  Topic Endpoints can optionally support priority message delivery; all messages of a higher priority are delivered before any messages of a lower priority. A Priority object contains information about the number and size of the messages with a particular priority in the Topic Endpoint.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| priority|x| topicEndpointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointName** | **string**| The name of the Topic Endpoint. | 
  **priority** | **string**| The level of the Priority, from 9 (highest) to 0 (lowest). | 
 **optional** | ***TopicEndpointApiGetMsgVpnTopicEndpointPriorityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TopicEndpointApiGetMsgVpnTopicEndpointPriorityOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointPriorityResponse**](MsgVpnTopicEndpointPriorityResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnTopicEndpointTxFlow**
> MsgVpnTopicEndpointTxFlowResponse GetMsgVpnTopicEndpointTxFlow(ctx, msgVpnName, topicEndpointName, flowId, optional)
Get a Topic Endpoint Transmit Flow object.

Get a Topic Endpoint Transmit Flow object.  Topic Endpoint Transmit Flows are used by clients to consume Guaranteed messages from a Topic Endpoint.   Attribute|Identifying|Deprecated :---|:---:|:---: flowId|x| msgVpnName|x| topicEndpointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointName** | **string**| The name of the Topic Endpoint. | 
  **flowId** | **string**| The identifier (ID) of the Flow. | 
 **optional** | ***TopicEndpointApiGetMsgVpnTopicEndpointTxFlowOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TopicEndpointApiGetMsgVpnTopicEndpointTxFlowOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointTxFlowResponse**](MsgVpnTopicEndpointTxFlowResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnTopicEndpointTxFlows**
> MsgVpnTopicEndpointTxFlowsResponse GetMsgVpnTopicEndpointTxFlows(ctx, msgVpnName, topicEndpointName, optional)
Get a list of Topic Endpoint Transmit Flow objects.

Get a list of Topic Endpoint Transmit Flow objects.  Topic Endpoint Transmit Flows are used by clients to consume Guaranteed messages from a Topic Endpoint.   Attribute|Identifying|Deprecated :---|:---:|:---: flowId|x| msgVpnName|x| topicEndpointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointName** | **string**| The name of the Topic Endpoint. | 
 **optional** | ***TopicEndpointApiGetMsgVpnTopicEndpointTxFlowsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TopicEndpointApiGetMsgVpnTopicEndpointTxFlowsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointTxFlowsResponse**](MsgVpnTopicEndpointTxFlowsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnTopicEndpoints**
> MsgVpnTopicEndpointsResponse GetMsgVpnTopicEndpoints(ctx, msgVpnName, optional)
Get a list of Topic Endpoint objects.

Get a list of Topic Endpoint objects.  A Topic Endpoint attracts messages published to a topic for which the Topic Endpoint has a matching topic subscription. The topic subscription for the Topic Endpoint is specified in the client request to bind a Flow to that Topic Endpoint. Queues are significantly more flexible than Topic Endpoints and are the recommended approach for most applications. The use of Topic Endpoints should be restricted to JMS applications.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| topicEndpointName|x| virtualRouter||x    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

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

