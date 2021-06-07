# MailSlurp\WebhookControllerApi

All URIs are relative to *https://api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebhook**](WebhookControllerApi#CreateWebhook) | **Post** /inboxes/{inboxId}/webhooks | Attach a WebHook URL to an inbox
[**DeleteWebhook**](WebhookControllerApi#DeleteWebhook) | **Delete** /inboxes/{inboxId}/webhooks/{webhookId} | Delete and disable a Webhook for an Inbox
[**GetAllWebhooks**](WebhookControllerApi#GetAllWebhooks) | **Get** /webhooks/paginated | List Webhooks Paginated
[**GetTestWebhookPayload**](WebhookControllerApi#GetTestWebhookPayload) | **Get** /webhooks/test | Get test webhook payload example. Response content depends on eventName passed. Uses &#x60;EMAIL_RECEIVED&#x60; as default.
[**GetTestWebhookPayloadNewAttachment**](WebhookControllerApi#GetTestWebhookPayloadNewAttachment) | **Get** /webhooks/test/new-attachment-payload | Get webhook test payload for new attachment event
[**GetTestWebhookPayloadNewContact**](WebhookControllerApi#GetTestWebhookPayloadNewContact) | **Get** /webhooks/test/new-contact-payload | Get webhook test payload for new contact event
[**GetTestWebhookPayloadNewEmail**](WebhookControllerApi#GetTestWebhookPayloadNewEmail) | **Get** /webhooks/test/new-email-payload | Get webhook test payload for new email event
[**GetWebhook**](WebhookControllerApi#GetWebhook) | **Get** /webhooks/{webhookId} | Get a webhook for an Inbox
[**GetWebhooks**](WebhookControllerApi#GetWebhooks) | **Get** /inboxes/{inboxId}/webhooks | Get all webhooks for an Inbox
[**SendTestData**](WebhookControllerApi#SendTestData) | **Post** /webhooks/{webhookId}/test | Send webhook test data



## CreateWebhook

> WebhookDto CreateWebhook(ctx, inboxId, webhookOptions)

Attach a WebHook URL to an inbox

Get notified whenever an inbox receives an email via a WebHook URL. An emailID will be posted to this URL every time an email is received for this inbox. The URL must be publicly reachable by the MailSlurp server. You can provide basicAuth values if you wish to secure this endpoint.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()| inboxId | 
**webhookOptions** | [**CreateWebhookOptions**](CreateWebhookOptions)| webhookOptions | 

### Return type

[**WebhookDto**](WebhookDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DeleteWebhook

> DeleteWebhook(ctx, inboxId, webhookId)

Delete and disable a Webhook for an Inbox

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()| inboxId | 
**webhookId** | [**string**]()| webhookId | 

### Return type

 (empty response body)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetAllWebhooks

> PageWebhookProjection GetAllWebhooks(ctx, optional)

List Webhooks Paginated

List webhooks in paginated form. Allows for page index, page size, and sort direction.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllWebhooksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllWebhooksOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Optional page index in list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]

### Return type

[**PageWebhookProjection**](PageWebhookProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetTestWebhookPayload

> AbstractWebhookPayload GetTestWebhookPayload(ctx, optional)

Get test webhook payload example. Response content depends on eventName passed. Uses `EMAIL_RECEIVED` as default.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetTestWebhookPayloadOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTestWebhookPayloadOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventName** | **optional.String**| eventName | 

### Return type

[**AbstractWebhookPayload**](AbstractWebhookPayload)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetTestWebhookPayloadNewAttachment

> WebhookNewAttachmentPayload GetTestWebhookPayloadNewAttachment(ctx, )

Get webhook test payload for new attachment event

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**WebhookNewAttachmentPayload**](WebhookNewAttachmentPayload)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetTestWebhookPayloadNewContact

> WebhookNewContactPayload GetTestWebhookPayloadNewContact(ctx, )

Get webhook test payload for new contact event

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**WebhookNewContactPayload**](WebhookNewContactPayload)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetTestWebhookPayloadNewEmail

> WebhookNewEmailPayload GetTestWebhookPayloadNewEmail(ctx, )

Get webhook test payload for new email event

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**WebhookNewEmailPayload**](WebhookNewEmailPayload)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetWebhook

> WebhookDto GetWebhook(ctx, webhookId)

Get a webhook for an Inbox

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | [**string**]()| webhookId | 

### Return type

[**WebhookDto**](WebhookDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetWebhooks

> []WebhookDto GetWebhooks(ctx, inboxId)

Get all webhooks for an Inbox

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()| inboxId | 

### Return type

[**[]WebhookDto**](WebhookDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## SendTestData

> WebhookTestResult SendTestData(ctx, webhookId)

Send webhook test data

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | [**string**]()| webhookId | 

### Return type

[**WebhookTestResult**](WebhookTestResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

