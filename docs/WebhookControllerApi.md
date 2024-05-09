# MailSlurp\WebhookControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountWebhook**](WebhookControllerApi#CreateAccountWebhook) | **Post** /webhooks | Attach a WebHook URL to an inbox
[**CreateWebhook**](WebhookControllerApi#CreateWebhook) | **Post** /inboxes/{inboxId}/webhooks | Attach a WebHook URL to an inbox
[**CreateWebhookForPhoneNumber**](WebhookControllerApi#CreateWebhookForPhoneNumber) | **Post** /phone/numbers/{phoneNumberId}/webhooks | Attach a WebHook URL to a phone number
[**DeleteAllWebhooks**](WebhookControllerApi#DeleteAllWebhooks) | **Delete** /webhooks | Delete all webhooks
[**DeleteWebhook**](WebhookControllerApi#DeleteWebhook) | **Delete** /inboxes/{inboxId}/webhooks/{webhookId} | Delete and disable a Webhook for an Inbox
[**DeleteWebhookById**](WebhookControllerApi#DeleteWebhookById) | **Delete** /webhooks/{webhookId} | Delete a webhook
[**GetAllAccountWebhooks**](WebhookControllerApi#GetAllAccountWebhooks) | **Get** /webhooks/account/paginated | List account webhooks Paginated
[**GetAllWebhookResults**](WebhookControllerApi#GetAllWebhookResults) | **Get** /webhooks/results | Get results for all webhooks
[**GetAllWebhooks**](WebhookControllerApi#GetAllWebhooks) | **Get** /webhooks/paginated | List Webhooks Paginated
[**GetInboxWebhooksPaginated**](WebhookControllerApi#GetInboxWebhooksPaginated) | **Get** /inboxes/{inboxId}/webhooks/paginated | Get paginated webhooks for an Inbox
[**GetJsonSchemaForWebhookEvent**](WebhookControllerApi#GetJsonSchemaForWebhookEvent) | **Post** /webhooks/schema | 
[**GetJsonSchemaForWebhookPayload**](WebhookControllerApi#GetJsonSchemaForWebhookPayload) | **Post** /webhooks/{webhookId}/schema | 
[**GetPhoneNumberWebhooksPaginated**](WebhookControllerApi#GetPhoneNumberWebhooksPaginated) | **Get** /phone/numbers/{phoneId}/webhooks/paginated | Get paginated webhooks for a phone number
[**GetTestWebhookPayload**](WebhookControllerApi#GetTestWebhookPayload) | **Get** /webhooks/test | 
[**GetTestWebhookPayloadBounce**](WebhookControllerApi#GetTestWebhookPayloadBounce) | **Get** /webhooks/test/email-bounce-payload | 
[**GetTestWebhookPayloadBounceRecipient**](WebhookControllerApi#GetTestWebhookPayloadBounceRecipient) | **Get** /webhooks/test/email-bounce-recipient-payload | 
[**GetTestWebhookPayloadDeliveryStatus**](WebhookControllerApi#GetTestWebhookPayloadDeliveryStatus) | **Get** /webhooks/test/delivery-status-payload | Get webhook test payload for delivery status event
[**GetTestWebhookPayloadEmailOpened**](WebhookControllerApi#GetTestWebhookPayloadEmailOpened) | **Get** /webhooks/test/email-opened-payload | 
[**GetTestWebhookPayloadEmailRead**](WebhookControllerApi#GetTestWebhookPayloadEmailRead) | **Get** /webhooks/test/email-read-payload | 
[**GetTestWebhookPayloadForWebhook**](WebhookControllerApi#GetTestWebhookPayloadForWebhook) | **Post** /webhooks/{webhookId}/example | 
[**GetTestWebhookPayloadNewAttachment**](WebhookControllerApi#GetTestWebhookPayloadNewAttachment) | **Get** /webhooks/test/new-attachment-payload | Get webhook test payload for new attachment event
[**GetTestWebhookPayloadNewContact**](WebhookControllerApi#GetTestWebhookPayloadNewContact) | **Get** /webhooks/test/new-contact-payload | Get webhook test payload for new contact event
[**GetTestWebhookPayloadNewEmail**](WebhookControllerApi#GetTestWebhookPayloadNewEmail) | **Get** /webhooks/test/new-email-payload | Get webhook test payload for new email event
[**GetTestWebhookPayloadNewSms**](WebhookControllerApi#GetTestWebhookPayloadNewSms) | **Get** /webhooks/test/new-sms-payload | Get webhook test payload for new sms event
[**GetWebhook**](WebhookControllerApi#GetWebhook) | **Get** /webhooks/{webhookId} | Get a webhook
[**GetWebhookResult**](WebhookControllerApi#GetWebhookResult) | **Get** /webhooks/results/{webhookResultId} | Get a webhook result for a webhook
[**GetWebhookResults**](WebhookControllerApi#GetWebhookResults) | **Get** /webhooks/{webhookId}/results | Get a webhook results for a webhook
[**GetWebhookResultsCount**](WebhookControllerApi#GetWebhookResultsCount) | **Get** /webhooks/{webhookId}/results/count | Get a webhook results count for a webhook
[**GetWebhookResultsUnseenErrorCount**](WebhookControllerApi#GetWebhookResultsUnseenErrorCount) | **Get** /webhooks/results/unseen-count | Get count of unseen webhook results with error status
[**GetWebhooks**](WebhookControllerApi#GetWebhooks) | **Get** /inboxes/{inboxId}/webhooks | Get all webhooks for an Inbox
[**RedriveAllWebhookResults**](WebhookControllerApi#RedriveAllWebhookResults) | **Post** /webhooks/results/redrive | Redrive all webhook results that have failed status
[**RedriveWebhookResult**](WebhookControllerApi#RedriveWebhookResult) | **Post** /webhooks/results/{webhookResultId}/redrive | Get a webhook result and try to resend the original webhook payload
[**SendTestData**](WebhookControllerApi#SendTestData) | **Post** /webhooks/{webhookId}/test | Send webhook test data
[**UpdateWebhookHeaders**](WebhookControllerApi#UpdateWebhookHeaders) | **Put** /webhooks/{webhookId}/headers | Update a webhook request headers
[**VerifyWebhookSignature**](WebhookControllerApi#VerifyWebhookSignature) | **Post** /webhooks/verify | Verify a webhook payload signature
[**WaitForWebhookResults**](WebhookControllerApi#WaitForWebhookResults) | **Get** /webhooks/{webhookId}/wait | Wait for webhook results for a webhook



## CreateAccountWebhook

> WebhookDto CreateAccountWebhook(ctx, createWebhookOptions)

Attach a WebHook URL to an inbox

Get notified of account level events such as bounce and bounce recipient.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createWebhookOptions** | [**CreateWebhookOptions**](CreateWebhookOptions)|  | 

### Return type

[**WebhookDto**](WebhookDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CreateWebhook

> WebhookDto CreateWebhook(ctx, inboxId, createWebhookOptions)

Attach a WebHook URL to an inbox

Get notified whenever an inbox receives an email via a WebHook URL. An emailID will be posted to this URL every time an email is received for this inbox. The URL must be publicly reachable by the MailSlurp server. You can provide basicAuth values if you wish to secure this endpoint.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()|  | 
**createWebhookOptions** | [**CreateWebhookOptions**](CreateWebhookOptions)|  | 

### Return type

[**WebhookDto**](WebhookDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CreateWebhookForPhoneNumber

> WebhookDto CreateWebhookForPhoneNumber(ctx, phoneNumberId, createWebhookOptions)

Attach a WebHook URL to a phone number

Get notified whenever a phone number receives an SMS via a WebHook URL.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneNumberId** | [**string**]()|  | 
**createWebhookOptions** | [**CreateWebhookOptions**](CreateWebhookOptions)|  | 

### Return type

[**WebhookDto**](WebhookDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DeleteAllWebhooks

> DeleteAllWebhooks(ctx, optional)

Delete all webhooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteAllWebhooksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteAllWebhooksOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **before** | **optional.Time**| before | 

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
**inboxId** | [**string**]()|  | 
**webhookId** | [**string**]()|  | 

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


## DeleteWebhookById

> DeleteWebhookById(ctx, webhookId)

Delete a webhook

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | [**string**]()|  | 

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


## GetAllAccountWebhooks

> PageWebhookProjection GetAllAccountWebhooks(ctx, optional)

List account webhooks Paginated

List account webhooks in paginated form. Allows for page index, page size, and sort direction.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllAccountWebhooksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllAccountWebhooksOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Optional page index in list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size for paginated result list. | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to DESC]
 **eventType** | **optional.String**| Optional event type | 
 **since** | **optional.Time**| Filter by created at after the given timestamp | 
 **before** | **optional.Time**| Filter by created at before the given timestamp | 

### Return type

[**PageWebhookProjection**](PageWebhookProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

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
 **page** | **optional.Int32**| Optional page index in list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **searchFilter** | **optional.String**| Optional search filter | 
 **since** | **optional.Time**| Filter by created at after the given timestamp | 
 **before** | **optional.Time**| Filter by created at before the given timestamp | 
 **unseenOnly** | **optional.Bool**| Filter for unseen exceptions only | 
 **resultType** | **optional.String**| Filter by result type | 
 **eventName** | **optional.String**| Filter by event name | 
 **minStatusCode** | **optional.Int32**| Minimum response status | 
 **maxStatusCode** | **optional.Int32**| Maximum response status | 
 **inboxId** | [**optional.Interface of string**]()| Inbox ID | 
 **smsId** | [**optional.Interface of string**]()| Sms ID | 
 **attachmentId** | [**optional.Interface of string**]()| Attachment ID | 
 **emailId** | [**optional.Interface of string**]()| Email ID | 
 **phoneId** | [**optional.Interface of string**]()| Phone ID | 

### Return type

[**PageWebhookResult**](PageWebhookResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

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
 **size** | **optional.Int32**| Optional page size for paginated result list. | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to DESC]
 **searchFilter** | **optional.String**| Optional search filter | 
 **since** | **optional.Time**| Filter by created at after the given timestamp | 
 **inboxId** | [**optional.Interface of string**]()| Filter by inboxId | 
 **phoneId** | [**optional.Interface of string**]()| Filter by phoneId | 
 **before** | **optional.Time**| Filter by created at before the given timestamp | 

### Return type

[**PageWebhookProjection**](PageWebhookProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

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
**inboxId** | [**string**]()|  | 
 **optional** | ***GetInboxWebhooksPaginatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetInboxWebhooksPaginatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Optional page index in list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **searchFilter** | **optional.String**| Optional search filter | 
 **since** | **optional.Time**| Filter by created at after the given timestamp | 
 **before** | **optional.Time**| Filter by created at before the given timestamp | 

### Return type

[**PageWebhookProjection**](PageWebhookProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetJsonSchemaForWebhookEvent

> JsonSchemaDto GetJsonSchemaForWebhookEvent(ctx, event)



Get JSON Schema definition for webhook payload by event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**event** | **string**|  | 

### Return type

[**JsonSchemaDto**](JSONSchemaDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

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
**webhookId** | [**string**]()|  | 

### Return type

[**JsonSchemaDto**](JSONSchemaDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetPhoneNumberWebhooksPaginated

> PageWebhookProjection GetPhoneNumberWebhooksPaginated(ctx, phoneId, optional)

Get paginated webhooks for a phone number

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneId** | [**string**]()|  | 
 **optional** | ***GetPhoneNumberWebhooksPaginatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPhoneNumberWebhooksPaginatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Optional page index in list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **since** | **optional.Time**| Filter by created at after the given timestamp | 
 **before** | **optional.Time**| Filter by created at before the given timestamp | 

### Return type

[**PageWebhookProjection**](PageWebhookProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

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
 **eventName** | **optional.String**|  | 

### Return type

[**AbstractWebhookPayload**](AbstractWebhookPayload)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetTestWebhookPayloadBounce

> WebhookBouncePayload GetTestWebhookPayloadBounce(ctx, )



Get webhook test payload for bounce

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**WebhookBouncePayload**](WebhookBouncePayload)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetTestWebhookPayloadBounceRecipient

> WebhookBounceRecipientPayload GetTestWebhookPayloadBounceRecipient(ctx, )



Get webhook test payload for bounce recipient

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**WebhookBounceRecipientPayload**](WebhookBounceRecipientPayload)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetTestWebhookPayloadDeliveryStatus

> WebhookDeliveryStatusPayload GetTestWebhookPayloadDeliveryStatus(ctx, )

Get webhook test payload for delivery status event

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**WebhookDeliveryStatusPayload**](WebhookDeliveryStatusPayload)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

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
- **Accept**: */*

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
- **Accept**: */*

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
**webhookId** | [**string**]()|  | 

### Return type

[**AbstractWebhookPayload**](AbstractWebhookPayload)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

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
- **Accept**: */*

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
- **Accept**: */*

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
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetTestWebhookPayloadNewSms

> WebhookNewSmsPayload GetTestWebhookPayloadNewSms(ctx, )

Get webhook test payload for new sms event

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**WebhookNewSmsPayload**](WebhookNewSmsPayload)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetWebhook

> WebhookDto GetWebhook(ctx, webhookId)

Get a webhook

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | [**string**]()|  | 

### Return type

[**WebhookDto**](WebhookDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

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
- **Accept**: */*

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

 **page** | **optional.Int32**| Optional page index in list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **searchFilter** | **optional.String**| Optional search filter | 
 **since** | **optional.Time**| Filter by created at after the given timestamp | 
 **before** | **optional.Time**| Filter by created at before the given timestamp | 
 **unseenOnly** | **optional.Bool**| Filter for unseen exceptions only | 
 **resultType** | **optional.String**| Filter by result type | 
 **eventName** | **optional.String**| Filter by event name | 
 **minStatusCode** | **optional.Int32**| Minimum response status | 
 **maxStatusCode** | **optional.Int32**| Maximum response status | 
 **inboxId** | [**optional.Interface of string**]()| Inbox ID | 
 **smsId** | [**optional.Interface of string**]()| Sms ID | 
 **attachmentId** | [**optional.Interface of string**]()| Attachment ID | 
 **emailId** | [**optional.Interface of string**]()| Email ID | 
 **phoneId** | [**optional.Interface of string**]()| Phone ID | 

### Return type

[**PageWebhookResult**](PageWebhookResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetWebhookResultsCount

> CountDto GetWebhookResultsCount(ctx, webhookId)

Get a webhook results count for a webhook

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | [**string**]()| ID of webhook to get results for | 

### Return type

[**CountDto**](CountDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetWebhookResultsUnseenErrorCount

> UnseenErrorCountDto GetWebhookResultsUnseenErrorCount(ctx, )

Get count of unseen webhook results with error status

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**UnseenErrorCountDto**](UnseenErrorCountDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

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
**inboxId** | [**string**]()|  | 

### Return type

[**[]WebhookDto**](WebhookDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## RedriveAllWebhookResults

> WebhookRedriveAllResult RedriveAllWebhookResults(ctx, )

Redrive all webhook results that have failed status

Allows you to resend webhook payloads for any recorded webhook result that failed to deliver the payload.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**WebhookRedriveAllResult**](WebhookRedriveAllResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

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
- **Accept**: */*

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
**webhookId** | [**string**]()|  | 

### Return type

[**WebhookTestResult**](WebhookTestResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## UpdateWebhookHeaders

> WebhookDto UpdateWebhookHeaders(ctx, webhookId, webhookHeaders)

Update a webhook request headers

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | [**string**]()|  | 
**webhookHeaders** | [**WebhookHeaders**](WebhookHeaders)|  | 

### Return type

[**WebhookDto**](WebhookDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## VerifyWebhookSignature

> VerifyWebhookSignatureResults VerifyWebhookSignature(ctx, verifyWebhookSignatureOptions)

Verify a webhook payload signature

Verify a webhook payload using the messageId and signature. This allows you to be sure that MailSlurp sent the payload and not another server.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**verifyWebhookSignatureOptions** | [**VerifyWebhookSignatureOptions**](VerifyWebhookSignatureOptions)|  | 

### Return type

[**VerifyWebhookSignatureResults**](VerifyWebhookSignatureResults)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## WaitForWebhookResults

> []WebhookResultDto WaitForWebhookResults(ctx, webhookId, expectedCount, timeout)

Wait for webhook results for a webhook

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | [**string**]()| ID of webhook to get results for | 
**expectedCount** | **int32**| Expected result count | 
**timeout** | **int32**| Max time to wait in milliseconds | 

### Return type

[**[]WebhookResultDto**](WebhookResultDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

