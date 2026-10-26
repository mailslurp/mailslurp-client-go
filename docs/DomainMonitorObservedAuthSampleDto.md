# DomainMonitorObservedAuthSampleDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**MonitorId** | **string** |  | 
**Domain** | **string** |  | 
**RecipientAddress** | **string** |  | 
**SourceMessageId** | Pointer to **string** |  | [optional] 
**MessageId** | Pointer to **string** |  | [optional] 
**FromAddress** | Pointer to **string** |  | [optional] 
**FromDomain** | Pointer to **string** |  | [optional] 
**ReturnPath** | Pointer to **string** |  | [optional] 
**ReturnPathDomain** | Pointer to **string** |  | [optional] 
**SpfStatus** | **string** |  | 
**SpfEvaluatedDomain** | Pointer to **string** |  | [optional] 
**SpfSourceIp** | Pointer to **string** |  | [optional] 
**DkimStatus** | **string** |  | 
**DkimSignatures** | [**[]DomainMonitorDkimSignatureDto**](DomainMonitorDkimSignatureDto) |  | 
**DmarcStatus** | **string** |  | 
**DmarcAlignedIdentifier** | Pointer to **string** |  | [optional] 
**TrustedAuthenticationResultsSource** | Pointer to **string** |  | [optional] 
**ProviderHints** | **[]string** |  | 
**Accepted** | **bool** |  | 
**ValidationError** | Pointer to **string** |  | [optional] 
**ReceivedAt** | [**time.Time**](time.Time) |  | 
**CreatedAt** | [**time.Time**](time.Time) |  | 
**UpdatedAt** | [**time.Time**](time.Time) |  | 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


