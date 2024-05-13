# SentEmailDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of sent email | 
**UserId** | **string** | User ID | 
**InboxId** | **string** | Inbox ID email was sent from | 
**DomainId** | Pointer to **string** | Domain ID | [optional] 
**To** | Pointer to **[]string** | Recipients email was sent to | [optional] 
**From** | Pointer to **string** | Sent from address | [optional] 
**ReplyTo** | Pointer to **string** |  | [optional] 
**Cc** | Pointer to **[]string** |  | [optional] 
**Bcc** | Pointer to **[]string** |  | [optional] 
**Attachments** | Pointer to **[]string** | Array of IDs of attachments that were sent with this email | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**BodyMD5Hash** | Pointer to **string** | MD5 Hash | [optional] 
**Body** | Pointer to **string** | Sent email body | [optional] 
**ToContacts** | Pointer to **[]string** |  | [optional] 
**ToGroup** | Pointer to **string** |  | [optional] 
**Charset** | Pointer to **string** |  | [optional] 
**IsHTML** | Pointer to **bool** |  | [optional] 
**SentAt** | [**time.Time**](time.Time) |  | 
**PixelIds** | Pointer to **[]string** |  | [optional] 
**MessageId** | Pointer to **string** |  | [optional] 
**MessageIds** | Pointer to **[]string** |  | [optional] 
**VirtualSend** | Pointer to **bool** |  | [optional] 
**TemplateId** | Pointer to **string** |  | [optional] 
**TemplateVariables** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Headers** | Pointer to **map[string]string** |  | [optional] 
**Html** | **bool** |  | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


