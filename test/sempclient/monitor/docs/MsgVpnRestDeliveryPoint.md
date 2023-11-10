# MsgVpnRestDeliveryPoint

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientName** | **string** | The name of the Client for the REST Delivery Point. | [optional] [default to null]
**ClientProfileName** | **string** | The name of the Client Profile for the REST Delivery Point. | [optional] [default to null]
**Enabled** | [****bool**](*bool.md) | Indicates whether the REST Delivery Point is enabled. | [optional] [default to null]
**LastFailureReason** | **string** | The reason for the last REST Delivery Point failure. | [optional] [default to null]
**LastFailureTime** | **int32** | The timestamp of the last REST Delivery Point failure. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). | [optional] [default to null]
**MsgVpnName** | **string** | The name of the Message VPN. | [optional] [default to null]
**RestDeliveryPointName** | **string** | The name of the REST Delivery Point. | [optional] [default to null]
**Service** | **string** | The name of the service that this REST Delivery Point connects to. Internally the broker does not use this value; it is informational only. Available since 2.19. | [optional] [default to null]
**TimeConnectionsBlocked** | **int64** | The percentage of time the REST Delivery Point connections are blocked from transmitting data. | [optional] [default to null]
**Up** | [****bool**](*bool.md) | Indicates whether the operational state of the REST Delivery Point is up. | [optional] [default to null]
**Vendor** | **string** | The name of the vendor that this REST Delivery Point connects to. Internally the broker does not use this value; it is informational only. Available since 2.19. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

