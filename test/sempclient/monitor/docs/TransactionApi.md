# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/monitor*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMsgVpnTransaction**](TransactionApi.md#GetMsgVpnTransaction) | **Get** /msgVpns/{msgVpnName}/transactions/{xid} | Get a Replicated Local Transaction or XA Transaction object.
[**GetMsgVpnTransactionConsumerMsg**](TransactionApi.md#GetMsgVpnTransactionConsumerMsg) | **Get** /msgVpns/{msgVpnName}/transactions/{xid}/consumerMsgs/{msgId} | Get a Transaction Consumer Message object.
[**GetMsgVpnTransactionConsumerMsgs**](TransactionApi.md#GetMsgVpnTransactionConsumerMsgs) | **Get** /msgVpns/{msgVpnName}/transactions/{xid}/consumerMsgs | Get a list of Transaction Consumer Message objects.
[**GetMsgVpnTransactionPublisherMsg**](TransactionApi.md#GetMsgVpnTransactionPublisherMsg) | **Get** /msgVpns/{msgVpnName}/transactions/{xid}/publisherMsgs/{msgId} | Get a Transaction Publisher Message object.
[**GetMsgVpnTransactionPublisherMsgs**](TransactionApi.md#GetMsgVpnTransactionPublisherMsgs) | **Get** /msgVpns/{msgVpnName}/transactions/{xid}/publisherMsgs | Get a list of Transaction Publisher Message objects.
[**GetMsgVpnTransactions**](TransactionApi.md#GetMsgVpnTransactions) | **Get** /msgVpns/{msgVpnName}/transactions | Get a list of Replicated Local Transaction or XA Transaction objects.

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

# **GetMsgVpnTransactionConsumerMsg**
> MsgVpnTransactionConsumerMsgResponse GetMsgVpnTransactionConsumerMsg(ctx, msgVpnName, xid, msgId, optional)
Get a Transaction Consumer Message object.

Get a Transaction Consumer Message object.  A Transaction Consumer Message is a message that will be consumed as part of this Transaction once the Transaction is committed.   Attribute|Identifying|Deprecated :---|:---:|:---: msgId|x| msgVpnName|x| xid|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **xid** | **string**| The identifier (ID) of the Transaction. | 
  **msgId** | **string**| The identifier (ID) of the Message. | 
 **optional** | ***TransactionApiGetMsgVpnTransactionConsumerMsgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TransactionApiGetMsgVpnTransactionConsumerMsgOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTransactionConsumerMsgResponse**](MsgVpnTransactionConsumerMsgResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnTransactionConsumerMsgs**
> MsgVpnTransactionConsumerMsgsResponse GetMsgVpnTransactionConsumerMsgs(ctx, msgVpnName, xid, optional)
Get a list of Transaction Consumer Message objects.

Get a list of Transaction Consumer Message objects.  A Transaction Consumer Message is a message that will be consumed as part of this Transaction once the Transaction is committed.   Attribute|Identifying|Deprecated :---|:---:|:---: msgId|x| msgVpnName|x| xid|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **xid** | **string**| The identifier (ID) of the Transaction. | 
 **optional** | ***TransactionApiGetMsgVpnTransactionConsumerMsgsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TransactionApiGetMsgVpnTransactionConsumerMsgsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTransactionConsumerMsgsResponse**](MsgVpnTransactionConsumerMsgsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnTransactionPublisherMsg**
> MsgVpnTransactionPublisherMsgResponse GetMsgVpnTransactionPublisherMsg(ctx, msgVpnName, xid, msgId, optional)
Get a Transaction Publisher Message object.

Get a Transaction Publisher Message object.  A Transaction Publisher Message is a message that will be published as part of this Transaction once the Transaction is committed.   Attribute|Identifying|Deprecated :---|:---:|:---: msgId|x| msgVpnName|x| xid|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **xid** | **string**| The identifier (ID) of the Transaction. | 
  **msgId** | **string**| The identifier (ID) of the Message. | 
 **optional** | ***TransactionApiGetMsgVpnTransactionPublisherMsgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TransactionApiGetMsgVpnTransactionPublisherMsgOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTransactionPublisherMsgResponse**](MsgVpnTransactionPublisherMsgResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnTransactionPublisherMsgs**
> MsgVpnTransactionPublisherMsgsResponse GetMsgVpnTransactionPublisherMsgs(ctx, msgVpnName, xid, optional)
Get a list of Transaction Publisher Message objects.

Get a list of Transaction Publisher Message objects.  A Transaction Publisher Message is a message that will be published as part of this Transaction once the Transaction is committed.   Attribute|Identifying|Deprecated :---|:---:|:---: msgId|x| msgVpnName|x| xid|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **xid** | **string**| The identifier (ID) of the Transaction. | 
 **optional** | ***TransactionApiGetMsgVpnTransactionPublisherMsgsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TransactionApiGetMsgVpnTransactionPublisherMsgsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTransactionPublisherMsgsResponse**](MsgVpnTransactionPublisherMsgsResponse.md)

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

