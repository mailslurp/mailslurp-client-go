# MailSlurp\ApiFreeInboxControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](ApiFreeInboxControllerApi#Create) | **Post** /tools/free-inbox | 
[**Delete**](ApiFreeInboxControllerApi#Delete) | **Delete** /tools/free-inbox/{id} | 
[**Read**](ApiFreeInboxControllerApi#Read) | **Get** /tools/free-inbox/{id} | 



## Create

> CreatedFreeInboxDto Create(ctx, createFreeInboxOptions)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createFreeInboxOptions** | [**CreateFreeInboxOptions**](CreateFreeInboxOptions)|  | 

### Return type

[**CreatedFreeInboxDto**](CreatedFreeInboxDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## Delete

> Delete(ctx, id, xFreeInboxToken)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**xFreeInboxToken** | **string**|  | 

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


## Read

> FreeInboxDto Read(ctx, id, xFreeInboxToken)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**xFreeInboxToken** | **string**|  | 

### Return type

[**FreeInboxDto**](FreeInboxDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

