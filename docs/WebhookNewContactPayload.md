# WebhookNewContactPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | **string** | Idempotent message ID. Store this ID locally or in a database to prevent message duplication. | [optional] 
**WebhookId** | **string** | ID of webhook entity being triggered | [optional] 
**WebhookName** | **string** | Name of the webhook being triggered | [optional] 
**EventName** | **string** | Name of the event type webhook is being triggered for. | [optional] 
**ContactId** | **string** |  | [optional] 
**GroupId** | **string** |  | [optional] 
**FirstName** | **string** |  | [optional] 
**LastName** | **string** |  | [optional] 
**Company** | **string** |  | [optional] 
**PrimaryEmailAddress** | **string** |  | [optional] 
**EmailAddresses** | **[]string** |  | [optional] 
**Tags** | **[]string** |  | [optional] 
**MetaData** | [**map[string]interface{}**]() |  | [optional] 
**OptOut** | **bool** |  | [optional] 
**CreatedAt** | [**time.Time**](time.Time) |  | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


