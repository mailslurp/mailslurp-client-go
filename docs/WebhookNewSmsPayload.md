# WebhookNewSmsPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | **string** | Idempotent message ID. Store this ID locally or in a database to prevent message duplication. | 
**WebhookId** | **string** | ID of webhook entity being triggered | 
**EventName** | **string** | Name of the event type webhook is being triggered for. | 
**WebhookName** | Pointer to **string** | Name of the webhook being triggered | [optional] 
**SmsId** | **string** | ID of SMS message | 
**UserId** | **string** | User ID of event | 
**PhoneNumber** | **string** | ID of phone number receiving SMS | 
**ToNumber** | **string** | Recipient phone number | 
**FromNumber** | **string** | Sender phone number | 
**Body** | **string** | SMS message body | 
**Read** | **bool** | SMS has been read | 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


