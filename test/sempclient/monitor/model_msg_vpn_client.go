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

type MsgVpnClient struct {
	// The name of the access control list (ACL) profile of the Client.
	AclProfileName string `json:"aclProfileName,omitempty"`
	// The name of the original MsgVpn which the client signaled in. Available since 2.14.
	AliasedFromMsgVpnName string `json:"aliasedFromMsgVpnName,omitempty"`
	// The number of Client bind failures due to endpoint being already bound.
	AlreadyBoundBindFailureCount int64 `json:"alreadyBoundBindFailureCount,omitempty"`
	// The name of the authorization group of the Client.
	AuthorizationGroupName string `json:"authorizationGroupName,omitempty"`
	// The one minute average of the message rate received from the Client, in bytes per second (B/sec).
	AverageRxByteRate int64 `json:"averageRxByteRate,omitempty"`
	// The one minute average of the message rate received from the Client, in messages per second (msg/sec).
	AverageRxMsgRate int64 `json:"averageRxMsgRate,omitempty"`
	// The one minute average of the message rate transmitted to the Client, in bytes per second (B/sec).
	AverageTxByteRate int64 `json:"averageTxByteRate,omitempty"`
	// The one minute average of the message rate transmitted to the Client, in messages per second (msg/sec).
	AverageTxMsgRate int64 `json:"averageTxMsgRate,omitempty"`
	// The number of Client requests to bind to an endpoint.
	BindRequestCount int64 `json:"bindRequestCount,omitempty"`
	// The number of successful Client requests to bind to an endpoint.
	BindSuccessCount int64 `json:"bindSuccessCount,omitempty"`
	// The IP address and port of the Client.
	ClientAddress string `json:"clientAddress,omitempty"`
	// The identifier (ID) of the Client.
	ClientId int32 `json:"clientId,omitempty"`
	// The name of the Client.
	ClientName string `json:"clientName,omitempty"`
	// The name of the client profile of the Client.
	ClientProfileName string `json:"clientProfileName,omitempty"`
	// The client username of the Client used for authorization.
	ClientUsername string `json:"clientUsername,omitempty"`
	// The amount of client control messages received from the Client, in bytes (B).
	ControlRxByteCount int64 `json:"controlRxByteCount,omitempty"`
	// The number of client control messages received from the Client.
	ControlRxMsgCount int64 `json:"controlRxMsgCount,omitempty"`
	// The amount of client control messages transmitted to the Client, in bytes (B).
	ControlTxByteCount int64 `json:"controlTxByteCount,omitempty"`
	// The number of client control messages transmitted to the Client.
	ControlTxMsgCount int64 `json:"controlTxMsgCount,omitempty"`
	// The number of Client bind failures due to being denied cut-through forwarding.
	CutThroughDeniedBindFailureCount int64 `json:"cutThroughDeniedBindFailureCount,omitempty"`
	// The amount of client data messages received from the Client, in bytes (B).
	DataRxByteCount int64 `json:"dataRxByteCount,omitempty"`
	// The number of client data messages received from the Client.
	DataRxMsgCount int64 `json:"dataRxMsgCount,omitempty"`
	// The amount of client data messages transmitted to the Client, in bytes (B).
	DataTxByteCount int64 `json:"dataTxByteCount,omitempty"`
	// The number of client data messages transmitted to the Client.
	DataTxMsgCount int64 `json:"dataTxMsgCount,omitempty"`
	// The description text of the Client.
	Description string `json:"description,omitempty"`
	// The number of Client bind failures due to endpoint being disabled.
	DisabledBindFailureCount int64 `json:"disabledBindFailureCount,omitempty"`
	// The priority of the Client's subscriptions for receiving deliver-to-one (DTO) messages published on the local broker.
	DtoLocalPriority int32 `json:"dtoLocalPriority,omitempty"`
	// The priority of the Client's subscriptions for receiving deliver-to-one (DTO) messages published on a remote broker.
	DtoNetworkPriority int32 `json:"dtoNetworkPriority,omitempty"`
	// Indicates whether message eliding is enabled for the Client.
	Eliding **bool `json:"eliding,omitempty"`
	// The number of topics requiring message eliding for the Client.
	ElidingTopicCount int32 `json:"elidingTopicCount,omitempty"`
	// The peak number of topics requiring message eliding for the Client.
	ElidingTopicPeakCount int32 `json:"elidingTopicPeakCount,omitempty"`
	// The number of Client bind failures due to being denied guaranteed messaging.
	GuaranteedDeniedBindFailureCount int64 `json:"guaranteedDeniedBindFailureCount,omitempty"`
	// The number of Client bind failures due to an invalid selector.
	InvalidSelectorBindFailureCount int64 `json:"invalidSelectorBindFailureCount,omitempty"`
	// Indicates whether keepalive messages from the Client are received by the broker. Applicable for SMF and MQTT clients only. Available since 2.19.
	Keepalive **bool `json:"keepalive,omitempty"`
	// The maximum period of time the broker will accept inactivity from the Client before disconnecting, in seconds. Available since 2.19.
	KeepaliveEffectiveTimeout int32 `json:"keepaliveEffectiveTimeout,omitempty"`
	// Indicates whether the large-message event has been raised for the Client.
	LargeMsgEventRaised **bool `json:"largeMsgEventRaised,omitempty"`
	// The number of login request messages received from the Client.
	LoginRxMsgCount int64 `json:"loginRxMsgCount,omitempty"`
	// The number of login response messages transmitted to the Client.
	LoginTxMsgCount int64 `json:"loginTxMsgCount,omitempty"`
	// The number of Client bind failures due to the endpoint maximum bind count being exceeded.
	MaxBindCountExceededBindFailureCount int64 `json:"maxBindCountExceededBindFailureCount,omitempty"`
	// Indicates whether the max-eliding-topic-count event has been raised for the Client.
	MaxElidingTopicCountEventRaised **bool `json:"maxElidingTopicCountEventRaised,omitempty"`
	// The number of MQTT connect acknowledgment (CONNACK) refused response packets transmitted to the Client.
	MqttConnackErrorTxCount int64 `json:"mqttConnackErrorTxCount,omitempty"`
	// The number of MQTT connect acknowledgment (CONNACK) accepted response packets transmitted to the Client.
	MqttConnackTxCount int64 `json:"mqttConnackTxCount,omitempty"`
	// The number of MQTT connect (CONNECT) request packets received from the Client.
	MqttConnectRxCount int64 `json:"mqttConnectRxCount,omitempty"`
	// The number of MQTT disconnect (DISCONNECT) request packets received from the Client.
	MqttDisconnectRxCount int64 `json:"mqttDisconnectRxCount,omitempty"`
	// The number of MQTT ping request (PINGREQ) packets received from the Client.
	MqttPingreqRxCount int64 `json:"mqttPingreqRxCount,omitempty"`
	// The number of MQTT ping response (PINGRESP) packets transmitted to the Client.
	MqttPingrespTxCount int64 `json:"mqttPingrespTxCount,omitempty"`
	// The number of MQTT publish acknowledgement (PUBACK) response packets received from the Client.
	MqttPubackRxCount int64 `json:"mqttPubackRxCount,omitempty"`
	// The number of MQTT publish acknowledgement (PUBACK) response packets transmitted to the Client.
	MqttPubackTxCount int64 `json:"mqttPubackTxCount,omitempty"`
	// The number of MQTT publish complete (PUBCOMP) packets transmitted to the Client in response to a PUBREL packet. These packets are the fourth and final packet of a QoS 2 protocol exchange.
	MqttPubcompTxCount int64 `json:"mqttPubcompTxCount,omitempty"`
	// The number of MQTT publish message (PUBLISH) request packets received from the Client for QoS 0 message delivery.
	MqttPublishQos0RxCount int64 `json:"mqttPublishQos0RxCount,omitempty"`
	// The number of MQTT publish message (PUBLISH) request packets transmitted to the Client for QoS 0 message delivery.
	MqttPublishQos0TxCount int64 `json:"mqttPublishQos0TxCount,omitempty"`
	// The number of MQTT publish message (PUBLISH) request packets received from the Client for QoS 1 message delivery.
	MqttPublishQos1RxCount int64 `json:"mqttPublishQos1RxCount,omitempty"`
	// The number of MQTT publish message (PUBLISH) request packets transmitted to the Client for QoS 1 message delivery.
	MqttPublishQos1TxCount int64 `json:"mqttPublishQos1TxCount,omitempty"`
	// The number of MQTT publish message (PUBLISH) request packets received from the Client for QoS 2 message delivery.
	MqttPublishQos2RxCount int64 `json:"mqttPublishQos2RxCount,omitempty"`
	// The number of MQTT publish received (PUBREC) packets transmitted to the Client in response to a PUBLISH packet with QoS 2. These packets are the second packet of a QoS 2 protocol exchange.
	MqttPubrecTxCount int64 `json:"mqttPubrecTxCount,omitempty"`
	// The number of MQTT publish release (PUBREL) packets received from the Client in response to a PUBREC packet. These packets are the third packet of a QoS 2 protocol exchange.
	MqttPubrelRxCount int64 `json:"mqttPubrelRxCount,omitempty"`
	// The number of MQTT subscribe acknowledgement (SUBACK) failure response packets transmitted to the Client.
	MqttSubackErrorTxCount int64 `json:"mqttSubackErrorTxCount,omitempty"`
	// The number of MQTT subscribe acknowledgement (SUBACK) response packets transmitted to the Client.
	MqttSubackTxCount int64 `json:"mqttSubackTxCount,omitempty"`
	// The number of MQTT subscribe (SUBSCRIBE) request packets received from the Client to create one or more topic subscriptions.
	MqttSubscribeRxCount int64 `json:"mqttSubscribeRxCount,omitempty"`
	// The number of MQTT unsubscribe acknowledgement (UNSUBACK) response packets transmitted to the Client.
	MqttUnsubackTxCount int64 `json:"mqttUnsubackTxCount,omitempty"`
	// The number of MQTT unsubscribe (UNSUBSCRIBE) request packets received from the Client to remove one or more topic subscriptions.
	MqttUnsubscribeRxCount int64 `json:"mqttUnsubscribeRxCount,omitempty"`
	// The number of messages from the Client discarded due to message spool congestion primarily caused by message promotion.
	MsgSpoolCongestionRxDiscardedMsgCount int64 `json:"msgSpoolCongestionRxDiscardedMsgCount,omitempty"`
	// The number of messages from the Client discarded by the message spool.
	MsgSpoolRxDiscardedMsgCount int64 `json:"msgSpoolRxDiscardedMsgCount,omitempty"`
	// The name of the Message VPN.
	MsgVpnName string `json:"msgVpnName,omitempty"`
	// Indicates whether not to deliver messages to the Client if it published them.
	NoLocalDelivery **bool `json:"noLocalDelivery,omitempty"`
	// The number of messages from the Client discarded due to no matching subscription found.
	NoSubscriptionMatchRxDiscardedMsgCount int64 `json:"noSubscriptionMatchRxDiscardedMsgCount,omitempty"`
	// The original value of the client username used for Client authentication.
	OriginalClientUsername string `json:"originalClientUsername,omitempty"`
	// The number of Client bind failures due to other reasons.
	OtherBindFailureCount int64 `json:"otherBindFailureCount,omitempty"`
	// The platform the Client application software was built for, which may include the OS and API type.
	Platform string `json:"platform,omitempty"`
	// The number of messages from the Client discarded due to the publish topic being denied by the Access Control List (ACL) profile.
	PublishTopicAclRxDiscardedMsgCount int64 `json:"publishTopicAclRxDiscardedMsgCount,omitempty"`
	// The amount of HTTP request messages received from the Client, in bytes (B).
	RestHttpRequestRxByteCount int64 `json:"restHttpRequestRxByteCount,omitempty"`
	// The number of HTTP request messages received from the Client.
	RestHttpRequestRxMsgCount int64 `json:"restHttpRequestRxMsgCount,omitempty"`
	// The amount of HTTP request messages transmitted to the Client, in bytes (B).
	RestHttpRequestTxByteCount int64 `json:"restHttpRequestTxByteCount,omitempty"`
	// The number of HTTP request messages transmitted to the Client.
	RestHttpRequestTxMsgCount int64 `json:"restHttpRequestTxMsgCount,omitempty"`
	// The number of HTTP client/server error response messages received from the Client.
	RestHttpResponseErrorRxMsgCount int64 `json:"restHttpResponseErrorRxMsgCount,omitempty"`
	// The number of HTTP client/server error response messages transmitted to the Client.
	RestHttpResponseErrorTxMsgCount int64 `json:"restHttpResponseErrorTxMsgCount,omitempty"`
	// The amount of HTTP response messages received from the Client, in bytes (B).
	RestHttpResponseRxByteCount int64 `json:"restHttpResponseRxByteCount,omitempty"`
	// The number of HTTP response messages received from the Client.
	RestHttpResponseRxMsgCount int64 `json:"restHttpResponseRxMsgCount,omitempty"`
	// The number of HTTP successful response messages received from the Client.
	RestHttpResponseSuccessRxMsgCount int64 `json:"restHttpResponseSuccessRxMsgCount,omitempty"`
	// The number of HTTP successful response messages transmitted to the Client.
	RestHttpResponseSuccessTxMsgCount int64 `json:"restHttpResponseSuccessTxMsgCount,omitempty"`
	// The number of HTTP wait for reply timeout response messages received from the Client.
	RestHttpResponseTimeoutRxMsgCount int64 `json:"restHttpResponseTimeoutRxMsgCount,omitempty"`
	// The number of HTTP wait for reply timeout response messages transmitted to the Client.
	RestHttpResponseTimeoutTxMsgCount int64 `json:"restHttpResponseTimeoutTxMsgCount,omitempty"`
	// The amount of HTTP response messages transmitted to the Client, in bytes (B).
	RestHttpResponseTxByteCount int64 `json:"restHttpResponseTxByteCount,omitempty"`
	// The number of HTTP response messages transmitted to the Client.
	RestHttpResponseTxMsgCount int64 `json:"restHttpResponseTxMsgCount,omitempty"`
	// The amount of messages received from the Client, in bytes (B).
	RxByteCount int64 `json:"rxByteCount,omitempty"`
	// The current message rate received from the Client, in bytes per second (B/sec).
	RxByteRate int64 `json:"rxByteRate,omitempty"`
	// The number of messages discarded during reception from the Client.
	RxDiscardedMsgCount int64 `json:"rxDiscardedMsgCount,omitempty"`
	// The number of messages received from the Client.
	RxMsgCount int64 `json:"rxMsgCount,omitempty"`
	// The current message rate received from the Client, in messages per second (msg/sec).
	RxMsgRate int64 `json:"rxMsgRate,omitempty"`
	// The timestamp of when the Client will be disconnected by the broker. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). Available since 2.13.
	ScheduledDisconnectTime int32 `json:"scheduledDisconnectTime,omitempty"`
	// Indicates whether the Client is a slow subscriber and blocks for a few seconds when receiving messages.
	SlowSubscriber **bool `json:"slowSubscriber,omitempty"`
	// The date the Client application software was built.
	SoftwareDate string `json:"softwareDate,omitempty"`
	// The version of the Client application software.
	SoftwareVersion string `json:"softwareVersion,omitempty"`
	// The description of the TLS cipher used by the Client, which may include cipher name, key exchange and encryption algorithms.
	TlsCipherDescription string `json:"tlsCipherDescription,omitempty"`
	// Indicates whether the Client TLS connection was downgraded to plain-text to increase performance.
	TlsDowngradedToPlainText **bool `json:"tlsDowngradedToPlainText,omitempty"`
	// The version of TLS used by the Client.
	TlsVersion string `json:"tlsVersion,omitempty"`
	// The number of messages from the Client discarded due to an error while parsing the publish topic.
	TopicParseErrorRxDiscardedMsgCount int64 `json:"topicParseErrorRxDiscardedMsgCount,omitempty"`
	// The amount of messages transmitted to the Client, in bytes (B).
	TxByteCount int64 `json:"txByteCount,omitempty"`
	// The current message rate transmitted to the Client, in bytes per second (B/sec).
	TxByteRate int64 `json:"txByteRate,omitempty"`
	// The number of messages discarded during transmission to the Client.
	TxDiscardedMsgCount int64 `json:"txDiscardedMsgCount,omitempty"`
	// The number of messages transmitted to the Client.
	TxMsgCount int64 `json:"txMsgCount,omitempty"`
	// The current message rate transmitted to the Client, in messages per second (msg/sec).
	TxMsgRate int64 `json:"txMsgRate,omitempty"`
	// The amount of time in seconds since the Client connected.
	Uptime int32 `json:"uptime,omitempty"`
	// The user description for the Client, which may include computer name and process ID.
	User string `json:"user,omitempty"`
	// The virtual router used by the Client. The allowed values and their meaning are:  <pre> \"primary\" - The Client is using the primary virtual router. \"backup\" - The Client is using the backup virtual router. \"internal\" - The Client is using the internal virtual router. \"unknown\" - The Client virtual router is unknown. </pre> 
	VirtualRouter string `json:"virtualRouter,omitempty"`
	// The maximum web transport timeout for the Client being inactive, in seconds.
	WebInactiveTimeout int32 `json:"webInactiveTimeout,omitempty"`
	// The maximum web transport message payload size which excludes the size of the message header, in bytes.
	WebMaxPayload int64 `json:"webMaxPayload,omitempty"`
	// The number of messages from the Client discarded due to an error while parsing the web message.
	WebParseErrorRxDiscardedMsgCount int64 `json:"webParseErrorRxDiscardedMsgCount,omitempty"`
	// The remaining web transport timeout for the Client being inactive, in seconds.
	WebRemainingTimeout int32 `json:"webRemainingTimeout,omitempty"`
	// The amount of web transport messages received from the Client, in bytes (B).
	WebRxByteCount int64 `json:"webRxByteCount,omitempty"`
	// The type of encoding used during reception from the Client. The allowed values and their meaning are:  <pre> \"binary\" - The Client is using binary encoding. \"base64\" - The Client is using base64 encoding. \"illegal\" - The Client is using an illegal encoding type. </pre> 
	WebRxEncoding string `json:"webRxEncoding,omitempty"`
	// The number of web transport messages received from the Client.
	WebRxMsgCount int64 `json:"webRxMsgCount,omitempty"`
	// The type of web transport used during reception from the Client. The allowed values and their meaning are:  <pre> \"ws-binary\" - The Client is using WebSocket binary transport. \"http-binary-streaming\" - The Client is using HTTP binary streaming transport. \"http-binary\" - The Client is using HTTP binary transport. \"http-base64\" - The Client is using HTTP base64 transport. </pre> 
	WebRxProtocol string `json:"webRxProtocol,omitempty"`
	// The number of web transport requests received from the Client (HTTP only). Not available for WebSockets.
	WebRxRequestCount int64 `json:"webRxRequestCount,omitempty"`
	// The number of web transport responses transmitted to the Client on the receive connection (HTTP only). Not available for WebSockets.
	WebRxResponseCount int64 `json:"webRxResponseCount,omitempty"`
	// The TCP state of the receive connection from the Client. When fully operational, should be: established. See RFC 793 for further details. The allowed values and their meaning are:  <pre> \"closed\" - No connection state at all. \"listen\" - Waiting for a connection request from any remote TCP and port. \"syn-sent\" - Waiting for a matching connection request after having sent a connection request. \"syn-received\" - Waiting for a confirming connection request acknowledgment after having both received and sent a connection request. \"established\" - An open connection, data received can be delivered to the user. \"close-wait\" - Waiting for a connection termination request from the local user. \"fin-wait-1\" - Waiting for a connection termination request from the remote TCP, or an acknowledgment of the connection termination request previously sent. \"closing\" - Waiting for a connection termination request acknowledgment from the remote TCP. \"last-ack\" - Waiting for an acknowledgment of the connection termination request previously sent to the remote TCP. \"fin-wait-2\" - Waiting for a connection termination request from the remote TCP. \"time-wait\" - Waiting for enough time to pass to be sure the remote TCP received the acknowledgment of its connection termination request. </pre> 
	WebRxTcpState string `json:"webRxTcpState,omitempty"`
	// The description of the TLS cipher received from the Client, which may include cipher name, key exchange and encryption algorithms.
	WebRxTlsCipherDescription string `json:"webRxTlsCipherDescription,omitempty"`
	// The version of TLS used during reception from the Client.
	WebRxTlsVersion string `json:"webRxTlsVersion,omitempty"`
	// The identifier (ID) of the web transport session for the Client.
	WebSessionId string `json:"webSessionId,omitempty"`
	// The amount of web transport messages transmitted to the Client, in bytes (B).
	WebTxByteCount int64 `json:"webTxByteCount,omitempty"`
	// The type of encoding used during transmission to the Client. The allowed values and their meaning are:  <pre> \"binary\" - The Client is using binary encoding. \"base64\" - The Client is using base64 encoding. \"illegal\" - The Client is using an illegal encoding type. </pre> 
	WebTxEncoding string `json:"webTxEncoding,omitempty"`
	// The number of web transport messages transmitted to the Client.
	WebTxMsgCount int64 `json:"webTxMsgCount,omitempty"`
	// The type of web transport used during transmission to the Client. The allowed values and their meaning are:  <pre> \"ws-binary\" - The Client is using WebSocket binary transport. \"http-binary-streaming\" - The Client is using HTTP binary streaming transport. \"http-binary\" - The Client is using HTTP binary transport. \"http-base64\" - The Client is using HTTP base64 transport. </pre> 
	WebTxProtocol string `json:"webTxProtocol,omitempty"`
	// The number of web transport requests transmitted to the Client (HTTP only). Not available for WebSockets.
	WebTxRequestCount int64 `json:"webTxRequestCount,omitempty"`
	// The number of web transport responses received from the Client on the transmit connection (HTTP only). Not available for WebSockets.
	WebTxResponseCount int64 `json:"webTxResponseCount,omitempty"`
	// The TCP state of the transmit connection to the Client. When fully operational, should be: established. See RFC 793 for further details. The allowed values and their meaning are:  <pre> \"closed\" - No connection state at all. \"listen\" - Waiting for a connection request from any remote TCP and port. \"syn-sent\" - Waiting for a matching connection request after having sent a connection request. \"syn-received\" - Waiting for a confirming connection request acknowledgment after having both received and sent a connection request. \"established\" - An open connection, data received can be delivered to the user. \"close-wait\" - Waiting for a connection termination request from the local user. \"fin-wait-1\" - Waiting for a connection termination request from the remote TCP, or an acknowledgment of the connection termination request previously sent. \"closing\" - Waiting for a connection termination request acknowledgment from the remote TCP. \"last-ack\" - Waiting for an acknowledgment of the connection termination request previously sent to the remote TCP. \"fin-wait-2\" - Waiting for a connection termination request from the remote TCP. \"time-wait\" - Waiting for enough time to pass to be sure the remote TCP received the acknowledgment of its connection termination request. </pre> 
	WebTxTcpState string `json:"webTxTcpState,omitempty"`
	// The description of the TLS cipher transmitted to the Client, which may include cipher name, key exchange and encryption algorithms.
	WebTxTlsCipherDescription string `json:"webTxTlsCipherDescription,omitempty"`
	// The version of TLS used during transmission to the Client.
	WebTxTlsVersion string `json:"webTxTlsVersion,omitempty"`
}
