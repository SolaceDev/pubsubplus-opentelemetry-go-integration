# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/config*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMsgVpnAclProfile**](AclProfileApi.md#CreateMsgVpnAclProfile) | **Post** /msgVpns/{msgVpnName}/aclProfiles | Create an ACL Profile object.
[**CreateMsgVpnAclProfileClientConnectException**](AclProfileApi.md#CreateMsgVpnAclProfileClientConnectException) | **Post** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/clientConnectExceptions | Create a Client Connect Exception object.
[**CreateMsgVpnAclProfilePublishException**](AclProfileApi.md#CreateMsgVpnAclProfilePublishException) | **Post** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishExceptions | Create a Publish Topic Exception object.
[**CreateMsgVpnAclProfilePublishTopicException**](AclProfileApi.md#CreateMsgVpnAclProfilePublishTopicException) | **Post** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishTopicExceptions | Create a Publish Topic Exception object.
[**CreateMsgVpnAclProfileSubscribeException**](AclProfileApi.md#CreateMsgVpnAclProfileSubscribeException) | **Post** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeExceptions | Create a Subscribe Topic Exception object.
[**CreateMsgVpnAclProfileSubscribeShareNameException**](AclProfileApi.md#CreateMsgVpnAclProfileSubscribeShareNameException) | **Post** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeShareNameExceptions | Create a Subscribe Share Name Exception object.
[**CreateMsgVpnAclProfileSubscribeTopicException**](AclProfileApi.md#CreateMsgVpnAclProfileSubscribeTopicException) | **Post** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeTopicExceptions | Create a Subscribe Topic Exception object.
[**DeleteMsgVpnAclProfile**](AclProfileApi.md#DeleteMsgVpnAclProfile) | **Delete** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName} | Delete an ACL Profile object.
[**DeleteMsgVpnAclProfileClientConnectException**](AclProfileApi.md#DeleteMsgVpnAclProfileClientConnectException) | **Delete** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/clientConnectExceptions/{clientConnectExceptionAddress} | Delete a Client Connect Exception object.
[**DeleteMsgVpnAclProfilePublishException**](AclProfileApi.md#DeleteMsgVpnAclProfilePublishException) | **Delete** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishExceptions/{topicSyntax},{publishExceptionTopic} | Delete a Publish Topic Exception object.
[**DeleteMsgVpnAclProfilePublishTopicException**](AclProfileApi.md#DeleteMsgVpnAclProfilePublishTopicException) | **Delete** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishTopicExceptions/{publishTopicExceptionSyntax},{publishTopicException} | Delete a Publish Topic Exception object.
[**DeleteMsgVpnAclProfileSubscribeException**](AclProfileApi.md#DeleteMsgVpnAclProfileSubscribeException) | **Delete** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeExceptions/{topicSyntax},{subscribeExceptionTopic} | Delete a Subscribe Topic Exception object.
[**DeleteMsgVpnAclProfileSubscribeShareNameException**](AclProfileApi.md#DeleteMsgVpnAclProfileSubscribeShareNameException) | **Delete** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeShareNameExceptions/{subscribeShareNameExceptionSyntax},{subscribeShareNameException} | Delete a Subscribe Share Name Exception object.
[**DeleteMsgVpnAclProfileSubscribeTopicException**](AclProfileApi.md#DeleteMsgVpnAclProfileSubscribeTopicException) | **Delete** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeTopicExceptions/{subscribeTopicExceptionSyntax},{subscribeTopicException} | Delete a Subscribe Topic Exception object.
[**GetMsgVpnAclProfile**](AclProfileApi.md#GetMsgVpnAclProfile) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName} | Get an ACL Profile object.
[**GetMsgVpnAclProfileClientConnectException**](AclProfileApi.md#GetMsgVpnAclProfileClientConnectException) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/clientConnectExceptions/{clientConnectExceptionAddress} | Get a Client Connect Exception object.
[**GetMsgVpnAclProfileClientConnectExceptions**](AclProfileApi.md#GetMsgVpnAclProfileClientConnectExceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/clientConnectExceptions | Get a list of Client Connect Exception objects.
[**GetMsgVpnAclProfilePublishException**](AclProfileApi.md#GetMsgVpnAclProfilePublishException) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishExceptions/{topicSyntax},{publishExceptionTopic} | Get a Publish Topic Exception object.
[**GetMsgVpnAclProfilePublishExceptions**](AclProfileApi.md#GetMsgVpnAclProfilePublishExceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishExceptions | Get a list of Publish Topic Exception objects.
[**GetMsgVpnAclProfilePublishTopicException**](AclProfileApi.md#GetMsgVpnAclProfilePublishTopicException) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishTopicExceptions/{publishTopicExceptionSyntax},{publishTopicException} | Get a Publish Topic Exception object.
[**GetMsgVpnAclProfilePublishTopicExceptions**](AclProfileApi.md#GetMsgVpnAclProfilePublishTopicExceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishTopicExceptions | Get a list of Publish Topic Exception objects.
[**GetMsgVpnAclProfileSubscribeException**](AclProfileApi.md#GetMsgVpnAclProfileSubscribeException) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeExceptions/{topicSyntax},{subscribeExceptionTopic} | Get a Subscribe Topic Exception object.
[**GetMsgVpnAclProfileSubscribeExceptions**](AclProfileApi.md#GetMsgVpnAclProfileSubscribeExceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeExceptions | Get a list of Subscribe Topic Exception objects.
[**GetMsgVpnAclProfileSubscribeShareNameException**](AclProfileApi.md#GetMsgVpnAclProfileSubscribeShareNameException) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeShareNameExceptions/{subscribeShareNameExceptionSyntax},{subscribeShareNameException} | Get a Subscribe Share Name Exception object.
[**GetMsgVpnAclProfileSubscribeShareNameExceptions**](AclProfileApi.md#GetMsgVpnAclProfileSubscribeShareNameExceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeShareNameExceptions | Get a list of Subscribe Share Name Exception objects.
[**GetMsgVpnAclProfileSubscribeTopicException**](AclProfileApi.md#GetMsgVpnAclProfileSubscribeTopicException) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeTopicExceptions/{subscribeTopicExceptionSyntax},{subscribeTopicException} | Get a Subscribe Topic Exception object.
[**GetMsgVpnAclProfileSubscribeTopicExceptions**](AclProfileApi.md#GetMsgVpnAclProfileSubscribeTopicExceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeTopicExceptions | Get a list of Subscribe Topic Exception objects.
[**GetMsgVpnAclProfiles**](AclProfileApi.md#GetMsgVpnAclProfiles) | **Get** /msgVpns/{msgVpnName}/aclProfiles | Get a list of ACL Profile objects.
[**ReplaceMsgVpnAclProfile**](AclProfileApi.md#ReplaceMsgVpnAclProfile) | **Put** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName} | Replace an ACL Profile object.
[**UpdateMsgVpnAclProfile**](AclProfileApi.md#UpdateMsgVpnAclProfile) | **Patch** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName} | Update an ACL Profile object.

# **CreateMsgVpnAclProfile**
> MsgVpnAclProfileResponse CreateMsgVpnAclProfile(ctx, body, msgVpnName, optional)
Create an ACL Profile object.

Create an ACL Profile object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  An ACL Profile controls whether an authenticated client is permitted to establish a connection with the message broker or permitted to publish and subscribe to specific topics.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: aclProfileName|x|x|||| msgVpnName|x||x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnAclProfile**](MsgVpnAclProfile.md)| The ACL Profile object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***AclProfileApiCreateMsgVpnAclProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AclProfileApiCreateMsgVpnAclProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileResponse**](MsgVpnAclProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnAclProfileClientConnectException**
> MsgVpnAclProfileClientConnectExceptionResponse CreateMsgVpnAclProfileClientConnectException(ctx, body, msgVpnName, aclProfileName, optional)
Create a Client Connect Exception object.

Create a Client Connect Exception object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Client Connect Exception is an exception to the default action to take when a client using the ACL Profile connects to the Message VPN. Exceptions must be expressed as an IP address/netmask in CIDR form.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: aclProfileName|x||x||| clientConnectExceptionAddress|x|x|||| msgVpnName|x||x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnAclProfileClientConnectException**](MsgVpnAclProfileClientConnectException.md)| The Client Connect Exception object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
 **optional** | ***AclProfileApiCreateMsgVpnAclProfileClientConnectExceptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AclProfileApiCreateMsgVpnAclProfileClientConnectExceptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileClientConnectExceptionResponse**](MsgVpnAclProfileClientConnectExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnAclProfilePublishException**
> MsgVpnAclProfilePublishExceptionResponse CreateMsgVpnAclProfilePublishException(ctx, body, msgVpnName, aclProfileName, optional)
Create a Publish Topic Exception object.

Create a Publish Topic Exception object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Publish Topic Exception is an exception to the default action to take when a client using the ACL Profile publishes to a topic in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: aclProfileName|x||x||x| msgVpnName|x||x||x| publishExceptionTopic|x|x|||x| topicSyntax|x|x|||x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been deprecated since 2.14. Replaced by publishTopicExceptions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnAclProfilePublishException**](MsgVpnAclProfilePublishException.md)| The Publish Topic Exception object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
 **optional** | ***AclProfileApiCreateMsgVpnAclProfilePublishExceptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AclProfileApiCreateMsgVpnAclProfilePublishExceptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfilePublishExceptionResponse**](MsgVpnAclProfilePublishExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnAclProfilePublishTopicException**
> MsgVpnAclProfilePublishTopicExceptionResponse CreateMsgVpnAclProfilePublishTopicException(ctx, body, msgVpnName, aclProfileName, optional)
Create a Publish Topic Exception object.

Create a Publish Topic Exception object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Publish Topic Exception is an exception to the default action to take when a client using the ACL Profile publishes to a topic in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: aclProfileName|x||x||| msgVpnName|x||x||| publishTopicException|x|x|||| publishTopicExceptionSyntax|x|x||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnAclProfilePublishTopicException**](MsgVpnAclProfilePublishTopicException.md)| The Publish Topic Exception object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
 **optional** | ***AclProfileApiCreateMsgVpnAclProfilePublishTopicExceptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AclProfileApiCreateMsgVpnAclProfilePublishTopicExceptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfilePublishTopicExceptionResponse**](MsgVpnAclProfilePublishTopicExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnAclProfileSubscribeException**
> MsgVpnAclProfileSubscribeExceptionResponse CreateMsgVpnAclProfileSubscribeException(ctx, body, msgVpnName, aclProfileName, optional)
Create a Subscribe Topic Exception object.

Create a Subscribe Topic Exception object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Subscribe Topic Exception is an exception to the default action to take when a client using the ACL Profile subscribes to a topic in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: aclProfileName|x||x||x| msgVpnName|x||x||x| subscribeExceptionTopic|x|x|||x| topicSyntax|x|x|||x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been deprecated since 2.14. Replaced by subscribeTopicExceptions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnAclProfileSubscribeException**](MsgVpnAclProfileSubscribeException.md)| The Subscribe Topic Exception object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
 **optional** | ***AclProfileApiCreateMsgVpnAclProfileSubscribeExceptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AclProfileApiCreateMsgVpnAclProfileSubscribeExceptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileSubscribeExceptionResponse**](MsgVpnAclProfileSubscribeExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnAclProfileSubscribeShareNameException**
> MsgVpnAclProfileSubscribeShareNameExceptionResponse CreateMsgVpnAclProfileSubscribeShareNameException(ctx, body, msgVpnName, aclProfileName, optional)
Create a Subscribe Share Name Exception object.

Create a Subscribe Share Name Exception object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Subscribe Share Name Exception is an exception to the default action to take when a client using the ACL Profile subscribes to a share-name subscription in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: aclProfileName|x||x||| msgVpnName|x||x||| subscribeShareNameException|x|x|||| subscribeShareNameExceptionSyntax|x|x||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnAclProfileSubscribeShareNameException**](MsgVpnAclProfileSubscribeShareNameException.md)| The Subscribe Share Name Exception object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
 **optional** | ***AclProfileApiCreateMsgVpnAclProfileSubscribeShareNameExceptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AclProfileApiCreateMsgVpnAclProfileSubscribeShareNameExceptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileSubscribeShareNameExceptionResponse**](MsgVpnAclProfileSubscribeShareNameExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnAclProfileSubscribeTopicException**
> MsgVpnAclProfileSubscribeTopicExceptionResponse CreateMsgVpnAclProfileSubscribeTopicException(ctx, body, msgVpnName, aclProfileName, optional)
Create a Subscribe Topic Exception object.

Create a Subscribe Topic Exception object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Subscribe Topic Exception is an exception to the default action to take when a client using the ACL Profile subscribes to a topic in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: aclProfileName|x||x||| msgVpnName|x||x||| subscribeTopicException|x|x|||| subscribeTopicExceptionSyntax|x|x||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnAclProfileSubscribeTopicException**](MsgVpnAclProfileSubscribeTopicException.md)| The Subscribe Topic Exception object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
 **optional** | ***AclProfileApiCreateMsgVpnAclProfileSubscribeTopicExceptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AclProfileApiCreateMsgVpnAclProfileSubscribeTopicExceptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileSubscribeTopicExceptionResponse**](MsgVpnAclProfileSubscribeTopicExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnAclProfile**
> SempMetaOnlyResponse DeleteMsgVpnAclProfile(ctx, msgVpnName, aclProfileName)
Delete an ACL Profile object.

Delete an ACL Profile object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  An ACL Profile controls whether an authenticated client is permitted to establish a connection with the message broker or permitted to publish and subscribe to specific topics.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnAclProfileClientConnectException**
> SempMetaOnlyResponse DeleteMsgVpnAclProfileClientConnectException(ctx, msgVpnName, aclProfileName, clientConnectExceptionAddress)
Delete a Client Connect Exception object.

Delete a Client Connect Exception object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Client Connect Exception is an exception to the default action to take when a client using the ACL Profile connects to the Message VPN. Exceptions must be expressed as an IP address/netmask in CIDR form.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
  **clientConnectExceptionAddress** | **string**| The IP address/netmask of the client connect exception in canonical CIDR form. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnAclProfilePublishException**
> SempMetaOnlyResponse DeleteMsgVpnAclProfilePublishException(ctx, msgVpnName, aclProfileName, topicSyntax, publishExceptionTopic)
Delete a Publish Topic Exception object.

Delete a Publish Topic Exception object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Publish Topic Exception is an exception to the default action to take when a client using the ACL Profile publishes to a topic in the Message VPN. Exceptions must be expressed as a topic.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been deprecated since 2.14. Replaced by publishTopicExceptions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
  **topicSyntax** | **string**| The syntax of the topic for the exception to the default action taken. | 
  **publishExceptionTopic** | **string**| The topic for the exception to the default action taken. May include wildcard characters. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnAclProfilePublishTopicException**
> SempMetaOnlyResponse DeleteMsgVpnAclProfilePublishTopicException(ctx, msgVpnName, aclProfileName, publishTopicExceptionSyntax, publishTopicException)
Delete a Publish Topic Exception object.

Delete a Publish Topic Exception object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Publish Topic Exception is an exception to the default action to take when a client using the ACL Profile publishes to a topic in the Message VPN. Exceptions must be expressed as a topic.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
  **publishTopicExceptionSyntax** | **string**| The syntax of the topic for the exception to the default action taken. | 
  **publishTopicException** | **string**| The topic for the exception to the default action taken. May include wildcard characters. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnAclProfileSubscribeException**
> SempMetaOnlyResponse DeleteMsgVpnAclProfileSubscribeException(ctx, msgVpnName, aclProfileName, topicSyntax, subscribeExceptionTopic)
Delete a Subscribe Topic Exception object.

Delete a Subscribe Topic Exception object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Subscribe Topic Exception is an exception to the default action to take when a client using the ACL Profile subscribes to a topic in the Message VPN. Exceptions must be expressed as a topic.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been deprecated since 2.14. Replaced by subscribeTopicExceptions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
  **topicSyntax** | **string**| The syntax of the topic for the exception to the default action taken. | 
  **subscribeExceptionTopic** | **string**| The topic for the exception to the default action taken. May include wildcard characters. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnAclProfileSubscribeShareNameException**
> SempMetaOnlyResponse DeleteMsgVpnAclProfileSubscribeShareNameException(ctx, msgVpnName, aclProfileName, subscribeShareNameExceptionSyntax, subscribeShareNameException)
Delete a Subscribe Share Name Exception object.

Delete a Subscribe Share Name Exception object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Subscribe Share Name Exception is an exception to the default action to take when a client using the ACL Profile subscribes to a share-name subscription in the Message VPN. Exceptions must be expressed as a topic.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
  **subscribeShareNameExceptionSyntax** | **string**| The syntax of the subscribe share name for the exception to the default action taken. | 
  **subscribeShareNameException** | **string**| The subscribe share name exception to the default action taken. May include wildcard characters. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnAclProfileSubscribeTopicException**
> SempMetaOnlyResponse DeleteMsgVpnAclProfileSubscribeTopicException(ctx, msgVpnName, aclProfileName, subscribeTopicExceptionSyntax, subscribeTopicException)
Delete a Subscribe Topic Exception object.

Delete a Subscribe Topic Exception object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Subscribe Topic Exception is an exception to the default action to take when a client using the ACL Profile subscribes to a topic in the Message VPN. Exceptions must be expressed as a topic.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
  **subscribeTopicExceptionSyntax** | **string**| The syntax of the topic for the exception to the default action taken. | 
  **subscribeTopicException** | **string**| The topic for the exception to the default action taken. May include wildcard characters. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAclProfile**
> MsgVpnAclProfileResponse GetMsgVpnAclProfile(ctx, msgVpnName, aclProfileName, optional)
Get an ACL Profile object.

Get an ACL Profile object.  An ACL Profile controls whether an authenticated client is permitted to establish a connection with the message broker or permitted to publish and subscribe to specific topics.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: aclProfileName|x||| msgVpnName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
 **optional** | ***AclProfileApiGetMsgVpnAclProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AclProfileApiGetMsgVpnAclProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileResponse**](MsgVpnAclProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAclProfileClientConnectException**
> MsgVpnAclProfileClientConnectExceptionResponse GetMsgVpnAclProfileClientConnectException(ctx, msgVpnName, aclProfileName, clientConnectExceptionAddress, optional)
Get a Client Connect Exception object.

Get a Client Connect Exception object.  A Client Connect Exception is an exception to the default action to take when a client using the ACL Profile connects to the Message VPN. Exceptions must be expressed as an IP address/netmask in CIDR form.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: aclProfileName|x||| clientConnectExceptionAddress|x||| msgVpnName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
  **clientConnectExceptionAddress** | **string**| The IP address/netmask of the client connect exception in canonical CIDR form. | 
 **optional** | ***AclProfileApiGetMsgVpnAclProfileClientConnectExceptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AclProfileApiGetMsgVpnAclProfileClientConnectExceptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileClientConnectExceptionResponse**](MsgVpnAclProfileClientConnectExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAclProfileClientConnectExceptions**
> MsgVpnAclProfileClientConnectExceptionsResponse GetMsgVpnAclProfileClientConnectExceptions(ctx, msgVpnName, aclProfileName, optional)
Get a list of Client Connect Exception objects.

Get a list of Client Connect Exception objects.  A Client Connect Exception is an exception to the default action to take when a client using the ACL Profile connects to the Message VPN. Exceptions must be expressed as an IP address/netmask in CIDR form.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: aclProfileName|x||| clientConnectExceptionAddress|x||| msgVpnName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
 **optional** | ***AclProfileApiGetMsgVpnAclProfileClientConnectExceptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AclProfileApiGetMsgVpnAclProfileClientConnectExceptionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileClientConnectExceptionsResponse**](MsgVpnAclProfileClientConnectExceptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAclProfilePublishException**
> MsgVpnAclProfilePublishExceptionResponse GetMsgVpnAclProfilePublishException(ctx, msgVpnName, aclProfileName, topicSyntax, publishExceptionTopic, optional)
Get a Publish Topic Exception object.

Get a Publish Topic Exception object.  A Publish Topic Exception is an exception to the default action to take when a client using the ACL Profile publishes to a topic in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: aclProfileName|x||x| msgVpnName|x||x| publishExceptionTopic|x||x| topicSyntax|x||x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been deprecated since 2.14. Replaced by publishTopicExceptions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
  **topicSyntax** | **string**| The syntax of the topic for the exception to the default action taken. | 
  **publishExceptionTopic** | **string**| The topic for the exception to the default action taken. May include wildcard characters. | 
 **optional** | ***AclProfileApiGetMsgVpnAclProfilePublishExceptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AclProfileApiGetMsgVpnAclProfilePublishExceptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfilePublishExceptionResponse**](MsgVpnAclProfilePublishExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAclProfilePublishExceptions**
> MsgVpnAclProfilePublishExceptionsResponse GetMsgVpnAclProfilePublishExceptions(ctx, msgVpnName, aclProfileName, optional)
Get a list of Publish Topic Exception objects.

Get a list of Publish Topic Exception objects.  A Publish Topic Exception is an exception to the default action to take when a client using the ACL Profile publishes to a topic in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: aclProfileName|x||x| msgVpnName|x||x| publishExceptionTopic|x||x| topicSyntax|x||x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been deprecated since 2.14. Replaced by publishTopicExceptions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
 **optional** | ***AclProfileApiGetMsgVpnAclProfilePublishExceptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AclProfileApiGetMsgVpnAclProfilePublishExceptionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfilePublishExceptionsResponse**](MsgVpnAclProfilePublishExceptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAclProfilePublishTopicException**
> MsgVpnAclProfilePublishTopicExceptionResponse GetMsgVpnAclProfilePublishTopicException(ctx, msgVpnName, aclProfileName, publishTopicExceptionSyntax, publishTopicException, optional)
Get a Publish Topic Exception object.

Get a Publish Topic Exception object.  A Publish Topic Exception is an exception to the default action to take when a client using the ACL Profile publishes to a topic in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: aclProfileName|x||| msgVpnName|x||| publishTopicException|x||| publishTopicExceptionSyntax|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
  **publishTopicExceptionSyntax** | **string**| The syntax of the topic for the exception to the default action taken. | 
  **publishTopicException** | **string**| The topic for the exception to the default action taken. May include wildcard characters. | 
 **optional** | ***AclProfileApiGetMsgVpnAclProfilePublishTopicExceptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AclProfileApiGetMsgVpnAclProfilePublishTopicExceptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfilePublishTopicExceptionResponse**](MsgVpnAclProfilePublishTopicExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAclProfilePublishTopicExceptions**
> MsgVpnAclProfilePublishTopicExceptionsResponse GetMsgVpnAclProfilePublishTopicExceptions(ctx, msgVpnName, aclProfileName, optional)
Get a list of Publish Topic Exception objects.

Get a list of Publish Topic Exception objects.  A Publish Topic Exception is an exception to the default action to take when a client using the ACL Profile publishes to a topic in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: aclProfileName|x||| msgVpnName|x||| publishTopicException|x||| publishTopicExceptionSyntax|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
 **optional** | ***AclProfileApiGetMsgVpnAclProfilePublishTopicExceptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AclProfileApiGetMsgVpnAclProfilePublishTopicExceptionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfilePublishTopicExceptionsResponse**](MsgVpnAclProfilePublishTopicExceptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAclProfileSubscribeException**
> MsgVpnAclProfileSubscribeExceptionResponse GetMsgVpnAclProfileSubscribeException(ctx, msgVpnName, aclProfileName, topicSyntax, subscribeExceptionTopic, optional)
Get a Subscribe Topic Exception object.

Get a Subscribe Topic Exception object.  A Subscribe Topic Exception is an exception to the default action to take when a client using the ACL Profile subscribes to a topic in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: aclProfileName|x||x| msgVpnName|x||x| subscribeExceptionTopic|x||x| topicSyntax|x||x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been deprecated since 2.14. Replaced by subscribeTopicExceptions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
  **topicSyntax** | **string**| The syntax of the topic for the exception to the default action taken. | 
  **subscribeExceptionTopic** | **string**| The topic for the exception to the default action taken. May include wildcard characters. | 
 **optional** | ***AclProfileApiGetMsgVpnAclProfileSubscribeExceptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AclProfileApiGetMsgVpnAclProfileSubscribeExceptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileSubscribeExceptionResponse**](MsgVpnAclProfileSubscribeExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAclProfileSubscribeExceptions**
> MsgVpnAclProfileSubscribeExceptionsResponse GetMsgVpnAclProfileSubscribeExceptions(ctx, msgVpnName, aclProfileName, optional)
Get a list of Subscribe Topic Exception objects.

Get a list of Subscribe Topic Exception objects.  A Subscribe Topic Exception is an exception to the default action to take when a client using the ACL Profile subscribes to a topic in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: aclProfileName|x||x| msgVpnName|x||x| subscribeExceptionTopic|x||x| topicSyntax|x||x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been deprecated since 2.14. Replaced by subscribeTopicExceptions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
 **optional** | ***AclProfileApiGetMsgVpnAclProfileSubscribeExceptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AclProfileApiGetMsgVpnAclProfileSubscribeExceptionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileSubscribeExceptionsResponse**](MsgVpnAclProfileSubscribeExceptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAclProfileSubscribeShareNameException**
> MsgVpnAclProfileSubscribeShareNameExceptionResponse GetMsgVpnAclProfileSubscribeShareNameException(ctx, msgVpnName, aclProfileName, subscribeShareNameExceptionSyntax, subscribeShareNameException, optional)
Get a Subscribe Share Name Exception object.

Get a Subscribe Share Name Exception object.  A Subscribe Share Name Exception is an exception to the default action to take when a client using the ACL Profile subscribes to a share-name subscription in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: aclProfileName|x||| msgVpnName|x||| subscribeShareNameException|x||| subscribeShareNameExceptionSyntax|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
  **subscribeShareNameExceptionSyntax** | **string**| The syntax of the subscribe share name for the exception to the default action taken. | 
  **subscribeShareNameException** | **string**| The subscribe share name exception to the default action taken. May include wildcard characters. | 
 **optional** | ***AclProfileApiGetMsgVpnAclProfileSubscribeShareNameExceptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AclProfileApiGetMsgVpnAclProfileSubscribeShareNameExceptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileSubscribeShareNameExceptionResponse**](MsgVpnAclProfileSubscribeShareNameExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAclProfileSubscribeShareNameExceptions**
> MsgVpnAclProfileSubscribeShareNameExceptionsResponse GetMsgVpnAclProfileSubscribeShareNameExceptions(ctx, msgVpnName, aclProfileName, optional)
Get a list of Subscribe Share Name Exception objects.

Get a list of Subscribe Share Name Exception objects.  A Subscribe Share Name Exception is an exception to the default action to take when a client using the ACL Profile subscribes to a share-name subscription in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: aclProfileName|x||| msgVpnName|x||| subscribeShareNameException|x||| subscribeShareNameExceptionSyntax|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
 **optional** | ***AclProfileApiGetMsgVpnAclProfileSubscribeShareNameExceptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AclProfileApiGetMsgVpnAclProfileSubscribeShareNameExceptionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileSubscribeShareNameExceptionsResponse**](MsgVpnAclProfileSubscribeShareNameExceptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAclProfileSubscribeTopicException**
> MsgVpnAclProfileSubscribeTopicExceptionResponse GetMsgVpnAclProfileSubscribeTopicException(ctx, msgVpnName, aclProfileName, subscribeTopicExceptionSyntax, subscribeTopicException, optional)
Get a Subscribe Topic Exception object.

Get a Subscribe Topic Exception object.  A Subscribe Topic Exception is an exception to the default action to take when a client using the ACL Profile subscribes to a topic in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: aclProfileName|x||| msgVpnName|x||| subscribeTopicException|x||| subscribeTopicExceptionSyntax|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
  **subscribeTopicExceptionSyntax** | **string**| The syntax of the topic for the exception to the default action taken. | 
  **subscribeTopicException** | **string**| The topic for the exception to the default action taken. May include wildcard characters. | 
 **optional** | ***AclProfileApiGetMsgVpnAclProfileSubscribeTopicExceptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AclProfileApiGetMsgVpnAclProfileSubscribeTopicExceptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileSubscribeTopicExceptionResponse**](MsgVpnAclProfileSubscribeTopicExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAclProfileSubscribeTopicExceptions**
> MsgVpnAclProfileSubscribeTopicExceptionsResponse GetMsgVpnAclProfileSubscribeTopicExceptions(ctx, msgVpnName, aclProfileName, optional)
Get a list of Subscribe Topic Exception objects.

Get a list of Subscribe Topic Exception objects.  A Subscribe Topic Exception is an exception to the default action to take when a client using the ACL Profile subscribes to a topic in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: aclProfileName|x||| msgVpnName|x||| subscribeTopicException|x||| subscribeTopicExceptionSyntax|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
 **optional** | ***AclProfileApiGetMsgVpnAclProfileSubscribeTopicExceptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AclProfileApiGetMsgVpnAclProfileSubscribeTopicExceptionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileSubscribeTopicExceptionsResponse**](MsgVpnAclProfileSubscribeTopicExceptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAclProfiles**
> MsgVpnAclProfilesResponse GetMsgVpnAclProfiles(ctx, msgVpnName, optional)
Get a list of ACL Profile objects.

Get a list of ACL Profile objects.  An ACL Profile controls whether an authenticated client is permitted to establish a connection with the message broker or permitted to publish and subscribe to specific topics.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: aclProfileName|x||| msgVpnName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***AclProfileApiGetMsgVpnAclProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AclProfileApiGetMsgVpnAclProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfilesResponse**](MsgVpnAclProfilesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnAclProfile**
> MsgVpnAclProfileResponse ReplaceMsgVpnAclProfile(ctx, body, msgVpnName, aclProfileName, optional)
Replace an ACL Profile object.

Replace an ACL Profile object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  An ACL Profile controls whether an authenticated client is permitted to establish a connection with the message broker or permitted to publish and subscribe to specific topics.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- aclProfileName|x||x||||| msgVpnName|x||x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnAclProfile**](MsgVpnAclProfile.md)| The ACL Profile object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
 **optional** | ***AclProfileApiReplaceMsgVpnAclProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AclProfileApiReplaceMsgVpnAclProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileResponse**](MsgVpnAclProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnAclProfile**
> MsgVpnAclProfileResponse UpdateMsgVpnAclProfile(ctx, body, msgVpnName, aclProfileName, optional)
Update an ACL Profile object.

Update an ACL Profile object. Any attribute missing from the request will be left unchanged.  An ACL Profile controls whether an authenticated client is permitted to establish a connection with the message broker or permitted to publish and subscribe to specific topics.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- aclProfileName|x|x||||| msgVpnName|x|x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnAclProfile**](MsgVpnAclProfile.md)| The ACL Profile object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
 **optional** | ***AclProfileApiUpdateMsgVpnAclProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AclProfileApiUpdateMsgVpnAclProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileResponse**](MsgVpnAclProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

