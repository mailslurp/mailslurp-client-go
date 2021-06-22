# WebhookDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BasicAuth** | **bool** | Does webhook expect basic authentication? If true it means you created this webhook with a username and password. MailSlurp will use these in the URL to authenticate itself. | [optional] 
**CreatedAt** | [**time.Time**](time.Time) | When the webhook was created | [optional] 
**EventName** | **string** |  | [optional] 
**Id** | **string** | ID of the Webhook | [optional] 
**InboxId** | **string** | The inbox that the Webhook will be triggered by | [optional] 
**Method** | **string** | HTTP method that your server endpoint must listen for | [optional] 
**Name** | **string** | Name of the webhook | [optional] 
**PayloadJsonSchema** | **string** | JSON Schema for the payload that will be sent to your URL via the HTTP method described. | [optional] 
**UpdatedAt** | [**time.Time**](time.Time) |  | 
**Url** | **string** | URL of your server that the webhook will be sent to. The schema of the JSON that is sent is described by the payloadJsonSchema. | [optional] 
**UserId** | **string** | User ID of the Webhook | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


