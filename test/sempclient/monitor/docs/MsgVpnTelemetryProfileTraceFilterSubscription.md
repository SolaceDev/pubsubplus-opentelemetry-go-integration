# MsgVpnTelemetryProfileTraceFilterSubscription

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MsgVpnName** | **string** | The name of the Message VPN. | [optional] [default to null]
**Subscription** | **string** | Messages matching this subscription will follow this filter&#x27;s configuration. | [optional] [default to null]
**SubscriptionSyntax** | **string** | The syntax of the trace filter subscription. The allowed values and their meaning are:  &lt;pre&gt; \&quot;smf\&quot; - Subscription uses SMF syntax. \&quot;mqtt\&quot; - Subscription uses MQTT syntax. &lt;/pre&gt;  | [optional] [default to null]
**TelemetryProfileName** | **string** | The name of the Telemetry Profile. | [optional] [default to null]
**TraceFilterName** | **string** | A name used to identify the trace filter. Consider a name that describes the subscriptions contained within the filter, such as the name of the application and/or the scenario in which the trace filter might be enabled, such as \&quot;appNameDebug\&quot;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

