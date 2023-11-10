# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/config*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMsgVpn**](MsgVpnApi.md#CreateMsgVpn) | **Post** /msgVpns | Create a Message VPN object.
[**CreateMsgVpnAclProfile**](MsgVpnApi.md#CreateMsgVpnAclProfile) | **Post** /msgVpns/{msgVpnName}/aclProfiles | Create an ACL Profile object.
[**CreateMsgVpnAclProfileClientConnectException**](MsgVpnApi.md#CreateMsgVpnAclProfileClientConnectException) | **Post** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/clientConnectExceptions | Create a Client Connect Exception object.
[**CreateMsgVpnAclProfilePublishException**](MsgVpnApi.md#CreateMsgVpnAclProfilePublishException) | **Post** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishExceptions | Create a Publish Topic Exception object.
[**CreateMsgVpnAclProfilePublishTopicException**](MsgVpnApi.md#CreateMsgVpnAclProfilePublishTopicException) | **Post** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishTopicExceptions | Create a Publish Topic Exception object.
[**CreateMsgVpnAclProfileSubscribeException**](MsgVpnApi.md#CreateMsgVpnAclProfileSubscribeException) | **Post** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeExceptions | Create a Subscribe Topic Exception object.
[**CreateMsgVpnAclProfileSubscribeShareNameException**](MsgVpnApi.md#CreateMsgVpnAclProfileSubscribeShareNameException) | **Post** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeShareNameExceptions | Create a Subscribe Share Name Exception object.
[**CreateMsgVpnAclProfileSubscribeTopicException**](MsgVpnApi.md#CreateMsgVpnAclProfileSubscribeTopicException) | **Post** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeTopicExceptions | Create a Subscribe Topic Exception object.
[**CreateMsgVpnAuthenticationOauthProfile**](MsgVpnApi.md#CreateMsgVpnAuthenticationOauthProfile) | **Post** /msgVpns/{msgVpnName}/authenticationOauthProfiles | Create an OAuth Profile object.
[**CreateMsgVpnAuthenticationOauthProfileClientRequiredClaim**](MsgVpnApi.md#CreateMsgVpnAuthenticationOauthProfileClientRequiredClaim) | **Post** /msgVpns/{msgVpnName}/authenticationOauthProfiles/{oauthProfileName}/clientRequiredClaims | Create a Required Claim object.
[**CreateMsgVpnAuthenticationOauthProfileResourceServerRequiredClaim**](MsgVpnApi.md#CreateMsgVpnAuthenticationOauthProfileResourceServerRequiredClaim) | **Post** /msgVpns/{msgVpnName}/authenticationOauthProfiles/{oauthProfileName}/resourceServerRequiredClaims | Create a Required Claim object.
[**CreateMsgVpnAuthenticationOauthProvider**](MsgVpnApi.md#CreateMsgVpnAuthenticationOauthProvider) | **Post** /msgVpns/{msgVpnName}/authenticationOauthProviders | Create an OAuth Provider object.
[**CreateMsgVpnAuthorizationGroup**](MsgVpnApi.md#CreateMsgVpnAuthorizationGroup) | **Post** /msgVpns/{msgVpnName}/authorizationGroups | Create an Authorization Group object.
[**CreateMsgVpnBridge**](MsgVpnApi.md#CreateMsgVpnBridge) | **Post** /msgVpns/{msgVpnName}/bridges | Create a Bridge object.
[**CreateMsgVpnBridgeRemoteMsgVpn**](MsgVpnApi.md#CreateMsgVpnBridgeRemoteMsgVpn) | **Post** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns | Create a Remote Message VPN object.
[**CreateMsgVpnBridgeRemoteSubscription**](MsgVpnApi.md#CreateMsgVpnBridgeRemoteSubscription) | **Post** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteSubscriptions | Create a Remote Subscription object.
[**CreateMsgVpnBridgeTlsTrustedCommonName**](MsgVpnApi.md#CreateMsgVpnBridgeTlsTrustedCommonName) | **Post** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/tlsTrustedCommonNames | Create a Trusted Common Name object.
[**CreateMsgVpnCertMatchingRule**](MsgVpnApi.md#CreateMsgVpnCertMatchingRule) | **Post** /msgVpns/{msgVpnName}/certMatchingRules | Create a Certificate Matching Rule object.
[**CreateMsgVpnCertMatchingRuleAttributeFilter**](MsgVpnApi.md#CreateMsgVpnCertMatchingRuleAttributeFilter) | **Post** /msgVpns/{msgVpnName}/certMatchingRules/{ruleName}/attributeFilters | Create a Certificate Matching Rule Attribute Filter object.
[**CreateMsgVpnCertMatchingRuleCondition**](MsgVpnApi.md#CreateMsgVpnCertMatchingRuleCondition) | **Post** /msgVpns/{msgVpnName}/certMatchingRules/{ruleName}/conditions | Create a Certificate Matching Rule Condition object.
[**CreateMsgVpnClientProfile**](MsgVpnApi.md#CreateMsgVpnClientProfile) | **Post** /msgVpns/{msgVpnName}/clientProfiles | Create a Client Profile object.
[**CreateMsgVpnClientUsername**](MsgVpnApi.md#CreateMsgVpnClientUsername) | **Post** /msgVpns/{msgVpnName}/clientUsernames | Create a Client Username object.
[**CreateMsgVpnClientUsernameAttribute**](MsgVpnApi.md#CreateMsgVpnClientUsernameAttribute) | **Post** /msgVpns/{msgVpnName}/clientUsernames/{clientUsername}/attributes | Create a Client Username Attribute object.
[**CreateMsgVpnDistributedCache**](MsgVpnApi.md#CreateMsgVpnDistributedCache) | **Post** /msgVpns/{msgVpnName}/distributedCaches | Create a Distributed Cache object.
[**CreateMsgVpnDistributedCacheCluster**](MsgVpnApi.md#CreateMsgVpnDistributedCacheCluster) | **Post** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters | Create a Cache Cluster object.
[**CreateMsgVpnDistributedCacheClusterGlobalCachingHomeCluster**](MsgVpnApi.md#CreateMsgVpnDistributedCacheClusterGlobalCachingHomeCluster) | **Post** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters | Create a Home Cache Cluster object.
[**CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix**](MsgVpnApi.md#CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix) | **Post** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName}/topicPrefixes | Create a Topic Prefix object.
[**CreateMsgVpnDistributedCacheClusterInstance**](MsgVpnApi.md#CreateMsgVpnDistributedCacheClusterInstance) | **Post** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances | Create a Cache Instance object.
[**CreateMsgVpnDistributedCacheClusterTopic**](MsgVpnApi.md#CreateMsgVpnDistributedCacheClusterTopic) | **Post** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/topics | Create a Topic object.
[**CreateMsgVpnDmrBridge**](MsgVpnApi.md#CreateMsgVpnDmrBridge) | **Post** /msgVpns/{msgVpnName}/dmrBridges | Create a DMR Bridge object.
[**CreateMsgVpnJndiConnectionFactory**](MsgVpnApi.md#CreateMsgVpnJndiConnectionFactory) | **Post** /msgVpns/{msgVpnName}/jndiConnectionFactories | Create a JNDI Connection Factory object.
[**CreateMsgVpnJndiQueue**](MsgVpnApi.md#CreateMsgVpnJndiQueue) | **Post** /msgVpns/{msgVpnName}/jndiQueues | Create a JNDI Queue object.
[**CreateMsgVpnJndiTopic**](MsgVpnApi.md#CreateMsgVpnJndiTopic) | **Post** /msgVpns/{msgVpnName}/jndiTopics | Create a JNDI Topic object.
[**CreateMsgVpnMqttRetainCache**](MsgVpnApi.md#CreateMsgVpnMqttRetainCache) | **Post** /msgVpns/{msgVpnName}/mqttRetainCaches | Create an MQTT Retain Cache object.
[**CreateMsgVpnMqttSession**](MsgVpnApi.md#CreateMsgVpnMqttSession) | **Post** /msgVpns/{msgVpnName}/mqttSessions | Create an MQTT Session object.
[**CreateMsgVpnMqttSessionSubscription**](MsgVpnApi.md#CreateMsgVpnMqttSessionSubscription) | **Post** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions | Create a Subscription object.
[**CreateMsgVpnQueue**](MsgVpnApi.md#CreateMsgVpnQueue) | **Post** /msgVpns/{msgVpnName}/queues | Create a Queue object.
[**CreateMsgVpnQueueSubscription**](MsgVpnApi.md#CreateMsgVpnQueueSubscription) | **Post** /msgVpns/{msgVpnName}/queues/{queueName}/subscriptions | Create a Queue Subscription object.
[**CreateMsgVpnQueueTemplate**](MsgVpnApi.md#CreateMsgVpnQueueTemplate) | **Post** /msgVpns/{msgVpnName}/queueTemplates | Create a Queue Template object.
[**CreateMsgVpnReplayLog**](MsgVpnApi.md#CreateMsgVpnReplayLog) | **Post** /msgVpns/{msgVpnName}/replayLogs | Create a Replay Log object.
[**CreateMsgVpnReplayLogTopicFilterSubscription**](MsgVpnApi.md#CreateMsgVpnReplayLogTopicFilterSubscription) | **Post** /msgVpns/{msgVpnName}/replayLogs/{replayLogName}/topicFilterSubscriptions | Create a Topic Filter Subscription object.
[**CreateMsgVpnReplicatedTopic**](MsgVpnApi.md#CreateMsgVpnReplicatedTopic) | **Post** /msgVpns/{msgVpnName}/replicatedTopics | Create a Replicated Topic object.
[**CreateMsgVpnRestDeliveryPoint**](MsgVpnApi.md#CreateMsgVpnRestDeliveryPoint) | **Post** /msgVpns/{msgVpnName}/restDeliveryPoints | Create a REST Delivery Point object.
[**CreateMsgVpnRestDeliveryPointQueueBinding**](MsgVpnApi.md#CreateMsgVpnRestDeliveryPointQueueBinding) | **Post** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings | Create a Queue Binding object.
[**CreateMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader**](MsgVpnApi.md#CreateMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader) | **Post** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName}/protectedRequestHeaders | Create a Protected Request Header object.
[**CreateMsgVpnRestDeliveryPointQueueBindingRequestHeader**](MsgVpnApi.md#CreateMsgVpnRestDeliveryPointQueueBindingRequestHeader) | **Post** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName}/requestHeaders | Create a Request Header object.
[**CreateMsgVpnRestDeliveryPointRestConsumer**](MsgVpnApi.md#CreateMsgVpnRestDeliveryPointRestConsumer) | **Post** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers | Create a REST Consumer object.
[**CreateMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim**](MsgVpnApi.md#CreateMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim) | **Post** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/oauthJwtClaims | Create a Claim object.
[**CreateMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName**](MsgVpnApi.md#CreateMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName) | **Post** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/tlsTrustedCommonNames | Create a Trusted Common Name object.
[**CreateMsgVpnSequencedTopic**](MsgVpnApi.md#CreateMsgVpnSequencedTopic) | **Post** /msgVpns/{msgVpnName}/sequencedTopics | Create a Sequenced Topic object.
[**CreateMsgVpnTelemetryProfile**](MsgVpnApi.md#CreateMsgVpnTelemetryProfile) | **Post** /msgVpns/{msgVpnName}/telemetryProfiles | Create a Telemetry Profile object.
[**CreateMsgVpnTelemetryProfileReceiverAclConnectException**](MsgVpnApi.md#CreateMsgVpnTelemetryProfileReceiverAclConnectException) | **Post** /msgVpns/{msgVpnName}/telemetryProfiles/{telemetryProfileName}/receiverAclConnectExceptions | Create a Receiver ACL Connect Exception object.
[**CreateMsgVpnTelemetryProfileTraceFilter**](MsgVpnApi.md#CreateMsgVpnTelemetryProfileTraceFilter) | **Post** /msgVpns/{msgVpnName}/telemetryProfiles/{telemetryProfileName}/traceFilters | Create a Trace Filter object.
[**CreateMsgVpnTelemetryProfileTraceFilterSubscription**](MsgVpnApi.md#CreateMsgVpnTelemetryProfileTraceFilterSubscription) | **Post** /msgVpns/{msgVpnName}/telemetryProfiles/{telemetryProfileName}/traceFilters/{traceFilterName}/subscriptions | Create a Telemetry Trace Filter Subscription object.
[**CreateMsgVpnTopicEndpoint**](MsgVpnApi.md#CreateMsgVpnTopicEndpoint) | **Post** /msgVpns/{msgVpnName}/topicEndpoints | Create a Topic Endpoint object.
[**CreateMsgVpnTopicEndpointTemplate**](MsgVpnApi.md#CreateMsgVpnTopicEndpointTemplate) | **Post** /msgVpns/{msgVpnName}/topicEndpointTemplates | Create a Topic Endpoint Template object.
[**DeleteMsgVpn**](MsgVpnApi.md#DeleteMsgVpn) | **Delete** /msgVpns/{msgVpnName} | Delete a Message VPN object.
[**DeleteMsgVpnAclProfile**](MsgVpnApi.md#DeleteMsgVpnAclProfile) | **Delete** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName} | Delete an ACL Profile object.
[**DeleteMsgVpnAclProfileClientConnectException**](MsgVpnApi.md#DeleteMsgVpnAclProfileClientConnectException) | **Delete** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/clientConnectExceptions/{clientConnectExceptionAddress} | Delete a Client Connect Exception object.
[**DeleteMsgVpnAclProfilePublishException**](MsgVpnApi.md#DeleteMsgVpnAclProfilePublishException) | **Delete** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishExceptions/{topicSyntax},{publishExceptionTopic} | Delete a Publish Topic Exception object.
[**DeleteMsgVpnAclProfilePublishTopicException**](MsgVpnApi.md#DeleteMsgVpnAclProfilePublishTopicException) | **Delete** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishTopicExceptions/{publishTopicExceptionSyntax},{publishTopicException} | Delete a Publish Topic Exception object.
[**DeleteMsgVpnAclProfileSubscribeException**](MsgVpnApi.md#DeleteMsgVpnAclProfileSubscribeException) | **Delete** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeExceptions/{topicSyntax},{subscribeExceptionTopic} | Delete a Subscribe Topic Exception object.
[**DeleteMsgVpnAclProfileSubscribeShareNameException**](MsgVpnApi.md#DeleteMsgVpnAclProfileSubscribeShareNameException) | **Delete** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeShareNameExceptions/{subscribeShareNameExceptionSyntax},{subscribeShareNameException} | Delete a Subscribe Share Name Exception object.
[**DeleteMsgVpnAclProfileSubscribeTopicException**](MsgVpnApi.md#DeleteMsgVpnAclProfileSubscribeTopicException) | **Delete** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeTopicExceptions/{subscribeTopicExceptionSyntax},{subscribeTopicException} | Delete a Subscribe Topic Exception object.
[**DeleteMsgVpnAuthenticationOauthProfile**](MsgVpnApi.md#DeleteMsgVpnAuthenticationOauthProfile) | **Delete** /msgVpns/{msgVpnName}/authenticationOauthProfiles/{oauthProfileName} | Delete an OAuth Profile object.
[**DeleteMsgVpnAuthenticationOauthProfileClientRequiredClaim**](MsgVpnApi.md#DeleteMsgVpnAuthenticationOauthProfileClientRequiredClaim) | **Delete** /msgVpns/{msgVpnName}/authenticationOauthProfiles/{oauthProfileName}/clientRequiredClaims/{clientRequiredClaimName} | Delete a Required Claim object.
[**DeleteMsgVpnAuthenticationOauthProfileResourceServerRequiredClaim**](MsgVpnApi.md#DeleteMsgVpnAuthenticationOauthProfileResourceServerRequiredClaim) | **Delete** /msgVpns/{msgVpnName}/authenticationOauthProfiles/{oauthProfileName}/resourceServerRequiredClaims/{resourceServerRequiredClaimName} | Delete a Required Claim object.
[**DeleteMsgVpnAuthenticationOauthProvider**](MsgVpnApi.md#DeleteMsgVpnAuthenticationOauthProvider) | **Delete** /msgVpns/{msgVpnName}/authenticationOauthProviders/{oauthProviderName} | Delete an OAuth Provider object.
[**DeleteMsgVpnAuthorizationGroup**](MsgVpnApi.md#DeleteMsgVpnAuthorizationGroup) | **Delete** /msgVpns/{msgVpnName}/authorizationGroups/{authorizationGroupName} | Delete an Authorization Group object.
[**DeleteMsgVpnBridge**](MsgVpnApi.md#DeleteMsgVpnBridge) | **Delete** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter} | Delete a Bridge object.
[**DeleteMsgVpnBridgeRemoteMsgVpn**](MsgVpnApi.md#DeleteMsgVpnBridgeRemoteMsgVpn) | **Delete** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns/{remoteMsgVpnName},{remoteMsgVpnLocation},{remoteMsgVpnInterface} | Delete a Remote Message VPN object.
[**DeleteMsgVpnBridgeRemoteSubscription**](MsgVpnApi.md#DeleteMsgVpnBridgeRemoteSubscription) | **Delete** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteSubscriptions/{remoteSubscriptionTopic} | Delete a Remote Subscription object.
[**DeleteMsgVpnBridgeTlsTrustedCommonName**](MsgVpnApi.md#DeleteMsgVpnBridgeTlsTrustedCommonName) | **Delete** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/tlsTrustedCommonNames/{tlsTrustedCommonName} | Delete a Trusted Common Name object.
[**DeleteMsgVpnCertMatchingRule**](MsgVpnApi.md#DeleteMsgVpnCertMatchingRule) | **Delete** /msgVpns/{msgVpnName}/certMatchingRules/{ruleName} | Delete a Certificate Matching Rule object.
[**DeleteMsgVpnCertMatchingRuleAttributeFilter**](MsgVpnApi.md#DeleteMsgVpnCertMatchingRuleAttributeFilter) | **Delete** /msgVpns/{msgVpnName}/certMatchingRules/{ruleName}/attributeFilters/{filterName} | Delete a Certificate Matching Rule Attribute Filter object.
[**DeleteMsgVpnCertMatchingRuleCondition**](MsgVpnApi.md#DeleteMsgVpnCertMatchingRuleCondition) | **Delete** /msgVpns/{msgVpnName}/certMatchingRules/{ruleName}/conditions/{source} | Delete a Certificate Matching Rule Condition object.
[**DeleteMsgVpnClientProfile**](MsgVpnApi.md#DeleteMsgVpnClientProfile) | **Delete** /msgVpns/{msgVpnName}/clientProfiles/{clientProfileName} | Delete a Client Profile object.
[**DeleteMsgVpnClientUsername**](MsgVpnApi.md#DeleteMsgVpnClientUsername) | **Delete** /msgVpns/{msgVpnName}/clientUsernames/{clientUsername} | Delete a Client Username object.
[**DeleteMsgVpnClientUsernameAttribute**](MsgVpnApi.md#DeleteMsgVpnClientUsernameAttribute) | **Delete** /msgVpns/{msgVpnName}/clientUsernames/{clientUsername}/attributes/{attributeName},{attributeValue} | Delete a Client Username Attribute object.
[**DeleteMsgVpnDistributedCache**](MsgVpnApi.md#DeleteMsgVpnDistributedCache) | **Delete** /msgVpns/{msgVpnName}/distributedCaches/{cacheName} | Delete a Distributed Cache object.
[**DeleteMsgVpnDistributedCacheCluster**](MsgVpnApi.md#DeleteMsgVpnDistributedCacheCluster) | **Delete** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName} | Delete a Cache Cluster object.
[**DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeCluster**](MsgVpnApi.md#DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeCluster) | **Delete** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName} | Delete a Home Cache Cluster object.
[**DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix**](MsgVpnApi.md#DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix) | **Delete** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName}/topicPrefixes/{topicPrefix} | Delete a Topic Prefix object.
[**DeleteMsgVpnDistributedCacheClusterInstance**](MsgVpnApi.md#DeleteMsgVpnDistributedCacheClusterInstance) | **Delete** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName} | Delete a Cache Instance object.
[**DeleteMsgVpnDistributedCacheClusterTopic**](MsgVpnApi.md#DeleteMsgVpnDistributedCacheClusterTopic) | **Delete** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/topics/{topic} | Delete a Topic object.
[**DeleteMsgVpnDmrBridge**](MsgVpnApi.md#DeleteMsgVpnDmrBridge) | **Delete** /msgVpns/{msgVpnName}/dmrBridges/{remoteNodeName} | Delete a DMR Bridge object.
[**DeleteMsgVpnJndiConnectionFactory**](MsgVpnApi.md#DeleteMsgVpnJndiConnectionFactory) | **Delete** /msgVpns/{msgVpnName}/jndiConnectionFactories/{connectionFactoryName} | Delete a JNDI Connection Factory object.
[**DeleteMsgVpnJndiQueue**](MsgVpnApi.md#DeleteMsgVpnJndiQueue) | **Delete** /msgVpns/{msgVpnName}/jndiQueues/{queueName} | Delete a JNDI Queue object.
[**DeleteMsgVpnJndiTopic**](MsgVpnApi.md#DeleteMsgVpnJndiTopic) | **Delete** /msgVpns/{msgVpnName}/jndiTopics/{topicName} | Delete a JNDI Topic object.
[**DeleteMsgVpnMqttRetainCache**](MsgVpnApi.md#DeleteMsgVpnMqttRetainCache) | **Delete** /msgVpns/{msgVpnName}/mqttRetainCaches/{cacheName} | Delete an MQTT Retain Cache object.
[**DeleteMsgVpnMqttSession**](MsgVpnApi.md#DeleteMsgVpnMqttSession) | **Delete** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter} | Delete an MQTT Session object.
[**DeleteMsgVpnMqttSessionSubscription**](MsgVpnApi.md#DeleteMsgVpnMqttSessionSubscription) | **Delete** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions/{subscriptionTopic} | Delete a Subscription object.
[**DeleteMsgVpnQueue**](MsgVpnApi.md#DeleteMsgVpnQueue) | **Delete** /msgVpns/{msgVpnName}/queues/{queueName} | Delete a Queue object.
[**DeleteMsgVpnQueueSubscription**](MsgVpnApi.md#DeleteMsgVpnQueueSubscription) | **Delete** /msgVpns/{msgVpnName}/queues/{queueName}/subscriptions/{subscriptionTopic} | Delete a Queue Subscription object.
[**DeleteMsgVpnQueueTemplate**](MsgVpnApi.md#DeleteMsgVpnQueueTemplate) | **Delete** /msgVpns/{msgVpnName}/queueTemplates/{queueTemplateName} | Delete a Queue Template object.
[**DeleteMsgVpnReplayLog**](MsgVpnApi.md#DeleteMsgVpnReplayLog) | **Delete** /msgVpns/{msgVpnName}/replayLogs/{replayLogName} | Delete a Replay Log object.
[**DeleteMsgVpnReplayLogTopicFilterSubscription**](MsgVpnApi.md#DeleteMsgVpnReplayLogTopicFilterSubscription) | **Delete** /msgVpns/{msgVpnName}/replayLogs/{replayLogName}/topicFilterSubscriptions/{topicFilterSubscription} | Delete a Topic Filter Subscription object.
[**DeleteMsgVpnReplicatedTopic**](MsgVpnApi.md#DeleteMsgVpnReplicatedTopic) | **Delete** /msgVpns/{msgVpnName}/replicatedTopics/{replicatedTopic} | Delete a Replicated Topic object.
[**DeleteMsgVpnRestDeliveryPoint**](MsgVpnApi.md#DeleteMsgVpnRestDeliveryPoint) | **Delete** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName} | Delete a REST Delivery Point object.
[**DeleteMsgVpnRestDeliveryPointQueueBinding**](MsgVpnApi.md#DeleteMsgVpnRestDeliveryPointQueueBinding) | **Delete** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName} | Delete a Queue Binding object.
[**DeleteMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader**](MsgVpnApi.md#DeleteMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader) | **Delete** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName}/protectedRequestHeaders/{headerName} | Delete a Protected Request Header object.
[**DeleteMsgVpnRestDeliveryPointQueueBindingRequestHeader**](MsgVpnApi.md#DeleteMsgVpnRestDeliveryPointQueueBindingRequestHeader) | **Delete** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName}/requestHeaders/{headerName} | Delete a Request Header object.
[**DeleteMsgVpnRestDeliveryPointRestConsumer**](MsgVpnApi.md#DeleteMsgVpnRestDeliveryPointRestConsumer) | **Delete** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName} | Delete a REST Consumer object.
[**DeleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim**](MsgVpnApi.md#DeleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim) | **Delete** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/oauthJwtClaims/{oauthJwtClaimName} | Delete a Claim object.
[**DeleteMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName**](MsgVpnApi.md#DeleteMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName) | **Delete** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/tlsTrustedCommonNames/{tlsTrustedCommonName} | Delete a Trusted Common Name object.
[**DeleteMsgVpnSequencedTopic**](MsgVpnApi.md#DeleteMsgVpnSequencedTopic) | **Delete** /msgVpns/{msgVpnName}/sequencedTopics/{sequencedTopic} | Delete a Sequenced Topic object.
[**DeleteMsgVpnTelemetryProfile**](MsgVpnApi.md#DeleteMsgVpnTelemetryProfile) | **Delete** /msgVpns/{msgVpnName}/telemetryProfiles/{telemetryProfileName} | Delete a Telemetry Profile object.
[**DeleteMsgVpnTelemetryProfileReceiverAclConnectException**](MsgVpnApi.md#DeleteMsgVpnTelemetryProfileReceiverAclConnectException) | **Delete** /msgVpns/{msgVpnName}/telemetryProfiles/{telemetryProfileName}/receiverAclConnectExceptions/{receiverAclConnectExceptionAddress} | Delete a Receiver ACL Connect Exception object.
[**DeleteMsgVpnTelemetryProfileTraceFilter**](MsgVpnApi.md#DeleteMsgVpnTelemetryProfileTraceFilter) | **Delete** /msgVpns/{msgVpnName}/telemetryProfiles/{telemetryProfileName}/traceFilters/{traceFilterName} | Delete a Trace Filter object.
[**DeleteMsgVpnTelemetryProfileTraceFilterSubscription**](MsgVpnApi.md#DeleteMsgVpnTelemetryProfileTraceFilterSubscription) | **Delete** /msgVpns/{msgVpnName}/telemetryProfiles/{telemetryProfileName}/traceFilters/{traceFilterName}/subscriptions/{subscription},{subscriptionSyntax} | Delete a Telemetry Trace Filter Subscription object.
[**DeleteMsgVpnTopicEndpoint**](MsgVpnApi.md#DeleteMsgVpnTopicEndpoint) | **Delete** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName} | Delete a Topic Endpoint object.
[**DeleteMsgVpnTopicEndpointTemplate**](MsgVpnApi.md#DeleteMsgVpnTopicEndpointTemplate) | **Delete** /msgVpns/{msgVpnName}/topicEndpointTemplates/{topicEndpointTemplateName} | Delete a Topic Endpoint Template object.
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
[**GetMsgVpnClientProfile**](MsgVpnApi.md#GetMsgVpnClientProfile) | **Get** /msgVpns/{msgVpnName}/clientProfiles/{clientProfileName} | Get a Client Profile object.
[**GetMsgVpnClientProfiles**](MsgVpnApi.md#GetMsgVpnClientProfiles) | **Get** /msgVpns/{msgVpnName}/clientProfiles | Get a list of Client Profile objects.
[**GetMsgVpnClientUsername**](MsgVpnApi.md#GetMsgVpnClientUsername) | **Get** /msgVpns/{msgVpnName}/clientUsernames/{clientUsername} | Get a Client Username object.
[**GetMsgVpnClientUsernameAttribute**](MsgVpnApi.md#GetMsgVpnClientUsernameAttribute) | **Get** /msgVpns/{msgVpnName}/clientUsernames/{clientUsername}/attributes/{attributeName},{attributeValue} | Get a Client Username Attribute object.
[**GetMsgVpnClientUsernameAttributes**](MsgVpnApi.md#GetMsgVpnClientUsernameAttributes) | **Get** /msgVpns/{msgVpnName}/clientUsernames/{clientUsername}/attributes | Get a list of Client Username Attribute objects.
[**GetMsgVpnClientUsernames**](MsgVpnApi.md#GetMsgVpnClientUsernames) | **Get** /msgVpns/{msgVpnName}/clientUsernames | Get a list of Client Username objects.
[**GetMsgVpnDistributedCache**](MsgVpnApi.md#GetMsgVpnDistributedCache) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName} | Get a Distributed Cache object.
[**GetMsgVpnDistributedCacheCluster**](MsgVpnApi.md#GetMsgVpnDistributedCacheCluster) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName} | Get a Cache Cluster object.
[**GetMsgVpnDistributedCacheClusterGlobalCachingHomeCluster**](MsgVpnApi.md#GetMsgVpnDistributedCacheClusterGlobalCachingHomeCluster) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName} | Get a Home Cache Cluster object.
[**GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix**](MsgVpnApi.md#GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName}/topicPrefixes/{topicPrefix} | Get a Topic Prefix object.
[**GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixes**](MsgVpnApi.md#GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixes) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName}/topicPrefixes | Get a list of Topic Prefix objects.
[**GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusters**](MsgVpnApi.md#GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusters) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters | Get a list of Home Cache Cluster objects.
[**GetMsgVpnDistributedCacheClusterInstance**](MsgVpnApi.md#GetMsgVpnDistributedCacheClusterInstance) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName} | Get a Cache Instance object.
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
[**GetMsgVpnQueueSubscription**](MsgVpnApi.md#GetMsgVpnQueueSubscription) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/subscriptions/{subscriptionTopic} | Get a Queue Subscription object.
[**GetMsgVpnQueueSubscriptions**](MsgVpnApi.md#GetMsgVpnQueueSubscriptions) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/subscriptions | Get a list of Queue Subscription objects.
[**GetMsgVpnQueueTemplate**](MsgVpnApi.md#GetMsgVpnQueueTemplate) | **Get** /msgVpns/{msgVpnName}/queueTemplates/{queueTemplateName} | Get a Queue Template object.
[**GetMsgVpnQueueTemplates**](MsgVpnApi.md#GetMsgVpnQueueTemplates) | **Get** /msgVpns/{msgVpnName}/queueTemplates | Get a list of Queue Template objects.
[**GetMsgVpnQueues**](MsgVpnApi.md#GetMsgVpnQueues) | **Get** /msgVpns/{msgVpnName}/queues | Get a list of Queue objects.
[**GetMsgVpnReplayLog**](MsgVpnApi.md#GetMsgVpnReplayLog) | **Get** /msgVpns/{msgVpnName}/replayLogs/{replayLogName} | Get a Replay Log object.
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
[**GetMsgVpnSequencedTopic**](MsgVpnApi.md#GetMsgVpnSequencedTopic) | **Get** /msgVpns/{msgVpnName}/sequencedTopics/{sequencedTopic} | Get a Sequenced Topic object.
[**GetMsgVpnSequencedTopics**](MsgVpnApi.md#GetMsgVpnSequencedTopics) | **Get** /msgVpns/{msgVpnName}/sequencedTopics | Get a list of Sequenced Topic objects.
[**GetMsgVpnTelemetryProfile**](MsgVpnApi.md#GetMsgVpnTelemetryProfile) | **Get** /msgVpns/{msgVpnName}/telemetryProfiles/{telemetryProfileName} | Get a Telemetry Profile object.
[**GetMsgVpnTelemetryProfileReceiverAclConnectException**](MsgVpnApi.md#GetMsgVpnTelemetryProfileReceiverAclConnectException) | **Get** /msgVpns/{msgVpnName}/telemetryProfiles/{telemetryProfileName}/receiverAclConnectExceptions/{receiverAclConnectExceptionAddress} | Get a Receiver ACL Connect Exception object.
[**GetMsgVpnTelemetryProfileReceiverAclConnectExceptions**](MsgVpnApi.md#GetMsgVpnTelemetryProfileReceiverAclConnectExceptions) | **Get** /msgVpns/{msgVpnName}/telemetryProfiles/{telemetryProfileName}/receiverAclConnectExceptions | Get a list of Receiver ACL Connect Exception objects.
[**GetMsgVpnTelemetryProfileTraceFilter**](MsgVpnApi.md#GetMsgVpnTelemetryProfileTraceFilter) | **Get** /msgVpns/{msgVpnName}/telemetryProfiles/{telemetryProfileName}/traceFilters/{traceFilterName} | Get a Trace Filter object.
[**GetMsgVpnTelemetryProfileTraceFilterSubscription**](MsgVpnApi.md#GetMsgVpnTelemetryProfileTraceFilterSubscription) | **Get** /msgVpns/{msgVpnName}/telemetryProfiles/{telemetryProfileName}/traceFilters/{traceFilterName}/subscriptions/{subscription},{subscriptionSyntax} | Get a Telemetry Trace Filter Subscription object.
[**GetMsgVpnTelemetryProfileTraceFilterSubscriptions**](MsgVpnApi.md#GetMsgVpnTelemetryProfileTraceFilterSubscriptions) | **Get** /msgVpns/{msgVpnName}/telemetryProfiles/{telemetryProfileName}/traceFilters/{traceFilterName}/subscriptions | Get a list of Telemetry Trace Filter Subscription objects.
[**GetMsgVpnTelemetryProfileTraceFilters**](MsgVpnApi.md#GetMsgVpnTelemetryProfileTraceFilters) | **Get** /msgVpns/{msgVpnName}/telemetryProfiles/{telemetryProfileName}/traceFilters | Get a list of Trace Filter objects.
[**GetMsgVpnTelemetryProfiles**](MsgVpnApi.md#GetMsgVpnTelemetryProfiles) | **Get** /msgVpns/{msgVpnName}/telemetryProfiles | Get a list of Telemetry Profile objects.
[**GetMsgVpnTopicEndpoint**](MsgVpnApi.md#GetMsgVpnTopicEndpoint) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName} | Get a Topic Endpoint object.
[**GetMsgVpnTopicEndpointTemplate**](MsgVpnApi.md#GetMsgVpnTopicEndpointTemplate) | **Get** /msgVpns/{msgVpnName}/topicEndpointTemplates/{topicEndpointTemplateName} | Get a Topic Endpoint Template object.
[**GetMsgVpnTopicEndpointTemplates**](MsgVpnApi.md#GetMsgVpnTopicEndpointTemplates) | **Get** /msgVpns/{msgVpnName}/topicEndpointTemplates | Get a list of Topic Endpoint Template objects.
[**GetMsgVpnTopicEndpoints**](MsgVpnApi.md#GetMsgVpnTopicEndpoints) | **Get** /msgVpns/{msgVpnName}/topicEndpoints | Get a list of Topic Endpoint objects.
[**GetMsgVpns**](MsgVpnApi.md#GetMsgVpns) | **Get** /msgVpns | Get a list of Message VPN objects.
[**ReplaceMsgVpn**](MsgVpnApi.md#ReplaceMsgVpn) | **Put** /msgVpns/{msgVpnName} | Replace a Message VPN object.
[**ReplaceMsgVpnAclProfile**](MsgVpnApi.md#ReplaceMsgVpnAclProfile) | **Put** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName} | Replace an ACL Profile object.
[**ReplaceMsgVpnAuthenticationOauthProfile**](MsgVpnApi.md#ReplaceMsgVpnAuthenticationOauthProfile) | **Put** /msgVpns/{msgVpnName}/authenticationOauthProfiles/{oauthProfileName} | Replace an OAuth Profile object.
[**ReplaceMsgVpnAuthenticationOauthProvider**](MsgVpnApi.md#ReplaceMsgVpnAuthenticationOauthProvider) | **Put** /msgVpns/{msgVpnName}/authenticationOauthProviders/{oauthProviderName} | Replace an OAuth Provider object.
[**ReplaceMsgVpnAuthorizationGroup**](MsgVpnApi.md#ReplaceMsgVpnAuthorizationGroup) | **Put** /msgVpns/{msgVpnName}/authorizationGroups/{authorizationGroupName} | Replace an Authorization Group object.
[**ReplaceMsgVpnBridge**](MsgVpnApi.md#ReplaceMsgVpnBridge) | **Put** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter} | Replace a Bridge object.
[**ReplaceMsgVpnBridgeRemoteMsgVpn**](MsgVpnApi.md#ReplaceMsgVpnBridgeRemoteMsgVpn) | **Put** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns/{remoteMsgVpnName},{remoteMsgVpnLocation},{remoteMsgVpnInterface} | Replace a Remote Message VPN object.
[**ReplaceMsgVpnCertMatchingRule**](MsgVpnApi.md#ReplaceMsgVpnCertMatchingRule) | **Put** /msgVpns/{msgVpnName}/certMatchingRules/{ruleName} | Replace a Certificate Matching Rule object.
[**ReplaceMsgVpnCertMatchingRuleAttributeFilter**](MsgVpnApi.md#ReplaceMsgVpnCertMatchingRuleAttributeFilter) | **Put** /msgVpns/{msgVpnName}/certMatchingRules/{ruleName}/attributeFilters/{filterName} | Replace a Certificate Matching Rule Attribute Filter object.
[**ReplaceMsgVpnClientProfile**](MsgVpnApi.md#ReplaceMsgVpnClientProfile) | **Put** /msgVpns/{msgVpnName}/clientProfiles/{clientProfileName} | Replace a Client Profile object.
[**ReplaceMsgVpnClientUsername**](MsgVpnApi.md#ReplaceMsgVpnClientUsername) | **Put** /msgVpns/{msgVpnName}/clientUsernames/{clientUsername} | Replace a Client Username object.
[**ReplaceMsgVpnDistributedCache**](MsgVpnApi.md#ReplaceMsgVpnDistributedCache) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName} | Replace a Distributed Cache object.
[**ReplaceMsgVpnDistributedCacheCluster**](MsgVpnApi.md#ReplaceMsgVpnDistributedCacheCluster) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName} | Replace a Cache Cluster object.
[**ReplaceMsgVpnDistributedCacheClusterInstance**](MsgVpnApi.md#ReplaceMsgVpnDistributedCacheClusterInstance) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName} | Replace a Cache Instance object.
[**ReplaceMsgVpnDmrBridge**](MsgVpnApi.md#ReplaceMsgVpnDmrBridge) | **Put** /msgVpns/{msgVpnName}/dmrBridges/{remoteNodeName} | Replace a DMR Bridge object.
[**ReplaceMsgVpnJndiConnectionFactory**](MsgVpnApi.md#ReplaceMsgVpnJndiConnectionFactory) | **Put** /msgVpns/{msgVpnName}/jndiConnectionFactories/{connectionFactoryName} | Replace a JNDI Connection Factory object.
[**ReplaceMsgVpnJndiQueue**](MsgVpnApi.md#ReplaceMsgVpnJndiQueue) | **Put** /msgVpns/{msgVpnName}/jndiQueues/{queueName} | Replace a JNDI Queue object.
[**ReplaceMsgVpnJndiTopic**](MsgVpnApi.md#ReplaceMsgVpnJndiTopic) | **Put** /msgVpns/{msgVpnName}/jndiTopics/{topicName} | Replace a JNDI Topic object.
[**ReplaceMsgVpnMqttRetainCache**](MsgVpnApi.md#ReplaceMsgVpnMqttRetainCache) | **Put** /msgVpns/{msgVpnName}/mqttRetainCaches/{cacheName} | Replace an MQTT Retain Cache object.
[**ReplaceMsgVpnMqttSession**](MsgVpnApi.md#ReplaceMsgVpnMqttSession) | **Put** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter} | Replace an MQTT Session object.
[**ReplaceMsgVpnMqttSessionSubscription**](MsgVpnApi.md#ReplaceMsgVpnMqttSessionSubscription) | **Put** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions/{subscriptionTopic} | Replace a Subscription object.
[**ReplaceMsgVpnQueue**](MsgVpnApi.md#ReplaceMsgVpnQueue) | **Put** /msgVpns/{msgVpnName}/queues/{queueName} | Replace a Queue object.
[**ReplaceMsgVpnQueueTemplate**](MsgVpnApi.md#ReplaceMsgVpnQueueTemplate) | **Put** /msgVpns/{msgVpnName}/queueTemplates/{queueTemplateName} | Replace a Queue Template object.
[**ReplaceMsgVpnReplayLog**](MsgVpnApi.md#ReplaceMsgVpnReplayLog) | **Put** /msgVpns/{msgVpnName}/replayLogs/{replayLogName} | Replace a Replay Log object.
[**ReplaceMsgVpnReplicatedTopic**](MsgVpnApi.md#ReplaceMsgVpnReplicatedTopic) | **Put** /msgVpns/{msgVpnName}/replicatedTopics/{replicatedTopic} | Replace a Replicated Topic object.
[**ReplaceMsgVpnRestDeliveryPoint**](MsgVpnApi.md#ReplaceMsgVpnRestDeliveryPoint) | **Put** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName} | Replace a REST Delivery Point object.
[**ReplaceMsgVpnRestDeliveryPointQueueBinding**](MsgVpnApi.md#ReplaceMsgVpnRestDeliveryPointQueueBinding) | **Put** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName} | Replace a Queue Binding object.
[**ReplaceMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader**](MsgVpnApi.md#ReplaceMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader) | **Put** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName}/protectedRequestHeaders/{headerName} | Replace a Protected Request Header object.
[**ReplaceMsgVpnRestDeliveryPointQueueBindingRequestHeader**](MsgVpnApi.md#ReplaceMsgVpnRestDeliveryPointQueueBindingRequestHeader) | **Put** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName}/requestHeaders/{headerName} | Replace a Request Header object.
[**ReplaceMsgVpnRestDeliveryPointRestConsumer**](MsgVpnApi.md#ReplaceMsgVpnRestDeliveryPointRestConsumer) | **Put** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName} | Replace a REST Consumer object.
[**ReplaceMsgVpnTelemetryProfile**](MsgVpnApi.md#ReplaceMsgVpnTelemetryProfile) | **Put** /msgVpns/{msgVpnName}/telemetryProfiles/{telemetryProfileName} | Replace a Telemetry Profile object.
[**ReplaceMsgVpnTelemetryProfileTraceFilter**](MsgVpnApi.md#ReplaceMsgVpnTelemetryProfileTraceFilter) | **Put** /msgVpns/{msgVpnName}/telemetryProfiles/{telemetryProfileName}/traceFilters/{traceFilterName} | Replace a Trace Filter object.
[**ReplaceMsgVpnTopicEndpoint**](MsgVpnApi.md#ReplaceMsgVpnTopicEndpoint) | **Put** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName} | Replace a Topic Endpoint object.
[**ReplaceMsgVpnTopicEndpointTemplate**](MsgVpnApi.md#ReplaceMsgVpnTopicEndpointTemplate) | **Put** /msgVpns/{msgVpnName}/topicEndpointTemplates/{topicEndpointTemplateName} | Replace a Topic Endpoint Template object.
[**UpdateMsgVpn**](MsgVpnApi.md#UpdateMsgVpn) | **Patch** /msgVpns/{msgVpnName} | Update a Message VPN object.
[**UpdateMsgVpnAclProfile**](MsgVpnApi.md#UpdateMsgVpnAclProfile) | **Patch** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName} | Update an ACL Profile object.
[**UpdateMsgVpnAuthenticationOauthProfile**](MsgVpnApi.md#UpdateMsgVpnAuthenticationOauthProfile) | **Patch** /msgVpns/{msgVpnName}/authenticationOauthProfiles/{oauthProfileName} | Update an OAuth Profile object.
[**UpdateMsgVpnAuthenticationOauthProvider**](MsgVpnApi.md#UpdateMsgVpnAuthenticationOauthProvider) | **Patch** /msgVpns/{msgVpnName}/authenticationOauthProviders/{oauthProviderName} | Update an OAuth Provider object.
[**UpdateMsgVpnAuthorizationGroup**](MsgVpnApi.md#UpdateMsgVpnAuthorizationGroup) | **Patch** /msgVpns/{msgVpnName}/authorizationGroups/{authorizationGroupName} | Update an Authorization Group object.
[**UpdateMsgVpnBridge**](MsgVpnApi.md#UpdateMsgVpnBridge) | **Patch** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter} | Update a Bridge object.
[**UpdateMsgVpnBridgeRemoteMsgVpn**](MsgVpnApi.md#UpdateMsgVpnBridgeRemoteMsgVpn) | **Patch** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns/{remoteMsgVpnName},{remoteMsgVpnLocation},{remoteMsgVpnInterface} | Update a Remote Message VPN object.
[**UpdateMsgVpnCertMatchingRule**](MsgVpnApi.md#UpdateMsgVpnCertMatchingRule) | **Patch** /msgVpns/{msgVpnName}/certMatchingRules/{ruleName} | Update a Certificate Matching Rule object.
[**UpdateMsgVpnCertMatchingRuleAttributeFilter**](MsgVpnApi.md#UpdateMsgVpnCertMatchingRuleAttributeFilter) | **Patch** /msgVpns/{msgVpnName}/certMatchingRules/{ruleName}/attributeFilters/{filterName} | Update a Certificate Matching Rule Attribute Filter object.
[**UpdateMsgVpnClientProfile**](MsgVpnApi.md#UpdateMsgVpnClientProfile) | **Patch** /msgVpns/{msgVpnName}/clientProfiles/{clientProfileName} | Update a Client Profile object.
[**UpdateMsgVpnClientUsername**](MsgVpnApi.md#UpdateMsgVpnClientUsername) | **Patch** /msgVpns/{msgVpnName}/clientUsernames/{clientUsername} | Update a Client Username object.
[**UpdateMsgVpnDistributedCache**](MsgVpnApi.md#UpdateMsgVpnDistributedCache) | **Patch** /msgVpns/{msgVpnName}/distributedCaches/{cacheName} | Update a Distributed Cache object.
[**UpdateMsgVpnDistributedCacheCluster**](MsgVpnApi.md#UpdateMsgVpnDistributedCacheCluster) | **Patch** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName} | Update a Cache Cluster object.
[**UpdateMsgVpnDistributedCacheClusterInstance**](MsgVpnApi.md#UpdateMsgVpnDistributedCacheClusterInstance) | **Patch** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName} | Update a Cache Instance object.
[**UpdateMsgVpnDmrBridge**](MsgVpnApi.md#UpdateMsgVpnDmrBridge) | **Patch** /msgVpns/{msgVpnName}/dmrBridges/{remoteNodeName} | Update a DMR Bridge object.
[**UpdateMsgVpnJndiConnectionFactory**](MsgVpnApi.md#UpdateMsgVpnJndiConnectionFactory) | **Patch** /msgVpns/{msgVpnName}/jndiConnectionFactories/{connectionFactoryName} | Update a JNDI Connection Factory object.
[**UpdateMsgVpnJndiQueue**](MsgVpnApi.md#UpdateMsgVpnJndiQueue) | **Patch** /msgVpns/{msgVpnName}/jndiQueues/{queueName} | Update a JNDI Queue object.
[**UpdateMsgVpnJndiTopic**](MsgVpnApi.md#UpdateMsgVpnJndiTopic) | **Patch** /msgVpns/{msgVpnName}/jndiTopics/{topicName} | Update a JNDI Topic object.
[**UpdateMsgVpnMqttRetainCache**](MsgVpnApi.md#UpdateMsgVpnMqttRetainCache) | **Patch** /msgVpns/{msgVpnName}/mqttRetainCaches/{cacheName} | Update an MQTT Retain Cache object.
[**UpdateMsgVpnMqttSession**](MsgVpnApi.md#UpdateMsgVpnMqttSession) | **Patch** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter} | Update an MQTT Session object.
[**UpdateMsgVpnMqttSessionSubscription**](MsgVpnApi.md#UpdateMsgVpnMqttSessionSubscription) | **Patch** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions/{subscriptionTopic} | Update a Subscription object.
[**UpdateMsgVpnQueue**](MsgVpnApi.md#UpdateMsgVpnQueue) | **Patch** /msgVpns/{msgVpnName}/queues/{queueName} | Update a Queue object.
[**UpdateMsgVpnQueueTemplate**](MsgVpnApi.md#UpdateMsgVpnQueueTemplate) | **Patch** /msgVpns/{msgVpnName}/queueTemplates/{queueTemplateName} | Update a Queue Template object.
[**UpdateMsgVpnReplayLog**](MsgVpnApi.md#UpdateMsgVpnReplayLog) | **Patch** /msgVpns/{msgVpnName}/replayLogs/{replayLogName} | Update a Replay Log object.
[**UpdateMsgVpnReplicatedTopic**](MsgVpnApi.md#UpdateMsgVpnReplicatedTopic) | **Patch** /msgVpns/{msgVpnName}/replicatedTopics/{replicatedTopic} | Update a Replicated Topic object.
[**UpdateMsgVpnRestDeliveryPoint**](MsgVpnApi.md#UpdateMsgVpnRestDeliveryPoint) | **Patch** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName} | Update a REST Delivery Point object.
[**UpdateMsgVpnRestDeliveryPointQueueBinding**](MsgVpnApi.md#UpdateMsgVpnRestDeliveryPointQueueBinding) | **Patch** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName} | Update a Queue Binding object.
[**UpdateMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader**](MsgVpnApi.md#UpdateMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader) | **Patch** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName}/protectedRequestHeaders/{headerName} | Update a Protected Request Header object.
[**UpdateMsgVpnRestDeliveryPointQueueBindingRequestHeader**](MsgVpnApi.md#UpdateMsgVpnRestDeliveryPointQueueBindingRequestHeader) | **Patch** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName}/requestHeaders/{headerName} | Update a Request Header object.
[**UpdateMsgVpnRestDeliveryPointRestConsumer**](MsgVpnApi.md#UpdateMsgVpnRestDeliveryPointRestConsumer) | **Patch** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName} | Update a REST Consumer object.
[**UpdateMsgVpnTelemetryProfile**](MsgVpnApi.md#UpdateMsgVpnTelemetryProfile) | **Patch** /msgVpns/{msgVpnName}/telemetryProfiles/{telemetryProfileName} | Update a Telemetry Profile object.
[**UpdateMsgVpnTelemetryProfileTraceFilter**](MsgVpnApi.md#UpdateMsgVpnTelemetryProfileTraceFilter) | **Patch** /msgVpns/{msgVpnName}/telemetryProfiles/{telemetryProfileName}/traceFilters/{traceFilterName} | Update a Trace Filter object.
[**UpdateMsgVpnTopicEndpoint**](MsgVpnApi.md#UpdateMsgVpnTopicEndpoint) | **Patch** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName} | Update a Topic Endpoint object.
[**UpdateMsgVpnTopicEndpointTemplate**](MsgVpnApi.md#UpdateMsgVpnTopicEndpointTemplate) | **Patch** /msgVpns/{msgVpnName}/topicEndpointTemplates/{topicEndpointTemplateName} | Update a Topic Endpoint Template object.

# **CreateMsgVpn**
> MsgVpnResponse CreateMsgVpn(ctx, body, optional)
Create a Message VPN object.

Create a Message VPN object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates via config-sync.  Message VPNs (Virtual Private Networks) allow for the segregation of topic space and clients. They also group clients connecting to a network of message brokers, such that messages published within a particular group are only visible to that group's clients.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: authenticationOauthDefaultProviderName|||||x| bridgingTlsServerCertEnforceTrustedCommonNameEnabled|||||x| distributedCacheManagementEnabled|||||x| msgVpnName|x|x|||| replicationBridgeAuthenticationBasicPassword||||x||x replicationBridgeAuthenticationClientCertContent||||x||x replicationBridgeAuthenticationClientCertPassword||||x|| replicationEnabledQueueBehavior||||x|| restTlsServerCertEnforceTrustedCommonNameEnabled|||||x|    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThreshold|clearPercent|setPercent|clearValue, setValue EventThreshold|clearValue|setValue|clearPercent, setPercent EventThreshold|setPercent|clearPercent|clearValue, setValue EventThreshold|setValue|clearValue|clearPercent, setPercent EventThresholdByValue|clearValue|setValue| EventThresholdByValue|setValue|clearValue| MsgVpn|authenticationBasicProfileName|authenticationBasicType| MsgVpn|authorizationProfileName|authorizationType| MsgVpn|eventPublishTopicFormatMqttEnabled|eventPublishTopicFormatSmfEnabled| MsgVpn|eventPublishTopicFormatSmfEnabled|eventPublishTopicFormatMqttEnabled| MsgVpn|replicationBridgeAuthenticationBasicClientUsername|replicationBridgeAuthenticationBasicPassword| MsgVpn|replicationBridgeAuthenticationBasicPassword|replicationBridgeAuthenticationBasicClientUsername| MsgVpn|replicationBridgeAuthenticationClientCertPassword|replicationBridgeAuthenticationClientCertContent| MsgVpn|replicationEnabledQueueBehavior|replicationEnabled|    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpn**](MsgVpn.md)| The Message VPN object&#x27;s attributes. | 
 **optional** | ***MsgVpnApiCreateMsgVpnOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnResponse**](MsgVpnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnAclProfile**
> MsgVpnAclProfileResponse CreateMsgVpnAclProfile(ctx, body, msgVpnName, optional)
Create an ACL Profile object.

Create an ACL Profile object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  An ACL Profile controls whether an authenticated client is permitted to establish a connection with the message broker or permitted to publish and subscribe to specific topics.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: aclProfileName|x|x|||| msgVpnName|x||x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnAclProfile**](MsgVpnAclProfile.md)| The ACL Profile object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MsgVpnApiCreateMsgVpnAclProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnAclProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileResponse**](MsgVpnAclProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnAclProfileClientConnectException**
> MsgVpnAclProfileClientConnectExceptionResponse CreateMsgVpnAclProfileClientConnectException(ctx, body, msgVpnName, aclProfileName, optional)
Create a Client Connect Exception object.

Create a Client Connect Exception object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Client Connect Exception is an exception to the default action to take when a client using the ACL Profile connects to the Message VPN. Exceptions must be expressed as an IP address/netmask in CIDR form.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: aclProfileName|x||x||| clientConnectExceptionAddress|x|x|||| msgVpnName|x||x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnAclProfileClientConnectException**](MsgVpnAclProfileClientConnectException.md)| The Client Connect Exception object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
 **optional** | ***MsgVpnApiCreateMsgVpnAclProfileClientConnectExceptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnAclProfileClientConnectExceptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileClientConnectExceptionResponse**](MsgVpnAclProfileClientConnectExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnAclProfilePublishException**
> MsgVpnAclProfilePublishExceptionResponse CreateMsgVpnAclProfilePublishException(ctx, body, msgVpnName, aclProfileName, optional)
Create a Publish Topic Exception object.

Create a Publish Topic Exception object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Publish Topic Exception is an exception to the default action to take when a client using the ACL Profile publishes to a topic in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: aclProfileName|x||x||x| msgVpnName|x||x||x| publishExceptionTopic|x|x|||x| topicSyntax|x|x|||x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been deprecated since 2.14. Replaced by publishTopicExceptions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnAclProfilePublishException**](MsgVpnAclProfilePublishException.md)| The Publish Topic Exception object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
 **optional** | ***MsgVpnApiCreateMsgVpnAclProfilePublishExceptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnAclProfilePublishExceptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfilePublishExceptionResponse**](MsgVpnAclProfilePublishExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnAclProfilePublishTopicException**
> MsgVpnAclProfilePublishTopicExceptionResponse CreateMsgVpnAclProfilePublishTopicException(ctx, body, msgVpnName, aclProfileName, optional)
Create a Publish Topic Exception object.

Create a Publish Topic Exception object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Publish Topic Exception is an exception to the default action to take when a client using the ACL Profile publishes to a topic in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: aclProfileName|x||x||| msgVpnName|x||x||| publishTopicException|x|x|||| publishTopicExceptionSyntax|x|x||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnAclProfilePublishTopicException**](MsgVpnAclProfilePublishTopicException.md)| The Publish Topic Exception object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
 **optional** | ***MsgVpnApiCreateMsgVpnAclProfilePublishTopicExceptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnAclProfilePublishTopicExceptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfilePublishTopicExceptionResponse**](MsgVpnAclProfilePublishTopicExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnAclProfileSubscribeException**
> MsgVpnAclProfileSubscribeExceptionResponse CreateMsgVpnAclProfileSubscribeException(ctx, body, msgVpnName, aclProfileName, optional)
Create a Subscribe Topic Exception object.

Create a Subscribe Topic Exception object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Subscribe Topic Exception is an exception to the default action to take when a client using the ACL Profile subscribes to a topic in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: aclProfileName|x||x||x| msgVpnName|x||x||x| subscribeExceptionTopic|x|x|||x| topicSyntax|x|x|||x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been deprecated since 2.14. Replaced by subscribeTopicExceptions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnAclProfileSubscribeException**](MsgVpnAclProfileSubscribeException.md)| The Subscribe Topic Exception object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
 **optional** | ***MsgVpnApiCreateMsgVpnAclProfileSubscribeExceptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnAclProfileSubscribeExceptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileSubscribeExceptionResponse**](MsgVpnAclProfileSubscribeExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnAclProfileSubscribeShareNameException**
> MsgVpnAclProfileSubscribeShareNameExceptionResponse CreateMsgVpnAclProfileSubscribeShareNameException(ctx, body, msgVpnName, aclProfileName, optional)
Create a Subscribe Share Name Exception object.

Create a Subscribe Share Name Exception object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Subscribe Share Name Exception is an exception to the default action to take when a client using the ACL Profile subscribes to a share-name subscription in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: aclProfileName|x||x||| msgVpnName|x||x||| subscribeShareNameException|x|x|||| subscribeShareNameExceptionSyntax|x|x||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnAclProfileSubscribeShareNameException**](MsgVpnAclProfileSubscribeShareNameException.md)| The Subscribe Share Name Exception object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
 **optional** | ***MsgVpnApiCreateMsgVpnAclProfileSubscribeShareNameExceptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnAclProfileSubscribeShareNameExceptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileSubscribeShareNameExceptionResponse**](MsgVpnAclProfileSubscribeShareNameExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnAclProfileSubscribeTopicException**
> MsgVpnAclProfileSubscribeTopicExceptionResponse CreateMsgVpnAclProfileSubscribeTopicException(ctx, body, msgVpnName, aclProfileName, optional)
Create a Subscribe Topic Exception object.

Create a Subscribe Topic Exception object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Subscribe Topic Exception is an exception to the default action to take when a client using the ACL Profile subscribes to a topic in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: aclProfileName|x||x||| msgVpnName|x||x||| subscribeTopicException|x|x|||| subscribeTopicExceptionSyntax|x|x||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnAclProfileSubscribeTopicException**](MsgVpnAclProfileSubscribeTopicException.md)| The Subscribe Topic Exception object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
 **optional** | ***MsgVpnApiCreateMsgVpnAclProfileSubscribeTopicExceptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnAclProfileSubscribeTopicExceptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileSubscribeTopicExceptionResponse**](MsgVpnAclProfileSubscribeTopicExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnAuthenticationOauthProfile**
> MsgVpnAuthenticationOauthProfileResponse CreateMsgVpnAuthenticationOauthProfile(ctx, body, msgVpnName, optional)
Create an OAuth Profile object.

Create an OAuth Profile object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  OAuth profiles specify how to securely authenticate to an OAuth provider.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: clientSecret||||x||x msgVpnName|x||x||| oauthProfileName|x|x||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.25.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnAuthenticationOauthProfile**](MsgVpnAuthenticationOauthProfile.md)| The OAuth Profile object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MsgVpnApiCreateMsgVpnAuthenticationOauthProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnAuthenticationOauthProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthenticationOauthProfileResponse**](MsgVpnAuthenticationOauthProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnAuthenticationOauthProfileClientRequiredClaim**
> MsgVpnAuthenticationOauthProfileClientRequiredClaimResponse CreateMsgVpnAuthenticationOauthProfileClientRequiredClaim(ctx, body, msgVpnName, oauthProfileName, optional)
Create a Required Claim object.

Create a Required Claim object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  Additional claims to be verified in the ID token.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: clientRequiredClaimName|x|x|||| clientRequiredClaimValue||x|||| msgVpnName|x||x||| oauthProfileName|x||x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.25.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnAuthenticationOauthProfileClientRequiredClaim**](MsgVpnAuthenticationOauthProfileClientRequiredClaim.md)| The Required Claim object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
 **optional** | ***MsgVpnApiCreateMsgVpnAuthenticationOauthProfileClientRequiredClaimOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnAuthenticationOauthProfileClientRequiredClaimOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthenticationOauthProfileClientRequiredClaimResponse**](MsgVpnAuthenticationOauthProfileClientRequiredClaimResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnAuthenticationOauthProfileResourceServerRequiredClaim**
> MsgVpnAuthenticationOauthProfileResourceServerRequiredClaimResponse CreateMsgVpnAuthenticationOauthProfileResourceServerRequiredClaim(ctx, body, msgVpnName, oauthProfileName, optional)
Create a Required Claim object.

Create a Required Claim object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  Additional claims to be verified in the access token.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x||| oauthProfileName|x||x||| resourceServerRequiredClaimName|x|x|||| resourceServerRequiredClaimValue||x||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.25.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnAuthenticationOauthProfileResourceServerRequiredClaim**](MsgVpnAuthenticationOauthProfileResourceServerRequiredClaim.md)| The Required Claim object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
 **optional** | ***MsgVpnApiCreateMsgVpnAuthenticationOauthProfileResourceServerRequiredClaimOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnAuthenticationOauthProfileResourceServerRequiredClaimOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthenticationOauthProfileResourceServerRequiredClaimResponse**](MsgVpnAuthenticationOauthProfileResourceServerRequiredClaimResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnAuthenticationOauthProvider**
> MsgVpnAuthenticationOauthProviderResponse CreateMsgVpnAuthenticationOauthProvider(ctx, body, msgVpnName, optional)
Create an OAuth Provider object.

Create an OAuth Provider object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  OAuth Providers contain information about the issuer of an OAuth token that is needed to validate the token and derive a client username from it.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: audienceClaimName|||||x| audienceClaimSource|||||x| audienceClaimValue|||||x| audienceValidationEnabled|||||x| authorizationGroupClaimName|||||x| authorizationGroupClaimSource|||||x| authorizationGroupEnabled|||||x| disconnectOnTokenExpirationEnabled|||||x| enabled|||||x| jwksRefreshInterval|||||x| jwksUri|||||x| msgVpnName|x||x||x| oauthProviderName|x|x|||x| tokenIgnoreTimeLimitsEnabled|||||x| tokenIntrospectionParameterName|||||x| tokenIntrospectionPassword||||x|x|x tokenIntrospectionTimeout|||||x| tokenIntrospectionUri|||||x| tokenIntrospectionUsername|||||x| usernameClaimName|||||x| usernameClaimSource|||||x| usernameValidateEnabled|||||x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been deprecated since 2.25. Replaced by authenticationOauthProfiles.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnAuthenticationOauthProvider**](MsgVpnAuthenticationOauthProvider.md)| The OAuth Provider object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MsgVpnApiCreateMsgVpnAuthenticationOauthProviderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnAuthenticationOauthProviderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthenticationOauthProviderResponse**](MsgVpnAuthenticationOauthProviderResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnAuthorizationGroup**
> MsgVpnAuthorizationGroupResponse CreateMsgVpnAuthorizationGroup(ctx, body, msgVpnName, optional)
Create an Authorization Group object.

Create an Authorization Group object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  To use client authorization groups configured on an external server to provide client authorizations, Authorization Group objects must be created on the Message VPN that match the authorization groups provisioned on the external server. These objects must be configured with the client profiles and ACL profiles that will be assigned to the clients that belong to those authorization groups. A newly created group is placed at the end of the group list which is the lowest priority.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: authorizationGroupName|x|x|||| msgVpnName|x||x||| orderAfterAuthorizationGroupName||||x|| orderBeforeAuthorizationGroupName||||x||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- MsgVpnAuthorizationGroup|orderAfterAuthorizationGroupName||orderBeforeAuthorizationGroupName MsgVpnAuthorizationGroup|orderBeforeAuthorizationGroupName||orderAfterAuthorizationGroupName    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnAuthorizationGroup**](MsgVpnAuthorizationGroup.md)| The Authorization Group object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MsgVpnApiCreateMsgVpnAuthorizationGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnAuthorizationGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthorizationGroupResponse**](MsgVpnAuthorizationGroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnBridge**
> MsgVpnBridgeResponse CreateMsgVpnBridge(ctx, body, msgVpnName, optional)
Create a Bridge object.

Create a Bridge object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  Bridges can be used to link two Message VPNs so that messages published to one Message VPN that match the topic subscriptions set for the bridge are also delivered to the linked Message VPN.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: bridgeName|x|x|||| bridgeVirtualRouter|x|x|||| msgVpnName|x||x||| remoteAuthenticationBasicPassword||||x||x remoteAuthenticationClientCertContent||||x||x remoteAuthenticationClientCertPassword||||x||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- MsgVpnBridge|remoteAuthenticationBasicClientUsername|remoteAuthenticationBasicPassword| MsgVpnBridge|remoteAuthenticationBasicPassword|remoteAuthenticationBasicClientUsername| MsgVpnBridge|remoteAuthenticationClientCertPassword|remoteAuthenticationClientCertContent|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnBridge**](MsgVpnBridge.md)| The Bridge object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MsgVpnApiCreateMsgVpnBridgeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnBridgeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeResponse**](MsgVpnBridgeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnBridgeRemoteMsgVpn**
> MsgVpnBridgeRemoteMsgVpnResponse CreateMsgVpnBridgeRemoteMsgVpn(ctx, body, msgVpnName, bridgeName, bridgeVirtualRouter, optional)
Create a Remote Message VPN object.

Create a Remote Message VPN object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  The Remote Message VPN is the Message VPN that the Bridge connects to.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: bridgeName|x||x||| bridgeVirtualRouter|x||x||| msgVpnName|x||x||| password||||x||x remoteMsgVpnInterface|x||||| remoteMsgVpnLocation|x|x|||| remoteMsgVpnName|x|x||||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- MsgVpnBridgeRemoteMsgVpn|clientUsername|password| MsgVpnBridgeRemoteMsgVpn|password|clientUsername|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnBridgeRemoteMsgVpn**](MsgVpnBridgeRemoteMsgVpn.md)| The Remote Message VPN object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **bridgeName** | **string**| The name of the Bridge. | 
  **bridgeVirtualRouter** | **string**| The virtual router of the Bridge. | 
 **optional** | ***MsgVpnApiCreateMsgVpnBridgeRemoteMsgVpnOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnBridgeRemoteMsgVpnOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeRemoteMsgVpnResponse**](MsgVpnBridgeRemoteMsgVpnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnBridgeRemoteSubscription**
> MsgVpnBridgeRemoteSubscriptionResponse CreateMsgVpnBridgeRemoteSubscription(ctx, body, msgVpnName, bridgeName, bridgeVirtualRouter, optional)
Create a Remote Subscription object.

Create a Remote Subscription object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Remote Subscription is a topic subscription used by the Message VPN Bridge to attract messages from the remote message broker.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: bridgeName|x||x||| bridgeVirtualRouter|x||x||| deliverAlwaysEnabled||x|||| msgVpnName|x||x||| remoteSubscriptionTopic|x|x||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnBridgeRemoteSubscription**](MsgVpnBridgeRemoteSubscription.md)| The Remote Subscription object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **bridgeName** | **string**| The name of the Bridge. | 
  **bridgeVirtualRouter** | **string**| The virtual router of the Bridge. | 
 **optional** | ***MsgVpnApiCreateMsgVpnBridgeRemoteSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnBridgeRemoteSubscriptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeRemoteSubscriptionResponse**](MsgVpnBridgeRemoteSubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnBridgeTlsTrustedCommonName**
> MsgVpnBridgeTlsTrustedCommonNameResponse CreateMsgVpnBridgeTlsTrustedCommonName(ctx, body, msgVpnName, bridgeName, bridgeVirtualRouter, optional)
Create a Trusted Common Name object.

Create a Trusted Common Name object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  The Trusted Common Names for the Bridge are used by encrypted transports to verify the name in the certificate presented by the remote node. They must include the common name of the remote node's server certificate or client certificate, depending upon the initiator of the connection.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: bridgeName|x||x||x| bridgeVirtualRouter|x||x||x| msgVpnName|x||x||x| tlsTrustedCommonName|x|x|||x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been deprecated since 2.18. Common Name validation has been replaced by Server Certificate Name validation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnBridgeTlsTrustedCommonName**](MsgVpnBridgeTlsTrustedCommonName.md)| The Trusted Common Name object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **bridgeName** | **string**| The name of the Bridge. | 
  **bridgeVirtualRouter** | **string**| The virtual router of the Bridge. | 
 **optional** | ***MsgVpnApiCreateMsgVpnBridgeTlsTrustedCommonNameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnBridgeTlsTrustedCommonNameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeTlsTrustedCommonNameResponse**](MsgVpnBridgeTlsTrustedCommonNameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnCertMatchingRule**
> MsgVpnCertMatchingRuleResponse CreateMsgVpnCertMatchingRule(ctx, body, msgVpnName, optional)
Create a Certificate Matching Rule object.

Create a Certificate Matching Rule object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Cert Matching Rule is a collection of conditions and attribute filters that all have to be satisfied for certificate to be acceptable as authentication for a given username.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x||| ruleName|x|x||||    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.27.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnCertMatchingRule**](MsgVpnCertMatchingRule.md)| The Certificate Matching Rule object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MsgVpnApiCreateMsgVpnCertMatchingRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnCertMatchingRuleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnCertMatchingRuleResponse**](MsgVpnCertMatchingRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnCertMatchingRuleAttributeFilter**
> MsgVpnCertMatchingRuleAttributeFilterResponse CreateMsgVpnCertMatchingRuleAttributeFilter(ctx, body, msgVpnName, ruleName, optional)
Create a Certificate Matching Rule Attribute Filter object.

Create a Certificate Matching Rule Attribute Filter object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Cert Matching Rule Attribute Filter compares a username attribute to a string.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: filterName|x|x|||| msgVpnName|x||x||| ruleName|x||x|||    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.28.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnCertMatchingRuleAttributeFilter**](MsgVpnCertMatchingRuleAttributeFilter.md)| The Certificate Matching Rule Attribute Filter object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **ruleName** | **string**| The name of the rule. | 
 **optional** | ***MsgVpnApiCreateMsgVpnCertMatchingRuleAttributeFilterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnCertMatchingRuleAttributeFilterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnCertMatchingRuleAttributeFilterResponse**](MsgVpnCertMatchingRuleAttributeFilterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnCertMatchingRuleCondition**
> MsgVpnCertMatchingRuleConditionResponse CreateMsgVpnCertMatchingRuleCondition(ctx, body, msgVpnName, ruleName, optional)
Create a Certificate Matching Rule Condition object.

Create a Certificate Matching Rule Condition object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Cert Matching Rule Condition compares data extracted from a certificate to a username attribute or an expression.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x||| ruleName|x||x||| source|x|x||||    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.27.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnCertMatchingRuleCondition**](MsgVpnCertMatchingRuleCondition.md)| The Certificate Matching Rule Condition object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **ruleName** | **string**| The name of the rule. | 
 **optional** | ***MsgVpnApiCreateMsgVpnCertMatchingRuleConditionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnCertMatchingRuleConditionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnCertMatchingRuleConditionResponse**](MsgVpnCertMatchingRuleConditionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnClientProfile**
> MsgVpnClientProfileResponse CreateMsgVpnClientProfile(ctx, body, msgVpnName, optional)
Create a Client Profile object.

Create a Client Profile object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  Client Profiles are used to assign common configuration properties to clients that have been successfully authorized.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: allowCutThroughForwardingEnabled|||||x| apiQueueManagementCopyFromOnCreateName|||||x| apiTopicEndpointManagementCopyFromOnCreateName|||||x| clientProfileName|x|x|||| msgVpnName|x||x|||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThreshold|clearPercent|setPercent|clearValue, setValue EventThreshold|clearValue|setValue|clearPercent, setPercent EventThreshold|setPercent|clearPercent|clearValue, setValue EventThreshold|setValue|clearValue|clearPercent, setPercent EventThresholdByPercent|clearPercent|setPercent| EventThresholdByPercent|setPercent|clearPercent|    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnClientProfile**](MsgVpnClientProfile.md)| The Client Profile object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MsgVpnApiCreateMsgVpnClientProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnClientProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientProfileResponse**](MsgVpnClientProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnClientUsername**
> MsgVpnClientUsernameResponse CreateMsgVpnClientUsername(ctx, body, msgVpnName, optional)
Create a Client Username object.

Create a Client Username object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A client is only authorized to connect to a Message VPN that is associated with a Client Username that the client has been assigned.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: clientUsername|x|x|||| msgVpnName|x||x||| password||||x||x    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnClientUsername**](MsgVpnClientUsername.md)| The Client Username object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MsgVpnApiCreateMsgVpnClientUsernameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnClientUsernameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientUsernameResponse**](MsgVpnClientUsernameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnClientUsernameAttribute**
> MsgVpnClientUsernameAttributeResponse CreateMsgVpnClientUsernameAttribute(ctx, body, msgVpnName, clientUsername, optional)
Create a Client Username Attribute object.

Create a Client Username Attribute object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A ClientUsername Attribute is a key+value pair that can be used to locate a client username, for example when using client certificate mapping.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: attributeName|x|x|||| attributeValue|x|x|||| clientUsername|x||x||| msgVpnName|x||x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.27.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnClientUsernameAttribute**](MsgVpnClientUsernameAttribute.md)| The Client Username Attribute object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **clientUsername** | **string**| The name of the Client Username. | 
 **optional** | ***MsgVpnApiCreateMsgVpnClientUsernameAttributeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnClientUsernameAttributeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientUsernameAttributeResponse**](MsgVpnClientUsernameAttributeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnDistributedCache**
> MsgVpnDistributedCacheResponse CreateMsgVpnDistributedCache(ctx, body, msgVpnName, optional)
Create a Distributed Cache object.

Create a Distributed Cache object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Distributed Cache is a collection of one or more Cache Clusters that belong to the same Message VPN. Each Cache Cluster in a Distributed Cache is configured to subscribe to a different set of topics. This effectively divides up the configured topic space, to provide scaling to very large topic spaces or very high cached message throughput.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: cacheName|x|x|||| msgVpnName|x||x|||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- MsgVpnDistributedCache|scheduledDeleteMsgDayList|scheduledDeleteMsgTimeList| MsgVpnDistributedCache|scheduledDeleteMsgTimeList|scheduledDeleteMsgDayList|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnDistributedCache**](MsgVpnDistributedCache.md)| The Distributed Cache object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MsgVpnApiCreateMsgVpnDistributedCacheOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnDistributedCacheOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheResponse**](MsgVpnDistributedCacheResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnDistributedCacheCluster**
> MsgVpnDistributedCacheClusterResponse CreateMsgVpnDistributedCacheCluster(ctx, body, msgVpnName, cacheName, optional)
Create a Cache Cluster object.

Create a Cache Cluster object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Cache Cluster is a collection of one or more Cache Instances that subscribe to exactly the same topics. Cache Instances are grouped together in a Cache Cluster for the purpose of fault tolerance and load balancing. As published messages are received, the message broker message bus sends these live data messages to the Cache Instances in the Cache Cluster. This enables client cache requests to be served by any of Cache Instances in the Cache Cluster.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: cacheName|x||x||| clusterName|x|x|||| msgVpnName|x||x|||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThresholdByPercent|clearPercent|setPercent| EventThresholdByPercent|setPercent|clearPercent| EventThresholdByValue|clearValue|setValue| EventThresholdByValue|setValue|clearValue|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnDistributedCacheCluster**](MsgVpnDistributedCacheCluster.md)| The Cache Cluster object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **cacheName** | **string**| The name of the Distributed Cache. | 
 **optional** | ***MsgVpnApiCreateMsgVpnDistributedCacheClusterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnDistributedCacheClusterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterResponse**](MsgVpnDistributedCacheClusterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnDistributedCacheClusterGlobalCachingHomeCluster**
> MsgVpnDistributedCacheClusterGlobalCachingHomeClusterResponse CreateMsgVpnDistributedCacheClusterGlobalCachingHomeCluster(ctx, body, msgVpnName, cacheName, clusterName, optional)
Create a Home Cache Cluster object.

Create a Home Cache Cluster object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Home Cache Cluster is a Cache Cluster that is the \"definitive\" Cache Cluster for a given topic in the context of the Global Caching feature.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: cacheName|x||x||| clusterName|x||x||| homeClusterName|x|x|||| msgVpnName|x||x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnDistributedCacheClusterGlobalCachingHomeCluster**](MsgVpnDistributedCacheClusterGlobalCachingHomeCluster.md)| The Home Cache Cluster object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **cacheName** | **string**| The name of the Distributed Cache. | 
  **clusterName** | **string**| The name of the Cache Cluster. | 
 **optional** | ***MsgVpnApiCreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterGlobalCachingHomeClusterResponse**](MsgVpnDistributedCacheClusterGlobalCachingHomeClusterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix**
> MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixResponse CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix(ctx, body, msgVpnName, cacheName, clusterName, homeClusterName, optional)
Create a Topic Prefix object.

Create a Topic Prefix object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Topic Prefix is a prefix for a global topic that is available from the containing Home Cache Cluster.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: cacheName|x||x||| clusterName|x||x||| homeClusterName|x||x||| msgVpnName|x||x||| topicPrefix|x|x||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix**](MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix.md)| The Topic Prefix object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **cacheName** | **string**| The name of the Distributed Cache. | 
  **clusterName** | **string**| The name of the Cache Cluster. | 
  **homeClusterName** | **string**| The name of the remote Home Cache Cluster. | 
 **optional** | ***MsgVpnApiCreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixResponse**](MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnDistributedCacheClusterInstance**
> MsgVpnDistributedCacheClusterInstanceResponse CreateMsgVpnDistributedCacheClusterInstance(ctx, body, msgVpnName, cacheName, clusterName, optional)
Create a Cache Instance object.

Create a Cache Instance object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Cache Instance is a single Cache process that belongs to a single Cache Cluster. A Cache Instance object provisioned on the broker is used to disseminate configuration information to the Cache process. Cache Instances listen for and cache live data messages that match the topic subscriptions configured for their parent Cache Cluster.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: cacheName|x||x||| clusterName|x||x||| instanceName|x|x|||| msgVpnName|x||x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnDistributedCacheClusterInstance**](MsgVpnDistributedCacheClusterInstance.md)| The Cache Instance object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **cacheName** | **string**| The name of the Distributed Cache. | 
  **clusterName** | **string**| The name of the Cache Cluster. | 
 **optional** | ***MsgVpnApiCreateMsgVpnDistributedCacheClusterInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnDistributedCacheClusterInstanceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterInstanceResponse**](MsgVpnDistributedCacheClusterInstanceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnDistributedCacheClusterTopic**
> MsgVpnDistributedCacheClusterTopicResponse CreateMsgVpnDistributedCacheClusterTopic(ctx, body, msgVpnName, cacheName, clusterName, optional)
Create a Topic object.

Create a Topic object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  The Cache Instances that belong to the containing Cache Cluster will cache any messages published to topics that match a Topic Subscription.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: cacheName|x||x||| clusterName|x||x||| msgVpnName|x||x||| topic|x|x||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnDistributedCacheClusterTopic**](MsgVpnDistributedCacheClusterTopic.md)| The Topic object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **cacheName** | **string**| The name of the Distributed Cache. | 
  **clusterName** | **string**| The name of the Cache Cluster. | 
 **optional** | ***MsgVpnApiCreateMsgVpnDistributedCacheClusterTopicOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnDistributedCacheClusterTopicOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterTopicResponse**](MsgVpnDistributedCacheClusterTopicResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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
 **optional** | ***MsgVpnApiCreateMsgVpnDmrBridgeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnDmrBridgeOpts struct
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

# **CreateMsgVpnJndiConnectionFactory**
> MsgVpnJndiConnectionFactoryResponse CreateMsgVpnJndiConnectionFactory(ctx, body, msgVpnName, optional)
Create a JNDI Connection Factory object.

Create a JNDI Connection Factory object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  The message broker provides an internal JNDI store for provisioned Connection Factory objects that clients can access through JNDI lookups.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: connectionFactoryName|x|x|||| msgVpnName|x||x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.2.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnJndiConnectionFactory**](MsgVpnJndiConnectionFactory.md)| The JNDI Connection Factory object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MsgVpnApiCreateMsgVpnJndiConnectionFactoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnJndiConnectionFactoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnJndiConnectionFactoryResponse**](MsgVpnJndiConnectionFactoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnJndiQueue**
> MsgVpnJndiQueueResponse CreateMsgVpnJndiQueue(ctx, body, msgVpnName, optional)
Create a JNDI Queue object.

Create a JNDI Queue object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  The message broker provides an internal JNDI store for provisioned Queue objects that clients can access through JNDI lookups.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x||| queueName|x|x||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.2.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnJndiQueue**](MsgVpnJndiQueue.md)| The JNDI Queue object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MsgVpnApiCreateMsgVpnJndiQueueOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnJndiQueueOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnJndiQueueResponse**](MsgVpnJndiQueueResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnJndiTopic**
> MsgVpnJndiTopicResponse CreateMsgVpnJndiTopic(ctx, body, msgVpnName, optional)
Create a JNDI Topic object.

Create a JNDI Topic object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  The message broker provides an internal JNDI store for provisioned Topic objects that clients can access through JNDI lookups.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x||| topicName|x|x||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.2.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnJndiTopic**](MsgVpnJndiTopic.md)| The JNDI Topic object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MsgVpnApiCreateMsgVpnJndiTopicOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnJndiTopicOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnJndiTopicResponse**](MsgVpnJndiTopicResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnMqttRetainCache**
> MsgVpnMqttRetainCacheResponse CreateMsgVpnMqttRetainCache(ctx, body, msgVpnName, optional)
Create an MQTT Retain Cache object.

Create an MQTT Retain Cache object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  Using MQTT retained messages allows publishing MQTT clients to indicate that a message must be stored for later delivery to subscribing clients when those subscribing clients add subscriptions matching the retained message's topic. An MQTT Retain Cache processes all retained messages for a Message VPN.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: cacheName|x|x|||| msgVpnName|x||x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnMqttRetainCache**](MsgVpnMqttRetainCache.md)| The MQTT Retain Cache object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MsgVpnApiCreateMsgVpnMqttRetainCacheOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnMqttRetainCacheOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnMqttRetainCacheResponse**](MsgVpnMqttRetainCacheResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnMqttSession**
> MsgVpnMqttSessionResponse CreateMsgVpnMqttSession(ctx, body, msgVpnName, optional)
Create an MQTT Session object.

Create an MQTT Session object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  An MQTT Session object is a virtual representation of an MQTT client connection. An MQTT session holds the state of an MQTT client (that is, it is used to contain a client's QoS 0 and QoS 1 subscription sets and any undelivered QoS 1 messages).   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: mqttSessionClientId|x|x|||| mqttSessionVirtualRouter|x|x|||| msgVpnName|x||x|||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThreshold|clearPercent|setPercent|clearValue, setValue EventThreshold|clearValue|setValue|clearPercent, setPercent EventThreshold|setPercent|clearPercent|clearValue, setValue EventThreshold|setValue|clearValue|clearPercent, setPercent    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.1.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnMqttSession**](MsgVpnMqttSession.md)| The MQTT Session object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MsgVpnApiCreateMsgVpnMqttSessionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnMqttSessionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnMqttSessionResponse**](MsgVpnMqttSessionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnMqttSessionSubscription**
> MsgVpnMqttSessionSubscriptionResponse CreateMsgVpnMqttSessionSubscription(ctx, body, msgVpnName, mqttSessionClientId, mqttSessionVirtualRouter, optional)
Create a Subscription object.

Create a Subscription object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  An MQTT session contains a client's QoS 0 and QoS 1 subscription sets. On creation, a subscription defaults to QoS 0.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: mqttSessionClientId|x||x||| mqttSessionVirtualRouter|x||x||| msgVpnName|x||x||| subscriptionTopic|x|x||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.1.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnMqttSessionSubscription**](MsgVpnMqttSessionSubscription.md)| The Subscription object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **mqttSessionClientId** | **string**| The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet. | 
  **mqttSessionVirtualRouter** | **string**| The virtual router of the MQTT Session. | 
 **optional** | ***MsgVpnApiCreateMsgVpnMqttSessionSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnMqttSessionSubscriptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnMqttSessionSubscriptionResponse**](MsgVpnMqttSessionSubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnQueue**
> MsgVpnQueueResponse CreateMsgVpnQueue(ctx, body, msgVpnName, optional)
Create a Queue object.

Create a Queue object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Queue acts as both a destination that clients can publish messages to, and as an endpoint that clients can bind consumers to and consume messages from.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x||| queueName|x|x||||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThreshold|clearPercent|setPercent|clearValue, setValue EventThreshold|clearValue|setValue|clearPercent, setPercent EventThreshold|setPercent|clearPercent|clearValue, setValue EventThreshold|setValue|clearValue|clearPercent, setPercent    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnQueue**](MsgVpnQueue.md)| The Queue object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MsgVpnApiCreateMsgVpnQueueOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnQueueOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueueResponse**](MsgVpnQueueResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnQueueSubscription**
> MsgVpnQueueSubscriptionResponse CreateMsgVpnQueueSubscription(ctx, body, msgVpnName, queueName, optional)
Create a Queue Subscription object.

Create a Queue Subscription object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  One or more Queue Subscriptions can be added to a durable queue so that Guaranteed messages published to matching topics are also delivered to and spooled by the queue.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x||| queueName|x||x||| subscriptionTopic|x|x||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnQueueSubscription**](MsgVpnQueueSubscription.md)| The Queue Subscription object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The name of the Queue. | 
 **optional** | ***MsgVpnApiCreateMsgVpnQueueSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnQueueSubscriptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueueSubscriptionResponse**](MsgVpnQueueSubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnQueueTemplate**
> MsgVpnQueueTemplateResponse CreateMsgVpnQueueTemplate(ctx, body, msgVpnName, optional)
Create a Queue Template object.

Create a Queue Template object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Queue Template provides a mechanism for specifying the initial state for client created queues.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x||| queueTemplateName|x|x||||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThreshold|clearPercent|setPercent|clearValue, setValue EventThreshold|clearValue|setValue|clearPercent, setPercent EventThreshold|setPercent|clearPercent|clearValue, setValue EventThreshold|setValue|clearValue|clearPercent, setPercent    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnQueueTemplate**](MsgVpnQueueTemplate.md)| The Queue Template object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MsgVpnApiCreateMsgVpnQueueTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnQueueTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueueTemplateResponse**](MsgVpnQueueTemplateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnReplayLog**
> MsgVpnReplayLogResponse CreateMsgVpnReplayLog(ctx, body, msgVpnName, optional)
Create a Replay Log object.

Create a Replay Log object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  When the Message Replay feature is enabled, message brokers store persistent messages in a Replay Log. These messages are kept until the log is full, after which the oldest messages are removed to free up space for new messages.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x||| replayLogName|x|x||||    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.10.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnReplayLog**](MsgVpnReplayLog.md)| The Replay Log object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MsgVpnApiCreateMsgVpnReplayLogOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnReplayLogOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnReplayLogResponse**](MsgVpnReplayLogResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnReplayLogTopicFilterSubscription**
> MsgVpnReplayLogTopicFilterSubscriptionResponse CreateMsgVpnReplayLogTopicFilterSubscription(ctx, body, msgVpnName, replayLogName, optional)
Create a Topic Filter Subscription object.

Create a Topic Filter Subscription object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  One or more Subscriptions can be added to a replay-log so that only guaranteed messages published to matching topics are stored in the Replay Log.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x||| replayLogName|x||x||| topicFilterSubscription|x|x||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.27.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnReplayLogTopicFilterSubscription**](MsgVpnReplayLogTopicFilterSubscription.md)| The Topic Filter Subscription object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **replayLogName** | **string**| The name of the Replay Log. | 
 **optional** | ***MsgVpnApiCreateMsgVpnReplayLogTopicFilterSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnReplayLogTopicFilterSubscriptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnReplayLogTopicFilterSubscriptionResponse**](MsgVpnReplayLogTopicFilterSubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnReplicatedTopic**
> MsgVpnReplicatedTopicResponse CreateMsgVpnReplicatedTopic(ctx, body, msgVpnName, optional)
Create a Replicated Topic object.

Create a Replicated Topic object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  To indicate which messages should be replicated between the active and standby site, a Replicated Topic subscription must be configured on a Message VPN. If a published message matches both a replicated topic and an endpoint on the active site, then the message is replicated to the standby site.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x||| replicatedTopic|x|x||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.1.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnReplicatedTopic**](MsgVpnReplicatedTopic.md)| The Replicated Topic object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MsgVpnApiCreateMsgVpnReplicatedTopicOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnReplicatedTopicOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnReplicatedTopicResponse**](MsgVpnReplicatedTopicResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnRestDeliveryPoint**
> MsgVpnRestDeliveryPointResponse CreateMsgVpnRestDeliveryPoint(ctx, body, msgVpnName, optional)
Create a REST Delivery Point object.

Create a REST Delivery Point object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A REST Delivery Point manages delivery of messages from queues to a named list of REST Consumers.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x||| restDeliveryPointName|x|x||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnRestDeliveryPoint**](MsgVpnRestDeliveryPoint.md)| The REST Delivery Point object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MsgVpnApiCreateMsgVpnRestDeliveryPointOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnRestDeliveryPointOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointResponse**](MsgVpnRestDeliveryPointResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnRestDeliveryPointQueueBinding**
> MsgVpnRestDeliveryPointQueueBindingResponse CreateMsgVpnRestDeliveryPointQueueBinding(ctx, body, msgVpnName, restDeliveryPointName, optional)
Create a Queue Binding object.

Create a Queue Binding object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Queue Binding for a REST Delivery Point attracts messages to be delivered to REST consumers. If the queue does not exist it can be created subsequently, and once the queue is operational the broker performs the queue binding. Removing the queue binding does not delete the queue itself. Similarly, removing the queue does not remove the queue binding, which fails until the queue is recreated or the queue binding is deleted.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x||| queueBindingName|x|x|||| restDeliveryPointName|x||x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnRestDeliveryPointQueueBinding**](MsgVpnRestDeliveryPointQueueBinding.md)| The Queue Binding object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
 **optional** | ***MsgVpnApiCreateMsgVpnRestDeliveryPointQueueBindingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnRestDeliveryPointQueueBindingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointQueueBindingResponse**](MsgVpnRestDeliveryPointQueueBindingResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader**
> MsgVpnRestDeliveryPointQueueBindingProtectedRequestHeaderResponse CreateMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader(ctx, body, msgVpnName, restDeliveryPointName, queueBindingName, optional)
Create a Protected Request Header object.

Create a Protected Request Header object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A protected request header to be added to the HTTP request. Unlike a non-protected request header, the header value cannot be displayed after it is set.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: headerName|x|x|||| headerValue||||x||x msgVpnName|x||x||| queueBindingName|x||x||| restDeliveryPointName|x||x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.30.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader**](MsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader.md)| The Protected Request Header object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **queueBindingName** | **string**| The name of a queue in the Message VPN. | 
 **optional** | ***MsgVpnApiCreateMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeaderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeaderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointQueueBindingProtectedRequestHeaderResponse**](MsgVpnRestDeliveryPointQueueBindingProtectedRequestHeaderResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnRestDeliveryPointQueueBindingRequestHeader**
> MsgVpnRestDeliveryPointQueueBindingRequestHeaderResponse CreateMsgVpnRestDeliveryPointQueueBindingRequestHeader(ctx, body, msgVpnName, restDeliveryPointName, queueBindingName, optional)
Create a Request Header object.

Create a Request Header object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A request header to be added to the HTTP request.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: headerName|x|x|||| msgVpnName|x||x||| queueBindingName|x||x||| restDeliveryPointName|x||x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.23.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnRestDeliveryPointQueueBindingRequestHeader**](MsgVpnRestDeliveryPointQueueBindingRequestHeader.md)| The Request Header object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **queueBindingName** | **string**| The name of a queue in the Message VPN. | 
 **optional** | ***MsgVpnApiCreateMsgVpnRestDeliveryPointQueueBindingRequestHeaderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnRestDeliveryPointQueueBindingRequestHeaderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointQueueBindingRequestHeaderResponse**](MsgVpnRestDeliveryPointQueueBindingRequestHeaderResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnRestDeliveryPointRestConsumer**
> MsgVpnRestDeliveryPointRestConsumerResponse CreateMsgVpnRestDeliveryPointRestConsumer(ctx, body, msgVpnName, restDeliveryPointName, optional)
Create a REST Consumer object.

Create a REST Consumer object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  REST Consumer objects establish HTTP connectivity to REST consumer applications who wish to receive messages from a broker.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: authenticationAwsSecretAccessKey||||x||x authenticationClientCertContent||||x||x authenticationClientCertPassword||||x|| authenticationHttpBasicPassword||||x||x authenticationHttpHeaderValue||||x||x authenticationOauthClientSecret||||x||x authenticationOauthJwtSecretKey||||x||x msgVpnName|x||x||| restConsumerName|x|x|||| restDeliveryPointName|x||x|||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- MsgVpnRestDeliveryPointRestConsumer|authenticationClientCertPassword|authenticationClientCertContent| MsgVpnRestDeliveryPointRestConsumer|authenticationHttpBasicPassword|authenticationHttpBasicUsername| MsgVpnRestDeliveryPointRestConsumer|authenticationHttpBasicUsername|authenticationHttpBasicPassword| MsgVpnRestDeliveryPointRestConsumer|remotePort|tlsEnabled| MsgVpnRestDeliveryPointRestConsumer|tlsEnabled|remotePort|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnRestDeliveryPointRestConsumer**](MsgVpnRestDeliveryPointRestConsumer.md)| The REST Consumer object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
 **optional** | ***MsgVpnApiCreateMsgVpnRestDeliveryPointRestConsumerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnRestDeliveryPointRestConsumerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointRestConsumerResponse**](MsgVpnRestDeliveryPointRestConsumerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim**
> MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimResponse CreateMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim(ctx, body, msgVpnName, restDeliveryPointName, restConsumerName, optional)
Create a Claim object.

Create a Claim object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Claim is added to the JWT sent to the OAuth token request endpoint.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x||| oauthJwtClaimName|x|x|||| oauthJwtClaimValue||x|||| restConsumerName|x||x||| restDeliveryPointName|x||x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.21.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnRestDeliveryPointRestConsumerOauthJwtClaim**](MsgVpnRestDeliveryPointRestConsumerOauthJwtClaim.md)| The Claim object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **restConsumerName** | **string**| The name of the REST Consumer. | 
 **optional** | ***MsgVpnApiCreateMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimResponse**](MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName**
> MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameResponse CreateMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName(ctx, body, msgVpnName, restDeliveryPointName, restConsumerName, optional)
Create a Trusted Common Name object.

Create a Trusted Common Name object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  The Trusted Common Names for the REST Consumer are used by encrypted transports to verify the name in the certificate presented by the remote REST consumer. They must include the common name of the remote REST consumer's server certificate.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x||x| restConsumerName|x||x||x| restDeliveryPointName|x||x||x| tlsTrustedCommonName|x|x|||x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been deprecated since (will be deprecated in next SEMP version). Common Name validation has been replaced by Server Certificate Name validation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName**](MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName.md)| The Trusted Common Name object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **restConsumerName** | **string**| The name of the REST Consumer. | 
 **optional** | ***MsgVpnApiCreateMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameResponse**](MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnSequencedTopic**
> MsgVpnSequencedTopicResponse CreateMsgVpnSequencedTopic(ctx, body, msgVpnName, optional)
Create a Sequenced Topic object.

Create a Sequenced Topic object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Sequenced Topic is a topic subscription for which any matching messages received on the Message VPN are assigned a sequence number that is monotonically increased by a value of one per message.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x||| sequencedTopic|x|x||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnSequencedTopic**](MsgVpnSequencedTopic.md)| The Sequenced Topic object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MsgVpnApiCreateMsgVpnSequencedTopicOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnSequencedTopicOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnSequencedTopicResponse**](MsgVpnSequencedTopicResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnTelemetryProfile**
> MsgVpnTelemetryProfileResponse CreateMsgVpnTelemetryProfile(ctx, body, msgVpnName, optional)
Create a Telemetry Profile object.

Create a Telemetry Profile object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  Using the Telemetry Profile allows trace spans to be generated as messages are processed by the broker. The generated spans are stored persistently on the broker and may be consumed by the Solace receiver component of an OpenTelemetry Collector.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x||| telemetryProfileName|x|x||||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThreshold|clearPercent|setPercent|clearValue, setValue EventThreshold|clearValue|setValue|clearPercent, setPercent EventThreshold|setPercent|clearPercent|clearValue, setValue EventThreshold|setValue|clearValue|clearPercent, setPercent    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.31.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnTelemetryProfile**](MsgVpnTelemetryProfile.md)| The Telemetry Profile object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MsgVpnApiCreateMsgVpnTelemetryProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnTelemetryProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTelemetryProfileResponse**](MsgVpnTelemetryProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnTelemetryProfileReceiverAclConnectException**
> MsgVpnTelemetryProfileReceiverAclConnectExceptionResponse CreateMsgVpnTelemetryProfileReceiverAclConnectException(ctx, body, msgVpnName, telemetryProfileName, optional)
Create a Receiver ACL Connect Exception object.

Create a Receiver ACL Connect Exception object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Receiver ACL Connect Exception is an exception to the default action to take when a receiver connects to the broker. Exceptions must be expressed as an IP address/netmask in CIDR form.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x||| receiverAclConnectExceptionAddress|x|x|||| telemetryProfileName|x||x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.31.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnTelemetryProfileReceiverAclConnectException**](MsgVpnTelemetryProfileReceiverAclConnectException.md)| The Receiver ACL Connect Exception object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **telemetryProfileName** | **string**| The name of the Telemetry Profile. | 
 **optional** | ***MsgVpnApiCreateMsgVpnTelemetryProfileReceiverAclConnectExceptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnTelemetryProfileReceiverAclConnectExceptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTelemetryProfileReceiverAclConnectExceptionResponse**](MsgVpnTelemetryProfileReceiverAclConnectExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnTelemetryProfileTraceFilter**
> MsgVpnTelemetryProfileTraceFilterResponse CreateMsgVpnTelemetryProfileTraceFilter(ctx, body, msgVpnName, telemetryProfileName, optional)
Create a Trace Filter object.

Create a Trace Filter object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Trace Filter controls which messages received by the broker will be traced. If an incoming message matches an enabled tracing filter's subscription, the message will be traced as it passes through the broker.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x||| telemetryProfileName|x||x||| traceFilterName|x|x||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.31.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnTelemetryProfileTraceFilter**](MsgVpnTelemetryProfileTraceFilter.md)| The Trace Filter object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **telemetryProfileName** | **string**| The name of the Telemetry Profile. | 
 **optional** | ***MsgVpnApiCreateMsgVpnTelemetryProfileTraceFilterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnTelemetryProfileTraceFilterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTelemetryProfileTraceFilterResponse**](MsgVpnTelemetryProfileTraceFilterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnTelemetryProfileTraceFilterSubscription**
> MsgVpnTelemetryProfileTraceFilterSubscriptionResponse CreateMsgVpnTelemetryProfileTraceFilterSubscription(ctx, body, msgVpnName, telemetryProfileName, traceFilterName, optional)
Create a Telemetry Trace Filter Subscription object.

Create a Telemetry Trace Filter Subscription object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  Trace filter subscriptions control which messages will be attracted by the tracing filter.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x||| subscription|x|x|||| subscriptionSyntax|x|x|||| telemetryProfileName|x||x||| traceFilterName|x||x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.31.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnTelemetryProfileTraceFilterSubscription**](MsgVpnTelemetryProfileTraceFilterSubscription.md)| The Telemetry Trace Filter Subscription object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **telemetryProfileName** | **string**| The name of the Telemetry Profile. | 
  **traceFilterName** | **string**| A name used to identify the trace filter. Consider a name that describes the subscriptions contained within the filter, such as the name of the application and/or the scenario in which the trace filter might be enabled, such as \&quot;appNameDebug\&quot;. | 
 **optional** | ***MsgVpnApiCreateMsgVpnTelemetryProfileTraceFilterSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnTelemetryProfileTraceFilterSubscriptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTelemetryProfileTraceFilterSubscriptionResponse**](MsgVpnTelemetryProfileTraceFilterSubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnTopicEndpoint**
> MsgVpnTopicEndpointResponse CreateMsgVpnTopicEndpoint(ctx, body, msgVpnName, optional)
Create a Topic Endpoint object.

Create a Topic Endpoint object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Topic Endpoint attracts messages published to a topic for which the Topic Endpoint has a matching topic subscription. The topic subscription for the Topic Endpoint is specified in the client request to bind a Flow to that Topic Endpoint. Queues are significantly more flexible than Topic Endpoints and are the recommended approach for most applications. The use of Topic Endpoints should be restricted to JMS applications.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x||| topicEndpointName|x|x||||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThreshold|clearPercent|setPercent|clearValue, setValue EventThreshold|clearValue|setValue|clearPercent, setPercent EventThreshold|setPercent|clearPercent|clearValue, setValue EventThreshold|setValue|clearValue|clearPercent, setPercent    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.1.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnTopicEndpoint**](MsgVpnTopicEndpoint.md)| The Topic Endpoint object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MsgVpnApiCreateMsgVpnTopicEndpointOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnTopicEndpointOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointResponse**](MsgVpnTopicEndpointResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMsgVpnTopicEndpointTemplate**
> MsgVpnTopicEndpointTemplateResponse CreateMsgVpnTopicEndpointTemplate(ctx, body, msgVpnName, optional)
Create a Topic Endpoint Template object.

Create a Topic Endpoint Template object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Topic Endpoint Template provides a mechanism for specifying the initial state for client created topic endpoints.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x||| topicEndpointTemplateName|x|x||||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThreshold|clearPercent|setPercent|clearValue, setValue EventThreshold|clearValue|setValue|clearPercent, setPercent EventThreshold|setPercent|clearPercent|clearValue, setValue EventThreshold|setValue|clearValue|clearPercent, setPercent    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnTopicEndpointTemplate**](MsgVpnTopicEndpointTemplate.md)| The Topic Endpoint Template object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MsgVpnApiCreateMsgVpnTopicEndpointTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiCreateMsgVpnTopicEndpointTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointTemplateResponse**](MsgVpnTopicEndpointTemplateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpn**
> SempMetaOnlyResponse DeleteMsgVpn(ctx, msgVpnName)
Delete a Message VPN object.

Delete a Message VPN object. The deletion of instances of this object are synchronized to HA mates via config-sync.  Message VPNs (Virtual Private Networks) allow for the segregation of topic space and clients. They also group clients connecting to a network of message brokers, such that messages published within a particular group are only visible to that group's clients.  A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnAclProfile**
> SempMetaOnlyResponse DeleteMsgVpnAclProfile(ctx, msgVpnName, aclProfileName)
Delete an ACL Profile object.

Delete an ACL Profile object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  An ACL Profile controls whether an authenticated client is permitted to establish a connection with the message broker or permitted to publish and subscribe to specific topics.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnAclProfileClientConnectException**
> SempMetaOnlyResponse DeleteMsgVpnAclProfileClientConnectException(ctx, msgVpnName, aclProfileName, clientConnectExceptionAddress)
Delete a Client Connect Exception object.

Delete a Client Connect Exception object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Client Connect Exception is an exception to the default action to take when a client using the ACL Profile connects to the Message VPN. Exceptions must be expressed as an IP address/netmask in CIDR form.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
  **clientConnectExceptionAddress** | **string**| The IP address/netmask of the client connect exception in canonical CIDR form. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnAclProfilePublishException**
> SempMetaOnlyResponse DeleteMsgVpnAclProfilePublishException(ctx, msgVpnName, aclProfileName, topicSyntax, publishExceptionTopic)
Delete a Publish Topic Exception object.

Delete a Publish Topic Exception object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Publish Topic Exception is an exception to the default action to take when a client using the ACL Profile publishes to a topic in the Message VPN. Exceptions must be expressed as a topic.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been deprecated since 2.14. Replaced by publishTopicExceptions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
  **topicSyntax** | **string**| The syntax of the topic for the exception to the default action taken. | 
  **publishExceptionTopic** | **string**| The topic for the exception to the default action taken. May include wildcard characters. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnAclProfilePublishTopicException**
> SempMetaOnlyResponse DeleteMsgVpnAclProfilePublishTopicException(ctx, msgVpnName, aclProfileName, publishTopicExceptionSyntax, publishTopicException)
Delete a Publish Topic Exception object.

Delete a Publish Topic Exception object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Publish Topic Exception is an exception to the default action to take when a client using the ACL Profile publishes to a topic in the Message VPN. Exceptions must be expressed as a topic.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
  **publishTopicExceptionSyntax** | **string**| The syntax of the topic for the exception to the default action taken. | 
  **publishTopicException** | **string**| The topic for the exception to the default action taken. May include wildcard characters. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnAclProfileSubscribeException**
> SempMetaOnlyResponse DeleteMsgVpnAclProfileSubscribeException(ctx, msgVpnName, aclProfileName, topicSyntax, subscribeExceptionTopic)
Delete a Subscribe Topic Exception object.

Delete a Subscribe Topic Exception object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Subscribe Topic Exception is an exception to the default action to take when a client using the ACL Profile subscribes to a topic in the Message VPN. Exceptions must be expressed as a topic.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been deprecated since 2.14. Replaced by subscribeTopicExceptions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
  **topicSyntax** | **string**| The syntax of the topic for the exception to the default action taken. | 
  **subscribeExceptionTopic** | **string**| The topic for the exception to the default action taken. May include wildcard characters. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnAclProfileSubscribeShareNameException**
> SempMetaOnlyResponse DeleteMsgVpnAclProfileSubscribeShareNameException(ctx, msgVpnName, aclProfileName, subscribeShareNameExceptionSyntax, subscribeShareNameException)
Delete a Subscribe Share Name Exception object.

Delete a Subscribe Share Name Exception object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Subscribe Share Name Exception is an exception to the default action to take when a client using the ACL Profile subscribes to a share-name subscription in the Message VPN. Exceptions must be expressed as a topic.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
  **subscribeShareNameExceptionSyntax** | **string**| The syntax of the subscribe share name for the exception to the default action taken. | 
  **subscribeShareNameException** | **string**| The subscribe share name exception to the default action taken. May include wildcard characters. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnAclProfileSubscribeTopicException**
> SempMetaOnlyResponse DeleteMsgVpnAclProfileSubscribeTopicException(ctx, msgVpnName, aclProfileName, subscribeTopicExceptionSyntax, subscribeTopicException)
Delete a Subscribe Topic Exception object.

Delete a Subscribe Topic Exception object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Subscribe Topic Exception is an exception to the default action to take when a client using the ACL Profile subscribes to a topic in the Message VPN. Exceptions must be expressed as a topic.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
  **subscribeTopicExceptionSyntax** | **string**| The syntax of the topic for the exception to the default action taken. | 
  **subscribeTopicException** | **string**| The topic for the exception to the default action taken. May include wildcard characters. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnAuthenticationOauthProfile**
> SempMetaOnlyResponse DeleteMsgVpnAuthenticationOauthProfile(ctx, msgVpnName, oauthProfileName)
Delete an OAuth Profile object.

Delete an OAuth Profile object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  OAuth profiles specify how to securely authenticate to an OAuth provider.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.25.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **oauthProfileName** | **string**| The name of the OAuth profile. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnAuthenticationOauthProfileClientRequiredClaim**
> SempMetaOnlyResponse DeleteMsgVpnAuthenticationOauthProfileClientRequiredClaim(ctx, msgVpnName, oauthProfileName, clientRequiredClaimName)
Delete a Required Claim object.

Delete a Required Claim object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  Additional claims to be verified in the ID token.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.25.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
  **clientRequiredClaimName** | **string**| The name of the ID token claim to verify. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnAuthenticationOauthProfileResourceServerRequiredClaim**
> SempMetaOnlyResponse DeleteMsgVpnAuthenticationOauthProfileResourceServerRequiredClaim(ctx, msgVpnName, oauthProfileName, resourceServerRequiredClaimName)
Delete a Required Claim object.

Delete a Required Claim object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  Additional claims to be verified in the access token.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.25.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
  **resourceServerRequiredClaimName** | **string**| The name of the access token claim to verify. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnAuthenticationOauthProvider**
> SempMetaOnlyResponse DeleteMsgVpnAuthenticationOauthProvider(ctx, msgVpnName, oauthProviderName)
Delete an OAuth Provider object.

Delete an OAuth Provider object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  OAuth Providers contain information about the issuer of an OAuth token that is needed to validate the token and derive a client username from it.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been deprecated since 2.25. Replaced by authenticationOauthProfiles.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **oauthProviderName** | **string**| The name of the OAuth Provider. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnAuthorizationGroup**
> SempMetaOnlyResponse DeleteMsgVpnAuthorizationGroup(ctx, msgVpnName, authorizationGroupName)
Delete an Authorization Group object.

Delete an Authorization Group object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  To use client authorization groups configured on an external server to provide client authorizations, Authorization Group objects must be created on the Message VPN that match the authorization groups provisioned on the external server. These objects must be configured with the client profiles and ACL profiles that will be assigned to the clients that belong to those authorization groups. A newly created group is placed at the end of the group list which is the lowest priority.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **authorizationGroupName** | **string**| The name of the Authorization Group. For LDAP groups, special care is needed if the group name contains special characters such as &#x27;#&#x27;, &#x27;+&#x27;, &#x27;;&#x27;, &#x27;&#x3D;&#x27; as the value of the group name returned from the LDAP server might prepend those characters with &#x27;\\&#x27;. For example a group name called &#x27;test#,lab,com&#x27; will be returned from the LDAP server as &#x27;test\\#,lab,com&#x27;. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnBridge**
> SempMetaOnlyResponse DeleteMsgVpnBridge(ctx, msgVpnName, bridgeName, bridgeVirtualRouter)
Delete a Bridge object.

Delete a Bridge object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  Bridges can be used to link two Message VPNs so that messages published to one Message VPN that match the topic subscriptions set for the bridge are also delivered to the linked Message VPN.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **bridgeName** | **string**| The name of the Bridge. | 
  **bridgeVirtualRouter** | **string**| The virtual router of the Bridge. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnBridgeRemoteMsgVpn**
> SempMetaOnlyResponse DeleteMsgVpnBridgeRemoteMsgVpn(ctx, msgVpnName, bridgeName, bridgeVirtualRouter, remoteMsgVpnName, remoteMsgVpnLocation, remoteMsgVpnInterface)
Delete a Remote Message VPN object.

Delete a Remote Message VPN object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  The Remote Message VPN is the Message VPN that the Bridge connects to.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

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

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnBridgeRemoteSubscription**
> SempMetaOnlyResponse DeleteMsgVpnBridgeRemoteSubscription(ctx, msgVpnName, bridgeName, bridgeVirtualRouter, remoteSubscriptionTopic)
Delete a Remote Subscription object.

Delete a Remote Subscription object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Remote Subscription is a topic subscription used by the Message VPN Bridge to attract messages from the remote message broker.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **bridgeName** | **string**| The name of the Bridge. | 
  **bridgeVirtualRouter** | **string**| The virtual router of the Bridge. | 
  **remoteSubscriptionTopic** | **string**| The topic of the Bridge remote subscription. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnBridgeTlsTrustedCommonName**
> SempMetaOnlyResponse DeleteMsgVpnBridgeTlsTrustedCommonName(ctx, msgVpnName, bridgeName, bridgeVirtualRouter, tlsTrustedCommonName)
Delete a Trusted Common Name object.

Delete a Trusted Common Name object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  The Trusted Common Names for the Bridge are used by encrypted transports to verify the name in the certificate presented by the remote node. They must include the common name of the remote node's server certificate or client certificate, depending upon the initiator of the connection.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been deprecated since 2.18. Common Name validation has been replaced by Server Certificate Name validation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **bridgeName** | **string**| The name of the Bridge. | 
  **bridgeVirtualRouter** | **string**| The virtual router of the Bridge. | 
  **tlsTrustedCommonName** | **string**| The expected trusted common name of the remote certificate. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnCertMatchingRule**
> SempMetaOnlyResponse DeleteMsgVpnCertMatchingRule(ctx, msgVpnName, ruleName)
Delete a Certificate Matching Rule object.

Delete a Certificate Matching Rule object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Cert Matching Rule is a collection of conditions and attribute filters that all have to be satisfied for certificate to be acceptable as authentication for a given username.  A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.27.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **ruleName** | **string**| The name of the rule. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnCertMatchingRuleAttributeFilter**
> SempMetaOnlyResponse DeleteMsgVpnCertMatchingRuleAttributeFilter(ctx, msgVpnName, ruleName, filterName)
Delete a Certificate Matching Rule Attribute Filter object.

Delete a Certificate Matching Rule Attribute Filter object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Cert Matching Rule Attribute Filter compares a username attribute to a string.  A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.28.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **ruleName** | **string**| The name of the rule. | 
  **filterName** | **string**| The name of the filter. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnCertMatchingRuleCondition**
> SempMetaOnlyResponse DeleteMsgVpnCertMatchingRuleCondition(ctx, msgVpnName, ruleName, source)
Delete a Certificate Matching Rule Condition object.

Delete a Certificate Matching Rule Condition object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Cert Matching Rule Condition compares data extracted from a certificate to a username attribute or an expression.  A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.27.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **ruleName** | **string**| The name of the rule. | 
  **source** | **string**| Certificate field to be compared with the Attribute. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnClientProfile**
> SempMetaOnlyResponse DeleteMsgVpnClientProfile(ctx, msgVpnName, clientProfileName)
Delete a Client Profile object.

Delete a Client Profile object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  Client Profiles are used to assign common configuration properties to clients that have been successfully authorized.  A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **clientProfileName** | **string**| The name of the Client Profile. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnClientUsername**
> SempMetaOnlyResponse DeleteMsgVpnClientUsername(ctx, msgVpnName, clientUsername)
Delete a Client Username object.

Delete a Client Username object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A client is only authorized to connect to a Message VPN that is associated with a Client Username that the client has been assigned.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **clientUsername** | **string**| The name of the Client Username. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnClientUsernameAttribute**
> SempMetaOnlyResponse DeleteMsgVpnClientUsernameAttribute(ctx, msgVpnName, clientUsername, attributeName, attributeValue)
Delete a Client Username Attribute object.

Delete a Client Username Attribute object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A ClientUsername Attribute is a key+value pair that can be used to locate a client username, for example when using client certificate mapping.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.27.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **clientUsername** | **string**| The name of the Client Username. | 
  **attributeName** | **string**| The name of the Attribute. | 
  **attributeValue** | **string**| The value of the Attribute. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnDistributedCache**
> SempMetaOnlyResponse DeleteMsgVpnDistributedCache(ctx, msgVpnName, cacheName)
Delete a Distributed Cache object.

Delete a Distributed Cache object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Distributed Cache is a collection of one or more Cache Clusters that belong to the same Message VPN. Each Cache Cluster in a Distributed Cache is configured to subscribe to a different set of topics. This effectively divides up the configured topic space, to provide scaling to very large topic spaces or very high cached message throughput.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **cacheName** | **string**| The name of the Distributed Cache. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnDistributedCacheCluster**
> SempMetaOnlyResponse DeleteMsgVpnDistributedCacheCluster(ctx, msgVpnName, cacheName, clusterName)
Delete a Cache Cluster object.

Delete a Cache Cluster object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Cache Cluster is a collection of one or more Cache Instances that subscribe to exactly the same topics. Cache Instances are grouped together in a Cache Cluster for the purpose of fault tolerance and load balancing. As published messages are received, the message broker message bus sends these live data messages to the Cache Instances in the Cache Cluster. This enables client cache requests to be served by any of Cache Instances in the Cache Cluster.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **cacheName** | **string**| The name of the Distributed Cache. | 
  **clusterName** | **string**| The name of the Cache Cluster. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeCluster**
> SempMetaOnlyResponse DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeCluster(ctx, msgVpnName, cacheName, clusterName, homeClusterName)
Delete a Home Cache Cluster object.

Delete a Home Cache Cluster object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Home Cache Cluster is a Cache Cluster that is the \"definitive\" Cache Cluster for a given topic in the context of the Global Caching feature.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **cacheName** | **string**| The name of the Distributed Cache. | 
  **clusterName** | **string**| The name of the Cache Cluster. | 
  **homeClusterName** | **string**| The name of the remote Home Cache Cluster. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix**
> SempMetaOnlyResponse DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix(ctx, msgVpnName, cacheName, clusterName, homeClusterName, topicPrefix)
Delete a Topic Prefix object.

Delete a Topic Prefix object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Topic Prefix is a prefix for a global topic that is available from the containing Home Cache Cluster.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **cacheName** | **string**| The name of the Distributed Cache. | 
  **clusterName** | **string**| The name of the Cache Cluster. | 
  **homeClusterName** | **string**| The name of the remote Home Cache Cluster. | 
  **topicPrefix** | **string**| A topic prefix for global topics available from the remote Home Cache Cluster. A wildcard (/&gt;) is implied at the end of the prefix. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnDistributedCacheClusterInstance**
> SempMetaOnlyResponse DeleteMsgVpnDistributedCacheClusterInstance(ctx, msgVpnName, cacheName, clusterName, instanceName)
Delete a Cache Instance object.

Delete a Cache Instance object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Cache Instance is a single Cache process that belongs to a single Cache Cluster. A Cache Instance object provisioned on the broker is used to disseminate configuration information to the Cache process. Cache Instances listen for and cache live data messages that match the topic subscriptions configured for their parent Cache Cluster.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **cacheName** | **string**| The name of the Distributed Cache. | 
  **clusterName** | **string**| The name of the Cache Cluster. | 
  **instanceName** | **string**| The name of the Cache Instance. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnDistributedCacheClusterTopic**
> SempMetaOnlyResponse DeleteMsgVpnDistributedCacheClusterTopic(ctx, msgVpnName, cacheName, clusterName, topic)
Delete a Topic object.

Delete a Topic object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  The Cache Instances that belong to the containing Cache Cluster will cache any messages published to topics that match a Topic Subscription.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **cacheName** | **string**| The name of the Distributed Cache. | 
  **clusterName** | **string**| The name of the Cache Cluster. | 
  **topic** | **string**| The value of the Topic in the form a/b/c. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
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

# **DeleteMsgVpnJndiConnectionFactory**
> SempMetaOnlyResponse DeleteMsgVpnJndiConnectionFactory(ctx, msgVpnName, connectionFactoryName)
Delete a JNDI Connection Factory object.

Delete a JNDI Connection Factory object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  The message broker provides an internal JNDI store for provisioned Connection Factory objects that clients can access through JNDI lookups.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.2.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **connectionFactoryName** | **string**| The name of the JMS Connection Factory. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnJndiQueue**
> SempMetaOnlyResponse DeleteMsgVpnJndiQueue(ctx, msgVpnName, queueName)
Delete a JNDI Queue object.

Delete a JNDI Queue object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  The message broker provides an internal JNDI store for provisioned Queue objects that clients can access through JNDI lookups.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.2.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The JNDI name of the JMS Queue. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnJndiTopic**
> SempMetaOnlyResponse DeleteMsgVpnJndiTopic(ctx, msgVpnName, topicName)
Delete a JNDI Topic object.

Delete a JNDI Topic object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  The message broker provides an internal JNDI store for provisioned Topic objects that clients can access through JNDI lookups.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.2.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicName** | **string**| The JNDI name of the JMS Topic. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnMqttRetainCache**
> SempMetaOnlyResponse DeleteMsgVpnMqttRetainCache(ctx, msgVpnName, cacheName)
Delete an MQTT Retain Cache object.

Delete an MQTT Retain Cache object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  Using MQTT retained messages allows publishing MQTT clients to indicate that a message must be stored for later delivery to subscribing clients when those subscribing clients add subscriptions matching the retained message's topic. An MQTT Retain Cache processes all retained messages for a Message VPN.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **cacheName** | **string**| The name of the MQTT Retain Cache. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnMqttSession**
> SempMetaOnlyResponse DeleteMsgVpnMqttSession(ctx, msgVpnName, mqttSessionClientId, mqttSessionVirtualRouter)
Delete an MQTT Session object.

Delete an MQTT Session object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  An MQTT Session object is a virtual representation of an MQTT client connection. An MQTT session holds the state of an MQTT client (that is, it is used to contain a client's QoS 0 and QoS 1 subscription sets and any undelivered QoS 1 messages).  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.1.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **mqttSessionClientId** | **string**| The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet. | 
  **mqttSessionVirtualRouter** | **string**| The virtual router of the MQTT Session. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnMqttSessionSubscription**
> SempMetaOnlyResponse DeleteMsgVpnMqttSessionSubscription(ctx, msgVpnName, mqttSessionClientId, mqttSessionVirtualRouter, subscriptionTopic)
Delete a Subscription object.

Delete a Subscription object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  An MQTT session contains a client's QoS 0 and QoS 1 subscription sets. On creation, a subscription defaults to QoS 0.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.1.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **mqttSessionClientId** | **string**| The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet. | 
  **mqttSessionVirtualRouter** | **string**| The virtual router of the MQTT Session. | 
  **subscriptionTopic** | **string**| The MQTT subscription topic. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnQueue**
> SempMetaOnlyResponse DeleteMsgVpnQueue(ctx, msgVpnName, queueName)
Delete a Queue object.

Delete a Queue object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Queue acts as both a destination that clients can publish messages to, and as an endpoint that clients can bind consumers to and consume messages from.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The name of the Queue. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnQueueSubscription**
> SempMetaOnlyResponse DeleteMsgVpnQueueSubscription(ctx, msgVpnName, queueName, subscriptionTopic)
Delete a Queue Subscription object.

Delete a Queue Subscription object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  One or more Queue Subscriptions can be added to a durable queue so that Guaranteed messages published to matching topics are also delivered to and spooled by the queue.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The name of the Queue. | 
  **subscriptionTopic** | **string**| The topic of the Subscription. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnQueueTemplate**
> SempMetaOnlyResponse DeleteMsgVpnQueueTemplate(ctx, msgVpnName, queueTemplateName)
Delete a Queue Template object.

Delete a Queue Template object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Queue Template provides a mechanism for specifying the initial state for client created queues.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueTemplateName** | **string**| The name of the Queue Template. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnReplayLog**
> SempMetaOnlyResponse DeleteMsgVpnReplayLog(ctx, msgVpnName, replayLogName)
Delete a Replay Log object.

Delete a Replay Log object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  When the Message Replay feature is enabled, message brokers store persistent messages in a Replay Log. These messages are kept until the log is full, after which the oldest messages are removed to free up space for new messages.  A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.10.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **replayLogName** | **string**| The name of the Replay Log. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnReplayLogTopicFilterSubscription**
> SempMetaOnlyResponse DeleteMsgVpnReplayLogTopicFilterSubscription(ctx, msgVpnName, replayLogName, topicFilterSubscription)
Delete a Topic Filter Subscription object.

Delete a Topic Filter Subscription object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  One or more Subscriptions can be added to a replay-log so that only guaranteed messages published to matching topics are stored in the Replay Log.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.27.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **replayLogName** | **string**| The name of the Replay Log. | 
  **topicFilterSubscription** | **string**| The topic of the Subscription. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnReplicatedTopic**
> SempMetaOnlyResponse DeleteMsgVpnReplicatedTopic(ctx, msgVpnName, replicatedTopic)
Delete a Replicated Topic object.

Delete a Replicated Topic object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  To indicate which messages should be replicated between the active and standby site, a Replicated Topic subscription must be configured on a Message VPN. If a published message matches both a replicated topic and an endpoint on the active site, then the message is replicated to the standby site.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.1.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **replicatedTopic** | **string**| The topic for applying replication. Published messages matching this topic will be replicated to the standby site. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnRestDeliveryPoint**
> SempMetaOnlyResponse DeleteMsgVpnRestDeliveryPoint(ctx, msgVpnName, restDeliveryPointName)
Delete a REST Delivery Point object.

Delete a REST Delivery Point object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A REST Delivery Point manages delivery of messages from queues to a named list of REST Consumers.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnRestDeliveryPointQueueBinding**
> SempMetaOnlyResponse DeleteMsgVpnRestDeliveryPointQueueBinding(ctx, msgVpnName, restDeliveryPointName, queueBindingName)
Delete a Queue Binding object.

Delete a Queue Binding object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Queue Binding for a REST Delivery Point attracts messages to be delivered to REST consumers. If the queue does not exist it can be created subsequently, and once the queue is operational the broker performs the queue binding. Removing the queue binding does not delete the queue itself. Similarly, removing the queue does not remove the queue binding, which fails until the queue is recreated or the queue binding is deleted.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **queueBindingName** | **string**| The name of a queue in the Message VPN. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader**
> SempMetaOnlyResponse DeleteMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader(ctx, msgVpnName, restDeliveryPointName, queueBindingName, headerName)
Delete a Protected Request Header object.

Delete a Protected Request Header object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A protected request header to be added to the HTTP request. Unlike a non-protected request header, the header value cannot be displayed after it is set.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.30.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **queueBindingName** | **string**| The name of a queue in the Message VPN. | 
  **headerName** | **string**| The name of the protected HTTP request header. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnRestDeliveryPointQueueBindingRequestHeader**
> SempMetaOnlyResponse DeleteMsgVpnRestDeliveryPointQueueBindingRequestHeader(ctx, msgVpnName, restDeliveryPointName, queueBindingName, headerName)
Delete a Request Header object.

Delete a Request Header object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A request header to be added to the HTTP request.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.23.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **queueBindingName** | **string**| The name of a queue in the Message VPN. | 
  **headerName** | **string**| The name of the HTTP request header. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnRestDeliveryPointRestConsumer**
> SempMetaOnlyResponse DeleteMsgVpnRestDeliveryPointRestConsumer(ctx, msgVpnName, restDeliveryPointName, restConsumerName)
Delete a REST Consumer object.

Delete a REST Consumer object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  REST Consumer objects establish HTTP connectivity to REST consumer applications who wish to receive messages from a broker.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **restConsumerName** | **string**| The name of the REST Consumer. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim**
> SempMetaOnlyResponse DeleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim(ctx, msgVpnName, restDeliveryPointName, restConsumerName, oauthJwtClaimName)
Delete a Claim object.

Delete a Claim object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Claim is added to the JWT sent to the OAuth token request endpoint.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.21.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **restConsumerName** | **string**| The name of the REST Consumer. | 
  **oauthJwtClaimName** | **string**| The name of the additional claim. Cannot be \&quot;exp\&quot;, \&quot;iat\&quot;, or \&quot;jti\&quot;. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName**
> SempMetaOnlyResponse DeleteMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName(ctx, msgVpnName, restDeliveryPointName, restConsumerName, tlsTrustedCommonName)
Delete a Trusted Common Name object.

Delete a Trusted Common Name object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  The Trusted Common Names for the REST Consumer are used by encrypted transports to verify the name in the certificate presented by the remote REST consumer. They must include the common name of the remote REST consumer's server certificate.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been deprecated since (will be deprecated in next SEMP version). Common Name validation has been replaced by Server Certificate Name validation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **restConsumerName** | **string**| The name of the REST Consumer. | 
  **tlsTrustedCommonName** | **string**| The expected trusted common name of the remote certificate. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnSequencedTopic**
> SempMetaOnlyResponse DeleteMsgVpnSequencedTopic(ctx, msgVpnName, sequencedTopic)
Delete a Sequenced Topic object.

Delete a Sequenced Topic object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Sequenced Topic is a topic subscription for which any matching messages received on the Message VPN are assigned a sequence number that is monotonically increased by a value of one per message.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **sequencedTopic** | **string**| Topic for applying sequence numbers. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnTelemetryProfile**
> SempMetaOnlyResponse DeleteMsgVpnTelemetryProfile(ctx, msgVpnName, telemetryProfileName)
Delete a Telemetry Profile object.

Delete a Telemetry Profile object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  Using the Telemetry Profile allows trace spans to be generated as messages are processed by the broker. The generated spans are stored persistently on the broker and may be consumed by the Solace receiver component of an OpenTelemetry Collector.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.31.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **telemetryProfileName** | **string**| The name of the Telemetry Profile. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnTelemetryProfileReceiverAclConnectException**
> SempMetaOnlyResponse DeleteMsgVpnTelemetryProfileReceiverAclConnectException(ctx, msgVpnName, telemetryProfileName, receiverAclConnectExceptionAddress)
Delete a Receiver ACL Connect Exception object.

Delete a Receiver ACL Connect Exception object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Receiver ACL Connect Exception is an exception to the default action to take when a receiver connects to the broker. Exceptions must be expressed as an IP address/netmask in CIDR form.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.31.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **telemetryProfileName** | **string**| The name of the Telemetry Profile. | 
  **receiverAclConnectExceptionAddress** | **string**| The IP address/netmask of the receiver connect exception in CIDR form. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnTelemetryProfileTraceFilter**
> SempMetaOnlyResponse DeleteMsgVpnTelemetryProfileTraceFilter(ctx, msgVpnName, telemetryProfileName, traceFilterName)
Delete a Trace Filter object.

Delete a Trace Filter object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Trace Filter controls which messages received by the broker will be traced. If an incoming message matches an enabled tracing filter's subscription, the message will be traced as it passes through the broker.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.31.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **telemetryProfileName** | **string**| The name of the Telemetry Profile. | 
  **traceFilterName** | **string**| A name used to identify the trace filter. Consider a name that describes the subscriptions contained within the filter, such as the name of the application and/or the scenario in which the trace filter might be enabled, such as \&quot;appNameDebug\&quot;. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnTelemetryProfileTraceFilterSubscription**
> SempMetaOnlyResponse DeleteMsgVpnTelemetryProfileTraceFilterSubscription(ctx, msgVpnName, telemetryProfileName, traceFilterName, subscription, subscriptionSyntax)
Delete a Telemetry Trace Filter Subscription object.

Delete a Telemetry Trace Filter Subscription object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  Trace filter subscriptions control which messages will be attracted by the tracing filter.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.31.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **telemetryProfileName** | **string**| The name of the Telemetry Profile. | 
  **traceFilterName** | **string**| A name used to identify the trace filter. Consider a name that describes the subscriptions contained within the filter, such as the name of the application and/or the scenario in which the trace filter might be enabled, such as \&quot;appNameDebug\&quot;. | 
  **subscription** | **string**| Messages matching this subscription will follow this filter&#x27;s configuration. | 
  **subscriptionSyntax** | **string**| The syntax of the trace filter subscription. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnTopicEndpoint**
> SempMetaOnlyResponse DeleteMsgVpnTopicEndpoint(ctx, msgVpnName, topicEndpointName)
Delete a Topic Endpoint object.

Delete a Topic Endpoint object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Topic Endpoint attracts messages published to a topic for which the Topic Endpoint has a matching topic subscription. The topic subscription for the Topic Endpoint is specified in the client request to bind a Flow to that Topic Endpoint. Queues are significantly more flexible than Topic Endpoints and are the recommended approach for most applications. The use of Topic Endpoints should be restricted to JMS applications.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.1.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointName** | **string**| The name of the Topic Endpoint. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsgVpnTopicEndpointTemplate**
> SempMetaOnlyResponse DeleteMsgVpnTopicEndpointTemplate(ctx, msgVpnName, topicEndpointTemplateName)
Delete a Topic Endpoint Template object.

Delete a Topic Endpoint Template object. The deletion of instances of this object are synchronized to HA mates and replication sites via config-sync.  A Topic Endpoint Template provides a mechanism for specifying the initial state for client created topic endpoints.  A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointTemplateName** | **string**| The name of the Topic Endpoint Template. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpn**
> MsgVpnResponse GetMsgVpn(ctx, msgVpnName, optional)
Get a Message VPN object.

Get a Message VPN object.  Message VPNs (Virtual Private Networks) allow for the segregation of topic space and clients. They also group clients connecting to a network of message brokers, such that messages published within a particular group are only visible to that group's clients.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: authenticationOauthDefaultProviderName|||x| bridgingTlsServerCertEnforceTrustedCommonNameEnabled|||x| distributedCacheManagementEnabled|||x| msgVpnName|x||| replicationBridgeAuthenticationBasicPassword||x||x replicationBridgeAuthenticationClientCertContent||x||x replicationBridgeAuthenticationClientCertPassword||x|| replicationEnabledQueueBehavior||x|| restTlsServerCertEnforceTrustedCommonNameEnabled|||x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.0.

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

 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get an ACL Profile object.  An ACL Profile controls whether an authenticated client is permitted to establish a connection with the message broker or permitted to publish and subscribe to specific topics.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: aclProfileName|x||| msgVpnName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.0.

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


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a Client Connect Exception object.  A Client Connect Exception is an exception to the default action to take when a client using the ACL Profile connects to the Message VPN. Exceptions must be expressed as an IP address/netmask in CIDR form.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: aclProfileName|x||| clientConnectExceptionAddress|x||| msgVpnName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.0.

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



 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Client Connect Exception objects.  A Client Connect Exception is an exception to the default action to take when a client using the ACL Profile connects to the Message VPN. Exceptions must be expressed as an IP address/netmask in CIDR form.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: aclProfileName|x||| clientConnectExceptionAddress|x||| msgVpnName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.0.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a Publish Topic Exception object.  A Publish Topic Exception is an exception to the default action to take when a client using the ACL Profile publishes to a topic in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: aclProfileName|x||x| msgVpnName|x||x| publishExceptionTopic|x||x| topicSyntax|x||x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been deprecated since 2.14. Replaced by publishTopicExceptions.

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




 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Publish Topic Exception objects.  A Publish Topic Exception is an exception to the default action to take when a client using the ACL Profile publishes to a topic in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: aclProfileName|x||x| msgVpnName|x||x| publishExceptionTopic|x||x| topicSyntax|x||x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been deprecated since 2.14. Replaced by publishTopicExceptions.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a Publish Topic Exception object.  A Publish Topic Exception is an exception to the default action to take when a client using the ACL Profile publishes to a topic in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: aclProfileName|x||| msgVpnName|x||| publishTopicException|x||| publishTopicExceptionSyntax|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.14.

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




 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Publish Topic Exception objects.  A Publish Topic Exception is an exception to the default action to take when a client using the ACL Profile publishes to a topic in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: aclProfileName|x||| msgVpnName|x||| publishTopicException|x||| publishTopicExceptionSyntax|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.14.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a Subscribe Topic Exception object.  A Subscribe Topic Exception is an exception to the default action to take when a client using the ACL Profile subscribes to a topic in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: aclProfileName|x||x| msgVpnName|x||x| subscribeExceptionTopic|x||x| topicSyntax|x||x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been deprecated since 2.14. Replaced by subscribeTopicExceptions.

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




 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Subscribe Topic Exception objects.  A Subscribe Topic Exception is an exception to the default action to take when a client using the ACL Profile subscribes to a topic in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: aclProfileName|x||x| msgVpnName|x||x| subscribeExceptionTopic|x||x| topicSyntax|x||x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been deprecated since 2.14. Replaced by subscribeTopicExceptions.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a Subscribe Share Name Exception object.  A Subscribe Share Name Exception is an exception to the default action to take when a client using the ACL Profile subscribes to a share-name subscription in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: aclProfileName|x||| msgVpnName|x||| subscribeShareNameException|x||| subscribeShareNameExceptionSyntax|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.14.

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




 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Subscribe Share Name Exception objects.  A Subscribe Share Name Exception is an exception to the default action to take when a client using the ACL Profile subscribes to a share-name subscription in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: aclProfileName|x||| msgVpnName|x||| subscribeShareNameException|x||| subscribeShareNameExceptionSyntax|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.14.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a Subscribe Topic Exception object.  A Subscribe Topic Exception is an exception to the default action to take when a client using the ACL Profile subscribes to a topic in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: aclProfileName|x||| msgVpnName|x||| subscribeTopicException|x||| subscribeTopicExceptionSyntax|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.14.

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




 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Subscribe Topic Exception objects.  A Subscribe Topic Exception is an exception to the default action to take when a client using the ACL Profile subscribes to a topic in the Message VPN. Exceptions must be expressed as a topic.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: aclProfileName|x||| msgVpnName|x||| subscribeTopicException|x||| subscribeTopicExceptionSyntax|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.14.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of ACL Profile objects.  An ACL Profile controls whether an authenticated client is permitted to establish a connection with the message broker or permitted to publish and subscribe to specific topics.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: aclProfileName|x||| msgVpnName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.0.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get an OAuth Profile object.  OAuth profiles specify how to securely authenticate to an OAuth provider.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: clientSecret||x||x msgVpnName|x||| oauthProfileName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.25.

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


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a Required Claim object.  Additional claims to be verified in the ID token.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: clientRequiredClaimName|x||| msgVpnName|x||| oauthProfileName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.25.

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



 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Required Claim objects.  Additional claims to be verified in the ID token.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: clientRequiredClaimName|x||| msgVpnName|x||| oauthProfileName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.25.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a Required Claim object.  Additional claims to be verified in the access token.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| oauthProfileName|x||| resourceServerRequiredClaimName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.25.

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



 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Required Claim objects.  Additional claims to be verified in the access token.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| oauthProfileName|x||| resourceServerRequiredClaimName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.25.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of OAuth Profile objects.  OAuth profiles specify how to securely authenticate to an OAuth provider.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: clientSecret||x||x msgVpnName|x||| oauthProfileName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.25.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get an OAuth Provider object.  OAuth Providers contain information about the issuer of an OAuth token that is needed to validate the token and derive a client username from it.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: audienceClaimName|||x| audienceClaimSource|||x| audienceClaimValue|||x| audienceValidationEnabled|||x| authorizationGroupClaimName|||x| authorizationGroupClaimSource|||x| authorizationGroupEnabled|||x| disconnectOnTokenExpirationEnabled|||x| enabled|||x| jwksRefreshInterval|||x| jwksUri|||x| msgVpnName|x||x| oauthProviderName|x||x| tokenIgnoreTimeLimitsEnabled|||x| tokenIntrospectionParameterName|||x| tokenIntrospectionPassword||x|x|x tokenIntrospectionTimeout|||x| tokenIntrospectionUri|||x| tokenIntrospectionUsername|||x| usernameClaimName|||x| usernameClaimSource|||x| usernameValidateEnabled|||x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been deprecated since 2.25. Replaced by authenticationOauthProfiles.

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


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of OAuth Provider objects.  OAuth Providers contain information about the issuer of an OAuth token that is needed to validate the token and derive a client username from it.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: audienceClaimName|||x| audienceClaimSource|||x| audienceClaimValue|||x| audienceValidationEnabled|||x| authorizationGroupClaimName|||x| authorizationGroupClaimSource|||x| authorizationGroupEnabled|||x| disconnectOnTokenExpirationEnabled|||x| enabled|||x| jwksRefreshInterval|||x| jwksUri|||x| msgVpnName|x||x| oauthProviderName|x||x| tokenIgnoreTimeLimitsEnabled|||x| tokenIntrospectionParameterName|||x| tokenIntrospectionPassword||x|x|x tokenIntrospectionTimeout|||x| tokenIntrospectionUri|||x| tokenIntrospectionUsername|||x| usernameClaimName|||x| usernameClaimSource|||x| usernameValidateEnabled|||x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been deprecated since 2.25. Replaced by authenticationOauthProfiles.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get an Authorization Group object.  To use client authorization groups configured on an external server to provide client authorizations, Authorization Group objects must be created on the Message VPN that match the authorization groups provisioned on the external server. These objects must be configured with the client profiles and ACL profiles that will be assigned to the clients that belong to those authorization groups. A newly created group is placed at the end of the group list which is the lowest priority.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: authorizationGroupName|x||| msgVpnName|x||| orderAfterAuthorizationGroupName||x|| orderBeforeAuthorizationGroupName||x||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.0.

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


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Authorization Group objects.  To use client authorization groups configured on an external server to provide client authorizations, Authorization Group objects must be created on the Message VPN that match the authorization groups provisioned on the external server. These objects must be configured with the client profiles and ACL profiles that will be assigned to the clients that belong to those authorization groups. A newly created group is placed at the end of the group list which is the lowest priority.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: authorizationGroupName|x||| msgVpnName|x||| orderAfterAuthorizationGroupName||x|| orderBeforeAuthorizationGroupName||x||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.0.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a Bridge object.  Bridges can be used to link two Message VPNs so that messages published to one Message VPN that match the topic subscriptions set for the bridge are also delivered to the linked Message VPN.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: bridgeName|x||| bridgeVirtualRouter|x||| msgVpnName|x||| remoteAuthenticationBasicPassword||x||x remoteAuthenticationClientCertContent||x||x remoteAuthenticationClientCertPassword||x||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.0.

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



 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeResponse**](MsgVpnBridgeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnBridgeRemoteMsgVpn**
> MsgVpnBridgeRemoteMsgVpnResponse GetMsgVpnBridgeRemoteMsgVpn(ctx, msgVpnName, bridgeName, bridgeVirtualRouter, remoteMsgVpnName, remoteMsgVpnLocation, remoteMsgVpnInterface, optional)
Get a Remote Message VPN object.

Get a Remote Message VPN object.  The Remote Message VPN is the Message VPN that the Bridge connects to.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: bridgeName|x||| bridgeVirtualRouter|x||| msgVpnName|x||| password||x||x remoteMsgVpnInterface|x||| remoteMsgVpnLocation|x||| remoteMsgVpnName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.0.

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






 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Remote Message VPN objects.  The Remote Message VPN is the Message VPN that the Bridge connects to.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: bridgeName|x||| bridgeVirtualRouter|x||| msgVpnName|x||| password||x||x remoteMsgVpnInterface|x||| remoteMsgVpnLocation|x||| remoteMsgVpnName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.0.

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



 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a Remote Subscription object.  A Remote Subscription is a topic subscription used by the Message VPN Bridge to attract messages from the remote message broker.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: bridgeName|x||| bridgeVirtualRouter|x||| msgVpnName|x||| remoteSubscriptionTopic|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.0.

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




 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Remote Subscription objects.  A Remote Subscription is a topic subscription used by the Message VPN Bridge to attract messages from the remote message broker.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: bridgeName|x||| bridgeVirtualRouter|x||| msgVpnName|x||| remoteSubscriptionTopic|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.0.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a Trusted Common Name object.  The Trusted Common Names for the Bridge are used by encrypted transports to verify the name in the certificate presented by the remote node. They must include the common name of the remote node's server certificate or client certificate, depending upon the initiator of the connection.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: bridgeName|x||x| bridgeVirtualRouter|x||x| msgVpnName|x||x| tlsTrustedCommonName|x||x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been deprecated since 2.18. Common Name validation has been replaced by Server Certificate Name validation.

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




 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Trusted Common Name objects.  The Trusted Common Names for the Bridge are used by encrypted transports to verify the name in the certificate presented by the remote node. They must include the common name of the remote node's server certificate or client certificate, depending upon the initiator of the connection.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: bridgeName|x||x| bridgeVirtualRouter|x||x| msgVpnName|x||x| tlsTrustedCommonName|x||x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been deprecated since 2.18. Common Name validation has been replaced by Server Certificate Name validation.

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



 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Bridge objects.  Bridges can be used to link two Message VPNs so that messages published to one Message VPN that match the topic subscriptions set for the bridge are also delivered to the linked Message VPN.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: bridgeName|x||| bridgeVirtualRouter|x||| msgVpnName|x||| remoteAuthenticationBasicPassword||x||x remoteAuthenticationClientCertContent||x||x remoteAuthenticationClientCertPassword||x||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.0.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a Certificate Matching Rule object.  A Cert Matching Rule is a collection of conditions and attribute filters that all have to be satisfied for certificate to be acceptable as authentication for a given username.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| ruleName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.27.

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


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a Certificate Matching Rule Attribute Filter object.  A Cert Matching Rule Attribute Filter compares a username attribute to a string.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: filterName|x||| msgVpnName|x||| ruleName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.28.

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



 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Certificate Matching Rule Attribute Filter objects.  A Cert Matching Rule Attribute Filter compares a username attribute to a string.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: filterName|x||| msgVpnName|x||| ruleName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.28.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a Certificate Matching Rule Condition object.  A Cert Matching Rule Condition compares data extracted from a certificate to a username attribute or an expression.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| ruleName|x||| source|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.27.

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



 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Certificate Matching Rule Condition objects.  A Cert Matching Rule Condition compares data extracted from a certificate to a username attribute or an expression.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| ruleName|x||| source|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.27.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Certificate Matching Rule objects.  A Cert Matching Rule is a collection of conditions and attribute filters that all have to be satisfied for certificate to be acceptable as authentication for a given username.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| ruleName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.27.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

# **GetMsgVpnClientProfile**
> MsgVpnClientProfileResponse GetMsgVpnClientProfile(ctx, msgVpnName, clientProfileName, optional)
Get a Client Profile object.

Get a Client Profile object.  Client Profiles are used to assign common configuration properties to clients that have been successfully authorized.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: allowCutThroughForwardingEnabled|||x| apiQueueManagementCopyFromOnCreateName|||x| apiTopicEndpointManagementCopyFromOnCreateName|||x| clientProfileName|x||| msgVpnName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.0.

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


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Client Profile objects.  Client Profiles are used to assign common configuration properties to clients that have been successfully authorized.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: allowCutThroughForwardingEnabled|||x| apiQueueManagementCopyFromOnCreateName|||x| apiTopicEndpointManagementCopyFromOnCreateName|||x| clientProfileName|x||| msgVpnName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.0.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

# **GetMsgVpnClientUsername**
> MsgVpnClientUsernameResponse GetMsgVpnClientUsername(ctx, msgVpnName, clientUsername, optional)
Get a Client Username object.

Get a Client Username object.  A client is only authorized to connect to a Message VPN that is associated with a Client Username that the client has been assigned.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: clientUsername|x||| msgVpnName|x||| password||x||x    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.0.

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


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a Client Username Attribute object.  A ClientUsername Attribute is a key+value pair that can be used to locate a client username, for example when using client certificate mapping.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: attributeName|x||| attributeValue|x||| clientUsername|x||| msgVpnName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.27.

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




 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Client Username Attribute objects.  A ClientUsername Attribute is a key+value pair that can be used to locate a client username, for example when using client certificate mapping.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: attributeName|x||| attributeValue|x||| clientUsername|x||| msgVpnName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.27.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Client Username objects.  A client is only authorized to connect to a Message VPN that is associated with a Client Username that the client has been assigned.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: clientUsername|x||| msgVpnName|x||| password||x||x    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.0.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

# **GetMsgVpnDistributedCache**
> MsgVpnDistributedCacheResponse GetMsgVpnDistributedCache(ctx, msgVpnName, cacheName, optional)
Get a Distributed Cache object.

Get a Distributed Cache object.  A Distributed Cache is a collection of one or more Cache Clusters that belong to the same Message VPN. Each Cache Cluster in a Distributed Cache is configured to subscribe to a different set of topics. This effectively divides up the configured topic space, to provide scaling to very large topic spaces or very high cached message throughput.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: cacheName|x||| msgVpnName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

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


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a Cache Cluster object.  A Cache Cluster is a collection of one or more Cache Instances that subscribe to exactly the same topics. Cache Instances are grouped together in a Cache Cluster for the purpose of fault tolerance and load balancing. As published messages are received, the message broker message bus sends these live data messages to the Cache Instances in the Cache Cluster. This enables client cache requests to be served by any of Cache Instances in the Cache Cluster.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: cacheName|x||| clusterName|x||| msgVpnName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

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



 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a Home Cache Cluster object.  A Home Cache Cluster is a Cache Cluster that is the \"definitive\" Cache Cluster for a given topic in the context of the Global Caching feature.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: cacheName|x||| clusterName|x||| homeClusterName|x||| msgVpnName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

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




 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a Topic Prefix object.  A Topic Prefix is a prefix for a global topic that is available from the containing Home Cache Cluster.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: cacheName|x||| clusterName|x||| homeClusterName|x||| msgVpnName|x||| topicPrefix|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

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





 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Topic Prefix objects.  A Topic Prefix is a prefix for a global topic that is available from the containing Home Cache Cluster.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: cacheName|x||| clusterName|x||| homeClusterName|x||| msgVpnName|x||| topicPrefix|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Home Cache Cluster objects.  A Home Cache Cluster is a Cache Cluster that is the \"definitive\" Cache Cluster for a given topic in the context of the Global Caching feature.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: cacheName|x||| clusterName|x||| homeClusterName|x||| msgVpnName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a Cache Instance object.  A Cache Instance is a single Cache process that belongs to a single Cache Cluster. A Cache Instance object provisioned on the broker is used to disseminate configuration information to the Cache process. Cache Instances listen for and cache live data messages that match the topic subscriptions configured for their parent Cache Cluster.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: cacheName|x||| clusterName|x||| instanceName|x||| msgVpnName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

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




 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Cache Instance objects.  A Cache Instance is a single Cache process that belongs to a single Cache Cluster. A Cache Instance object provisioned on the broker is used to disseminate configuration information to the Cache process. Cache Instances listen for and cache live data messages that match the topic subscriptions configured for their parent Cache Cluster.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: cacheName|x||| clusterName|x||| instanceName|x||| msgVpnName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a Topic object.  The Cache Instances that belong to the containing Cache Cluster will cache any messages published to topics that match a Topic Subscription.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: cacheName|x||| clusterName|x||| msgVpnName|x||| topic|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

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




 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Topic objects.  The Cache Instances that belong to the containing Cache Cluster will cache any messages published to topics that match a Topic Subscription.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: cacheName|x||| clusterName|x||| msgVpnName|x||| topic|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Cache Cluster objects.  A Cache Cluster is a collection of one or more Cache Instances that subscribe to exactly the same topics. Cache Instances are grouped together in a Cache Cluster for the purpose of fault tolerance and load balancing. As published messages are received, the message broker message bus sends these live data messages to the Cache Instances in the Cache Cluster. This enables client cache requests to be served by any of Cache Instances in the Cache Cluster.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: cacheName|x||| clusterName|x||| msgVpnName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Distributed Cache objects.  A Distributed Cache is a collection of one or more Cache Clusters that belong to the same Message VPN. Each Cache Cluster in a Distributed Cache is configured to subscribe to a different set of topics. This effectively divides up the configured topic space, to provide scaling to very large topic spaces or very high cached message throughput.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: cacheName|x||| msgVpnName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a DMR Bridge object.  A DMR Bridge is required to establish a data channel over a corresponding external link to the remote node for a given Message VPN. Each DMR Bridge identifies which external link the Message VPN should use, and what the name of the equivalent Message VPN at the remote node is.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| remoteNodeName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

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
 **optional** | ***MsgVpnApiGetMsgVpnDmrBridgesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnDmrBridgesOpts struct
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

# **GetMsgVpnJndiConnectionFactories**
> MsgVpnJndiConnectionFactoriesResponse GetMsgVpnJndiConnectionFactories(ctx, msgVpnName, optional)
Get a list of JNDI Connection Factory objects.

Get a list of JNDI Connection Factory objects.  The message broker provides an internal JNDI store for provisioned Connection Factory objects that clients can access through JNDI lookups.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: connectionFactoryName|x||| msgVpnName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.2.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a JNDI Connection Factory object.  The message broker provides an internal JNDI store for provisioned Connection Factory objects that clients can access through JNDI lookups.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: connectionFactoryName|x||| msgVpnName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.2.

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


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a JNDI Queue object.  The message broker provides an internal JNDI store for provisioned Queue objects that clients can access through JNDI lookups.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| queueName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.2.

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


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of JNDI Queue objects.  The message broker provides an internal JNDI store for provisioned Queue objects that clients can access through JNDI lookups.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| queueName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.2.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a JNDI Topic object.  The message broker provides an internal JNDI store for provisioned Topic objects that clients can access through JNDI lookups.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| topicName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.2.

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


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of JNDI Topic objects.  The message broker provides an internal JNDI store for provisioned Topic objects that clients can access through JNDI lookups.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| topicName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.2.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get an MQTT Retain Cache object.  Using MQTT retained messages allows publishing MQTT clients to indicate that a message must be stored for later delivery to subscribing clients when those subscribing clients add subscriptions matching the retained message's topic. An MQTT Retain Cache processes all retained messages for a Message VPN.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: cacheName|x||| msgVpnName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

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


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of MQTT Retain Cache objects.  Using MQTT retained messages allows publishing MQTT clients to indicate that a message must be stored for later delivery to subscribing clients when those subscribing clients add subscriptions matching the retained message's topic. An MQTT Retain Cache processes all retained messages for a Message VPN.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: cacheName|x||| msgVpnName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.11.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get an MQTT Session object.  An MQTT Session object is a virtual representation of an MQTT client connection. An MQTT session holds the state of an MQTT client (that is, it is used to contain a client's QoS 0 and QoS 1 subscription sets and any undelivered QoS 1 messages).   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: mqttSessionClientId|x||| mqttSessionVirtualRouter|x||| msgVpnName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.1.

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



 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a Subscription object.  An MQTT session contains a client's QoS 0 and QoS 1 subscription sets. On creation, a subscription defaults to QoS 0.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: mqttSessionClientId|x||| mqttSessionVirtualRouter|x||| msgVpnName|x||| subscriptionTopic|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.1.

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




 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Subscription objects.  An MQTT session contains a client's QoS 0 and QoS 1 subscription sets. On creation, a subscription defaults to QoS 0.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: mqttSessionClientId|x||| mqttSessionVirtualRouter|x||| msgVpnName|x||| subscriptionTopic|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.1.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of MQTT Session objects.  An MQTT Session object is a virtual representation of an MQTT client connection. An MQTT session holds the state of an MQTT client (that is, it is used to contain a client's QoS 0 and QoS 1 subscription sets and any undelivered QoS 1 messages).   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: mqttSessionClientId|x||| mqttSessionVirtualRouter|x||| msgVpnName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.1.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a Queue object.  A Queue acts as both a destination that clients can publish messages to, and as an endpoint that clients can bind consumers to and consume messages from.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| queueName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.0.

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


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueueResponse**](MsgVpnQueueResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnQueueSubscription**
> MsgVpnQueueSubscriptionResponse GetMsgVpnQueueSubscription(ctx, msgVpnName, queueName, subscriptionTopic, optional)
Get a Queue Subscription object.

Get a Queue Subscription object.  One or more Queue Subscriptions can be added to a durable queue so that Guaranteed messages published to matching topics are also delivered to and spooled by the queue.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| queueName|x||| subscriptionTopic|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.0.

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



 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Queue Subscription objects.  One or more Queue Subscriptions can be added to a durable queue so that Guaranteed messages published to matching topics are also delivered to and spooled by the queue.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| queueName|x||| subscriptionTopic|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.0.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a Queue Template object.  A Queue Template provides a mechanism for specifying the initial state for client created queues.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| queueTemplateName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.14.

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


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Queue Template objects.  A Queue Template provides a mechanism for specifying the initial state for client created queues.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| queueTemplateName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.14.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

# **GetMsgVpnQueues**
> MsgVpnQueuesResponse GetMsgVpnQueues(ctx, msgVpnName, optional)
Get a list of Queue objects.

Get a list of Queue objects.  A Queue acts as both a destination that clients can publish messages to, and as an endpoint that clients can bind consumers to and consume messages from.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| queueName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.0.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a Replay Log object.  When the Message Replay feature is enabled, message brokers store persistent messages in a Replay Log. These messages are kept until the log is full, after which the oldest messages are removed to free up space for new messages.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| replayLogName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.10.

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


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnReplayLogResponse**](MsgVpnReplayLogResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnReplayLogTopicFilterSubscription**
> MsgVpnReplayLogTopicFilterSubscriptionResponse GetMsgVpnReplayLogTopicFilterSubscription(ctx, msgVpnName, replayLogName, topicFilterSubscription, optional)
Get a Topic Filter Subscription object.

Get a Topic Filter Subscription object.  One or more Subscriptions can be added to a replay-log so that only guaranteed messages published to matching topics are stored in the Replay Log.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| replayLogName|x||| topicFilterSubscription|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.27.

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



 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Topic Filter Subscription objects.  One or more Subscriptions can be added to a replay-log so that only guaranteed messages published to matching topics are stored in the Replay Log.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| replayLogName|x||| topicFilterSubscription|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.27.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Replay Log objects.  When the Message Replay feature is enabled, message brokers store persistent messages in a Replay Log. These messages are kept until the log is full, after which the oldest messages are removed to free up space for new messages.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| replayLogName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.10.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a Replicated Topic object.  To indicate which messages should be replicated between the active and standby site, a Replicated Topic subscription must be configured on a Message VPN. If a published message matches both a replicated topic and an endpoint on the active site, then the message is replicated to the standby site.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| replicatedTopic|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.1.

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


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Replicated Topic objects.  To indicate which messages should be replicated between the active and standby site, a Replicated Topic subscription must be configured on a Message VPN. If a published message matches both a replicated topic and an endpoint on the active site, then the message is replicated to the standby site.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| replicatedTopic|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.1.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a REST Delivery Point object.  A REST Delivery Point manages delivery of messages from queues to a named list of REST Consumers.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| restDeliveryPointName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.0.

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


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a Queue Binding object.  A Queue Binding for a REST Delivery Point attracts messages to be delivered to REST consumers. If the queue does not exist it can be created subsequently, and once the queue is operational the broker performs the queue binding. Removing the queue binding does not delete the queue itself. Similarly, removing the queue does not remove the queue binding, which fails until the queue is recreated or the queue binding is deleted.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| queueBindingName|x||| restDeliveryPointName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.0.

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



 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a Protected Request Header object.  A protected request header to be added to the HTTP request. Unlike a non-protected request header, the header value cannot be displayed after it is set.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: headerName|x||| headerValue||x||x msgVpnName|x||| queueBindingName|x||| restDeliveryPointName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.30.

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




 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Protected Request Header objects.  A protected request header to be added to the HTTP request. Unlike a non-protected request header, the header value cannot be displayed after it is set.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: headerName|x||| headerValue||x||x msgVpnName|x||| queueBindingName|x||| restDeliveryPointName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.30.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a Request Header object.  A request header to be added to the HTTP request.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: headerName|x||| msgVpnName|x||| queueBindingName|x||| restDeliveryPointName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.23.

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




 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Request Header objects.  A request header to be added to the HTTP request.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: headerName|x||| msgVpnName|x||| queueBindingName|x||| restDeliveryPointName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.23.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Queue Binding objects.  A Queue Binding for a REST Delivery Point attracts messages to be delivered to REST consumers. If the queue does not exist it can be created subsequently, and once the queue is operational the broker performs the queue binding. Removing the queue binding does not delete the queue itself. Similarly, removing the queue does not remove the queue binding, which fails until the queue is recreated or the queue binding is deleted.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| queueBindingName|x||| restDeliveryPointName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.0.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a REST Consumer object.  REST Consumer objects establish HTTP connectivity to REST consumer applications who wish to receive messages from a broker.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: authenticationAwsSecretAccessKey||x||x authenticationClientCertContent||x||x authenticationClientCertPassword||x|| authenticationHttpBasicPassword||x||x authenticationHttpHeaderValue||x||x authenticationOauthClientSecret||x||x authenticationOauthJwtSecretKey||x||x msgVpnName|x||| restConsumerName|x||| restDeliveryPointName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.0.

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



 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a Claim object.  A Claim is added to the JWT sent to the OAuth token request endpoint.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| oauthJwtClaimName|x||| restConsumerName|x||| restDeliveryPointName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.21.

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




 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Claim objects.  A Claim is added to the JWT sent to the OAuth token request endpoint.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| oauthJwtClaimName|x||| restConsumerName|x||| restDeliveryPointName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.21.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a Trusted Common Name object.  The Trusted Common Names for the REST Consumer are used by encrypted transports to verify the name in the certificate presented by the remote REST consumer. They must include the common name of the remote REST consumer's server certificate.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||x| restConsumerName|x||x| restDeliveryPointName|x||x| tlsTrustedCommonName|x||x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been deprecated since (will be deprecated in next SEMP version). Common Name validation has been replaced by Server Certificate Name validation.

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




 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Trusted Common Name objects.  The Trusted Common Names for the REST Consumer are used by encrypted transports to verify the name in the certificate presented by the remote REST consumer. They must include the common name of the remote REST consumer's server certificate.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||x| restConsumerName|x||x| restDeliveryPointName|x||x| tlsTrustedCommonName|x||x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been deprecated since (will be deprecated in next SEMP version). Common Name validation has been replaced by Server Certificate Name validation.

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



 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of REST Consumer objects.  REST Consumer objects establish HTTP connectivity to REST consumer applications who wish to receive messages from a broker.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: authenticationAwsSecretAccessKey||x||x authenticationClientCertContent||x||x authenticationClientCertPassword||x|| authenticationHttpBasicPassword||x||x authenticationHttpHeaderValue||x||x authenticationOauthClientSecret||x||x authenticationOauthJwtSecretKey||x||x msgVpnName|x||| restConsumerName|x||| restDeliveryPointName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.0.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of REST Delivery Point objects.  A REST Delivery Point manages delivery of messages from queues to a named list of REST Consumers.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| restDeliveryPointName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.0.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

# **GetMsgVpnSequencedTopic**
> MsgVpnSequencedTopicResponse GetMsgVpnSequencedTopic(ctx, msgVpnName, sequencedTopic, optional)
Get a Sequenced Topic object.

Get a Sequenced Topic object.  A Sequenced Topic is a topic subscription for which any matching messages received on the Message VPN are assigned a sequence number that is monotonically increased by a value of one per message.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| sequencedTopic|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **sequencedTopic** | **string**| Topic for applying sequence numbers. | 
 **optional** | ***MsgVpnApiGetMsgVpnSequencedTopicOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnSequencedTopicOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnSequencedTopicResponse**](MsgVpnSequencedTopicResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnSequencedTopics**
> MsgVpnSequencedTopicsResponse GetMsgVpnSequencedTopics(ctx, msgVpnName, optional)
Get a list of Sequenced Topic objects.

Get a list of Sequenced Topic objects.  A Sequenced Topic is a topic subscription for which any matching messages received on the Message VPN are assigned a sequence number that is monotonically increased by a value of one per message.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| sequencedTopic|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MsgVpnApiGetMsgVpnSequencedTopicsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiGetMsgVpnSequencedTopicsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnSequencedTopicsResponse**](MsgVpnSequencedTopicsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnTelemetryProfile**
> MsgVpnTelemetryProfileResponse GetMsgVpnTelemetryProfile(ctx, msgVpnName, telemetryProfileName, optional)
Get a Telemetry Profile object.

Get a Telemetry Profile object.  Using the Telemetry Profile allows trace spans to be generated as messages are processed by the broker. The generated spans are stored persistently on the broker and may be consumed by the Solace receiver component of an OpenTelemetry Collector.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| telemetryProfileName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.31.

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


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a Receiver ACL Connect Exception object.  A Receiver ACL Connect Exception is an exception to the default action to take when a receiver connects to the broker. Exceptions must be expressed as an IP address/netmask in CIDR form.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| receiverAclConnectExceptionAddress|x||| telemetryProfileName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.31.

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



 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Receiver ACL Connect Exception objects.  A Receiver ACL Connect Exception is an exception to the default action to take when a receiver connects to the broker. Exceptions must be expressed as an IP address/netmask in CIDR form.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| receiverAclConnectExceptionAddress|x||| telemetryProfileName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.31.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a Trace Filter object.  A Trace Filter controls which messages received by the broker will be traced. If an incoming message matches an enabled tracing filter's subscription, the message will be traced as it passes through the broker.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| telemetryProfileName|x||| traceFilterName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.31.

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



 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a Telemetry Trace Filter Subscription object.  Trace filter subscriptions control which messages will be attracted by the tracing filter.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| subscription|x||| subscriptionSyntax|x||| telemetryProfileName|x||| traceFilterName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.31.

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





 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Telemetry Trace Filter Subscription objects.  Trace filter subscriptions control which messages will be attracted by the tracing filter.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| subscription|x||| subscriptionSyntax|x||| telemetryProfileName|x||| traceFilterName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.31.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Trace Filter objects.  A Trace Filter controls which messages received by the broker will be traced. If an incoming message matches an enabled tracing filter's subscription, the message will be traced as it passes through the broker.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| telemetryProfileName|x||| traceFilterName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.31.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Telemetry Profile objects.  Using the Telemetry Profile allows trace spans to be generated as messages are processed by the broker. The generated spans are stored persistently on the broker and may be consumed by the Solace receiver component of an OpenTelemetry Collector.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| telemetryProfileName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.31.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a Topic Endpoint object.  A Topic Endpoint attracts messages published to a topic for which the Topic Endpoint has a matching topic subscription. The topic subscription for the Topic Endpoint is specified in the client request to bind a Flow to that Topic Endpoint. Queues are significantly more flexible than Topic Endpoints and are the recommended approach for most applications. The use of Topic Endpoints should be restricted to JMS applications.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| topicEndpointName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.1.

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


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointResponse**](MsgVpnTopicEndpointResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnTopicEndpointTemplate**
> MsgVpnTopicEndpointTemplateResponse GetMsgVpnTopicEndpointTemplate(ctx, msgVpnName, topicEndpointTemplateName, optional)
Get a Topic Endpoint Template object.

Get a Topic Endpoint Template object.  A Topic Endpoint Template provides a mechanism for specifying the initial state for client created topic endpoints.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| topicEndpointTemplateName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.14.

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


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

Get a list of Topic Endpoint Template objects.  A Topic Endpoint Template provides a mechanism for specifying the initial state for client created topic endpoints.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| topicEndpointTemplateName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.14.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

# **GetMsgVpnTopicEndpoints**
> MsgVpnTopicEndpointsResponse GetMsgVpnTopicEndpoints(ctx, msgVpnName, optional)
Get a list of Topic Endpoint objects.

Get a list of Topic Endpoint objects.  A Topic Endpoint attracts messages published to a topic for which the Topic Endpoint has a matching topic subscription. The topic subscription for the Topic Endpoint is specified in the client request to bind a Flow to that Topic Endpoint. Queues are significantly more flexible than Topic Endpoints and are the recommended approach for most applications. The use of Topic Endpoints should be restricted to JMS applications.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| topicEndpointName|x|||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.1.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

# **GetMsgVpns**
> MsgVpnsResponse GetMsgVpns(ctx, optional)
Get a list of Message VPN objects.

Get a list of Message VPN objects.  Message VPNs (Virtual Private Networks) allow for the segregation of topic space and clients. They also group clients connecting to a network of message brokers, such that messages published within a particular group are only visible to that group's clients.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: authenticationOauthDefaultProviderName|||x| bridgingTlsServerCertEnforceTrustedCommonNameEnabled|||x| distributedCacheManagementEnabled|||x| msgVpnName|x||| replicationBridgeAuthenticationBasicPassword||x||x replicationBridgeAuthenticationClientCertContent||x||x replicationBridgeAuthenticationClientCertPassword||x|| replicationEnabledQueueBehavior||x|| restTlsServerCertEnforceTrustedCommonNameEnabled|||x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.0.

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
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

# **ReplaceMsgVpn**
> MsgVpnResponse ReplaceMsgVpn(ctx, body, msgVpnName, optional)
Replace a Message VPN object.

Replace a Message VPN object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  Message VPNs (Virtual Private Networks) allow for the segregation of topic space and clients. They also group clients connecting to a network of message brokers, such that messages published within a particular group are only visible to that group's clients.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- authenticationOauthDefaultProviderName|||||||x| bridgingTlsServerCertEnforceTrustedCommonNameEnabled|||||||x| distributedCacheManagementEnabled|||||||x| msgVpnName|x||x||||| replicationBridgeAuthenticationBasicPassword||||x||||x replicationBridgeAuthenticationClientCertContent||||x||||x replicationBridgeAuthenticationClientCertPassword||||x|||| replicationEnabledQueueBehavior||||x|||| restTlsServerCertEnforceTrustedCommonNameEnabled|||||||x| serviceAmqpPlainTextListenPort||||||x|| serviceAmqpTlsListenPort||||||x|| serviceMqttPlainTextListenPort||||||x|| serviceMqttTlsListenPort||||||x|| serviceMqttTlsWebSocketListenPort||||||x|| serviceMqttWebSocketListenPort||||||x|| serviceRestIncomingPlainTextListenPort||||||x|| serviceRestIncomingTlsListenPort||||||x||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThreshold|clearPercent|setPercent|clearValue, setValue EventThreshold|clearValue|setValue|clearPercent, setPercent EventThreshold|setPercent|clearPercent|clearValue, setValue EventThreshold|setValue|clearValue|clearPercent, setPercent EventThresholdByValue|clearValue|setValue| EventThresholdByValue|setValue|clearValue| MsgVpn|authenticationBasicProfileName|authenticationBasicType| MsgVpn|authorizationProfileName|authorizationType| MsgVpn|eventPublishTopicFormatMqttEnabled|eventPublishTopicFormatSmfEnabled| MsgVpn|eventPublishTopicFormatSmfEnabled|eventPublishTopicFormatMqttEnabled| MsgVpn|replicationBridgeAuthenticationBasicClientUsername|replicationBridgeAuthenticationBasicPassword| MsgVpn|replicationBridgeAuthenticationBasicPassword|replicationBridgeAuthenticationBasicClientUsername| MsgVpn|replicationBridgeAuthenticationClientCertPassword|replicationBridgeAuthenticationClientCertContent| MsgVpn|replicationEnabledQueueBehavior|replicationEnabled|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation. Requests which include the following attributes require greater access scope/level:   Attribute|Access Scope/Level :---|:---: alias|global/read-write authenticationBasicEnabled|global/read-write authenticationBasicProfileName|global/read-write authenticationBasicRadiusDomain|global/read-write authenticationBasicType|global/read-write authenticationClientCertAllowApiProvidedUsernameEnabled|global/read-write authenticationClientCertCertificateMatchingRulesEnabled|global/read-write authenticationClientCertEnabled|global/read-write authenticationClientCertMaxChainDepth|global/read-write authenticationClientCertRevocationCheckMode|global/read-write authenticationClientCertUsernameSource|global/read-write authenticationClientCertValidateDateEnabled|global/read-write authenticationKerberosAllowApiProvidedUsernameEnabled|global/read-write authenticationKerberosEnabled|global/read-write authenticationOauthEnabled|global/read-write bridgingTlsServerCertEnforceTrustedCommonNameEnabled|global/read-write bridgingTlsServerCertMaxChainDepth|global/read-write bridgingTlsServerCertValidateDateEnabled|global/read-write bridgingTlsServerCertValidateNameEnabled|global/read-write dmrEnabled|global/read-write exportSubscriptionsEnabled|global/read-write maxConnectionCount|global/read-write maxEgressFlowCount|global/read-write maxEndpointCount|global/read-write maxIngressFlowCount|global/read-write maxMsgSpoolUsage|global/read-write maxSubscriptionCount|global/read-write maxTransactedSessionCount|global/read-write maxTransactionCount|global/read-write mqttRetainMaxMemory|global/read-write replicationBridgeAuthenticationBasicClientUsername|global/read-write replicationBridgeAuthenticationBasicPassword|global/read-write replicationBridgeAuthenticationClientCertContent|global/read-write replicationBridgeAuthenticationClientCertPassword|global/read-write replicationBridgeAuthenticationScheme|global/read-write replicationBridgeCompressedDataEnabled|global/read-write replicationBridgeEgressFlowWindowSize|global/read-write replicationBridgeRetryDelay|global/read-write replicationBridgeTlsEnabled|global/read-write replicationBridgeUnidirectionalClientProfileName|global/read-write replicationEnabled|global/read-write replicationEnabledQueueBehavior|global/read-write replicationQueueMaxMsgSpoolUsage|global/read-write replicationRole|global/read-write restTlsServerCertEnforceTrustedCommonNameEnabled|global/read-write restTlsServerCertMaxChainDepth|global/read-write restTlsServerCertValidateDateEnabled|global/read-write restTlsServerCertValidateNameEnabled|global/read-write sempOverMsgBusAdminClientEnabled|global/read-write sempOverMsgBusAdminDistributedCacheEnabled|global/read-write sempOverMsgBusAdminEnabled|global/read-write sempOverMsgBusEnabled|global/read-write sempOverMsgBusShowEnabled|global/read-write serviceAmqpMaxConnectionCount|global/read-write serviceAmqpPlainTextListenPort|global/read-write serviceAmqpTlsListenPort|global/read-write serviceMqttMaxConnectionCount|global/read-write serviceMqttPlainTextListenPort|global/read-write serviceMqttTlsListenPort|global/read-write serviceMqttTlsWebSocketListenPort|global/read-write serviceMqttWebSocketListenPort|global/read-write serviceRestIncomingMaxConnectionCount|global/read-write serviceRestIncomingPlainTextListenPort|global/read-write serviceRestIncomingTlsListenPort|global/read-write serviceRestOutgoingMaxConnectionCount|global/read-write serviceSmfMaxConnectionCount|global/read-write serviceWebMaxConnectionCount|global/read-write    This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpn**](MsgVpn.md)| The Message VPN object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MsgVpnApiReplaceMsgVpnOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiReplaceMsgVpnOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnResponse**](MsgVpnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnAclProfile**
> MsgVpnAclProfileResponse ReplaceMsgVpnAclProfile(ctx, body, msgVpnName, aclProfileName, optional)
Replace an ACL Profile object.

Replace an ACL Profile object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  An ACL Profile controls whether an authenticated client is permitted to establish a connection with the message broker or permitted to publish and subscribe to specific topics.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- aclProfileName|x||x||||| msgVpnName|x||x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnAclProfile**](MsgVpnAclProfile.md)| The ACL Profile object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
 **optional** | ***MsgVpnApiReplaceMsgVpnAclProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiReplaceMsgVpnAclProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileResponse**](MsgVpnAclProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnAuthenticationOauthProfile**
> MsgVpnAuthenticationOauthProfileResponse ReplaceMsgVpnAuthenticationOauthProfile(ctx, body, msgVpnName, oauthProfileName, optional)
Replace an OAuth Profile object.

Replace an OAuth Profile object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  OAuth profiles specify how to securely authenticate to an OAuth provider.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- clientSecret||||x||||x msgVpnName|x||x||||| oauthProfileName|x||x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.25.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnAuthenticationOauthProfile**](MsgVpnAuthenticationOauthProfile.md)| The OAuth Profile object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
 **optional** | ***MsgVpnApiReplaceMsgVpnAuthenticationOauthProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiReplaceMsgVpnAuthenticationOauthProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthenticationOauthProfileResponse**](MsgVpnAuthenticationOauthProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnAuthenticationOauthProvider**
> MsgVpnAuthenticationOauthProviderResponse ReplaceMsgVpnAuthenticationOauthProvider(ctx, body, msgVpnName, oauthProviderName, optional)
Replace an OAuth Provider object.

Replace an OAuth Provider object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  OAuth Providers contain information about the issuer of an OAuth token that is needed to validate the token and derive a client username from it.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- audienceClaimName|||||||x| audienceClaimSource|||||||x| audienceClaimValue|||||||x| audienceValidationEnabled|||||||x| authorizationGroupClaimName|||||||x| authorizationGroupClaimSource|||||||x| authorizationGroupEnabled|||||||x| disconnectOnTokenExpirationEnabled|||||||x| enabled|||||||x| jwksRefreshInterval|||||||x| jwksUri|||||||x| msgVpnName|x||x||||x| oauthProviderName|x||x||||x| tokenIgnoreTimeLimitsEnabled|||||||x| tokenIntrospectionParameterName|||||||x| tokenIntrospectionPassword||||x|||x|x tokenIntrospectionTimeout|||||||x| tokenIntrospectionUri|||||||x| tokenIntrospectionUsername|||||||x| usernameClaimName|||||||x| usernameClaimSource|||||||x| usernameValidateEnabled|||||||x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been deprecated since 2.25. Replaced by authenticationOauthProfiles.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnAuthenticationOauthProvider**](MsgVpnAuthenticationOauthProvider.md)| The OAuth Provider object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **oauthProviderName** | **string**| The name of the OAuth Provider. | 
 **optional** | ***MsgVpnApiReplaceMsgVpnAuthenticationOauthProviderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiReplaceMsgVpnAuthenticationOauthProviderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthenticationOauthProviderResponse**](MsgVpnAuthenticationOauthProviderResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnAuthorizationGroup**
> MsgVpnAuthorizationGroupResponse ReplaceMsgVpnAuthorizationGroup(ctx, body, msgVpnName, authorizationGroupName, optional)
Replace an Authorization Group object.

Replace an Authorization Group object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  To use client authorization groups configured on an external server to provide client authorizations, Authorization Group objects must be created on the Message VPN that match the authorization groups provisioned on the external server. These objects must be configured with the client profiles and ACL profiles that will be assigned to the clients that belong to those authorization groups. A newly created group is placed at the end of the group list which is the lowest priority.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- aclProfileName||||||x|| authorizationGroupName|x||x||||| clientProfileName||||||x|| msgVpnName|x||x||||| orderAfterAuthorizationGroupName||||x|||| orderBeforeAuthorizationGroupName||||x||||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- MsgVpnAuthorizationGroup|orderAfterAuthorizationGroupName||orderBeforeAuthorizationGroupName MsgVpnAuthorizationGroup|orderBeforeAuthorizationGroupName||orderAfterAuthorizationGroupName    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnAuthorizationGroup**](MsgVpnAuthorizationGroup.md)| The Authorization Group object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **authorizationGroupName** | **string**| The name of the Authorization Group. For LDAP groups, special care is needed if the group name contains special characters such as &#x27;#&#x27;, &#x27;+&#x27;, &#x27;;&#x27;, &#x27;&#x3D;&#x27; as the value of the group name returned from the LDAP server might prepend those characters with &#x27;\\&#x27;. For example a group name called &#x27;test#,lab,com&#x27; will be returned from the LDAP server as &#x27;test\\#,lab,com&#x27;. | 
 **optional** | ***MsgVpnApiReplaceMsgVpnAuthorizationGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiReplaceMsgVpnAuthorizationGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthorizationGroupResponse**](MsgVpnAuthorizationGroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnBridge**
> MsgVpnBridgeResponse ReplaceMsgVpnBridge(ctx, body, msgVpnName, bridgeName, bridgeVirtualRouter, optional)
Replace a Bridge object.

Replace a Bridge object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  Bridges can be used to link two Message VPNs so that messages published to one Message VPN that match the topic subscriptions set for the bridge are also delivered to the linked Message VPN.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- bridgeName|x||x||||| bridgeVirtualRouter|x||x||||| maxTtl||||||x|| msgVpnName|x||x||||| remoteAuthenticationBasicClientUsername||||||x|| remoteAuthenticationBasicPassword||||x||x||x remoteAuthenticationClientCertContent||||x||x||x remoteAuthenticationClientCertPassword||||x||x|| remoteAuthenticationScheme||||||x|| remoteDeliverToOnePriority||||||x||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- MsgVpnBridge|remoteAuthenticationBasicClientUsername|remoteAuthenticationBasicPassword| MsgVpnBridge|remoteAuthenticationBasicPassword|remoteAuthenticationBasicClientUsername| MsgVpnBridge|remoteAuthenticationClientCertPassword|remoteAuthenticationClientCertContent|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnBridge**](MsgVpnBridge.md)| The Bridge object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **bridgeName** | **string**| The name of the Bridge. | 
  **bridgeVirtualRouter** | **string**| The virtual router of the Bridge. | 
 **optional** | ***MsgVpnApiReplaceMsgVpnBridgeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiReplaceMsgVpnBridgeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeResponse**](MsgVpnBridgeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnBridgeRemoteMsgVpn**
> MsgVpnBridgeRemoteMsgVpnResponse ReplaceMsgVpnBridgeRemoteMsgVpn(ctx, body, msgVpnName, bridgeName, bridgeVirtualRouter, remoteMsgVpnName, remoteMsgVpnLocation, remoteMsgVpnInterface, optional)
Replace a Remote Message VPN object.

Replace a Remote Message VPN object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  The Remote Message VPN is the Message VPN that the Bridge connects to.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- bridgeName|x||x||||| bridgeVirtualRouter|x||x||||| clientUsername||||||x|| compressedDataEnabled||||||x|| egressFlowWindowSize||||||x|| msgVpnName|x||x||||| password||||x||x||x remoteMsgVpnInterface|x||x||||| remoteMsgVpnLocation|x||x||||| remoteMsgVpnName|x||x||||| tlsEnabled||||||x||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- MsgVpnBridgeRemoteMsgVpn|clientUsername|password| MsgVpnBridgeRemoteMsgVpn|password|clientUsername|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnBridgeRemoteMsgVpn**](MsgVpnBridgeRemoteMsgVpn.md)| The Remote Message VPN object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **bridgeName** | **string**| The name of the Bridge. | 
  **bridgeVirtualRouter** | **string**| The virtual router of the Bridge. | 
  **remoteMsgVpnName** | **string**| The name of the remote Message VPN. | 
  **remoteMsgVpnLocation** | **string**| The location of the remote Message VPN as either an FQDN with port, IP address with port, or virtual router name (starting with \&quot;v:\&quot;). | 
  **remoteMsgVpnInterface** | **string**| The physical interface on the local Message VPN host for connecting to the remote Message VPN. By default, an interface is chosen automatically (recommended), but if specified, &#x60;remoteMsgVpnLocation&#x60; must not be a virtual router name. | 
 **optional** | ***MsgVpnApiReplaceMsgVpnBridgeRemoteMsgVpnOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiReplaceMsgVpnBridgeRemoteMsgVpnOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------







 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeRemoteMsgVpnResponse**](MsgVpnBridgeRemoteMsgVpnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnCertMatchingRule**
> MsgVpnCertMatchingRuleResponse ReplaceMsgVpnCertMatchingRule(ctx, body, msgVpnName, ruleName, optional)
Replace a Certificate Matching Rule object.

Replace a Certificate Matching Rule object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  A Cert Matching Rule is a collection of conditions and attribute filters that all have to be satisfied for certificate to be acceptable as authentication for a given username.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- msgVpnName|x||x||||| ruleName|x||x|||||    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.27.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnCertMatchingRule**](MsgVpnCertMatchingRule.md)| The Certificate Matching Rule object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **ruleName** | **string**| The name of the rule. | 
 **optional** | ***MsgVpnApiReplaceMsgVpnCertMatchingRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiReplaceMsgVpnCertMatchingRuleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnCertMatchingRuleResponse**](MsgVpnCertMatchingRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnCertMatchingRuleAttributeFilter**
> MsgVpnCertMatchingRuleAttributeFilterResponse ReplaceMsgVpnCertMatchingRuleAttributeFilter(ctx, body, msgVpnName, ruleName, filterName, optional)
Replace a Certificate Matching Rule Attribute Filter object.

Replace a Certificate Matching Rule Attribute Filter object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  A Cert Matching Rule Attribute Filter compares a username attribute to a string.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- filterName|x||x||||| msgVpnName|x||x||||| ruleName|x||x|||||    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.28.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnCertMatchingRuleAttributeFilter**](MsgVpnCertMatchingRuleAttributeFilter.md)| The Certificate Matching Rule Attribute Filter object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **ruleName** | **string**| The name of the rule. | 
  **filterName** | **string**| The name of the filter. | 
 **optional** | ***MsgVpnApiReplaceMsgVpnCertMatchingRuleAttributeFilterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiReplaceMsgVpnCertMatchingRuleAttributeFilterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnCertMatchingRuleAttributeFilterResponse**](MsgVpnCertMatchingRuleAttributeFilterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnClientProfile**
> MsgVpnClientProfileResponse ReplaceMsgVpnClientProfile(ctx, body, msgVpnName, clientProfileName, optional)
Replace a Client Profile object.

Replace a Client Profile object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  Client Profiles are used to assign common configuration properties to clients that have been successfully authorized.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- allowCutThroughForwardingEnabled|||||||x| apiQueueManagementCopyFromOnCreateName|||||||x| apiTopicEndpointManagementCopyFromOnCreateName|||||||x| clientProfileName|x||x||||| msgVpnName|x||x|||||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThreshold|clearPercent|setPercent|clearValue, setValue EventThreshold|clearValue|setValue|clearPercent, setPercent EventThreshold|setPercent|clearPercent|clearValue, setValue EventThreshold|setValue|clearValue|clearPercent, setPercent EventThresholdByPercent|clearPercent|setPercent| EventThresholdByPercent|setPercent|clearPercent|    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnClientProfile**](MsgVpnClientProfile.md)| The Client Profile object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **clientProfileName** | **string**| The name of the Client Profile. | 
 **optional** | ***MsgVpnApiReplaceMsgVpnClientProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiReplaceMsgVpnClientProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientProfileResponse**](MsgVpnClientProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnClientUsername**
> MsgVpnClientUsernameResponse ReplaceMsgVpnClientUsername(ctx, body, msgVpnName, clientUsername, optional)
Replace a Client Username object.

Replace a Client Username object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  A client is only authorized to connect to a Message VPN that is associated with a Client Username that the client has been assigned.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- aclProfileName||||||x|| clientProfileName||||||x|| clientUsername|x||x||||| msgVpnName|x||x||||| password||||x||||x    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnClientUsername**](MsgVpnClientUsername.md)| The Client Username object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **clientUsername** | **string**| The name of the Client Username. | 
 **optional** | ***MsgVpnApiReplaceMsgVpnClientUsernameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiReplaceMsgVpnClientUsernameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientUsernameResponse**](MsgVpnClientUsernameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnDistributedCache**
> MsgVpnDistributedCacheResponse ReplaceMsgVpnDistributedCache(ctx, body, msgVpnName, cacheName, optional)
Replace a Distributed Cache object.

Replace a Distributed Cache object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  A Distributed Cache is a collection of one or more Cache Clusters that belong to the same Message VPN. Each Cache Cluster in a Distributed Cache is configured to subscribe to a different set of topics. This effectively divides up the configured topic space, to provide scaling to very large topic spaces or very high cached message throughput.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- cacheName|x||x||||| cacheVirtualRouter||x|||||| msgVpnName|x||x|||||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- MsgVpnDistributedCache|scheduledDeleteMsgDayList|scheduledDeleteMsgTimeList| MsgVpnDistributedCache|scheduledDeleteMsgTimeList|scheduledDeleteMsgDayList|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnDistributedCache**](MsgVpnDistributedCache.md)| The Distributed Cache object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **cacheName** | **string**| The name of the Distributed Cache. | 
 **optional** | ***MsgVpnApiReplaceMsgVpnDistributedCacheOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiReplaceMsgVpnDistributedCacheOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheResponse**](MsgVpnDistributedCacheResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnDistributedCacheCluster**
> MsgVpnDistributedCacheClusterResponse ReplaceMsgVpnDistributedCacheCluster(ctx, body, msgVpnName, cacheName, clusterName, optional)
Replace a Cache Cluster object.

Replace a Cache Cluster object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  A Cache Cluster is a collection of one or more Cache Instances that subscribe to exactly the same topics. Cache Instances are grouped together in a Cache Cluster for the purpose of fault tolerance and load balancing. As published messages are received, the message broker message bus sends these live data messages to the Cache Instances in the Cache Cluster. This enables client cache requests to be served by any of Cache Instances in the Cache Cluster.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- cacheName|x||x||||| clusterName|x||x||||| deliverToOneOverrideEnabled||||||x|| msgVpnName|x||x|||||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThresholdByPercent|clearPercent|setPercent| EventThresholdByPercent|setPercent|clearPercent| EventThresholdByValue|clearValue|setValue| EventThresholdByValue|setValue|clearValue|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnDistributedCacheCluster**](MsgVpnDistributedCacheCluster.md)| The Cache Cluster object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **cacheName** | **string**| The name of the Distributed Cache. | 
  **clusterName** | **string**| The name of the Cache Cluster. | 
 **optional** | ***MsgVpnApiReplaceMsgVpnDistributedCacheClusterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiReplaceMsgVpnDistributedCacheClusterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterResponse**](MsgVpnDistributedCacheClusterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnDistributedCacheClusterInstance**
> MsgVpnDistributedCacheClusterInstanceResponse ReplaceMsgVpnDistributedCacheClusterInstance(ctx, body, msgVpnName, cacheName, clusterName, instanceName, optional)
Replace a Cache Instance object.

Replace a Cache Instance object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  A Cache Instance is a single Cache process that belongs to a single Cache Cluster. A Cache Instance object provisioned on the broker is used to disseminate configuration information to the Cache process. Cache Instances listen for and cache live data messages that match the topic subscriptions configured for their parent Cache Cluster.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- cacheName|x||x||||| clusterName|x||x||||| instanceName|x||x||||| msgVpnName|x||x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnDistributedCacheClusterInstance**](MsgVpnDistributedCacheClusterInstance.md)| The Cache Instance object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **cacheName** | **string**| The name of the Distributed Cache. | 
  **clusterName** | **string**| The name of the Cache Cluster. | 
  **instanceName** | **string**| The name of the Cache Instance. | 
 **optional** | ***MsgVpnApiReplaceMsgVpnDistributedCacheClusterInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiReplaceMsgVpnDistributedCacheClusterInstanceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterInstanceResponse**](MsgVpnDistributedCacheClusterInstanceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
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
 **optional** | ***MsgVpnApiReplaceMsgVpnDmrBridgeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiReplaceMsgVpnDmrBridgeOpts struct
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

# **ReplaceMsgVpnJndiConnectionFactory**
> MsgVpnJndiConnectionFactoryResponse ReplaceMsgVpnJndiConnectionFactory(ctx, body, msgVpnName, connectionFactoryName, optional)
Replace a JNDI Connection Factory object.

Replace a JNDI Connection Factory object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  The message broker provides an internal JNDI store for provisioned Connection Factory objects that clients can access through JNDI lookups.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- connectionFactoryName|x||x||||| msgVpnName|x||x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.2.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnJndiConnectionFactory**](MsgVpnJndiConnectionFactory.md)| The JNDI Connection Factory object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **connectionFactoryName** | **string**| The name of the JMS Connection Factory. | 
 **optional** | ***MsgVpnApiReplaceMsgVpnJndiConnectionFactoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiReplaceMsgVpnJndiConnectionFactoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnJndiConnectionFactoryResponse**](MsgVpnJndiConnectionFactoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnJndiQueue**
> MsgVpnJndiQueueResponse ReplaceMsgVpnJndiQueue(ctx, body, msgVpnName, queueName, optional)
Replace a JNDI Queue object.

Replace a JNDI Queue object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  The message broker provides an internal JNDI store for provisioned Queue objects that clients can access through JNDI lookups.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- msgVpnName|x||x||||| queueName|x||x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.2.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnJndiQueue**](MsgVpnJndiQueue.md)| The JNDI Queue object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The JNDI name of the JMS Queue. | 
 **optional** | ***MsgVpnApiReplaceMsgVpnJndiQueueOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiReplaceMsgVpnJndiQueueOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnJndiQueueResponse**](MsgVpnJndiQueueResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnJndiTopic**
> MsgVpnJndiTopicResponse ReplaceMsgVpnJndiTopic(ctx, body, msgVpnName, topicName, optional)
Replace a JNDI Topic object.

Replace a JNDI Topic object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  The message broker provides an internal JNDI store for provisioned Topic objects that clients can access through JNDI lookups.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- msgVpnName|x||x||||| topicName|x||x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.2.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnJndiTopic**](MsgVpnJndiTopic.md)| The JNDI Topic object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicName** | **string**| The JNDI name of the JMS Topic. | 
 **optional** | ***MsgVpnApiReplaceMsgVpnJndiTopicOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiReplaceMsgVpnJndiTopicOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnJndiTopicResponse**](MsgVpnJndiTopicResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnMqttRetainCache**
> MsgVpnMqttRetainCacheResponse ReplaceMsgVpnMqttRetainCache(ctx, body, msgVpnName, cacheName, optional)
Replace an MQTT Retain Cache object.

Replace an MQTT Retain Cache object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  Using MQTT retained messages allows publishing MQTT clients to indicate that a message must be stored for later delivery to subscribing clients when those subscribing clients add subscriptions matching the retained message's topic. An MQTT Retain Cache processes all retained messages for a Message VPN.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- cacheName|x||x||||| msgVpnName|x||x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnMqttRetainCache**](MsgVpnMqttRetainCache.md)| The MQTT Retain Cache object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **cacheName** | **string**| The name of the MQTT Retain Cache. | 
 **optional** | ***MsgVpnApiReplaceMsgVpnMqttRetainCacheOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiReplaceMsgVpnMqttRetainCacheOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnMqttRetainCacheResponse**](MsgVpnMqttRetainCacheResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnMqttSession**
> MsgVpnMqttSessionResponse ReplaceMsgVpnMqttSession(ctx, body, msgVpnName, mqttSessionClientId, mqttSessionVirtualRouter, optional)
Replace an MQTT Session object.

Replace an MQTT Session object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  An MQTT Session object is a virtual representation of an MQTT client connection. An MQTT session holds the state of an MQTT client (that is, it is used to contain a client's QoS 0 and QoS 1 subscription sets and any undelivered QoS 1 messages).   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- mqttSessionClientId|x||x||||| mqttSessionVirtualRouter|x||x||||| msgVpnName|x||x||||| owner||||||x|| queueRejectMsgToSenderOnDiscardBehavior||||||x||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThreshold|clearPercent|setPercent|clearValue, setValue EventThreshold|clearValue|setValue|clearPercent, setPercent EventThreshold|setPercent|clearPercent|clearValue, setValue EventThreshold|setValue|clearValue|clearPercent, setPercent    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.1.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnMqttSession**](MsgVpnMqttSession.md)| The MQTT Session object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **mqttSessionClientId** | **string**| The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet. | 
  **mqttSessionVirtualRouter** | **string**| The virtual router of the MQTT Session. | 
 **optional** | ***MsgVpnApiReplaceMsgVpnMqttSessionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiReplaceMsgVpnMqttSessionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnMqttSessionResponse**](MsgVpnMqttSessionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnMqttSessionSubscription**
> MsgVpnMqttSessionSubscriptionResponse ReplaceMsgVpnMqttSessionSubscription(ctx, body, msgVpnName, mqttSessionClientId, mqttSessionVirtualRouter, subscriptionTopic, optional)
Replace a Subscription object.

Replace a Subscription object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  An MQTT session contains a client's QoS 0 and QoS 1 subscription sets. On creation, a subscription defaults to QoS 0.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- mqttSessionClientId|x||x||||| mqttSessionVirtualRouter|x||x||||| msgVpnName|x||x||||| subscriptionTopic|x||x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.1.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnMqttSessionSubscription**](MsgVpnMqttSessionSubscription.md)| The Subscription object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **mqttSessionClientId** | **string**| The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet. | 
  **mqttSessionVirtualRouter** | **string**| The virtual router of the MQTT Session. | 
  **subscriptionTopic** | **string**| The MQTT subscription topic. | 
 **optional** | ***MsgVpnApiReplaceMsgVpnMqttSessionSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiReplaceMsgVpnMqttSessionSubscriptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnMqttSessionSubscriptionResponse**](MsgVpnMqttSessionSubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnQueue**
> MsgVpnQueueResponse ReplaceMsgVpnQueue(ctx, body, msgVpnName, queueName, optional)
Replace a Queue object.

Replace a Queue object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  A Queue acts as both a destination that clients can publish messages to, and as an endpoint that clients can bind consumers to and consume messages from.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- accessType||||||x|| msgVpnName|x||x||||| owner||||||x|| permission||||||x|| queueName|x||x||||| redeliveryDelayEnabled||||||x|| redeliveryDelayInitialInterval||||||x|| redeliveryDelayMaxInterval||||||x|| redeliveryDelayMultiplier||||||x|| rejectMsgToSenderOnDiscardBehavior||||||x|| respectMsgPriorityEnabled||||||x||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThreshold|clearPercent|setPercent|clearValue, setValue EventThreshold|clearValue|setValue|clearPercent, setPercent EventThreshold|setPercent|clearPercent|clearValue, setValue EventThreshold|setValue|clearValue|clearPercent, setPercent    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnQueue**](MsgVpnQueue.md)| The Queue object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The name of the Queue. | 
 **optional** | ***MsgVpnApiReplaceMsgVpnQueueOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiReplaceMsgVpnQueueOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueueResponse**](MsgVpnQueueResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnQueueTemplate**
> MsgVpnQueueTemplateResponse ReplaceMsgVpnQueueTemplate(ctx, body, msgVpnName, queueTemplateName, optional)
Replace a Queue Template object.

Replace a Queue Template object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  A Queue Template provides a mechanism for specifying the initial state for client created queues.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- msgVpnName|x||x||||| queueTemplateName|x||x||||| rejectMsgToSenderOnDiscardBehavior||||||x||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThreshold|clearPercent|setPercent|clearValue, setValue EventThreshold|clearValue|setValue|clearPercent, setPercent EventThreshold|setPercent|clearPercent|clearValue, setValue EventThreshold|setValue|clearValue|clearPercent, setPercent    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnQueueTemplate**](MsgVpnQueueTemplate.md)| The Queue Template object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueTemplateName** | **string**| The name of the Queue Template. | 
 **optional** | ***MsgVpnApiReplaceMsgVpnQueueTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiReplaceMsgVpnQueueTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueueTemplateResponse**](MsgVpnQueueTemplateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnReplayLog**
> MsgVpnReplayLogResponse ReplaceMsgVpnReplayLog(ctx, body, msgVpnName, replayLogName, optional)
Replace a Replay Log object.

Replace a Replay Log object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  When the Message Replay feature is enabled, message brokers store persistent messages in a Replay Log. These messages are kept until the log is full, after which the oldest messages are removed to free up space for new messages.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- msgVpnName|x||x||||| replayLogName|x||x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.10.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnReplayLog**](MsgVpnReplayLog.md)| The Replay Log object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **replayLogName** | **string**| The name of the Replay Log. | 
 **optional** | ***MsgVpnApiReplaceMsgVpnReplayLogOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiReplaceMsgVpnReplayLogOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnReplayLogResponse**](MsgVpnReplayLogResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnReplicatedTopic**
> MsgVpnReplicatedTopicResponse ReplaceMsgVpnReplicatedTopic(ctx, body, msgVpnName, replicatedTopic, optional)
Replace a Replicated Topic object.

Replace a Replicated Topic object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  To indicate which messages should be replicated between the active and standby site, a Replicated Topic subscription must be configured on a Message VPN. If a published message matches both a replicated topic and an endpoint on the active site, then the message is replicated to the standby site.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- msgVpnName|x||x||||| replicatedTopic|x||x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.1.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnReplicatedTopic**](MsgVpnReplicatedTopic.md)| The Replicated Topic object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **replicatedTopic** | **string**| The topic for applying replication. Published messages matching this topic will be replicated to the standby site. | 
 **optional** | ***MsgVpnApiReplaceMsgVpnReplicatedTopicOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiReplaceMsgVpnReplicatedTopicOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnReplicatedTopicResponse**](MsgVpnReplicatedTopicResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnRestDeliveryPoint**
> MsgVpnRestDeliveryPointResponse ReplaceMsgVpnRestDeliveryPoint(ctx, body, msgVpnName, restDeliveryPointName, optional)
Replace a REST Delivery Point object.

Replace a REST Delivery Point object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  A REST Delivery Point manages delivery of messages from queues to a named list of REST Consumers.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- clientProfileName||||||x|| msgVpnName|x||x||||| restDeliveryPointName|x||x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnRestDeliveryPoint**](MsgVpnRestDeliveryPoint.md)| The REST Delivery Point object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
 **optional** | ***MsgVpnApiReplaceMsgVpnRestDeliveryPointOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiReplaceMsgVpnRestDeliveryPointOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointResponse**](MsgVpnRestDeliveryPointResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnRestDeliveryPointQueueBinding**
> MsgVpnRestDeliveryPointQueueBindingResponse ReplaceMsgVpnRestDeliveryPointQueueBinding(ctx, body, msgVpnName, restDeliveryPointName, queueBindingName, optional)
Replace a Queue Binding object.

Replace a Queue Binding object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  A Queue Binding for a REST Delivery Point attracts messages to be delivered to REST consumers. If the queue does not exist it can be created subsequently, and once the queue is operational the broker performs the queue binding. Removing the queue binding does not delete the queue itself. Similarly, removing the queue does not remove the queue binding, which fails until the queue is recreated or the queue binding is deleted.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- msgVpnName|x||x||||| queueBindingName|x||x||||| restDeliveryPointName|x||x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnRestDeliveryPointQueueBinding**](MsgVpnRestDeliveryPointQueueBinding.md)| The Queue Binding object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **queueBindingName** | **string**| The name of a queue in the Message VPN. | 
 **optional** | ***MsgVpnApiReplaceMsgVpnRestDeliveryPointQueueBindingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiReplaceMsgVpnRestDeliveryPointQueueBindingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointQueueBindingResponse**](MsgVpnRestDeliveryPointQueueBindingResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader**
> MsgVpnRestDeliveryPointQueueBindingProtectedRequestHeaderResponse ReplaceMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader(ctx, body, msgVpnName, restDeliveryPointName, queueBindingName, headerName, optional)
Replace a Protected Request Header object.

Replace a Protected Request Header object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  A protected request header to be added to the HTTP request. Unlike a non-protected request header, the header value cannot be displayed after it is set.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- headerName|x||x||||| headerValue||||x||||x msgVpnName|x||x||||| queueBindingName|x||x||||| restDeliveryPointName|x||x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.30.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader**](MsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader.md)| The Protected Request Header object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **queueBindingName** | **string**| The name of a queue in the Message VPN. | 
  **headerName** | **string**| The name of the protected HTTP request header. | 
 **optional** | ***MsgVpnApiReplaceMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeaderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiReplaceMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeaderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointQueueBindingProtectedRequestHeaderResponse**](MsgVpnRestDeliveryPointQueueBindingProtectedRequestHeaderResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnRestDeliveryPointQueueBindingRequestHeader**
> MsgVpnRestDeliveryPointQueueBindingRequestHeaderResponse ReplaceMsgVpnRestDeliveryPointQueueBindingRequestHeader(ctx, body, msgVpnName, restDeliveryPointName, queueBindingName, headerName, optional)
Replace a Request Header object.

Replace a Request Header object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  A request header to be added to the HTTP request.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- headerName|x||x||||| msgVpnName|x||x||||| queueBindingName|x||x||||| restDeliveryPointName|x||x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.23.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnRestDeliveryPointQueueBindingRequestHeader**](MsgVpnRestDeliveryPointQueueBindingRequestHeader.md)| The Request Header object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **queueBindingName** | **string**| The name of a queue in the Message VPN. | 
  **headerName** | **string**| The name of the HTTP request header. | 
 **optional** | ***MsgVpnApiReplaceMsgVpnRestDeliveryPointQueueBindingRequestHeaderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiReplaceMsgVpnRestDeliveryPointQueueBindingRequestHeaderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointQueueBindingRequestHeaderResponse**](MsgVpnRestDeliveryPointQueueBindingRequestHeaderResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnRestDeliveryPointRestConsumer**
> MsgVpnRestDeliveryPointRestConsumerResponse ReplaceMsgVpnRestDeliveryPointRestConsumer(ctx, body, msgVpnName, restDeliveryPointName, restConsumerName, optional)
Replace a REST Consumer object.

Replace a REST Consumer object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  REST Consumer objects establish HTTP connectivity to REST consumer applications who wish to receive messages from a broker.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- authenticationAwsSecretAccessKey||||x||||x authenticationClientCertContent||||x||x||x authenticationClientCertPassword||||x||x|| authenticationHttpBasicPassword||||x||x||x authenticationHttpBasicUsername||||||x|| authenticationHttpHeaderValue||||x||||x authenticationOauthClientId||||||x|| authenticationOauthClientScope||||||x|| authenticationOauthClientSecret||||x||x||x authenticationOauthClientTokenEndpoint||||||x|| authenticationOauthClientTokenExpiryDefault||||||x|| authenticationOauthJwtSecretKey||||x||x||x authenticationOauthJwtTokenEndpoint||||||x|| authenticationOauthJwtTokenExpiryDefault||||||x|| authenticationScheme||||||x|| msgVpnName|x||x||||| outgoingConnectionCount||||||x|| remoteHost||||||x|| remotePort||||||x|| restConsumerName|x||x||||| restDeliveryPointName|x||x||||| tlsCipherSuiteList||||||x|| tlsEnabled||||||x||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- MsgVpnRestDeliveryPointRestConsumer|authenticationClientCertPassword|authenticationClientCertContent| MsgVpnRestDeliveryPointRestConsumer|authenticationHttpBasicPassword|authenticationHttpBasicUsername| MsgVpnRestDeliveryPointRestConsumer|authenticationHttpBasicUsername|authenticationHttpBasicPassword| MsgVpnRestDeliveryPointRestConsumer|remotePort|tlsEnabled| MsgVpnRestDeliveryPointRestConsumer|tlsEnabled|remotePort|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnRestDeliveryPointRestConsumer**](MsgVpnRestDeliveryPointRestConsumer.md)| The REST Consumer object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **restConsumerName** | **string**| The name of the REST Consumer. | 
 **optional** | ***MsgVpnApiReplaceMsgVpnRestDeliveryPointRestConsumerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiReplaceMsgVpnRestDeliveryPointRestConsumerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointRestConsumerResponse**](MsgVpnRestDeliveryPointRestConsumerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnTelemetryProfile**
> MsgVpnTelemetryProfileResponse ReplaceMsgVpnTelemetryProfile(ctx, body, msgVpnName, telemetryProfileName, optional)
Replace a Telemetry Profile object.

Replace a Telemetry Profile object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  Using the Telemetry Profile allows trace spans to be generated as messages are processed by the broker. The generated spans are stored persistently on the broker and may be consumed by the Solace receiver component of an OpenTelemetry Collector.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- msgVpnName|x||x||||| telemetryProfileName|x||x|||||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThreshold|clearPercent|setPercent|clearValue, setValue EventThreshold|clearValue|setValue|clearPercent, setPercent EventThreshold|setPercent|clearPercent|clearValue, setValue EventThreshold|setValue|clearValue|clearPercent, setPercent    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.31.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnTelemetryProfile**](MsgVpnTelemetryProfile.md)| The Telemetry Profile object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **telemetryProfileName** | **string**| The name of the Telemetry Profile. | 
 **optional** | ***MsgVpnApiReplaceMsgVpnTelemetryProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiReplaceMsgVpnTelemetryProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTelemetryProfileResponse**](MsgVpnTelemetryProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnTelemetryProfileTraceFilter**
> MsgVpnTelemetryProfileTraceFilterResponse ReplaceMsgVpnTelemetryProfileTraceFilter(ctx, body, msgVpnName, telemetryProfileName, traceFilterName, optional)
Replace a Trace Filter object.

Replace a Trace Filter object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  A Trace Filter controls which messages received by the broker will be traced. If an incoming message matches an enabled tracing filter's subscription, the message will be traced as it passes through the broker.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- msgVpnName|x||x||||| telemetryProfileName|x||x||||| traceFilterName|x||x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.31.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnTelemetryProfileTraceFilter**](MsgVpnTelemetryProfileTraceFilter.md)| The Trace Filter object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **telemetryProfileName** | **string**| The name of the Telemetry Profile. | 
  **traceFilterName** | **string**| A name used to identify the trace filter. Consider a name that describes the subscriptions contained within the filter, such as the name of the application and/or the scenario in which the trace filter might be enabled, such as \&quot;appNameDebug\&quot;. | 
 **optional** | ***MsgVpnApiReplaceMsgVpnTelemetryProfileTraceFilterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiReplaceMsgVpnTelemetryProfileTraceFilterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTelemetryProfileTraceFilterResponse**](MsgVpnTelemetryProfileTraceFilterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnTopicEndpoint**
> MsgVpnTopicEndpointResponse ReplaceMsgVpnTopicEndpoint(ctx, body, msgVpnName, topicEndpointName, optional)
Replace a Topic Endpoint object.

Replace a Topic Endpoint object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  A Topic Endpoint attracts messages published to a topic for which the Topic Endpoint has a matching topic subscription. The topic subscription for the Topic Endpoint is specified in the client request to bind a Flow to that Topic Endpoint. Queues are significantly more flexible than Topic Endpoints and are the recommended approach for most applications. The use of Topic Endpoints should be restricted to JMS applications.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- accessType||||||x|| msgVpnName|x||x||||| owner||||||x|| permission||||||x|| redeliveryDelayEnabled||||||x|| redeliveryDelayInitialInterval||||||x|| redeliveryDelayMaxInterval||||||x|| redeliveryDelayMultiplier||||||x|| rejectMsgToSenderOnDiscardBehavior||||||x|| respectMsgPriorityEnabled||||||x|| topicEndpointName|x||x|||||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThreshold|clearPercent|setPercent|clearValue, setValue EventThreshold|clearValue|setValue|clearPercent, setPercent EventThreshold|setPercent|clearPercent|clearValue, setValue EventThreshold|setValue|clearValue|clearPercent, setPercent    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.1.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnTopicEndpoint**](MsgVpnTopicEndpoint.md)| The Topic Endpoint object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointName** | **string**| The name of the Topic Endpoint. | 
 **optional** | ***MsgVpnApiReplaceMsgVpnTopicEndpointOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiReplaceMsgVpnTopicEndpointOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointResponse**](MsgVpnTopicEndpointResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMsgVpnTopicEndpointTemplate**
> MsgVpnTopicEndpointTemplateResponse ReplaceMsgVpnTopicEndpointTemplate(ctx, body, msgVpnName, topicEndpointTemplateName, optional)
Replace a Topic Endpoint Template object.

Replace a Topic Endpoint Template object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  A Topic Endpoint Template provides a mechanism for specifying the initial state for client created topic endpoints.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- msgVpnName|x||x||||| rejectMsgToSenderOnDiscardBehavior||||||x|| topicEndpointTemplateName|x||x|||||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThreshold|clearPercent|setPercent|clearValue, setValue EventThreshold|clearValue|setValue|clearPercent, setPercent EventThreshold|setPercent|clearPercent|clearValue, setValue EventThreshold|setValue|clearValue|clearPercent, setPercent    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnTopicEndpointTemplate**](MsgVpnTopicEndpointTemplate.md)| The Topic Endpoint Template object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointTemplateName** | **string**| The name of the Topic Endpoint Template. | 
 **optional** | ***MsgVpnApiReplaceMsgVpnTopicEndpointTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiReplaceMsgVpnTopicEndpointTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointTemplateResponse**](MsgVpnTopicEndpointTemplateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpn**
> MsgVpnResponse UpdateMsgVpn(ctx, body, msgVpnName, optional)
Update a Message VPN object.

Update a Message VPN object. Any attribute missing from the request will be left unchanged.  Message VPNs (Virtual Private Networks) allow for the segregation of topic space and clients. They also group clients connecting to a network of message brokers, such that messages published within a particular group are only visible to that group's clients.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- authenticationOauthDefaultProviderName||||||x| bridgingTlsServerCertEnforceTrustedCommonNameEnabled||||||x| distributedCacheManagementEnabled||||||x| msgVpnName|x|x||||| replicationBridgeAuthenticationBasicPassword|||x||||x replicationBridgeAuthenticationClientCertContent|||x||||x replicationBridgeAuthenticationClientCertPassword|||x|||| replicationEnabledQueueBehavior|||x|||| restTlsServerCertEnforceTrustedCommonNameEnabled||||||x| serviceAmqpPlainTextListenPort|||||x|| serviceAmqpTlsListenPort|||||x|| serviceMqttPlainTextListenPort|||||x|| serviceMqttTlsListenPort|||||x|| serviceMqttTlsWebSocketListenPort|||||x|| serviceMqttWebSocketListenPort|||||x|| serviceRestIncomingPlainTextListenPort|||||x|| serviceRestIncomingTlsListenPort|||||x||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThreshold|clearPercent|setPercent|clearValue, setValue EventThreshold|clearValue|setValue|clearPercent, setPercent EventThreshold|setPercent|clearPercent|clearValue, setValue EventThreshold|setValue|clearValue|clearPercent, setPercent EventThresholdByValue|clearValue|setValue| EventThresholdByValue|setValue|clearValue| MsgVpn|authenticationBasicProfileName|authenticationBasicType| MsgVpn|authorizationProfileName|authorizationType| MsgVpn|eventPublishTopicFormatMqttEnabled|eventPublishTopicFormatSmfEnabled| MsgVpn|eventPublishTopicFormatSmfEnabled|eventPublishTopicFormatMqttEnabled| MsgVpn|replicationBridgeAuthenticationBasicClientUsername|replicationBridgeAuthenticationBasicPassword| MsgVpn|replicationBridgeAuthenticationBasicPassword|replicationBridgeAuthenticationBasicClientUsername| MsgVpn|replicationBridgeAuthenticationClientCertPassword|replicationBridgeAuthenticationClientCertContent| MsgVpn|replicationEnabledQueueBehavior|replicationEnabled|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation. Requests which include the following attributes require greater access scope/level:   Attribute|Access Scope/Level :---|:---: alias|global/read-write authenticationBasicEnabled|global/read-write authenticationBasicProfileName|global/read-write authenticationBasicRadiusDomain|global/read-write authenticationBasicType|global/read-write authenticationClientCertAllowApiProvidedUsernameEnabled|global/read-write authenticationClientCertCertificateMatchingRulesEnabled|global/read-write authenticationClientCertEnabled|global/read-write authenticationClientCertMaxChainDepth|global/read-write authenticationClientCertRevocationCheckMode|global/read-write authenticationClientCertUsernameSource|global/read-write authenticationClientCertValidateDateEnabled|global/read-write authenticationKerberosAllowApiProvidedUsernameEnabled|global/read-write authenticationKerberosEnabled|global/read-write authenticationOauthEnabled|global/read-write bridgingTlsServerCertEnforceTrustedCommonNameEnabled|global/read-write bridgingTlsServerCertMaxChainDepth|global/read-write bridgingTlsServerCertValidateDateEnabled|global/read-write bridgingTlsServerCertValidateNameEnabled|global/read-write dmrEnabled|global/read-write exportSubscriptionsEnabled|global/read-write maxConnectionCount|global/read-write maxEgressFlowCount|global/read-write maxEndpointCount|global/read-write maxIngressFlowCount|global/read-write maxMsgSpoolUsage|global/read-write maxSubscriptionCount|global/read-write maxTransactedSessionCount|global/read-write maxTransactionCount|global/read-write mqttRetainMaxMemory|global/read-write replicationBridgeAuthenticationBasicClientUsername|global/read-write replicationBridgeAuthenticationBasicPassword|global/read-write replicationBridgeAuthenticationClientCertContent|global/read-write replicationBridgeAuthenticationClientCertPassword|global/read-write replicationBridgeAuthenticationScheme|global/read-write replicationBridgeCompressedDataEnabled|global/read-write replicationBridgeEgressFlowWindowSize|global/read-write replicationBridgeRetryDelay|global/read-write replicationBridgeTlsEnabled|global/read-write replicationBridgeUnidirectionalClientProfileName|global/read-write replicationEnabled|global/read-write replicationEnabledQueueBehavior|global/read-write replicationQueueMaxMsgSpoolUsage|global/read-write replicationRole|global/read-write restTlsServerCertEnforceTrustedCommonNameEnabled|global/read-write restTlsServerCertMaxChainDepth|global/read-write restTlsServerCertValidateDateEnabled|global/read-write restTlsServerCertValidateNameEnabled|global/read-write sempOverMsgBusAdminClientEnabled|global/read-write sempOverMsgBusAdminDistributedCacheEnabled|global/read-write sempOverMsgBusAdminEnabled|global/read-write sempOverMsgBusEnabled|global/read-write sempOverMsgBusShowEnabled|global/read-write serviceAmqpMaxConnectionCount|global/read-write serviceAmqpPlainTextListenPort|global/read-write serviceAmqpTlsListenPort|global/read-write serviceMqttMaxConnectionCount|global/read-write serviceMqttPlainTextListenPort|global/read-write serviceMqttTlsListenPort|global/read-write serviceMqttTlsWebSocketListenPort|global/read-write serviceMqttWebSocketListenPort|global/read-write serviceRestIncomingMaxConnectionCount|global/read-write serviceRestIncomingPlainTextListenPort|global/read-write serviceRestIncomingTlsListenPort|global/read-write serviceRestOutgoingMaxConnectionCount|global/read-write serviceSmfMaxConnectionCount|global/read-write serviceWebMaxConnectionCount|global/read-write    This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpn**](MsgVpn.md)| The Message VPN object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***MsgVpnApiUpdateMsgVpnOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiUpdateMsgVpnOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnResponse**](MsgVpnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnAclProfile**
> MsgVpnAclProfileResponse UpdateMsgVpnAclProfile(ctx, body, msgVpnName, aclProfileName, optional)
Update an ACL Profile object.

Update an ACL Profile object. Any attribute missing from the request will be left unchanged.  An ACL Profile controls whether an authenticated client is permitted to establish a connection with the message broker or permitted to publish and subscribe to specific topics.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- aclProfileName|x|x||||| msgVpnName|x|x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnAclProfile**](MsgVpnAclProfile.md)| The ACL Profile object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **aclProfileName** | **string**| The name of the ACL Profile. | 
 **optional** | ***MsgVpnApiUpdateMsgVpnAclProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiUpdateMsgVpnAclProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileResponse**](MsgVpnAclProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnAuthenticationOauthProfile**
> MsgVpnAuthenticationOauthProfileResponse UpdateMsgVpnAuthenticationOauthProfile(ctx, body, msgVpnName, oauthProfileName, optional)
Update an OAuth Profile object.

Update an OAuth Profile object. Any attribute missing from the request will be left unchanged.  OAuth profiles specify how to securely authenticate to an OAuth provider.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- clientSecret|||x||||x msgVpnName|x|x||||| oauthProfileName|x|x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.25.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnAuthenticationOauthProfile**](MsgVpnAuthenticationOauthProfile.md)| The OAuth Profile object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **oauthProfileName** | **string**| The name of the OAuth profile. | 
 **optional** | ***MsgVpnApiUpdateMsgVpnAuthenticationOauthProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiUpdateMsgVpnAuthenticationOauthProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthenticationOauthProfileResponse**](MsgVpnAuthenticationOauthProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnAuthenticationOauthProvider**
> MsgVpnAuthenticationOauthProviderResponse UpdateMsgVpnAuthenticationOauthProvider(ctx, body, msgVpnName, oauthProviderName, optional)
Update an OAuth Provider object.

Update an OAuth Provider object. Any attribute missing from the request will be left unchanged.  OAuth Providers contain information about the issuer of an OAuth token that is needed to validate the token and derive a client username from it.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- audienceClaimName||||||x| audienceClaimSource||||||x| audienceClaimValue||||||x| audienceValidationEnabled||||||x| authorizationGroupClaimName||||||x| authorizationGroupClaimSource||||||x| authorizationGroupEnabled||||||x| disconnectOnTokenExpirationEnabled||||||x| enabled||||||x| jwksRefreshInterval||||||x| jwksUri||||||x| msgVpnName|x|x||||x| oauthProviderName|x|x||||x| tokenIgnoreTimeLimitsEnabled||||||x| tokenIntrospectionParameterName||||||x| tokenIntrospectionPassword|||x|||x|x tokenIntrospectionTimeout||||||x| tokenIntrospectionUri||||||x| tokenIntrospectionUsername||||||x| usernameClaimName||||||x| usernameClaimSource||||||x| usernameValidateEnabled||||||x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been deprecated since 2.25. Replaced by authenticationOauthProfiles.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnAuthenticationOauthProvider**](MsgVpnAuthenticationOauthProvider.md)| The OAuth Provider object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **oauthProviderName** | **string**| The name of the OAuth Provider. | 
 **optional** | ***MsgVpnApiUpdateMsgVpnAuthenticationOauthProviderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiUpdateMsgVpnAuthenticationOauthProviderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthenticationOauthProviderResponse**](MsgVpnAuthenticationOauthProviderResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnAuthorizationGroup**
> MsgVpnAuthorizationGroupResponse UpdateMsgVpnAuthorizationGroup(ctx, body, msgVpnName, authorizationGroupName, optional)
Update an Authorization Group object.

Update an Authorization Group object. Any attribute missing from the request will be left unchanged.  To use client authorization groups configured on an external server to provide client authorizations, Authorization Group objects must be created on the Message VPN that match the authorization groups provisioned on the external server. These objects must be configured with the client profiles and ACL profiles that will be assigned to the clients that belong to those authorization groups. A newly created group is placed at the end of the group list which is the lowest priority.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- aclProfileName|||||x|| authorizationGroupName|x|x||||| clientProfileName|||||x|| msgVpnName|x|x||||| orderAfterAuthorizationGroupName|||x|||| orderBeforeAuthorizationGroupName|||x||||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- MsgVpnAuthorizationGroup|orderAfterAuthorizationGroupName||orderBeforeAuthorizationGroupName MsgVpnAuthorizationGroup|orderBeforeAuthorizationGroupName||orderAfterAuthorizationGroupName    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnAuthorizationGroup**](MsgVpnAuthorizationGroup.md)| The Authorization Group object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **authorizationGroupName** | **string**| The name of the Authorization Group. For LDAP groups, special care is needed if the group name contains special characters such as &#x27;#&#x27;, &#x27;+&#x27;, &#x27;;&#x27;, &#x27;&#x3D;&#x27; as the value of the group name returned from the LDAP server might prepend those characters with &#x27;\\&#x27;. For example a group name called &#x27;test#,lab,com&#x27; will be returned from the LDAP server as &#x27;test\\#,lab,com&#x27;. | 
 **optional** | ***MsgVpnApiUpdateMsgVpnAuthorizationGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiUpdateMsgVpnAuthorizationGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthorizationGroupResponse**](MsgVpnAuthorizationGroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnBridge**
> MsgVpnBridgeResponse UpdateMsgVpnBridge(ctx, body, msgVpnName, bridgeName, bridgeVirtualRouter, optional)
Update a Bridge object.

Update a Bridge object. Any attribute missing from the request will be left unchanged.  Bridges can be used to link two Message VPNs so that messages published to one Message VPN that match the topic subscriptions set for the bridge are also delivered to the linked Message VPN.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- bridgeName|x|x||||| bridgeVirtualRouter|x|x||||| maxTtl|||||x|| msgVpnName|x|x||||| remoteAuthenticationBasicClientUsername|||||x|| remoteAuthenticationBasicPassword|||x||x||x remoteAuthenticationClientCertContent|||x||x||x remoteAuthenticationClientCertPassword|||x||x|| remoteAuthenticationScheme|||||x|| remoteDeliverToOnePriority|||||x||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- MsgVpnBridge|remoteAuthenticationBasicClientUsername|remoteAuthenticationBasicPassword| MsgVpnBridge|remoteAuthenticationBasicPassword|remoteAuthenticationBasicClientUsername| MsgVpnBridge|remoteAuthenticationClientCertPassword|remoteAuthenticationClientCertContent|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnBridge**](MsgVpnBridge.md)| The Bridge object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **bridgeName** | **string**| The name of the Bridge. | 
  **bridgeVirtualRouter** | **string**| The virtual router of the Bridge. | 
 **optional** | ***MsgVpnApiUpdateMsgVpnBridgeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiUpdateMsgVpnBridgeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeResponse**](MsgVpnBridgeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnBridgeRemoteMsgVpn**
> MsgVpnBridgeRemoteMsgVpnResponse UpdateMsgVpnBridgeRemoteMsgVpn(ctx, body, msgVpnName, bridgeName, bridgeVirtualRouter, remoteMsgVpnName, remoteMsgVpnLocation, remoteMsgVpnInterface, optional)
Update a Remote Message VPN object.

Update a Remote Message VPN object. Any attribute missing from the request will be left unchanged.  The Remote Message VPN is the Message VPN that the Bridge connects to.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- bridgeName|x|x||||| bridgeVirtualRouter|x|x||||| clientUsername|||||x|| compressedDataEnabled|||||x|| egressFlowWindowSize|||||x|| msgVpnName|x|x||||| password|||x||x||x remoteMsgVpnInterface|x|x||||| remoteMsgVpnLocation|x|x||||| remoteMsgVpnName|x|x||||| tlsEnabled|||||x||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- MsgVpnBridgeRemoteMsgVpn|clientUsername|password| MsgVpnBridgeRemoteMsgVpn|password|clientUsername|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnBridgeRemoteMsgVpn**](MsgVpnBridgeRemoteMsgVpn.md)| The Remote Message VPN object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **bridgeName** | **string**| The name of the Bridge. | 
  **bridgeVirtualRouter** | **string**| The virtual router of the Bridge. | 
  **remoteMsgVpnName** | **string**| The name of the remote Message VPN. | 
  **remoteMsgVpnLocation** | **string**| The location of the remote Message VPN as either an FQDN with port, IP address with port, or virtual router name (starting with \&quot;v:\&quot;). | 
  **remoteMsgVpnInterface** | **string**| The physical interface on the local Message VPN host for connecting to the remote Message VPN. By default, an interface is chosen automatically (recommended), but if specified, &#x60;remoteMsgVpnLocation&#x60; must not be a virtual router name. | 
 **optional** | ***MsgVpnApiUpdateMsgVpnBridgeRemoteMsgVpnOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiUpdateMsgVpnBridgeRemoteMsgVpnOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------







 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeRemoteMsgVpnResponse**](MsgVpnBridgeRemoteMsgVpnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnCertMatchingRule**
> MsgVpnCertMatchingRuleResponse UpdateMsgVpnCertMatchingRule(ctx, body, msgVpnName, ruleName, optional)
Update a Certificate Matching Rule object.

Update a Certificate Matching Rule object. Any attribute missing from the request will be left unchanged.  A Cert Matching Rule is a collection of conditions and attribute filters that all have to be satisfied for certificate to be acceptable as authentication for a given username.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- msgVpnName|x|x||||| ruleName|x|x|||||    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.27.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnCertMatchingRule**](MsgVpnCertMatchingRule.md)| The Certificate Matching Rule object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **ruleName** | **string**| The name of the rule. | 
 **optional** | ***MsgVpnApiUpdateMsgVpnCertMatchingRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiUpdateMsgVpnCertMatchingRuleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnCertMatchingRuleResponse**](MsgVpnCertMatchingRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnCertMatchingRuleAttributeFilter**
> MsgVpnCertMatchingRuleAttributeFilterResponse UpdateMsgVpnCertMatchingRuleAttributeFilter(ctx, body, msgVpnName, ruleName, filterName, optional)
Update a Certificate Matching Rule Attribute Filter object.

Update a Certificate Matching Rule Attribute Filter object. Any attribute missing from the request will be left unchanged.  A Cert Matching Rule Attribute Filter compares a username attribute to a string.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- filterName|x|x||||| msgVpnName|x|x||||| ruleName|x|x|||||    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.28.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnCertMatchingRuleAttributeFilter**](MsgVpnCertMatchingRuleAttributeFilter.md)| The Certificate Matching Rule Attribute Filter object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **ruleName** | **string**| The name of the rule. | 
  **filterName** | **string**| The name of the filter. | 
 **optional** | ***MsgVpnApiUpdateMsgVpnCertMatchingRuleAttributeFilterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiUpdateMsgVpnCertMatchingRuleAttributeFilterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnCertMatchingRuleAttributeFilterResponse**](MsgVpnCertMatchingRuleAttributeFilterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnClientProfile**
> MsgVpnClientProfileResponse UpdateMsgVpnClientProfile(ctx, body, msgVpnName, clientProfileName, optional)
Update a Client Profile object.

Update a Client Profile object. Any attribute missing from the request will be left unchanged.  Client Profiles are used to assign common configuration properties to clients that have been successfully authorized.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- allowCutThroughForwardingEnabled||||||x| apiQueueManagementCopyFromOnCreateName||||||x| apiTopicEndpointManagementCopyFromOnCreateName||||||x| clientProfileName|x|x||||| msgVpnName|x|x|||||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThreshold|clearPercent|setPercent|clearValue, setValue EventThreshold|clearValue|setValue|clearPercent, setPercent EventThreshold|setPercent|clearPercent|clearValue, setValue EventThreshold|setValue|clearValue|clearPercent, setPercent EventThresholdByPercent|clearPercent|setPercent| EventThresholdByPercent|setPercent|clearPercent|    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnClientProfile**](MsgVpnClientProfile.md)| The Client Profile object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **clientProfileName** | **string**| The name of the Client Profile. | 
 **optional** | ***MsgVpnApiUpdateMsgVpnClientProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiUpdateMsgVpnClientProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientProfileResponse**](MsgVpnClientProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnClientUsername**
> MsgVpnClientUsernameResponse UpdateMsgVpnClientUsername(ctx, body, msgVpnName, clientUsername, optional)
Update a Client Username object.

Update a Client Username object. Any attribute missing from the request will be left unchanged.  A client is only authorized to connect to a Message VPN that is associated with a Client Username that the client has been assigned.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- aclProfileName|||||x|| clientProfileName|||||x|| clientUsername|x|x||||| msgVpnName|x|x||||| password|||x||||x    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnClientUsername**](MsgVpnClientUsername.md)| The Client Username object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **clientUsername** | **string**| The name of the Client Username. | 
 **optional** | ***MsgVpnApiUpdateMsgVpnClientUsernameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiUpdateMsgVpnClientUsernameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientUsernameResponse**](MsgVpnClientUsernameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnDistributedCache**
> MsgVpnDistributedCacheResponse UpdateMsgVpnDistributedCache(ctx, body, msgVpnName, cacheName, optional)
Update a Distributed Cache object.

Update a Distributed Cache object. Any attribute missing from the request will be left unchanged.  A Distributed Cache is a collection of one or more Cache Clusters that belong to the same Message VPN. Each Cache Cluster in a Distributed Cache is configured to subscribe to a different set of topics. This effectively divides up the configured topic space, to provide scaling to very large topic spaces or very high cached message throughput.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- cacheName|x|x||||| cacheVirtualRouter||x||||| msgVpnName|x|x|||||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- MsgVpnDistributedCache|scheduledDeleteMsgDayList|scheduledDeleteMsgTimeList| MsgVpnDistributedCache|scheduledDeleteMsgTimeList|scheduledDeleteMsgDayList|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnDistributedCache**](MsgVpnDistributedCache.md)| The Distributed Cache object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **cacheName** | **string**| The name of the Distributed Cache. | 
 **optional** | ***MsgVpnApiUpdateMsgVpnDistributedCacheOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiUpdateMsgVpnDistributedCacheOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheResponse**](MsgVpnDistributedCacheResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnDistributedCacheCluster**
> MsgVpnDistributedCacheClusterResponse UpdateMsgVpnDistributedCacheCluster(ctx, body, msgVpnName, cacheName, clusterName, optional)
Update a Cache Cluster object.

Update a Cache Cluster object. Any attribute missing from the request will be left unchanged.  A Cache Cluster is a collection of one or more Cache Instances that subscribe to exactly the same topics. Cache Instances are grouped together in a Cache Cluster for the purpose of fault tolerance and load balancing. As published messages are received, the message broker message bus sends these live data messages to the Cache Instances in the Cache Cluster. This enables client cache requests to be served by any of Cache Instances in the Cache Cluster.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- cacheName|x|x||||| clusterName|x|x||||| deliverToOneOverrideEnabled|||||x|| msgVpnName|x|x|||||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThresholdByPercent|clearPercent|setPercent| EventThresholdByPercent|setPercent|clearPercent| EventThresholdByValue|clearValue|setValue| EventThresholdByValue|setValue|clearValue|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnDistributedCacheCluster**](MsgVpnDistributedCacheCluster.md)| The Cache Cluster object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **cacheName** | **string**| The name of the Distributed Cache. | 
  **clusterName** | **string**| The name of the Cache Cluster. | 
 **optional** | ***MsgVpnApiUpdateMsgVpnDistributedCacheClusterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiUpdateMsgVpnDistributedCacheClusterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterResponse**](MsgVpnDistributedCacheClusterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnDistributedCacheClusterInstance**
> MsgVpnDistributedCacheClusterInstanceResponse UpdateMsgVpnDistributedCacheClusterInstance(ctx, body, msgVpnName, cacheName, clusterName, instanceName, optional)
Update a Cache Instance object.

Update a Cache Instance object. Any attribute missing from the request will be left unchanged.  A Cache Instance is a single Cache process that belongs to a single Cache Cluster. A Cache Instance object provisioned on the broker is used to disseminate configuration information to the Cache process. Cache Instances listen for and cache live data messages that match the topic subscriptions configured for their parent Cache Cluster.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- cacheName|x|x||||| clusterName|x|x||||| instanceName|x|x||||| msgVpnName|x|x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnDistributedCacheClusterInstance**](MsgVpnDistributedCacheClusterInstance.md)| The Cache Instance object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **cacheName** | **string**| The name of the Distributed Cache. | 
  **clusterName** | **string**| The name of the Cache Cluster. | 
  **instanceName** | **string**| The name of the Cache Instance. | 
 **optional** | ***MsgVpnApiUpdateMsgVpnDistributedCacheClusterInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiUpdateMsgVpnDistributedCacheClusterInstanceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterInstanceResponse**](MsgVpnDistributedCacheClusterInstanceResponse.md)

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
 **optional** | ***MsgVpnApiUpdateMsgVpnDmrBridgeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiUpdateMsgVpnDmrBridgeOpts struct
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

# **UpdateMsgVpnJndiConnectionFactory**
> MsgVpnJndiConnectionFactoryResponse UpdateMsgVpnJndiConnectionFactory(ctx, body, msgVpnName, connectionFactoryName, optional)
Update a JNDI Connection Factory object.

Update a JNDI Connection Factory object. Any attribute missing from the request will be left unchanged.  The message broker provides an internal JNDI store for provisioned Connection Factory objects that clients can access through JNDI lookups.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- connectionFactoryName|x|x||||| msgVpnName|x|x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.2.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnJndiConnectionFactory**](MsgVpnJndiConnectionFactory.md)| The JNDI Connection Factory object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **connectionFactoryName** | **string**| The name of the JMS Connection Factory. | 
 **optional** | ***MsgVpnApiUpdateMsgVpnJndiConnectionFactoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiUpdateMsgVpnJndiConnectionFactoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnJndiConnectionFactoryResponse**](MsgVpnJndiConnectionFactoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnJndiQueue**
> MsgVpnJndiQueueResponse UpdateMsgVpnJndiQueue(ctx, body, msgVpnName, queueName, optional)
Update a JNDI Queue object.

Update a JNDI Queue object. Any attribute missing from the request will be left unchanged.  The message broker provides an internal JNDI store for provisioned Queue objects that clients can access through JNDI lookups.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- msgVpnName|x|x||||| queueName|x|x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.2.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnJndiQueue**](MsgVpnJndiQueue.md)| The JNDI Queue object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The JNDI name of the JMS Queue. | 
 **optional** | ***MsgVpnApiUpdateMsgVpnJndiQueueOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiUpdateMsgVpnJndiQueueOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnJndiQueueResponse**](MsgVpnJndiQueueResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnJndiTopic**
> MsgVpnJndiTopicResponse UpdateMsgVpnJndiTopic(ctx, body, msgVpnName, topicName, optional)
Update a JNDI Topic object.

Update a JNDI Topic object. Any attribute missing from the request will be left unchanged.  The message broker provides an internal JNDI store for provisioned Topic objects that clients can access through JNDI lookups.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- msgVpnName|x|x||||| topicName|x|x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.2.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnJndiTopic**](MsgVpnJndiTopic.md)| The JNDI Topic object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicName** | **string**| The JNDI name of the JMS Topic. | 
 **optional** | ***MsgVpnApiUpdateMsgVpnJndiTopicOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiUpdateMsgVpnJndiTopicOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnJndiTopicResponse**](MsgVpnJndiTopicResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnMqttRetainCache**
> MsgVpnMqttRetainCacheResponse UpdateMsgVpnMqttRetainCache(ctx, body, msgVpnName, cacheName, optional)
Update an MQTT Retain Cache object.

Update an MQTT Retain Cache object. Any attribute missing from the request will be left unchanged.  Using MQTT retained messages allows publishing MQTT clients to indicate that a message must be stored for later delivery to subscribing clients when those subscribing clients add subscriptions matching the retained message's topic. An MQTT Retain Cache processes all retained messages for a Message VPN.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- cacheName|x|x||||| msgVpnName|x|x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnMqttRetainCache**](MsgVpnMqttRetainCache.md)| The MQTT Retain Cache object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **cacheName** | **string**| The name of the MQTT Retain Cache. | 
 **optional** | ***MsgVpnApiUpdateMsgVpnMqttRetainCacheOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiUpdateMsgVpnMqttRetainCacheOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnMqttRetainCacheResponse**](MsgVpnMqttRetainCacheResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnMqttSession**
> MsgVpnMqttSessionResponse UpdateMsgVpnMqttSession(ctx, body, msgVpnName, mqttSessionClientId, mqttSessionVirtualRouter, optional)
Update an MQTT Session object.

Update an MQTT Session object. Any attribute missing from the request will be left unchanged.  An MQTT Session object is a virtual representation of an MQTT client connection. An MQTT session holds the state of an MQTT client (that is, it is used to contain a client's QoS 0 and QoS 1 subscription sets and any undelivered QoS 1 messages).   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- mqttSessionClientId|x|x||||| mqttSessionVirtualRouter|x|x||||| msgVpnName|x|x||||| owner|||||x|| queueRejectMsgToSenderOnDiscardBehavior|||||x||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThreshold|clearPercent|setPercent|clearValue, setValue EventThreshold|clearValue|setValue|clearPercent, setPercent EventThreshold|setPercent|clearPercent|clearValue, setValue EventThreshold|setValue|clearValue|clearPercent, setPercent    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.1.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnMqttSession**](MsgVpnMqttSession.md)| The MQTT Session object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **mqttSessionClientId** | **string**| The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet. | 
  **mqttSessionVirtualRouter** | **string**| The virtual router of the MQTT Session. | 
 **optional** | ***MsgVpnApiUpdateMsgVpnMqttSessionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiUpdateMsgVpnMqttSessionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnMqttSessionResponse**](MsgVpnMqttSessionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnMqttSessionSubscription**
> MsgVpnMqttSessionSubscriptionResponse UpdateMsgVpnMqttSessionSubscription(ctx, body, msgVpnName, mqttSessionClientId, mqttSessionVirtualRouter, subscriptionTopic, optional)
Update a Subscription object.

Update a Subscription object. Any attribute missing from the request will be left unchanged.  An MQTT session contains a client's QoS 0 and QoS 1 subscription sets. On creation, a subscription defaults to QoS 0.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- mqttSessionClientId|x|x||||| mqttSessionVirtualRouter|x|x||||| msgVpnName|x|x||||| subscriptionTopic|x|x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.1.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnMqttSessionSubscription**](MsgVpnMqttSessionSubscription.md)| The Subscription object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **mqttSessionClientId** | **string**| The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet. | 
  **mqttSessionVirtualRouter** | **string**| The virtual router of the MQTT Session. | 
  **subscriptionTopic** | **string**| The MQTT subscription topic. | 
 **optional** | ***MsgVpnApiUpdateMsgVpnMqttSessionSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiUpdateMsgVpnMqttSessionSubscriptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnMqttSessionSubscriptionResponse**](MsgVpnMqttSessionSubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnQueue**
> MsgVpnQueueResponse UpdateMsgVpnQueue(ctx, body, msgVpnName, queueName, optional)
Update a Queue object.

Update a Queue object. Any attribute missing from the request will be left unchanged.  A Queue acts as both a destination that clients can publish messages to, and as an endpoint that clients can bind consumers to and consume messages from.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- accessType|||||x|| msgVpnName|x|x||||| owner|||||x|| permission|||||x|| queueName|x|x||||| redeliveryDelayEnabled|||||x|| redeliveryDelayInitialInterval|||||x|| redeliveryDelayMaxInterval|||||x|| redeliveryDelayMultiplier|||||x|| rejectMsgToSenderOnDiscardBehavior|||||x|| respectMsgPriorityEnabled|||||x||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThreshold|clearPercent|setPercent|clearValue, setValue EventThreshold|clearValue|setValue|clearPercent, setPercent EventThreshold|setPercent|clearPercent|clearValue, setValue EventThreshold|setValue|clearValue|clearPercent, setPercent    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnQueue**](MsgVpnQueue.md)| The Queue object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueName** | **string**| The name of the Queue. | 
 **optional** | ***MsgVpnApiUpdateMsgVpnQueueOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiUpdateMsgVpnQueueOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueueResponse**](MsgVpnQueueResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnQueueTemplate**
> MsgVpnQueueTemplateResponse UpdateMsgVpnQueueTemplate(ctx, body, msgVpnName, queueTemplateName, optional)
Update a Queue Template object.

Update a Queue Template object. Any attribute missing from the request will be left unchanged.  A Queue Template provides a mechanism for specifying the initial state for client created queues.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- msgVpnName|x|x||||| queueTemplateName|x|x||||| rejectMsgToSenderOnDiscardBehavior|||||x||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThreshold|clearPercent|setPercent|clearValue, setValue EventThreshold|clearValue|setValue|clearPercent, setPercent EventThreshold|setPercent|clearPercent|clearValue, setValue EventThreshold|setValue|clearValue|clearPercent, setPercent    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnQueueTemplate**](MsgVpnQueueTemplate.md)| The Queue Template object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **queueTemplateName** | **string**| The name of the Queue Template. | 
 **optional** | ***MsgVpnApiUpdateMsgVpnQueueTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiUpdateMsgVpnQueueTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnQueueTemplateResponse**](MsgVpnQueueTemplateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnReplayLog**
> MsgVpnReplayLogResponse UpdateMsgVpnReplayLog(ctx, body, msgVpnName, replayLogName, optional)
Update a Replay Log object.

Update a Replay Log object. Any attribute missing from the request will be left unchanged.  When the Message Replay feature is enabled, message brokers store persistent messages in a Replay Log. These messages are kept until the log is full, after which the oldest messages are removed to free up space for new messages.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- msgVpnName|x|x||||| replayLogName|x|x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.10.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnReplayLog**](MsgVpnReplayLog.md)| The Replay Log object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **replayLogName** | **string**| The name of the Replay Log. | 
 **optional** | ***MsgVpnApiUpdateMsgVpnReplayLogOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiUpdateMsgVpnReplayLogOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnReplayLogResponse**](MsgVpnReplayLogResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnReplicatedTopic**
> MsgVpnReplicatedTopicResponse UpdateMsgVpnReplicatedTopic(ctx, body, msgVpnName, replicatedTopic, optional)
Update a Replicated Topic object.

Update a Replicated Topic object. Any attribute missing from the request will be left unchanged.  To indicate which messages should be replicated between the active and standby site, a Replicated Topic subscription must be configured on a Message VPN. If a published message matches both a replicated topic and an endpoint on the active site, then the message is replicated to the standby site.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- msgVpnName|x|x||||| replicatedTopic|x|x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.1.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnReplicatedTopic**](MsgVpnReplicatedTopic.md)| The Replicated Topic object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **replicatedTopic** | **string**| The topic for applying replication. Published messages matching this topic will be replicated to the standby site. | 
 **optional** | ***MsgVpnApiUpdateMsgVpnReplicatedTopicOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiUpdateMsgVpnReplicatedTopicOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnReplicatedTopicResponse**](MsgVpnReplicatedTopicResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnRestDeliveryPoint**
> MsgVpnRestDeliveryPointResponse UpdateMsgVpnRestDeliveryPoint(ctx, body, msgVpnName, restDeliveryPointName, optional)
Update a REST Delivery Point object.

Update a REST Delivery Point object. Any attribute missing from the request will be left unchanged.  A REST Delivery Point manages delivery of messages from queues to a named list of REST Consumers.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- clientProfileName|||||x|| msgVpnName|x|x||||| restDeliveryPointName|x|x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnRestDeliveryPoint**](MsgVpnRestDeliveryPoint.md)| The REST Delivery Point object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
 **optional** | ***MsgVpnApiUpdateMsgVpnRestDeliveryPointOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiUpdateMsgVpnRestDeliveryPointOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointResponse**](MsgVpnRestDeliveryPointResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnRestDeliveryPointQueueBinding**
> MsgVpnRestDeliveryPointQueueBindingResponse UpdateMsgVpnRestDeliveryPointQueueBinding(ctx, body, msgVpnName, restDeliveryPointName, queueBindingName, optional)
Update a Queue Binding object.

Update a Queue Binding object. Any attribute missing from the request will be left unchanged.  A Queue Binding for a REST Delivery Point attracts messages to be delivered to REST consumers. If the queue does not exist it can be created subsequently, and once the queue is operational the broker performs the queue binding. Removing the queue binding does not delete the queue itself. Similarly, removing the queue does not remove the queue binding, which fails until the queue is recreated or the queue binding is deleted.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- msgVpnName|x|x||||| queueBindingName|x|x||||| restDeliveryPointName|x|x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnRestDeliveryPointQueueBinding**](MsgVpnRestDeliveryPointQueueBinding.md)| The Queue Binding object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **queueBindingName** | **string**| The name of a queue in the Message VPN. | 
 **optional** | ***MsgVpnApiUpdateMsgVpnRestDeliveryPointQueueBindingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiUpdateMsgVpnRestDeliveryPointQueueBindingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointQueueBindingResponse**](MsgVpnRestDeliveryPointQueueBindingResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader**
> MsgVpnRestDeliveryPointQueueBindingProtectedRequestHeaderResponse UpdateMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader(ctx, body, msgVpnName, restDeliveryPointName, queueBindingName, headerName, optional)
Update a Protected Request Header object.

Update a Protected Request Header object. Any attribute missing from the request will be left unchanged.  A protected request header to be added to the HTTP request. Unlike a non-protected request header, the header value cannot be displayed after it is set.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- headerName|x|x||||| headerValue|||x||||x msgVpnName|x|x||||| queueBindingName|x|x||||| restDeliveryPointName|x|x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.30.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader**](MsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader.md)| The Protected Request Header object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **queueBindingName** | **string**| The name of a queue in the Message VPN. | 
  **headerName** | **string**| The name of the protected HTTP request header. | 
 **optional** | ***MsgVpnApiUpdateMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeaderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiUpdateMsgVpnRestDeliveryPointQueueBindingProtectedRequestHeaderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointQueueBindingProtectedRequestHeaderResponse**](MsgVpnRestDeliveryPointQueueBindingProtectedRequestHeaderResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnRestDeliveryPointQueueBindingRequestHeader**
> MsgVpnRestDeliveryPointQueueBindingRequestHeaderResponse UpdateMsgVpnRestDeliveryPointQueueBindingRequestHeader(ctx, body, msgVpnName, restDeliveryPointName, queueBindingName, headerName, optional)
Update a Request Header object.

Update a Request Header object. Any attribute missing from the request will be left unchanged.  A request header to be added to the HTTP request.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- headerName|x|x||||| msgVpnName|x|x||||| queueBindingName|x|x||||| restDeliveryPointName|x|x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.23.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnRestDeliveryPointQueueBindingRequestHeader**](MsgVpnRestDeliveryPointQueueBindingRequestHeader.md)| The Request Header object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **queueBindingName** | **string**| The name of a queue in the Message VPN. | 
  **headerName** | **string**| The name of the HTTP request header. | 
 **optional** | ***MsgVpnApiUpdateMsgVpnRestDeliveryPointQueueBindingRequestHeaderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiUpdateMsgVpnRestDeliveryPointQueueBindingRequestHeaderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointQueueBindingRequestHeaderResponse**](MsgVpnRestDeliveryPointQueueBindingRequestHeaderResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnRestDeliveryPointRestConsumer**
> MsgVpnRestDeliveryPointRestConsumerResponse UpdateMsgVpnRestDeliveryPointRestConsumer(ctx, body, msgVpnName, restDeliveryPointName, restConsumerName, optional)
Update a REST Consumer object.

Update a REST Consumer object. Any attribute missing from the request will be left unchanged.  REST Consumer objects establish HTTP connectivity to REST consumer applications who wish to receive messages from a broker.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- authenticationAwsSecretAccessKey|||x||||x authenticationClientCertContent|||x||x||x authenticationClientCertPassword|||x||x|| authenticationHttpBasicPassword|||x||x||x authenticationHttpBasicUsername|||||x|| authenticationHttpHeaderValue|||x||||x authenticationOauthClientId|||||x|| authenticationOauthClientScope|||||x|| authenticationOauthClientSecret|||x||x||x authenticationOauthClientTokenEndpoint|||||x|| authenticationOauthClientTokenExpiryDefault|||||x|| authenticationOauthJwtSecretKey|||x||x||x authenticationOauthJwtTokenEndpoint|||||x|| authenticationOauthJwtTokenExpiryDefault|||||x|| authenticationScheme|||||x|| msgVpnName|x|x||||| outgoingConnectionCount|||||x|| remoteHost|||||x|| remotePort|||||x|| restConsumerName|x|x||||| restDeliveryPointName|x|x||||| tlsCipherSuiteList|||||x|| tlsEnabled|||||x||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- MsgVpnRestDeliveryPointRestConsumer|authenticationClientCertPassword|authenticationClientCertContent| MsgVpnRestDeliveryPointRestConsumer|authenticationHttpBasicPassword|authenticationHttpBasicUsername| MsgVpnRestDeliveryPointRestConsumer|authenticationHttpBasicUsername|authenticationHttpBasicPassword| MsgVpnRestDeliveryPointRestConsumer|remotePort|tlsEnabled| MsgVpnRestDeliveryPointRestConsumer|tlsEnabled|remotePort|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.0.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnRestDeliveryPointRestConsumer**](MsgVpnRestDeliveryPointRestConsumer.md)| The REST Consumer object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **restDeliveryPointName** | **string**| The name of the REST Delivery Point. | 
  **restConsumerName** | **string**| The name of the REST Consumer. | 
 **optional** | ***MsgVpnApiUpdateMsgVpnRestDeliveryPointRestConsumerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiUpdateMsgVpnRestDeliveryPointRestConsumerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnRestDeliveryPointRestConsumerResponse**](MsgVpnRestDeliveryPointRestConsumerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnTelemetryProfile**
> MsgVpnTelemetryProfileResponse UpdateMsgVpnTelemetryProfile(ctx, body, msgVpnName, telemetryProfileName, optional)
Update a Telemetry Profile object.

Update a Telemetry Profile object. Any attribute missing from the request will be left unchanged.  Using the Telemetry Profile allows trace spans to be generated as messages are processed by the broker. The generated spans are stored persistently on the broker and may be consumed by the Solace receiver component of an OpenTelemetry Collector.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- msgVpnName|x|x||||| telemetryProfileName|x|x|||||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThreshold|clearPercent|setPercent|clearValue, setValue EventThreshold|clearValue|setValue|clearPercent, setPercent EventThreshold|setPercent|clearPercent|clearValue, setValue EventThreshold|setValue|clearValue|clearPercent, setPercent    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.31.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnTelemetryProfile**](MsgVpnTelemetryProfile.md)| The Telemetry Profile object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **telemetryProfileName** | **string**| The name of the Telemetry Profile. | 
 **optional** | ***MsgVpnApiUpdateMsgVpnTelemetryProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiUpdateMsgVpnTelemetryProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTelemetryProfileResponse**](MsgVpnTelemetryProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnTelemetryProfileTraceFilter**
> MsgVpnTelemetryProfileTraceFilterResponse UpdateMsgVpnTelemetryProfileTraceFilter(ctx, body, msgVpnName, telemetryProfileName, traceFilterName, optional)
Update a Trace Filter object.

Update a Trace Filter object. Any attribute missing from the request will be left unchanged.  A Trace Filter controls which messages received by the broker will be traced. If an incoming message matches an enabled tracing filter's subscription, the message will be traced as it passes through the broker.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- msgVpnName|x|x||||| telemetryProfileName|x|x||||| traceFilterName|x|x|||||    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.31.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnTelemetryProfileTraceFilter**](MsgVpnTelemetryProfileTraceFilter.md)| The Trace Filter object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **telemetryProfileName** | **string**| The name of the Telemetry Profile. | 
  **traceFilterName** | **string**| A name used to identify the trace filter. Consider a name that describes the subscriptions contained within the filter, such as the name of the application and/or the scenario in which the trace filter might be enabled, such as \&quot;appNameDebug\&quot;. | 
 **optional** | ***MsgVpnApiUpdateMsgVpnTelemetryProfileTraceFilterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiUpdateMsgVpnTelemetryProfileTraceFilterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTelemetryProfileTraceFilterResponse**](MsgVpnTelemetryProfileTraceFilterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnTopicEndpoint**
> MsgVpnTopicEndpointResponse UpdateMsgVpnTopicEndpoint(ctx, body, msgVpnName, topicEndpointName, optional)
Update a Topic Endpoint object.

Update a Topic Endpoint object. Any attribute missing from the request will be left unchanged.  A Topic Endpoint attracts messages published to a topic for which the Topic Endpoint has a matching topic subscription. The topic subscription for the Topic Endpoint is specified in the client request to bind a Flow to that Topic Endpoint. Queues are significantly more flexible than Topic Endpoints and are the recommended approach for most applications. The use of Topic Endpoints should be restricted to JMS applications.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- accessType|||||x|| msgVpnName|x|x||||| owner|||||x|| permission|||||x|| redeliveryDelayEnabled|||||x|| redeliveryDelayInitialInterval|||||x|| redeliveryDelayMaxInterval|||||x|| redeliveryDelayMultiplier|||||x|| rejectMsgToSenderOnDiscardBehavior|||||x|| respectMsgPriorityEnabled|||||x|| topicEndpointName|x|x|||||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThreshold|clearPercent|setPercent|clearValue, setValue EventThreshold|clearValue|setValue|clearPercent, setPercent EventThreshold|setPercent|clearPercent|clearValue, setValue EventThreshold|setValue|clearValue|clearPercent, setPercent    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.1.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnTopicEndpoint**](MsgVpnTopicEndpoint.md)| The Topic Endpoint object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointName** | **string**| The name of the Topic Endpoint. | 
 **optional** | ***MsgVpnApiUpdateMsgVpnTopicEndpointOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiUpdateMsgVpnTopicEndpointOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointResponse**](MsgVpnTopicEndpointResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgVpnTopicEndpointTemplate**
> MsgVpnTopicEndpointTemplateResponse UpdateMsgVpnTopicEndpointTemplate(ctx, body, msgVpnName, topicEndpointTemplateName, optional)
Update a Topic Endpoint Template object.

Update a Topic Endpoint Template object. Any attribute missing from the request will be left unchanged.  A Topic Endpoint Template provides a mechanism for specifying the initial state for client created topic endpoints.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- msgVpnName|x|x||||| rejectMsgToSenderOnDiscardBehavior|||||x|| topicEndpointTemplateName|x|x|||||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThreshold|clearPercent|setPercent|clearValue, setValue EventThreshold|clearValue|setValue|clearPercent, setPercent EventThreshold|setPercent|clearPercent|clearValue, setValue EventThreshold|setValue|clearValue|clearPercent, setPercent    A SEMP client authorized with a minimum access scope/level of \"vpn/read-write\" is required to perform this operation.  This has been available since 2.14.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MsgVpnTopicEndpointTemplate**](MsgVpnTopicEndpointTemplate.md)| The Topic Endpoint Template object&#x27;s attributes. | 
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **topicEndpointTemplateName** | **string**| The name of the Topic Endpoint Template. | 
 **optional** | ***MsgVpnApiUpdateMsgVpnTopicEndpointTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MsgVpnApiUpdateMsgVpnTopicEndpointTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTopicEndpointTemplateResponse**](MsgVpnTopicEndpointTemplateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

