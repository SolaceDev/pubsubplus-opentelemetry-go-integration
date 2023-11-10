# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/monitor*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVirtualHostname**](VirtualHostnameApi.md#GetVirtualHostname) | **Get** /virtualHostnames/{virtualHostname} | Get a Virtual Hostname object.
[**GetVirtualHostnames**](VirtualHostnameApi.md#GetVirtualHostnames) | **Get** /virtualHostnames | Get a list of Virtual Hostname objects.

# **GetVirtualHostname**
> VirtualHostnameResponse GetVirtualHostname(ctx, virtualHostname, optional)
Get a Virtual Hostname object.

Get a Virtual Hostname object.  A Virtual Hostname is a provisioned object on a message broker that contains a Virtual Hostname to Message VPN mapping.  Clients which connect to a global (as opposed to per Message VPN) port and provides this hostname will be directed to its corresponding Message VPN. A case-insentive match is performed on the full client-provided hostname against the configured virtual-hostname.  This mechanism is only supported for hostnames provided through the Server Name Indication (SNI) extension of TLS.   Attribute|Identifying|Deprecated :---|:---:|:---: virtualHostname|x|    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.17.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **virtualHostname** | **string**| The virtual hostname. | 
 **optional** | ***VirtualHostnameApiGetVirtualHostnameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VirtualHostnameApiGetVirtualHostnameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**VirtualHostnameResponse**](VirtualHostnameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVirtualHostnames**
> VirtualHostnamesResponse GetVirtualHostnames(ctx, optional)
Get a list of Virtual Hostname objects.

Get a list of Virtual Hostname objects.  A Virtual Hostname is a provisioned object on a message broker that contains a Virtual Hostname to Message VPN mapping.  Clients which connect to a global (as opposed to per Message VPN) port and provides this hostname will be directed to its corresponding Message VPN. A case-insentive match is performed on the full client-provided hostname against the configured virtual-hostname.  This mechanism is only supported for hostnames provided through the Server Name Indication (SNI) extension of TLS.   Attribute|Identifying|Deprecated :---|:---:|:---: virtualHostname|x|    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.17.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VirtualHostnameApiGetVirtualHostnamesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VirtualHostnameApiGetVirtualHostnamesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**VirtualHostnamesResponse**](VirtualHostnamesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

