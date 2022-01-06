# WebhookEmailReadPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | **string** | Idempotent message ID. Store this ID locally or in a database to prevent message duplication. | [optional] 
**WebhookId** | **string** | ID of webhook entity being triggered | [optional] 
**EventName** | **string** | Name of the event type webhook is being triggered for. | [optional] 
**WebhookName** | **string** | Name of the webhook being triggered | [optional] 
**EmailId** | **string** | ID of the email that was received. Use this ID for fetching the email with the &#x60;EmailController&#x60;. | [optional] 
**InboxId** | **string** | Id of the inbox that received an email | [optional] 
**EmailIsRead** | **bool** | Is the email read | [optional] 
**CreatedAt** | [**time.Time**](time.Time) | Date time of event creation | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


