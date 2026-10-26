# MailSlurp\AuthV2Api

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthV2AccountSwitch**](AuthV2Api#AuthV2AccountSwitch) | **Post** /auth/v2/account/switch | Switch the account bound to an Auth v2 session
[**AuthV2Bootstrap**](AuthV2Api#AuthV2Bootstrap) | **Post** /auth/v2/bootstrap | Bootstrap an existing Auth v2 session
[**AuthV2EnvironmentSwitch**](AuthV2Api#AuthV2EnvironmentSwitch) | **Post** /auth/v2/environment/{targetUserId}/switch | Switch the environment bound to an Auth v2 session
[**AuthV2Logout**](AuthV2Api#AuthV2Logout) | **Post** /auth/v2/logout | Revoke an Auth v2 session
[**AuthV2PasswordLogin**](AuthV2Api#AuthV2PasswordLogin) | **Post** /auth/v2/password/login | Create an Auth v2 session using a password
[**AuthV2ProviderCallback**](AuthV2Api#AuthV2ProviderCallback) | **Get** /auth/v2/callback/{providerKind} | Redeem a staging provider handoff and continue to the app
[**EmailCallback**](AuthV2Api#EmailCallback) | **Get** /auth/v2/callback/email | 
[**EmailCallbackComplete**](AuthV2Api#EmailCallbackComplete) | **Post** /auth/v2/callback/email/complete | 
[**EmailCallbackContinue**](AuthV2Api#EmailCallbackContinue) | **Get** /auth/v2/callback/email/continue | 
[**EmailCallbackContinueScript**](AuthV2Api#EmailCallbackContinueScript) | **Get** /auth/v2/callback/email/continue.js | 
[**EmailCallbackHead**](AuthV2Api#EmailCallbackHead) | **Head** /auth/v2/callback/email | 



## AuthV2AccountSwitch

> AuthV2BootstrapDto AuthV2AccountSwitch(ctx, switchAccountOptions)

Switch the account bound to an Auth v2 session

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**switchAccountOptions** | [**SwitchAccountOptions**](SwitchAccountOptions)|  | 

### Return type

[**AuthV2BootstrapDto**](AuthV2BootstrapDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AuthV2Bootstrap

> AuthV2BootstrapDto AuthV2Bootstrap(ctx, optional)

Bootstrap an existing Auth v2 session

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthV2BootstrapOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AuthV2BootstrapOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authV2BootstrapOptions** | [**optional.Interface of AuthV2BootstrapOptions**](AuthV2BootstrapOptions)|  | 

### Return type

[**AuthV2BootstrapDto**](AuthV2BootstrapDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AuthV2EnvironmentSwitch

> AuthV2BootstrapDto AuthV2EnvironmentSwitch(ctx, targetUserId)

Switch the environment bound to an Auth v2 session

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetUserId** | [**string**]()|  | 

### Return type

[**AuthV2BootstrapDto**](AuthV2BootstrapDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AuthV2Logout

> AuthV2Logout(ctx, )

Revoke an Auth v2 session

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


## AuthV2PasswordLogin

> AuthV2BootstrapDto AuthV2PasswordLogin(ctx, authV2PasswordLoginOptions)

Create an Auth v2 session using a password

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authV2PasswordLoginOptions** | [**AuthV2PasswordLoginOptions**](AuthV2PasswordLoginOptions)|  | 

### Return type

[**AuthV2BootstrapDto**](AuthV2BootstrapDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AuthV2ProviderCallback

> AuthV2ProviderCallback(ctx, providerKind, optional)

Redeem a staging provider handoff and continue to the app

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerKind** | **string**|  | 
 **optional** | ***AuthV2ProviderCallbackOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AuthV2ProviderCallbackOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **code** | **optional.String**|  | 

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


## EmailCallback

> EmailCallback(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EmailCallbackOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EmailCallbackOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **optional.String**|  | 

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


## EmailCallbackComplete

> EmailCallbackComplete(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EmailCallbackCompleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EmailCallbackCompleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attempt** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## EmailCallbackContinue

> string EmailCallbackContinue(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EmailCallbackContinueOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EmailCallbackContinueOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attempt** | **optional.String**|  | 

### Return type

**string**

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## EmailCallbackContinueScript

> string EmailCallbackContinueScript(ctx, )



### Required Parameters

This endpoint does not need any parameter.

### Return type

**string**

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/javascript;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## EmailCallbackHead

> EmailCallbackHead(ctx, )



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

