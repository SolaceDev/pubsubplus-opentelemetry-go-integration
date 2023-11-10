# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/config*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMsgVpnJndiConnectionFactory**](JndiApi.md#CreateMsgVpnJndiConnectionFactory) | **Post** /msgVpns/{msgVpnName}/jndiConnectionFactories | Create a JNDI Connection Factory object.
[**CreateMsgVpnJndiQueue**](JndiApi.md#CreateMsgVpnJndiQueue) | **Post** /msgVpns/{msgVpnName}/jndiQueues | Create a JNDI Queue object.
[**CreateMsgVpnJndiTopic**](JndiApi.md#CreateMsgVpnJndiTopic) | **Post** /msgVpns/{msgVpnName}/jndiTopics | Create a JNDI Topic object.
[**DeleteMsgVpnJndiConnectionFactory**](JndiApi.md#DeleteMsgVpnJndiConnectionFactory) | **Delete** /msgVpns/{msgVpnName}/jndiConnectionFactories/{connectionFactoryName} | Delete a JNDI Connection Factory object.
[**DeleteMsgVpnJndiQueue**](JndiApi.md#DeleteMsgVpnJndiQueue) | **Delete** /msgVpns/{msgVpnName}/jndiQueues/{queueName} | Delete a JNDI Queue object.
[**DeleteMsgVpnJndiTopic**](JndiApi.md#DeleteMsgVpnJndiTopic) | **Delete** /msgVpns/{msgVpnName}/jndiTopics/{topicName} | Delete a JNDI Topic object.
[**GetMsgVpnJndiConnectionFactories**](JndiApi.md#GetMsgVpnJndiConnectionFactories) | **Get** /msgVpns/{msgVpnName}/jndiConnectionFactories | Get a list of JNDI Connection Factory objects.
[**GetMsgVpnJndiConnectionFactory**](JndiApi.md#GetMsgVpnJndiConnectionFactory) | **Get** /msgVpns/{msgVpnName}/jndiConnectionFactories/{connectionFactoryName} | Get a JNDI Connection Factory object.
[**GetMsgVpnJndiQueue**](JndiApi.md#GetMsgVpnJndiQueue) | **Get** /msgVpns/{msgVpnName}/jndiQueues/{queueName} | Get a JNDI Queue object.
[**GetMsgVpnJndiQueues**](JndiApi.md#GetMsgVpnJndiQueues) | **Get** /msgVpns/{msgVpnName}/jndiQueues | Get a list of JNDI Queue objects.
[**GetMsgVpnJndiTopic**](JndiApi.md#GetMsgVpnJndiTopic) | **Get** /msgVpns/{msgVpnName}/jndiTopics/{topicName} | Get a JNDI Topic object.
[**GetMsgVpnJndiTopics**](JndiApi.md#GetMsgVpnJndiTopics) | **Get** /msgVpns/{msgVpnName}/jndiTopics | Get a list of JNDI Topic objects.
[**ReplaceMsgVpnJndiConnectionFactory**](JndiApi.md#ReplaceMsgVpnJndiConnectionFactory) | **Put** /msgVpns/{msgVpnName}/jndiConnectionFactories/{connectionFactoryName} | Replace a JNDI Connection Factory object.
[**ReplaceMsgVpnJndiQueue**](JndiApi.md#ReplaceMsgVpnJndiQueue) | **Put** /msgVpns/{msgVpnName}/jndiQueues/{queueName} | Replace a JNDI Queue object.
[**ReplaceMsgVpnJndiTopic**](JndiApi.md#ReplaceMsgVpnJndiTopic) | **Put** /msgVpns/{msgVpnName}/jndiTopics/{topicName} | Replace a JNDI Topic object.
[**UpdateMsgVpnJndiConnectionFactory**](JndiApi.md#UpdateMsgVpnJndiConnectionFactory) | **Patch** /msgVpns/{msgVpnName}/jndiConnectionFactories/{connectionFactoryName} | Update a JNDI Connection Factory object.
[**UpdateMsgVpnJndiQueue**](JndiApi.md#UpdateMsgVpnJndiQueue) | **Patch** /msgVpns/{msgVpnName}/jndiQueues/{queueName} | Update a JNDI Queue object.
[**UpdateMsgVpnJndiTopic**](JndiApi.md#UpdateMsgVpnJndiTopic) | **Patch** /msgVpns/{msgVpnName}/jndiTopics/{topicName} | Update a JNDI Topic object.

# **CreateMsgVpnJndiConnectionFactory**
> MsgVpnJndiConnectionFactoryResponse CreateMsgVpnJndiConnectionFactory(ctx, body, msgVpnName, optional)
Create a JNDI Connection Factory object.

Create a JNDI Connection Factory object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  The message broker provides an internal JNDI store for provisioned Connection Factory objects that clients can access through JNDI lookups.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: connectionFactoryName|x|x|||| msgVpnName|x||x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.2.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnJndiConnectionFactory**](MsgVpnJndiConnectionFactory.md)| The JNDI Connection Factory object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***JndiApiCreateMsgVpnJndiConnectionFactoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JndiApiCreateMsgVpnJndiConnectionFactoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnJndiConnectionFactoryResponse**](MsgVpnJndiConnectionFactoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnJndiQueue**
> MsgVpnJndiQueueResponse CreateMsgVpnJndiQueue(ctx, body, msgVpnName, optional)
Create a JNDI Queue object.

Create a JNDI Queue object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  The message broker provides an internal JNDI store for provisioned Queue objects that clients can access through JNDI lookups.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x||| queueName|x|x||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.2.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnJndiQueue**](MsgVpnJndiQueue.md)| The JNDI Queue object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***JndiApiCreateMsgVpnJndiQueueOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JndiApiCreateMsgVpnJndiQueueOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnJndiQueueResponse**](MsgVpnJndiQueueResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnJndiTopic**
> MsgVpnJndiTopicResponse CreateMsgVpnJndiTopic(ctx, body, msgVpnName, optional)
Create a JNDI Topic object.

Create a JNDI Topic object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  The message broker provides an internal JNDI store for provisioned Topic objects that clients can access through JNDI lookups.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x||| topicName|x|x||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.2.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnJndiTopic**](MsgVpnJndiTopic.md)| The JNDI Topic object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***JndiApiCreateMsgVpnJndiTopicOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JndiApiCreateMsgVpnJndiTopicOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnJndiTopicResponse**](MsgVpnJndiTopicResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnJndiConnectionFactory**
> SempMetaOnlyResponse DeleteMsgVpnJndiConnectionFactory(ctx, msgVpnName, connectionFactoryName)
Delete a JNDI Connection Factory object.

Delete a JNDI Connection Factory object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  The message broker provides an internal JNDI store for provisioned Connection Factory objects that clients can access through JNDI lookups.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.2.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **connectionFactoryName** | **string**| The name of the JMS Connection Factory. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnJndiQueue**
> SempMetaOnlyResponse DeleteMsgVpnJndiQueue(ctx, msgVpnName, queueName)
Delete a JNDI Queue object.

Delete a JNDI Queue object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  The message broker provides an internal JNDI store for provisioned Queue objects that clients can access through JNDI lookups.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.2.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The JNDI name of the JMS Queue. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnJndiTopic**
> SempMetaOnlyResponse DeleteMsgVpnJndiTopic(ctx, msgVpnName, topicName)
Delete a JNDI Topic object.

Delete a JNDI Topic object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  The message broker provides an internal JNDI store for provisioned Topic objects that clients can access through JNDI lookups.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.2.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicName** | **string**| The JNDI name of the JMS Topic. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnJndiConnectionFactories**
> MsgVpnJndiConnectionFactoriesResponse GetMsgVpnJndiConnectionFactories(ctx, msgVpnName, optional)
Get a list of JNDI Connection Factory objects.

Get a list of JNDI Connection Factory objects.  The message broker provides an internal JNDI store for provisioned Connection Factory objects that clients can access through JNDI lookups.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: connectionFactoryName|x||| msgVpnName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.2.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***JndiApiGetMsgVpnJndiConnectionFactoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JndiApiGetMsgVpnJndiConnectionFactoriesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnJndiConnectionFactoriesResponse**](MsgVpnJndiConnectionFactoriesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnJndiConnectionFactory**
> MsgVpnJndiConnectionFactoryResponse GetMsgVpnJndiConnectionFactory(ctx, msgVpnName, connectionFactoryName, optional)
Get a JNDI Connection Factory object.

Get a JNDI Connection Factory object.  The message broker provides an internal JNDI store for provisioned Connection Factory objects that clients can access through JNDI lookups.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: connectionFactoryName|x||| msgVpnName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.2.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **connectionFactoryName** | **string**| The name of the JMS Connection Factory. | 
 **optional** | ***JndiApiGetMsgVpnJndiConnectionFactoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JndiApiGetMsgVpnJndiConnectionFactoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnJndiConnectionFactoryResponse**](MsgVpnJndiConnectionFactoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnJndiQueue**
> MsgVpnJndiQueueResponse GetMsgVpnJndiQueue(ctx, msgVpnName, queueName, optional)
Get a JNDI Queue object.

Get a JNDI Queue object.  The message broker provides an internal JNDI store for provisioned Queue objects that clients can access through JNDI lookups.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| queueName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.2.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The JNDI name of the JMS Queue. | 
 **optional** | ***JndiApiGetMsgVpnJndiQueueOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JndiApiGetMsgVpnJndiQueueOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnJndiQueueResponse**](MsgVpnJndiQueueResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnJndiQueues**
> MsgVpnJndiQueuesResponse GetMsgVpnJndiQueues(ctx, msgVpnName, optional)
Get a list of JNDI Queue objects.

Get a list of JNDI Queue objects.  The message broker provides an internal JNDI store for provisioned Queue objects that clients can access through JNDI lookups.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| queueName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.2.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***JndiApiGetMsgVpnJndiQueuesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JndiApiGetMsgVpnJndiQueuesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnJndiQueuesResponse**](MsgVpnJndiQueuesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnJndiTopic**
> MsgVpnJndiTopicResponse GetMsgVpnJndiTopic(ctx, msgVpnName, topicName, optional)
Get a JNDI Topic object.

Get a JNDI Topic object.  The message broker provides an internal JNDI store for provisioned Topic objects that clients can access through JNDI lookups.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| topicName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.2.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicName** | **string**| The JNDI name of the JMS Topic. | 
 **optional** | ***JndiApiGetMsgVpnJndiTopicOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JndiApiGetMsgVpnJndiTopicOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnJndiTopicResponse**](MsgVpnJndiTopicResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnJndiTopics**
> MsgVpnJndiTopicsResponse GetMsgVpnJndiTopics(ctx, msgVpnName, optional)
Get a list of JNDI Topic objects.

Get a list of JNDI Topic objects.  The message broker provides an internal JNDI store for provisioned Topic objects that clients can access through JNDI lookups.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| topicName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.2.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***JndiApiGetMsgVpnJndiTopicsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JndiApiGetMsgVpnJndiTopicsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnJndiTopicsResponse**](MsgVpnJndiTopicsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnJndiConnectionFactory**
> MsgVpnJndiConnectionFactoryResponse ReplaceMsgVpnJndiConnectionFactory(ctx, body, msgVpnName, connectionFactoryName, optional)
Replace a JNDI Connection Factory object.

Replace a JNDI Connection Factory object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  The message broker provides an internal JNDI store for provisioned Connection Factory objects that clients can access through JNDI lookups.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- connectionFactoryName|x||x||||| msgVpnName|x||x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.2.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnJndiConnectionFactory**](MsgVpnJndiConnectionFactory.md)| The JNDI Connection Factory object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **connectionFactoryName** | **string**| The name of the JMS Connection Factory. | 
 **optional** | ***JndiApiReplaceMsgVpnJndiConnectionFactoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JndiApiReplaceMsgVpnJndiConnectionFactoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnJndiConnectionFactoryResponse**](MsgVpnJndiConnectionFactoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnJndiQueue**
> MsgVpnJndiQueueResponse ReplaceMsgVpnJndiQueue(ctx, body, msgVpnName, queueName, optional)
Replace a JNDI Queue object.

Replace a JNDI Queue object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  The message broker provides an internal JNDI store for provisioned Queue objects that clients can access through JNDI lookups.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- msgVpnName|x||x||||| queueName|x||x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.2.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnJndiQueue**](MsgVpnJndiQueue.md)| The JNDI Queue object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The JNDI name of the JMS Queue. | 
 **optional** | ***JndiApiReplaceMsgVpnJndiQueueOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JndiApiReplaceMsgVpnJndiQueueOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnJndiQueueResponse**](MsgVpnJndiQueueResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnJndiTopic**
> MsgVpnJndiTopicResponse ReplaceMsgVpnJndiTopic(ctx, body, msgVpnName, topicName, optional)
Replace a JNDI Topic object.

Replace a JNDI Topic object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  The message broker provides an internal JNDI store for provisioned Topic objects that clients can access through JNDI lookups.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- msgVpnName|x||x||||| topicName|x||x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.2.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnJndiTopic**](MsgVpnJndiTopic.md)| The JNDI Topic object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicName** | **string**| The JNDI name of the JMS Topic. | 
 **optional** | ***JndiApiReplaceMsgVpnJndiTopicOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JndiApiReplaceMsgVpnJndiTopicOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnJndiTopicResponse**](MsgVpnJndiTopicResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnJndiConnectionFactory**
> MsgVpnJndiConnectionFactoryResponse UpdateMsgVpnJndiConnectionFactory(ctx, body, msgVpnName, connectionFactoryName, optional)
Update a JNDI Connection Factory object.

Update a JNDI Connection Factory object. Any attribute missing from the request will be left unchanged.  The message broker provides an internal JNDI store for provisioned Connection Factory objects that clients can access through JNDI lookups.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- connectionFactoryName|x|x||||| msgVpnName|x|x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.2.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnJndiConnectionFactory**](MsgVpnJndiConnectionFactory.md)| The JNDI Connection Factory object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **connectionFactoryName** | **string**| The name of the JMS Connection Factory. | 
 **optional** | ***JndiApiUpdateMsgVpnJndiConnectionFactoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JndiApiUpdateMsgVpnJndiConnectionFactoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnJndiConnectionFactoryResponse**](MsgVpnJndiConnectionFactoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnJndiQueue**
> MsgVpnJndiQueueResponse UpdateMsgVpnJndiQueue(ctx, body, msgVpnName, queueName, optional)
Update a JNDI Queue object.

Update a JNDI Queue object. Any attribute missing from the request will be left unchanged.  The message broker provides an internal JNDI store for provisioned Queue objects that clients can access through JNDI lookups.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- msgVpnName|x|x||||| queueName|x|x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.2.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnJndiQueue**](MsgVpnJndiQueue.md)| The JNDI Queue object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The JNDI name of the JMS Queue. | 
 **optional** | ***JndiApiUpdateMsgVpnJndiQueueOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JndiApiUpdateMsgVpnJndiQueueOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnJndiQueueResponse**](MsgVpnJndiQueueResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnJndiTopic**
> MsgVpnJndiTopicResponse UpdateMsgVpnJndiTopic(ctx, body, msgVpnName, topicName, optional)
Update a JNDI Topic object.

Update a JNDI Topic object. Any attribute missing from the request will be left unchanged.  The message broker provides an internal JNDI store for provisioned Topic objects that clients can access through JNDI lookups.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- msgVpnName|x|x||||| topicName|x|x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.2.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnJndiTopic**](MsgVpnJndiTopic.md)| The JNDI Topic object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicName** | **string**| The JNDI name of the JMS Topic. | 
 **optional** | ***JndiApiUpdateMsgVpnJndiTopicOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JndiApiUpdateMsgVpnJndiTopicOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnJndiTopicResponse**](MsgVpnJndiTopicResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

