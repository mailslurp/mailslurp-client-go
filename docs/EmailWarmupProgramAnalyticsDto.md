# EmailWarmupProgramAnalyticsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProgramId** | **string** |  | 
**FromDate** | **string** |  | 
**ThroughDate** | **string** |  | 
**Progress** | [**EmailWarmupProgramProgressDto**](EmailWarmupProgramProgressDto) |  | 
**TotalScheduledSends** | **int64** |  | 
**TotalAcceptedSends** | **int64** |  | 
**TotalObservedSends** | **int64** |  | 
**TotalFailedSends** | **int64** |  | 
**AttentionRequiredSends** | **int64** |  | 
**DailyActivity** | [**[]EmailWarmupDailyActivityDto**](EmailWarmupDailyActivityDto) |  | 
**PlacementChecks** | [**[]EmailWarmupPlacementCheckDto**](EmailWarmupPlacementCheckDto) |  | 
**Readiness** | Pointer to [**EmailWarmupReadinessDto**](EmailWarmupReadinessDto) |  | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


