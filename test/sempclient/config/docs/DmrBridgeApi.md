# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/config*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMsgVpnDmrBridge**](DmrBridgeApi.md#CreateMsgVpnDmrBridge) | **Post** /msgVpns/{msgVpnName}/dmrBridges | Create a DMR Bridge object.
[**DeleteMsgVpnDmrBridge**](DmrBridgeApi.md#DeleteMsgVpnDmrBridge) | **Delete** /msgVpns/{msgVpnName}/dmrBridges/{remoteNodeName} | Delete a DMR Bridge object.
[**GetMsgVpnDmrBridge**](DmrBridgeApi.md#GetMsgVpnDmrBridge) | **Get** /msgVpns/{msgVpnName}/dmrBridges/{remoteNodeName} | Get a DMR Bridge object.
[**GetMsgVpnDmrBridges**](DmrBridgeApi.md#GetMsgVpnDmrBridges) | **Get** /msgVpns/{msgVpnName}/dmrBridges | Get a list of DMR Bridge objects.
[**ReplaceMsgVpnDmrBridge**](DmrBridgeApi.md#ReplaceMsgVpnDmrBridge) | **Put** /msgVpns/{msgVpnName}/dmrBridges/{remoteNodeName} | Replace a DMR Bridge object.
[**UpdateMsgVpnDmrBridge**](DmrBridgeApi.md#UpdateMsgVpnDmrBridge) | **Patch** /msgVpns/{msgVpnName}/dmrBridges/{remoteNodeName} | Update a DMR Bridge object.

# **CreateMsgVpnDmrBridge**
> MsgVpnDmrBridgeResponse CreateMsgVpnDmrBridge(ctx, body, msgVpnName, optional)
Create a DMR Bridge object.

Create a DMR Bridge object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A DMR Bridge is required to establish a data channel over a corresponding external link to the remote node for a given Message VPN. Each DMR Bridge identifies which external link the Message VPN should use, and what the name of the equivalent Message VPN at the remote node is.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x||| remoteNodeName|x|x||||    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnDmrBridge**](MsgVpnDmrBridge.md)| The DMR Bridge object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***DmrBridgeApiCreateMsgVpnDmrBridgeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DmrBridgeApiCreateMsgVpnDmrBridgeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDmrBridgeResponse**](MsgVpnDmrBridgeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnDmrBridge**
> SempMetaOnlyResponse DeleteMsgVpnDmrBridge(ctx, msgVpnName, remoteNodeName)
Delete a DMR Bridge object.

Delete a DMR Bridge object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A DMR Bridge is required to establish a data channel over a corresponding external link to the remote node for a given Message VPN. Each DMR Bridge identifies which external link the Message VPN should use, and what the name of the equivalent Message VPN at the remote node is.  A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **remoteNodeName** | **string**| The name of the node at the remote end of the DMR Bridge. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnDmrBridge**
> MsgVpnDmrBridgeResponse GetMsgVpnDmrBridge(ctx, msgVpnName, remoteNodeName, optional)
Get a DMR Bridge object.

Get a DMR Bridge object.  A DMR Bridge is required to establish a data channel over a corresponding external link to the remote node for a given Message VPN. Each DMR Bridge identifies which external link the Message VPN should use, and what the name of the equivalent Message VPN at the remote node is.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| remoteNodeName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **remoteNodeName** | **string**| The name of the node at the remote end of the DMR Bridge. | 
 **optional** | ***DmrBridgeApiGetMsgVpnDmrBridgeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DmrBridgeApiGetMsgVpnDmrBridgeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDmrBridgeResponse**](MsgVpnDmrBridgeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnDmrBridges**
> MsgVpnDmrBridgesResponse GetMsgVpnDmrBridges(ctx, msgVpnName, optional)
Get a list of DMR Bridge objects.

Get a list of DMR Bridge objects.  A DMR Bridge is required to establish a data channel over a corresponding external link to the remote node for a given Message VPN. Each DMR Bridge identifies which external link the Message VPN should use, and what the name of the equivalent Message VPN at the remote node is.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| remoteNodeName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***DmrBridgeApiGetMsgVpnDmrBridgesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DmrBridgeApiGetMsgVpnDmrBridgesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDmrBridgesResponse**](MsgVpnDmrBridgesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnDmrBridge**
> MsgVpnDmrBridgeResponse ReplaceMsgVpnDmrBridge(ctx, body, msgVpnName, remoteNodeName, optional)
Replace a DMR Bridge object.

Replace a DMR Bridge object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  A DMR Bridge is required to establish a data channel over a corresponding external link to the remote node for a given Message VPN. Each DMR Bridge identifies which external link the Message VPN should use, and what the name of the equivalent Message VPN at the remote node is.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- msgVpnName|x||x||||| remoteNodeName|x||x|||||    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnDmrBridge**](MsgVpnDmrBridge.md)| The DMR Bridge object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **remoteNodeName** | **string**| The name of the node at the remote end of the DMR Bridge. | 
 **optional** | ***DmrBridgeApiReplaceMsgVpnDmrBridgeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DmrBridgeApiReplaceMsgVpnDmrBridgeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDmrBridgeResponse**](MsgVpnDmrBridgeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnDmrBridge**
> MsgVpnDmrBridgeResponse UpdateMsgVpnDmrBridge(ctx, body, msgVpnName, remoteNodeName, optional)
Update a DMR Bridge object.

Update a DMR Bridge object. Any attribute missing from the request will be left unchanged.  A DMR Bridge is required to establish a data channel over a corresponding external link to the remote node for a given Message VPN. Each DMR Bridge identifies which external link the Message VPN should use, and what the name of the equivalent Message VPN at the remote node is.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- msgVpnName|x|x||||| remoteNodeName|x|x|||||    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnDmrBridge**](MsgVpnDmrBridge.md)| The DMR Bridge object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **remoteNodeName** | **string**| The name of the node at the remote end of the DMR Bridge. | 
 **optional** | ***DmrBridgeApiUpdateMsgVpnDmrBridgeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DmrBridgeApiUpdateMsgVpnDmrBridgeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDmrBridgeResponse**](MsgVpnDmrBridgeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

