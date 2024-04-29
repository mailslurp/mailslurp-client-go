# MailSlurp\BounceControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilterBouncedRecipient**](BounceControllerApi#FilterBouncedRecipient) | **Post** /bounce/filter-recipients | Filter a list of email recipients and remove those who have bounced
[**GetAccountBounceBlockStatus**](BounceControllerApi#GetAccountBounceBlockStatus) | **Get** /bounce/account-block | Can account send email
[**GetBouncedEmail**](BounceControllerApi#GetBouncedEmail) | **Get** /bounce/emails/{id} | Get a bounced email.
[**GetBouncedEmails**](BounceControllerApi#GetBouncedEmails) | **Get** /bounce/emails | Get paginated list of bounced emails.
[**GetBouncedRecipient**](BounceControllerApi#GetBouncedRecipient) | **Get** /bounce/recipients/{id} | Get a bounced email.
[**GetBouncedRecipients**](BounceControllerApi#GetBouncedRecipients) | **Get** /bounce/recipients | Get paginated list of bounced recipients.
[**GetComplaint**](BounceControllerApi#GetComplaint) | **Get** /bounce/complaints/{id} | Get complaint
[**GetComplaints**](BounceControllerApi#GetComplaints) | **Get** /bounce/complaints | Get paginated list of complaints.
[**GetListUnsubscribeRecipients**](BounceControllerApi#GetListUnsubscribeRecipients) | **Get** /bounce/list-unsubscribe-recipients | Get paginated list of unsubscribed recipients.



## FilterBouncedRecipient

> FilterBouncedRecipientsResult FilterBouncedRecipient(ctx, filterBouncedRecipientsOptions)

Filter a list of email recipients and remove those who have bounced

Prevent email sending errors by remove recipients who have resulted in past email bounces or complaints

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filterBouncedRecipientsOptions** | [**FilterBouncedRecipientsOptions**](FilterBouncedRecipientsOptions)|  | 

### Return type

[**FilterBouncedRecipientsResult**](FilterBouncedRecipientsResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetAccountBounceBlockStatus

> AccountBounceBlockDto GetAccountBounceBlockStatus(ctx, )

Can account send email

Check if account block status prevents sending

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**AccountBounceBlockDto**](AccountBounceBlockDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetBouncedEmail

> BouncedEmailDto GetBouncedEmail(ctx, id)

Get a bounced email.

Bounced emails are email you have sent that were rejected by a recipient

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()| ID of the bounced email to fetch | 

### Return type

[**BouncedEmailDto**](BouncedEmailDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

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
 **page** | **optional.Int32**| Optional page index | [default to 0]
 **size** | **optional.Int32**| Optional page size  | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **since** | **optional.Time**| Filter by created at after the given timestamp | 
 **before** | **optional.Time**| Filter by created at before the given timestamp | 

### Return type

[**PageBouncedEmail**](PageBouncedEmail)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetBouncedRecipient

> BouncedRecipientDto GetBouncedRecipient(ctx, id)

Get a bounced email.

Bounced emails are email you have sent that were rejected by a recipient

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()| ID of the bounced recipient | 

### Return type

[**BouncedRecipientDto**](BouncedRecipientDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

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
 **since** | **optional.Time**| Filter by created at after the given timestamp | 
 **before** | **optional.Time**| Filter by created at before the given timestamp | 

### Return type

[**PageBouncedRecipients**](PageBouncedRecipients)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetComplaint

> Complaint GetComplaint(ctx, id)

Get complaint

Get complaint

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()| ID of the complaint | 

### Return type

[**Complaint**](Complaint)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetComplaints

> PageComplaint GetComplaints(ctx, optional)

Get paginated list of complaints.

SMTP complaints made against your account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetComplaintsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetComplaintsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Optional page index  | [default to 0]
 **size** | **optional.Int32**| Optional page size  | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **since** | **optional.Time**| Filter by created at after the given timestamp | 
 **before** | **optional.Time**| Filter by created at before the given timestamp | 

### Return type

[**PageComplaint**](PageComplaint)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetListUnsubscribeRecipients

> PageListUnsubscribeRecipients GetListUnsubscribeRecipients(ctx, optional)

Get paginated list of unsubscribed recipients.

Unsubscribed recipient have unsubscribed from a mailing list for a user or domain and cannot be contacted again.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetListUnsubscribeRecipientsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetListUnsubscribeRecipientsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Optional page index | [default to 0]
 **size** | **optional.Int32**| Optional page size  | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **domainId** | [**optional.Interface of string**]()| Filter by domainId | 

### Return type

[**PageListUnsubscribeRecipients**](PageListUnsubscribeRecipients)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

