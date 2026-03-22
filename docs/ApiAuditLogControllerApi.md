# MailSlurp\ApiAuditLogControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuditLogByEventId**](ApiAuditLogControllerApi#GetAuditLogByEventId) | **Get** /audit-logs/{eventId} | 
[**GetAuditLogs**](ApiAuditLogControllerApi#GetAuditLogs) | **Get** /audit-logs | 
[**SearchAuditLogs**](ApiAuditLogControllerApi#SearchAuditLogs) | **Post** /audit-logs/search | 



## GetAuditLogByEventId

> AuditLogEventDto GetAuditLogByEventId(ctx, eventId, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string**|  | 
 **optional** | ***GetAuditLogByEventIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAuditLogByEventIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **since** | **optional.Time**|  | 
 **before** | **optional.Time**|  | 

### Return type

[**AuditLogEventDto**](AuditLogEventDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetAuditLogs

> AuditLogPageDto GetAuditLogs(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAuditLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAuditLogsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **since** | **optional.Time**|  | 
 **before** | **optional.Time**|  | 
 **action** | **optional.String**|  | 
 **userId** | [**optional.Interface of string**]()|  | 
 **actorUserId** | [**optional.Interface of string**]()|  | 
 **targetUserId** | [**optional.Interface of string**]()|  | 
 **resourceType** | **optional.String**|  | 
 **resourceId** | **optional.String**|  | 
 **outcome** | **optional.String**|  | 
 **requestId** | **optional.String**|  | 
 **ipAddress** | **optional.String**|  | 
 **pageSize** | **optional.Int32**|  | 
 **cursor** | **optional.String**|  | 

### Return type

[**AuditLogPageDto**](AuditLogPageDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## SearchAuditLogs

> AuditLogPageDto SearchAuditLogs(ctx, auditLogSearchOptions)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**auditLogSearchOptions** | [**AuditLogSearchOptions**](AuditLogSearchOptions)|  | 

### Return type

[**AuditLogPageDto**](AuditLogPageDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

