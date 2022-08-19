# WebhookDeliveryStatusPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | **string** | Idempotent message ID. Store this ID locally or in a database to prevent message duplication. | 
**WebhookId** | **string** | ID of webhook entity being triggered | 
**EventName** | **string** | Name of the event type webhook is being triggered for. | 
**WebhookName** | **string** | Name of the webhook being triggered | [optional] 
**Id** | **string** | ID of delivery status | 
**UserId** | **string** | User ID of event | 
**SentId** | Pointer to **string** | ID of sent email | [optional] 
**RemoteMtaIp** | Pointer to **string** | IP address of the remote Mail Transfer Agent | [optional] 
**InboxId** | Pointer to **string** | Id of the inbox | [optional] 
**ReportingMta** | Pointer to **string** | Mail Transfer Agent reporting delivery status | [optional] 
**Recipients** | Pointer to **[]string** | Recipients for delivery | [optional] 
**SmtpResponse** | Pointer to **string** | SMTP server response message | [optional] 
**SmtpStatusCode** | Pointer to **int32** | SMTP server status | [optional] 
**ProcessingTimeMillis** | Pointer to **int64** | Time in milliseconds for delivery processing | [optional] 
**Received** | Pointer to [**time.Time**](time.Time) | Time event was received | [optional] 
**Subject** | Pointer to **string** | Email subject | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


