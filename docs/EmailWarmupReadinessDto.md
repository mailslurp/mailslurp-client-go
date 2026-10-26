# EmailWarmupReadinessDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Summary** | **string** |  | 
**ReadyForDailyVolume** | **int32** |  | 
**GoalDailyVolume** | **int32** |  | 
**EstimatedReadyAt** | Pointer to **string** |  | [optional] 
**Today** | [**EmailWarmupTodayDto**](EmailWarmupTodayDto) |  | 
**Pacing** | [**EmailWarmupPacingDto**](EmailWarmupPacingDto) |  | 
**PlacementDeltas** | [**[]EmailWarmupPlacementDeltaDto**](EmailWarmupPlacementDeltaDto) |  | 
**DomainHealth** | [**EmailWarmupDomainHealthDto**](EmailWarmupDomainHealthDto) |  | 
**PostmasterSignals** | Pointer to [**EmailWarmupPostmasterSignalsDto**](EmailWarmupPostmasterSignalsDto) |  | [optional] 
**Recommendations** | **[]string** |  | 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


