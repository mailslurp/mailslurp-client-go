# SentEmailDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of sent email | [optional] 
**UserId** | **string** | User ID | [optional] 
**InboxId** | **string** | Inbox ID email was sent from | [optional] 
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
**SentAt** | [**time.Time**](time.Time) |  | [optional] 
**PixelIds** | **[]string** |  | [optional] 
**Html** | **bool** |  | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


