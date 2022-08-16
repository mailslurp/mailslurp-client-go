# WebhookEmailOpenedPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | **string** | Idempotent message ID. Store this ID locally or in a database to prevent message duplication. | 
**WebhookId** | **string** | ID of webhook entity being triggered | 
**EventName** | **string** | Name of the event type webhook is being triggered for. | 
**WebhookName** | **string** | Name of the webhook being triggered | [optional] 
**InboxId** | **string** | Id of the inbox | 
**PixelId** | **string** | ID of the tracking pixel | 
**SentEmailId** | **string** | ID of sent email | 
**Recipient** | **string** | Email address for the recipient of the tracking pixel | 
**CreatedAt** | [**time.Time**](time.Time) | Date time of event creation | 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


