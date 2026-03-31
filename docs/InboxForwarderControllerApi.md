# MailSlurp\InboxForwarderControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewInboxForwarder**](InboxForwarderControllerApi#CreateNewInboxForwarder) | **Post** /forwarders | Create an inbox forwarder
[**DeleteInboxForwarder**](InboxForwarderControllerApi#DeleteInboxForwarder) | **Delete** /forwarders/{id} | Delete an inbox forwarder
[**DeleteInboxForwarders**](InboxForwarderControllerApi#DeleteInboxForwarders) | **Delete** /forwarders | Delete inbox forwarders
[**GetAllInboxForwarderEvents**](InboxForwarderControllerApi#GetAllInboxForwarderEvents) | **Get** /forwarders/events | Get all inbox forwarder events
[**GetForwarderEvent**](InboxForwarderControllerApi#GetForwarderEvent) | **Get** /forwarders/events/{eventId} | Get a forwarder event
[**GetInboxForwarder**](InboxForwarderControllerApi#GetInboxForwarder) | **Get** /forwarders/{id} | Get an inbox forwarder
[**GetInboxForwarderEvent**](InboxForwarderControllerApi#GetInboxForwarderEvent) | **Get** /forwarders/{id}/events/{eventId} | Get an inbox forwarder event
[**GetInboxForwarderEvents**](InboxForwarderControllerApi#GetInboxForwarderEvents) | **Get** /forwarders/{id}/events | Get an inbox forwarder event list
[**GetInboxForwarders**](InboxForwarderControllerApi#GetInboxForwarders) | **Get** /forwarders | List inbox forwarders
[**TestInboxForwarder**](InboxForwarderControllerApi#TestInboxForwarder) | **Post** /forwarders/{id}/test | Test an inbox forwarder
[**TestInboxForwardersForInbox**](InboxForwarderControllerApi#TestInboxForwardersForInbox) | **Put** /forwarders | Test inbox forwarders for inbox
[**TestNewInboxForwarder**](InboxForwarderControllerApi#TestNewInboxForwarder) | **Patch** /forwarders | Test new inbox forwarder
[**UpdateInboxForwarder**](InboxForwarderControllerApi#UpdateInboxForwarder) | **Put** /forwarders/{id} | Update an inbox forwarder



## CreateNewInboxForwarder

> InboxForwarderDto CreateNewInboxForwarder(ctx, createInboxForwarderOptions, optional)

Create an inbox forwarder

Create a new inbox rule for forwarding, blocking, and allowing emails when sending and receiving

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createInboxForwarderOptions** | [**CreateInboxForwarderOptions**](CreateInboxForwarderOptions)|  | 
 **optional** | ***CreateNewInboxForwarderOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateNewInboxForwarderOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inboxId** | [**optional.Interface of string**]()| Inbox id to attach forwarder to | 

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


## GetAllInboxForwarderEvents

> PageInboxForwarderEvents GetAllInboxForwarderEvents(ctx, optional)

Get all inbox forwarder events

Get all inbox forwarder events

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllInboxForwarderEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllInboxForwarderEventsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Optional page index in inbox forwarder event list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in inbox forwarder event list pagination | [default to 20]
 **inboxId** | [**optional.Interface of string**]()| Optional inbox ID to filter for | 
 **emailId** | [**optional.Interface of string**]()| Optional email ID to filter for | 
 **sentId** | [**optional.Interface of string**]()| Optional sent ID to filter for | 
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]

### Return type

[**PageInboxForwarderEvents**](PageInboxForwarderEvents)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetForwarderEvent

> InboxForwarderEventDto GetForwarderEvent(ctx, eventId)

Get a forwarder event

Get forwarder event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | [**string**]()| ID of inbox forwarder event | 

### Return type

[**InboxForwarderEventDto**](InboxForwarderEventDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetInboxForwarder

> InboxForwarderDto GetInboxForwarder(ctx, id)

Get an inbox forwarder

Get inbox forwarder

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


## GetInboxForwarderEvent

> InboxForwarderEventDto GetInboxForwarderEvent(ctx, id, eventId)

Get an inbox forwarder event

Get inbox forwarder event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()| ID of inbox forwarder | 
**eventId** | [**string**]()| ID of inbox forwarder event | 

### Return type

[**InboxForwarderEventDto**](InboxForwarderEventDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetInboxForwarderEvents

> PageInboxForwarderEvents GetInboxForwarderEvents(ctx, id, optional)

Get an inbox forwarder event list

Get inbox forwarder events

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()| ID of inbox forwarder | 
 **optional** | ***GetInboxForwarderEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetInboxForwarderEventsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Optional page index in inbox forwarder event list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in inbox forwarder event list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]

### Return type

[**PageInboxForwarderEvents**](PageInboxForwarderEvents)

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


## UpdateInboxForwarder

> InboxForwarderDto UpdateInboxForwarder(ctx, id, createInboxForwarderOptions)

Update an inbox forwarder

Update inbox forwarder

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()| ID of inbox forwarder | 
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

