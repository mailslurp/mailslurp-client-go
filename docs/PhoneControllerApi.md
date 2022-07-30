# MailSlurp\PhoneControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEmergencyAddress**](PhoneControllerApi#CreateEmergencyAddress) | **Post** /phone/emergency-addresses | 
[**DeleteEmergencyAddress**](PhoneControllerApi#DeleteEmergencyAddress) | **Delete** /phone/emergency-addresses/{addressId} | 
[**DeletePhoneNumber**](PhoneControllerApi#DeletePhoneNumber) | **Delete** /phone/numbers/{phoneNumberId} | 
[**GetEmergencyAddress**](PhoneControllerApi#GetEmergencyAddress) | **Get** /phone/emergency-addresses/{addressId} | 
[**GetEmergencyAddresses**](PhoneControllerApi#GetEmergencyAddresses) | **Get** /phone/emergency-addresses | 
[**GetPhoneNumber**](PhoneControllerApi#GetPhoneNumber) | **Get** /phone/numbers/{phoneNumberId} | 
[**GetPhoneNumbers**](PhoneControllerApi#GetPhoneNumbers) | **Get** /phone/numbers | 
[**GetPhonePlans**](PhoneControllerApi#GetPhonePlans) | **Get** /phone/plans | 
[**TestPhoneNumberSendSms**](PhoneControllerApi#TestPhoneNumberSendSms) | **Post** /phone/numbers/{phoneNumberId}/test | 



## CreateEmergencyAddress

> EmergencyAddress CreateEmergencyAddress(ctx, createEmergencyAddressOptions)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createEmergencyAddressOptions** | [**CreateEmergencyAddressOptions**](CreateEmergencyAddressOptions)|  | 

### Return type

[**EmergencyAddress**](EmergencyAddress)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DeleteEmergencyAddress

> EmptyResponseDto DeleteEmergencyAddress(ctx, addressId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addressId** | [**string**]()|  | 

### Return type

[**EmptyResponseDto**](EmptyResponseDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DeletePhoneNumber

> DeletePhoneNumber(ctx, phoneNumberId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneNumberId** | [**string**]()|  | 

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


## GetEmergencyAddress

> EmergencyAddress GetEmergencyAddress(ctx, addressId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addressId** | [**string**]()|  | 

### Return type

[**EmergencyAddress**](EmergencyAddress)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetEmergencyAddresses

> []EmergencyAddressDto GetEmergencyAddresses(ctx, )



### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]EmergencyAddressDto**](EmergencyAddressDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetPhoneNumber

> PhoneNumberDto GetPhoneNumber(ctx, phoneNumberId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneNumberId** | [**string**]()|  | 

### Return type

[**PhoneNumberDto**](PhoneNumberDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetPhoneNumbers

> PagePhoneNumberProjection GetPhoneNumbers(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetPhoneNumbersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPhoneNumbersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Optional page index for list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size for list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **since** | **optional.Time**| Filter by created at after the given timestamp | 
 **before** | **optional.Time**| Filter by created at before the given timestamp | 

### Return type

[**PagePhoneNumberProjection**](PagePhoneNumberProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetPhonePlans

> []PhonePlanDto GetPhonePlans(ctx, )



### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]PhonePlanDto**](PhonePlanDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## TestPhoneNumberSendSms

> TestPhoneNumberSendSms(ctx, phoneNumberId, testPhoneNumberOptions)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneNumberId** | [**string**]()|  | 
**testPhoneNumberOptions** | [**TestPhoneNumberOptions**](TestPhoneNumberOptions)|  | 

### Return type

 (empty response body)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

