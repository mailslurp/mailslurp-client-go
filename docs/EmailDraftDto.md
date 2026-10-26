# EmailDraftDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the draft | 
**UserId** | **string** | User ID | 
**InboxId** | **string** | Inbox ID the draft will send from | 
**To** | Pointer to **[]string** | Draft recipients | [optional] 
**Cc** | Pointer to **[]string** | Draft CC recipients | [optional] 
**Bcc** | Pointer to **[]string** | Draft BCC recipients | [optional] 
**From** | Pointer to **string** | Draft sender | [optional] 
**Sender** | Pointer to [**Sender**](Sender) |  | [optional] 
**Recipients** | Pointer to [**EmailRecipients**](EmailRecipients) |  | [optional] 
**Subject** | Pointer to **string** | Draft subject | [optional] 
**BodyExcerpt** | Pointer to **string** | Draft body preview | [optional] 
**TextExcerpt** | Pointer to **string** | Draft text preview | [optional] 
**Attachments** | Pointer to **[]string** | Attachment IDs selected for the draft | [optional] 
**SendOptions** | [**SendEmailOptions**](SendEmailOptions) |  | 
**CreatedAt** | [**time.Time**](time.Time) | Draft creation time | 
**UpdatedAt** | [**time.Time**](time.Time) | Draft last update time | 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


