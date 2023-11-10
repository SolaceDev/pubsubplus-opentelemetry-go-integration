# MsgVpnMqttSession

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MqttSessionClientId** | **string** | The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet. | [optional] [default to null]
**MqttSessionVirtualRouter** | **string** | The virtual router of the MQTT Session. The allowed values and their meaning are:  &lt;pre&gt; \&quot;primary\&quot; - The MQTT Session belongs to the primary virtual router. \&quot;backup\&quot; - The MQTT Session belongs to the backup virtual router. \&quot;auto\&quot; - The MQTT Session is automatically assigned a virtual router at creation, depending on the broker&#x27;s active-standby role. &lt;/pre&gt;  | [optional] [default to null]
**MsgVpnName** | **string** | The name of the Message VPN. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

