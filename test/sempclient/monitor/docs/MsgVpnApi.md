# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/monitor*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMsgVpn**](MsgVpnApi.md#GetMsgVpn) | **Get** /msgVpns/{msgVpnName} | Get a Message VPN object.
[**GetMsgVpnAclProfile**](MsgVpnApi.md#GetMsgVpnAclProfile) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName} | Get an ACL Profile object.
[**GetMsgVpnAclProfileClientConnectException**](MsgVpnApi.md#GetMsgVpnAclProfileClientConnectException) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/clientConnectExceptions/{clientConnectExceptionAddress} | Get a Client Connect Exception object.
[**GetMsgVpnAclProfileClientConnectExceptions**](MsgVpnApi.md#GetMsgVpnAclProfileClientConnectExceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/clientConnectExceptions | Get a list of Client Connect Exception objects.
[**GetMsgVpnAclProfilePublishException**](MsgVpnApi.md#GetMsgVpnAclProfilePublishException) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishExceptions/{topicSyntax},{publishExceptionTopic} | Get a Publish Topic Exception object.
[**GetMsgVpnAclProfilePublishExceptions**](MsgVpnApi.md#GetMsgVpnAclProfilePublishExceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishExceptions | Get a list of Publish Topic Exception objects.
[**GetMsgVpnAclProfilePublishTopicException**](MsgVpnApi.md#GetMsgVpnAclProfilePublishTopicException) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishTopicExceptions/{publishTopicExceptionSyntax},{publishTopicException} | Get a Publish Topic Exception object.
[**GetMsgVpnAclProfilePublishTopicExceptions**](MsgVpnApi.md#GetMsgVpnAclProfilePublishTopicExceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishTopicExceptions | Get a list of Publish Topic Exception objects.
[**GetMsgVpnAclProfileSubscribeException**](MsgVpnApi.md#GetMsgVpnAclProfileSubscribeException) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeExceptions/{topicSyntax},{subscribeExceptionTopic} | Get a Subscribe Topic Exception object.
[**GetMsgVpnAclProfileSubscribeExceptions**](MsgVpnApi.md#GetMsgVpnAclProfileSubscribeExceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeExceptions | Get a list of Subscribe Topic Exception objects.
[**GetMsgVpnAclProfileSubscribeShareNameException**](MsgVpnApi.md#GetMsgVpnAclProfileSubscribeShareNameException) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeShareNameExceptions/{subscribeShareNameExceptionSyntax},{subscribeShareNameException} | Get a Subscribe Share Name Exception object.
[**GetMsgVpnAclProfileSubscribeShareNameExceptions**](MsgVpnApi.md#GetMsgVpnAclProfileSubscribeShareNameExceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeShareNameExceptions | Get a list of Subscribe Share Name Exception objects.
[**GetMsgVpnAclProfileSubscribeTopicException**](MsgVpnApi.md#GetMsgVpnAclProfileSubscribeTopicException) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeTopicExceptions/{subscribeTopicExceptionSyntax},{subscribeTopicException} | Get a Subscribe Topic Exception object.
[**GetMsgVpnAclProfileSubscribeTopicExceptions**](MsgVpnApi.md#GetMsgVpnAclProfileSubscribeTopicExceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeTopicExceptions | Get a list of Subscribe Topic Exception objects.
[**GetMsgVpnAclProfiles**](MsgVpnApi.md#GetMsgVpnAclProfiles) | **Get** /msgVpns/{msgVpnName}/aclProfiles | Get a list of ACL Profile objects.
[**GetMsgVpnAuthenticationOauthProfile**](MsgVpnApi.md#GetMsgVpnAuthenticationOauthProfile) | **Get** /msgVpns/{msgVpnName}/authenticationOauthProfiles/{oauthProfileName} | Get an OAuth Profile object.
[**GetMsgVpnAuthenticationOauthProfileClientRequiredClaim**](MsgVpnApi.md#GetMsgVpnAuthenticationOauthProfileClientRequiredClaim) | **Get** /msgVpns/{msgVpnName}/authenticationOauthProfiles/{oauthProfileName}/clientRequiredClaims/{clientRequiredClaimName} | Get a Required Claim object.
[**GetMsgVpnAuthenticationOauthProfileClientRequiredClaims**](MsgVpnApi.md#GetMsgVpnAuthenticationOauthProfileClientRequiredClaims) | **Get** /msgVpns/{msgVpnName}/authenticationOauthProfiles/{oauthProfileName}/clientRequiredClaims | Get a list of Required Claim objects.
[**GetMsgVpnAuthenticationOauthProfileResourceServerRequiredClaim**](MsgVpnApi.md#GetMsgVpnAuthenticationOauthProfileResourceServerRequiredClaim) | **Get** /msgVpns/{msgVpnName}/authenticationOauthProfiles/{oauthProfileName}/resourceServerRequiredClaims/{resourceServerRequiredClaimName} | Get a Required Claim object.
[**GetMsgVpnAuthenticationOauthProfileResourceServerRequiredClaims**](MsgVpnApi.md#GetMsgVpnAuthenticationOauthProfileResourceServerRequiredClaims) | **Get** /msgVpns/{msgVpnName}/authenticationOauthProfiles/{oauthProfileName}/resourceServerRequiredClaims | Get a list of Required Claim objects.
[**GetMsgVpnAuthenticationOauthProfiles**](MsgVpnApi.md#GetMsgVpnAuthenticationOauthProfiles) | **Get** /msgVpns/{msgVpnName}/authenticationOauthProfiles | Get a list of OAuth Profile objects.
[**GetMsgVpnAuthenticationOauthProvider**](MsgVpnApi.md#GetMsgVpnAuthenticationOauthProvider) | **Get** /msgVpns/{msgVpnName}/authenticationOauthProviders/{oauthProviderName} | Get an OAuth Provider object.
[**GetMsgVpnAuthenticationOauthProviders**](MsgVpnApi.md#GetMsgVpnAuthenticationOauthProviders) | **Get** /msgVpns/{msgVpnName}/authenticationOauthProviders | Get a list of OAuth Provider objects.
[**GetMsgVpnAuthorizationGroup**](MsgVpnApi.md#GetMsgVpnAuthorizationGroup) | **Get** /msgVpns/{msgVpnName}/authorizationGroups/{authorizationGroupName} | Get an Authorization Group object.
[**GetMsgVpnAuthorizationGroups**](MsgVpnApi.md#GetMsgVpnAuthorizationGroups) | **Get** /msgVpns/{msgVpnName}/authorizationGroups | Get a list of Authorization Group objects.
[**GetMsgVpnBridge**](MsgVpnApi.md#GetMsgVpnBridge) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter} | Get a Bridge object.
[**GetMsgVpnBridgeLocalSubscription**](MsgVpnApi.md#GetMsgVpnBridgeLocalSubscription) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/localSubscriptions/{localSubscriptionTopic} | Get a Bridge Local Subscriptions object.
[**GetMsgVpnBridgeLocalSubscriptions**](MsgVpnApi.md#GetMsgVpnBridgeLocalSubscriptions) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/localSubscriptions | Get a list of Bridge Local Subscriptions objects.
[**GetMsgVpnBridgeRemoteMsgVpn**](MsgVpnApi.md#GetMsgVpnBridgeRemoteMsgVpn) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns/{remoteMsgVpnName},{remoteMsgVpnLocation},{remoteMsgVpnInterface} | Get a Remote Message VPN object.
[**GetMsgVpnBridgeRemoteMsgVpns**](MsgVpnApi.md#GetMsgVpnBridgeRemoteMsgVpns) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns | Get a list of Remote Message VPN objects.
[**GetMsgVpnBridgeRemoteSubscription**](MsgVpnApi.md#GetMsgVpnBridgeRemoteSubscription) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteSubscriptions/{remoteSubscriptionTopic} | Get a Remote Subscription object.
[**GetMsgVpnBridgeRemoteSubscriptions**](MsgVpnApi.md#GetMsgVpnBridgeRemoteSubscriptions) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteSubscriptions | Get a list of Remote Subscription objects.
[**GetMsgVpnBridgeTlsTrustedCommonName**](MsgVpnApi.md#GetMsgVpnBridgeTlsTrustedCommonName) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/tlsTrustedCommonNames/{tlsTrustedCommonName} | Get a Trusted Common Name object.
[**GetMsgVpnBridgeTlsTrustedCommonNames**](MsgVpnApi.md#GetMsgVpnBridgeTlsTrustedCommonNames) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/tlsTrustedCommonNames | Get a list of Trusted Common Name objects.
[**GetMsgVpnBridges**](MsgVpnApi.md#GetMsgVpnBridges) | **Get** /msgVpns/{msgVpnName}/bridges | Get a list of Bridge objects.
[**GetMsgVpnCertMatchingRule**](MsgVpnApi.md#GetMsgVpnCertMatchingRule) | **Get** /msgVpns/{msgVpnName}/certMatchingRules/{ruleName} | Get a Certificate Matching Rule object.
[**GetMsgVpnCertMatchingRuleAttributeFilter**](MsgVpnApi.md#GetMsgVpnCertMatchingRuleAttributeFilter) | **Get** /msgVpns/{msgVpnName}/certMatchingRules/{ruleName}/attributeFilters/{filterName} | Get a Certificate Matching Rule Attribute Filter object.
[**GetMsgVpnCertMatchingRuleAttributeFilters**](MsgVpnApi.md#GetMsgVpnCertMatchingRuleAttributeFilters) | **Get** /msgVpns/{msgVpnName}/certMatchingRules/{ruleName}/attributeFilters | Get a list of Certificate Matching Rule Attribute Filter objects.
[**GetMsgVpnCertMatchingRuleCondition**](MsgVpnApi.md#GetMsgVpnCertMatchingRuleCondition) | **Get** /msgVpns/{msgVpnName}/certMatchingRules/{ruleName}/conditions/{source} | Get a Certificate Matching Rule Condition object.
[**GetMsgVpnCertMatchingRuleConditions**](MsgVpnApi.md#GetMsgVpnCertMatchingRuleConditions) | **Get** /msgVpns/{msgVpnName}/certMatchingRules/{ruleName}/conditions | Get a list of Certificate Matching Rule Condition objects.
[**GetMsgVpnCertMatchingRules**](MsgVpnApi.md#GetMsgVpnCertMatchingRules) | **Get** /msgVpns/{msgVpnName}/certMatchingRules | Get a list of Certificate Matching Rule objects.
[**GetMsgVpnClient**](MsgVpnApi.md#GetMsgVpnClient) | **Get** /msgVpns/{msgVpnName}/clients/{clientName} | Get a Client object.
[**GetMsgVpnClientConnection**](MsgVpnApi.md#GetMsgVpnClientConnection) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/connections/{clientAddress} | Get a Client Connection object.
[**GetMsgVpnClientConnections**](MsgVpnApi.md#GetMsgVpnClientConnections) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/connections | Get a list of Client Connection objects.
[**GetMsgVpnClientProfile**](MsgVpnApi.md#GetMsgVpnClientProfile) | **Get** /msgVpns/{msgVpnName}/clientProfiles/{clientProfileName} | Get a Client Profile object.
[**GetMsgVpnClientProfiles**](MsgVpnApi.md#GetMsgVpnClientProfiles) | **Get** /msgVpns/{msgVpnName}/clientProfiles | Get a list of Client Profile objects.
[**GetMsgVpnClientRxFlow**](MsgVpnApi.md#GetMsgVpnClientRxFlow) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/rxFlows/{flowId} | Get a Client Receive Flow object.
[**GetMsgVpnClientRxFlows**](MsgVpnApi.md#GetMsgVpnClientRxFlows) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/rxFlows | Get a list of Client Receive Flow objects.
[**GetMsgVpnClientSubscription**](MsgVpnApi.md#GetMsgVpnClientSubscription) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/subscriptions/{subscriptionTopic} | Get a Client Subscription object.
[**GetMsgVpnClientSubscriptions**](MsgVpnApi.md#GetMsgVpnClientSubscriptions) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/subscriptions | Get a list of Client Subscription objects.
[**GetMsgVpnClientTransactedSession**](MsgVpnApi.md#GetMsgVpnClientTransactedSession) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/transactedSessions/{sessionName} | Get a Client Transacted Session object.
[**GetMsgVpnClientTransactedSessions**](MsgVpnApi.md#GetMsgVpnClientTransactedSessions) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/transactedSessions | Get a list of Client Transacted Session objects.
[**GetMsgVpnClientTxFlow**](MsgVpnApi.md#GetMsgVpnClientTxFlow) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/txFlows/{flowId} | Get a Client Transmit Flow object.
[**GetMsgVpnClientTxFlows**](MsgVpnApi.md#GetMsgVpnClientTxFlows) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/txFlows | Get a list of Client Transmit Flow objects.
[**GetMsgVpnClientUsername**](MsgVpnApi.md#GetMsgVpnClientUsername) | **Get** /msgVpns/{msgVpnName}/clientUsernames/{clientUsername} | Get a Client Username object.
[**GetMsgVpnClientUsernameAttribute**](MsgVpnApi.md#GetMsgVpnClientUsernameAttribute) | **Get** /msgVpns/{msgVpnName}/clientUsernames/{clientUsername}/attributes/{attributeName},{attributeValue} | Get a Client Username Attribute object.
[**GetMsgVpnClientUsernameAttributes**](MsgVpnApi.md#GetMsgVpnClientUsernameAttributes) | **Get** /msgVpns/{msgVpnName}/clientUsernames/{clientUsername}/attributes | Get a list of Client Username Attribute objects.
[**GetMsgVpnClientUsernames**](MsgVpnApi.md#GetMsgVpnClientUsernames) | **Get** /msgVpns/{msgVpnName}/clientUsernames | Get a list of Client Username objects.
[**GetMsgVpnClients**](MsgVpnApi.md#GetMsgVpnClients) | **Get** /msgVpns/{msgVpnName}/clients | Get a list of Client objects.
[**GetMsgVpnConfigSyncRemoteNode**](MsgVpnApi.md#GetMsgVpnConfigSyncRemoteNode) | **Get** /msgVpns/{msgVpnName}/configSyncRemoteNodes/{remoteNodeName} | Get a Config Sync Remote Node object.
[**GetMsgVpnConfigSyncRemoteNodes**](MsgVpnApi.md#GetMsgVpnConfigSyncRemoteNodes) | **Get** /msgVpns/{msgVpnName}/configSyncRemoteNodes | Get a list of Config Sync Remote Node objects.
[**GetMsgVpnDistributedCache**](MsgVpnApi.md#GetMsgVpnDistributedCache) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName} | Get a Distributed Cache object.
[**GetMsgVpnDistributedCacheCluster**](MsgVpnApi.md#GetMsgVpnDistributedCacheCluster) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName} | Get a Cache Cluster object.
[**GetMsgVpnDistributedCacheClusterGlobalCachingHomeCluster**](MsgVpnApi.md#GetMsgVpnDistributedCacheClusterGlobalCachingHomeCluster) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName} | Get a Home Cache Cluster object.
[**GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix**](MsgVpnApi.md#GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName}/topicPrefixes/{topicPrefix} | Get a Topic Prefix object.
[**GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixes**](MsgVpnApi.md#GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixes) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName}/topicPrefixes | Get a list of Topic Prefix objects.
[**GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusters**](MsgVpnApi.md#GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusters) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters | Get a list of Home Cache Cluster objects.
[**GetMsgVpnDistributedCacheClusterInstance**](MsgVpnApi.md#GetMsgVpnDistributedCacheClusterInstance) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName} | Get a Cache Instance object.
[**GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster**](MsgVpnApi.md#GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/remoteGlobalCachingHomeClusters/{homeClusterName} | Get a Remote Home Cache Cluster object.
[**GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusters**](MsgVpnApi.md#GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusters) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/remoteGlobalCachingHomeClusters | Get a list of Remote Home Cache Cluster objects.
[**GetMsgVpnDistributedCacheClusterInstanceRemoteTopic**](MsgVpnApi.md#GetMsgVpnDistributedCacheClusterInstanceRemoteTopic) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/remoteTopics/{topic} | Get a Remote Topic object.
[**GetMsgVpnDistributedCacheClusterInstanceRemoteTopics**](MsgVpnApi.md#GetMsgVpnDistributedCacheClusterInstanceRemoteTopics) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/remoteTopics | Get a list of Remote Topic objects.
[**GetMsgVpnDistributedCacheClusterInstances**](MsgVpnApi.md#GetMsgVpnDistributedCacheClusterInstances) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances | Get a list of Cache Instance objects.
[**GetMsgVpnDistributedCacheClusterTopic**](MsgVpnApi.md#GetMsgVpnDistributedCacheClusterTopic) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/topics/{topic} | Get a Topic object.
[**GetMsgVpnDistributedCacheClusterTopics**](MsgVpnApi.md#GetMsgVpnDistributedCacheClusterTopics) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/topics | Get a list of Topic objects.
[**GetMsgVpnDistributedCacheClusters**](MsgVpnApi.md#GetMsgVpnDistributedCacheClusters) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters | Get a list of Cache Cluster objects.
[**GetMsgVpnDistributedCaches**](MsgVpnApi.md#GetMsgVpnDistributedCaches) | **Get** /msgVpns/{msgVpnName}/distributedCaches | Get a list of Distributed Cache objects.
[**GetMsgVpnDmrBridge**](MsgVpnApi.md#GetMsgVpnDmrBridge) | **Get** /msgVpns/{msgVpnName}/dmrBridges/{remoteNodeName} | Get a DMR Bridge object.
[**GetMsgVpnDmrBridges**](MsgVpnApi.md#GetMsgVpnDmrBridges) | **Get** /msgVpns/{msgVpnName}/dmrBridges | Get a list of DMR Bridge objects.
[**GetMsgVpnJndiConnectionFactories**](MsgVpnApi.md#GetMsgVpnJndiConnectionFactories) | **Get** /msgVpns/{msgVpnName}/jndiConnectionFactories | Get a list of JNDI Connection Factory objects.
[**GetMsgVpnJndiConnectionFactory**](MsgVpnApi.md#GetMsgVpnJndiConnectionFactory) | **Get** /msgVpns/{msgVpnName}/jndiConnectionFactories/{connectionFactoryName} | Get a JNDI Connection Factory object.
[**GetMsgVpnJndiQueue**](MsgVpnApi.md#GetMsgVpnJndiQueue) | **Get** /msgVpns/{msgVpnName}/jndiQueues/{queueName} | Get a JNDI Queue object.
[**GetMsgVpnJndiQueues**](MsgVpnApi.md#GetMsgVpnJndiQueues) | **Get** /msgVpns/{msgVpnName}/jndiQueues | Get a list of JNDI Queue objects.
[**GetMsgVpnJndiTopic**](MsgVpnApi.md#GetMsgVpnJndiTopic) | **Get** /msgVpns/{msgVpnName}/jndiTopics/{topicName} | Get a JNDI Topic object.
[**GetMsgVpnJndiTopics**](MsgVpnApi.md#GetMsgVpnJndiTopics) | **Get** /msgVpns/{msgVpnName}/jndiTopics | Get a list of JNDI Topic objects.
[**GetMsgVpnMqttRetainCache**](MsgVpnApi.md#GetMsgVpnMqttRetainCache) | **Get** /msgVpns/{msgVpnName}/mqttRetainCaches/{cacheName} | Get an MQTT Retain Cache object.
[**GetMsgVpnMqttRetainCaches**](MsgVpnApi.md#GetMsgVpnMqttRetainCaches) | **Get** /msgVpns/{msgVpnName}/mqttRetainCaches | Get a list of MQTT Retain Cache objects.
[**GetMsgVpnMqttSession**](MsgVpnApi.md#GetMsgVpnMqttSession) | **Get** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter} | Get an MQTT Session object.
[**GetMsgVpnMqttSessionSubscription**](MsgVpnApi.md#GetMsgVpnMqttSessionSubscription) | **Get** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions/{subscriptionTopic} | Get a Subscription object.
[**GetMsgVpnMqttSessionSubscriptions**](MsgVpnApi.md#GetMsgVpnMqttSessionSubscriptions) | **Get** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions | Get a list of Subscription objects.
[**GetMsgVpnMqttSessions**](MsgVpnApi.md#GetMsgVpnMqttSessions) | **Get** /msgVpns/{msgVpnName}/mqttSessions | Get a list of MQTT Session objects.
[**GetMsgVpnQueue**](MsgVpnApi.md#GetMsgVpnQueue) | **Get** /msgVpns/{msgVpnName}/queues/{queueName} | Get a Queue object.
[**GetMsgVpnQueueMsg**](MsgVpnApi.md#GetMsgVpnQueueMsg) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/msgs/{msgId} | Get a Queue Message object.
[**GetMsgVpnQueueMsgs**](MsgVpnApi.md#GetMsgVpnQueueMsgs) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/msgs | Get a list of Queue Message objects.
[**GetMsgVpnQueuePriorities**](MsgVpnApi.md#GetMsgVpnQueuePriorities) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/priorities | Get a list of Queue Priority objects.
[**GetMsgVpnQueuePriority**](MsgVpnApi.md#GetMsgVpnQueuePriority) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/priorities/{priority} | Get a Queue Priority object.
[**GetMsgVpnQueueSubscription**](MsgVpnApi.md#GetMsgVpnQueueSubscription) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/subscriptions/{subscriptionTopic} | Get a Queue Subscription object.
[**GetMsgVpnQueueSubscriptions**](MsgVpnApi.md#GetMsgVpnQueueSubscriptions) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/subscriptions | Get a list of Queue Subscription objects.
[**GetMsgVpnQueueTemplate**](MsgVpnApi.md#GetMsgVpnQueueTemplate) | **Get** /msgVpns/{msgVpnName}/queueTemplates/{queueTemplateName} | Get a Queue Template object.
[**GetMsgVpnQueueTemplates**](MsgVpnApi.md#GetMsgVpnQueueTemplates) | **Get** /msgVpns/{msgVpnName}/queueTemplates | Get a list of Queue Template objects.
[**GetMsgVpnQueueTxFlow**](MsgVpnApi.md#GetMsgVpnQueueTxFlow) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/txFlows/{flowId} | Get a Queue Transmit Flow object.
[**GetMsgVpnQueueTxFlows**](MsgVpnApi.md#GetMsgVpnQueueTxFlows) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/txFlows | Get a list of Queue Transmit Flow objects.
[**GetMsgVpnQueues**](MsgVpnApi.md#GetMsgVpnQueues) | **Get** /msgVpns/{msgVpnName}/queues | Get a list of Queue objects.
[**GetMsgVpnReplayLog**](MsgVpnApi.md#GetMsgVpnReplayLog) | **Get** /msgVpns/{msgVpnName}/replayLogs/{replayLogName} | Get a Replay Log object.
[**GetMsgVpnReplayLogMsg**](MsgVpnApi.md#GetMsgVpnReplayLogMsg) | **Get** /msgVpns/{msgVpnName}/replayLogs/{replayLogName}/msgs/{msgId} | Get a Message object.
[**GetMsgVpnReplayLogMsgs**](MsgVpnApi.md#GetMsgVpnReplayLogMsgs) | **Get** /msgVpns/{msgVpnName}/replayLogs/{replayLogName}/msgs | Get a list of Message objects.
[**GetMsgVpnReplayLogTopicFilterSubscription**](MsgVpnApi.md#GetMsgVpnReplayLogTopicFilterSubscription) | **Get** /msgVpns/{msgVpnName}/replayLogs/{replayLogName}/topicFilterSubscriptions/{topicFilterSubscription} | Get a Topic Filter Subscription object.
[**GetMsgVpnReplayLogTopicFilterSubscriptions**](MsgVpnApi.md#GetMsgVpnReplayLogTopicFilterSubscriptions) | **Get** /msgVpns/{msgVpnName}/replayLogs/{replayLogName}/topicFilterSubscriptions | Get a list of Topic Filter Subscription objects.
[**GetMsgVpnReplayLogs**](MsgVpnApi.md#GetMsgVpnReplayLogs) | **Get** /msgVpns/{msgVpnName}/replayLogs | Get a list of Replay Log objects.
[**GetMsgVpnReplicatedTopic**](MsgVpnApi.md#GetMsgVpnReplicatedTopic) | **Get** /msgVpns/{msgVpnName}/replicatedTopics/{replicatedTopic} | Get a Replicated Topic object.
[**GetMsgVpnReplicatedTopics**](MsgVpnApi.md#GetMsgVpnReplicatedTopics) | **Get** /msgVpns/{msgVpnName}/replicatedTopics | Get a list of Replicated Topic objects.
[**GetMsgVpnRestDeliveryPoint**](MsgVpnApi.md#GetMsgVpnRestDeliveryPoint) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName} | Get a REST Delivery Point object.
[**GetMsgVpnRestDeliveryPointQueueBinding**](MsgVpnApi.md#GetMsgVpnRestDeliveryPointQueueBinding) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName} | Get a Queue Binding object.
[**GetMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader**](MsgVpnApi.md#GetMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName}/protectedRequestHeaders/{headerName} | Get a Protected Request Header object.
[**GetMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeaders**](MsgVpnApi.md#GetMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeaders) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName}/protectedRequestHeaders | Get a list of Protected Request Header objects.
[**GetMsgVpnRestDeliveryPointQueueBindingRequestHeader**](MsgVpnApi.md#GetMsgVpnRestDeliveryPointQueueBindingRequestHeader) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName}/requestHeaders/{headerName} | Get a Request Header object.
[**GetMsgVpnRestDeliveryPointQueueBindingRequestHeaders**](MsgVpnApi.md#GetMsgVpnRestDeliveryPointQueueBindingRequestHeaders) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName}/requestHeaders | Get a list of Request Header objects.
[**GetMsgVpnRestDeliveryPointQueueBindings**](MsgVpnApi.md#GetMsgVpnRestDeliveryPointQueueBindings) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings | Get a list of Queue Binding objects.
[**GetMsgVpnRestDeliveryPointRestConsumer**](MsgVpnApi.md#GetMsgVpnRestDeliveryPointRestConsumer) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName} | Get a REST Consumer object.
[**GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim**](MsgVpnApi.md#GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/oauthJwtClaims/{oauthJwtClaimName} | Get a Claim object.
[**GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaims**](MsgVpnApi.md#GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaims) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/oauthJwtClaims | Get a list of Claim objects.
[**GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName**](MsgVpnApi.md#GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/tlsTrustedCommonNames/{tlsTrustedCommonName} | Get a Trusted Common Name object.
[**GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNames**](MsgVpnApi.md#GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNames) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/tlsTrustedCommonNames | Get a list of Trusted Common Name objects.
[**GetMsgVpnRestDeliveryPointRestConsumers**](MsgVpnApi.md#GetMsgVpnRestDeliveryPointRestConsumers) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers | Get a list of REST Consumer objects.
[**GetMsgVpnRestDeliveryPoints**](MsgVpnApi.md#GetMsgVpnRestDeliveryPoints) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints | Get a list of REST Delivery Point objects.
[**GetMsgVpnTelemetryProfile**](MsgVpnApi.md#GetMsgVpnTelemetryProfile) | **Get** /msgVpns/{msgVpnName}/telemetryProfiles/{telemetryProfileName} | Get a Telemetry Profile object.
[**GetMsgVpnTelemetryProfileReceiverAclConnectException**](MsgVpnApi.md#GetMsgVpnTelemetryProfileReceiverAclConnectException) | **Get** /msgVpns/{msgVpnName}/telemetryProfiles/{telemetryProfileName}/receiverAclConnectExceptions/{receiverAclConnectExceptionAddress} | Get a Receiver ACL Connect Exception object.
[**GetMsgVpnTelemetryProfileReceiverAclConnectExceptions**](MsgVpnApi.md#GetMsgVpnTelemetryProfileReceiverAclConnectExceptions) | **Get** /msgVpns/{msgVpnName}/telemetryProfiles/{telemetryProfileName}/receiverAclConnectExceptions | Get a list of Receiver ACL Connect Exception objects.
[**GetMsgVpnTelemetryProfileTraceFilter**](MsgVpnApi.md#GetMsgVpnTelemetryProfileTraceFilter) | **Get** /msgVpns/{msgVpnName}/telemetryProfiles/{telemetryProfileName}/traceFilters/{traceFilterName} | Get a Trace Filter object.
[**GetMsgVpnTelemetryProfileTraceFilterSubscription**](MsgVpnApi.md#GetMsgVpnTelemetryProfileTraceFilterSubscription) | **Get** /msgVpns/{msgVpnName}/telemetryProfiles/{telemetryProfileName}/traceFilters/{traceFilterName}/subscriptions/{subscription},{subscriptionSyntax} | Get a Telemetry Trace Filter Subscription object.
[**GetMsgVpnTelemetryProfileTraceFilterSubscriptions**](MsgVpnApi.md#GetMsgVpnTelemetryProfileTraceFilterSubscriptions) | **Get** /msgVpns/{msgVpnName}/telemetryProfiles/{telemetryProfileName}/traceFilters/{traceFilterName}/subscriptions | Get a list of Telemetry Trace Filter Subscription objects.
[**GetMsgVpnTelemetryProfileTraceFilters**](MsgVpnApi.md#GetMsgVpnTelemetryProfileTraceFilters) | **Get** /msgVpns/{msgVpnName}/telemetryProfiles/{telemetryProfileName}/traceFilters | Get a list of Trace Filter objects.
[**GetMsgVpnTelemetryProfiles**](MsgVpnApi.md#GetMsgVpnTelemetryProfiles) | **Get** /msgVpns/{msgVpnName}/telemetryProfiles | Get a list of Telemetry Profile objects.
[**GetMsgVpnTopicEndpoint**](MsgVpnApi.md#GetMsgVpnTopicEndpoint) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName} | Get a Topic Endpoint object.
[**GetMsgVpnTopicEndpointMsg**](MsgVpnApi.md#GetMsgVpnTopicEndpointMsg) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/msgs/{msgId} | Get a Topic Endpoint Message object.
[**GetMsgVpnTopicEndpointMsgs**](MsgVpnApi.md#GetMsgVpnTopicEndpointMsgs) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/msgs | Get a list of Topic Endpoint Message objects.
[**GetMsgVpnTopicEndpointPriorities**](MsgVpnApi.md#GetMsgVpnTopicEndpointPriorities) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/priorities | Get a list of Topic Endpoint Priority objects.
[**GetMsgVpnTopicEndpointPriority**](MsgVpnApi.md#GetMsgVpnTopicEndpointPriority) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/priorities/{priority} | Get a Topic Endpoint Priority object.
[**GetMsgVpnTopicEndpointTemplate**](MsgVpnApi.md#GetMsgVpnTopicEndpointTemplate) | **Get** /msgVpns/{msgVpnName}/topicEndpointTemplates/{topicEndpointTemplateName} | Get a Topic Endpoint Template object.
[**GetMsgVpnTopicEndpointTemplates**](MsgVpnApi.md#GetMsgVpnTopicEndpointTemplates) | **Get** /msgVpns/{msgVpnName}/topicEndpointTemplates | Get a list of Topic Endpoint Template objects.
[**GetMsgVpnTopicEndpointTxFlow**](MsgVpnApi.md#GetMsgVpnTopicEndpointTxFlow) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/txFlows/{flowId} | Get a Topic Endpoint Transmit Flow object.
[**GetMsgVpnTopicEndpointTxFlows**](MsgVpnApi.md#GetMsgVpnTopicEndpointTxFlows) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/txFlows | Get a list of Topic Endpoint Transmit Flow objects.
[**GetMsgVpnTopicEndpoints**](MsgVpnApi.md#GetMsgVpnTopicEndpoints) | **Get** /msgVpns/{msgVpnName}/topicEndpoints | Get a list of Topic Endpoint objects.
[**GetMsgVpnTransaction**](MsgVpnApi.md#GetMsgVpnTransaction) | **Get** /msgVpns/{msgVpnName}/transactions/{xid} | Get a Replicated Local Transaction or XA Transaction object.
[**GetMsgVpnTransactionConsumerMsg**](MsgVpnApi.md#GetMsgVpnTransactionConsumerMsg) | **Get** /msgVpns/{msgVpnName}/transactions/{xid}/consumerMsgs/{msgId} | Get a Transaction Consumer Message object.
[**GetMsgVpnTransactionConsumerMsgs**](MsgVpnApi.md#GetMsgVpnTransactionConsumerMsgs) | **Get** /msgVpns/{msgVpnName}/transactions/{xid}/consumerMsgs | Get a list of Transaction Consumer Message objects.
[**GetMsgVpnTransactionPublisherMsg**](MsgVpnApi.md#GetMsgVpnTransactionPublisherMsg) | **Get** /msgVpns/{msgVpnName}/transactions/{xid}/publisherMsgs/{msgId} | Get a Transaction Publisher Message object.
[**GetMsgVpnTransactionPublisherMsgs**](MsgVpnApi.md#GetMsgVpnTransactionPublisherMsgs) | **Get** /msgVpns/{msgVpnName}/transactions/{xid}/publisherMsgs | Get a list of Transaction Publisher Message objects.
[**GetMsgVpnTransactions**](MsgVpnApi.md#GetMsgVpnTransactions) | **Get** /msgVpns/{msgVpnName}/transactions | Get a list of Replicated Local Transaction or XA Transaction objects.
[**GetMsgVpns**](MsgVpnApi.md#GetMsgVpns) | **Get** /msgVpns | Get a list of Message VPN objects.

# **GetMsgVpn**
> MsgVpnResponse GetMsgVpn(ctx, msgVpnName, optional)
Get a Message VPN object.

Get a Message VPN object.  Message VPNs (Virtual Private Networks) allow for the segregation of topic space and clients. They also group clients connecting to a network of message brokers, such that messages published within a particular group are only visible to that group's clients.   Attribute|Identifying|Deprecated :---|:---:|:---: authenticationOauthDefaultProviderName||x bridgingTlsServerCertEnforceTrustedCommonNameEnabled||x configSyncLocalKey||x configSyncLocalLastResult||x configSyncLocalRole||x configSyncLocalState||x configSyncLocalTimeInState||x counter.controlRxByteCount||x counter.controlRxMsgCount||x counter.controlTxByteCount||x counter.controlTxMsgCount||x counter.dataRxByteCount||x counter.dataRxMsgCount||x counter.dataTxByteCount||x counter.dataTxMsgCount||x counter.discardedRxMsgCount||x counter.discardedTxMsgCount||x counter.loginRxMsgCount||x counter.loginTxMsgCount||x counter.msgSpoolRxMsgCount||x counter.msgSpoolTxMsgCount||x counter.tlsRxByteCount||x counter.tlsTxByteCount||x distributedCacheManagementEnabled||x msgVpnName|x| rate.averageRxByteRate||x rate.averageRxMsgRate||x rate.averageTxByteRate||x rate.averageTxMsgRate||x rate.rxByteRate||x rate.rxMsgRate||x rate.tlsAverageRxByteRate||x rate.tlsAverageTxByteRate||x rate.tlsRxByteRate||x rate.tlsTxByteRate||x rate.txByteRate||x rate.txMsgRate||x restTlsServerCertEnforceTrustedCommonNameEnabled||x    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MsgVpnApiGetMsgVpnOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnOpts struct
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

# **GetMsgVpnAclProfile**
> MsgVpnAclProfileResponse GetMsgVpnAclProfile(ctx, msgVpnName, aclProfileName, optional)
Get an ACL Profile object.

Get an ACL Profile object.  An ACL Profile controls whether an authenticated client is permitted to establish a connection with the message broker or permitted to publish and subscribe to specific topics.   Attribute|Identifying|Deprecated :---|:---:|:---: aclProfileName|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
 **optional** | ***MsgVpnApiGetMsgVpnAclProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnAclProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileResponse**](MsgVpnAclProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAclProfileClientConnectException**
> MsgVpnAclProfileClientConnectExceptionResponse GetMsgVpnAclProfileClientConnectException(ctx, msgVpnName, aclProfileName, clientConnectExceptionAddress, optional)
Get a Client Connect Exception object.

Get a Client Connect Exception object.  A Client Connect Exception is an exception to the default action to take when a client using the ACL Profile connects to the Message VPN. Exceptions must be expressed as an IP address/netmask in CIDR form.   Attribute|Identifying|Deprecated :---|:---:|:---: aclProfileName|x| clientConnectExceptionAddress|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
  **clientConnectExceptionAddress** | **string**| The IP address/netmask of the client connect exception in canonical CIDR form. | 
 **optional** | ***MsgVpnApiGetMsgVpnAclProfileClientConnectExceptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnAclProfileClientConnectExceptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileClientConnectExceptionResponse**](MsgVpnAclProfileClientConnectExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAclProfileClientConnectExceptions**
> MsgVpnAclProfileClientConnectExceptionsResponse GetMsgVpnAclProfileClientConnectExceptions(ctx, msgVpnName, aclProfileName, optional)
Get a list of Client Connect Exception objects.

Get a list of Client Connect Exception objects.  A Client Connect Exception is an exception to the default action to take when a client using the ACL Profile connects to the Message VPN. Exceptions must be expressed as an IP address/netmask in CIDR form.   Attribute|Identifying|Deprecated :---|:---:|:---: aclProfileName|x| clientConnectExceptionAddress|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
 **optional** | ***MsgVpnApiGetMsgVpnAclProfileClientConnectExceptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnAclProfileClientConnectExceptionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileClientConnectExceptionsResponse**](MsgVpnAclProfileClientConnectExceptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAclProfilePublishException**
> MsgVpnAclProfilePublishExceptionResponse GetMsgVpnAclProfilePublishException(ctx, msgVpnName, aclProfileName, topicSyntax, publishExceptionTopic, optional)
Get a Publish Topic Exception object.

Get a Publish Topic Exception object.  A Publish Topic Exception is an exception to the default action to take when a client using the ACL Profile publishes to a topic in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Deprecated :---|:---:|:---: aclProfileName|x|x msgVpnName|x|x publishExceptionTopic|x|x topicSyntax|x|x    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been deprecated since 2.14. Replaced by publishTopicExceptions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
  **topicSyntax** | **string**| The syntax of the topic for the exception to the default action taken. | 
  **publishExceptionTopic** | **string**| The topic for the exception to the default action taken. May include wildcard characters. | 
 **optional** | ***MsgVpnApiGetMsgVpnAclProfilePublishExceptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnAclProfilePublishExceptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfilePublishExceptionResponse**](MsgVpnAclProfilePublishExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAclProfilePublishExceptions**
> MsgVpnAclProfilePublishExceptionsResponse GetMsgVpnAclProfilePublishExceptions(ctx, msgVpnName, aclProfileName, optional)
Get a list of Publish Topic Exception objects.

Get a list of Publish Topic Exception objects.  A Publish Topic Exception is an exception to the default action to take when a client using the ACL Profile publishes to a topic in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Deprecated :---|:---:|:---: aclProfileName|x|x msgVpnName|x|x publishExceptionTopic|x|x topicSyntax|x|x    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been deprecated since 2.14. Replaced by publishTopicExceptions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
 **optional** | ***MsgVpnApiGetMsgVpnAclProfilePublishExceptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnAclProfilePublishExceptionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfilePublishExceptionsResponse**](MsgVpnAclProfilePublishExceptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAclProfilePublishTopicException**
> MsgVpnAclProfilePublishTopicExceptionResponse GetMsgVpnAclProfilePublishTopicException(ctx, msgVpnName, aclProfileName, publishTopicExceptionSyntax, publishTopicException, optional)
Get a Publish Topic Exception object.

Get a Publish Topic Exception object.  A Publish Topic Exception is an exception to the default action to take when a client using the ACL Profile publishes to a topic in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Deprecated :---|:---:|:---: aclProfileName|x| msgVpnName|x| publishTopicException|x| publishTopicExceptionSyntax|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
  **publishTopicExceptionSyntax** | **string**| The syntax of the topic for the exception to the default action taken. | 
  **publishTopicException** | **string**| The topic for the exception to the default action taken. May include wildcard characters. | 
 **optional** | ***MsgVpnApiGetMsgVpnAclProfilePublishTopicExceptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnAclProfilePublishTopicExceptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfilePublishTopicExceptionResponse**](MsgVpnAclProfilePublishTopicExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAclProfilePublishTopicExceptions**
> MsgVpnAclProfilePublishTopicExceptionsResponse GetMsgVpnAclProfilePublishTopicExceptions(ctx, msgVpnName, aclProfileName, optional)
Get a list of Publish Topic Exception objects.

Get a list of Publish Topic Exception objects.  A Publish Topic Exception is an exception to the default action to take when a client using the ACL Profile publishes to a topic in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Deprecated :---|:---:|:---: aclProfileName|x| msgVpnName|x| publishTopicException|x| publishTopicExceptionSyntax|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
 **optional** | ***MsgVpnApiGetMsgVpnAclProfilePublishTopicExceptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnAclProfilePublishTopicExceptionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfilePublishTopicExceptionsResponse**](MsgVpnAclProfilePublishTopicExceptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAclProfileSubscribeException**
> MsgVpnAclProfileSubscribeExceptionResponse GetMsgVpnAclProfileSubscribeException(ctx, msgVpnName, aclProfileName, topicSyntax, subscribeExceptionTopic, optional)
Get a Subscribe Topic Exception object.

Get a Subscribe Topic Exception object.  A Subscribe Topic Exception is an exception to the default action to take when a client using the ACL Profile subscribes to a topic in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Deprecated :---|:---:|:---: aclProfileName|x|x msgVpnName|x|x subscribeExceptionTopic|x|x topicSyntax|x|x    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been deprecated since 2.14. Replaced by subscribeTopicExceptions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
  **topicSyntax** | **string**| The syntax of the topic for the exception to the default action taken. | 
  **subscribeExceptionTopic** | **string**| The topic for the exception to the default action taken. May include wildcard characters. | 
 **optional** | ***MsgVpnApiGetMsgVpnAclProfileSubscribeExceptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnAclProfileSubscribeExceptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileSubscribeExceptionResponse**](MsgVpnAclProfileSubscribeExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAclProfileSubscribeExceptions**
> MsgVpnAclProfileSubscribeExceptionsResponse GetMsgVpnAclProfileSubscribeExceptions(ctx, msgVpnName, aclProfileName, optional)
Get a list of Subscribe Topic Exception objects.

Get a list of Subscribe Topic Exception objects.  A Subscribe Topic Exception is an exception to the default action to take when a client using the ACL Profile subscribes to a topic in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Deprecated :---|:---:|:---: aclProfileName|x|x msgVpnName|x|x subscribeExceptionTopic|x|x topicSyntax|x|x    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been deprecated since 2.14. Replaced by subscribeTopicExceptions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
 **optional** | ***MsgVpnApiGetMsgVpnAclProfileSubscribeExceptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnAclProfileSubscribeExceptionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileSubscribeExceptionsResponse**](MsgVpnAclProfileSubscribeExceptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAclProfileSubscribeShareNameException**
> MsgVpnAclProfileSubscribeShareNameExceptionResponse GetMsgVpnAclProfileSubscribeShareNameException(ctx, msgVpnName, aclProfileName, subscribeShareNameExceptionSyntax, subscribeShareNameException, optional)
Get a Subscribe Share Name Exception object.

Get a Subscribe Share Name Exception object.  A Subscribe Share Name Exception is an exception to the default action to take when a client using the ACL Profile subscribes to a share-name subscription in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Deprecated :---|:---:|:---: aclProfileName|x| msgVpnName|x| subscribeShareNameException|x| subscribeShareNameExceptionSyntax|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
  **subscribeShareNameExceptionSyntax** | **string**| The syntax of the subscribe share name for the exception to the default action taken. | 
  **subscribeShareNameException** | **string**| The subscribe share name exception to the default action taken. May include wildcard characters. | 
 **optional** | ***MsgVpnApiGetMsgVpnAclProfileSubscribeShareNameExceptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnAclProfileSubscribeShareNameExceptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileSubscribeShareNameExceptionResponse**](MsgVpnAclProfileSubscribeShareNameExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAclProfileSubscribeShareNameExceptions**
> MsgVpnAclProfileSubscribeShareNameExceptionsResponse GetMsgVpnAclProfileSubscribeShareNameExceptions(ctx, msgVpnName, aclProfileName, optional)
Get a list of Subscribe Share Name Exception objects.

Get a list of Subscribe Share Name Exception objects.  A Subscribe Share Name Exception is an exception to the default action to take when a client using the ACL Profile subscribes to a share-name subscription in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Deprecated :---|:---:|:---: aclProfileName|x| msgVpnName|x| subscribeShareNameException|x| subscribeShareNameExceptionSyntax|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
 **optional** | ***MsgVpnApiGetMsgVpnAclProfileSubscribeShareNameExceptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnAclProfileSubscribeShareNameExceptionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileSubscribeShareNameExceptionsResponse**](MsgVpnAclProfileSubscribeShareNameExceptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAclProfileSubscribeTopicException**
> MsgVpnAclProfileSubscribeTopicExceptionResponse GetMsgVpnAclProfileSubscribeTopicException(ctx, msgVpnName, aclProfileName, subscribeTopicExceptionSyntax, subscribeTopicException, optional)
Get a Subscribe Topic Exception object.

Get a Subscribe Topic Exception object.  A Subscribe Topic Exception is an exception to the default action to take when a client using the ACL Profile subscribes to a topic in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Deprecated :---|:---:|:---: aclProfileName|x| msgVpnName|x| subscribeTopicException|x| subscribeTopicExceptionSyntax|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
  **subscribeTopicExceptionSyntax** | **string**| The syntax of the topic for the exception to the default action taken. | 
  **subscribeTopicException** | **string**| The topic for the exception to the default action taken. May include wildcard characters. | 
 **optional** | ***MsgVpnApiGetMsgVpnAclProfileSubscribeTopicExceptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnAclProfileSubscribeTopicExceptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileSubscribeTopicExceptionResponse**](MsgVpnAclProfileSubscribeTopicExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAclProfileSubscribeTopicExceptions**
> MsgVpnAclProfileSubscribeTopicExceptionsResponse GetMsgVpnAclProfileSubscribeTopicExceptions(ctx, msgVpnName, aclProfileName, optional)
Get a list of Subscribe Topic Exception objects.

Get a list of Subscribe Topic Exception objects.  A Subscribe Topic Exception is an exception to the default action to take when a client using the ACL Profile subscribes to a topic in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Deprecated :---|:---:|:---: aclProfileName|x| msgVpnName|x| subscribeTopicException|x| subscribeTopicExceptionSyntax|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
 **optional** | ***MsgVpnApiGetMsgVpnAclProfileSubscribeTopicExceptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnAclProfileSubscribeTopicExceptionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileSubscribeTopicExceptionsResponse**](MsgVpnAclProfileSubscribeTopicExceptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAclProfiles**
> MsgVpnAclProfilesResponse GetMsgVpnAclProfiles(ctx, msgVpnName, optional)
Get a list of ACL Profile objects.

Get a list of ACL Profile objects.  An ACL Profile controls whether an authenticated client is permitted to establish a connection with the message broker or permitted to publish and subscribe to specific topics.   Attribute|Identifying|Deprecated :---|:---:|:---: aclProfileName|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MsgVpnApiGetMsgVpnAclProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnAclProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfilesResponse**](MsgVpnAclProfilesResponse.md)

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
 **optional** | ***MsgVpnApiGetMsgVpnAuthenticationOauthProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnAuthenticationOauthProfileOpts struct
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

# **GetMsgVpnAuthenticationOauthProfileClientRequiredClaim**
> MsgVpnAuthenticationOauthProfileClientRequiredClaimResponse GetMsgVpnAuthenticationOauthProfileClientRequiredClaim(ctx, msgVpnName, oauthProfileName, clientRequiredClaimName, optional)
Get a Required Claim object.

Get a Required Claim object.  Additional claims to be verified in the ID token.   Attribute|Identifying|Deprecated :---|:---:|:---: clientRequiredClaimName|x| msgVpnName|x| oauthProfileName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.25.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
  **clientRequiredClaimName** | **string**| The name of the ID token claim to verify. | 
 **optional** | ***MsgVpnApiGetMsgVpnAuthenticationOauthProfileClientRequiredClaimOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnAuthenticationOauthProfileClientRequiredClaimOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthenticationOauthProfileClientRequiredClaimResponse**](MsgVpnAuthenticationOauthProfileClientRequiredClaimResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAuthenticationOauthProfileClientRequiredClaims**
> MsgVpnAuthenticationOauthProfileClientRequiredClaimsResponse GetMsgVpnAuthenticationOauthProfileClientRequiredClaims(ctx, msgVpnName, oauthProfileName, optional)
Get a list of Required Claim objects.

Get a list of Required Claim objects.  Additional claims to be verified in the ID token.   Attribute|Identifying|Deprecated :---|:---:|:---: clientRequiredClaimName|x| msgVpnName|x| oauthProfileName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.25.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
 **optional** | ***MsgVpnApiGetMsgVpnAuthenticationOauthProfileClientRequiredClaimsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnAuthenticationOauthProfileClientRequiredClaimsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthenticationOauthProfileClientRequiredClaimsResponse**](MsgVpnAuthenticationOauthProfileClientRequiredClaimsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAuthenticationOauthProfileResourceServerRequiredClaim**
> MsgVpnAuthenticationOauthProfileResourceServerRequiredClaimResponse GetMsgVpnAuthenticationOauthProfileResourceServerRequiredClaim(ctx, msgVpnName, oauthProfileName, resourceServerRequiredClaimName, optional)
Get a Required Claim object.

Get a Required Claim object.  Additional claims to be verified in the access token.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| oauthProfileName|x| resourceServerRequiredClaimName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.25.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
  **resourceServerRequiredClaimName** | **string**| The name of the access token claim to verify. | 
 **optional** | ***MsgVpnApiGetMsgVpnAuthenticationOauthProfileResourceServerRequiredClaimOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnAuthenticationOauthProfileResourceServerRequiredClaimOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthenticationOauthProfileResourceServerRequiredClaimResponse**](MsgVpnAuthenticationOauthProfileResourceServerRequiredClaimResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAuthenticationOauthProfileResourceServerRequiredClaims**
> MsgVpnAuthenticationOauthProfileResourceServerRequiredClaimsResponse GetMsgVpnAuthenticationOauthProfileResourceServerRequiredClaims(ctx, msgVpnName, oauthProfileName, optional)
Get a list of Required Claim objects.

Get a list of Required Claim objects.  Additional claims to be verified in the access token.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| oauthProfileName|x| resourceServerRequiredClaimName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.25.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
 **optional** | ***MsgVpnApiGetMsgVpnAuthenticationOauthProfileResourceServerRequiredClaimsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnAuthenticationOauthProfileResourceServerRequiredClaimsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthenticationOauthProfileResourceServerRequiredClaimsResponse**](MsgVpnAuthenticationOauthProfileResourceServerRequiredClaimsResponse.md)

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
 **optional** | ***MsgVpnApiGetMsgVpnAuthenticationOauthProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnAuthenticationOauthProfilesOpts struct
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

Get an OAuth Provider object.  OAuth Providers contain information about the issuer of an OAuth token that is needed to validate the token and derive a client username from it.   Attribute|Identifying|Deprecated :---|:---:|:---: audienceClaimName||x audienceClaimSource||x audienceClaimValue||x audienceValidationEnabled||x authenticationSuccessCount||x authorizationGroupClaimName||x authorizationGroupClaimSource||x authorizationGroupEnabled||x disconnectOnTokenExpirationEnabled||x enabled||x jwksLastRefreshFailureReason||x jwksLastRefreshFailureTime||x jwksLastRefreshTime||x jwksNextScheduledRefreshTime||x jwksRefreshFailureCount||x jwksRefreshInterval||x jwksUri||x loginFailureIncorrectAudienceValueCount||x loginFailureInvalidAudienceValueCount||x loginFailureInvalidAuthorizationGroupValueCount||x loginFailureInvalidJwtSignatureCount||x loginFailureInvalidUsernameValueCount||x loginFailureMismatchedUsernameCount||x loginFailureMissingAudienceCount||x loginFailureMissingJwkCount||x loginFailureMissingOrInvalidTokenCount||x loginFailureMissingUsernameCount||x loginFailureTokenExpiredCount||x loginFailureTokenIntrospectionErroredCount||x loginFailureTokenIntrospectionFailureCount||x loginFailureTokenIntrospectionHttpsErrorCount||x loginFailureTokenIntrospectionInvalidCount||x loginFailureTokenIntrospectionTimeoutCount||x loginFailureTokenNotValidYetCount||x loginFailureUnsupportedAlgCount||x missingAuthorizationGroupCount||x msgVpnName|x|x oauthProviderName|x|x tokenIgnoreTimeLimitsEnabled||x tokenIntrospectionAverageTime||x tokenIntrospectionLastFailureReason||x tokenIntrospectionLastFailureTime||x tokenIntrospectionParameterName||x tokenIntrospectionSuccessCount||x tokenIntrospectionTimeout||x tokenIntrospectionUri||x tokenIntrospectionUsername||x usernameClaimName||x usernameClaimSource||x usernameValidateEnabled||x    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been deprecated since 2.25. Replaced by authenticationOauthProfiles.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **oauthProviderName** | **string**| The name of the OAuth Provider. | 
 **optional** | ***MsgVpnApiGetMsgVpnAuthenticationOauthProviderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnAuthenticationOauthProviderOpts struct
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

Get a list of OAuth Provider objects.  OAuth Providers contain information about the issuer of an OAuth token that is needed to validate the token and derive a client username from it.   Attribute|Identifying|Deprecated :---|:---:|:---: audienceClaimName||x audienceClaimSource||x audienceClaimValue||x audienceValidationEnabled||x authenticationSuccessCount||x authorizationGroupClaimName||x authorizationGroupClaimSource||x authorizationGroupEnabled||x disconnectOnTokenExpirationEnabled||x enabled||x jwksLastRefreshFailureReason||x jwksLastRefreshFailureTime||x jwksLastRefreshTime||x jwksNextScheduledRefreshTime||x jwksRefreshFailureCount||x jwksRefreshInterval||x jwksUri||x loginFailureIncorrectAudienceValueCount||x loginFailureInvalidAudienceValueCount||x loginFailureInvalidAuthorizationGroupValueCount||x loginFailureInvalidJwtSignatureCount||x loginFailureInvalidUsernameValueCount||x loginFailureMismatchedUsernameCount||x loginFailureMissingAudienceCount||x loginFailureMissingJwkCount||x loginFailureMissingOrInvalidTokenCount||x loginFailureMissingUsernameCount||x loginFailureTokenExpiredCount||x loginFailureTokenIntrospectionErroredCount||x loginFailureTokenIntrospectionFailureCount||x loginFailureTokenIntrospectionHttpsErrorCount||x loginFailureTokenIntrospectionInvalidCount||x loginFailureTokenIntrospectionTimeoutCount||x loginFailureTokenNotValidYetCount||x loginFailureUnsupportedAlgCount||x missingAuthorizationGroupCount||x msgVpnName|x|x oauthProviderName|x|x tokenIgnoreTimeLimitsEnabled||x tokenIntrospectionAverageTime||x tokenIntrospectionLastFailureReason||x tokenIntrospectionLastFailureTime||x tokenIntrospectionParameterName||x tokenIntrospectionSuccessCount||x tokenIntrospectionTimeout||x tokenIntrospectionUri||x tokenIntrospectionUsername||x usernameClaimName||x usernameClaimSource||x usernameValidateEnabled||x    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been deprecated since 2.25. Replaced by authenticationOauthProfiles.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MsgVpnApiGetMsgVpnAuthenticationOauthProvidersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnAuthenticationOauthProvidersOpts struct
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

# **GetMsgVpnAuthorizationGroup**
> MsgVpnAuthorizationGroupResponse GetMsgVpnAuthorizationGroup(ctx, msgVpnName, authorizationGroupName, optional)
Get an Authorization Group object.

Get an Authorization Group object.  To use client authorization groups configured on an external server to provide client authorizations, Authorization Group objects must be created on the Message VPN that match the authorization groups provisioned on the external server. These objects must be configured with the client profiles and ACL profiles that will be assigned to the clients that belong to those authorization groups. A newly created group is placed at the end of the group list which is the lowest priority.   Attribute|Identifying|Deprecated :---|:---:|:---: authorizationGroupName|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **authorizationGroupName** | **string**| The name of the Authorization Group. For LDAP groups, special care is needed if the group name contains special characters such as &#x27;#&#x27;, &#x27;+&#x27;, &#x27;;&#x27;, &#x27;&#x3D;&#x27; as the value of the group name returned from the LDAP server might prepend those characters with &#x27;\\&#x27;. For example a group name called &#x27;test#,lab,com&#x27; will be returned from the LDAP server as &#x27;test\\#,lab,com&#x27;. | 
 **optional** | ***MsgVpnApiGetMsgVpnAuthorizationGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnAuthorizationGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthorizationGroupResponse**](MsgVpnAuthorizationGroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnAuthorizationGroups**
> MsgVpnAuthorizationGroupsResponse GetMsgVpnAuthorizationGroups(ctx, msgVpnName, optional)
Get a list of Authorization Group objects.

Get a list of Authorization Group objects.  To use client authorization groups configured on an external server to provide client authorizations, Authorization Group objects must be created on the Message VPN that match the authorization groups provisioned on the external server. These objects must be configured with the client profiles and ACL profiles that will be assigned to the clients that belong to those authorization groups. A newly created group is placed at the end of the group list which is the lowest priority.   Attribute|Identifying|Deprecated :---|:---:|:---: authorizationGroupName|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MsgVpnApiGetMsgVpnAuthorizationGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnAuthorizationGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthorizationGroupsResponse**](MsgVpnAuthorizationGroupsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnBridge**
> MsgVpnBridgeResponse GetMsgVpnBridge(ctx, msgVpnName, bridgeName, bridgeVirtualRouter, optional)
Get a Bridge object.

Get a Bridge object.  Bridges can be used to link two Message VPNs so that messages published to one Message VPN that match the topic subscriptions set for the bridge are also delivered to the linked Message VPN.   Attribute|Identifying|Deprecated :---|:---:|:---: bridgeName|x| bridgeVirtualRouter|x| counter.controlRxByteCount||x counter.controlRxMsgCount||x counter.controlTxByteCount||x counter.controlTxMsgCount||x counter.dataRxByteCount||x counter.dataRxMsgCount||x counter.dataTxByteCount||x counter.dataTxMsgCount||x counter.discardedRxMsgCount||x counter.discardedTxMsgCount||x counter.loginRxMsgCount||x counter.loginTxMsgCount||x counter.msgSpoolRxMsgCount||x counter.rxByteCount||x counter.rxMsgCount||x counter.txByteCount||x counter.txMsgCount||x msgVpnName|x| rate.averageRxByteRate||x rate.averageRxMsgRate||x rate.averageTxByteRate||x rate.averageTxMsgRate||x rate.rxByteRate||x rate.rxMsgRate||x rate.txByteRate||x rate.txMsgRate||x    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **bridgeName** | **string**| The name of the Bridge. | 
  **bridgeVirtualRouter** | **string**| The virtual router of the Bridge. | 
 **optional** | ***MsgVpnApiGetMsgVpnBridgeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnBridgeOpts struct
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

# **GetMsgVpnBridgeLocalSubscription**
> MsgVpnBridgeLocalSubscriptionResponse GetMsgVpnBridgeLocalSubscription(ctx, msgVpnName, bridgeName, bridgeVirtualRouter, localSubscriptionTopic, optional)
Get a Bridge Local Subscriptions object.

Get a Bridge Local Subscriptions object.  A Local Subscription is a topic subscription used by a remote Message VPN Bridge to attract messages from this broker.   Attribute|Identifying|Deprecated :---|:---:|:---: bridgeName|x| bridgeVirtualRouter|x| localSubscriptionTopic|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **bridgeName** | **string**| The name of the Bridge. | 
  **bridgeVirtualRouter** | **string**| The virtual router of the Bridge. | 
  **localSubscriptionTopic** | **string**| The topic of the Bridge local subscription. | 
 **optional** | ***MsgVpnApiGetMsgVpnBridgeLocalSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnBridgeLocalSubscriptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeLocalSubscriptionResponse**](MsgVpnBridgeLocalSubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnBridgeLocalSubscriptions**
> MsgVpnBridgeLocalSubscriptionsResponse GetMsgVpnBridgeLocalSubscriptions(ctx, msgVpnName, bridgeName, bridgeVirtualRouter, optional)
Get a list of Bridge Local Subscriptions objects.

Get a list of Bridge Local Subscriptions objects.  A Local Subscription is a topic subscription used by a remote Message VPN Bridge to attract messages from this broker.   Attribute|Identifying|Deprecated :---|:---:|:---: bridgeName|x| bridgeVirtualRouter|x| localSubscriptionTopic|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **bridgeName** | **string**| The name of the Bridge. | 
  **bridgeVirtualRouter** | **string**| The virtual router of the Bridge. | 
 **optional** | ***MsgVpnApiGetMsgVpnBridgeLocalSubscriptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnBridgeLocalSubscriptionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeLocalSubscriptionsResponse**](MsgVpnBridgeLocalSubscriptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnBridgeRemoteMsgVpn**
> MsgVpnBridgeRemoteMsgVpnResponse GetMsgVpnBridgeRemoteMsgVpn(ctx, msgVpnName, bridgeName, bridgeVirtualRouter, remoteMsgVpnName, remoteMsgVpnLocation, remoteMsgVpnInterface, optional)
Get a Remote Message VPN object.

Get a Remote Message VPN object.  The Remote Message VPN is the Message VPN that the Bridge connects to.   Attribute|Identifying|Deprecated :---|:---:|:---: bridgeName|x| bridgeVirtualRouter|x| msgVpnName|x| remoteMsgVpnInterface|x| remoteMsgVpnLocation|x| remoteMsgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **bridgeName** | **string**| The name of the Bridge. | 
  **bridgeVirtualRouter** | **string**| The virtual router of the Bridge. | 
  **remoteMsgVpnName** | **string**| The name of the remote Message VPN. | 
  **remoteMsgVpnLocation** | **string**| The location of the remote Message VPN as either an FQDN with port, IP address with port, or virtual router name (starting with \&quot;v:\&quot;). | 
  **remoteMsgVpnInterface** | **string**| The physical interface on the local Message VPN host for connecting to the remote Message VPN. By default, an interface is chosen automatically (recommended), but if specified, &#x60;remoteMsgVpnLocation&#x60; must not be a virtual router name. | 
 **optional** | ***MsgVpnApiGetMsgVpnBridgeRemoteMsgVpnOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnBridgeRemoteMsgVpnOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeRemoteMsgVpnResponse**](MsgVpnBridgeRemoteMsgVpnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnBridgeRemoteMsgVpns**
> MsgVpnBridgeRemoteMsgVpnsResponse GetMsgVpnBridgeRemoteMsgVpns(ctx, msgVpnName, bridgeName, bridgeVirtualRouter, optional)
Get a list of Remote Message VPN objects.

Get a list of Remote Message VPN objects.  The Remote Message VPN is the Message VPN that the Bridge connects to.   Attribute|Identifying|Deprecated :---|:---:|:---: bridgeName|x| bridgeVirtualRouter|x| msgVpnName|x| remoteMsgVpnInterface|x| remoteMsgVpnLocation|x| remoteMsgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **bridgeName** | **string**| The name of the Bridge. | 
  **bridgeVirtualRouter** | **string**| The virtual router of the Bridge. | 
 **optional** | ***MsgVpnApiGetMsgVpnBridgeRemoteMsgVpnsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnBridgeRemoteMsgVpnsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeRemoteMsgVpnsResponse**](MsgVpnBridgeRemoteMsgVpnsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnBridgeRemoteSubscription**
> MsgVpnBridgeRemoteSubscriptionResponse GetMsgVpnBridgeRemoteSubscription(ctx, msgVpnName, bridgeName, bridgeVirtualRouter, remoteSubscriptionTopic, optional)
Get a Remote Subscription object.

Get a Remote Subscription object.  A Remote Subscription is a topic subscription used by the Message VPN Bridge to attract messages from the remote message broker.   Attribute|Identifying|Deprecated :---|:---:|:---: bridgeName|x| bridgeVirtualRouter|x| msgVpnName|x| remoteSubscriptionTopic|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **bridgeName** | **string**| The name of the Bridge. | 
  **bridgeVirtualRouter** | **string**| The virtual router of the Bridge. | 
  **remoteSubscriptionTopic** | **string**| The topic of the Bridge remote subscription. | 
 **optional** | ***MsgVpnApiGetMsgVpnBridgeRemoteSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnBridgeRemoteSubscriptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeRemoteSubscriptionResponse**](MsgVpnBridgeRemoteSubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnBridgeRemoteSubscriptions**
> MsgVpnBridgeRemoteSubscriptionsResponse GetMsgVpnBridgeRemoteSubscriptions(ctx, msgVpnName, bridgeName, bridgeVirtualRouter, optional)
Get a list of Remote Subscription objects.

Get a list of Remote Subscription objects.  A Remote Subscription is a topic subscription used by the Message VPN Bridge to attract messages from the remote message broker.   Attribute|Identifying|Deprecated :---|:---:|:---: bridgeName|x| bridgeVirtualRouter|x| msgVpnName|x| remoteSubscriptionTopic|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **bridgeName** | **string**| The name of the Bridge. | 
  **bridgeVirtualRouter** | **string**| The virtual router of the Bridge. | 
 **optional** | ***MsgVpnApiGetMsgVpnBridgeRemoteSubscriptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnBridgeRemoteSubscriptionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeRemoteSubscriptionsResponse**](MsgVpnBridgeRemoteSubscriptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnBridgeTlsTrustedCommonName**
> MsgVpnBridgeTlsTrustedCommonNameResponse GetMsgVpnBridgeTlsTrustedCommonName(ctx, msgVpnName, bridgeName, bridgeVirtualRouter, tlsTrustedCommonName, optional)
Get a Trusted Common Name object.

Get a Trusted Common Name object.  The Trusted Common Names for the Bridge are used by encrypted transports to verify the name in the certificate presented by the remote node. They must include the common name of the remote node's server certificate or client certificate, depending upon the initiator of the connection.   Attribute|Identifying|Deprecated :---|:---:|:---: bridgeName|x|x bridgeVirtualRouter|x|x msgVpnName|x|x tlsTrustedCommonName|x|x    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been deprecated since 2.18. Common Name validation has been replaced by Server Certificate Name validation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **bridgeName** | **string**| The name of the Bridge. | 
  **bridgeVirtualRouter** | **string**| The virtual router of the Bridge. | 
  **tlsTrustedCommonName** | **string**| The expected trusted common name of the remote certificate. | 
 **optional** | ***MsgVpnApiGetMsgVpnBridgeTlsTrustedCommonNameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnBridgeTlsTrustedCommonNameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeTlsTrustedCommonNameResponse**](MsgVpnBridgeTlsTrustedCommonNameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnBridgeTlsTrustedCommonNames**
> MsgVpnBridgeTlsTrustedCommonNamesResponse GetMsgVpnBridgeTlsTrustedCommonNames(ctx, msgVpnName, bridgeName, bridgeVirtualRouter, optional)
Get a list of Trusted Common Name objects.

Get a list of Trusted Common Name objects.  The Trusted Common Names for the Bridge are used by encrypted transports to verify the name in the certificate presented by the remote node. They must include the common name of the remote node's server certificate or client certificate, depending upon the initiator of the connection.   Attribute|Identifying|Deprecated :---|:---:|:---: bridgeName|x|x bridgeVirtualRouter|x|x msgVpnName|x|x tlsTrustedCommonName|x|x    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been deprecated since 2.18. Common Name validation has been replaced by Server Certificate Name validation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **bridgeName** | **string**| The name of the Bridge. | 
  **bridgeVirtualRouter** | **string**| The virtual router of the Bridge. | 
 **optional** | ***MsgVpnApiGetMsgVpnBridgeTlsTrustedCommonNamesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnBridgeTlsTrustedCommonNamesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeTlsTrustedCommonNamesResponse**](MsgVpnBridgeTlsTrustedCommonNamesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnBridges**
> MsgVpnBridgesResponse GetMsgVpnBridges(ctx, msgVpnName, optional)
Get a list of Bridge objects.

Get a list of Bridge objects.  Bridges can be used to link two Message VPNs so that messages published to one Message VPN that match the topic subscriptions set for the bridge are also delivered to the linked Message VPN.   Attribute|Identifying|Deprecated :---|:---:|:---: bridgeName|x| bridgeVirtualRouter|x| counter.controlRxByteCount||x counter.controlRxMsgCount||x counter.controlTxByteCount||x counter.controlTxMsgCount||x counter.dataRxByteCount||x counter.dataRxMsgCount||x counter.dataTxByteCount||x counter.dataTxMsgCount||x counter.discardedRxMsgCount||x counter.discardedTxMsgCount||x counter.loginRxMsgCount||x counter.loginTxMsgCount||x counter.msgSpoolRxMsgCount||x counter.rxByteCount||x counter.rxMsgCount||x counter.txByteCount||x counter.txMsgCount||x msgVpnName|x| rate.averageRxByteRate||x rate.averageRxMsgRate||x rate.averageTxByteRate||x rate.averageTxMsgRate||x rate.rxByteRate||x rate.rxMsgRate||x rate.txByteRate||x rate.txMsgRate||x    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MsgVpnApiGetMsgVpnBridgesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnBridgesOpts struct
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

# **GetMsgVpnCertMatchingRule**
> MsgVpnCertMatchingRuleResponse GetMsgVpnCertMatchingRule(ctx, msgVpnName, ruleName, optional)
Get a Certificate Matching Rule object.

Get a Certificate Matching Rule object.  A Cert Matching Rule is a collection of conditions and attribute filters that all have to be satisfied for certificate to be acceptable as authentication for a given username.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| ruleName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.27.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **ruleName** | **string**| The name of the rule. | 
 **optional** | ***MsgVpnApiGetMsgVpnCertMatchingRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnCertMatchingRuleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnCertMatchingRuleResponse**](MsgVpnCertMatchingRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnCertMatchingRuleAttributeFilter**
> MsgVpnCertMatchingRuleAttributeFilterResponse GetMsgVpnCertMatchingRuleAttributeFilter(ctx, msgVpnName, ruleName, filterName, optional)
Get a Certificate Matching Rule Attribute Filter object.

Get a Certificate Matching Rule Attribute Filter object.  A Cert Matching Rule Attribute Filter compares a username attribute to a string.   Attribute|Identifying|Deprecated :---|:---:|:---: filterName|x| msgVpnName|x| ruleName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.28.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **ruleName** | **string**| The name of the rule. | 
  **filterName** | **string**| The name of the filter. | 
 **optional** | ***MsgVpnApiGetMsgVpnCertMatchingRuleAttributeFilterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnCertMatchingRuleAttributeFilterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnCertMatchingRuleAttributeFilterResponse**](MsgVpnCertMatchingRuleAttributeFilterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnCertMatchingRuleAttributeFilters**
> MsgVpnCertMatchingRuleAttributeFiltersResponse GetMsgVpnCertMatchingRuleAttributeFilters(ctx, msgVpnName, ruleName, optional)
Get a list of Certificate Matching Rule Attribute Filter objects.

Get a list of Certificate Matching Rule Attribute Filter objects.  A Cert Matching Rule Attribute Filter compares a username attribute to a string.   Attribute|Identifying|Deprecated :---|:---:|:---: filterName|x| msgVpnName|x| ruleName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.28.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **ruleName** | **string**| The name of the rule. | 
 **optional** | ***MsgVpnApiGetMsgVpnCertMatchingRuleAttributeFiltersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnCertMatchingRuleAttributeFiltersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnCertMatchingRuleAttributeFiltersResponse**](MsgVpnCertMatchingRuleAttributeFiltersResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnCertMatchingRuleCondition**
> MsgVpnCertMatchingRuleConditionResponse GetMsgVpnCertMatchingRuleCondition(ctx, msgVpnName, ruleName, source, optional)
Get a Certificate Matching Rule Condition object.

Get a Certificate Matching Rule Condition object.  A Cert Matching Rule Condition compares data extracted from a certificate to a username attribute or an expression.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| ruleName|x| source|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.27.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **ruleName** | **string**| The name of the rule. | 
  **source** | **string**| Certificate field to be compared with the Attribute. | 
 **optional** | ***MsgVpnApiGetMsgVpnCertMatchingRuleConditionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnCertMatchingRuleConditionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnCertMatchingRuleConditionResponse**](MsgVpnCertMatchingRuleConditionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnCertMatchingRuleConditions**
> MsgVpnCertMatchingRuleConditionsResponse GetMsgVpnCertMatchingRuleConditions(ctx, msgVpnName, ruleName, optional)
Get a list of Certificate Matching Rule Condition objects.

Get a list of Certificate Matching Rule Condition objects.  A Cert Matching Rule Condition compares data extracted from a certificate to a username attribute or an expression.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| ruleName|x| source|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.27.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **ruleName** | **string**| The name of the rule. | 
 **optional** | ***MsgVpnApiGetMsgVpnCertMatchingRuleConditionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnCertMatchingRuleConditionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnCertMatchingRuleConditionsResponse**](MsgVpnCertMatchingRuleConditionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnCertMatchingRules**
> MsgVpnCertMatchingRulesResponse GetMsgVpnCertMatchingRules(ctx, msgVpnName, optional)
Get a list of Certificate Matching Rule objects.

Get a list of Certificate Matching Rule objects.  A Cert Matching Rule is a collection of conditions and attribute filters that all have to be satisfied for certificate to be acceptable as authentication for a given username.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| ruleName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.27.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MsgVpnApiGetMsgVpnCertMatchingRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnCertMatchingRulesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnCertMatchingRulesResponse**](MsgVpnCertMatchingRulesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnClient**
> MsgVpnClientResponse GetMsgVpnClient(ctx, msgVpnName, clientName, optional)
Get a Client object.

Get a Client object.  Applications or devices that connect to message brokers to send and/or receive messages are represented as Clients.   Attribute|Identifying|Deprecated :---|:---:|:---: clientName|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **clientName** | **string**| The name of the Client. | 
 **optional** | ***MsgVpnApiGetMsgVpnClientOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnClientOpts struct
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

# **GetMsgVpnClientConnection**
> MsgVpnClientConnectionResponse GetMsgVpnClientConnection(ctx, msgVpnName, clientName, clientAddress, optional)
Get a Client Connection object.

Get a Client Connection object.  A Client Connection represents the Transmission Control Protocol (TCP) connection the Client uses to communicate with the message broker.   Attribute|Identifying|Deprecated :---|:---:|:---: clientAddress|x| clientName|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **clientName** | **string**| The name of the Client. | 
  **clientAddress** | **string**| The IP address and TCP port on the Client side of the Client Connection. | 
 **optional** | ***MsgVpnApiGetMsgVpnClientConnectionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnClientConnectionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientConnectionResponse**](MsgVpnClientConnectionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnClientConnections**
> MsgVpnClientConnectionsResponse GetMsgVpnClientConnections(ctx, msgVpnName, clientName, optional)
Get a list of Client Connection objects.

Get a list of Client Connection objects.  A Client Connection represents the Transmission Control Protocol (TCP) connection the Client uses to communicate with the message broker.   Attribute|Identifying|Deprecated :---|:---:|:---: clientAddress|x| clientName|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **clientName** | **string**| The name of the Client. | 
 **optional** | ***MsgVpnApiGetMsgVpnClientConnectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnClientConnectionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientConnectionsResponse**](MsgVpnClientConnectionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnClientProfile**
> MsgVpnClientProfileResponse GetMsgVpnClientProfile(ctx, msgVpnName, clientProfileName, optional)
Get a Client Profile object.

Get a Client Profile object.  Client Profiles are used to assign common configuration properties to clients that have been successfully authorized.   Attribute|Identifying|Deprecated :---|:---:|:---: allowCutThroughForwardingEnabled||x apiQueueManagementCopyFromOnCreateName||x apiTopicEndpointManagementCopyFromOnCreateName||x clientProfileName|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **clientProfileName** | **string**| The name of the Client Profile. | 
 **optional** | ***MsgVpnApiGetMsgVpnClientProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnClientProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientProfileResponse**](MsgVpnClientProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnClientProfiles**
> MsgVpnClientProfilesResponse GetMsgVpnClientProfiles(ctx, msgVpnName, optional)
Get a list of Client Profile objects.

Get a list of Client Profile objects.  Client Profiles are used to assign common configuration properties to clients that have been successfully authorized.   Attribute|Identifying|Deprecated :---|:---:|:---: allowCutThroughForwardingEnabled||x apiQueueManagementCopyFromOnCreateName||x apiTopicEndpointManagementCopyFromOnCreateName||x clientProfileName|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MsgVpnApiGetMsgVpnClientProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnClientProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientProfilesResponse**](MsgVpnClientProfilesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnClientRxFlow**
> MsgVpnClientRxFlowResponse GetMsgVpnClientRxFlow(ctx, msgVpnName, clientName, flowId, optional)
Get a Client Receive Flow object.

Get a Client Receive Flow object.  Client Receive Flows are used by clients to publish Guaranteed messages to a message broker.   Attribute|Identifying|Deprecated :---|:---:|:---: clientName|x| flowId|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **clientName** | **string**| The name of the Client. | 
  **flowId** | **string**| The identifier (ID) of the flow. | 
 **optional** | ***MsgVpnApiGetMsgVpnClientRxFlowOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnClientRxFlowOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientRxFlowResponse**](MsgVpnClientRxFlowResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnClientRxFlows**
> MsgVpnClientRxFlowsResponse GetMsgVpnClientRxFlows(ctx, msgVpnName, clientName, optional)
Get a list of Client Receive Flow objects.

Get a list of Client Receive Flow objects.  Client Receive Flows are used by clients to publish Guaranteed messages to a message broker.   Attribute|Identifying|Deprecated :---|:---:|:---: clientName|x| flowId|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **clientName** | **string**| The name of the Client. | 
 **optional** | ***MsgVpnApiGetMsgVpnClientRxFlowsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnClientRxFlowsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientRxFlowsResponse**](MsgVpnClientRxFlowsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnClientSubscription**
> MsgVpnClientSubscriptionResponse GetMsgVpnClientSubscription(ctx, msgVpnName, clientName, subscriptionTopic, optional)
Get a Client Subscription object.

Get a Client Subscription object.  Once clients are authenticated on the message broker they can add and remove Client Subscriptions for Direct messages published to the Message VPN to which they have connected.   Attribute|Identifying|Deprecated :---|:---:|:---: clientName|x| msgVpnName|x| subscriptionTopic|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **clientName** | **string**| The name of the Client. | 
  **subscriptionTopic** | **string**| The topic of the Subscription. | 
 **optional** | ***MsgVpnApiGetMsgVpnClientSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnClientSubscriptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientSubscriptionResponse**](MsgVpnClientSubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnClientSubscriptions**
> MsgVpnClientSubscriptionsResponse GetMsgVpnClientSubscriptions(ctx, msgVpnName, clientName, optional)
Get a list of Client Subscription objects.

Get a list of Client Subscription objects.  Once clients are authenticated on the message broker they can add and remove Client Subscriptions for Direct messages published to the Message VPN to which they have connected.   Attribute|Identifying|Deprecated :---|:---:|:---: clientName|x| msgVpnName|x| subscriptionTopic|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **clientName** | **string**| The name of the Client. | 
 **optional** | ***MsgVpnApiGetMsgVpnClientSubscriptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnClientSubscriptionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientSubscriptionsResponse**](MsgVpnClientSubscriptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnClientTransactedSession**
> MsgVpnClientTransactedSessionResponse GetMsgVpnClientTransactedSession(ctx, msgVpnName, clientName, sessionName, optional)
Get a Client Transacted Session object.

Get a Client Transacted Session object.  Transacted Sessions enable clients to group multiple message send and/or receive operations together in single, atomic units known as local transactions.   Attribute|Identifying|Deprecated :---|:---:|:---: clientName|x| msgVpnName|x| sessionName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **clientName** | **string**| The name of the Client. | 
  **sessionName** | **string**| The name of the Transacted Session. | 
 **optional** | ***MsgVpnApiGetMsgVpnClientTransactedSessionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnClientTransactedSessionOpts struct
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

Get a list of Client Transacted Session objects.  Transacted Sessions enable clients to group multiple message send and/or receive operations together in single, atomic units known as local transactions.   Attribute|Identifying|Deprecated :---|:---:|:---: clientName|x| msgVpnName|x| sessionName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **clientName** | **string**| The name of the Client. | 
 **optional** | ***MsgVpnApiGetMsgVpnClientTransactedSessionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnClientTransactedSessionsOpts struct
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

# **GetMsgVpnClientTxFlow**
> MsgVpnClientTxFlowResponse GetMsgVpnClientTxFlow(ctx, msgVpnName, clientName, flowId, optional)
Get a Client Transmit Flow object.

Get a Client Transmit Flow object.  Client Transmit Flows are used by clients to consume Guaranteed messages from a message broker.   Attribute|Identifying|Deprecated :---|:---:|:---: clientName|x| flowId|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **clientName** | **string**| The name of the Client. | 
  **flowId** | **string**| The identifier (ID) of the flow. | 
 **optional** | ***MsgVpnApiGetMsgVpnClientTxFlowOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnClientTxFlowOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientTxFlowResponse**](MsgVpnClientTxFlowResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnClientTxFlows**
> MsgVpnClientTxFlowsResponse GetMsgVpnClientTxFlows(ctx, msgVpnName, clientName, optional)
Get a list of Client Transmit Flow objects.

Get a list of Client Transmit Flow objects.  Client Transmit Flows are used by clients to consume Guaranteed messages from a message broker.   Attribute|Identifying|Deprecated :---|:---:|:---: clientName|x| flowId|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **clientName** | **string**| The name of the Client. | 
 **optional** | ***MsgVpnApiGetMsgVpnClientTxFlowsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnClientTxFlowsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientTxFlowsResponse**](MsgVpnClientTxFlowsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnClientUsername**
> MsgVpnClientUsernameResponse GetMsgVpnClientUsername(ctx, msgVpnName, clientUsername, optional)
Get a Client Username object.

Get a Client Username object.  A client is only authorized to connect to a Message VPN that is associated with a Client Username that the client has been assigned.   Attribute|Identifying|Deprecated :---|:---:|:---: clientUsername|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **clientUsername** | **string**| The name of the Client Username. | 
 **optional** | ***MsgVpnApiGetMsgVpnClientUsernameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnClientUsernameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientUsernameResponse**](MsgVpnClientUsernameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnClientUsernameAttribute**
> MsgVpnClientUsernameAttributeResponse GetMsgVpnClientUsernameAttribute(ctx, msgVpnName, clientUsername, attributeName, attributeValue, optional)
Get a Client Username Attribute object.

Get a Client Username Attribute object.  A ClientUsername Attribute is a key+value pair that can be used to locate a client username, for example when using client certificate mapping.   Attribute|Identifying|Deprecated :---|:---:|:---: attributeName|x| attributeValue|x| clientUsername|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.27.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **clientUsername** | **string**| The name of the Client Username. | 
  **attributeName** | **string**| The name of the Attribute. | 
  **attributeValue** | **string**| The value of the Attribute. | 
 **optional** | ***MsgVpnApiGetMsgVpnClientUsernameAttributeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnClientUsernameAttributeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientUsernameAttributeResponse**](MsgVpnClientUsernameAttributeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnClientUsernameAttributes**
> MsgVpnClientUsernameAttributesResponse GetMsgVpnClientUsernameAttributes(ctx, msgVpnName, clientUsername, optional)
Get a list of Client Username Attribute objects.

Get a list of Client Username Attribute objects.  A ClientUsername Attribute is a key+value pair that can be used to locate a client username, for example when using client certificate mapping.   Attribute|Identifying|Deprecated :---|:---:|:---: attributeName|x| attributeValue|x| clientUsername|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.27.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **clientUsername** | **string**| The name of the Client Username. | 
 **optional** | ***MsgVpnApiGetMsgVpnClientUsernameAttributesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnClientUsernameAttributesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientUsernameAttributesResponse**](MsgVpnClientUsernameAttributesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnClientUsernames**
> MsgVpnClientUsernamesResponse GetMsgVpnClientUsernames(ctx, msgVpnName, optional)
Get a list of Client Username objects.

Get a list of Client Username objects.  A client is only authorized to connect to a Message VPN that is associated with a Client Username that the client has been assigned.   Attribute|Identifying|Deprecated :---|:---:|:---: clientUsername|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MsgVpnApiGetMsgVpnClientUsernamesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnClientUsernamesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientUsernamesResponse**](MsgVpnClientUsernamesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnClients**
> MsgVpnClientsResponse GetMsgVpnClients(ctx, msgVpnName, optional)
Get a list of Client objects.

Get a list of Client objects.  Applications or devices that connect to message brokers to send and/or receive messages are represented as Clients.   Attribute|Identifying|Deprecated :---|:---:|:---: clientName|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MsgVpnApiGetMsgVpnClientsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnClientsOpts struct
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

# **GetMsgVpnConfigSyncRemoteNode**
> MsgVpnConfigSyncRemoteNodeResponse GetMsgVpnConfigSyncRemoteNode(ctx, msgVpnName, remoteNodeName, optional)
Get a Config Sync Remote Node object.

Get a Config Sync Remote Node object.  A Config Sync Remote Node object contains information about the status of the table for this Message VPN with respect to a remote node.   Attribute|Identifying|Deprecated :---|:---:|:---: lastMsgRxTime||x msgVpnName|x|x remoteNodeName|x|x role||x stale||x state||x timeInState||x    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been deprecated since 2.22. This attribute has been deprecated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **remoteNodeName** | **string**| The name of the Config Sync Remote Node. | 
 **optional** | ***MsgVpnApiGetMsgVpnConfigSyncRemoteNodeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnConfigSyncRemoteNodeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnConfigSyncRemoteNodeResponse**](MsgVpnConfigSyncRemoteNodeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnConfigSyncRemoteNodes**
> MsgVpnConfigSyncRemoteNodesResponse GetMsgVpnConfigSyncRemoteNodes(ctx, msgVpnName, optional)
Get a list of Config Sync Remote Node objects.

Get a list of Config Sync Remote Node objects.  A Config Sync Remote Node object contains information about the status of the table for this Message VPN with respect to a remote node.   Attribute|Identifying|Deprecated :---|:---:|:---: lastMsgRxTime||x msgVpnName|x|x remoteNodeName|x|x role||x stale||x state||x timeInState||x    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been deprecated since 2.22. This attribute has been deprecated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MsgVpnApiGetMsgVpnConfigSyncRemoteNodesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnConfigSyncRemoteNodesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnConfigSyncRemoteNodesResponse**](MsgVpnConfigSyncRemoteNodesResponse.md)

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
 **optional** | ***MsgVpnApiGetMsgVpnDistributedCacheOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnDistributedCacheOpts struct
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
 **optional** | ***MsgVpnApiGetMsgVpnDistributedCacheClusterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnDistributedCacheClusterOpts struct
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

# **GetMsgVpnDistributedCacheClusterGlobalCachingHomeCluster**
> MsgVpnDistributedCacheClusterGlobalCachingHomeClusterResponse GetMsgVpnDistributedCacheClusterGlobalCachingHomeCluster(ctx, msgVpnName, cacheName, clusterName, homeClusterName, optional)
Get a Home Cache Cluster object.

Get a Home Cache Cluster object.  A Home Cache Cluster is a Cache Cluster that is the \"definitive\" Cache Cluster for a given topic in the context of the Global Caching feature.   Attribute|Identifying|Deprecated :---|:---:|:---: cacheName|x| clusterName|x| homeClusterName|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **cacheName** | **string**| The name of the Distributed Cache. | 
  **clusterName** | **string**| The name of the Cache Cluster. | 
  **homeClusterName** | **string**| The name of the remote Home Cache Cluster. | 
 **optional** | ***MsgVpnApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterGlobalCachingHomeClusterResponse**](MsgVpnDistributedCacheClusterGlobalCachingHomeClusterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix**
> MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixResponse GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix(ctx, msgVpnName, cacheName, clusterName, homeClusterName, topicPrefix, optional)
Get a Topic Prefix object.

Get a Topic Prefix object.  A Topic Prefix is a prefix for a global topic that is available from the containing Home Cache Cluster.   Attribute|Identifying|Deprecated :---|:---:|:---: cacheName|x| clusterName|x| homeClusterName|x| msgVpnName|x| topicPrefix|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **cacheName** | **string**| The name of the Distributed Cache. | 
  **clusterName** | **string**| The name of the Cache Cluster. | 
  **homeClusterName** | **string**| The name of the remote Home Cache Cluster. | 
  **topicPrefix** | **string**| A topic prefix for global topics available from the remote Home Cache Cluster. A wildcard (/&gt;) is implied at the end of the prefix. | 
 **optional** | ***MsgVpnApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixResponse**](MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixes**
> MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesResponse GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixes(ctx, msgVpnName, cacheName, clusterName, homeClusterName, optional)
Get a list of Topic Prefix objects.

Get a list of Topic Prefix objects.  A Topic Prefix is a prefix for a global topic that is available from the containing Home Cache Cluster.   Attribute|Identifying|Deprecated :---|:---:|:---: cacheName|x| clusterName|x| homeClusterName|x| msgVpnName|x| topicPrefix|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **cacheName** | **string**| The name of the Distributed Cache. | 
  **clusterName** | **string**| The name of the Cache Cluster. | 
  **homeClusterName** | **string**| The name of the remote Home Cache Cluster. | 
 **optional** | ***MsgVpnApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesResponse**](MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusters**
> MsgVpnDistributedCacheClusterGlobalCachingHomeClustersResponse GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusters(ctx, msgVpnName, cacheName, clusterName, optional)
Get a list of Home Cache Cluster objects.

Get a list of Home Cache Cluster objects.  A Home Cache Cluster is a Cache Cluster that is the \"definitive\" Cache Cluster for a given topic in the context of the Global Caching feature.   Attribute|Identifying|Deprecated :---|:---:|:---: cacheName|x| clusterName|x| homeClusterName|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **cacheName** | **string**| The name of the Distributed Cache. | 
  **clusterName** | **string**| The name of the Cache Cluster. | 
 **optional** | ***MsgVpnApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClustersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClustersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterGlobalCachingHomeClustersResponse**](MsgVpnDistributedCacheClusterGlobalCachingHomeClustersResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnDistributedCacheClusterInstance**
> MsgVpnDistributedCacheClusterInstanceResponse GetMsgVpnDistributedCacheClusterInstance(ctx, msgVpnName, cacheName, clusterName, instanceName, optional)
Get a Cache Instance object.

Get a Cache Instance object.  A Cache Instance is a single Cache process that belongs to a single Cache Cluster. A Cache Instance object provisioned on the broker is used to disseminate configuration information to the Cache process. Cache Instances listen for and cache live data messages that match the topic subscriptions configured for their parent Cache Cluster.   Attribute|Identifying|Deprecated :---|:---:|:---: cacheName|x| clusterName|x| counter.msgCount||x counter.msgPeakCount||x counter.requestQueueDepthCount||x counter.requestQueueDepthPeakCount||x counter.topicCount||x counter.topicPeakCount||x instanceName|x| msgVpnName|x| rate.averageDataRxBytePeakRate||x rate.averageDataRxByteRate||x rate.averageDataRxMsgPeakRate||x rate.averageDataRxMsgRate||x rate.averageDataTxMsgPeakRate||x rate.averageDataTxMsgRate||x rate.averageRequestRxPeakRate||x rate.averageRequestRxRate||x rate.dataRxBytePeakRate||x rate.dataRxByteRate||x rate.dataRxMsgPeakRate||x rate.dataRxMsgRate||x rate.dataTxMsgPeakRate||x rate.dataTxMsgRate||x rate.requestRxPeakRate||x rate.requestRxRate||x    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **cacheName** | **string**| The name of the Distributed Cache. | 
  **clusterName** | **string**| The name of the Cache Cluster. | 
  **instanceName** | **string**| The name of the Cache Instance. | 
 **optional** | ***MsgVpnApiGetMsgVpnDistributedCacheClusterInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnDistributedCacheClusterInstanceOpts struct
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

# **GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster**
> MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterResponse GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster(ctx, msgVpnName, cacheName, clusterName, instanceName, homeClusterName, optional)
Get a Remote Home Cache Cluster object.

Get a Remote Home Cache Cluster object.  A Remote Home Cache Cluster is a Home Cache Cluster that the Cache Instance is communicating with in the context of the Global Caching feature.   Attribute|Identifying|Deprecated :---|:---:|:---: cacheName|x| clusterName|x| homeClusterName|x| instanceName|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **cacheName** | **string**| The name of the Distributed Cache. | 
  **clusterName** | **string**| The name of the Cache Cluster. | 
  **instanceName** | **string**| The name of the Cache Instance. | 
  **homeClusterName** | **string**| The name of the remote Home Cache Cluster. | 
 **optional** | ***MsgVpnApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterResponse**](MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusters**
> MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersResponse GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusters(ctx, msgVpnName, cacheName, clusterName, instanceName, optional)
Get a list of Remote Home Cache Cluster objects.

Get a list of Remote Home Cache Cluster objects.  A Remote Home Cache Cluster is a Home Cache Cluster that the Cache Instance is communicating with in the context of the Global Caching feature.   Attribute|Identifying|Deprecated :---|:---:|:---: cacheName|x| clusterName|x| homeClusterName|x| instanceName|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **cacheName** | **string**| The name of the Distributed Cache. | 
  **clusterName** | **string**| The name of the Cache Cluster. | 
  **instanceName** | **string**| The name of the Cache Instance. | 
 **optional** | ***MsgVpnApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersResponse**](MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnDistributedCacheClusterInstanceRemoteTopic**
> MsgVpnDistributedCacheClusterInstanceRemoteTopicResponse GetMsgVpnDistributedCacheClusterInstanceRemoteTopic(ctx, msgVpnName, cacheName, clusterName, instanceName, topic, optional)
Get a Remote Topic object.

Get a Remote Topic object.  A Remote Topic is a topic for which the Cache Instance has cached messages.   Attribute|Identifying|Deprecated :---|:---:|:---: cacheName|x| clusterName|x| instanceName|x| msgVpnName|x| topic|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **cacheName** | **string**| The name of the Distributed Cache. | 
  **clusterName** | **string**| The name of the Cache Cluster. | 
  **instanceName** | **string**| The name of the Cache Instance. | 
  **topic** | **string**| The value of the remote Topic. | 
 **optional** | ***MsgVpnApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterInstanceRemoteTopicResponse**](MsgVpnDistributedCacheClusterInstanceRemoteTopicResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnDistributedCacheClusterInstanceRemoteTopics**
> MsgVpnDistributedCacheClusterInstanceRemoteTopicsResponse GetMsgVpnDistributedCacheClusterInstanceRemoteTopics(ctx, msgVpnName, cacheName, clusterName, instanceName, optional)
Get a list of Remote Topic objects.

Get a list of Remote Topic objects.  A Remote Topic is a topic for which the Cache Instance has cached messages.   Attribute|Identifying|Deprecated :---|:---:|:---: cacheName|x| clusterName|x| instanceName|x| msgVpnName|x| topic|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **cacheName** | **string**| The name of the Distributed Cache. | 
  **clusterName** | **string**| The name of the Cache Cluster. | 
  **instanceName** | **string**| The name of the Cache Instance. | 
 **optional** | ***MsgVpnApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterInstanceRemoteTopicsResponse**](MsgVpnDistributedCacheClusterInstanceRemoteTopicsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnDistributedCacheClusterInstances**
> MsgVpnDistributedCacheClusterInstancesResponse GetMsgVpnDistributedCacheClusterInstances(ctx, msgVpnName, cacheName, clusterName, optional)
Get a list of Cache Instance objects.

Get a list of Cache Instance objects.  A Cache Instance is a single Cache process that belongs to a single Cache Cluster. A Cache Instance object provisioned on the broker is used to disseminate configuration information to the Cache process. Cache Instances listen for and cache live data messages that match the topic subscriptions configured for their parent Cache Cluster.   Attribute|Identifying|Deprecated :---|:---:|:---: cacheName|x| clusterName|x| counter.msgCount||x counter.msgPeakCount||x counter.requestQueueDepthCount||x counter.requestQueueDepthPeakCount||x counter.topicCount||x counter.topicPeakCount||x instanceName|x| msgVpnName|x| rate.averageDataRxBytePeakRate||x rate.averageDataRxByteRate||x rate.averageDataRxMsgPeakRate||x rate.averageDataRxMsgRate||x rate.averageDataTxMsgPeakRate||x rate.averageDataTxMsgRate||x rate.averageRequestRxPeakRate||x rate.averageRequestRxRate||x rate.dataRxBytePeakRate||x rate.dataRxByteRate||x rate.dataRxMsgPeakRate||x rate.dataRxMsgRate||x rate.dataTxMsgPeakRate||x rate.dataTxMsgRate||x rate.requestRxPeakRate||x rate.requestRxRate||x    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **cacheName** | **string**| The name of the Distributed Cache. | 
  **clusterName** | **string**| The name of the Cache Cluster. | 
 **optional** | ***MsgVpnApiGetMsgVpnDistributedCacheClusterInstancesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnDistributedCacheClusterInstancesOpts struct
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

# **GetMsgVpnDistributedCacheClusterTopic**
> MsgVpnDistributedCacheClusterTopicResponse GetMsgVpnDistributedCacheClusterTopic(ctx, msgVpnName, cacheName, clusterName, topic, optional)
Get a Topic object.

Get a Topic object.  The Cache Instances that belong to the containing Cache Cluster will cache any messages published to topics that match a Topic Subscription.   Attribute|Identifying|Deprecated :---|:---:|:---: cacheName|x| clusterName|x| msgVpnName|x| topic|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **cacheName** | **string**| The name of the Distributed Cache. | 
  **clusterName** | **string**| The name of the Cache Cluster. | 
  **topic** | **string**| The value of the Topic in the form a/b/c. | 
 **optional** | ***MsgVpnApiGetMsgVpnDistributedCacheClusterTopicOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnDistributedCacheClusterTopicOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterTopicResponse**](MsgVpnDistributedCacheClusterTopicResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnDistributedCacheClusterTopics**
> MsgVpnDistributedCacheClusterTopicsResponse GetMsgVpnDistributedCacheClusterTopics(ctx, msgVpnName, cacheName, clusterName, optional)
Get a list of Topic objects.

Get a list of Topic objects.  The Cache Instances that belong to the containing Cache Cluster will cache any messages published to topics that match a Topic Subscription.   Attribute|Identifying|Deprecated :---|:---:|:---: cacheName|x| clusterName|x| msgVpnName|x| topic|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **cacheName** | **string**| The name of the Distributed Cache. | 
  **clusterName** | **string**| The name of the Cache Cluster. | 
 **optional** | ***MsgVpnApiGetMsgVpnDistributedCacheClusterTopicsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnDistributedCacheClusterTopicsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterTopicsResponse**](MsgVpnDistributedCacheClusterTopicsResponse.md)

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
 **optional** | ***MsgVpnApiGetMsgVpnDistributedCacheClustersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnDistributedCacheClustersOpts struct
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
 **optional** | ***MsgVpnApiGetMsgVpnDistributedCachesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnDistributedCachesOpts struct
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

# **GetMsgVpnDmrBridge**
> MsgVpnDmrBridgeResponse GetMsgVpnDmrBridge(ctx, msgVpnName, remoteNodeName, optional)
Get a DMR Bridge object.

Get a DMR Bridge object.  A DMR Bridge is required to establish a data channel over a corresponding external link to the remote node for a given Message VPN. Each DMR Bridge identifies which external link the Message VPN should use, and what the name of the equivalent Message VPN at the remote node is.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| remoteNodeName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **remoteNodeName** | **string**| The name of the node at the remote end of the DMR Bridge. | 
 **optional** | ***MsgVpnApiGetMsgVpnDmrBridgeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnDmrBridgeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

Get a list of DMR Bridge objects.  A DMR Bridge is required to establish a data channel over a corresponding external link to the remote node for a given Message VPN. Each DMR Bridge identifies which external link the Message VPN should use, and what the name of the equivalent Message VPN at the remote node is.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| remoteNodeName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MsgVpnApiGetMsgVpnDmrBridgesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnDmrBridgesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
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

# **GetMsgVpnJndiConnectionFactories**
> MsgVpnJndiConnectionFactoriesResponse GetMsgVpnJndiConnectionFactories(ctx, msgVpnName, optional)
Get a list of JNDI Connection Factory objects.

Get a list of JNDI Connection Factory objects.  The message broker provides an internal JNDI store for provisioned Connection Factory objects that clients can access through JNDI lookups.   Attribute|Identifying|Deprecated :---|:---:|:---: connectionFactoryName|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MsgVpnApiGetMsgVpnJndiConnectionFactoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnJndiConnectionFactoriesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnJndiConnectionFactoriesResponse**](MsgVpnJndiConnectionFactoriesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnJndiConnectionFactory**
> MsgVpnJndiConnectionFactoryResponse GetMsgVpnJndiConnectionFactory(ctx, msgVpnName, connectionFactoryName, optional)
Get a JNDI Connection Factory object.

Get a JNDI Connection Factory object.  The message broker provides an internal JNDI store for provisioned Connection Factory objects that clients can access through JNDI lookups.   Attribute|Identifying|Deprecated :---|:---:|:---: connectionFactoryName|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **connectionFactoryName** | **string**| The name of the JMS Connection Factory. | 
 **optional** | ***MsgVpnApiGetMsgVpnJndiConnectionFactoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnJndiConnectionFactoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnJndiConnectionFactoryResponse**](MsgVpnJndiConnectionFactoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnJndiQueue**
> MsgVpnJndiQueueResponse GetMsgVpnJndiQueue(ctx, msgVpnName, queueName, optional)
Get a JNDI Queue object.

Get a JNDI Queue object.  The message broker provides an internal JNDI store for provisioned Queue objects that clients can access through JNDI lookups.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| queueName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The JNDI name of the JMS Queue. | 
 **optional** | ***MsgVpnApiGetMsgVpnJndiQueueOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnJndiQueueOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnJndiQueueResponse**](MsgVpnJndiQueueResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnJndiQueues**
> MsgVpnJndiQueuesResponse GetMsgVpnJndiQueues(ctx, msgVpnName, optional)
Get a list of JNDI Queue objects.

Get a list of JNDI Queue objects.  The message broker provides an internal JNDI store for provisioned Queue objects that clients can access through JNDI lookups.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| queueName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MsgVpnApiGetMsgVpnJndiQueuesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnJndiQueuesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnJndiQueuesResponse**](MsgVpnJndiQueuesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnJndiTopic**
> MsgVpnJndiTopicResponse GetMsgVpnJndiTopic(ctx, msgVpnName, topicName, optional)
Get a JNDI Topic object.

Get a JNDI Topic object.  The message broker provides an internal JNDI store for provisioned Topic objects that clients can access through JNDI lookups.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| topicName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicName** | **string**| The JNDI name of the JMS Topic. | 
 **optional** | ***MsgVpnApiGetMsgVpnJndiTopicOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnJndiTopicOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnJndiTopicResponse**](MsgVpnJndiTopicResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnJndiTopics**
> MsgVpnJndiTopicsResponse GetMsgVpnJndiTopics(ctx, msgVpnName, optional)
Get a list of JNDI Topic objects.

Get a list of JNDI Topic objects.  The message broker provides an internal JNDI store for provisioned Topic objects that clients can access through JNDI lookups.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| topicName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MsgVpnApiGetMsgVpnJndiTopicsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnJndiTopicsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnJndiTopicsResponse**](MsgVpnJndiTopicsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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
 **optional** | ***MsgVpnApiGetMsgVpnMqttRetainCacheOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnMqttRetainCacheOpts struct
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
 **optional** | ***MsgVpnApiGetMsgVpnMqttRetainCachesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnMqttRetainCachesOpts struct
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
 **optional** | ***MsgVpnApiGetMsgVpnMqttSessionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnMqttSessionOpts struct
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
 **optional** | ***MsgVpnApiGetMsgVpnMqttSessionSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnMqttSessionSubscriptionOpts struct
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
 **optional** | ***MsgVpnApiGetMsgVpnMqttSessionSubscriptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnMqttSessionSubscriptionsOpts struct
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
 **optional** | ***MsgVpnApiGetMsgVpnMqttSessionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnMqttSessionsOpts struct
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

Get a Queue object.  A Queue acts as both a destination that clients can publish messages to, and as an endpoint that clients can bind consumers to and consume messages from.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| queueName|x| virtualRouter||x    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The name of the Queue. | 
 **optional** | ***MsgVpnApiGetMsgVpnQueueOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnQueueOpts struct
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

Get a Queue Message object.  A Queue Message is a packet of information sent from producers to consumers using the Queue.   Attribute|Identifying|Deprecated :---|:---:|:---: msgId|x| msgVpnName|x| queueName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The name of the Queue. | 
  **msgId** | **string**| The identifier (ID) of the Message. | 
 **optional** | ***MsgVpnApiGetMsgVpnQueueMsgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnQueueMsgOpts struct
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

Get a list of Queue Message objects.  A Queue Message is a packet of information sent from producers to consumers using the Queue.   Attribute|Identifying|Deprecated :---|:---:|:---: msgId|x| msgVpnName|x| queueName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The name of the Queue. | 
 **optional** | ***MsgVpnApiGetMsgVpnQueueMsgsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnQueueMsgsOpts struct
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

# **GetMsgVpnQueuePriorities**
> MsgVpnQueuePrioritiesResponse GetMsgVpnQueuePriorities(ctx, msgVpnName, queueName, optional)
Get a list of Queue Priority objects.

Get a list of Queue Priority objects.  Queues can optionally support priority message delivery; all messages of a higher priority are delivered before any messages of a lower priority. A Priority object contains information about the number and size of the messages with a particular priority in the Queue.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| priority|x| queueName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The name of the Queue. | 
 **optional** | ***MsgVpnApiGetMsgVpnQueuePrioritiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnQueuePrioritiesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueuePrioritiesResponse**](MsgVpnQueuePrioritiesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnQueuePriority**
> MsgVpnQueuePriorityResponse GetMsgVpnQueuePriority(ctx, msgVpnName, queueName, priority, optional)
Get a Queue Priority object.

Get a Queue Priority object.  Queues can optionally support priority message delivery; all messages of a higher priority are delivered before any messages of a lower priority. A Priority object contains information about the number and size of the messages with a particular priority in the Queue.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| priority|x| queueName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The name of the Queue. | 
  **priority** | **string**| The level of the Priority, from 9 (highest) to 0 (lowest). | 
 **optional** | ***MsgVpnApiGetMsgVpnQueuePriorityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnQueuePriorityOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueuePriorityResponse**](MsgVpnQueuePriorityResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnQueueSubscription**
> MsgVpnQueueSubscriptionResponse GetMsgVpnQueueSubscription(ctx, msgVpnName, queueName, subscriptionTopic, optional)
Get a Queue Subscription object.

Get a Queue Subscription object.  One or more Queue Subscriptions can be added to a durable queue so that Guaranteed messages published to matching topics are also delivered to and spooled by the queue.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| queueName|x| subscriptionTopic|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The name of the Queue. | 
  **subscriptionTopic** | **string**| The topic of the Subscription. | 
 **optional** | ***MsgVpnApiGetMsgVpnQueueSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnQueueSubscriptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueueSubscriptionResponse**](MsgVpnQueueSubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnQueueSubscriptions**
> MsgVpnQueueSubscriptionsResponse GetMsgVpnQueueSubscriptions(ctx, msgVpnName, queueName, optional)
Get a list of Queue Subscription objects.

Get a list of Queue Subscription objects.  One or more Queue Subscriptions can be added to a durable queue so that Guaranteed messages published to matching topics are also delivered to and spooled by the queue.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| queueName|x| subscriptionTopic|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The name of the Queue. | 
 **optional** | ***MsgVpnApiGetMsgVpnQueueSubscriptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnQueueSubscriptionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueueSubscriptionsResponse**](MsgVpnQueueSubscriptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnQueueTemplate**
> MsgVpnQueueTemplateResponse GetMsgVpnQueueTemplate(ctx, msgVpnName, queueTemplateName, optional)
Get a Queue Template object.

Get a Queue Template object.  A Queue Template provides a mechanism for specifying the initial state for client created queues.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| queueTemplateName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueTemplateName** | **string**| The name of the Queue Template. | 
 **optional** | ***MsgVpnApiGetMsgVpnQueueTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnQueueTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueueTemplateResponse**](MsgVpnQueueTemplateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnQueueTemplates**
> MsgVpnQueueTemplatesResponse GetMsgVpnQueueTemplates(ctx, msgVpnName, optional)
Get a list of Queue Template objects.

Get a list of Queue Template objects.  A Queue Template provides a mechanism for specifying the initial state for client created queues.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| queueTemplateName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MsgVpnApiGetMsgVpnQueueTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnQueueTemplatesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueueTemplatesResponse**](MsgVpnQueueTemplatesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnQueueTxFlow**
> MsgVpnQueueTxFlowResponse GetMsgVpnQueueTxFlow(ctx, msgVpnName, queueName, flowId, optional)
Get a Queue Transmit Flow object.

Get a Queue Transmit Flow object.  Queue Transmit Flows are used by clients to consume Guaranteed messages from a Queue.   Attribute|Identifying|Deprecated :---|:---:|:---: flowId|x| msgVpnName|x| queueName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The name of the Queue. | 
  **flowId** | **string**| The identifier (ID) of the Flow. | 
 **optional** | ***MsgVpnApiGetMsgVpnQueueTxFlowOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnQueueTxFlowOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueueTxFlowResponse**](MsgVpnQueueTxFlowResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnQueueTxFlows**
> MsgVpnQueueTxFlowsResponse GetMsgVpnQueueTxFlows(ctx, msgVpnName, queueName, optional)
Get a list of Queue Transmit Flow objects.

Get a list of Queue Transmit Flow objects.  Queue Transmit Flows are used by clients to consume Guaranteed messages from a Queue.   Attribute|Identifying|Deprecated :---|:---:|:---: flowId|x| msgVpnName|x| queueName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The name of the Queue. | 
 **optional** | ***MsgVpnApiGetMsgVpnQueueTxFlowsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnQueueTxFlowsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueueTxFlowsResponse**](MsgVpnQueueTxFlowsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnQueues**
> MsgVpnQueuesResponse GetMsgVpnQueues(ctx, msgVpnName, optional)
Get a list of Queue objects.

Get a list of Queue objects.  A Queue acts as both a destination that clients can publish messages to, and as an endpoint that clients can bind consumers to and consume messages from.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| queueName|x| virtualRouter||x    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MsgVpnApiGetMsgVpnQueuesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnQueuesOpts struct
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
 **optional** | ***MsgVpnApiGetMsgVpnReplayLogOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnReplayLogOpts struct
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

# **GetMsgVpnReplayLogMsg**
> MsgVpnReplayLogMsgResponse GetMsgVpnReplayLogMsg(ctx, msgVpnName, replayLogName, msgId, optional)
Get a Message object.

Get a Message object.  A Message is a packet of information sent from producers to consumers. Messages are the central units of information that clients exchange using the message broker and which are cached in the Replay Log.   Attribute|Identifying|Deprecated :---|:---:|:---: msgId|x| msgVpnName|x| replayLogName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **replayLogName** | **string**| The name of the Replay Log. | 
  **msgId** | **string**| The identifier (ID) of the message. | 
 **optional** | ***MsgVpnApiGetMsgVpnReplayLogMsgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnReplayLogMsgOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnReplayLogMsgResponse**](MsgVpnReplayLogMsgResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnReplayLogMsgs**
> MsgVpnReplayLogMsgsResponse GetMsgVpnReplayLogMsgs(ctx, msgVpnName, replayLogName, optional)
Get a list of Message objects.

Get a list of Message objects.  A Message is a packet of information sent from producers to consumers. Messages are the central units of information that clients exchange using the message broker and which are cached in the Replay Log.   Attribute|Identifying|Deprecated :---|:---:|:---: msgId|x| msgVpnName|x| replayLogName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **replayLogName** | **string**| The name of the Replay Log. | 
 **optional** | ***MsgVpnApiGetMsgVpnReplayLogMsgsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnReplayLogMsgsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnReplayLogMsgsResponse**](MsgVpnReplayLogMsgsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnReplayLogTopicFilterSubscription**
> MsgVpnReplayLogTopicFilterSubscriptionResponse GetMsgVpnReplayLogTopicFilterSubscription(ctx, msgVpnName, replayLogName, topicFilterSubscription, optional)
Get a Topic Filter Subscription object.

Get a Topic Filter Subscription object.  One or more Subscriptions can be added to a replay-log so that only guaranteed messages published to matching topics are stored in the Replay Log.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| replayLogName|x| topicFilterSubscription|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.27.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **replayLogName** | **string**| The name of the Replay Log. | 
  **topicFilterSubscription** | **string**| The topic of the Subscription. | 
 **optional** | ***MsgVpnApiGetMsgVpnReplayLogTopicFilterSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnReplayLogTopicFilterSubscriptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnReplayLogTopicFilterSubscriptionResponse**](MsgVpnReplayLogTopicFilterSubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnReplayLogTopicFilterSubscriptions**
> MsgVpnReplayLogTopicFilterSubscriptionsResponse GetMsgVpnReplayLogTopicFilterSubscriptions(ctx, msgVpnName, replayLogName, optional)
Get a list of Topic Filter Subscription objects.

Get a list of Topic Filter Subscription objects.  One or more Subscriptions can be added to a replay-log so that only guaranteed messages published to matching topics are stored in the Replay Log.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| replayLogName|x| topicFilterSubscription|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.27.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **replayLogName** | **string**| The name of the Replay Log. | 
 **optional** | ***MsgVpnApiGetMsgVpnReplayLogTopicFilterSubscriptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnReplayLogTopicFilterSubscriptionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnReplayLogTopicFilterSubscriptionsResponse**](MsgVpnReplayLogTopicFilterSubscriptionsResponse.md)

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
 **optional** | ***MsgVpnApiGetMsgVpnReplayLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnReplayLogsOpts struct
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

# **GetMsgVpnReplicatedTopic**
> MsgVpnReplicatedTopicResponse GetMsgVpnReplicatedTopic(ctx, msgVpnName, replicatedTopic, optional)
Get a Replicated Topic object.

Get a Replicated Topic object.  To indicate which messages should be replicated between the active and standby site, a Replicated Topic subscription must be configured on a Message VPN. If a published message matches both a replicated topic and an endpoint on the active site, then the message is replicated to the standby site.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| replicatedTopic|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **replicatedTopic** | **string**| The topic for applying replication. Published messages matching this topic will be replicated to the standby site. | 
 **optional** | ***MsgVpnApiGetMsgVpnReplicatedTopicOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnReplicatedTopicOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnReplicatedTopicResponse**](MsgVpnReplicatedTopicResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnReplicatedTopics**
> MsgVpnReplicatedTopicsResponse GetMsgVpnReplicatedTopics(ctx, msgVpnName, optional)
Get a list of Replicated Topic objects.

Get a list of Replicated Topic objects.  To indicate which messages should be replicated between the active and standby site, a Replicated Topic subscription must be configured on a Message VPN. If a published message matches both a replicated topic and an endpoint on the active site, then the message is replicated to the standby site.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| replicatedTopic|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MsgVpnApiGetMsgVpnReplicatedTopicsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnReplicatedTopicsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnReplicatedTopicsResponse**](MsgVpnReplicatedTopicsResponse.md)

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
 **optional** | ***MsgVpnApiGetMsgVpnRestDeliveryPointOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnRestDeliveryPointOpts struct
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

# **GetMsgVpnRestDeliveryPointQueueBinding**
> MsgVpnRestDeliveryPointQueueBindingResponse GetMsgVpnRestDeliveryPointQueueBinding(ctx, msgVpnName, restDeliveryPointName, queueBindingName, optional)
Get a Queue Binding object.

Get a Queue Binding object.  A Queue Binding for a REST Delivery Point attracts messages to be delivered to REST consumers. If the queue does not exist it can be created subsequently, and once the queue is operational the broker performs the queue binding. Removing the queue binding does not delete the queue itself. Similarly, removing the queue does not remove the queue binding, which fails until the queue is recreated or the queue binding is deleted.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| queueBindingName|x| restDeliveryPointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **queueBindingName** | **string**| The name of a queue in the Message VPN. | 
 **optional** | ***MsgVpnApiGetMsgVpnRestDeliveryPointQueueBindingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnRestDeliveryPointQueueBindingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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

Get a Protected Request Header object.  A protected request header to be added to the HTTP request. Unlike a non-protected request header, the header value cannot be displayed after it is set.   Attribute|Identifying|Deprecated :---|:---:|:---: headerName|x| msgVpnName|x| queueBindingName|x| restDeliveryPointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.30.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **queueBindingName** | **string**| The name of a queue in the Message VPN. | 
  **headerName** | **string**| The name of the protected HTTP request header. | 
 **optional** | ***MsgVpnApiGetMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeaderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeaderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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

Get a list of Protected Request Header objects.  A protected request header to be added to the HTTP request. Unlike a non-protected request header, the header value cannot be displayed after it is set.   Attribute|Identifying|Deprecated :---|:---:|:---: headerName|x| msgVpnName|x| queueBindingName|x| restDeliveryPointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.30.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **queueBindingName** | **string**| The name of a queue in the Message VPN. | 
 **optional** | ***MsgVpnApiGetMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeadersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeadersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
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

Get a Request Header object.  A request header to be added to the HTTP request.   Attribute|Identifying|Deprecated :---|:---:|:---: headerName|x| msgVpnName|x| queueBindingName|x| restDeliveryPointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.23.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **queueBindingName** | **string**| The name of a queue in the Message VPN. | 
  **headerName** | **string**| The name of the HTTP request header. | 
 **optional** | ***MsgVpnApiGetMsgVpnRestDeliveryPointQueueBindingRequestHeaderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnRestDeliveryPointQueueBindingRequestHeaderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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

Get a list of Request Header objects.  A request header to be added to the HTTP request.   Attribute|Identifying|Deprecated :---|:---:|:---: headerName|x| msgVpnName|x| queueBindingName|x| restDeliveryPointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.23.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **queueBindingName** | **string**| The name of a queue in the Message VPN. | 
 **optional** | ***MsgVpnApiGetMsgVpnRestDeliveryPointQueueBindingRequestHeadersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnRestDeliveryPointQueueBindingRequestHeadersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
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

Get a list of Queue Binding objects.  A Queue Binding for a REST Delivery Point attracts messages to be delivered to REST consumers. If the queue does not exist it can be created subsequently, and once the queue is operational the broker performs the queue binding. Removing the queue binding does not delete the queue itself. Similarly, removing the queue does not remove the queue binding, which fails until the queue is recreated or the queue binding is deleted.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| queueBindingName|x| restDeliveryPointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
 **optional** | ***MsgVpnApiGetMsgVpnRestDeliveryPointQueueBindingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnRestDeliveryPointQueueBindingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
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

Get a REST Consumer object.  REST Consumer objects establish HTTP connectivity to REST consumer applications who wish to receive messages from a broker.   Attribute|Identifying|Deprecated :---|:---:|:---: counter.httpRequestConnectionCloseTxMsgCount||x counter.httpRequestOutstandingTxMsgCount||x counter.httpRequestTimedOutTxMsgCount||x counter.httpRequestTxByteCount||x counter.httpRequestTxMsgCount||x counter.httpResponseErrorRxMsgCount||x counter.httpResponseRxByteCount||x counter.httpResponseRxMsgCount||x counter.httpResponseSuccessRxMsgCount||x msgVpnName|x| restConsumerName|x| restDeliveryPointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **restConsumerName** | **string**| The name of the REST Consumer. | 
 **optional** | ***MsgVpnApiGetMsgVpnRestDeliveryPointRestConsumerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnRestDeliveryPointRestConsumerOpts struct
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

# **GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim**
> MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimResponse GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim(ctx, msgVpnName, restDeliveryPointName, restConsumerName, oauthJwtClaimName, optional)
Get a Claim object.

Get a Claim object.  A Claim is added to the JWT sent to the OAuth token request endpoint.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| oauthJwtClaimName|x| restConsumerName|x| restDeliveryPointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.21.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **restConsumerName** | **string**| The name of the REST Consumer. | 
  **oauthJwtClaimName** | **string**| The name of the additional claim. Cannot be \&quot;exp\&quot;, \&quot;iat\&quot;, or \&quot;jti\&quot;. | 
 **optional** | ***MsgVpnApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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

Get a list of Claim objects.  A Claim is added to the JWT sent to the OAuth token request endpoint.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| oauthJwtClaimName|x| restConsumerName|x| restDeliveryPointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.21.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **restConsumerName** | **string**| The name of the REST Consumer. | 
 **optional** | ***MsgVpnApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
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

Get a Trusted Common Name object.  The Trusted Common Names for the REST Consumer are used by encrypted transports to verify the name in the certificate presented by the remote REST consumer. They must include the common name of the remote REST consumer's server certificate.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x|x restConsumerName|x|x restDeliveryPointName|x|x tlsTrustedCommonName|x|x    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been deprecated since 2.17. Common Name validation has been replaced by Server Certificate Name validation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **restConsumerName** | **string**| The name of the REST Consumer. | 
  **tlsTrustedCommonName** | **string**| The expected trusted common name of the remote certificate. | 
 **optional** | ***MsgVpnApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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

Get a list of Trusted Common Name objects.  The Trusted Common Names for the REST Consumer are used by encrypted transports to verify the name in the certificate presented by the remote REST consumer. They must include the common name of the remote REST consumer's server certificate.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x|x restConsumerName|x|x restDeliveryPointName|x|x tlsTrustedCommonName|x|x    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been deprecated since 2.17. Common Name validation has been replaced by Server Certificate Name validation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **restConsumerName** | **string**| The name of the REST Consumer. | 
 **optional** | ***MsgVpnApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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

Get a list of REST Consumer objects.  REST Consumer objects establish HTTP connectivity to REST consumer applications who wish to receive messages from a broker.   Attribute|Identifying|Deprecated :---|:---:|:---: counter.httpRequestConnectionCloseTxMsgCount||x counter.httpRequestOutstandingTxMsgCount||x counter.httpRequestTimedOutTxMsgCount||x counter.httpRequestTxByteCount||x counter.httpRequestTxMsgCount||x counter.httpResponseErrorRxMsgCount||x counter.httpResponseRxByteCount||x counter.httpResponseRxMsgCount||x counter.httpResponseSuccessRxMsgCount||x msgVpnName|x| restConsumerName|x| restDeliveryPointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
 **optional** | ***MsgVpnApiGetMsgVpnRestDeliveryPointRestConsumersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnRestDeliveryPointRestConsumersOpts struct
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
 **optional** | ***MsgVpnApiGetMsgVpnRestDeliveryPointsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnRestDeliveryPointsOpts struct
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

# **GetMsgVpnTelemetryProfile**
> MsgVpnTelemetryProfileResponse GetMsgVpnTelemetryProfile(ctx, msgVpnName, telemetryProfileName, optional)
Get a Telemetry Profile object.

Get a Telemetry Profile object.  Using the Telemetry Profile allows trace spans to be generated as messages are processed by the broker. The generated spans are stored persistently on the broker and may be consumed by the Solace receiver component of an OpenTelemetry Collector.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| telemetryProfileName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.31.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **telemetryProfileName** | **string**| The name of the Telemetry Profile. | 
 **optional** | ***MsgVpnApiGetMsgVpnTelemetryProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnTelemetryProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTelemetryProfileResponse**](MsgVpnTelemetryProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnTelemetryProfileReceiverAclConnectException**
> MsgVpnTelemetryProfileReceiverAclConnectExceptionResponse GetMsgVpnTelemetryProfileReceiverAclConnectException(ctx, msgVpnName, telemetryProfileName, receiverAclConnectExceptionAddress, optional)
Get a Receiver ACL Connect Exception object.

Get a Receiver ACL Connect Exception object.  A Receiver ACL Connect Exception is an exception to the default action to take when a receiver connects to the broker. Exceptions must be expressed as an IP address/netmask in CIDR form.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| receiverAclConnectExceptionAddress|x| telemetryProfileName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.31.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **telemetryProfileName** | **string**| The name of the Telemetry Profile. | 
  **receiverAclConnectExceptionAddress** | **string**| The IP address/netmask of the receiver connect exception in CIDR form. | 
 **optional** | ***MsgVpnApiGetMsgVpnTelemetryProfileReceiverAclConnectExceptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnTelemetryProfileReceiverAclConnectExceptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTelemetryProfileReceiverAclConnectExceptionResponse**](MsgVpnTelemetryProfileReceiverAclConnectExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnTelemetryProfileReceiverAclConnectExceptions**
> MsgVpnTelemetryProfileReceiverAclConnectExceptionsResponse GetMsgVpnTelemetryProfileReceiverAclConnectExceptions(ctx, msgVpnName, telemetryProfileName, optional)
Get a list of Receiver ACL Connect Exception objects.

Get a list of Receiver ACL Connect Exception objects.  A Receiver ACL Connect Exception is an exception to the default action to take when a receiver connects to the broker. Exceptions must be expressed as an IP address/netmask in CIDR form.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| receiverAclConnectExceptionAddress|x| telemetryProfileName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.31.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **telemetryProfileName** | **string**| The name of the Telemetry Profile. | 
 **optional** | ***MsgVpnApiGetMsgVpnTelemetryProfileReceiverAclConnectExceptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnTelemetryProfileReceiverAclConnectExceptionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTelemetryProfileReceiverAclConnectExceptionsResponse**](MsgVpnTelemetryProfileReceiverAclConnectExceptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnTelemetryProfileTraceFilter**
> MsgVpnTelemetryProfileTraceFilterResponse GetMsgVpnTelemetryProfileTraceFilter(ctx, msgVpnName, telemetryProfileName, traceFilterName, optional)
Get a Trace Filter object.

Get a Trace Filter object.  A Trace Filter controls which messages received by the broker will be traced. If an incoming message matches an enabled tracing filter's subscription, the message will be traced as it passes through the broker.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| telemetryProfileName|x| traceFilterName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.31.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **telemetryProfileName** | **string**| The name of the Telemetry Profile. | 
  **traceFilterName** | **string**| A name used to identify the trace filter. Consider a name that describes the subscriptions contained within the filter, such as the name of the application and/or the scenario in which the trace filter might be enabled, such as \&quot;appNameDebug\&quot;. | 
 **optional** | ***MsgVpnApiGetMsgVpnTelemetryProfileTraceFilterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnTelemetryProfileTraceFilterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTelemetryProfileTraceFilterResponse**](MsgVpnTelemetryProfileTraceFilterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnTelemetryProfileTraceFilterSubscription**
> MsgVpnTelemetryProfileTraceFilterSubscriptionResponse GetMsgVpnTelemetryProfileTraceFilterSubscription(ctx, msgVpnName, telemetryProfileName, traceFilterName, subscription, subscriptionSyntax, optional)
Get a Telemetry Trace Filter Subscription object.

Get a Telemetry Trace Filter Subscription object.  Trace filter subscriptions control which messages will be attracted by the tracing filter.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| subscription|x| subscriptionSyntax|x| telemetryProfileName|x| traceFilterName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.31.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **telemetryProfileName** | **string**| The name of the Telemetry Profile. | 
  **traceFilterName** | **string**| A name used to identify the trace filter. Consider a name that describes the subscriptions contained within the filter, such as the name of the application and/or the scenario in which the trace filter might be enabled, such as \&quot;appNameDebug\&quot;. | 
  **subscription** | **string**| Messages matching this subscription will follow this filter&#x27;s configuration. | 
  **subscriptionSyntax** | **string**| The syntax of the trace filter subscription. | 
 **optional** | ***MsgVpnApiGetMsgVpnTelemetryProfileTraceFilterSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnTelemetryProfileTraceFilterSubscriptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTelemetryProfileTraceFilterSubscriptionResponse**](MsgVpnTelemetryProfileTraceFilterSubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnTelemetryProfileTraceFilterSubscriptions**
> MsgVpnTelemetryProfileTraceFilterSubscriptionsResponse GetMsgVpnTelemetryProfileTraceFilterSubscriptions(ctx, msgVpnName, telemetryProfileName, traceFilterName, optional)
Get a list of Telemetry Trace Filter Subscription objects.

Get a list of Telemetry Trace Filter Subscription objects.  Trace filter subscriptions control which messages will be attracted by the tracing filter.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| subscription|x| subscriptionSyntax|x| telemetryProfileName|x| traceFilterName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.31.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **telemetryProfileName** | **string**| The name of the Telemetry Profile. | 
  **traceFilterName** | **string**| A name used to identify the trace filter. Consider a name that describes the subscriptions contained within the filter, such as the name of the application and/or the scenario in which the trace filter might be enabled, such as \&quot;appNameDebug\&quot;. | 
 **optional** | ***MsgVpnApiGetMsgVpnTelemetryProfileTraceFilterSubscriptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnTelemetryProfileTraceFilterSubscriptionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTelemetryProfileTraceFilterSubscriptionsResponse**](MsgVpnTelemetryProfileTraceFilterSubscriptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnTelemetryProfileTraceFilters**
> MsgVpnTelemetryProfileTraceFiltersResponse GetMsgVpnTelemetryProfileTraceFilters(ctx, msgVpnName, telemetryProfileName, optional)
Get a list of Trace Filter objects.

Get a list of Trace Filter objects.  A Trace Filter controls which messages received by the broker will be traced. If an incoming message matches an enabled tracing filter's subscription, the message will be traced as it passes through the broker.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| telemetryProfileName|x| traceFilterName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.31.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **telemetryProfileName** | **string**| The name of the Telemetry Profile. | 
 **optional** | ***MsgVpnApiGetMsgVpnTelemetryProfileTraceFiltersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnTelemetryProfileTraceFiltersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTelemetryProfileTraceFiltersResponse**](MsgVpnTelemetryProfileTraceFiltersResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnTelemetryProfiles**
> MsgVpnTelemetryProfilesResponse GetMsgVpnTelemetryProfiles(ctx, msgVpnName, optional)
Get a list of Telemetry Profile objects.

Get a list of Telemetry Profile objects.  Using the Telemetry Profile allows trace spans to be generated as messages are processed by the broker. The generated spans are stored persistently on the broker and may be consumed by the Solace receiver component of an OpenTelemetry Collector.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| telemetryProfileName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.31.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MsgVpnApiGetMsgVpnTelemetryProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnTelemetryProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTelemetryProfilesResponse**](MsgVpnTelemetryProfilesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnTopicEndpoint**
> MsgVpnTopicEndpointResponse GetMsgVpnTopicEndpoint(ctx, msgVpnName, topicEndpointName, optional)
Get a Topic Endpoint object.

Get a Topic Endpoint object.  A Topic Endpoint attracts messages published to a topic for which the Topic Endpoint has a matching topic subscription. The topic subscription for the Topic Endpoint is specified in the client request to bind a Flow to that Topic Endpoint. Queues are significantly more flexible than Topic Endpoints and are the recommended approach for most applications. The use of Topic Endpoints should be restricted to JMS applications.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| topicEndpointName|x| virtualRouter||x    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointName** | **string**| The name of the Topic Endpoint. | 
 **optional** | ***MsgVpnApiGetMsgVpnTopicEndpointOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnTopicEndpointOpts struct
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

Get a Topic Endpoint Message object.  A Topic Endpoint Message is a packet of information sent from producers to consumers using the Topic Endpoint.   Attribute|Identifying|Deprecated :---|:---:|:---: msgId|x| msgVpnName|x| topicEndpointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointName** | **string**| The name of the Topic Endpoint. | 
  **msgId** | **string**| The identifier (ID) of the Message. | 
 **optional** | ***MsgVpnApiGetMsgVpnTopicEndpointMsgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnTopicEndpointMsgOpts struct
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

Get a list of Topic Endpoint Message objects.  A Topic Endpoint Message is a packet of information sent from producers to consumers using the Topic Endpoint.   Attribute|Identifying|Deprecated :---|:---:|:---: msgId|x| msgVpnName|x| topicEndpointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointName** | **string**| The name of the Topic Endpoint. | 
 **optional** | ***MsgVpnApiGetMsgVpnTopicEndpointMsgsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnTopicEndpointMsgsOpts struct
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

# **GetMsgVpnTopicEndpointPriorities**
> MsgVpnTopicEndpointPrioritiesResponse GetMsgVpnTopicEndpointPriorities(ctx, msgVpnName, topicEndpointName, optional)
Get a list of Topic Endpoint Priority objects.

Get a list of Topic Endpoint Priority objects.  Topic Endpoints can optionally support priority message delivery; all messages of a higher priority are delivered before any messages of a lower priority. A Priority object contains information about the number and size of the messages with a particular priority in the Topic Endpoint.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| priority|x| topicEndpointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointName** | **string**| The name of the Topic Endpoint. | 
 **optional** | ***MsgVpnApiGetMsgVpnTopicEndpointPrioritiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnTopicEndpointPrioritiesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointPrioritiesResponse**](MsgVpnTopicEndpointPrioritiesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnTopicEndpointPriority**
> MsgVpnTopicEndpointPriorityResponse GetMsgVpnTopicEndpointPriority(ctx, msgVpnName, topicEndpointName, priority, optional)
Get a Topic Endpoint Priority object.

Get a Topic Endpoint Priority object.  Topic Endpoints can optionally support priority message delivery; all messages of a higher priority are delivered before any messages of a lower priority. A Priority object contains information about the number and size of the messages with a particular priority in the Topic Endpoint.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| priority|x| topicEndpointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointName** | **string**| The name of the Topic Endpoint. | 
  **priority** | **string**| The level of the Priority, from 9 (highest) to 0 (lowest). | 
 **optional** | ***MsgVpnApiGetMsgVpnTopicEndpointPriorityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnTopicEndpointPriorityOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointPriorityResponse**](MsgVpnTopicEndpointPriorityResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnTopicEndpointTemplate**
> MsgVpnTopicEndpointTemplateResponse GetMsgVpnTopicEndpointTemplate(ctx, msgVpnName, topicEndpointTemplateName, optional)
Get a Topic Endpoint Template object.

Get a Topic Endpoint Template object.  A Topic Endpoint Template provides a mechanism for specifying the initial state for client created topic endpoints.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| topicEndpointTemplateName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointTemplateName** | **string**| The name of the Topic Endpoint Template. | 
 **optional** | ***MsgVpnApiGetMsgVpnTopicEndpointTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnTopicEndpointTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointTemplateResponse**](MsgVpnTopicEndpointTemplateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnTopicEndpointTemplates**
> MsgVpnTopicEndpointTemplatesResponse GetMsgVpnTopicEndpointTemplates(ctx, msgVpnName, optional)
Get a list of Topic Endpoint Template objects.

Get a list of Topic Endpoint Template objects.  A Topic Endpoint Template provides a mechanism for specifying the initial state for client created topic endpoints.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| topicEndpointTemplateName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MsgVpnApiGetMsgVpnTopicEndpointTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnTopicEndpointTemplatesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointTemplatesResponse**](MsgVpnTopicEndpointTemplatesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnTopicEndpointTxFlow**
> MsgVpnTopicEndpointTxFlowResponse GetMsgVpnTopicEndpointTxFlow(ctx, msgVpnName, topicEndpointName, flowId, optional)
Get a Topic Endpoint Transmit Flow object.

Get a Topic Endpoint Transmit Flow object.  Topic Endpoint Transmit Flows are used by clients to consume Guaranteed messages from a Topic Endpoint.   Attribute|Identifying|Deprecated :---|:---:|:---: flowId|x| msgVpnName|x| topicEndpointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointName** | **string**| The name of the Topic Endpoint. | 
  **flowId** | **string**| The identifier (ID) of the Flow. | 
 **optional** | ***MsgVpnApiGetMsgVpnTopicEndpointTxFlowOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnTopicEndpointTxFlowOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointTxFlowResponse**](MsgVpnTopicEndpointTxFlowResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnTopicEndpointTxFlows**
> MsgVpnTopicEndpointTxFlowsResponse GetMsgVpnTopicEndpointTxFlows(ctx, msgVpnName, topicEndpointName, optional)
Get a list of Topic Endpoint Transmit Flow objects.

Get a list of Topic Endpoint Transmit Flow objects.  Topic Endpoint Transmit Flows are used by clients to consume Guaranteed messages from a Topic Endpoint.   Attribute|Identifying|Deprecated :---|:---:|:---: flowId|x| msgVpnName|x| topicEndpointName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointName** | **string**| The name of the Topic Endpoint. | 
 **optional** | ***MsgVpnApiGetMsgVpnTopicEndpointTxFlowsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnTopicEndpointTxFlowsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointTxFlowsResponse**](MsgVpnTopicEndpointTxFlowsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnTopicEndpoints**
> MsgVpnTopicEndpointsResponse GetMsgVpnTopicEndpoints(ctx, msgVpnName, optional)
Get a list of Topic Endpoint objects.

Get a list of Topic Endpoint objects.  A Topic Endpoint attracts messages published to a topic for which the Topic Endpoint has a matching topic subscription. The topic subscription for the Topic Endpoint is specified in the client request to bind a Flow to that Topic Endpoint. Queues are significantly more flexible than Topic Endpoints and are the recommended approach for most applications. The use of Topic Endpoints should be restricted to JMS applications.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| topicEndpointName|x| virtualRouter||x    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.12.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MsgVpnApiGetMsgVpnTopicEndpointsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnTopicEndpointsOpts struct
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
 **optional** | ***MsgVpnApiGetMsgVpnTransactionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnTransactionOpts struct
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
 **optional** | ***MsgVpnApiGetMsgVpnTransactionConsumerMsgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnTransactionConsumerMsgOpts struct
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
 **optional** | ***MsgVpnApiGetMsgVpnTransactionConsumerMsgsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnTransactionConsumerMsgsOpts struct
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
 **optional** | ***MsgVpnApiGetMsgVpnTransactionPublisherMsgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnTransactionPublisherMsgOpts struct
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
 **optional** | ***MsgVpnApiGetMsgVpnTransactionPublisherMsgsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnTransactionPublisherMsgsOpts struct
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
 **optional** | ***MsgVpnApiGetMsgVpnTransactionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnTransactionsOpts struct
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

Get a list of Message VPN objects.  Message VPNs (Virtual Private Networks) allow for the segregation of topic space and clients. They also group clients connecting to a network of message brokers, such that messages published within a particular group are only visible to that group's clients.   Attribute|Identifying|Deprecated :---|:---:|:---: authenticationOauthDefaultProviderName||x bridgingTlsServerCertEnforceTrustedCommonNameEnabled||x configSyncLocalKey||x configSyncLocalLastResult||x configSyncLocalRole||x configSyncLocalState||x configSyncLocalTimeInState||x counter.controlRxByteCount||x counter.controlRxMsgCount||x counter.controlTxByteCount||x counter.controlTxMsgCount||x counter.dataRxByteCount||x counter.dataRxMsgCount||x counter.dataTxByteCount||x counter.dataTxMsgCount||x counter.discardedRxMsgCount||x counter.discardedTxMsgCount||x counter.loginRxMsgCount||x counter.loginTxMsgCount||x counter.msgSpoolRxMsgCount||x counter.msgSpoolTxMsgCount||x counter.tlsRxByteCount||x counter.tlsTxByteCount||x distributedCacheManagementEnabled||x msgVpnName|x| rate.averageRxByteRate||x rate.averageRxMsgRate||x rate.averageTxByteRate||x rate.averageTxMsgRate||x rate.rxByteRate||x rate.rxMsgRate||x rate.tlsAverageRxByteRate||x rate.tlsAverageTxByteRate||x rate.tlsRxByteRate||x rate.tlsTxByteRate||x rate.txByteRate||x rate.txMsgRate||x restTlsServerCertEnforceTrustedCommonNameEnabled||x    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MsgVpnApiGetMsgVpnsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnsOpts struct
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

