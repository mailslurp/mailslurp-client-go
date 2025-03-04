# EmailPreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the email entity | 
**InboxId** | Pointer to **string** | ID of the inbox that received the email | [optional] 
**DomainId** | Pointer to **string** | ID of the domain that received the email | [optional] 
**Subject** | Pointer to **string** | The subject line of the email message as specified by SMTP subject header | [optional] 
**To** | Pointer to **[]string** | List of &#x60;To&#x60; recipient email addresses that the email was addressed to. See recipients object for names. | 
**From** | Pointer to **string** | Who the email was sent from. An email address - see fromName for the sender name. | [optional] 
**Bcc** | Pointer to **[]string** | List of &#x60;BCC&#x60; recipients email addresses that the email was addressed to. See recipients object for names. | [optional] 
**Cc** | Pointer to **[]string** | List of &#x60;CC&#x60; recipients email addresses that the email was addressed to. See recipients object for names. | [optional] 
**CreatedAt** | [**time.Time**](time.Time) | When was the email received by MailSlurp | 
**Read** | **bool** | Read flag. Has the email ever been viewed in the dashboard or fetched via the API with a hydrated body? If so the email is marked as read. Paginated results do not affect read status. Read status is different to email opened event as it depends on your own account accessing the email. Email opened is determined by tracking pixels sent to other uses if enable during sending. You can listened for both email read and email opened events using webhooks. | 
**Attachments** | Pointer to **[]string** | List of IDs of attachments found in the email. Use these IDs with the Inbox and Email Controllers to download attachments and attachment meta data such as filesize, name, extension. | [optional] 
**ThreadId** | Pointer to **string** | MailSlurp thread ID for email chain that enables lookup for In-Reply-To and References fields. | [optional] 
**MessageId** | Pointer to **string** | RFC 5322 Message-ID header value without angle brackets. | [optional] 
**InReplyTo** | Pointer to **string** | Parsed value of In-Reply-To header. A Message-ID in a thread. | [optional] 
**Sender** | Pointer to [**Sender**](Sender) |  | [optional] 
**Recipients** | Pointer to [**EmailRecipients**](EmailRecipients) |  | [optional] 
**Favourite** | Pointer to **bool** |  | [optional] 
**BodyPartContentTypes** | Pointer to **[]string** |  | [optional] 
**PlusAddress** | Pointer to **string** |  | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


