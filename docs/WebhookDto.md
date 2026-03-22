# WebhookDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the Webhook | 
**UserId** | **string** | User ID of the Webhook | 
**BasicAuth** | **bool** | Does webhook expect basic authentication? If true it means you created this webhook with a username and password. MailSlurp will use these in the URL to authenticate itself. | 
**Name** | Pointer to **string** | Name of the webhook | [optional] 
**PhoneId** | Pointer to **string** | The phoneNumberId that the Webhook will be triggered by. If null then webhook triggered at account level or inbox level if inboxId set | [optional] 
**InboxId** | Pointer to **string** | The inbox that the Webhook will be triggered by. If null then webhook triggered at account level or phone level if phoneId set | [optional] 
**RequestBodyTemplate** | Pointer to **string** | Request body template for HTTP request that will be sent for the webhook. Use Moustache style template variables to insert values from the original event payload. | [optional] 
**Url** | **string** | URL of your server that the webhook will be sent to. The schema of the JSON that is sent is described by the payloadJsonSchema. | 
**Method** | **string** | HTTP method that your server endpoint must listen for | 
**PayloadJsonSchema** | **string** | Deprecated. Fetch JSON Schema for webhook using the getJsonSchemaForWebhookPayload method | 
**CreatedAt** | Pointer to [**time.Time**](time.Time) | When the webhook was created | 
**UpdatedAt** | [**time.Time**](time.Time) |  | 
**EventName** | Pointer to **string** | Webhook trigger event name | [optional] 
**RequestHeaders** | [**WebhookHeaders**](WebhookHeaders) |  | [optional] 
**AiTransformId** | Pointer to **string** | ID of AI transformer for payload | [optional] 
**IgnoreInsecureSslCertificates** | Pointer to **bool** | Should notifier ignore insecure SSL certificates | [optional] 
**UseStaticIpRange** | Pointer to **bool** | Should notifier use static IP range when sending webhook payload | [optional] 
**HealthStatus** | Pointer to **string** | Webhook health | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


