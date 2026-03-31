# UpdateDeliverabilityTestOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Optional updated name | [optional] 
**Description** | Pointer to **string** | Optional updated description | [optional] 
**StartAt** | Pointer to [**time.Time**](time.Time) | Optional updated receive-window start time. Only applied while test is not terminal. | [optional] 
**MaxDurationSeconds** | Pointer to **int64** | Optional updated timeout in seconds | [optional] 
**ClearMaxDuration** | Pointer to **bool** | Set true to clear timeout. If true, maxDurationSeconds is ignored for this request. | [optional] 
**SuccessThresholdPercent** | Pointer to **float64** | Optional updated acceptable success threshold percentage (0,100]. | [optional] 
**ClearSuccessThreshold** | Pointer to **bool** | Set true to clear success threshold. If true, successThresholdPercent is ignored for this request. | [optional] 
**Expectations** | Pointer to [**[]DeliverabilityExpectation**](DeliverabilityExpectation) | Optional replacement expectations | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


