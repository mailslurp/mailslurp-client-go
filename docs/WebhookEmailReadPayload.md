# WebhookEmailReadPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | **string** | Idempotent message ID. Store this ID locally or in a database to prevent message duplication. | 
**WebhookId** | **string** | ID of webhook entity being triggered | 
**EventName** | **string** | Name of the event type webhook is being triggered for. | 
**WebhookName** | Pointer to **string** | Name of the webhook being triggered | [optional] 
**EmailId** | **string** | ID of the email that was received. Use this ID for fetching the email with the &#x60;EmailController&#x60;. | 
**InboxId** | **string** | Id of the inbox | 
**EmailIsRead** | **bool** | Is the email read | 
**CreatedAt** | [**time.Time**](time.Time) | Date time of event creation | 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


