# MailSlurp\ImapControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImapServerFetch**](ImapControllerApi#ImapServerFetch) | **Post** /imap/server/fetch | Fetch message in an inbox
[**ImapServerGet**](ImapControllerApi#ImapServerGet) | **Post** /imap/server/get | Get a message by email ID
[**ImapServerList**](ImapControllerApi#ImapServerList) | **Post** /imap/server/list | List messages in an inbox
[**ImapServerSearch**](ImapControllerApi#ImapServerSearch) | **Post** /imap/server/search | Search messages in an inbox
[**ImapServerStatus**](ImapControllerApi#ImapServerStatus) | **Post** /imap/server/status | Get status for mailbox
[**ImapServerUpdateFlags**](ImapControllerApi#ImapServerUpdateFlags) | **Post** /imap/server/update-flags | 



## ImapServerFetch

> ImapServerFetchResult ImapServerFetch(ctx, seqNum, optional)

Fetch message in an inbox

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seqNum** | **int64**|  | 
 **optional** | ***ImapServerFetchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ImapServerFetchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inboxId** | [**optional.Interface of string**]()| Inbox ID to search | 

### Return type

[**ImapServerFetchResult**](ImapServerFetchResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## ImapServerGet

> ImapServerGetResult ImapServerGet(ctx, emailId, optional)

Get a message by email ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()| Email ID to get | 
 **optional** | ***ImapServerGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ImapServerGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inboxId** | [**optional.Interface of string**]()| Inbox ID to search | 

### Return type

[**ImapServerGetResult**](ImapServerGetResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## ImapServerList

> ImapServerListResult ImapServerList(ctx, imapServerListOptions, optional)

List messages in an inbox

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imapServerListOptions** | [**ImapServerListOptions**](ImapServerListOptions)|  | 
 **optional** | ***ImapServerListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ImapServerListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inboxId** | [**optional.Interface of string**]()| Inbox ID to list | 

### Return type

[**ImapServerListResult**](ImapServerListResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## ImapServerSearch

> ImapServerSearchResult ImapServerSearch(ctx, imapServerSearchOptions, optional)

Search messages in an inbox

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imapServerSearchOptions** | [**ImapServerSearchOptions**](ImapServerSearchOptions)|  | 
 **optional** | ***ImapServerSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ImapServerSearchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inboxId** | [**optional.Interface of string**]()| Inbox ID to search | 

### Return type

[**ImapServerSearchResult**](ImapServerSearchResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## ImapServerStatus

> ImapServerStatusResult ImapServerStatus(ctx, imapServerStatusOptions, optional)

Get status for mailbox

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imapServerStatusOptions** | [**ImapServerStatusOptions**](ImapServerStatusOptions)|  | 
 **optional** | ***ImapServerStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ImapServerStatusOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inboxId** | [**optional.Interface of string**]()| Inbox ID to list | 

### Return type

[**ImapServerStatusResult**](ImapServerStatusResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## ImapServerUpdateFlags

> ImapServerUpdateFlags(ctx, imapUpdateFlagsOptions, optional)



Update message flags

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imapUpdateFlagsOptions** | [**ImapUpdateFlagsOptions**](ImapUpdateFlagsOptions)|  | 
 **optional** | ***ImapServerUpdateFlagsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ImapServerUpdateFlagsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inboxId** | [**optional.Interface of string**]()|  | 

### Return type

 (empty response body)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

