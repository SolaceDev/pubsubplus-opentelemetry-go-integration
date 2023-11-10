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

type MsgVpnTopicEndpoint struct {
	// The access type for delivering messages to consumer flows bound to the Topic Endpoint. The allowed values and their meaning are:  <pre> \"exclusive\" - Exclusive delivery of messages to the first bound consumer flow. \"non-exclusive\" - Non-exclusive delivery of messages to bound consumer flows in a round-robin (if partition count is zero) or partitioned (if partition count is non-zero) fashion. </pre> 
	AccessType string `json:"accessType,omitempty"`
	// The number of Topic Endpoint bind failures due to being already bound.
	AlreadyBoundBindFailureCount int64 `json:"alreadyBoundBindFailureCount,omitempty"`
	// The one minute average of the bind request rate received by the Topic Endpoint, in binds per second (binds/sec). Available since 2.25.
	AverageBindRequestRate int64 `json:"averageBindRequestRate,omitempty"`
	// The one minute average of the message rate received by the Topic Endpoint, in bytes per second (B/sec).
	AverageRxByteRate int64 `json:"averageRxByteRate,omitempty"`
	// The one minute average of the message rate received by the Topic Endpoint, in messages per second (msg/sec).
	AverageRxMsgRate int64 `json:"averageRxMsgRate,omitempty"`
	// The one minute average of the message rate transmitted by the Topic Endpoint, in bytes per second (B/sec).
	AverageTxByteRate int64 `json:"averageTxByteRate,omitempty"`
	// The one minute average of the message rate transmitted by the Topic Endpoint, in messages per second (msg/sec).
	AverageTxMsgRate int64 `json:"averageTxMsgRate,omitempty"`
	// The number of consumer requests to bind to the Topic Endpoint.
	BindRequestCount int64 `json:"bindRequestCount,omitempty"`
	// The current bind request rate received by the Topic Endpoint, in binds per second (binds/sec). Available since 2.25.
	BindRequestRate int64 `json:"bindRequestRate,omitempty"`
	// The number of successful consumer requests to bind to the Topic Endpoint.
	BindSuccessCount int64 `json:"bindSuccessCount,omitempty"`
	// The forwarding mode of the Topic Endpoint at bind time. The allowed values and their meaning are:  <pre> \"store-and-forward\" - Deliver messages using the guaranteed data path. \"cut-through\" - Deliver messages using the direct and guaranteed data paths for lower latency. </pre> 
	BindTimeForwardingMode string `json:"bindTimeForwardingMode,omitempty"`
	// The number of guaranteed messages discarded by the Topic Endpoint due to being denied by the Client Profile.
	ClientProfileDeniedDiscardedMsgCount int64 `json:"clientProfileDeniedDiscardedMsgCount,omitempty"`
	// Indicates whether the propagation of consumer acknowledgements (ACKs) received on the active replication Message VPN to the standby replication Message VPN is enabled.
	ConsumerAckPropagationEnabled **bool `json:"consumerAckPropagationEnabled,omitempty"`
	// Indicates whether the Topic Endpoint was created by a management API (CLI or SEMP).
	CreatedByManagement **bool `json:"createdByManagement,omitempty"`
	// The name of the Dead Message Queue (DMQ) used by the Topic Endpoint.
	DeadMsgQueue string `json:"deadMsgQueue,omitempty"`
	// The number of guaranteed messages deleted from the Topic Endpoint.
	DeletedMsgCount int64 `json:"deletedMsgCount,omitempty"`
	// Enable or disable the ability for client applications to query the message delivery count of messages received from the Topic Endpoint. This is a controlled availability feature. Please contact support to find out if this feature is supported for your use case. Available since 2.19.
	DeliveryCountEnabled **bool `json:"deliveryCountEnabled,omitempty"`
	// The delay, in seconds, to apply to messages arriving on the Topic Endpoint before the messages are eligible for delivery. Available since 2.22.
	DeliveryDelay int64 `json:"deliveryDelay,omitempty"`
	// The number of guaranteed messages discarded by the Topic Endpoint due to a destination group error.
	DestinationGroupErrorDiscardedMsgCount int64 `json:"destinationGroupErrorDiscardedMsgCount,omitempty"`
	// The destination topic of the Topic Endpoint.
	DestinationTopic string `json:"destinationTopic,omitempty"`
	// The number of Topic Endpoint bind failures due to being disabled.
	DisabledBindFailureCount int64 `json:"disabledBindFailureCount,omitempty"`
	// The number of guaranteed messages discarded by the Topic Endpoint due to it being disabled.
	DisabledDiscardedMsgCount int64 `json:"disabledDiscardedMsgCount,omitempty"`
	// Indicates whether the Topic Endpoint is durable and not temporary.
	Durable **bool `json:"durable,omitempty"`
	// Indicates whether the transmission of messages from the Topic Endpoint is enabled.
	EgressEnabled **bool `json:"egressEnabled,omitempty"`
	EventBindCountThreshold *EventThreshold `json:"eventBindCountThreshold,omitempty"`
	EventRejectLowPriorityMsgLimitThreshold *EventThreshold `json:"eventRejectLowPriorityMsgLimitThreshold,omitempty"`
	EventSpoolUsageThreshold *EventThreshold `json:"eventSpoolUsageThreshold,omitempty"`
	// The highest identifier (ID) of guaranteed messages in the Topic Endpoint that were acknowledged.
	HighestAckedMsgId int64 `json:"highestAckedMsgId,omitempty"`
	// The highest identifier (ID) of guaranteed messages in the Topic Endpoint.
	HighestMsgId int64 `json:"highestMsgId,omitempty"`
	// The number of acknowledgement messages received by the Topic Endpoint that are in the process of updating and deleting associated guaranteed messages.
	InProgressAckMsgCount int64 `json:"inProgressAckMsgCount,omitempty"`
	// Indicates whether the reception of messages to the Topic Endpoint is enabled.
	IngressEnabled **bool `json:"ingressEnabled,omitempty"`
	// The number of Topic Endpoint bind failures due to an invalid selector.
	InvalidSelectorBindFailureCount int64 `json:"invalidSelectorBindFailureCount,omitempty"`
	// The timestamp of the last completed replay for the Topic Endpoint. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time).
	LastReplayCompleteTime int32 `json:"lastReplayCompleteTime,omitempty"`
	// The reason for the last replay failure for the Topic Endpoint.
	LastReplayFailureReason string `json:"lastReplayFailureReason,omitempty"`
	// The timestamp of the last replay failure for the Topic Endpoint. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time).
	LastReplayFailureTime int32 `json:"lastReplayFailureTime,omitempty"`
	// The timestamp of the last replay started for the Topic Endpoint. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time).
	LastReplayStartTime int32 `json:"lastReplayStartTime,omitempty"`
	// The timestamp of the last replayed message transmitted by the Topic Endpoint. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time).
	LastReplayedMsgTxTime int32 `json:"lastReplayedMsgTxTime,omitempty"`
	// The identifier (ID) of the last message examined by the Topic Endpoint selector.
	LastSelectorExaminedMsgId int64 `json:"lastSelectorExaminedMsgId,omitempty"`
	// The identifier (ID) of the last guaranteed message spooled in the Topic Endpoint.
	LastSpooledMsgId int64 `json:"lastSpooledMsgId,omitempty"`
	// The number of guaranteed messages discarded by the Topic Endpoint due to low priority message congestion control.
	LowPriorityMsgCongestionDiscardedMsgCount int64 `json:"lowPriorityMsgCongestionDiscardedMsgCount,omitempty"`
	// The state of the low priority message congestion in the Topic Endpoint. The allowed values and their meaning are:  <pre> \"disabled\" - Messages are not being checked for priority. \"not-congested\" - Low priority messages are being stored and delivered. \"congested\" - Low priority messages are being discarded. </pre> 
	LowPriorityMsgCongestionState string `json:"lowPriorityMsgCongestionState,omitempty"`
	// The lowest identifier (ID) of guaranteed messages in the Topic Endpoint that were acknowledged.
	LowestAckedMsgId int64 `json:"lowestAckedMsgId,omitempty"`
	// The lowest identifier (ID) of guaranteed messages in the Topic Endpoint.
	LowestMsgId int64 `json:"lowestMsgId,omitempty"`
	// The maximum number of consumer flows that can bind to the Topic Endpoint.
	MaxBindCount int64 `json:"maxBindCount,omitempty"`
	// The number of Topic Endpoint bind failures due to the maximum bind count being exceeded.
	MaxBindCountExceededBindFailureCount int64 `json:"maxBindCountExceededBindFailureCount,omitempty"`
	// The maximum number of messages delivered but not acknowledged per flow for the Topic Endpoint.
	MaxDeliveredUnackedMsgsPerFlow int64 `json:"maxDeliveredUnackedMsgsPerFlow,omitempty"`
	// The effective maximum number of consumer flows that can bind to the Topic Endpoint.
	MaxEffectiveBindCount int32 `json:"maxEffectiveBindCount,omitempty"`
	// The maximum message size allowed in the Topic Endpoint, in bytes (B).
	MaxMsgSize int32 `json:"maxMsgSize,omitempty"`
	// The number of guaranteed messages discarded by the Topic Endpoint due to the maximum message size being exceeded.
	MaxMsgSizeExceededDiscardedMsgCount int64 `json:"maxMsgSizeExceededDiscardedMsgCount,omitempty"`
	// The number of guaranteed messages discarded by the Topic Endpoint due to the maximum message spool usage being exceeded.
	MaxMsgSpoolUsageExceededDiscardedMsgCount int64 `json:"maxMsgSpoolUsageExceededDiscardedMsgCount,omitempty"`
	// The maximum number of times the Topic Endpoint will attempt redelivery of a message prior to it being discarded or moved to the DMQ. A value of 0 means to retry forever.
	MaxRedeliveryCount int64 `json:"maxRedeliveryCount,omitempty"`
	// The number of guaranteed messages discarded by the Topic Endpoint due to the maximum redelivery attempts being exceeded.
	MaxRedeliveryExceededDiscardedMsgCount int64 `json:"maxRedeliveryExceededDiscardedMsgCount,omitempty"`
	// The number of guaranteed messages discarded by the Topic Endpoint due to the maximum redelivery attempts being exceeded and failing to move to the Dead Message Queue (DMQ).
	MaxRedeliveryExceededToDmqFailedMsgCount int64 `json:"maxRedeliveryExceededToDmqFailedMsgCount,omitempty"`
	// The number of guaranteed messages moved to the Dead Message Queue (DMQ) by the Topic Endpoint due to the maximum redelivery attempts being exceeded.
	MaxRedeliveryExceededToDmqMsgCount int64 `json:"maxRedeliveryExceededToDmqMsgCount,omitempty"`
	// The maximum message spool usage allowed by the Topic Endpoint, in megabytes (MB). A value of 0 only allows spooling of the last message received and disables quota checking.
	MaxSpoolUsage int64 `json:"maxSpoolUsage,omitempty"`
	// The maximum time in seconds a message can stay in the Topic Endpoint when `respectTtlEnabled` is `\"true\"`. A message expires when the lesser of the sender assigned time-to-live (TTL) in the message and the `maxTtl` configured for the Topic Endpoint, is exceeded. A value of 0 disables expiry.
	MaxTtl int64 `json:"maxTtl,omitempty"`
	// The number of guaranteed messages discarded by the Topic Endpoint due to the maximum time-to-live (TTL) in hops being exceeded. The TTL hop count is incremented when the message crosses a bridge.
	MaxTtlExceededDiscardedMsgCount int64 `json:"maxTtlExceededDiscardedMsgCount,omitempty"`
	// The number of guaranteed messages discarded by the Topic Endpoint due to the maximum time-to-live (TTL) timestamp expiring.
	MaxTtlExpiredDiscardedMsgCount int64 `json:"maxTtlExpiredDiscardedMsgCount,omitempty"`
	// The number of guaranteed messages discarded by the Topic Endpoint due to the maximum time-to-live (TTL) timestamp expiring and failing to move to the Dead Message Queue (DMQ).
	MaxTtlExpiredToDmqFailedMsgCount int64 `json:"maxTtlExpiredToDmqFailedMsgCount,omitempty"`
	// The number of guaranteed messages moved to the Dead Message Queue (DMQ) by the Topic Endpoint due to the maximum time-to-live (TTL) timestamp expiring.
	MaxTtlExpiredToDmqMsgCount int64 `json:"maxTtlExpiredToDmqMsgCount,omitempty"`
	// The message spool peak usage by the Topic Endpoint, in bytes (B).
	MsgSpoolPeakUsage int64 `json:"msgSpoolPeakUsage,omitempty"`
	// The message spool usage by the Topic Endpoint, in bytes (B).
	MsgSpoolUsage int64 `json:"msgSpoolUsage,omitempty"`
	// The name of the Message VPN.
	MsgVpnName string `json:"msgVpnName,omitempty"`
	// The name of the network topic for the Topic Endpoint.
	NetworkTopic string `json:"networkTopic,omitempty"`
	// The number of guaranteed messages discarded by the Topic Endpoint due to no local delivery being requested.
	NoLocalDeliveryDiscardedMsgCount int64 `json:"noLocalDeliveryDiscardedMsgCount,omitempty"`
	// The number of Topic Endpoint bind failures due to other reasons.
	OtherBindFailureCount int64 `json:"otherBindFailureCount,omitempty"`
	// The Client Username that owns the Topic Endpoint and has permission equivalent to `\"delete\"`.
	Owner string `json:"owner,omitempty"`
	// The permission level for all consumers of the Topic Endpoint, excluding the owner. The allowed values and their meaning are:  <pre> \"no-access\" - Disallows all access. \"read-only\" - Read-only access to the messages. \"consume\" - Consume (read and remove) messages. \"modify-topic\" - Consume messages or modify the topic/selector. \"delete\" - Consume messages, modify the topic/selector or delete the Client created endpoint altogether. </pre> 
	Permission string `json:"permission,omitempty"`
	// The number of guaranteed messages transmitted by the Topic Endpoint for redelivery.
	RedeliveredMsgCount int64 `json:"redeliveredMsgCount,omitempty"`
	// Enable or disable a message redelivery delay. When false, messages are redelivered as-soon-as-possible.  When true, messages are redelivered according to the initial, max and multiplier.  This should only be enabled when redelivery is enabled. Available since 2.33.
	RedeliveryDelayEnabled **bool `json:"redeliveryDelayEnabled,omitempty"`
	// The delay to be used between the first 2 redelivery attempts.  This value is in milliseconds. Available since 2.33.
	RedeliveryDelayInitialInterval int32 `json:"redeliveryDelayInitialInterval,omitempty"`
	// The maximum delay to be used between any 2 redelivery attempts.  This value is in milliseconds.  Due to technical limitations, some redelivery attempt delays may slightly exceed this value. Available since 2.33.
	RedeliveryDelayMaxInterval int32 `json:"redeliveryDelayMaxInterval,omitempty"`
	// The amount each delay interval is multiplied by after each failed delivery attempt.  This number is in a fixed-point decimal format in which you must divide by 100 to get the floating point value. For example, a value of 125 would cause the delay to be multiplied by 1.25. Available since 2.33.
	RedeliveryDelayMultiplier int32 `json:"redeliveryDelayMultiplier,omitempty"`
	// Enable or disable message redelivery. When enabled, the number of redelivery attempts is controlled by maxRedeliveryCount. When disabled, the message will never be delivered from the topic-endpoint more than once. Available since 2.18.
	RedeliveryEnabled **bool `json:"redeliveryEnabled,omitempty"`
	// Indicates whether the checking of low priority messages against the `rejectLowPriorityMsgLimit` is enabled.
	RejectLowPriorityMsgEnabled **bool `json:"rejectLowPriorityMsgEnabled,omitempty"`
	// The number of messages of any priority in the Topic Endpoint above which low priority messages are not admitted but higher priority messages are allowed.
	RejectLowPriorityMsgLimit int64 `json:"rejectLowPriorityMsgLimit,omitempty"`
	// Determines when to return negative acknowledgements (NACKs) to sending clients on message discards. Note that NACKs cause the message to not be delivered to any destination and Transacted Session commits to fail. The allowed values and their meaning are:  <pre> \"always\" - Always return a negative acknowledgment (NACK) to the sending client on message discard. \"when-topic-endpoint-enabled\" - Only return a negative acknowledgment (NACK) to the sending client on message discard when the Topic Endpoint is enabled. \"never\" - Never return a negative acknowledgment (NACK) to the sending client on message discard. </pre> 
	RejectMsgToSenderOnDiscardBehavior string `json:"rejectMsgToSenderOnDiscardBehavior,omitempty"`
	// The number of replays that failed for the Topic Endpoint.
	ReplayFailureCount int64 `json:"replayFailureCount,omitempty"`
	// The number of replays started for the Topic Endpoint.
	ReplayStartCount int64 `json:"replayStartCount,omitempty"`
	// The state of replay for the Topic Endpoint. The allowed values and their meaning are:  <pre> \"initializing\" - All messages are being deleted from the endpoint before replay starts. \"active\" - Subscription matching logged messages are being replayed to the endpoint. \"pending-complete\" - Replay is complete, but final accounting is in progress. \"complete\" - Replay and all related activities are complete. \"failed\" - Replay has failed and is waiting for an unbind response. </pre> 
	ReplayState string `json:"replayState,omitempty"`
	// The number of replays that succeeded for the Topic Endpoint.
	ReplaySuccessCount int64 `json:"replaySuccessCount,omitempty"`
	// The number of replayed messages transmitted by the Topic Endpoint and acked by all consumers.
	ReplayedAckedMsgCount int64 `json:"replayedAckedMsgCount,omitempty"`
	// The number of replayed messages transmitted by the Topic Endpoint.
	ReplayedTxMsgCount int64 `json:"replayedTxMsgCount,omitempty"`
	// The number of acknowledgement messages propagated by the Topic Endpoint to the replication standby remote Message VPN.
	ReplicationActiveAckPropTxMsgCount int64 `json:"replicationActiveAckPropTxMsgCount,omitempty"`
	// The number of propagated acknowledgement messages received by the Topic Endpoint from the replication active remote Message VPN.
	ReplicationStandbyAckPropRxMsgCount int64 `json:"replicationStandbyAckPropRxMsgCount,omitempty"`
	// The number of messages acknowledged in the Topic Endpoint by acknowledgement propagation from the replication active remote Message VPN.
	ReplicationStandbyAckedByAckPropMsgCount int64 `json:"replicationStandbyAckedByAckPropMsgCount,omitempty"`
	// The number of messages received by the Topic Endpoint from the replication active remote Message VPN.
	ReplicationStandbyRxMsgCount int64 `json:"replicationStandbyRxMsgCount,omitempty"`
	// Indicates whether message priorities are respected. When enabled, messages contained in the Topic Endpoint are delivered in priority order, from 9 (highest) to 0 (lowest).
	RespectMsgPriorityEnabled **bool `json:"respectMsgPriorityEnabled,omitempty"`
	// Indicates whether the time-to-live (TTL) for messages in the Topic Endpoint is respected. When enabled, expired messages are discarded or moved to the DMQ.
	RespectTtlEnabled **bool `json:"respectTtlEnabled,omitempty"`
	// The current message rate received by the Topic Endpoint, in bytes per second (B/sec).
	RxByteRate int32 `json:"rxByteRate,omitempty"`
	// The current message rate received by the Topic Endpoint, in messages per second (msg/sec).
	RxMsgRate int64 `json:"rxMsgRate,omitempty"`
	// Indicates whether the Topic Endpoint has a selector to filter received messages.
	RxSelector **bool `json:"rxSelector,omitempty"`
	// The value of the receive selector for the Topic Endpoint.
	Selector string `json:"selector,omitempty"`
	// The number of guaranteed messages examined by the Topic Endpoint selector.
	SelectorExaminedMsgCount int64 `json:"selectorExaminedMsgCount,omitempty"`
	// The number of guaranteed messages for which the Topic Endpoint selector matched.
	SelectorMatchedMsgCount int64 `json:"selectorMatchedMsgCount,omitempty"`
	// The number of guaranteed messages for which the Topic Endpoint selector did not match.
	SelectorNotMatchedMsgCount int64 `json:"selectorNotMatchedMsgCount,omitempty"`
	// The amount of guaranteed messages that were spooled in the Topic Endpoint, in bytes (B).
	SpooledByteCount int64 `json:"spooledByteCount,omitempty"`
	// The number of guaranteed messages that were spooled in the Topic Endpoint.
	SpooledMsgCount int64 `json:"spooledMsgCount,omitempty"`
	// The name of the Topic Endpoint.
	TopicEndpointName string `json:"topicEndpointName,omitempty"`
	// The number of guaranteed messages that were retransmitted by the Topic Endpoint at the transport layer as part of a single delivery attempt. Available since 2.18.
	TransportRetransmitMsgCount int64 `json:"transportRetransmitMsgCount,omitempty"`
	// The current message rate transmitted by the Topic Endpoint, in bytes per second (B/sec).
	TxByteRate int64 `json:"txByteRate,omitempty"`
	// The current message rate transmitted by the Topic Endpoint, in messages per second (msg/sec).
	TxMsgRate int64 `json:"txMsgRate,omitempty"`
	// The number of guaranteed messages in the Topic Endpoint that have been transmitted but not acknowledged by all consumers.
	TxUnackedMsgCount int64 `json:"txUnackedMsgCount,omitempty"`
	// The virtual router used by the Topic Endpoint. The allowed values and their meaning are:  <pre> \"primary\" - The endpoint belongs to the primary virtual router. \"backup\" - The endpoint belongs to the backup virtual router. </pre>  Deprecated since 2.31. This attribute has been deprecated. When Guaranteed Messaging is active, this value is always the virtual router for which Guaranteed Messaging is enabled. Otherwise, this value should be ignored.
	VirtualRouter string `json:"virtualRouter,omitempty"`
}
