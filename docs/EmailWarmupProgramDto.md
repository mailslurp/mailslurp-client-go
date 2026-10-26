# EmailWarmupProgramDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ConnectionId** | **string** |  | 
**SenderEmail** | **string** |  | 
**Name** | **string** |  | 
**CustomerIntent** | **string** |  | 
**ExecutionPolicy** | **string** |  | 
**ProgramMode** | **string** |  | 
**Status** | **string** |  | 
**Blockers** | **[]string** |  | 
**CurrentDailyVolume** | **int32** |  | 
**GoalDailyVolume** | **int32** |  | 
**SchedulePreset** | **string** |  | 
**RampDurationDays** | **int32** |  | 
**MaintenanceEnabled** | **bool** |  | 
**MaintenanceDailyVolume** | **int32** |  | 
**SendWindowStart** | **string** |  | 
**SendWindowEnd** | **string** |  | 
**TimeZone** | **string** |  | 
**RequestedCoverage** | **string** |  | 
**ResolvedCoverage** | **string** |  | 
**BaselineTestId** | Pointer to **string** |  | [optional] 
**BaselineSource** | **string** |  | 
**ContentSource** | **string** |  | 
**Content** | [**EmailWarmupContentOptions**](EmailWarmupContentOptions) |  | 
**ContentHash** | **string** |  | 
**SenderOwnershipAttestedAt** | [**time.Time**](time.Time) |  | 
**PermissionBasedSendingAttestedAt** | [**time.Time**](time.Time) |  | 
**AttestationPolicyVersion** | **string** |  | 
**PolicyVersion** | **string** |  | 
**ContentPolicyVersion** | **string** |  | 
**RecipientPolicyVersion** | **string** |  | 
**NextActionAt** | Pointer to [**time.Time**](time.Time) |  | [optional] 
**StartedAt** | Pointer to [**time.Time**](time.Time) |  | [optional] 
**CompletedAt** | Pointer to [**time.Time**](time.Time) |  | [optional] 
**CompletionCode** | Pointer to **string** |  | [optional] 
**Version** | **int64** |  | 
**CreatedAt** | [**time.Time**](time.Time) |  | 
**UpdatedAt** | [**time.Time**](time.Time) |  | 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


