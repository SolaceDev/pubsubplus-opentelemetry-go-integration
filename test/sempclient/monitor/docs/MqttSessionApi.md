# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/monitor*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMsgVpnMqttSession**](MqttSessionApi.md#GetMsgVpnMqttSession) | **Get** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter} | Get an MQTT Session object.
[**GetMsgVpnMqttSessionSubscription**](MqttSessionApi.md#GetMsgVpnMqttSessionSubscription) | **Get** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions/{subscriptionTopic} | Get a Subscription object.
[**GetMsgVpnMqttSessionSubscriptions**](MqttSessionApi.md#GetMsgVpnMqttSessionSubscriptions) | **Get** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions | Get a list of Subscription objects.
[**GetMsgVpnMqttSessions**](MqttSessionApi.md#GetMsgVpnMqttSessions) | **Get** /msgVpns/{msgVpnName}/mqttSessions | Get a list of MQTT Session objects.

# **GetMsgVpnMqttSession**
> MsgVpnMqttSessionResponse GetMsgVpnMqttSession(ctx, msgVpnName, mqttSessionClientId, mqttSessionVirtualRouter, optional)
Get an MQTT Session object.

Get an MQTT Session object.  An MQTT Session object is a virtual representation of an MQTT client connection. An MQTT session holds the state of an MQTT client (that is, it is used to contain a client's QoS 0 and QoS 1 subscription sets and any undelivered QoS 1 messages).   Attribute|Identifying|Deprecated :---|:---:|:---: counter.mqttConnackErrorTxCount||x counter.mqttConnackTxCount||x counter.mqttConnectRxCount||x counter.mqttDisconnectRxCount||x counter.mqttPubcompTxCount||x counter.mqttPublishQos0RxCount||x counter.mqttPublishQos0TxCount||x counter.mqttPublishQos1RxCount||x counter.mqttPublishQos1TxCount||x counter.mqttPublishQos2RxCount||x counter.mqttPubrecTxCount||x counter.mqttPubrelRxCount||x mqttSessionClientId|x| mqttSessionVirtualRouter|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **mqttSessionClientId** | **string**| The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet. | 
  **mqttSessionVirtualRouter** | **string**| The virtual router of the MQTT Session. | 
 **optional** | ***MqttSessionApiGetMsgVpnMqttSessionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MqttSessionApiGetMsgVpnMqttSessionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnMqttSessionResponse**](MsgVpnMqttSessionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnMqttSessionSubscription**
> MsgVpnMqttSessionSubscriptionResponse GetMsgVpnMqttSessionSubscription(ctx, msgVpnName, mqttSessionClientId, mqttSessionVirtualRouter, subscriptionTopic, optional)
Get a Subscription object.

Get a Subscription object.  An MQTT session contains a client's QoS 0 and QoS 1 subscription sets. On creation, a subscription defaults to QoS 0.   Attribute|Identifying|Deprecated :---|:---:|:---: mqttSessionClientId|x| mqttSessionVirtualRouter|x| msgVpnName|x| subscriptionTopic|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **mqttSessionClientId** | **string**| The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet. | 
  **mqttSessionVirtualRouter** | **string**| The virtual router of the MQTT Session. | 
  **subscriptionTopic** | **string**| The MQTT subscription topic. | 
 **optional** | ***MqttSessionApiGetMsgVpnMqttSessionSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MqttSessionApiGetMsgVpnMqttSessionSubscriptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnMqttSessionSubscriptionResponse**](MsgVpnMqttSessionSubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnMqttSessionSubscriptions**
> MsgVpnMqttSessionSubscriptionsResponse GetMsgVpnMqttSessionSubscriptions(ctx, msgVpnName, mqttSessionClientId, mqttSessionVirtualRouter, optional)
Get a list of Subscription objects.

Get a list of Subscription objects.  An MQTT session contains a client's QoS 0 and QoS 1 subscription sets. On creation, a subscription defaults to QoS 0.   Attribute|Identifying|Deprecated :---|:---:|:---: mqttSessionClientId|x| mqttSessionVirtualRouter|x| msgVpnName|x| subscriptionTopic|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **mqttSessionClientId** | **string**| The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet. | 
  **mqttSessionVirtualRouter** | **string**| The virtual router of the MQTT Session. | 
 **optional** | ***MqttSessionApiGetMsgVpnMqttSessionSubscriptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MqttSessionApiGetMsgVpnMqttSessionSubscriptionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnMqttSessionSubscriptionsResponse**](MsgVpnMqttSessionSubscriptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnMqttSessions**
> MsgVpnMqttSessionsResponse GetMsgVpnMqttSessions(ctx, msgVpnName, optional)
Get a list of MQTT Session objects.

Get a list of MQTT Session objects.  An MQTT Session object is a virtual representation of an MQTT client connection. An MQTT session holds the state of an MQTT client (that is, it is used to contain a client's QoS 0 and QoS 1 subscription sets and any undelivered QoS 1 messages).   Attribute|Identifying|Deprecated :---|:---:|:---: counter.mqttConnackErrorTxCount||x counter.mqttConnackTxCount||x counter.mqttConnectRxCount||x counter.mqttDisconnectRxCount||x counter.mqttPubcompTxCount||x counter.mqttPublishQos0RxCount||x counter.mqttPublishQos0TxCount||x counter.mqttPublishQos1RxCount||x counter.mqttPublishQos1TxCount||x counter.mqttPublishQos2RxCount||x counter.mqttPubrecTxCount||x counter.mqttPubrelRxCount||x mqttSessionClientId|x| mqttSessionVirtualRouter|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MqttSessionApiGetMsgVpnMqttSessionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MqttSessionApiGetMsgVpnMqttSessionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnMqttSessionsResponse**](MsgVpnMqttSessionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

