# EventThreshold

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClearPercent** | **int64** | The clear threshold for the value of this counter as a percentage of its maximum value. Falling below this value will trigger a corresponding event. This attribute may not be returned in a GET. | [optional] [default to null]
**ClearValue** | **int64** | The clear threshold for the absolute value of this counter. Falling below this value will trigger a corresponding event. This attribute may not be returned in a GET. | [optional] [default to null]
**SetPercent** | **int64** | The set threshold for the value of this counter as a percentage of its maximum value. Exceeding this value will trigger a corresponding event. This attribute may not be returned in a GET. | [optional] [default to null]
**SetValue** | **int64** | The set threshold for the absolute value of this counter. Exceeding this value will trigger a corresponding event. This attribute may not be returned in a GET. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

