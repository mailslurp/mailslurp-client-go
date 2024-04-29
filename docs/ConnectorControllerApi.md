# MailSlurp\ConnectorControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConnector**](ConnectorControllerApi#CreateConnector) | **Post** /connectors | Create an inbox connector
[**CreateConnectorImapConnection**](ConnectorControllerApi#CreateConnectorImapConnection) | **Post** /connectors/{id}/imap | Create an inbox connector IMAP connection
[**CreateConnectorSmtpConnection**](ConnectorControllerApi#CreateConnectorSmtpConnection) | **Post** /connectors/{id}/smtp | Create an inbox connector SMTP connection
[**DeleteAllConnector**](ConnectorControllerApi#DeleteAllConnector) | **Delete** /connectors | Delete all inbox connectors
[**DeleteConnector**](ConnectorControllerApi#DeleteConnector) | **Delete** /connectors/{id} | Delete an inbox connector
[**DeleteConnectorImapConnection**](ConnectorControllerApi#DeleteConnectorImapConnection) | **Delete** /connectors/{id}/imap | Delete an inbox connector IMAP connection
[**DeleteConnectorSmtpConnection**](ConnectorControllerApi#DeleteConnectorSmtpConnection) | **Delete** /connectors/{id}/smtp | Delete an inbox connector SMTP connection
[**GetAllConnectorSyncEvents**](ConnectorControllerApi#GetAllConnectorSyncEvents) | **Get** /connectors/events | Get all inbox connector sync events
[**GetConnector**](ConnectorControllerApi#GetConnector) | **Get** /connectors/{id} | Get an inbox connector
[**GetConnectorSyncEvent**](ConnectorControllerApi#GetConnectorSyncEvent) | **Get** /connectors/events/{id} | Get an inbox connector sync event
[**GetConnectorSyncEvents**](ConnectorControllerApi#GetConnectorSyncEvents) | **Get** /connectors/{id}/events | Get an inbox connector sync events
[**GetConnectors**](ConnectorControllerApi#GetConnectors) | **Get** /connectors | Get inbox connectors
[**SyncConnector**](ConnectorControllerApi#SyncConnector) | **Post** /connectors/{id}/sync | Sync an inbox connector
[**UpdateConnector**](ConnectorControllerApi#UpdateConnector) | **Put** /connectors/{id} | Update an inbox connector



## CreateConnector

> ConnectorDto CreateConnector(ctx, createConnectorOptions)

Create an inbox connector

Sync emails between external mailboxes and MailSlurp inboxes

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createConnectorOptions** | [**CreateConnectorOptions**](CreateConnectorOptions)|  | 

### Return type

[**ConnectorDto**](ConnectorDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CreateConnectorImapConnection

> ConnectorImapConnectionDto CreateConnectorImapConnection(ctx, id, createConnectorImapConnectionOptions)

Create an inbox connector IMAP connection

Allows the reading of emails in an external mailbox and syncing to a MailSlurp inbox

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()|  | 
**createConnectorImapConnectionOptions** | [**CreateConnectorImapConnectionOptions**](CreateConnectorImapConnectionOptions)|  | 

### Return type

[**ConnectorImapConnectionDto**](ConnectorImapConnectionDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CreateConnectorSmtpConnection

> ConnectorSmtpConnectionDto CreateConnectorSmtpConnection(ctx, id, createConnectorSmtpConnectionOptions)

Create an inbox connector SMTP connection

Allows sending via connector and email is routed to connected inbox and sent via SMTP

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()|  | 
**createConnectorSmtpConnectionOptions** | [**CreateConnectorSmtpConnectionOptions**](CreateConnectorSmtpConnectionOptions)|  | 

### Return type

[**ConnectorSmtpConnectionDto**](ConnectorSmtpConnectionDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DeleteAllConnector

> DeleteAllConnector(ctx, )

Delete all inbox connectors

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


## DeleteConnector

> DeleteConnector(ctx, id)

Delete an inbox connector

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()|  | 

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


## DeleteConnectorImapConnection

> DeleteConnectorImapConnection(ctx, id)

Delete an inbox connector IMAP connection

Delete IMAP connection for external inbox

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()|  | 

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


## DeleteConnectorSmtpConnection

> DeleteConnectorSmtpConnection(ctx, id)

Delete an inbox connector SMTP connection

Delete SMTP connection for external inbox

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()|  | 

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


## GetAllConnectorSyncEvents

> PageConnectorSyncEvents GetAllConnectorSyncEvents(ctx, optional)

Get all inbox connector sync events

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllConnectorSyncEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllConnectorSyncEventsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Optional page index in connector list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in connector list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **since** | **optional.Time**| Filter by created at after the given timestamp | 
 **before** | **optional.Time**| Filter by created at before the given timestamp | 

### Return type

[**PageConnectorSyncEvents**](PageConnectorSyncEvents)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetConnector

> ConnectorDto GetConnector(ctx, id)

Get an inbox connector

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()|  | 

### Return type

[**ConnectorDto**](ConnectorDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetConnectorSyncEvent

> ConnectorSyncEventDto GetConnectorSyncEvent(ctx, id)

Get an inbox connector sync event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()|  | 

### Return type

[**ConnectorSyncEventDto**](ConnectorSyncEventDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetConnectorSyncEvents

> PageConnectorSyncEvents GetConnectorSyncEvents(ctx, id, optional)

Get an inbox connector sync events

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()|  | 
 **optional** | ***GetConnectorSyncEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetConnectorSyncEventsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Optional page index in connector list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in connector list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **since** | **optional.Time**| Filter by created at after the given timestamp | 
 **before** | **optional.Time**| Filter by created at before the given timestamp | 

### Return type

[**PageConnectorSyncEvents**](PageConnectorSyncEvents)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetConnectors

> PageConnector GetConnectors(ctx, optional)

Get inbox connectors

List inbox connectors that sync external emails to MailSlurp inboxes

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetConnectorsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetConnectorsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Optional page index in connector list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in connector list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **since** | **optional.Time**| Filter by created at after the given timestamp | 
 **before** | **optional.Time**| Filter by created at before the given timestamp | 

### Return type

[**PageConnector**](PageConnector)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## SyncConnector

> ConnectorSyncRequestResult SyncConnector(ctx, id)

Sync an inbox connector

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()|  | 

### Return type

[**ConnectorSyncRequestResult**](ConnectorSyncRequestResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## UpdateConnector

> ConnectorDto UpdateConnector(ctx, id, createConnectorOptions)

Update an inbox connector

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()|  | 
**createConnectorOptions** | [**CreateConnectorOptions**](CreateConnectorOptions)|  | 

### Return type

[**ConnectorDto**](ConnectorDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

