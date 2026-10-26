# MailSlurp\MissedSmsControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllMissedSmsMessages**](MissedSmsControllerApi#GetAllMissedSmsMessages) | **Get** /missed-sms | Get all missed SMS messages in paginated format
[**GetMissedSmsCount**](MissedSmsControllerApi#GetMissedSmsCount) | **Get** /missed-sms/count | Get missed SMS count
[**GetMissedSmsMessage**](MissedSmsControllerApi#GetMissedSmsMessage) | **Get** /missed-sms/{missedSmsId} | Get missed SMS content



## GetAllMissedSmsMessages

> PageMissedSmsProjection GetAllMissedSmsMessages(ctx, optional)

Get all missed SMS messages in paginated format

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllMissedSmsMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllMissedSmsMessagesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phoneNumber** | [**optional.Interface of string**]()| Optional receiving phone number to filter missed SMS for | 
 **page** | **optional.Int32**| Optional page index in list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **since** | **optional.Time**| Optional filter missed SMS after given date time | 
 **before** | **optional.Time**| Optional filter missed SMS before given date time | 
 **search** | **optional.String**| Optional search filter | 

### Return type

[**PageMissedSmsProjection**](PageMissedSmsProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetMissedSmsCount

> CountDto GetMissedSmsCount(ctx, )

Get missed SMS count

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**CountDto**](CountDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetMissedSmsMessage

> MissedSmsDto GetMissedSmsMessage(ctx, missedSmsId)

Get missed SMS content

Returns a missed SMS with full content.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**missedSmsId** | [**string**]()|  | 

### Return type

[**MissedSmsDto**](MissedSmsDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

