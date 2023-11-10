# MsgVpnReplicatedTopic

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MsgVpnName** | **string** | The name of the Message VPN. | [optional] [default to null]
**ReplicatedTopic** | **string** | The topic for applying replication. Published messages matching this topic will be replicated to the standby site. | [optional] [default to null]
**ReplicationMode** | **string** | The replication mode for the Replicated Topic. The allowed values and their meaning are:  &lt;pre&gt; \&quot;sync\&quot; - Messages are acknowledged when replicated (spooled remotely). \&quot;async\&quot; - Messages are acknowledged when pending replication (spooled locally). &lt;/pre&gt;  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

