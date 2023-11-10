# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/monitor*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMsgVpnRestDeliveryPoint**](RestDeliveryPointApi.md#GetMsgVpnRestDeliveryPoint) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName} | Get a REST Delivery Point object.
[**GetMsgVpnRestDeliveryPointQueueBinding**](RestDeliveryPointApi.md#GetMsgVpnRestDeliveryPointQueueBinding) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName} | Get a Queue Binding object.
[**GetMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader**](RestDeliveryPointApi.md#GetMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName}/protectedRequestHeaders/{headerName} | Get a Protected Request Header object.
[**GetMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeaders**](RestDeliveryPointApi.md#GetMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeaders) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName}/protectedRequestHeaders | Get a list of Protected Request Header objects.
[**GetMsgVpnRestDeliveryPointQueueBindingRequestHeader**](RestDeliveryPointApi.md#GetMsgVpnRestDeliveryPointQueueBindingRequestHeader) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName}/requestHeaders/{headerName} | Get a Request Header object.
[**GetMsgVpnRestDeliveryPointQueueBindingRequestHeaders**](RestDeliveryPointApi.md#GetMsgVpnRestDeliveryPointQueueBindingRequestHeaders) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName}/requestHeaders | Get a list of Request Header objects.
[**GetMsgVpnRestDeliveryPointQueueBindings**](RestDeliveryPointApi.md#GetMsgVpnRestDeliveryPointQueueBindings) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings | Get a list of Queue Binding objects.
[**GetMsgVpnRestDeliveryPointRestConsumer**](RestDeliveryPointApi.md#GetMsgVpnRestDeliveryPointRestConsumer) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName} | Get a REST Consumer object.
[**GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim**](RestDeliveryPointApi.md#GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/oauthJwtClaims/{oauthJwtClaimName} | Get a Claim object.
[**GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaims**](RestDeliveryPointApi.md#GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaims) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/oauthJwtClaims | Get a list of Claim objects.
[**GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName**](RestDeliveryPointApi.md#GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/tlsTrustedCommonNames/{tlsTrustedCommonName} | Get a Trusted Common Name object.
[**GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNames**](RestDeliveryPointApi.md#GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNames) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/tlsTrustedCommonNames | Get a list of Trusted Common Name objects.
[**GetMsgVpnRestDeliveryPointRestConsumers**](RestDeliveryPointApi.md#GetMsgVpnRestDeliveryPointRestConsumers) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers | Get a list of REST Consumer objects.
[**GetMsgVpnRestDeliveryPoints**](RestDeliveryPointApi.md#GetMsgVpnRestDeliveryPoints) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints | Get a list of REST Delivery Point objects.

# **GetMsgVpnRestDeliveryPoint**
> MsgVpnRestDeliveryPointResponse GetMsgVpnRestDeliveryPoint(ctx, msgVpnName, restDeliveryPointName, optional)
Get a REST Delivery Point object.

Get a REST Delivery Point object.  A REST Delivery Point manages delivery of messages from queues to a named list of REST Consumers.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| restDeliveryPointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
 **optional** | ***RestDeliveryPointApiGetMsgVpnRestDeliveryPointOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RestDeliveryPointApiGetMsgVpnRestDeliveryPointOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointResponse**](MsgVpnRestDeliveryPointResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnRestDeliveryPointQueueBinding**
> MsgVpnRestDeliveryPointQueueBindingResponse GetMsgVpnRestDeliveryPointQueueBinding(ctx, msgVpnName, restDeliveryPointName, queueBindingName, optional)
Get a Queue Binding object.

Get a Queue Binding object.  A Queue Binding for a REST Delivery Point attracts messages to be delivered to REST consumers. If the queue does not exist it can be created subsequently, and once the queue is operational the broker performs the queue binding. Removing the queue binding does not delete the queue itself. Similarly, removing the queue does not remove the queue binding, which fails until the queue is recreated or the queue binding is deleted.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| queueBindingName|x| restDeliveryPointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **queueBindingName** | **string**| The name of a queue in the Message VPN. | 
 **optional** | ***RestDeliveryPointApiGetMsgVpnRestDeliveryPointQueueBindingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RestDeliveryPointApiGetMsgVpnRestDeliveryPointQueueBindingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointQueueBindingResponse**](MsgVpnRestDeliveryPointQueueBindingResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader**
> MsgVpnRestDeliveryPointQueueBindingProtectedRequestHeaderResponse GetMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader(ctx, msgVpnName, restDeliveryPointName, queueBindingName, headerName, optional)
Get a Protected Request Header object.

Get a Protected Request Header object.  A protected request header to be added to the HTTP request. Unlike a non-protected request header, the header value cannot be displayed after it is set.   Attribute|Identifying|Deprecated :---|:---:|:---: headerName|x| msgVpnName|x| queueBindingName|x| restDeliveryPointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.30.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **queueBindingName** | **string**| The name of a queue in the Message VPN. | 
  **headerName** | **string**| The name of the protected HTTP request header. | 
 **optional** | ***RestDeliveryPointApiGetMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeaderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RestDeliveryPointApiGetMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeaderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointQueueBindingProtectedRequestHeaderResponse**](MsgVpnRestDeliveryPointQueueBindingProtectedRequestHeaderResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeaders**
> MsgVpnRestDeliveryPointQueueBindingProtectedRequestHeadersResponse GetMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeaders(ctx, msgVpnName, restDeliveryPointName, queueBindingName, optional)
Get a list of Protected Request Header objects.

Get a list of Protected Request Header objects.  A protected request header to be added to the HTTP request. Unlike a non-protected request header, the header value cannot be displayed after it is set.   Attribute|Identifying|Deprecated :---|:---:|:---: headerName|x| msgVpnName|x| queueBindingName|x| restDeliveryPointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.30.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **queueBindingName** | **string**| The name of a queue in the Message VPN. | 
 **optional** | ***RestDeliveryPointApiGetMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeadersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RestDeliveryPointApiGetMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeadersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointQueueBindingProtectedRequestHeadersResponse**](MsgVpnRestDeliveryPointQueueBindingProtectedRequestHeadersResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnRestDeliveryPointQueueBindingRequestHeader**
> MsgVpnRestDeliveryPointQueueBindingRequestHeaderResponse GetMsgVpnRestDeliveryPointQueueBindingRequestHeader(ctx, msgVpnName, restDeliveryPointName, queueBindingName, headerName, optional)
Get a Request Header object.

Get a Request Header object.  A request header to be added to the HTTP request.   Attribute|Identifying|Deprecated :---|:---:|:---: headerName|x| msgVpnName|x| queueBindingName|x| restDeliveryPointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.23.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **queueBindingName** | **string**| The name of a queue in the Message VPN. | 
  **headerName** | **string**| The name of the HTTP request header. | 
 **optional** | ***RestDeliveryPointApiGetMsgVpnRestDeliveryPointQueueBindingRequestHeaderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RestDeliveryPointApiGetMsgVpnRestDeliveryPointQueueBindingRequestHeaderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointQueueBindingRequestHeaderResponse**](MsgVpnRestDeliveryPointQueueBindingRequestHeaderResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnRestDeliveryPointQueueBindingRequestHeaders**
> MsgVpnRestDeliveryPointQueueBindingRequestHeadersResponse GetMsgVpnRestDeliveryPointQueueBindingRequestHeaders(ctx, msgVpnName, restDeliveryPointName, queueBindingName, optional)
Get a list of Request Header objects.

Get a list of Request Header objects.  A request header to be added to the HTTP request.   Attribute|Identifying|Deprecated :---|:---:|:---: headerName|x| msgVpnName|x| queueBindingName|x| restDeliveryPointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.23.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **queueBindingName** | **string**| The name of a queue in the Message VPN. | 
 **optional** | ***RestDeliveryPointApiGetMsgVpnRestDeliveryPointQueueBindingRequestHeadersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RestDeliveryPointApiGetMsgVpnRestDeliveryPointQueueBindingRequestHeadersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointQueueBindingRequestHeadersResponse**](MsgVpnRestDeliveryPointQueueBindingRequestHeadersResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnRestDeliveryPointQueueBindings**
> MsgVpnRestDeliveryPointQueueBindingsResponse GetMsgVpnRestDeliveryPointQueueBindings(ctx, msgVpnName, restDeliveryPointName, optional)
Get a list of Queue Binding objects.

Get a list of Queue Binding objects.  A Queue Binding for a REST Delivery Point attracts messages to be delivered to REST consumers. If the queue does not exist it can be created subsequently, and once the queue is operational the broker performs the queue binding. Removing the queue binding does not delete the queue itself. Similarly, removing the queue does not remove the queue binding, which fails until the queue is recreated or the queue binding is deleted.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| queueBindingName|x| restDeliveryPointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
 **optional** | ***RestDeliveryPointApiGetMsgVpnRestDeliveryPointQueueBindingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RestDeliveryPointApiGetMsgVpnRestDeliveryPointQueueBindingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointQueueBindingsResponse**](MsgVpnRestDeliveryPointQueueBindingsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnRestDeliveryPointRestConsumer**
> MsgVpnRestDeliveryPointRestConsumerResponse GetMsgVpnRestDeliveryPointRestConsumer(ctx, msgVpnName, restDeliveryPointName, restConsumerName, optional)
Get a REST Consumer object.

Get a REST Consumer object.  REST Consumer objects establish HTTP connectivity to REST consumer applications who wish to receive messages from a broker.   Attribute|Identifying|Deprecated :---|:---:|:---: counter.httpRequestConnectionCloseTxMsgCount||x counter.httpRequestOutstandingTxMsgCount||x counter.httpRequestTimedOutTxMsgCount||x counter.httpRequestTxByteCount||x counter.httpRequestTxMsgCount||x counter.httpResponseErrorRxMsgCount||x counter.httpResponseRxByteCount||x counter.httpResponseRxMsgCount||x counter.httpResponseSuccessRxMsgCount||x msgVpnName|x| restConsumerName|x| restDeliveryPointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **restConsumerName** | **string**| The name of the REST Consumer. | 
 **optional** | ***RestDeliveryPointApiGetMsgVpnRestDeliveryPointRestConsumerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RestDeliveryPointApiGetMsgVpnRestDeliveryPointRestConsumerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointRestConsumerResponse**](MsgVpnRestDeliveryPointRestConsumerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim**
> MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimResponse GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim(ctx, msgVpnName, restDeliveryPointName, restConsumerName, oauthJwtClaimName, optional)
Get a Claim object.

Get a Claim object.  A Claim is added to the JWT sent to the OAuth token request endpoint.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| oauthJwtClaimName|x| restConsumerName|x| restDeliveryPointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.21.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **restConsumerName** | **string**| The name of the REST Consumer. | 
  **oauthJwtClaimName** | **string**| The name of the additional claim. Cannot be \&quot;exp\&quot;, \&quot;iat\&quot;, or \&quot;jti\&quot;. | 
 **optional** | ***RestDeliveryPointApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RestDeliveryPointApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimResponse**](MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaims**
> MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsResponse GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaims(ctx, msgVpnName, restDeliveryPointName, restConsumerName, optional)
Get a list of Claim objects.

Get a list of Claim objects.  A Claim is added to the JWT sent to the OAuth token request endpoint.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| oauthJwtClaimName|x| restConsumerName|x| restDeliveryPointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.21.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **restConsumerName** | **string**| The name of the REST Consumer. | 
 **optional** | ***RestDeliveryPointApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RestDeliveryPointApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsResponse**](MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName**
> MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameResponse GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName(ctx, msgVpnName, restDeliveryPointName, restConsumerName, tlsTrustedCommonName, optional)
Get a Trusted Common Name object.

Get a Trusted Common Name object.  The Trusted Common Names for the REST Consumer are used by encrypted transports to verify the name in the certificate presented by the remote REST consumer. They must include the common name of the remote REST consumer's server certificate.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x|x restConsumerName|x|x restDeliveryPointName|x|x tlsTrustedCommonName|x|x    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been deprecated since 2.17. Common Name validation has been replaced by Server Certificate Name validation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **restConsumerName** | **string**| The name of the REST Consumer. | 
  **tlsTrustedCommonName** | **string**| The expected trusted common name of the remote certificate. | 
 **optional** | ***RestDeliveryPointApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RestDeliveryPointApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameResponse**](MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNames**
> MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesResponse GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNames(ctx, msgVpnName, restDeliveryPointName, restConsumerName, optional)
Get a list of Trusted Common Name objects.

Get a list of Trusted Common Name objects.  The Trusted Common Names for the REST Consumer are used by encrypted transports to verify the name in the certificate presented by the remote REST consumer. They must include the common name of the remote REST consumer's server certificate.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x|x restConsumerName|x|x restDeliveryPointName|x|x tlsTrustedCommonName|x|x    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been deprecated since 2.17. Common Name validation has been replaced by Server Certificate Name validation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **restConsumerName** | **string**| The name of the REST Consumer. | 
 **optional** | ***RestDeliveryPointApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RestDeliveryPointApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesResponse**](MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnRestDeliveryPointRestConsumers**
> MsgVpnRestDeliveryPointRestConsumersResponse GetMsgVpnRestDeliveryPointRestConsumers(ctx, msgVpnName, restDeliveryPointName, optional)
Get a list of REST Consumer objects.

Get a list of REST Consumer objects.  REST Consumer objects establish HTTP connectivity to REST consumer applications who wish to receive messages from a broker.   Attribute|Identifying|Deprecated :---|:---:|:---: counter.httpRequestConnectionCloseTxMsgCount||x counter.httpRequestOutstandingTxMsgCount||x counter.httpRequestTimedOutTxMsgCount||x counter.httpRequestTxByteCount||x counter.httpRequestTxMsgCount||x counter.httpResponseErrorRxMsgCount||x counter.httpResponseRxByteCount||x counter.httpResponseRxMsgCount||x counter.httpResponseSuccessRxMsgCount||x msgVpnName|x| restConsumerName|x| restDeliveryPointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
 **optional** | ***RestDeliveryPointApiGetMsgVpnRestDeliveryPointRestConsumersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RestDeliveryPointApiGetMsgVpnRestDeliveryPointRestConsumersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointRestConsumersResponse**](MsgVpnRestDeliveryPointRestConsumersResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnRestDeliveryPoints**
> MsgVpnRestDeliveryPointsResponse GetMsgVpnRestDeliveryPoints(ctx, msgVpnName, optional)
Get a list of REST Delivery Point objects.

Get a list of REST Delivery Point objects.  A REST Delivery Point manages delivery of messages from queues to a named list of REST Consumers.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| restDeliveryPointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***RestDeliveryPointApiGetMsgVpnRestDeliveryPointsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RestDeliveryPointApiGetMsgVpnRestDeliveryPointsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointsResponse**](MsgVpnRestDeliveryPointsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

