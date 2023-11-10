# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/monitor*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMsgVpnAuthenticationOauthProvider**](AuthenticationOauthProviderApi.md#GetMsgVpnAuthenticationOauthProvider) | **Get** /msgVpns/{msgVpnName}/authenticationOauthProviders/{oauthProviderName} | Get an OAuth Provider object.
[**GetMsgVpnAuthenticationOauthProviders**](AuthenticationOauthProviderApi.md#GetMsgVpnAuthenticationOauthProviders) | **Get** /msgVpns/{msgVpnName}/authenticationOauthProviders | Get a list of OAuth Provider objects.

# **GetMsgVpnAuthenticationOauthProvider**
> MsgVpnAuthenticationOauthProviderResponse GetMsgVpnAuthenticationOauthProvider(ctx, msgVpnName, oauthProviderName, optional)
Get an OAuth Provider object.

Get an OAuth Provider object.  OAuth Providers contain information about the issuer of an OAuth token that is needed to validate the token and derive a client username from it.   Attribute|Identifying|Deprecated :---|:---:|:---: audienceClaimName||x audienceClaimSource||x audienceClaimValue||x audienceValidationEnabled||x authenticationSuccessCount||x authorizationGroupClaimName||x authorizationGroupClaimSource||x authorizationGroupEnabled||x disconnectOnTokenExpirationEnabled||x enabled||x jwksLastRefreshFailureReason||x jwksLastRefreshFailureTime||x jwksLastRefreshTime||x jwksNextScheduledRefreshTime||x jwksRefreshFailureCount||x jwksRefreshInterval||x jwksUri||x loginFailureIncorrectAudienceValueCount||x loginFailureInvalidAudienceValueCount||x loginFailureInvalidAuthorizationGroupValueCount||x loginFailureInvalidJwtSignatureCount||x loginFailureInvalidUsernameValueCount||x loginFailureMismatchedUsernameCount||x loginFailureMissingAudienceCount||x loginFailureMissingJwkCount||x loginFailureMissingOrInvalidTokenCount||x loginFailureMissingUsernameCount||x loginFailureTokenExpiredCount||x loginFailureTokenIntrospectionErroredCount||x loginFailureTokenIntrospectionFailureCount||x loginFailureTokenIntrospectionHttpsErrorCount||x loginFailureTokenIntrospectionInvalidCount||x loginFailureTokenIntrospectionTimeoutCount||x loginFailureTokenNotValidYetCount||x loginFailureUnsupportedAlgCount||x missingAuthorizationGroupCount||x msgVpnName|x|x oauthProviderName|x|x tokenIgnoreTimeLimitsEnabled||x tokenIntrospectionAverageTime||x tokenIntrospectionLastFailureReason||x tokenIntrospectionLastFailureTime||x tokenIntrospectionParameterName||x tokenIntrospectionSuccessCount||x tokenIntrospectionTimeout||x tokenIntrospectionUri||x tokenIntrospectionUsername||x usernameClaimName||x usernameClaimSource||x usernameValidateEnabled||x    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been deprecated since 2.25. Replaced by authenticationOauthProfiles.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **oauthProviderName** | **string**| The name of the OAuth Provider. | 
 **optional** | ***AuthenticationOauthProviderApiGetMsgVpnAuthenticationOauthProviderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationOauthProviderApiGetMsgVpnAuthenticationOauthProviderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthenticationOauthProviderResponse**](MsgVpnAuthenticationOauthProviderResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAuthenticationOauthProviders**
> MsgVpnAuthenticationOauthProvidersResponse GetMsgVpnAuthenticationOauthProviders(ctx, msgVpnName, optional)
Get a list of OAuth Provider objects.

Get a list of OAuth Provider objects.  OAuth Providers contain information about the issuer of an OAuth token that is needed to validate the token and derive a client username from it.   Attribute|Identifying|Deprecated :---|:---:|:---: audienceClaimName||x audienceClaimSource||x audienceClaimValue||x audienceValidationEnabled||x authenticationSuccessCount||x authorizationGroupClaimName||x authorizationGroupClaimSource||x authorizationGroupEnabled||x disconnectOnTokenExpirationEnabled||x enabled||x jwksLastRefreshFailureReason||x jwksLastRefreshFailureTime||x jwksLastRefreshTime||x jwksNextScheduledRefreshTime||x jwksRefreshFailureCount||x jwksRefreshInterval||x jwksUri||x loginFailureIncorrectAudienceValueCount||x loginFailureInvalidAudienceValueCount||x loginFailureInvalidAuthorizationGroupValueCount||x loginFailureInvalidJwtSignatureCount||x loginFailureInvalidUsernameValueCount||x loginFailureMismatchedUsernameCount||x loginFailureMissingAudienceCount||x loginFailureMissingJwkCount||x loginFailureMissingOrInvalidTokenCount||x loginFailureMissingUsernameCount||x loginFailureTokenExpiredCount||x loginFailureTokenIntrospectionErroredCount||x loginFailureTokenIntrospectionFailureCount||x loginFailureTokenIntrospectionHttpsErrorCount||x loginFailureTokenIntrospectionInvalidCount||x loginFailureTokenIntrospectionTimeoutCount||x loginFailureTokenNotValidYetCount||x loginFailureUnsupportedAlgCount||x missingAuthorizationGroupCount||x msgVpnName|x|x oauthProviderName|x|x tokenIgnoreTimeLimitsEnabled||x tokenIntrospectionAverageTime||x tokenIntrospectionLastFailureReason||x tokenIntrospectionLastFailureTime||x tokenIntrospectionParameterName||x tokenIntrospectionSuccessCount||x tokenIntrospectionTimeout||x tokenIntrospectionUri||x tokenIntrospectionUsername||x usernameClaimName||x usernameClaimSource||x usernameValidateEnabled||x    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been deprecated since 2.25. Replaced by authenticationOauthProfiles.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***AuthenticationOauthProviderApiGetMsgVpnAuthenticationOauthProvidersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthenticationOauthProviderApiGetMsgVpnAuthenticationOauthProvidersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthenticationOauthProvidersResponse**](MsgVpnAuthenticationOauthProvidersResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

