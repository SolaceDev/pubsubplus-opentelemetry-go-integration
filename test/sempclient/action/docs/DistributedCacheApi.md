# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/action*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DoMsgVpnDistributedCacheClusterInstanceBackupCachedMsgs**](DistributedCacheApi.md#DoMsgVpnDistributedCacheClusterInstanceBackupCachedMsgs) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/backupCachedMsgs | Backup cached messages of the Cache Instance to disk.
[**DoMsgVpnDistributedCacheClusterInstanceCancelBackupCachedMsgs**](DistributedCacheApi.md#DoMsgVpnDistributedCacheClusterInstanceCancelBackupCachedMsgs) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/cancelBackupCachedMsgs | Cancel the backup of cached messages from the Cache Instance.
[**DoMsgVpnDistributedCacheClusterInstanceCancelRestoreCachedMsgs**](DistributedCacheApi.md#DoMsgVpnDistributedCacheClusterInstanceCancelRestoreCachedMsgs) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/cancelRestoreCachedMsgs | Cancel the restore of cached messages to the Cache Instance.
[**DoMsgVpnDistributedCacheClusterInstanceClearEvent**](DistributedCacheApi.md#DoMsgVpnDistributedCacheClusterInstanceClearEvent) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/clearEvent | Clear an event for the Cache Instance so it can be generated anew.
[**DoMsgVpnDistributedCacheClusterInstanceClearStats**](DistributedCacheApi.md#DoMsgVpnDistributedCacheClusterInstanceClearStats) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/clearStats | Clear the statistics for the Cache Instance.
[**DoMsgVpnDistributedCacheClusterInstanceDeleteMsgs**](DistributedCacheApi.md#DoMsgVpnDistributedCacheClusterInstanceDeleteMsgs) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/deleteMsgs | Delete messages covered by the given topic in the Cache Instance.
[**DoMsgVpnDistributedCacheClusterInstanceRestoreCachedMsgs**](DistributedCacheApi.md#DoMsgVpnDistributedCacheClusterInstanceRestoreCachedMsgs) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/restoreCachedMsgs | Restore cached messages for the Cache Instance from disk.
[**DoMsgVpnDistributedCacheClusterInstanceStart**](DistributedCacheApi.md#DoMsgVpnDistributedCacheClusterInstanceStart) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/start | Start the Cache Instance.
[**GetMsgVpnDistributedCache**](DistributedCacheApi.md#GetMsgVpnDistributedCache) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName} | Get a Distributed Cache object.
[**GetMsgVpnDistributedCacheCluster**](DistributedCacheApi.md#GetMsgVpnDistributedCacheCluster) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName} | Get a Cache Cluster object.
[**GetMsgVpnDistributedCacheClusterInstance**](DistributedCacheApi.md#GetMsgVpnDistributedCacheClusterInstance) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName} | Get a Cache Instance object.
[**GetMsgVpnDistributedCacheClusterInstances**](DistributedCacheApi.md#GetMsgVpnDistributedCacheClusterInstances) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances | Get a list of Cache Instance objects.
[**GetMsgVpnDistributedCacheClusters**](DistributedCacheApi.md#GetMsgVpnDistributedCacheClusters) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters | Get a list of Cache Cluster objects.
[**GetMsgVpnDistributedCaches**](DistributedCacheApi.md#GetMsgVpnDistributedCaches) | **Get** /msgVpns/{msgVpnName}/distributedCaches | Get a list of Distributed Cache objects.

# **DoMsgVpnDistributedCacheClusterInstanceBackupCachedMsgs**
> SempMetaOnlyResponse DoMsgVpnDistributedCacheClusterInstanceBackupCachedMsgs(ctx, body, msgVpnName, cacheName, clusterName, instanceName)
Backup cached messages of the Cache Instance to disk.

Backup cached messages of the Cache Instance to disk.    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnDistributedCacheClusterInstanceBackupCachedMsgs**](MsgVpnDistributedCacheClusterInstanceBackupCachedMsgs.md)| The Backup Cached Messages action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **cacheName** | **string**| The name of the Distributed Cache. | 
  **clusterName** | **string**| The name of the Cache Cluster. | 
  **instanceName** | **string**| The name of the Cache Instance. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnDistributedCacheClusterInstanceCancelBackupCachedMsgs**
> SempMetaOnlyResponse DoMsgVpnDistributedCacheClusterInstanceCancelBackupCachedMsgs(ctx, body, msgVpnName, cacheName, clusterName, instanceName)
Cancel the backup of cached messages from the Cache Instance.

Cancel the backup of cached messages from the Cache Instance.    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnDistributedCacheClusterInstanceCancelBackupCachedMsgs**](MsgVpnDistributedCacheClusterInstanceCancelBackupCachedMsgs.md)| The Cancel Backup Cached Messages action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **cacheName** | **string**| The name of the Distributed Cache. | 
  **clusterName** | **string**| The name of the Cache Cluster. | 
  **instanceName** | **string**| The name of the Cache Instance. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnDistributedCacheClusterInstanceCancelRestoreCachedMsgs**
> SempMetaOnlyResponse DoMsgVpnDistributedCacheClusterInstanceCancelRestoreCachedMsgs(ctx, body, msgVpnName, cacheName, clusterName, instanceName)
Cancel the restore of cached messages to the Cache Instance.

Cancel the restore of cached messages to the Cache Instance.    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnDistributedCacheClusterInstanceCancelRestoreCachedMsgs**](MsgVpnDistributedCacheClusterInstanceCancelRestoreCachedMsgs.md)| The Cancel Restore Cached Messages action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **cacheName** | **string**| The name of the Distributed Cache. | 
  **clusterName** | **string**| The name of the Cache Cluster. | 
  **instanceName** | **string**| The name of the Cache Instance. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnDistributedCacheClusterInstanceClearEvent**
> SempMetaOnlyResponse DoMsgVpnDistributedCacheClusterInstanceClearEvent(ctx, body, msgVpnName, cacheName, clusterName, instanceName)
Clear an event for the Cache Instance so it can be generated anew.

Clear an event for the Cache Instance so it can be generated anew.   Attribute|Required|Deprecated :---|:---:|:---: eventName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnDistributedCacheClusterInstanceClearEvent**](MsgVpnDistributedCacheClusterInstanceClearEvent.md)| The Clear Event action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **cacheName** | **string**| The name of the Distributed Cache. | 
  **clusterName** | **string**| The name of the Cache Cluster. | 
  **instanceName** | **string**| The name of the Cache Instance. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnDistributedCacheClusterInstanceClearStats**
> SempMetaOnlyResponse DoMsgVpnDistributedCacheClusterInstanceClearStats(ctx, body, msgVpnName, cacheName, clusterName, instanceName)
Clear the statistics for the Cache Instance.

Clear the statistics for the Cache Instance.    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnDistributedCacheClusterInstanceClearStats**](MsgVpnDistributedCacheClusterInstanceClearStats.md)| The Clear Stats action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **cacheName** | **string**| The name of the Distributed Cache. | 
  **clusterName** | **string**| The name of the Cache Cluster. | 
  **instanceName** | **string**| The name of the Cache Instance. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnDistributedCacheClusterInstanceDeleteMsgs**
> SempMetaOnlyResponse DoMsgVpnDistributedCacheClusterInstanceDeleteMsgs(ctx, body, msgVpnName, cacheName, clusterName, instanceName)
Delete messages covered by the given topic in the Cache Instance.

Delete messages covered by the given topic in the Cache Instance.   Attribute|Required|Deprecated :---|:---:|:---: topic|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnDistributedCacheClusterInstanceDeleteMsgs**](MsgVpnDistributedCacheClusterInstanceDeleteMsgs.md)| The Delete Messages action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **cacheName** | **string**| The name of the Distributed Cache. | 
  **clusterName** | **string**| The name of the Cache Cluster. | 
  **instanceName** | **string**| The name of the Cache Instance. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnDistributedCacheClusterInstanceRestoreCachedMsgs**
> SempMetaOnlyResponse DoMsgVpnDistributedCacheClusterInstanceRestoreCachedMsgs(ctx, body, msgVpnName, cacheName, clusterName, instanceName)
Restore cached messages for the Cache Instance from disk.

Restore cached messages for the Cache Instance from disk.    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnDistributedCacheClusterInstanceRestoreCachedMsgs**](MsgVpnDistributedCacheClusterInstanceRestoreCachedMsgs.md)| The Restore Cached Messages action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **cacheName** | **string**| The name of the Distributed Cache. | 
  **clusterName** | **string**| The name of the Cache Cluster. | 
  **instanceName** | **string**| The name of the Cache Instance. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnDistributedCacheClusterInstanceStart**
> SempMetaOnlyResponse DoMsgVpnDistributedCacheClusterInstanceStart(ctx, body, msgVpnName, cacheName, clusterName, instanceName)
Start the Cache Instance.

Start the Cache Instance.    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnDistributedCacheClusterInstanceStart**](MsgVpnDistributedCacheClusterInstanceStart.md)| The Start Cache Instance action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **cacheName** | **string**| The name of the Distributed Cache. | 
  **clusterName** | **string**| The name of the Cache Cluster. | 
  **instanceName** | **string**| The name of the Cache Instance. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnDistributedCache**
> MsgVpnDistributedCacheResponse GetMsgVpnDistributedCache(ctx, msgVpnName, cacheName, optional)
Get a Distributed Cache object.

Get a Distributed Cache object.  A Distributed Cache is a collection of one or more Cache Clusters that belong to the same Message VPN. Each Cache Cluster in a Distributed Cache is configured to subscribe to a different set of topics. This effectively divides up the configured topic space, to provide scaling to very large topic spaces or very high cached message throughput.   Attribute|Identifying|Deprecated :---|:---:|:---: cacheName|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **cacheName** | **string**| The name of the Distributed Cache. | 
 **optional** | ***DistributedCacheApiGetMsgVpnDistributedCacheOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DistributedCacheApiGetMsgVpnDistributedCacheOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheResponse**](MsgVpnDistributedCacheResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnDistributedCacheCluster**
> MsgVpnDistributedCacheClusterResponse GetMsgVpnDistributedCacheCluster(ctx, msgVpnName, cacheName, clusterName, optional)
Get a Cache Cluster object.

Get a Cache Cluster object.  A Cache Cluster is a collection of one or more Cache Instances that subscribe to exactly the same topics. Cache Instances are grouped together in a Cache Cluster for the purpose of fault tolerance and load balancing. As published messages are received, the message broker message bus sends these live data messages to the Cache Instances in the Cache Cluster. This enables client cache requests to be served by any of Cache Instances in the Cache Cluster.   Attribute|Identifying|Deprecated :---|:---:|:---: cacheName|x| clusterName|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **cacheName** | **string**| The name of the Distributed Cache. | 
  **clusterName** | **string**| The name of the Cache Cluster. | 
 **optional** | ***DistributedCacheApiGetMsgVpnDistributedCacheClusterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DistributedCacheApiGetMsgVpnDistributedCacheClusterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterResponse**](MsgVpnDistributedCacheClusterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnDistributedCacheClusterInstance**
> MsgVpnDistributedCacheClusterInstanceResponse GetMsgVpnDistributedCacheClusterInstance(ctx, msgVpnName, cacheName, clusterName, instanceName, optional)
Get a Cache Instance object.

Get a Cache Instance object.  A Cache Instance is a single Cache process that belongs to a single Cache Cluster. A Cache Instance object provisioned on the broker is used to disseminate configuration information to the Cache process. Cache Instances listen for and cache live data messages that match the topic subscriptions configured for their parent Cache Cluster.   Attribute|Identifying|Deprecated :---|:---:|:---: cacheName|x| clusterName|x| instanceName|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **cacheName** | **string**| The name of the Distributed Cache. | 
  **clusterName** | **string**| The name of the Cache Cluster. | 
  **instanceName** | **string**| The name of the Cache Instance. | 
 **optional** | ***DistributedCacheApiGetMsgVpnDistributedCacheClusterInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DistributedCacheApiGetMsgVpnDistributedCacheClusterInstanceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterInstanceResponse**](MsgVpnDistributedCacheClusterInstanceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnDistributedCacheClusterInstances**
> MsgVpnDistributedCacheClusterInstancesResponse GetMsgVpnDistributedCacheClusterInstances(ctx, msgVpnName, cacheName, clusterName, optional)
Get a list of Cache Instance objects.

Get a list of Cache Instance objects.  A Cache Instance is a single Cache process that belongs to a single Cache Cluster. A Cache Instance object provisioned on the broker is used to disseminate configuration information to the Cache process. Cache Instances listen for and cache live data messages that match the topic subscriptions configured for their parent Cache Cluster.   Attribute|Identifying|Deprecated :---|:---:|:---: cacheName|x| clusterName|x| instanceName|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **cacheName** | **string**| The name of the Distributed Cache. | 
  **clusterName** | **string**| The name of the Cache Cluster. | 
 **optional** | ***DistributedCacheApiGetMsgVpnDistributedCacheClusterInstancesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DistributedCacheApiGetMsgVpnDistributedCacheClusterInstancesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterInstancesResponse**](MsgVpnDistributedCacheClusterInstancesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnDistributedCacheClusters**
> MsgVpnDistributedCacheClustersResponse GetMsgVpnDistributedCacheClusters(ctx, msgVpnName, cacheName, optional)
Get a list of Cache Cluster objects.

Get a list of Cache Cluster objects.  A Cache Cluster is a collection of one or more Cache Instances that subscribe to exactly the same topics. Cache Instances are grouped together in a Cache Cluster for the purpose of fault tolerance and load balancing. As published messages are received, the message broker message bus sends these live data messages to the Cache Instances in the Cache Cluster. This enables client cache requests to be served by any of Cache Instances in the Cache Cluster.   Attribute|Identifying|Deprecated :---|:---:|:---: cacheName|x| clusterName|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **cacheName** | **string**| The name of the Distributed Cache. | 
 **optional** | ***DistributedCacheApiGetMsgVpnDistributedCacheClustersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DistributedCacheApiGetMsgVpnDistributedCacheClustersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClustersResponse**](MsgVpnDistributedCacheClustersResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnDistributedCaches**
> MsgVpnDistributedCachesResponse GetMsgVpnDistributedCaches(ctx, msgVpnName, optional)
Get a list of Distributed Cache objects.

Get a list of Distributed Cache objects.  A Distributed Cache is a collection of one or more Cache Clusters that belong to the same Message VPN. Each Cache Cluster in a Distributed Cache is configured to subscribe to a different set of topics. This effectively divides up the configured topic space, to provide scaling to very large topic spaces or very high cached message throughput.   Attribute|Identifying|Deprecated :---|:---:|:---: cacheName|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***DistributedCacheApiGetMsgVpnDistributedCachesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DistributedCacheApiGetMsgVpnDistributedCachesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCachesResponse**](MsgVpnDistributedCachesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

