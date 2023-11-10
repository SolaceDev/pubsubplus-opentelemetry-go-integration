
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

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"fmt"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type DistributedCacheApiService service
/*
DistributedCacheApiService Get a Distributed Cache object.
Get a Distributed Cache object.  A Distributed Cache is a collection of one or more Cache Clusters that belong to the same Message VPN. Each Cache Cluster in a Distributed Cache is configured to subscribe to a different set of topics. This effectively divides up the configured topic space, to provide scaling to very large topic spaces or very high cached message throughput.   Attribute|Identifying|Deprecated :---|:---:|:---: cacheName|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/read-only\&quot; is required to perform this operation.  This has been available since 2.11.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param cacheName The name of the Distributed Cache.
 * @param optional nil or *DistributedCacheApiGetMsgVpnDistributedCacheOpts - Optional Parameters:
     * @param "Select_" (optional.Interface of []string) -  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter.
@return MsgVpnDistributedCacheResponse
*/

type DistributedCacheApiGetMsgVpnDistributedCacheOpts struct {
    Select_ optional.Interface
}

func (a *DistributedCacheApiService) GetMsgVpnDistributedCache(ctx context.Context, msgVpnName string, cacheName string, localVarOptionals *DistributedCacheApiGetMsgVpnDistributedCacheOpts) (MsgVpnDistributedCacheResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue MsgVpnDistributedCacheResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/msgVpns/{msgVpnName}/distributedCaches/{cacheName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cacheName"+"}", fmt.Sprintf("%v", cacheName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Select_.IsSet() {
		localVarQueryParams.Add("select", parameterToString(localVarOptionals.Select_.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v MsgVpnDistributedCacheResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
DistributedCacheApiService Get a Cache Cluster object.
Get a Cache Cluster object.  A Cache Cluster is a collection of one or more Cache Instances that subscribe to exactly the same topics. Cache Instances are grouped together in a Cache Cluster for the purpose of fault tolerance and load balancing. As published messages are received, the message broker message bus sends these live data messages to the Cache Instances in the Cache Cluster. This enables client cache requests to be served by any of Cache Instances in the Cache Cluster.   Attribute|Identifying|Deprecated :---|:---:|:---: cacheName|x| clusterName|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/read-only\&quot; is required to perform this operation.  This has been available since 2.11.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param cacheName The name of the Distributed Cache.
 * @param clusterName The name of the Cache Cluster.
 * @param optional nil or *DistributedCacheApiGetMsgVpnDistributedCacheClusterOpts - Optional Parameters:
     * @param "Select_" (optional.Interface of []string) -  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter.
@return MsgVpnDistributedCacheClusterResponse
*/

type DistributedCacheApiGetMsgVpnDistributedCacheClusterOpts struct {
    Select_ optional.Interface
}

func (a *DistributedCacheApiService) GetMsgVpnDistributedCacheCluster(ctx context.Context, msgVpnName string, cacheName string, clusterName string, localVarOptionals *DistributedCacheApiGetMsgVpnDistributedCacheClusterOpts) (MsgVpnDistributedCacheClusterResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue MsgVpnDistributedCacheClusterResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cacheName"+"}", fmt.Sprintf("%v", cacheName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", fmt.Sprintf("%v", clusterName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Select_.IsSet() {
		localVarQueryParams.Add("select", parameterToString(localVarOptionals.Select_.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v MsgVpnDistributedCacheClusterResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
DistributedCacheApiService Get a Home Cache Cluster object.
Get a Home Cache Cluster object.  A Home Cache Cluster is a Cache Cluster that is the \&quot;definitive\&quot; Cache Cluster for a given topic in the context of the Global Caching feature.   Attribute|Identifying|Deprecated :---|:---:|:---: cacheName|x| clusterName|x| homeClusterName|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/read-only\&quot; is required to perform this operation.  This has been available since 2.11.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param cacheName The name of the Distributed Cache.
 * @param clusterName The name of the Cache Cluster.
 * @param homeClusterName The name of the remote Home Cache Cluster.
 * @param optional nil or *DistributedCacheApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterOpts - Optional Parameters:
     * @param "Select_" (optional.Interface of []string) -  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter.
@return MsgVpnDistributedCacheClusterGlobalCachingHomeClusterResponse
*/

type DistributedCacheApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterOpts struct {
    Select_ optional.Interface
}

func (a *DistributedCacheApiService) GetMsgVpnDistributedCacheClusterGlobalCachingHomeCluster(ctx context.Context, msgVpnName string, cacheName string, clusterName string, homeClusterName string, localVarOptionals *DistributedCacheApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterOpts) (MsgVpnDistributedCacheClusterGlobalCachingHomeClusterResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue MsgVpnDistributedCacheClusterGlobalCachingHomeClusterResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cacheName"+"}", fmt.Sprintf("%v", cacheName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", fmt.Sprintf("%v", clusterName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"homeClusterName"+"}", fmt.Sprintf("%v", homeClusterName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Select_.IsSet() {
		localVarQueryParams.Add("select", parameterToString(localVarOptionals.Select_.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v MsgVpnDistributedCacheClusterGlobalCachingHomeClusterResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
DistributedCacheApiService Get a Topic Prefix object.
Get a Topic Prefix object.  A Topic Prefix is a prefix for a global topic that is available from the containing Home Cache Cluster.   Attribute|Identifying|Deprecated :---|:---:|:---: cacheName|x| clusterName|x| homeClusterName|x| msgVpnName|x| topicPrefix|x|    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/read-only\&quot; is required to perform this operation.  This has been available since 2.11.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param cacheName The name of the Distributed Cache.
 * @param clusterName The name of the Cache Cluster.
 * @param homeClusterName The name of the remote Home Cache Cluster.
 * @param topicPrefix A topic prefix for global topics available from the remote Home Cache Cluster. A wildcard (/&gt;) is implied at the end of the prefix.
 * @param optional nil or *DistributedCacheApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixOpts - Optional Parameters:
     * @param "Select_" (optional.Interface of []string) -  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter.
@return MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixResponse
*/

type DistributedCacheApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixOpts struct {
    Select_ optional.Interface
}

func (a *DistributedCacheApiService) GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix(ctx context.Context, msgVpnName string, cacheName string, clusterName string, homeClusterName string, topicPrefix string, localVarOptionals *DistributedCacheApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixOpts) (MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName}/topicPrefixes/{topicPrefix}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cacheName"+"}", fmt.Sprintf("%v", cacheName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", fmt.Sprintf("%v", clusterName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"homeClusterName"+"}", fmt.Sprintf("%v", homeClusterName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"topicPrefix"+"}", fmt.Sprintf("%v", topicPrefix), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Select_.IsSet() {
		localVarQueryParams.Add("select", parameterToString(localVarOptionals.Select_.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
DistributedCacheApiService Get a list of Topic Prefix objects.
Get a list of Topic Prefix objects.  A Topic Prefix is a prefix for a global topic that is available from the containing Home Cache Cluster.   Attribute|Identifying|Deprecated :---|:---:|:---: cacheName|x| clusterName|x| homeClusterName|x| msgVpnName|x| topicPrefix|x|    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/read-only\&quot; is required to perform this operation.  This has been available since 2.11.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param cacheName The name of the Distributed Cache.
 * @param clusterName The name of the Cache Cluster.
 * @param homeClusterName The name of the remote Home Cache Cluster.
 * @param optional nil or *DistributedCacheApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesOpts - Optional Parameters:
     * @param "Count" (optional.Int32) -  Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter.
     * @param "Cursor" (optional.String) -  The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter.
     * @param "Where" (optional.Interface of []string) -  Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter.
     * @param "Select_" (optional.Interface of []string) -  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter.
@return MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesResponse
*/

type DistributedCacheApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesOpts struct {
    Count optional.Int32
    Cursor optional.String
    Where optional.Interface
    Select_ optional.Interface
}

func (a *DistributedCacheApiService) GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixes(ctx context.Context, msgVpnName string, cacheName string, clusterName string, homeClusterName string, localVarOptionals *DistributedCacheApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesOpts) (MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName}/topicPrefixes"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cacheName"+"}", fmt.Sprintf("%v", cacheName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", fmt.Sprintf("%v", clusterName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"homeClusterName"+"}", fmt.Sprintf("%v", homeClusterName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Count.IsSet() {
		localVarQueryParams.Add("count", parameterToString(localVarOptionals.Count.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Cursor.IsSet() {
		localVarQueryParams.Add("cursor", parameterToString(localVarOptionals.Cursor.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Where.IsSet() {
		localVarQueryParams.Add("where", parameterToString(localVarOptionals.Where.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.Select_.IsSet() {
		localVarQueryParams.Add("select", parameterToString(localVarOptionals.Select_.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
DistributedCacheApiService Get a list of Home Cache Cluster objects.
Get a list of Home Cache Cluster objects.  A Home Cache Cluster is a Cache Cluster that is the \&quot;definitive\&quot; Cache Cluster for a given topic in the context of the Global Caching feature.   Attribute|Identifying|Deprecated :---|:---:|:---: cacheName|x| clusterName|x| homeClusterName|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/read-only\&quot; is required to perform this operation.  This has been available since 2.11.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param cacheName The name of the Distributed Cache.
 * @param clusterName The name of the Cache Cluster.
 * @param optional nil or *DistributedCacheApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClustersOpts - Optional Parameters:
     * @param "Count" (optional.Int32) -  Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter.
     * @param "Cursor" (optional.String) -  The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter.
     * @param "Where" (optional.Interface of []string) -  Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter.
     * @param "Select_" (optional.Interface of []string) -  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter.
@return MsgVpnDistributedCacheClusterGlobalCachingHomeClustersResponse
*/

type DistributedCacheApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClustersOpts struct {
    Count optional.Int32
    Cursor optional.String
    Where optional.Interface
    Select_ optional.Interface
}

func (a *DistributedCacheApiService) GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusters(ctx context.Context, msgVpnName string, cacheName string, clusterName string, localVarOptionals *DistributedCacheApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClustersOpts) (MsgVpnDistributedCacheClusterGlobalCachingHomeClustersResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue MsgVpnDistributedCacheClusterGlobalCachingHomeClustersResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cacheName"+"}", fmt.Sprintf("%v", cacheName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", fmt.Sprintf("%v", clusterName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Count.IsSet() {
		localVarQueryParams.Add("count", parameterToString(localVarOptionals.Count.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Cursor.IsSet() {
		localVarQueryParams.Add("cursor", parameterToString(localVarOptionals.Cursor.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Where.IsSet() {
		localVarQueryParams.Add("where", parameterToString(localVarOptionals.Where.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.Select_.IsSet() {
		localVarQueryParams.Add("select", parameterToString(localVarOptionals.Select_.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v MsgVpnDistributedCacheClusterGlobalCachingHomeClustersResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
DistributedCacheApiService Get a Cache Instance object.
Get a Cache Instance object.  A Cache Instance is a single Cache process that belongs to a single Cache Cluster. A Cache Instance object provisioned on the broker is used to disseminate configuration information to the Cache process. Cache Instances listen for and cache live data messages that match the topic subscriptions configured for their parent Cache Cluster.   Attribute|Identifying|Deprecated :---|:---:|:---: cacheName|x| clusterName|x| counter.msgCount||x counter.msgPeakCount||x counter.requestQueueDepthCount||x counter.requestQueueDepthPeakCount||x counter.topicCount||x counter.topicPeakCount||x instanceName|x| msgVpnName|x| rate.averageDataRxBytePeakRate||x rate.averageDataRxByteRate||x rate.averageDataRxMsgPeakRate||x rate.averageDataRxMsgRate||x rate.averageDataTxMsgPeakRate||x rate.averageDataTxMsgRate||x rate.averageRequestRxPeakRate||x rate.averageRequestRxRate||x rate.dataRxBytePeakRate||x rate.dataRxByteRate||x rate.dataRxMsgPeakRate||x rate.dataRxMsgRate||x rate.dataTxMsgPeakRate||x rate.dataTxMsgRate||x rate.requestRxPeakRate||x rate.requestRxRate||x    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/read-only\&quot; is required to perform this operation.  This has been available since 2.11.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param cacheName The name of the Distributed Cache.
 * @param clusterName The name of the Cache Cluster.
 * @param instanceName The name of the Cache Instance.
 * @param optional nil or *DistributedCacheApiGetMsgVpnDistributedCacheClusterInstanceOpts - Optional Parameters:
     * @param "Select_" (optional.Interface of []string) -  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter.
@return MsgVpnDistributedCacheClusterInstanceResponse
*/

type DistributedCacheApiGetMsgVpnDistributedCacheClusterInstanceOpts struct {
    Select_ optional.Interface
}

func (a *DistributedCacheApiService) GetMsgVpnDistributedCacheClusterInstance(ctx context.Context, msgVpnName string, cacheName string, clusterName string, instanceName string, localVarOptionals *DistributedCacheApiGetMsgVpnDistributedCacheClusterInstanceOpts) (MsgVpnDistributedCacheClusterInstanceResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue MsgVpnDistributedCacheClusterInstanceResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cacheName"+"}", fmt.Sprintf("%v", cacheName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", fmt.Sprintf("%v", clusterName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"instanceName"+"}", fmt.Sprintf("%v", instanceName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Select_.IsSet() {
		localVarQueryParams.Add("select", parameterToString(localVarOptionals.Select_.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v MsgVpnDistributedCacheClusterInstanceResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
DistributedCacheApiService Get a Remote Home Cache Cluster object.
Get a Remote Home Cache Cluster object.  A Remote Home Cache Cluster is a Home Cache Cluster that the Cache Instance is communicating with in the context of the Global Caching feature.   Attribute|Identifying|Deprecated :---|:---:|:---: cacheName|x| clusterName|x| homeClusterName|x| instanceName|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/read-only\&quot; is required to perform this operation.  This has been available since 2.11.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param cacheName The name of the Distributed Cache.
 * @param clusterName The name of the Cache Cluster.
 * @param instanceName The name of the Cache Instance.
 * @param homeClusterName The name of the remote Home Cache Cluster.
 * @param optional nil or *DistributedCacheApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterOpts - Optional Parameters:
     * @param "Select_" (optional.Interface of []string) -  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter.
@return MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterResponse
*/

type DistributedCacheApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterOpts struct {
    Select_ optional.Interface
}

func (a *DistributedCacheApiService) GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster(ctx context.Context, msgVpnName string, cacheName string, clusterName string, instanceName string, homeClusterName string, localVarOptionals *DistributedCacheApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterOpts) (MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/remoteGlobalCachingHomeClusters/{homeClusterName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cacheName"+"}", fmt.Sprintf("%v", cacheName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", fmt.Sprintf("%v", clusterName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"instanceName"+"}", fmt.Sprintf("%v", instanceName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"homeClusterName"+"}", fmt.Sprintf("%v", homeClusterName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Select_.IsSet() {
		localVarQueryParams.Add("select", parameterToString(localVarOptionals.Select_.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
DistributedCacheApiService Get a list of Remote Home Cache Cluster objects.
Get a list of Remote Home Cache Cluster objects.  A Remote Home Cache Cluster is a Home Cache Cluster that the Cache Instance is communicating with in the context of the Global Caching feature.   Attribute|Identifying|Deprecated :---|:---:|:---: cacheName|x| clusterName|x| homeClusterName|x| instanceName|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/read-only\&quot; is required to perform this operation.  This has been available since 2.11.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param cacheName The name of the Distributed Cache.
 * @param clusterName The name of the Cache Cluster.
 * @param instanceName The name of the Cache Instance.
 * @param optional nil or *DistributedCacheApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersOpts - Optional Parameters:
     * @param "Count" (optional.Int32) -  Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter.
     * @param "Cursor" (optional.String) -  The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter.
     * @param "Where" (optional.Interface of []string) -  Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter.
     * @param "Select_" (optional.Interface of []string) -  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter.
@return MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersResponse
*/

type DistributedCacheApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersOpts struct {
    Count optional.Int32
    Cursor optional.String
    Where optional.Interface
    Select_ optional.Interface
}

func (a *DistributedCacheApiService) GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusters(ctx context.Context, msgVpnName string, cacheName string, clusterName string, instanceName string, localVarOptionals *DistributedCacheApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersOpts) (MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/remoteGlobalCachingHomeClusters"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cacheName"+"}", fmt.Sprintf("%v", cacheName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", fmt.Sprintf("%v", clusterName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"instanceName"+"}", fmt.Sprintf("%v", instanceName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Count.IsSet() {
		localVarQueryParams.Add("count", parameterToString(localVarOptionals.Count.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Cursor.IsSet() {
		localVarQueryParams.Add("cursor", parameterToString(localVarOptionals.Cursor.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Where.IsSet() {
		localVarQueryParams.Add("where", parameterToString(localVarOptionals.Where.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.Select_.IsSet() {
		localVarQueryParams.Add("select", parameterToString(localVarOptionals.Select_.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
DistributedCacheApiService Get a Remote Topic object.
Get a Remote Topic object.  A Remote Topic is a topic for which the Cache Instance has cached messages.   Attribute|Identifying|Deprecated :---|:---:|:---: cacheName|x| clusterName|x| instanceName|x| msgVpnName|x| topic|x|    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/read-only\&quot; is required to perform this operation.  This has been available since 2.11.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param cacheName The name of the Distributed Cache.
 * @param clusterName The name of the Cache Cluster.
 * @param instanceName The name of the Cache Instance.
 * @param topic The value of the remote Topic.
 * @param optional nil or *DistributedCacheApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicOpts - Optional Parameters:
     * @param "Select_" (optional.Interface of []string) -  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter.
@return MsgVpnDistributedCacheClusterInstanceRemoteTopicResponse
*/

type DistributedCacheApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicOpts struct {
    Select_ optional.Interface
}

func (a *DistributedCacheApiService) GetMsgVpnDistributedCacheClusterInstanceRemoteTopic(ctx context.Context, msgVpnName string, cacheName string, clusterName string, instanceName string, topic string, localVarOptionals *DistributedCacheApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicOpts) (MsgVpnDistributedCacheClusterInstanceRemoteTopicResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue MsgVpnDistributedCacheClusterInstanceRemoteTopicResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/remoteTopics/{topic}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cacheName"+"}", fmt.Sprintf("%v", cacheName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", fmt.Sprintf("%v", clusterName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"instanceName"+"}", fmt.Sprintf("%v", instanceName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"topic"+"}", fmt.Sprintf("%v", topic), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Select_.IsSet() {
		localVarQueryParams.Add("select", parameterToString(localVarOptionals.Select_.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v MsgVpnDistributedCacheClusterInstanceRemoteTopicResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
DistributedCacheApiService Get a list of Remote Topic objects.
Get a list of Remote Topic objects.  A Remote Topic is a topic for which the Cache Instance has cached messages.   Attribute|Identifying|Deprecated :---|:---:|:---: cacheName|x| clusterName|x| instanceName|x| msgVpnName|x| topic|x|    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/read-only\&quot; is required to perform this operation.  This has been available since 2.11.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param cacheName The name of the Distributed Cache.
 * @param clusterName The name of the Cache Cluster.
 * @param instanceName The name of the Cache Instance.
 * @param optional nil or *DistributedCacheApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicsOpts - Optional Parameters:
     * @param "Count" (optional.Int32) -  Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter.
     * @param "Cursor" (optional.String) -  The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter.
     * @param "Where" (optional.Interface of []string) -  Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter.
     * @param "Select_" (optional.Interface of []string) -  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter.
@return MsgVpnDistributedCacheClusterInstanceRemoteTopicsResponse
*/

type DistributedCacheApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicsOpts struct {
    Count optional.Int32
    Cursor optional.String
    Where optional.Interface
    Select_ optional.Interface
}

func (a *DistributedCacheApiService) GetMsgVpnDistributedCacheClusterInstanceRemoteTopics(ctx context.Context, msgVpnName string, cacheName string, clusterName string, instanceName string, localVarOptionals *DistributedCacheApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicsOpts) (MsgVpnDistributedCacheClusterInstanceRemoteTopicsResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue MsgVpnDistributedCacheClusterInstanceRemoteTopicsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/remoteTopics"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cacheName"+"}", fmt.Sprintf("%v", cacheName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", fmt.Sprintf("%v", clusterName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"instanceName"+"}", fmt.Sprintf("%v", instanceName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Count.IsSet() {
		localVarQueryParams.Add("count", parameterToString(localVarOptionals.Count.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Cursor.IsSet() {
		localVarQueryParams.Add("cursor", parameterToString(localVarOptionals.Cursor.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Where.IsSet() {
		localVarQueryParams.Add("where", parameterToString(localVarOptionals.Where.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.Select_.IsSet() {
		localVarQueryParams.Add("select", parameterToString(localVarOptionals.Select_.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v MsgVpnDistributedCacheClusterInstanceRemoteTopicsResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
DistributedCacheApiService Get a list of Cache Instance objects.
Get a list of Cache Instance objects.  A Cache Instance is a single Cache process that belongs to a single Cache Cluster. A Cache Instance object provisioned on the broker is used to disseminate configuration information to the Cache process. Cache Instances listen for and cache live data messages that match the topic subscriptions configured for their parent Cache Cluster.   Attribute|Identifying|Deprecated :---|:---:|:---: cacheName|x| clusterName|x| counter.msgCount||x counter.msgPeakCount||x counter.requestQueueDepthCount||x counter.requestQueueDepthPeakCount||x counter.topicCount||x counter.topicPeakCount||x instanceName|x| msgVpnName|x| rate.averageDataRxBytePeakRate||x rate.averageDataRxByteRate||x rate.averageDataRxMsgPeakRate||x rate.averageDataRxMsgRate||x rate.averageDataTxMsgPeakRate||x rate.averageDataTxMsgRate||x rate.averageRequestRxPeakRate||x rate.averageRequestRxRate||x rate.dataRxBytePeakRate||x rate.dataRxByteRate||x rate.dataRxMsgPeakRate||x rate.dataRxMsgRate||x rate.dataTxMsgPeakRate||x rate.dataTxMsgRate||x rate.requestRxPeakRate||x rate.requestRxRate||x    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/read-only\&quot; is required to perform this operation.  This has been available since 2.11.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param cacheName The name of the Distributed Cache.
 * @param clusterName The name of the Cache Cluster.
 * @param optional nil or *DistributedCacheApiGetMsgVpnDistributedCacheClusterInstancesOpts - Optional Parameters:
     * @param "Count" (optional.Int32) -  Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter.
     * @param "Cursor" (optional.String) -  The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter.
     * @param "Where" (optional.Interface of []string) -  Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter.
     * @param "Select_" (optional.Interface of []string) -  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter.
@return MsgVpnDistributedCacheClusterInstancesResponse
*/

type DistributedCacheApiGetMsgVpnDistributedCacheClusterInstancesOpts struct {
    Count optional.Int32
    Cursor optional.String
    Where optional.Interface
    Select_ optional.Interface
}

func (a *DistributedCacheApiService) GetMsgVpnDistributedCacheClusterInstances(ctx context.Context, msgVpnName string, cacheName string, clusterName string, localVarOptionals *DistributedCacheApiGetMsgVpnDistributedCacheClusterInstancesOpts) (MsgVpnDistributedCacheClusterInstancesResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue MsgVpnDistributedCacheClusterInstancesResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cacheName"+"}", fmt.Sprintf("%v", cacheName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", fmt.Sprintf("%v", clusterName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Count.IsSet() {
		localVarQueryParams.Add("count", parameterToString(localVarOptionals.Count.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Cursor.IsSet() {
		localVarQueryParams.Add("cursor", parameterToString(localVarOptionals.Cursor.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Where.IsSet() {
		localVarQueryParams.Add("where", parameterToString(localVarOptionals.Where.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.Select_.IsSet() {
		localVarQueryParams.Add("select", parameterToString(localVarOptionals.Select_.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v MsgVpnDistributedCacheClusterInstancesResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
DistributedCacheApiService Get a Topic object.
Get a Topic object.  The Cache Instances that belong to the containing Cache Cluster will cache any messages published to topics that match a Topic Subscription.   Attribute|Identifying|Deprecated :---|:---:|:---: cacheName|x| clusterName|x| msgVpnName|x| topic|x|    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/read-only\&quot; is required to perform this operation.  This has been available since 2.11.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param cacheName The name of the Distributed Cache.
 * @param clusterName The name of the Cache Cluster.
 * @param topic The value of the Topic in the form a/b/c.
 * @param optional nil or *DistributedCacheApiGetMsgVpnDistributedCacheClusterTopicOpts - Optional Parameters:
     * @param "Select_" (optional.Interface of []string) -  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter.
@return MsgVpnDistributedCacheClusterTopicResponse
*/

type DistributedCacheApiGetMsgVpnDistributedCacheClusterTopicOpts struct {
    Select_ optional.Interface
}

func (a *DistributedCacheApiService) GetMsgVpnDistributedCacheClusterTopic(ctx context.Context, msgVpnName string, cacheName string, clusterName string, topic string, localVarOptionals *DistributedCacheApiGetMsgVpnDistributedCacheClusterTopicOpts) (MsgVpnDistributedCacheClusterTopicResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue MsgVpnDistributedCacheClusterTopicResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/topics/{topic}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cacheName"+"}", fmt.Sprintf("%v", cacheName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", fmt.Sprintf("%v", clusterName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"topic"+"}", fmt.Sprintf("%v", topic), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Select_.IsSet() {
		localVarQueryParams.Add("select", parameterToString(localVarOptionals.Select_.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v MsgVpnDistributedCacheClusterTopicResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
DistributedCacheApiService Get a list of Topic objects.
Get a list of Topic objects.  The Cache Instances that belong to the containing Cache Cluster will cache any messages published to topics that match a Topic Subscription.   Attribute|Identifying|Deprecated :---|:---:|:---: cacheName|x| clusterName|x| msgVpnName|x| topic|x|    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/read-only\&quot; is required to perform this operation.  This has been available since 2.11.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param cacheName The name of the Distributed Cache.
 * @param clusterName The name of the Cache Cluster.
 * @param optional nil or *DistributedCacheApiGetMsgVpnDistributedCacheClusterTopicsOpts - Optional Parameters:
     * @param "Count" (optional.Int32) -  Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter.
     * @param "Cursor" (optional.String) -  The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter.
     * @param "Where" (optional.Interface of []string) -  Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter.
     * @param "Select_" (optional.Interface of []string) -  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter.
@return MsgVpnDistributedCacheClusterTopicsResponse
*/

type DistributedCacheApiGetMsgVpnDistributedCacheClusterTopicsOpts struct {
    Count optional.Int32
    Cursor optional.String
    Where optional.Interface
    Select_ optional.Interface
}

func (a *DistributedCacheApiService) GetMsgVpnDistributedCacheClusterTopics(ctx context.Context, msgVpnName string, cacheName string, clusterName string, localVarOptionals *DistributedCacheApiGetMsgVpnDistributedCacheClusterTopicsOpts) (MsgVpnDistributedCacheClusterTopicsResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue MsgVpnDistributedCacheClusterTopicsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/topics"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cacheName"+"}", fmt.Sprintf("%v", cacheName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", fmt.Sprintf("%v", clusterName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Count.IsSet() {
		localVarQueryParams.Add("count", parameterToString(localVarOptionals.Count.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Cursor.IsSet() {
		localVarQueryParams.Add("cursor", parameterToString(localVarOptionals.Cursor.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Where.IsSet() {
		localVarQueryParams.Add("where", parameterToString(localVarOptionals.Where.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.Select_.IsSet() {
		localVarQueryParams.Add("select", parameterToString(localVarOptionals.Select_.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v MsgVpnDistributedCacheClusterTopicsResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
DistributedCacheApiService Get a list of Cache Cluster objects.
Get a list of Cache Cluster objects.  A Cache Cluster is a collection of one or more Cache Instances that subscribe to exactly the same topics. Cache Instances are grouped together in a Cache Cluster for the purpose of fault tolerance and load balancing. As published messages are received, the message broker message bus sends these live data messages to the Cache Instances in the Cache Cluster. This enables client cache requests to be served by any of Cache Instances in the Cache Cluster.   Attribute|Identifying|Deprecated :---|:---:|:---: cacheName|x| clusterName|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/read-only\&quot; is required to perform this operation.  This has been available since 2.11.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param cacheName The name of the Distributed Cache.
 * @param optional nil or *DistributedCacheApiGetMsgVpnDistributedCacheClustersOpts - Optional Parameters:
     * @param "Count" (optional.Int32) -  Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter.
     * @param "Cursor" (optional.String) -  The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter.
     * @param "Where" (optional.Interface of []string) -  Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter.
     * @param "Select_" (optional.Interface of []string) -  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter.
@return MsgVpnDistributedCacheClustersResponse
*/

type DistributedCacheApiGetMsgVpnDistributedCacheClustersOpts struct {
    Count optional.Int32
    Cursor optional.String
    Where optional.Interface
    Select_ optional.Interface
}

func (a *DistributedCacheApiService) GetMsgVpnDistributedCacheClusters(ctx context.Context, msgVpnName string, cacheName string, localVarOptionals *DistributedCacheApiGetMsgVpnDistributedCacheClustersOpts) (MsgVpnDistributedCacheClustersResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue MsgVpnDistributedCacheClustersResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cacheName"+"}", fmt.Sprintf("%v", cacheName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Count.IsSet() {
		localVarQueryParams.Add("count", parameterToString(localVarOptionals.Count.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Cursor.IsSet() {
		localVarQueryParams.Add("cursor", parameterToString(localVarOptionals.Cursor.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Where.IsSet() {
		localVarQueryParams.Add("where", parameterToString(localVarOptionals.Where.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.Select_.IsSet() {
		localVarQueryParams.Add("select", parameterToString(localVarOptionals.Select_.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v MsgVpnDistributedCacheClustersResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
DistributedCacheApiService Get a list of Distributed Cache objects.
Get a list of Distributed Cache objects.  A Distributed Cache is a collection of one or more Cache Clusters that belong to the same Message VPN. Each Cache Cluster in a Distributed Cache is configured to subscribe to a different set of topics. This effectively divides up the configured topic space, to provide scaling to very large topic spaces or very high cached message throughput.   Attribute|Identifying|Deprecated :---|:---:|:---: cacheName|x| msgVpnName|x|    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/read-only\&quot; is required to perform this operation.  This has been available since 2.11.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param optional nil or *DistributedCacheApiGetMsgVpnDistributedCachesOpts - Optional Parameters:
     * @param "Count" (optional.Int32) -  Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter.
     * @param "Cursor" (optional.String) -  The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter.
     * @param "Where" (optional.Interface of []string) -  Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter.
     * @param "Select_" (optional.Interface of []string) -  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter.
@return MsgVpnDistributedCachesResponse
*/

type DistributedCacheApiGetMsgVpnDistributedCachesOpts struct {
    Count optional.Int32
    Cursor optional.String
    Where optional.Interface
    Select_ optional.Interface
}

func (a *DistributedCacheApiService) GetMsgVpnDistributedCaches(ctx context.Context, msgVpnName string, localVarOptionals *DistributedCacheApiGetMsgVpnDistributedCachesOpts) (MsgVpnDistributedCachesResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue MsgVpnDistributedCachesResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/msgVpns/{msgVpnName}/distributedCaches"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Count.IsSet() {
		localVarQueryParams.Add("count", parameterToString(localVarOptionals.Count.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Cursor.IsSet() {
		localVarQueryParams.Add("cursor", parameterToString(localVarOptionals.Cursor.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Where.IsSet() {
		localVarQueryParams.Add("where", parameterToString(localVarOptionals.Where.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.Select_.IsSet() {
		localVarQueryParams.Add("select", parameterToString(localVarOptionals.Select_.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v MsgVpnDistributedCachesResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
