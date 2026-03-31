# DeliverabilitySimulationJobDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TestId** | **string** |  | 
**Scope** | **string** |  | 
**Status** | **string** |  | 
**CreatedAt** | [**time.Time**](time.Time) |  | 
**UpdatedAt** | [**time.Time**](time.Time) |  | 
**StartedAt** | Pointer to [**time.Time**](time.Time) |  | [optional] 
**CompletedAt** | Pointer to [**time.Time**](time.Time) |  | [optional] 
**TotalTargets** | **int64** |  | 
**TotalPlannedSends** | **int64** |  | 
**NextSendIndex** | **int64** |  | 
**SentCount** | **int64** |  | 
**SuccessCount** | **int64** |  | 
**FailureCount** | **int64** |  | 
**ProgressPercent** | **float64** |  | 
**ActiveTargetLabel** | Pointer to **string** |  | [optional] 
**EstimatedRemainingMs** | Pointer to **int64** |  | [optional] 
**ConfigSnapshot** | [**DeliverabilitySimulationJobConfigDto**](DeliverabilitySimulationJobConfigDto) |  | 
**ErrorSummary** | [**DeliverabilitySimulationJobErrorSummaryDto**](DeliverabilitySimulationJobErrorSummaryDto) |  | 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


