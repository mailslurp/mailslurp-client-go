# EmailThreadProjection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | Pointer to **string** | Thread topic subject | [optional] 
**From** | Pointer to **string** | From sender | [optional] 
**Id** | **string** | ID of email thread | 
**UserId** | **string** | User ID | 
**InboxId** | Pointer to **string** | Inbox ID | [optional] 
**Recipients** | Pointer to [**EmailRecipientsProjection**](EmailRecipientsProjection) |  | [optional] 
**To** | **[]string** | To recipients | 
**CreatedAt** | [**time.Time**](time.Time) | Created at DateTime | 
**Cc** | Pointer to **[]string** | CC recipients | [optional] 
**Bcc** | Pointer to **[]string** | BCC recipients | [optional] 
**UpdatedAt** | [**time.Time**](time.Time) | Updated at DateTime | 
**Sender** | Pointer to [**SenderProjection**](SenderProjection) |  | [optional] 
**LastCreatedAt** | Pointer to [**time.Time**](time.Time) | Last email created time | [optional] 
**HasAttachments** | **bool** | Has attachments | 
**LastSender** | Pointer to [**SenderProjection**](SenderProjection) |  | [optional] 
**LastFrom** | Pointer to **string** | Last sender | [optional] 
**Unread** | **bool** | Has unread | 
**MessageCount** | **int32** | Number of messages in the thread | 
**LastBodyExcerpt** | Pointer to **string** | Last body excerpt | [optional] 
**LastTextExcerpt** | Pointer to **string** | Last text excerpt | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


