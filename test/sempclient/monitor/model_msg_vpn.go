/*
 * SEMP (Solace Element Management Protocol)
 *
 * SEMP (starting in `v2`, see note 1) is a RESTful API for configuring, monitoring, and administering a Solace PubSub+ broker.  SEMP uses URIs to address manageable **resources** of the Solace PubSub+ broker. Resources are individual **objects**, **collections** of objects, or (exclusively in the action API) **actions**. This document applies to the following API:   API|Base Path|Purpose|Comments :---|:---|:---|:--- Monitoring|/SEMP/v2/monitor|Querying operational parameters|See note 2    The following APIs are also available:   API|Base Path|Purpose|Comments :---|:---|:---|:--- Action|/SEMP/v2/action|Performing actions|See note 2 Configuration|/SEMP/v2/config|Reading and writing config state|See note 2    Resources are always nouns, with individual objects being singular and collections being plural.  Objects within a collection are identified by an `obj-id`, which follows the collection name with the form `collection-name/obj-id`.  Actions within an object are identified by an `action-id`, which follows the object name with the form `obj-id/action-id`.  Some examples:  ``` /SEMP/v2/config/msgVpns                        ; MsgVpn collection /SEMP/v2/config/msgVpns/a                      ; MsgVpn object named \"a\" /SEMP/v2/config/msgVpns/a/queues               ; Queue collection in MsgVpn \"a\" /SEMP/v2/config/msgVpns/a/queues/b             ; Queue object named \"b\" in MsgVpn \"a\" /SEMP/v2/action/msgVpns/a/queues/b/startReplay ; Action that starts a replay on Queue \"b\" in MsgVpn \"a\" /SEMP/v2/monitor/msgVpns/a/clients             ; Client collection in MsgVpn \"a\" /SEMP/v2/monitor/msgVpns/a/clients/c           ; Client object named \"c\" in MsgVpn \"a\" ```  ## Collection Resources  Collections are unordered lists of objects (unless described as otherwise), and are described by JSON arrays. Each item in the array represents an object in the same manner as the individual object would normally be represented. In the configuration API, the creation of a new object is done through its collection resource.  ## Object and Action Resources  Objects are composed of attributes, actions, collections, and other objects. They are described by JSON objects as name/value pairs. The collections and actions of an object are not contained directly in the object's JSON content; rather the content includes an attribute containing a URI which points to the collections and actions. These contained resources must be managed through this URI. At a minimum, every object has one or more identifying attributes, and its own `uri` attribute which contains the URI pointing to itself.  Actions are also composed of attributes, and are described by JSON objects as name/value pairs. Unlike objects, however, they are not members of a collection and cannot be retrieved, only performed. Actions only exist in the action API.  Attributes in an object or action may have any combination of the following properties:   Property|Meaning|Comments :---|:---|:--- Identifying|Attribute is involved in unique identification of the object, and appears in its URI| Const|Attribute value can only be chosen during object creation| Required|Attribute must be provided in the request| Read-Only|Attribute can only be read, not written.|See note 3 Write-Only|Attribute can only be written, not read, unless the attribute is also opaque|See the documentation for the opaque property Requires-Disable|Attribute cannot be changed while the object (or the relevant part of the object) is administratively enabled| Auto-Disable|Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as one or more attributes will be temporarily disabled to apply the change| Deprecated|Attribute is deprecated, and will disappear in the next SEMP version| Opaque|Attribute can be set or retrieved in opaque form when the `opaquePassword` query parameter is present|See the `opaquePassword` query parameter documentation    In some requests, certain attributes may only be provided in certain combinations with other attributes:   Relationship|Meaning :---|:--- Requires|Attribute may only be changed by a request if a particular attribute or combination of attributes is also provided in the request Conflicts|Attribute may only be provided in a request if a particular attribute or combination of attributes is not also provided in the request    In the monitoring API, any non-identifying attribute may not be returned in a GET.  ## HTTP Methods  The following HTTP methods manipulate resources in accordance with these general principles. Note that some methods are only used in certain APIs:   Method|Resource|Meaning|Request Body|Response Body|Notes :---|:---|:---|:---|:---|:--- POST|Collection|Create object|Initial attribute values|Object attributes and metadata|Absent attributes are set to default. If object already exists, a 400 error is returned PUT|Object|Update object|New attribute values|Object attributes and metadata|If does not exist, the object is first created. Absent attributes are set to default, with certain exceptions (see note 4) PUT|Action|Performs action|Action arguments|Action metadata| PATCH|Object|Update object|New attribute values|Object attributes and metadata|Absent attributes are left unchanged. If the object does not exist, a 404 error is returned DELETE|Object|Delete object|Empty|Object metadata|If the object does not exist, a 404 is returned GET|Object|Get object|Empty|Object attributes and metadata|If the object does not exist, a 404 is returned GET|Collection|Get collection|Empty|Object attributes and collection metadata|If the collection is empty, then an empty collection is returned with a 200 code    ## Common Query Parameters  The following are some common query parameters that are supported by many method/URI combinations. Individual URIs may document additional parameters. Note that multiple query parameters can be used together in a single URI, separated by the ampersand character. For example:  ``` ; Request for the MsgVpns collection using two hypothetical query parameters ; \"q1\" and \"q2\" with values \"val1\" and \"val2\" respectively /SEMP/v2/monitor/msgVpns?q1=val1&q2=val2 ```  ### select  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. Use this query parameter to limit the size of the returned data for each returned object, return only those fields that are desired, or exclude fields that are not desired.  The value of `select` is a comma-separated list of attribute names. If the list contains attribute names that are not prefaced by `-`, only those attributes are included in the response. If the list contains attribute names that are prefaced by `-`, those attributes are excluded from the response. If the list contains both types, then the difference of the first set of attributes and the second set of attributes is returned. If the list is empty (i.e. `select=`), it is treated the same as if no `select` was provided: all attribute are returned.  All attributes that are prefaced by `-` must follow all attributes that are not prefaced by `-`. In addition, each attribute name in the list must match at least one attribute in the object.  Names may include the `*` wildcard (zero or more characters). Nested attribute names are supported using periods (e.g. `parentName.childName`).  Some examples:  ``` ; List of all MsgVpn names /SEMP/v2/monitor/msgVpns?select=msgVpnName ; List of all MsgVpn and their attributes except for their names /SEMP/v2/monitor/msgVpns?select=-msgVpnName ; Authentication attributes of MsgVpn \"finance\" /SEMP/v2/monitor/msgVpns/finance?select=authentication%2A ; All attributes of MsgVpn \"finance\" except for authentication attributes /SEMP/v2/monitor/msgVpns/finance?select=-authentication%2A ; Access related attributes of Queue \"orderQ\" of MsgVpn \"finance\" /SEMP/v2/monitor/msgVpns/finance/queues/orderQ?select=owner,permission ```  ### where  Include in the response only objects where certain conditions are true. Use this query parameter to limit which objects are returned to those whose attribute values meet the given conditions.  The value of `where` is a comma-separated list of expressions. All expressions must be true for the object to be included in the response. Each expression takes the form:  ``` expression  = attribute-name OP value OP          = '==' | '!=' | '<' | '>' | '<=' | '>=' ```  `value` may be a number, string, `true`, or `false`, as appropriate for the type of `attribute-name`. Greater-than and less-than comparisons only work for numbers. A `*` in a string `value` is interpreted as a wildcard (zero or more characters). Some examples:  ``` ; Only enabled MsgVpns /SEMP/v2/monitor/msgVpns?where=enabled%3D%3Dtrue ; Only MsgVpns using basic non-LDAP authentication /SEMP/v2/monitor/msgVpns?where=authenticationBasicEnabled%3D%3Dtrue,authenticationBasicType%21%3Dldap ; Only MsgVpns that allow more than 100 client connections /SEMP/v2/monitor/msgVpns?where=maxConnectionCount%3E100 ; Only MsgVpns with msgVpnName starting with \"B\": /SEMP/v2/monitor/msgVpns?where=msgVpnName%3D%3DB%2A ```  ### count  Limit the count of objects in the response. This can be useful to limit the size of the response for large collections. The minimum value for `count` is `1` and the default is `10`. There is also a per-collection maximum value to limit request handling time.  `count` does not guarantee that a minimum number of objects will be returned. A page may contain fewer than `count` objects or even be empty. Additional objects may nonetheless be available for retrieval on subsequent pages. See the `cursor` query parameter documentation for more information on paging.  For example: ``` ; Up to 25 MsgVpns /SEMP/v2/monitor/msgVpns?count=25 ```  ### cursor  The cursor, or position, for the next page of objects. Cursors are opaque data that should not be created or interpreted by SEMP clients, and should only be used as described below.  When a request is made for a collection and there may be additional objects available for retrieval that are not included in the initial response, the response will include a `cursorQuery` field containing a cursor. The value of this field can be specified in the `cursor` query parameter of a subsequent request to retrieve the next page of objects.  Applications must continue to use the `cursorQuery` if one is provided in order to retrieve the full set of objects associated with the request, even if a page contains fewer than the requested number of objects (see the `count` query parameter documentation) or is empty.  ### opaquePassword  Attributes with the opaque property are also write-only and so cannot normally be retrieved in a GET. However, when a password is provided in the `opaquePassword` query parameter, attributes with the opaque property are retrieved in a GET in opaque form, encrypted with this password. The query parameter can also be used on a POST, PATCH, or PUT to set opaque attributes using opaque attribute values retrieved in a GET, so long as:  1. the same password that was used to retrieve the opaque attribute values is provided; and  2. the broker to which the request is being sent has the same major and minor SEMP version as the broker that produced the opaque attribute values.  The password provided in the query parameter must be a minimum of 8 characters and a maximum of 128 characters.  The query parameter can only be used in the configuration API, and only over HTTPS.  ## Authentication  When a client makes its first SEMPv2 request, it must supply a username and password using HTTP Basic authentication, or an OAuth token or tokens using HTTP Bearer authentication.  When HTTP Basic authentication is used, the broker returns a cookie containing a session key. The client can omit the username and password from subsequent requests, because the broker can use the session cookie for authentication instead. When the session expires or is deleted, the client must provide the username and password again, and the broker creates a new session.  There are a limited number of session slots available on the broker. The broker returns 529 No SEMP Session Available if it is not able to allocate a session.  If certain attributes—such as a user's password—are changed, the broker automatically deletes the affected sessions. These attributes are documented below. However, changes in external user configuration data stored on a RADIUS or LDAP server do not trigger the broker to delete the associated session(s), therefore you must do this manually, if required.  A client can retrieve its current session information using the /about/user endpoint and delete its own session using the /about/user/logout endpoint. A client with appropriate permissions can also manage all sessions using the /sessions endpoint.  Sessions are not created when authenticating with an OAuth token or tokens using HTTP Bearer authentication. If a session cookie is provided, it is ignored.  ## Help  Visit [our website](https://solace.com) to learn more about Solace.  You can also download the SEMP API specifications by clicking [here](https://solace.com/downloads/).  If you need additional support, please contact us at [support@solace.com](mailto:support@solace.com).  ## Notes  Note|Description :---:|:--- 1|This specification defines SEMP starting in \"v2\", and not the original SEMP \"v1\" interface. Request and response formats between \"v1\" and \"v2\" are entirely incompatible, although both protocols share a common port configuration on the Solace PubSub+ broker. They are differentiated by the initial portion of the URI path, one of either \"/SEMP/\" or \"/SEMP/v2/\" 2|This API is partially implemented. Only a subset of all objects are available. 3|Read-only attributes may appear in POST and PUT/PATCH requests. However, if a read-only attribute is not marked as identifying, it will be ignored during a PUT/PATCH. 4|On a PUT, if the SEMP user is not authorized to modify the attribute, its value is left unchanged rather than set to default. In addition, the values of write-only attributes are not set to their defaults on a PUT, except in the following two cases: there is a mutual requires relationship with another non-write-only attribute, both attributes are absent from the request, and the non-write-only attribute is not currently set to its default value; or the attribute is also opaque and the `opaquePassword` query parameter is provided in the request.  
 *
 * API version: 2.34
 * Contact: support@solace.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package monitor

type MsgVpn struct {
	// The name of another Message VPN which this Message VPN is an alias for. Available since 2.14.
	Alias string `json:"alias,omitempty"`
	// Indicates whether basic authentication is enabled for clients connecting to the Message VPN.
	AuthenticationBasicEnabled **bool `json:"authenticationBasicEnabled,omitempty"`
	// The name of the RADIUS or LDAP Profile to use for basic authentication.
	AuthenticationBasicProfileName string `json:"authenticationBasicProfileName,omitempty"`
	// The RADIUS domain to use for basic authentication.
	AuthenticationBasicRadiusDomain string `json:"authenticationBasicRadiusDomain,omitempty"`
	// The type of basic authentication to use for clients connecting to the Message VPN. The allowed values and their meaning are:  <pre> \"internal\" - Internal database. Authentication is against Client Usernames. \"ldap\" - LDAP authentication. An LDAP profile name must be provided. \"radius\" - RADIUS authentication. A RADIUS profile name must be provided. \"none\" - No authentication. Anonymous login allowed. </pre> 
	AuthenticationBasicType string `json:"authenticationBasicType,omitempty"`
	// Indicates whether a client is allowed to specify a Client Username via the API connect method. When disabled, the certificate CN (Common Name) is always used.
	AuthenticationClientCertAllowApiProvidedUsernameEnabled **bool `json:"authenticationClientCertAllowApiProvidedUsernameEnabled,omitempty"`
	// Indicates whether client certificate matching rules are enabled in the Message VPN. Available since 2.27.
	AuthenticationClientCertCertificateMatchingRulesEnabled **bool `json:"authenticationClientCertCertificateMatchingRulesEnabled,omitempty"`
	// Indicates whether client certificate authentication is enabled in the Message VPN.
	AuthenticationClientCertEnabled **bool `json:"authenticationClientCertEnabled,omitempty"`
	// The maximum depth for a client certificate chain. The depth of a chain is defined as the number of signing CA certificates that are present in the chain back to a trusted self-signed root CA certificate.
	AuthenticationClientCertMaxChainDepth int64 `json:"authenticationClientCertMaxChainDepth,omitempty"`
	// The desired behavior for client certificate revocation checking. The allowed values and their meaning are:  <pre> \"allow-all\" - Allow the client to authenticate, the result of client certificate revocation check is ignored. \"allow-unknown\" - Allow the client to authenticate even if the revocation status of his certificate cannot be determined. \"allow-valid\" - Allow the client to authenticate only when the revocation check returned an explicit positive response. </pre> 
	AuthenticationClientCertRevocationCheckMode string `json:"authenticationClientCertRevocationCheckMode,omitempty"`
	// The field from the client certificate to use as the client username. The allowed values and their meaning are:  <pre> \"certificate-thumbprint\" - The username is computed as the SHA-1 hash over the entire DER-encoded contents of the client certificate. \"common-name\" - The username is extracted from the certificate's first instance of the Common Name attribute in the Subject DN. \"common-name-last\" - The username is extracted from the certificate's last instance of the Common Name attribute in the Subject DN. \"subject-alternate-name-msupn\" - The username is extracted from the certificate's Other Name type of the Subject Alternative Name and must have the msUPN signature. \"uid\" - The username is extracted from the certificate's first instance of the User Identifier attribute in the Subject DN. \"uid-last\" - The username is extracted from the certificate's last instance of the User Identifier attribute in the Subject DN. </pre> 
	AuthenticationClientCertUsernameSource string `json:"authenticationClientCertUsernameSource,omitempty"`
	// Indicates whether the \"Not Before\" and \"Not After\" validity dates in the client certificate are checked.
	AuthenticationClientCertValidateDateEnabled **bool `json:"authenticationClientCertValidateDateEnabled,omitempty"`
	// Indicates whether a client is allowed to specify a Client Username via the API connect method. When disabled, the Kerberos Principal name is always used.
	AuthenticationKerberosAllowApiProvidedUsernameEnabled **bool `json:"authenticationKerberosAllowApiProvidedUsernameEnabled,omitempty"`
	// Indicates whether Kerberos authentication is enabled in the Message VPN.
	AuthenticationKerberosEnabled **bool `json:"authenticationKerberosEnabled,omitempty"`
	// The name of the profile to use when the client does not supply a profile name. Available since 2.25.
	AuthenticationOauthDefaultProfileName string `json:"authenticationOauthDefaultProfileName,omitempty"`
	// The name of the provider to use when the client does not supply a provider name. Deprecated since 2.25. authenticationOauthDefaultProviderName and authenticationOauthProviders replaced by authenticationOauthDefaultProfileName and authenticationOauthProfiles.
	AuthenticationOauthDefaultProviderName string `json:"authenticationOauthDefaultProviderName,omitempty"`
	// Indicates whether OAuth authentication is enabled. Available since 2.13.
	AuthenticationOauthEnabled **bool `json:"authenticationOauthEnabled,omitempty"`
	// The name of the attribute that is retrieved from the LDAP server as part of the LDAP search when authorizing a client connecting to the Message VPN.
	AuthorizationLdapGroupMembershipAttributeName string `json:"authorizationLdapGroupMembershipAttributeName,omitempty"`
	// Indicates whether client-username domain trimming for LDAP lookups of client connections is enabled. Available since 2.13.
	AuthorizationLdapTrimClientUsernameDomainEnabled **bool `json:"authorizationLdapTrimClientUsernameDomainEnabled,omitempty"`
	// The name of the LDAP Profile to use for client authorization.
	AuthorizationProfileName string `json:"authorizationProfileName,omitempty"`
	// The type of authorization to use for clients connecting to the Message VPN. The allowed values and their meaning are:  <pre> \"ldap\" - LDAP authorization. \"internal\" - Internal authorization. </pre> 
	AuthorizationType string `json:"authorizationType,omitempty"`
	// The one minute average of the bind request rate received by the MsgVpn, in binds per second (binds/sec). Available since 2.25.
	AverageBindRequestRate int64 `json:"averageBindRequestRate,omitempty"`
	// The one minute average of the message rate received by the Message VPN, in bytes per second (B/sec). Available since 2.13.
	AverageRxByteRate int64 `json:"averageRxByteRate,omitempty"`
	// The one minute average of the compressed message rate received by the Message VPN, in bytes per second (B/sec). Available since 2.12.
	AverageRxCompressedByteRate int64 `json:"averageRxCompressedByteRate,omitempty"`
	// The one minute average of the message rate received by the Message VPN, in messages per second (msg/sec). Available since 2.13.
	AverageRxMsgRate int64 `json:"averageRxMsgRate,omitempty"`
	// The one minute average of the uncompressed message rate received by the Message VPN, in bytes per second (B/sec). Available since 2.12.
	AverageRxUncompressedByteRate int64 `json:"averageRxUncompressedByteRate,omitempty"`
	// The one minute average of the message rate transmitted by the Message VPN, in bytes per second (B/sec). Available since 2.13.
	AverageTxByteRate int64 `json:"averageTxByteRate,omitempty"`
	// The one minute average of the compressed message rate transmitted by the Message VPN, in bytes per second (B/sec). Available since 2.12.
	AverageTxCompressedByteRate int64 `json:"averageTxCompressedByteRate,omitempty"`
	// The one minute average of the message rate transmitted by the Message VPN, in messages per second (msg/sec). Available since 2.13.
	AverageTxMsgRate int64 `json:"averageTxMsgRate,omitempty"`
	// The one minute average of the uncompressed message rate transmitted by the Message VPN, in bytes per second (B/sec). Available since 2.12.
	AverageTxUncompressedByteRate int64 `json:"averageTxUncompressedByteRate,omitempty"`
	// The current bind request rate received by the MsgVpn, in binds per second (binds/sec). Available since 2.25.
	BindRequestRate int64 `json:"bindRequestRate,omitempty"`
	// Indicates whether the Common Name (CN) in the server certificate from the remote broker is validated for the Bridge. Deprecated since 2.18. Common Name validation has been replaced by Server Certificate Name validation.
	BridgingTlsServerCertEnforceTrustedCommonNameEnabled **bool `json:"bridgingTlsServerCertEnforceTrustedCommonNameEnabled,omitempty"`
	// The maximum depth for a server certificate chain. The depth of a chain is defined as the number of signing CA certificates that are present in the chain back to a trusted self-signed root CA certificate.
	BridgingTlsServerCertMaxChainDepth int64 `json:"bridgingTlsServerCertMaxChainDepth,omitempty"`
	// Indicates whether the \"Not Before\" and \"Not After\" validity dates in the server certificate are checked.
	BridgingTlsServerCertValidateDateEnabled **bool `json:"bridgingTlsServerCertValidateDateEnabled,omitempty"`
	// Enable or disable the standard TLS authentication mechanism of verifying the name used to connect to the bridge. If enabled, the name used to connect to the bridge is checked against the names specified in the certificate returned by the remote router. Legacy Common Name validation is not performed if Server Certificate Name Validation is enabled, even if Common Name validation is also enabled. Available since 2.18.
	BridgingTlsServerCertValidateNameEnabled **bool `json:"bridgingTlsServerCertValidateNameEnabled,omitempty"`
	// The key for the config sync table of the local Message VPN. Deprecated since 2.22. This attribute has been deprecated.
	ConfigSyncLocalKey string `json:"configSyncLocalKey,omitempty"`
	// The result of the last operation on the config sync table of the local Message VPN. Deprecated since 2.22. This attribute has been replaced by 'lastResult' in the ConfigSyncLocalDatabaseRow object.
	ConfigSyncLocalLastResult string `json:"configSyncLocalLastResult,omitempty"`
	// The role of the config sync table of the local Message VPN. The allowed values and their meaning are:  <pre> \"unknown\" - The role is unknown. \"primary\" - Acts as the primary source of config data. \"replica\" - Acts as a replica of the primary config data. </pre>  Deprecated since 2.22. This attribute has been replaced by 'role' in the ConfigSyncLocalDatabaseRow object.
	ConfigSyncLocalRole string `json:"configSyncLocalRole,omitempty"`
	// The state of the config sync table of the local Message VPN. The allowed values and their meaning are:  <pre> \"unknown\" - The state is unknown. \"in-sync\" - The config data is synchronized between Message VPNs. \"reconciling\" - The config data is reconciling between Message VPNs. \"blocked\" - The config data is blocked from reconciling due to an error. \"out-of-sync\" - The config data is out of sync between Message VPNs. \"down\" - The state is down due to configuration. </pre>  Deprecated since 2.22. This attribute has been replaced by 'syncStatus' in the ConfigSyncLocalDatabaseRow object.
	ConfigSyncLocalState string `json:"configSyncLocalState,omitempty"`
	// The amount of time in seconds the config sync table of the local Message VPN has been in the current state. Deprecated since 2.22. This attribute has been replaced by 'timeInState' in the ConfigSyncLocalDatabaseRow object.
	ConfigSyncLocalTimeInState int32 `json:"configSyncLocalTimeInState,omitempty"`
	// The amount of client control messages received from clients by the Message VPN, in bytes (B). Available since 2.13.
	ControlRxByteCount int64 `json:"controlRxByteCount,omitempty"`
	// The number of client control messages received from clients by the Message VPN. Available since 2.13.
	ControlRxMsgCount int64 `json:"controlRxMsgCount,omitempty"`
	// The amount of client control messages transmitted to clients by the Message VPN, in bytes (B). Available since 2.13.
	ControlTxByteCount int64 `json:"controlTxByteCount,omitempty"`
	// The number of client control messages transmitted to clients by the Message VPN. Available since 2.13.
	ControlTxMsgCount int64 `json:"controlTxMsgCount,omitempty"`
	Counter *MsgVpnCounter `json:"counter,omitempty"`
	// The amount of client data messages received from clients by the Message VPN, in bytes (B). Available since 2.13.
	DataRxByteCount int64 `json:"dataRxByteCount,omitempty"`
	// The number of client data messages received from clients by the Message VPN. Available since 2.13.
	DataRxMsgCount int64 `json:"dataRxMsgCount,omitempty"`
	// The amount of client data messages transmitted to clients by the Message VPN, in bytes (B). Available since 2.13.
	DataTxByteCount int64 `json:"dataTxByteCount,omitempty"`
	// The number of client data messages transmitted to clients by the Message VPN. Available since 2.13.
	DataTxMsgCount int64 `json:"dataTxMsgCount,omitempty"`
	// The number of messages discarded during reception by the Message VPN. Available since 2.13.
	DiscardedRxMsgCount int64 `json:"discardedRxMsgCount,omitempty"`
	// The number of messages discarded during transmission by the Message VPN. Available since 2.13.
	DiscardedTxMsgCount int64 `json:"discardedTxMsgCount,omitempty"`
	// Indicates whether managing of cache instances over the message bus is enabled in the Message VPN. Deprecated since 2.28. Distributed cache mangement is now redundancy aware and thus no longer requires administrative intervention for operational state.
	DistributedCacheManagementEnabled **bool `json:"distributedCacheManagementEnabled,omitempty"`
	// Indicates whether Dynamic Message Routing (DMR) is enabled for the Message VPN.
	DmrEnabled **bool `json:"dmrEnabled,omitempty"`
	// Indicates whether the Message VPN is enabled.
	Enabled **bool `json:"enabled,omitempty"`
	EventConnectionCountThreshold *EventThreshold `json:"eventConnectionCountThreshold,omitempty"`
	EventEgressFlowCountThreshold *EventThreshold `json:"eventEgressFlowCountThreshold,omitempty"`
	EventEgressMsgRateThreshold *EventThresholdByValue `json:"eventEgressMsgRateThreshold,omitempty"`
	EventEndpointCountThreshold *EventThreshold `json:"eventEndpointCountThreshold,omitempty"`
	EventIngressFlowCountThreshold *EventThreshold `json:"eventIngressFlowCountThreshold,omitempty"`
	EventIngressMsgRateThreshold *EventThresholdByValue `json:"eventIngressMsgRateThreshold,omitempty"`
	// Exceeding this message size in kilobytes (KB) triggers a corresponding Event in the Message VPN.
	EventLargeMsgThreshold int64 `json:"eventLargeMsgThreshold,omitempty"`
	// The value of the prefix applied to all published Events in the Message VPN.
	EventLogTag string `json:"eventLogTag,omitempty"`
	EventMsgSpoolUsageThreshold *EventThreshold `json:"eventMsgSpoolUsageThreshold,omitempty"`
	// Indicates whether client Events are published in the Message VPN.
	EventPublishClientEnabled **bool `json:"eventPublishClientEnabled,omitempty"`
	// Indicates whether Message VPN Events are published in the Message VPN.
	EventPublishMsgVpnEnabled **bool `json:"eventPublishMsgVpnEnabled,omitempty"`
	// The mode of subscription Events published in the Message VPN. The allowed values and their meaning are:  <pre> \"off\" - Disable client level event message publishing. \"on-with-format-v1\" - Enable client level event message publishing with format v1. \"on-with-no-unsubscribe-events-on-disconnect-format-v1\" - As \"on-with-format-v1\", but unsubscribe events are not generated when a client disconnects. Unsubscribe events are still raised when a client explicitly unsubscribes from its subscriptions. \"on-with-format-v2\" - Enable client level event message publishing with format v2. \"on-with-no-unsubscribe-events-on-disconnect-format-v2\" - As \"on-with-format-v2\", but unsubscribe events are not generated when a client disconnects. Unsubscribe events are still raised when a client explicitly unsubscribes from its subscriptions. </pre> 
	EventPublishSubscriptionMode string `json:"eventPublishSubscriptionMode,omitempty"`
	// Indicates whether Message VPN Events are published in the MQTT format.
	EventPublishTopicFormatMqttEnabled **bool `json:"eventPublishTopicFormatMqttEnabled,omitempty"`
	// Indicates whether Message VPN Events are published in the SMF format.
	EventPublishTopicFormatSmfEnabled **bool `json:"eventPublishTopicFormatSmfEnabled,omitempty"`
	EventServiceAmqpConnectionCountThreshold *EventThreshold `json:"eventServiceAmqpConnectionCountThreshold,omitempty"`
	EventServiceMqttConnectionCountThreshold *EventThreshold `json:"eventServiceMqttConnectionCountThreshold,omitempty"`
	EventServiceRestIncomingConnectionCountThreshold *EventThreshold `json:"eventServiceRestIncomingConnectionCountThreshold,omitempty"`
	EventServiceSmfConnectionCountThreshold *EventThreshold `json:"eventServiceSmfConnectionCountThreshold,omitempty"`
	EventServiceWebConnectionCountThreshold *EventThreshold `json:"eventServiceWebConnectionCountThreshold,omitempty"`
	EventSubscriptionCountThreshold *EventThreshold `json:"eventSubscriptionCountThreshold,omitempty"`
	EventTransactedSessionCountThreshold *EventThreshold `json:"eventTransactedSessionCountThreshold,omitempty"`
	EventTransactionCountThreshold *EventThreshold `json:"eventTransactionCountThreshold,omitempty"`
	// Indicates whether exports of subscriptions to other routers in the network over neighbour links is enabled in the Message VPN.
	ExportSubscriptionsEnabled **bool `json:"exportSubscriptionsEnabled,omitempty"`
	// The reason for the Message VPN failure.
	FailureReason string `json:"failureReason,omitempty"`
	// Indicates whether the JNDI access for clients is enabled in the Message VPN.
	JndiEnabled **bool `json:"jndiEnabled,omitempty"`
	// The number of login request messages received by the Message VPN. Available since 2.13.
	LoginRxMsgCount int64 `json:"loginRxMsgCount,omitempty"`
	// The number of login response messages transmitted by the Message VPN. Available since 2.13.
	LoginTxMsgCount int64 `json:"loginTxMsgCount,omitempty"`
	// The maximum number of client connections to the Message VPN.
	MaxConnectionCount int64 `json:"maxConnectionCount,omitempty"`
	// The effective maximum number of Queues and Topic Endpoints allowed in the Message VPN.
	MaxEffectiveEndpointCount int32 `json:"maxEffectiveEndpointCount,omitempty"`
	// The effective maximum number of receive flows allowed in the Message VPN.
	MaxEffectiveRxFlowCount int32 `json:"maxEffectiveRxFlowCount,omitempty"`
	// The effective maximum number of subscriptions allowed in the Message VPN.
	MaxEffectiveSubscriptionCount int64 `json:"maxEffectiveSubscriptionCount,omitempty"`
	// The effective maximum number of transacted sessions allowed in the Message VPN.
	MaxEffectiveTransactedSessionCount int32 `json:"maxEffectiveTransactedSessionCount,omitempty"`
	// The effective maximum number of transactions allowed in the Message VPN.
	MaxEffectiveTransactionCount int32 `json:"maxEffectiveTransactionCount,omitempty"`
	// The effective maximum number of transmit flows allowed in the Message VPN.
	MaxEffectiveTxFlowCount int32 `json:"maxEffectiveTxFlowCount,omitempty"`
	// The maximum number of transmit flows that can be created in the Message VPN.
	MaxEgressFlowCount int64 `json:"maxEgressFlowCount,omitempty"`
	// The maximum number of Queues and Topic Endpoints that can be created in the Message VPN.
	MaxEndpointCount int64 `json:"maxEndpointCount,omitempty"`
	// The maximum number of receive flows that can be created in the Message VPN.
	MaxIngressFlowCount int64 `json:"maxIngressFlowCount,omitempty"`
	// The maximum message spool usage by the Message VPN, in megabytes.
	MaxMsgSpoolUsage int64 `json:"maxMsgSpoolUsage,omitempty"`
	// The maximum number of local client subscriptions that can be added to the Message VPN. This limit is not enforced when a subscription is added using a management interface, such as CLI or SEMP.
	MaxSubscriptionCount int64 `json:"maxSubscriptionCount,omitempty"`
	// The maximum number of transacted sessions that can be created in the Message VPN.
	MaxTransactedSessionCount int64 `json:"maxTransactedSessionCount,omitempty"`
	// The maximum number of transactions that can be created in the Message VPN.
	MaxTransactionCount int64 `json:"maxTransactionCount,omitempty"`
	// The maximum total memory usage of the MQTT Retain feature for this Message VPN, in MB. If the maximum memory is reached, any arriving retain messages that require more memory are discarded. A value of -1 indicates that the memory is bounded only by the global max memory limit. A value of 0 prevents MQTT Retain from becoming operational.
	MqttRetainMaxMemory int32 `json:"mqttRetainMaxMemory,omitempty"`
	// The number of message replays that are currently active in the Message VPN.
	MsgReplayActiveCount int32 `json:"msgReplayActiveCount,omitempty"`
	// The number of message replays that are currently failed in the Message VPN.
	MsgReplayFailedCount int32 `json:"msgReplayFailedCount,omitempty"`
	// The number of message replays that are currently initializing in the Message VPN.
	MsgReplayInitializingCount int32 `json:"msgReplayInitializingCount,omitempty"`
	// The number of message replays that are pending complete in the Message VPN.
	MsgReplayPendingCompleteCount int32 `json:"msgReplayPendingCompleteCount,omitempty"`
	// The current number of messages spooled (persisted in the Message Spool) in the Message VPN. Available since 2.14.
	MsgSpoolMsgCount int64 `json:"msgSpoolMsgCount,omitempty"`
	// The number of guaranteed messages received by the Message VPN. Available since 2.13.
	MsgSpoolRxMsgCount int64 `json:"msgSpoolRxMsgCount,omitempty"`
	// The number of guaranteed messages transmitted by the Message VPN. One message to multiple clients is counted as one message. Available since 2.13.
	MsgSpoolTxMsgCount int64 `json:"msgSpoolTxMsgCount,omitempty"`
	// The current message spool usage by the Message VPN, in bytes (B).
	MsgSpoolUsage int64 `json:"msgSpoolUsage,omitempty"`
	// The name of the Message VPN.
	MsgVpnName string `json:"msgVpnName,omitempty"`
	Rate *MsgVpnRate `json:"rate,omitempty"`
	// The acknowledgement (ACK) propagation interval for the replication Bridge, in number of replicated messages. Available since 2.12.
	ReplicationAckPropagationIntervalMsgCount int64 `json:"replicationAckPropagationIntervalMsgCount,omitempty"`
	// The number of acknowledgement messages propagated to the replication standby remote Message VPN. Available since 2.12.
	ReplicationActiveAckPropTxMsgCount int64 `json:"replicationActiveAckPropTxMsgCount,omitempty"`
	// The number of async messages queued to the replication standby remote Message VPN. Available since 2.12.
	ReplicationActiveAsyncQueuedMsgCount int64 `json:"replicationActiveAsyncQueuedMsgCount,omitempty"`
	// The number of messages consumed in the replication active local Message VPN. Available since 2.12.
	ReplicationActiveLocallyConsumedMsgCount int64 `json:"replicationActiveLocallyConsumedMsgCount,omitempty"`
	// The peak amount of time in seconds the message flow has been congested to the replication standby remote Message VPN. Available since 2.12.
	ReplicationActiveMateFlowCongestedPeakTime int32 `json:"replicationActiveMateFlowCongestedPeakTime,omitempty"`
	// The peak amount of time in seconds the message flow has not been congested to the replication standby remote Message VPN. Available since 2.12.
	ReplicationActiveMateFlowNotCongestedPeakTime int32 `json:"replicationActiveMateFlowNotCongestedPeakTime,omitempty"`
	// The number of promoted messages queued to the replication standby remote Message VPN. Available since 2.12.
	ReplicationActivePromotedQueuedMsgCount int64 `json:"replicationActivePromotedQueuedMsgCount,omitempty"`
	// The number of reconcile request messages received from the replication standby remote Message VPN. Available since 2.12.
	ReplicationActiveReconcileRequestRxMsgCount int64 `json:"replicationActiveReconcileRequestRxMsgCount,omitempty"`
	// The peak amount of time in seconds sync replication has been eligible to the replication standby remote Message VPN. Available since 2.12.
	ReplicationActiveSyncEligiblePeakTime int32 `json:"replicationActiveSyncEligiblePeakTime,omitempty"`
	// The peak amount of time in seconds sync replication has been ineligible to the replication standby remote Message VPN. Available since 2.12.
	ReplicationActiveSyncIneligiblePeakTime int32 `json:"replicationActiveSyncIneligiblePeakTime,omitempty"`
	// The number of sync messages queued as async to the replication standby remote Message VPN. Available since 2.12.
	ReplicationActiveSyncQueuedAsAsyncMsgCount int64 `json:"replicationActiveSyncQueuedAsAsyncMsgCount,omitempty"`
	// The number of sync messages queued to the replication standby remote Message VPN. Available since 2.12.
	ReplicationActiveSyncQueuedMsgCount int64 `json:"replicationActiveSyncQueuedMsgCount,omitempty"`
	// The number of sync replication ineligible transitions to the replication standby remote Message VPN. Available since 2.12.
	ReplicationActiveTransitionToSyncIneligibleCount int64 `json:"replicationActiveTransitionToSyncIneligibleCount,omitempty"`
	// The Client Username the replication Bridge uses to login to the remote Message VPN. Available since 2.12.
	ReplicationBridgeAuthenticationBasicClientUsername string `json:"replicationBridgeAuthenticationBasicClientUsername,omitempty"`
	// The authentication scheme for the replication Bridge in the Message VPN. The allowed values and their meaning are:  <pre> \"basic\" - Basic Authentication Scheme (via username and password). \"client-certificate\" - Client Certificate Authentication Scheme (via certificate file or content). </pre>  Available since 2.12.
	ReplicationBridgeAuthenticationScheme string `json:"replicationBridgeAuthenticationScheme,omitempty"`
	// Indicates whether the local replication Bridge is bound to the Queue in the remote Message VPN. Available since 2.12.
	ReplicationBridgeBoundToQueue **bool `json:"replicationBridgeBoundToQueue,omitempty"`
	// Indicates whether compression is used for the replication Bridge. Available since 2.12.
	ReplicationBridgeCompressedDataEnabled **bool `json:"replicationBridgeCompressedDataEnabled,omitempty"`
	// The size of the window used for guaranteed messages published to the replication Bridge, in messages. Available since 2.12.
	ReplicationBridgeEgressFlowWindowSize int64 `json:"replicationBridgeEgressFlowWindowSize,omitempty"`
	// The name of the local replication Bridge in the Message VPN. Available since 2.12.
	ReplicationBridgeName string `json:"replicationBridgeName,omitempty"`
	// The number of seconds that must pass before retrying the replication Bridge connection. Available since 2.12.
	ReplicationBridgeRetryDelay int64 `json:"replicationBridgeRetryDelay,omitempty"`
	// Indicates whether encryption (TLS) is enabled for the replication Bridge connection. Available since 2.12.
	ReplicationBridgeTlsEnabled **bool `json:"replicationBridgeTlsEnabled,omitempty"`
	// The Client Profile for the unidirectional replication Bridge in the Message VPN. It is used only for the TCP parameters. Available since 2.12.
	ReplicationBridgeUnidirectionalClientProfileName string `json:"replicationBridgeUnidirectionalClientProfileName,omitempty"`
	// Indicates whether the local replication Bridge is operationally up in the Message VPN. Available since 2.12.
	ReplicationBridgeUp **bool `json:"replicationBridgeUp,omitempty"`
	// Indicates whether replication is enabled for the Message VPN. Available since 2.12.
	ReplicationEnabled **bool `json:"replicationEnabled,omitempty"`
	// Indicates whether the remote replication Bridge is bound to the Queue in the Message VPN. Available since 2.12.
	ReplicationQueueBound **bool `json:"replicationQueueBound,omitempty"`
	// The maximum message spool usage by the replication Bridge local Queue (quota), in megabytes. Available since 2.12.
	ReplicationQueueMaxMsgSpoolUsage int64 `json:"replicationQueueMaxMsgSpoolUsage,omitempty"`
	// Indicates whether messages discarded on this replication Bridge Queue are rejected back to the sender. Available since 2.12.
	ReplicationQueueRejectMsgToSenderOnDiscardEnabled **bool `json:"replicationQueueRejectMsgToSenderOnDiscardEnabled,omitempty"`
	// Indicates whether guaranteed messages published to synchronously replicated Topics are rejected back to the sender when synchronous replication becomes ineligible. Available since 2.12.
	ReplicationRejectMsgWhenSyncIneligibleEnabled **bool `json:"replicationRejectMsgWhenSyncIneligibleEnabled,omitempty"`
	// The name of the remote replication Bridge in the Message VPN. Available since 2.12.
	ReplicationRemoteBridgeName string `json:"replicationRemoteBridgeName,omitempty"`
	// Indicates whether the remote replication Bridge is operationally up in the Message VPN. Available since 2.12.
	ReplicationRemoteBridgeUp **bool `json:"replicationRemoteBridgeUp,omitempty"`
	// The replication role for the Message VPN. The allowed values and their meaning are:  <pre> \"active\" - Assume the Active role in replication for the Message VPN. \"standby\" - Assume the Standby role in replication for the Message VPN. </pre>  Available since 2.12.
	ReplicationRole string `json:"replicationRole,omitempty"`
	// The number of acknowledgement messages received out of sequence from the replication active remote Message VPN. Available since 2.12.
	ReplicationStandbyAckPropOutOfSeqRxMsgCount int64 `json:"replicationStandbyAckPropOutOfSeqRxMsgCount,omitempty"`
	// The number of acknowledgement messages received from the replication active remote Message VPN. Available since 2.12.
	ReplicationStandbyAckPropRxMsgCount int64 `json:"replicationStandbyAckPropRxMsgCount,omitempty"`
	// The number of reconcile request messages transmitted to the replication active remote Message VPN. Available since 2.12.
	ReplicationStandbyReconcileRequestTxMsgCount int64 `json:"replicationStandbyReconcileRequestTxMsgCount,omitempty"`
	// The number of messages received from the replication active remote Message VPN. Available since 2.12.
	ReplicationStandbyRxMsgCount int64 `json:"replicationStandbyRxMsgCount,omitempty"`
	// The number of transaction requests received from the replication active remote Message VPN. Available since 2.12.
	ReplicationStandbyTransactionRequestCount int64 `json:"replicationStandbyTransactionRequestCount,omitempty"`
	// The number of transaction requests received from the replication active remote Message VPN that failed. Available since 2.12.
	ReplicationStandbyTransactionRequestFailureCount int64 `json:"replicationStandbyTransactionRequestFailureCount,omitempty"`
	// The number of transaction requests received from the replication active remote Message VPN that succeeded. Available since 2.12.
	ReplicationStandbyTransactionRequestSuccessCount int64 `json:"replicationStandbyTransactionRequestSuccessCount,omitempty"`
	// Indicates whether sync replication is eligible in the Message VPN. Available since 2.12.
	ReplicationSyncEligible **bool `json:"replicationSyncEligible,omitempty"`
	// Indicates whether synchronous or asynchronous replication mode is used for all transactions within the Message VPN. The allowed values and their meaning are:  <pre> \"sync\" - Messages are acknowledged when replicated (spooled remotely). \"async\" - Messages are acknowledged when pending replication (spooled locally). </pre>  Available since 2.12.
	ReplicationTransactionMode string `json:"replicationTransactionMode,omitempty"`
	// Indicates whether the Common Name (CN) in the server certificate from the remote REST Consumer is validated. Deprecated since 2.17. Common Name validation has been replaced by Server Certificate Name validation.
	RestTlsServerCertEnforceTrustedCommonNameEnabled **bool `json:"restTlsServerCertEnforceTrustedCommonNameEnabled,omitempty"`
	// The maximum depth for a REST Consumer server certificate chain. The depth of a chain is defined as the number of signing CA certificates that are present in the chain back to a trusted self-signed root CA certificate.
	RestTlsServerCertMaxChainDepth int64 `json:"restTlsServerCertMaxChainDepth,omitempty"`
	// Indicates whether the \"Not Before\" and \"Not After\" validity dates in the REST Consumer server certificate are checked.
	RestTlsServerCertValidateDateEnabled **bool `json:"restTlsServerCertValidateDateEnabled,omitempty"`
	// Enable or disable the standard TLS authentication mechanism of verifying the name used to connect to the remote REST Consumer. If enabled, the name used to connect to the remote REST Consumer is checked against the names specified in the certificate returned by the remote router. Legacy Common Name validation is not performed if Server Certificate Name Validation is enabled, even if Common Name validation is also enabled. Available since 2.17.
	RestTlsServerCertValidateNameEnabled **bool `json:"restTlsServerCertValidateNameEnabled,omitempty"`
	// The amount of messages received from clients by the Message VPN, in bytes (B). Available since 2.12.
	RxByteCount int64 `json:"rxByteCount,omitempty"`
	// The current message rate received by the Message VPN, in bytes per second (B/sec). Available since 2.13.
	RxByteRate int64 `json:"rxByteRate,omitempty"`
	// The amount of compressed messages received by the Message VPN, in bytes (B). Available since 2.12.
	RxCompressedByteCount int64 `json:"rxCompressedByteCount,omitempty"`
	// The current compressed message rate received by the Message VPN, in bytes per second (B/sec). Available since 2.12.
	RxCompressedByteRate int64 `json:"rxCompressedByteRate,omitempty"`
	// The compression ratio for messages received by the message VPN. Available since 2.12.
	RxCompressionRatio string `json:"rxCompressionRatio,omitempty"`
	// The number of messages received from clients by the Message VPN. Available since 2.12.
	RxMsgCount int64 `json:"rxMsgCount,omitempty"`
	// The current message rate received by the Message VPN, in messages per second (msg/sec). Available since 2.13.
	RxMsgRate int64 `json:"rxMsgRate,omitempty"`
	// The amount of uncompressed messages received by the Message VPN, in bytes (B). Available since 2.12.
	RxUncompressedByteCount int64 `json:"rxUncompressedByteCount,omitempty"`
	// The current uncompressed message rate received by the Message VPN, in bytes per second (B/sec). Available since 2.12.
	RxUncompressedByteRate int64 `json:"rxUncompressedByteRate,omitempty"`
	// Indicates whether the \"admin\" level \"client\" commands are enabled for SEMP over the message bus in the Message VPN.
	SempOverMsgBusAdminClientEnabled **bool `json:"sempOverMsgBusAdminClientEnabled,omitempty"`
	// Indicates whether the \"admin\" level \"Distributed Cache\" commands are enabled for SEMP over the message bus in the Message VPN.
	SempOverMsgBusAdminDistributedCacheEnabled **bool `json:"sempOverMsgBusAdminDistributedCacheEnabled,omitempty"`
	// Indicates whether the \"admin\" level commands are enabled for SEMP over the message bus in the Message VPN.
	SempOverMsgBusAdminEnabled **bool `json:"sempOverMsgBusAdminEnabled,omitempty"`
	// Indicates whether SEMP over the message bus is enabled in the Message VPN.
	SempOverMsgBusEnabled **bool `json:"sempOverMsgBusEnabled,omitempty"`
	// Indicates whether the \"show\" level commands are enabled for SEMP over the message bus in the Message VPN.
	SempOverMsgBusShowEnabled **bool `json:"sempOverMsgBusShowEnabled,omitempty"`
	// The maximum number of AMQP client connections that can be simultaneously connected to the Message VPN. This value may be higher than supported by the platform.
	ServiceAmqpMaxConnectionCount int64 `json:"serviceAmqpMaxConnectionCount,omitempty"`
	// Indicates whether the AMQP Service is compressed in the Message VPN.
	ServiceAmqpPlainTextCompressed **bool `json:"serviceAmqpPlainTextCompressed,omitempty"`
	// Indicates whether the AMQP Service is enabled in the Message VPN.
	ServiceAmqpPlainTextEnabled **bool `json:"serviceAmqpPlainTextEnabled,omitempty"`
	// The reason for the AMQP Service failure in the Message VPN.
	ServiceAmqpPlainTextFailureReason string `json:"serviceAmqpPlainTextFailureReason,omitempty"`
	// The port number for plain-text AMQP clients that connect to the Message VPN. The port must be unique across the message backbone. A value of 0 means that the listen-port is unassigned and cannot be enabled.
	ServiceAmqpPlainTextListenPort int64 `json:"serviceAmqpPlainTextListenPort,omitempty"`
	// Indicates whether the AMQP Service is operationally up in the Message VPN.
	ServiceAmqpPlainTextUp **bool `json:"serviceAmqpPlainTextUp,omitempty"`
	// Indicates whether the TLS related AMQP Service is compressed in the Message VPN.
	ServiceAmqpTlsCompressed **bool `json:"serviceAmqpTlsCompressed,omitempty"`
	// Indicates whether encryption (TLS) is enabled for AMQP clients in the Message VPN.
	ServiceAmqpTlsEnabled **bool `json:"serviceAmqpTlsEnabled,omitempty"`
	// The reason for the TLS related AMQP Service failure in the Message VPN.
	ServiceAmqpTlsFailureReason string `json:"serviceAmqpTlsFailureReason,omitempty"`
	// The port number for AMQP clients that connect to the Message VPN over TLS. The port must be unique across the message backbone. A value of 0 means that the listen-port is unassigned and cannot be enabled.
	ServiceAmqpTlsListenPort int64 `json:"serviceAmqpTlsListenPort,omitempty"`
	// Indicates whether the TLS related AMQP Service is operationally up in the Message VPN.
	ServiceAmqpTlsUp **bool `json:"serviceAmqpTlsUp,omitempty"`
	// Determines when to request a client certificate from an incoming MQTT client connecting via a TLS port. The allowed values and their meaning are:  <pre> \"always\" - Always ask for a client certificate regardless of the \"message-vpn > authentication > client-certificate > shutdown\" configuration. \"never\" - Never ask for a client certificate regardless of the \"message-vpn > authentication > client-certificate > shutdown\" configuration. \"when-enabled-in-message-vpn\" - Only ask for a client-certificate if client certificate authentication is enabled under \"message-vpn >  authentication > client-certificate > shutdown\". </pre>  Available since 2.21.
	ServiceMqttAuthenticationClientCertRequest string `json:"serviceMqttAuthenticationClientCertRequest,omitempty"`
	// The maximum number of MQTT client connections that can be simultaneously connected to the Message VPN.
	ServiceMqttMaxConnectionCount int64 `json:"serviceMqttMaxConnectionCount,omitempty"`
	// Indicates whether the MQTT Service is compressed in the Message VPN.
	ServiceMqttPlainTextCompressed **bool `json:"serviceMqttPlainTextCompressed,omitempty"`
	// Indicates whether the MQTT Service is enabled in the Message VPN.
	ServiceMqttPlainTextEnabled **bool `json:"serviceMqttPlainTextEnabled,omitempty"`
	// The reason for the MQTT Service failure in the Message VPN.
	ServiceMqttPlainTextFailureReason string `json:"serviceMqttPlainTextFailureReason,omitempty"`
	// The port number for plain-text MQTT clients that connect to the Message VPN. The port must be unique across the message backbone. A value of 0 means that the listen-port is unassigned and cannot be enabled.
	ServiceMqttPlainTextListenPort int64 `json:"serviceMqttPlainTextListenPort,omitempty"`
	// Indicates whether the MQTT Service is operationally up in the Message VPN.
	ServiceMqttPlainTextUp **bool `json:"serviceMqttPlainTextUp,omitempty"`
	// Indicates whether the TLS related MQTT Service is compressed in the Message VPN.
	ServiceMqttTlsCompressed **bool `json:"serviceMqttTlsCompressed,omitempty"`
	// Indicates whether encryption (TLS) is enabled for MQTT clients in the Message VPN.
	ServiceMqttTlsEnabled **bool `json:"serviceMqttTlsEnabled,omitempty"`
	// The reason for the TLS related MQTT Service failure in the Message VPN.
	ServiceMqttTlsFailureReason string `json:"serviceMqttTlsFailureReason,omitempty"`
	// The port number for MQTT clients that connect to the Message VPN over TLS. The port must be unique across the message backbone. A value of 0 means that the listen-port is unassigned and cannot be enabled.
	ServiceMqttTlsListenPort int64 `json:"serviceMqttTlsListenPort,omitempty"`
	// Indicates whether the TLS related MQTT Service is operationally up in the Message VPN.
	ServiceMqttTlsUp **bool `json:"serviceMqttTlsUp,omitempty"`
	// Indicates whether the TLS related Web transport MQTT Service is compressed in the Message VPN.
	ServiceMqttTlsWebSocketCompressed **bool `json:"serviceMqttTlsWebSocketCompressed,omitempty"`
	// Indicates whether encryption (TLS) is enabled for MQTT Web clients in the Message VPN.
	ServiceMqttTlsWebSocketEnabled **bool `json:"serviceMqttTlsWebSocketEnabled,omitempty"`
	// The reason for the TLS related Web transport MQTT Service failure in the Message VPN.
	ServiceMqttTlsWebSocketFailureReason string `json:"serviceMqttTlsWebSocketFailureReason,omitempty"`
	// The port number for MQTT clients that connect to the Message VPN using WebSocket over TLS. The port must be unique across the message backbone. A value of 0 means that the listen-port is unassigned and cannot be enabled.
	ServiceMqttTlsWebSocketListenPort int64 `json:"serviceMqttTlsWebSocketListenPort,omitempty"`
	// Indicates whether the TLS related Web transport MQTT Service is operationally up in the Message VPN.
	ServiceMqttTlsWebSocketUp **bool `json:"serviceMqttTlsWebSocketUp,omitempty"`
	// Indicates whether the Web transport related MQTT Service is compressed in the Message VPN.
	ServiceMqttWebSocketCompressed **bool `json:"serviceMqttWebSocketCompressed,omitempty"`
	// Indicates whether the Web transport for the SMF Service is enabled in the Message VPN.
	ServiceMqttWebSocketEnabled **bool `json:"serviceMqttWebSocketEnabled,omitempty"`
	// The reason for the Web transport related MQTT Service failure in the Message VPN.
	ServiceMqttWebSocketFailureReason string `json:"serviceMqttWebSocketFailureReason,omitempty"`
	// The port number for plain-text MQTT clients that connect to the Message VPN using WebSocket. The port must be unique across the message backbone. A value of 0 means that the listen-port is unassigned and cannot be enabled.
	ServiceMqttWebSocketListenPort int64 `json:"serviceMqttWebSocketListenPort,omitempty"`
	// Indicates whether the Web transport related MQTT Service is operationally up in the Message VPN.
	ServiceMqttWebSocketUp **bool `json:"serviceMqttWebSocketUp,omitempty"`
	// Determines when to request a client certificate from an incoming REST Producer connecting via a TLS port. The allowed values and their meaning are:  <pre> \"always\" - Always ask for a client certificate regardless of the \"message-vpn > authentication > client-certificate > shutdown\" configuration. \"never\" - Never ask for a client certificate regardless of the \"message-vpn > authentication > client-certificate > shutdown\" configuration. \"when-enabled-in-message-vpn\" - Only ask for a client-certificate if client certificate authentication is enabled under \"message-vpn >  authentication > client-certificate > shutdown\". </pre>  Available since 2.21.
	ServiceRestIncomingAuthenticationClientCertRequest string `json:"serviceRestIncomingAuthenticationClientCertRequest,omitempty"`
	// The handling of Authorization headers for incoming REST connections. The allowed values and their meaning are:  <pre> \"drop\" - Do not attach the Authorization header to the message as a user property. This configuration is most secure. \"forward\" - Forward the Authorization header, attaching it to the message as a user property in the same way as other headers. For best security, use the drop setting. \"legacy\" - If the Authorization header was used for authentication to the broker, do not attach it to the message. If the Authorization header was not used for authentication to the broker, attach it to the message as a user property in the same way as other headers. For best security, use the drop setting. </pre>  Available since 2.19.
	ServiceRestIncomingAuthorizationHeaderHandling string `json:"serviceRestIncomingAuthorizationHeaderHandling,omitempty"`
	// The maximum number of REST incoming client connections that can be simultaneously connected to the Message VPN. This value may be higher than supported by the platform.
	ServiceRestIncomingMaxConnectionCount int64 `json:"serviceRestIncomingMaxConnectionCount,omitempty"`
	// Indicates whether the incoming REST Service is compressed in the Message VPN.
	ServiceRestIncomingPlainTextCompressed **bool `json:"serviceRestIncomingPlainTextCompressed,omitempty"`
	// Indicates whether the REST Service is enabled in the Message VPN for incoming clients.
	ServiceRestIncomingPlainTextEnabled **bool `json:"serviceRestIncomingPlainTextEnabled,omitempty"`
	// The reason for the incoming REST Service failure in the Message VPN.
	ServiceRestIncomingPlainTextFailureReason string `json:"serviceRestIncomingPlainTextFailureReason,omitempty"`
	// The port number for incoming plain-text REST clients that connect to the Message VPN. The port must be unique across the message backbone. A value of 0 means that the listen-port is unassigned and cannot be enabled.
	ServiceRestIncomingPlainTextListenPort int64 `json:"serviceRestIncomingPlainTextListenPort,omitempty"`
	// Indicates whether the incoming REST Service is operationally up in the Message VPN.
	ServiceRestIncomingPlainTextUp **bool `json:"serviceRestIncomingPlainTextUp,omitempty"`
	// Indicates whether the TLS related incoming REST Service is compressed in the Message VPN.
	ServiceRestIncomingTlsCompressed **bool `json:"serviceRestIncomingTlsCompressed,omitempty"`
	// Indicates whether encryption (TLS) is enabled for incoming REST clients in the Message VPN.
	ServiceRestIncomingTlsEnabled **bool `json:"serviceRestIncomingTlsEnabled,omitempty"`
	// The reason for the TLS related incoming REST Service failure in the Message VPN.
	ServiceRestIncomingTlsFailureReason string `json:"serviceRestIncomingTlsFailureReason,omitempty"`
	// The port number for incoming REST clients that connect to the Message VPN over TLS. The port must be unique across the message backbone. A value of 0 means that the listen-port is unassigned and cannot be enabled.
	ServiceRestIncomingTlsListenPort int64 `json:"serviceRestIncomingTlsListenPort,omitempty"`
	// Indicates whether the TLS related incoming REST Service is operationally up in the Message VPN.
	ServiceRestIncomingTlsUp **bool `json:"serviceRestIncomingTlsUp,omitempty"`
	// The REST service mode for incoming REST clients that connect to the Message VPN. The allowed values and their meaning are:  <pre> \"gateway\" - Act as a message gateway through which REST messages are propagated. \"messaging\" - Act as a message broker on which REST messages are queued. </pre> 
	ServiceRestMode string `json:"serviceRestMode,omitempty"`
	// The maximum number of REST Consumer (outgoing) client connections that can be simultaneously connected to the Message VPN.
	ServiceRestOutgoingMaxConnectionCount int64 `json:"serviceRestOutgoingMaxConnectionCount,omitempty"`
	// The maximum number of SMF client connections that can be simultaneously connected to the Message VPN. This value may be higher than supported by the platform.
	ServiceSmfMaxConnectionCount int64 `json:"serviceSmfMaxConnectionCount,omitempty"`
	// Indicates whether the SMF Service is enabled in the Message VPN.
	ServiceSmfPlainTextEnabled **bool `json:"serviceSmfPlainTextEnabled,omitempty"`
	// The reason for the SMF Service failure in the Message VPN.
	ServiceSmfPlainTextFailureReason string `json:"serviceSmfPlainTextFailureReason,omitempty"`
	// Indicates whether the SMF Service is operationally up in the Message VPN.
	ServiceSmfPlainTextUp **bool `json:"serviceSmfPlainTextUp,omitempty"`
	// Indicates whether encryption (TLS) is enabled for SMF clients in the Message VPN.
	ServiceSmfTlsEnabled **bool `json:"serviceSmfTlsEnabled,omitempty"`
	// The reason for the TLS related SMF Service failure in the Message VPN.
	ServiceSmfTlsFailureReason string `json:"serviceSmfTlsFailureReason,omitempty"`
	// Indicates whether the TLS related SMF Service is operationally up in the Message VPN.
	ServiceSmfTlsUp **bool `json:"serviceSmfTlsUp,omitempty"`
	// Determines when to request a client certificate from a Web Transport client connecting via a TLS port. The allowed values and their meaning are:  <pre> \"always\" - Always ask for a client certificate regardless of the \"message-vpn > authentication > client-certificate > shutdown\" configuration. \"never\" - Never ask for a client certificate regardless of the \"message-vpn > authentication > client-certificate > shutdown\" configuration. \"when-enabled-in-message-vpn\" - Only ask for a client-certificate if client certificate authentication is enabled under \"message-vpn >  authentication > client-certificate > shutdown\". </pre>  Available since 2.21.
	ServiceWebAuthenticationClientCertRequest string `json:"serviceWebAuthenticationClientCertRequest,omitempty"`
	// The maximum number of Web Transport client connections that can be simultaneously connected to the Message VPN. This value may be higher than supported by the platform.
	ServiceWebMaxConnectionCount int64 `json:"serviceWebMaxConnectionCount,omitempty"`
	// Indicates whether the Web transport for the SMF Service is enabled in the Message VPN.
	ServiceWebPlainTextEnabled **bool `json:"serviceWebPlainTextEnabled,omitempty"`
	// The reason for the Web transport related SMF Service failure in the Message VPN.
	ServiceWebPlainTextFailureReason string `json:"serviceWebPlainTextFailureReason,omitempty"`
	// Indicates whether the Web transport for the SMF Service is operationally up in the Message VPN.
	ServiceWebPlainTextUp **bool `json:"serviceWebPlainTextUp,omitempty"`
	// Indicates whether TLS is enabled for SMF clients in the Message VPN that use the Web transport.
	ServiceWebTlsEnabled **bool `json:"serviceWebTlsEnabled,omitempty"`
	// The reason for the TLS related Web transport SMF Service failure in the Message VPN.
	ServiceWebTlsFailureReason string `json:"serviceWebTlsFailureReason,omitempty"`
	// Indicates whether the TLS related Web transport SMF Service is operationally up in the Message VPN.
	ServiceWebTlsUp **bool `json:"serviceWebTlsUp,omitempty"`
	// The operational state of the local Message VPN. The allowed values and their meaning are:  <pre> \"up\" - The Message VPN is operationally up. \"down\" - The Message VPN is operationally down. \"standby\" - The Message VPN is operationally replication standby. </pre> 
	State string `json:"state,omitempty"`
	// The progress of the subscription export task, in percent complete.
	SubscriptionExportProgress int64 `json:"subscriptionExportProgress,omitempty"`
	// Indicates whether the Message VPN is the system manager for handling system level SEMP get requests and system level event publishing.
	SystemManager **bool `json:"systemManager,omitempty"`
	// Indicates whether SMF clients connected to the Message VPN are allowed to downgrade their connections from TLS to plain text.
	TlsAllowDowngradeToPlainTextEnabled **bool `json:"tlsAllowDowngradeToPlainTextEnabled,omitempty"`
	// The one minute average of the TLS message rate received by the Message VPN, in bytes per second (B/sec). Available since 2.13.
	TlsAverageRxByteRate int64 `json:"tlsAverageRxByteRate,omitempty"`
	// The one minute average of the TLS message rate transmitted by the Message VPN, in bytes per second (B/sec). Available since 2.13.
	TlsAverageTxByteRate int64 `json:"tlsAverageTxByteRate,omitempty"`
	// The amount of TLS messages received by the Message VPN, in bytes (B). Available since 2.13.
	TlsRxByteCount int64 `json:"tlsRxByteCount,omitempty"`
	// The current TLS message rate received by the Message VPN, in bytes per second (B/sec). Available since 2.13.
	TlsRxByteRate int64 `json:"tlsRxByteRate,omitempty"`
	// The amount of TLS messages transmitted by the Message VPN, in bytes (B). Available since 2.13.
	TlsTxByteCount int64 `json:"tlsTxByteCount,omitempty"`
	// The current TLS message rate transmitted by the Message VPN, in bytes per second (B/sec). Available since 2.13.
	TlsTxByteRate int64 `json:"tlsTxByteRate,omitempty"`
	// The amount of messages transmitted to clients by the Message VPN, in bytes (B). Available since 2.12.
	TxByteCount int64 `json:"txByteCount,omitempty"`
	// The current message rate transmitted by the Message VPN, in bytes per second (B/sec). Available since 2.13.
	TxByteRate int64 `json:"txByteRate,omitempty"`
	// The amount of compressed messages transmitted by the Message VPN, in bytes (B). Available since 2.12.
	TxCompressedByteCount int64 `json:"txCompressedByteCount,omitempty"`
	// The current compressed message rate transmitted by the Message VPN, in bytes per second (B/sec). Available since 2.12.
	TxCompressedByteRate int64 `json:"txCompressedByteRate,omitempty"`
	// The compression ratio for messages transmitted by the message VPN. Available since 2.12.
	TxCompressionRatio string `json:"txCompressionRatio,omitempty"`
	// The number of messages transmitted to clients by the Message VPN. Available since 2.12.
	TxMsgCount int64 `json:"txMsgCount,omitempty"`
	// The current message rate transmitted by the Message VPN, in messages per second (msg/sec). Available since 2.13.
	TxMsgRate int64 `json:"txMsgRate,omitempty"`
	// The amount of uncompressed messages transmitted by the Message VPN, in bytes (B). Available since 2.12.
	TxUncompressedByteCount int64 `json:"txUncompressedByteCount,omitempty"`
	// The current uncompressed message rate transmitted by the Message VPN, in bytes per second (B/sec). Available since 2.12.
	TxUncompressedByteRate int64 `json:"txUncompressedByteRate,omitempty"`
}
