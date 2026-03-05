# MailSlurp\TemplateControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTemplate**](TemplateControllerApi#CreateTemplate) | **Post** /templates | Create a Template
[**DeleteTemplate**](TemplateControllerApi#DeleteTemplate) | **Delete** /templates/{templateId} | Delete email template
[**GetAllTemplates**](TemplateControllerApi#GetAllTemplates) | **Get** /templates/paginated | List templates
[**GetTemplate**](TemplateControllerApi#GetTemplate) | **Get** /templates/{templateId} | Get template
[**GetTemplatePreviewHtml**](TemplateControllerApi#GetTemplatePreviewHtml) | **Get** /templates/{templateId}/preview/html | Get template preview HTML
[**GetTemplatePreviewJson**](TemplateControllerApi#GetTemplatePreviewJson) | **Get** /templates/{templateId}/preview/json | Get template preview Json
[**GetTemplates**](TemplateControllerApi#GetTemplates) | **Get** /templates | List templates
[**UpdateTemplate**](TemplateControllerApi#UpdateTemplate) | **Put** /templates/{templateId} | Update template



## CreateTemplate

> TemplateDto CreateTemplate(ctx, createTemplateOptions)

Create a Template

Create an email template with variables for use with templated transactional emails.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createTemplateOptions** | [**CreateTemplateOptions**](CreateTemplateOptions)|  | 

### Return type

[**TemplateDto**](TemplateDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DeleteTemplate

> DeleteTemplate(ctx, templateId)

Delete email template

Delete template

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | [**string**]()| Template ID | 

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


## GetAllTemplates

> PageTemplateProjection GetAllTemplates(ctx, optional)

List templates

Get all templates in paginated format

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllTemplatesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Optional page index in list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **since** | **optional.Time**| Filter by created at after the given timestamp | 
 **before** | **optional.Time**| Filter by created at before the given timestamp | 

### Return type

[**PageTemplateProjection**](PageTemplateProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetTemplate

> TemplateDto GetTemplate(ctx, templateId)

Get template

Get email template

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | [**string**]()| Template ID | 

### Return type

[**TemplateDto**](TemplateDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetTemplatePreviewHtml

> string GetTemplatePreviewHtml(ctx, templateId)

Get template preview HTML

Get email template preview with passed template variables in HTML format for browsers. Pass template variables as query params.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | [**string**]()| Template ID | 

### Return type

**string**

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html;charset=utf-8, text/html

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetTemplatePreviewJson

> TemplatePreview GetTemplatePreviewJson(ctx, templateId)

Get template preview Json

Get email template preview with passed template variables in JSON format. Pass template variables as query params.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | [**string**]()| Template ID | 

### Return type

[**TemplatePreview**](TemplatePreview)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetTemplates

> []TemplateProjection GetTemplates(ctx, )

List templates

Get all templates

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]TemplateProjection**](TemplateProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## UpdateTemplate

> TemplateDto UpdateTemplate(ctx, templateId, createTemplateOptions)

Update template

Update email template

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | [**string**]()| Template ID | 
**createTemplateOptions** | [**CreateTemplateOptions**](CreateTemplateOptions)|  | 

### Return type

[**TemplateDto**](TemplateDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

