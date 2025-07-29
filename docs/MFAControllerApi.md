# MailSlurp\MFAControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTotpDeviceForBase32SecretKey**](MFAControllerApi#CreateTotpDeviceForBase32SecretKey) | **Post** /mfa/totp/device/base32SecretKey | Create a TOTP device from an base32 secret key
[**CreateTotpDeviceForCustom**](MFAControllerApi#CreateTotpDeviceForCustom) | **Post** /mfa/totp/device/custom | Create a TOTP device from custom options
[**CreateTotpDeviceForOtpAuthUrl**](MFAControllerApi#CreateTotpDeviceForOtpAuthUrl) | **Post** /mfa/totp/device/otpAuthUrl | Create a TOTP device from an OTP Auth URL
[**GetTotpDevice**](MFAControllerApi#GetTotpDevice) | **Get** /mfa/totp/device/{id} | Get a TOTP device by ID
[**GetTotpDeviceBy**](MFAControllerApi#GetTotpDeviceBy) | **Get** /mfa/totp/device/by | Get a TOTP device by username, issuer, or name. Returns empty if not found.
[**GetTotpDeviceCode**](MFAControllerApi#GetTotpDeviceCode) | **Get** /mfa/totp/device/{id}/code | Get a TOTP device code by device ID



## CreateTotpDeviceForBase32SecretKey

> TotpDeviceDto CreateTotpDeviceForBase32SecretKey(ctx, createTotpDeviceBase32SecretKeyOptions)

Create a TOTP device from an base32 secret key

Create a virtual TOTP device for a given secret key. This is usually present as an alternative login option when pairing OTP devices.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createTotpDeviceBase32SecretKeyOptions** | [**CreateTotpDeviceBase32SecretKeyOptions**](CreateTotpDeviceBase32SecretKeyOptions)|  | 

### Return type

[**TotpDeviceDto**](TotpDeviceDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CreateTotpDeviceForCustom

> TotpDeviceDto CreateTotpDeviceForCustom(ctx, createTotpDeviceCustomOptions)

Create a TOTP device from custom options

Create a virtual TOTP device for custom options specifying all parameters of the TOTP device.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createTotpDeviceCustomOptions** | [**CreateTotpDeviceCustomOptions**](CreateTotpDeviceCustomOptions)|  | 

### Return type

[**TotpDeviceDto**](TotpDeviceDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CreateTotpDeviceForOtpAuthUrl

> TotpDeviceDto CreateTotpDeviceForOtpAuthUrl(ctx, createTotpDeviceOtpAuthUrlOptions)

Create a TOTP device from an OTP Auth URL

Create a virtual TOTP device for a given OTP Auth URL such as otpauth://totp/MyApp:alice@example.com?secret=ABC123&issuer=MyApp&period=30&digits=6&algorithm=SHA1. You can provider overrides in the options for each component of the URL.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createTotpDeviceOtpAuthUrlOptions** | [**CreateTotpDeviceOtpAuthUrlOptions**](CreateTotpDeviceOtpAuthUrlOptions)|  | 

### Return type

[**TotpDeviceDto**](TotpDeviceDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetTotpDevice

> TotpDeviceDto GetTotpDevice(ctx, id)

Get a TOTP device by ID

Get Time-Based One-Time Password (TOTP) device by its ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()|  | 

### Return type

[**TotpDeviceDto**](TotpDeviceDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetTotpDeviceBy

> TotpDeviceOptionalDto GetTotpDeviceBy(ctx, optional)

Get a TOTP device by username, issuer, or name. Returns empty if not found.

Get Time-Based One-Time Password (TOTP) device by its username and issuer mapping.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetTotpDeviceByOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTotpDeviceByOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Optional name filter | 
 **issuer** | **optional.String**| Optional issuer filter | 
 **username** | **optional.String**| Optional username filter | 

### Return type

[**TotpDeviceOptionalDto**](TotpDeviceOptionalDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetTotpDeviceCode

> TotpDeviceCodeDto GetTotpDeviceCode(ctx, id, optional)

Get a TOTP device code by device ID

Get Time-Based One-Time Password for a device by its ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()| ID of the TOTP device to get the code for | 
 **optional** | ***GetTotpDeviceCodeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTotpDeviceCodeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **at** | **optional.Time**| Optional time to get code for. If not provided, current time is used. | 
 **minSecondsUntilExpire** | **optional.Int32**| Optional minimum time until code expires. Will hold thread open until period reached. | [default to 5]

### Return type

[**TotpDeviceCodeDto**](TotpDeviceCodeDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

