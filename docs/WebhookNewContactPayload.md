# WebhookNewContactPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | **string** |  | [optional] 
**ContactId** | **string** |  | 
**CreatedAt** | [**time.Time**](time.Time) |  | 
**EmailAddresses** | **[]string** |  | 
**EventName** | **string** | Name of the event type webhook is being triggered for. | [optional] 
**FirstName** | **string** |  | [optional] 
**GroupId** | **string** |  | [optional] 
**LastName** | **string** |  | [optional] 
**MessageId** | **string** | Idempotent message ID. Store this ID locally or in a database to prevent message duplication. | [optional] 
**MetaData** | [**map[string]interface{}**]() |  | [optional] 
**OptOut** | **bool** |  | [optional] 
**PrimaryEmailAddress** | **string** |  | [optional] 
**Tags** | **[]string** |  | 
**WebhookId** | **string** | ID of webhook entity being triggered | [optional] 
**WebhookName** | **string** | Name of the webhook being triggered | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


