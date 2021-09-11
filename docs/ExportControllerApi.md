# MailSlurp\ExportControllerApi

All URIs are relative to *https://api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportEntities**](ExportControllerApi#ExportEntities) | **Get** /export | Export inboxes link callable via browser
[**GetExportLink**](ExportControllerApi#GetExportLink) | **Post** /export | Get export link



## ExportEntities

> string ExportEntities(ctx, apiKey, exportType, outputFormat, optional)

Export inboxes link callable via browser

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiKey** | **string**| apiKey | 
**exportType** | **string**| exportType | 
**outputFormat** | **string**| outputFormat | 
 **optional** | ***ExportEntitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExportEntitiesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **createdEarliestTime** | **optional.Time**| createdEarliestTime | 
 **createdOldestTime** | **optional.Time**| createdOldestTime | 
 **excludePreviouslyExported** | **optional.Bool**| excludePreviouslyExported | 
 **filter** | **optional.String**| filter | 
 **listSeparatorToken** | **optional.String**| listSeparatorToken | 

### Return type

**string**

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
**exportType** | **string**| exportType | 
**exportOptions** | [**ExportOptions**](ExportOptions)| exportOptions | 
 **optional** | ***GetExportLinkOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetExportLinkOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiKey** | **optional.String**| apiKey | 

### Return type

[**ExportLink**](ExportLink)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

