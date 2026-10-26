# CreateDeliverabilityTestOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Optional name for the test | [optional] 
**Description** | Pointer to **string** | Optional description | [optional] 
**Scope** | **string** | Entity scope to evaluate | 
**StartAt** | Pointer to [**time.Time**](time.Time) | UTC instant when the receive window starts. Defaults to now if omitted. | [optional] 
**MaxDurationSeconds** | Pointer to **int64** | Optional timeout in seconds after startAt. If not all entities match before timeout the test transitions to FAILED. | [optional] 
**SuccessThresholdPercent** | Pointer to **float64** | Optional acceptable success threshold percentage (0,100]. If set, a timed-out test can complete successfully when matchedEntities/totalEntities reaches this percentage. | [optional] 
**Selector** | [**DeliverabilitySelectorOptions**](DeliverabilitySelectorOptions) |  | 
**Expectations** | [**[]DeliverabilityExpectation**](DeliverabilityExpectation) | One or more expectations to evaluate for each entity | 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


