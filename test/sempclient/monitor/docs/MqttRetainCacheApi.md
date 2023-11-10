# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/monitor*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMsgVpnMqttRetainCache**](MqttRetainCacheApi.md#GetMsgVpnMqttRetainCache) | **Get** /msgVpns/{msgVpnName}/mqttRetainCaches/{cacheName} | Get an MQTT Retain Cache object.
[**GetMsgVpnMqttRetainCaches**](MqttRetainCacheApi.md#GetMsgVpnMqttRetainCaches) | **Get** /msgVpns/{msgVpnName}/mqttRetainCaches | Get a list of MQTT Retain Cache objects.

# **GetMsgVpnMqttRetainCache**
> MsgVpnMqttRetainCacheResponse GetMsgVpnMqttRetainCache(ctx, msgVpnName, cacheName, optional)
Get an MQTT Retain Cache object.

Get an MQTT Retain Cache object.  Using MQTT retained messages allows publishing MQTT clients to indicate that a message must be stored for later delivery to subscribing clients when those subscribing clients add subscriptions matching the retained message's topic. An MQTT Retain Cache processes all retained messages for a Message VPN.   Attribute|Identifying|Deprecated :---|:---:|:---: cacheName|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **cacheName** | **string**| The name of the MQTT Retain Cache. | 
 **optional** | ***MqttRetainCacheApiGetMsgVpnMqttRetainCacheOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MqttRetainCacheApiGetMsgVpnMqttRetainCacheOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnMqttRetainCacheResponse**](MsgVpnMqttRetainCacheResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnMqttRetainCaches**
> MsgVpnMqttRetainCachesResponse GetMsgVpnMqttRetainCaches(ctx, msgVpnName, optional)
Get a list of MQTT Retain Cache objects.

Get a list of MQTT Retain Cache objects.  Using MQTT retained messages allows publishing MQTT clients to indicate that a message must be stored for later delivery to subscribing clients when those subscribing clients add subscriptions matching the retained message's topic. An MQTT Retain Cache processes all retained messages for a Message VPN.   Attribute|Identifying|Deprecated :---|:---:|:---: cacheName|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MqttRetainCacheApiGetMsgVpnMqttRetainCachesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MqttRetainCacheApiGetMsgVpnMqttRetainCachesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnMqttRetainCachesResponse**](MsgVpnMqttRetainCachesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

