# MailSlurp\ApiInternalControllerApi

All URIs are relative to *https://api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSamlUserOrCreate**](ApiInternalControllerApi#GetSamlUserOrCreate) | **Post** /internal/saml/user | 



## GetSamlUserOrCreate

> UserDto GetSamlUserOrCreate(ctx, key, getOrCreateSamlUserOptions)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**|  | 
**getOrCreateSamlUserOptions** | [**GetOrCreateSamlUserOptions**](GetOrCreateSamlUserOptions)|  | 

### Return type

[**UserDto**](UserDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

