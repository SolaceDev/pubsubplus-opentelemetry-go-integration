# MsgVpnRestDeliveryPointQueueBinding

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewayReplaceTargetAuthorityEnabled** | [****bool**](*bool.md) | Indicates whether the authority for the request-target is replaced with that configured for the REST Consumer remote. | [optional] [default to null]
**LastFailureReason** | **string** | The reason for the last REST Delivery Point queue binding failure. | [optional] [default to null]
**LastFailureTime** | **int32** | The timestamp of the last REST Delivery Point queue binding failure. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). | [optional] [default to null]
**MsgVpnName** | **string** | The name of the Message VPN. | [optional] [default to null]
**PostRequestTarget** | **string** | The request-target string being used when sending requests to a REST Consumer. | [optional] [default to null]
**QueueBindingName** | **string** | The name of a queue in the Message VPN. | [optional] [default to null]
**RequestTargetEvaluation** | **string** | The type of evaluation to perform on the request target. The allowed values and their meaning are:  &lt;pre&gt; \&quot;none\&quot; - Do not evaluate substitution expressions on the request target. \&quot;substitution-expressions\&quot; - Evaluate substitution expressions on the request target. &lt;/pre&gt;  Available since 2.23. | [optional] [default to null]
**RestDeliveryPointName** | **string** | The name of the REST Delivery Point. | [optional] [default to null]
**Up** | [****bool**](*bool.md) | Indicates whether the operational state of the REST Delivery Point queue binding is up. | [optional] [default to null]
**Uptime** | **int64** | The amount of time in seconds since the REST Delivery Point queue binding was up. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

