# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/config*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCertAuthority**](CertAuthorityApi.md#CreateCertAuthority) | **Post** /certAuthorities | Create a Certificate Authority object.
[**CreateCertAuthorityOcspTlsTrustedCommonName**](CertAuthorityApi.md#CreateCertAuthorityOcspTlsTrustedCommonName) | **Post** /certAuthorities/{certAuthorityName}/ocspTlsTrustedCommonNames | Create an OCSP Responder Trusted Common Name object.
[**DeleteCertAuthority**](CertAuthorityApi.md#DeleteCertAuthority) | **Delete** /certAuthorities/{certAuthorityName} | Delete a Certificate Authority object.
[**DeleteCertAuthorityOcspTlsTrustedCommonName**](CertAuthorityApi.md#DeleteCertAuthorityOcspTlsTrustedCommonName) | **Delete** /certAuthorities/{certAuthorityName}/ocspTlsTrustedCommonNames/{ocspTlsTrustedCommonName} | Delete an OCSP Responder Trusted Common Name object.
[**GetCertAuthorities**](CertAuthorityApi.md#GetCertAuthorities) | **Get** /certAuthorities | Get a list of Certificate Authority objects.
[**GetCertAuthority**](CertAuthorityApi.md#GetCertAuthority) | **Get** /certAuthorities/{certAuthorityName} | Get a Certificate Authority object.
[**GetCertAuthorityOcspTlsTrustedCommonName**](CertAuthorityApi.md#GetCertAuthorityOcspTlsTrustedCommonName) | **Get** /certAuthorities/{certAuthorityName}/ocspTlsTrustedCommonNames/{ocspTlsTrustedCommonName} | Get an OCSP Responder Trusted Common Name object.
[**GetCertAuthorityOcspTlsTrustedCommonNames**](CertAuthorityApi.md#GetCertAuthorityOcspTlsTrustedCommonNames) | **Get** /certAuthorities/{certAuthorityName}/ocspTlsTrustedCommonNames | Get a list of OCSP Responder Trusted Common Name objects.
[**ReplaceCertAuthority**](CertAuthorityApi.md#ReplaceCertAuthority) | **Put** /certAuthorities/{certAuthorityName} | Replace a Certificate Authority object.
[**UpdateCertAuthority**](CertAuthorityApi.md#UpdateCertAuthority) | **Patch** /certAuthorities/{certAuthorityName} | Update a Certificate Authority object.

# **CreateCertAuthority**
> CertAuthorityResponse CreateCertAuthority(ctx, body, optional)
Create a Certificate Authority object.

Create a Certificate Authority object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates via config-sync.  Clients can authenticate with the message broker over TLS by presenting a valid client certificate. The message broker authenticates the client certificate by constructing a full certificate chain (from the client certificate to intermediate CAs to a configured root CA). The intermediate CAs in this chain can be provided by the client, or configured in the message broker. The root CA must be configured on the message broker.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: certAuthorityName|x|x|||x| certContent|||||x| crlDayList|||||x| crlTimeList|||||x| crlUrl|||||x| ocspNonResponderCertEnabled|||||x| ocspOverrideUrl|||||x| ocspTimeout|||||x| revocationCheckEnabled|||||x|    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- CertAuthority|crlDayList|crlTimeList| CertAuthority|crlTimeList|crlDayList|    A SEMP client authorized with a minimum access scope/level of \"global/admin\" is required to perform this operation.  This has been deprecated since 2.19. Replaced by clientCertAuthorities and domainCertAuthorities.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CertAuthority**](CertAuthority.md)| The Certificate Authority object&#x27;s attributes. | 
 **optional** | ***CertAuthorityApiCreateCertAuthorityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CertAuthorityApiCreateCertAuthorityOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**CertAuthorityResponse**](CertAuthorityResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCertAuthorityOcspTlsTrustedCommonName**
> CertAuthorityOcspTlsTrustedCommonNameResponse CreateCertAuthorityOcspTlsTrustedCommonName(ctx, body, certAuthorityName, optional)
Create an OCSP Responder Trusted Common Name object.

Create an OCSP Responder Trusted Common Name object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates via config-sync.  When an OCSP override URL is configured, the OCSP responder will be required to sign the OCSP responses with certificates issued to these Trusted Common Names. A maximum of 8 common names can be configured as valid response signers.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: certAuthorityName|x||x||x| ocspTlsTrustedCommonName|x|x|||x|    A SEMP client authorized with a minimum access scope/level of \"global/admin\" is required to perform this operation.  This has been deprecated since 2.19. Replaced by clientCertAuthorities.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CertAuthorityOcspTlsTrustedCommonName**](CertAuthorityOcspTlsTrustedCommonName.md)| The OCSP Responder Trusted Common Name object&#x27;s attributes. | 
  **certAuthorityName** | **string**| The name of the Certificate Authority. | 
 **optional** | ***CertAuthorityApiCreateCertAuthorityOcspTlsTrustedCommonNameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CertAuthorityApiCreateCertAuthorityOcspTlsTrustedCommonNameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**CertAuthorityOcspTlsTrustedCommonNameResponse**](CertAuthorityOcspTlsTrustedCommonNameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCertAuthority**
> SempMetaOnlyResponse DeleteCertAuthority(ctx, certAuthorityName)
Delete a Certificate Authority object.

Delete a Certificate Authority object. The deletion of instances of this object are synchronized to HA mates via config-sync.  Clients can authenticate with the message broker over TLS by presenting a valid client certificate. The message broker authenticates the client certificate by constructing a full certificate chain (from the client certificate to intermediate CAs to a configured root CA). The intermediate CAs in this chain can be provided by the client, or configured in the message broker. The root CA must be configured on the message broker.  A SEMP client authorized with a minimum access scope/level of \"global/admin\" is required to perform this operation.  This has been deprecated since 2.19. Replaced by clientCertAuthorities and domainCertAuthorities.

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

# **DeleteCertAuthorityOcspTlsTrustedCommonName**
> SempMetaOnlyResponse DeleteCertAuthorityOcspTlsTrustedCommonName(ctx, certAuthorityName, ocspTlsTrustedCommonName)
Delete an OCSP Responder Trusted Common Name object.

Delete an OCSP Responder Trusted Common Name object. The deletion of instances of this object are synchronized to HA mates via config-sync.  When an OCSP override URL is configured, the OCSP responder will be required to sign the OCSP responses with certificates issued to these Trusted Common Names. A maximum of 8 common names can be configured as valid response signers.  A SEMP client authorized with a minimum access scope/level of \"global/admin\" is required to perform this operation.  This has been deprecated since 2.19. Replaced by clientCertAuthorities.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **certAuthorityName** | **string**| The name of the Certificate Authority. | 
  **ocspTlsTrustedCommonName** | **string**| The expected Trusted Common Name of the OCSP responder remote certificate. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCertAuthorities**
> CertAuthoritiesResponse GetCertAuthorities(ctx, optional)
Get a list of Certificate Authority objects.

Get a list of Certificate Authority objects.  Clients can authenticate with the message broker over TLS by presenting a valid client certificate. The message broker authenticates the client certificate by constructing a full certificate chain (from the client certificate to intermediate CAs to a configured root CA). The intermediate CAs in this chain can be provided by the client, or configured in the message broker. The root CA must be configured on the message broker.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: certAuthorityName|x||x| certContent|||x| crlDayList|||x| crlTimeList|||x| crlUrl|||x| ocspNonResponderCertEnabled|||x| ocspOverrideUrl|||x| ocspTimeout|||x| revocationCheckEnabled|||x|    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been deprecated since 2.19. Replaced by clientCertAuthorities and domainCertAuthorities.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a Certificate Authority object.  Clients can authenticate with the message broker over TLS by presenting a valid client certificate. The message broker authenticates the client certificate by constructing a full certificate chain (from the client certificate to intermediate CAs to a configured root CA). The intermediate CAs in this chain can be provided by the client, or configured in the message broker. The root CA must be configured on the message broker.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: certAuthorityName|x||x| certContent|||x| crlDayList|||x| crlTimeList|||x| crlUrl|||x| ocspNonResponderCertEnabled|||x| ocspOverrideUrl|||x| ocspTimeout|||x| revocationCheckEnabled|||x|    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been deprecated since 2.19. Replaced by clientCertAuthorities and domainCertAuthorities.

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

 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**CertAuthorityResponse**](CertAuthorityResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCertAuthorityOcspTlsTrustedCommonName**
> CertAuthorityOcspTlsTrustedCommonNameResponse GetCertAuthorityOcspTlsTrustedCommonName(ctx, certAuthorityName, ocspTlsTrustedCommonName, optional)
Get an OCSP Responder Trusted Common Name object.

Get an OCSP Responder Trusted Common Name object.  When an OCSP override URL is configured, the OCSP responder will be required to sign the OCSP responses with certificates issued to these Trusted Common Names. A maximum of 8 common names can be configured as valid response signers.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: certAuthorityName|x||x| ocspTlsTrustedCommonName|x||x|    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been deprecated since 2.19. Replaced by clientCertAuthorities.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **certAuthorityName** | **string**| The name of the Certificate Authority. | 
  **ocspTlsTrustedCommonName** | **string**| The expected Trusted Common Name of the OCSP responder remote certificate. | 
 **optional** | ***CertAuthorityApiGetCertAuthorityOcspTlsTrustedCommonNameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CertAuthorityApiGetCertAuthorityOcspTlsTrustedCommonNameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**CertAuthorityOcspTlsTrustedCommonNameResponse**](CertAuthorityOcspTlsTrustedCommonNameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCertAuthorityOcspTlsTrustedCommonNames**
> CertAuthorityOcspTlsTrustedCommonNamesResponse GetCertAuthorityOcspTlsTrustedCommonNames(ctx, certAuthorityName, optional)
Get a list of OCSP Responder Trusted Common Name objects.

Get a list of OCSP Responder Trusted Common Name objects.  When an OCSP override URL is configured, the OCSP responder will be required to sign the OCSP responses with certificates issued to these Trusted Common Names. A maximum of 8 common names can be configured as valid response signers.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: certAuthorityName|x||x| ocspTlsTrustedCommonName|x||x|    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been deprecated since 2.19. Replaced by clientCertAuthorities.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **certAuthorityName** | **string**| The name of the Certificate Authority. | 
 **optional** | ***CertAuthorityApiGetCertAuthorityOcspTlsTrustedCommonNamesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CertAuthorityApiGetCertAuthorityOcspTlsTrustedCommonNamesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**CertAuthorityOcspTlsTrustedCommonNamesResponse**](CertAuthorityOcspTlsTrustedCommonNamesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceCertAuthority**
> CertAuthorityResponse ReplaceCertAuthority(ctx, body, certAuthorityName, optional)
Replace a Certificate Authority object.

Replace a Certificate Authority object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  Clients can authenticate with the message broker over TLS by presenting a valid client certificate. The message broker authenticates the client certificate by constructing a full certificate chain (from the client certificate to intermediate CAs to a configured root CA). The intermediate CAs in this chain can be provided by the client, or configured in the message broker. The root CA must be configured on the message broker.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- certAuthorityName|x||x||||x| certContent|||||||x| crlDayList|||||||x| crlTimeList|||||||x| crlUrl||||||x|x| ocspNonResponderCertEnabled|||||||x| ocspOverrideUrl|||||||x| ocspTimeout|||||||x| revocationCheckEnabled|||||||x|    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- CertAuthority|crlDayList|crlTimeList| CertAuthority|crlTimeList|crlDayList|    A SEMP client authorized with a minimum access scope/level of \"global/admin\" is required to perform this operation.  This has been deprecated since 2.19. Replaced by clientCertAuthorities and domainCertAuthorities.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CertAuthority**](CertAuthority.md)| The Certificate Authority object&#x27;s attributes. | 
  **certAuthorityName** | **string**| The name of the Certificate Authority. | 
 **optional** | ***CertAuthorityApiReplaceCertAuthorityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CertAuthorityApiReplaceCertAuthorityOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**CertAuthorityResponse**](CertAuthorityResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCertAuthority**
> CertAuthorityResponse UpdateCertAuthority(ctx, body, certAuthorityName, optional)
Update a Certificate Authority object.

Update a Certificate Authority object. Any attribute missing from the request will be left unchanged.  Clients can authenticate with the message broker over TLS by presenting a valid client certificate. The message broker authenticates the client certificate by constructing a full certificate chain (from the client certificate to intermediate CAs to a configured root CA). The intermediate CAs in this chain can be provided by the client, or configured in the message broker. The root CA must be configured on the message broker.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- certAuthorityName|x|x||||x| certContent||||||x| crlDayList||||||x| crlTimeList||||||x| crlUrl|||||x|x| ocspNonResponderCertEnabled||||||x| ocspOverrideUrl||||||x| ocspTimeout||||||x| revocationCheckEnabled||||||x|    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- CertAuthority|crlDayList|crlTimeList| CertAuthority|crlTimeList|crlDayList|    A SEMP client authorized with a minimum access scope/level of \"global/admin\" is required to perform this operation.  This has been deprecated since 2.19. Replaced by clientCertAuthorities and domainCertAuthorities.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CertAuthority**](CertAuthority.md)| The Certificate Authority object&#x27;s attributes. | 
  **certAuthorityName** | **string**| The name of the Certificate Authority. | 
 **optional** | ***CertAuthorityApiUpdateCertAuthorityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CertAuthorityApiUpdateCertAuthorityOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**CertAuthorityResponse**](CertAuthorityResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

