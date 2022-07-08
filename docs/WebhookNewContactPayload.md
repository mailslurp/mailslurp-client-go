# WebhookNewContactPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | **string** | Idempotent message ID. Store this ID locally or in a database to prevent message duplication. | 
**WebhookId** | **string** | ID of webhook entity being triggered | 
**WebhookName** | **string** | Name of the webhook being triggered | [optional] 
**EventName** | **string** | Name of the event type webhook is being triggered for. | 
**ContactId** | **string** |  | 
**GroupId** | **string** |  | [optional] 
**FirstName** | **string** |  | [optional] 
**LastName** | **string** |  | [optional] 
**Company** | **string** |  | [optional] 
**PrimaryEmailAddress** | **string** |  | [optional] 
**EmailAddresses** | **[]string** |  | 
**Tags** | **[]string** |  | 
**MetaData** | [**map[string]interface{}**]() |  | [optional] 
**OptOut** | **bool** |  | 
**CreatedAt** | [**time.Time**](time.Time) |  | 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


