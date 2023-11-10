# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/monitor*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMsgVpnClientUsername**](ClientUsernameApi.md#GetMsgVpnClientUsername) | **Get** /msgVpns/{msgVpnName}/clientUsernames/{clientUsername} | Get a Client Username object.
[**GetMsgVpnClientUsernameAttribute**](ClientUsernameApi.md#GetMsgVpnClientUsernameAttribute) | **Get** /msgVpns/{msgVpnName}/clientUsernames/{clientUsername}/attributes/{attributeName},{attributeValue} | Get a Client Username Attribute object.
[**GetMsgVpnClientUsernameAttributes**](ClientUsernameApi.md#GetMsgVpnClientUsernameAttributes) | **Get** /msgVpns/{msgVpnName}/clientUsernames/{clientUsername}/attributes | Get a list of Client Username Attribute objects.
[**GetMsgVpnClientUsernames**](ClientUsernameApi.md#GetMsgVpnClientUsernames) | **Get** /msgVpns/{msgVpnName}/clientUsernames | Get a list of Client Username objects.

# **GetMsgVpnClientUsername**
> MsgVpnClientUsernameResponse GetMsgVpnClientUsername(ctx, msgVpnName, clientUsername, optional)
Get a Client Username object.

Get a Client Username object.  A client is only authorized to connect to a Message VPN that is associated with a Client Username that the client has been assigned.   Attribute|Identifying|Deprecated :---|:---:|:---: clientUsername|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **clientUsername** | **string**| The name of the Client Username. | 
 **optional** | ***ClientUsernameApiGetMsgVpnClientUsernameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientUsernameApiGetMsgVpnClientUsernameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientUsernameResponse**](MsgVpnClientUsernameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnClientUsernameAttribute**
> MsgVpnClientUsernameAttributeResponse GetMsgVpnClientUsernameAttribute(ctx, msgVpnName, clientUsername, attributeName, attributeValue, optional)
Get a Client Username Attribute object.

Get a Client Username Attribute object.  A ClientUsername Attribute is a key+value pair that can be used to locate a client username, for example when using client certificate mapping.   Attribute|Identifying|Deprecated :---|:---:|:---: attributeName|x| attributeValue|x| clientUsername|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.27.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **clientUsername** | **string**| The name of the Client Username. | 
  **attributeName** | **string**| The name of the Attribute. | 
  **attributeValue** | **string**| The value of the Attribute. | 
 **optional** | ***ClientUsernameApiGetMsgVpnClientUsernameAttributeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientUsernameApiGetMsgVpnClientUsernameAttributeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientUsernameAttributeResponse**](MsgVpnClientUsernameAttributeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnClientUsernameAttributes**
> MsgVpnClientUsernameAttributesResponse GetMsgVpnClientUsernameAttributes(ctx, msgVpnName, clientUsername, optional)
Get a list of Client Username Attribute objects.

Get a list of Client Username Attribute objects.  A ClientUsername Attribute is a key+value pair that can be used to locate a client username, for example when using client certificate mapping.   Attribute|Identifying|Deprecated :---|:---:|:---: attributeName|x| attributeValue|x| clientUsername|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.27.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **clientUsername** | **string**| The name of the Client Username. | 
 **optional** | ***ClientUsernameApiGetMsgVpnClientUsernameAttributesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientUsernameApiGetMsgVpnClientUsernameAttributesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientUsernameAttributesResponse**](MsgVpnClientUsernameAttributesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnClientUsernames**
> MsgVpnClientUsernamesResponse GetMsgVpnClientUsernames(ctx, msgVpnName, optional)
Get a list of Client Username objects.

Get a list of Client Username objects.  A client is only authorized to connect to a Message VPN that is associated with a Client Username that the client has been assigned.   Attribute|Identifying|Deprecated :---|:---:|:---: clientUsername|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***ClientUsernameApiGetMsgVpnClientUsernamesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientUsernameApiGetMsgVpnClientUsernamesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientUsernamesResponse**](MsgVpnClientUsernamesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

