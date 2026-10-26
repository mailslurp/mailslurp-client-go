# PostmasterTrafficStatsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Domain** | **string** |  | 
**ReportDate** | Pointer to **string** |  | [optional] 
**UserReportedSpamRatio** | Pointer to **float64** |  | [optional] 
**DomainReputation** | Pointer to **string** |  | [optional] 
**IpReputations** | [**[]PostmasterIpReputationDto**](PostmasterIpReputationDto) |  | 
**SpfSuccessRatio** | Pointer to **float64** |  | [optional] 
**DkimSuccessRatio** | Pointer to **float64** |  | [optional] 
**DmarcSuccessRatio** | Pointer to **float64** |  | [optional] 
**OutboundEncryptionRatio** | Pointer to **float64** |  | [optional] 
**InboundEncryptionRatio** | Pointer to **float64** |  | [optional] 
**Payload** | [**map[string]interface{}**]() |  | 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


