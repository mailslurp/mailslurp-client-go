# DomainPreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Domain** | **string** |  | 
**CatchAllInboxId** | Pointer to **string** |  | [optional] 
**CreatedAt** | [**time.Time**](time.Time) |  | 
**DomainType** | **string** | Type of domain. Dictates type of inbox that can be created with domain. HTTP means inboxes are processed using SES while SMTP inboxes use a custom SMTP mail server. SMTP does not support sending so use HTTP for sending emails. | 
**IsVerified** | **bool** |  | 
**HasMissingRecords** | **bool** |  | 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


