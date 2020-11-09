# MailSlurp\SentEmailsControllerApi

All URIs are relative to *https://api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSentEmail**](SentEmailsControllerApi.md#GetSentEmail) | **Get** /sent/{id} | Get sent email receipt
[**GetSentEmails**](SentEmailsControllerApi.md#GetSentEmails) | **Get** /sent | Get all sent emails in paginated form



## GetSentEmail

> SentEmailDto GetSentEmail(ctx, id)

Get sent email receipt

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| id | 

### Return type

[**SentEmailDto**](SentEmailDto.md)

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSentEmails

> PageSentEmailProjection GetSentEmails(ctx, optional)

Get all sent emails in paginated form

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetSentEmailsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSentEmailsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inboxId** | [**optional.Interface of string**](.md)| Optional inboxId to filter sender of sent emails by | 
 **page** | **optional.Int32**| Optional page index in inbox sent email list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in inbox sent email list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]

### Return type

[**PageSentEmailProjection**](Page«SentEmailProjection».md)

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

