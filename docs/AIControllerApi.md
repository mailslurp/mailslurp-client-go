# MailSlurp\AIControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTransformer**](AIControllerApi#CreateTransformer) | **Post** /ai/transformer | Create a transformer for reuse in automations
[**CreateTransformerMappings**](AIControllerApi#CreateTransformerMappings) | **Post** /ai/transformer/mappings | Create transformer mapping
[**DeleteAllTransformerMappings**](AIControllerApi#DeleteAllTransformerMappings) | **Delete** /ai/transformer/mappings | Delete all transformer mapping
[**DeleteTransformer**](AIControllerApi#DeleteTransformer) | **Delete** /ai/transformer/{id} | Delete a transformer
[**DeleteTransformerMapping**](AIControllerApi#DeleteTransformerMapping) | **Delete** /ai/transformer/mappings/{id} | Delete transformer mapping
[**DeleteTransformers**](AIControllerApi#DeleteTransformers) | **Delete** /ai/transformer | Delete all transformers
[**GenerateStructuredContentFromAttachment**](AIControllerApi#GenerateStructuredContentFromAttachment) | **Post** /ai/structured-content/attachment | Generate structured content for an attachment
[**GenerateStructuredContentFromEmail**](AIControllerApi#GenerateStructuredContentFromEmail) | **Post** /ai/structured-content/email | Generate structured content for an email
[**GenerateStructuredContentFromSms**](AIControllerApi#GenerateStructuredContentFromSms) | **Post** /ai/structured-content/sms | Generate structured content for a TXT message
[**GetTransformer**](AIControllerApi#GetTransformer) | **Get** /ai/transformer/{id} | Get a transformer
[**GetTransformerMapping**](AIControllerApi#GetTransformerMapping) | **Get** /ai/transformer/mappings/{id} | Get transformer mapping
[**GetTransformerMappings**](AIControllerApi#GetTransformerMappings) | **Get** /ai/transformer/mappings | Get transformer mappings
[**GetTransformerResult**](AIControllerApi#GetTransformerResult) | **Get** /ai/transformer/results/{id} | Get transformer result
[**GetTransformerResults**](AIControllerApi#GetTransformerResults) | **Get** /ai/transformer/results | Get transformer results
[**GetTransformers**](AIControllerApi#GetTransformers) | **Get** /ai/transformer | List transformers
[**InvokeTransformer**](AIControllerApi#InvokeTransformer) | **Post** /ai/transformer/invoke | Invoke a transformer
[**ValidateStructuredOutputSchema**](AIControllerApi#ValidateStructuredOutputSchema) | **Post** /ai/structured-content/validate | Validate structured content schema



## CreateTransformer

> AiTransformDto CreateTransformer(ctx, aiTransformCreateOptions)

Create a transformer for reuse in automations

Save an AI transform instructions and schema for use with webhooks and automations

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aiTransformCreateOptions** | [**AiTransformCreateOptions**](AiTransformCreateOptions)|  | 

### Return type

[**AiTransformDto**](AITransformDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CreateTransformerMappings

> AiTransformMappingDto CreateTransformerMappings(ctx, createAiTransformerMappingOptions)

Create transformer mapping

Create AI transformer mappings to other entities

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createAiTransformerMappingOptions** | [**CreateAiTransformerMappingOptions**](CreateAiTransformerMappingOptions)|  | 

### Return type

[**AiTransformMappingDto**](AITransformMappingDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DeleteAllTransformerMappings

> DeleteAllTransformerMappings(ctx, )

Delete all transformer mapping

Delete all AI transformer mappings

### Required Parameters

This endpoint does not need any parameter.

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


## DeleteTransformer

> DeleteTransformer(ctx, id)

Delete a transformer

Delete an AI transformer and schemas by ID

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


## DeleteTransformerMapping

> DeleteTransformerMapping(ctx, id)

Delete transformer mapping

Delete an AI transformer mapping

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()| ID of transform mapping | 

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


## DeleteTransformers

> DeleteTransformers(ctx, )

Delete all transformers

Delete all AI transformers and schemas

### Required Parameters

This endpoint does not need any parameter.

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


## GenerateStructuredContentFromAttachment

> StructuredContentResultDto GenerateStructuredContentFromAttachment(ctx, generateStructuredContentAttachmentOptions)

Generate structured content for an attachment

Use output schemas to extract data from an attachment using AI

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**generateStructuredContentAttachmentOptions** | [**GenerateStructuredContentAttachmentOptions**](GenerateStructuredContentAttachmentOptions)|  | 

### Return type

[**StructuredContentResultDto**](StructuredContentResultDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GenerateStructuredContentFromEmail

> StructuredContentResultDto GenerateStructuredContentFromEmail(ctx, generateStructuredContentEmailOptions)

Generate structured content for an email

Use output schemas to extract data from an email using AI

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**generateStructuredContentEmailOptions** | [**GenerateStructuredContentEmailOptions**](GenerateStructuredContentEmailOptions)|  | 

### Return type

[**StructuredContentResultDto**](StructuredContentResultDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GenerateStructuredContentFromSms

> StructuredContentResultDto GenerateStructuredContentFromSms(ctx, generateStructuredContentSmsOptions)

Generate structured content for a TXT message

Use output schemas to extract data from an SMS using AI

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**generateStructuredContentSmsOptions** | [**GenerateStructuredContentSmsOptions**](GenerateStructuredContentSmsOptions)|  | 

### Return type

[**StructuredContentResultDto**](StructuredContentResultDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetTransformer

> AiTransformDto GetTransformer(ctx, id)

Get a transformer

Get AI transformer and schemas by ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()|  | 

### Return type

[**AiTransformDto**](AITransformDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetTransformerMapping

> AiTransformMappingDto GetTransformerMapping(ctx, id)

Get transformer mapping

Get an AI transformer mapping

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()| ID of transform mapping | 

### Return type

[**AiTransformMappingDto**](AITransformMappingDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetTransformerMappings

> PageAiTransformMappingProjection GetTransformerMappings(ctx, optional)

Get transformer mappings

Get AI transformer mappings to other entities

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetTransformerMappingsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTransformerMappingsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiTransformId** | [**optional.Interface of string**]()|  | 
 **entityId** | [**optional.Interface of string**]()|  | 
 **entityType** | **optional.String**|  | 
 **page** | **optional.Int32**|  | [default to 0]
 **size** | **optional.Int32**| Optional page size. Maximum size is 100. Use page index and sort to page through larger results | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]

### Return type

[**PageAiTransformMappingProjection**](PageAITransformMappingProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetTransformerResult

> AiTransformResultDto GetTransformerResult(ctx, id)

Get transformer result

Get AI transformer result

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()| ID of transform result | 

### Return type

[**AiTransformResultDto**](AITransformResultDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetTransformerResults

> PageAiTransformResultProjection GetTransformerResults(ctx, optional)

Get transformer results

Get AI transformer results

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetTransformerResultsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTransformerResultsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiTransformId** | [**optional.Interface of string**]()|  | 
 **aiTransformMappingId** | [**optional.Interface of string**]()|  | 
 **entityId** | [**optional.Interface of string**]()|  | 
 **entityType** | **optional.String**|  | 
 **page** | **optional.Int32**|  | [default to 0]
 **size** | **optional.Int32**| Optional page size. Maximum size is 100. Use page index and sort to page through larger results | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]

### Return type

[**PageAiTransformResultProjection**](PageAITransformResultProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetTransformers

> PageAiTransformProjection GetTransformers(ctx, optional)

List transformers

List all AI transforms

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetTransformersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTransformersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**|  | [default to 0]
 **size** | **optional.Int32**| Optional page size in SMS list pagination. Maximum size is 100. Use page index and sort to page through larger results | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **include** | [**optional.Interface of []string**](string)| Optional list of IDs to include in result | 

### Return type

[**PageAiTransformProjection**](PageAITransformProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## InvokeTransformer

> ConditionalStructuredContentResult InvokeTransformer(ctx, invokeTransformerOptions)

Invoke a transformer

Execute an AI transformer to generate structured content

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invokeTransformerOptions** | [**InvokeTransformerOptions**](InvokeTransformerOptions)|  | 

### Return type

[**ConditionalStructuredContentResult**](ConditionalStructuredContentResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## ValidateStructuredOutputSchema

> StructuredOutputSchemaValidation ValidateStructuredOutputSchema(ctx, structuredOutputSchema)

Validate structured content schema

Check if a schema is valid and can be used to extract data using AI

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**structuredOutputSchema** | [**StructuredOutputSchema**](StructuredOutputSchema)|  | 

### Return type

[**StructuredOutputSchemaValidation**](StructuredOutputSchemaValidation)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

