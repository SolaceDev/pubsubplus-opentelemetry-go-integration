# MsgVpnTransactionConsumerMsg

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndpointName** | **string** | The name of the Queue or Topic Endpoint source. | [optional] [default to null]
**EndpointType** | **string** | The type of endpoint source. The allowed values and their meaning are:  &lt;pre&gt; \&quot;queue\&quot; - The Message is from a Queue. \&quot;topic-endpoint\&quot; - The Message is from a Topic Endpoint. &lt;/pre&gt;  | [optional] [default to null]
**MsgId** | **int64** | The identifier (ID) of the Message. | [optional] [default to null]
**MsgVpnName** | **string** | The name of the Message VPN. | [optional] [default to null]
**ReplicationGroupMsgId** | **string** | An ID that uniquely identifies this message within this replication group. Available since 2.21. | [optional] [default to null]
**Xid** | **string** | The identifier (ID) of the Transaction. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

