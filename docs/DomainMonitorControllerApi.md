# MailSlurp\DomainMonitorControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDomainMonitor**](DomainMonitorControllerApi#CreateDomainMonitor) | **Post** /domain-monitor/monitors | Create domain monitor
[**CreateDomainMonitorAlertSink**](DomainMonitorControllerApi#CreateDomainMonitorAlertSink) | **Post** /domain-monitor/monitors/{monitorId}/alert-sinks | Create alert sink for monitor
[**DeleteDomainMonitor**](DomainMonitorControllerApi#DeleteDomainMonitor) | **Delete** /domain-monitor/monitors/{monitorId} | Delete domain monitor
[**DeleteDomainMonitorAlertSink**](DomainMonitorControllerApi#DeleteDomainMonitorAlertSink) | **Delete** /domain-monitor/monitors/{monitorId}/alert-sinks/{sinkId} | Delete monitor alert sink
[**GetDomainMonitor**](DomainMonitorControllerApi#GetDomainMonitor) | **Get** /domain-monitor/monitors/{monitorId} | Get domain monitor
[**GetDomainMonitorAlertSinks**](DomainMonitorControllerApi#GetDomainMonitorAlertSinks) | **Get** /domain-monitor/monitors/{monitorId}/alert-sinks | List alert sinks for monitor
[**GetDomainMonitorInsights**](DomainMonitorControllerApi#GetDomainMonitorInsights) | **Get** /domain-monitor/monitors/{monitorId}/insights | Get monitor insights
[**GetDomainMonitorRuns**](DomainMonitorControllerApi#GetDomainMonitorRuns) | **Get** /domain-monitor/monitors/{monitorId}/runs | List monitor runs
[**GetDomainMonitorSeries**](DomainMonitorControllerApi#GetDomainMonitorSeries) | **Get** /domain-monitor/monitors/{monitorId}/series | Get monitor trend series
[**GetDomainMonitors**](DomainMonitorControllerApi#GetDomainMonitors) | **Get** /domain-monitor/monitors | List domain monitors
[**RunDomainMonitorNow**](DomainMonitorControllerApi#RunDomainMonitorNow) | **Post** /domain-monitor/monitors/{monitorId}/run-now | Run monitor now
[**RunDueDomainMonitors**](DomainMonitorControllerApi#RunDueDomainMonitors) | **Post** /domain-monitor/monitors/run-due | Run due monitors for user
[**UpdateDomainMonitor**](DomainMonitorControllerApi#UpdateDomainMonitor) | **Put** /domain-monitor/monitors/{monitorId} | Update domain monitor



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

