# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/action*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DoCertAuthorityRefreshCrl**](CertAuthorityApi.md#DoCertAuthorityRefreshCrl) | **Put** /certAuthorities/{certAuthorityName}/refreshCrl | Refresh the CRL file for the Certificate Authority.
[**GetCertAuthorities**](CertAuthorityApi.md#GetCertAuthorities) | **Get** /certAuthorities | Get a list of Certificate Authority objects.
[**GetCertAuthority**](CertAuthorityApi.md#GetCertAuthority) | **Get** /certAuthorities/{certAuthorityName} | Get a Certificate Authority object.

# **DoCertAuthorityRefreshCrl**
> SempMetaOnlyResponse DoCertAuthorityRefreshCrl(ctx, body, certAuthorityName)
Refresh the CRL file for the Certificate Authority.

Refresh the CRL file for the Certificate Authority.    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been deprecated since 2.19. Replaced by clientCertAuthorities.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CertAuthorityRefreshCrl**](CertAuthorityRefreshCrl.md)| The Refresh CRL action&#x27;s attributes. | 
  **certAuthorityName** | **string**| The name of the Certificate Authority. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCertAuthorities**
> CertAuthoritiesResponse GetCertAuthorities(ctx, optional)
Get a list of Certificate Authority objects.

Get a list of Certificate Authority objects.  Clients can authenticate with the message broker over TLS by presenting a valid client certificate. The message broker authenticates the client certificate by constructing a full certificate chain (from the client certificate to intermediate CAs to a configured root CA). The intermediate CAs in this chain can be provided by the client, or configured in the message broker. The root CA must be configured on the message broker.   Attribute|Identifying|Deprecated :---|:---:|:---: certAuthorityName|x|x    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been deprecated since 2.19. Replaced by clientCertAuthorities and domainCertAuthorities.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CertAuthorityApiGetCertAuthoritiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CertAuthorityApiGetCertAuthoritiesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**CertAuthoritiesResponse**](CertAuthoritiesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCertAuthority**
> CertAuthorityResponse GetCertAuthority(ctx, certAuthorityName, optional)
Get a Certificate Authority object.

Get a Certificate Authority object.  Clients can authenticate with the message broker over TLS by presenting a valid client certificate. The message broker authenticates the client certificate by constructing a full certificate chain (from the client certificate to intermediate CAs to a configured root CA). The intermediate CAs in this chain can be provided by the client, or configured in the message broker. The root CA must be configured on the message broker.   Attribute|Identifying|Deprecated :---|:---:|:---: certAuthorityName|x|x    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been deprecated since 2.19. Replaced by clientCertAuthorities and domainCertAuthorities.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **certAuthorityName** | **string**| The name of the Certificate Authority. | 
 **optional** | ***CertAuthorityApiGetCertAuthorityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CertAuthorityApiGetCertAuthorityOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**CertAuthorityResponse**](CertAuthorityResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

