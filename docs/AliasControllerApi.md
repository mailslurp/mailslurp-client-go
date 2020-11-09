# MailSlurp\AliasControllerApi

All URIs are relative to *https://api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAlias**](AliasControllerApi.md#CreateAlias) | **Post** /aliases | Create an email alias
[**CreateAnonymousAlias**](AliasControllerApi.md#CreateAnonymousAlias) | **Post** /aliases/anonymous | Create an anonymous email alias
[**DeleteAlias**](AliasControllerApi.md#DeleteAlias) | **Delete** /aliases/{aliasId} | Delete an owned alias
[**GetAlias**](AliasControllerApi.md#GetAlias) | **Get** /aliases/{aliasId} | Get an email alias
[**GetAliases**](AliasControllerApi.md#GetAliases) | **Get** /aliases | Get all email aliases
[**UpdateAlias**](AliasControllerApi.md#UpdateAlias) | **Put** /aliases/{aliasId} | Update an owned alias



## CreateAlias

> CreateAlias(ctx, createOwnedAliasOptions)

Create an email alias

Create an email alias belonging to a user ID. To create anonymous aliases use the `createAnonymousAlias` method.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createOwnedAliasOptions** | [**CreateOwnedAliasOptions**](CreateOwnedAliasOptions.md)| createOwnedAliasOptions | 

### Return type

 (empty response body)

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAnonymousAlias

> Alias CreateAnonymousAlias(ctx, createAnonymousAliasOptions)

Create an anonymous email alias

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createAnonymousAliasOptions** | [**CreateAnonymousAliasOptions**](CreateAnonymousAliasOptions.md)| createAnonymousAliasOptions | 

### Return type

[**Alias**](Alias.md)

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlias

> DeleteAlias(ctx, aliasId)

Delete an owned alias

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aliasId** | [**string**](.md)| aliasId | 

### Return type

 (empty response body)

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlias

> Alias GetAlias(ctx, aliasId)

Get an email alias

Get an email alias by ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aliasId** | [**string**](.md)| aliasId | 

### Return type

[**Alias**](Alias.md)

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAliases

> PageAlias GetAliases(ctx, optional)

Get all email aliases

Get all email aliases in paginated form

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAliasesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAliasesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Optional page index in alias list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in alias list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]

### Return type

[**PageAlias**](Page«Alias».md)

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAlias

> UpdateAlias(ctx, aliasId, createOwnedAliasOptions)

Update an owned alias

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aliasId** | [**string**](.md)| aliasId | 
**createOwnedAliasOptions** | [**CreateOwnedAliasOptions**](CreateOwnedAliasOptions.md)| createOwnedAliasOptions | 

### Return type

 (empty response body)

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

