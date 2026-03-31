# MailSlurp\DomainMonitorControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompareDomainMonitorRuns**](DomainMonitorControllerApi#CompareDomainMonitorRuns) | **Get** /domain-monitor/monitors/{monitorId}/runs/{runId}/compare/{otherRunId} | Compare two monitor runs
[**CreateDomainMonitor**](DomainMonitorControllerApi#CreateDomainMonitor) | **Post** /domain-monitor/monitors | Create domain monitor
[**CreateDomainMonitorAlertSink**](DomainMonitorControllerApi#CreateDomainMonitorAlertSink) | **Post** /domain-monitor/monitors/{monitorId}/alert-sinks | Create alert sink for monitor
[**DeleteDomainMonitor**](DomainMonitorControllerApi#DeleteDomainMonitor) | **Delete** /domain-monitor/monitors/{monitorId} | Delete domain monitor
[**DeleteDomainMonitorAlertSink**](DomainMonitorControllerApi#DeleteDomainMonitorAlertSink) | **Delete** /domain-monitor/monitors/{monitorId}/alert-sinks/{sinkId} | Delete monitor alert sink
[**GetDomainMonitor**](DomainMonitorControllerApi#GetDomainMonitor) | **Get** /domain-monitor/monitors/{monitorId} | Get domain monitor
[**GetDomainMonitorAlertSinks**](DomainMonitorControllerApi#GetDomainMonitorAlertSinks) | **Get** /domain-monitor/monitors/{monitorId}/alert-sinks | List alert sinks for monitor
[**GetDomainMonitorAuthStack**](DomainMonitorControllerApi#GetDomainMonitorAuthStack) | **Get** /domain-monitor/monitors/{monitorId}/auth-stack | Get current auth stack for monitor domain
[**GetDomainMonitorInsights**](DomainMonitorControllerApi#GetDomainMonitorInsights) | **Get** /domain-monitor/monitors/{monitorId}/insights | Get monitor insights
[**GetDomainMonitorRun**](DomainMonitorControllerApi#GetDomainMonitorRun) | **Get** /domain-monitor/monitors/{monitorId}/runs/{runId} | Get monitor run
[**GetDomainMonitorRuns**](DomainMonitorControllerApi#GetDomainMonitorRuns) | **Get** /domain-monitor/monitors/{monitorId}/runs | List monitor runs
[**GetDomainMonitorSeries**](DomainMonitorControllerApi#GetDomainMonitorSeries) | **Get** /domain-monitor/monitors/{monitorId}/series | Get monitor trend series
[**GetDomainMonitorSummary**](DomainMonitorControllerApi#GetDomainMonitorSummary) | **Get** /domain-monitor/monitors/{monitorId}/summary | Get domain monitor summary
[**GetDomainMonitors**](DomainMonitorControllerApi#GetDomainMonitors) | **Get** /domain-monitor/monitors | List domain monitors
[**RunDomainMonitorNow**](DomainMonitorControllerApi#RunDomainMonitorNow) | **Post** /domain-monitor/monitors/{monitorId}/run-now | Run monitor now
[**RunDueDomainMonitors**](DomainMonitorControllerApi#RunDueDomainMonitors) | **Post** /domain-monitor/monitors/run-due | Run due monitors for user
[**UpdateDomainMonitor**](DomainMonitorControllerApi#UpdateDomainMonitor) | **Put** /domain-monitor/monitors/{monitorId} | Update domain monitor



## CompareDomainMonitorRuns

> DomainMonitorRunComparisonDto CompareDomainMonitorRuns(ctx, monitorId, runId, otherRunId)

Compare two monitor runs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorId** | [**string**]()|  | 
**runId** | [**string**]()|  | 
**otherRunId** | [**string**]()|  | 

### Return type

[**DomainMonitorRunComparisonDto**](DomainMonitorRunComparisonDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CreateDomainMonitor

> DomainMonitorDto CreateDomainMonitor(ctx, createDomainMonitorOptions)

Create domain monitor

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createDomainMonitorOptions** | [**CreateDomainMonitorOptions**](CreateDomainMonitorOptions)|  | 

### Return type

[**DomainMonitorDto**](DomainMonitorDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CreateDomainMonitorAlertSink

> DomainMonitorAlertSinkDto CreateDomainMonitorAlertSink(ctx, monitorId, createDomainMonitorAlertSinkOptions)

Create alert sink for monitor

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorId** | [**string**]()|  | 
**createDomainMonitorAlertSinkOptions** | [**CreateDomainMonitorAlertSinkOptions**](CreateDomainMonitorAlertSinkOptions)|  | 

### Return type

[**DomainMonitorAlertSinkDto**](DomainMonitorAlertSinkDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DeleteDomainMonitor

> DeleteDomainMonitor(ctx, monitorId)

Delete domain monitor

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorId** | [**string**]()|  | 

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


## DeleteDomainMonitorAlertSink

> DeleteDomainMonitorAlertSink(ctx, monitorId, sinkId)

Delete monitor alert sink

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorId** | [**string**]()|  | 
**sinkId** | [**string**]()|  | 

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


## GetDomainMonitor

> DomainMonitorDto GetDomainMonitor(ctx, monitorId)

Get domain monitor

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorId** | [**string**]()|  | 

### Return type

[**DomainMonitorDto**](DomainMonitorDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDomainMonitorAlertSinks

> []DomainMonitorAlertSinkDto GetDomainMonitorAlertSinks(ctx, monitorId)

List alert sinks for monitor

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorId** | [**string**]()|  | 

### Return type

[**[]DomainMonitorAlertSinkDto**](DomainMonitorAlertSinkDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDomainMonitorAuthStack

> CheckEmailAuthStackResults GetDomainMonitorAuthStack(ctx, monitorId, optional)

Get current auth stack for monitor domain

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorId** | [**string**]()|  | 
 **optional** | ***GetDomainMonitorAuthStackOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDomainMonitorAuthStackOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dkimSelector** | **optional.String**|  | 

### Return type

[**CheckEmailAuthStackResults**](CheckEmailAuthStackResults)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDomainMonitorInsights

> DomainMonitorInsightsDto GetDomainMonitorInsights(ctx, monitorId, optional)

Get monitor insights

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorId** | [**string**]()|  | 
 **optional** | ***GetDomainMonitorInsightsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDomainMonitorInsightsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **since** | **optional.Time**|  | 
 **before** | **optional.Time**|  | 

### Return type

[**DomainMonitorInsightsDto**](DomainMonitorInsightsDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDomainMonitorRun

> DomainMonitorRunDto GetDomainMonitorRun(ctx, monitorId, runId)

Get monitor run

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorId** | [**string**]()|  | 
**runId** | [**string**]()|  | 

### Return type

[**DomainMonitorRunDto**](DomainMonitorRunDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDomainMonitorRuns

> []DomainMonitorRunDto GetDomainMonitorRuns(ctx, monitorId, optional)

List monitor runs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorId** | [**string**]()|  | 
 **optional** | ***GetDomainMonitorRunsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDomainMonitorRunsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **since** | **optional.Time**|  | 
 **before** | **optional.Time**|  | 
 **status** | **optional.String**|  | 
 **limit** | **optional.Int32**|  | 

### Return type

[**[]DomainMonitorRunDto**](DomainMonitorRunDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDomainMonitorSeries

> DomainMonitorSeriesDto GetDomainMonitorSeries(ctx, monitorId, optional)

Get monitor trend series

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorId** | [**string**]()|  | 
 **optional** | ***GetDomainMonitorSeriesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDomainMonitorSeriesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **since** | **optional.Time**|  | 
 **before** | **optional.Time**|  | 
 **bucket** | **optional.String**|  | [default to DAY]

### Return type

[**DomainMonitorSeriesDto**](DomainMonitorSeriesDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDomainMonitorSummary

> DomainMonitorSummaryDto GetDomainMonitorSummary(ctx, monitorId, optional)

Get domain monitor summary

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorId** | [**string**]()|  | 
 **optional** | ***GetDomainMonitorSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDomainMonitorSummaryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dkimSelector** | **optional.String**|  | 

### Return type

[**DomainMonitorSummaryDto**](DomainMonitorSummaryDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDomainMonitors

> []DomainMonitorDto GetDomainMonitors(ctx, )

List domain monitors

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]DomainMonitorDto**](DomainMonitorDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## RunDomainMonitorNow

> DomainMonitorRunNowResult RunDomainMonitorNow(ctx, monitorId)

Run monitor now

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorId** | [**string**]()|  | 

### Return type

[**DomainMonitorRunNowResult**](DomainMonitorRunNowResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## RunDueDomainMonitors

> DomainMonitorRunDueResult RunDueDomainMonitors(ctx, optional)

Run due monitors for user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RunDueDomainMonitorsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RunDueDomainMonitorsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **maxRuns** | **optional.Int32**|  | 

### Return type

[**DomainMonitorRunDueResult**](DomainMonitorRunDueResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## UpdateDomainMonitor

> DomainMonitorDto UpdateDomainMonitor(ctx, monitorId, updateDomainMonitorOptions)

Update domain monitor

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorId** | [**string**]()|  | 
**updateDomainMonitorOptions** | [**UpdateDomainMonitorOptions**](UpdateDomainMonitorOptions)|  | 

### Return type

[**DomainMonitorDto**](DomainMonitorDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

