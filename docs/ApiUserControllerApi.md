# MailSlurp\ApiUserControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetJsonPropertyAsString**](ApiUserControllerApi#GetJsonPropertyAsString) | **Post** /user/json/pluck | 
[**GetUserInfo**](ApiUserControllerApi#GetUserInfo) | **Get** /user/info | 



## GetJsonPropertyAsString

> string GetJsonPropertyAsString(ctx, property, body)



Utility function to extract properties from JSON objects in language where this is cumbersome.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**property** | **string**| JSON property name or dot separated path selector such as &#x60;a.b.c&#x60; | 
**body** | **map[string]interface{}**|  | 

### Return type

**string**

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetUserInfo

> UserInfoDto GetUserInfo(ctx, )



### Required Parameters

This endpoint does not need any parameter.

### Return type

[**UserInfoDto**](UserInfoDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

