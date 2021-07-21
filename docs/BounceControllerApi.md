# MailSlurp\BounceControllerApi

All URIs are relative to *https://api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBouncedEmail**](BounceControllerApi#GetBouncedEmail) | **Get** /bounce/emails/{id} | Get a bounced email.
[**GetBouncedEmails**](BounceControllerApi#GetBouncedEmails) | **Get** /bounce/emails | Get paginated list of bounced emails.
[**GetBouncedRecipient**](BounceControllerApi#GetBouncedRecipient) | **Get** /bounce/recipients/{id} | Get a bounced email.
[**GetBouncedRecipients**](BounceControllerApi#GetBouncedRecipients) | **Get** /bounce/recipients | Get paginated list of bounced recipients.



## GetBouncedEmail

> Bounce GetBouncedEmail(ctx, id)

Get a bounced email.

Bounced emails are email you have sent that were rejected by a recipient

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()| ID of the bounced email to fetch | 

### Return type

[**Bounce**](Bounce)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetBouncedEmails

> PageBouncedEmail GetBouncedEmails(ctx, optional)

Get paginated list of bounced emails.

Bounced emails are email you have sent that were rejected by a recipient

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetBouncedEmailsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetBouncedEmailsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Optional page index  | [default to 0]
 **size** | **optional.Int32**| Optional page size  | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]

### Return type

[**PageBouncedEmail**](PageBouncedEmail)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetBouncedRecipient

> BounceRecipient GetBouncedRecipient(ctx, id)

Get a bounced email.

Bounced emails are email you have sent that were rejected by a recipient

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()| ID of the bounced recipient | 

### Return type

[**BounceRecipient**](BounceRecipient)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetBouncedRecipients

> PageBouncedRecipients GetBouncedRecipients(ctx, optional)

Get paginated list of bounced recipients.

Bounced recipients are email addresses that you have sent emails to that did not accept the sent email. Once a recipient is bounced you cannot send emails to that address.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetBouncedRecipientsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetBouncedRecipientsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Optional page index  | [default to 0]
 **size** | **optional.Int32**| Optional page size  | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]

### Return type

[**PageBouncedRecipients**](PageBouncedRecipients)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

