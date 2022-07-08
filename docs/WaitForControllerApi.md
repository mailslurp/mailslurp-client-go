# MailSlurp\WaitForControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WaitFor**](WaitForControllerApi#WaitFor) | **Post** /waitFor | Wait for an email to match the provided filter conditions such as subject contains keyword.
[**WaitForEmailCount**](WaitForControllerApi#WaitForEmailCount) | **Get** /waitForEmailCount | Wait for and return count number of emails. Hold connection until inbox count matches expected or timeout occurs
[**WaitForLatestEmail**](WaitForControllerApi#WaitForLatestEmail) | **Get** /waitForLatestEmail | Fetch inbox&#39;s latest email or if empty wait for an email to arrive
[**WaitForMatchingEmails**](WaitForControllerApi#WaitForMatchingEmails) | **Post** /waitForMatchingEmails | Wait or return list of emails that match simple matching patterns
[**WaitForMatchingFirstEmail**](WaitForControllerApi#WaitForMatchingFirstEmail) | **Post** /waitForMatchingFirstEmail | Wait for or return the first email that matches provided MatchOptions array
[**WaitForNthEmail**](WaitForControllerApi#WaitForNthEmail) | **Get** /waitForNthEmail | Wait for or fetch the email with a given index in the inbox specified. If index doesn&#39;t exist waits for it to exist or timeout to occur.
[**WaitForSms**](WaitForControllerApi#WaitForSms) | **Post** /waitForSms | Wait for an SMS message to match the provided filter conditions such as body contains keyword.



## WaitFor

> []EmailPreview WaitFor(ctx, waitForConditions)

Wait for an email to match the provided filter conditions such as subject contains keyword.

Generic waitFor method that will wait until an inbox meets given conditions or return immediately if already met

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**waitForConditions** | [**WaitForConditions**](WaitForConditions)|  | 

### Return type

[**[]EmailPreview**](EmailPreview)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## WaitForEmailCount

> []EmailPreview WaitForEmailCount(ctx, inboxId, count, optional)

Wait for and return count number of emails. Hold connection until inbox count matches expected or timeout occurs

If inbox contains count or more emails at time of request then return count worth of emails. If not wait until the count is reached and return those or return an error if timeout is exceeded.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()| Id of the inbox we are fetching emails from | 
**count** | **int32**| Number of emails to wait for. Must be greater that 1 | 
 **optional** | ***WaitForEmailCountOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WaitForEmailCountOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **timeout** | **optional.Int64**| Max milliseconds to wait | 
 **unreadOnly** | **optional.Bool**| Optional filter for unread only | [default to false]
 **before** | **optional.Time**| Filter for emails that were received before the given timestamp | 
 **since** | **optional.Time**| Filter for emails that were received after the given timestamp | 
 **sort** | **optional.String**| Sort direction | 
 **delay** | **optional.Int64**| Max milliseconds delay between calls | 

### Return type

[**[]EmailPreview**](EmailPreview)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


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
 **inboxId** | [**optional.Interface of string**]()| Id of the inbox we are fetching emails from | 
 **timeout** | **optional.Int64**| Max milliseconds to wait | 
 **unreadOnly** | **optional.Bool**| Optional filter for unread only. | [default to false]
 **before** | **optional.Time**| Filter for emails that were before after the given timestamp | 
 **since** | **optional.Time**| Filter for emails that were received after the given timestamp | 
 **sort** | **optional.String**| Sort direction | 
 **delay** | **optional.Int64**| Max milliseconds delay between calls | 

### Return type

[**Email**](Email)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## WaitForMatchingEmails

> []EmailPreview WaitForMatchingEmails(ctx, inboxId, count, matchOptions, optional)

Wait or return list of emails that match simple matching patterns

Perform a search of emails in an inbox with the given patterns. If results match expected count then return or else retry the search until results are found or timeout is reached. Match options allow simple CONTAINS or EQUALS filtering on SUBJECT, TO, BCC, CC, and FROM. See the `MatchOptions` object for options. An example payload is `{ matches: [{field: 'SUBJECT',should:'CONTAIN',value:'needle'}] }`. You can use an array of matches and they will be applied sequentially to filter out emails. If you want to perform matches and extractions of content using Regex patterns see the EmailController `getEmailContentMatch` method.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()| Id of the inbox we are fetching emails from | 
**count** | **int32**| Number of emails to wait for. Must be greater or equal to 1 | 
**matchOptions** | [**MatchOptions**](MatchOptions)|  | 
 **optional** | ***WaitForMatchingEmailsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WaitForMatchingEmailsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **before** | **optional.Time**| Filter for emails that were received before the given timestamp | 
 **since** | **optional.Time**| Filter for emails that were received after the given timestamp | 
 **sort** | **optional.String**| Sort direction | 
 **delay** | **optional.Int64**| Max milliseconds delay between calls | 
 **timeout** | **optional.Int64**| Max milliseconds to wait | 
 **unreadOnly** | **optional.Bool**| Optional filter for unread only | [default to false]

### Return type

[**[]EmailPreview**](EmailPreview)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## WaitForMatchingFirstEmail

> Email WaitForMatchingFirstEmail(ctx, inboxId, matchOptions, optional)

Wait for or return the first email that matches provided MatchOptions array

Perform a search of emails in an inbox with the given patterns. If a result if found then return or else retry the search until a result is found or timeout is reached. Match options allow simple CONTAINS or EQUALS filtering on SUBJECT, TO, BCC, CC, and FROM. See the `MatchOptions` object for options. An example payload is `{ matches: [{field: 'SUBJECT',should:'CONTAIN',value:'needle'}] }`. You can use an array of matches and they will be applied sequentially to filter out emails. If you want to perform matches and extractions of content using Regex patterns see the EmailController `getEmailContentMatch` method.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()| Id of the inbox we are matching an email for | 
**matchOptions** | [**MatchOptions**](MatchOptions)|  | 
 **optional** | ***WaitForMatchingFirstEmailOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WaitForMatchingFirstEmailOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **timeout** | **optional.Int64**| Max milliseconds to wait | 
 **unreadOnly** | **optional.Bool**| Optional filter for unread only | [default to false]
 **since** | **optional.Time**| Filter for emails that were received after the given timestamp | 
 **before** | **optional.Time**| Filter for emails that were received before the given timestamp | 
 **sort** | **optional.String**| Sort direction | 
 **delay** | **optional.Int64**| Max milliseconds delay between calls | 

### Return type

[**Email**](Email)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## WaitForNthEmail

> Email WaitForNthEmail(ctx, optional)

Wait for or fetch the email with a given index in the inbox specified. If index doesn't exist waits for it to exist or timeout to occur.

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
 **inboxId** | [**optional.Interface of string**]()| Id of the inbox you are fetching emails from | 
 **index** | **optional.Int32**| Zero based index of the email to wait for. If an inbox has 1 email already and you want to wait for the 2nd email pass index&#x3D;1 | [default to 0]
 **timeout** | **optional.Int64**| Max milliseconds to wait for the nth email if not already present | 
 **unreadOnly** | **optional.Bool**| Optional filter for unread only | [default to false]
 **since** | **optional.Time**| Filter for emails that were received after the given timestamp | 
 **before** | **optional.Time**| Filter for emails that were received before the given timestamp | 
 **sort** | **optional.String**| Sort direction | 
 **delay** | **optional.Int64**| Max milliseconds delay between calls | 

### Return type

[**Email**](Email)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## WaitForSms

> []SmsPreview WaitForSms(ctx, waitForSmsConditions)

Wait for an SMS message to match the provided filter conditions such as body contains keyword.

Generic waitFor method that will wait until a phone number meets given conditions or return immediately if already met

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**waitForSmsConditions** | [**WaitForSmsConditions**](WaitForSmsConditions)|  | 

### Return type

[**[]SmsPreview**](SmsPreview)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

