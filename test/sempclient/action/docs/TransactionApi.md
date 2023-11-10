# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/action*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DoMsgVpnTransactionCommit**](TransactionApi.md#DoMsgVpnTransactionCommit) | **Put** /msgVpns/{msgVpnName}/transactions/{xid}/commit | Commit the Transaction.
[**DoMsgVpnTransactionDelete**](TransactionApi.md#DoMsgVpnTransactionDelete) | **Put** /msgVpns/{msgVpnName}/transactions/{xid}/delete | Delete the Transaction.
[**DoMsgVpnTransactionRollback**](TransactionApi.md#DoMsgVpnTransactionRollback) | **Put** /msgVpns/{msgVpnName}/transactions/{xid}/rollback | Rollback the Transaction.
[**GetMsgVpnTransaction**](TransactionApi.md#GetMsgVpnTransaction) | **Get** /msgVpns/{msgVpnName}/transactions/{xid} | Get a Replicated Local Transaction or XA Transaction object.
[**GetMsgVpnTransactions**](TransactionApi.md#GetMsgVpnTransactions) | **Get** /msgVpns/{msgVpnName}/transactions | Get a list of Replicated Local Transaction or XA Transaction objects.

# **DoMsgVpnTransactionCommit**
> SempMetaOnlyResponse DoMsgVpnTransactionCommit(ctx, body, msgVpnName, xid)
Commit the Transaction.

Commit the Transaction.    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnTransactionCommit**](MsgVpnTransactionCommit.md)| The Commit action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **xid** | **string**| The identifier (ID) of the Transaction. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnTransactionDelete**
> SempMetaOnlyResponse DoMsgVpnTransactionDelete(ctx, body, msgVpnName, xid)
Delete the Transaction.

Delete the Transaction.    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnTransactionDelete**](MsgVpnTransactionDelete.md)| The Delete action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **xid** | **string**| The identifier (ID) of the Transaction. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnTransactionRollback**
> SempMetaOnlyResponse DoMsgVpnTransactionRollback(ctx, body, msgVpnName, xid)
Rollback the Transaction.

Rollback the Transaction.    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnTransactionRollback**](MsgVpnTransactionRollback.md)| The Rollback action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **xid** | **string**| The identifier (ID) of the Transaction. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnTransaction**
> MsgVpnTransactionResponse GetMsgVpnTransaction(ctx, msgVpnName, xid, optional)
Get a Replicated Local Transaction or XA Transaction object.

Get a Replicated Local Transaction or XA Transaction object.  Transactions can be used to group a set of Guaranteed messages to be published or consumed or both as an atomic unit of work.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| xid|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **xid** | **string**| The identifier (ID) of the Transaction. | 
 **optional** | ***TransactionApiGetMsgVpnTransactionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TransactionApiGetMsgVpnTransactionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTransactionResponse**](MsgVpnTransactionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnTransactions**
> MsgVpnTransactionsResponse GetMsgVpnTransactions(ctx, msgVpnName, optional)
Get a list of Replicated Local Transaction or XA Transaction objects.

Get a list of Replicated Local Transaction or XA Transaction objects.  Transactions can be used to group a set of Guaranteed messages to be published or consumed or both as an atomic unit of work.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| xid|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***TransactionApiGetMsgVpnTransactionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TransactionApiGetMsgVpnTransactionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTransactionsResponse**](MsgVpnTransactionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

