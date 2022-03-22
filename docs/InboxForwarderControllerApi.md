# MailSlurp\InboxForwarderControllerApi

All URIs are relative to *https://api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewInboxForwarder**](InboxForwarderControllerApi#CreateNewInboxForwarder) | **Post** /forwarders | Create an inbox forwarder
[**DeleteInboxForwarder**](InboxForwarderControllerApi#DeleteInboxForwarder) | **Delete** /forwarders/{id} | Delete an inbox forwarder
[**DeleteInboxForwarders**](InboxForwarderControllerApi#DeleteInboxForwarders) | **Delete** /forwarders | Delete inbox forwarders
[**GetInboxForwarder**](InboxForwarderControllerApi#GetInboxForwarder) | **Get** /forwarders/{id} | Get an inbox forwarder
[**GetInboxForwarders**](InboxForwarderControllerApi#GetInboxForwarders) | **Get** /forwarders | List inbox forwarders
[**TestInboxForwarder**](InboxForwarderControllerApi#TestInboxForwarder) | **Post** /forwarders/{id}/test | Test an inbox forwarder
[**TestInboxForwardersForInbox**](InboxForwarderControllerApi#TestInboxForwardersForInbox) | **Put** /forwarders | Test inbox forwarders for inbox
[**TestNewInboxForwarder**](InboxForwarderControllerApi#TestNewInboxForwarder) | **Patch** /forwarders | Test new inbox forwarder



## CreateNewInboxForwarder

> InboxForwarderDto CreateNewInboxForwarder(ctx, inboxId, createInboxForwarderOptions)

Create an inbox forwarder

Create a new inbox rule for forwarding, blocking, and allowing emails when sending and receiving

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()| Inbox id to attach forwarder to | 
**createInboxForwarderOptions** | [**CreateInboxForwarderOptions**](CreateInboxForwarderOptions)|  | 

### Return type

[**InboxForwarderDto**](InboxForwarderDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DeleteInboxForwarder

> DeleteInboxForwarder(ctx, id)

Delete an inbox forwarder

Delete inbox forwarder

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()| ID of inbox forwarder | 

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


## DeleteInboxForwarders

> DeleteInboxForwarders(ctx, optional)

Delete inbox forwarders

Delete inbox forwarders. Accepts optional inboxId filter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteInboxForwardersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteInboxForwardersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inboxId** | [**optional.Interface of string**]()| Optional inbox id to attach forwarder to | 

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


## GetInboxForwarder

> InboxForwarderDto GetInboxForwarder(ctx, id)

Get an inbox forwarder

Get inbox ruleset

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()| ID of inbox forwarder | 

### Return type

[**InboxForwarderDto**](InboxForwarderDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetInboxForwarders

> PageInboxForwarderDto GetInboxForwarders(ctx, optional)

List inbox forwarders

List all forwarders attached to an inbox

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetInboxForwardersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetInboxForwardersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inboxId** | [**optional.Interface of string**]()| Optional inbox id to get forwarders from | 
 **page** | **optional.Int32**| Optional page index in inbox forwarder list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in inbox forwarder list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **searchFilter** | **optional.String**| Optional search filter | 
 **since** | **optional.Time**| Filter by created at after the given timestamp | 
 **before** | **optional.Time**| Filter by created at before the given timestamp | 

### Return type

[**PageInboxForwarderDto**](PageInboxForwarderDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## TestInboxForwarder

> InboxForwarderTestResult TestInboxForwarder(ctx, id, inboxForwarderTestOptions)

Test an inbox forwarder

Test an inbox forwarder

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()| ID of inbox forwarder | 
**inboxForwarderTestOptions** | [**InboxForwarderTestOptions**](InboxForwarderTestOptions)|  | 

### Return type

[**InboxForwarderTestResult**](InboxForwarderTestResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## TestInboxForwardersForInbox

> InboxForwarderTestResult TestInboxForwardersForInbox(ctx, inboxId, inboxForwarderTestOptions)

Test inbox forwarders for inbox

Test inbox forwarders for inbox

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()| ID of inbox | 
**inboxForwarderTestOptions** | [**InboxForwarderTestOptions**](InboxForwarderTestOptions)|  | 

### Return type

[**InboxForwarderTestResult**](InboxForwarderTestResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## TestNewInboxForwarder

> InboxForwarderTestResult TestNewInboxForwarder(ctx, testNewInboxForwarderOptions)

Test new inbox forwarder

Test new inbox forwarder

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testNewInboxForwarderOptions** | [**TestNewInboxForwarderOptions**](TestNewInboxForwarderOptions)|  | 

### Return type

[**InboxForwarderTestResult**](InboxForwarderTestResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

