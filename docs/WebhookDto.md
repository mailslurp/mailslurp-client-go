# WebhookDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the Webhook | 
**UserId** | **string** | User ID of the Webhook | 
**BasicAuth** | **bool** | Does webhook expect basic authentication? If true it means you created this webhook with a username and password. MailSlurp will use these in the URL to authenticate itself. | 
**Name** | **string** | Name of the webhook | [optional] 
**InboxId** | **string** | The inbox that the Webhook will be triggered by. If null then webhook triggered at account level | [optional] 
**Url** | **string** | URL of your server that the webhook will be sent to. The schema of the JSON that is sent is described by the payloadJsonSchema. | 
**Method** | **string** | HTTP method that your server endpoint must listen for | 
**PayloadJsonSchema** | **string** | Deprecated. Fetch JSON Schema for webhook using the getJsonSchemaForWebhookPayload method | 
**CreatedAt** | [**time.Time**](time.Time) | When the webhook was created | 
**UpdatedAt** | [**time.Time**](time.Time) |  | 
**EventName** | **string** |  | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


