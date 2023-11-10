# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/config*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDmrCluster**](DmrClusterApi.md#CreateDmrCluster) | **Post** /dmrClusters | Create a Cluster object.
[**CreateDmrClusterCertMatchingRule**](DmrClusterApi.md#CreateDmrClusterCertMatchingRule) | **Post** /dmrClusters/{dmrClusterName}/certMatchingRules | Create a Certificate Matching Rule object.
[**CreateDmrClusterCertMatchingRuleAttributeFilter**](DmrClusterApi.md#CreateDmrClusterCertMatchingRuleAttributeFilter) | **Post** /dmrClusters/{dmrClusterName}/certMatchingRules/{ruleName}/attributeFilters | Create a Certificate Matching Rule Attribute Filter object.
[**CreateDmrClusterCertMatchingRuleCondition**](DmrClusterApi.md#CreateDmrClusterCertMatchingRuleCondition) | **Post** /dmrClusters/{dmrClusterName}/certMatchingRules/{ruleName}/conditions | Create a Certificate Matching Rule Condition object.
[**CreateDmrClusterLink**](DmrClusterApi.md#CreateDmrClusterLink) | **Post** /dmrClusters/{dmrClusterName}/links | Create a Link object.
[**CreateDmrClusterLinkAttribute**](DmrClusterApi.md#CreateDmrClusterLinkAttribute) | **Post** /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/attributes | Create a Link Attribute object.
[**CreateDmrClusterLinkRemoteAddress**](DmrClusterApi.md#CreateDmrClusterLinkRemoteAddress) | **Post** /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/remoteAddresses | Create a Remote Address object.
[**CreateDmrClusterLinkTlsTrustedCommonName**](DmrClusterApi.md#CreateDmrClusterLinkTlsTrustedCommonName) | **Post** /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/tlsTrustedCommonNames | Create a Trusted Common Name object.
[**DeleteDmrCluster**](DmrClusterApi.md#DeleteDmrCluster) | **Delete** /dmrClusters/{dmrClusterName} | Delete a Cluster object.
[**DeleteDmrClusterCertMatchingRule**](DmrClusterApi.md#DeleteDmrClusterCertMatchingRule) | **Delete** /dmrClusters/{dmrClusterName}/certMatchingRules/{ruleName} | Delete a Certificate Matching Rule object.
[**DeleteDmrClusterCertMatchingRuleAttributeFilter**](DmrClusterApi.md#DeleteDmrClusterCertMatchingRuleAttributeFilter) | **Delete** /dmrClusters/{dmrClusterName}/certMatchingRules/{ruleName}/attributeFilters/{filterName} | Delete a Certificate Matching Rule Attribute Filter object.
[**DeleteDmrClusterCertMatchingRuleCondition**](DmrClusterApi.md#DeleteDmrClusterCertMatchingRuleCondition) | **Delete** /dmrClusters/{dmrClusterName}/certMatchingRules/{ruleName}/conditions/{source} | Delete a Certificate Matching Rule Condition object.
[**DeleteDmrClusterLink**](DmrClusterApi.md#DeleteDmrClusterLink) | **Delete** /dmrClusters/{dmrClusterName}/links/{remoteNodeName} | Delete a Link object.
[**DeleteDmrClusterLinkAttribute**](DmrClusterApi.md#DeleteDmrClusterLinkAttribute) | **Delete** /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/attributes/{attributeName},{attributeValue} | Delete a Link Attribute object.
[**DeleteDmrClusterLinkRemoteAddress**](DmrClusterApi.md#DeleteDmrClusterLinkRemoteAddress) | **Delete** /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/remoteAddresses/{remoteAddress} | Delete a Remote Address object.
[**DeleteDmrClusterLinkTlsTrustedCommonName**](DmrClusterApi.md#DeleteDmrClusterLinkTlsTrustedCommonName) | **Delete** /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/tlsTrustedCommonNames/{tlsTrustedCommonName} | Delete a Trusted Common Name object.
[**GetDmrCluster**](DmrClusterApi.md#GetDmrCluster) | **Get** /dmrClusters/{dmrClusterName} | Get a Cluster object.
[**GetDmrClusterCertMatchingRule**](DmrClusterApi.md#GetDmrClusterCertMatchingRule) | **Get** /dmrClusters/{dmrClusterName}/certMatchingRules/{ruleName} | Get a Certificate Matching Rule object.
[**GetDmrClusterCertMatchingRuleAttributeFilter**](DmrClusterApi.md#GetDmrClusterCertMatchingRuleAttributeFilter) | **Get** /dmrClusters/{dmrClusterName}/certMatchingRules/{ruleName}/attributeFilters/{filterName} | Get a Certificate Matching Rule Attribute Filter object.
[**GetDmrClusterCertMatchingRuleAttributeFilters**](DmrClusterApi.md#GetDmrClusterCertMatchingRuleAttributeFilters) | **Get** /dmrClusters/{dmrClusterName}/certMatchingRules/{ruleName}/attributeFilters | Get a list of Certificate Matching Rule Attribute Filter objects.
[**GetDmrClusterCertMatchingRuleCondition**](DmrClusterApi.md#GetDmrClusterCertMatchingRuleCondition) | **Get** /dmrClusters/{dmrClusterName}/certMatchingRules/{ruleName}/conditions/{source} | Get a Certificate Matching Rule Condition object.
[**GetDmrClusterCertMatchingRuleConditions**](DmrClusterApi.md#GetDmrClusterCertMatchingRuleConditions) | **Get** /dmrClusters/{dmrClusterName}/certMatchingRules/{ruleName}/conditions | Get a list of Certificate Matching Rule Condition objects.
[**GetDmrClusterCertMatchingRules**](DmrClusterApi.md#GetDmrClusterCertMatchingRules) | **Get** /dmrClusters/{dmrClusterName}/certMatchingRules | Get a list of Certificate Matching Rule objects.
[**GetDmrClusterLink**](DmrClusterApi.md#GetDmrClusterLink) | **Get** /dmrClusters/{dmrClusterName}/links/{remoteNodeName} | Get a Link object.
[**GetDmrClusterLinkAttribute**](DmrClusterApi.md#GetDmrClusterLinkAttribute) | **Get** /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/attributes/{attributeName},{attributeValue} | Get a Link Attribute object.
[**GetDmrClusterLinkAttributes**](DmrClusterApi.md#GetDmrClusterLinkAttributes) | **Get** /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/attributes | Get a list of Link Attribute objects.
[**GetDmrClusterLinkRemoteAddress**](DmrClusterApi.md#GetDmrClusterLinkRemoteAddress) | **Get** /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/remoteAddresses/{remoteAddress} | Get a Remote Address object.
[**GetDmrClusterLinkRemoteAddresses**](DmrClusterApi.md#GetDmrClusterLinkRemoteAddresses) | **Get** /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/remoteAddresses | Get a list of Remote Address objects.
[**GetDmrClusterLinkTlsTrustedCommonName**](DmrClusterApi.md#GetDmrClusterLinkTlsTrustedCommonName) | **Get** /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/tlsTrustedCommonNames/{tlsTrustedCommonName} | Get a Trusted Common Name object.
[**GetDmrClusterLinkTlsTrustedCommonNames**](DmrClusterApi.md#GetDmrClusterLinkTlsTrustedCommonNames) | **Get** /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/tlsTrustedCommonNames | Get a list of Trusted Common Name objects.
[**GetDmrClusterLinks**](DmrClusterApi.md#GetDmrClusterLinks) | **Get** /dmrClusters/{dmrClusterName}/links | Get a list of Link objects.
[**GetDmrClusters**](DmrClusterApi.md#GetDmrClusters) | **Get** /dmrClusters | Get a list of Cluster objects.
[**ReplaceDmrCluster**](DmrClusterApi.md#ReplaceDmrCluster) | **Put** /dmrClusters/{dmrClusterName} | Replace a Cluster object.
[**ReplaceDmrClusterCertMatchingRule**](DmrClusterApi.md#ReplaceDmrClusterCertMatchingRule) | **Put** /dmrClusters/{dmrClusterName}/certMatchingRules/{ruleName} | Replace a Certificate Matching Rule object.
[**ReplaceDmrClusterCertMatchingRuleAttributeFilter**](DmrClusterApi.md#ReplaceDmrClusterCertMatchingRuleAttributeFilter) | **Put** /dmrClusters/{dmrClusterName}/certMatchingRules/{ruleName}/attributeFilters/{filterName} | Replace a Certificate Matching Rule Attribute Filter object.
[**ReplaceDmrClusterLink**](DmrClusterApi.md#ReplaceDmrClusterLink) | **Put** /dmrClusters/{dmrClusterName}/links/{remoteNodeName} | Replace a Link object.
[**UpdateDmrCluster**](DmrClusterApi.md#UpdateDmrCluster) | **Patch** /dmrClusters/{dmrClusterName} | Update a Cluster object.
[**UpdateDmrClusterCertMatchingRule**](DmrClusterApi.md#UpdateDmrClusterCertMatchingRule) | **Patch** /dmrClusters/{dmrClusterName}/certMatchingRules/{ruleName} | Update a Certificate Matching Rule object.
[**UpdateDmrClusterCertMatchingRuleAttributeFilter**](DmrClusterApi.md#UpdateDmrClusterCertMatchingRuleAttributeFilter) | **Patch** /dmrClusters/{dmrClusterName}/certMatchingRules/{ruleName}/attributeFilters/{filterName} | Update a Certificate Matching Rule Attribute Filter object.
[**UpdateDmrClusterLink**](DmrClusterApi.md#UpdateDmrClusterLink) | **Patch** /dmrClusters/{dmrClusterName}/links/{remoteNodeName} | Update a Link object.

# **CreateDmrCluster**
> DmrClusterResponse CreateDmrCluster(ctx, body, optional)
Create a Cluster object.

Create a Cluster object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates via config-sync.  A Cluster is a provisioned object on a message broker that contains global DMR configuration parameters.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: authenticationBasicPassword||||x||x authenticationClientCertContent||||x||x authenticationClientCertPassword||||x|| dmrClusterName|x|x|||| nodeName|||x||| tlsServerCertEnforceTrustedCommonNameEnabled|||||x|    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- DmrCluster|authenticationClientCertPassword|authenticationClientCertContent|    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DmrCluster**](DmrCluster.md)| The Cluster object&#x27;s attributes. | 
 **optional** | ***DmrClusterApiCreateDmrClusterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DmrClusterApiCreateDmrClusterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterResponse**](DmrClusterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDmrClusterCertMatchingRule**
> DmrClusterCertMatchingRuleResponse CreateDmrClusterCertMatchingRule(ctx, body, dmrClusterName, optional)
Create a Certificate Matching Rule object.

Create a Certificate Matching Rule object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates via config-sync.  A Cert Matching Rule is a collection of conditions and attribute filters that all have to be satisfied for certificate to be acceptable as authentication for a given link.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: dmrClusterName|x||x||| ruleName|x|x||||    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.28.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DmrClusterCertMatchingRule**](DmrClusterCertMatchingRule.md)| The Certificate Matching Rule object&#x27;s attributes. | 
  **dmrClusterName** | **string**| The name of the Cluster. | 
 **optional** | ***DmrClusterApiCreateDmrClusterCertMatchingRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DmrClusterApiCreateDmrClusterCertMatchingRuleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterCertMatchingRuleResponse**](DmrClusterCertMatchingRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDmrClusterCertMatchingRuleAttributeFilter**
> DmrClusterCertMatchingRuleAttributeFilterResponse CreateDmrClusterCertMatchingRuleAttributeFilter(ctx, body, dmrClusterName, ruleName, optional)
Create a Certificate Matching Rule Attribute Filter object.

Create a Certificate Matching Rule Attribute Filter object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates via config-sync.  A Cert Matching Rule Attribute Filter compares a link attribute to a string.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: dmrClusterName|x||x||| filterName|x|x|||| ruleName|x||x|||    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.28.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DmrClusterCertMatchingRuleAttributeFilter**](DmrClusterCertMatchingRuleAttributeFilter.md)| The Certificate Matching Rule Attribute Filter object&#x27;s attributes. | 
  **dmrClusterName** | **string**| The name of the Cluster. | 
  **ruleName** | **string**| The name of the rule. | 
 **optional** | ***DmrClusterApiCreateDmrClusterCertMatchingRuleAttributeFilterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DmrClusterApiCreateDmrClusterCertMatchingRuleAttributeFilterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterCertMatchingRuleAttributeFilterResponse**](DmrClusterCertMatchingRuleAttributeFilterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDmrClusterCertMatchingRuleCondition**
> DmrClusterCertMatchingRuleConditionResponse CreateDmrClusterCertMatchingRuleCondition(ctx, body, dmrClusterName, ruleName, optional)
Create a Certificate Matching Rule Condition object.

Create a Certificate Matching Rule Condition object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates via config-sync.  A Cert Matching Rule Condition compares data extracted from a certificate to a link attribute or an expression.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: dmrClusterName|x||x||| ruleName|x||x||| source|x|x||||    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.28.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DmrClusterCertMatchingRuleCondition**](DmrClusterCertMatchingRuleCondition.md)| The Certificate Matching Rule Condition object&#x27;s attributes. | 
  **dmrClusterName** | **string**| The name of the Cluster. | 
  **ruleName** | **string**| The name of the rule. | 
 **optional** | ***DmrClusterApiCreateDmrClusterCertMatchingRuleConditionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DmrClusterApiCreateDmrClusterCertMatchingRuleConditionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterCertMatchingRuleConditionResponse**](DmrClusterCertMatchingRuleConditionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDmrClusterLink**
> DmrClusterLinkResponse CreateDmrClusterLink(ctx, body, dmrClusterName, optional)
Create a Link object.

Create a Link object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates via config-sync.  A Link connects nodes (either within a Cluster or between two different Clusters) and allows them to exchange topology information, subscriptions and data.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: authenticationBasicPassword||||x||x dmrClusterName|x||x||| remoteNodeName|x|x||||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThreshold|clearPercent|setPercent|clearValue, setValue EventThreshold|clearValue|setValue|clearPercent, setPercent EventThreshold|setPercent|clearPercent|clearValue, setValue EventThreshold|setValue|clearValue|clearPercent, setPercent    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DmrClusterLink**](DmrClusterLink.md)| The Link object&#x27;s attributes. | 
  **dmrClusterName** | **string**| The name of the Cluster. | 
 **optional** | ***DmrClusterApiCreateDmrClusterLinkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DmrClusterApiCreateDmrClusterLinkOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterLinkResponse**](DmrClusterLinkResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDmrClusterLinkAttribute**
> DmrClusterLinkAttributeResponse CreateDmrClusterLinkAttribute(ctx, body, dmrClusterName, remoteNodeName, optional)
Create a Link Attribute object.

Create a Link Attribute object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates via config-sync.  A Link Attribute is a key+value pair that can be used to locate a DMR Cluster Link, for example when using client certificate mapping.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: attributeName|x|x|||| attributeValue|x|x|||| dmrClusterName|x||x||| remoteNodeName|x||x|||    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.28.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DmrClusterLinkAttribute**](DmrClusterLinkAttribute.md)| The Link Attribute object&#x27;s attributes. | 
  **dmrClusterName** | **string**| The name of the Cluster. | 
  **remoteNodeName** | **string**| The name of the node at the remote end of the Link. | 
 **optional** | ***DmrClusterApiCreateDmrClusterLinkAttributeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DmrClusterApiCreateDmrClusterLinkAttributeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterLinkAttributeResponse**](DmrClusterLinkAttributeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDmrClusterLinkRemoteAddress**
> DmrClusterLinkRemoteAddressResponse CreateDmrClusterLinkRemoteAddress(ctx, body, dmrClusterName, remoteNodeName, optional)
Create a Remote Address object.

Create a Remote Address object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates via config-sync.  Each Remote Address, consisting of a FQDN or IP address and optional port, is used to connect to the remote node for this Link. Up to 4 addresses may be provided for each Link, and will be tried on a round-robin basis.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: dmrClusterName|x||x||| remoteAddress|x|x|||| remoteNodeName|x||x|||    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DmrClusterLinkRemoteAddress**](DmrClusterLinkRemoteAddress.md)| The Remote Address object&#x27;s attributes. | 
  **dmrClusterName** | **string**| The name of the Cluster. | 
  **remoteNodeName** | **string**| The name of the node at the remote end of the Link. | 
 **optional** | ***DmrClusterApiCreateDmrClusterLinkRemoteAddressOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DmrClusterApiCreateDmrClusterLinkRemoteAddressOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterLinkRemoteAddressResponse**](DmrClusterLinkRemoteAddressResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDmrClusterLinkTlsTrustedCommonName**
> DmrClusterLinkTlsTrustedCommonNameResponse CreateDmrClusterLinkTlsTrustedCommonName(ctx, body, dmrClusterName, remoteNodeName, optional)
Create a Trusted Common Name object.

Create a Trusted Common Name object. Any attribute missing from the request will be set to its default value. The creation of instances of this object are synchronized to HA mates via config-sync.  The Trusted Common Names for the Link are used by encrypted transports to verify the name in the certificate presented by the remote node. They must include the common name of the remote node's server certificate or client certificate, depending upon the initiator of the connection.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: dmrClusterName|x||x||x| remoteNodeName|x||x||x| tlsTrustedCommonName|x|x|||x|    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been deprecated since 2.18. Common Name validation has been replaced by Server Certificate Name validation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DmrClusterLinkTlsTrustedCommonName**](DmrClusterLinkTlsTrustedCommonName.md)| The Trusted Common Name object&#x27;s attributes. | 
  **dmrClusterName** | **string**| The name of the Cluster. | 
  **remoteNodeName** | **string**| The name of the node at the remote end of the Link. | 
 **optional** | ***DmrClusterApiCreateDmrClusterLinkTlsTrustedCommonNameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DmrClusterApiCreateDmrClusterLinkTlsTrustedCommonNameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterLinkTlsTrustedCommonNameResponse**](DmrClusterLinkTlsTrustedCommonNameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDmrCluster**
> SempMetaOnlyResponse DeleteDmrCluster(ctx, dmrClusterName)
Delete a Cluster object.

Delete a Cluster object. The deletion of instances of this object are synchronized to HA mates via config-sync.  A Cluster is a provisioned object on a message broker that contains global DMR configuration parameters.  A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dmrClusterName** | **string**| The name of the Cluster. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDmrClusterCertMatchingRule**
> SempMetaOnlyResponse DeleteDmrClusterCertMatchingRule(ctx, dmrClusterName, ruleName)
Delete a Certificate Matching Rule object.

Delete a Certificate Matching Rule object. The deletion of instances of this object are synchronized to HA mates via config-sync.  A Cert Matching Rule is a collection of conditions and attribute filters that all have to be satisfied for certificate to be acceptable as authentication for a given link.  A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.28.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dmrClusterName** | **string**| The name of the Cluster. | 
  **ruleName** | **string**| The name of the rule. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDmrClusterCertMatchingRuleAttributeFilter**
> SempMetaOnlyResponse DeleteDmrClusterCertMatchingRuleAttributeFilter(ctx, dmrClusterName, ruleName, filterName)
Delete a Certificate Matching Rule Attribute Filter object.

Delete a Certificate Matching Rule Attribute Filter object. The deletion of instances of this object are synchronized to HA mates via config-sync.  A Cert Matching Rule Attribute Filter compares a link attribute to a string.  A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.28.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dmrClusterName** | **string**| The name of the Cluster. | 
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

# **DeleteDmrClusterCertMatchingRuleCondition**
> SempMetaOnlyResponse DeleteDmrClusterCertMatchingRuleCondition(ctx, dmrClusterName, ruleName, source)
Delete a Certificate Matching Rule Condition object.

Delete a Certificate Matching Rule Condition object. The deletion of instances of this object are synchronized to HA mates via config-sync.  A Cert Matching Rule Condition compares data extracted from a certificate to a link attribute or an expression.  A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.28.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dmrClusterName** | **string**| The name of the Cluster. | 
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

# **DeleteDmrClusterLink**
> SempMetaOnlyResponse DeleteDmrClusterLink(ctx, dmrClusterName, remoteNodeName)
Delete a Link object.

Delete a Link object. The deletion of instances of this object are synchronized to HA mates via config-sync.  A Link connects nodes (either within a Cluster or between two different Clusters) and allows them to exchange topology information, subscriptions and data.  A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dmrClusterName** | **string**| The name of the Cluster. | 
  **remoteNodeName** | **string**| The name of the node at the remote end of the Link. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDmrClusterLinkAttribute**
> SempMetaOnlyResponse DeleteDmrClusterLinkAttribute(ctx, dmrClusterName, remoteNodeName, attributeName, attributeValue)
Delete a Link Attribute object.

Delete a Link Attribute object. The deletion of instances of this object are synchronized to HA mates via config-sync.  A Link Attribute is a key+value pair that can be used to locate a DMR Cluster Link, for example when using client certificate mapping.  A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.28.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dmrClusterName** | **string**| The name of the Cluster. | 
  **remoteNodeName** | **string**| The name of the node at the remote end of the Link. | 
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

# **DeleteDmrClusterLinkRemoteAddress**
> SempMetaOnlyResponse DeleteDmrClusterLinkRemoteAddress(ctx, dmrClusterName, remoteNodeName, remoteAddress)
Delete a Remote Address object.

Delete a Remote Address object. The deletion of instances of this object are synchronized to HA mates via config-sync.  Each Remote Address, consisting of a FQDN or IP address and optional port, is used to connect to the remote node for this Link. Up to 4 addresses may be provided for each Link, and will be tried on a round-robin basis.  A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dmrClusterName** | **string**| The name of the Cluster. | 
  **remoteNodeName** | **string**| The name of the node at the remote end of the Link. | 
  **remoteAddress** | **string**| The FQDN or IP address (and optional port) of the remote node. If a port is not provided, it will vary based on the transport encoding: 55555 (plain-text), 55443 (encrypted), or 55003 (compressed). | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDmrClusterLinkTlsTrustedCommonName**
> SempMetaOnlyResponse DeleteDmrClusterLinkTlsTrustedCommonName(ctx, dmrClusterName, remoteNodeName, tlsTrustedCommonName)
Delete a Trusted Common Name object.

Delete a Trusted Common Name object. The deletion of instances of this object are synchronized to HA mates via config-sync.  The Trusted Common Names for the Link are used by encrypted transports to verify the name in the certificate presented by the remote node. They must include the common name of the remote node's server certificate or client certificate, depending upon the initiator of the connection.  A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been deprecated since 2.18. Common Name validation has been replaced by Server Certificate Name validation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dmrClusterName** | **string**| The name of the Cluster. | 
  **remoteNodeName** | **string**| The name of the node at the remote end of the Link. | 
  **tlsTrustedCommonName** | **string**| The expected trusted common name of the remote certificate. | 

### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDmrCluster**
> DmrClusterResponse GetDmrCluster(ctx, dmrClusterName, optional)
Get a Cluster object.

Get a Cluster object.  A Cluster is a provisioned object on a message broker that contains global DMR configuration parameters.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: authenticationBasicPassword||x||x authenticationClientCertContent||x||x authenticationClientCertPassword||x|| dmrClusterName|x||| tlsServerCertEnforceTrustedCommonNameEnabled|||x|    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dmrClusterName** | **string**| The name of the Cluster. | 
 **optional** | ***DmrClusterApiGetDmrClusterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DmrClusterApiGetDmrClusterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterResponse**](DmrClusterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDmrClusterCertMatchingRule**
> DmrClusterCertMatchingRuleResponse GetDmrClusterCertMatchingRule(ctx, dmrClusterName, ruleName, optional)
Get a Certificate Matching Rule object.

Get a Certificate Matching Rule object.  A Cert Matching Rule is a collection of conditions and attribute filters that all have to be satisfied for certificate to be acceptable as authentication for a given link.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: dmrClusterName|x||| ruleName|x|||    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.28.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dmrClusterName** | **string**| The name of the Cluster. | 
  **ruleName** | **string**| The name of the rule. | 
 **optional** | ***DmrClusterApiGetDmrClusterCertMatchingRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DmrClusterApiGetDmrClusterCertMatchingRuleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterCertMatchingRuleResponse**](DmrClusterCertMatchingRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDmrClusterCertMatchingRuleAttributeFilter**
> DmrClusterCertMatchingRuleAttributeFilterResponse GetDmrClusterCertMatchingRuleAttributeFilter(ctx, dmrClusterName, ruleName, filterName, optional)
Get a Certificate Matching Rule Attribute Filter object.

Get a Certificate Matching Rule Attribute Filter object.  A Cert Matching Rule Attribute Filter compares a link attribute to a string.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: dmrClusterName|x||| filterName|x||| ruleName|x|||    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.28.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dmrClusterName** | **string**| The name of the Cluster. | 
  **ruleName** | **string**| The name of the rule. | 
  **filterName** | **string**| The name of the filter. | 
 **optional** | ***DmrClusterApiGetDmrClusterCertMatchingRuleAttributeFilterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DmrClusterApiGetDmrClusterCertMatchingRuleAttributeFilterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterCertMatchingRuleAttributeFilterResponse**](DmrClusterCertMatchingRuleAttributeFilterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDmrClusterCertMatchingRuleAttributeFilters**
> DmrClusterCertMatchingRuleAttributeFiltersResponse GetDmrClusterCertMatchingRuleAttributeFilters(ctx, dmrClusterName, ruleName, optional)
Get a list of Certificate Matching Rule Attribute Filter objects.

Get a list of Certificate Matching Rule Attribute Filter objects.  A Cert Matching Rule Attribute Filter compares a link attribute to a string.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: dmrClusterName|x||| filterName|x||| ruleName|x|||    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.28.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dmrClusterName** | **string**| The name of the Cluster. | 
  **ruleName** | **string**| The name of the rule. | 
 **optional** | ***DmrClusterApiGetDmrClusterCertMatchingRuleAttributeFiltersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DmrClusterApiGetDmrClusterCertMatchingRuleAttributeFiltersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterCertMatchingRuleAttributeFiltersResponse**](DmrClusterCertMatchingRuleAttributeFiltersResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDmrClusterCertMatchingRuleCondition**
> DmrClusterCertMatchingRuleConditionResponse GetDmrClusterCertMatchingRuleCondition(ctx, dmrClusterName, ruleName, source, optional)
Get a Certificate Matching Rule Condition object.

Get a Certificate Matching Rule Condition object.  A Cert Matching Rule Condition compares data extracted from a certificate to a link attribute or an expression.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: dmrClusterName|x||| ruleName|x||| source|x|||    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.28.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dmrClusterName** | **string**| The name of the Cluster. | 
  **ruleName** | **string**| The name of the rule. | 
  **source** | **string**| Certificate field to be compared with the Attribute. | 
 **optional** | ***DmrClusterApiGetDmrClusterCertMatchingRuleConditionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DmrClusterApiGetDmrClusterCertMatchingRuleConditionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterCertMatchingRuleConditionResponse**](DmrClusterCertMatchingRuleConditionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDmrClusterCertMatchingRuleConditions**
> DmrClusterCertMatchingRuleConditionsResponse GetDmrClusterCertMatchingRuleConditions(ctx, dmrClusterName, ruleName, optional)
Get a list of Certificate Matching Rule Condition objects.

Get a list of Certificate Matching Rule Condition objects.  A Cert Matching Rule Condition compares data extracted from a certificate to a link attribute or an expression.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: dmrClusterName|x||| ruleName|x||| source|x|||    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.28.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dmrClusterName** | **string**| The name of the Cluster. | 
  **ruleName** | **string**| The name of the rule. | 
 **optional** | ***DmrClusterApiGetDmrClusterCertMatchingRuleConditionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DmrClusterApiGetDmrClusterCertMatchingRuleConditionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterCertMatchingRuleConditionsResponse**](DmrClusterCertMatchingRuleConditionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDmrClusterCertMatchingRules**
> DmrClusterCertMatchingRulesResponse GetDmrClusterCertMatchingRules(ctx, dmrClusterName, optional)
Get a list of Certificate Matching Rule objects.

Get a list of Certificate Matching Rule objects.  A Cert Matching Rule is a collection of conditions and attribute filters that all have to be satisfied for certificate to be acceptable as authentication for a given link.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: dmrClusterName|x||| ruleName|x|||    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.28.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dmrClusterName** | **string**| The name of the Cluster. | 
 **optional** | ***DmrClusterApiGetDmrClusterCertMatchingRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DmrClusterApiGetDmrClusterCertMatchingRulesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterCertMatchingRulesResponse**](DmrClusterCertMatchingRulesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDmrClusterLink**
> DmrClusterLinkResponse GetDmrClusterLink(ctx, dmrClusterName, remoteNodeName, optional)
Get a Link object.

Get a Link object.  A Link connects nodes (either within a Cluster or between two different Clusters) and allows them to exchange topology information, subscriptions and data.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: authenticationBasicPassword||x||x dmrClusterName|x||| remoteNodeName|x|||    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dmrClusterName** | **string**| The name of the Cluster. | 
  **remoteNodeName** | **string**| The name of the node at the remote end of the Link. | 
 **optional** | ***DmrClusterApiGetDmrClusterLinkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DmrClusterApiGetDmrClusterLinkOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterLinkResponse**](DmrClusterLinkResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDmrClusterLinkAttribute**
> DmrClusterLinkAttributeResponse GetDmrClusterLinkAttribute(ctx, dmrClusterName, remoteNodeName, attributeName, attributeValue, optional)
Get a Link Attribute object.

Get a Link Attribute object.  A Link Attribute is a key+value pair that can be used to locate a DMR Cluster Link, for example when using client certificate mapping.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: attributeName|x||| attributeValue|x||| dmrClusterName|x||| remoteNodeName|x|||    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.28.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dmrClusterName** | **string**| The name of the Cluster. | 
  **remoteNodeName** | **string**| The name of the node at the remote end of the Link. | 
  **attributeName** | **string**| The name of the Attribute. | 
  **attributeValue** | **string**| The value of the Attribute. | 
 **optional** | ***DmrClusterApiGetDmrClusterLinkAttributeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DmrClusterApiGetDmrClusterLinkAttributeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterLinkAttributeResponse**](DmrClusterLinkAttributeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDmrClusterLinkAttributes**
> DmrClusterLinkAttributesResponse GetDmrClusterLinkAttributes(ctx, dmrClusterName, remoteNodeName, optional)
Get a list of Link Attribute objects.

Get a list of Link Attribute objects.  A Link Attribute is a key+value pair that can be used to locate a DMR Cluster Link, for example when using client certificate mapping.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: attributeName|x||| attributeValue|x||| dmrClusterName|x||| remoteNodeName|x|||    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.28.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dmrClusterName** | **string**| The name of the Cluster. | 
  **remoteNodeName** | **string**| The name of the node at the remote end of the Link. | 
 **optional** | ***DmrClusterApiGetDmrClusterLinkAttributesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DmrClusterApiGetDmrClusterLinkAttributesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterLinkAttributesResponse**](DmrClusterLinkAttributesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDmrClusterLinkRemoteAddress**
> DmrClusterLinkRemoteAddressResponse GetDmrClusterLinkRemoteAddress(ctx, dmrClusterName, remoteNodeName, remoteAddress, optional)
Get a Remote Address object.

Get a Remote Address object.  Each Remote Address, consisting of a FQDN or IP address and optional port, is used to connect to the remote node for this Link. Up to 4 addresses may be provided for each Link, and will be tried on a round-robin basis.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: dmrClusterName|x||| remoteAddress|x||| remoteNodeName|x|||    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dmrClusterName** | **string**| The name of the Cluster. | 
  **remoteNodeName** | **string**| The name of the node at the remote end of the Link. | 
  **remoteAddress** | **string**| The FQDN or IP address (and optional port) of the remote node. If a port is not provided, it will vary based on the transport encoding: 55555 (plain-text), 55443 (encrypted), or 55003 (compressed). | 
 **optional** | ***DmrClusterApiGetDmrClusterLinkRemoteAddressOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DmrClusterApiGetDmrClusterLinkRemoteAddressOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterLinkRemoteAddressResponse**](DmrClusterLinkRemoteAddressResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDmrClusterLinkRemoteAddresses**
> DmrClusterLinkRemoteAddressesResponse GetDmrClusterLinkRemoteAddresses(ctx, dmrClusterName, remoteNodeName, optional)
Get a list of Remote Address objects.

Get a list of Remote Address objects.  Each Remote Address, consisting of a FQDN or IP address and optional port, is used to connect to the remote node for this Link. Up to 4 addresses may be provided for each Link, and will be tried on a round-robin basis.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: dmrClusterName|x||| remoteAddress|x||| remoteNodeName|x|||    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dmrClusterName** | **string**| The name of the Cluster. | 
  **remoteNodeName** | **string**| The name of the node at the remote end of the Link. | 
 **optional** | ***DmrClusterApiGetDmrClusterLinkRemoteAddressesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DmrClusterApiGetDmrClusterLinkRemoteAddressesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterLinkRemoteAddressesResponse**](DmrClusterLinkRemoteAddressesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDmrClusterLinkTlsTrustedCommonName**
> DmrClusterLinkTlsTrustedCommonNameResponse GetDmrClusterLinkTlsTrustedCommonName(ctx, dmrClusterName, remoteNodeName, tlsTrustedCommonName, optional)
Get a Trusted Common Name object.

Get a Trusted Common Name object.  The Trusted Common Names for the Link are used by encrypted transports to verify the name in the certificate presented by the remote node. They must include the common name of the remote node's server certificate or client certificate, depending upon the initiator of the connection.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: dmrClusterName|x||x| remoteNodeName|x||x| tlsTrustedCommonName|x||x|    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been deprecated since 2.18. Common Name validation has been replaced by Server Certificate Name validation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dmrClusterName** | **string**| The name of the Cluster. | 
  **remoteNodeName** | **string**| The name of the node at the remote end of the Link. | 
  **tlsTrustedCommonName** | **string**| The expected trusted common name of the remote certificate. | 
 **optional** | ***DmrClusterApiGetDmrClusterLinkTlsTrustedCommonNameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DmrClusterApiGetDmrClusterLinkTlsTrustedCommonNameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterLinkTlsTrustedCommonNameResponse**](DmrClusterLinkTlsTrustedCommonNameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDmrClusterLinkTlsTrustedCommonNames**
> DmrClusterLinkTlsTrustedCommonNamesResponse GetDmrClusterLinkTlsTrustedCommonNames(ctx, dmrClusterName, remoteNodeName, optional)
Get a list of Trusted Common Name objects.

Get a list of Trusted Common Name objects.  The Trusted Common Names for the Link are used by encrypted transports to verify the name in the certificate presented by the remote node. They must include the common name of the remote node's server certificate or client certificate, depending upon the initiator of the connection.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: dmrClusterName|x||x| remoteNodeName|x||x| tlsTrustedCommonName|x||x|    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been deprecated since 2.18. Common Name validation has been replaced by Server Certificate Name validation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dmrClusterName** | **string**| The name of the Cluster. | 
  **remoteNodeName** | **string**| The name of the node at the remote end of the Link. | 
 **optional** | ***DmrClusterApiGetDmrClusterLinkTlsTrustedCommonNamesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DmrClusterApiGetDmrClusterLinkTlsTrustedCommonNamesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterLinkTlsTrustedCommonNamesResponse**](DmrClusterLinkTlsTrustedCommonNamesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDmrClusterLinks**
> DmrClusterLinksResponse GetDmrClusterLinks(ctx, dmrClusterName, optional)
Get a list of Link objects.

Get a list of Link objects.  A Link connects nodes (either within a Cluster or between two different Clusters) and allows them to exchange topology information, subscriptions and data.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: authenticationBasicPassword||x||x dmrClusterName|x||| remoteNodeName|x|||    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dmrClusterName** | **string**| The name of the Cluster. | 
 **optional** | ***DmrClusterApiGetDmrClusterLinksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DmrClusterApiGetDmrClusterLinksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterLinksResponse**](DmrClusterLinksResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDmrClusters**
> DmrClustersResponse GetDmrClusters(ctx, optional)
Get a list of Cluster objects.

Get a list of Cluster objects.  A Cluster is a provisioned object on a message broker that contains global DMR configuration parameters.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: authenticationBasicPassword||x||x authenticationClientCertContent||x||x authenticationClientCertPassword||x|| dmrClusterName|x||| tlsServerCertEnforceTrustedCommonNameEnabled|||x|    A SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DmrClusterApiGetDmrClustersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DmrClusterApiGetDmrClustersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClustersResponse**](DmrClustersResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceDmrCluster**
> DmrClusterResponse ReplaceDmrCluster(ctx, body, dmrClusterName, optional)
Replace a Cluster object.

Replace a Cluster object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  A Cluster is a provisioned object on a message broker that contains global DMR configuration parameters.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- authenticationBasicPassword||||x||x||x authenticationClientCertContent||||x||x||x authenticationClientCertPassword||||x||x|| directOnlyEnabled||x|||||| dmrClusterName|x||x||||| nodeName|||x||||| tlsServerCertEnforceTrustedCommonNameEnabled|||||||x|    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- DmrCluster|authenticationClientCertPassword|authenticationClientCertContent|    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DmrCluster**](DmrCluster.md)| The Cluster object&#x27;s attributes. | 
  **dmrClusterName** | **string**| The name of the Cluster. | 
 **optional** | ***DmrClusterApiReplaceDmrClusterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DmrClusterApiReplaceDmrClusterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterResponse**](DmrClusterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceDmrClusterCertMatchingRule**
> DmrClusterCertMatchingRuleResponse ReplaceDmrClusterCertMatchingRule(ctx, body, dmrClusterName, ruleName, optional)
Replace a Certificate Matching Rule object.

Replace a Certificate Matching Rule object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  A Cert Matching Rule is a collection of conditions and attribute filters that all have to be satisfied for certificate to be acceptable as authentication for a given link.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- dmrClusterName|x||x||||| ruleName|x||x|||||    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.28.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DmrClusterCertMatchingRule**](DmrClusterCertMatchingRule.md)| The Certificate Matching Rule object&#x27;s attributes. | 
  **dmrClusterName** | **string**| The name of the Cluster. | 
  **ruleName** | **string**| The name of the rule. | 
 **optional** | ***DmrClusterApiReplaceDmrClusterCertMatchingRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DmrClusterApiReplaceDmrClusterCertMatchingRuleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterCertMatchingRuleResponse**](DmrClusterCertMatchingRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceDmrClusterCertMatchingRuleAttributeFilter**
> DmrClusterCertMatchingRuleAttributeFilterResponse ReplaceDmrClusterCertMatchingRuleAttributeFilter(ctx, body, dmrClusterName, ruleName, filterName, optional)
Replace a Certificate Matching Rule Attribute Filter object.

Replace a Certificate Matching Rule Attribute Filter object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  A Cert Matching Rule Attribute Filter compares a link attribute to a string.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- dmrClusterName|x||x||||| filterName|x||x||||| ruleName|x||x|||||    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.28.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DmrClusterCertMatchingRuleAttributeFilter**](DmrClusterCertMatchingRuleAttributeFilter.md)| The Certificate Matching Rule Attribute Filter object&#x27;s attributes. | 
  **dmrClusterName** | **string**| The name of the Cluster. | 
  **ruleName** | **string**| The name of the rule. | 
  **filterName** | **string**| The name of the filter. | 
 **optional** | ***DmrClusterApiReplaceDmrClusterCertMatchingRuleAttributeFilterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DmrClusterApiReplaceDmrClusterCertMatchingRuleAttributeFilterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterCertMatchingRuleAttributeFilterResponse**](DmrClusterCertMatchingRuleAttributeFilterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceDmrClusterLink**
> DmrClusterLinkResponse ReplaceDmrClusterLink(ctx, body, dmrClusterName, remoteNodeName, optional)
Replace a Link object.

Replace a Link object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  A Link connects nodes (either within a Cluster or between two different Clusters) and allows them to exchange topology information, subscriptions and data.   Attribute|Identifying|Const|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:---|:--- authenticationBasicPassword||||x||x||x authenticationScheme||||||x|| dmrClusterName|x||x||||| egressFlowWindowSize||||||x|| initiator||||||x|| remoteNodeName|x||x||||| span||||||x|| transportCompressedEnabled||||||x|| transportTlsEnabled||||||x||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThreshold|clearPercent|setPercent|clearValue, setValue EventThreshold|clearValue|setValue|clearPercent, setPercent EventThreshold|setPercent|clearPercent|clearValue, setValue EventThreshold|setValue|clearValue|clearPercent, setPercent    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DmrClusterLink**](DmrClusterLink.md)| The Link object&#x27;s attributes. | 
  **dmrClusterName** | **string**| The name of the Cluster. | 
  **remoteNodeName** | **string**| The name of the node at the remote end of the Link. | 
 **optional** | ***DmrClusterApiReplaceDmrClusterLinkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DmrClusterApiReplaceDmrClusterLinkOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterLinkResponse**](DmrClusterLinkResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDmrCluster**
> DmrClusterResponse UpdateDmrCluster(ctx, body, dmrClusterName, optional)
Update a Cluster object.

Update a Cluster object. Any attribute missing from the request will be left unchanged.  A Cluster is a provisioned object on a message broker that contains global DMR configuration parameters.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- authenticationBasicPassword|||x||x||x authenticationClientCertContent|||x||x||x authenticationClientCertPassword|||x||x|| directOnlyEnabled||x||||| dmrClusterName|x|x||||| nodeName||x||||| tlsServerCertEnforceTrustedCommonNameEnabled||||||x|    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- DmrCluster|authenticationClientCertPassword|authenticationClientCertContent|    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DmrCluster**](DmrCluster.md)| The Cluster object&#x27;s attributes. | 
  **dmrClusterName** | **string**| The name of the Cluster. | 
 **optional** | ***DmrClusterApiUpdateDmrClusterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DmrClusterApiUpdateDmrClusterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterResponse**](DmrClusterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDmrClusterCertMatchingRule**
> DmrClusterCertMatchingRuleResponse UpdateDmrClusterCertMatchingRule(ctx, body, dmrClusterName, ruleName, optional)
Update a Certificate Matching Rule object.

Update a Certificate Matching Rule object. Any attribute missing from the request will be left unchanged.  A Cert Matching Rule is a collection of conditions and attribute filters that all have to be satisfied for certificate to be acceptable as authentication for a given link.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- dmrClusterName|x|x||||| ruleName|x|x|||||    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.28.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DmrClusterCertMatchingRule**](DmrClusterCertMatchingRule.md)| The Certificate Matching Rule object&#x27;s attributes. | 
  **dmrClusterName** | **string**| The name of the Cluster. | 
  **ruleName** | **string**| The name of the rule. | 
 **optional** | ***DmrClusterApiUpdateDmrClusterCertMatchingRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DmrClusterApiUpdateDmrClusterCertMatchingRuleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterCertMatchingRuleResponse**](DmrClusterCertMatchingRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDmrClusterCertMatchingRuleAttributeFilter**
> DmrClusterCertMatchingRuleAttributeFilterResponse UpdateDmrClusterCertMatchingRuleAttributeFilter(ctx, body, dmrClusterName, ruleName, filterName, optional)
Update a Certificate Matching Rule Attribute Filter object.

Update a Certificate Matching Rule Attribute Filter object. Any attribute missing from the request will be left unchanged.  A Cert Matching Rule Attribute Filter compares a link attribute to a string.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- dmrClusterName|x|x||||| filterName|x|x||||| ruleName|x|x|||||    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.28.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DmrClusterCertMatchingRuleAttributeFilter**](DmrClusterCertMatchingRuleAttributeFilter.md)| The Certificate Matching Rule Attribute Filter object&#x27;s attributes. | 
  **dmrClusterName** | **string**| The name of the Cluster. | 
  **ruleName** | **string**| The name of the rule. | 
  **filterName** | **string**| The name of the filter. | 
 **optional** | ***DmrClusterApiUpdateDmrClusterCertMatchingRuleAttributeFilterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DmrClusterApiUpdateDmrClusterCertMatchingRuleAttributeFilterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterCertMatchingRuleAttributeFilterResponse**](DmrClusterCertMatchingRuleAttributeFilterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDmrClusterLink**
> DmrClusterLinkResponse UpdateDmrClusterLink(ctx, body, dmrClusterName, remoteNodeName, optional)
Update a Link object.

Update a Link object. Any attribute missing from the request will be left unchanged.  A Link connects nodes (either within a Cluster or between two different Clusters) and allows them to exchange topology information, subscriptions and data.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Auto-Disable|Deprecated|Opaque :---|:---|:---|:---|:---|:---|:---|:--- authenticationBasicPassword|||x||x||x authenticationScheme|||||x|| dmrClusterName|x|x||||| egressFlowWindowSize|||||x|| initiator|||||x|| remoteNodeName|x|x||||| span|||||x|| transportCompressedEnabled|||||x|| transportTlsEnabled|||||x||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- EventThreshold|clearPercent|setPercent|clearValue, setValue EventThreshold|clearValue|setValue|clearPercent, setPercent EventThreshold|setPercent|clearPercent|clearValue, setValue EventThreshold|setValue|clearValue|clearPercent, setPercent    A SEMP client authorized with a minimum access scope/level of \"global/read-write\" is required to perform this operation.  This has been available since 2.11.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DmrClusterLink**](DmrClusterLink.md)| The Link object&#x27;s attributes. | 
  **dmrClusterName** | **string**| The name of the Cluster. | 
  **remoteNodeName** | **string**| The name of the node at the remote end of the Link. | 
 **optional** | ***DmrClusterApiUpdateDmrClusterLinkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DmrClusterApiUpdateDmrClusterLinkOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **optional.**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterLinkResponse**](DmrClusterLinkResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

