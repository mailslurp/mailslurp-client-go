# MailSlurp\WebhookControllerApi

All URIs are relative to *https://api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebhook**](WebhookControllerApi#CreateWebhook) | **Post** /inboxes/{inboxId}/webhooks | Attach a WebHook URL to an inbox
[**DeleteAllWebhooks**](WebhookControllerApi#DeleteAllWebhooks) | **Delete** /webhooks | Delete all webhooks
[**DeleteWebhook**](WebhookControllerApi#DeleteWebhook) | **Delete** /inboxes/{inboxId}/webhooks/{webhookId} | Delete and disable a Webhook for an Inbox
[**GetAllWebhookResults**](WebhookControllerApi#GetAllWebhookResults) | **Get** /webhooks/results | Get results for all webhooks
[**GetAllWebhooks**](WebhookControllerApi#GetAllWebhooks) | **Get** /webhooks/paginated | List Webhooks Paginated
[**GetInboxWebhooksPaginated**](WebhookControllerApi#GetInboxWebhooksPaginated) | **Get** /inboxes/{inboxId}/webhooks/paginated | Get paginated webhooks for an Inbox
[**GetJsonSchemaForWebhookPayload**](WebhookControllerApi#GetJsonSchemaForWebhookPayload) | **Post** /webhooks/{webhookId}/schema | Get JSON Schema definition for webhook payload
[**GetTestWebhookPayload**](WebhookControllerApi#GetTestWebhookPayload) | **Get** /webhooks/test | Get test webhook payload example. Response content depends on eventName passed. Uses &#x60;EMAIL_RECEIVED&#x60; as default.
[**GetTestWebhookPayloadEmailOpened**](WebhookControllerApi#GetTestWebhookPayloadEmailOpened) | **Get** /webhooks/test/email-opened-payload | Get webhook test payload for email opened event
[**GetTestWebhookPayloadEmailRead**](WebhookControllerApi#GetTestWebhookPayloadEmailRead) | **Get** /webhooks/test/email-read-payload | Get webhook test payload for email opened event
[**GetTestWebhookPayloadForWebhook**](WebhookControllerApi#GetTestWebhookPayloadForWebhook) | **Post** /webhooks/{webhookId}/example | Get example payload for webhook
[**GetTestWebhookPayloadNewAttachment**](WebhookControllerApi#GetTestWebhookPayloadNewAttachment) | **Get** /webhooks/test/new-attachment-payload | Get webhook test payload for new attachment event
[**GetTestWebhookPayloadNewContact**](WebhookControllerApi#GetTestWebhookPayloadNewContact) | **Get** /webhooks/test/new-contact-payload | Get webhook test payload for new contact event
[**GetTestWebhookPayloadNewEmail**](WebhookControllerApi#GetTestWebhookPayloadNewEmail) | **Get** /webhooks/test/new-email-payload | Get webhook test payload for new email event
[**GetWebhook**](WebhookControllerApi#GetWebhook) | **Get** /webhooks/{webhookId} | Get a webhook for an Inbox
[**GetWebhookResult**](WebhookControllerApi#GetWebhookResult) | **Get** /webhooks/results/{webhookResultId} | Get a webhook result for a webhook
[**GetWebhookResults**](WebhookControllerApi#GetWebhookResults) | **Get** /webhooks/{webhookId}/results | Get a webhook results for a webhook
[**GetWebhookResultsUnseenErrorCount**](WebhookControllerApi#GetWebhookResultsUnseenErrorCount) | **Get** /webhooks/results/unseen-count | Get count of unseen webhook results with error status
[**GetWebhooks**](WebhookControllerApi#GetWebhooks) | **Get** /inboxes/{inboxId}/webhooks | Get all webhooks for an Inbox
[**RedriveWebhookResult**](WebhookControllerApi#RedriveWebhookResult) | **Post** /webhooks/results/{webhookResultId}/redrive | Get a webhook result and try to resend the original webhook payload
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


## DeleteAllWebhooks

> DeleteAllWebhooks(ctx, )

Delete all webhooks

### Required Parameters

This endpoint does not need any parameter.

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


## GetAllWebhookResults

> PageWebhookResult GetAllWebhookResults(ctx, optional)

Get results for all webhooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllWebhookResultsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllWebhookResultsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **before** | **optional.Time**| Filter by created at before the given timestamp | 
 **page** | **optional.Int32**| Optional page index in list pagination | [default to 0]
 **searchFilter** | **optional.String**| Optional search filter | 
 **since** | **optional.Time**| Filter by created at after the given timestamp | 
 **size** | **optional.Int32**| Optional page size in list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **unseenOnly** | **optional.Bool**| Filter for unseen exceptions only | 

### Return type

[**PageWebhookResult**](PageWebhookResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
 **before** | **optional.Time**| Filter by created at before the given timestamp | 
 **page** | **optional.Int32**| Optional page index in list pagination | [default to 0]
 **searchFilter** | **optional.String**| Optional search filter | 
 **since** | **optional.Time**| Filter by created at after the given timestamp | 
 **size** | **optional.Int32**| Optional page size for paginated result list. | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to DESC]

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


## GetInboxWebhooksPaginated

> PageWebhookProjection GetInboxWebhooksPaginated(ctx, inboxId, optional)

Get paginated webhooks for an Inbox

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()| inboxId | 
 **optional** | ***GetInboxWebhooksPaginatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetInboxWebhooksPaginatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **before** | **optional.Time**| Filter by created at before the given timestamp | 
 **page** | **optional.Int32**| Optional page index in list pagination | [default to 0]
 **searchFilter** | **optional.String**| Optional search filter | 
 **since** | **optional.Time**| Filter by created at after the given timestamp | 
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


## GetJsonSchemaForWebhookPayload

> JsonSchemaDto GetJsonSchemaForWebhookPayload(ctx, webhookId)

Get JSON Schema definition for webhook payload

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | [**string**]()| webhookId | 

### Return type

[**JsonSchemaDto**](JSONSchemaDto)

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


## GetTestWebhookPayloadEmailOpened

> WebhookEmailOpenedPayload GetTestWebhookPayloadEmailOpened(ctx, )

Get webhook test payload for email opened event

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**WebhookEmailOpenedPayload**](WebhookEmailOpenedPayload)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetTestWebhookPayloadEmailRead

> WebhookEmailReadPayload GetTestWebhookPayloadEmailRead(ctx, )

Get webhook test payload for email opened event

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**WebhookEmailReadPayload**](WebhookEmailReadPayload)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetTestWebhookPayloadForWebhook

> AbstractWebhookPayload GetTestWebhookPayloadForWebhook(ctx, webhookId)

Get example payload for webhook

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | [**string**]()| webhookId | 

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


## GetWebhookResult

> WebhookResultDto GetWebhookResult(ctx, webhookResultId)

Get a webhook result for a webhook

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookResultId** | [**string**]()| Webhook Result ID | 

### Return type

[**WebhookResultDto**](WebhookResultDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetWebhookResults

> PageWebhookResult GetWebhookResults(ctx, webhookId, optional)

Get a webhook results for a webhook

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | [**string**]()| ID of webhook to get results for | 
 **optional** | ***GetWebhookResultsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetWebhookResultsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **before** | **optional.Time**| Filter by created at before the given timestamp | 
 **page** | **optional.Int32**| Optional page index in list pagination | [default to 0]
 **searchFilter** | **optional.String**| Optional search filter | 
 **since** | **optional.Time**| Filter by created at after the given timestamp | 
 **size** | **optional.Int32**| Optional page size in list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **unseenOnly** | **optional.Bool**| Filter for unseen exceptions only | 

### Return type

[**PageWebhookResult**](PageWebhookResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetWebhookResultsUnseenErrorCount

> UnseenErrorCountDto GetWebhookResultsUnseenErrorCount(ctx, inboxId)

Get count of unseen webhook results with error status

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()| inboxId | 

### Return type

[**UnseenErrorCountDto**](UnseenErrorCountDto)

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


## RedriveWebhookResult

> WebhookRedriveResult RedriveWebhookResult(ctx, webhookResultId)

Get a webhook result and try to resend the original webhook payload

Allows you to resend a webhook payload that was already sent. Webhooks that fail are retried automatically for 24 hours and then put in a dead letter queue. You can retry results manually using this method.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookResultId** | [**string**]()| Webhook Result ID | 

### Return type

[**WebhookRedriveResult**](WebhookRedriveResult)

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

