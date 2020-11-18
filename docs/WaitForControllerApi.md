# MailSlurp\WaitForControllerApi

All URIs are relative to *https://api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WaitFor**](WaitForControllerApi.md#WaitFor) | **Post** /waitFor | Wait for conditions to be met
[**WaitForEmailCount**](WaitForControllerApi.md#WaitForEmailCount) | **Get** /waitForEmailCount | Wait for and return count number of emails 
[**WaitForLatestEmail**](WaitForControllerApi.md#WaitForLatestEmail) | **Get** /waitForLatestEmail | Fetch inbox&#39;s latest email or if empty wait for an email to arrive
[**WaitForMatchingEmail**](WaitForControllerApi.md#WaitForMatchingEmail) | **Post** /waitForMatchingEmails | Wait or return list of emails that match simple matching patterns
[**WaitForNthEmail**](WaitForControllerApi.md#WaitForNthEmail) | **Get** /waitForNthEmail | Wait for or fetch the email with a given index in the inbox specified



## WaitFor

> []EmailPreview WaitFor(ctx, optional)

Wait for conditions to be met

Generic waitFor method that will wait until an inbox meets given conditions or return immediately if already met

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WaitForOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WaitForOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **waitForConditions** | [**optional.Interface of WaitForConditions**](WaitForConditions.md)| Conditions to wait for | 

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


## WaitForEmailCount

> []EmailPreview WaitForEmailCount(ctx, optional)

Wait for and return count number of emails 

If inbox contains count or more emails at time of request then return count worth of emails. If not wait until the count is reached and return those or return an error if timeout is exceeded.

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

Fetch inbox's latest email or if empty wait for an email to arrive

Will return either the last received email or wait for an email to arrive and return that. If you need to wait for an email for a non-empty inbox set `unreadOnly=true` or see the other receive methods such as `waitForNthEmail` or `waitForEmailCount`.

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
 **unreadOnly** | **optional.Bool**| Optional filter for unread only. | [default to false]

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

Perform a search of emails in an inbox with the given patterns. If results match expected count then return or else retry the search until results are found or timeout is reached. Match options allow simple CONTAINS or EQUALS filtering on SUBJECT, TO, BCC, CC, and FROM. See the `MatchOptions` object for options.

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

If nth email is already present in inbox then return it. If not hold the connection open until timeout expires or the nth email is received and returned.

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
 **index** | **optional.Int32**| Zero based index of the email to wait for. If an inbox has 1 email already and you want to wait for the 2nd email pass index&#x3D;1 | [default to 0]
 **timeout** | **optional.Int64**| Max milliseconds to wait for the nth email if not already present | 
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

