# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/monitor*

Method | HTTP request | Description
------------- | ------------- | -------------
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

# **GetOauthProfile**
> OauthProfileResponse GetOauthProfile(ctx, oauthProfileName, optional)
Get an OAuth Profile object.

Get an OAuth Profile object.  OAuth profiles specify how to securely authenticate to an OAuth provider.   Attribute|Identifying|Deprecated :---|:---:|:---: oauthProfileName|x|    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.24.

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

Get a Group Access Level object.  The name of a group as it exists on the OAuth server being used to authenticate SEMP users.   Attribute|Identifying|Deprecated :---|:---:|:---: groupName|x| oauthProfileName|x|    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.24.

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

Get a Message VPN Access-Level Exception object.  Message VPN access-level exceptions for members of this group.   Attribute|Identifying|Deprecated :---|:---:|:---: groupName|x| msgVpnName|x| oauthProfileName|x|    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.24.

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

Get a list of Message VPN Access-Level Exception objects.  Message VPN access-level exceptions for members of this group.   Attribute|Identifying|Deprecated :---|:---:|:---: groupName|x| msgVpnName|x| oauthProfileName|x|    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.24.

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

Get a list of Group Access Level objects.  The name of a group as it exists on the OAuth server being used to authenticate SEMP users.   Attribute|Identifying|Deprecated :---|:---:|:---: groupName|x| oauthProfileName|x|    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.24.

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

Get an Allowed Host Value object.  A valid hostname for this broker in OAuth redirects.   Attribute|Identifying|Deprecated :---|:---:|:---: allowedHost|x| oauthProfileName|x|    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.24.

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

Get a list of Allowed Host Value objects.  A valid hostname for this broker in OAuth redirects.   Attribute|Identifying|Deprecated :---|:---:|:---: allowedHost|x| oauthProfileName|x|    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.24.

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

Get an Authorization Parameter object.  Additional parameters to be passed to the OAuth authorization endpoint.   Attribute|Identifying|Deprecated :---|:---:|:---: authorizationParameterName|x| oauthProfileName|x|    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.24.

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

Get a list of Authorization Parameter objects.  Additional parameters to be passed to the OAuth authorization endpoint.   Attribute|Identifying|Deprecated :---|:---:|:---: authorizationParameterName|x| oauthProfileName|x|    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.24.

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

Get a Required Claim object.  Additional claims to be verified in the ID token.   Attribute|Identifying|Deprecated :---|:---:|:---: clientRequiredClaimName|x| oauthProfileName|x|    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.24.

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

Get a list of Required Claim objects.  Additional claims to be verified in the ID token.   Attribute|Identifying|Deprecated :---|:---:|:---: clientRequiredClaimName|x| oauthProfileName|x|    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.24.

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

Get a Message VPN Access-Level Exception object.  Default message VPN access-level exceptions.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| oauthProfileName|x|    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.24.

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

Get a list of Message VPN Access-Level Exception objects.  Default message VPN access-level exceptions.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| oauthProfileName|x|    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.24.

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

Get a Required Claim object.  Additional claims to be verified in the access token.   Attribute|Identifying|Deprecated :---|:---:|:---: oauthProfileName|x| resourceServerRequiredClaimName|x|    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.24.

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

Get a list of Required Claim objects.  Additional claims to be verified in the access token.   Attribute|Identifying|Deprecated :---|:---:|:---: oauthProfileName|x| resourceServerRequiredClaimName|x|    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.24.

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

Get a list of OAuth Profile objects.  OAuth profiles specify how to securely authenticate to an OAuth provider.   Attribute|Identifying|Deprecated :---|:---:|:---: oauthProfileName|x|    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.24.

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

