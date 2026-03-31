# CampaignProbeInsightsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProbeId** | **string** |  | 
**Since** | [**time.Time**](time.Time) |  | 
**Before** | [**time.Time**](time.Time) |  | 
**TotalRuns** | **int32** |  | 
**HealthyRuns** | **int32** |  | 
**WarningRuns** | **int32** |  | 
**FailedRuns** | **int32** |  | 
**SuccessRate** | **float64** |  | 
**AverageHealthScore** | **float64** |  | 
**CurrentHealthyStreak** | **int32** |  | 
**BestHealthyStreak** | **int32** |  | 
**GoodPerformanceSignals** | **[]string** |  | 
**AttentionSignals** | **[]string** |  | 
**TopFindings** | **[]string** |  | 
**LatestRun** | [**CampaignProbeRunDto**](CampaignProbeRunDto) |  | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


