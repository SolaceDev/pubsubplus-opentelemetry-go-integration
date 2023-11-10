# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/monitor*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfigSyncLocalDatabaseRow**](ConfigSyncLocalDatabaseRowApi.md#GetConfigSyncLocalDatabaseRow) | **Get** /configSyncLocalDatabaseRows/{type},{name} | Get a Config Sync Local Database object.
[**GetConfigSyncLocalDatabaseRows**](ConfigSyncLocalDatabaseRowApi.md#GetConfigSyncLocalDatabaseRows) | **Get** /configSyncLocalDatabaseRows | Get a list of Config Sync Local Database objects.

# **GetConfigSyncLocalDatabaseRow**
> ConfigSyncLocalDatabaseRowResponse GetConfigSyncLocalDatabaseRow(ctx, type_, name, optional)
Get a Config Sync Local Database object.

Get a Config Sync Local Database object.  Config Sync Local Database Rows contains information about the status of the table for this Broker or a local Message VPN.   Attribute|Identifying|Deprecated :---|:---:|:---: name|x| type|x|    A SEMP client authorized with a minimum access scope/level of \"global/none\" is required to perform this operation.  This has been available since 2.22.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**| The type of the row. Can be one of \&quot;router\&quot; or \&quot;vpn\&quot;. There is one \&quot;router\&quot; row and one row for each configured \&quot;vpn\&quot;. Each row represents a table of information that is synchronized between Config Sync and replication mates. | 
  **name** | **string**| The name is \&quot;site\&quot; when the row type is \&quot;router\&quot;, otherwise it is the Message VPN name. | 
 **optional** | ***ConfigSyncLocalDatabaseRowApiGetConfigSyncLocalDatabaseRowOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigSyncLocalDatabaseRowApiGetConfigSyncLocalDatabaseRowOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**ConfigSyncLocalDatabaseRowResponse**](ConfigSyncLocalDatabaseRowResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConfigSyncLocalDatabaseRows**
> ConfigSyncLocalDatabaseRowsResponse GetConfigSyncLocalDatabaseRows(ctx, optional)
Get a list of Config Sync Local Database objects.

Get a list of Config Sync Local Database objects.  Config Sync Local Database Rows contains information about the status of the table for this Broker or a local Message VPN.   Attribute|Identifying|Deprecated :---|:---:|:---: name|x| type|x|    A SEMP client authorized with a minimum access scope/level of \"global/none\" is required to perform this operation.  This has been available since 2.22.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigSyncLocalDatabaseRowApiGetConfigSyncLocalDatabaseRowsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigSyncLocalDatabaseRowApiGetConfigSyncLocalDatabaseRowsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**ConfigSyncLocalDatabaseRowsResponse**](ConfigSyncLocalDatabaseRowsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

