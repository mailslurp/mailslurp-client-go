# MailSlurp\PhoneControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEmergencyAddress**](PhoneControllerApi#CreateEmergencyAddress) | **Post** /phone/emergency-addresses | Create an emergency address
[**CreatePhoneNumber**](PhoneControllerApi#CreatePhoneNumber) | **Post** /phone | Add phone number to your account. Only works if you have already added a plan and an initial phone number in your account and acknowledged the pricing and terms of service by enabling API phone creation.
[**DeleteEmergencyAddress**](PhoneControllerApi#DeleteEmergencyAddress) | **Delete** /phone/emergency-addresses/{addressId} | Delete an emergency address
[**DeletePhoneNumber**](PhoneControllerApi#DeletePhoneNumber) | **Delete** /phone/numbers/{phoneNumberId} | Delete a phone number
[**GetAllPhoneMessageThreads**](PhoneControllerApi#GetAllPhoneMessageThreads) | **Get** /phone/numbers/message-threads | Get the latest messages for all phones
[**GetConsentStatus**](PhoneControllerApi#GetConsentStatus) | **Get** /phone/consent | Get consent status
[**GetEmergencyAddress**](PhoneControllerApi#GetEmergencyAddress) | **Get** /phone/emergency-addresses/{addressId} | Get an emergency address
[**GetEmergencyAddresses**](PhoneControllerApi#GetEmergencyAddresses) | **Get** /phone/emergency-addresses | Get emergency addresses
[**GetPhoneMessageThreadItems**](PhoneControllerApi#GetPhoneMessageThreadItems) | **Get** /phone/numbers/{phoneNumberId}/message-threads/{otherNumber} | Get messages in a phone thread
[**GetPhoneMessageThreads**](PhoneControllerApi#GetPhoneMessageThreads) | **Get** /phone/numbers/{phoneNumberId}/message-threads | Get the latest message preview for a thread
[**GetPhoneNumber**](PhoneControllerApi#GetPhoneNumber) | **Get** /phone/numbers/{phoneNumberId} | Get a phone number by ID
[**GetPhoneNumbers**](PhoneControllerApi#GetPhoneNumbers) | **Get** /phone/numbers | Get phone numbers
[**GetPhonePlans**](PhoneControllerApi#GetPhonePlans) | **Get** /phone/plans | Get phone plans
[**GetPhonePlansAvailability**](PhoneControllerApi#GetPhonePlansAvailability) | **Get** /phone/plans/availability | Get phone plans availability
[**GetSentSmsByPhoneNumber**](PhoneControllerApi#GetSentSmsByPhoneNumber) | **Get** /phone/numbers/{phoneNumberId}/sms-sent | List sent TXT messages for a phone number
[**GetSmsByPhoneNumber**](PhoneControllerApi#GetSmsByPhoneNumber) | **Get** /phone/numbers/{phoneNumberId}/sms | List SMS messages for a phone number
[**SendSmsFromPhoneNumber**](PhoneControllerApi#SendSmsFromPhoneNumber) | **Post** /phone/numbers/{phoneNumberId}/sms | Send TXT message from a phone number
[**SetConsentStatus**](PhoneControllerApi#SetConsentStatus) | **Post** /phone/consent | Set consent status
[**SetPhoneFavourited**](PhoneControllerApi#SetPhoneFavourited) | **Put** /phone/numbers/{phoneNumberId}/favourite | Set phone favourited state
[**TestPhoneNumberSendSms**](PhoneControllerApi#TestPhoneNumberSendSms) | **Post** /phone/numbers/{phoneNumberId}/test | Test sending an SMS to a number
[**UpdatePhoneNumber**](PhoneControllerApi#UpdatePhoneNumber) | **Put** /phone/numbers/{phoneNumberId} | Update a phone number
[**ValidatePhoneNumber**](PhoneControllerApi#ValidatePhoneNumber) | **Post** /phone/validate | Verify validity of a phone number



## CreateEmergencyAddress

> EmergencyAddress CreateEmergencyAddress(ctx, createEmergencyAddressOptions)

Create an emergency address

Add an emergency address to a phone number

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


## CreatePhoneNumber

> PhoneNumberDto CreatePhoneNumber(ctx, createPhoneNumberOptions)

Add phone number to your account. Only works if you have already added a plan and an initial phone number in your account and acknowledged the pricing and terms of service by enabling API phone creation.

Create new phone number

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createPhoneNumberOptions** | [**CreatePhoneNumberOptions**](CreatePhoneNumberOptions)|  | 

### Return type

[**PhoneNumberDto**](PhoneNumberDto)

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

Delete an emergency address

Delete an emergency address

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

Delete a phone number

Remove phone number from account

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


## GetAllPhoneMessageThreads

> PagePhoneMessageThreadProjection GetAllPhoneMessageThreads(ctx, optional)

Get the latest messages for all phones

List all message threads for all phones

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllPhoneMessageThreadsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllPhoneMessageThreadsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**|  | [default to 0]
 **size** | **optional.Int32**|  | [default to 20]

### Return type

[**PagePhoneMessageThreadProjection**](PagePhoneMessageThreadProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetConsentStatus

> ConsentStatusDto GetConsentStatus(ctx, )

Get consent status

Get the status of phone usage consent

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ConsentStatusDto**](ConsentStatusDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetEmergencyAddress

> EmergencyAddress GetEmergencyAddress(ctx, addressId)

Get an emergency address

Fetch an emergency address by ID

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

Get emergency addresses

List emergency addresses

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


## GetPhoneMessageThreadItems

> PagePhoneMessageThreadItemProjection GetPhoneMessageThreadItems(ctx, phoneNumberId, otherNumber, optional)

Get messages in a phone thread

List message thread messages for a phone message thread

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneNumberId** | [**string**]()|  | 
**otherNumber** | **string**|  | 
 **optional** | ***GetPhoneMessageThreadItemsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPhoneMessageThreadItemsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**|  | [default to 0]
 **size** | **optional.Int32**|  | [default to 20]

### Return type

[**PagePhoneMessageThreadItemProjection**](PagePhoneMessageThreadItemProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetPhoneMessageThreads

> PagePhoneMessageThreadProjection GetPhoneMessageThreads(ctx, phoneNumberId, optional)

Get the latest message preview for a thread

List message threads for a phone

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneNumberId** | [**string**]()|  | 
 **optional** | ***GetPhoneMessageThreadsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPhoneMessageThreadsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**|  | [default to 0]
 **size** | **optional.Int32**|  | [default to 20]

### Return type

[**PagePhoneMessageThreadProjection**](PagePhoneMessageThreadProjection)

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

Get a phone number by ID

Get a phone number by ID

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

Get phone numbers

List phone numbers for account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetPhoneNumbersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPhoneNumbersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phoneCountry** | **optional.String**| Optional phone country | 
 **page** | **optional.Int32**| Optional page index for list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size for list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **since** | **optional.Time**| Filter by created at after the given timestamp | 
 **before** | **optional.Time**| Filter by created at before the given timestamp | 
 **search** | **optional.String**| Optional search filter | 
 **include** | [**optional.Interface of []string**](string)| Optional phoneIds to include in result | 
 **favourite** | **optional.Bool**| Optionally filter results for favourites only | [default to false]

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

Get phone plans

Get phone number plans

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


## GetPhonePlansAvailability

> PhonePlanAvailability GetPhonePlansAvailability(ctx, )

Get phone plans availability

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**PhonePlanAvailability**](PhonePlanAvailability)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetSentSmsByPhoneNumber

> PageSentSmsProjection GetSentSmsByPhoneNumber(ctx, phoneNumberId, optional)

List sent TXT messages for a phone number

Get sent SMS messages for a phone number

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneNumberId** | [**string**]()|  | 
 **optional** | ***GetSentSmsByPhoneNumberOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSentSmsByPhoneNumberOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Optional page index in SMS list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in SMS list pagination. Maximum size is 100. Use page index and sort to page through larger results | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **since** | **optional.Time**| Optional filter SMSs received after given date time | 
 **before** | **optional.Time**| Optional filter SMSs received before given date time | 
 **search** | **optional.String**| Optional search filter | 

### Return type

[**PageSentSmsProjection**](PageSentSmsProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetSmsByPhoneNumber

> PageSmsProjection GetSmsByPhoneNumber(ctx, phoneNumberId, optional)

List SMS messages for a phone number

Get SMS messages for a phone number

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneNumberId** | [**string**]()|  | 
 **optional** | ***GetSmsByPhoneNumberOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSmsByPhoneNumberOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Optional page index in SMS list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in SMS list pagination. Maximum size is 100. Use page index and sort to page through larger results | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **unreadOnly** | **optional.Bool**| Optional filter for unread SMS only. All SMS are considered unread until they are viewed in the dashboard or requested directly | [default to false]
 **since** | **optional.Time**| Optional filter SMSs received after given date time | 
 **before** | **optional.Time**| Optional filter SMSs received before given date time | 
 **search** | **optional.String**| Optional search filter | 
 **favourite** | **optional.Bool**| Optionally filter results for favourites only | [default to false]

### Return type

[**PageSmsProjection**](PageSmsProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## SendSmsFromPhoneNumber

> SentSmsDto SendSmsFromPhoneNumber(ctx, phoneNumberId, smsSendOptions)

Send TXT message from a phone number

Send SMS from a phone number

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneNumberId** | [**string**]()|  | 
**smsSendOptions** | [**SmsSendOptions**](SmsSendOptions)|  | 

### Return type

[**SentSmsDto**](SentSmsDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## SetConsentStatus

> ConsentStatusDto SetConsentStatus(ctx, agree)

Set consent status

Give or revoke consent for phone usage

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agree** | **bool**|  | 

### Return type

[**ConsentStatusDto**](ConsentStatusDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## SetPhoneFavourited

> PhoneNumberDto SetPhoneFavourited(ctx, phoneNumberId, setPhoneFavouritedOptions)

Set phone favourited state

Set and return new favorite state for a phone

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneNumberId** | [**string**]()| ID of phone to set favourite state | 
**setPhoneFavouritedOptions** | [**SetPhoneFavouritedOptions**](SetPhoneFavouritedOptions)|  | 

### Return type

[**PhoneNumberDto**](PhoneNumberDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## TestPhoneNumberSendSms

> TestPhoneNumberSendSms(ctx, phoneNumberId, testPhoneNumberOptions, optional)

Test sending an SMS to a number

Test a phone number by sending an SMS to it

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneNumberId** | [**string**]()|  | 
**testPhoneNumberOptions** | [**TestPhoneNumberOptions**](TestPhoneNumberOptions)|  | 
 **optional** | ***TestPhoneNumberSendSmsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TestPhoneNumberSendSmsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xTestId** | **optional.String**|  | 

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


## UpdatePhoneNumber

> PhoneNumberDto UpdatePhoneNumber(ctx, phoneNumberId, updatePhoneNumberOptions)

Update a phone number

Set field for phone number

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneNumberId** | [**string**]()| ID of phone to set favourite state | 
**updatePhoneNumberOptions** | [**UpdatePhoneNumberOptions**](UpdatePhoneNumberOptions)|  | 

### Return type

[**PhoneNumberDto**](PhoneNumberDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## ValidatePhoneNumber

> PhoneNumberValidationDto ValidatePhoneNumber(ctx, validatePhoneNumberOptions)

Verify validity of a phone number

Validate a phone number

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**validatePhoneNumberOptions** | [**ValidatePhoneNumberOptions**](ValidatePhoneNumberOptions)|  | 

### Return type

[**PhoneNumberValidationDto**](PhoneNumberValidationDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

