# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/action*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DoAboutUserLogout**](AllApi.md#DoAboutUserLogout) | **Put** /about/user/logout | Logout of the current session.
[**DoCertAuthorityRefreshCrl**](AllApi.md#DoCertAuthorityRefreshCrl) | **Put** /certAuthorities/{certAuthorityName}/refreshCrl | Refresh the CRL file for the Certificate Authority.
[**DoClientCertAuthorityRefreshCrl**](AllApi.md#DoClientCertAuthorityRefreshCrl) | **Put** /clientCertAuthorities/{certAuthorityName}/refreshCrl | Refresh the CRL file for the Client Certificate Authority.
[**DoConfigSyncAssertLeaderMsgVpn**](AllApi.md#DoConfigSyncAssertLeaderMsgVpn) | **Put** /configSyncAssertLeaderMsgVpn | Assert leadership of the specified Config Sync table, forcing any other leader&#x27;s content to be overwritten with our own. Use whenever a High Availability pair fall out of sync. Config Sync must be a leader for the selected table.
[**DoConfigSyncAssertLeaderRouter**](AllApi.md#DoConfigSyncAssertLeaderRouter) | **Put** /configSyncAssertLeaderRouter | Assert leadership of the specified Config Sync table, forcing any other leader&#x27;s content to be overwritten with our own. Use whenever a High Availability pair fall out of sync. Config Sync must be a leader for the selected table.
[**DoConfigSyncResyncFollowerMsgVpn**](AllApi.md#DoConfigSyncResyncFollowerMsgVpn) | **Put** /configSyncResyncFollowerMsgVpn | Resync the selected Config Sync table, forcing this follower&#x27;s content to be overwritten with that from a leader. Config Sync must be a follower for the selected table.
[**DoConfigSyncResyncLeaderMsgVpn**](AllApi.md#DoConfigSyncResyncLeaderMsgVpn) | **Put** /configSyncResyncLeaderMsgVpn | Resync the selected Config Sync table, forcing this leader&#x27;s content to be overwritten with that from a leader. Config Sync must be a leader for the selected table.
[**DoConfigSyncResyncLeaderRouter**](AllApi.md#DoConfigSyncResyncLeaderRouter) | **Put** /configSyncResyncLeaderRouter | Resync the selected Config Sync table, forcing this leader&#x27;s content to be overwritten with that from a leader. Config Sync must be a leader for the selected table.
[**DoGuaranteedMsgingDefragmentMsgSpoolFilesStart**](AllApi.md#DoGuaranteedMsgingDefragmentMsgSpoolFilesStart) | **Put** /guaranteedMsgingDefragmentMsgSpoolFilesStart | Start a spool file defragmentation run.
[**DoGuaranteedMsgingDefragmentMsgSpoolFilesStop**](AllApi.md#DoGuaranteedMsgingDefragmentMsgSpoolFilesStop) | **Put** /guaranteedMsgingDefragmentMsgSpoolFilesStop | Stop a spool file defragmentation run.
[**DoMsgVpnAuthenticationOauthProfileClearStats**](AllApi.md#DoMsgVpnAuthenticationOauthProfileClearStats) | **Put** /msgVpns/{msgVpnName}/authenticationOauthProfiles/{oauthProfileName}/clearStats | Clear the statistics for the OAuth Profile.
[**DoMsgVpnAuthenticationOauthProviderClearStats**](AllApi.md#DoMsgVpnAuthenticationOauthProviderClearStats) | **Put** /msgVpns/{msgVpnName}/authenticationOauthProviders/{oauthProviderName}/clearStats | Clear the statistics for the OAuth Provider.
[**DoMsgVpnBridgeClearEvent**](AllApi.md#DoMsgVpnBridgeClearEvent) | **Put** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/clearEvent | Clear an event for the Bridge so it can be generated anew.
[**DoMsgVpnBridgeClearStats**](AllApi.md#DoMsgVpnBridgeClearStats) | **Put** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/clearStats | Clear the statistics for the Bridge.
[**DoMsgVpnBridgeDisconnect**](AllApi.md#DoMsgVpnBridgeDisconnect) | **Put** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/disconnect | Disconnect the Bridge.
[**DoMsgVpnClearMsgSpoolStats**](AllApi.md#DoMsgVpnClearMsgSpoolStats) | **Put** /msgVpns/{msgVpnName}/clearMsgSpoolStats | Clear the message spool statistics for the Message VPN.
[**DoMsgVpnClearReplicationStats**](AllApi.md#DoMsgVpnClearReplicationStats) | **Put** /msgVpns/{msgVpnName}/clearReplicationStats | Clear the replication statistics for the Message VPN.
[**DoMsgVpnClearServiceStats**](AllApi.md#DoMsgVpnClearServiceStats) | **Put** /msgVpns/{msgVpnName}/clearServiceStats | Clear the service statistics for the Message VPN.
[**DoMsgVpnClearStats**](AllApi.md#DoMsgVpnClearStats) | **Put** /msgVpns/{msgVpnName}/clearStats | Clear the client statistics for the Message VPN.
[**DoMsgVpnClientClearEvent**](AllApi.md#DoMsgVpnClientClearEvent) | **Put** /msgVpns/{msgVpnName}/clients/{clientName}/clearEvent | Clear an event for the Client so it can be generated anew.
[**DoMsgVpnClientClearStats**](AllApi.md#DoMsgVpnClientClearStats) | **Put** /msgVpns/{msgVpnName}/clients/{clientName}/clearStats | Clear the statistics for the Client.
[**DoMsgVpnClientDisconnect**](AllApi.md#DoMsgVpnClientDisconnect) | **Put** /msgVpns/{msgVpnName}/clients/{clientName}/disconnect | Disconnect the Client.
[**DoMsgVpnClientTransactedSessionDelete**](AllApi.md#DoMsgVpnClientTransactedSessionDelete) | **Put** /msgVpns/{msgVpnName}/clients/{clientName}/transactedSessions/{sessionName}/delete | Delete the Transacted Session.
[**DoMsgVpnDistributedCacheClusterInstanceBackupCachedMsgs**](AllApi.md#DoMsgVpnDistributedCacheClusterInstanceBackupCachedMsgs) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/backupCachedMsgs | Backup cached messages of the Cache Instance to disk.
[**DoMsgVpnDistributedCacheClusterInstanceCancelBackupCachedMsgs**](AllApi.md#DoMsgVpnDistributedCacheClusterInstanceCancelBackupCachedMsgs) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/cancelBackupCachedMsgs | Cancel the backup of cached messages from the Cache Instance.
[**DoMsgVpnDistributedCacheClusterInstanceCancelRestoreCachedMsgs**](AllApi.md#DoMsgVpnDistributedCacheClusterInstanceCancelRestoreCachedMsgs) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/cancelRestoreCachedMsgs | Cancel the restore of cached messages to the Cache Instance.
[**DoMsgVpnDistributedCacheClusterInstanceClearEvent**](AllApi.md#DoMsgVpnDistributedCacheClusterInstanceClearEvent) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/clearEvent | Clear an event for the Cache Instance so it can be generated anew.
[**DoMsgVpnDistributedCacheClusterInstanceClearStats**](AllApi.md#DoMsgVpnDistributedCacheClusterInstanceClearStats) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/clearStats | Clear the statistics for the Cache Instance.
[**DoMsgVpnDistributedCacheClusterInstanceDeleteMsgs**](AllApi.md#DoMsgVpnDistributedCacheClusterInstanceDeleteMsgs) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/deleteMsgs | Delete messages covered by the given topic in the Cache Instance.
[**DoMsgVpnDistributedCacheClusterInstanceRestoreCachedMsgs**](AllApi.md#DoMsgVpnDistributedCacheClusterInstanceRestoreCachedMsgs) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/restoreCachedMsgs | Restore cached messages for the Cache Instance from disk.
[**DoMsgVpnDistributedCacheClusterInstanceStart**](AllApi.md#DoMsgVpnDistributedCacheClusterInstanceStart) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/start | Start the Cache Instance.
[**DoMsgVpnMqttSessionClearStats**](AllApi.md#DoMsgVpnMqttSessionClearStats) | **Put** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/clearStats | Clear the statistics for the MQTT Session.
[**DoMsgVpnQueueCancelReplay**](AllApi.md#DoMsgVpnQueueCancelReplay) | **Put** /msgVpns/{msgVpnName}/queues/{queueName}/cancelReplay | Cancel the replay of messages to the Queue.
[**DoMsgVpnQueueClearStats**](AllApi.md#DoMsgVpnQueueClearStats) | **Put** /msgVpns/{msgVpnName}/queues/{queueName}/clearStats | Clear the statistics for the Queue.
[**DoMsgVpnQueueCopyMsgFromQueue**](AllApi.md#DoMsgVpnQueueCopyMsgFromQueue) | **Put** /msgVpns/{msgVpnName}/queues/{queueName}/copyMsgFromQueue | Copy a message from another Queue to this Queue.
[**DoMsgVpnQueueCopyMsgFromReplayLog**](AllApi.md#DoMsgVpnQueueCopyMsgFromReplayLog) | **Put** /msgVpns/{msgVpnName}/queues/{queueName}/copyMsgFromReplayLog | Copy a message from a Replay Log to this Queue.
[**DoMsgVpnQueueCopyMsgFromTopicEndpoint**](AllApi.md#DoMsgVpnQueueCopyMsgFromTopicEndpoint) | **Put** /msgVpns/{msgVpnName}/queues/{queueName}/copyMsgFromTopicEndpoint | Copy a message from a Topic Endpoint to this Queue.
[**DoMsgVpnQueueDeleteMsgs**](AllApi.md#DoMsgVpnQueueDeleteMsgs) | **Put** /msgVpns/{msgVpnName}/queues/{queueName}/deleteMsgs | Delete all spooled messages from the Queue.
[**DoMsgVpnQueueMsgDelete**](AllApi.md#DoMsgVpnQueueMsgDelete) | **Put** /msgVpns/{msgVpnName}/queues/{queueName}/msgs/{msgId}/delete | Delete the Message from the Queue.
[**DoMsgVpnQueueStartReplay**](AllApi.md#DoMsgVpnQueueStartReplay) | **Put** /msgVpns/{msgVpnName}/queues/{queueName}/startReplay | Start the replay of messages to the Queue.
[**DoMsgVpnReplayLogTrimLoggedMsgs**](AllApi.md#DoMsgVpnReplayLogTrimLoggedMsgs) | **Put** /msgVpns/{msgVpnName}/replayLogs/{replayLogName}/trimLoggedMsgs | Trim (delete) messages from the Replay Log.
[**DoMsgVpnRestDeliveryPointRestConsumerClearStats**](AllApi.md#DoMsgVpnRestDeliveryPointRestConsumerClearStats) | **Put** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/clearStats | Clear the statistics for the REST Consumer.
[**DoMsgVpnTopicEndpointCancelReplay**](AllApi.md#DoMsgVpnTopicEndpointCancelReplay) | **Put** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/cancelReplay | Cancel the replay of messages to the Topic Endpoint.
[**DoMsgVpnTopicEndpointClearStats**](AllApi.md#DoMsgVpnTopicEndpointClearStats) | **Put** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/clearStats | Clear the statistics for the Topic Endpoint.
[**DoMsgVpnTopicEndpointCopyMsgFromQueue**](AllApi.md#DoMsgVpnTopicEndpointCopyMsgFromQueue) | **Put** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/copyMsgFromQueue | Copy a message from a Queue to this Topic Endpoint.
[**DoMsgVpnTopicEndpointCopyMsgFromReplayLog**](AllApi.md#DoMsgVpnTopicEndpointCopyMsgFromReplayLog) | **Put** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/copyMsgFromReplayLog | Copy a message from a Replay Log to this Topic Endpoint.
[**DoMsgVpnTopicEndpointCopyMsgFromTopicEndpoint**](AllApi.md#DoMsgVpnTopicEndpointCopyMsgFromTopicEndpoint) | **Put** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/copyMsgFromTopicEndpoint | Copy a message from another Topic Endpoint to this Topic Endpoint.
[**DoMsgVpnTopicEndpointDeleteMsgs**](AllApi.md#DoMsgVpnTopicEndpointDeleteMsgs) | **Put** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/deleteMsgs | Delete all spooled messages from the Topic Endpoint.
[**DoMsgVpnTopicEndpointMsgDelete**](AllApi.md#DoMsgVpnTopicEndpointMsgDelete) | **Put** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/msgs/{msgId}/delete | Delete the Message from the Topic Endpoint.
[**DoMsgVpnTopicEndpointStartReplay**](AllApi.md#DoMsgVpnTopicEndpointStartReplay) | **Put** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/startReplay | Start the replay of messages to the Topic Endpoint.
[**DoMsgVpnTransactionCommit**](AllApi.md#DoMsgVpnTransactionCommit) | **Put** /msgVpns/{msgVpnName}/transactions/{xid}/commit | Commit the Transaction.
[**DoMsgVpnTransactionDelete**](AllApi.md#DoMsgVpnTransactionDelete) | **Put** /msgVpns/{msgVpnName}/transactions/{xid}/delete | Delete the Transaction.
[**DoMsgVpnTransactionRollback**](AllApi.md#DoMsgVpnTransactionRollback) | **Put** /msgVpns/{msgVpnName}/transactions/{xid}/rollback | Rollback the Transaction.
[**DoOauthProfileClearStats**](AllApi.md#DoOauthProfileClearStats) | **Put** /oauthProfiles/{oauthProfileName}/clearStats | Clear the statistics for the OAuth Profile.
[**DoSessionDelete**](AllApi.md#DoSessionDelete) | **Put** /sessions/{sessionUsername},{sessionId}/delete | Delete the SEMP session.
[**GetAbout**](AllApi.md#GetAbout) | **Get** /about | Get an About object.
[**GetAboutApi**](AllApi.md#GetAboutApi) | **Get** /about/api | Get an API Description object.
[**GetAboutUser**](AllApi.md#GetAboutUser) | **Get** /about/user | Get a User object.
[**GetAboutUserMsgVpn**](AllApi.md#GetAboutUserMsgVpn) | **Get** /about/user/msgVpns/{msgVpnName} | Get a User Message VPN object.
[**GetAboutUserMsgVpns**](AllApi.md#GetAboutUserMsgVpns) | **Get** /about/user/msgVpns | Get a list of User Message VPN objects.
[**GetBroker**](AllApi.md#GetBroker) | **Get** / | Get a Broker object.
[**GetCertAuthorities**](AllApi.md#GetCertAuthorities) | **Get** /certAuthorities | Get a list of Certificate Authority objects.
[**GetCertAuthority**](AllApi.md#GetCertAuthority) | **Get** /certAuthorities/{certAuthorityName} | Get a Certificate Authority object.
[**GetClientCertAuthorities**](AllApi.md#GetClientCertAuthorities) | **Get** /clientCertAuthorities | Get a list of Client Certificate Authority objects.
[**GetClientCertAuthority**](AllApi.md#GetClientCertAuthority) | **Get** /clientCertAuthorities/{certAuthorityName} | Get a Client Certificate Authority object.
[**GetMsgVpn**](AllApi.md#GetMsgVpn) | **Get** /msgVpns/{msgVpnName} | Get a Message VPN object.
[**GetMsgVpnAuthenticationOauthProfile**](AllApi.md#GetMsgVpnAuthenticationOauthProfile) | **Get** /msgVpns/{msgVpnName}/authenticationOauthProfiles/{oauthProfileName} | Get an OAuth Profile object.
[**GetMsgVpnAuthenticationOauthProfiles**](AllApi.md#GetMsgVpnAuthenticationOauthProfiles) | **Get** /msgVpns/{msgVpnName}/authenticationOauthProfiles | Get a list of OAuth Profile objects.
[**GetMsgVpnAuthenticationOauthProvider**](AllApi.md#GetMsgVpnAuthenticationOauthProvider) | **Get** /msgVpns/{msgVpnName}/authenticationOauthProviders/{oauthProviderName} | Get an OAuth Provider object.
[**GetMsgVpnAuthenticationOauthProviders**](AllApi.md#GetMsgVpnAuthenticationOauthProviders) | **Get** /msgVpns/{msgVpnName}/authenticationOauthProviders | Get a list of OAuth Provider objects.
[**GetMsgVpnBridge**](AllApi.md#GetMsgVpnBridge) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter} | Get a Bridge object.
[**GetMsgVpnBridges**](AllApi.md#GetMsgVpnBridges) | **Get** /msgVpns/{msgVpnName}/bridges | Get a list of Bridge objects.
[**GetMsgVpnClient**](AllApi.md#GetMsgVpnClient) | **Get** /msgVpns/{msgVpnName}/clients/{clientName} | Get a Client object.
[**GetMsgVpnClientTransactedSession**](AllApi.md#GetMsgVpnClientTransactedSession) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/transactedSessions/{sessionName} | Get a Client Transacted Session object.
[**GetMsgVpnClientTransactedSessions**](AllApi.md#GetMsgVpnClientTransactedSessions) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/transactedSessions | Get a list of Client Transacted Session objects.
[**GetMsgVpnClients**](AllApi.md#GetMsgVpnClients) | **Get** /msgVpns/{msgVpnName}/clients | Get a list of Client objects.
[**GetMsgVpnDistributedCache**](AllApi.md#GetMsgVpnDistributedCache) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName} | Get a Distributed Cache object.
[**GetMsgVpnDistributedCacheCluster**](AllApi.md#GetMsgVpnDistributedCacheCluster) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName} | Get a Cache Cluster object.
[**GetMsgVpnDistributedCacheClusterInstance**](AllApi.md#GetMsgVpnDistributedCacheClusterInstance) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName} | Get a Cache Instance object.
[**GetMsgVpnDistributedCacheClusterInstances**](AllApi.md#GetMsgVpnDistributedCacheClusterInstances) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances | Get a list of Cache Instance objects.
[**GetMsgVpnDistributedCacheClusters**](AllApi.md#GetMsgVpnDistributedCacheClusters) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters | Get a list of Cache Cluster objects.
[**GetMsgVpnDistributedCaches**](AllApi.md#GetMsgVpnDistributedCaches) | **Get** /msgVpns/{msgVpnName}/distributedCaches | Get a list of Distributed Cache objects.
[**GetMsgVpnMqttSession**](AllApi.md#GetMsgVpnMqttSession) | **Get** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter} | Get an MQTT Session object.
[**GetMsgVpnMqttSessions**](AllApi.md#GetMsgVpnMqttSessions) | **Get** /msgVpns/{msgVpnName}/mqttSessions | Get a list of MQTT Session objects.
[**GetMsgVpnQueue**](AllApi.md#GetMsgVpnQueue) | **Get** /msgVpns/{msgVpnName}/queues/{queueName} | Get a Queue object.
[**GetMsgVpnQueueMsg**](AllApi.md#GetMsgVpnQueueMsg) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/msgs/{msgId} | Get a Queue Message object.
[**GetMsgVpnQueueMsgs**](AllApi.md#GetMsgVpnQueueMsgs) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/msgs | Get a list of Queue Message objects.
[**GetMsgVpnQueues**](AllApi.md#GetMsgVpnQueues) | **Get** /msgVpns/{msgVpnName}/queues | Get a list of Queue objects.
[**GetMsgVpnReplayLog**](AllApi.md#GetMsgVpnReplayLog) | **Get** /msgVpns/{msgVpnName}/replayLogs/{replayLogName} | Get a Replay Log object.
[**GetMsgVpnReplayLogs**](AllApi.md#GetMsgVpnReplayLogs) | **Get** /msgVpns/{msgVpnName}/replayLogs | Get a list of Replay Log objects.
[**GetMsgVpnRestDeliveryPoint**](AllApi.md#GetMsgVpnRestDeliveryPoint) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName} | Get a REST Delivery Point object.
[**GetMsgVpnRestDeliveryPointRestConsumer**](AllApi.md#GetMsgVpnRestDeliveryPointRestConsumer) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName} | Get a REST Consumer object.
[**GetMsgVpnRestDeliveryPointRestConsumers**](AllApi.md#GetMsgVpnRestDeliveryPointRestConsumers) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers | Get a list of REST Consumer objects.
[**GetMsgVpnRestDeliveryPoints**](AllApi.md#GetMsgVpnRestDeliveryPoints) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints | Get a list of REST Delivery Point objects.
[**GetMsgVpnTopicEndpoint**](AllApi.md#GetMsgVpnTopicEndpoint) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName} | Get a Topic Endpoint object.
[**GetMsgVpnTopicEndpointMsg**](AllApi.md#GetMsgVpnTopicEndpointMsg) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/msgs/{msgId} | Get a Topic Endpoint Message object.
[**GetMsgVpnTopicEndpointMsgs**](AllApi.md#GetMsgVpnTopicEndpointMsgs) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/msgs | Get a list of Topic Endpoint Message objects.
[**GetMsgVpnTopicEndpoints**](AllApi.md#GetMsgVpnTopicEndpoints) | **Get** /msgVpns/{msgVpnName}/topicEndpoints | Get a list of Topic Endpoint objects.
[**GetMsgVpnTransaction**](AllApi.md#GetMsgVpnTransaction) | **Get** /msgVpns/{msgVpnName}/transactions/{xid} | Get a Replicated Local Transaction or XA Transaction object.
[**GetMsgVpnTransactions**](AllApi.md#GetMsgVpnTransactions) | **Get** /msgVpns/{msgVpnName}/transactions | Get a list of Replicated Local Transaction or XA Transaction objects.
[**GetMsgVpns**](AllApi.md#GetMsgVpns) | **Get** /msgVpns | Get a list of Message VPN objects.
[**GetOauthProfile**](AllApi.md#GetOauthProfile) | **Get** /oauthProfiles/{oauthProfileName} | Get an OAuth Profile object.
[**GetOauthProfiles**](AllApi.md#GetOauthProfiles) | **Get** /oauthProfiles | Get a list of OAuth Profile objects.
[**GetSession**](AllApi.md#GetSession) | **Get** /sessions/{sessionUsername},{sessionId} | Get a SEMP Session object.
[**GetSessions**](AllApi.md#GetSessions) | **Get** /sessions | Get a list of SEMP Session objects.

# **DoAboutUserLogout**
> SempMetaOnlyResponse DoAboutUserLogout(ctx, body)
Logout of the current session.

Logout of the current session.    A SEMP client authorized with a minimum access scope/level of \"global/none\" is required to perform this operation.  This has been available since 2.21.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AboutUserLogout**](AboutUserLogout.md)| The Logout action&#x27;s attributes. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoCertAuthorityRefreshCrl**
> SempMetaOnlyResponse DoCertAuthorityRefreshCrl(ctx, body, certAuthorityName)
Refresh the CRL file for the Certificate Authority.

Refresh the CRL file for the Certificate Authority.    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been deprecated since 2.19. Replaced by clientCertAuthorities.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CertAuthorityRefreshCrl**](CertAuthorityRefreshCrl.md)| The Refresh CRL action&#x27;s attributes. | 
  **certAuthorityName** | **string**| The name of the Certificate Authority. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoClientCertAuthorityRefreshCrl**
> SempMetaOnlyResponse DoClientCertAuthorityRefreshCrl(ctx, body, certAuthorityName)
Refresh the CRL file for the Client Certificate Authority.

Refresh the CRL file for the Client Certificate Authority.    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.19.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ClientCertAuthorityRefreshCrl**](ClientCertAuthorityRefreshCrl.md)| The Refresh CRL action&#x27;s attributes. | 
  **certAuthorityName** | **string**| The name of the Certificate Authority. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoConfigSyncAssertLeaderMsgVpn**
> SempMetaOnlyResponse DoConfigSyncAssertLeaderMsgVpn(ctx, body)
Assert leadership of the specified Config Sync table, forcing any other leader's content to be overwritten with our own. Use whenever a High Availability pair fall out of sync. Config Sync must be a leader for the selected table.

Assert leadership of the specified Config Sync table, forcing any other leader's content to be overwritten with our own. Use whenever a High Availability pair fall out of sync. Config Sync must be a leader for the selected table.   Attribute|Required|Deprecated :---|:---:|:---: msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.22.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConfigSyncAssertLeaderMsgVpn**](ConfigSyncAssertLeaderMsgVpn.md)| The Assert Leader Message VPN action&#x27;s attributes. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoConfigSyncAssertLeaderRouter**
> SempMetaOnlyResponse DoConfigSyncAssertLeaderRouter(ctx, body)
Assert leadership of the specified Config Sync table, forcing any other leader's content to be overwritten with our own. Use whenever a High Availability pair fall out of sync. Config Sync must be a leader for the selected table.

Assert leadership of the specified Config Sync table, forcing any other leader's content to be overwritten with our own. Use whenever a High Availability pair fall out of sync. Config Sync must be a leader for the selected table.    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.22.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConfigSyncAssertLeaderRouter**](ConfigSyncAssertLeaderRouter.md)| The Assert Leader Router action&#x27;s attributes. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoConfigSyncResyncFollowerMsgVpn**
> SempMetaOnlyResponse DoConfigSyncResyncFollowerMsgVpn(ctx, body)
Resync the selected Config Sync table, forcing this follower's content to be overwritten with that from a leader. Config Sync must be a follower for the selected table.

Resync the selected Config Sync table, forcing this follower's content to be overwritten with that from a leader. Config Sync must be a follower for the selected table.   Attribute|Required|Deprecated :---|:---:|:---: msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.22.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConfigSyncResyncFollowerMsgVpn**](ConfigSyncResyncFollowerMsgVpn.md)| The Resync Follower Message VPN action&#x27;s attributes. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoConfigSyncResyncLeaderMsgVpn**
> SempMetaOnlyResponse DoConfigSyncResyncLeaderMsgVpn(ctx, body)
Resync the selected Config Sync table, forcing this leader's content to be overwritten with that from a leader. Config Sync must be a leader for the selected table.

Resync the selected Config Sync table, forcing this leader's content to be overwritten with that from a leader. Config Sync must be a leader for the selected table.   Attribute|Required|Deprecated :---|:---:|:---: msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.22.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConfigSyncResyncLeaderMsgVpn**](ConfigSyncResyncLeaderMsgVpn.md)| The Resync Leader Message VPN action&#x27;s attributes. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoConfigSyncResyncLeaderRouter**
> SempMetaOnlyResponse DoConfigSyncResyncLeaderRouter(ctx, body)
Resync the selected Config Sync table, forcing this leader's content to be overwritten with that from a leader. Config Sync must be a leader for the selected table.

Resync the selected Config Sync table, forcing this leader's content to be overwritten with that from a leader. Config Sync must be a leader for the selected table.    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.22.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConfigSyncResyncLeaderRouter**](ConfigSyncResyncLeaderRouter.md)| The Resync Leader Router action&#x27;s attributes. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoGuaranteedMsgingDefragmentMsgSpoolFilesStart**
> SempMetaOnlyResponse DoGuaranteedMsgingDefragmentMsgSpoolFilesStart(ctx, body)
Start a spool file defragmentation run.

Start a spool file defragmentation run.    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.18.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GuaranteedMsgingDefragmentMsgSpoolFilesStart**](GuaranteedMsgingDefragmentMsgSpoolFilesStart.md)| The Start Defragmentation action&#x27;s attributes. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoGuaranteedMsgingDefragmentMsgSpoolFilesStop**
> SempMetaOnlyResponse DoGuaranteedMsgingDefragmentMsgSpoolFilesStop(ctx, body)
Stop a spool file defragmentation run.

Stop a spool file defragmentation run.    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.18.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GuaranteedMsgingDefragmentMsgSpoolFilesStop**](GuaranteedMsgingDefragmentMsgSpoolFilesStop.md)| The Stop Defragmentation action&#x27;s attributes. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnAuthenticationOauthProfileClearStats**
> SempMetaOnlyResponse DoMsgVpnAuthenticationOauthProfileClearStats(ctx, body, msgVpnName, oauthProfileName)
Clear the statistics for the OAuth Profile.

Clear the statistics for the OAuth Profile.    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.25.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnAuthenticationOauthProfileClearStats**](MsgVpnAuthenticationOauthProfileClearStats.md)| The Clear Stats action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **oauthProfileName** | **string**| The name of the OAuth profile. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnAuthenticationOauthProviderClearStats**
> SempMetaOnlyResponse DoMsgVpnAuthenticationOauthProviderClearStats(ctx, body, msgVpnName, oauthProviderName)
Clear the statistics for the OAuth Provider.

Clear the statistics for the OAuth Provider.    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been deprecated since 2.25. authenticationOauthProviders replaced by authenticationOauthProfiles.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnAuthenticationOauthProviderClearStats**](MsgVpnAuthenticationOauthProviderClearStats.md)| The Clear Stats action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **oauthProviderName** | **string**| The name of the OAuth Provider. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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

# **DoMsgVpnClearMsgSpoolStats**
> SempMetaOnlyResponse DoMsgVpnClearMsgSpoolStats(ctx, body, msgVpnName)
Clear the message spool statistics for the Message VPN.

Clear the message spool statistics for the Message VPN.    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnClearMsgSpoolStats**](MsgVpnClearMsgSpoolStats.md)| The Clear Stats action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnClearReplicationStats**
> SempMetaOnlyResponse DoMsgVpnClearReplicationStats(ctx, body, msgVpnName)
Clear the replication statistics for the Message VPN.

Clear the replication statistics for the Message VPN.    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnClearReplicationStats**](MsgVpnClearReplicationStats.md)| The Clear Stats action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnClearServiceStats**
> SempMetaOnlyResponse DoMsgVpnClearServiceStats(ctx, body, msgVpnName)
Clear the service statistics for the Message VPN.

Clear the service statistics for the Message VPN.    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnClearServiceStats**](MsgVpnClearServiceStats.md)| The Clear Stats action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnClearStats**
> SempMetaOnlyResponse DoMsgVpnClearStats(ctx, body, msgVpnName)
Clear the client statistics for the Message VPN.

Clear the client statistics for the Message VPN.    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnClearStats**](MsgVpnClearStats.md)| The Clear Stats action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnClientClearEvent**
> SempMetaOnlyResponse DoMsgVpnClientClearEvent(ctx, body, msgVpnName, clientName)
Clear an event for the Client so it can be generated anew.

Clear an event for the Client so it can be generated anew.   Attribute|Required|Deprecated :---|:---:|:---: eventName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnClientClearEvent**](MsgVpnClientClearEvent.md)| The Clear Event action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **clientName** | **string**| The name of the Client. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnClientClearStats**
> SempMetaOnlyResponse DoMsgVpnClientClearStats(ctx, body, msgVpnName, clientName)
Clear the statistics for the Client.

Clear the statistics for the Client.    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnClientClearStats**](MsgVpnClientClearStats.md)| The Clear Stats action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **clientName** | **string**| The name of the Client. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnClientDisconnect**
> SempMetaOnlyResponse DoMsgVpnClientDisconnect(ctx, body, msgVpnName, clientName)
Disconnect the Client.

Disconnect the Client.    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnClientDisconnect**](MsgVpnClientDisconnect.md)| The Disconnect action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **clientName** | **string**| The name of the Client. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnClientTransactedSessionDelete**
> SempMetaOnlyResponse DoMsgVpnClientTransactedSessionDelete(ctx, body, msgVpnName, clientName, sessionName)
Delete the Transacted Session.

Delete the Transacted Session.    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnClientTransactedSessionDelete**](MsgVpnClientTransactedSessionDelete.md)| The Delete action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **clientName** | **string**| The name of the Client. | 
  **sessionName** | **string**| The name of the Transacted Session. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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

# **DoMsgVpnMqttSessionClearStats**
> SempMetaOnlyResponse DoMsgVpnMqttSessionClearStats(ctx, body, msgVpnName, mqttSessionClientId, mqttSessionVirtualRouter)
Clear the statistics for the MQTT Session.

Clear the statistics for the MQTT Session.    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnMqttSessionClearStats**](MsgVpnMqttSessionClearStats.md)| The Clear Stats action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **mqttSessionClientId** | **string**| The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet. | 
  **mqttSessionVirtualRouter** | **string**| The virtual router of the MQTT Session. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnQueueCancelReplay**
> SempMetaOnlyResponse DoMsgVpnQueueCancelReplay(ctx, body, msgVpnName, queueName)
Cancel the replay of messages to the Queue.

Cancel the replay of messages to the Queue.    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnQueueCancelReplay**](MsgVpnQueueCancelReplay.md)| The Cancel Replay action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The name of the Queue. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnQueueClearStats**
> SempMetaOnlyResponse DoMsgVpnQueueClearStats(ctx, body, msgVpnName, queueName)
Clear the statistics for the Queue.

Clear the statistics for the Queue.    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnQueueClearStats**](MsgVpnQueueClearStats.md)| The Clear Stats action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The name of the Queue. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnQueueCopyMsgFromQueue**
> SempMetaOnlyResponse DoMsgVpnQueueCopyMsgFromQueue(ctx, body, msgVpnName, queueName)
Copy a message from another Queue to this Queue.

Copy a message from another Queue to this Queue.   Attribute|Required|Deprecated :---|:---:|:---: replicationGroupMsgId|x| sourceQueueName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.29.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnQueueCopyMsgFromQueue**](MsgVpnQueueCopyMsgFromQueue.md)| The Copy Message From Queue action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The name of the Queue. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnQueueCopyMsgFromReplayLog**
> SempMetaOnlyResponse DoMsgVpnQueueCopyMsgFromReplayLog(ctx, body, msgVpnName, queueName)
Copy a message from a Replay Log to this Queue.

Copy a message from a Replay Log to this Queue.   Attribute|Required|Deprecated :---|:---:|:---: replicationGroupMsgId|x| sourceReplayLogName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.29.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnQueueCopyMsgFromReplayLog**](MsgVpnQueueCopyMsgFromReplayLog.md)| The Copy Message From Replay Log action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The name of the Queue. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnQueueCopyMsgFromTopicEndpoint**
> SempMetaOnlyResponse DoMsgVpnQueueCopyMsgFromTopicEndpoint(ctx, body, msgVpnName, queueName)
Copy a message from a Topic Endpoint to this Queue.

Copy a message from a Topic Endpoint to this Queue.   Attribute|Required|Deprecated :---|:---:|:---: replicationGroupMsgId|x| sourceTopicEndpointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.29.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnQueueCopyMsgFromTopicEndpoint**](MsgVpnQueueCopyMsgFromTopicEndpoint.md)| The Copy Message From Topic Endpoint action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The name of the Queue. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnQueueDeleteMsgs**
> SempMetaOnlyResponse DoMsgVpnQueueDeleteMsgs(ctx, body, msgVpnName, queueName)
Delete all spooled messages from the Queue.

Delete all spooled messages from the Queue.    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.28.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnQueueDeleteMsgs**](MsgVpnQueueDeleteMsgs.md)| The Delete All Messages action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The name of the Queue. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnQueueMsgDelete**
> SempMetaOnlyResponse DoMsgVpnQueueMsgDelete(ctx, body, msgVpnName, queueName, msgId)
Delete the Message from the Queue.

Delete the Message from the Queue.    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnQueueMsgDelete**](MsgVpnQueueMsgDelete.md)| The Delete action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The name of the Queue. | 
  **msgId** | **string**| The identifier (ID) of the Message. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnQueueStartReplay**
> SempMetaOnlyResponse DoMsgVpnQueueStartReplay(ctx, body, msgVpnName, queueName)
Start the replay of messages to the Queue.

Start the replay of messages to the Queue.    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnQueueStartReplay**](MsgVpnQueueStartReplay.md)| The Start Replay action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The name of the Queue. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnReplayLogTrimLoggedMsgs**
> SempMetaOnlyResponse DoMsgVpnReplayLogTrimLoggedMsgs(ctx, body, msgVpnName, replayLogName)
Trim (delete) messages from the Replay Log.

Trim (delete) messages from the Replay Log.   Attribute|Required|Deprecated :---|:---:|:---: olderThanTime|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnReplayLogTrimLoggedMsgs**](MsgVpnReplayLogTrimLoggedMsgs.md)| The Trim Logged Messages action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **replayLogName** | **string**| The name of the Replay Log. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnRestDeliveryPointRestConsumerClearStats**
> SempMetaOnlyResponse DoMsgVpnRestDeliveryPointRestConsumerClearStats(ctx, body, msgVpnName, restDeliveryPointName, restConsumerName)
Clear the statistics for the REST Consumer.

Clear the statistics for the REST Consumer.    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnRestDeliveryPointRestConsumerClearStats**](MsgVpnRestDeliveryPointRestConsumerClearStats.md)| The Clear Stats action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **restConsumerName** | **string**| The name of the REST Consumer. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnTopicEndpointCancelReplay**
> SempMetaOnlyResponse DoMsgVpnTopicEndpointCancelReplay(ctx, body, msgVpnName, topicEndpointName)
Cancel the replay of messages to the Topic Endpoint.

Cancel the replay of messages to the Topic Endpoint.    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnTopicEndpointCancelReplay**](MsgVpnTopicEndpointCancelReplay.md)| The Cancel Replay action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointName** | **string**| The name of the Topic Endpoint. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnTopicEndpointClearStats**
> SempMetaOnlyResponse DoMsgVpnTopicEndpointClearStats(ctx, body, msgVpnName, topicEndpointName)
Clear the statistics for the Topic Endpoint.

Clear the statistics for the Topic Endpoint.    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnTopicEndpointClearStats**](MsgVpnTopicEndpointClearStats.md)| The Clear Stats action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointName** | **string**| The name of the Topic Endpoint. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnTopicEndpointCopyMsgFromQueue**
> SempMetaOnlyResponse DoMsgVpnTopicEndpointCopyMsgFromQueue(ctx, body, msgVpnName, topicEndpointName)
Copy a message from a Queue to this Topic Endpoint.

Copy a message from a Queue to this Topic Endpoint.   Attribute|Required|Deprecated :---|:---:|:---: replicationGroupMsgId|x| sourceQueueName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.29.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnTopicEndpointCopyMsgFromQueue**](MsgVpnTopicEndpointCopyMsgFromQueue.md)| The Copy Message From Queue action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointName** | **string**| The name of the Topic Endpoint. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnTopicEndpointCopyMsgFromReplayLog**
> SempMetaOnlyResponse DoMsgVpnTopicEndpointCopyMsgFromReplayLog(ctx, body, msgVpnName, topicEndpointName)
Copy a message from a Replay Log to this Topic Endpoint.

Copy a message from a Replay Log to this Topic Endpoint.   Attribute|Required|Deprecated :---|:---:|:---: replicationGroupMsgId|x| sourceReplayLogName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.29.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnTopicEndpointCopyMsgFromReplayLog**](MsgVpnTopicEndpointCopyMsgFromReplayLog.md)| The Copy Message From Replay Log action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointName** | **string**| The name of the Topic Endpoint. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnTopicEndpointCopyMsgFromTopicEndpoint**
> SempMetaOnlyResponse DoMsgVpnTopicEndpointCopyMsgFromTopicEndpoint(ctx, body, msgVpnName, topicEndpointName)
Copy a message from another Topic Endpoint to this Topic Endpoint.

Copy a message from another Topic Endpoint to this Topic Endpoint.   Attribute|Required|Deprecated :---|:---:|:---: replicationGroupMsgId|x| sourceTopicEndpointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.29.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnTopicEndpointCopyMsgFromTopicEndpoint**](MsgVpnTopicEndpointCopyMsgFromTopicEndpoint.md)| The Copy Message From Topic Endpoint action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointName** | **string**| The name of the Topic Endpoint. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnTopicEndpointDeleteMsgs**
> SempMetaOnlyResponse DoMsgVpnTopicEndpointDeleteMsgs(ctx, body, msgVpnName, topicEndpointName)
Delete all spooled messages from the Topic Endpoint.

Delete all spooled messages from the Topic Endpoint.    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.28.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnTopicEndpointDeleteMsgs**](MsgVpnTopicEndpointDeleteMsgs.md)| The Delete All Messages action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointName** | **string**| The name of the Topic Endpoint. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnTopicEndpointMsgDelete**
> SempMetaOnlyResponse DoMsgVpnTopicEndpointMsgDelete(ctx, body, msgVpnName, topicEndpointName, msgId)
Delete the Message from the Topic Endpoint.

Delete the Message from the Topic Endpoint.    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnTopicEndpointMsgDelete**](MsgVpnTopicEndpointMsgDelete.md)| The Delete action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointName** | **string**| The name of the Topic Endpoint. | 
  **msgId** | **string**| The identifier (ID) of the Message. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoMsgVpnTopicEndpointStartReplay**
> SempMetaOnlyResponse DoMsgVpnTopicEndpointStartReplay(ctx, body, msgVpnName, topicEndpointName)
Start the replay of messages to the Topic Endpoint.

Start the replay of messages to the Topic Endpoint.    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnTopicEndpointStartReplay**](MsgVpnTopicEndpointStartReplay.md)| The Start Replay action&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointName** | **string**| The name of the Topic Endpoint. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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

# **DoOauthProfileClearStats**
> SempMetaOnlyResponse DoOauthProfileClearStats(ctx, body, oauthProfileName)
Clear the statistics for the OAuth Profile.

Clear the statistics for the OAuth Profile.    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OauthProfileClearStats**](OauthProfileClearStats.md)| The Clear Stats action&#x27;s attributes. | 
  **oauthProfileName** | **string**| The name of the OAuth profile. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoSessionDelete**
> SempMetaOnlyResponse DoSessionDelete(ctx, body, sessionUsername, sessionId)
Delete the SEMP session.

Delete the SEMP session.    A SEMP client authorized with a minimum access scope/level of \"global/admin\" is required to perform this operation.  This has been available since 2.21.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SessionDelete**](SessionDelete.md)| The Delete action&#x27;s attributes. | 
  **sessionUsername** | **string**| The username used for authorization. | 
  **sessionId** | **string**| The unique identifier for the session. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAbout**
> AboutResponse GetAbout(ctx, optional)
Get an About object.

Get an About object.  This provides metadata about the SEMP API, such as the version of the API supported by the broker.    A SEMP client authorized with a minimum access scope/level of \"global/none\" is required to perform this operation.  This has been available since 2.13.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AllApiGetAboutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetAboutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**AboutResponse**](AboutResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAboutApi**
> AboutApiResponse GetAboutApi(ctx, optional)
Get an API Description object.

Get an API Description object.  The API Description object provides metadata about the SEMP API.    A SEMP client authorized with a minimum access scope/level of \"global/none\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AllApiGetAboutApiOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetAboutApiOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**AboutApiResponse**](AboutApiResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAboutUser**
> AboutUserResponse GetAboutUser(ctx, optional)
Get a User object.

Get a User object.  Session and access level information about the user accessing the SEMP API.    A SEMP client authorized with a minimum access scope/level of \"global/none\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AllApiGetAboutUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetAboutUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**AboutUserResponse**](AboutUserResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAboutUserMsgVpn**
> AboutUserMsgVpnResponse GetAboutUserMsgVpn(ctx, msgVpnName, optional)
Get a User Message VPN object.

Get a User Message VPN object.  This provides information about the Message VPN access level for the username used to access the SEMP API.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"global/none\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***AllApiGetAboutUserMsgVpnOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetAboutUserMsgVpnOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**AboutUserMsgVpnResponse**](AboutUserMsgVpnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAboutUserMsgVpns**
> AboutUserMsgVpnsResponse GetAboutUserMsgVpns(ctx, optional)
Get a list of User Message VPN objects.

Get a list of User Message VPN objects.  This provides information about the Message VPN access level for the username used to access the SEMP API.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"global/none\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AllApiGetAboutUserMsgVpnsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetAboutUserMsgVpnsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**AboutUserMsgVpnsResponse**](AboutUserMsgVpnsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBroker**
> BrokerResponse GetBroker(ctx, optional)
Get a Broker object.

Get a Broker object.  This object contains global configuration for the message broker.    A SEMP client authorized with a minimum access scope/level of \"global/none\" is required to perform this operation.  This has been available since 2.13.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AllApiGetBrokerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetBrokerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**BrokerResponse**](BrokerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCertAuthorities**
> CertAuthoritiesResponse GetCertAuthorities(ctx, optional)
Get a list of Certificate Authority objects.

Get a list of Certificate Authority objects.  Clients can authenticate with the message broker over TLS by presenting a valid client certificate. The message broker authenticates the client certificate by constructing a full certificate chain (from the client certificate to intermediate CAs to a configured root CA). The intermediate CAs in this chain can be provided by the client, or configured in the message broker. The root CA must be configured on the message broker.   Attribute|Identifying|Deprecated :---|:---:|:---: certAuthorityName|x|x    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been deprecated since 2.19. Replaced by clientCertAuthorities and domainCertAuthorities.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AllApiGetCertAuthoritiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetCertAuthoritiesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**CertAuthoritiesResponse**](CertAuthoritiesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCertAuthority**
> CertAuthorityResponse GetCertAuthority(ctx, certAuthorityName, optional)
Get a Certificate Authority object.

Get a Certificate Authority object.  Clients can authenticate with the message broker over TLS by presenting a valid client certificate. The message broker authenticates the client certificate by constructing a full certificate chain (from the client certificate to intermediate CAs to a configured root CA). The intermediate CAs in this chain can be provided by the client, or configured in the message broker. The root CA must be configured on the message broker.   Attribute|Identifying|Deprecated :---|:---:|:---: certAuthorityName|x|x    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been deprecated since 2.19. Replaced by clientCertAuthorities and domainCertAuthorities.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **certAuthorityName** | **string**| The name of the Certificate Authority. | 
 **optional** | ***AllApiGetCertAuthorityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetCertAuthorityOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**CertAuthorityResponse**](CertAuthorityResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClientCertAuthorities**
> ClientCertAuthoritiesResponse GetClientCertAuthorities(ctx, optional)
Get a list of Client Certificate Authority objects.

Get a list of Client Certificate Authority objects.  Clients can authenticate with the message broker over TLS by presenting a valid client certificate. The message broker authenticates the client certificate by constructing a full certificate chain (from the client certificate to intermediate CAs to a configured root CA). The intermediate CAs in this chain can be provided by the client, or configured in the message broker. The root CA must be configured on the message broker.   Attribute|Identifying|Deprecated :---|:---:|:---: certAuthorityName|x|    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.19.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AllApiGetClientCertAuthoritiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetClientCertAuthoritiesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**ClientCertAuthoritiesResponse**](ClientCertAuthoritiesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClientCertAuthority**
> ClientCertAuthorityResponse GetClientCertAuthority(ctx, certAuthorityName, optional)
Get a Client Certificate Authority object.

Get a Client Certificate Authority object.  Clients can authenticate with the message broker over TLS by presenting a valid client certificate. The message broker authenticates the client certificate by constructing a full certificate chain (from the client certificate to intermediate CAs to a configured root CA). The intermediate CAs in this chain can be provided by the client, or configured in the message broker. The root CA must be configured on the message broker.   Attribute|Identifying|Deprecated :---|:---:|:---: certAuthorityName|x|    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.19.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **certAuthorityName** | **string**| The name of the Certificate Authority. | 
 **optional** | ***AllApiGetClientCertAuthorityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetClientCertAuthorityOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**ClientCertAuthorityResponse**](ClientCertAuthorityResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpn**
> MsgVpnResponse GetMsgVpn(ctx, msgVpnName, optional)
Get a Message VPN object.

Get a Message VPN object.  Message VPNs (Virtual Private Networks) allow for the segregation of topic space and clients. They also group clients connecting to a network of message brokers, such that messages published within a particular group are only visible to that group's clients.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***AllApiGetMsgVpnOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetMsgVpnOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnResponse**](MsgVpnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAuthenticationOauthProfile**
> MsgVpnAuthenticationOauthProfileResponse GetMsgVpnAuthenticationOauthProfile(ctx, msgVpnName, oauthProfileName, optional)
Get an OAuth Profile object.

Get an OAuth Profile object.  OAuth profiles specify how to securely authenticate to an OAuth provider.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| oauthProfileName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.25.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
 **optional** | ***AllApiGetMsgVpnAuthenticationOauthProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetMsgVpnAuthenticationOauthProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthenticationOauthProfileResponse**](MsgVpnAuthenticationOauthProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAuthenticationOauthProfiles**
> MsgVpnAuthenticationOauthProfilesResponse GetMsgVpnAuthenticationOauthProfiles(ctx, msgVpnName, optional)
Get a list of OAuth Profile objects.

Get a list of OAuth Profile objects.  OAuth profiles specify how to securely authenticate to an OAuth provider.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| oauthProfileName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.25.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***AllApiGetMsgVpnAuthenticationOauthProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetMsgVpnAuthenticationOauthProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthenticationOauthProfilesResponse**](MsgVpnAuthenticationOauthProfilesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAuthenticationOauthProvider**
> MsgVpnAuthenticationOauthProviderResponse GetMsgVpnAuthenticationOauthProvider(ctx, msgVpnName, oauthProviderName, optional)
Get an OAuth Provider object.

Get an OAuth Provider object.  OAuth Providers contain information about the issuer of an OAuth token that is needed to validate the token and derive a client username from it.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x|x oauthProviderName|x|x    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been deprecated since 2.25. Replaced by authenticationOauthProfiles.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **oauthProviderName** | **string**| The name of the OAuth Provider. | 
 **optional** | ***AllApiGetMsgVpnAuthenticationOauthProviderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetMsgVpnAuthenticationOauthProviderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthenticationOauthProviderResponse**](MsgVpnAuthenticationOauthProviderResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAuthenticationOauthProviders**
> MsgVpnAuthenticationOauthProvidersResponse GetMsgVpnAuthenticationOauthProviders(ctx, msgVpnName, optional)
Get a list of OAuth Provider objects.

Get a list of OAuth Provider objects.  OAuth Providers contain information about the issuer of an OAuth token that is needed to validate the token and derive a client username from it.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x|x oauthProviderName|x|x    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been deprecated since 2.25. Replaced by authenticationOauthProfiles.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***AllApiGetMsgVpnAuthenticationOauthProvidersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetMsgVpnAuthenticationOauthProvidersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthenticationOauthProvidersResponse**](MsgVpnAuthenticationOauthProvidersResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
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
 **optional** | ***AllApiGetMsgVpnBridgeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetMsgVpnBridgeOpts struct
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
 **optional** | ***AllApiGetMsgVpnBridgesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetMsgVpnBridgesOpts struct
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

# **GetMsgVpnClient**
> MsgVpnClientResponse GetMsgVpnClient(ctx, msgVpnName, clientName, optional)
Get a Client object.

Get a Client object.  Applications or devices that connect to message brokers to send and/or receive messages are represented as Clients.   Attribute|Identifying|Deprecated :---|:---:|:---: clientName|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **clientName** | **string**| The name of the Client. | 
 **optional** | ***AllApiGetMsgVpnClientOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetMsgVpnClientOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientResponse**](MsgVpnClientResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnClientTransactedSession**
> MsgVpnClientTransactedSessionResponse GetMsgVpnClientTransactedSession(ctx, msgVpnName, clientName, sessionName, optional)
Get a Client Transacted Session object.

Get a Client Transacted Session object.  Transacted Sessions enable clients to group multiple message send and/or receive operations together in single, atomic units known as local transactions.   Attribute|Identifying|Deprecated :---|:---:|:---: clientName|x| msgVpnName|x| sessionName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **clientName** | **string**| The name of the Client. | 
  **sessionName** | **string**| The name of the Transacted Session. | 
 **optional** | ***AllApiGetMsgVpnClientTransactedSessionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetMsgVpnClientTransactedSessionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientTransactedSessionResponse**](MsgVpnClientTransactedSessionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnClientTransactedSessions**
> MsgVpnClientTransactedSessionsResponse GetMsgVpnClientTransactedSessions(ctx, msgVpnName, clientName, optional)
Get a list of Client Transacted Session objects.

Get a list of Client Transacted Session objects.  Transacted Sessions enable clients to group multiple message send and/or receive operations together in single, atomic units known as local transactions.   Attribute|Identifying|Deprecated :---|:---:|:---: clientName|x| msgVpnName|x| sessionName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **clientName** | **string**| The name of the Client. | 
 **optional** | ***AllApiGetMsgVpnClientTransactedSessionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetMsgVpnClientTransactedSessionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientTransactedSessionsResponse**](MsgVpnClientTransactedSessionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnClients**
> MsgVpnClientsResponse GetMsgVpnClients(ctx, msgVpnName, optional)
Get a list of Client objects.

Get a list of Client objects.  Applications or devices that connect to message brokers to send and/or receive messages are represented as Clients.   Attribute|Identifying|Deprecated :---|:---:|:---: clientName|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***AllApiGetMsgVpnClientsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetMsgVpnClientsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientsResponse**](MsgVpnClientsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
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
 **optional** | ***AllApiGetMsgVpnDistributedCacheOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetMsgVpnDistributedCacheOpts struct
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
 **optional** | ***AllApiGetMsgVpnDistributedCacheClusterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetMsgVpnDistributedCacheClusterOpts struct
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
 **optional** | ***AllApiGetMsgVpnDistributedCacheClusterInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetMsgVpnDistributedCacheClusterInstanceOpts struct
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
 **optional** | ***AllApiGetMsgVpnDistributedCacheClusterInstancesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetMsgVpnDistributedCacheClusterInstancesOpts struct
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
 **optional** | ***AllApiGetMsgVpnDistributedCacheClustersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetMsgVpnDistributedCacheClustersOpts struct
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
 **optional** | ***AllApiGetMsgVpnDistributedCachesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetMsgVpnDistributedCachesOpts struct
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

# **GetMsgVpnMqttSession**
> MsgVpnMqttSessionResponse GetMsgVpnMqttSession(ctx, msgVpnName, mqttSessionClientId, mqttSessionVirtualRouter, optional)
Get an MQTT Session object.

Get an MQTT Session object.  An MQTT Session object is a virtual representation of an MQTT client connection. An MQTT session holds the state of an MQTT client (that is, it is used to contain a client's QoS 0 and QoS 1 subscription sets and any undelivered QoS 1 messages).   Attribute|Identifying|Deprecated :---|:---:|:---: mqttSessionClientId|x| mqttSessionVirtualRouter|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **mqttSessionClientId** | **string**| The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet. | 
  **mqttSessionVirtualRouter** | **string**| The virtual router of the MQTT Session. | 
 **optional** | ***AllApiGetMsgVpnMqttSessionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetMsgVpnMqttSessionOpts struct
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

# **GetMsgVpnMqttSessions**
> MsgVpnMqttSessionsResponse GetMsgVpnMqttSessions(ctx, msgVpnName, optional)
Get a list of MQTT Session objects.

Get a list of MQTT Session objects.  An MQTT Session object is a virtual representation of an MQTT client connection. An MQTT session holds the state of an MQTT client (that is, it is used to contain a client's QoS 0 and QoS 1 subscription sets and any undelivered QoS 1 messages).   Attribute|Identifying|Deprecated :---|:---:|:---: mqttSessionClientId|x| mqttSessionVirtualRouter|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***AllApiGetMsgVpnMqttSessionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetMsgVpnMqttSessionsOpts struct
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

# **GetMsgVpnQueue**
> MsgVpnQueueResponse GetMsgVpnQueue(ctx, msgVpnName, queueName, optional)
Get a Queue object.

Get a Queue object.  A Queue acts as both a destination that clients can publish messages to, and as an endpoint that clients can bind consumers to and consume messages from.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| queueName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The name of the Queue. | 
 **optional** | ***AllApiGetMsgVpnQueueOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetMsgVpnQueueOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueueResponse**](MsgVpnQueueResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnQueueMsg**
> MsgVpnQueueMsgResponse GetMsgVpnQueueMsg(ctx, msgVpnName, queueName, msgId, optional)
Get a Queue Message object.

Get a Queue Message object.  A Queue Message is a packet of information sent from producers to consumers using the Queue.   Attribute|Identifying|Deprecated :---|:---:|:---: msgId|x| msgVpnName|x| queueName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The name of the Queue. | 
  **msgId** | **string**| The identifier (ID) of the Message. | 
 **optional** | ***AllApiGetMsgVpnQueueMsgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetMsgVpnQueueMsgOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueueMsgResponse**](MsgVpnQueueMsgResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnQueueMsgs**
> MsgVpnQueueMsgsResponse GetMsgVpnQueueMsgs(ctx, msgVpnName, queueName, optional)
Get a list of Queue Message objects.

Get a list of Queue Message objects.  A Queue Message is a packet of information sent from producers to consumers using the Queue.   Attribute|Identifying|Deprecated :---|:---:|:---: msgId|x| msgVpnName|x| queueName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The name of the Queue. | 
 **optional** | ***AllApiGetMsgVpnQueueMsgsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetMsgVpnQueueMsgsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueueMsgsResponse**](MsgVpnQueueMsgsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnQueues**
> MsgVpnQueuesResponse GetMsgVpnQueues(ctx, msgVpnName, optional)
Get a list of Queue objects.

Get a list of Queue objects.  A Queue acts as both a destination that clients can publish messages to, and as an endpoint that clients can bind consumers to and consume messages from.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| queueName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***AllApiGetMsgVpnQueuesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetMsgVpnQueuesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueuesResponse**](MsgVpnQueuesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnReplayLog**
> MsgVpnReplayLogResponse GetMsgVpnReplayLog(ctx, msgVpnName, replayLogName, optional)
Get a Replay Log object.

Get a Replay Log object.  When the Message Replay feature is enabled, message brokers store persistent messages in a Replay Log. These messages are kept until the log is full, after which the oldest messages are removed to free up space for new messages.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| replayLogName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **replayLogName** | **string**| The name of the Replay Log. | 
 **optional** | ***AllApiGetMsgVpnReplayLogOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetMsgVpnReplayLogOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnReplayLogResponse**](MsgVpnReplayLogResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnReplayLogs**
> MsgVpnReplayLogsResponse GetMsgVpnReplayLogs(ctx, msgVpnName, optional)
Get a list of Replay Log objects.

Get a list of Replay Log objects.  When the Message Replay feature is enabled, message brokers store persistent messages in a Replay Log. These messages are kept until the log is full, after which the oldest messages are removed to free up space for new messages.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| replayLogName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***AllApiGetMsgVpnReplayLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetMsgVpnReplayLogsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnReplayLogsResponse**](MsgVpnReplayLogsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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
 **optional** | ***AllApiGetMsgVpnRestDeliveryPointOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetMsgVpnRestDeliveryPointOpts struct
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

# **GetMsgVpnRestDeliveryPointRestConsumer**
> MsgVpnRestDeliveryPointRestConsumerResponse GetMsgVpnRestDeliveryPointRestConsumer(ctx, msgVpnName, restDeliveryPointName, restConsumerName, optional)
Get a REST Consumer object.

Get a REST Consumer object.  REST Consumer objects establish HTTP connectivity to REST consumer applications who wish to receive messages from a broker.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| restConsumerName|x| restDeliveryPointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **restConsumerName** | **string**| The name of the REST Consumer. | 
 **optional** | ***AllApiGetMsgVpnRestDeliveryPointRestConsumerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetMsgVpnRestDeliveryPointRestConsumerOpts struct
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

# **GetMsgVpnRestDeliveryPointRestConsumers**
> MsgVpnRestDeliveryPointRestConsumersResponse GetMsgVpnRestDeliveryPointRestConsumers(ctx, msgVpnName, restDeliveryPointName, optional)
Get a list of REST Consumer objects.

Get a list of REST Consumer objects.  REST Consumer objects establish HTTP connectivity to REST consumer applications who wish to receive messages from a broker.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| restConsumerName|x| restDeliveryPointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
 **optional** | ***AllApiGetMsgVpnRestDeliveryPointRestConsumersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetMsgVpnRestDeliveryPointRestConsumersOpts struct
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
 **optional** | ***AllApiGetMsgVpnRestDeliveryPointsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetMsgVpnRestDeliveryPointsOpts struct
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

# **GetMsgVpnTopicEndpoint**
> MsgVpnTopicEndpointResponse GetMsgVpnTopicEndpoint(ctx, msgVpnName, topicEndpointName, optional)
Get a Topic Endpoint object.

Get a Topic Endpoint object.  A Topic Endpoint attracts messages published to a topic for which the Topic Endpoint has a matching topic subscription. The topic subscription for the Topic Endpoint is specified in the client request to bind a Flow to that Topic Endpoint. Queues are significantly more flexible than Topic Endpoints and are the recommended approach for most applications. The use of Topic Endpoints should be restricted to JMS applications.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| topicEndpointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointName** | **string**| The name of the Topic Endpoint. | 
 **optional** | ***AllApiGetMsgVpnTopicEndpointOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetMsgVpnTopicEndpointOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointResponse**](MsgVpnTopicEndpointResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnTopicEndpointMsg**
> MsgVpnTopicEndpointMsgResponse GetMsgVpnTopicEndpointMsg(ctx, msgVpnName, topicEndpointName, msgId, optional)
Get a Topic Endpoint Message object.

Get a Topic Endpoint Message object.  A Topic Endpoint Message is a packet of information sent from producers to consumers using the Topic Endpoint.   Attribute|Identifying|Deprecated :---|:---:|:---: msgId|x| msgVpnName|x| topicEndpointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointName** | **string**| The name of the Topic Endpoint. | 
  **msgId** | **string**| The identifier (ID) of the Message. | 
 **optional** | ***AllApiGetMsgVpnTopicEndpointMsgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetMsgVpnTopicEndpointMsgOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointMsgResponse**](MsgVpnTopicEndpointMsgResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnTopicEndpointMsgs**
> MsgVpnTopicEndpointMsgsResponse GetMsgVpnTopicEndpointMsgs(ctx, msgVpnName, topicEndpointName, optional)
Get a list of Topic Endpoint Message objects.

Get a list of Topic Endpoint Message objects.  A Topic Endpoint Message is a packet of information sent from producers to consumers using the Topic Endpoint.   Attribute|Identifying|Deprecated :---|:---:|:---: msgId|x| msgVpnName|x| topicEndpointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointName** | **string**| The name of the Topic Endpoint. | 
 **optional** | ***AllApiGetMsgVpnTopicEndpointMsgsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetMsgVpnTopicEndpointMsgsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointMsgsResponse**](MsgVpnTopicEndpointMsgsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnTopicEndpoints**
> MsgVpnTopicEndpointsResponse GetMsgVpnTopicEndpoints(ctx, msgVpnName, optional)
Get a list of Topic Endpoint objects.

Get a list of Topic Endpoint objects.  A Topic Endpoint attracts messages published to a topic for which the Topic Endpoint has a matching topic subscription. The topic subscription for the Topic Endpoint is specified in the client request to bind a Flow to that Topic Endpoint. Queues are significantly more flexible than Topic Endpoints and are the recommended approach for most applications. The use of Topic Endpoints should be restricted to JMS applications.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| topicEndpointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***AllApiGetMsgVpnTopicEndpointsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetMsgVpnTopicEndpointsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointsResponse**](MsgVpnTopicEndpointsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
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
 **optional** | ***AllApiGetMsgVpnTransactionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetMsgVpnTransactionOpts struct
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
 **optional** | ***AllApiGetMsgVpnTransactionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetMsgVpnTransactionsOpts struct
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

# **GetMsgVpns**
> MsgVpnsResponse GetMsgVpns(ctx, optional)
Get a list of Message VPN objects.

Get a list of Message VPN objects.  Message VPNs (Virtual Private Networks) allow for the segregation of topic space and clients. They also group clients connecting to a network of message brokers, such that messages published within a particular group are only visible to that group's clients.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AllApiGetMsgVpnsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetMsgVpnsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnsResponse**](MsgVpnsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOauthProfile**
> OauthProfileResponse GetOauthProfile(ctx, oauthProfileName, optional)
Get an OAuth Profile object.

Get an OAuth Profile object.  OAuth profiles specify how to securely authenticate to an OAuth provider.   Attribute|Identifying|Deprecated :---|:---:|:---: oauthProfileName|x|    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
 **optional** | ***AllApiGetOauthProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetOauthProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**OauthProfileResponse**](OauthProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOauthProfiles**
> OauthProfilesResponse GetOauthProfiles(ctx, optional)
Get a list of OAuth Profile objects.

Get a list of OAuth Profile objects.  OAuth profiles specify how to securely authenticate to an OAuth provider.   Attribute|Identifying|Deprecated :---|:---:|:---: oauthProfileName|x|    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.24.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AllApiGetOauthProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetOauthProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**OauthProfilesResponse**](OauthProfilesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSession**
> SessionResponse GetSession(ctx, sessionUsername, sessionId, optional)
Get a SEMP Session object.

Get a SEMP Session object.  Administrative sessions for configuration and monitoring.   Attribute|Identifying|Deprecated :---|:---:|:---: sessionId|x| sessionUsername|x|    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.21.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sessionUsername** | **string**| The username used for authorization. | 
  **sessionId** | **string**| The unique identifier for the session. | 
 **optional** | ***AllApiGetSessionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetSessionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**SessionResponse**](SessionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSessions**
> SessionsResponse GetSessions(ctx, optional)
Get a list of SEMP Session objects.

Get a list of SEMP Session objects.  Administrative sessions for configuration and monitoring.   Attribute|Identifying|Deprecated :---|:---:|:---: sessionId|x| sessionUsername|x|    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.21.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AllApiGetSessionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllApiGetSessionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**SessionsResponse**](SessionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

