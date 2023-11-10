# MsgVpnClientRxFlow

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientName** | **string** | The name of the Client. | [optional] [default to null]
**ConnectTime** | **int32** | The timestamp of when the Flow from the Client connected. | [optional] [default to null]
**DestinationGroupErrorDiscardedMsgCount** | **int64** | The number of guaranteed messages from the Flow discarded due to a destination group error. | [optional] [default to null]
**DuplicateDiscardedMsgCount** | **int64** | The number of guaranteed messages from the Flow discarded due to being a duplicate. | [optional] [default to null]
**EndpointDisabledDiscardedMsgCount** | **int64** | The number of guaranteed messages from the Flow discarded due to an eligible endpoint destination being disabled. | [optional] [default to null]
**EndpointUsageExceededDiscardedMsgCount** | **int64** | The number of guaranteed messages from the Flow discarded due to an eligible endpoint destination having its maximum message spool usage exceeded. | [optional] [default to null]
**ErroredDiscardedMsgCount** | **int64** | The number of guaranteed messages from the Flow discarded due to errors being detected. | [optional] [default to null]
**FlowId** | **int64** | The identifier (ID) of the flow. | [optional] [default to null]
**FlowName** | **string** | The name of the Flow. | [optional] [default to null]
**GuaranteedMsgCount** | **int64** | The number of guaranteed messages from the Flow. | [optional] [default to null]
**LastRxMsgId** | **int64** | The identifier (ID) of the last message received on the Flow. | [optional] [default to null]
**LocalMsgCountExceededDiscardedMsgCount** | **int64** | The number of guaranteed messages from the Flow discarded due to the maximum number of messages allowed on the broker being exceeded. | [optional] [default to null]
**LowPriorityMsgCongestionDiscardedMsgCount** | **int64** | The number of guaranteed messages from the Flow discarded due to congestion of low priority messages. | [optional] [default to null]
**MaxMsgSizeExceededDiscardedMsgCount** | **int64** | The number of guaranteed messages from the Flow discarded due to the maximum allowed message size being exceeded. | [optional] [default to null]
**MsgVpnName** | **string** | The name of the Message VPN. | [optional] [default to null]
**NoEligibleDestinationsDiscardedMsgCount** | **int64** | The number of guaranteed messages from the Flow discarded due to there being no eligible endpoint destination. | [optional] [default to null]
**NoLocalDeliveryDiscardedMsgCount** | **int64** | The number of guaranteed messages from the Flow discarded due to no local delivery being requested. | [optional] [default to null]
**NotCompatibleWithForwardingModeDiscardedMsgCount** | **int64** | The number of guaranteed messages from the Flow discarded due to being incompatible with the forwarding mode of an eligible endpoint destination. | [optional] [default to null]
**OutOfOrderDiscardedMsgCount** | **int64** | The number of guaranteed messages from the Flow discarded due to being received out of order. | [optional] [default to null]
**PublishAclDeniedDiscardedMsgCount** | **int64** | The number of guaranteed messages from the Flow discarded due to being denied by the access control list (ACL) profile for the published topic. | [optional] [default to null]
**PublisherId** | **int64** | The identifier (ID) of the publisher for the Flow. | [optional] [default to null]
**QueueNotFoundDiscardedMsgCount** | **int64** | The number of guaranteed messages from the Flow discarded due to the destination queue not being found. | [optional] [default to null]
**ReplicationStandbyDiscardedMsgCount** | **int64** | The number of guaranteed messages from the Flow discarded due to the Message VPN being in the replication standby state. | [optional] [default to null]
**SessionName** | **string** | The name of the transacted session on the Flow. | [optional] [default to null]
**SmfTtlExceededDiscardedMsgCount** | **int64** | The number of guaranteed messages from the Flow discarded due to the message time-to-live (TTL) count being exceeded. The message TTL count is the maximum number of times the message can cross a bridge between Message VPNs. | [optional] [default to null]
**SpoolFileLimitExceededDiscardedMsgCount** | **int64** | The number of guaranteed messages from the Flow discarded due to all available message spool file resources being used. | [optional] [default to null]
**SpoolNotReadyDiscardedMsgCount** | **int64** | The number of guaranteed messages from the Flow discarded due to the message spool being not ready. | [optional] [default to null]
**SpoolToAdbFailDiscardedMsgCount** | **int64** | The number of guaranteed messages from the Flow discarded due to a failure while spooling to the Assured Delivery Blade (ADB). | [optional] [default to null]
**SpoolToDiskFailDiscardedMsgCount** | **int64** | The number of guaranteed messages from the Flow discarded due to a failure while spooling to the disk. | [optional] [default to null]
**SpoolUsageExceededDiscardedMsgCount** | **int64** | The number of guaranteed messages from the Flow discarded due to the maximum message spool usage being exceeded. | [optional] [default to null]
**SyncReplicationIneligibleDiscardedMsgCount** | **int64** | The number of guaranteed messages from the Flow discarded due to synchronous replication being ineligible. | [optional] [default to null]
**UserProfileDeniedGuaranteedDiscardedMsgCount** | **int64** | The number of guaranteed messages from the Flow discarded due to being denied by the client profile. | [optional] [default to null]
**WindowSize** | **int32** | The size of the window used for guaranteed messages sent on the Flow, in messages. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

