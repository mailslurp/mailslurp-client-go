# WebhookBouncePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | **string** | Idempotent message ID. Store this ID locally or in a database to prevent message duplication. | 
**WebhookId** | **string** | ID of webhook entity being triggered | 
**EventName** | **string** | Name of the event type webhook is being triggered for. | 
**WebhookName** | Pointer to **string** | Name of the webhook being triggered | [optional] 
**BounceId** | **string** | ID of the bounce email record. Use the ID with the bounce controller to view more information | 
**SentToRecipients** | Pointer to **[]string** | Email sent to recipients | [optional] 
**Sender** | **string** | Sender causing bounce | 
**BounceRecipients** | Pointer to **[]string** | Email addresses that resulted in a bounce or email being rejected. Please save these recipients and avoid emailing them in the future to maintain your reputation. | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


