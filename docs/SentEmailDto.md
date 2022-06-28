# SentEmailDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of sent email | 
**UserId** | **string** | User ID | 
**InboxId** | **string** | Inbox ID email was sent from | 
**To** | **[]string** | Recipients email was sent to | [optional] 
**From** | **string** |  | [optional] 
**ReplyTo** | **string** |  | [optional] 
**Cc** | **[]string** |  | [optional] 
**Bcc** | **[]string** |  | [optional] 
**Attachments** | **[]string** | Array of IDs of attachments that were sent with this email | [optional] 
**Subject** | **string** |  | [optional] 
**BodyMD5Hash** | **string** | MD5 Hash | [optional] 
**Body** | **string** |  | [optional] 
**Charset** | **string** |  | [optional] 
**IsHTML** | **bool** |  | [optional] 
**SentAt** | [**time.Time**](time.Time) |  | 
**PixelIds** | **[]string** |  | [optional] 
**MessageId** | **string** |  | [optional] 
**VirtualSend** | **bool** |  | [optional] 
**Html** | **bool** |  | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


