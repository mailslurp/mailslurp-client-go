# \CommonActionsControllerApi

All URIs are relative to *https://api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewEmailAddress**](CommonActionsControllerApi.md#CreateNewEmailAddress) | **Post** /newEmailAddress | Create new random inbox
[**EmptyInbox**](CommonActionsControllerApi.md#EmptyInbox) | **Delete** /emptyInbox | Delete all emails in an inbox
[**SendEmailSimple**](CommonActionsControllerApi.md#SendEmailSimple) | **Post** /sendEmail | Send an email from a random email address
[**WaitForEmailCount**](CommonActionsControllerApi.md#WaitForEmailCount) | **Get** /waitForEmailCount | Wait for and return count number of emails 
[**WaitForLatestEmail**](CommonActionsControllerApi.md#WaitForLatestEmail) | **Get** /waitForLatestEmail | Fetch inbox&#39;s latest email or if empty wait for email to arrive
[**WaitForMatchingEmail**](CommonActionsControllerApi.md#WaitForMatchingEmail) | **Post** /waitForMatchingEmails | Wait or return list of emails that match simple matching patterns
[**WaitForNthEmail**](CommonActionsControllerApi.md#WaitForNthEmail) | **Get** /waitForNthEmail | Wait for or fetch the email with a given index in the inbox specified



## CreateNewEmailAddress

> Inbox CreateNewEmailAddress(ctx, )

Create new random inbox

Returns an Inbox with an `id` and an `emailAddress`

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**Inbox**](Inbox.md)

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmptyInbox

> EmptyInbox(ctx, inboxId)

Delete all emails in an inbox

Deletes all emails

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**](.md)| inboxId | 

### Return type

 (empty response body)

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendEmailSimple

> SendEmailSimple(ctx, sendEmailOptions)

Send an email from a random email address

To specify an email address first create an inbox and use that with the other send email methods

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sendEmailOptions** | [**SendEmailOptions**](SendEmailOptions.md)| sendEmailOptions | 

### Return type

 (empty response body)

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WaitForEmailCount

> []EmailPreview WaitForEmailCount(ctx, optional)

Wait for and return count number of emails 

Will only wait if count is greater that the found emails in given inbox.If you need to wait for an email for a non-empty inbox see the other receive methods.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WaitForEmailCountOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WaitForEmailCountOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| Number of emails to wait for. Must be greater that 1 | 
 **inboxId** | [**optional.Interface of string**](.md)| Id of the inbox we are fetching emails from | 
 **timeout** | **optional.Int64**| Max milliseconds to wait | 
 **unreadOnly** | **optional.Bool**| Optional filter for unread only | [default to false]

### Return type

[**[]EmailPreview**](EmailPreview.md)

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WaitForLatestEmail

> Email WaitForLatestEmail(ctx, optional)

Fetch inbox's latest email or if empty wait for email to arrive

Will return either the last received email or wait for an email to arrive and return that. If you need to wait for an email for a non-empty inbox see the other receive methods.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WaitForLatestEmailOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WaitForLatestEmailOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inboxId** | [**optional.Interface of string**](.md)| Id of the inbox we are fetching emails from | 
 **timeout** | **optional.Int64**| Max milliseconds to wait | 
 **unreadOnly** | **optional.Bool**| Optional filter for unread only | [default to false]

### Return type

[**Email**](Email.md)

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WaitForMatchingEmail

> []EmailPreview WaitForMatchingEmail(ctx, matchOptions, optional)

Wait or return list of emails that match simple matching patterns

Results must also meet provided count. Match options allow simple CONTAINS or EQUALS filtering on SUBJECT, TO, BCC, CC, and FROM.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**matchOptions** | [**MatchOptions**](MatchOptions.md)| matchOptions | 
 **optional** | ***WaitForMatchingEmailOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WaitForMatchingEmailOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| Number of emails to wait for. Must be greater that 1 | 
 **inboxId** | [**optional.Interface of string**](.md)| Id of the inbox we are fetching emails from | 
 **timeout** | **optional.Int64**| Max milliseconds to wait | 
 **unreadOnly** | **optional.Bool**| Optional filter for unread only | [default to false]

### Return type

[**[]EmailPreview**](EmailPreview.md)

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WaitForNthEmail

> Email WaitForNthEmail(ctx, optional)

Wait for or fetch the email with a given index in the inbox specified

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WaitForNthEmailOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WaitForNthEmailOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inboxId** | [**optional.Interface of string**](.md)| Id of the inbox you are fetching emails from | 
 **index** | **optional.Int32**| Zero based index of the email to wait for | 
 **timeout** | **optional.Int64**| Max milliseconds to wait | 
 **unreadOnly** | **optional.Bool**| Optional filter for unread only | [default to false]

### Return type

[**Email**](Email.md)

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

