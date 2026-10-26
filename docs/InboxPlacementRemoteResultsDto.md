# InboxPlacementRemoteResultsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Test** | [**InboxPlacementRemoteTestSummaryDto**](InboxPlacementRemoteTestSummaryDto) |  | 
**Deliverability** | [**InboxPlacementDeliverabilitySummaryDto**](InboxPlacementDeliverabilitySummaryDto) |  | 
**SetupAnalysis** | **map[string]string** |  | 
**ContentAnalysis** | **map[string]string** |  | 
**Analysis** | Pointer to [**InboxPlacementAnalysisBundleDto**](InboxPlacementAnalysisBundleDto) |  | [optional] 
**ProviderAnalysis** | Pointer to [**InboxPlacementProviderAnalysisDto**](InboxPlacementProviderAnalysisDto) |  | [optional] 
**Targets** | [**[]InboxPlacementTargetResultDto**](InboxPlacementTargetResultDto) |  | 
**Scores** | [**[]map[string]string**](map) |  | 
**Share** | Pointer to **map[string]string** |  | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


