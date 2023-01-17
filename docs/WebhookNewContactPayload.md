# WebhookNewContactPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | **string** | Idempotent message ID. Store this ID locally or in a database to prevent message duplication. | 
**WebhookId** | **string** | ID of webhook entity being triggered | 
**WebhookName** | Pointer to **string** | Name of the webhook being triggered | [optional] 
**EventName** | **string** | Name of the event type webhook is being triggered for. | 
**ContactId** | **string** | Contact ID | 
**GroupId** | Pointer to **string** | Contact group ID | [optional] 
**FirstName** | Pointer to **string** | Contact first name | [optional] 
**LastName** | Pointer to **string** | Contact last name | [optional] 
**Company** | Pointer to **string** | Contact company name | [optional] 
**PrimaryEmailAddress** | Pointer to **string** | Primary email address for contact | [optional] 
**EmailAddresses** | **[]string** | Email addresses for contact | 
**Tags** | **[]string** | Tags for contact | 
**MetaData** | Pointer to [**map[string]interface{}**]() |  | [optional] 
**OptOut** | **bool** | Has contact opted out of emails | 
**CreatedAt** | [**time.Time**](time.Time) | Date time of event creation | 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


