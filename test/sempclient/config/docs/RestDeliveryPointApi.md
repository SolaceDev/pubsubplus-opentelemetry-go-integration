# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/config*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMsgVpnRestDeliveryPoint**](RestDeliveryPointApi.md#CreateMsgVpnRestDeliveryPoint) | **Post** /msgVpns/{msgVpnName}/restDeliveryPoints | Create a REST Delivery Point object.
[**CreateMsgVpnRestDeliveryPointQueueBinding**](RestDeliveryPointApi.md#CreateMsgVpnRestDeliveryPointQueueBinding) | **Post** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings | Create a Queue Binding object.
[**CreateMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader**](RestDeliveryPointApi.md#CreateMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader) | **Post** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName}/protectedRequestHeaders | Create a Protected Request Header object.
[**CreateMsgVpnRestDeliveryPointQueueBindingRequestHeader**](RestDeliveryPointApi.md#CreateMsgVpnRestDeliveryPointQueueBindingRequestHeader) | **Post** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName}/requestHeaders | Create a Request Header object.
[**CreateMsgVpnRestDeliveryPointRestConsumer**](RestDeliveryPointApi.md#CreateMsgVpnRestDeliveryPointRestConsumer) | **Post** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers | Create a REST Consumer object.
[**CreateMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim**](RestDeliveryPointApi.md#CreateMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim) | **Post** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/oauthJwtClaims | Create a Claim object.
[**CreateMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName**](RestDeliveryPointApi.md#CreateMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName) | **Post** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/tlsTrustedCommonNames | Create a Trusted Common Name object.
[**DeleteMsgVpnRestDeliveryPoint**](RestDeliveryPointApi.md#DeleteMsgVpnRestDeliveryPoint) | **Delete** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName} | Delete a REST Delivery Point object.
[**DeleteMsgVpnRestDeliveryPointQueueBinding**](RestDeliveryPointApi.md#DeleteMsgVpnRestDeliveryPointQueueBinding) | **Delete** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName} | Delete a Queue Binding object.
[**DeleteMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader**](RestDeliveryPointApi.md#DeleteMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader) | **Delete** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName}/protectedRequestHeaders/{headerName} | Delete a Protected Request Header object.
[**DeleteMsgVpnRestDeliveryPointQueueBindingRequestHeader**](RestDeliveryPointApi.md#DeleteMsgVpnRestDeliveryPointQueueBindingRequestHeader) | **Delete** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName}/requestHeaders/{headerName} | Delete a Request Header object.
[**DeleteMsgVpnRestDeliveryPointRestConsumer**](RestDeliveryPointApi.md#DeleteMsgVpnRestDeliveryPointRestConsumer) | **Delete** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName} | Delete a REST Consumer object.
[**DeleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim**](RestDeliveryPointApi.md#DeleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim) | **Delete** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/oauthJwtClaims/{oauthJwtClaimName} | Delete a Claim object.
[**DeleteMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName**](RestDeliveryPointApi.md#DeleteMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName) | **Delete** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/tlsTrustedCommonNames/{tlsTrustedCommonName} | Delete a Trusted Common Name object.
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
[**ReplaceMsgVpnRestDeliveryPoint**](RestDeliveryPointApi.md#ReplaceMsgVpnRestDeliveryPoint) | **Put** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName} | Replace a REST Delivery Point object.
[**ReplaceMsgVpnRestDeliveryPointQueueBinding**](RestDeliveryPointApi.md#ReplaceMsgVpnRestDeliveryPointQueueBinding) | **Put** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName} | Replace a Queue Binding object.
[**ReplaceMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader**](RestDeliveryPointApi.md#ReplaceMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader) | **Put** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName}/protectedRequestHeaders/{headerName} | Replace a Protected Request Header object.
[**ReplaceMsgVpnRestDeliveryPointQueueBindingRequestHeader**](RestDeliveryPointApi.md#ReplaceMsgVpnRestDeliveryPointQueueBindingRequestHeader) | **Put** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName}/requestHeaders/{headerName} | Replace a Request Header object.
[**ReplaceMsgVpnRestDeliveryPointRestConsumer**](RestDeliveryPointApi.md#ReplaceMsgVpnRestDeliveryPointRestConsumer) | **Put** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName} | Replace a REST Consumer object.
[**UpdateMsgVpnRestDeliveryPoint**](RestDeliveryPointApi.md#UpdateMsgVpnRestDeliveryPoint) | **Patch** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName} | Update a REST Delivery Point object.
[**UpdateMsgVpnRestDeliveryPointQueueBinding**](RestDeliveryPointApi.md#UpdateMsgVpnRestDeliveryPointQueueBinding) | **Patch** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName} | Update a Queue Binding object.
[**UpdateMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader**](RestDeliveryPointApi.md#UpdateMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader) | **Patch** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName}/protectedRequestHeaders/{headerName} | Update a Protected Request Header object.
[**UpdateMsgVpnRestDeliveryPointQueueBindingRequestHeader**](RestDeliveryPointApi.md#UpdateMsgVpnRestDeliveryPointQueueBindingRequestHeader) | **Patch** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName}/requestHeaders/{headerName} | Update a Request Header object.
[**UpdateMsgVpnRestDeliveryPointRestConsumer**](RestDeliveryPointApi.md#UpdateMsgVpnRestDeliveryPointRestConsumer) | **Patch** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName} | Update a REST Consumer object.

# **CreateMsgVpnRestDeliveryPoint**
> MsgVpnRestDeliveryPointResponse CreateMsgVpnRestDeliveryPoint(ctx, body, msgVpnName, optional)
Create a REST Delivery Point object.

Create a REST Delivery Point object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A REST Delivery Point manages delivery of messages from queues to a named list of REST Consumers.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x||| restDeliveryPointName|x|x||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnRestDeliveryPoint**](MsgVpnRestDeliveryPoint.md)| The REST Delivery Point object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***RestDeliveryPointApiCreateMsgVpnRestDeliveryPointOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RestDeliveryPointApiCreateMsgVpnRestDeliveryPointOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointResponse**](MsgVpnRestDeliveryPointResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnRestDeliveryPointQueueBinding**
> MsgVpnRestDeliveryPointQueueBindingResponse CreateMsgVpnRestDeliveryPointQueueBinding(ctx, body, msgVpnName, restDeliveryPointName, optional)
Create a Queue Binding object.

Create a Queue Binding object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Queue Binding for a REST Delivery Point attracts messages to be delivered to REST consumers. If the queue does not exist it can be created subsequently, and once the queue is operational the broker performs the queue binding. Removing the queue binding does not delete the queue itself. Similarly, removing the queue does not remove the queue binding, which fails until the queue is recreated or the queue binding is deleted.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x||| queueBindingName|x|x|||| restDeliveryPointName|x||x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnRestDeliveryPointQueueBinding**](MsgVpnRestDeliveryPointQueueBinding.md)| The Queue Binding object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
 **optional** | ***RestDeliveryPointApiCreateMsgVpnRestDeliveryPointQueueBindingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RestDeliveryPointApiCreateMsgVpnRestDeliveryPointQueueBindingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointQueueBindingResponse**](MsgVpnRestDeliveryPointQueueBindingResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader**
> MsgVpnRestDeliveryPointQueueBindingProtectedRequestHeaderResponse CreateMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader(ctx, body, msgVpnName, restDeliveryPointName, queueBindingName, optional)
Create a Protected Request Header object.

Create a Protected Request Header object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A protected request header to be added to the HTTP request. Unlike a non-protected request header, the header value cannot be displayed after it is set.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: headerName|x|x|||| headerValue||||x||x msgVpnName|x||x||| queueBindingName|x||x||| restDeliveryPointName|x||x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.30.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader**](MsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader.md)| The Protected Request Header object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **queueBindingName** | **string**| The name of a queue in the Message VPN. | 
 **optional** | ***RestDeliveryPointApiCreateMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeaderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RestDeliveryPointApiCreateMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeaderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointQueueBindingProtectedRequestHeaderResponse**](MsgVpnRestDeliveryPointQueueBindingProtectedRequestHeaderResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnRestDeliveryPointQueueBindingRequestHeader**
> MsgVpnRestDeliveryPointQueueBindingRequestHeaderResponse CreateMsgVpnRestDeliveryPointQueueBindingRequestHeader(ctx, body, msgVpnName, restDeliveryPointName, queueBindingName, optional)
Create a Request Header object.

Create a Request Header object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A request header to be added to the HTTP request.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: headerName|x|x|||| msgVpnName|x||x||| queueBindingName|x||x||| restDeliveryPointName|x||x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.23.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnRestDeliveryPointQueueBindingRequestHeader**](MsgVpnRestDeliveryPointQueueBindingRequestHeader.md)| The Request Header object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **queueBindingName** | **string**| The name of a queue in the Message VPN. | 
 **optional** | ***RestDeliveryPointApiCreateMsgVpnRestDeliveryPointQueueBindingRequestHeaderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RestDeliveryPointApiCreateMsgVpnRestDeliveryPointQueueBindingRequestHeaderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointQueueBindingRequestHeaderResponse**](MsgVpnRestDeliveryPointQueueBindingRequestHeaderResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnRestDeliveryPointRestConsumer**
> MsgVpnRestDeliveryPointRestConsumerResponse CreateMsgVpnRestDeliveryPointRestConsumer(ctx, body, msgVpnName, restDeliveryPointName, optional)
Create a REST Consumer object.

Create a REST Consumer object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  REST Consumer objects establish HTTP connectivity to REST consumer applications who wish to receive messages from a broker.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: authenticationAwsSecretAccessKey||||x||x authenticationClientCertContent||||x||x authenticationClientCertPassword||||x|| authenticationHttpBasicPassword||||x||x authenticationHttpHeaderValue||||x||x authenticationOauthClientSecret||||x||x authenticationOauthJwtSecretKey||||x||x msgVpnName|x||x||| restConsumerName|x|x|||| restDeliveryPointName|x||x|||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- MsgVpnRestDeliveryPointRestConsumer|authenticationClientCertPassword|authenticationClientCertContent| MsgVpnRestDeliveryPointRestConsumer|authenticationHttpBasicPassword|authenticationHttpBasicUsername| MsgVpnRestDeliveryPointRestConsumer|authenticationHttpBasicUsername|authenticationHttpBasicPassword| MsgVpnRestDeliveryPointRestConsumer|remotePort|tlsEnabled| MsgVpnRestDeliveryPointRestConsumer|tlsEnabled|remotePort|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnRestDeliveryPointRestConsumer**](MsgVpnRestDeliveryPointRestConsumer.md)| The REST Consumer object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
 **optional** | ***RestDeliveryPointApiCreateMsgVpnRestDeliveryPointRestConsumerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RestDeliveryPointApiCreateMsgVpnRestDeliveryPointRestConsumerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointRestConsumerResponse**](MsgVpnRestDeliveryPointRestConsumerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim**
> MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimResponse CreateMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim(ctx, body, msgVpnName, restDeliveryPointName, restConsumerName, optional)
Create a Claim object.

Create a Claim object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Claim is added to the JWT sent to the OAuth token request endpoint.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x||| oauthJwtClaimName|x|x|||| oauthJwtClaimValue||x|||| restConsumerName|x||x||| restDeliveryPointName|x||x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.21.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnRestDeliveryPointRestConsumerOauthJwtClaim**](MsgVpnRestDeliveryPointRestConsumerOauthJwtClaim.md)| The Claim object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **restConsumerName** | **string**| The name of the REST Consumer. | 
 **optional** | ***RestDeliveryPointApiCreateMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RestDeliveryPointApiCreateMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimResponse**](MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName**
> MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameResponse CreateMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName(ctx, body, msgVpnName, restDeliveryPointName, restConsumerName, optional)
Create a Trusted Common Name object.

Create a Trusted Common Name object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  The Trusted Common Names for the REST Consumer are used by encrypted transports to verify the name in the certificate presented by the remote REST consumer. They must include the common name of the remote REST consumer's server certificate.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x||x| restConsumerName|x||x||x| restDeliveryPointName|x||x||x| tlsTrustedCommonName|x|x|||x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been deprecated since (will be deprecated in next SEMP version). Common Name validation has been replaced by Server Certificate Name validation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName**](MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName.md)| The Trusted Common Name object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **restConsumerName** | **string**| The name of the REST Consumer. | 
 **optional** | ***RestDeliveryPointApiCreateMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RestDeliveryPointApiCreateMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameResponse**](MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnRestDeliveryPoint**
> SempMetaOnlyResponse DeleteMsgVpnRestDeliveryPoint(ctx, msgVpnName, restDeliveryPointName)
Delete a REST Delivery Point object.

Delete a REST Delivery Point object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A REST Delivery Point manages delivery of messages from queues to a named list of REST Consumers.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnRestDeliveryPointQueueBinding**
> SempMetaOnlyResponse DeleteMsgVpnRestDeliveryPointQueueBinding(ctx, msgVpnName, restDeliveryPointName, queueBindingName)
Delete a Queue Binding object.

Delete a Queue Binding object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Queue Binding for a REST Delivery Point attracts messages to be delivered to REST consumers. If the queue does not exist it can be created subsequently, and once the queue is operational the broker performs the queue binding. Removing the queue binding does not delete the queue itself. Similarly, removing the queue does not remove the queue binding, which fails until the queue is recreated or the queue binding is deleted.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **queueBindingName** | **string**| The name of a queue in the Message VPN. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader**
> SempMetaOnlyResponse DeleteMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader(ctx, msgVpnName, restDeliveryPointName, queueBindingName, headerName)
Delete a Protected Request Header object.

Delete a Protected Request Header object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A protected request header to be added to the HTTP request. Unlike a non-protected request header, the header value cannot be displayed after it is set.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.30.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **queueBindingName** | **string**| The name of a queue in the Message VPN. | 
  **headerName** | **string**| The name of the protected HTTP request header. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnRestDeliveryPointQueueBindingRequestHeader**
> SempMetaOnlyResponse DeleteMsgVpnRestDeliveryPointQueueBindingRequestHeader(ctx, msgVpnName, restDeliveryPointName, queueBindingName, headerName)
Delete a Request Header object.

Delete a Request Header object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A request header to be added to the HTTP request.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.23.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **queueBindingName** | **string**| The name of a queue in the Message VPN. | 
  **headerName** | **string**| The name of the HTTP request header. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnRestDeliveryPointRestConsumer**
> SempMetaOnlyResponse DeleteMsgVpnRestDeliveryPointRestConsumer(ctx, msgVpnName, restDeliveryPointName, restConsumerName)
Delete a REST Consumer object.

Delete a REST Consumer object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  REST Consumer objects establish HTTP connectivity to REST consumer applications who wish to receive messages from a broker.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **restConsumerName** | **string**| The name of the REST Consumer. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim**
> SempMetaOnlyResponse DeleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim(ctx, msgVpnName, restDeliveryPointName, restConsumerName, oauthJwtClaimName)
Delete a Claim object.

Delete a Claim object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Claim is added to the JWT sent to the OAuth token request endpoint.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.21.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **restConsumerName** | **string**| The name of the REST Consumer. | 
  **oauthJwtClaimName** | **string**| The name of the additional claim. Cannot be \&quot;exp\&quot;, \&quot;iat\&quot;, or \&quot;jti\&quot;. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName**
> SempMetaOnlyResponse DeleteMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName(ctx, msgVpnName, restDeliveryPointName, restConsumerName, tlsTrustedCommonName)
Delete a Trusted Common Name object.

Delete a Trusted Common Name object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  The Trusted Common Names for the REST Consumer are used by encrypted transports to verify the name in the certificate presented by the remote REST consumer. They must include the common name of the remote REST consumer's server certificate.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been deprecated since (will be deprecated in next SEMP version). Common Name validation has been replaced by Server Certificate Name validation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **restConsumerName** | **string**| The name of the REST Consumer. | 
  **tlsTrustedCommonName** | **string**| The expected trusted common name of the remote certificate. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnRestDeliveryPoint**
> MsgVpnRestDeliveryPointResponse GetMsgVpnRestDeliveryPoint(ctx, msgVpnName, restDeliveryPointName, optional)
Get a REST Delivery Point object.

Get a REST Delivery Point object.  A REST Delivery Point manages delivery of messages from queues to a named list of REST Consumers.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| restDeliveryPointName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.0.

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


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a Queue Binding object.  A Queue Binding for a REST Delivery Point attracts messages to be delivered to REST consumers. If the queue does not exist it can be created subsequently, and once the queue is operational the broker performs the queue binding. Removing the queue binding does not delete the queue itself. Similarly, removing the queue does not remove the queue binding, which fails until the queue is recreated or the queue binding is deleted.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| queueBindingName|x||| restDeliveryPointName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.0.

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



 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a Protected Request Header object.  A protected request header to be added to the HTTP request. Unlike a non-protected request header, the header value cannot be displayed after it is set.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: headerName|x||| headerValue||x||x msgVpnName|x||| queueBindingName|x||| restDeliveryPointName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.30.

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




 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Protected Request Header objects.  A protected request header to be added to the HTTP request. Unlike a non-protected request header, the header value cannot be displayed after it is set.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: headerName|x||| headerValue||x||x msgVpnName|x||| queueBindingName|x||| restDeliveryPointName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.30.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a Request Header object.  A request header to be added to the HTTP request.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: headerName|x||| msgVpnName|x||| queueBindingName|x||| restDeliveryPointName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.23.

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




 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Request Header objects.  A request header to be added to the HTTP request.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: headerName|x||| msgVpnName|x||| queueBindingName|x||| restDeliveryPointName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.23.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Queue Binding objects.  A Queue Binding for a REST Delivery Point attracts messages to be delivered to REST consumers. If the queue does not exist it can be created subsequently, and once the queue is operational the broker performs the queue binding. Removing the queue binding does not delete the queue itself. Similarly, removing the queue does not remove the queue binding, which fails until the queue is recreated or the queue binding is deleted.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| queueBindingName|x||| restDeliveryPointName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.0.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a REST Consumer object.  REST Consumer objects establish HTTP connectivity to REST consumer applications who wish to receive messages from a broker.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: authenticationAwsSecretAccessKey||x||x authenticationClientCertContent||x||x authenticationClientCertPassword||x|| authenticationHttpBasicPassword||x||x authenticationHttpHeaderValue||x||x authenticationOauthClientSecret||x||x authenticationOauthJwtSecretKey||x||x msgVpnName|x||| restConsumerName|x||| restDeliveryPointName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.0.

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



 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a Claim object.  A Claim is added to the JWT sent to the OAuth token request endpoint.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| oauthJwtClaimName|x||| restConsumerName|x||| restDeliveryPointName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.21.

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




 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Claim objects.  A Claim is added to the JWT sent to the OAuth token request endpoint.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| oauthJwtClaimName|x||| restConsumerName|x||| restDeliveryPointName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.21.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a Trusted Common Name object.  The Trusted Common Names for the REST Consumer are used by encrypted transports to verify the name in the certificate presented by the remote REST consumer. They must include the common name of the remote REST consumer's server certificate.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||x| restConsumerName|x||x| restDeliveryPointName|x||x| tlsTrustedCommonName|x||x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been deprecated since (will be deprecated in next SEMP version). Common Name validation has been replaced by Server Certificate Name validation.

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




 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Trusted Common Name objects.  The Trusted Common Names for the REST Consumer are used by encrypted transports to verify the name in the certificate presented by the remote REST consumer. They must include the common name of the remote REST consumer's server certificate.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||x| restConsumerName|x||x| restDeliveryPointName|x||x| tlsTrustedCommonName|x||x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been deprecated since (will be deprecated in next SEMP version). Common Name validation has been replaced by Server Certificate Name validation.

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



 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of REST Consumer objects.  REST Consumer objects establish HTTP connectivity to REST consumer applications who wish to receive messages from a broker.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: authenticationAwsSecretAccessKey||x||x authenticationClientCertContent||x||x authenticationClientCertPassword||x|| authenticationHttpBasicPassword||x||x authenticationHttpHeaderValue||x||x authenticationOauthClientSecret||x||x authenticationOauthJwtSecretKey||x||x msgVpnName|x||| restConsumerName|x||| restDeliveryPointName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.0.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of REST Delivery Point objects.  A REST Delivery Point manages delivery of messages from queues to a named list of REST Consumers.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| restDeliveryPointName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.0.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

# **ReplaceMsgVpnRestDeliveryPoint**
> MsgVpnRestDeliveryPointResponse ReplaceMsgVpnRestDeliveryPoint(ctx, body, msgVpnName, restDeliveryPointName, optional)
Replace a REST Delivery Point object.

Replace a REST Delivery Point object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  A REST Delivery Point manages delivery of messages from queues to a named list of REST Consumers.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- clientProfileName||||||x|| msgVpnName|x||x||||| restDeliveryPointName|x||x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnRestDeliveryPoint**](MsgVpnRestDeliveryPoint.md)| The REST Delivery Point object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
 **optional** | ***RestDeliveryPointApiReplaceMsgVpnRestDeliveryPointOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RestDeliveryPointApiReplaceMsgVpnRestDeliveryPointOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointResponse**](MsgVpnRestDeliveryPointResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnRestDeliveryPointQueueBinding**
> MsgVpnRestDeliveryPointQueueBindingResponse ReplaceMsgVpnRestDeliveryPointQueueBinding(ctx, body, msgVpnName, restDeliveryPointName, queueBindingName, optional)
Replace a Queue Binding object.

Replace a Queue Binding object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  A Queue Binding for a REST Delivery Point attracts messages to be delivered to REST consumers. If the queue does not exist it can be created subsequently, and once the queue is operational the broker performs the queue binding. Removing the queue binding does not delete the queue itself. Similarly, removing the queue does not remove the queue binding, which fails until the queue is recreated or the queue binding is deleted.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- msgVpnName|x||x||||| queueBindingName|x||x||||| restDeliveryPointName|x||x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnRestDeliveryPointQueueBinding**](MsgVpnRestDeliveryPointQueueBinding.md)| The Queue Binding object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **queueBindingName** | **string**| The name of a queue in the Message VPN. | 
 **optional** | ***RestDeliveryPointApiReplaceMsgVpnRestDeliveryPointQueueBindingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RestDeliveryPointApiReplaceMsgVpnRestDeliveryPointQueueBindingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointQueueBindingResponse**](MsgVpnRestDeliveryPointQueueBindingResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader**
> MsgVpnRestDeliveryPointQueueBindingProtectedRequestHeaderResponse ReplaceMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader(ctx, body, msgVpnName, restDeliveryPointName, queueBindingName, headerName, optional)
Replace a Protected Request Header object.

Replace a Protected Request Header object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  A protected request header to be added to the HTTP request. Unlike a non-protected request header, the header value cannot be displayed after it is set.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- headerName|x||x||||| headerValue||||x||||x msgVpnName|x||x||||| queueBindingName|x||x||||| restDeliveryPointName|x||x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.30.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader**](MsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader.md)| The Protected Request Header object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **queueBindingName** | **string**| The name of a queue in the Message VPN. | 
  **headerName** | **string**| The name of the protected HTTP request header. | 
 **optional** | ***RestDeliveryPointApiReplaceMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeaderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RestDeliveryPointApiReplaceMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeaderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointQueueBindingProtectedRequestHeaderResponse**](MsgVpnRestDeliveryPointQueueBindingProtectedRequestHeaderResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnRestDeliveryPointQueueBindingRequestHeader**
> MsgVpnRestDeliveryPointQueueBindingRequestHeaderResponse ReplaceMsgVpnRestDeliveryPointQueueBindingRequestHeader(ctx, body, msgVpnName, restDeliveryPointName, queueBindingName, headerName, optional)
Replace a Request Header object.

Replace a Request Header object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  A request header to be added to the HTTP request.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- headerName|x||x||||| msgVpnName|x||x||||| queueBindingName|x||x||||| restDeliveryPointName|x||x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.23.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnRestDeliveryPointQueueBindingRequestHeader**](MsgVpnRestDeliveryPointQueueBindingRequestHeader.md)| The Request Header object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **queueBindingName** | **string**| The name of a queue in the Message VPN. | 
  **headerName** | **string**| The name of the HTTP request header. | 
 **optional** | ***RestDeliveryPointApiReplaceMsgVpnRestDeliveryPointQueueBindingRequestHeaderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RestDeliveryPointApiReplaceMsgVpnRestDeliveryPointQueueBindingRequestHeaderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointQueueBindingRequestHeaderResponse**](MsgVpnRestDeliveryPointQueueBindingRequestHeaderResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnRestDeliveryPointRestConsumer**
> MsgVpnRestDeliveryPointRestConsumerResponse ReplaceMsgVpnRestDeliveryPointRestConsumer(ctx, body, msgVpnName, restDeliveryPointName, restConsumerName, optional)
Replace a REST Consumer object.

Replace a REST Consumer object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  REST Consumer objects establish HTTP connectivity to REST consumer applications who wish to receive messages from a broker.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- authenticationAwsSecretAccessKey||||x||||x authenticationClientCertContent||||x||x||x authenticationClientCertPassword||||x||x|| authenticationHttpBasicPassword||||x||x||x authenticationHttpBasicUsername||||||x|| authenticationHttpHeaderValue||||x||||x authenticationOauthClientId||||||x|| authenticationOauthClientScope||||||x|| authenticationOauthClientSecret||||x||x||x authenticationOauthClientTokenEndpoint||||||x|| authenticationOauthClientTokenExpiryDefault||||||x|| authenticationOauthJwtSecretKey||||x||x||x authenticationOauthJwtTokenEndpoint||||||x|| authenticationOauthJwtTokenExpiryDefault||||||x|| authenticationScheme||||||x|| msgVpnName|x||x||||| outgoingConnectionCount||||||x|| remoteHost||||||x|| remotePort||||||x|| restConsumerName|x||x||||| restDeliveryPointName|x||x||||| tlsCipherSuiteList||||||x|| tlsEnabled||||||x||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- MsgVpnRestDeliveryPointRestConsumer|authenticationClientCertPassword|authenticationClientCertContent| MsgVpnRestDeliveryPointRestConsumer|authenticationHttpBasicPassword|authenticationHttpBasicUsername| MsgVpnRestDeliveryPointRestConsumer|authenticationHttpBasicUsername|authenticationHttpBasicPassword| MsgVpnRestDeliveryPointRestConsumer|remotePort|tlsEnabled| MsgVpnRestDeliveryPointRestConsumer|tlsEnabled|remotePort|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnRestDeliveryPointRestConsumer**](MsgVpnRestDeliveryPointRestConsumer.md)| The REST Consumer object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **restConsumerName** | **string**| The name of the REST Consumer. | 
 **optional** | ***RestDeliveryPointApiReplaceMsgVpnRestDeliveryPointRestConsumerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RestDeliveryPointApiReplaceMsgVpnRestDeliveryPointRestConsumerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointRestConsumerResponse**](MsgVpnRestDeliveryPointRestConsumerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnRestDeliveryPoint**
> MsgVpnRestDeliveryPointResponse UpdateMsgVpnRestDeliveryPoint(ctx, body, msgVpnName, restDeliveryPointName, optional)
Update a REST Delivery Point object.

Update a REST Delivery Point object. Any attribute missing from the request will be left unchanged.  A REST Delivery Point manages delivery of messages from queues to a named list of REST Consumers.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- clientProfileName|||||x|| msgVpnName|x|x||||| restDeliveryPointName|x|x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnRestDeliveryPoint**](MsgVpnRestDeliveryPoint.md)| The REST Delivery Point object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
 **optional** | ***RestDeliveryPointApiUpdateMsgVpnRestDeliveryPointOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RestDeliveryPointApiUpdateMsgVpnRestDeliveryPointOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointResponse**](MsgVpnRestDeliveryPointResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnRestDeliveryPointQueueBinding**
> MsgVpnRestDeliveryPointQueueBindingResponse UpdateMsgVpnRestDeliveryPointQueueBinding(ctx, body, msgVpnName, restDeliveryPointName, queueBindingName, optional)
Update a Queue Binding object.

Update a Queue Binding object. Any attribute missing from the request will be left unchanged.  A Queue Binding for a REST Delivery Point attracts messages to be delivered to REST consumers. If the queue does not exist it can be created subsequently, and once the queue is operational the broker performs the queue binding. Removing the queue binding does not delete the queue itself. Similarly, removing the queue does not remove the queue binding, which fails until the queue is recreated or the queue binding is deleted.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- msgVpnName|x|x||||| queueBindingName|x|x||||| restDeliveryPointName|x|x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnRestDeliveryPointQueueBinding**](MsgVpnRestDeliveryPointQueueBinding.md)| The Queue Binding object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **queueBindingName** | **string**| The name of a queue in the Message VPN. | 
 **optional** | ***RestDeliveryPointApiUpdateMsgVpnRestDeliveryPointQueueBindingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RestDeliveryPointApiUpdateMsgVpnRestDeliveryPointQueueBindingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointQueueBindingResponse**](MsgVpnRestDeliveryPointQueueBindingResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader**
> MsgVpnRestDeliveryPointQueueBindingProtectedRequestHeaderResponse UpdateMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader(ctx, body, msgVpnName, restDeliveryPointName, queueBindingName, headerName, optional)
Update a Protected Request Header object.

Update a Protected Request Header object. Any attribute missing from the request will be left unchanged.  A protected request header to be added to the HTTP request. Unlike a non-protected request header, the header value cannot be displayed after it is set.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- headerName|x|x||||| headerValue|||x||||x msgVpnName|x|x||||| queueBindingName|x|x||||| restDeliveryPointName|x|x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.30.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader**](MsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader.md)| The Protected Request Header object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **queueBindingName** | **string**| The name of a queue in the Message VPN. | 
  **headerName** | **string**| The name of the protected HTTP request header. | 
 **optional** | ***RestDeliveryPointApiUpdateMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeaderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RestDeliveryPointApiUpdateMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeaderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointQueueBindingProtectedRequestHeaderResponse**](MsgVpnRestDeliveryPointQueueBindingProtectedRequestHeaderResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnRestDeliveryPointQueueBindingRequestHeader**
> MsgVpnRestDeliveryPointQueueBindingRequestHeaderResponse UpdateMsgVpnRestDeliveryPointQueueBindingRequestHeader(ctx, body, msgVpnName, restDeliveryPointName, queueBindingName, headerName, optional)
Update a Request Header object.

Update a Request Header object. Any attribute missing from the request will be left unchanged.  A request header to be added to the HTTP request.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- headerName|x|x||||| msgVpnName|x|x||||| queueBindingName|x|x||||| restDeliveryPointName|x|x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.23.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnRestDeliveryPointQueueBindingRequestHeader**](MsgVpnRestDeliveryPointQueueBindingRequestHeader.md)| The Request Header object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **queueBindingName** | **string**| The name of a queue in the Message VPN. | 
  **headerName** | **string**| The name of the HTTP request header. | 
 **optional** | ***RestDeliveryPointApiUpdateMsgVpnRestDeliveryPointQueueBindingRequestHeaderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RestDeliveryPointApiUpdateMsgVpnRestDeliveryPointQueueBindingRequestHeaderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointQueueBindingRequestHeaderResponse**](MsgVpnRestDeliveryPointQueueBindingRequestHeaderResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnRestDeliveryPointRestConsumer**
> MsgVpnRestDeliveryPointRestConsumerResponse UpdateMsgVpnRestDeliveryPointRestConsumer(ctx, body, msgVpnName, restDeliveryPointName, restConsumerName, optional)
Update a REST Consumer object.

Update a REST Consumer object. Any attribute missing from the request will be left unchanged.  REST Consumer objects establish HTTP connectivity to REST consumer applications who wish to receive messages from a broker.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- authenticationAwsSecretAccessKey|||x||||x authenticationClientCertContent|||x||x||x authenticationClientCertPassword|||x||x|| authenticationHttpBasicPassword|||x||x||x authenticationHttpBasicUsername|||||x|| authenticationHttpHeaderValue|||x||||x authenticationOauthClientId|||||x|| authenticationOauthClientScope|||||x|| authenticationOauthClientSecret|||x||x||x authenticationOauthClientTokenEndpoint|||||x|| authenticationOauthClientTokenExpiryDefault|||||x|| authenticationOauthJwtSecretKey|||x||x||x authenticationOauthJwtTokenEndpoint|||||x|| authenticationOauthJwtTokenExpiryDefault|||||x|| authenticationScheme|||||x|| msgVpnName|x|x||||| outgoingConnectionCount|||||x|| remoteHost|||||x|| remotePort|||||x|| restConsumerName|x|x||||| restDeliveryPointName|x|x||||| tlsCipherSuiteList|||||x|| tlsEnabled|||||x||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- MsgVpnRestDeliveryPointRestConsumer|authenticationClientCertPassword|authenticationClientCertContent| MsgVpnRestDeliveryPointRestConsumer|authenticationHttpBasicPassword|authenticationHttpBasicUsername| MsgVpnRestDeliveryPointRestConsumer|authenticationHttpBasicUsername|authenticationHttpBasicPassword| MsgVpnRestDeliveryPointRestConsumer|remotePort|tlsEnabled| MsgVpnRestDeliveryPointRestConsumer|tlsEnabled|remotePort|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnRestDeliveryPointRestConsumer**](MsgVpnRestDeliveryPointRestConsumer.md)| The REST Consumer object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **restConsumerName** | **string**| The name of the REST Consumer. | 
 **optional** | ***RestDeliveryPointApiUpdateMsgVpnRestDeliveryPointRestConsumerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RestDeliveryPointApiUpdateMsgVpnRestDeliveryPointRestConsumerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointRestConsumerResponse**](MsgVpnRestDeliveryPointRestConsumerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

