# Go API client for action

SEMP (starting in `v2`, see note 1) is a RESTful API for configuring, monitoring, and administering a Solace PubSub+ broker.  SEMP uses URIs to address manageable **resources** of the Solace PubSub+ broker. Resources are individual **objects**, **collections** of objects, or (exclusively in the action API) **actions**. This document applies to the following API:   API|Base Path|Purpose|Comments :---|:---|:---|:--- Action|/SEMP/v2/action|Performing actions|See note 2    The following APIs are also available:   API|Base Path|Purpose|Comments :---|:---|:---|:--- Configuration|/SEMP/v2/config|Reading and writing config state|See note 2 Monitoring|/SEMP/v2/monitor|Querying operational parameters|See note 2    Resources are always nouns, with individual objects being singular and collections being plural.  Objects within a collection are identified by an `obj-id`, which follows the collection name with the form `collection-name/obj-id`.  Actions within an object are identified by an `action-id`, which follows the object name with the form `obj-id/action-id`.  Some examples:  ``` /SEMP/v2/config/msgVpns                        ; MsgVpn collection /SEMP/v2/config/msgVpns/a                      ; MsgVpn object named \"a\" /SEMP/v2/config/msgVpns/a/queues               ; Queue collection in MsgVpn \"a\" /SEMP/v2/config/msgVpns/a/queues/b             ; Queue object named \"b\" in MsgVpn \"a\" /SEMP/v2/action/msgVpns/a/queues/b/startReplay ; Action that starts a replay on Queue \"b\" in MsgVpn \"a\" /SEMP/v2/monitor/msgVpns/a/clients             ; Client collection in MsgVpn \"a\" /SEMP/v2/monitor/msgVpns/a/clients/c           ; Client object named \"c\" in MsgVpn \"a\" ```  ## Collection Resources  Collections are unordered lists of objects (unless described as otherwise), and are described by JSON arrays. Each item in the array represents an object in the same manner as the individual object would normally be represented. In the configuration API, the creation of a new object is done through its collection resource.  ## Object and Action Resources  Objects are composed of attributes, actions, collections, and other objects. They are described by JSON objects as name/value pairs. The collections and actions of an object are not contained directly in the object's JSON content; rather the content includes an attribute containing a URI which points to the collections and actions. These contained resources must be managed through this URI. At a minimum, every object has one or more identifying attributes, and its own `uri` attribute which contains the URI pointing to itself.  Actions are also composed of attributes, and are described by JSON objects as name/value pairs. Unlike objects, however, they are not members of a collection and cannot be retrieved, only performed. Actions only exist in the action API.  Attributes in an object or action may have any combination of the following properties:   Property|Meaning|Comments :---|:---|:--- Identifying|Attribute is involved in unique identification of the object, and appears in its URI| Const|Attribute value can only be chosen during object creation| Required|Attribute must be provided in the request| Read-Only|Attribute can only be read, not written.|See note 3 Write-Only|Attribute can only be written, not read, unless the attribute is also opaque|See the documentation for the opaque property Requires-Disable|Attribute cannot be changed while the object (or the relevant part of the object) is administratively enabled| Auto-Disable|Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as one or more attributes will be temporarily disabled to apply the change| Deprecated|Attribute is deprecated, and will disappear in the next SEMP version| Opaque|Attribute can be set or retrieved in opaque form when the `opaquePassword` query parameter is present|See the `opaquePassword` query parameter documentation    In some requests, certain attributes may only be provided in certain combinations with other attributes:   Relationship|Meaning :---|:--- Requires|Attribute may only be changed by a request if a particular attribute or combination of attributes is also provided in the request Conflicts|Attribute may only be provided in a request if a particular attribute or combination of attributes is not also provided in the request    In the monitoring API, any non-identifying attribute may not be returned in a GET.  ## HTTP Methods  The following HTTP methods manipulate resources in accordance with these general principles. Note that some methods are only used in certain APIs:   Method|Resource|Meaning|Request Body|Response Body|Notes :---|:---|:---|:---|:---|:--- POST|Collection|Create object|Initial attribute values|Object attributes and metadata|Absent attributes are set to default. If object already exists, a 400 error is returned PUT|Object|Update object|New attribute values|Object attributes and metadata|If does not exist, the object is first created. Absent attributes are set to default, with certain exceptions (see note 4) PUT|Action|Performs action|Action arguments|Action metadata| PATCH|Object|Update object|New attribute values|Object attributes and metadata|Absent attributes are left unchanged. If the object does not exist, a 404 error is returned DELETE|Object|Delete object|Empty|Object metadata|If the object does not exist, a 404 is returned GET|Object|Get object|Empty|Object attributes and metadata|If the object does not exist, a 404 is returned GET|Collection|Get collection|Empty|Object attributes and collection metadata|If the collection is empty, then an empty collection is returned with a 200 code    ## Common Query Parameters  The following are some common query parameters that are supported by many method/URI combinations. Individual URIs may document additional parameters. Note that multiple query parameters can be used together in a single URI, separated by the ampersand character. For example:  ``` ; Request for the MsgVpns collection using two hypothetical query parameters ; \"q1\" and \"q2\" with values \"val1\" and \"val2\" respectively /SEMP/v2/action/msgVpns?q1=val1&q2=val2 ```  ### select  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. Use this query parameter to limit the size of the returned data for each returned object, return only those fields that are desired, or exclude fields that are not desired.  The value of `select` is a comma-separated list of attribute names. If the list contains attribute names that are not prefaced by `-`, only those attributes are included in the response. If the list contains attribute names that are prefaced by `-`, those attributes are excluded from the response. If the list contains both types, then the difference of the first set of attributes and the second set of attributes is returned. If the list is empty (i.e. `select=`), it is treated the same as if no `select` was provided: all attribute are returned.  All attributes that are prefaced by `-` must follow all attributes that are not prefaced by `-`. In addition, each attribute name in the list must match at least one attribute in the object.  Names may include the `*` wildcard (zero or more characters). Nested attribute names are supported using periods (e.g. `parentName.childName`).  Some examples:  ``` ; List of all MsgVpn names /SEMP/v2/action/msgVpns?select=msgVpnName ; List of all MsgVpn and their attributes except for their names /SEMP/v2/action/msgVpns?select=-msgVpnName ; Authentication attributes of MsgVpn \"finance\" /SEMP/v2/action/msgVpns/finance?select=authentication%2A ; All attributes of MsgVpn \"finance\" except for authentication attributes /SEMP/v2/action/msgVpns/finance?select=-authentication%2A ; Access related attributes of Queue \"orderQ\" of MsgVpn \"finance\" /SEMP/v2/action/msgVpns/finance/queues/orderQ?select=owner,permission ```  ### where  Include in the response only objects where certain conditions are true. Use this query parameter to limit which objects are returned to those whose attribute values meet the given conditions.  The value of `where` is a comma-separated list of expressions. All expressions must be true for the object to be included in the response. Each expression takes the form:  ``` expression  = attribute-name OP value OP          = '==' | '!=' | '<' | '>' | '<=' | '>=' ```  `value` may be a number, string, `true`, or `false`, as appropriate for the type of `attribute-name`. Greater-than and less-than comparisons only work for numbers. A `*` in a string `value` is interpreted as a wildcard (zero or more characters). Some examples:  ``` ; Only enabled MsgVpns /SEMP/v2/action/msgVpns?where=enabled%3D%3Dtrue ; Only MsgVpns using basic non-LDAP authentication /SEMP/v2/action/msgVpns?where=authenticationBasicEnabled%3D%3Dtrue,authenticationBasicType%21%3Dldap ; Only MsgVpns that allow more than 100 client connections /SEMP/v2/action/msgVpns?where=maxConnectionCount%3E100 ; Only MsgVpns with msgVpnName starting with \"B\": /SEMP/v2/action/msgVpns?where=msgVpnName%3D%3DB%2A ```  ### count  Limit the count of objects in the response. This can be useful to limit the size of the response for large collections. The minimum value for `count` is `1` and the default is `10`. There is also a per-collection maximum value to limit request handling time.  `count` does not guarantee that a minimum number of objects will be returned. A page may contain fewer than `count` objects or even be empty. Additional objects may nonetheless be available for retrieval on subsequent pages. See the `cursor` query parameter documentation for more information on paging.  For example: ``` ; Up to 25 MsgVpns /SEMP/v2/action/msgVpns?count=25 ```  ### cursor  The cursor, or position, for the next page of objects. Cursors are opaque data that should not be created or interpreted by SEMP clients, and should only be used as described below.  When a request is made for a collection and there may be additional objects available for retrieval that are not included in the initial response, the response will include a `cursorQuery` field containing a cursor. The value of this field can be specified in the `cursor` query parameter of a subsequent request to retrieve the next page of objects.  Applications must continue to use the `cursorQuery` if one is provided in order to retrieve the full set of objects associated with the request, even if a page contains fewer than the requested number of objects (see the `count` query parameter documentation) or is empty.  ### opaquePassword  Attributes with the opaque property are also write-only and so cannot normally be retrieved in a GET. However, when a password is provided in the `opaquePassword` query parameter, attributes with the opaque property are retrieved in a GET in opaque form, encrypted with this password. The query parameter can also be used on a POST, PATCH, or PUT to set opaque attributes using opaque attribute values retrieved in a GET, so long as:  1. the same password that was used to retrieve the opaque attribute values is provided; and  2. the broker to which the request is being sent has the same major and minor SEMP version as the broker that produced the opaque attribute values.  The password provided in the query parameter must be a minimum of 8 characters and a maximum of 128 characters.  The query parameter can only be used in the configuration API, and only over HTTPS.  ## Authentication  When a client makes its first SEMPv2 request, it must supply a username and password using HTTP Basic authentication, or an OAuth token or tokens using HTTP Bearer authentication.  When HTTP Basic authentication is used, the broker returns a cookie containing a session key. The client can omit the username and password from subsequent requests, because the broker can use the session cookie for authentication instead. When the session expires or is deleted, the client must provide the username and password again, and the broker creates a new session.  There are a limited number of session slots available on the broker. The broker returns 529 No SEMP Session Available if it is not able to allocate a session.  If certain attributes—such as a user's password—are changed, the broker automatically deletes the affected sessions. These attributes are documented below. However, changes in external user configuration data stored on a RADIUS or LDAP server do not trigger the broker to delete the associated session(s), therefore you must do this manually, if required.  A client can retrieve its current session information using the /about/user endpoint and delete its own session using the /about/user/logout endpoint. A client with appropriate permissions can also manage all sessions using the /sessions endpoint.  Sessions are not created when authenticating with an OAuth token or tokens using HTTP Bearer authentication. If a session cookie is provided, it is ignored.  ## Help  Visit [our website](https://solace.com) to learn more about Solace.  You can also download the SEMP API specifications by clicking [here](https://solace.com/downloads/).  If you need additional support, please contact us at [support@solace.com](mailto:support@solace.com).  ## Notes  Note|Description :---:|:--- 1|This specification defines SEMP starting in \"v2\", and not the original SEMP \"v1\" interface. Request and response formats between \"v1\" and \"v2\" are entirely incompatible, although both protocols share a common port configuration on the Solace PubSub+ broker. They are differentiated by the initial portion of the URI path, one of either \"/SEMP/\" or \"/SEMP/v2/\" 2|This API is partially implemented. Only a subset of all objects are available. 3|Read-only attributes may appear in POST and PUT/PATCH requests. However, if a read-only attribute is not marked as identifying, it will be ignored during a PUT/PATCH. 4|On a PUT, if the SEMP user is not authorized to modify the attribute, its value is left unchanged rather than set to default. In addition, the values of write-only attributes are not set to their defaults on a PUT, except in the following two cases: there is a mutual requires relationship with another non-write-only attribute, both attributes are absent from the request, and the non-write-only attribute is not currently set to its default value; or the attribute is also opaque and the `opaquePassword` query parameter is provided in the request.  

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 2.34
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen
For more information, please visit [http://www.solace.com](http://www.solace.com)

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./action"
```

## Documentation for API Endpoints

All URIs are relative to *http://www.solace.com/SEMP/v2/action*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AboutApi* | [**DoAboutUserLogout**](docs/AboutApi.md#doaboutuserlogout) | **Put** /about/user/logout | Logout of the current session.
*AboutApi* | [**GetAbout**](docs/AboutApi.md#getabout) | **Get** /about | Get an About object.
*AboutApi* | [**GetAboutApi**](docs/AboutApi.md#getaboutapi) | **Get** /about/api | Get an API Description object.
*AboutApi* | [**GetAboutUser**](docs/AboutApi.md#getaboutuser) | **Get** /about/user | Get a User object.
*AboutApi* | [**GetAboutUserMsgVpn**](docs/AboutApi.md#getaboutusermsgvpn) | **Get** /about/user/msgVpns/{msgVpnName} | Get a User Message VPN object.
*AboutApi* | [**GetAboutUserMsgVpns**](docs/AboutApi.md#getaboutusermsgvpns) | **Get** /about/user/msgVpns | Get a list of User Message VPN objects.
*AllApi* | [**DoAboutUserLogout**](docs/AllApi.md#doaboutuserlogout) | **Put** /about/user/logout | Logout of the current session.
*AllApi* | [**DoCertAuthorityRefreshCrl**](docs/AllApi.md#docertauthorityrefreshcrl) | **Put** /certAuthorities/{certAuthorityName}/refreshCrl | Refresh the CRL file for the Certificate Authority.
*AllApi* | [**DoClientCertAuthorityRefreshCrl**](docs/AllApi.md#doclientcertauthorityrefreshcrl) | **Put** /clientCertAuthorities/{certAuthorityName}/refreshCrl | Refresh the CRL file for the Client Certificate Authority.
*AllApi* | [**DoConfigSyncAssertLeaderMsgVpn**](docs/AllApi.md#doconfigsyncassertleadermsgvpn) | **Put** /configSyncAssertLeaderMsgVpn | Assert leadership of the specified Config Sync table, forcing any other leader&#x27;s content to be overwritten with our own. Use whenever a High Availability pair fall out of sync. Config Sync must be a leader for the selected table.
*AllApi* | [**DoConfigSyncAssertLeaderRouter**](docs/AllApi.md#doconfigsyncassertleaderrouter) | **Put** /configSyncAssertLeaderRouter | Assert leadership of the specified Config Sync table, forcing any other leader&#x27;s content to be overwritten with our own. Use whenever a High Availability pair fall out of sync. Config Sync must be a leader for the selected table.
*AllApi* | [**DoConfigSyncResyncFollowerMsgVpn**](docs/AllApi.md#doconfigsyncresyncfollowermsgvpn) | **Put** /configSyncResyncFollowerMsgVpn | Resync the selected Config Sync table, forcing this follower&#x27;s content to be overwritten with that from a leader. Config Sync must be a follower for the selected table.
*AllApi* | [**DoConfigSyncResyncLeaderMsgVpn**](docs/AllApi.md#doconfigsyncresyncleadermsgvpn) | **Put** /configSyncResyncLeaderMsgVpn | Resync the selected Config Sync table, forcing this leader&#x27;s content to be overwritten with that from a leader. Config Sync must be a leader for the selected table.
*AllApi* | [**DoConfigSyncResyncLeaderRouter**](docs/AllApi.md#doconfigsyncresyncleaderrouter) | **Put** /configSyncResyncLeaderRouter | Resync the selected Config Sync table, forcing this leader&#x27;s content to be overwritten with that from a leader. Config Sync must be a leader for the selected table.
*AllApi* | [**DoGuaranteedMsgingDefragmentMsgSpoolFilesStart**](docs/AllApi.md#doguaranteedmsgingdefragmentmsgspoolfilesstart) | **Put** /guaranteedMsgingDefragmentMsgSpoolFilesStart | Start a spool file defragmentation run.
*AllApi* | [**DoGuaranteedMsgingDefragmentMsgSpoolFilesStop**](docs/AllApi.md#doguaranteedmsgingdefragmentmsgspoolfilesstop) | **Put** /guaranteedMsgingDefragmentMsgSpoolFilesStop | Stop a spool file defragmentation run.
*AllApi* | [**DoMsgVpnAuthenticationOauthProfileClearStats**](docs/AllApi.md#domsgvpnauthenticationoauthprofileclearstats) | **Put** /msgVpns/{msgVpnName}/authenticationOauthProfiles/{oauthProfileName}/clearStats | Clear the statistics for the OAuth Profile.
*AllApi* | [**DoMsgVpnAuthenticationOauthProviderClearStats**](docs/AllApi.md#domsgvpnauthenticationoauthproviderclearstats) | **Put** /msgVpns/{msgVpnName}/authenticationOauthProviders/{oauthProviderName}/clearStats | Clear the statistics for the OAuth Provider.
*AllApi* | [**DoMsgVpnBridgeClearEvent**](docs/AllApi.md#domsgvpnbridgeclearevent) | **Put** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/clearEvent | Clear an event for the Bridge so it can be generated anew.
*AllApi* | [**DoMsgVpnBridgeClearStats**](docs/AllApi.md#domsgvpnbridgeclearstats) | **Put** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/clearStats | Clear the statistics for the Bridge.
*AllApi* | [**DoMsgVpnBridgeDisconnect**](docs/AllApi.md#domsgvpnbridgedisconnect) | **Put** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/disconnect | Disconnect the Bridge.
*AllApi* | [**DoMsgVpnClearMsgSpoolStats**](docs/AllApi.md#domsgvpnclearmsgspoolstats) | **Put** /msgVpns/{msgVpnName}/clearMsgSpoolStats | Clear the message spool statistics for the Message VPN.
*AllApi* | [**DoMsgVpnClearReplicationStats**](docs/AllApi.md#domsgvpnclearreplicationstats) | **Put** /msgVpns/{msgVpnName}/clearReplicationStats | Clear the replication statistics for the Message VPN.
*AllApi* | [**DoMsgVpnClearServiceStats**](docs/AllApi.md#domsgvpnclearservicestats) | **Put** /msgVpns/{msgVpnName}/clearServiceStats | Clear the service statistics for the Message VPN.
*AllApi* | [**DoMsgVpnClearStats**](docs/AllApi.md#domsgvpnclearstats) | **Put** /msgVpns/{msgVpnName}/clearStats | Clear the client statistics for the Message VPN.
*AllApi* | [**DoMsgVpnClientClearEvent**](docs/AllApi.md#domsgvpnclientclearevent) | **Put** /msgVpns/{msgVpnName}/clients/{clientName}/clearEvent | Clear an event for the Client so it can be generated anew.
*AllApi* | [**DoMsgVpnClientClearStats**](docs/AllApi.md#domsgvpnclientclearstats) | **Put** /msgVpns/{msgVpnName}/clients/{clientName}/clearStats | Clear the statistics for the Client.
*AllApi* | [**DoMsgVpnClientDisconnect**](docs/AllApi.md#domsgvpnclientdisconnect) | **Put** /msgVpns/{msgVpnName}/clients/{clientName}/disconnect | Disconnect the Client.
*AllApi* | [**DoMsgVpnClientTransactedSessionDelete**](docs/AllApi.md#domsgvpnclienttransactedsessiondelete) | **Put** /msgVpns/{msgVpnName}/clients/{clientName}/transactedSessions/{sessionName}/delete | Delete the Transacted Session.
*AllApi* | [**DoMsgVpnDistributedCacheClusterInstanceBackupCachedMsgs**](docs/AllApi.md#domsgvpndistributedcacheclusterinstancebackupcachedmsgs) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/backupCachedMsgs | Backup cached messages of the Cache Instance to disk.
*AllApi* | [**DoMsgVpnDistributedCacheClusterInstanceCancelBackupCachedMsgs**](docs/AllApi.md#domsgvpndistributedcacheclusterinstancecancelbackupcachedmsgs) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/cancelBackupCachedMsgs | Cancel the backup of cached messages from the Cache Instance.
*AllApi* | [**DoMsgVpnDistributedCacheClusterInstanceCancelRestoreCachedMsgs**](docs/AllApi.md#domsgvpndistributedcacheclusterinstancecancelrestorecachedmsgs) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/cancelRestoreCachedMsgs | Cancel the restore of cached messages to the Cache Instance.
*AllApi* | [**DoMsgVpnDistributedCacheClusterInstanceClearEvent**](docs/AllApi.md#domsgvpndistributedcacheclusterinstanceclearevent) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/clearEvent | Clear an event for the Cache Instance so it can be generated anew.
*AllApi* | [**DoMsgVpnDistributedCacheClusterInstanceClearStats**](docs/AllApi.md#domsgvpndistributedcacheclusterinstanceclearstats) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/clearStats | Clear the statistics for the Cache Instance.
*AllApi* | [**DoMsgVpnDistributedCacheClusterInstanceDeleteMsgs**](docs/AllApi.md#domsgvpndistributedcacheclusterinstancedeletemsgs) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/deleteMsgs | Delete messages covered by the given topic in the Cache Instance.
*AllApi* | [**DoMsgVpnDistributedCacheClusterInstanceRestoreCachedMsgs**](docs/AllApi.md#domsgvpndistributedcacheclusterinstancerestorecachedmsgs) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/restoreCachedMsgs | Restore cached messages for the Cache Instance from disk.
*AllApi* | [**DoMsgVpnDistributedCacheClusterInstanceStart**](docs/AllApi.md#domsgvpndistributedcacheclusterinstancestart) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/start | Start the Cache Instance.
*AllApi* | [**DoMsgVpnMqttSessionClearStats**](docs/AllApi.md#domsgvpnmqttsessionclearstats) | **Put** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/clearStats | Clear the statistics for the MQTT Session.
*AllApi* | [**DoMsgVpnQueueCancelReplay**](docs/AllApi.md#domsgvpnqueuecancelreplay) | **Put** /msgVpns/{msgVpnName}/queues/{queueName}/cancelReplay | Cancel the replay of messages to the Queue.
*AllApi* | [**DoMsgVpnQueueClearStats**](docs/AllApi.md#domsgvpnqueueclearstats) | **Put** /msgVpns/{msgVpnName}/queues/{queueName}/clearStats | Clear the statistics for the Queue.
*AllApi* | [**DoMsgVpnQueueCopyMsgFromQueue**](docs/AllApi.md#domsgvpnqueuecopymsgfromqueue) | **Put** /msgVpns/{msgVpnName}/queues/{queueName}/copyMsgFromQueue | Copy a message from another Queue to this Queue.
*AllApi* | [**DoMsgVpnQueueCopyMsgFromReplayLog**](docs/AllApi.md#domsgvpnqueuecopymsgfromreplaylog) | **Put** /msgVpns/{msgVpnName}/queues/{queueName}/copyMsgFromReplayLog | Copy a message from a Replay Log to this Queue.
*AllApi* | [**DoMsgVpnQueueCopyMsgFromTopicEndpoint**](docs/AllApi.md#domsgvpnqueuecopymsgfromtopicendpoint) | **Put** /msgVpns/{msgVpnName}/queues/{queueName}/copyMsgFromTopicEndpoint | Copy a message from a Topic Endpoint to this Queue.
*AllApi* | [**DoMsgVpnQueueDeleteMsgs**](docs/AllApi.md#domsgvpnqueuedeletemsgs) | **Put** /msgVpns/{msgVpnName}/queues/{queueName}/deleteMsgs | Delete all spooled messages from the Queue.
*AllApi* | [**DoMsgVpnQueueMsgDelete**](docs/AllApi.md#domsgvpnqueuemsgdelete) | **Put** /msgVpns/{msgVpnName}/queues/{queueName}/msgs/{msgId}/delete | Delete the Message from the Queue.
*AllApi* | [**DoMsgVpnQueueStartReplay**](docs/AllApi.md#domsgvpnqueuestartreplay) | **Put** /msgVpns/{msgVpnName}/queues/{queueName}/startReplay | Start the replay of messages to the Queue.
*AllApi* | [**DoMsgVpnReplayLogTrimLoggedMsgs**](docs/AllApi.md#domsgvpnreplaylogtrimloggedmsgs) | **Put** /msgVpns/{msgVpnName}/replayLogs/{replayLogName}/trimLoggedMsgs | Trim (delete) messages from the Replay Log.
*AllApi* | [**DoMsgVpnRestDeliveryPointRestConsumerClearStats**](docs/AllApi.md#domsgvpnrestdeliverypointrestconsumerclearstats) | **Put** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/clearStats | Clear the statistics for the REST Consumer.
*AllApi* | [**DoMsgVpnTopicEndpointCancelReplay**](docs/AllApi.md#domsgvpntopicendpointcancelreplay) | **Put** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/cancelReplay | Cancel the replay of messages to the Topic Endpoint.
*AllApi* | [**DoMsgVpnTopicEndpointClearStats**](docs/AllApi.md#domsgvpntopicendpointclearstats) | **Put** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/clearStats | Clear the statistics for the Topic Endpoint.
*AllApi* | [**DoMsgVpnTopicEndpointCopyMsgFromQueue**](docs/AllApi.md#domsgvpntopicendpointcopymsgfromqueue) | **Put** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/copyMsgFromQueue | Copy a message from a Queue to this Topic Endpoint.
*AllApi* | [**DoMsgVpnTopicEndpointCopyMsgFromReplayLog**](docs/AllApi.md#domsgvpntopicendpointcopymsgfromreplaylog) | **Put** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/copyMsgFromReplayLog | Copy a message from a Replay Log to this Topic Endpoint.
*AllApi* | [**DoMsgVpnTopicEndpointCopyMsgFromTopicEndpoint**](docs/AllApi.md#domsgvpntopicendpointcopymsgfromtopicendpoint) | **Put** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/copyMsgFromTopicEndpoint | Copy a message from another Topic Endpoint to this Topic Endpoint.
*AllApi* | [**DoMsgVpnTopicEndpointDeleteMsgs**](docs/AllApi.md#domsgvpntopicendpointdeletemsgs) | **Put** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/deleteMsgs | Delete all spooled messages from the Topic Endpoint.
*AllApi* | [**DoMsgVpnTopicEndpointMsgDelete**](docs/AllApi.md#domsgvpntopicendpointmsgdelete) | **Put** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/msgs/{msgId}/delete | Delete the Message from the Topic Endpoint.
*AllApi* | [**DoMsgVpnTopicEndpointStartReplay**](docs/AllApi.md#domsgvpntopicendpointstartreplay) | **Put** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/startReplay | Start the replay of messages to the Topic Endpoint.
*AllApi* | [**DoMsgVpnTransactionCommit**](docs/AllApi.md#domsgvpntransactioncommit) | **Put** /msgVpns/{msgVpnName}/transactions/{xid}/commit | Commit the Transaction.
*AllApi* | [**DoMsgVpnTransactionDelete**](docs/AllApi.md#domsgvpntransactiondelete) | **Put** /msgVpns/{msgVpnName}/transactions/{xid}/delete | Delete the Transaction.
*AllApi* | [**DoMsgVpnTransactionRollback**](docs/AllApi.md#domsgvpntransactionrollback) | **Put** /msgVpns/{msgVpnName}/transactions/{xid}/rollback | Rollback the Transaction.
*AllApi* | [**DoOauthProfileClearStats**](docs/AllApi.md#dooauthprofileclearstats) | **Put** /oauthProfiles/{oauthProfileName}/clearStats | Clear the statistics for the OAuth Profile.
*AllApi* | [**DoSessionDelete**](docs/AllApi.md#dosessiondelete) | **Put** /sessions/{sessionUsername},{sessionId}/delete | Delete the SEMP session.
*AllApi* | [**GetAbout**](docs/AllApi.md#getabout) | **Get** /about | Get an About object.
*AllApi* | [**GetAboutApi**](docs/AllApi.md#getaboutapi) | **Get** /about/api | Get an API Description object.
*AllApi* | [**GetAboutUser**](docs/AllApi.md#getaboutuser) | **Get** /about/user | Get a User object.
*AllApi* | [**GetAboutUserMsgVpn**](docs/AllApi.md#getaboutusermsgvpn) | **Get** /about/user/msgVpns/{msgVpnName} | Get a User Message VPN object.
*AllApi* | [**GetAboutUserMsgVpns**](docs/AllApi.md#getaboutusermsgvpns) | **Get** /about/user/msgVpns | Get a list of User Message VPN objects.
*AllApi* | [**GetBroker**](docs/AllApi.md#getbroker) | **Get** / | Get a Broker object.
*AllApi* | [**GetCertAuthorities**](docs/AllApi.md#getcertauthorities) | **Get** /certAuthorities | Get a list of Certificate Authority objects.
*AllApi* | [**GetCertAuthority**](docs/AllApi.md#getcertauthority) | **Get** /certAuthorities/{certAuthorityName} | Get a Certificate Authority object.
*AllApi* | [**GetClientCertAuthorities**](docs/AllApi.md#getclientcertauthorities) | **Get** /clientCertAuthorities | Get a list of Client Certificate Authority objects.
*AllApi* | [**GetClientCertAuthority**](docs/AllApi.md#getclientcertauthority) | **Get** /clientCertAuthorities/{certAuthorityName} | Get a Client Certificate Authority object.
*AllApi* | [**GetMsgVpn**](docs/AllApi.md#getmsgvpn) | **Get** /msgVpns/{msgVpnName} | Get a Message VPN object.
*AllApi* | [**GetMsgVpnAuthenticationOauthProfile**](docs/AllApi.md#getmsgvpnauthenticationoauthprofile) | **Get** /msgVpns/{msgVpnName}/authenticationOauthProfiles/{oauthProfileName} | Get an OAuth Profile object.
*AllApi* | [**GetMsgVpnAuthenticationOauthProfiles**](docs/AllApi.md#getmsgvpnauthenticationoauthprofiles) | **Get** /msgVpns/{msgVpnName}/authenticationOauthProfiles | Get a list of OAuth Profile objects.
*AllApi* | [**GetMsgVpnAuthenticationOauthProvider**](docs/AllApi.md#getmsgvpnauthenticationoauthprovider) | **Get** /msgVpns/{msgVpnName}/authenticationOauthProviders/{oauthProviderName} | Get an OAuth Provider object.
*AllApi* | [**GetMsgVpnAuthenticationOauthProviders**](docs/AllApi.md#getmsgvpnauthenticationoauthproviders) | **Get** /msgVpns/{msgVpnName}/authenticationOauthProviders | Get a list of OAuth Provider objects.
*AllApi* | [**GetMsgVpnBridge**](docs/AllApi.md#getmsgvpnbridge) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter} | Get a Bridge object.
*AllApi* | [**GetMsgVpnBridges**](docs/AllApi.md#getmsgvpnbridges) | **Get** /msgVpns/{msgVpnName}/bridges | Get a list of Bridge objects.
*AllApi* | [**GetMsgVpnClient**](docs/AllApi.md#getmsgvpnclient) | **Get** /msgVpns/{msgVpnName}/clients/{clientName} | Get a Client object.
*AllApi* | [**GetMsgVpnClientTransactedSession**](docs/AllApi.md#getmsgvpnclienttransactedsession) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/transactedSessions/{sessionName} | Get a Client Transacted Session object.
*AllApi* | [**GetMsgVpnClientTransactedSessions**](docs/AllApi.md#getmsgvpnclienttransactedsessions) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/transactedSessions | Get a list of Client Transacted Session objects.
*AllApi* | [**GetMsgVpnClients**](docs/AllApi.md#getmsgvpnclients) | **Get** /msgVpns/{msgVpnName}/clients | Get a list of Client objects.
*AllApi* | [**GetMsgVpnDistributedCache**](docs/AllApi.md#getmsgvpndistributedcache) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName} | Get a Distributed Cache object.
*AllApi* | [**GetMsgVpnDistributedCacheCluster**](docs/AllApi.md#getmsgvpndistributedcachecluster) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName} | Get a Cache Cluster object.
*AllApi* | [**GetMsgVpnDistributedCacheClusterInstance**](docs/AllApi.md#getmsgvpndistributedcacheclusterinstance) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName} | Get a Cache Instance object.
*AllApi* | [**GetMsgVpnDistributedCacheClusterInstances**](docs/AllApi.md#getmsgvpndistributedcacheclusterinstances) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances | Get a list of Cache Instance objects.
*AllApi* | [**GetMsgVpnDistributedCacheClusters**](docs/AllApi.md#getmsgvpndistributedcacheclusters) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters | Get a list of Cache Cluster objects.
*AllApi* | [**GetMsgVpnDistributedCaches**](docs/AllApi.md#getmsgvpndistributedcaches) | **Get** /msgVpns/{msgVpnName}/distributedCaches | Get a list of Distributed Cache objects.
*AllApi* | [**GetMsgVpnMqttSession**](docs/AllApi.md#getmsgvpnmqttsession) | **Get** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter} | Get an MQTT Session object.
*AllApi* | [**GetMsgVpnMqttSessions**](docs/AllApi.md#getmsgvpnmqttsessions) | **Get** /msgVpns/{msgVpnName}/mqttSessions | Get a list of MQTT Session objects.
*AllApi* | [**GetMsgVpnQueue**](docs/AllApi.md#getmsgvpnqueue) | **Get** /msgVpns/{msgVpnName}/queues/{queueName} | Get a Queue object.
*AllApi* | [**GetMsgVpnQueueMsg**](docs/AllApi.md#getmsgvpnqueuemsg) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/msgs/{msgId} | Get a Queue Message object.
*AllApi* | [**GetMsgVpnQueueMsgs**](docs/AllApi.md#getmsgvpnqueuemsgs) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/msgs | Get a list of Queue Message objects.
*AllApi* | [**GetMsgVpnQueues**](docs/AllApi.md#getmsgvpnqueues) | **Get** /msgVpns/{msgVpnName}/queues | Get a list of Queue objects.
*AllApi* | [**GetMsgVpnReplayLog**](docs/AllApi.md#getmsgvpnreplaylog) | **Get** /msgVpns/{msgVpnName}/replayLogs/{replayLogName} | Get a Replay Log object.
*AllApi* | [**GetMsgVpnReplayLogs**](docs/AllApi.md#getmsgvpnreplaylogs) | **Get** /msgVpns/{msgVpnName}/replayLogs | Get a list of Replay Log objects.
*AllApi* | [**GetMsgVpnRestDeliveryPoint**](docs/AllApi.md#getmsgvpnrestdeliverypoint) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName} | Get a REST Delivery Point object.
*AllApi* | [**GetMsgVpnRestDeliveryPointRestConsumer**](docs/AllApi.md#getmsgvpnrestdeliverypointrestconsumer) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName} | Get a REST Consumer object.
*AllApi* | [**GetMsgVpnRestDeliveryPointRestConsumers**](docs/AllApi.md#getmsgvpnrestdeliverypointrestconsumers) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers | Get a list of REST Consumer objects.
*AllApi* | [**GetMsgVpnRestDeliveryPoints**](docs/AllApi.md#getmsgvpnrestdeliverypoints) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints | Get a list of REST Delivery Point objects.
*AllApi* | [**GetMsgVpnTopicEndpoint**](docs/AllApi.md#getmsgvpntopicendpoint) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName} | Get a Topic Endpoint object.
*AllApi* | [**GetMsgVpnTopicEndpointMsg**](docs/AllApi.md#getmsgvpntopicendpointmsg) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/msgs/{msgId} | Get a Topic Endpoint Message object.
*AllApi* | [**GetMsgVpnTopicEndpointMsgs**](docs/AllApi.md#getmsgvpntopicendpointmsgs) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/msgs | Get a list of Topic Endpoint Message objects.
*AllApi* | [**GetMsgVpnTopicEndpoints**](docs/AllApi.md#getmsgvpntopicendpoints) | **Get** /msgVpns/{msgVpnName}/topicEndpoints | Get a list of Topic Endpoint objects.
*AllApi* | [**GetMsgVpnTransaction**](docs/AllApi.md#getmsgvpntransaction) | **Get** /msgVpns/{msgVpnName}/transactions/{xid} | Get a Replicated Local Transaction or XA Transaction object.
*AllApi* | [**GetMsgVpnTransactions**](docs/AllApi.md#getmsgvpntransactions) | **Get** /msgVpns/{msgVpnName}/transactions | Get a list of Replicated Local Transaction or XA Transaction objects.
*AllApi* | [**GetMsgVpns**](docs/AllApi.md#getmsgvpns) | **Get** /msgVpns | Get a list of Message VPN objects.
*AllApi* | [**GetOauthProfile**](docs/AllApi.md#getoauthprofile) | **Get** /oauthProfiles/{oauthProfileName} | Get an OAuth Profile object.
*AllApi* | [**GetOauthProfiles**](docs/AllApi.md#getoauthprofiles) | **Get** /oauthProfiles | Get a list of OAuth Profile objects.
*AllApi* | [**GetSession**](docs/AllApi.md#getsession) | **Get** /sessions/{sessionUsername},{sessionId} | Get a SEMP Session object.
*AllApi* | [**GetSessions**](docs/AllApi.md#getsessions) | **Get** /sessions | Get a list of SEMP Session objects.
*AuthenticationOauthProfileApi* | [**DoMsgVpnAuthenticationOauthProfileClearStats**](docs/AuthenticationOauthProfileApi.md#domsgvpnauthenticationoauthprofileclearstats) | **Put** /msgVpns/{msgVpnName}/authenticationOauthProfiles/{oauthProfileName}/clearStats | Clear the statistics for the OAuth Profile.
*AuthenticationOauthProfileApi* | [**GetMsgVpnAuthenticationOauthProfile**](docs/AuthenticationOauthProfileApi.md#getmsgvpnauthenticationoauthprofile) | **Get** /msgVpns/{msgVpnName}/authenticationOauthProfiles/{oauthProfileName} | Get an OAuth Profile object.
*AuthenticationOauthProfileApi* | [**GetMsgVpnAuthenticationOauthProfiles**](docs/AuthenticationOauthProfileApi.md#getmsgvpnauthenticationoauthprofiles) | **Get** /msgVpns/{msgVpnName}/authenticationOauthProfiles | Get a list of OAuth Profile objects.
*AuthenticationOauthProviderApi* | [**DoMsgVpnAuthenticationOauthProviderClearStats**](docs/AuthenticationOauthProviderApi.md#domsgvpnauthenticationoauthproviderclearstats) | **Put** /msgVpns/{msgVpnName}/authenticationOauthProviders/{oauthProviderName}/clearStats | Clear the statistics for the OAuth Provider.
*AuthenticationOauthProviderApi* | [**GetMsgVpnAuthenticationOauthProvider**](docs/AuthenticationOauthProviderApi.md#getmsgvpnauthenticationoauthprovider) | **Get** /msgVpns/{msgVpnName}/authenticationOauthProviders/{oauthProviderName} | Get an OAuth Provider object.
*AuthenticationOauthProviderApi* | [**GetMsgVpnAuthenticationOauthProviders**](docs/AuthenticationOauthProviderApi.md#getmsgvpnauthenticationoauthproviders) | **Get** /msgVpns/{msgVpnName}/authenticationOauthProviders | Get a list of OAuth Provider objects.
*BridgeApi* | [**DoMsgVpnBridgeClearEvent**](docs/BridgeApi.md#domsgvpnbridgeclearevent) | **Put** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/clearEvent | Clear an event for the Bridge so it can be generated anew.
*BridgeApi* | [**DoMsgVpnBridgeClearStats**](docs/BridgeApi.md#domsgvpnbridgeclearstats) | **Put** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/clearStats | Clear the statistics for the Bridge.
*BridgeApi* | [**DoMsgVpnBridgeDisconnect**](docs/BridgeApi.md#domsgvpnbridgedisconnect) | **Put** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/disconnect | Disconnect the Bridge.
*BridgeApi* | [**GetMsgVpnBridge**](docs/BridgeApi.md#getmsgvpnbridge) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter} | Get a Bridge object.
*BridgeApi* | [**GetMsgVpnBridges**](docs/BridgeApi.md#getmsgvpnbridges) | **Get** /msgVpns/{msgVpnName}/bridges | Get a list of Bridge objects.
*CertAuthorityApi* | [**DoCertAuthorityRefreshCrl**](docs/CertAuthorityApi.md#docertauthorityrefreshcrl) | **Put** /certAuthorities/{certAuthorityName}/refreshCrl | Refresh the CRL file for the Certificate Authority.
*CertAuthorityApi* | [**GetCertAuthorities**](docs/CertAuthorityApi.md#getcertauthorities) | **Get** /certAuthorities | Get a list of Certificate Authority objects.
*CertAuthorityApi* | [**GetCertAuthority**](docs/CertAuthorityApi.md#getcertauthority) | **Get** /certAuthorities/{certAuthorityName} | Get a Certificate Authority object.
*ClientApi* | [**DoMsgVpnClientClearEvent**](docs/ClientApi.md#domsgvpnclientclearevent) | **Put** /msgVpns/{msgVpnName}/clients/{clientName}/clearEvent | Clear an event for the Client so it can be generated anew.
*ClientApi* | [**DoMsgVpnClientClearStats**](docs/ClientApi.md#domsgvpnclientclearstats) | **Put** /msgVpns/{msgVpnName}/clients/{clientName}/clearStats | Clear the statistics for the Client.
*ClientApi* | [**DoMsgVpnClientDisconnect**](docs/ClientApi.md#domsgvpnclientdisconnect) | **Put** /msgVpns/{msgVpnName}/clients/{clientName}/disconnect | Disconnect the Client.
*ClientApi* | [**DoMsgVpnClientTransactedSessionDelete**](docs/ClientApi.md#domsgvpnclienttransactedsessiondelete) | **Put** /msgVpns/{msgVpnName}/clients/{clientName}/transactedSessions/{sessionName}/delete | Delete the Transacted Session.
*ClientApi* | [**GetMsgVpnClient**](docs/ClientApi.md#getmsgvpnclient) | **Get** /msgVpns/{msgVpnName}/clients/{clientName} | Get a Client object.
*ClientApi* | [**GetMsgVpnClientTransactedSession**](docs/ClientApi.md#getmsgvpnclienttransactedsession) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/transactedSessions/{sessionName} | Get a Client Transacted Session object.
*ClientApi* | [**GetMsgVpnClientTransactedSessions**](docs/ClientApi.md#getmsgvpnclienttransactedsessions) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/transactedSessions | Get a list of Client Transacted Session objects.
*ClientApi* | [**GetMsgVpnClients**](docs/ClientApi.md#getmsgvpnclients) | **Get** /msgVpns/{msgVpnName}/clients | Get a list of Client objects.
*ClientCertAuthorityApi* | [**DoClientCertAuthorityRefreshCrl**](docs/ClientCertAuthorityApi.md#doclientcertauthorityrefreshcrl) | **Put** /clientCertAuthorities/{certAuthorityName}/refreshCrl | Refresh the CRL file for the Client Certificate Authority.
*ClientCertAuthorityApi* | [**GetClientCertAuthorities**](docs/ClientCertAuthorityApi.md#getclientcertauthorities) | **Get** /clientCertAuthorities | Get a list of Client Certificate Authority objects.
*ClientCertAuthorityApi* | [**GetClientCertAuthority**](docs/ClientCertAuthorityApi.md#getclientcertauthority) | **Get** /clientCertAuthorities/{certAuthorityName} | Get a Client Certificate Authority object.
*DistributedCacheApi* | [**DoMsgVpnDistributedCacheClusterInstanceBackupCachedMsgs**](docs/DistributedCacheApi.md#domsgvpndistributedcacheclusterinstancebackupcachedmsgs) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/backupCachedMsgs | Backup cached messages of the Cache Instance to disk.
*DistributedCacheApi* | [**DoMsgVpnDistributedCacheClusterInstanceCancelBackupCachedMsgs**](docs/DistributedCacheApi.md#domsgvpndistributedcacheclusterinstancecancelbackupcachedmsgs) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/cancelBackupCachedMsgs | Cancel the backup of cached messages from the Cache Instance.
*DistributedCacheApi* | [**DoMsgVpnDistributedCacheClusterInstanceCancelRestoreCachedMsgs**](docs/DistributedCacheApi.md#domsgvpndistributedcacheclusterinstancecancelrestorecachedmsgs) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/cancelRestoreCachedMsgs | Cancel the restore of cached messages to the Cache Instance.
*DistributedCacheApi* | [**DoMsgVpnDistributedCacheClusterInstanceClearEvent**](docs/DistributedCacheApi.md#domsgvpndistributedcacheclusterinstanceclearevent) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/clearEvent | Clear an event for the Cache Instance so it can be generated anew.
*DistributedCacheApi* | [**DoMsgVpnDistributedCacheClusterInstanceClearStats**](docs/DistributedCacheApi.md#domsgvpndistributedcacheclusterinstanceclearstats) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/clearStats | Clear the statistics for the Cache Instance.
*DistributedCacheApi* | [**DoMsgVpnDistributedCacheClusterInstanceDeleteMsgs**](docs/DistributedCacheApi.md#domsgvpndistributedcacheclusterinstancedeletemsgs) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/deleteMsgs | Delete messages covered by the given topic in the Cache Instance.
*DistributedCacheApi* | [**DoMsgVpnDistributedCacheClusterInstanceRestoreCachedMsgs**](docs/DistributedCacheApi.md#domsgvpndistributedcacheclusterinstancerestorecachedmsgs) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/restoreCachedMsgs | Restore cached messages for the Cache Instance from disk.
*DistributedCacheApi* | [**DoMsgVpnDistributedCacheClusterInstanceStart**](docs/DistributedCacheApi.md#domsgvpndistributedcacheclusterinstancestart) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/start | Start the Cache Instance.
*DistributedCacheApi* | [**GetMsgVpnDistributedCache**](docs/DistributedCacheApi.md#getmsgvpndistributedcache) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName} | Get a Distributed Cache object.
*DistributedCacheApi* | [**GetMsgVpnDistributedCacheCluster**](docs/DistributedCacheApi.md#getmsgvpndistributedcachecluster) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName} | Get a Cache Cluster object.
*DistributedCacheApi* | [**GetMsgVpnDistributedCacheClusterInstance**](docs/DistributedCacheApi.md#getmsgvpndistributedcacheclusterinstance) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName} | Get a Cache Instance object.
*DistributedCacheApi* | [**GetMsgVpnDistributedCacheClusterInstances**](docs/DistributedCacheApi.md#getmsgvpndistributedcacheclusterinstances) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances | Get a list of Cache Instance objects.
*DistributedCacheApi* | [**GetMsgVpnDistributedCacheClusters**](docs/DistributedCacheApi.md#getmsgvpndistributedcacheclusters) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters | Get a list of Cache Cluster objects.
*DistributedCacheApi* | [**GetMsgVpnDistributedCaches**](docs/DistributedCacheApi.md#getmsgvpndistributedcaches) | **Get** /msgVpns/{msgVpnName}/distributedCaches | Get a list of Distributed Cache objects.
*MqttSessionApi* | [**DoMsgVpnMqttSessionClearStats**](docs/MqttSessionApi.md#domsgvpnmqttsessionclearstats) | **Put** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/clearStats | Clear the statistics for the MQTT Session.
*MqttSessionApi* | [**GetMsgVpnMqttSession**](docs/MqttSessionApi.md#getmsgvpnmqttsession) | **Get** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter} | Get an MQTT Session object.
*MqttSessionApi* | [**GetMsgVpnMqttSessions**](docs/MqttSessionApi.md#getmsgvpnmqttsessions) | **Get** /msgVpns/{msgVpnName}/mqttSessions | Get a list of MQTT Session objects.
*MsgVpnApi* | [**DoMsgVpnAuthenticationOauthProfileClearStats**](docs/MsgVpnApi.md#domsgvpnauthenticationoauthprofileclearstats) | **Put** /msgVpns/{msgVpnName}/authenticationOauthProfiles/{oauthProfileName}/clearStats | Clear the statistics for the OAuth Profile.
*MsgVpnApi* | [**DoMsgVpnAuthenticationOauthProviderClearStats**](docs/MsgVpnApi.md#domsgvpnauthenticationoauthproviderclearstats) | **Put** /msgVpns/{msgVpnName}/authenticationOauthProviders/{oauthProviderName}/clearStats | Clear the statistics for the OAuth Provider.
*MsgVpnApi* | [**DoMsgVpnBridgeClearEvent**](docs/MsgVpnApi.md#domsgvpnbridgeclearevent) | **Put** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/clearEvent | Clear an event for the Bridge so it can be generated anew.
*MsgVpnApi* | [**DoMsgVpnBridgeClearStats**](docs/MsgVpnApi.md#domsgvpnbridgeclearstats) | **Put** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/clearStats | Clear the statistics for the Bridge.
*MsgVpnApi* | [**DoMsgVpnBridgeDisconnect**](docs/MsgVpnApi.md#domsgvpnbridgedisconnect) | **Put** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/disconnect | Disconnect the Bridge.
*MsgVpnApi* | [**DoMsgVpnClearMsgSpoolStats**](docs/MsgVpnApi.md#domsgvpnclearmsgspoolstats) | **Put** /msgVpns/{msgVpnName}/clearMsgSpoolStats | Clear the message spool statistics for the Message VPN.
*MsgVpnApi* | [**DoMsgVpnClearReplicationStats**](docs/MsgVpnApi.md#domsgvpnclearreplicationstats) | **Put** /msgVpns/{msgVpnName}/clearReplicationStats | Clear the replication statistics for the Message VPN.
*MsgVpnApi* | [**DoMsgVpnClearServiceStats**](docs/MsgVpnApi.md#domsgvpnclearservicestats) | **Put** /msgVpns/{msgVpnName}/clearServiceStats | Clear the service statistics for the Message VPN.
*MsgVpnApi* | [**DoMsgVpnClearStats**](docs/MsgVpnApi.md#domsgvpnclearstats) | **Put** /msgVpns/{msgVpnName}/clearStats | Clear the client statistics for the Message VPN.
*MsgVpnApi* | [**DoMsgVpnClientClearEvent**](docs/MsgVpnApi.md#domsgvpnclientclearevent) | **Put** /msgVpns/{msgVpnName}/clients/{clientName}/clearEvent | Clear an event for the Client so it can be generated anew.
*MsgVpnApi* | [**DoMsgVpnClientClearStats**](docs/MsgVpnApi.md#domsgvpnclientclearstats) | **Put** /msgVpns/{msgVpnName}/clients/{clientName}/clearStats | Clear the statistics for the Client.
*MsgVpnApi* | [**DoMsgVpnClientDisconnect**](docs/MsgVpnApi.md#domsgvpnclientdisconnect) | **Put** /msgVpns/{msgVpnName}/clients/{clientName}/disconnect | Disconnect the Client.
*MsgVpnApi* | [**DoMsgVpnClientTransactedSessionDelete**](docs/MsgVpnApi.md#domsgvpnclienttransactedsessiondelete) | **Put** /msgVpns/{msgVpnName}/clients/{clientName}/transactedSessions/{sessionName}/delete | Delete the Transacted Session.
*MsgVpnApi* | [**DoMsgVpnDistributedCacheClusterInstanceBackupCachedMsgs**](docs/MsgVpnApi.md#domsgvpndistributedcacheclusterinstancebackupcachedmsgs) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/backupCachedMsgs | Backup cached messages of the Cache Instance to disk.
*MsgVpnApi* | [**DoMsgVpnDistributedCacheClusterInstanceCancelBackupCachedMsgs**](docs/MsgVpnApi.md#domsgvpndistributedcacheclusterinstancecancelbackupcachedmsgs) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/cancelBackupCachedMsgs | Cancel the backup of cached messages from the Cache Instance.
*MsgVpnApi* | [**DoMsgVpnDistributedCacheClusterInstanceCancelRestoreCachedMsgs**](docs/MsgVpnApi.md#domsgvpndistributedcacheclusterinstancecancelrestorecachedmsgs) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/cancelRestoreCachedMsgs | Cancel the restore of cached messages to the Cache Instance.
*MsgVpnApi* | [**DoMsgVpnDistributedCacheClusterInstanceClearEvent**](docs/MsgVpnApi.md#domsgvpndistributedcacheclusterinstanceclearevent) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/clearEvent | Clear an event for the Cache Instance so it can be generated anew.
*MsgVpnApi* | [**DoMsgVpnDistributedCacheClusterInstanceClearStats**](docs/MsgVpnApi.md#domsgvpndistributedcacheclusterinstanceclearstats) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/clearStats | Clear the statistics for the Cache Instance.
*MsgVpnApi* | [**DoMsgVpnDistributedCacheClusterInstanceDeleteMsgs**](docs/MsgVpnApi.md#domsgvpndistributedcacheclusterinstancedeletemsgs) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/deleteMsgs | Delete messages covered by the given topic in the Cache Instance.
*MsgVpnApi* | [**DoMsgVpnDistributedCacheClusterInstanceRestoreCachedMsgs**](docs/MsgVpnApi.md#domsgvpndistributedcacheclusterinstancerestorecachedmsgs) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/restoreCachedMsgs | Restore cached messages for the Cache Instance from disk.
*MsgVpnApi* | [**DoMsgVpnDistributedCacheClusterInstanceStart**](docs/MsgVpnApi.md#domsgvpndistributedcacheclusterinstancestart) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/start | Start the Cache Instance.
*MsgVpnApi* | [**DoMsgVpnMqttSessionClearStats**](docs/MsgVpnApi.md#domsgvpnmqttsessionclearstats) | **Put** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/clearStats | Clear the statistics for the MQTT Session.
*MsgVpnApi* | [**DoMsgVpnQueueCancelReplay**](docs/MsgVpnApi.md#domsgvpnqueuecancelreplay) | **Put** /msgVpns/{msgVpnName}/queues/{queueName}/cancelReplay | Cancel the replay of messages to the Queue.
*MsgVpnApi* | [**DoMsgVpnQueueClearStats**](docs/MsgVpnApi.md#domsgvpnqueueclearstats) | **Put** /msgVpns/{msgVpnName}/queues/{queueName}/clearStats | Clear the statistics for the Queue.
*MsgVpnApi* | [**DoMsgVpnQueueCopyMsgFromQueue**](docs/MsgVpnApi.md#domsgvpnqueuecopymsgfromqueue) | **Put** /msgVpns/{msgVpnName}/queues/{queueName}/copyMsgFromQueue | Copy a message from another Queue to this Queue.
*MsgVpnApi* | [**DoMsgVpnQueueCopyMsgFromReplayLog**](docs/MsgVpnApi.md#domsgvpnqueuecopymsgfromreplaylog) | **Put** /msgVpns/{msgVpnName}/queues/{queueName}/copyMsgFromReplayLog | Copy a message from a Replay Log to this Queue.
*MsgVpnApi* | [**DoMsgVpnQueueCopyMsgFromTopicEndpoint**](docs/MsgVpnApi.md#domsgvpnqueuecopymsgfromtopicendpoint) | **Put** /msgVpns/{msgVpnName}/queues/{queueName}/copyMsgFromTopicEndpoint | Copy a message from a Topic Endpoint to this Queue.
*MsgVpnApi* | [**DoMsgVpnQueueDeleteMsgs**](docs/MsgVpnApi.md#domsgvpnqueuedeletemsgs) | **Put** /msgVpns/{msgVpnName}/queues/{queueName}/deleteMsgs | Delete all spooled messages from the Queue.
*MsgVpnApi* | [**DoMsgVpnQueueMsgDelete**](docs/MsgVpnApi.md#domsgvpnqueuemsgdelete) | **Put** /msgVpns/{msgVpnName}/queues/{queueName}/msgs/{msgId}/delete | Delete the Message from the Queue.
*MsgVpnApi* | [**DoMsgVpnQueueStartReplay**](docs/MsgVpnApi.md#domsgvpnqueuestartreplay) | **Put** /msgVpns/{msgVpnName}/queues/{queueName}/startReplay | Start the replay of messages to the Queue.
*MsgVpnApi* | [**DoMsgVpnReplayLogTrimLoggedMsgs**](docs/MsgVpnApi.md#domsgvpnreplaylogtrimloggedmsgs) | **Put** /msgVpns/{msgVpnName}/replayLogs/{replayLogName}/trimLoggedMsgs | Trim (delete) messages from the Replay Log.
*MsgVpnApi* | [**DoMsgVpnRestDeliveryPointRestConsumerClearStats**](docs/MsgVpnApi.md#domsgvpnrestdeliverypointrestconsumerclearstats) | **Put** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/clearStats | Clear the statistics for the REST Consumer.
*MsgVpnApi* | [**DoMsgVpnTopicEndpointCancelReplay**](docs/MsgVpnApi.md#domsgvpntopicendpointcancelreplay) | **Put** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/cancelReplay | Cancel the replay of messages to the Topic Endpoint.
*MsgVpnApi* | [**DoMsgVpnTopicEndpointClearStats**](docs/MsgVpnApi.md#domsgvpntopicendpointclearstats) | **Put** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/clearStats | Clear the statistics for the Topic Endpoint.
*MsgVpnApi* | [**DoMsgVpnTopicEndpointCopyMsgFromQueue**](docs/MsgVpnApi.md#domsgvpntopicendpointcopymsgfromqueue) | **Put** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/copyMsgFromQueue | Copy a message from a Queue to this Topic Endpoint.
*MsgVpnApi* | [**DoMsgVpnTopicEndpointCopyMsgFromReplayLog**](docs/MsgVpnApi.md#domsgvpntopicendpointcopymsgfromreplaylog) | **Put** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/copyMsgFromReplayLog | Copy a message from a Replay Log to this Topic Endpoint.
*MsgVpnApi* | [**DoMsgVpnTopicEndpointCopyMsgFromTopicEndpoint**](docs/MsgVpnApi.md#domsgvpntopicendpointcopymsgfromtopicendpoint) | **Put** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/copyMsgFromTopicEndpoint | Copy a message from another Topic Endpoint to this Topic Endpoint.
*MsgVpnApi* | [**DoMsgVpnTopicEndpointDeleteMsgs**](docs/MsgVpnApi.md#domsgvpntopicendpointdeletemsgs) | **Put** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/deleteMsgs | Delete all spooled messages from the Topic Endpoint.
*MsgVpnApi* | [**DoMsgVpnTopicEndpointMsgDelete**](docs/MsgVpnApi.md#domsgvpntopicendpointmsgdelete) | **Put** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/msgs/{msgId}/delete | Delete the Message from the Topic Endpoint.
*MsgVpnApi* | [**DoMsgVpnTopicEndpointStartReplay**](docs/MsgVpnApi.md#domsgvpntopicendpointstartreplay) | **Put** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/startReplay | Start the replay of messages to the Topic Endpoint.
*MsgVpnApi* | [**DoMsgVpnTransactionCommit**](docs/MsgVpnApi.md#domsgvpntransactioncommit) | **Put** /msgVpns/{msgVpnName}/transactions/{xid}/commit | Commit the Transaction.
*MsgVpnApi* | [**DoMsgVpnTransactionDelete**](docs/MsgVpnApi.md#domsgvpntransactiondelete) | **Put** /msgVpns/{msgVpnName}/transactions/{xid}/delete | Delete the Transaction.
*MsgVpnApi* | [**DoMsgVpnTransactionRollback**](docs/MsgVpnApi.md#domsgvpntransactionrollback) | **Put** /msgVpns/{msgVpnName}/transactions/{xid}/rollback | Rollback the Transaction.
*MsgVpnApi* | [**GetMsgVpn**](docs/MsgVpnApi.md#getmsgvpn) | **Get** /msgVpns/{msgVpnName} | Get a Message VPN object.
*MsgVpnApi* | [**GetMsgVpnAuthenticationOauthProfile**](docs/MsgVpnApi.md#getmsgvpnauthenticationoauthprofile) | **Get** /msgVpns/{msgVpnName}/authenticationOauthProfiles/{oauthProfileName} | Get an OAuth Profile object.
*MsgVpnApi* | [**GetMsgVpnAuthenticationOauthProfiles**](docs/MsgVpnApi.md#getmsgvpnauthenticationoauthprofiles) | **Get** /msgVpns/{msgVpnName}/authenticationOauthProfiles | Get a list of OAuth Profile objects.
*MsgVpnApi* | [**GetMsgVpnAuthenticationOauthProvider**](docs/MsgVpnApi.md#getmsgvpnauthenticationoauthprovider) | **Get** /msgVpns/{msgVpnName}/authenticationOauthProviders/{oauthProviderName} | Get an OAuth Provider object.
*MsgVpnApi* | [**GetMsgVpnAuthenticationOauthProviders**](docs/MsgVpnApi.md#getmsgvpnauthenticationoauthproviders) | **Get** /msgVpns/{msgVpnName}/authenticationOauthProviders | Get a list of OAuth Provider objects.
*MsgVpnApi* | [**GetMsgVpnBridge**](docs/MsgVpnApi.md#getmsgvpnbridge) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter} | Get a Bridge object.
*MsgVpnApi* | [**GetMsgVpnBridges**](docs/MsgVpnApi.md#getmsgvpnbridges) | **Get** /msgVpns/{msgVpnName}/bridges | Get a list of Bridge objects.
*MsgVpnApi* | [**GetMsgVpnClient**](docs/MsgVpnApi.md#getmsgvpnclient) | **Get** /msgVpns/{msgVpnName}/clients/{clientName} | Get a Client object.
*MsgVpnApi* | [**GetMsgVpnClientTransactedSession**](docs/MsgVpnApi.md#getmsgvpnclienttransactedsession) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/transactedSessions/{sessionName} | Get a Client Transacted Session object.
*MsgVpnApi* | [**GetMsgVpnClientTransactedSessions**](docs/MsgVpnApi.md#getmsgvpnclienttransactedsessions) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/transactedSessions | Get a list of Client Transacted Session objects.
*MsgVpnApi* | [**GetMsgVpnClients**](docs/MsgVpnApi.md#getmsgvpnclients) | **Get** /msgVpns/{msgVpnName}/clients | Get a list of Client objects.
*MsgVpnApi* | [**GetMsgVpnDistributedCache**](docs/MsgVpnApi.md#getmsgvpndistributedcache) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName} | Get a Distributed Cache object.
*MsgVpnApi* | [**GetMsgVpnDistributedCacheCluster**](docs/MsgVpnApi.md#getmsgvpndistributedcachecluster) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName} | Get a Cache Cluster object.
*MsgVpnApi* | [**GetMsgVpnDistributedCacheClusterInstance**](docs/MsgVpnApi.md#getmsgvpndistributedcacheclusterinstance) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName} | Get a Cache Instance object.
*MsgVpnApi* | [**GetMsgVpnDistributedCacheClusterInstances**](docs/MsgVpnApi.md#getmsgvpndistributedcacheclusterinstances) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances | Get a list of Cache Instance objects.
*MsgVpnApi* | [**GetMsgVpnDistributedCacheClusters**](docs/MsgVpnApi.md#getmsgvpndistributedcacheclusters) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters | Get a list of Cache Cluster objects.
*MsgVpnApi* | [**GetMsgVpnDistributedCaches**](docs/MsgVpnApi.md#getmsgvpndistributedcaches) | **Get** /msgVpns/{msgVpnName}/distributedCaches | Get a list of Distributed Cache objects.
*MsgVpnApi* | [**GetMsgVpnMqttSession**](docs/MsgVpnApi.md#getmsgvpnmqttsession) | **Get** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter} | Get an MQTT Session object.
*MsgVpnApi* | [**GetMsgVpnMqttSessions**](docs/MsgVpnApi.md#getmsgvpnmqttsessions) | **Get** /msgVpns/{msgVpnName}/mqttSessions | Get a list of MQTT Session objects.
*MsgVpnApi* | [**GetMsgVpnQueue**](docs/MsgVpnApi.md#getmsgvpnqueue) | **Get** /msgVpns/{msgVpnName}/queues/{queueName} | Get a Queue object.
*MsgVpnApi* | [**GetMsgVpnQueueMsg**](docs/MsgVpnApi.md#getmsgvpnqueuemsg) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/msgs/{msgId} | Get a Queue Message object.
*MsgVpnApi* | [**GetMsgVpnQueueMsgs**](docs/MsgVpnApi.md#getmsgvpnqueuemsgs) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/msgs | Get a list of Queue Message objects.
*MsgVpnApi* | [**GetMsgVpnQueues**](docs/MsgVpnApi.md#getmsgvpnqueues) | **Get** /msgVpns/{msgVpnName}/queues | Get a list of Queue objects.
*MsgVpnApi* | [**GetMsgVpnReplayLog**](docs/MsgVpnApi.md#getmsgvpnreplaylog) | **Get** /msgVpns/{msgVpnName}/replayLogs/{replayLogName} | Get a Replay Log object.
*MsgVpnApi* | [**GetMsgVpnReplayLogs**](docs/MsgVpnApi.md#getmsgvpnreplaylogs) | **Get** /msgVpns/{msgVpnName}/replayLogs | Get a list of Replay Log objects.
*MsgVpnApi* | [**GetMsgVpnRestDeliveryPoint**](docs/MsgVpnApi.md#getmsgvpnrestdeliverypoint) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName} | Get a REST Delivery Point object.
*MsgVpnApi* | [**GetMsgVpnRestDeliveryPointRestConsumer**](docs/MsgVpnApi.md#getmsgvpnrestdeliverypointrestconsumer) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName} | Get a REST Consumer object.
*MsgVpnApi* | [**GetMsgVpnRestDeliveryPointRestConsumers**](docs/MsgVpnApi.md#getmsgvpnrestdeliverypointrestconsumers) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers | Get a list of REST Consumer objects.
*MsgVpnApi* | [**GetMsgVpnRestDeliveryPoints**](docs/MsgVpnApi.md#getmsgvpnrestdeliverypoints) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints | Get a list of REST Delivery Point objects.
*MsgVpnApi* | [**GetMsgVpnTopicEndpoint**](docs/MsgVpnApi.md#getmsgvpntopicendpoint) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName} | Get a Topic Endpoint object.
*MsgVpnApi* | [**GetMsgVpnTopicEndpointMsg**](docs/MsgVpnApi.md#getmsgvpntopicendpointmsg) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/msgs/{msgId} | Get a Topic Endpoint Message object.
*MsgVpnApi* | [**GetMsgVpnTopicEndpointMsgs**](docs/MsgVpnApi.md#getmsgvpntopicendpointmsgs) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/msgs | Get a list of Topic Endpoint Message objects.
*MsgVpnApi* | [**GetMsgVpnTopicEndpoints**](docs/MsgVpnApi.md#getmsgvpntopicendpoints) | **Get** /msgVpns/{msgVpnName}/topicEndpoints | Get a list of Topic Endpoint objects.
*MsgVpnApi* | [**GetMsgVpnTransaction**](docs/MsgVpnApi.md#getmsgvpntransaction) | **Get** /msgVpns/{msgVpnName}/transactions/{xid} | Get a Replicated Local Transaction or XA Transaction object.
*MsgVpnApi* | [**GetMsgVpnTransactions**](docs/MsgVpnApi.md#getmsgvpntransactions) | **Get** /msgVpns/{msgVpnName}/transactions | Get a list of Replicated Local Transaction or XA Transaction objects.
*MsgVpnApi* | [**GetMsgVpns**](docs/MsgVpnApi.md#getmsgvpns) | **Get** /msgVpns | Get a list of Message VPN objects.
*OauthProfileApi* | [**DoOauthProfileClearStats**](docs/OauthProfileApi.md#dooauthprofileclearstats) | **Put** /oauthProfiles/{oauthProfileName}/clearStats | Clear the statistics for the OAuth Profile.
*OauthProfileApi* | [**GetOauthProfile**](docs/OauthProfileApi.md#getoauthprofile) | **Get** /oauthProfiles/{oauthProfileName} | Get an OAuth Profile object.
*OauthProfileApi* | [**GetOauthProfiles**](docs/OauthProfileApi.md#getoauthprofiles) | **Get** /oauthProfiles | Get a list of OAuth Profile objects.
*QueueApi* | [**DoMsgVpnQueueCancelReplay**](docs/QueueApi.md#domsgvpnqueuecancelreplay) | **Put** /msgVpns/{msgVpnName}/queues/{queueName}/cancelReplay | Cancel the replay of messages to the Queue.
*QueueApi* | [**DoMsgVpnQueueClearStats**](docs/QueueApi.md#domsgvpnqueueclearstats) | **Put** /msgVpns/{msgVpnName}/queues/{queueName}/clearStats | Clear the statistics for the Queue.
*QueueApi* | [**DoMsgVpnQueueCopyMsgFromQueue**](docs/QueueApi.md#domsgvpnqueuecopymsgfromqueue) | **Put** /msgVpns/{msgVpnName}/queues/{queueName}/copyMsgFromQueue | Copy a message from another Queue to this Queue.
*QueueApi* | [**DoMsgVpnQueueCopyMsgFromReplayLog**](docs/QueueApi.md#domsgvpnqueuecopymsgfromreplaylog) | **Put** /msgVpns/{msgVpnName}/queues/{queueName}/copyMsgFromReplayLog | Copy a message from a Replay Log to this Queue.
*QueueApi* | [**DoMsgVpnQueueCopyMsgFromTopicEndpoint**](docs/QueueApi.md#domsgvpnqueuecopymsgfromtopicendpoint) | **Put** /msgVpns/{msgVpnName}/queues/{queueName}/copyMsgFromTopicEndpoint | Copy a message from a Topic Endpoint to this Queue.
*QueueApi* | [**DoMsgVpnQueueDeleteMsgs**](docs/QueueApi.md#domsgvpnqueuedeletemsgs) | **Put** /msgVpns/{msgVpnName}/queues/{queueName}/deleteMsgs | Delete all spooled messages from the Queue.
*QueueApi* | [**DoMsgVpnQueueMsgDelete**](docs/QueueApi.md#domsgvpnqueuemsgdelete) | **Put** /msgVpns/{msgVpnName}/queues/{queueName}/msgs/{msgId}/delete | Delete the Message from the Queue.
*QueueApi* | [**DoMsgVpnQueueStartReplay**](docs/QueueApi.md#domsgvpnqueuestartreplay) | **Put** /msgVpns/{msgVpnName}/queues/{queueName}/startReplay | Start the replay of messages to the Queue.
*QueueApi* | [**GetMsgVpnQueue**](docs/QueueApi.md#getmsgvpnqueue) | **Get** /msgVpns/{msgVpnName}/queues/{queueName} | Get a Queue object.
*QueueApi* | [**GetMsgVpnQueueMsg**](docs/QueueApi.md#getmsgvpnqueuemsg) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/msgs/{msgId} | Get a Queue Message object.
*QueueApi* | [**GetMsgVpnQueueMsgs**](docs/QueueApi.md#getmsgvpnqueuemsgs) | **Get** /msgVpns/{msgVpnName}/queues/{queueName}/msgs | Get a list of Queue Message objects.
*QueueApi* | [**GetMsgVpnQueues**](docs/QueueApi.md#getmsgvpnqueues) | **Get** /msgVpns/{msgVpnName}/queues | Get a list of Queue objects.
*ReplayLogApi* | [**DoMsgVpnReplayLogTrimLoggedMsgs**](docs/ReplayLogApi.md#domsgvpnreplaylogtrimloggedmsgs) | **Put** /msgVpns/{msgVpnName}/replayLogs/{replayLogName}/trimLoggedMsgs | Trim (delete) messages from the Replay Log.
*ReplayLogApi* | [**GetMsgVpnReplayLog**](docs/ReplayLogApi.md#getmsgvpnreplaylog) | **Get** /msgVpns/{msgVpnName}/replayLogs/{replayLogName} | Get a Replay Log object.
*ReplayLogApi* | [**GetMsgVpnReplayLogs**](docs/ReplayLogApi.md#getmsgvpnreplaylogs) | **Get** /msgVpns/{msgVpnName}/replayLogs | Get a list of Replay Log objects.
*RestDeliveryPointApi* | [**DoMsgVpnRestDeliveryPointRestConsumerClearStats**](docs/RestDeliveryPointApi.md#domsgvpnrestdeliverypointrestconsumerclearstats) | **Put** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/clearStats | Clear the statistics for the REST Consumer.
*RestDeliveryPointApi* | [**GetMsgVpnRestDeliveryPoint**](docs/RestDeliveryPointApi.md#getmsgvpnrestdeliverypoint) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName} | Get a REST Delivery Point object.
*RestDeliveryPointApi* | [**GetMsgVpnRestDeliveryPointRestConsumer**](docs/RestDeliveryPointApi.md#getmsgvpnrestdeliverypointrestconsumer) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName} | Get a REST Consumer object.
*RestDeliveryPointApi* | [**GetMsgVpnRestDeliveryPointRestConsumers**](docs/RestDeliveryPointApi.md#getmsgvpnrestdeliverypointrestconsumers) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers | Get a list of REST Consumer objects.
*RestDeliveryPointApi* | [**GetMsgVpnRestDeliveryPoints**](docs/RestDeliveryPointApi.md#getmsgvpnrestdeliverypoints) | **Get** /msgVpns/{msgVpnName}/restDeliveryPoints | Get a list of REST Delivery Point objects.
*SessionApi* | [**DoSessionDelete**](docs/SessionApi.md#dosessiondelete) | **Put** /sessions/{sessionUsername},{sessionId}/delete | Delete the SEMP session.
*SessionApi* | [**GetSession**](docs/SessionApi.md#getsession) | **Get** /sessions/{sessionUsername},{sessionId} | Get a SEMP Session object.
*SessionApi* | [**GetSessions**](docs/SessionApi.md#getsessions) | **Get** /sessions | Get a list of SEMP Session objects.
*TopicEndpointApi* | [**DoMsgVpnTopicEndpointCancelReplay**](docs/TopicEndpointApi.md#domsgvpntopicendpointcancelreplay) | **Put** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/cancelReplay | Cancel the replay of messages to the Topic Endpoint.
*TopicEndpointApi* | [**DoMsgVpnTopicEndpointClearStats**](docs/TopicEndpointApi.md#domsgvpntopicendpointclearstats) | **Put** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/clearStats | Clear the statistics for the Topic Endpoint.
*TopicEndpointApi* | [**DoMsgVpnTopicEndpointCopyMsgFromQueue**](docs/TopicEndpointApi.md#domsgvpntopicendpointcopymsgfromqueue) | **Put** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/copyMsgFromQueue | Copy a message from a Queue to this Topic Endpoint.
*TopicEndpointApi* | [**DoMsgVpnTopicEndpointCopyMsgFromReplayLog**](docs/TopicEndpointApi.md#domsgvpntopicendpointcopymsgfromreplaylog) | **Put** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/copyMsgFromReplayLog | Copy a message from a Replay Log to this Topic Endpoint.
*TopicEndpointApi* | [**DoMsgVpnTopicEndpointCopyMsgFromTopicEndpoint**](docs/TopicEndpointApi.md#domsgvpntopicendpointcopymsgfromtopicendpoint) | **Put** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/copyMsgFromTopicEndpoint | Copy a message from another Topic Endpoint to this Topic Endpoint.
*TopicEndpointApi* | [**DoMsgVpnTopicEndpointDeleteMsgs**](docs/TopicEndpointApi.md#domsgvpntopicendpointdeletemsgs) | **Put** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/deleteMsgs | Delete all spooled messages from the Topic Endpoint.
*TopicEndpointApi* | [**DoMsgVpnTopicEndpointMsgDelete**](docs/TopicEndpointApi.md#domsgvpntopicendpointmsgdelete) | **Put** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/msgs/{msgId}/delete | Delete the Message from the Topic Endpoint.
*TopicEndpointApi* | [**DoMsgVpnTopicEndpointStartReplay**](docs/TopicEndpointApi.md#domsgvpntopicendpointstartreplay) | **Put** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/startReplay | Start the replay of messages to the Topic Endpoint.
*TopicEndpointApi* | [**GetMsgVpnTopicEndpoint**](docs/TopicEndpointApi.md#getmsgvpntopicendpoint) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName} | Get a Topic Endpoint object.
*TopicEndpointApi* | [**GetMsgVpnTopicEndpointMsg**](docs/TopicEndpointApi.md#getmsgvpntopicendpointmsg) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/msgs/{msgId} | Get a Topic Endpoint Message object.
*TopicEndpointApi* | [**GetMsgVpnTopicEndpointMsgs**](docs/TopicEndpointApi.md#getmsgvpntopicendpointmsgs) | **Get** /msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/msgs | Get a list of Topic Endpoint Message objects.
*TopicEndpointApi* | [**GetMsgVpnTopicEndpoints**](docs/TopicEndpointApi.md#getmsgvpntopicendpoints) | **Get** /msgVpns/{msgVpnName}/topicEndpoints | Get a list of Topic Endpoint objects.
*TransactionApi* | [**DoMsgVpnTransactionCommit**](docs/TransactionApi.md#domsgvpntransactioncommit) | **Put** /msgVpns/{msgVpnName}/transactions/{xid}/commit | Commit the Transaction.
*TransactionApi* | [**DoMsgVpnTransactionDelete**](docs/TransactionApi.md#domsgvpntransactiondelete) | **Put** /msgVpns/{msgVpnName}/transactions/{xid}/delete | Delete the Transaction.
*TransactionApi* | [**DoMsgVpnTransactionRollback**](docs/TransactionApi.md#domsgvpntransactionrollback) | **Put** /msgVpns/{msgVpnName}/transactions/{xid}/rollback | Rollback the Transaction.
*TransactionApi* | [**GetMsgVpnTransaction**](docs/TransactionApi.md#getmsgvpntransaction) | **Get** /msgVpns/{msgVpnName}/transactions/{xid} | Get a Replicated Local Transaction or XA Transaction object.
*TransactionApi* | [**GetMsgVpnTransactions**](docs/TransactionApi.md#getmsgvpntransactions) | **Get** /msgVpns/{msgVpnName}/transactions | Get a list of Replicated Local Transaction or XA Transaction objects.

## Documentation For Models

 - [About](docs/About.md)
 - [AboutApi](docs/AboutApi.md)
 - [AboutApiLinks](docs/AboutApiLinks.md)
 - [AboutApiResponse](docs/AboutApiResponse.md)
 - [AboutLinks](docs/AboutLinks.md)
 - [AboutResponse](docs/AboutResponse.md)
 - [AboutUser](docs/AboutUser.md)
 - [AboutUserLinks](docs/AboutUserLinks.md)
 - [AboutUserLogout](docs/AboutUserLogout.md)
 - [AboutUserMsgVpn](docs/AboutUserMsgVpn.md)
 - [AboutUserMsgVpnLinks](docs/AboutUserMsgVpnLinks.md)
 - [AboutUserMsgVpnResponse](docs/AboutUserMsgVpnResponse.md)
 - [AboutUserMsgVpnsResponse](docs/AboutUserMsgVpnsResponse.md)
 - [AboutUserResponse](docs/AboutUserResponse.md)
 - [Broker](docs/Broker.md)
 - [BrokerLinks](docs/BrokerLinks.md)
 - [BrokerResponse](docs/BrokerResponse.md)
 - [CertAuthoritiesResponse](docs/CertAuthoritiesResponse.md)
 - [CertAuthority](docs/CertAuthority.md)
 - [CertAuthorityLinks](docs/CertAuthorityLinks.md)
 - [CertAuthorityRefreshCrl](docs/CertAuthorityRefreshCrl.md)
 - [CertAuthorityResponse](docs/CertAuthorityResponse.md)
 - [ClientCertAuthoritiesResponse](docs/ClientCertAuthoritiesResponse.md)
 - [ClientCertAuthority](docs/ClientCertAuthority.md)
 - [ClientCertAuthorityLinks](docs/ClientCertAuthorityLinks.md)
 - [ClientCertAuthorityRefreshCrl](docs/ClientCertAuthorityRefreshCrl.md)
 - [ClientCertAuthorityResponse](docs/ClientCertAuthorityResponse.md)
 - [ConfigSyncAssertLeaderMsgVpn](docs/ConfigSyncAssertLeaderMsgVpn.md)
 - [ConfigSyncAssertLeaderRouter](docs/ConfigSyncAssertLeaderRouter.md)
 - [ConfigSyncResyncFollowerMsgVpn](docs/ConfigSyncResyncFollowerMsgVpn.md)
 - [ConfigSyncResyncLeaderMsgVpn](docs/ConfigSyncResyncLeaderMsgVpn.md)
 - [ConfigSyncResyncLeaderRouter](docs/ConfigSyncResyncLeaderRouter.md)
 - [GuaranteedMsgingDefragmentMsgSpoolFilesStart](docs/GuaranteedMsgingDefragmentMsgSpoolFilesStart.md)
 - [GuaranteedMsgingDefragmentMsgSpoolFilesStop](docs/GuaranteedMsgingDefragmentMsgSpoolFilesStop.md)
 - [MsgVpn](docs/MsgVpn.md)
 - [MsgVpnAuthenticationOauthProfile](docs/MsgVpnAuthenticationOauthProfile.md)
 - [MsgVpnAuthenticationOauthProfileClearStats](docs/MsgVpnAuthenticationOauthProfileClearStats.md)
 - [MsgVpnAuthenticationOauthProfileLinks](docs/MsgVpnAuthenticationOauthProfileLinks.md)
 - [MsgVpnAuthenticationOauthProfileResponse](docs/MsgVpnAuthenticationOauthProfileResponse.md)
 - [MsgVpnAuthenticationOauthProfilesResponse](docs/MsgVpnAuthenticationOauthProfilesResponse.md)
 - [MsgVpnAuthenticationOauthProvider](docs/MsgVpnAuthenticationOauthProvider.md)
 - [MsgVpnAuthenticationOauthProviderClearStats](docs/MsgVpnAuthenticationOauthProviderClearStats.md)
 - [MsgVpnAuthenticationOauthProviderLinks](docs/MsgVpnAuthenticationOauthProviderLinks.md)
 - [MsgVpnAuthenticationOauthProviderResponse](docs/MsgVpnAuthenticationOauthProviderResponse.md)
 - [MsgVpnAuthenticationOauthProvidersResponse](docs/MsgVpnAuthenticationOauthProvidersResponse.md)
 - [MsgVpnBridge](docs/MsgVpnBridge.md)
 - [MsgVpnBridgeClearEvent](docs/MsgVpnBridgeClearEvent.md)
 - [MsgVpnBridgeClearStats](docs/MsgVpnBridgeClearStats.md)
 - [MsgVpnBridgeDisconnect](docs/MsgVpnBridgeDisconnect.md)
 - [MsgVpnBridgeLinks](docs/MsgVpnBridgeLinks.md)
 - [MsgVpnBridgeResponse](docs/MsgVpnBridgeResponse.md)
 - [MsgVpnBridgesResponse](docs/MsgVpnBridgesResponse.md)
 - [MsgVpnClearMsgSpoolStats](docs/MsgVpnClearMsgSpoolStats.md)
 - [MsgVpnClearReplicationStats](docs/MsgVpnClearReplicationStats.md)
 - [MsgVpnClearServiceStats](docs/MsgVpnClearServiceStats.md)
 - [MsgVpnClearStats](docs/MsgVpnClearStats.md)
 - [MsgVpnClient](docs/MsgVpnClient.md)
 - [MsgVpnClientClearEvent](docs/MsgVpnClientClearEvent.md)
 - [MsgVpnClientClearStats](docs/MsgVpnClientClearStats.md)
 - [MsgVpnClientDisconnect](docs/MsgVpnClientDisconnect.md)
 - [MsgVpnClientLinks](docs/MsgVpnClientLinks.md)
 - [MsgVpnClientResponse](docs/MsgVpnClientResponse.md)
 - [MsgVpnClientTransactedSession](docs/MsgVpnClientTransactedSession.md)
 - [MsgVpnClientTransactedSessionDelete](docs/MsgVpnClientTransactedSessionDelete.md)
 - [MsgVpnClientTransactedSessionLinks](docs/MsgVpnClientTransactedSessionLinks.md)
 - [MsgVpnClientTransactedSessionResponse](docs/MsgVpnClientTransactedSessionResponse.md)
 - [MsgVpnClientTransactedSessionsResponse](docs/MsgVpnClientTransactedSessionsResponse.md)
 - [MsgVpnClientsResponse](docs/MsgVpnClientsResponse.md)
 - [MsgVpnDistributedCache](docs/MsgVpnDistributedCache.md)
 - [MsgVpnDistributedCacheCluster](docs/MsgVpnDistributedCacheCluster.md)
 - [MsgVpnDistributedCacheClusterInstance](docs/MsgVpnDistributedCacheClusterInstance.md)
 - [MsgVpnDistributedCacheClusterInstanceBackupCachedMsgs](docs/MsgVpnDistributedCacheClusterInstanceBackupCachedMsgs.md)
 - [MsgVpnDistributedCacheClusterInstanceCancelBackupCachedMsgs](docs/MsgVpnDistributedCacheClusterInstanceCancelBackupCachedMsgs.md)
 - [MsgVpnDistributedCacheClusterInstanceCancelRestoreCachedMsgs](docs/MsgVpnDistributedCacheClusterInstanceCancelRestoreCachedMsgs.md)
 - [MsgVpnDistributedCacheClusterInstanceClearEvent](docs/MsgVpnDistributedCacheClusterInstanceClearEvent.md)
 - [MsgVpnDistributedCacheClusterInstanceClearStats](docs/MsgVpnDistributedCacheClusterInstanceClearStats.md)
 - [MsgVpnDistributedCacheClusterInstanceDeleteMsgs](docs/MsgVpnDistributedCacheClusterInstanceDeleteMsgs.md)
 - [MsgVpnDistributedCacheClusterInstanceLinks](docs/MsgVpnDistributedCacheClusterInstanceLinks.md)
 - [MsgVpnDistributedCacheClusterInstanceResponse](docs/MsgVpnDistributedCacheClusterInstanceResponse.md)
 - [MsgVpnDistributedCacheClusterInstanceRestoreCachedMsgs](docs/MsgVpnDistributedCacheClusterInstanceRestoreCachedMsgs.md)
 - [MsgVpnDistributedCacheClusterInstanceStart](docs/MsgVpnDistributedCacheClusterInstanceStart.md)
 - [MsgVpnDistributedCacheClusterInstancesResponse](docs/MsgVpnDistributedCacheClusterInstancesResponse.md)
 - [MsgVpnDistributedCacheClusterLinks](docs/MsgVpnDistributedCacheClusterLinks.md)
 - [MsgVpnDistributedCacheClusterResponse](docs/MsgVpnDistributedCacheClusterResponse.md)
 - [MsgVpnDistributedCacheClustersResponse](docs/MsgVpnDistributedCacheClustersResponse.md)
 - [MsgVpnDistributedCacheLinks](docs/MsgVpnDistributedCacheLinks.md)
 - [MsgVpnDistributedCacheResponse](docs/MsgVpnDistributedCacheResponse.md)
 - [MsgVpnDistributedCachesResponse](docs/MsgVpnDistributedCachesResponse.md)
 - [MsgVpnLinks](docs/MsgVpnLinks.md)
 - [MsgVpnMqttSession](docs/MsgVpnMqttSession.md)
 - [MsgVpnMqttSessionClearStats](docs/MsgVpnMqttSessionClearStats.md)
 - [MsgVpnMqttSessionLinks](docs/MsgVpnMqttSessionLinks.md)
 - [MsgVpnMqttSessionResponse](docs/MsgVpnMqttSessionResponse.md)
 - [MsgVpnMqttSessionsResponse](docs/MsgVpnMqttSessionsResponse.md)
 - [MsgVpnQueue](docs/MsgVpnQueue.md)
 - [MsgVpnQueueCancelReplay](docs/MsgVpnQueueCancelReplay.md)
 - [MsgVpnQueueClearStats](docs/MsgVpnQueueClearStats.md)
 - [MsgVpnQueueCopyMsgFromQueue](docs/MsgVpnQueueCopyMsgFromQueue.md)
 - [MsgVpnQueueCopyMsgFromReplayLog](docs/MsgVpnQueueCopyMsgFromReplayLog.md)
 - [MsgVpnQueueCopyMsgFromTopicEndpoint](docs/MsgVpnQueueCopyMsgFromTopicEndpoint.md)
 - [MsgVpnQueueDeleteMsgs](docs/MsgVpnQueueDeleteMsgs.md)
 - [MsgVpnQueueLinks](docs/MsgVpnQueueLinks.md)
 - [MsgVpnQueueMsg](docs/MsgVpnQueueMsg.md)
 - [MsgVpnQueueMsgDelete](docs/MsgVpnQueueMsgDelete.md)
 - [MsgVpnQueueMsgLinks](docs/MsgVpnQueueMsgLinks.md)
 - [MsgVpnQueueMsgResponse](docs/MsgVpnQueueMsgResponse.md)
 - [MsgVpnQueueMsgsResponse](docs/MsgVpnQueueMsgsResponse.md)
 - [MsgVpnQueueResponse](docs/MsgVpnQueueResponse.md)
 - [MsgVpnQueueStartReplay](docs/MsgVpnQueueStartReplay.md)
 - [MsgVpnQueuesResponse](docs/MsgVpnQueuesResponse.md)
 - [MsgVpnReplayLog](docs/MsgVpnReplayLog.md)
 - [MsgVpnReplayLogLinks](docs/MsgVpnReplayLogLinks.md)
 - [MsgVpnReplayLogResponse](docs/MsgVpnReplayLogResponse.md)
 - [MsgVpnReplayLogTrimLoggedMsgs](docs/MsgVpnReplayLogTrimLoggedMsgs.md)
 - [MsgVpnReplayLogsResponse](docs/MsgVpnReplayLogsResponse.md)
 - [MsgVpnResponse](docs/MsgVpnResponse.md)
 - [MsgVpnRestDeliveryPoint](docs/MsgVpnRestDeliveryPoint.md)
 - [MsgVpnRestDeliveryPointLinks](docs/MsgVpnRestDeliveryPointLinks.md)
 - [MsgVpnRestDeliveryPointResponse](docs/MsgVpnRestDeliveryPointResponse.md)
 - [MsgVpnRestDeliveryPointRestConsumer](docs/MsgVpnRestDeliveryPointRestConsumer.md)
 - [MsgVpnRestDeliveryPointRestConsumerClearStats](docs/MsgVpnRestDeliveryPointRestConsumerClearStats.md)
 - [MsgVpnRestDeliveryPointRestConsumerLinks](docs/MsgVpnRestDeliveryPointRestConsumerLinks.md)
 - [MsgVpnRestDeliveryPointRestConsumerResponse](docs/MsgVpnRestDeliveryPointRestConsumerResponse.md)
 - [MsgVpnRestDeliveryPointRestConsumersResponse](docs/MsgVpnRestDeliveryPointRestConsumersResponse.md)
 - [MsgVpnRestDeliveryPointsResponse](docs/MsgVpnRestDeliveryPointsResponse.md)
 - [MsgVpnTopicEndpoint](docs/MsgVpnTopicEndpoint.md)
 - [MsgVpnTopicEndpointCancelReplay](docs/MsgVpnTopicEndpointCancelReplay.md)
 - [MsgVpnTopicEndpointClearStats](docs/MsgVpnTopicEndpointClearStats.md)
 - [MsgVpnTopicEndpointCopyMsgFromQueue](docs/MsgVpnTopicEndpointCopyMsgFromQueue.md)
 - [MsgVpnTopicEndpointCopyMsgFromReplayLog](docs/MsgVpnTopicEndpointCopyMsgFromReplayLog.md)
 - [MsgVpnTopicEndpointCopyMsgFromTopicEndpoint](docs/MsgVpnTopicEndpointCopyMsgFromTopicEndpoint.md)
 - [MsgVpnTopicEndpointDeleteMsgs](docs/MsgVpnTopicEndpointDeleteMsgs.md)
 - [MsgVpnTopicEndpointLinks](docs/MsgVpnTopicEndpointLinks.md)
 - [MsgVpnTopicEndpointMsg](docs/MsgVpnTopicEndpointMsg.md)
 - [MsgVpnTopicEndpointMsgDelete](docs/MsgVpnTopicEndpointMsgDelete.md)
 - [MsgVpnTopicEndpointMsgLinks](docs/MsgVpnTopicEndpointMsgLinks.md)
 - [MsgVpnTopicEndpointMsgResponse](docs/MsgVpnTopicEndpointMsgResponse.md)
 - [MsgVpnTopicEndpointMsgsResponse](docs/MsgVpnTopicEndpointMsgsResponse.md)
 - [MsgVpnTopicEndpointResponse](docs/MsgVpnTopicEndpointResponse.md)
 - [MsgVpnTopicEndpointStartReplay](docs/MsgVpnTopicEndpointStartReplay.md)
 - [MsgVpnTopicEndpointsResponse](docs/MsgVpnTopicEndpointsResponse.md)
 - [MsgVpnTransaction](docs/MsgVpnTransaction.md)
 - [MsgVpnTransactionCommit](docs/MsgVpnTransactionCommit.md)
 - [MsgVpnTransactionDelete](docs/MsgVpnTransactionDelete.md)
 - [MsgVpnTransactionLinks](docs/MsgVpnTransactionLinks.md)
 - [MsgVpnTransactionResponse](docs/MsgVpnTransactionResponse.md)
 - [MsgVpnTransactionRollback](docs/MsgVpnTransactionRollback.md)
 - [MsgVpnTransactionsResponse](docs/MsgVpnTransactionsResponse.md)
 - [MsgVpnsResponse](docs/MsgVpnsResponse.md)
 - [OauthProfile](docs/OauthProfile.md)
 - [OauthProfileClearStats](docs/OauthProfileClearStats.md)
 - [OauthProfileLinks](docs/OauthProfileLinks.md)
 - [OauthProfileResponse](docs/OauthProfileResponse.md)
 - [OauthProfilesResponse](docs/OauthProfilesResponse.md)
 - [SempError](docs/SempError.md)
 - [SempMeta](docs/SempMeta.md)
 - [SempMetaOnlyResponse](docs/SempMetaOnlyResponse.md)
 - [SempPaging](docs/SempPaging.md)
 - [SempRequest](docs/SempRequest.md)
 - [Session](docs/Session.md)
 - [SessionDelete](docs/SessionDelete.md)
 - [SessionLinks](docs/SessionLinks.md)
 - [SessionResponse](docs/SessionResponse.md)
 - [SessionsResponse](docs/SessionsResponse.md)

## Documentation For Authorization

## basicAuth
- **Type**: HTTP basic authentication

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
	UserName: "username",
	Password: "password",
})
r, err := client.Service.Operation(auth, args)
```

## Author

support@solace.com
