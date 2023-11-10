# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/config*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMsgVpnTopicEndpointTemplate**](TopicEndpointTemplateApi.md#CreateMsgVpnTopicEndpointTemplate) | **Post** /msgVpns/{msgVpnName}/topicEndpointTemplates | Create a Topic Endpoint Template object.
[**DeleteMsgVpnTopicEndpointTemplate**](TopicEndpointTemplateApi.md#DeleteMsgVpnTopicEndpointTemplate) | **Delete** /msgVpns/{msgVpnName}/topicEndpointTemplates/{topicEndpointTemplateName} | Delete a Topic Endpoint Template object.
[**GetMsgVpnTopicEndpointTemplate**](TopicEndpointTemplateApi.md#GetMsgVpnTopicEndpointTemplate) | **Get** /msgVpns/{msgVpnName}/topicEndpointTemplates/{topicEndpointTemplateName} | Get a Topic Endpoint Template object.
[**GetMsgVpnTopicEndpointTemplates**](TopicEndpointTemplateApi.md#GetMsgVpnTopicEndpointTemplates) | **Get** /msgVpns/{msgVpnName}/topicEndpointTemplates | Get a list of Topic Endpoint Template objects.
[**ReplaceMsgVpnTopicEndpointTemplate**](TopicEndpointTemplateApi.md#ReplaceMsgVpnTopicEndpointTemplate) | **Put** /msgVpns/{msgVpnName}/topicEndpointTemplates/{topicEndpointTemplateName} | Replace a Topic Endpoint Template object.
[**UpdateMsgVpnTopicEndpointTemplate**](TopicEndpointTemplateApi.md#UpdateMsgVpnTopicEndpointTemplate) | **Patch** /msgVpns/{msgVpnName}/topicEndpointTemplates/{topicEndpointTemplateName} | Update a Topic Endpoint Template object.

# **CreateMsgVpnTopicEndpointTemplate**
> MsgVpnTopicEndpointTemplateResponse CreateMsgVpnTopicEndpointTemplate(ctx, body, msgVpnName, optional)
Create a Topic Endpoint Template object.

Create a Topic Endpoint Template object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Topic Endpoint Template provides a mechanism for specifying the initial state for client created topic endpoints.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x||| topicEndpointTemplateName|x|x||||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThreshold|clearPercent|setPercent|clearValue, setValue EventThreshold|clearValue|setValue|clearPercent, setPercent EventThreshold|setPercent|clearPercent|clearValue, setValue EventThreshold|setValue|clearValue|clearPercent, setPercent    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnTopicEndpointTemplate**](MsgVpnTopicEndpointTemplate.md)| The Topic Endpoint Template object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***TopicEndpointTemplateApiCreateMsgVpnTopicEndpointTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TopicEndpointTemplateApiCreateMsgVpnTopicEndpointTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointTemplateResponse**](MsgVpnTopicEndpointTemplateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnTopicEndpointTemplate**
> SempMetaOnlyResponse DeleteMsgVpnTopicEndpointTemplate(ctx, msgVpnName, topicEndpointTemplateName)
Delete a Topic Endpoint Template object.

Delete a Topic Endpoint Template object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Topic Endpoint Template provides a mechanism for specifying the initial state for client created topic endpoints.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointTemplateName** | **string**| The name of the Topic Endpoint Template. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnTopicEndpointTemplate**
> MsgVpnTopicEndpointTemplateResponse GetMsgVpnTopicEndpointTemplate(ctx, msgVpnName, topicEndpointTemplateName, optional)
Get a Topic Endpoint Template object.

Get a Topic Endpoint Template object.  A Topic Endpoint Template provides a mechanism for specifying the initial state for client created topic endpoints.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| topicEndpointTemplateName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointTemplateName** | **string**| The name of the Topic Endpoint Template. | 
 **optional** | ***TopicEndpointTemplateApiGetMsgVpnTopicEndpointTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TopicEndpointTemplateApiGetMsgVpnTopicEndpointTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointTemplateResponse**](MsgVpnTopicEndpointTemplateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnTopicEndpointTemplates**
> MsgVpnTopicEndpointTemplatesResponse GetMsgVpnTopicEndpointTemplates(ctx, msgVpnName, optional)
Get a list of Topic Endpoint Template objects.

Get a list of Topic Endpoint Template objects.  A Topic Endpoint Template provides a mechanism for specifying the initial state for client created topic endpoints.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| topicEndpointTemplateName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***TopicEndpointTemplateApiGetMsgVpnTopicEndpointTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TopicEndpointTemplateApiGetMsgVpnTopicEndpointTemplatesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointTemplatesResponse**](MsgVpnTopicEndpointTemplatesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnTopicEndpointTemplate**
> MsgVpnTopicEndpointTemplateResponse ReplaceMsgVpnTopicEndpointTemplate(ctx, body, msgVpnName, topicEndpointTemplateName, optional)
Replace a Topic Endpoint Template object.

Replace a Topic Endpoint Template object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  A Topic Endpoint Template provides a mechanism for specifying the initial state for client created topic endpoints.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- msgVpnName|x||x||||| rejectMsgToSenderOnDiscardBehavior||||||x|| topicEndpointTemplateName|x||x|||||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThreshold|clearPercent|setPercent|clearValue, setValue EventThreshold|clearValue|setValue|clearPercent, setPercent EventThreshold|setPercent|clearPercent|clearValue, setValue EventThreshold|setValue|clearValue|clearPercent, setPercent    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnTopicEndpointTemplate**](MsgVpnTopicEndpointTemplate.md)| The Topic Endpoint Template object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointTemplateName** | **string**| The name of the Topic Endpoint Template. | 
 **optional** | ***TopicEndpointTemplateApiReplaceMsgVpnTopicEndpointTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TopicEndpointTemplateApiReplaceMsgVpnTopicEndpointTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointTemplateResponse**](MsgVpnTopicEndpointTemplateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnTopicEndpointTemplate**
> MsgVpnTopicEndpointTemplateResponse UpdateMsgVpnTopicEndpointTemplate(ctx, body, msgVpnName, topicEndpointTemplateName, optional)
Update a Topic Endpoint Template object.

Update a Topic Endpoint Template object. Any attribute missing from the request will be left unchanged.  A Topic Endpoint Template provides a mechanism for specifying the initial state for client created topic endpoints.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- msgVpnName|x|x||||| rejectMsgToSenderOnDiscardBehavior|||||x|| topicEndpointTemplateName|x|x|||||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThreshold|clearPercent|setPercent|clearValue, setValue EventThreshold|clearValue|setValue|clearPercent, setPercent EventThreshold|setPercent|clearPercent|clearValue, setValue EventThreshold|setValue|clearValue|clearPercent, setPercent    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnTopicEndpointTemplate**](MsgVpnTopicEndpointTemplate.md)| The Topic Endpoint Template object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointTemplateName** | **string**| The name of the Topic Endpoint Template. | 
 **optional** | ***TopicEndpointTemplateApiUpdateMsgVpnTopicEndpointTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TopicEndpointTemplateApiUpdateMsgVpnTopicEndpointTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointTemplateResponse**](MsgVpnTopicEndpointTemplateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

