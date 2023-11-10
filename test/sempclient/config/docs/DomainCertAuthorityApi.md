# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/config*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDomainCertAuthority**](DomainCertAuthorityApi.md#CreateDomainCertAuthority) | **Post** /domainCertAuthorities | Create a Domain Certificate Authority object.
[**DeleteDomainCertAuthority**](DomainCertAuthorityApi.md#DeleteDomainCertAuthority) | **Delete** /domainCertAuthorities/{certAuthorityName} | Delete a Domain Certificate Authority object.
[**GetDomainCertAuthorities**](DomainCertAuthorityApi.md#GetDomainCertAuthorities) | **Get** /domainCertAuthorities | Get a list of Domain Certificate Authority objects.
[**GetDomainCertAuthority**](DomainCertAuthorityApi.md#GetDomainCertAuthority) | **Get** /domainCertAuthorities/{certAuthorityName} | Get a Domain Certificate Authority object.
[**ReplaceDomainCertAuthority**](DomainCertAuthorityApi.md#ReplaceDomainCertAuthority) | **Put** /domainCertAuthorities/{certAuthorityName} | Replace a Domain Certificate Authority object.
[**UpdateDomainCertAuthority**](DomainCertAuthorityApi.md#UpdateDomainCertAuthority) | **Patch** /domainCertAuthorities/{certAuthorityName} | Update a Domain Certificate Authority object.

# **CreateDomainCertAuthority**
> DomainCertAuthorityResponse CreateDomainCertAuthority(ctx, body, optional)
Create a Domain Certificate Authority object.

Create a Domain Certificate Authority object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates via config-sync.  Certificate Authorities trusted for domain verification.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: certAuthorityName|x|x||||    A SEMP client authorized with a minimum access scope/level of \"global/admin\" is required to perform this operation.  This has been available since 2.19.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DomainCertAuthority**](DomainCertAuthority.md)| The Domain Certificate Authority object&#x27;s attributes. | 
 **optional** | ***DomainCertAuthorityApiCreateDomainCertAuthorityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomainCertAuthorityApiCreateDomainCertAuthorityOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DomainCertAuthorityResponse**](DomainCertAuthorityResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDomainCertAuthority**
> SempMetaOnlyResponse DeleteDomainCertAuthority(ctx, certAuthorityName)
Delete a Domain Certificate Authority object.

Delete a Domain Certificate Authority object. The deletion of instances of this object are synchronized to HA mates via config-sync.  Certificate Authorities trusted for domain verification.  A SEMP client authorized with a minimum access scope/level of \"global/admin\" is required to perform this operation.  This has been available since 2.19.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **certAuthorityName** | **string**| The name of the Certificate Authority. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDomainCertAuthorities**
> DomainCertAuthoritiesResponse GetDomainCertAuthorities(ctx, optional)
Get a list of Domain Certificate Authority objects.

Get a list of Domain Certificate Authority objects.  Certificate Authorities trusted for domain verification.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: certAuthorityName|x|||    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.19.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DomainCertAuthorityApiGetDomainCertAuthoritiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomainCertAuthorityApiGetDomainCertAuthoritiesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DomainCertAuthoritiesResponse**](DomainCertAuthoritiesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDomainCertAuthority**
> DomainCertAuthorityResponse GetDomainCertAuthority(ctx, certAuthorityName, optional)
Get a Domain Certificate Authority object.

Get a Domain Certificate Authority object.  Certificate Authorities trusted for domain verification.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: certAuthorityName|x|||    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.19.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **certAuthorityName** | **string**| The name of the Certificate Authority. | 
 **optional** | ***DomainCertAuthorityApiGetDomainCertAuthorityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomainCertAuthorityApiGetDomainCertAuthorityOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DomainCertAuthorityResponse**](DomainCertAuthorityResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceDomainCertAuthority**
> DomainCertAuthorityResponse ReplaceDomainCertAuthority(ctx, body, certAuthorityName, optional)
Replace a Domain Certificate Authority object.

Replace a Domain Certificate Authority object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  Certificate Authorities trusted for domain verification.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- certAuthorityName|x||x|||||    A SEMP client authorized with a minimum access scope/level of \"global/admin\" is required to perform this operation.  This has been available since 2.19.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DomainCertAuthority**](DomainCertAuthority.md)| The Domain Certificate Authority object&#x27;s attributes. | 
  **certAuthorityName** | **string**| The name of the Certificate Authority. | 
 **optional** | ***DomainCertAuthorityApiReplaceDomainCertAuthorityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomainCertAuthorityApiReplaceDomainCertAuthorityOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DomainCertAuthorityResponse**](DomainCertAuthorityResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDomainCertAuthority**
> DomainCertAuthorityResponse UpdateDomainCertAuthority(ctx, body, certAuthorityName, optional)
Update a Domain Certificate Authority object.

Update a Domain Certificate Authority object. Any attribute missing from the request will be left unchanged.  Certificate Authorities trusted for domain verification.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- certAuthorityName|x|x|||||    A SEMP client authorized with a minimum access scope/level of \"global/admin\" is required to perform this operation.  This has been available since 2.19.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DomainCertAuthority**](DomainCertAuthority.md)| The Domain Certificate Authority object&#x27;s attributes. | 
  **certAuthorityName** | **string**| The name of the Certificate Authority. | 
 **optional** | ***DomainCertAuthorityApiUpdateDomainCertAuthorityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomainCertAuthorityApiUpdateDomainCertAuthorityOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DomainCertAuthorityResponse**](DomainCertAuthorityResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

