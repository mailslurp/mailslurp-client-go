# DomainMonitorRunEvidenceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScoreVersion** | **string** |  | 
**ScoreScope** | **string** |  | 
**ObservedMessageReceivedAt** | Pointer to [**time.Time**](time.Time) |  | [optional] 
**OutboundIp** | Pointer to **string** |  | [optional] 
**OutboundIpSource** | Pointer to **string** |  | [optional] 
**ObservedDkimSelector** | Pointer to **string** |  | [optional] 
**ObservedSpfStatus** | Pointer to **string** |  | [optional] 
**ObservedDkimStatus** | Pointer to **string** |  | [optional] 
**ObservedDmarcStatus** | Pointer to **string** |  | [optional] 
**Blacklist** | Pointer to [**CheckEmailBlacklistResults**](CheckEmailBlacklistResults) |  | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


