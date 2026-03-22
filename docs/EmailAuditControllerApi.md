# MailSlurp\EmailAuditControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompareEmailAudits**](EmailAuditControllerApi#CompareEmailAudits) | **Get** /email-audits/{auditId}/compare/{otherAuditId} | Compare two email audits
[**CreateEmailAudit**](EmailAuditControllerApi#CreateEmailAudit) | **Post** /email-audits | Create email audit
[**DeleteEmailAudit**](EmailAuditControllerApi#DeleteEmailAudit) | **Delete** /email-audits/{auditId} | Delete email audit
[**GetEmailAudit**](EmailAuditControllerApi#GetEmailAudit) | **Get** /email-audits/{auditId} | Get email audit
[**GetEmailAudits**](EmailAuditControllerApi#GetEmailAudits) | **Get** /email-audits | List email audits



## CompareEmailAudits

> EmailAuditComparisonDto CompareEmailAudits(ctx, auditId, otherAuditId)

Compare two email audits

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**auditId** | [**string**]()|  | 
**otherAuditId** | [**string**]()|  | 

### Return type

[**EmailAuditComparisonDto**](EmailAuditComparisonDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CreateEmailAudit

> EmailAuditDto CreateEmailAudit(ctx, createEmailAuditOptions)

Create email audit

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createEmailAuditOptions** | [**CreateEmailAuditOptions**](CreateEmailAuditOptions)|  | 

### Return type

[**EmailAuditDto**](EmailAuditDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DeleteEmailAudit

> DeleteEmailAudit(ctx, auditId)

Delete email audit

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**auditId** | [**string**]()|  | 

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


## GetEmailAudit

> EmailAuditDto GetEmailAudit(ctx, auditId)

Get email audit

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**auditId** | [**string**]()|  | 

### Return type

[**EmailAuditDto**](EmailAuditDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetEmailAudits

> []EmailAuditDto GetEmailAudits(ctx, optional)

List email audits

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetEmailAuditsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEmailAuditsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailId** | [**optional.Interface of string**]()|  | 
 **since** | **optional.Time**|  | 
 **before** | **optional.Time**|  | 
 **limit** | **optional.Int32**|  | 

### Return type

[**[]EmailAuditDto**](EmailAuditDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

