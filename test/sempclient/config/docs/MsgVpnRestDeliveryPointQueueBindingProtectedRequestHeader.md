# MsgVpnRestDeliveryPointQueueBindingProtectedRequestHeader

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HeaderName** | **string** | The name of the protected HTTP request header. | [optional] [default to null]
**HeaderValue** | **string** | The value of the protected HTTP request header. Unlike a non-protected request header, this value cannot be displayed after it is set, and does not support substitution expressions. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is &#x60;\&quot;\&quot;&#x60;. | [optional] [default to null]
**MsgVpnName** | **string** | The name of the Message VPN. | [optional] [default to null]
**QueueBindingName** | **string** | The name of a queue in the Message VPN. | [optional] [default to null]
**RestDeliveryPointName** | **string** | The name of the REST Delivery Point. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

