# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/monitor*

Method | HTTP request | Description
------------- | ------------- | -------------
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

# **GetMsgVpnAclProfile**
> MsgVpnAclProfileResponse GetMsgVpnAclProfile(ctx, msgVpnName, aclProfileName, optional)
Get an ACL Profile object.

Get an ACL Profile object.  An ACL Profile controls whether an authenticated client is permitted to establish a connection with the message broker or permitted to publish and subscribe to specific topics.   Attribute|Identifying|Deprecated :---|:---:|:---: aclProfileName|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

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

Get a Client Connect Exception object.  A Client Connect Exception is an exception to the default action to take when a client using the ACL Profile connects to the Message VPN. Exceptions must be expressed as an IP address/netmask in CIDR form.   Attribute|Identifying|Deprecated :---|:---:|:---: aclProfileName|x| clientConnectExceptionAddress|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

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

Get a list of Client Connect Exception objects.  A Client Connect Exception is an exception to the default action to take when a client using the ACL Profile connects to the Message VPN. Exceptions must be expressed as an IP address/netmask in CIDR form.   Attribute|Identifying|Deprecated :---|:---:|:---: aclProfileName|x| clientConnectExceptionAddress|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

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

Get a Publish Topic Exception object.  A Publish Topic Exception is an exception to the default action to take when a client using the ACL Profile publishes to a topic in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Deprecated :---|:---:|:---: aclProfileName|x|x msgVpnName|x|x publishExceptionTopic|x|x topicSyntax|x|x    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been deprecated since 2.14. Replaced by publishTopicExceptions.

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

Get a list of Publish Topic Exception objects.  A Publish Topic Exception is an exception to the default action to take when a client using the ACL Profile publishes to a topic in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Deprecated :---|:---:|:---: aclProfileName|x|x msgVpnName|x|x publishExceptionTopic|x|x topicSyntax|x|x    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been deprecated since 2.14. Replaced by publishTopicExceptions.

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

Get a Publish Topic Exception object.  A Publish Topic Exception is an exception to the default action to take when a client using the ACL Profile publishes to a topic in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Deprecated :---|:---:|:---: aclProfileName|x| msgVpnName|x| publishTopicException|x| publishTopicExceptionSyntax|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.14.

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

Get a list of Publish Topic Exception objects.  A Publish Topic Exception is an exception to the default action to take when a client using the ACL Profile publishes to a topic in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Deprecated :---|:---:|:---: aclProfileName|x| msgVpnName|x| publishTopicException|x| publishTopicExceptionSyntax|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.14.

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

Get a Subscribe Topic Exception object.  A Subscribe Topic Exception is an exception to the default action to take when a client using the ACL Profile subscribes to a topic in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Deprecated :---|:---:|:---: aclProfileName|x|x msgVpnName|x|x subscribeExceptionTopic|x|x topicSyntax|x|x    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been deprecated since 2.14. Replaced by subscribeTopicExceptions.

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

Get a list of Subscribe Topic Exception objects.  A Subscribe Topic Exception is an exception to the default action to take when a client using the ACL Profile subscribes to a topic in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Deprecated :---|:---:|:---: aclProfileName|x|x msgVpnName|x|x subscribeExceptionTopic|x|x topicSyntax|x|x    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been deprecated since 2.14. Replaced by subscribeTopicExceptions.

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

Get a Subscribe Share Name Exception object.  A Subscribe Share Name Exception is an exception to the default action to take when a client using the ACL Profile subscribes to a share-name subscription in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Deprecated :---|:---:|:---: aclProfileName|x| msgVpnName|x| subscribeShareNameException|x| subscribeShareNameExceptionSyntax|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.14.

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

Get a list of Subscribe Share Name Exception objects.  A Subscribe Share Name Exception is an exception to the default action to take when a client using the ACL Profile subscribes to a share-name subscription in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Deprecated :---|:---:|:---: aclProfileName|x| msgVpnName|x| subscribeShareNameException|x| subscribeShareNameExceptionSyntax|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.14.

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

Get a Subscribe Topic Exception object.  A Subscribe Topic Exception is an exception to the default action to take when a client using the ACL Profile subscribes to a topic in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Deprecated :---|:---:|:---: aclProfileName|x| msgVpnName|x| subscribeTopicException|x| subscribeTopicExceptionSyntax|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.14.

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

Get a list of Subscribe Topic Exception objects.  A Subscribe Topic Exception is an exception to the default action to take when a client using the ACL Profile subscribes to a topic in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Deprecated :---|:---:|:---: aclProfileName|x| msgVpnName|x| subscribeTopicException|x| subscribeTopicExceptionSyntax|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.14.

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

Get a list of ACL Profile objects.  An ACL Profile controls whether an authenticated client is permitted to establish a connection with the message broker or permitted to publish and subscribe to specific topics.   Attribute|Identifying|Deprecated :---|:---:|:---: aclProfileName|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

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

