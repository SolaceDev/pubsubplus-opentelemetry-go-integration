# MsgVpnBridgeLocalSubscription

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BridgeName** | **string** | The name of the Bridge. | [optional] [default to null]
**BridgeVirtualRouter** | **string** | The virtual router of the Bridge. The allowed values and their meaning are:  &lt;pre&gt; \&quot;primary\&quot; - The Bridge is used for the primary virtual router. \&quot;backup\&quot; - The Bridge is used for the backup virtual router. \&quot;auto\&quot; - The Bridge is automatically assigned a virtual router at creation, depending on the broker&#x27;s active-standby role. &lt;/pre&gt;  | [optional] [default to null]
**DtoPriority** | **string** | The priority of the Bridge local subscription topic for receiving deliver-to-one (DTO) messages. The allowed values and their meaning are:  &lt;pre&gt; \&quot;p1\&quot; - The 1st or highest priority. \&quot;p2\&quot; - The 2nd highest priority. \&quot;p3\&quot; - The 3rd highest priority. \&quot;p4\&quot; - The 4th highest priority. \&quot;da\&quot; - Ignore priority and deliver always. &lt;/pre&gt;  | [optional] [default to null]
**LocalSubscriptionTopic** | **string** | The topic of the Bridge local subscription. | [optional] [default to null]
**MsgVpnName** | **string** | The name of the Message VPN. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

