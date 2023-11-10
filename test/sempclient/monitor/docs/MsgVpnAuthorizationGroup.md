# MsgVpnAuthorizationGroup

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AclProfileName** | **string** | The ACL Profile of the Authorization Group. | [optional] [default to null]
**AuthorizationGroupName** | **string** | The name of the Authorization Group. For LDAP groups, special care is needed if the group name contains special characters such as &#x27;#&#x27;, &#x27;+&#x27;, &#x27;;&#x27;, &#x27;&#x3D;&#x27; as the value of the group name returned from the LDAP server might prepend those characters with &#x27;\\&#x27;. For example a group name called &#x27;test#,lab,com&#x27; will be returned from the LDAP server as &#x27;test\\#,lab,com&#x27;. | [optional] [default to null]
**ClientProfileName** | **string** | The Client Profile of the Authorization Group. | [optional] [default to null]
**Enabled** | [****bool**](*bool.md) | Indicates whether the Authorization Group is enabled. | [optional] [default to null]
**MsgVpnName** | **string** | The name of the Message VPN. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

