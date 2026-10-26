# MailSlurp\ApiEmailWarmupControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEmailWarmupProgram**](ApiEmailWarmupControllerApi#CreateEmailWarmupProgram) | **Post** /email-warmup/programs | Create an email warmup program
[**DeleteEmailWarmupProgram**](ApiEmailWarmupControllerApi#DeleteEmailWarmupProgram) | **Delete** /email-warmup/programs/{programId} | Delete an email warmup program
[**DeleteEmailWarmupSenderSeats**](ApiEmailWarmupControllerApi#DeleteEmailWarmupSenderSeats) | **Delete** /email-warmup/programs/sender-seats | Cancel purchased email warmup sender seats
[**GetEmailWarmupOverview**](ApiEmailWarmupControllerApi#GetEmailWarmupOverview) | **Get** /email-warmup/programs/overview | Get operational email warmup activity and program progress
[**GetEmailWarmupPlacementChecks**](ApiEmailWarmupControllerApi#GetEmailWarmupPlacementChecks) | **Get** /email-warmup/programs/{programId}/placement-checks | List Gmail and Outlook placement checks for an email warmup program
[**GetEmailWarmupProgram**](ApiEmailWarmupControllerApi#GetEmailWarmupProgram) | **Get** /email-warmup/programs/{programId} | Get an email warmup program
[**GetEmailWarmupProgramAnalytics**](ApiEmailWarmupControllerApi#GetEmailWarmupProgramAnalytics) | **Get** /email-warmup/programs/{programId}/analytics | Get email warmup program progress and daily activity
[**GetEmailWarmupPrograms**](ApiEmailWarmupControllerApi#GetEmailWarmupPrograms) | **Get** /email-warmup/programs | List email warmup programs
[**GetEmailWarmupReadiness**](ApiEmailWarmupControllerApi#GetEmailWarmupReadiness) | **Get** /email-warmup/programs/{programId}/readiness | Get warmup readiness, adaptive pacing, and sender-health signals
[**GetEmailWarmupSlots**](ApiEmailWarmupControllerApi#GetEmailWarmupSlots) | **Get** /email-warmup/programs/{programId}/slots | List scheduled email warmup slots
[**GetEmailWarmupUsage**](ApiEmailWarmupControllerApi#GetEmailWarmupUsage) | **Get** /email-warmup/programs/usage | Get email warmup sender capacity, included credits, and billable usage
[**PauseEmailWarmupProgram**](ApiEmailWarmupControllerApi#PauseEmailWarmupProgram) | **Post** /email-warmup/programs/{programId}/pause | Pause an email warmup program
[**ResumeEmailWarmupProgram**](ApiEmailWarmupControllerApi#ResumeEmailWarmupProgram) | **Post** /email-warmup/programs/{programId}/resume | Resume an email warmup program
[**RetryEmailWarmupProgram**](ApiEmailWarmupControllerApi#RetryEmailWarmupProgram) | **Post** /email-warmup/programs/{programId}/retry | Retry a system-paused email warmup program
[**StartEmailWarmupProgram**](ApiEmailWarmupControllerApi#StartEmailWarmupProgram) | **Post** /email-warmup/programs/{programId}/start | Start an email warmup program
[**StopEmailWarmupProgram**](ApiEmailWarmupControllerApi#StopEmailWarmupProgram) | **Post** /email-warmup/programs/{programId}/stop | Stop an email warmup program
[**UpdateEmailWarmupProgram**](ApiEmailWarmupControllerApi#UpdateEmailWarmupProgram) | **Put** /email-warmup/programs/{programId} | Update a draft email warmup program
[**UpdateEmailWarmupSenderSeats**](ApiEmailWarmupControllerApi#UpdateEmailWarmupSenderSeats) | **Put** /email-warmup/programs/sender-seats | Update purchased email warmup sender seats



## CreateEmailWarmupProgram

> EmailWarmupProgramDto CreateEmailWarmupProgram(ctx, createEmailWarmupProgramOptions)

Create an email warmup program

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createEmailWarmupProgramOptions** | [**CreateEmailWarmupProgramOptions**](CreateEmailWarmupProgramOptions)|  | 

### Return type

[**EmailWarmupProgramDto**](EmailWarmupProgramDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DeleteEmailWarmupProgram

> DeleteEmailWarmupProgram(ctx, programId, emailWarmupProgramCommandOptions)

Delete an email warmup program

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**programId** | [**string**]()|  | 
**emailWarmupProgramCommandOptions** | [**EmailWarmupProgramCommandOptions**](EmailWarmupProgramCommandOptions)|  | 

### Return type

 (empty response body)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DeleteEmailWarmupSenderSeats

> EmailWarmupUsageSummaryDto DeleteEmailWarmupSenderSeats(ctx, )

Cancel purchased email warmup sender seats

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**EmailWarmupUsageSummaryDto**](EmailWarmupUsageSummaryDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetEmailWarmupOverview

> EmailWarmupOverviewDto GetEmailWarmupOverview(ctx, optional)

Get operational email warmup activity and program progress

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetEmailWarmupOverviewOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEmailWarmupOverviewOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **days** | **optional.Int32**|  | [default to 30]

### Return type

[**EmailWarmupOverviewDto**](EmailWarmupOverviewDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetEmailWarmupPlacementChecks

> []EmailWarmupPlacementCheckDto GetEmailWarmupPlacementChecks(ctx, programId)

List Gmail and Outlook placement checks for an email warmup program

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**programId** | [**string**]()|  | 

### Return type

[**[]EmailWarmupPlacementCheckDto**](EmailWarmupPlacementCheckDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetEmailWarmupProgram

> EmailWarmupProgramDto GetEmailWarmupProgram(ctx, programId)

Get an email warmup program

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**programId** | [**string**]()|  | 

### Return type

[**EmailWarmupProgramDto**](EmailWarmupProgramDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetEmailWarmupProgramAnalytics

> EmailWarmupProgramAnalyticsDto GetEmailWarmupProgramAnalytics(ctx, programId, optional)

Get email warmup program progress and daily activity

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**programId** | [**string**]()|  | 
 **optional** | ***GetEmailWarmupProgramAnalyticsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEmailWarmupProgramAnalyticsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **days** | **optional.Int32**|  | [default to 30]

### Return type

[**EmailWarmupProgramAnalyticsDto**](EmailWarmupProgramAnalyticsDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetEmailWarmupPrograms

> EmailWarmupProgramPageDto GetEmailWarmupPrograms(ctx, optional)

List email warmup programs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetEmailWarmupProgramsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEmailWarmupProgramsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**|  | [default to 0]
 **size** | **optional.Int32**|  | [default to 20]

### Return type

[**EmailWarmupProgramPageDto**](EmailWarmupProgramPageDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetEmailWarmupReadiness

> EmailWarmupReadinessDto GetEmailWarmupReadiness(ctx, programId)

Get warmup readiness, adaptive pacing, and sender-health signals

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**programId** | [**string**]()|  | 

### Return type

[**EmailWarmupReadinessDto**](EmailWarmupReadinessDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetEmailWarmupSlots

> EmailWarmupSendSlotPageDto GetEmailWarmupSlots(ctx, programId, optional)

List scheduled email warmup slots

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**programId** | [**string**]()|  | 
 **optional** | ***GetEmailWarmupSlotsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEmailWarmupSlotsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**|  | [default to 0]
 **size** | **optional.Int32**|  | [default to 50]

### Return type

[**EmailWarmupSendSlotPageDto**](EmailWarmupSendSlotPageDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetEmailWarmupUsage

> EmailWarmupUsageSummaryDto GetEmailWarmupUsage(ctx, )

Get email warmup sender capacity, included credits, and billable usage

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**EmailWarmupUsageSummaryDto**](EmailWarmupUsageSummaryDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## PauseEmailWarmupProgram

> EmailWarmupProgramDto PauseEmailWarmupProgram(ctx, programId, emailWarmupProgramCommandOptions)

Pause an email warmup program

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**programId** | [**string**]()|  | 
**emailWarmupProgramCommandOptions** | [**EmailWarmupProgramCommandOptions**](EmailWarmupProgramCommandOptions)|  | 

### Return type

[**EmailWarmupProgramDto**](EmailWarmupProgramDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## ResumeEmailWarmupProgram

> EmailWarmupProgramDto ResumeEmailWarmupProgram(ctx, programId, emailWarmupProgramCommandOptions)

Resume an email warmup program

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**programId** | [**string**]()|  | 
**emailWarmupProgramCommandOptions** | [**EmailWarmupProgramCommandOptions**](EmailWarmupProgramCommandOptions)|  | 

### Return type

[**EmailWarmupProgramDto**](EmailWarmupProgramDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## RetryEmailWarmupProgram

> EmailWarmupProgramDto RetryEmailWarmupProgram(ctx, programId, emailWarmupProgramCommandOptions)

Retry a system-paused email warmup program

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**programId** | [**string**]()|  | 
**emailWarmupProgramCommandOptions** | [**EmailWarmupProgramCommandOptions**](EmailWarmupProgramCommandOptions)|  | 

### Return type

[**EmailWarmupProgramDto**](EmailWarmupProgramDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## StartEmailWarmupProgram

> EmailWarmupProgramDto StartEmailWarmupProgram(ctx, programId, emailWarmupProgramCommandOptions)

Start an email warmup program

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**programId** | [**string**]()|  | 
**emailWarmupProgramCommandOptions** | [**EmailWarmupProgramCommandOptions**](EmailWarmupProgramCommandOptions)|  | 

### Return type

[**EmailWarmupProgramDto**](EmailWarmupProgramDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## StopEmailWarmupProgram

> EmailWarmupProgramDto StopEmailWarmupProgram(ctx, programId, emailWarmupProgramCommandOptions)

Stop an email warmup program

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**programId** | [**string**]()|  | 
**emailWarmupProgramCommandOptions** | [**EmailWarmupProgramCommandOptions**](EmailWarmupProgramCommandOptions)|  | 

### Return type

[**EmailWarmupProgramDto**](EmailWarmupProgramDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## UpdateEmailWarmupProgram

> EmailWarmupProgramDto UpdateEmailWarmupProgram(ctx, programId, updateEmailWarmupProgramOptions)

Update a draft email warmup program

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**programId** | [**string**]()|  | 
**updateEmailWarmupProgramOptions** | [**UpdateEmailWarmupProgramOptions**](UpdateEmailWarmupProgramOptions)|  | 

### Return type

[**EmailWarmupProgramDto**](EmailWarmupProgramDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## UpdateEmailWarmupSenderSeats

> EmailWarmupUsageSummaryDto UpdateEmailWarmupSenderSeats(ctx, updateEmailWarmupSenderSeatsOptions)

Update purchased email warmup sender seats

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**updateEmailWarmupSenderSeatsOptions** | [**UpdateEmailWarmupSenderSeatsOptions**](UpdateEmailWarmupSenderSeatsOptions)|  | 

### Return type

[**EmailWarmupUsageSummaryDto**](EmailWarmupUsageSummaryDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

