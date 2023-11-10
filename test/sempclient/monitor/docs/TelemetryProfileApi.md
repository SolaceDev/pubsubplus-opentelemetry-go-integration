# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/monitor*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMsgVpnTelemetryProfile**](TelemetryProfileApi.md#GetMsgVpnTelemetryProfile) | **Get** /msgVpns/{msgVpnName}/telemetryProfiles/{telemetryProfileName} | Get a Telemetry Profile object.
[**GetMsgVpnTelemetryProfileReceiverAclConnectException**](TelemetryProfileApi.md#GetMsgVpnTelemetryProfileReceiverAclConnectException) | **Get** /msgVpns/{msgVpnName}/telemetryProfiles/{telemetryProfileName}/receiverAclConnectExceptions/{receiverAclConnectExceptionAddress} | Get a Receiver ACL Connect Exception object.
[**GetMsgVpnTelemetryProfileReceiverAclConnectExceptions**](TelemetryProfileApi.md#GetMsgVpnTelemetryProfileReceiverAclConnectExceptions) | **Get** /msgVpns/{msgVpnName}/telemetryProfiles/{telemetryProfileName}/receiverAclConnectExceptions | Get a list of Receiver ACL Connect Exception objects.
[**GetMsgVpnTelemetryProfileTraceFilter**](TelemetryProfileApi.md#GetMsgVpnTelemetryProfileTraceFilter) | **Get** /msgVpns/{msgVpnName}/telemetryProfiles/{telemetryProfileName}/traceFilters/{traceFilterName} | Get a Trace Filter object.
[**GetMsgVpnTelemetryProfileTraceFilterSubscription**](TelemetryProfileApi.md#GetMsgVpnTelemetryProfileTraceFilterSubscription) | **Get** /msgVpns/{msgVpnName}/telemetryProfiles/{telemetryProfileName}/traceFilters/{traceFilterName}/subscriptions/{subscription},{subscriptionSyntax} | Get a Telemetry Trace Filter Subscription object.
[**GetMsgVpnTelemetryProfileTraceFilterSubscriptions**](TelemetryProfileApi.md#GetMsgVpnTelemetryProfileTraceFilterSubscriptions) | **Get** /msgVpns/{msgVpnName}/telemetryProfiles/{telemetryProfileName}/traceFilters/{traceFilterName}/subscriptions | Get a list of Telemetry Trace Filter Subscription objects.
[**GetMsgVpnTelemetryProfileTraceFilters**](TelemetryProfileApi.md#GetMsgVpnTelemetryProfileTraceFilters) | **Get** /msgVpns/{msgVpnName}/telemetryProfiles/{telemetryProfileName}/traceFilters | Get a list of Trace Filter objects.
[**GetMsgVpnTelemetryProfiles**](TelemetryProfileApi.md#GetMsgVpnTelemetryProfiles) | **Get** /msgVpns/{msgVpnName}/telemetryProfiles | Get a list of Telemetry Profile objects.

# **GetMsgVpnTelemetryProfile**
> MsgVpnTelemetryProfileResponse GetMsgVpnTelemetryProfile(ctx, msgVpnName, telemetryProfileName, optional)
Get a Telemetry Profile object.

Get a Telemetry Profile object.  Using the Telemetry Profile allows trace spans to be generated as messages are processed by the broker. The generated spans are stored persistently on the broker and may be consumed by the Solace receiver component of an OpenTelemetry Collector.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| telemetryProfileName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.31.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **telemetryProfileName** | **string**| The name of the Telemetry Profile. | 
 **optional** | ***TelemetryProfileApiGetMsgVpnTelemetryProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TelemetryProfileApiGetMsgVpnTelemetryProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTelemetryProfileResponse**](MsgVpnTelemetryProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnTelemetryProfileReceiverAclConnectException**
> MsgVpnTelemetryProfileReceiverAclConnectExceptionResponse GetMsgVpnTelemetryProfileReceiverAclConnectException(ctx, msgVpnName, telemetryProfileName, receiverAclConnectExceptionAddress, optional)
Get a Receiver ACL Connect Exception object.

Get a Receiver ACL Connect Exception object.  A Receiver ACL Connect Exception is an exception to the default action to take when a receiver connects to the broker. Exceptions must be expressed as an IP address/netmask in CIDR form.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| receiverAclConnectExceptionAddress|x| telemetryProfileName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.31.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **telemetryProfileName** | **string**| The name of the Telemetry Profile. | 
  **receiverAclConnectExceptionAddress** | **string**| The IP address/netmask of the receiver connect exception in CIDR form. | 
 **optional** | ***TelemetryProfileApiGetMsgVpnTelemetryProfileReceiverAclConnectExceptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TelemetryProfileApiGetMsgVpnTelemetryProfileReceiverAclConnectExceptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTelemetryProfileReceiverAclConnectExceptionResponse**](MsgVpnTelemetryProfileReceiverAclConnectExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnTelemetryProfileReceiverAclConnectExceptions**
> MsgVpnTelemetryProfileReceiverAclConnectExceptionsResponse GetMsgVpnTelemetryProfileReceiverAclConnectExceptions(ctx, msgVpnName, telemetryProfileName, optional)
Get a list of Receiver ACL Connect Exception objects.

Get a list of Receiver ACL Connect Exception objects.  A Receiver ACL Connect Exception is an exception to the default action to take when a receiver connects to the broker. Exceptions must be expressed as an IP address/netmask in CIDR form.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| receiverAclConnectExceptionAddress|x| telemetryProfileName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.31.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **telemetryProfileName** | **string**| The name of the Telemetry Profile. | 
 **optional** | ***TelemetryProfileApiGetMsgVpnTelemetryProfileReceiverAclConnectExceptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TelemetryProfileApiGetMsgVpnTelemetryProfileReceiverAclConnectExceptionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTelemetryProfileReceiverAclConnectExceptionsResponse**](MsgVpnTelemetryProfileReceiverAclConnectExceptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnTelemetryProfileTraceFilter**
> MsgVpnTelemetryProfileTraceFilterResponse GetMsgVpnTelemetryProfileTraceFilter(ctx, msgVpnName, telemetryProfileName, traceFilterName, optional)
Get a Trace Filter object.

Get a Trace Filter object.  A Trace Filter controls which messages received by the broker will be traced. If an incoming message matches an enabled tracing filter's subscription, the message will be traced as it passes through the broker.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| telemetryProfileName|x| traceFilterName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.31.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **telemetryProfileName** | **string**| The name of the Telemetry Profile. | 
  **traceFilterName** | **string**| A name used to identify the trace filter. Consider a name that describes the subscriptions contained within the filter, such as the name of the application and/or the scenario in which the trace filter might be enabled, such as \&quot;appNameDebug\&quot;. | 
 **optional** | ***TelemetryProfileApiGetMsgVpnTelemetryProfileTraceFilterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TelemetryProfileApiGetMsgVpnTelemetryProfileTraceFilterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTelemetryProfileTraceFilterResponse**](MsgVpnTelemetryProfileTraceFilterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnTelemetryProfileTraceFilterSubscription**
> MsgVpnTelemetryProfileTraceFilterSubscriptionResponse GetMsgVpnTelemetryProfileTraceFilterSubscription(ctx, msgVpnName, telemetryProfileName, traceFilterName, subscription, subscriptionSyntax, optional)
Get a Telemetry Trace Filter Subscription object.

Get a Telemetry Trace Filter Subscription object.  Trace filter subscriptions control which messages will be attracted by the tracing filter.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| subscription|x| subscriptionSyntax|x| telemetryProfileName|x| traceFilterName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.31.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **telemetryProfileName** | **string**| The name of the Telemetry Profile. | 
  **traceFilterName** | **string**| A name used to identify the trace filter. Consider a name that describes the subscriptions contained within the filter, such as the name of the application and/or the scenario in which the trace filter might be enabled, such as \&quot;appNameDebug\&quot;. | 
  **subscription** | **string**| Messages matching this subscription will follow this filter&#x27;s configuration. | 
  **subscriptionSyntax** | **string**| The syntax of the trace filter subscription. | 
 **optional** | ***TelemetryProfileApiGetMsgVpnTelemetryProfileTraceFilterSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TelemetryProfileApiGetMsgVpnTelemetryProfileTraceFilterSubscriptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTelemetryProfileTraceFilterSubscriptionResponse**](MsgVpnTelemetryProfileTraceFilterSubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnTelemetryProfileTraceFilterSubscriptions**
> MsgVpnTelemetryProfileTraceFilterSubscriptionsResponse GetMsgVpnTelemetryProfileTraceFilterSubscriptions(ctx, msgVpnName, telemetryProfileName, traceFilterName, optional)
Get a list of Telemetry Trace Filter Subscription objects.

Get a list of Telemetry Trace Filter Subscription objects.  Trace filter subscriptions control which messages will be attracted by the tracing filter.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| subscription|x| subscriptionSyntax|x| telemetryProfileName|x| traceFilterName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.31.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **telemetryProfileName** | **string**| The name of the Telemetry Profile. | 
  **traceFilterName** | **string**| A name used to identify the trace filter. Consider a name that describes the subscriptions contained within the filter, such as the name of the application and/or the scenario in which the trace filter might be enabled, such as \&quot;appNameDebug\&quot;. | 
 **optional** | ***TelemetryProfileApiGetMsgVpnTelemetryProfileTraceFilterSubscriptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TelemetryProfileApiGetMsgVpnTelemetryProfileTraceFilterSubscriptionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTelemetryProfileTraceFilterSubscriptionsResponse**](MsgVpnTelemetryProfileTraceFilterSubscriptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnTelemetryProfileTraceFilters**
> MsgVpnTelemetryProfileTraceFiltersResponse GetMsgVpnTelemetryProfileTraceFilters(ctx, msgVpnName, telemetryProfileName, optional)
Get a list of Trace Filter objects.

Get a list of Trace Filter objects.  A Trace Filter controls which messages received by the broker will be traced. If an incoming message matches an enabled tracing filter's subscription, the message will be traced as it passes through the broker.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| telemetryProfileName|x| traceFilterName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.31.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
  **telemetryProfileName** | **string**| The name of the Telemetry Profile. | 
 **optional** | ***TelemetryProfileApiGetMsgVpnTelemetryProfileTraceFiltersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TelemetryProfileApiGetMsgVpnTelemetryProfileTraceFiltersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTelemetryProfileTraceFiltersResponse**](MsgVpnTelemetryProfileTraceFiltersResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMsgVpnTelemetryProfiles**
> MsgVpnTelemetryProfilesResponse GetMsgVpnTelemetryProfiles(ctx, msgVpnName, optional)
Get a list of Telemetry Profile objects.

Get a list of Telemetry Profile objects.  Using the Telemetry Profile allows trace spans to be generated as messages are processed by the broker. The generated spans are stored persistently on the broker and may be consumed by the Solace receiver component of an OpenTelemetry Collector.   Attribute|Identifying|Deprecated :---|:---:|:---: msgVpnName|x| telemetryProfileName|x|    A SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.  This has been available since 2.31.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msgVpnName** | **string**| The name of the Message VPN. | 
 **optional** | ***TelemetryProfileApiGetMsgVpnTelemetryProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TelemetryProfileApiGetMsgVpnTelemetryProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **optional.String**| The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | [**optional.Interface of []string**](string.md)| Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnTelemetryProfilesResponse**](MsgVpnTelemetryProfilesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

