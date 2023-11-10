# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/monitor*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetStandardDomainCertAuthorities**](StandardDomainCertAuthorityApi.md#GetStandardDomainCertAuthorities) | **Get** /standardDomainCertAuthorities | Get a list of Standard Domain Certificate Authority objects.
[**GetStandardDomainCertAuthority**](StandardDomainCertAuthorityApi.md#GetStandardDomainCertAuthority) | **Get** /standardDomainCertAuthorities/{certAuthorityName} | Get a Standard Domain Certificate Authority object.

# **GetStandardDomainCertAuthorities**
> StandardDomainCertAuthoritiesResponse GetStandardDomainCertAuthorities(ctx, optional)
Get a list of Standard Domain Certificate Authority objects.

Get a list of Standard Domain Certificate Authority objects.  Standard Certificate Authorities trusted for domain verification.   Attribute|Identifying|Deprecated :---|:---:|:---: certAuthorityName|x|    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.19.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StandardDomainCertAuthorityApiGetStandardDomainCertAuthoritiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StandardDomainCertAuthorityApiGetStandardDomainCertAuthoritiesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**StandardDomainCertAuthoritiesResponse**](StandardDomainCertAuthoritiesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStandardDomainCertAuthority**
> StandardDomainCertAuthorityResponse GetStandardDomainCertAuthority(ctx, certAuthorityName, optional)
Get a Standard Domain Certificate Authority object.

Get a Standard Domain Certificate Authority object.  Standard Certificate Authorities trusted for domain verification.   Attribute|Identifying|Deprecated :---|:---:|:---: certAuthorityName|x|    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.19.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **certAuthorityName** | **string**| The name of the Certificate Authority. | 
 **optional** | ***StandardDomainCertAuthorityApiGetStandardDomainCertAuthorityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StandardDomainCertAuthorityApiGetStandardDomainCertAuthorityOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**StandardDomainCertAuthorityResponse**](StandardDomainCertAuthorityResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

