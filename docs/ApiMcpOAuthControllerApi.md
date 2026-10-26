# MailSlurp\ApiMcpOAuthControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthorizationServerMetadata**](ApiMcpOAuthControllerApi#AuthorizationServerMetadata) | **Get** /.well-known/oauth-authorization-server/mcp | 
[**AuthorizationServerMetadata1**](ApiMcpOAuthControllerApi#AuthorizationServerMetadata1) | **Get** /.well-known/oauth-authorization-server | 
[**Authorize**](ApiMcpOAuthControllerApi#Authorize) | **Get** /mcp/oauth/authorize | 
[**ProtectedResourceMetadata**](ApiMcpOAuthControllerApi#ProtectedResourceMetadata) | **Get** /.well-known/oauth-protected-resource/mcp | 
[**ProtectedResourceMetadata1**](ApiMcpOAuthControllerApi#ProtectedResourceMetadata1) | **Get** /.well-known/oauth-protected-resource | 
[**Register**](ApiMcpOAuthControllerApi#Register) | **Post** /mcp/oauth/register | 
[**Token**](ApiMcpOAuthControllerApi#Token) | **Post** /mcp/oauth/token | 



## AuthorizationServerMetadata

> map[string]map[string]interface{} AuthorizationServerMetadata(ctx, )



### Required Parameters

This endpoint does not need any parameter.

### Return type

**map[string]map[string]interface{}**

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AuthorizationServerMetadata1

> map[string]map[string]interface{} AuthorizationServerMetadata1(ctx, )



### Required Parameters

This endpoint does not need any parameter.

### Return type

**map[string]map[string]interface{}**

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## Authorize

> map[string]interface{} Authorize(ctx, clientId, redirectUri, responseType, codeChallenge, codeChallengeMethod, resource, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string**|  | 
**redirectUri** | **string**|  | 
**responseType** | **string**|  | 
**codeChallenge** | **string**|  | 
**codeChallengeMethod** | **string**|  | 
**resource** | **string**|  | 
 **optional** | ***AuthorizeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AuthorizeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **scope** | **optional.String**|  | 
 **state** | **optional.String**|  | 

### Return type

**map[string]interface{}**

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## ProtectedResourceMetadata

> map[string]map[string]interface{} ProtectedResourceMetadata(ctx, )



### Required Parameters

This endpoint does not need any parameter.

### Return type

**map[string]map[string]interface{}**

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## ProtectedResourceMetadata1

> map[string]map[string]interface{} ProtectedResourceMetadata1(ctx, )



### Required Parameters

This endpoint does not need any parameter.

### Return type

**map[string]map[string]interface{}**

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## Register

> map[string]interface{} Register(ctx, mcpOAuthClientRegistrationOptions)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mcpOAuthClientRegistrationOptions** | [**McpOAuthClientRegistrationOptions**](McpOAuthClientRegistrationOptions)|  | 

### Return type

**map[string]interface{}**

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## Token

> map[string]interface{} Token(ctx, grantType, clientId, resource, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**grantType** | **string**|  | 
**clientId** | **string**|  | 
**resource** | **string**|  | 
 **optional** | ***TokenOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TokenOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **code** | **optional.String**|  | 
 **redirectUri** | **optional.String**|  | 
 **codeVerifier** | **optional.String**|  | 
 **refreshToken** | **optional.String**|  | 

### Return type

**map[string]interface{}**

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

