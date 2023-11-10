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

type MsgVpnAuthenticationOauthProfile struct {
	// Indicates whether the profile is active. An enabled profile may not be active if discovery is not complete, if there is no issuer specified, or if there is another profile with the same issuer. Available since 2.26.
	Active **bool `json:"active,omitempty"`
	// The name of the groups claim. If non-empty, the specified claim will be used to determine groups for authorization. If empty, the authorizationType attribute of the Message VPN will be used to determine authorization.
	AuthorizationGroupsClaimName string `json:"authorizationGroupsClaimName,omitempty"`
	// The format of the authorization groups claim value when it is a string. The allowed values and their meaning are:  <pre> \"single\" - When the claim is a string, it is interpreted as a single group. \"space-delimited\" - When the claim is a string, it is interpreted as a space-delimited list of groups, similar to the \"scope\" claim. </pre>  Available since 2.32.
	AuthorizationGroupsClaimStringFormat string `json:"authorizationGroupsClaimStringFormat,omitempty"`
	// The OAuth client id.
	ClientId string `json:"clientId,omitempty"`
	// The required value for the TYP field in the ID token header.
	ClientRequiredType string `json:"clientRequiredType,omitempty"`
	// Enable or disable verification of the TYP field in the ID token header.
	ClientValidateTypeEnabled **bool `json:"clientValidateTypeEnabled,omitempty"`
	// Enable or disable the disconnection of clients when their tokens expire. Changing this value does not affect existing clients, only new client connections.
	DisconnectOnTokenExpirationEnabled **bool `json:"disconnectOnTokenExpirationEnabled,omitempty"`
	// The reason for the last discovery endpoint refresh failure.
	DiscoveryLastRefreshFailureReason string `json:"discoveryLastRefreshFailureReason,omitempty"`
	// The timestamp of the last discovery endpoint refresh failure. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time).
	DiscoveryLastRefreshFailureTime int32 `json:"discoveryLastRefreshFailureTime,omitempty"`
	// The timestamp of the last discovery endpoint refresh success. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time).
	DiscoveryLastRefreshTime int32 `json:"discoveryLastRefreshTime,omitempty"`
	// The timestamp of the next scheduled discovery endpoint refresh. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time).
	DiscoveryNextScheduledRefreshTime int32 `json:"discoveryNextScheduledRefreshTime,omitempty"`
	// The number of discovery endpoint refresh failures.
	DiscoveryRefreshFailureCount int64 `json:"discoveryRefreshFailureCount,omitempty"`
	// Enable or disable the OAuth profile.
	Enabled **bool `json:"enabled,omitempty"`
	// The OpenID Connect discovery endpoint or OAuth Authorization Server Metadata endpoint.
	EndpointDiscovery string `json:"endpointDiscovery,omitempty"`
	// The number of seconds between discovery endpoint requests.
	EndpointDiscoveryRefreshInterval int32 `json:"endpointDiscoveryRefreshInterval,omitempty"`
	// The OAuth introspection endpoint.
	EndpointIntrospection string `json:"endpointIntrospection,omitempty"`
	// The operational OAuth introspection endpoint.
	EndpointIntrospectionOperational string `json:"endpointIntrospectionOperational,omitempty"`
	// The maximum time in seconds a token introspection request is allowed to take.
	EndpointIntrospectionTimeout int32 `json:"endpointIntrospectionTimeout,omitempty"`
	// The OAuth JWKS endpoint.
	EndpointJwks string `json:"endpointJwks,omitempty"`
	// The operational OAuth JWKS endpoint.
	EndpointJwksOperational string `json:"endpointJwksOperational,omitempty"`
	// The number of seconds between JWKS endpoint requests.
	EndpointJwksRefreshInterval int32 `json:"endpointJwksRefreshInterval,omitempty"`
	// The OpenID Connect Userinfo endpoint.
	EndpointUserinfo string `json:"endpointUserinfo,omitempty"`
	// The operational OpenID Connect Userinfo endpoint.
	EndpointUserinfoOperational string `json:"endpointUserinfoOperational,omitempty"`
	// The maximum time in seconds a userinfo request is allowed to take.
	EndpointUserinfoTimeout int32 `json:"endpointUserinfoTimeout,omitempty"`
	// The number of requests with an expired OAuth token.
	ExpiredTokenCount int64 `json:"expiredTokenCount,omitempty"`
	// The number of times the groups were successfully found in the ID token or access token for request authentication.
	GroupsFoundInTokenCount int64 `json:"groupsFoundInTokenCount,omitempty"`
	// The reason the profile is not active. The allowed values and their meaning are:  <pre> \"msg-vpn-disabled\" - The Message VPN is disabled. \"oauth-disabled\" - OAuth is disabled for the Message VPN. \"profile-disabled\" - The OAuth profile is disabled. \"missing-issuer\" - The issuer has not been discovered or configured. \"duplicate-issuer\" - Another OAuth profile in the Message VPN already has the same issuer. \"none\" - The OAuth profile is active. </pre>  Available since 2.26.
	InactiveReason string `json:"inactiveReason,omitempty"`
	// The one minute average of the time required to complete a token introspection, in milliseconds (ms).
	IntrospectionAverageTime int32 `json:"introspectionAverageTime,omitempty"`
	// The reason for the introspection endpoint request failure.
	IntrospectionLastFailureReason string `json:"introspectionLastFailureReason,omitempty"`
	// The timestamp of the last introspection endpoint request failure. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time).
	IntrospectionLastFailureTime int32 `json:"introspectionLastFailureTime,omitempty"`
	// The number of failures during request authentication due to missing introspection configuration (a introspection request was required but no introspection endpoint was configured).
	IntrospectionMissingCount int64 `json:"introspectionMissingCount,omitempty"`
	// The number of introspection request made from the broker during request authentication for this OAuth profile where the configured groups claim wasn't found in the access token or the introspection response.
	IntrospectionMissingGroupsCount int64 `json:"introspectionMissingGroupsCount,omitempty"`
	// The number of introspection requests made from the broker during request authentication for this OAuth profile where the configured username claim wasn't found in the access token or introspection response.
	IntrospectionMissingUsernameCount int64 `json:"introspectionMissingUsernameCount,omitempty"`
	// The number of requests made to the introspection endpoint during request authentication.
	IntrospectionRequestCount int64 `json:"introspectionRequestCount,omitempty"`
	// The number of introspection responses during request authentication that couldn't be parsed.
	IntrospectionResponseInvalidCount int64 `json:"introspectionResponseInvalidCount,omitempty"`
	// The number of introspection requests made from the broker during request authentication for this OAuth profile with 200 status responses.
	IntrospectionStatusOkCount int64 `json:"introspectionStatusOkCount,omitempty"`
	// The number of introspection requests made from the broker during request authentication for this OAuth profile with status responses other than 200.
	IntrospectionStatusOtherCount int64 `json:"introspectionStatusOtherCount,omitempty"`
	// The number of introspection responses indicating that the provided token was not active.
	IntrospectionTokenNotActiveCount int64 `json:"introspectionTokenNotActiveCount,omitempty"`
	// The number of requests with an invalid OAuth token.
	InvalidTokenCount int64 `json:"invalidTokenCount,omitempty"`
	// The Issuer Identifier for the OAuth provider.
	Issuer string `json:"issuer,omitempty"`
	// The operational Issuer Identifier for the OAuth provider.
	IssuerOperational string `json:"issuerOperational,omitempty"`
	// The reason for the last JWKS public key refresh failure.
	JwksLastRefreshFailureReason string `json:"jwksLastRefreshFailureReason,omitempty"`
	// The timestamp of the last JWKS public key refresh failure. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time).
	JwksLastRefreshFailureTime int32 `json:"jwksLastRefreshFailureTime,omitempty"`
	// The timestamp of the last JWKS public key refresh success. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time).
	JwksLastRefreshTime int32 `json:"jwksLastRefreshTime,omitempty"`
	// The timestamp of the next scheduled JWKS public key refresh. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time).
	JwksNextScheduledRefreshTime int32 `json:"jwksNextScheduledRefreshTime,omitempty"`
	// The number of JWKS public key refresh failures.
	JwksRefreshFailureCount int64 `json:"jwksRefreshFailureCount,omitempty"`
	// Enable or disable whether the API provided MQTT client username will be validated against the username calculated from the token(s). When enabled, connection attempts by MQTT clients are rejected if they differ. Note that this value only applies to MQTT clients; SMF client usernames will not be validated.
	MqttUsernameValidateEnabled **bool `json:"mqttUsernameValidateEnabled,omitempty"`
	// The name of the Message VPN.
	MsgVpnName string `json:"msgVpnName,omitempty"`
	// The name of the OAuth profile.
	OauthProfileName string `json:"oauthProfileName,omitempty"`
	// The OAuth role of the broker. The allowed values and their meaning are:  <pre> \"client\" - The broker is in the OAuth client role. \"resource-server\" - The broker is in the OAuth resource server role. </pre> 
	OauthRole string `json:"oauthRole,omitempty"`
	// The number of requests (successful and unsuccessful) using this OAuth profile.
	RequestCount int64 `json:"requestCount,omitempty"`
	// Enable or disable parsing of the access token as a JWT.
	ResourceServerParseAccessTokenEnabled **bool `json:"resourceServerParseAccessTokenEnabled,omitempty"`
	// The required audience value.
	ResourceServerRequiredAudience string `json:"resourceServerRequiredAudience,omitempty"`
	// The required issuer value.
	ResourceServerRequiredIssuer string `json:"resourceServerRequiredIssuer,omitempty"`
	// A space-separated list of scopes that must be present in the scope claim.
	ResourceServerRequiredScope string `json:"resourceServerRequiredScope,omitempty"`
	// The required TYP value.
	ResourceServerRequiredType string `json:"resourceServerRequiredType,omitempty"`
	// Enable or disable verification of the audience claim in the access token or introspection response.
	ResourceServerValidateAudienceEnabled **bool `json:"resourceServerValidateAudienceEnabled,omitempty"`
	// Enable or disable verification of the issuer claim in the access token or introspection response.
	ResourceServerValidateIssuerEnabled **bool `json:"resourceServerValidateIssuerEnabled,omitempty"`
	// Enable or disable verification of the scope claim in the access token or introspection response.
	ResourceServerValidateScopeEnabled **bool `json:"resourceServerValidateScopeEnabled,omitempty"`
	// Enable or disable verification of the TYP field in the access token header.
	ResourceServerValidateTypeEnabled **bool `json:"resourceServerValidateTypeEnabled,omitempty"`
	// The number of successful authentications using this OAuth profile.
	SuccessCount int64 `json:"successCount,omitempty"`
	// The one minute average of the time required to complete a token userinfo request, in milliseconds (ms).
	UserinfoAverageTime int32 `json:"userinfoAverageTime,omitempty"`
	// The reason for the userinfo endpoint request failure.
	UserinfoLastFailureReason string `json:"userinfoLastFailureReason,omitempty"`
	// The timestamp of the last userinfo endpoint request failure. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time).
	UserinfoLastFailureTime int32 `json:"userinfoLastFailureTime,omitempty"`
	// The number of failures due to missing Userinfo configuration (a Userinfo request was required but no Userinfo endpoint was configured) during request authentication.
	UserinfoMissingCount int64 `json:"userinfoMissingCount,omitempty"`
	// The number of Userinfo request made from the broker during request authentication for this OAuth profile where the configured groups claim wasn't found in the ID token or the Userinfo response.
	UserinfoMissingGroupsCount int64 `json:"userinfoMissingGroupsCount,omitempty"`
	// The number of Userinfo requests made from the broker during request authentication for this OAuth profile where the configured username claim wasn't found in the ID token or Userinfo response.
	UserinfoMissingUsernameCount int64 `json:"userinfoMissingUsernameCount,omitempty"`
	// The number of requests made to the Userinfo endpoint during request authentication.
	UserinfoRequestCount int64 `json:"userinfoRequestCount,omitempty"`
	// The number of Userinfo requests made from the broker during request authentication for this OAuth profile with responses that couldn't be parsed.
	UserinfoResponseInvalidCount int64 `json:"userinfoResponseInvalidCount,omitempty"`
	// The number of Userinfo requests made from the broker during request authentication for this OAuth profile with 200 status responses.
	UserinfoStatusOkCount int64 `json:"userinfoStatusOkCount,omitempty"`
	// The number of Userinfo requests made from the broker during request authentication for this OAuth profile with status responses other than 200.
	UserinfoStatusOtherCount int64 `json:"userinfoStatusOtherCount,omitempty"`
	// The number of Userinfo requests made from the broker during request authentication for this OAuth profile with subject claims that did not match the subject from the ID token.
	UserinfoSubjectMismatchCount int64 `json:"userinfoSubjectMismatchCount,omitempty"`
	// The name of the username claim.
	UsernameClaimName string `json:"usernameClaimName,omitempty"`
	// The number of time the username was successfully found in the ID token or access token for request authentication.
	UsernameFoundInTokenCount int64 `json:"usernameFoundInTokenCount,omitempty"`
}
