# DomainMonitorDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**UserId** | **string** |  | 
**Domain** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**IntervalSeconds** | Pointer to **int64** |  | [optional] 
**Enabled** | **bool** |  | 
**SchedulingEnabled** | **bool** |  | 
**LastStatus** | Pointer to **string** |  | [optional] 
**HealthScore** | Pointer to **int32** |  | [optional] 
**LastRunAt** | Pointer to [**time.Time**](time.Time) |  | [optional] 
**NextRunAt** | Pointer to [**time.Time**](time.Time) |  | [optional] 
**EmailVerification** | Pointer to [**DomainMonitorEmailVerificationDto**](DomainMonitorEmailVerificationDto) |  | [optional] 
**CreatedAt** | [**time.Time**](time.Time) |  | 
**UpdatedAt** | [**time.Time**](time.Time) |  | 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


