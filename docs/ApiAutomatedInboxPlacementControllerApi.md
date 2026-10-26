# MailSlurp\ApiAutomatedInboxPlacementControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAutomatedInboxPlacementConnection**](ApiAutomatedInboxPlacementControllerApi#CreateAutomatedInboxPlacementConnection) | **Post** /inbox-placement-tests/automations/connections | Create an encrypted sender connection
[**CreateAutomatedInboxPlacementSchedule**](ApiAutomatedInboxPlacementControllerApi#CreateAutomatedInboxPlacementSchedule) | **Post** /inbox-placement-tests/automations/schedules | Create an automated inbox placement schedule
[**DeleteAutomatedInboxPlacementConnection**](ApiAutomatedInboxPlacementControllerApi#DeleteAutomatedInboxPlacementConnection) | **Delete** /inbox-placement-tests/automations/connections/{connectionId} | Disconnect a sender connection
[**DeleteAutomatedInboxPlacementSchedule**](ApiAutomatedInboxPlacementControllerApi#DeleteAutomatedInboxPlacementSchedule) | **Delete** /inbox-placement-tests/automations/schedules/{scheduleId} | Delete an automated inbox placement schedule
[**GetAutomatedInboxPlacementConnection**](ApiAutomatedInboxPlacementControllerApi#GetAutomatedInboxPlacementConnection) | **Get** /inbox-placement-tests/automations/connections/{connectionId} | Get a sender connection
[**GetAutomatedInboxPlacementConnections**](ApiAutomatedInboxPlacementControllerApi#GetAutomatedInboxPlacementConnections) | **Get** /inbox-placement-tests/automations/connections | List sender connections
[**GetAutomatedInboxPlacementExecutionEvents**](ApiAutomatedInboxPlacementControllerApi#GetAutomatedInboxPlacementExecutionEvents) | **Get** /inbox-placement-tests/automations/schedules/{scheduleId}/executions/{executionId}/events | List non-secret execution events
[**GetAutomatedInboxPlacementExecutions**](ApiAutomatedInboxPlacementControllerApi#GetAutomatedInboxPlacementExecutions) | **Get** /inbox-placement-tests/automations/schedules/{scheduleId}/executions | List schedule executions
[**GetAutomatedInboxPlacementProviders**](ApiAutomatedInboxPlacementControllerApi#GetAutomatedInboxPlacementProviders) | **Get** /inbox-placement-tests/automations/providers | List automated inbox placement sender providers
[**GetAutomatedInboxPlacementSchedule**](ApiAutomatedInboxPlacementControllerApi#GetAutomatedInboxPlacementSchedule) | **Get** /inbox-placement-tests/automations/schedules/{scheduleId} | Get an automated inbox placement schedule
[**GetAutomatedInboxPlacementScheduleAnalyticsSeries**](ApiAutomatedInboxPlacementControllerApi#GetAutomatedInboxPlacementScheduleAnalyticsSeries) | **Get** /inbox-placement-tests/automations/schedules/{scheduleId}/analytics/series | Get automation inbox placement analytics over time
[**GetAutomatedInboxPlacementSchedules**](ApiAutomatedInboxPlacementControllerApi#GetAutomatedInboxPlacementSchedules) | **Get** /inbox-placement-tests/automations/schedules | List automated inbox placement schedules
[**RotateAutomatedInboxPlacementConnectionCredentials**](ApiAutomatedInboxPlacementControllerApi#RotateAutomatedInboxPlacementConnectionCredentials) | **Post** /inbox-placement-tests/automations/connections/{connectionId}/credentials | Rotate sender connection credentials
[**RunAutomatedInboxPlacementSchedule**](ApiAutomatedInboxPlacementControllerApi#RunAutomatedInboxPlacementSchedule) | **Post** /inbox-placement-tests/automations/schedules/{scheduleId}/run | Run an automated inbox placement schedule now
[**UpdateAutomatedInboxPlacementConnection**](ApiAutomatedInboxPlacementControllerApi#UpdateAutomatedInboxPlacementConnection) | **Patch** /inbox-placement-tests/automations/connections/{connectionId} | Update a sender connection
[**UpdateAutomatedInboxPlacementSchedule**](ApiAutomatedInboxPlacementControllerApi#UpdateAutomatedInboxPlacementSchedule) | **Patch** /inbox-placement-tests/automations/schedules/{scheduleId} | Update an automated inbox placement schedule
[**VerifyAutomatedInboxPlacementConnection**](ApiAutomatedInboxPlacementControllerApi#VerifyAutomatedInboxPlacementConnection) | **Post** /inbox-placement-tests/automations/connections/{connectionId}/verify | Verify a sender connection



## CreateAutomatedInboxPlacementConnection

> AutomatedInboxPlacementConnectionDto CreateAutomatedInboxPlacementConnection(ctx, createAutomatedInboxPlacementConnectionOptions)

Create an encrypted sender connection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createAutomatedInboxPlacementConnectionOptions** | [**CreateAutomatedInboxPlacementConnectionOptions**](CreateAutomatedInboxPlacementConnectionOptions)|  | 

### Return type

[**AutomatedInboxPlacementConnectionDto**](AutomatedInboxPlacementConnectionDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CreateAutomatedInboxPlacementSchedule

> AutomatedInboxPlacementScheduleDto CreateAutomatedInboxPlacementSchedule(ctx, createAutomatedInboxPlacementScheduleOptions)

Create an automated inbox placement schedule

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createAutomatedInboxPlacementScheduleOptions** | [**CreateAutomatedInboxPlacementScheduleOptions**](CreateAutomatedInboxPlacementScheduleOptions)|  | 

### Return type

[**AutomatedInboxPlacementScheduleDto**](AutomatedInboxPlacementScheduleDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DeleteAutomatedInboxPlacementConnection

> DeleteAutomatedInboxPlacementConnection(ctx, connectionId)

Disconnect a sender connection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | [**string**]()|  | 

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


## DeleteAutomatedInboxPlacementSchedule

> DeleteAutomatedInboxPlacementSchedule(ctx, scheduleId)

Delete an automated inbox placement schedule

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scheduleId** | [**string**]()|  | 

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


## GetAutomatedInboxPlacementConnection

> AutomatedInboxPlacementConnectionDto GetAutomatedInboxPlacementConnection(ctx, connectionId)

Get a sender connection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | [**string**]()|  | 

### Return type

[**AutomatedInboxPlacementConnectionDto**](AutomatedInboxPlacementConnectionDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetAutomatedInboxPlacementConnections

> []AutomatedInboxPlacementConnectionDto GetAutomatedInboxPlacementConnections(ctx, )

List sender connections

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]AutomatedInboxPlacementConnectionDto**](AutomatedInboxPlacementConnectionDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetAutomatedInboxPlacementExecutionEvents

> []AutomatedInboxPlacementEventDto GetAutomatedInboxPlacementExecutionEvents(ctx, scheduleId, executionId)

List non-secret execution events

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scheduleId** | [**string**]()|  | 
**executionId** | [**string**]()|  | 

### Return type

[**[]AutomatedInboxPlacementEventDto**](AutomatedInboxPlacementEventDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetAutomatedInboxPlacementExecutions

> []AutomatedInboxPlacementExecutionDto GetAutomatedInboxPlacementExecutions(ctx, scheduleId, optional)

List schedule executions

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scheduleId** | [**string**]()|  | 
 **optional** | ***GetAutomatedInboxPlacementExecutionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAutomatedInboxPlacementExecutionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 50]

### Return type

[**[]AutomatedInboxPlacementExecutionDto**](AutomatedInboxPlacementExecutionDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetAutomatedInboxPlacementProviders

> []AutomatedInboxPlacementProviderDto GetAutomatedInboxPlacementProviders(ctx, )

List automated inbox placement sender providers

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]AutomatedInboxPlacementProviderDto**](AutomatedInboxPlacementProviderDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetAutomatedInboxPlacementSchedule

> AutomatedInboxPlacementScheduleDto GetAutomatedInboxPlacementSchedule(ctx, scheduleId)

Get an automated inbox placement schedule

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scheduleId** | [**string**]()|  | 

### Return type

[**AutomatedInboxPlacementScheduleDto**](AutomatedInboxPlacementScheduleDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetAutomatedInboxPlacementScheduleAnalyticsSeries

> InboxPlacementAnalyticsSeriesDto GetAutomatedInboxPlacementScheduleAnalyticsSeries(ctx, scheduleId, optional)

Get automation inbox placement analytics over time

Returns chart-ready metrics for billed placement runs created by this schedule only.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scheduleId** | [**string**]()|  | 
 **optional** | ***GetAutomatedInboxPlacementScheduleAnalyticsSeriesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAutomatedInboxPlacementScheduleAnalyticsSeriesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **since** | **optional.Time**|  | 
 **before** | **optional.Time**|  | 
 **bucket** | **optional.String**|  | [default to DAY]
 **runLimit** | **optional.Int32**|  | 

### Return type

[**InboxPlacementAnalyticsSeriesDto**](InboxPlacementAnalyticsSeriesDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetAutomatedInboxPlacementSchedules

> []AutomatedInboxPlacementScheduleDto GetAutomatedInboxPlacementSchedules(ctx, )

List automated inbox placement schedules

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]AutomatedInboxPlacementScheduleDto**](AutomatedInboxPlacementScheduleDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## RotateAutomatedInboxPlacementConnectionCredentials

> AutomatedInboxPlacementConnectionDto RotateAutomatedInboxPlacementConnectionCredentials(ctx, connectionId, rotateAutomatedInboxPlacementConnectionCredentialsOptions)

Rotate sender connection credentials

Verifies the replacement SMTP settings before atomically replacing the encrypted credential envelope. A failed verification leaves the existing connection unchanged.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | [**string**]()|  | 
**rotateAutomatedInboxPlacementConnectionCredentialsOptions** | [**RotateAutomatedInboxPlacementConnectionCredentialsOptions**](RotateAutomatedInboxPlacementConnectionCredentialsOptions)|  | 

### Return type

[**AutomatedInboxPlacementConnectionDto**](AutomatedInboxPlacementConnectionDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## RunAutomatedInboxPlacementSchedule

> AutomatedInboxPlacementExecutionDto RunAutomatedInboxPlacementSchedule(ctx, scheduleId)

Run an automated inbox placement schedule now

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scheduleId** | [**string**]()|  | 

### Return type

[**AutomatedInboxPlacementExecutionDto**](AutomatedInboxPlacementExecutionDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## UpdateAutomatedInboxPlacementConnection

> AutomatedInboxPlacementConnectionDto UpdateAutomatedInboxPlacementConnection(ctx, connectionId, updateAutomatedInboxPlacementConnectionOptions)

Update a sender connection

Updates non-secret connection metadata. Use the credentials endpoint to replace SMTP credentials or transport settings.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | [**string**]()|  | 
**updateAutomatedInboxPlacementConnectionOptions** | [**UpdateAutomatedInboxPlacementConnectionOptions**](UpdateAutomatedInboxPlacementConnectionOptions)|  | 

### Return type

[**AutomatedInboxPlacementConnectionDto**](AutomatedInboxPlacementConnectionDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## UpdateAutomatedInboxPlacementSchedule

> AutomatedInboxPlacementScheduleDto UpdateAutomatedInboxPlacementSchedule(ctx, scheduleId, updateAutomatedInboxPlacementScheduleOptions)

Update an automated inbox placement schedule

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scheduleId** | [**string**]()|  | 
**updateAutomatedInboxPlacementScheduleOptions** | [**UpdateAutomatedInboxPlacementScheduleOptions**](UpdateAutomatedInboxPlacementScheduleOptions)|  | 

### Return type

[**AutomatedInboxPlacementScheduleDto**](AutomatedInboxPlacementScheduleDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## VerifyAutomatedInboxPlacementConnection

> AutomatedInboxPlacementConnectionDto VerifyAutomatedInboxPlacementConnection(ctx, connectionId)

Verify a sender connection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | [**string**]()|  | 

### Return type

[**AutomatedInboxPlacementConnectionDto**](AutomatedInboxPlacementConnectionDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

