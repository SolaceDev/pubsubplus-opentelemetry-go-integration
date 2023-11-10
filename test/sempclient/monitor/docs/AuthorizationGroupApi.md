# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/monitor*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMsgVpnAuthorizationGroup**](AuthorizationGroupApi.md#GetMsgVpnAuthorizationGroup) | **Get** /msgVpns/{msgVpnName}/authorizationGroups/{authorizationGroupName} | Get an Authorization Group object.
[**GetMsgVpnAuthorizationGroups**](AuthorizationGroupApi.md#GetMsgVpnAuthorizationGroups) | **Get** /msgVpns/{msgVpnName}/authorizationGroups | Get a list of Authorization Group objects.

# **GetMsgVpnAuthorizationGroup**
> MsgVpnAuthorizationGroupResponse GetMsgVpnAuthorizationGroup(ctx, msgVpnName, authorizationGroupName, optional)
Get an Authorization Group object.

Get an Authorization Group object.  To use client authorization groups configured on an external server to provide client authorizations, Authorization Group objects must be created on the Message VPN that match the authorization groups provisioned on the external server. These objects must be configured with the client profiles and ACL profiles that will be assigned to the clients that belong to those authorization groups. A newly created group is placed at the end of the group list which is the lowest priority.   Attribute|Identifying|Deprecated :---|:---:|:---: authorizationGroupName|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **authorizationGroupName** | **string**| The name of the Authorization Group. For LDAP groups, special care is needed if the group name contains special characters such as &#x27;#&#x27;, &#x27;+&#x27;, &#x27;;&#x27;, &#x27;&#x3D;&#x27; as the value of the group name returned from the LDAP server might prepend those characters with &#x27;\\&#x27;. For example a group name called &#x27;test#,lab,com&#x27; will be returned from the LDAP server as &#x27;test\\#,lab,com&#x27;. | 
 **optional** | ***AuthorizationGroupApiGetMsgVpnAuthorizationGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthorizationGroupApiGetMsgVpnAuthorizationGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthorizationGroupResponse**](MsgVpnAuthorizationGroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAuthorizationGroups**
> MsgVpnAuthorizationGroupsResponse GetMsgVpnAuthorizationGroups(ctx, msgVpnName, optional)
Get a list of Authorization Group objects.

Get a list of Authorization Group objects.  To use client authorization groups configured on an external server to provide client authorizations, Authorization Group objects must be created on the Message VPN that match the authorization groups provisioned on the external server. These objects must be configured with the client profiles and ACL profiles that will be assigned to the clients that belong to those authorization groups. A newly created group is placed at the end of the group list which is the lowest priority.   Attribute|Identifying|Deprecated :---|:---:|:---: authorizationGroupName|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***AuthorizationGroupApiGetMsgVpnAuthorizationGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthorizationGroupApiGetMsgVpnAuthorizationGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthorizationGroupsResponse**](MsgVpnAuthorizationGroupsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

