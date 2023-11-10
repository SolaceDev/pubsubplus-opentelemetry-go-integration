# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/config*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMsgVpnReplicatedTopic**](ReplicatedTopicApi.md#CreateMsgVpnReplicatedTopic) | **Post** /msgVpns/{msgVpnName}/replicatedTopics | Create a Replicated Topic object.
[**DeleteMsgVpnReplicatedTopic**](ReplicatedTopicApi.md#DeleteMsgVpnReplicatedTopic) | **Delete** /msgVpns/{msgVpnName}/replicatedTopics/{replicatedTopic} | Delete a Replicated Topic object.
[**GetMsgVpnReplicatedTopic**](ReplicatedTopicApi.md#GetMsgVpnReplicatedTopic) | **Get** /msgVpns/{msgVpnName}/replicatedTopics/{replicatedTopic} | Get a Replicated Topic object.
[**GetMsgVpnReplicatedTopics**](ReplicatedTopicApi.md#GetMsgVpnReplicatedTopics) | **Get** /msgVpns/{msgVpnName}/replicatedTopics | Get a list of Replicated Topic objects.
[**ReplaceMsgVpnReplicatedTopic**](ReplicatedTopicApi.md#ReplaceMsgVpnReplicatedTopic) | **Put** /msgVpns/{msgVpnName}/replicatedTopics/{replicatedTopic} | Replace a Replicated Topic object.
[**UpdateMsgVpnReplicatedTopic**](ReplicatedTopicApi.md#UpdateMsgVpnReplicatedTopic) | **Patch** /msgVpns/{msgVpnName}/replicatedTopics/{replicatedTopic} | Update a Replicated Topic object.

# **CreateMsgVpnReplicatedTopic**
> MsgVpnReplicatedTopicResponse CreateMsgVpnReplicatedTopic(ctx, body, msgVpnName, optional)
Create a Replicated Topic object.

Create a Replicated Topic object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  To indicate which messages should be replicated between the active and standby site, a Replicated Topic subscription must be configured on a Message VPN. If a published message matches both a replicated topic and an endpoint on the active site, then the message is replicated to the standby site.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x||| replicatedTopic|x|x||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.1.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnReplicatedTopic**](MsgVpnReplicatedTopic.md)| The Replicated Topic object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***ReplicatedTopicApiCreateMsgVpnReplicatedTopicOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplicatedTopicApiCreateMsgVpnReplicatedTopicOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnReplicatedTopicResponse**](MsgVpnReplicatedTopicResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnReplicatedTopic**
> SempMetaOnlyResponse DeleteMsgVpnReplicatedTopic(ctx, msgVpnName, replicatedTopic)
Delete a Replicated Topic object.

Delete a Replicated Topic object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  To indicate which messages should be replicated between the active and standby site, a Replicated Topic subscription must be configured on a Message VPN. If a published message matches both a replicated topic and an endpoint on the active site, then the message is replicated to the standby site.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.1.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **replicatedTopic** | **string**| The topic for applying replication. Published messages matching this topic will be replicated to the standby site. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnReplicatedTopic**
> MsgVpnReplicatedTopicResponse GetMsgVpnReplicatedTopic(ctx, msgVpnName, replicatedTopic, optional)
Get a Replicated Topic object.

Get a Replicated Topic object.  To indicate which messages should be replicated between the active and standby site, a Replicated Topic subscription must be configured on a Message VPN. If a published message matches both a replicated topic and an endpoint on the active site, then the message is replicated to the standby site.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| replicatedTopic|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.1.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **replicatedTopic** | **string**| The topic for applying replication. Published messages matching this topic will be replicated to the standby site. | 
 **optional** | ***ReplicatedTopicApiGetMsgVpnReplicatedTopicOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplicatedTopicApiGetMsgVpnReplicatedTopicOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnReplicatedTopicResponse**](MsgVpnReplicatedTopicResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnReplicatedTopics**
> MsgVpnReplicatedTopicsResponse GetMsgVpnReplicatedTopics(ctx, msgVpnName, optional)
Get a list of Replicated Topic objects.

Get a list of Replicated Topic objects.  To indicate which messages should be replicated between the active and standby site, a Replicated Topic subscription must be configured on a Message VPN. If a published message matches both a replicated topic and an endpoint on the active site, then the message is replicated to the standby site.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| replicatedTopic|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.1.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***ReplicatedTopicApiGetMsgVpnReplicatedTopicsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplicatedTopicApiGetMsgVpnReplicatedTopicsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnReplicatedTopicsResponse**](MsgVpnReplicatedTopicsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnReplicatedTopic**
> MsgVpnReplicatedTopicResponse ReplaceMsgVpnReplicatedTopic(ctx, body, msgVpnName, replicatedTopic, optional)
Replace a Replicated Topic object.

Replace a Replicated Topic object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  To indicate which messages should be replicated between the active and standby site, a Replicated Topic subscription must be configured on a Message VPN. If a published message matches both a replicated topic and an endpoint on the active site, then the message is replicated to the standby site.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- msgVpnName|x||x||||| replicatedTopic|x||x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.1.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnReplicatedTopic**](MsgVpnReplicatedTopic.md)| The Replicated Topic object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **replicatedTopic** | **string**| The topic for applying replication. Published messages matching this topic will be replicated to the standby site. | 
 **optional** | ***ReplicatedTopicApiReplaceMsgVpnReplicatedTopicOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplicatedTopicApiReplaceMsgVpnReplicatedTopicOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnReplicatedTopicResponse**](MsgVpnReplicatedTopicResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnReplicatedTopic**
> MsgVpnReplicatedTopicResponse UpdateMsgVpnReplicatedTopic(ctx, body, msgVpnName, replicatedTopic, optional)
Update a Replicated Topic object.

Update a Replicated Topic object. Any attribute missing from the request will be left unchanged.  To indicate which messages should be replicated between the active and standby site, a Replicated Topic subscription must be configured on a Message VPN. If a published message matches both a replicated topic and an endpoint on the active site, then the message is replicated to the standby site.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- msgVpnName|x|x||||| replicatedTopic|x|x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.1.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnReplicatedTopic**](MsgVpnReplicatedTopic.md)| The Replicated Topic object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **replicatedTopic** | **string**| The topic for applying replication. Published messages matching this topic will be replicated to the standby site. | 
 **optional** | ***ReplicatedTopicApiUpdateMsgVpnReplicatedTopicOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplicatedTopicApiUpdateMsgVpnReplicatedTopicOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnReplicatedTopicResponse**](MsgVpnReplicatedTopicResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

