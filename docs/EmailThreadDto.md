# EmailThreadDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of email thread | 
**UserId** | **string** | User ID | 
**InboxId** | **string** | Inbox ID | [optional] 
**From** | **string** | From sender | [optional] 
**To** | **[]string** | To recipients | 
**Cc** | **[]string** | CC recipients | [optional] 
**Bcc** | **[]string** | BCC recipients | [optional] 
**Sender** | Pointer to [**Sender**](Sender) |  | [optional] 
**Recipients** | Pointer to [**EmailRecipients**](EmailRecipients) |  | [optional] 
**Subject** | **string** | Thread topic subject | [optional] 
**CreatedAt** | [**time.Time**](time.Time) | Created at DateTime | 
**UpdatedAt** | [**time.Time**](time.Time) | Updated at DateTime | 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


