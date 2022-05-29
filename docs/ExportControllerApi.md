# MailSlurp\ExportControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportEntities**](ExportControllerApi#ExportEntities) | **Get** /export | Export inboxes link callable via browser
[**GetExportLink**](ExportControllerApi#GetExportLink) | **Post** /export | Get export link



## ExportEntities

> []string ExportEntities(ctx, exportType, apiKey, outputFormat, optional)

Export inboxes link callable via browser

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exportType** | **string**|  | 
**apiKey** | **string**|  | 
**outputFormat** | **string**|  | 
 **optional** | ***ExportEntitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExportEntitiesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filter** | **optional.String**|  | 
 **listSeparatorToken** | **optional.String**|  | 
 **excludePreviouslyExported** | **optional.Bool**|  | 
 **createdEarliestTime** | **optional.Time**|  | 
 **createdOldestTime** | **optional.Time**|  | 

### Return type

**[]string**

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetExportLink

> ExportLink GetExportLink(ctx, exportType, exportOptions, optional)

Get export link

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exportType** | **string**|  | 
**exportOptions** | [**ExportOptions**](ExportOptions)|  | 
 **optional** | ***GetExportLinkOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetExportLinkOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiKey** | **optional.String**|  | 

### Return type

[**ExportLink**](ExportLink)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

