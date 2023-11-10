# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/config*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateClientCertAuthority**](ClientCertAuthorityApi.md#CreateClientCertAuthority) | **Post** /clientCertAuthorities | Create a Client Certificate Authority object.
[**CreateClientCertAuthorityOcspTlsTrustedCommonName**](ClientCertAuthorityApi.md#CreateClientCertAuthorityOcspTlsTrustedCommonName) | **Post** /clientCertAuthorities/{certAuthorityName}/ocspTlsTrustedCommonNames | Create an OCSP Responder Trusted Common Name object.
[**DeleteClientCertAuthority**](ClientCertAuthorityApi.md#DeleteClientCertAuthority) | **Delete** /clientCertAuthorities/{certAuthorityName} | Delete a Client Certificate Authority object.
[**DeleteClientCertAuthorityOcspTlsTrustedCommonName**](ClientCertAuthorityApi.md#DeleteClientCertAuthorityOcspTlsTrustedCommonName) | **Delete** /clientCertAuthorities/{certAuthorityName}/ocspTlsTrustedCommonNames/{ocspTlsTrustedCommonName} | Delete an OCSP Responder Trusted Common Name object.
[**GetClientCertAuthorities**](ClientCertAuthorityApi.md#GetClientCertAuthorities) | **Get** /clientCertAuthorities | Get a list of Client Certificate Authority objects.
[**GetClientCertAuthority**](ClientCertAuthorityApi.md#GetClientCertAuthority) | **Get** /clientCertAuthorities/{certAuthorityName} | Get a Client Certificate Authority object.
[**GetClientCertAuthorityOcspTlsTrustedCommonName**](ClientCertAuthorityApi.md#GetClientCertAuthorityOcspTlsTrustedCommonName) | **Get** /clientCertAuthorities/{certAuthorityName}/ocspTlsTrustedCommonNames/{ocspTlsTrustedCommonName} | Get an OCSP Responder Trusted Common Name object.
[**GetClientCertAuthorityOcspTlsTrustedCommonNames**](ClientCertAuthorityApi.md#GetClientCertAuthorityOcspTlsTrustedCommonNames) | **Get** /clientCertAuthorities/{certAuthorityName}/ocspTlsTrustedCommonNames | Get a list of OCSP Responder Trusted Common Name objects.
[**ReplaceClientCertAuthority**](ClientCertAuthorityApi.md#ReplaceClientCertAuthority) | **Put** /clientCertAuthorities/{certAuthorityName} | Replace a Client Certificate Authority object.
[**UpdateClientCertAuthority**](ClientCertAuthorityApi.md#UpdateClientCertAuthority) | **Patch** /clientCertAuthorities/{certAuthorityName} | Update a Client Certificate Authority object.

# **CreateClientCertAuthority**
> ClientCertAuthorityResponse CreateClientCertAuthority(ctx, body, optional)
Create a Client Certificate Authority object.

Create a Client Certificate Authority object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates via config-sync.  Clients can authenticate with the message broker over TLS by presenting a valid client certificate. The message broker authenticates the client certificate by constructing a full certificate chain (from the client certificate to intermediate CAs to a configured root CA). The intermediate CAs in this chain can be provided by the client, or configured in the message broker. The root CA must be configured on the message broker.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: certAuthorityName|x|x||||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- ClientCertAuthority|crlDayList|crlTimeList| ClientCertAuthority|crlTimeList|crlDayList|    A SEMP client authorized with a minimum access scope/level of \"global/admin\" is required to perform this operation.  This has been available since 2.19.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ClientCertAuthority**](ClientCertAuthority.md)| The Client Certificate Authority object&#x27;s attributes. | 
 **optional** | ***ClientCertAuthorityApiCreateClientCertAuthorityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientCertAuthorityApiCreateClientCertAuthorityOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**ClientCertAuthorityResponse**](ClientCertAuthorityResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateClientCertAuthorityOcspTlsTrustedCommonName**
> ClientCertAuthorityOcspTlsTrustedCommonNameResponse CreateClientCertAuthorityOcspTlsTrustedCommonName(ctx, body, certAuthorityName, optional)
Create an OCSP Responder Trusted Common Name object.

Create an OCSP Responder Trusted Common Name object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates via config-sync.  When an OCSP override URL is configured, the OCSP responder will be required to sign the OCSP responses with certificates issued to these Trusted Common Names. A maximum of 8 common names can be configured as valid response signers.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: certAuthorityName|x||x||| ocspTlsTrustedCommonName|x|x||||    A SEMP client authorized with a minimum access scope/level of \"global/admin\" is required to perform this operation.  This has been available since 2.19.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ClientCertAuthorityOcspTlsTrustedCommonName**](ClientCertAuthorityOcspTlsTrustedCommonName.md)| The OCSP Responder Trusted Common Name object&#x27;s attributes. | 
  **certAuthorityName** | **string**| The name of the Certificate Authority. | 
 **optional** | ***ClientCertAuthorityApiCreateClientCertAuthorityOcspTlsTrustedCommonNameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientCertAuthorityApiCreateClientCertAuthorityOcspTlsTrustedCommonNameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**ClientCertAuthorityOcspTlsTrustedCommonNameResponse**](ClientCertAuthorityOcspTlsTrustedCommonNameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteClientCertAuthority**
> SempMetaOnlyResponse DeleteClientCertAuthority(ctx, certAuthorityName)
Delete a Client Certificate Authority object.

Delete a Client Certificate Authority object. The deletion of instances of this object are synchronized to HA mates via config-sync.  Clients can authenticate with the message broker over TLS by presenting a valid client certificate. The message broker authenticates the client certificate by constructing a full certificate chain (from the client certificate to intermediate CAs to a configured root CA). The intermediate CAs in this chain can be provided by the client, or configured in the message broker. The root CA must be configured on the message broker.  A SEMP client authorized with a minimum access scope/level of \"global/admin\" is required to perform this operation.  This has been available since 2.19.

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

# **DeleteClientCertAuthorityOcspTlsTrustedCommonName**
> SempMetaOnlyResponse DeleteClientCertAuthorityOcspTlsTrustedCommonName(ctx, certAuthorityName, ocspTlsTrustedCommonName)
Delete an OCSP Responder Trusted Common Name object.

Delete an OCSP Responder Trusted Common Name object. The deletion of instances of this object are synchronized to HA mates via config-sync.  When an OCSP override URL is configured, the OCSP responder will be required to sign the OCSP responses with certificates issued to these Trusted Common Names. A maximum of 8 common names can be configured as valid response signers.  A SEMP client authorized with a minimum access scope/level of \"global/admin\" is required to perform this operation.  This has been available since 2.19.

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

# **GetClientCertAuthorities**
> ClientCertAuthoritiesResponse GetClientCertAuthorities(ctx, optional)
Get a list of Client Certificate Authority objects.

Get a list of Client Certificate Authority objects.  Clients can authenticate with the message broker over TLS by presenting a valid client certificate. The message broker authenticates the client certificate by constructing a full certificate chain (from the client certificate to intermediate CAs to a configured root CA). The intermediate CAs in this chain can be provided by the client, or configured in the message broker. The root CA must be configured on the message broker.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: certAuthorityName|x|||    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.19.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClientCertAuthorityApiGetClientCertAuthoritiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientCertAuthorityApiGetClientCertAuthoritiesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**ClientCertAuthoritiesResponse**](ClientCertAuthoritiesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClientCertAuthority**
> ClientCertAuthorityResponse GetClientCertAuthority(ctx, certAuthorityName, optional)
Get a Client Certificate Authority object.

Get a Client Certificate Authority object.  Clients can authenticate with the message broker over TLS by presenting a valid client certificate. The message broker authenticates the client certificate by constructing a full certificate chain (from the client certificate to intermediate CAs to a configured root CA). The intermediate CAs in this chain can be provided by the client, or configured in the message broker. The root CA must be configured on the message broker.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: certAuthorityName|x|||    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.19.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **certAuthorityName** | **string**| The name of the Certificate Authority. | 
 **optional** | ***ClientCertAuthorityApiGetClientCertAuthorityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientCertAuthorityApiGetClientCertAuthorityOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**ClientCertAuthorityResponse**](ClientCertAuthorityResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClientCertAuthorityOcspTlsTrustedCommonName**
> ClientCertAuthorityOcspTlsTrustedCommonNameResponse GetClientCertAuthorityOcspTlsTrustedCommonName(ctx, certAuthorityName, ocspTlsTrustedCommonName, optional)
Get an OCSP Responder Trusted Common Name object.

Get an OCSP Responder Trusted Common Name object.  When an OCSP override URL is configured, the OCSP responder will be required to sign the OCSP responses with certificates issued to these Trusted Common Names. A maximum of 8 common names can be configured as valid response signers.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: certAuthorityName|x||| ocspTlsTrustedCommonName|x|||    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.19.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **certAuthorityName** | **string**| The name of the Certificate Authority. | 
  **ocspTlsTrustedCommonName** | **string**| The expected Trusted Common Name of the OCSP responder remote certificate. | 
 **optional** | ***ClientCertAuthorityApiGetClientCertAuthorityOcspTlsTrustedCommonNameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientCertAuthorityApiGetClientCertAuthorityOcspTlsTrustedCommonNameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**ClientCertAuthorityOcspTlsTrustedCommonNameResponse**](ClientCertAuthorityOcspTlsTrustedCommonNameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClientCertAuthorityOcspTlsTrustedCommonNames**
> ClientCertAuthorityOcspTlsTrustedCommonNamesResponse GetClientCertAuthorityOcspTlsTrustedCommonNames(ctx, certAuthorityName, optional)
Get a list of OCSP Responder Trusted Common Name objects.

Get a list of OCSP Responder Trusted Common Name objects.  When an OCSP override URL is configured, the OCSP responder will be required to sign the OCSP responses with certificates issued to these Trusted Common Names. A maximum of 8 common names can be configured as valid response signers.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: certAuthorityName|x||| ocspTlsTrustedCommonName|x|||    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.19.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **certAuthorityName** | **string**| The name of the Certificate Authority. | 
 **optional** | ***ClientCertAuthorityApiGetClientCertAuthorityOcspTlsTrustedCommonNamesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientCertAuthorityApiGetClientCertAuthorityOcspTlsTrustedCommonNamesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**ClientCertAuthorityOcspTlsTrustedCommonNamesResponse**](ClientCertAuthorityOcspTlsTrustedCommonNamesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceClientCertAuthority**
> ClientCertAuthorityResponse ReplaceClientCertAuthority(ctx, body, certAuthorityName, optional)
Replace a Client Certificate Authority object.

Replace a Client Certificate Authority object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  Clients can authenticate with the message broker over TLS by presenting a valid client certificate. The message broker authenticates the client certificate by constructing a full certificate chain (from the client certificate to intermediate CAs to a configured root CA). The intermediate CAs in this chain can be provided by the client, or configured in the message broker. The root CA must be configured on the message broker.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- certAuthorityName|x||x||||| crlUrl||||||x||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- ClientCertAuthority|crlDayList|crlTimeList| ClientCertAuthority|crlTimeList|crlDayList|    A SEMP client authorized with a minimum access scope/level of \"global/admin\" is required to perform this operation.  This has been available since 2.19.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ClientCertAuthority**](ClientCertAuthority.md)| The Client Certificate Authority object&#x27;s attributes. | 
  **certAuthorityName** | **string**| The name of the Certificate Authority. | 
 **optional** | ***ClientCertAuthorityApiReplaceClientCertAuthorityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientCertAuthorityApiReplaceClientCertAuthorityOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**ClientCertAuthorityResponse**](ClientCertAuthorityResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateClientCertAuthority**
> ClientCertAuthorityResponse UpdateClientCertAuthority(ctx, body, certAuthorityName, optional)
Update a Client Certificate Authority object.

Update a Client Certificate Authority object. Any attribute missing from the request will be left unchanged.  Clients can authenticate with the message broker over TLS by presenting a valid client certificate. The message broker authenticates the client certificate by constructing a full certificate chain (from the client certificate to intermediate CAs to a configured root CA). The intermediate CAs in this chain can be provided by the client, or configured in the message broker. The root CA must be configured on the message broker.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- certAuthorityName|x|x||||| crlUrl|||||x||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- ClientCertAuthority|crlDayList|crlTimeList| ClientCertAuthority|crlTimeList|crlDayList|    A SEMP client authorized with a minimum access scope/level of \"global/admin\" is required to perform this operation.  This has been available since 2.19.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ClientCertAuthority**](ClientCertAuthority.md)| The Client Certificate Authority object&#x27;s attributes. | 
  **certAuthorityName** | **string**| The name of the Certificate Authority. | 
 **optional** | ***ClientCertAuthorityApiUpdateClientCertAuthorityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientCertAuthorityApiUpdateClientCertAuthorityOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**ClientCertAuthorityResponse**](ClientCertAuthorityResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

