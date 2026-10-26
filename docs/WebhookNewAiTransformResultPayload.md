# WebhookNewAiTransformResultPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | **string** | Idempotent message ID. Store this ID locally or in a database to prevent message duplication. | 
**WebhookId** | **string** | ID of webhook entity being triggered | 
**EventName** | **string** | Name of the event type webhook is being triggered for. | 
**WebhookName** | Pointer to **string** | Name of the webhook being triggered | [optional] 
**AiTransformResultId** | **string** | AI Transform ID of event | 
**UserId** | **string** | User ID of event | 
**AiTransformId** | **string** | ID of AI Transform | 
**AiTransformMappingId** | Pointer to **string** | ID of AI Transform mapping | [optional] 
**EntityId** | Pointer to **string** | ID of entity that triggered the transformation | [optional] 
**EntityType** | Pointer to **string** | Entity type that triggered the transformation | [optional] 
**Result** | Pointer to **string** | JSON string result of the AI transformation | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


