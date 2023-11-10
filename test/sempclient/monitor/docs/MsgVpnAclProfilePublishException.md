# MsgVpnAclProfilePublishException

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AclProfileName** | **string** | The name of the ACL Profile. Deprecated since 2.14. Replaced by publishTopicExceptions. | [optional] [default to null]
**MsgVpnName** | **string** | The name of the Message VPN. Deprecated since 2.14. Replaced by publishTopicExceptions. | [optional] [default to null]
**PublishExceptionTopic** | **string** | The topic for the exception to the default action taken. May include wildcard characters. Deprecated since 2.14. Replaced by publishTopicExceptions. | [optional] [default to null]
**TopicSyntax** | **string** | The syntax of the topic for the exception to the default action taken. The allowed values and their meaning are:  &lt;pre&gt; \&quot;smf\&quot; - Topic uses SMF syntax. \&quot;mqtt\&quot; - Topic uses MQTT syntax. &lt;/pre&gt;  Deprecated since 2.14. Replaced by publishTopicExceptions. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

