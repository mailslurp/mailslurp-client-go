# MailSlurp\ApiInboxPlacementTestControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfirmInboxPlacementTestSent**](ApiInboxPlacementTestControllerApi#ConfirmInboxPlacementTestSent) | **Post** /inbox-placement-tests/{id}/confirm-sent | Confirm an inbox placement test email has been sent
[**CreateInboxPlacementTest**](ApiInboxPlacementTestControllerApi#CreateInboxPlacementTest) | **Post** /inbox-placement-tests | Create a new direct-send inbox placement test
[**ExportInboxPlacementSeedList**](ApiInboxPlacementTestControllerApi#ExportInboxPlacementSeedList) | **Get** /inbox-placement-tests/{id}/seed-list.csv | Download inbox placement seed addresses as an email-platform CSV
[**ExportInboxPlacementTestResults**](ApiInboxPlacementTestControllerApi#ExportInboxPlacementTestResults) | **Get** /inbox-placement-tests/{id}/results/export | Export inbox placement test results as CSV
[**GetInboxPlacementAnalyticsBreakdown**](ApiInboxPlacementTestControllerApi#GetInboxPlacementAnalyticsBreakdown) | **Get** /inbox-placement-tests/analytics/breakdown | Get inbox placement analytics breakdowns
[**GetInboxPlacementAnalyticsSeries**](ApiInboxPlacementTestControllerApi#GetInboxPlacementAnalyticsSeries) | **Get** /inbox-placement-tests/analytics/series | Get inbox placement analytics time series
[**GetInboxPlacementPublicShare**](ApiInboxPlacementTestControllerApi#GetInboxPlacementPublicShare) | **Get** /inbox-placement-tests/share/{shareToken} | Get a public inbox placement share
[**GetInboxPlacementPublicShareAnalysis**](ApiInboxPlacementTestControllerApi#GetInboxPlacementPublicShareAnalysis) | **Get** /inbox-placement-tests/share/{shareToken}/analysis | Get public inbox placement share analysis
[**GetInboxPlacementTest**](ApiInboxPlacementTestControllerApi#GetInboxPlacementTest) | **Get** /inbox-placement-tests/{id} | Get inbox placement test run
[**GetInboxPlacementTestAnalysis**](ApiInboxPlacementTestControllerApi#GetInboxPlacementTestAnalysis) | **Get** /inbox-placement-tests/{id}/analysis | Get inbox placement run analysis
[**GetInboxPlacementTestResults**](ApiInboxPlacementTestControllerApi#GetInboxPlacementTestResults) | **Get** /inbox-placement-tests/{id}/results | Get inbox placement test results
[**GetInboxPlacementTests**](ApiInboxPlacementTestControllerApi#GetInboxPlacementTests) | **Get** /inbox-placement-tests | List inbox placement test runs
[**PrepareInboxPlacementTest**](ApiInboxPlacementTestControllerApi#PrepareInboxPlacementTest) | **Post** /inbox-placement-tests/prepare | Prepare a direct-send inbox placement test without consuming allowance



## ConfirmInboxPlacementTestSent

> InboxPlacementTestRunDto ConfirmInboxPlacementTestSent(ctx, id)

Confirm an inbox placement test email has been sent

Reconcile any messages that arrived after seed reservation, commit allowance, and move a prepared dashboard test into its active lifecycle. This operation is idempotent.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()|  | 

### Return type

[**InboxPlacementTestRunDto**](InboxPlacementTestRunDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CreateInboxPlacementTest

> InboxPlacementTestRunDto CreateInboxPlacementTest(ctx, createInboxPlacementTestOptions)

Create a new direct-send inbox placement test

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createInboxPlacementTestOptions** | [**CreateInboxPlacementTestOptions**](CreateInboxPlacementTestOptions)|  | 

### Return type

[**InboxPlacementTestRunDto**](InboxPlacementTestRunDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## ExportInboxPlacementSeedList

> ExportInboxPlacementSeedList(ctx, id, optional)

Download inbox placement seed addresses as an email-platform CSV

Export the current run's seed addresses with headers recognized by the selected email platform. Downloading a prepared seed list does not confirm the send or consume allowance.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()|  | 
 **optional** | ***ExportInboxPlacementSeedListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExportInboxPlacementSeedListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **provider** | **optional.String**|  | [default to GENERIC]

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


## ExportInboxPlacementTestResults

> ExportInboxPlacementTestResults(ctx, id)

Export inbox placement test results as CSV

Export one row per placement target, including run-level delivery totals, 0-10 placement scores, and target placement outcome fields.

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


## GetInboxPlacementAnalyticsBreakdown

> InboxPlacementAnalyticsBreakdownDto GetInboxPlacementAnalyticsBreakdown(ctx, optional)

Get inbox placement analytics breakdowns

Return aggregated folder, provider, sender-domain, and from-email breakdowns for inbox placement runs in a time range.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetInboxPlacementAnalyticsBreakdownOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetInboxPlacementAnalyticsBreakdownOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **since** | **optional.Time**|  | 
 **before** | **optional.Time**|  | 
 **senderDomain** | **optional.String**|  | 
 **fromEmail** | **optional.String**|  | 
 **limit** | **optional.Int32**|  | 

### Return type

[**InboxPlacementAnalyticsBreakdownDto**](InboxPlacementAnalyticsBreakdownDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetInboxPlacementAnalyticsSeries

> InboxPlacementAnalyticsSeriesDto GetInboxPlacementAnalyticsSeries(ctx, optional)

Get inbox placement analytics time series

Return chart-ready inbox placement metrics over time, with optional sender/domain filters and grouped comparisons for sender domains or from-email addresses.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetInboxPlacementAnalyticsSeriesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetInboxPlacementAnalyticsSeriesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **since** | **optional.Time**|  | 
 **before** | **optional.Time**|  | 
 **bucket** | **optional.String**|  | [default to DAY]
 **groupBy** | **optional.String**|  | [default to NONE]
 **senderDomain** | **optional.String**|  | 
 **fromEmail** | **optional.String**|  | 
 **runLimit** | **optional.Int32**|  | 
 **groupLimit** | **optional.Int32**|  | 

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


## GetInboxPlacementPublicShare

> InboxPlacementPublicShareDto GetInboxPlacementPublicShare(ctx, shareToken)

Get a public inbox placement share

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shareToken** | **string**|  | 

### Return type

[**InboxPlacementPublicShareDto**](InboxPlacementPublicShareDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetInboxPlacementPublicShareAnalysis

> InboxPlacementAnalysisBundleDto GetInboxPlacementPublicShareAnalysis(ctx, shareToken)

Get public inbox placement share analysis

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shareToken** | **string**|  | 

### Return type

[**InboxPlacementAnalysisBundleDto**](InboxPlacementAnalysisBundleDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetInboxPlacementTest

> InboxPlacementTestRunDto GetInboxPlacementTest(ctx, id)

Get inbox placement test run

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()|  | 

### Return type

[**InboxPlacementTestRunDto**](InboxPlacementTestRunDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetInboxPlacementTestAnalysis

> InboxPlacementAnalysisBundleDto GetInboxPlacementTestAnalysis(ctx, id)

Get inbox placement run analysis

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()|  | 

### Return type

[**InboxPlacementAnalysisBundleDto**](InboxPlacementAnalysisBundleDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetInboxPlacementTestResults

> InboxPlacementTestResultsDto GetInboxPlacementTestResults(ctx, id)

Get inbox placement test results

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()|  | 

### Return type

[**InboxPlacementTestResultsDto**](InboxPlacementTestResultsDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetInboxPlacementTests

> []InboxPlacementTestRunDto GetInboxPlacementTests(ctx, optional)

List inbox placement test runs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetInboxPlacementTestsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetInboxPlacementTestsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**|  | [default to 50]

### Return type

[**[]InboxPlacementTestRunDto**](InboxPlacementTestRunDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## PrepareInboxPlacementTest

> InboxPlacementTestRunDto PrepareInboxPlacementTest(ctx, createInboxPlacementTestOptions)

Prepare a direct-send inbox placement test without consuming allowance

Reserve seed addresses for the dashboard create flow. The run remains staged and allowance is not committed until send confirmation. Reuses a recent active preparation when possible.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createInboxPlacementTestOptions** | [**CreateInboxPlacementTestOptions**](CreateInboxPlacementTestOptions)|  | 

### Return type

[**InboxPlacementTestRunDto**](InboxPlacementTestRunDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

