# DeliverabilityTestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Scope** | **string** |  | 
**Selector** | [**DeliverabilitySelectorOptions**](DeliverabilitySelectorOptions) |  | 
**SelectedEntityCount** | **int64** |  | 
**Expectations** | [**[]DeliverabilityExpectation**](DeliverabilityExpectation) |  | 
**Status** | **string** |  | 
**StartAt** | [**time.Time**](time.Time) |  | 
**StartedAt** | Pointer to [**time.Time**](time.Time) |  | [optional] 
**CompletedAt** | Pointer to [**time.Time**](time.Time) |  | [optional] 
**MaxDurationSeconds** | Pointer to **int64** |  | [optional] 
**SuccessThresholdPercent** | Pointer to **float64** |  | [optional] 
**ThresholdStatus** | **string** |  | 
**ThresholdMet** | Pointer to **bool** |  | [optional] 
**LastEvaluatedAt** | Pointer to [**time.Time**](time.Time) |  | [optional] 
**TotalEntities** | **int64** |  | 
**MatchedEntities** | **int64** |  | 
**UnmatchedEntities** | **int64** |  | 
**CompletionPercentage** | **float64** |  | 
**TimedOut** | **bool** |  | 
**FailureReason** | Pointer to **string** |  | [optional] 
**CreatedAt** | [**time.Time**](time.Time) |  | 
**UpdatedAt** | [**time.Time**](time.Time) |  | 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


