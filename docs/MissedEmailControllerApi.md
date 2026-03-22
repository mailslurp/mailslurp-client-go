# MailSlurp\MissedEmailControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllMissedEmails**](MissedEmailControllerApi#GetAllMissedEmails) | **Get** /missed-emails | Get all MissedEmails in paginated format
[**GetAllUnknownMissedEmails**](MissedEmailControllerApi#GetAllUnknownMissedEmails) | **Get** /missed-emails/unknown | Get all unknown missed emails in paginated format
[**GetMissedEmail**](MissedEmailControllerApi#GetMissedEmail) | **Get** /missed-emails/{missedEmailId} | Get MissedEmail
[**RestoreMissedEmails**](MissedEmailControllerApi#RestoreMissedEmails) | **Post** /missed-emails/restore | Restore missed emails
[**WaitForNthMissedEmail**](MissedEmailControllerApi#WaitForNthMissedEmail) | **Get** /missed-emails/waitForNthMissedEmail | Wait for Nth missed email



## GetAllMissedEmails

> PageMissedEmailProjection GetAllMissedEmails(ctx, optional)

Get all MissedEmails in paginated format

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllMissedEmailsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllMissedEmailsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Optional page index in list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **searchFilter** | **optional.String**| Optional search filter | 
 **since** | **optional.Time**| Filter by created at after the given timestamp | 
 **before** | **optional.Time**| Filter by created at before the given timestamp | 
 **inboxId** | [**optional.Interface of string**]()| Optional inbox ID filter | 

### Return type

[**PageMissedEmailProjection**](PageMissedEmailProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetAllUnknownMissedEmails

> PageUnknownMissedEmailProjection GetAllUnknownMissedEmails(ctx, optional)

Get all unknown missed emails in paginated format

Unknown missed emails are emails that were sent to MailSlurp but could not be assigned to an existing inbox.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllUnknownMissedEmailsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllUnknownMissedEmailsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Optional page index in list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **searchFilter** | **optional.String**| Optional search filter | 
 **since** | **optional.Time**| Filter by created at after the given timestamp | 
 **before** | **optional.Time**| Filter by created at before the given timestamp | 
 **inboxId** | [**optional.Interface of string**]()| Optional inbox ID filter | 

### Return type

[**PageUnknownMissedEmailProjection**](PageUnknownMissedEmailProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetMissedEmail

> MissedEmailDto GetMissedEmail(ctx, missedEmailId)

Get MissedEmail

List emails that were missed due to plan limits.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**missedEmailId** | [**string**]()|  | 

### Return type

[**MissedEmailDto**](MissedEmailDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## RestoreMissedEmails

> RestoreMissedEmails(ctx, )

Restore missed emails

If emails were missed due to a plan limit they are saved as missed emails. If support team enables the canRestore flag these emails can be reload into your account using this method.

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


## WaitForNthMissedEmail

> MissedEmailDto WaitForNthMissedEmail(ctx, index, optional)

Wait for Nth missed email

Wait for 0 based index missed email

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **int32**| Zero based index of the email to wait for. If 1 missed email already and you want to wait for the 2nd email pass index&#x3D;1 | 
 **optional** | ***WaitForNthMissedEmailOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WaitForNthMissedEmailOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inboxId** | [**optional.Interface of string**]()| Optional inbox ID filter | 
 **timeout** | **optional.Int64**| Optional timeout milliseconds | 
 **since** | **optional.Time**| Filter by created at after the given timestamp | 
 **before** | **optional.Time**| Filter by created at before the given timestamp | 

### Return type

[**MissedEmailDto**](MissedEmailDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

