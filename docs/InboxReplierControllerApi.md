# MailSlurp\InboxReplierControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewInboxReplier**](InboxReplierControllerApi#CreateNewInboxReplier) | **Post** /repliers | Create an inbox replier
[**DeleteInboxReplier**](InboxReplierControllerApi#DeleteInboxReplier) | **Delete** /repliers/{id} | Delete an inbox replier
[**DeleteInboxRepliers**](InboxReplierControllerApi#DeleteInboxRepliers) | **Delete** /repliers | Delete inbox repliers
[**GetAllInboxReplierEvents**](InboxReplierControllerApi#GetAllInboxReplierEvents) | **Get** /repliers/events | Get inbox replier event list
[**GetInboxReplier**](InboxReplierControllerApi#GetInboxReplier) | **Get** /repliers/{id} | Get an inbox replier
[**GetInboxReplierEvents**](InboxReplierControllerApi#GetInboxReplierEvents) | **Get** /repliers/{id}/events | Get an inbox replier event list
[**GetInboxRepliers**](InboxReplierControllerApi#GetInboxRepliers) | **Get** /repliers | List inbox repliers
[**UpdateInboxReplier**](InboxReplierControllerApi#UpdateInboxReplier) | **Put** /repliers/{id} | Update an inbox replier



## CreateNewInboxReplier

> InboxReplierDto CreateNewInboxReplier(ctx, createInboxReplierOptions)

Create an inbox replier

Create a new inbox rule for reply toing, blocking, and allowing emails when sending and receiving

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createInboxReplierOptions** | [**CreateInboxReplierOptions**](CreateInboxReplierOptions)|  | 

### Return type

[**InboxReplierDto**](InboxReplierDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DeleteInboxReplier

> DeleteInboxReplier(ctx, id)

Delete an inbox replier

Delete inbox replier

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()| ID of inbox replier | 

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


## DeleteInboxRepliers

> DeleteInboxRepliers(ctx, optional)

Delete inbox repliers

Delete inbox repliers. Accepts optional inboxId filter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteInboxRepliersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteInboxRepliersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inboxId** | [**optional.Interface of string**]()| Optional inbox id to attach replier to | 

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


## GetAllInboxReplierEvents

> PageInboxReplierEvents GetAllInboxReplierEvents(ctx, optional)

Get inbox replier event list

Get all inbox ruleset events

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllInboxReplierEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllInboxReplierEventsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inboxReplierId** | [**optional.Interface of string**]()| ID of inbox replier | 
 **inboxId** | [**optional.Interface of string**]()| ID of inbox | 
 **emailId** | [**optional.Interface of string**]()| ID of email | 
 **sentId** | [**optional.Interface of string**]()| ID of sent | 
 **page** | **optional.Int32**| Optional page index in inbox replier event list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in inbox replier event list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]

### Return type

[**PageInboxReplierEvents**](PageInboxReplierEvents)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetInboxReplier

> InboxReplierDto GetInboxReplier(ctx, id)

Get an inbox replier

Get inbox ruleset

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()| ID of inbox replier | 

### Return type

[**InboxReplierDto**](InboxReplierDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetInboxReplierEvents

> PageInboxReplierEvents GetInboxReplierEvents(ctx, id, optional)

Get an inbox replier event list

Get inbox ruleset events

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()| ID of inbox replier | 
 **optional** | ***GetInboxReplierEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetInboxReplierEventsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Optional page index in inbox replier event list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in inbox replier event list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]

### Return type

[**PageInboxReplierEvents**](PageInboxReplierEvents)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetInboxRepliers

> PageInboxReplierDto GetInboxRepliers(ctx, optional)

List inbox repliers

List all repliers attached to an inbox

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetInboxRepliersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetInboxRepliersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inboxId** | [**optional.Interface of string**]()| Optional inbox id to get repliers from | 
 **page** | **optional.Int32**| Optional page index in inbox replier list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in inbox replier list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **since** | **optional.Time**| Filter by created at after the given timestamp | 
 **before** | **optional.Time**| Filter by created at before the given timestamp | 

### Return type

[**PageInboxReplierDto**](PageInboxReplierDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## UpdateInboxReplier

> InboxReplierDto UpdateInboxReplier(ctx, id, updateInboxReplierOptions)

Update an inbox replier

Update inbox ruleset

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()| ID of inbox replier | 
**updateInboxReplierOptions** | [**UpdateInboxReplierOptions**](UpdateInboxReplierOptions)|  | 

### Return type

[**InboxReplierDto**](InboxReplierDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

