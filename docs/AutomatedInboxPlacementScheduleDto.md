# AutomatedInboxPlacementScheduleDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ConnectionId** | **string** |  | 
**Name** | **string** |  | 
**Enabled** | **bool** |  | 
**IntervalMinutes** | **int64** |  | 
**RequestedSegment** | **string** |  | 
**SeedlistProfile** | Pointer to **string** |  | [optional] 
**SubjectTemplate** | **string** |  | 
**TextBodyTemplate** | Pointer to **string** |  | [optional] 
**HtmlBodyTemplate** | Pointer to **string** |  | [optional] 
**SourceRunId** | Pointer to **string** |  | [optional] 
**LastStatus** | Pointer to **string** |  | [optional] 
**ConsecutiveFailures** | **int32** |  | 
**LastRunAt** | Pointer to [**time.Time**](time.Time) |  | [optional] 
**NextRunAt** | Pointer to [**time.Time**](time.Time) |  | [optional] 
**CreatedAt** | [**time.Time**](time.Time) |  | 
**UpdatedAt** | [**time.Time**](time.Time) |  | 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


