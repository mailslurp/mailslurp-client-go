# MailSlurp\AliasControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAlias**](AliasControllerApi#CreateAlias) | **Post** /aliases | Create an email alias. Must be verified by clicking link inside verification email that will be sent to the address. Once verified the alias will be active.
[**DeleteAlias**](AliasControllerApi#DeleteAlias) | **Delete** /aliases/{aliasId} | Delete an email alias
[**GetAlias**](AliasControllerApi#GetAlias) | **Get** /aliases/{aliasId} | Get an email alias
[**GetAliasEmails**](AliasControllerApi#GetAliasEmails) | **Get** /aliases/{aliasId}/emails | Get emails for an alias
[**GetAliasThreads**](AliasControllerApi#GetAliasThreads) | **Get** /aliases/{aliasId}/threads | Get threads created for an alias
[**GetAliases**](AliasControllerApi#GetAliases) | **Get** /aliases | Get all email aliases you have created
[**GetThread**](AliasControllerApi#GetThread) | **Get** /aliases/threads/{threadId} | Get a thread
[**GetThreadsPaginated**](AliasControllerApi#GetThreadsPaginated) | **Get** /aliases/threads | Get all threads
[**ReplyToAliasEmail**](AliasControllerApi#ReplyToAliasEmail) | **Put** /aliases/{aliasId}/emails/{emailId} | Reply to an email
[**SendAliasEmail**](AliasControllerApi#SendAliasEmail) | **Post** /aliases/{aliasId}/emails | Send an email from an alias inbox
[**UpdateAlias**](AliasControllerApi#UpdateAlias) | **Put** /aliases/{aliasId} | Update an email alias



## CreateAlias

> AliasDto CreateAlias(ctx, createAliasOptions)

Create an email alias. Must be verified by clicking link inside verification email that will be sent to the address. Once verified the alias will be active.

Email aliases use a MailSlurp randomly generated email address (or a custom domain inbox that you provide) to mask or proxy a real email address. Emails sent to the alias address will be forwarded to the hidden email address it was created for. If you want to send a reply use the threadId attached

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createAliasOptions** | [**CreateAliasOptions**](CreateAliasOptions)|  | 

### Return type

[**AliasDto**](AliasDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DeleteAlias

> DeleteAlias(ctx, aliasId)

Delete an email alias

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aliasId** | [**string**]()|  | 

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


## GetAlias

> AliasDto GetAlias(ctx, aliasId)

Get an email alias

Get an email alias by ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aliasId** | [**string**]()|  | 

### Return type

[**AliasDto**](AliasDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetAliasEmails

> PageEmailProjection GetAliasEmails(ctx, aliasId, optional)

Get emails for an alias

Get paginated emails for an alias by ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aliasId** | [**string**]()|  | 
 **optional** | ***GetAliasEmailsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAliasEmailsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Optional page index alias email list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size alias email list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **since** | **optional.Time**| Optional filter by sent after given date time | 
 **before** | **optional.Time**| Optional filter by sent before given date time | 

### Return type

[**PageEmailProjection**](PageEmailProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetAliasThreads

> PageAliasThreadProjection GetAliasThreads(ctx, aliasId, optional)

Get threads created for an alias

Returns threads created for an email alias in paginated form

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aliasId** | [**string**]()|  | 
 **optional** | ***GetAliasThreadsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAliasThreadsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Optional page index in thread list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in thread list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **since** | **optional.Time**| Optional filter by sent after given date time | 
 **before** | **optional.Time**| Optional filter by sent before given date time | 

### Return type

[**PageAliasThreadProjection**](PageAliasThreadProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetAliases

> PageAlias GetAliases(ctx, optional)

Get all email aliases you have created

Get all email aliases in paginated form

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAliasesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAliasesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| Optional search term | 
 **page** | **optional.Int32**| Optional page index in alias list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in alias list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **since** | **optional.Time**| Filter by created at after the given timestamp | 
 **before** | **optional.Time**| Filter by created at before the given timestamp | 

### Return type

[**PageAlias**](PageAlias)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetThread

> AliasThreadProjection GetThread(ctx, threadId)

Get a thread

Return a thread associated with an alias

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | [**string**]()|  | 

### Return type

[**AliasThreadProjection**](AliasThreadProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetThreadsPaginated

> PageAliasThreadProjection GetThreadsPaginated(ctx, optional)

Get all threads

Returns threads created for all aliases in paginated form

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetThreadsPaginatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetThreadsPaginatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Optional page index in thread list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in thread list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **since** | **optional.Time**| Optional filter by sent after given date time | 
 **before** | **optional.Time**| Optional filter by sent before given date time | 

### Return type

[**PageAliasThreadProjection**](PageAliasThreadProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## ReplyToAliasEmail

> SentEmailDto ReplyToAliasEmail(ctx, aliasId, emailId, replyToAliasEmailOptions)

Reply to an email

Send the reply to the email sender or reply-to and include same subject cc bcc etc. Reply to an email and the contents will be sent with the existing subject to the emails `to`, `cc`, and `bcc`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aliasId** | [**string**]()| ID of the alias that email belongs to | 
**emailId** | [**string**]()| ID of the email that should be replied to | 
**replyToAliasEmailOptions** | [**ReplyToAliasEmailOptions**](ReplyToAliasEmailOptions)|  | 

### Return type

[**SentEmailDto**](SentEmailDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## SendAliasEmail

> SentEmailDto SendAliasEmail(ctx, aliasId, sendEmailOptions)

Send an email from an alias inbox

Send an email from an alias. Replies to the email will be forwarded to the alias masked email address

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aliasId** | [**string**]()|  | 
**sendEmailOptions** | [**SendEmailOptions**](SendEmailOptions)|  | 

### Return type

[**SentEmailDto**](SentEmailDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## UpdateAlias

> AliasDto UpdateAlias(ctx, aliasId, updateAliasOptions)

Update an email alias

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aliasId** | [**string**]()|  | 
**updateAliasOptions** | [**UpdateAliasOptions**](UpdateAliasOptions)|  | 

### Return type

[**AliasDto**](AliasDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

