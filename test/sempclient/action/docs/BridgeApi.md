# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/action*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DoMsgVpnBridgeClearEvent**](BridgeApi.md#DoMsgVpnBridgeClearEvent) | **Put** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/clearEvent | Clear an event for the Bridge so it can be generated anew.
[**DoMsgVpnBridgeClearStats**](BridgeApi.md#DoMsgVpnBridgeClearStats) | **Put** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/clearStats | Clear the statistics for the Bridge.
[**DoMsgVpnBridgeDisconnect**](BridgeApi.md#DoMsgVpnBridgeDisconnect) | **Put** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/disconnect | Disconnect the Bridge.
[**GetMsgVpnBridge**](BridgeApi.md#GetMsgVpnBridge) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter} | Get a Bridge object.
[**GetMsgVpnBridges**](BridgeApi.md#GetMsgVpnBridges) | **Get** /msgVpns/{msgVpnName}/bridges | Get a list of Bridge objects.

# **DoMsgVpnBridgeClearEvent**
> SempMetaOnlyResponse DoMsgVpnBridgeClearEvent(ctx, body, msgVpnName, bridgeName, bridgeVirtualRouter)
Clear an event for the Bridge so it can be generated anew.

Clear an event for the Bridge so it can be generated anew.   Attribute|Required|Deprecated :---|:---:|:---: eventName|x|    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnBridgeClearEvent**](MsgVpnBridgeClearEvent.md)| The Clear Event action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **bridgeName** | **string**| The name of the Bridge. | 
  **bridgeVirtualRouter** | **string**| The virtual router of the Bridge. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnBridgeClearStats**
> SempMetaOnlyResponse DoMsgVpnBridgeClearStats(ctx, body, msgVpnName, bridgeName, bridgeVirtualRouter)
Clear the statistics for the Bridge.

Clear the statistics for the Bridge.    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnBridgeClearStats**](MsgVpnBridgeClearStats.md)| The Clear Stats action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **bridgeName** | **string**| The name of the Bridge. | 
  **bridgeVirtualRouter** | **string**| The virtual router of the Bridge. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnBridgeDisconnect**
> SempMetaOnlyResponse DoMsgVpnBridgeDisconnect(ctx, body, msgVpnName, bridgeName, bridgeVirtualRouter)
Disconnect the Bridge.

Disconnect the Bridge.    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnBridgeDisconnect**](MsgVpnBridgeDisconnect.md)| The Disconnect action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **bridgeName** | **string**| The name of the Bridge. | 
  **bridgeVirtualRouter** | **string**| The virtual router of the Bridge. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnBridge**
> MsgVpnBridgeResponse GetMsgVpnBridge(ctx, msgVpnName, bridgeName, bridgeVirtualRouter, optional)
Get a Bridge object.

Get a Bridge object.  Bridges can be used to link two Message VPNs so that messages published to one Message VPN that match the topic subscriptions set for the bridge are also delivered to the linked Message VPN.   Attribute|Identifying|Deprecated :---|:---:|:---: bridgeName|x| bridgeVirtualRouter|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **bridgeName** | **string**| The name of the Bridge. | 
  **bridgeVirtualRouter** | **string**| The virtual router of the Bridge. | 
 **optional** | ***BridgeApiGetMsgVpnBridgeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BridgeApiGetMsgVpnBridgeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeResponse**](MsgVpnBridgeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnBridges**
> MsgVpnBridgesResponse GetMsgVpnBridges(ctx, msgVpnName, optional)
Get a list of Bridge objects.

Get a list of Bridge objects.  Bridges can be used to link two Message VPNs so that messages published to one Message VPN that match the topic subscriptions set for the bridge are also delivered to the linked Message VPN.   Attribute|Identifying|Deprecated :---|:---:|:---: bridgeName|x| bridgeVirtualRouter|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***BridgeApiGetMsgVpnBridgesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BridgeApiGetMsgVpnBridgesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgesResponse**](MsgVpnBridgesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

