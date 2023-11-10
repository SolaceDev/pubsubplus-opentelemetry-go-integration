# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/config*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMsgVpnAuthenticationOauthProfile**](AuthenticationOauthProfileApi.md#CreateMsgVpnAuthenticationOauthProfile) | **Post** /msgVpns/{msgVpnName}/authenticationOauthProfiles | Create an OAuth Profile object.
[**CreateMsgVpnAuthenticationOauthProfileClientRequiredClaim**](AuthenticationOauthProfileApi.md#CreateMsgVpnAuthenticationOauthProfileClientRequiredClaim) | **Post** /msgVpns/{msgVpnName}/authenticationOauthProfiles/{oauthProfileName}/clientRequiredClaims | Create a Required Claim object.
[**CreateMsgVpnAuthenticationOauthProfileResourceServerRequiredClaim**](AuthenticationOauthProfileApi.md#CreateMsgVpnAuthenticationOauthProfileResourceServerRequiredClaim) | **Post** /msgVpns/{msgVpnName}/authenticationOauthProfiles/{oauthProfileName}/resourceServerRequiredClaims | Create a Required Claim object.
[**DeleteMsgVpnAuthenticationOauthProfile**](AuthenticationOauthProfileApi.md#DeleteMsgVpnAuthenticationOauthProfile) | **Delete** /msgVpns/{msgVpnName}/authenticationOauthProfiles/{oauthProfileName} | Delete an OAuth Profile object.
[**DeleteMsgVpnAuthenticationOauthProfileClientRequiredClaim**](AuthenticationOauthProfileApi.md#DeleteMsgVpnAuthenticationOauthProfileClientRequiredClaim) | **Delete** /msgVpns/{msgVpnName}/authenticationOauthProfiles/{oauthProfileName}/clientRequiredClaims/{clientRequiredClaimName} | Delete a Required Claim object.
[**DeleteMsgVpnAuthenticationOauthProfileResourceServerRequiredClaim**](AuthenticationOauthProfileApi.md#DeleteMsgVpnAuthenticationOauthProfileResourceServerRequiredClaim) | **Delete** /msgVpns/{msgVpnName}/authenticationOauthProfiles/{oauthProfileName}/resourceServerRequiredClaims/{resourceServerRequiredClaimName} | Delete a Required Claim object.
[**GetMsgVpnAuthenticationOauthProfile**](AuthenticationOauthProfileApi.md#GetMsgVpnAuthenticationOauthProfile) | **Get** /msgVpns/{msgVpnName}/authenticationOauthProfiles/{oauthProfileName} | Get an OAuth Profile object.
[**GetMsgVpnAuthenticationOauthProfileClientRequiredClaim**](AuthenticationOauthProfileApi.md#GetMsgVpnAuthenticationOauthProfileClientRequiredClaim) | **Get** /msgVpns/{msgVpnName}/authenticationOauthProfiles/{oauthProfileName}/clientRequiredClaims/{clientRequiredClaimName} | Get a Required Claim object.
[**GetMsgVpnAuthenticationOauthProfileClientRequiredClaims**](AuthenticationOauthProfileApi.md#GetMsgVpnAuthenticationOauthProfileClientRequiredClaims) | **Get** /msgVpns/{msgVpnName}/authenticationOauthProfiles/{oauthProfileName}/clientRequiredClaims | Get a list of Required Claim objects.
[**GetMsgVpnAuthenticationOauthProfileResourceServerRequiredClaim**](AuthenticationOauthProfileApi.md#GetMsgVpnAuthenticationOauthProfileResourceServerRequiredClaim) | **Get** /msgVpns/{msgVpnName}/authenticationOauthProfiles/{oauthProfileName}/resourceServerRequiredClaims/{resourceServerRequiredClaimName} | Get a Required Claim object.
[**GetMsgVpnAuthenticationOauthProfileResourceServerRequiredClaims**](AuthenticationOauthProfileApi.md#GetMsgVpnAuthenticationOauthProfileResourceServerRequiredClaims) | **Get** /msgVpns/{msgVpnName}/authenticationOauthProfiles/{oauthProfileName}/resourceServerRequiredClaims | Get a list of Required Claim objects.
[**GetMsgVpnAuthenticationOauthProfiles**](AuthenticationOauthProfileApi.md#GetMsgVpnAuthenticationOauthProfiles) | **Get** /msgVpns/{msgVpnName}/authenticationOauthProfiles | Get a list of OAuth Profile objects.
[**ReplaceMsgVpnAuthenticationOauthProfile**](AuthenticationOauthProfileApi.md#ReplaceMsgVpnAuthenticationOauthProfile) | **Put** /msgVpns/{msgVpnName}/authenticationOauthProfiles/{oauthProfileName} | Replace an OAuth Profile object.
[**UpdateMsgVpnAuthenticationOauthProfile**](AuthenticationOauthProfileApi.md#UpdateMsgVpnAuthenticationOauthProfile) | **Patch** /msgVpns/{msgVpnName}/authenticationOauthProfiles/{oauthProfileName} | Update an OAuth Profile object.

# **CreateMsgVpnAuthenticationOauthProfile**
> MsgVpnAuthenticationOauthProfileResponse CreateMsgVpnAuthenticationOauthProfile(ctx, body, msgVpnName, optional)
Create an OAuth Profile object.

Create an OAuth Profile object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  OAuth profiles specify how to securely authenticate to an OAuth provider.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: clientSecret||||x||x msgVpnName|x||x||| oauthProfileName|x|x||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.25.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnAuthenticationOauthProfile**](MsgVpnAuthenticationOauthProfile.md)| The OAuth Profile object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***AuthenticationOauthProfileApiCreateMsgVpnAuthenticationOauthProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationOauthProfileApiCreateMsgVpnAuthenticationOauthProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthenticationOauthProfileResponse**](MsgVpnAuthenticationOauthProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnAuthenticationOauthProfileClientRequiredClaim**
> MsgVpnAuthenticationOauthProfileClientRequiredClaimResponse CreateMsgVpnAuthenticationOauthProfileClientRequiredClaim(ctx, body, msgVpnName, oauthProfileName, optional)
Create a Required Claim object.

Create a Required Claim object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  Additional claims to be verified in the ID token.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: clientRequiredClaimName|x|x|||| clientRequiredClaimValue||x|||| msgVpnName|x||x||| oauthProfileName|x||x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.25.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnAuthenticationOauthProfileClientRequiredClaim**](MsgVpnAuthenticationOauthProfileClientRequiredClaim.md)| The Required Claim object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
 **optional** | ***AuthenticationOauthProfileApiCreateMsgVpnAuthenticationOauthProfileClientRequiredClaimOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationOauthProfileApiCreateMsgVpnAuthenticationOauthProfileClientRequiredClaimOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthenticationOauthProfileClientRequiredClaimResponse**](MsgVpnAuthenticationOauthProfileClientRequiredClaimResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnAuthenticationOauthProfileResourceServerRequiredClaim**
> MsgVpnAuthenticationOauthProfileResourceServerRequiredClaimResponse CreateMsgVpnAuthenticationOauthProfileResourceServerRequiredClaim(ctx, body, msgVpnName, oauthProfileName, optional)
Create a Required Claim object.

Create a Required Claim object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  Additional claims to be verified in the access token.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x||| oauthProfileName|x||x||| resourceServerRequiredClaimName|x|x|||| resourceServerRequiredClaimValue||x||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.25.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnAuthenticationOauthProfileResourceServerRequiredClaim**](MsgVpnAuthenticationOauthProfileResourceServerRequiredClaim.md)| The Required Claim object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
 **optional** | ***AuthenticationOauthProfileApiCreateMsgVpnAuthenticationOauthProfileResourceServerRequiredClaimOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationOauthProfileApiCreateMsgVpnAuthenticationOauthProfileResourceServerRequiredClaimOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthenticationOauthProfileResourceServerRequiredClaimResponse**](MsgVpnAuthenticationOauthProfileResourceServerRequiredClaimResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnAuthenticationOauthProfile**
> SempMetaOnlyResponse DeleteMsgVpnAuthenticationOauthProfile(ctx, msgVpnName, oauthProfileName)
Delete an OAuth Profile object.

Delete an OAuth Profile object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  OAuth profiles specify how to securely authenticate to an OAuth provider.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.25.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **oauthProfileName** | **string**| The name of the OAuth profile. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnAuthenticationOauthProfileClientRequiredClaim**
> SempMetaOnlyResponse DeleteMsgVpnAuthenticationOauthProfileClientRequiredClaim(ctx, msgVpnName, oauthProfileName, clientRequiredClaimName)
Delete a Required Claim object.

Delete a Required Claim object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  Additional claims to be verified in the ID token.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.25.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
  **clientRequiredClaimName** | **string**| The name of the ID token claim to verify. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnAuthenticationOauthProfileResourceServerRequiredClaim**
> SempMetaOnlyResponse DeleteMsgVpnAuthenticationOauthProfileResourceServerRequiredClaim(ctx, msgVpnName, oauthProfileName, resourceServerRequiredClaimName)
Delete a Required Claim object.

Delete a Required Claim object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  Additional claims to be verified in the access token.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.25.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
  **resourceServerRequiredClaimName** | **string**| The name of the access token claim to verify. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAuthenticationOauthProfile**
> MsgVpnAuthenticationOauthProfileResponse GetMsgVpnAuthenticationOauthProfile(ctx, msgVpnName, oauthProfileName, optional)
Get an OAuth Profile object.

Get an OAuth Profile object.  OAuth profiles specify how to securely authenticate to an OAuth provider.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: clientSecret||x||x msgVpnName|x||| oauthProfileName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.25.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
 **optional** | ***AuthenticationOauthProfileApiGetMsgVpnAuthenticationOauthProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationOauthProfileApiGetMsgVpnAuthenticationOauthProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthenticationOauthProfileResponse**](MsgVpnAuthenticationOauthProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAuthenticationOauthProfileClientRequiredClaim**
> MsgVpnAuthenticationOauthProfileClientRequiredClaimResponse GetMsgVpnAuthenticationOauthProfileClientRequiredClaim(ctx, msgVpnName, oauthProfileName, clientRequiredClaimName, optional)
Get a Required Claim object.

Get a Required Claim object.  Additional claims to be verified in the ID token.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: clientRequiredClaimName|x||| msgVpnName|x||| oauthProfileName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.25.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
  **clientRequiredClaimName** | **string**| The name of the ID token claim to verify. | 
 **optional** | ***AuthenticationOauthProfileApiGetMsgVpnAuthenticationOauthProfileClientRequiredClaimOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationOauthProfileApiGetMsgVpnAuthenticationOauthProfileClientRequiredClaimOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthenticationOauthProfileClientRequiredClaimResponse**](MsgVpnAuthenticationOauthProfileClientRequiredClaimResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAuthenticationOauthProfileClientRequiredClaims**
> MsgVpnAuthenticationOauthProfileClientRequiredClaimsResponse GetMsgVpnAuthenticationOauthProfileClientRequiredClaims(ctx, msgVpnName, oauthProfileName, optional)
Get a list of Required Claim objects.

Get a list of Required Claim objects.  Additional claims to be verified in the ID token.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: clientRequiredClaimName|x||| msgVpnName|x||| oauthProfileName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.25.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
 **optional** | ***AuthenticationOauthProfileApiGetMsgVpnAuthenticationOauthProfileClientRequiredClaimsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationOauthProfileApiGetMsgVpnAuthenticationOauthProfileClientRequiredClaimsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthenticationOauthProfileClientRequiredClaimsResponse**](MsgVpnAuthenticationOauthProfileClientRequiredClaimsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAuthenticationOauthProfileResourceServerRequiredClaim**
> MsgVpnAuthenticationOauthProfileResourceServerRequiredClaimResponse GetMsgVpnAuthenticationOauthProfileResourceServerRequiredClaim(ctx, msgVpnName, oauthProfileName, resourceServerRequiredClaimName, optional)
Get a Required Claim object.

Get a Required Claim object.  Additional claims to be verified in the access token.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| oauthProfileName|x||| resourceServerRequiredClaimName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.25.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
  **resourceServerRequiredClaimName** | **string**| The name of the access token claim to verify. | 
 **optional** | ***AuthenticationOauthProfileApiGetMsgVpnAuthenticationOauthProfileResourceServerRequiredClaimOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationOauthProfileApiGetMsgVpnAuthenticationOauthProfileResourceServerRequiredClaimOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthenticationOauthProfileResourceServerRequiredClaimResponse**](MsgVpnAuthenticationOauthProfileResourceServerRequiredClaimResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAuthenticationOauthProfileResourceServerRequiredClaims**
> MsgVpnAuthenticationOauthProfileResourceServerRequiredClaimsResponse GetMsgVpnAuthenticationOauthProfileResourceServerRequiredClaims(ctx, msgVpnName, oauthProfileName, optional)
Get a list of Required Claim objects.

Get a list of Required Claim objects.  Additional claims to be verified in the access token.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| oauthProfileName|x||| resourceServerRequiredClaimName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.25.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
 **optional** | ***AuthenticationOauthProfileApiGetMsgVpnAuthenticationOauthProfileResourceServerRequiredClaimsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationOauthProfileApiGetMsgVpnAuthenticationOauthProfileResourceServerRequiredClaimsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthenticationOauthProfileResourceServerRequiredClaimsResponse**](MsgVpnAuthenticationOauthProfileResourceServerRequiredClaimsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAuthenticationOauthProfiles**
> MsgVpnAuthenticationOauthProfilesResponse GetMsgVpnAuthenticationOauthProfiles(ctx, msgVpnName, optional)
Get a list of OAuth Profile objects.

Get a list of OAuth Profile objects.  OAuth profiles specify how to securely authenticate to an OAuth provider.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: clientSecret||x||x msgVpnName|x||| oauthProfileName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.25.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***AuthenticationOauthProfileApiGetMsgVpnAuthenticationOauthProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationOauthProfileApiGetMsgVpnAuthenticationOauthProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthenticationOauthProfilesResponse**](MsgVpnAuthenticationOauthProfilesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnAuthenticationOauthProfile**
> MsgVpnAuthenticationOauthProfileResponse ReplaceMsgVpnAuthenticationOauthProfile(ctx, body, msgVpnName, oauthProfileName, optional)
Replace an OAuth Profile object.

Replace an OAuth Profile object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  OAuth profiles specify how to securely authenticate to an OAuth provider.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- clientSecret||||x||||x msgVpnName|x||x||||| oauthProfileName|x||x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.25.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnAuthenticationOauthProfile**](MsgVpnAuthenticationOauthProfile.md)| The OAuth Profile object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
 **optional** | ***AuthenticationOauthProfileApiReplaceMsgVpnAuthenticationOauthProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationOauthProfileApiReplaceMsgVpnAuthenticationOauthProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthenticationOauthProfileResponse**](MsgVpnAuthenticationOauthProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnAuthenticationOauthProfile**
> MsgVpnAuthenticationOauthProfileResponse UpdateMsgVpnAuthenticationOauthProfile(ctx, body, msgVpnName, oauthProfileName, optional)
Update an OAuth Profile object.

Update an OAuth Profile object. Any attribute missing from the request will be left unchanged.  OAuth profiles specify how to securely authenticate to an OAuth provider.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- clientSecret|||x||||x msgVpnName|x|x||||| oauthProfileName|x|x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.25.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnAuthenticationOauthProfile**](MsgVpnAuthenticationOauthProfile.md)| The OAuth Profile object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
 **optional** | ***AuthenticationOauthProfileApiUpdateMsgVpnAuthenticationOauthProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationOauthProfileApiUpdateMsgVpnAuthenticationOauthProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthenticationOauthProfileResponse**](MsgVpnAuthenticationOauthProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

