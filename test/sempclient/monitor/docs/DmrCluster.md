# DmrCluster

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationBasicEnabled** | [****bool**](*bool.md) | Indicates whether basic authentication is enabled for Cluster Links. | [optional] [default to null]
**AuthenticationBasicType** | **string** | The type of basic authentication to use for Cluster Links. The allowed values and their meaning are:  &lt;pre&gt; \&quot;internal\&quot; - Use locally configured password. \&quot;none\&quot; - No authentication. &lt;/pre&gt;  | [optional] [default to null]
**AuthenticationClientCertEnabled** | [****bool**](*bool.md) | Indicates whether client certificate authentication is enabled for Cluster Links. | [optional] [default to null]
**DirectOnlyEnabled** | [****bool**](*bool.md) | Indicates whether this cluster only supports direct messaging. If true, guaranteed messages will not be transmitted through the cluster. | [optional] [default to null]
**DmrClusterName** | **string** | The name of the Cluster. | [optional] [default to null]
**Enabled** | [****bool**](*bool.md) | Indicates whether the Cluster is enabled. | [optional] [default to null]
**FailureReason** | **string** | The failure reason for the Cluster being down. | [optional] [default to null]
**NodeName** | **string** | The name of this node in the Cluster. This is the name that this broker (or redundant group of brokers) is know by to other nodes in the Cluster. The name is chosen automatically to be either this broker&#x27;s Router Name or Mate Router Name, depending on which Active Standby Role (primary or backup) this broker plays in its redundancy group. | [optional] [default to null]
**SubscriptionDbBuildPercentage** | **int64** | Cluster Subscription Database build completion percentage. Available since 2.20. | [optional] [default to null]
**TlsServerCertEnforceTrustedCommonNameEnabled** | [****bool**](*bool.md) | Indicates whether the common name provided by the remote broker is enforced against the list of trusted common names configured for the Link. If enabled, the certificate&#x27;s common name must match one of the trusted common names for the Link to be accepted. Deprecated since 2.18. Common Name validation has been replaced by Server Certificate Name validation. | [optional] [default to null]
**TlsServerCertMaxChainDepth** | **int64** | The maximum allowed depth of a certificate chain. The depth of a chain is defined as the number of signing CA certificates that are present in the chain back to a trusted self-signed root CA certificate. | [optional] [default to null]
**TlsServerCertValidateDateEnabled** | [****bool**](*bool.md) | Indicates whether validation of the \&quot;Not Before\&quot; and \&quot;Not After\&quot; validity dates in the certificate is enabled. When disabled, the certificate is accepted even if the certificate is not valid based on these dates. | [optional] [default to null]
**TlsServerCertValidateNameEnabled** | [****bool**](*bool.md) | Enable or disable the standard TLS authentication mechanism of verifying the name used to connect to the bridge. If enabled, the name used to connect to the bridge is checked against the names specified in the certificate returned by the remote router. Legacy Common Name validation is not performed if Server Certificate Name Validation is enabled, even if Common Name validation is also enabled. Available since 2.18. | [optional] [default to null]
**Up** | [****bool**](*bool.md) | Indicates whether the Cluster is operationally up. | [optional] [default to null]
**Uptime** | **int64** | The amount of time in seconds since the Cluster was up. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

