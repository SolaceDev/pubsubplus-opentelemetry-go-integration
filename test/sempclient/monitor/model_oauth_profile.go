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

type OauthProfile struct {
	// The name of the groups claim.
	AccessLevelGroupsClaimName string `json:"accessLevelGroupsClaimName,omitempty"`
	// The format of the access level groups claim value when it is a string. The allowed values and their meaning are:  <pre> \"single\" - When the claim is a string, it is interpreted as a single group. \"space-delimited\" - When the claim is a string, it is interpreted as a space-delimited list of groups, similar to the \"scope\" claim. </pre>  Available since 2.32.
	AccessLevelGroupsClaimStringFormat string `json:"accessLevelGroupsClaimStringFormat,omitempty"`
	// Indicates whether the profile is active. An enabled profile may not be active if discovery is not complete, if there is no issuer specified, or if there is another profile with the same issuer. Available since 2.26.
	Active **bool `json:"active,omitempty"`
	// The OAuth client id.
	ClientId string `json:"clientId,omitempty"`
	// The OAuth redirect URI.
	ClientRedirectUri string `json:"clientRedirectUri,omitempty"`
	// The required value for the TYP field in the ID token header.
	ClientRequiredType string `json:"clientRequiredType,omitempty"`
	// The OAuth scope.
	ClientScope string `json:"clientScope,omitempty"`
	// Enable or disable verification of the TYP field in the ID token header.
	ClientValidateTypeEnabled **bool `json:"clientValidateTypeEnabled,omitempty"`
	// The number of requests to the broker OAuth completion endpoint with an expired state token.
	CompleteExpiredStateCount int64 `json:"completeExpiredStateCount,omitempty"`
	// The number of request to the broker OAuth completion endpoint with an invalid state token.
	CompleteInvalidStateCount int64 `json:"completeInvalidStateCount,omitempty"`
	// The number of requests to the broker OAuth completion endpoint (successful and unsuccessful).
	CompleteRequestCount int64 `json:"completeRequestCount,omitempty"`
	// The number of successful requests to the broker OAuth completion endpoint.  Successful requests have authenticated  the user and established a browser session.
	CompleteSuccessCount int64 `json:"completeSuccessCount,omitempty"`
	// The default global access level for this OAuth profile. The allowed values and their meaning are:  <pre> \"none\" - User has no access to global data. \"read-only\" - User has read-only access to global data. \"read-write\" - User has read-write access to most global data. \"admin\" - User has read-write access to all global data. </pre> 
	DefaultGlobalAccessLevel string `json:"defaultGlobalAccessLevel,omitempty"`
	// The default message VPN access level for the OAuth profile. The allowed values and their meaning are:  <pre> \"none\" - User has no access to a Message VPN. \"read-only\" - User has read-only access to a Message VPN. \"read-write\" - User has read-write access to most Message VPN settings. </pre> 
	DefaultMsgVpnAccessLevel string `json:"defaultMsgVpnAccessLevel,omitempty"`
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
	// The user friendly name for the OAuth profile.
	DisplayName string `json:"displayName,omitempty"`
	// Enable or disable the OAuth profile.
	Enabled **bool `json:"enabled,omitempty"`
	// The OAuth authorization endpoint.
	EndpointAuthorization string `json:"endpointAuthorization,omitempty"`
	// The operational OAuth authorization endpoint.
	EndpointAuthorizationOperational string `json:"endpointAuthorizationOperational,omitempty"`
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
	// The OAuth token endpoint.
	EndpointToken string `json:"endpointToken,omitempty"`
	// The operational OAuth token endpoint.
	EndpointTokenOperational string `json:"endpointTokenOperational,omitempty"`
	// The maximum time in seconds a token request is allowed to take.
	EndpointTokenTimeout int32 `json:"endpointTokenTimeout,omitempty"`
	// The OpenID Connect Userinfo endpoint.
	EndpointUserinfo string `json:"endpointUserinfo,omitempty"`
	// The operational OpenID Connect Userinfo endpoint.
	EndpointUserinfoOperational string `json:"endpointUserinfoOperational,omitempty"`
	// The maximum time in seconds a userinfo request is allowed to take.
	EndpointUserinfoTimeout int32 `json:"endpointUserinfoTimeout,omitempty"`
	// The reason the profile is not active. The allowed values and their meaning are:  <pre> \"profile-disabled\" - The OAuth profile is disabled. \"missing-issuer\" - The issuer has not been discovered or configured. \"duplicate-issuer\" - Another OAuth profile already has the same issuer. \"none\" - The OAuth profile is active. </pre>  Available since 2.26.
	InactiveReason string `json:"inactiveReason,omitempty"`
	// The number of requests to the broker OAuth initiation endpoint that had an invalid error_link_uri parameter.
	InitiateInvalidErrorLinkCount int64 `json:"initiateInvalidErrorLinkCount,omitempty"`
	// The number of requests to the broker OAuth initiation endpoint that did not have a valid Host header.  See the Allowed Host configuration setting.
	InitiateInvalidHostCount int64 `json:"initiateInvalidHostCount,omitempty"`
	// The number of requests to the broker OAuth initiation endpoint that had an invalid target_link_uri parameter.
	InitiateInvalidTargetLinkCount int64 `json:"initiateInvalidTargetLinkCount,omitempty"`
	// The number of requests to the broker OAuth initiation endpoint (successful and unsuccessful).
	InitiateRequestCount int64 `json:"initiateRequestCount,omitempty"`
	// The number of requests to the broker OAuth initiation endpoint that successfully redirected to the OAuth provider's authorization endpoint.
	InitiateSuccessCount int64 `json:"initiateSuccessCount,omitempty"`
	// Enable or disable interactive logins via this OAuth provider.
	InteractiveEnabled **bool `json:"interactiveEnabled,omitempty"`
	// The number of times the groups were successfully found in the ID token or access token for interactive authentication.
	InteractiveGroupsFoundInTokenCount int64 `json:"interactiveGroupsFoundInTokenCount,omitempty"`
	// The number of failures during interactive authentication due to missing introspection configuration (a introspection request was required but no introspection endpoint was configured).
	InteractiveIntrospectionMissingCount int64 `json:"interactiveIntrospectionMissingCount,omitempty"`
	// The number of introspection request made from the broker during interactive authentication for this OAuth profile  where the configured groups claim wasn't found in the access token or the introspection response.
	InteractiveIntrospectionMissingGroupsCount int64 `json:"interactiveIntrospectionMissingGroupsCount,omitempty"`
	// The number of introspection requests made from the broker during interactive authentication for this OAuth profile where the configured username claim wasn't found in the access token or introspection response.
	InteractiveIntrospectionMissingUsernameCount int64 `json:"interactiveIntrospectionMissingUsernameCount,omitempty"`
	// The number of requests made to the introspection endpoint during interactive authentication.
	InteractiveIntrospectionRequestCount int64 `json:"interactiveIntrospectionRequestCount,omitempty"`
	// The number of introspection responses during interactive authentication that couldn't be parsed.
	InteractiveIntrospectionResponseInvalidCount int64 `json:"interactiveIntrospectionResponseInvalidCount,omitempty"`
	// The number of introspection requests made from the broker during interactive authentication for this OAuth profile with 200 status responses.
	InteractiveIntrospectionStatusOkCount int64 `json:"interactiveIntrospectionStatusOkCount,omitempty"`
	// The number of introspection requests made from the broker during interactive authentication for this OAuth profile with status responses other than 200.
	InteractiveIntrospectionStatusOtherCount int64 `json:"interactiveIntrospectionStatusOtherCount,omitempty"`
	// The number of introspection responses indicating that the provided token was not active.
	InteractiveIntrospectionTokenNotActiveCount int64 `json:"interactiveIntrospectionTokenNotActiveCount,omitempty"`
	// The value of the prompt parameter provided to the OAuth authorization server for login requests where the session has expired.
	InteractivePromptForExpiredSession string `json:"interactivePromptForExpiredSession,omitempty"`
	// The value of the prompt parameter provided to the OAuth authorization server for login requests where the session is new or the user has explicitly logged out.
	InteractivePromptForNewSession string `json:"interactivePromptForNewSession,omitempty"`
	// The number of failures due to missing Userinfo configuration (a Userinfo request was required but no Userinfo endpoint was configured) during interactive authentication.
	InteractiveUserinfoMissingCount int64 `json:"interactiveUserinfoMissingCount,omitempty"`
	// The number of Userinfo request made from the broker during interactive authentication for this OAuth profile where the configured groups claim wasn't found in the ID token or the Userinfo response.
	InteractiveUserinfoMissingGroupsCount int64 `json:"interactiveUserinfoMissingGroupsCount,omitempty"`
	// The number of Userinfo requests made from the broker during interactive authentication for this OAuth profile where the configured username claim wasn't found in the ID token or Userinfo response.
	InteractiveUserinfoMissingUsernameCount int64 `json:"interactiveUserinfoMissingUsernameCount,omitempty"`
	// The number of requests made to the Userinfo endpoint during interactive authentication.
	InteractiveUserinfoRequestCount int64 `json:"interactiveUserinfoRequestCount,omitempty"`
	// The number of Userinfo requests made from the broker during interactive authentication for this OAuth profile with responses that couldn't be parsed.
	InteractiveUserinfoResponseInvalidCount int64 `json:"interactiveUserinfoResponseInvalidCount,omitempty"`
	// The number of Userinfo requests made from the broker during interactive authentication for this OAuth profile with 200 status responses.
	InteractiveUserinfoStatusOkCount int64 `json:"interactiveUserinfoStatusOkCount,omitempty"`
	// The number of Userinfo requests made from the broker during interactive authentication for this OAuth profile with status responses other than 200.
	InteractiveUserinfoStatusOtherCount int64 `json:"interactiveUserinfoStatusOtherCount,omitempty"`
	// The number of Userinfo requests made from the broker during interactive authentication for this OAuth profile with subject claims that did not match the subject from the ID token.
	InteractiveUserinfoSubjectMismatchCount int64 `json:"interactiveUserinfoSubjectMismatchCount,omitempty"`
	// The number of time the username was successfully found in the ID token or access token for interactive authentication.
	InteractiveUsernameFoundInTokenCount int64 `json:"interactiveUsernameFoundInTokenCount,omitempty"`
	// The one minute average of the time required to complete a token introspection, in milliseconds (ms).
	IntrospectionAverageTime int32 `json:"introspectionAverageTime,omitempty"`
	// The reason for the introspection endpoint request failure.
	IntrospectionLastFailureReason string `json:"introspectionLastFailureReason,omitempty"`
	// The timestamp of the last introspection endpoint request failure. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time).
	IntrospectionLastFailureTime int32 `json:"introspectionLastFailureTime,omitempty"`
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
	// The name of the OAuth profile.
	OauthProfileName string `json:"oauthProfileName,omitempty"`
	// The OAuth role of the broker. The allowed values and their meaning are:  <pre> \"client\" - The broker is in the OAuth client role. \"resource-server\" - The broker is in the OAuth resource server role. </pre> 
	OauthRole string `json:"oauthRole,omitempty"`
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
	// Enable or disable authentication of SEMP requests with OAuth tokens.
	SempEnabled **bool `json:"sempEnabled,omitempty"`
	// The number of SEMP requests with an expired OAuth token.
	SempExpiredTokenCount int64 `json:"sempExpiredTokenCount,omitempty"`
	// The number of times the groups were successfully found in the ID token or access token for SEMP request authentication.
	SempGroupsFoundInTokenCount int64 `json:"sempGroupsFoundInTokenCount,omitempty"`
	// The number of failures during SEMP request authentication due to missing introspection configuration (a introspection request was required but no introspection endpoint was configured).
	SempIntrospectionMissingCount int64 `json:"sempIntrospectionMissingCount,omitempty"`
	// The number of introspection request made from the broker during SEMP request authentication for this OAuth profile where the configured groups claim wasn't found in the access token or the introspection response.
	SempIntrospectionMissingGroupsCount int64 `json:"sempIntrospectionMissingGroupsCount,omitempty"`
	// The number of introspection requests made from the broker during SEMP request authentication for this OAuth profile where the configured username claim wasn't found in the access token or introspection response.
	SempIntrospectionMissingUsernameCount int64 `json:"sempIntrospectionMissingUsernameCount,omitempty"`
	// The number of requests made to the introspection endpoint during SEMP request authentication.
	SempIntrospectionRequestCount int64 `json:"sempIntrospectionRequestCount,omitempty"`
	// The number of introspection responses during SEMP request authentication that couldn't be parsed.
	SempIntrospectionResponseInvalidCount int64 `json:"sempIntrospectionResponseInvalidCount,omitempty"`
	// The number of introspection requests made from the broker during SEMP request authentication for this OAuth profile with 200 status responses.
	SempIntrospectionStatusOkCount int64 `json:"sempIntrospectionStatusOkCount,omitempty"`
	// The number of introspection requests made from the broker during SEMP request authentication for this OAuth profile with status responses other than 200.
	SempIntrospectionStatusOtherCount int64 `json:"sempIntrospectionStatusOtherCount,omitempty"`
	// The number of introspection responses indicating that the provided token was not active.
	SempIntrospectionTokenNotActiveCount int64 `json:"sempIntrospectionTokenNotActiveCount,omitempty"`
	// The number of SEMP requests with an invalid OAuth token.
	SempInvalidTokenCount int64 `json:"sempInvalidTokenCount,omitempty"`
	// The number of SEMP requests (successful and unsuccessful) using this OAuth profile.
	SempRequestCount int64 `json:"sempRequestCount,omitempty"`
	// The number of successful SEMP authentications using this OAuth profile.
	SempSuccessCount int64 `json:"sempSuccessCount,omitempty"`
	// The number of failures due to missing Userinfo configuration (a Userinfo request was required but no Userinfo endpoint was configured) during SEMP request authentication.
	SempUserinfoMissingCount int64 `json:"sempUserinfoMissingCount,omitempty"`
	// The number of Userinfo request made from the broker during SEMP request authentication for this OAuth profile where the configured groups claim wasn't found in the ID token or the Userinfo response.
	SempUserinfoMissingGroupsCount int64 `json:"sempUserinfoMissingGroupsCount,omitempty"`
	// The number of Userinfo requests made from the broker during SEMP request authentication for this OAuth profile where the configured username claim wasn't found in the ID token or Userinfo response.
	SempUserinfoMissingUsernameCount int64 `json:"sempUserinfoMissingUsernameCount,omitempty"`
	// The number of requests made to the Userinfo endpoint during SEMP request authentication.
	SempUserinfoRequestCount int64 `json:"sempUserinfoRequestCount,omitempty"`
	// The number of Userinfo requests made from the broker during SEMP request authentication for this OAuth profile with responses that couldn't be parsed.
	SempUserinfoResponseInvalidCount int64 `json:"sempUserinfoResponseInvalidCount,omitempty"`
	// The number of Userinfo requests made from the broker during SEMP request authentication for this OAuth profile with 200 status responses.
	SempUserinfoStatusOkCount int64 `json:"sempUserinfoStatusOkCount,omitempty"`
	// The number of Userinfo requests made from the broker during SEMP request authentication for this OAuth profile with status responses other than 200.
	SempUserinfoStatusOtherCount int64 `json:"sempUserinfoStatusOtherCount,omitempty"`
	// The number of Userinfo requests made from the broker during SEMP request authentication for this OAuth profile with subject claims that did not match the subject from the ID token.
	SempUserinfoSubjectMismatchCount int64 `json:"sempUserinfoSubjectMismatchCount,omitempty"`
	// The number of time the username was successfully found in the ID token or access token for SEMP request authentication.
	SempUsernameFoundInTokenCount int64 `json:"sempUsernameFoundInTokenCount,omitempty"`
	// The one minute average of the time required to complete a token request, in milliseconds (ms).
	TokenEndpointAverageTime int32 `json:"tokenEndpointAverageTime,omitempty"`
	// The number of token endpoint requests made from the broker for this OAuth profile that returned tokens that couldn't be verified.
	TokenEndpointInvalidTokenCount int64 `json:"tokenEndpointInvalidTokenCount,omitempty"`
	// The reason for the last token endpoint request failure.
	TokenEndpointLastFailureReason string `json:"tokenEndpointLastFailureReason,omitempty"`
	// The timestamp of the last token endpoint request failure. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time).
	TokenEndpointLastFailureTime int32 `json:"tokenEndpointLastFailureTime,omitempty"`
	// The number of token endpoint requests made from the broker for this OAuth profile that returned an unexpected error not accounted for in the other failure statistics.
	TokenEndpointOtherErrorCount int64 `json:"tokenEndpointOtherErrorCount,omitempty"`
	// The number of token endpoint requests made from the broker for this OAuth profile.
	TokenEndpointRequestCount int64 `json:"tokenEndpointRequestCount,omitempty"`
	// The number of token endpoint requests made from the broker for this OAuth profile with 400 status responses.
	TokenEndpointStatusBadRequestCount int64 `json:"tokenEndpointStatusBadRequestCount,omitempty"`
	// The number of token endpoint requests made from the broker for this OAuth profile with 200 status responses.
	TokenEndpointStatusOkCount int64 `json:"tokenEndpointStatusOkCount,omitempty"`
	// The number of token endpoint requests made from the broker for this OAuth profile with status responses other than 200 or 400.
	TokenEndpointStatusOtherCount int64 `json:"tokenEndpointStatusOtherCount,omitempty"`
	// The one minute average of the time required to complete a token userinfo request, in milliseconds (ms).
	UserinfoAverageTime int32 `json:"userinfoAverageTime,omitempty"`
	// The reason for the userinfo endpoint request failure.
	UserinfoLastFailureReason string `json:"userinfoLastFailureReason,omitempty"`
	// The timestamp of the last userinfo endpoint request failure. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time).
	UserinfoLastFailureTime int32 `json:"userinfoLastFailureTime,omitempty"`
	// The name of the username claim.
	UsernameClaimName string `json:"usernameClaimName,omitempty"`
}
