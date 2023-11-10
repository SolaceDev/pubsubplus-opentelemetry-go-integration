# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/config*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOauthProfile**](OauthProfileApi.md#CreateOauthProfile) | **Post** /oauthProfiles | Create an OAuth Profile object.
[**CreateOauthProfileAccessLevelGroup**](OauthProfileApi.md#CreateOauthProfileAccessLevelGroup) | **Post** /oauthProfiles/{oauthProfileName}/accessLevelGroups | Create a Group Access Level object.
[**CreateOauthProfileAccessLevelGroupMsgVpnAccessLevelException**](OauthProfileApi.md#CreateOauthProfileAccessLevelGroupMsgVpnAccessLevelException) | **Post** /oauthProfiles/{oauthProfileName}/accessLevelGroups/{groupName}/msgVpnAccessLevelExceptions | Create a Message VPN Access-Level Exception object.
[**CreateOauthProfileClientAllowedHost**](OauthProfileApi.md#CreateOauthProfileClientAllowedHost) | **Post** /oauthProfiles/{oauthProfileName}/clientAllowedHosts | Create an Allowed Host Value object.
[**CreateOauthProfileClientAuthorizationParameter**](OauthProfileApi.md#CreateOauthProfileClientAuthorizationParameter) | **Post** /oauthProfiles/{oauthProfileName}/clientAuthorizationParameters | Create an Authorization Parameter object.
[**CreateOauthProfileClientRequiredClaim**](OauthProfileApi.md#CreateOauthProfileClientRequiredClaim) | **Post** /oauthProfiles/{oauthProfileName}/clientRequiredClaims | Create a Required Claim object.
[**CreateOauthProfileDefaultMsgVpnAccessLevelException**](OauthProfileApi.md#CreateOauthProfileDefaultMsgVpnAccessLevelException) | **Post** /oauthProfiles/{oauthProfileName}/defaultMsgVpnAccessLevelExceptions | Create a Message VPN Access-Level Exception object.
[**CreateOauthProfileResourceServerRequiredClaim**](OauthProfileApi.md#CreateOauthProfileResourceServerRequiredClaim) | **Post** /oauthProfiles/{oauthProfileName}/resourceServerRequiredClaims | Create a Required Claim object.
[**DeleteOauthProfile**](OauthProfileApi.md#DeleteOauthProfile) | **Delete** /oauthProfiles/{oauthProfileName} | Delete an OAuth Profile object.
[**DeleteOauthProfileAccessLevelGroup**](OauthProfileApi.md#DeleteOauthProfileAccessLevelGroup) | **Delete** /oauthProfiles/{oauthProfileName}/accessLevelGroups/{groupName} | Delete a Group Access Level object.
[**DeleteOauthProfileAccessLevelGroupMsgVpnAccessLevelException**](OauthProfileApi.md#DeleteOauthProfileAccessLevelGroupMsgVpnAccessLevelException) | **Delete** /oauthProfiles/{oauthProfileName}/accessLevelGroups/{groupName}/msgVpnAccessLevelExceptions/{msgVpnName} | Delete a Message VPN Access-Level Exception object.
[**DeleteOauthProfileClientAllowedHost**](OauthProfileApi.md#DeleteOauthProfileClientAllowedHost) | **Delete** /oauthProfiles/{oauthProfileName}/clientAllowedHosts/{allowedHost} | Delete an Allowed Host Value object.
[**DeleteOauthProfileClientAuthorizationParameter**](OauthProfileApi.md#DeleteOauthProfileClientAuthorizationParameter) | **Delete** /oauthProfiles/{oauthProfileName}/clientAuthorizationParameters/{authorizationParameterName} | Delete an Authorization Parameter object.
[**DeleteOauthProfileClientRequiredClaim**](OauthProfileApi.md#DeleteOauthProfileClientRequiredClaim) | **Delete** /oauthProfiles/{oauthProfileName}/clientRequiredClaims/{clientRequiredClaimName} | Delete a Required Claim object.
[**DeleteOauthProfileDefaultMsgVpnAccessLevelException**](OauthProfileApi.md#DeleteOauthProfileDefaultMsgVpnAccessLevelException) | **Delete** /oauthProfiles/{oauthProfileName}/defaultMsgVpnAccessLevelExceptions/{msgVpnName} | Delete a Message VPN Access-Level Exception object.
[**DeleteOauthProfileResourceServerRequiredClaim**](OauthProfileApi.md#DeleteOauthProfileResourceServerRequiredClaim) | **Delete** /oauthProfiles/{oauthProfileName}/resourceServerRequiredClaims/{resourceServerRequiredClaimName} | Delete a Required Claim object.
[**GetOauthProfile**](OauthProfileApi.md#GetOauthProfile) | **Get** /oauthProfiles/{oauthProfileName} | Get an OAuth Profile object.
[**GetOauthProfileAccessLevelGroup**](OauthProfileApi.md#GetOauthProfileAccessLevelGroup) | **Get** /oauthProfiles/{oauthProfileName}/accessLevelGroups/{groupName} | Get a Group Access Level object.
[**GetOauthProfileAccessLevelGroupMsgVpnAccessLevelException**](OauthProfileApi.md#GetOauthProfileAccessLevelGroupMsgVpnAccessLevelException) | **Get** /oauthProfiles/{oauthProfileName}/accessLevelGroups/{groupName}/msgVpnAccessLevelExceptions/{msgVpnName} | Get a Message VPN Access-Level Exception object.
[**GetOauthProfileAccessLevelGroupMsgVpnAccessLevelExceptions**](OauthProfileApi.md#GetOauthProfileAccessLevelGroupMsgVpnAccessLevelExceptions) | **Get** /oauthProfiles/{oauthProfileName}/accessLevelGroups/{groupName}/msgVpnAccessLevelExceptions | Get a list of Message VPN Access-Level Exception objects.
[**GetOauthProfileAccessLevelGroups**](OauthProfileApi.md#GetOauthProfileAccessLevelGroups) | **Get** /oauthProfiles/{oauthProfileName}/accessLevelGroups | Get a list of Group Access Level objects.
[**GetOauthProfileClientAllowedHost**](OauthProfileApi.md#GetOauthProfileClientAllowedHost) | **Get** /oauthProfiles/{oauthProfileName}/clientAllowedHosts/{allowedHost} | Get an Allowed Host Value object.
[**GetOauthProfileClientAllowedHosts**](OauthProfileApi.md#GetOauthProfileClientAllowedHosts) | **Get** /oauthProfiles/{oauthProfileName}/clientAllowedHosts | Get a list of Allowed Host Value objects.
[**GetOauthProfileClientAuthorizationParameter**](OauthProfileApi.md#GetOauthProfileClientAuthorizationParameter) | **Get** /oauthProfiles/{oauthProfileName}/clientAuthorizationParameters/{authorizationParameterName} | Get an Authorization Parameter object.
[**GetOauthProfileClientAuthorizationParameters**](OauthProfileApi.md#GetOauthProfileClientAuthorizationParameters) | **Get** /oauthProfiles/{oauthProfileName}/clientAuthorizationParameters | Get a list of Authorization Parameter objects.
[**GetOauthProfileClientRequiredClaim**](OauthProfileApi.md#GetOauthProfileClientRequiredClaim) | **Get** /oauthProfiles/{oauthProfileName}/clientRequiredClaims/{clientRequiredClaimName} | Get a Required Claim object.
[**GetOauthProfileClientRequiredClaims**](OauthProfileApi.md#GetOauthProfileClientRequiredClaims) | **Get** /oauthProfiles/{oauthProfileName}/clientRequiredClaims | Get a list of Required Claim objects.
[**GetOauthProfileDefaultMsgVpnAccessLevelException**](OauthProfileApi.md#GetOauthProfileDefaultMsgVpnAccessLevelException) | **Get** /oauthProfiles/{oauthProfileName}/defaultMsgVpnAccessLevelExceptions/{msgVpnName} | Get a Message VPN Access-Level Exception object.
[**GetOauthProfileDefaultMsgVpnAccessLevelExceptions**](OauthProfileApi.md#GetOauthProfileDefaultMsgVpnAccessLevelExceptions) | **Get** /oauthProfiles/{oauthProfileName}/defaultMsgVpnAccessLevelExceptions | Get a list of Message VPN Access-Level Exception objects.
[**GetOauthProfileResourceServerRequiredClaim**](OauthProfileApi.md#GetOauthProfileResourceServerRequiredClaim) | **Get** /oauthProfiles/{oauthProfileName}/resourceServerRequiredClaims/{resourceServerRequiredClaimName} | Get a Required Claim object.
[**GetOauthProfileResourceServerRequiredClaims**](OauthProfileApi.md#GetOauthProfileResourceServerRequiredClaims) | **Get** /oauthProfiles/{oauthProfileName}/resourceServerRequiredClaims | Get a list of Required Claim objects.
[**GetOauthProfiles**](OauthProfileApi.md#GetOauthProfiles) | **Get** /oauthProfiles | Get a list of OAuth Profile objects.
[**ReplaceOauthProfile**](OauthProfileApi.md#ReplaceOauthProfile) | **Put** /oauthProfiles/{oauthProfileName} | Replace an OAuth Profile object.
[**ReplaceOauthProfileAccessLevelGroup**](OauthProfileApi.md#ReplaceOauthProfileAccessLevelGroup) | **Put** /oauthProfiles/{oauthProfileName}/accessLevelGroups/{groupName} | Replace a Group Access Level object.
[**ReplaceOauthProfileAccessLevelGroupMsgVpnAccessLevelException**](OauthProfileApi.md#ReplaceOauthProfileAccessLevelGroupMsgVpnAccessLevelException) | **Put** /oauthProfiles/{oauthProfileName}/accessLevelGroups/{groupName}/msgVpnAccessLevelExceptions/{msgVpnName} | Replace a Message VPN Access-Level Exception object.
[**ReplaceOauthProfileClientAuthorizationParameter**](OauthProfileApi.md#ReplaceOauthProfileClientAuthorizationParameter) | **Put** /oauthProfiles/{oauthProfileName}/clientAuthorizationParameters/{authorizationParameterName} | Replace an Authorization Parameter object.
[**ReplaceOauthProfileDefaultMsgVpnAccessLevelException**](OauthProfileApi.md#ReplaceOauthProfileDefaultMsgVpnAccessLevelException) | **Put** /oauthProfiles/{oauthProfileName}/defaultMsgVpnAccessLevelExceptions/{msgVpnName} | Replace a Message VPN Access-Level Exception object.
[**UpdateOauthProfile**](OauthProfileApi.md#UpdateOauthProfile) | **Patch** /oauthProfiles/{oauthProfileName} | Update an OAuth Profile object.
[**UpdateOauthProfileAccessLevelGroup**](OauthProfileApi.md#UpdateOauthProfileAccessLevelGroup) | **Patch** /oauthProfiles/{oauthProfileName}/accessLevelGroups/{groupName} | Update a Group Access Level object.
[**UpdateOauthProfileAccessLevelGroupMsgVpnAccessLevelException**](OauthProfileApi.md#UpdateOauthProfileAccessLevelGroupMsgVpnAccessLevelException) | **Patch** /oauthProfiles/{oauthProfileName}/accessLevelGroups/{groupName}/msgVpnAccessLevelExceptions/{msgVpnName} | Update a Message VPN Access-Level Exception object.
[**UpdateOauthProfileClientAuthorizationParameter**](OauthProfileApi.md#UpdateOauthProfileClientAuthorizationParameter) | **Patch** /oauthProfiles/{oauthProfileName}/clientAuthorizationParameters/{authorizationParameterName} | Update an Authorization Parameter object.
[**UpdateOauthProfileDefaultMsgVpnAccessLevelException**](OauthProfileApi.md#UpdateOauthProfileDefaultMsgVpnAccessLevelException) | **Patch** /oauthProfiles/{oauthProfileName}/defaultMsgVpnAccessLevelExceptions/{msgVpnName} | Update a Message VPN Access-Level Exception object.

# **CreateOauthProfile**
> OauthProfileResponse CreateOauthProfile(ctx, body, optional)
Create an OAuth Profile object.

Create an OAuth Profile object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates via config-sync.  OAuth profiles specify how to securely authenticate to an OAuth provider.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: clientSecret||||x||x oauthProfileName|x|x||||    A SEMP client authorized with a minimum access scope/level of \"global/admin\" is required to perform this operation.  This has been available since 2.24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OauthProfile**](OauthProfile.md)| The OAuth Profile object&#x27;s attributes. | 
 **optional** | ***OauthProfileApiCreateOauthProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthProfileApiCreateOauthProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**OauthProfileResponse**](OauthProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOauthProfileAccessLevelGroup**
> OauthProfileAccessLevelGroupResponse CreateOauthProfileAccessLevelGroup(ctx, body, oauthProfileName, optional)
Create a Group Access Level object.

Create a Group Access Level object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates via config-sync.  The name of a group as it exists on the OAuth server being used to authenticate SEMP users.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: groupName|x|x|||| oauthProfileName|x||x|||    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation. Requests which include the following attributes require greater access scope/level:   Attribute|Access Scope/Level :---|:---: globalAccessLevel|global/admin    This has been available since 2.24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OauthProfileAccessLevelGroup**](OauthProfileAccessLevelGroup.md)| The Group Access Level object&#x27;s attributes. | 
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
 **optional** | ***OauthProfileApiCreateOauthProfileAccessLevelGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthProfileApiCreateOauthProfileAccessLevelGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**OauthProfileAccessLevelGroupResponse**](OauthProfileAccessLevelGroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOauthProfileAccessLevelGroupMsgVpnAccessLevelException**
> OauthProfileAccessLevelGroupMsgVpnAccessLevelExceptionResponse CreateOauthProfileAccessLevelGroupMsgVpnAccessLevelException(ctx, body, oauthProfileName, groupName, optional)
Create a Message VPN Access-Level Exception object.

Create a Message VPN Access-Level Exception object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates via config-sync.  Message VPN access-level exceptions for members of this group.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: groupName|x||x||| msgVpnName|x|x|||| oauthProfileName|x||x|||    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OauthProfileAccessLevelGroupMsgVpnAccessLevelException**](OauthProfileAccessLevelGroupMsgVpnAccessLevelException.md)| The Message VPN Access-Level Exception object&#x27;s attributes. | 
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
  **groupName** | **string**| The name of the group. | 
 **optional** | ***OauthProfileApiCreateOauthProfileAccessLevelGroupMsgVpnAccessLevelExceptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthProfileApiCreateOauthProfileAccessLevelGroupMsgVpnAccessLevelExceptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**OauthProfileAccessLevelGroupMsgVpnAccessLevelExceptionResponse**](OauthProfileAccessLevelGroupMsgVpnAccessLevelExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOauthProfileClientAllowedHost**
> OauthProfileClientAllowedHostResponse CreateOauthProfileClientAllowedHost(ctx, body, oauthProfileName, optional)
Create an Allowed Host Value object.

Create an Allowed Host Value object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates via config-sync.  A valid hostname for this broker in OAuth redirects.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: allowedHost|x|x|||| oauthProfileName|x||x|||    A SEMP client authorized with a minimum access scope/level of \"global/admin\" is required to perform this operation.  This has been available since 2.24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OauthProfileClientAllowedHost**](OauthProfileClientAllowedHost.md)| The Allowed Host Value object&#x27;s attributes. | 
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
 **optional** | ***OauthProfileApiCreateOauthProfileClientAllowedHostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthProfileApiCreateOauthProfileClientAllowedHostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**OauthProfileClientAllowedHostResponse**](OauthProfileClientAllowedHostResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOauthProfileClientAuthorizationParameter**
> OauthProfileClientAuthorizationParameterResponse CreateOauthProfileClientAuthorizationParameter(ctx, body, oauthProfileName, optional)
Create an Authorization Parameter object.

Create an Authorization Parameter object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates via config-sync.  Additional parameters to be passed to the OAuth authorization endpoint.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: authorizationParameterName|x|x|||| oauthProfileName|x||x|||    A SEMP client authorized with a minimum access scope/level of \"global/admin\" is required to perform this operation.  This has been available since 2.24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OauthProfileClientAuthorizationParameter**](OauthProfileClientAuthorizationParameter.md)| The Authorization Parameter object&#x27;s attributes. | 
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
 **optional** | ***OauthProfileApiCreateOauthProfileClientAuthorizationParameterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthProfileApiCreateOauthProfileClientAuthorizationParameterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**OauthProfileClientAuthorizationParameterResponse**](OauthProfileClientAuthorizationParameterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOauthProfileClientRequiredClaim**
> OauthProfileClientRequiredClaimResponse CreateOauthProfileClientRequiredClaim(ctx, body, oauthProfileName, optional)
Create a Required Claim object.

Create a Required Claim object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates via config-sync.  Additional claims to be verified in the ID token.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: clientRequiredClaimName|x|x|||| clientRequiredClaimValue||x|||| oauthProfileName|x||x|||    A SEMP client authorized with a minimum access scope/level of \"global/admin\" is required to perform this operation.  This has been available since 2.24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OauthProfileClientRequiredClaim**](OauthProfileClientRequiredClaim.md)| The Required Claim object&#x27;s attributes. | 
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
 **optional** | ***OauthProfileApiCreateOauthProfileClientRequiredClaimOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthProfileApiCreateOauthProfileClientRequiredClaimOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**OauthProfileClientRequiredClaimResponse**](OauthProfileClientRequiredClaimResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOauthProfileDefaultMsgVpnAccessLevelException**
> OauthProfileDefaultMsgVpnAccessLevelExceptionResponse CreateOauthProfileDefaultMsgVpnAccessLevelException(ctx, body, oauthProfileName, optional)
Create a Message VPN Access-Level Exception object.

Create a Message VPN Access-Level Exception object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates via config-sync.  Default message VPN access-level exceptions.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: msgVpnName|x|x|||| oauthProfileName|x||x|||    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OauthProfileDefaultMsgVpnAccessLevelException**](OauthProfileDefaultMsgVpnAccessLevelException.md)| The Message VPN Access-Level Exception object&#x27;s attributes. | 
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
 **optional** | ***OauthProfileApiCreateOauthProfileDefaultMsgVpnAccessLevelExceptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthProfileApiCreateOauthProfileDefaultMsgVpnAccessLevelExceptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**OauthProfileDefaultMsgVpnAccessLevelExceptionResponse**](OauthProfileDefaultMsgVpnAccessLevelExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOauthProfileResourceServerRequiredClaim**
> OauthProfileResourceServerRequiredClaimResponse CreateOauthProfileResourceServerRequiredClaim(ctx, body, oauthProfileName, optional)
Create a Required Claim object.

Create a Required Claim object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates via config-sync.  Additional claims to be verified in the access token.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: oauthProfileName|x||x||| resourceServerRequiredClaimName|x|x|||| resourceServerRequiredClaimValue||x||||    A SEMP client authorized with a minimum access scope/level of \"global/admin\" is required to perform this operation.  This has been available since 2.24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OauthProfileResourceServerRequiredClaim**](OauthProfileResourceServerRequiredClaim.md)| The Required Claim object&#x27;s attributes. | 
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
 **optional** | ***OauthProfileApiCreateOauthProfileResourceServerRequiredClaimOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthProfileApiCreateOauthProfileResourceServerRequiredClaimOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**OauthProfileResourceServerRequiredClaimResponse**](OauthProfileResourceServerRequiredClaimResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOauthProfile**
> SempMetaOnlyResponse DeleteOauthProfile(ctx, oauthProfileName)
Delete an OAuth Profile object.

Delete an OAuth Profile object. The deletion of instances of this object are synchronized to HA mates via config-sync.  OAuth profiles specify how to securely authenticate to an OAuth provider.  A SEMP client authorized with a minimum access scope/level of \"global/admin\" is required to perform this operation.  This has been available since 2.24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **oauthProfileName** | **string**| The name of the OAuth profile. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOauthProfileAccessLevelGroup**
> SempMetaOnlyResponse DeleteOauthProfileAccessLevelGroup(ctx, oauthProfileName, groupName)
Delete a Group Access Level object.

Delete a Group Access Level object. The deletion of instances of this object are synchronized to HA mates via config-sync.  The name of a group as it exists on the OAuth server being used to authenticate SEMP users.  A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
  **groupName** | **string**| The name of the group. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOauthProfileAccessLevelGroupMsgVpnAccessLevelException**
> SempMetaOnlyResponse DeleteOauthProfileAccessLevelGroupMsgVpnAccessLevelException(ctx, oauthProfileName, groupName, msgVpnName)
Delete a Message VPN Access-Level Exception object.

Delete a Message VPN Access-Level Exception object. The deletion of instances of this object are synchronized to HA mates via config-sync.  Message VPN access-level exceptions for members of this group.  A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
  **groupName** | **string**| The name of the group. | 
  **msgVpnName** | **string**| The name of the message VPN. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOauthProfileClientAllowedHost**
> SempMetaOnlyResponse DeleteOauthProfileClientAllowedHost(ctx, oauthProfileName, allowedHost)
Delete an Allowed Host Value object.

Delete an Allowed Host Value object. The deletion of instances of this object are synchronized to HA mates via config-sync.  A valid hostname for this broker in OAuth redirects.  A SEMP client authorized with a minimum access scope/level of \"global/admin\" is required to perform this operation.  This has been available since 2.24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
  **allowedHost** | **string**| An allowed value for the Host header. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOauthProfileClientAuthorizationParameter**
> SempMetaOnlyResponse DeleteOauthProfileClientAuthorizationParameter(ctx, oauthProfileName, authorizationParameterName)
Delete an Authorization Parameter object.

Delete an Authorization Parameter object. The deletion of instances of this object are synchronized to HA mates via config-sync.  Additional parameters to be passed to the OAuth authorization endpoint.  A SEMP client authorized with a minimum access scope/level of \"global/admin\" is required to perform this operation.  This has been available since 2.24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
  **authorizationParameterName** | **string**| The name of the authorization parameter. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOauthProfileClientRequiredClaim**
> SempMetaOnlyResponse DeleteOauthProfileClientRequiredClaim(ctx, oauthProfileName, clientRequiredClaimName)
Delete a Required Claim object.

Delete a Required Claim object. The deletion of instances of this object are synchronized to HA mates via config-sync.  Additional claims to be verified in the ID token.  A SEMP client authorized with a minimum access scope/level of \"global/admin\" is required to perform this operation.  This has been available since 2.24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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

# **DeleteOauthProfileDefaultMsgVpnAccessLevelException**
> SempMetaOnlyResponse DeleteOauthProfileDefaultMsgVpnAccessLevelException(ctx, oauthProfileName, msgVpnName)
Delete a Message VPN Access-Level Exception object.

Delete a Message VPN Access-Level Exception object. The deletion of instances of this object are synchronized to HA mates via config-sync.  Default message VPN access-level exceptions.  A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
  **msgVpnName** | **string**| The name of the message VPN. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOauthProfileResourceServerRequiredClaim**
> SempMetaOnlyResponse DeleteOauthProfileResourceServerRequiredClaim(ctx, oauthProfileName, resourceServerRequiredClaimName)
Delete a Required Claim object.

Delete a Required Claim object. The deletion of instances of this object are synchronized to HA mates via config-sync.  Additional claims to be verified in the access token.  A SEMP client authorized with a minimum access scope/level of \"global/admin\" is required to perform this operation.  This has been available since 2.24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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

# **GetOauthProfile**
> OauthProfileResponse GetOauthProfile(ctx, oauthProfileName, optional)
Get an OAuth Profile object.

Get an OAuth Profile object.  OAuth profiles specify how to securely authenticate to an OAuth provider.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: clientSecret||x||x oauthProfileName|x|||    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
 **optional** | ***OauthProfileApiGetOauthProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthProfileApiGetOauthProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**OauthProfileResponse**](OauthProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOauthProfileAccessLevelGroup**
> OauthProfileAccessLevelGroupResponse GetOauthProfileAccessLevelGroup(ctx, oauthProfileName, groupName, optional)
Get a Group Access Level object.

Get a Group Access Level object.  The name of a group as it exists on the OAuth server being used to authenticate SEMP users.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: groupName|x||| oauthProfileName|x|||    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
  **groupName** | **string**| The name of the group. | 
 **optional** | ***OauthProfileApiGetOauthProfileAccessLevelGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthProfileApiGetOauthProfileAccessLevelGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**OauthProfileAccessLevelGroupResponse**](OauthProfileAccessLevelGroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOauthProfileAccessLevelGroupMsgVpnAccessLevelException**
> OauthProfileAccessLevelGroupMsgVpnAccessLevelExceptionResponse GetOauthProfileAccessLevelGroupMsgVpnAccessLevelException(ctx, oauthProfileName, groupName, msgVpnName, optional)
Get a Message VPN Access-Level Exception object.

Get a Message VPN Access-Level Exception object.  Message VPN access-level exceptions for members of this group.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: groupName|x||| msgVpnName|x||| oauthProfileName|x|||    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
  **groupName** | **string**| The name of the group. | 
  **msgVpnName** | **string**| The name of the message VPN. | 
 **optional** | ***OauthProfileApiGetOauthProfileAccessLevelGroupMsgVpnAccessLevelExceptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthProfileApiGetOauthProfileAccessLevelGroupMsgVpnAccessLevelExceptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**OauthProfileAccessLevelGroupMsgVpnAccessLevelExceptionResponse**](OauthProfileAccessLevelGroupMsgVpnAccessLevelExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOauthProfileAccessLevelGroupMsgVpnAccessLevelExceptions**
> OauthProfileAccessLevelGroupMsgVpnAccessLevelExceptionsResponse GetOauthProfileAccessLevelGroupMsgVpnAccessLevelExceptions(ctx, oauthProfileName, groupName, optional)
Get a list of Message VPN Access-Level Exception objects.

Get a list of Message VPN Access-Level Exception objects.  Message VPN access-level exceptions for members of this group.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: groupName|x||| msgVpnName|x||| oauthProfileName|x|||    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
  **groupName** | **string**| The name of the group. | 
 **optional** | ***OauthProfileApiGetOauthProfileAccessLevelGroupMsgVpnAccessLevelExceptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthProfileApiGetOauthProfileAccessLevelGroupMsgVpnAccessLevelExceptionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**OauthProfileAccessLevelGroupMsgVpnAccessLevelExceptionsResponse**](OauthProfileAccessLevelGroupMsgVpnAccessLevelExceptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOauthProfileAccessLevelGroups**
> OauthProfileAccessLevelGroupsResponse GetOauthProfileAccessLevelGroups(ctx, oauthProfileName, optional)
Get a list of Group Access Level objects.

Get a list of Group Access Level objects.  The name of a group as it exists on the OAuth server being used to authenticate SEMP users.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: groupName|x||| oauthProfileName|x|||    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
 **optional** | ***OauthProfileApiGetOauthProfileAccessLevelGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthProfileApiGetOauthProfileAccessLevelGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**OauthProfileAccessLevelGroupsResponse**](OauthProfileAccessLevelGroupsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOauthProfileClientAllowedHost**
> OauthProfileClientAllowedHostResponse GetOauthProfileClientAllowedHost(ctx, oauthProfileName, allowedHost, optional)
Get an Allowed Host Value object.

Get an Allowed Host Value object.  A valid hostname for this broker in OAuth redirects.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: allowedHost|x||| oauthProfileName|x|||    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
  **allowedHost** | **string**| An allowed value for the Host header. | 
 **optional** | ***OauthProfileApiGetOauthProfileClientAllowedHostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthProfileApiGetOauthProfileClientAllowedHostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**OauthProfileClientAllowedHostResponse**](OauthProfileClientAllowedHostResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOauthProfileClientAllowedHosts**
> OauthProfileClientAllowedHostsResponse GetOauthProfileClientAllowedHosts(ctx, oauthProfileName, optional)
Get a list of Allowed Host Value objects.

Get a list of Allowed Host Value objects.  A valid hostname for this broker in OAuth redirects.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: allowedHost|x||| oauthProfileName|x|||    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
 **optional** | ***OauthProfileApiGetOauthProfileClientAllowedHostsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthProfileApiGetOauthProfileClientAllowedHostsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**OauthProfileClientAllowedHostsResponse**](OauthProfileClientAllowedHostsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOauthProfileClientAuthorizationParameter**
> OauthProfileClientAuthorizationParameterResponse GetOauthProfileClientAuthorizationParameter(ctx, oauthProfileName, authorizationParameterName, optional)
Get an Authorization Parameter object.

Get an Authorization Parameter object.  Additional parameters to be passed to the OAuth authorization endpoint.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: authorizationParameterName|x||| oauthProfileName|x|||    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
  **authorizationParameterName** | **string**| The name of the authorization parameter. | 
 **optional** | ***OauthProfileApiGetOauthProfileClientAuthorizationParameterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthProfileApiGetOauthProfileClientAuthorizationParameterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**OauthProfileClientAuthorizationParameterResponse**](OauthProfileClientAuthorizationParameterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOauthProfileClientAuthorizationParameters**
> OauthProfileClientAuthorizationParametersResponse GetOauthProfileClientAuthorizationParameters(ctx, oauthProfileName, optional)
Get a list of Authorization Parameter objects.

Get a list of Authorization Parameter objects.  Additional parameters to be passed to the OAuth authorization endpoint.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: authorizationParameterName|x||| oauthProfileName|x|||    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
 **optional** | ***OauthProfileApiGetOauthProfileClientAuthorizationParametersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthProfileApiGetOauthProfileClientAuthorizationParametersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**OauthProfileClientAuthorizationParametersResponse**](OauthProfileClientAuthorizationParametersResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOauthProfileClientRequiredClaim**
> OauthProfileClientRequiredClaimResponse GetOauthProfileClientRequiredClaim(ctx, oauthProfileName, clientRequiredClaimName, optional)
Get a Required Claim object.

Get a Required Claim object.  Additional claims to be verified in the ID token.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: clientRequiredClaimName|x||| oauthProfileName|x|||    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
  **clientRequiredClaimName** | **string**| The name of the ID token claim to verify. | 
 **optional** | ***OauthProfileApiGetOauthProfileClientRequiredClaimOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthProfileApiGetOauthProfileClientRequiredClaimOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**OauthProfileClientRequiredClaimResponse**](OauthProfileClientRequiredClaimResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOauthProfileClientRequiredClaims**
> OauthProfileClientRequiredClaimsResponse GetOauthProfileClientRequiredClaims(ctx, oauthProfileName, optional)
Get a list of Required Claim objects.

Get a list of Required Claim objects.  Additional claims to be verified in the ID token.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: clientRequiredClaimName|x||| oauthProfileName|x|||    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
 **optional** | ***OauthProfileApiGetOauthProfileClientRequiredClaimsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthProfileApiGetOauthProfileClientRequiredClaimsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**OauthProfileClientRequiredClaimsResponse**](OauthProfileClientRequiredClaimsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOauthProfileDefaultMsgVpnAccessLevelException**
> OauthProfileDefaultMsgVpnAccessLevelExceptionResponse GetOauthProfileDefaultMsgVpnAccessLevelException(ctx, oauthProfileName, msgVpnName, optional)
Get a Message VPN Access-Level Exception object.

Get a Message VPN Access-Level Exception object.  Default message VPN access-level exceptions.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| oauthProfileName|x|||    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
  **msgVpnName** | **string**| The name of the message VPN. | 
 **optional** | ***OauthProfileApiGetOauthProfileDefaultMsgVpnAccessLevelExceptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthProfileApiGetOauthProfileDefaultMsgVpnAccessLevelExceptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**OauthProfileDefaultMsgVpnAccessLevelExceptionResponse**](OauthProfileDefaultMsgVpnAccessLevelExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOauthProfileDefaultMsgVpnAccessLevelExceptions**
> OauthProfileDefaultMsgVpnAccessLevelExceptionsResponse GetOauthProfileDefaultMsgVpnAccessLevelExceptions(ctx, oauthProfileName, optional)
Get a list of Message VPN Access-Level Exception objects.

Get a list of Message VPN Access-Level Exception objects.  Default message VPN access-level exceptions.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| oauthProfileName|x|||    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
 **optional** | ***OauthProfileApiGetOauthProfileDefaultMsgVpnAccessLevelExceptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthProfileApiGetOauthProfileDefaultMsgVpnAccessLevelExceptionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**OauthProfileDefaultMsgVpnAccessLevelExceptionsResponse**](OauthProfileDefaultMsgVpnAccessLevelExceptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOauthProfileResourceServerRequiredClaim**
> OauthProfileResourceServerRequiredClaimResponse GetOauthProfileResourceServerRequiredClaim(ctx, oauthProfileName, resourceServerRequiredClaimName, optional)
Get a Required Claim object.

Get a Required Claim object.  Additional claims to be verified in the access token.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: oauthProfileName|x||| resourceServerRequiredClaimName|x|||    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
  **resourceServerRequiredClaimName** | **string**| The name of the access token claim to verify. | 
 **optional** | ***OauthProfileApiGetOauthProfileResourceServerRequiredClaimOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthProfileApiGetOauthProfileResourceServerRequiredClaimOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**OauthProfileResourceServerRequiredClaimResponse**](OauthProfileResourceServerRequiredClaimResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOauthProfileResourceServerRequiredClaims**
> OauthProfileResourceServerRequiredClaimsResponse GetOauthProfileResourceServerRequiredClaims(ctx, oauthProfileName, optional)
Get a list of Required Claim objects.

Get a list of Required Claim objects.  Additional claims to be verified in the access token.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: oauthProfileName|x||| resourceServerRequiredClaimName|x|||    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
 **optional** | ***OauthProfileApiGetOauthProfileResourceServerRequiredClaimsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthProfileApiGetOauthProfileResourceServerRequiredClaimsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**OauthProfileResourceServerRequiredClaimsResponse**](OauthProfileResourceServerRequiredClaimsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOauthProfiles**
> OauthProfilesResponse GetOauthProfiles(ctx, optional)
Get a list of OAuth Profile objects.

Get a list of OAuth Profile objects.  OAuth profiles specify how to securely authenticate to an OAuth provider.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: clientSecret||x||x oauthProfileName|x|||    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OauthProfileApiGetOauthProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthProfileApiGetOauthProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**OauthProfilesResponse**](OauthProfilesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceOauthProfile**
> OauthProfileResponse ReplaceOauthProfile(ctx, body, oauthProfileName, optional)
Replace an OAuth Profile object.

Replace an OAuth Profile object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  OAuth profiles specify how to securely authenticate to an OAuth provider.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- clientSecret||||x||||x oauthProfileName|x||x|||||    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation. Requests which include the following attributes require greater access scope/level:   Attribute|Access Scope/Level :---|:---: accessLevelGroupsClaimName|global/admin accessLevelGroupsClaimStringFormat|global/admin clientId|global/admin clientRedirectUri|global/admin clientRequiredType|global/admin clientScope|global/admin clientSecret|global/admin clientValidateTypeEnabled|global/admin defaultGlobalAccessLevel|global/admin displayName|global/admin enabled|global/admin endpointAuthorization|global/admin endpointDiscovery|global/admin endpointDiscoveryRefreshInterval|global/admin endpointIntrospection|global/admin endpointIntrospectionTimeout|global/admin endpointJwks|global/admin endpointJwksRefreshInterval|global/admin endpointToken|global/admin endpointTokenTimeout|global/admin endpointUserinfo|global/admin endpointUserinfoTimeout|global/admin interactiveEnabled|global/admin interactivePromptForExpiredSession|global/admin interactivePromptForNewSession|global/admin issuer|global/admin oauthRole|global/admin resourceServerParseAccessTokenEnabled|global/admin resourceServerRequiredAudience|global/admin resourceServerRequiredIssuer|global/admin resourceServerRequiredScope|global/admin resourceServerRequiredType|global/admin resourceServerValidateAudienceEnabled|global/admin resourceServerValidateIssuerEnabled|global/admin resourceServerValidateScopeEnabled|global/admin resourceServerValidateTypeEnabled|global/admin sempEnabled|global/admin usernameClaimName|global/admin    This has been available since 2.24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OauthProfile**](OauthProfile.md)| The OAuth Profile object&#x27;s attributes. | 
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
 **optional** | ***OauthProfileApiReplaceOauthProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthProfileApiReplaceOauthProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**OauthProfileResponse**](OauthProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceOauthProfileAccessLevelGroup**
> OauthProfileAccessLevelGroupResponse ReplaceOauthProfileAccessLevelGroup(ctx, body, oauthProfileName, groupName, optional)
Replace a Group Access Level object.

Replace a Group Access Level object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  The name of a group as it exists on the OAuth server being used to authenticate SEMP users.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- groupName|x||x||||| oauthProfileName|x||x|||||    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation. Requests which include the following attributes require greater access scope/level:   Attribute|Access Scope/Level :---|:---: globalAccessLevel|global/admin    This has been available since 2.24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OauthProfileAccessLevelGroup**](OauthProfileAccessLevelGroup.md)| The Group Access Level object&#x27;s attributes. | 
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
  **groupName** | **string**| The name of the group. | 
 **optional** | ***OauthProfileApiReplaceOauthProfileAccessLevelGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthProfileApiReplaceOauthProfileAccessLevelGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**OauthProfileAccessLevelGroupResponse**](OauthProfileAccessLevelGroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceOauthProfileAccessLevelGroupMsgVpnAccessLevelException**
> OauthProfileAccessLevelGroupMsgVpnAccessLevelExceptionResponse ReplaceOauthProfileAccessLevelGroupMsgVpnAccessLevelException(ctx, body, oauthProfileName, groupName, msgVpnName, optional)
Replace a Message VPN Access-Level Exception object.

Replace a Message VPN Access-Level Exception object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  Message VPN access-level exceptions for members of this group.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- groupName|x||x||||| msgVpnName|x||x||||| oauthProfileName|x||x|||||    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OauthProfileAccessLevelGroupMsgVpnAccessLevelException**](OauthProfileAccessLevelGroupMsgVpnAccessLevelException.md)| The Message VPN Access-Level Exception object&#x27;s attributes. | 
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
  **groupName** | **string**| The name of the group. | 
  **msgVpnName** | **string**| The name of the message VPN. | 
 **optional** | ***OauthProfileApiReplaceOauthProfileAccessLevelGroupMsgVpnAccessLevelExceptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthProfileApiReplaceOauthProfileAccessLevelGroupMsgVpnAccessLevelExceptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**OauthProfileAccessLevelGroupMsgVpnAccessLevelExceptionResponse**](OauthProfileAccessLevelGroupMsgVpnAccessLevelExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceOauthProfileClientAuthorizationParameter**
> OauthProfileClientAuthorizationParameterResponse ReplaceOauthProfileClientAuthorizationParameter(ctx, body, oauthProfileName, authorizationParameterName, optional)
Replace an Authorization Parameter object.

Replace an Authorization Parameter object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  Additional parameters to be passed to the OAuth authorization endpoint.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- authorizationParameterName|x||x||||| oauthProfileName|x||x|||||    A SEMP client authorized with a minimum access scope/level of \"global/admin\" is required to perform this operation.  This has been available since 2.24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OauthProfileClientAuthorizationParameter**](OauthProfileClientAuthorizationParameter.md)| The Authorization Parameter object&#x27;s attributes. | 
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
  **authorizationParameterName** | **string**| The name of the authorization parameter. | 
 **optional** | ***OauthProfileApiReplaceOauthProfileClientAuthorizationParameterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthProfileApiReplaceOauthProfileClientAuthorizationParameterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**OauthProfileClientAuthorizationParameterResponse**](OauthProfileClientAuthorizationParameterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceOauthProfileDefaultMsgVpnAccessLevelException**
> OauthProfileDefaultMsgVpnAccessLevelExceptionResponse ReplaceOauthProfileDefaultMsgVpnAccessLevelException(ctx, body, oauthProfileName, msgVpnName, optional)
Replace a Message VPN Access-Level Exception object.

Replace a Message VPN Access-Level Exception object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  Default message VPN access-level exceptions.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- msgVpnName|x||x||||| oauthProfileName|x||x|||||    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OauthProfileDefaultMsgVpnAccessLevelException**](OauthProfileDefaultMsgVpnAccessLevelException.md)| The Message VPN Access-Level Exception object&#x27;s attributes. | 
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
  **msgVpnName** | **string**| The name of the message VPN. | 
 **optional** | ***OauthProfileApiReplaceOauthProfileDefaultMsgVpnAccessLevelExceptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthProfileApiReplaceOauthProfileDefaultMsgVpnAccessLevelExceptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**OauthProfileDefaultMsgVpnAccessLevelExceptionResponse**](OauthProfileDefaultMsgVpnAccessLevelExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOauthProfile**
> OauthProfileResponse UpdateOauthProfile(ctx, body, oauthProfileName, optional)
Update an OAuth Profile object.

Update an OAuth Profile object. Any attribute missing from the request will be left unchanged.  OAuth profiles specify how to securely authenticate to an OAuth provider.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- clientSecret|||x||||x oauthProfileName|x|x|||||    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation. Requests which include the following attributes require greater access scope/level:   Attribute|Access Scope/Level :---|:---: accessLevelGroupsClaimName|global/admin accessLevelGroupsClaimStringFormat|global/admin clientId|global/admin clientRedirectUri|global/admin clientRequiredType|global/admin clientScope|global/admin clientSecret|global/admin clientValidateTypeEnabled|global/admin defaultGlobalAccessLevel|global/admin displayName|global/admin enabled|global/admin endpointAuthorization|global/admin endpointDiscovery|global/admin endpointDiscoveryRefreshInterval|global/admin endpointIntrospection|global/admin endpointIntrospectionTimeout|global/admin endpointJwks|global/admin endpointJwksRefreshInterval|global/admin endpointToken|global/admin endpointTokenTimeout|global/admin endpointUserinfo|global/admin endpointUserinfoTimeout|global/admin interactiveEnabled|global/admin interactivePromptForExpiredSession|global/admin interactivePromptForNewSession|global/admin issuer|global/admin oauthRole|global/admin resourceServerParseAccessTokenEnabled|global/admin resourceServerRequiredAudience|global/admin resourceServerRequiredIssuer|global/admin resourceServerRequiredScope|global/admin resourceServerRequiredType|global/admin resourceServerValidateAudienceEnabled|global/admin resourceServerValidateIssuerEnabled|global/admin resourceServerValidateScopeEnabled|global/admin resourceServerValidateTypeEnabled|global/admin sempEnabled|global/admin usernameClaimName|global/admin    This has been available since 2.24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OauthProfile**](OauthProfile.md)| The OAuth Profile object&#x27;s attributes. | 
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
 **optional** | ***OauthProfileApiUpdateOauthProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthProfileApiUpdateOauthProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**OauthProfileResponse**](OauthProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOauthProfileAccessLevelGroup**
> OauthProfileAccessLevelGroupResponse UpdateOauthProfileAccessLevelGroup(ctx, body, oauthProfileName, groupName, optional)
Update a Group Access Level object.

Update a Group Access Level object. Any attribute missing from the request will be left unchanged.  The name of a group as it exists on the OAuth server being used to authenticate SEMP users.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- groupName|x|x||||| oauthProfileName|x|x|||||    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation. Requests which include the following attributes require greater access scope/level:   Attribute|Access Scope/Level :---|:---: globalAccessLevel|global/admin    This has been available since 2.24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OauthProfileAccessLevelGroup**](OauthProfileAccessLevelGroup.md)| The Group Access Level object&#x27;s attributes. | 
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
  **groupName** | **string**| The name of the group. | 
 **optional** | ***OauthProfileApiUpdateOauthProfileAccessLevelGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthProfileApiUpdateOauthProfileAccessLevelGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**OauthProfileAccessLevelGroupResponse**](OauthProfileAccessLevelGroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOauthProfileAccessLevelGroupMsgVpnAccessLevelException**
> OauthProfileAccessLevelGroupMsgVpnAccessLevelExceptionResponse UpdateOauthProfileAccessLevelGroupMsgVpnAccessLevelException(ctx, body, oauthProfileName, groupName, msgVpnName, optional)
Update a Message VPN Access-Level Exception object.

Update a Message VPN Access-Level Exception object. Any attribute missing from the request will be left unchanged.  Message VPN access-level exceptions for members of this group.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- groupName|x|x||||| msgVpnName|x|x||||| oauthProfileName|x|x|||||    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OauthProfileAccessLevelGroupMsgVpnAccessLevelException**](OauthProfileAccessLevelGroupMsgVpnAccessLevelException.md)| The Message VPN Access-Level Exception object&#x27;s attributes. | 
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
  **groupName** | **string**| The name of the group. | 
  **msgVpnName** | **string**| The name of the message VPN. | 
 **optional** | ***OauthProfileApiUpdateOauthProfileAccessLevelGroupMsgVpnAccessLevelExceptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthProfileApiUpdateOauthProfileAccessLevelGroupMsgVpnAccessLevelExceptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**OauthProfileAccessLevelGroupMsgVpnAccessLevelExceptionResponse**](OauthProfileAccessLevelGroupMsgVpnAccessLevelExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOauthProfileClientAuthorizationParameter**
> OauthProfileClientAuthorizationParameterResponse UpdateOauthProfileClientAuthorizationParameter(ctx, body, oauthProfileName, authorizationParameterName, optional)
Update an Authorization Parameter object.

Update an Authorization Parameter object. Any attribute missing from the request will be left unchanged.  Additional parameters to be passed to the OAuth authorization endpoint.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- authorizationParameterName|x|x||||| oauthProfileName|x|x|||||    A SEMP client authorized with a minimum access scope/level of \"global/admin\" is required to perform this operation.  This has been available since 2.24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OauthProfileClientAuthorizationParameter**](OauthProfileClientAuthorizationParameter.md)| The Authorization Parameter object&#x27;s attributes. | 
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
  **authorizationParameterName** | **string**| The name of the authorization parameter. | 
 **optional** | ***OauthProfileApiUpdateOauthProfileClientAuthorizationParameterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthProfileApiUpdateOauthProfileClientAuthorizationParameterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**OauthProfileClientAuthorizationParameterResponse**](OauthProfileClientAuthorizationParameterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOauthProfileDefaultMsgVpnAccessLevelException**
> OauthProfileDefaultMsgVpnAccessLevelExceptionResponse UpdateOauthProfileDefaultMsgVpnAccessLevelException(ctx, body, oauthProfileName, msgVpnName, optional)
Update a Message VPN Access-Level Exception object.

Update a Message VPN Access-Level Exception object. Any attribute missing from the request will be left unchanged.  Default message VPN access-level exceptions.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- msgVpnName|x|x||||| oauthProfileName|x|x|||||    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OauthProfileDefaultMsgVpnAccessLevelException**](OauthProfileDefaultMsgVpnAccessLevelException.md)| The Message VPN Access-Level Exception object&#x27;s attributes. | 
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
  **msgVpnName** | **string**| The name of the message VPN. | 
 **optional** | ***OauthProfileApiUpdateOauthProfileDefaultMsgVpnAccessLevelExceptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthProfileApiUpdateOauthProfileDefaultMsgVpnAccessLevelExceptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**OauthProfileDefaultMsgVpnAccessLevelExceptionResponse**](OauthProfileDefaultMsgVpnAccessLevelExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

