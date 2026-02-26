# EmailThreadProjection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | **string** | Thread topic subject | [optional] 
**UserId** | **string** | User ID | 
**InboxId** | **string** | Inbox ID | [optional] 
**Sender** | [**SenderProjection**](SenderProjection) |  | [optional] 
**Recipients** | [**EmailRecipientsProjection**](EmailRecipientsProjection) |  | [optional] 
**UpdatedAt** | [**time.Time**](time.Time) | Updated at DateTime | 
**CreatedAt** | [**time.Time**](time.Time) | Created at DateTime | 
**To** | **[]string** | To recipients | 
**Cc** | **[]string** | CC recipients | [optional] 
**Bcc** | **[]string** | BCC recipients | [optional] 
**HasAttachments** | **bool** | Has attachments | 
**LastCreatedAt** | [**time.Time**](time.Time) | Last email created time | [optional] 
**LastFrom** | **string** | Last sender | [optional] 
**LastSender** | [**SenderProjection**](SenderProjection) |  | [optional] 
**Unread** | **bool** | Has unread | 
**MessageCount** | **int32** | Number of messages in the thread | 
**LastBodyExcerpt** | **string** | Last body excerpt | [optional] 
**LastTextExcerpt** | **string** | Last text excerpt | [optional] 
**Id** | **string** | ID of email thread | 
**From** | **string** | From sender | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


