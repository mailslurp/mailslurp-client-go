# MailSlurp\ConsentControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckSendingConsentForEmailAddress**](ConsentControllerApi#CheckSendingConsentForEmailAddress) | **Get** /consent/opt-in/sending-consent | 
[**GetOptInIdentities**](ConsentControllerApi#GetOptInIdentities) | **Get** /consent/opt-in | 
[**RevokeOptInConsentForEmailAddress**](ConsentControllerApi#RevokeOptInConsentForEmailAddress) | **Delete** /consent/opt-in | 
[**SendOptInConsentForEmailAddress**](ConsentControllerApi#SendOptInConsentForEmailAddress) | **Post** /consent/opt-in/send | Send a verification code to a user once they have explicitly submitted their email address



## CheckSendingConsentForEmailAddress

> OptInSendingConsentDto CheckSendingConsentForEmailAddress(ctx, emailAddress)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailAddress** | **string**| Email address to check consent for | 

### Return type

[**OptInSendingConsentDto**](OptInSendingConsentDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetOptInIdentities

> PageOptInIdentityProjection GetOptInIdentities(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetOptInIdentitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetOptInIdentitiesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Optional page index in list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]

### Return type

[**PageOptInIdentityProjection**](PageOptInIdentityProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## RevokeOptInConsentForEmailAddress

> OptInSendingConsentDto RevokeOptInConsentForEmailAddress(ctx, emailAddress)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailAddress** | **string**| Email address to revoke consent for | 

### Return type

[**OptInSendingConsentDto**](OptInSendingConsentDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## SendOptInConsentForEmailAddress

> OptInConsentSendResult SendOptInConsentForEmailAddress(ctx, optInConsentOptions)

Send a verification code to a user once they have explicitly submitted their email address

Send double-opt in consent for an email address

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optInConsentOptions** | [**OptInConsentOptions**](OptInConsentOptions)|  | 

### Return type

[**OptInConsentSendResult**](OptInConsentSendResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

