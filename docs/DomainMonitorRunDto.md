# DomainMonitorRunDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**MonitorId** | **string** |  | 
**UserId** | **string** |  | 
**Domain** | **string** |  | 
**Status** | **string** |  | 
**TriggerSource** | **string** |  | 
**HealthScore** | **int32** |  | 
**TotalChecks** | **int32** |  | 
**PassingChecks** | **int32** |  | 
**FailingChecks** | **int32** |  | 
**SpfOk** | **bool** |  | 
**DmarcOk** | **bool** |  | 
**DmarcEnforced** | **bool** |  | 
**MxOk** | **bool** |  | 
**DkimOk** | **bool** |  | 
**BimiOk** | **bool** |  | 
**MtaStsOk** | **bool** |  | 
**TlsReportingOk** | **bool** |  | 
**BlacklistChecked** | **bool** |  | 
**BlacklistOk** | **bool** |  | 
**BlacklistListed** | **bool** |  | 
**Evidence** | Pointer to [**DomainMonitorRunEvidenceDto**](DomainMonitorRunEvidenceDto) |  | [optional] 
**Insights** | **[]string** |  | 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**CreatedAt** | [**time.Time**](time.Time) |  | 
**UpdatedAt** | [**time.Time**](time.Time) |  | 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


