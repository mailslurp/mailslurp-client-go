# InboxPlacementDomainMonitorSummaryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** | Sender-domain monitor state derived while analyzing an inbox placement run | 
**Domain** | Pointer to **string** |  | [optional] 
**ProviderManaged** | **bool** |  | 
**Source** | **string** | Source for an inbox-placement domain monitor summary | 
**Status** | Pointer to **string** |  | [optional] 
**HealthScore** | Pointer to **int32** |  | [optional] 
**TotalChecks** | Pointer to **int32** |  | [optional] 
**PassingChecks** | Pointer to **int32** |  | [optional] 
**FailingChecks** | Pointer to **int32** |  | [optional] 
**SpfOk** | Pointer to **bool** |  | [optional] 
**DmarcOk** | Pointer to **bool** |  | [optional] 
**DmarcEnforced** | Pointer to **bool** |  | [optional] 
**MxOk** | Pointer to **bool** |  | [optional] 
**DkimOk** | Pointer to **bool** |  | [optional] 
**BimiOk** | Pointer to **bool** |  | [optional] 
**MtaStsOk** | Pointer to **bool** |  | [optional] 
**TlsReportingOk** | Pointer to **bool** |  | [optional] 
**BlacklistChecked** | **bool** |  | 
**BlacklistOk** | **bool** |  | 
**BlacklistListed** | **bool** |  | 
**BlacklistStatus** | **string** |  | 
**ObservedSendingIp** | Pointer to **string** |  | [optional] 
**ScoreVersion** | **string** |  | 
**ScoreScope** | **string** |  | 
**Insights** | **[]string** |  | 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**MonitorId** | Pointer to **string** |  | [optional] 
**RunId** | Pointer to **string** |  | [optional] 
**CheckedAt** | Pointer to [**time.Time**](time.Time) |  | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


