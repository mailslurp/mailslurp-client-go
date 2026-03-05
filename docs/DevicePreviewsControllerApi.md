# MailSlurp\DevicePreviewsControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelDevicePreviewRun**](DevicePreviewsControllerApi#CancelDevicePreviewRun) | **Post** /emails/device-previews/{runId}/cancel | Cancel a running device preview run
[**CreateDevicePreviewFeedback**](DevicePreviewsControllerApi#CreateDevicePreviewFeedback) | **Post** /emails/device-previews/feedback | Create device preview feedback
[**CreateDevicePreviewRun**](DevicePreviewsControllerApi#CreateDevicePreviewRun) | **Post** /emails/{emailId}/device-previews | Create a new device preview run for an email
[**DeleteDevicePreviewRun**](DevicePreviewsControllerApi#DeleteDevicePreviewRun) | **Delete** /emails/device-previews/{runId} | Delete local device preview run data
[**EnsureDevicePreviewRun**](DevicePreviewsControllerApi#EnsureDevicePreviewRun) | **Put** /emails/{emailId}/device-previews/latest | Return active run for email or create one when none exists
[**GetDevicePreviewFeedback**](DevicePreviewsControllerApi#GetDevicePreviewFeedback) | **Get** /emails/device-previews/feedback/{feedbackId} | Get a single device preview feedback item
[**GetDevicePreviewFeedbackItems**](DevicePreviewsControllerApi#GetDevicePreviewFeedbackItems) | **Get** /emails/device-previews/feedback | List device preview feedback
[**GetDevicePreviewRun**](DevicePreviewsControllerApi#GetDevicePreviewRun) | **Get** /emails/device-previews/{runId} | Get device preview run status
[**GetDevicePreviewRunProviderProgress**](DevicePreviewsControllerApi#GetDevicePreviewRunProviderProgress) | **Get** /emails/device-previews/{runId}/providers/{provider} | Get provider-level progress for a device preview run
[**GetDevicePreviewRunResults**](DevicePreviewsControllerApi#GetDevicePreviewRunResults) | **Get** /emails/device-previews/{runId}/results | Get device preview run results
[**GetDevicePreviewRunScreenshot**](DevicePreviewsControllerApi#GetDevicePreviewRunScreenshot) | **Get** /emails/device-previews/{runId}/screenshots/{screenshotId}/image | Get a seeded device preview screenshot image
[**GetDevicePreviewRuns**](DevicePreviewsControllerApi#GetDevicePreviewRuns) | **Get** /emails/{emailId}/device-previews | List previous device preview runs for an email
[**GetDevicePreviewRunsForAccount**](DevicePreviewsControllerApi#GetDevicePreviewRunsForAccount) | **Get** /emails/device-previews | List previous device preview runs for account
[**GetDevicePreviewRunsOffsetPaginated**](DevicePreviewsControllerApi#GetDevicePreviewRunsOffsetPaginated) | **Get** /emails/{emailId}/device-previews/offset-paginated | List previous device preview runs for an email in paginated form
[**UpdateDevicePreviewFeedback**](DevicePreviewsControllerApi#UpdateDevicePreviewFeedback) | **Put** /emails/device-previews/feedback/{feedbackId} | Update device preview feedback



## CancelDevicePreviewRun

> CancelDevicePreviewRunResult CancelDevicePreviewRun(ctx, runId, optional)

Cancel a running device preview run

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | [**string**]()|  | 
 **optional** | ***CancelDevicePreviewRunOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CancelDevicePreviewRunOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cancelDevicePreviewRunOptions** | [**optional.Interface of CancelDevicePreviewRunOptions**](CancelDevicePreviewRunOptions)|  | 

### Return type

[**CancelDevicePreviewRunResult**](CancelDevicePreviewRunResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CreateDevicePreviewFeedback

> DevicePreviewFeedbackDto CreateDevicePreviewFeedback(ctx, createDevicePreviewFeedbackOptions)

Create device preview feedback

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createDevicePreviewFeedbackOptions** | [**CreateDevicePreviewFeedbackOptions**](CreateDevicePreviewFeedbackOptions)|  | 

### Return type

[**DevicePreviewFeedbackDto**](DevicePreviewFeedbackDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CreateDevicePreviewRun

> CreateDevicePreviewRunResult CreateDevicePreviewRun(ctx, emailId, optional)

Create a new device preview run for an email

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()|  | 
 **optional** | ***CreateDevicePreviewRunOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateDevicePreviewRunOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDevicePreviewOptions** | [**optional.Interface of CreateDevicePreviewOptions**](CreateDevicePreviewOptions)|  | 

### Return type

[**CreateDevicePreviewRunResult**](CreateDevicePreviewRunResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DeleteDevicePreviewRun

> DeleteDevicePreviewRunResult DeleteDevicePreviewRun(ctx, runId)

Delete local device preview run data

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | [**string**]()|  | 

### Return type

[**DeleteDevicePreviewRunResult**](DeleteDevicePreviewRunResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## EnsureDevicePreviewRun

> CreateDevicePreviewRunResult EnsureDevicePreviewRun(ctx, emailId, optional)

Return active run for email or create one when none exists

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()|  | 
 **optional** | ***EnsureDevicePreviewRunOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EnsureDevicePreviewRunOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDevicePreviewOptions** | [**optional.Interface of CreateDevicePreviewOptions**](CreateDevicePreviewOptions)|  | 

### Return type

[**CreateDevicePreviewRunResult**](CreateDevicePreviewRunResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDevicePreviewFeedback

> DevicePreviewFeedbackDto GetDevicePreviewFeedback(ctx, feedbackId)

Get a single device preview feedback item

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**feedbackId** | [**string**]()|  | 

### Return type

[**DevicePreviewFeedbackDto**](DevicePreviewFeedbackDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDevicePreviewFeedbackItems

> DevicePreviewFeedbackListDto GetDevicePreviewFeedbackItems(ctx, optional)

List device preview feedback

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetDevicePreviewFeedbackItemsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDevicePreviewFeedbackItemsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**|  | 
 **size** | **optional.Int32**|  | 
 **source** | **optional.String**|  | 
 **runId** | [**optional.Interface of string**]()|  | 
 **status** | **optional.String**|  | 
 **provider** | **optional.String**|  | 
 **category** | **optional.String**|  | 
 **search** | **optional.String**|  | 

### Return type

[**DevicePreviewFeedbackListDto**](DevicePreviewFeedbackListDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDevicePreviewRun

> DevicePreviewRunDto GetDevicePreviewRun(ctx, runId)

Get device preview run status

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | [**string**]()|  | 

### Return type

[**DevicePreviewRunDto**](DevicePreviewRunDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDevicePreviewRunProviderProgress

> DevicePreviewProviderProgressDto GetDevicePreviewRunProviderProgress(ctx, runId, provider)

Get provider-level progress for a device preview run

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | [**string**]()|  | 
**provider** | **string**|  | 

### Return type

[**DevicePreviewProviderProgressDto**](DevicePreviewProviderProgressDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDevicePreviewRunResults

> DevicePreviewRunResultsDto GetDevicePreviewRunResults(ctx, runId)

Get device preview run results

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | [**string**]()|  | 

### Return type

[**DevicePreviewRunResultsDto**](DevicePreviewRunResultsDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDevicePreviewRunScreenshot

> string GetDevicePreviewRunScreenshot(ctx, runId, screenshotId)

Get a seeded device preview screenshot image

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | [**string**]()|  | 
**screenshotId** | [**string**]()|  | 

### Return type

**string**

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDevicePreviewRuns

> []DevicePreviewRunDto GetDevicePreviewRuns(ctx, emailId, optional)

List previous device preview runs for an email

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()|  | 
 **optional** | ***GetDevicePreviewRunsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDevicePreviewRunsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 50]

### Return type

[**[]DevicePreviewRunDto**](DevicePreviewRunDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDevicePreviewRunsForAccount

> []DevicePreviewRunDto GetDevicePreviewRunsForAccount(ctx, optional)

List previous device preview runs for account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetDevicePreviewRunsForAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDevicePreviewRunsForAccountOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**|  | [default to 50]

### Return type

[**[]DevicePreviewRunDto**](DevicePreviewRunDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDevicePreviewRunsOffsetPaginated

> PageDevicePreviewRunProjection GetDevicePreviewRunsOffsetPaginated(ctx, emailId, optional)

List previous device preview runs for an email in paginated form

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()|  | 
 **optional** | ***GetDevicePreviewRunsOffsetPaginatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDevicePreviewRunsOffsetPaginatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Optional page index in list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size for paginated result list. | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to DESC]

### Return type

[**PageDevicePreviewRunProjection**](PageDevicePreviewRunProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## UpdateDevicePreviewFeedback

> DevicePreviewFeedbackDto UpdateDevicePreviewFeedback(ctx, feedbackId, updateDevicePreviewFeedbackOptions)

Update device preview feedback

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**feedbackId** | [**string**]()|  | 
**updateDevicePreviewFeedbackOptions** | [**UpdateDevicePreviewFeedbackOptions**](UpdateDevicePreviewFeedbackOptions)|  | 

### Return type

[**DevicePreviewFeedbackDto**](DevicePreviewFeedbackDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

