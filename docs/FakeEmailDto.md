# FakeEmailDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**EmailAddress** | **string** |  | 
**Sender** | Pointer to [**Sender**](Sender) |  | [optional] 
**Recipients** | Pointer to [**EmailRecipients**](EmailRecipients) |  | [optional] 
**AttachmentNames** | **[]string** |  | 
**Subject** | **string** |  | [optional] 
**Preview** | **string** |  | [optional] 
**Body** | **string** | use read content endpoints instead | 
**Seen** | **bool** |  | 
**CreatedAt** | [**time.Time**](time.Time) |  | 
**ContentType** | **string** |  | 
**BodyUrl** | **string** |  | 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


