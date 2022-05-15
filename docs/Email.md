# Email

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the email entity | 
**UserId** | **string** | ID of user that email belongs to | 
**InboxId** | **string** | ID of the inbox that received the email | 
**To** | **[]string** | List of &#x60;To&#x60; recipient email addresses that the email was addressed to. See recipients object for names. | 
**From** | **string** | Who the email was sent from. An email address - see fromName for the sender name. | [optional] 
**Sender** | [**Sender**](Sender) |  | [optional] 
**Recipients** | [**EmailRecipients**](EmailRecipients) |  | [optional] 
**ReplyTo** | **string** | The &#x60;replyTo&#x60; field on the received email message | [optional] 
**Cc** | **[]string** | List of &#x60;CC&#x60; recipients email addresses that the email was addressed to. See recipients object for names. | [optional] 
**Bcc** | **[]string** | List of &#x60;BCC&#x60; recipients email addresses that the email was addressed to. See recipients object for names. | [optional] 
**Headers** | **map[string]string** | Collection of SMTP headers attached to email | [optional] 
**Attachments** | **[]string** | List of IDs of attachments found in the email. Use these IDs with the Inbox and Email Controllers to download attachments and attachment meta data such as filesize, name, extension. | [optional] 
**Subject** | **string** | The subject line of the email message as specified by SMTP subject header | [optional] 
**Body** | **string** | The body of the email message as text parsed from the SMTP message body (does not include attachments). Fetch the raw content to access the SMTP message and use the attachments property to access attachments. The body is stored separately to the email entity so the body is not returned in paginated results only in full single email or wait requests. | [optional] 
**BodyExcerpt** | **string** | An excerpt of the body of the email message for quick preview . | [optional] 
**BodyMD5Hash** | **string** | A hash signature of the email message using MD5. Useful for comparing emails without fetching full body. | [optional] 
**IsHTML** | **bool** | Is the email body content type HTML? | [optional] 
**Charset** | **string** | Detected character set of the email body such as UTF-8 | [optional] 
**Analysis** | [**EmailAnalysis**](EmailAnalysis) |  | [optional] 
**CreatedAt** | [**time.Time**](time.Time) | When was the email received by MailSlurp | 
**UpdatedAt** | [**time.Time**](time.Time) | When was the email last updated | 
**Read** | **bool** | Read flag. Has the email ever been viewed in the dashboard or fetched via the API with a hydrated body? If so the email is marked as read. Paginated results do not affect read status. Read status is different to email opened event as it depends on your own account accessing the email. Email opened is determined by tracking pixels sent to other uses if enable during sending. You can listened for both email read and email opened events using webhooks. | 
**TeamAccess** | **bool** | Can the email be accessed by organization team members | 
**Html** | **bool** |  | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


