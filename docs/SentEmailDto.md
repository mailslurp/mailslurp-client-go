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
**Sender** | Pointer to [**Sender**](Sender) |  | [optional] 
**Recipients** | Pointer to [**EmailRecipients**](EmailRecipients) |  | [optional] 
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
**CreatedAt** | [**time.Time**](time.Time) |  | 
**PixelIds** | Pointer to **[]string** |  | [optional] 
**MessageId** | Pointer to **string** | RFC 5322 Message-ID header value without angle brackets. | [optional] 
**MessageIds** | Pointer to **[]string** |  | [optional] 
**VirtualSend** | Pointer to **bool** |  | [optional] 
**TemplateId** | Pointer to **string** |  | [optional] 
**TemplateVariables** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Headers** | Pointer to **map[string]string** |  | [optional] 
**ThreadId** | Pointer to **string** | MailSlurp thread ID for email chain that enables lookup for In-Reply-To and References fields. | [optional] 
**BodyExcerpt** | Pointer to **string** | An excerpt of the body of the email message for quick preview. Takes HTML content part if exists falls back to TEXT content part if not | [optional] 
**TextExcerpt** | Pointer to **string** | An excerpt of the body of the email message for quick preview. Takes TEXT content part if exists | [optional] 
**InReplyTo** | Pointer to **string** | Parsed value of In-Reply-To header. A Message-ID in a thread. | [optional] 
**Favourite** | Pointer to **bool** | Is email favourited | [optional] 
**SizeBytes** | Pointer to **int64** | Size of raw email message in bytes | [optional] 
**Html** | **bool** |  | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


