# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/action*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DoMsgVpnAuthenticationOauthProfileClearStats**](AuthenticationOauthProfileApi.md#DoMsgVpnAuthenticationOauthProfileClearStats) | **Put** /msgVpns/{msgVpnName}/authenticationOauthProfiles/{oauthProfileName}/clearStats | Clear the statistics for the OAuth Profile.
[**GetMsgVpnAuthenticationOauthProfile**](AuthenticationOauthProfileApi.md#GetMsgVpnAuthenticationOauthProfile) | **Get** /msgVpns/{msgVpnName}/authenticationOauthProfiles/{oauthProfileName} | Get an OAuth Profile object.
[**GetMsgVpnAuthenticationOauthProfiles**](AuthenticationOauthProfileApi.md#GetMsgVpnAuthenticationOauthProfiles) | **Get** /msgVpns/{msgVpnName}/authenticationOauthProfiles | Get a list of OAuth Profile objects.

# **DoMsgVpnAuthenticationOauthProfileClearStats**
> SempMetaOnlyResponse DoMsgVpnAuthenticationOauthProfileClearStats(ctx, body, msgVpnName, oauthProfileName)
Clear the statistics for the OAuth Profile.

Clear the statistics for the OAuth Profile.    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.25.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnAuthenticationOauthProfileClearStats**](MsgVpnAuthenticationOauthProfileClearStats.md)| The Clear Stats action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **oauthProfileName** | **string**| The name of the OAuth profile. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAuthenticationOauthProfile**
> MsgVpnAuthenticationOauthProfileResponse GetMsgVpnAuthenticationOauthProfile(ctx, msgVpnName, oauthProfileName, optional)
Get an OAuth Profile object.

Get an OAuth Profile object.  OAuth profiles specify how to securely authenticate to an OAuth provider.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| oauthProfileName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.25.

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


 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthenticationOauthProfileResponse**](MsgVpnAuthenticationOauthProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAuthenticationOauthProfiles**
> MsgVpnAuthenticationOauthProfilesResponse GetMsgVpnAuthenticationOauthProfiles(ctx, msgVpnName, optional)
Get a list of OAuth Profile objects.

Get a list of OAuth Profile objects.  OAuth profiles specify how to securely authenticate to an OAuth provider.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| oauthProfileName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.25.

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

