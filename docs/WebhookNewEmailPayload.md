# WebhookNewEmailPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | **string** | Idempotent message ID. Store this ID locally or in a database to prevent message duplication. | [optional] 
**WebhookId** | **string** | ID of webhook entity being triggered | [optional] 
**EventName** | **string** | Name of the event type webhook is being triggered for. | [optional] 
**WebhookName** | **string** | Name of the webhook being triggered | [optional] 
**InboxId** | **string** | Id of the inbox that received an email | [optional] 
**EmailId** | **string** | ID of the email that was received. Use this ID for fetching the email with the &#x60;EmailController&#x60;. | [optional] 
**CreatedAt** | [**time.Time**](time.Time) | Date time of event creation | [optional] 
**To** | **[]string** | List of &#x60;To&#x60; recipient email addresses that the email was addressed to. See recipients object for names. | [optional] 
**From** | **string** | Who the email was sent from. An email address - see fromName for the sender name. | [optional] 
**Cc** | **[]string** | List of &#x60;CC&#x60; recipients email addresses that the email was addressed to. See recipients object for names. | [optional] 
**Bcc** | **[]string** | List of &#x60;BCC&#x60; recipients email addresses that the email was addressed to. See recipients object for names. | [optional] 
**Subject** | **string** | The subject line of the email message as specified by SMTP subject header | [optional] 
**AttachmentMetaDatas** | [**[]AttachmentMetaData**](AttachmentMetaData) | List of attachment meta data objects if attachments present | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


