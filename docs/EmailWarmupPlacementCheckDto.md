# EmailWarmupPlacementCheckDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ProgramId** | **string** |  | 
**CheckpointDay** | **int32** |  | 
**Type** | **string** |  | 
**Status** | **string** |  | 
**ScheduledFor** | [**time.Time**](time.Time) |  | 
**InboxPlacementRunId** | Pointer to **string** |  | [optional] 
**TargetCount** | **int32** |  | 
**AcceptedCount** | **int32** |  | 
**InboxCount** | **int32** |  | 
**SpamCount** | **int32** |  | 
**NotReceivedCount** | **int32** |  | 
**InboxPercent** | **int32** |  | 
**ProviderResults** | [**[]EmailWarmupProviderPlacementDto**](EmailWarmupProviderPlacementDto) |  | 
**FailureCode** | Pointer to **string** |  | [optional] 
**CompletedAt** | Pointer to [**time.Time**](time.Time) |  | [optional] 
**CreatedAt** | [**time.Time**](time.Time) |  | 
**UpdatedAt** | [**time.Time**](time.Time) |  | 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


