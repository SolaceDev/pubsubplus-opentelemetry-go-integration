# MsgVpnConfigSyncRemoteNode

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastMsgRxTime** | **int32** | The amount of time in seconds since the last message was received from the config sync table of the remote Message VPN. Deprecated since 2.22. This attribute has been deprecated. | [optional] [default to null]
**MsgVpnName** | **string** | The name of the Message VPN. Deprecated since 2.22. This attribute has been deprecated. | [optional] [default to null]
**RemoteNodeName** | **string** | The name of the Config Sync Remote Node. Deprecated since 2.22. This attribute has been deprecated. | [optional] [default to null]
**Role** | **string** | The role of the config sync table of the remote Message VPN. The allowed values and their meaning are:  &lt;pre&gt; \&quot;unknown\&quot; - The role is unknown. \&quot;primary\&quot; - Acts as the primary source of config data. \&quot;replica\&quot; - Acts as a replica of the primary config data. &lt;/pre&gt;  Deprecated since 2.22. This attribute has been deprecated. | [optional] [default to null]
**Stale** | [****bool**](*bool.md) | Indicates whether the config sync table of the remote Message VPN is stale. Deprecated since 2.22. This attribute has been deprecated. | [optional] [default to null]
**State** | **string** | The state of the config sync table of the remote Message VPN. The allowed values and their meaning are:  &lt;pre&gt; \&quot;unknown\&quot; - The state is unknown. \&quot;in-sync\&quot; - The config data is synchronized between Message VPNs. \&quot;reconciling\&quot; - The config data is reconciling between Message VPNs. \&quot;blocked\&quot; - The config data is blocked from reconciling due to an error. \&quot;out-of-sync\&quot; - The config data is out of sync between Message VPNs. \&quot;down\&quot; - The state is down due to configuration. &lt;/pre&gt;  Deprecated since 2.22. This attribute has been deprecated. | [optional] [default to null]
**TimeInState** | **int32** | The amount of time in seconds the config sync table of the remote Message VPN has been in the current state. Deprecated since 2.22. This attribute has been deprecated. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

