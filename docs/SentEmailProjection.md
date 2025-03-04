# SentEmailProjection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | Pointer to **string** |  | [optional] 
**CreatedAt** | [**time.Time**](time.Time) |  | 
**Id** | **string** |  | 
**From** | Pointer to **string** |  | [optional] 
**Sender** | Pointer to [**Sender**](Sender) |  | [optional] 
**Recipients** | Pointer to [**EmailRecipients**](EmailRecipients) |  | [optional] 
**UserId** | **string** |  | 
**Attachments** | Pointer to **[]string** |  | [optional] 
**InboxId** | **string** |  | 
**To** | **[]string** |  | 
**Cc** | **[]string** |  | 
**Bcc** | **[]string** |  | 
**MessageId** | Pointer to **string** |  | [optional] 
**InReplyTo** | Pointer to **string** |  | [optional] 
**BodyExcerpt** | Pointer to **string** |  | [optional] 
**TextExcerpt** | Pointer to **string** |  | [optional] 
**BodyMD5Hash** | Pointer to **string** |  | [optional] 
**VirtualSend** | **bool** |  | 
**ThreadId** | Pointer to **string** |  | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


