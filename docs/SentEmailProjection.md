# SentEmailProjection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ThreadId** | Pointer to **string** |  | [optional] 
**InReplyTo** | Pointer to **string** |  | [optional] 
**From** | Pointer to **string** |  | [optional] 
**Sender** | Pointer to [**Sender**](Sender) |  | [optional] 
**Recipients** | Pointer to [**EmailRecipients**](EmailRecipients) |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**UserId** | **string** |  | 
**Attachments** | Pointer to **[]string** |  | [optional] 
**InboxId** | **string** |  | 
**CreatedAt** | [**time.Time**](time.Time) |  | 
**To** | **[]string** |  | [optional] 
**Cc** | **[]string** |  | [optional] 
**Bcc** | **[]string** |  | [optional] 
**MessageId** | Pointer to **string** |  | [optional] 
**VirtualSend** | **bool** |  | 
**BodyExcerpt** | Pointer to **string** |  | [optional] 
**TextExcerpt** | Pointer to **string** |  | [optional] 
**BodyMD5Hash** | Pointer to **string** |  | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


