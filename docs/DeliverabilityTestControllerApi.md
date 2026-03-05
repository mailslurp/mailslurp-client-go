# MailSlurp\DeliverabilityTestControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelDeliverabilitySimulationJob**](DeliverabilityTestControllerApi#CancelDeliverabilitySimulationJob) | **Post** /test/deliverability/{testId}/simulation-jobs/{jobId}/cancel | Cancel deliverability simulation job
[**CreateDeliverabilitySimulationJob**](DeliverabilityTestControllerApi#CreateDeliverabilitySimulationJob) | **Post** /test/deliverability/{testId}/simulation-jobs | Create deliverability simulation job
[**CreateDeliverabilityTest**](DeliverabilityTestControllerApi#CreateDeliverabilityTest) | **Post** /test/deliverability | Create deliverability/load test
[**DeleteDeliverabilityTest**](DeliverabilityTestControllerApi#DeleteDeliverabilityTest) | **Delete** /test/deliverability/{testId} | Delete deliverability/load test
[**DuplicateDeliverabilityTest**](DeliverabilityTestControllerApi#DuplicateDeliverabilityTest) | **Post** /test/deliverability/{testId}/duplicate | Duplicate deliverability/load test
[**ExportDeliverabilityTestReport**](DeliverabilityTestControllerApi#ExportDeliverabilityTestReport) | **Get** /test/deliverability/{testId}/report/export | Export deliverability/load test report as PDF
[**ExportDeliverabilityTestResults**](DeliverabilityTestControllerApi#ExportDeliverabilityTestResults) | **Get** /test/deliverability/{testId}/results/export | Export deliverability/load test entity results as CSV
[**GetDeliverabilityAnalyticsSeries**](DeliverabilityTestControllerApi#GetDeliverabilityAnalyticsSeries) | **Get** /test/deliverability/analytics/series | Get deliverability analytics time series
[**GetDeliverabilityFailureHotspots**](DeliverabilityTestControllerApi#GetDeliverabilityFailureHotspots) | **Get** /test/deliverability/analytics/hotspots | Get deliverability failure hotspots
[**GetDeliverabilitySimulationJob**](DeliverabilityTestControllerApi#GetDeliverabilitySimulationJob) | **Get** /test/deliverability/{testId}/simulation-jobs/{jobId} | Get deliverability simulation job
[**GetDeliverabilitySimulationJobEvents**](DeliverabilityTestControllerApi#GetDeliverabilitySimulationJobEvents) | **Get** /test/deliverability/{testId}/simulation-jobs/{jobId}/events | Get deliverability simulation job events
[**GetDeliverabilityTest**](DeliverabilityTestControllerApi#GetDeliverabilityTest) | **Get** /test/deliverability/{testId} | Get deliverability/load test
[**GetDeliverabilityTestResults**](DeliverabilityTestControllerApi#GetDeliverabilityTestResults) | **Get** /test/deliverability/{testId}/results | Get deliverability/load test entity results
[**GetDeliverabilityTests**](DeliverabilityTestControllerApi#GetDeliverabilityTests) | **Get** /test/deliverability | List deliverability/load tests
[**GetLatestDeliverabilitySimulationJob**](DeliverabilityTestControllerApi#GetLatestDeliverabilitySimulationJob) | **Get** /test/deliverability/{testId}/simulation-jobs/latest | Get latest deliverability simulation job
[**PauseDeliverabilitySimulationJob**](DeliverabilityTestControllerApi#PauseDeliverabilitySimulationJob) | **Post** /test/deliverability/{testId}/simulation-jobs/{jobId}/pause | Pause deliverability simulation job
[**PauseDeliverabilityTest**](DeliverabilityTestControllerApi#PauseDeliverabilityTest) | **Post** /test/deliverability/{testId}/pause | Pause deliverability/load test
[**PollDeliverabilityTestStatus**](DeliverabilityTestControllerApi#PollDeliverabilityTestStatus) | **Get** /test/deliverability/{testId}/status | Poll deliverability/load test status
[**ResumeDeliverabilitySimulationJob**](DeliverabilityTestControllerApi#ResumeDeliverabilitySimulationJob) | **Post** /test/deliverability/{testId}/simulation-jobs/{jobId}/resume | Resume deliverability simulation job
[**StartDeliverabilityTest**](DeliverabilityTestControllerApi#StartDeliverabilityTest) | **Post** /test/deliverability/{testId}/start | Start or resume deliverability/load test
[**StopDeliverabilityTest**](DeliverabilityTestControllerApi#StopDeliverabilityTest) | **Post** /test/deliverability/{testId}/stop | Stop deliverability/load test
[**UpdateDeliverabilityTest**](DeliverabilityTestControllerApi#UpdateDeliverabilityTest) | **Patch** /test/deliverability/{testId} | Update deliverability/load test



## CancelDeliverabilitySimulationJob

> DeliverabilitySimulationJobDto CancelDeliverabilitySimulationJob(ctx, testId, jobId)

Cancel deliverability simulation job

Cancel a running or paused simulation job.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testId** | [**string**]()|  | 
**jobId** | [**string**]()|  | 

### Return type

[**DeliverabilitySimulationJobDto**](DeliverabilitySimulationJobDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CreateDeliverabilitySimulationJob

> DeliverabilitySimulationJobDto CreateDeliverabilitySimulationJob(ctx, testId, createDeliverabilitySimulationJobOptions)

Create deliverability simulation job

Create and start a simulation job for a running deliverability test. Only one active simulation job is allowed per user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testId** | [**string**]()|  | 
**createDeliverabilitySimulationJobOptions** | [**CreateDeliverabilitySimulationJobOptions**](CreateDeliverabilitySimulationJobOptions)|  | 

### Return type

[**DeliverabilitySimulationJobDto**](DeliverabilitySimulationJobDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CreateDeliverabilityTest

> DeliverabilityTestDto CreateDeliverabilityTest(ctx, createDeliverabilityTestOptions)

Create deliverability/load test

Create a deliverability test for inboxes or phone numbers using ALL, PATTERN, or EXPLICIT selector scope.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createDeliverabilityTestOptions** | [**CreateDeliverabilityTestOptions**](CreateDeliverabilityTestOptions)|  | 

### Return type

[**DeliverabilityTestDto**](DeliverabilityTestDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DeleteDeliverabilityTest

> DeleteResult DeleteDeliverabilityTest(ctx, testId)

Delete deliverability/load test

Delete test and all persisted entity-level results.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testId** | [**string**]()|  | 

### Return type

[**DeleteResult**](DeleteResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DuplicateDeliverabilityTest

> DeliverabilityTestDto DuplicateDeliverabilityTest(ctx, testId)

Duplicate deliverability/load test

Create a fresh deliverability test using an existing test configuration, including selector scope, exclusions, and expectations.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testId** | [**string**]()|  | 

### Return type

[**DeliverabilityTestDto**](DeliverabilityTestDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## ExportDeliverabilityTestReport

> ExportDeliverabilityTestReport(ctx, testId)

Export deliverability/load test report as PDF

Export a PDF report for a terminal deliverability test (COMPLETE, FAILED, or STOPPED), including configuration, summary outcomes, and detailed entity-level results.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testId** | [**string**]()|  | 

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


## ExportDeliverabilityTestResults

> ExportDeliverabilityTestResults(ctx, testId, optional)

Export deliverability/load test entity results as CSV

Export per-entity deliverability results including expectation-level pass/fail counts. The latest status is evaluated with the same polling safeguards before export.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testId** | [**string**]()|  | 
 **optional** | ***ExportDeliverabilityTestResultsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExportDeliverabilityTestResultsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **matched** | **optional.Bool**|  | 

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


## GetDeliverabilityAnalyticsSeries

> DeliverabilityAnalyticsSeriesDto GetDeliverabilityAnalyticsSeries(ctx, optional)

Get deliverability analytics time series

Compare deliverability runs over a time range with bucketed chart metrics and run-level rows for table views.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetDeliverabilityAnalyticsSeriesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDeliverabilityAnalyticsSeriesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **since** | **optional.Time**|  | 
 **before** | **optional.Time**|  | 
 **scope** | **optional.String**|  | 
 **bucket** | **optional.String**|  | [default to DAY]
 **runLimit** | **optional.Int32**|  | 

### Return type

[**DeliverabilityAnalyticsSeriesDto**](DeliverabilityAnalyticsSeriesDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDeliverabilityFailureHotspots

> DeliverabilityFailureHotspotsDto GetDeliverabilityFailureHotspots(ctx, optional)

Get deliverability failure hotspots

Find commonly failing entities and phone country/variant dimensions across deliverability runs in a time range.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetDeliverabilityFailureHotspotsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDeliverabilityFailureHotspotsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **since** | **optional.Time**|  | 
 **before** | **optional.Time**|  | 
 **scope** | **optional.String**|  | 
 **limit** | **optional.Int32**|  | 

### Return type

[**DeliverabilityFailureHotspotsDto**](DeliverabilityFailureHotspotsDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDeliverabilitySimulationJob

> DeliverabilitySimulationJobDto GetDeliverabilitySimulationJob(ctx, testId, jobId)

Get deliverability simulation job

Get simulation job status and progress counters.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testId** | [**string**]()|  | 
**jobId** | [**string**]()|  | 

### Return type

[**DeliverabilitySimulationJobDto**](DeliverabilitySimulationJobDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDeliverabilitySimulationJobEvents

> DeliverabilitySimulationJobEventPageDto GetDeliverabilitySimulationJobEvents(ctx, testId, jobId, optional)

Get deliverability simulation job events

Get paged simulation events including send successes and failures.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testId** | [**string**]()|  | 
**jobId** | [**string**]()|  | 
 **optional** | ***GetDeliverabilitySimulationJobEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDeliverabilitySimulationJobEventsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**|  | [default to 0]
 **size** | **optional.Int32**|  | [default to 20]
 **sort** | **optional.String**|  | [default to DESC]

### Return type

[**DeliverabilitySimulationJobEventPageDto**](DeliverabilitySimulationJobEventPageDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDeliverabilityTest

> DeliverabilityTestDto GetDeliverabilityTest(ctx, testId)

Get deliverability/load test

Get deliverability test configuration and latest progress counters.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testId** | [**string**]()|  | 

### Return type

[**DeliverabilityTestDto**](DeliverabilityTestDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDeliverabilityTestResults

> DeliverabilityEntityResultPageDto GetDeliverabilityTestResults(ctx, testId, optional)

Get deliverability/load test entity results

Get paged per-entity expectation results with optional matched/unmatched filtering.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testId** | [**string**]()|  | 
 **optional** | ***GetDeliverabilityTestResultsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDeliverabilityTestResultsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **matched** | **optional.Bool**|  | 
 **page** | **optional.Int32**|  | [default to 0]
 **size** | **optional.Int32**|  | [default to 20]
 **sort** | **optional.String**|  | [default to ASC]

### Return type

[**DeliverabilityEntityResultPageDto**](DeliverabilityEntityResultPageDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDeliverabilityTests

> DeliverabilityTestPageDto GetDeliverabilityTests(ctx, optional)

List deliverability/load tests

List deliverability tests for the authenticated account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetDeliverabilityTestsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDeliverabilityTestsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page index | [default to 0]
 **size** | **optional.Int32**| Page size | [default to 20]
 **sort** | **optional.String**| Sort direction | [default to DESC]

### Return type

[**DeliverabilityTestPageDto**](DeliverabilityTestPageDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetLatestDeliverabilitySimulationJob

> DeliverabilitySimulationJobDto GetLatestDeliverabilitySimulationJob(ctx, testId)

Get latest deliverability simulation job

Get the most recent simulation job for a deliverability test.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testId** | [**string**]()|  | 

### Return type

[**DeliverabilitySimulationJobDto**](DeliverabilitySimulationJobDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## PauseDeliverabilitySimulationJob

> DeliverabilitySimulationJobDto PauseDeliverabilitySimulationJob(ctx, testId, jobId)

Pause deliverability simulation job

Pause a running simulation job.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testId** | [**string**]()|  | 
**jobId** | [**string**]()|  | 

### Return type

[**DeliverabilitySimulationJobDto**](DeliverabilitySimulationJobDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## PauseDeliverabilityTest

> DeliverabilityTestDto PauseDeliverabilityTest(ctx, testId)

Pause deliverability/load test

Pause a RUNNING or SCHEDULED deliverability test.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testId** | [**string**]()|  | 

### Return type

[**DeliverabilityTestDto**](DeliverabilityTestDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## PollDeliverabilityTestStatus

> DeliverabilityPollStatusResultDto PollDeliverabilityTestStatus(ctx, testId)

Poll deliverability/load test status

Poll test progress. Evaluation is throttled with a 5-second cache window to protect backing data stores.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testId** | [**string**]()|  | 

### Return type

[**DeliverabilityPollStatusResultDto**](DeliverabilityPollStatusResultDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## ResumeDeliverabilitySimulationJob

> DeliverabilitySimulationJobDto ResumeDeliverabilitySimulationJob(ctx, testId, jobId)

Resume deliverability simulation job

Resume a paused simulation job.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testId** | [**string**]()|  | 
**jobId** | [**string**]()|  | 

### Return type

[**DeliverabilitySimulationJobDto**](DeliverabilitySimulationJobDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## StartDeliverabilityTest

> DeliverabilityTestDto StartDeliverabilityTest(ctx, testId)

Start or resume deliverability/load test

Start a CREATED test or resume a PAUSED/SCHEDULED test.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testId** | [**string**]()|  | 

### Return type

[**DeliverabilityTestDto**](DeliverabilityTestDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## StopDeliverabilityTest

> DeliverabilityTestDto StopDeliverabilityTest(ctx, testId)

Stop deliverability/load test

Stop a deliverability test and mark it terminal.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testId** | [**string**]()|  | 

### Return type

[**DeliverabilityTestDto**](DeliverabilityTestDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## UpdateDeliverabilityTest

> DeliverabilityTestDto UpdateDeliverabilityTest(ctx, testId, updateDeliverabilityTestOptions)

Update deliverability/load test

Update metadata, timeout, and expectations for a non-running non-terminal test.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testId** | [**string**]()|  | 
**updateDeliverabilityTestOptions** | [**UpdateDeliverabilityTestOptions**](UpdateDeliverabilityTestOptions)|  | 

### Return type

[**DeliverabilityTestDto**](DeliverabilityTestDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

