# WebhookNewAttachmentPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | **string** | Idempotent message ID. Store this ID locally or in a database to prevent message duplication. | [optional] 
**WebhookId** | **string** | ID of webhook entity being triggered | [optional] 
**WebhookName** | **string** | Name of the webhook being triggered | [optional] 
**EventName** | **string** | Name of the event type webhook is being triggered for. | [optional] 
**AttachmentId** | **string** | ID of attachment. Use the &#x60;AttachmentController&#x60; to | [optional] 
**Name** | **string** | Filename of the attachment if present | [optional] 
**ContentType** | **string** | Content type of attachment such as &#39;image/png&#39; or &#39;application/pdf | [optional] 
**ContentLength** | **int64** | Size of attachment in bytes | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


