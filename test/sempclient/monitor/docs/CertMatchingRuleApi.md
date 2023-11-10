# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/monitor*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMsgVpnCertMatchingRule**](CertMatchingRuleApi.md#GetMsgVpnCertMatchingRule) | **Get** /msgVpns/{msgVpnName}/certMatchingRules/{ruleName} | Get a Certificate Matching Rule object.
[**GetMsgVpnCertMatchingRuleAttributeFilter**](CertMatchingRuleApi.md#GetMsgVpnCertMatchingRuleAttributeFilter) | **Get** /msgVpns/{msgVpnName}/certMatchingRules/{ruleName}/attributeFilters/{filterName} | Get a Certificate Matching Rule Attribute Filter object.
[**GetMsgVpnCertMatchingRuleAttributeFilters**](CertMatchingRuleApi.md#GetMsgVpnCertMatchingRuleAttributeFilters) | **Get** /msgVpns/{msgVpnName}/certMatchingRules/{ruleName}/attributeFilters | Get a list of Certificate Matching Rule Attribute Filter objects.
[**GetMsgVpnCertMatchingRuleCondition**](CertMatchingRuleApi.md#GetMsgVpnCertMatchingRuleCondition) | **Get** /msgVpns/{msgVpnName}/certMatchingRules/{ruleName}/conditions/{source} | Get a Certificate Matching Rule Condition object.
[**GetMsgVpnCertMatchingRuleConditions**](CertMatchingRuleApi.md#GetMsgVpnCertMatchingRuleConditions) | **Get** /msgVpns/{msgVpnName}/certMatchingRules/{ruleName}/conditions | Get a list of Certificate Matching Rule Condition objects.
[**GetMsgVpnCertMatchingRules**](CertMatchingRuleApi.md#GetMsgVpnCertMatchingRules) | **Get** /msgVpns/{msgVpnName}/certMatchingRules | Get a list of Certificate Matching Rule objects.

# **GetMsgVpnCertMatchingRule**
> MsgVpnCertMatchingRuleResponse GetMsgVpnCertMatchingRule(ctx, msgVpnName, ruleName, optional)
Get a Certificate Matching Rule object.

Get a Certificate Matching Rule object.  A Cert Matching Rule is a collection of conditions and attribute filters that all have to be satisfied for certificate to be acceptable as authentication for a given username.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| ruleName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.27.

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

Get a Certificate Matching Rule Attribute Filter object.  A Cert Matching Rule Attribute Filter compares a username attribute to a string.   Attribute|Identifying|Deprecated :---|:---:|:---: filterName|x| msgVpnName|x| ruleName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.28.

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

Get a list of Certificate Matching Rule Attribute Filter objects.  A Cert Matching Rule Attribute Filter compares a username attribute to a string.   Attribute|Identifying|Deprecated :---|:---:|:---: filterName|x| msgVpnName|x| ruleName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.28.

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

Get a Certificate Matching Rule Condition object.  A Cert Matching Rule Condition compares data extracted from a certificate to a username attribute or an expression.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| ruleName|x| source|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.27.

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

Get a list of Certificate Matching Rule Condition objects.  A Cert Matching Rule Condition compares data extracted from a certificate to a username attribute or an expression.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| ruleName|x| source|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.27.

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

Get a list of Certificate Matching Rule objects.  A Cert Matching Rule is a collection of conditions and attribute filters that all have to be satisfied for certificate to be acceptable as authentication for a given username.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| ruleName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.27.

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

