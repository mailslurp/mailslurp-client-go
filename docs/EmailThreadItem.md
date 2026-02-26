# EmailThreadItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemType** | **string** |  | 
**EntityId** | **string** |  | 
**BodyExcerpt** | Pointer to **string** |  | [optional] 
**TextExcerpt** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**To** | **[]string** |  | 
**From** | Pointer to **string** |  | [optional] 
**Bcc** | Pointer to **[]string** |  | [optional] 
**Cc** | Pointer to **[]string** |  | [optional] 
**Attachments** | Pointer to **[]string** |  | [optional] 
**CreatedAt** | [**time.Time**](time.Time) |  | 
**Read** | **bool** |  | 
**InReplyTo** | Pointer to **string** |  | [optional] 
**MessageId** | Pointer to **string** |  | [optional] 
**ThreadId** | Pointer to **string** |  | [optional] 
**Sender** | Pointer to [**Sender**](Sender) |  | [optional] 
**Recipients** | Pointer to [**EmailRecipients**](EmailRecipients) |  | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


