# WebhookEmailOpenedPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | [**time.Time**](time.Time) | Date time of event creation | [optional] 
**EventName** | **string** | Name of the event type webhook is being triggered for. | [optional] 
**InboxId** | **string** | Id of the inbox that received an email | [optional] 
**MessageId** | **string** | Idempotent message ID. Store this ID locally or in a database to prevent message duplication. | [optional] 
**PixelId** | **string** | ID of the tracking pixel | [optional] 
**Recipient** | **string** | Email address for the recipient of the tracking pixel | [optional] 
**SentEmailId** | **string** | ID of sent email | [optional] 
**WebhookId** | **string** | ID of webhook entity being triggered | [optional] 
**WebhookName** | **string** | Name of the webhook being triggered | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)

