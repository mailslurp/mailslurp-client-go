# MailSlurp\ConnectorControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConnector**](ConnectorControllerApi#CreateConnector) | **Post** /connectors | Create an inbox connector
[**CreateConnectorImapConnection**](ConnectorControllerApi#CreateConnectorImapConnection) | **Post** /connectors/{id}/imap | Create an inbox connector IMAP connection
[**CreateConnectorSmtpConnection**](ConnectorControllerApi#CreateConnectorSmtpConnection) | **Post** /connectors/{id}/smtp | Create an inbox connector SMTP connection
[**CreateConnectorSyncSettings**](ConnectorControllerApi#CreateConnectorSyncSettings) | **Post** /connectors/{id}/sync-settings | Create an inbox connector sync settings
[**CreateConnectorWithOptions**](ConnectorControllerApi#CreateConnectorWithOptions) | **Post** /connectors/withOptions | Create an inbox connector with options
[**DeleteAllConnector**](ConnectorControllerApi#DeleteAllConnector) | **Delete** /connectors | Delete all inbox connectors
[**DeleteConnector**](ConnectorControllerApi#DeleteConnector) | **Delete** /connectors/{id} | Delete an inbox connector
[**DeleteConnectorImapConnection**](ConnectorControllerApi#DeleteConnectorImapConnection) | **Delete** /connectors/{id}/imap | Delete an inbox connector IMAP connection
[**DeleteConnectorSmtpConnection**](ConnectorControllerApi#DeleteConnectorSmtpConnection) | **Delete** /connectors/{id}/smtp | Delete an inbox connector SMTP connection
[**DeleteConnectorSyncSettings**](ConnectorControllerApi#DeleteConnectorSyncSettings) | **Delete** /connectors/{id}/sync-settings | Create an inbox connector sync settings
[**GetAllConnectorEvents**](ConnectorControllerApi#GetAllConnectorEvents) | **Get** /connectors/events | Get all inbox connector events
[**GetConnector**](ConnectorControllerApi#GetConnector) | **Get** /connectors/{id} | Get an inbox connector
[**GetConnectorByEmailAddress**](ConnectorControllerApi#GetConnectorByEmailAddress) | **Get** /connectors/by-email-address | Get connector by email address
[**GetConnectorByInboxId**](ConnectorControllerApi#GetConnectorByInboxId) | **Get** /connectors/by-inbox-id | Get connector by inbox ID
[**GetConnectorByName**](ConnectorControllerApi#GetConnectorByName) | **Get** /connectors/by-name | Get connector by name
[**GetConnectorEvent**](ConnectorControllerApi#GetConnectorEvent) | **Get** /connectors/events/{id} | Get an inbox connector event
[**GetConnectorEvents**](ConnectorControllerApi#GetConnectorEvents) | **Get** /connectors/{id}/events | Get an inbox connector events
[**GetConnectorImapConnection**](ConnectorControllerApi#GetConnectorImapConnection) | **Get** /connectors/{id}/imap | Get an inbox connector IMAP connection
[**GetConnectorProviderSettings**](ConnectorControllerApi#GetConnectorProviderSettings) | **Get** /connectors/provider-settings | Get SMTP and IMAP connection settings for common mail providers
[**GetConnectorSmtpConnection**](ConnectorControllerApi#GetConnectorSmtpConnection) | **Get** /connectors/{id}/smtp | Get an inbox connector SMTP connection
[**GetConnectorSyncSettings**](ConnectorControllerApi#GetConnectorSyncSettings) | **Get** /connectors/{id}/sync-settings | Get an inbox connector sync settings
[**GetConnectors**](ConnectorControllerApi#GetConnectors) | **Get** /connectors | Get inbox connectors
[**SendEmailFromConnector**](ConnectorControllerApi#SendEmailFromConnector) | **Post** /connectors/{id}/send | Send from an inbox connector
[**SyncConnector**](ConnectorControllerApi#SyncConnector) | **Post** /connectors/{id}/sync | Sync an inbox connector
[**TestConnectorImapConnection**](ConnectorControllerApi#TestConnectorImapConnection) | **Post** /connectors/{id}/imap/test | Test an inbox connector IMAP connection
[**TestConnectorImapConnectionOptions**](ConnectorControllerApi#TestConnectorImapConnectionOptions) | **Post** /connectors/connections/imap/test | Test an inbox connector IMAP connection options
[**TestConnectorSmtpConnection**](ConnectorControllerApi#TestConnectorSmtpConnection) | **Post** /connectors/{id}/smtp/test | Test an inbox connector SMTP connection
[**TestConnectorSmtpConnectionOptions**](ConnectorControllerApi#TestConnectorSmtpConnectionOptions) | **Post** /connectors/connections/smtp/test | Test an inbox connector SMTP connection options
[**UpdateConnector**](ConnectorControllerApi#UpdateConnector) | **Put** /connectors/{id} | Update an inbox connector
[**UpdateConnectorImapConnection**](ConnectorControllerApi#UpdateConnectorImapConnection) | **Patch** /connectors/{id}/imap | Update an inbox connector IMAP connection
[**UpdateConnectorSmtpConnection**](ConnectorControllerApi#UpdateConnectorSmtpConnection) | **Patch** /connectors/{id}/smtp | Update an inbox connector SMTP connection



## CreateConnector

> ConnectorDto CreateConnector(ctx, createConnectorOptions, optional)

Create an inbox connector

Sync emails between external mailboxes and MailSlurp inboxes

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createConnectorOptions** | [**CreateConnectorOptions**](CreateConnectorOptions)|  | 
 **optional** | ***CreateConnectorOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateConnectorOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inboxId** | [**optional.Interface of string**]()| Optional inbox ID to associate with the connector | 

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


## CreateConnectorSyncSettings

> ConnectorSyncSettingsDto CreateConnectorSyncSettings(ctx, id, createConnectorSyncSettingsOptions)

Create an inbox connector sync settings

Configure automatic pull or emails from external inboxes using an interval or schedule

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()|  | 
**createConnectorSyncSettingsOptions** | [**CreateConnectorSyncSettingsOptions**](CreateConnectorSyncSettingsOptions)|  | 

### Return type

[**ConnectorSyncSettingsDto**](ConnectorSyncSettingsDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CreateConnectorWithOptions

> ConnectorDto CreateConnectorWithOptions(ctx, createConnectorWithOptions, optional)

Create an inbox connector with options

Sync emails between external mailboxes and MailSlurp inboxes

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createConnectorWithOptions** | [**CreateConnectorWithOptions**](CreateConnectorWithOptions)|  | 
 **optional** | ***CreateConnectorWithOptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateConnectorWithOptionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inboxId** | [**optional.Interface of string**]()| Optional inbox ID to associate with the connector | 

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


## DeleteConnectorSyncSettings

> DeleteConnectorSyncSettings(ctx, id)

Create an inbox connector sync settings

Configure automatic pull or emails from external inboxes using an interval or schedule

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


## GetAllConnectorEvents

> PageConnectorEvents GetAllConnectorEvents(ctx, optional)

Get all inbox connector events

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllConnectorEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllConnectorEventsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | [**optional.Interface of string**]()| Optional connector ID | 
 **page** | **optional.Int32**| Optional page index in connector list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in connector list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **since** | **optional.Time**| Filter by created at after the given timestamp | 
 **before** | **optional.Time**| Filter by created at before the given timestamp | 
 **eventType** | **optional.String**| Filter by event type | 

### Return type

[**PageConnectorEvents**](PageConnectorEvents)

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


## GetConnectorByEmailAddress

> OptionalConnectorDto GetConnectorByEmailAddress(ctx, emailAddress)

Get connector by email address

Find an inbox connector by email address

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailAddress** | **string**| Email address to search for connector by | 

### Return type

[**OptionalConnectorDto**](OptionalConnectorDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetConnectorByInboxId

> OptionalConnectorDto GetConnectorByInboxId(ctx, inboxId)

Get connector by inbox ID

Find an inbox connector by inbox ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()| Inbox ID to search for connector by | 

### Return type

[**OptionalConnectorDto**](OptionalConnectorDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetConnectorByName

> OptionalConnectorDto GetConnectorByName(ctx, name)

Get connector by name

Find an inbox connector by name

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Name to search for connector by | 

### Return type

[**OptionalConnectorDto**](OptionalConnectorDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetConnectorEvent

> ConnectorEventDto GetConnectorEvent(ctx, id)

Get an inbox connector event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()|  | 

### Return type

[**ConnectorEventDto**](ConnectorEventDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetConnectorEvents

> PageConnectorEvents GetConnectorEvents(ctx, id, optional)

Get an inbox connector events

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()|  | 
 **optional** | ***GetConnectorEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetConnectorEventsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Optional page index in connector list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in connector list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **since** | **optional.Time**| Filter by created at after the given timestamp | 
 **before** | **optional.Time**| Filter by created at before the given timestamp | 
 **eventType** | **optional.String**| Filter by event type | 

### Return type

[**PageConnectorEvents**](PageConnectorEvents)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetConnectorImapConnection

> OptionalConnectorImapConnectionDto GetConnectorImapConnection(ctx, id)

Get an inbox connector IMAP connection

Get IMAP connection for external inbox

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()|  | 

### Return type

[**OptionalConnectorImapConnectionDto**](OptionalConnectorImapConnectionDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetConnectorProviderSettings

> ConnectorProviderSettingsDto GetConnectorProviderSettings(ctx, )

Get SMTP and IMAP connection settings for common mail providers

Get common mail provider SMTP and IMAP connection settings

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ConnectorProviderSettingsDto**](ConnectorProviderSettingsDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetConnectorSmtpConnection

> OptionalConnectorSmtpConnectionDto GetConnectorSmtpConnection(ctx, id)

Get an inbox connector SMTP connection

Get SMTP connection for external inbox

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()|  | 

### Return type

[**OptionalConnectorSmtpConnectionDto**](OptionalConnectorSmtpConnectionDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetConnectorSyncSettings

> OptionalConnectorSyncSettingsDto GetConnectorSyncSettings(ctx, id)

Get an inbox connector sync settings

Get sync settings for connection with external inbox

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()|  | 

### Return type

[**OptionalConnectorSyncSettingsDto**](OptionalConnectorSyncSettingsDto)

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


## SendEmailFromConnector

> SentEmailDto SendEmailFromConnector(ctx, id, sendEmailOptions, optional)

Send from an inbox connector

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()|  | 
**sendEmailOptions** | [**SendEmailOptions**](SendEmailOptions)|  | 
 **optional** | ***SendEmailFromConnectorOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SendEmailFromConnectorOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **useFallback** | **optional.Bool**|  | 

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


## SyncConnector

> ConnectorSyncRequestResult SyncConnector(ctx, id, optional)

Sync an inbox connector

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()|  | 
 **optional** | ***SyncConnectorOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SyncConnectorOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **since** | **optional.Time**| Date to request emails since | 
 **folder** | **optional.String**| Which folder to sync emails with | 
 **logging** | **optional.Bool**| Enable or disable logging for the sync operation | 

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


## TestConnectorImapConnection

> ConnectorImapConnectionTestResult TestConnectorImapConnection(ctx, id, optional)

Test an inbox connector IMAP connection

Test the IMAP connection for a connector

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()|  | 
 **optional** | ***TestConnectorImapConnectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TestConnectorImapConnectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createConnectorImapConnectionOptions** | [**optional.Interface of CreateConnectorImapConnectionOptions**](CreateConnectorImapConnectionOptions)|  | 

### Return type

[**ConnectorImapConnectionTestResult**](ConnectorImapConnectionTestResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## TestConnectorImapConnectionOptions

> ConnectorImapConnectionTestResult TestConnectorImapConnectionOptions(ctx, createConnectorImapConnectionOptions)

Test an inbox connector IMAP connection options

Test the IMAP connection options for a connector

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createConnectorImapConnectionOptions** | [**CreateConnectorImapConnectionOptions**](CreateConnectorImapConnectionOptions)|  | 

### Return type

[**ConnectorImapConnectionTestResult**](ConnectorImapConnectionTestResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## TestConnectorSmtpConnection

> ConnectorSmtpConnectionTestResult TestConnectorSmtpConnection(ctx, id, optional)

Test an inbox connector SMTP connection

Test the SMTP connection for a connector

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()|  | 
 **optional** | ***TestConnectorSmtpConnectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TestConnectorSmtpConnectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createConnectorSmtpConnectionOptions** | [**optional.Interface of CreateConnectorSmtpConnectionOptions**](CreateConnectorSmtpConnectionOptions)|  | 

### Return type

[**ConnectorSmtpConnectionTestResult**](ConnectorSmtpConnectionTestResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## TestConnectorSmtpConnectionOptions

> ConnectorSmtpConnectionTestResult TestConnectorSmtpConnectionOptions(ctx, createConnectorSmtpConnectionOptions)

Test an inbox connector SMTP connection options

Test the SMTP connection options for a connector

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createConnectorSmtpConnectionOptions** | [**CreateConnectorSmtpConnectionOptions**](CreateConnectorSmtpConnectionOptions)|  | 

### Return type

[**ConnectorSmtpConnectionTestResult**](ConnectorSmtpConnectionTestResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
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


## UpdateConnectorImapConnection

> ConnectorImapConnectionDto UpdateConnectorImapConnection(ctx, id, createConnectorImapConnectionOptions)

Update an inbox connector IMAP connection

Update IMAP connection for external inbox

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


## UpdateConnectorSmtpConnection

> ConnectorSmtpConnectionDto UpdateConnectorSmtpConnection(ctx, id, createConnectorSmtpConnectionOptions)

Update an inbox connector SMTP connection

Update SMTP connection for external inbox

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

