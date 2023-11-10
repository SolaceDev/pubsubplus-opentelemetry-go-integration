# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/config*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVirtualHostname**](VirtualHostnameApi.md#CreateVirtualHostname) | **Post** /virtualHostnames | Create a Virtual Hostname object.
[**DeleteVirtualHostname**](VirtualHostnameApi.md#DeleteVirtualHostname) | **Delete** /virtualHostnames/{virtualHostname} | Delete a Virtual Hostname object.
[**GetVirtualHostname**](VirtualHostnameApi.md#GetVirtualHostname) | **Get** /virtualHostnames/{virtualHostname} | Get a Virtual Hostname object.
[**GetVirtualHostnames**](VirtualHostnameApi.md#GetVirtualHostnames) | **Get** /virtualHostnames | Get a list of Virtual Hostname objects.
[**ReplaceVirtualHostname**](VirtualHostnameApi.md#ReplaceVirtualHostname) | **Put** /virtualHostnames/{virtualHostname} | Replace a Virtual Hostname object.
[**UpdateVirtualHostname**](VirtualHostnameApi.md#UpdateVirtualHostname) | **Patch** /virtualHostnames/{virtualHostname} | Update a Virtual Hostname object.

# **CreateVirtualHostname**
> VirtualHostnameResponse CreateVirtualHostname(ctx, body, optional)
Create a Virtual Hostname object.

Create a Virtual Hostname object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates via config-sync.  A Virtual Hostname is a provisioned object on a message broker that contains a Virtual Hostname to Message VPN mapping.  Clients which connect to a global (as opposed to per Message VPN) port and provides this hostname will be directed to its corresponding Message VPN. A case-insentive match is performed on the full client-provided hostname against the configured virtual-hostname.  This mechanism is only supported for hostnames provided through the Server Name Indication (SNI) extension of TLS.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: virtualHostname|x|x||||    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.17.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VirtualHostname**](VirtualHostname.md)| The Virtual Hostname object&#x27;s attributes. | 
 **optional** | ***VirtualHostnameApiCreateVirtualHostnameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VirtualHostnameApiCreateVirtualHostnameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**VirtualHostnameResponse**](VirtualHostnameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVirtualHostname**
> SempMetaOnlyResponse DeleteVirtualHostname(ctx, virtualHostname)
Delete a Virtual Hostname object.

Delete a Virtual Hostname object. The deletion of instances of this object are synchronized to HA mates via config-sync.  A Virtual Hostname is a provisioned object on a message broker that contains a Virtual Hostname to Message VPN mapping.  Clients which connect to a global (as opposed to per Message VPN) port and provides this hostname will be directed to its corresponding Message VPN. A case-insentive match is performed on the full client-provided hostname against the configured virtual-hostname.  This mechanism is only supported for hostnames provided through the Server Name Indication (SNI) extension of TLS.  A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.17.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **virtualHostname** | **string**| The virtual hostname. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVirtualHostname**
> VirtualHostnameResponse GetVirtualHostname(ctx, virtualHostname, optional)
Get a Virtual Hostname object.

Get a Virtual Hostname object.  A Virtual Hostname is a provisioned object on a message broker that contains a Virtual Hostname to Message VPN mapping.  Clients which connect to a global (as opposed to per Message VPN) port and provides this hostname will be directed to its corresponding Message VPN. A case-insentive match is performed on the full client-provided hostname against the configured virtual-hostname.  This mechanism is only supported for hostnames provided through the Server Name Indication (SNI) extension of TLS.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: virtualHostname|x|||    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.17.

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

 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Virtual Hostname objects.  A Virtual Hostname is a provisioned object on a message broker that contains a Virtual Hostname to Message VPN mapping.  Clients which connect to a global (as opposed to per Message VPN) port and provides this hostname will be directed to its corresponding Message VPN. A case-insentive match is performed on the full client-provided hostname against the configured virtual-hostname.  This mechanism is only supported for hostnames provided through the Server Name Indication (SNI) extension of TLS.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: virtualHostname|x|||    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.17.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

# **ReplaceVirtualHostname**
> VirtualHostnameResponse ReplaceVirtualHostname(ctx, body, virtualHostname, optional)
Replace a Virtual Hostname object.

Replace a Virtual Hostname object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  A Virtual Hostname is a provisioned object on a message broker that contains a Virtual Hostname to Message VPN mapping.  Clients which connect to a global (as opposed to per Message VPN) port and provides this hostname will be directed to its corresponding Message VPN. A case-insentive match is performed on the full client-provided hostname against the configured virtual-hostname.  This mechanism is only supported for hostnames provided through the Server Name Indication (SNI) extension of TLS.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- virtualHostname|x||x|||||    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.17.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VirtualHostname**](VirtualHostname.md)| The Virtual Hostname object&#x27;s attributes. | 
  **virtualHostname** | **string**| The virtual hostname. | 
 **optional** | ***VirtualHostnameApiReplaceVirtualHostnameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VirtualHostnameApiReplaceVirtualHostnameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**VirtualHostnameResponse**](VirtualHostnameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVirtualHostname**
> VirtualHostnameResponse UpdateVirtualHostname(ctx, body, virtualHostname, optional)
Update a Virtual Hostname object.

Update a Virtual Hostname object. Any attribute missing from the request will be left unchanged.  A Virtual Hostname is a provisioned object on a message broker that contains a Virtual Hostname to Message VPN mapping.  Clients which connect to a global (as opposed to per Message VPN) port and provides this hostname will be directed to its corresponding Message VPN. A case-insentive match is performed on the full client-provided hostname against the configured virtual-hostname.  This mechanism is only supported for hostnames provided through the Server Name Indication (SNI) extension of TLS.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- virtualHostname|x|x|||||    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.17.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VirtualHostname**](VirtualHostname.md)| The Virtual Hostname object&#x27;s attributes. | 
  **virtualHostname** | **string**| The virtual hostname. | 
 **optional** | ***VirtualHostnameApiUpdateVirtualHostnameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VirtualHostnameApiUpdateVirtualHostnameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**VirtualHostnameResponse**](VirtualHostnameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

