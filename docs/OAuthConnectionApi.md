# MailSlurp\OAuthConnectionApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOAuthConnection**](OAuthConnectionApi#CreateOAuthConnection) | **Post** /oauth-connection | Create an OAuth connection
[**ExchangeAuthorizationTokenAndCreateOrUpdateInbox**](OAuthConnectionApi#ExchangeAuthorizationTokenAndCreateOrUpdateInbox) | **Post** /oauth-connection/oauth-exchange/google | Exchange authorization code for access token and create inbox



## CreateOAuthConnection

> CreateOAuthConnectionResult CreateOAuthConnection(ctx, redirectBase, oAuthConnectionType, optional)

Create an OAuth connection

Configure an inbox for OAuth sync with MailSlurp

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**redirectBase** | **string**|  | 
**oAuthConnectionType** | **string**|  | 
 **optional** | ***CreateOAuthConnectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateOAuthConnectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **emailAddress** | **optional.String**|  | 

### Return type

[**CreateOAuthConnectionResult**](CreateOAuthConnectionResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## ExchangeAuthorizationTokenAndCreateOrUpdateInbox

> CreateOAuthExchangeResult ExchangeAuthorizationTokenAndCreateOrUpdateInbox(ctx, authorizationCode, redirectUri)

Exchange authorization code for access token and create inbox

Exchange an OAuth code for an access token and create an inbox connection in MailSlurp

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorizationCode** | **string**|  | 
**redirectUri** | **string**|  | 

### Return type

[**CreateOAuthExchangeResult**](CreateOAuthExchangeResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

