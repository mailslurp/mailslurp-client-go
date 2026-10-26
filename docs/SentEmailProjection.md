# SentEmailProjection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | Pointer to **string** |  | [optional] 
**From** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**ThreadId** | Pointer to **string** |  | [optional] 
**UserId** | **string** |  | 
**InboxId** | **string** |  | 
**Recipients** | Pointer to [**EmailRecipients**](EmailRecipients) |  | [optional] 
**Attachments** | Pointer to **[]string** |  | [optional] 
**TextExcerpt** | Pointer to **string** |  | [optional] 
**BodyExcerpt** | Pointer to **string** |  | [optional] 
**To** | Pointer to **[]string** |  | [optional] 
**CreatedAt** | [**time.Time**](time.Time) |  | 
**Cc** | Pointer to **[]string** |  | [optional] 
**Bcc** | Pointer to **[]string** |  | [optional] 
**MessageId** | Pointer to **string** |  | [optional] 
**InReplyTo** | Pointer to **string** |  | [optional] 
**Sender** | Pointer to [**Sender**](Sender) |  | [optional] 
**BodyMD5Hash** | Pointer to **string** |  | [optional] 
**VirtualSend** | **bool** |  | 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


