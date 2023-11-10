# {{classname}}

All URIs are relative to *http://www.solace.com/SEMP/v2/config*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSystemInformation**](SystemInformationApi.md#GetSystemInformation) | **Get** /systemInformation | Get a System Information object.

# **GetSystemInformation**
> SystemInformationResponse GetSystemInformation(ctx, optional)
Get a System Information object.

Get a System Information object.  The System Information object provides metadata about the SEMP API.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: platform|||x| sempVersion|||x|    A SEMP client authorized with a minimum access scope/level of \"global/none\" is required to perform this operation.  This has been deprecated since 2.2. /systemInformation was replaced by /about/api.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemInformationApiGetSystemInformationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemInformationApiGetSystemInformationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **opaquePassword** | **optional.String**| Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See the documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | [**optional.Interface of []string**](string.md)| Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**SystemInformationResponse**](SystemInformationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

