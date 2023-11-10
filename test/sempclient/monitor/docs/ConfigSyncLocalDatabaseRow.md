# ConfigSyncLocalDatabaseRow

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastRequest** | **string** | The last series of commands exchanged between Config Sync sites. Note that this value is only updated during transitions to any syncStatus that is not \&quot;in-sync\&quot;. | [optional] [default to null]
**LastResult** | **string** | The result of the last exchange between Config Sync sites. | [optional] [default to null]
**Name** | **string** | The name is \&quot;site\&quot; when the row type is \&quot;router\&quot;, otherwise it is the Message VPN name. | [optional] [default to null]
**Role** | **string** | The row&#x27;s role relative to the local broker. The allowed values and their meaning are:  &lt;pre&gt; \&quot;unknown\&quot; - The role is unknown. \&quot;leader\&quot; - In HA deployments, the role of the event broker and Message VPNs in the Config Sync database of both HA mates is always Leader. \&quot;follower\&quot; - Only replication-enabled Message VPNs on standby replication mates can have a Follower role. &lt;/pre&gt;  | [optional] [default to null]
**SyncStatus** | **string** | The synchronization status of the row. The allowed values and their meaning are:  &lt;pre&gt; \&quot;unknown\&quot; - The state is unknown. \&quot;in-sync\&quot; - The config data is synchronized between Message VPNs. \&quot;reconciling\&quot; - The config data is reconciling between Message VPNs. \&quot;blocked\&quot; - The config data is blocked from reconciling due to an error. \&quot;out-of-sync\&quot; - The config data is out of sync between Message VPNs. \&quot;down\&quot; - The state is down due to configuration. &lt;/pre&gt;  | [optional] [default to null]
**TimeInState** | **int32** | The amount of time this row has remained in the shown syncStatus, in seconds. | [optional] [default to null]
**Type_** | **string** | The type of the row. Can be one of \&quot;router\&quot; or \&quot;vpn\&quot;. There is one \&quot;router\&quot; row and one row for each configured \&quot;vpn\&quot;. Each row represents a table of information that is synchronized between Config Sync and replication mates. The allowed values and their meaning are:  &lt;pre&gt; \&quot;router\&quot; - The Config Sync database row is for the Router. \&quot;vpn\&quot; - The Config Sync database row is for a Message VPN. &lt;/pre&gt;  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

