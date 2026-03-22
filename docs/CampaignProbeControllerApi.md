# MailSlurp\CampaignProbeControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompareCampaignProbeRuns**](CampaignProbeControllerApi#CompareCampaignProbeRuns) | **Get** /campaign-probe/probes/{probeId}/runs/{runId}/compare/{otherRunId} | Compare two campaign probe runs
[**CreateCampaignProbe**](CampaignProbeControllerApi#CreateCampaignProbe) | **Post** /campaign-probe/probes | Create campaign probe
[**DeleteCampaignProbe**](CampaignProbeControllerApi#DeleteCampaignProbe) | **Delete** /campaign-probe/probes/{probeId} | Delete campaign probe
[**GetCampaignProbe**](CampaignProbeControllerApi#GetCampaignProbe) | **Get** /campaign-probe/probes/{probeId} | Get campaign probe
[**GetCampaignProbeInsights**](CampaignProbeControllerApi#GetCampaignProbeInsights) | **Get** /campaign-probe/probes/{probeId}/insights | Get campaign probe insights
[**GetCampaignProbeRun**](CampaignProbeControllerApi#GetCampaignProbeRun) | **Get** /campaign-probe/probes/{probeId}/runs/{runId} | Get campaign probe run
[**GetCampaignProbeRuns**](CampaignProbeControllerApi#GetCampaignProbeRuns) | **Get** /campaign-probe/probes/{probeId}/runs | List campaign probe runs
[**GetCampaignProbeSeries**](CampaignProbeControllerApi#GetCampaignProbeSeries) | **Get** /campaign-probe/probes/{probeId}/series | Get campaign probe trend series
[**GetCampaignProbes**](CampaignProbeControllerApi#GetCampaignProbes) | **Get** /campaign-probe/probes | List campaign probes
[**RunCampaignProbeNow**](CampaignProbeControllerApi#RunCampaignProbeNow) | **Post** /campaign-probe/probes/{probeId}/run-now | Run campaign probe now
[**RunDueCampaignProbes**](CampaignProbeControllerApi#RunDueCampaignProbes) | **Post** /campaign-probe/probes/run-due | Run due campaign probes for user
[**UpdateCampaignProbe**](CampaignProbeControllerApi#UpdateCampaignProbe) | **Put** /campaign-probe/probes/{probeId} | Update campaign probe



## CompareCampaignProbeRuns

> CampaignProbeRunComparisonDto CompareCampaignProbeRuns(ctx, probeId, runId, otherRunId)

Compare two campaign probe runs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**probeId** | [**string**]()|  | 
**runId** | [**string**]()|  | 
**otherRunId** | [**string**]()|  | 

### Return type

[**CampaignProbeRunComparisonDto**](CampaignProbeRunComparisonDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CreateCampaignProbe

> CampaignProbeDto CreateCampaignProbe(ctx, createCampaignProbeOptions)

Create campaign probe

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createCampaignProbeOptions** | [**CreateCampaignProbeOptions**](CreateCampaignProbeOptions)|  | 

### Return type

[**CampaignProbeDto**](CampaignProbeDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DeleteCampaignProbe

> DeleteCampaignProbe(ctx, probeId)

Delete campaign probe

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**probeId** | [**string**]()|  | 

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


## GetCampaignProbe

> CampaignProbeDto GetCampaignProbe(ctx, probeId)

Get campaign probe

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**probeId** | [**string**]()|  | 

### Return type

[**CampaignProbeDto**](CampaignProbeDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetCampaignProbeInsights

> CampaignProbeInsightsDto GetCampaignProbeInsights(ctx, probeId, optional)

Get campaign probe insights

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**probeId** | [**string**]()|  | 
 **optional** | ***GetCampaignProbeInsightsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCampaignProbeInsightsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **since** | **optional.Time**|  | 
 **before** | **optional.Time**|  | 

### Return type

[**CampaignProbeInsightsDto**](CampaignProbeInsightsDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetCampaignProbeRun

> CampaignProbeRunDto GetCampaignProbeRun(ctx, probeId, runId)

Get campaign probe run

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**probeId** | [**string**]()|  | 
**runId** | [**string**]()|  | 

### Return type

[**CampaignProbeRunDto**](CampaignProbeRunDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetCampaignProbeRuns

> []CampaignProbeRunDto GetCampaignProbeRuns(ctx, probeId, optional)

List campaign probe runs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**probeId** | [**string**]()|  | 
 **optional** | ***GetCampaignProbeRunsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCampaignProbeRunsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **since** | **optional.Time**|  | 
 **before** | **optional.Time**|  | 
 **status** | **optional.String**|  | 
 **limit** | **optional.Int32**|  | 

### Return type

[**[]CampaignProbeRunDto**](CampaignProbeRunDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetCampaignProbeSeries

> CampaignProbeSeriesDto GetCampaignProbeSeries(ctx, probeId, optional)

Get campaign probe trend series

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**probeId** | [**string**]()|  | 
 **optional** | ***GetCampaignProbeSeriesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCampaignProbeSeriesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **since** | **optional.Time**|  | 
 **before** | **optional.Time**|  | 
 **bucket** | **optional.String**|  | [default to DAY]

### Return type

[**CampaignProbeSeriesDto**](CampaignProbeSeriesDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetCampaignProbes

> []CampaignProbeDto GetCampaignProbes(ctx, )

List campaign probes

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]CampaignProbeDto**](CampaignProbeDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## RunCampaignProbeNow

> CampaignProbeRunNowResult RunCampaignProbeNow(ctx, probeId, createCampaignProbeRunOptions)

Run campaign probe now

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**probeId** | [**string**]()|  | 
**createCampaignProbeRunOptions** | [**CreateCampaignProbeRunOptions**](CreateCampaignProbeRunOptions)|  | 

### Return type

[**CampaignProbeRunNowResult**](CampaignProbeRunNowResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## RunDueCampaignProbes

> CampaignProbeRunDueResult RunDueCampaignProbes(ctx, optional)

Run due campaign probes for user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RunDueCampaignProbesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RunDueCampaignProbesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **maxRuns** | **optional.Int32**|  | 

### Return type

[**CampaignProbeRunDueResult**](CampaignProbeRunDueResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## UpdateCampaignProbe

> CampaignProbeDto UpdateCampaignProbe(ctx, probeId, updateCampaignProbeOptions)

Update campaign probe

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**probeId** | [**string**]()|  | 
**updateCampaignProbeOptions** | [**UpdateCampaignProbeOptions**](UpdateCampaignProbeOptions)|  | 

### Return type

[**CampaignProbeDto**](CampaignProbeDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

