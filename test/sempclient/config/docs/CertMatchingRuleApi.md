# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/config*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMsgVpnCertMatchingRule**](CertMatchingRuleApi.md#CreateMsgVpnCertMatchingRule) | **Post** /msgVpns/{msgVpnName}/certMatchingRules | Create a Certificate Matching Rule object.
[**CreateMsgVpnCertMatchingRuleAttributeFilter**](CertMatchingRuleApi.md#CreateMsgVpnCertMatchingRuleAttributeFilter) | **Post** /msgVpns/{msgVpnName}/certMatchingRules/{ruleName}/attributeFilters | Create a Certificate Matching Rule Attribute Filter object.
[**CreateMsgVpnCertMatchingRuleCondition**](CertMatchingRuleApi.md#CreateMsgVpnCertMatchingRuleCondition) | **Post** /msgVpns/{msgVpnName}/certMatchingRules/{ruleName}/conditions | Create a Certificate Matching Rule Condition object.
[**DeleteMsgVpnCertMatchingRule**](CertMatchingRuleApi.md#DeleteMsgVpnCertMatchingRule) | **Delete** /msgVpns/{msgVpnName}/certMatchingRules/{ruleName} | Delete a Certificate Matching Rule object.
[**DeleteMsgVpnCertMatchingRuleAttributeFilter**](CertMatchingRuleApi.md#DeleteMsgVpnCertMatchingRuleAttributeFilter) | **Delete** /msgVpns/{msgVpnName}/certMatchingRules/{ruleName}/attributeFilters/{filterName} | Delete a Certificate Matching Rule Attribute Filter object.
[**DeleteMsgVpnCertMatchingRuleCondition**](CertMatchingRuleApi.md#DeleteMsgVpnCertMatchingRuleCondition) | **Delete** /msgVpns/{msgVpnName}/certMatchingRules/{ruleName}/conditions/{source} | Delete a Certificate Matching Rule Condition object.
[**GetMsgVpnCertMatchingRule**](CertMatchingRuleApi.md#GetMsgVpnCertMatchingRule) | **Get** /msgVpns/{msgVpnName}/certMatchingRules/{ruleName} | Get a Certificate Matching Rule object.
[**GetMsgVpnCertMatchingRuleAttributeFilter**](CertMatchingRuleApi.md#GetMsgVpnCertMatchingRuleAttributeFilter) | **Get** /msgVpns/{msgVpnName}/certMatchingRules/{ruleName}/attributeFilters/{filterName} | Get a Certificate Matching Rule Attribute Filter object.
[**GetMsgVpnCertMatchingRuleAttributeFilters**](CertMatchingRuleApi.md#GetMsgVpnCertMatchingRuleAttributeFilters) | **Get** /msgVpns/{msgVpnName}/certMatchingRules/{ruleName}/attributeFilters | Get a list of Certificate Matching Rule Attribute Filter objects.
[**GetMsgVpnCertMatchingRuleCondition**](CertMatchingRuleApi.md#GetMsgVpnCertMatchingRuleCondition) | **Get** /msgVpns/{msgVpnName}/certMatchingRules/{ruleName}/conditions/{source} | Get a Certificate Matching Rule Condition object.
[**GetMsgVpnCertMatchingRuleConditions**](CertMatchingRuleApi.md#GetMsgVpnCertMatchingRuleConditions) | **Get** /msgVpns/{msgVpnName}/certMatchingRules/{ruleName}/conditions | Get a list of Certificate Matching Rule Condition objects.
[**GetMsgVpnCertMatchingRules**](CertMatchingRuleApi.md#GetMsgVpnCertMatchingRules) | **Get** /msgVpns/{msgVpnName}/certMatchingRules | Get a list of Certificate Matching Rule objects.
[**ReplaceMsgVpnCertMatchingRule**](CertMatchingRuleApi.md#ReplaceMsgVpnCertMatchingRule) | **Put** /msgVpns/{msgVpnName}/certMatchingRules/{ruleName} | Replace a Certificate Matching Rule object.
[**ReplaceMsgVpnCertMatchingRuleAttributeFilter**](CertMatchingRuleApi.md#ReplaceMsgVpnCertMatchingRuleAttributeFilter) | **Put** /msgVpns/{msgVpnName}/certMatchingRules/{ruleName}/attributeFilters/{filterName} | Replace a Certificate Matching Rule Attribute Filter object.
[**UpdateMsgVpnCertMatchingRule**](CertMatchingRuleApi.md#UpdateMsgVpnCertMatchingRule) | **Patch** /msgVpns/{msgVpnName}/certMatchingRules/{ruleName} | Update a Certificate Matching Rule object.
[**UpdateMsgVpnCertMatchingRuleAttributeFilter**](CertMatchingRuleApi.md#UpdateMsgVpnCertMatchingRuleAttributeFilter) | **Patch** /msgVpns/{msgVpnName}/certMatchingRules/{ruleName}/attributeFilters/{filterName} | Update a Certificate Matching Rule Attribute Filter object.

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
 **optional** | ***CertMatchingRuleApiCreateMsgVpnCertMatchingRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CertMatchingRuleApiCreateMsgVpnCertMatchingRuleOpts struct
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
 **optional** | ***CertMatchingRuleApiCreateMsgVpnCertMatchingRuleAttributeFilterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CertMatchingRuleApiCreateMsgVpnCertMatchingRuleAttributeFilterOpts struct
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
 **optional** | ***CertMatchingRuleApiCreateMsgVpnCertMatchingRuleConditionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CertMatchingRuleApiCreateMsgVpnCertMatchingRuleConditionOpts struct
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
 **optional** | ***CertMatchingRuleApiGetMsgVpnCertMatchingRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CertMatchingRuleApiGetMsgVpnCertMatchingRuleOpts struct
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
 **optional** | ***CertMatchingRuleApiGetMsgVpnCertMatchingRuleAttributeFilterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CertMatchingRuleApiGetMsgVpnCertMatchingRuleAttributeFilterOpts struct
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
 **optional** | ***CertMatchingRuleApiGetMsgVpnCertMatchingRuleAttributeFiltersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CertMatchingRuleApiGetMsgVpnCertMatchingRuleAttributeFiltersOpts struct
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
 **optional** | ***CertMatchingRuleApiGetMsgVpnCertMatchingRuleConditionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CertMatchingRuleApiGetMsgVpnCertMatchingRuleConditionOpts struct
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
 **optional** | ***CertMatchingRuleApiGetMsgVpnCertMatchingRuleConditionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CertMatchingRuleApiGetMsgVpnCertMatchingRuleConditionsOpts struct
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
 **optional** | ***CertMatchingRuleApiGetMsgVpnCertMatchingRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CertMatchingRuleApiGetMsgVpnCertMatchingRulesOpts struct
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
 **optional** | ***CertMatchingRuleApiReplaceMsgVpnCertMatchingRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CertMatchingRuleApiReplaceMsgVpnCertMatchingRuleOpts struct
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
 **optional** | ***CertMatchingRuleApiReplaceMsgVpnCertMatchingRuleAttributeFilterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CertMatchingRuleApiReplaceMsgVpnCertMatchingRuleAttributeFilterOpts struct
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
 **optional** | ***CertMatchingRuleApiUpdateMsgVpnCertMatchingRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CertMatchingRuleApiUpdateMsgVpnCertMatchingRuleOpts struct
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
 **optional** | ***CertMatchingRuleApiUpdateMsgVpnCertMatchingRuleAttributeFilterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CertMatchingRuleApiUpdateMsgVpnCertMatchingRuleAttributeFilterOpts struct
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

