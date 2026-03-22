# MailSlurp\PhoneControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcquirePhonePoolLease**](PhoneControllerApi#AcquirePhonePoolLease) | **Post** /phone/pools/{poolId}/leases | Acquire phone pool lease
[**AddAllPhoneNumbersToPhonePool**](PhoneControllerApi#AddAllPhoneNumbersToPhonePool) | **Post** /phone/pools/{poolId}/numbers/add-all | Add all phone numbers to phone pool
[**AddPhoneNumbersToPhonePool**](PhoneControllerApi#AddPhoneNumbersToPhonePool) | **Post** /phone/pools/{poolId}/numbers | Add phone numbers to phone pool
[**CreateEmergencyAddress**](PhoneControllerApi#CreateEmergencyAddress) | **Post** /phone/emergency-addresses | Create an emergency address
[**CreatePhoneNumber**](PhoneControllerApi#CreatePhoneNumber) | **Post** /phone | Add phone number to your account. Only works if you have already added a plan and an initial phone number in your account and acknowledged the pricing and terms of service by enabling API phone creation.
[**CreatePhonePool**](PhoneControllerApi#CreatePhonePool) | **Post** /phone/pools | Create phone pool
[**CreatePhoneProvisioningJob**](PhoneControllerApi#CreatePhoneProvisioningJob) | **Post** /phone/provisioning/jobs | Create a phone provisioning job
[**DeleteAllPhoneNumber**](PhoneControllerApi#DeleteAllPhoneNumber) | **Delete** /phone/numbers | Delete all phone numbers
[**DeleteEmergencyAddress**](PhoneControllerApi#DeleteEmergencyAddress) | **Delete** /phone/emergency-addresses/{addressId} | Delete an emergency address
[**DeletePhoneMessageThreadItems**](PhoneControllerApi#DeletePhoneMessageThreadItems) | **Delete** /phone/numbers/{phoneNumberId}/message-threads/{otherNumber} | Delete messages in a phone thread
[**DeletePhoneNumber**](PhoneControllerApi#DeletePhoneNumber) | **Delete** /phone/numbers/{phoneNumberId} | Delete a phone number
[**DeletePhonePool**](PhoneControllerApi#DeletePhonePool) | **Delete** /phone/pools/{poolId} | Delete phone pool
[**GetAllPhoneMessageThreads**](PhoneControllerApi#GetAllPhoneMessageThreads) | **Get** /phone/numbers/message-threads | Get the latest messages for all phones
[**GetAllPhoneNumberReleases**](PhoneControllerApi#GetAllPhoneNumberReleases) | **Get** /phone/releases | Get all phone number releases
[**GetConsentStatus**](PhoneControllerApi#GetConsentStatus) | **Get** /phone/consent | Get consent status
[**GetEmergencyAddress**](PhoneControllerApi#GetEmergencyAddress) | **Get** /phone/emergency-addresses/{addressId} | Get an emergency address
[**GetEmergencyAddresses**](PhoneControllerApi#GetEmergencyAddresses) | **Get** /phone/emergency-addresses | Get emergency addresses
[**GetOrCreatePhonePool**](PhoneControllerApi#GetOrCreatePhonePool) | **Post** /phone/pools/get-or-create | Get or create phone pool
[**GetPhoneMessageThreadItems**](PhoneControllerApi#GetPhoneMessageThreadItems) | **Get** /phone/numbers/{phoneNumberId}/message-threads/{otherNumber} | Get messages in a phone thread
[**GetPhoneMessageThreads**](PhoneControllerApi#GetPhoneMessageThreads) | **Get** /phone/numbers/{phoneNumberId}/message-threads | Get the latest message preview for a thread
[**GetPhoneNumber**](PhoneControllerApi#GetPhoneNumber) | **Get** /phone/numbers/{phoneNumberId} | Get a phone number by ID
[**GetPhoneNumberByName**](PhoneControllerApi#GetPhoneNumberByName) | **Get** /phone/numbers/by-name | Get a phone number by name
[**GetPhoneNumberByPhoneNumber**](PhoneControllerApi#GetPhoneNumberByPhoneNumber) | **Get** /phone/numbers/by-phone-number | Get a phone number by phone number
[**GetPhoneNumberLineTypeIntelligence**](PhoneControllerApi#GetPhoneNumberLineTypeIntelligence) | **Post** /phone/validate/line-type-intelligence | Get line type intelligence for a phone number
[**GetPhoneNumberRelease**](PhoneControllerApi#GetPhoneNumberRelease) | **Get** /phone/releases/{releaseId} | Get phone number release
[**GetPhoneNumbers**](PhoneControllerApi#GetPhoneNumbers) | **Get** /phone/numbers | Get phone numbers
[**GetPhonePlans**](PhoneControllerApi#GetPhonePlans) | **Get** /phone/plans | Get phone plans
[**GetPhonePlansAvailability**](PhoneControllerApi#GetPhonePlansAvailability) | **Get** /phone/plans/availability | Get phone plans availability
[**GetPhonePool**](PhoneControllerApi#GetPhonePool) | **Get** /phone/pools/{poolId} | Get phone pool
[**GetPhonePoolByName**](PhoneControllerApi#GetPhonePoolByName) | **Get** /phone/pools/by-name | Get phone pool by name
[**GetPhonePools**](PhoneControllerApi#GetPhonePools) | **Get** /phone/pools | Get phone pools
[**GetPhoneProvisioningCapabilities**](PhoneControllerApi#GetPhoneProvisioningCapabilities) | **Get** /phone/provisioning/capabilities | Get phone provisioning capabilities
[**GetPhoneProvisioningJob**](PhoneControllerApi#GetPhoneProvisioningJob) | **Get** /phone/provisioning/jobs/{jobId} | Get phone provisioning job
[**GetPhoneSmsPrepaidCredit**](PhoneControllerApi#GetPhoneSmsPrepaidCredit) | **Get** /phone/sms-prepaid-credits/{creditId} | Get SMS prepaid credit
[**GetPhoneSmsPrepaidCredits**](PhoneControllerApi#GetPhoneSmsPrepaidCredits) | **Get** /phone/sms-prepaid-credits | Get SMS prepaid credits
[**GetPhoneSummary**](PhoneControllerApi#GetPhoneSummary) | **Get** /phone/summary | Get phone summary
[**GetSentSmsByPhoneNumber**](PhoneControllerApi#GetSentSmsByPhoneNumber) | **Get** /phone/numbers/{phoneNumberId}/sms-sent | List sent TXT messages for a phone number
[**GetSmsByPhoneNumber**](PhoneControllerApi#GetSmsByPhoneNumber) | **Get** /phone/numbers/{phoneNumberId}/sms | List SMS messages for a phone number
[**ReassignPhoneNumberRelease**](PhoneControllerApi#ReassignPhoneNumberRelease) | **Get** /phone/releases/{releaseId}/reassign | Reassign phone number release
[**ReleasePhonePoolLease**](PhoneControllerApi#ReleasePhonePoolLease) | **Delete** /phone/pools/{poolId}/leases/{leaseId} | Release phone pool lease
[**RemovePhoneNumberFromPhonePool**](PhoneControllerApi#RemovePhoneNumberFromPhonePool) | **Delete** /phone/pools/{poolId}/numbers/{phoneNumberId} | Remove phone number from phone pool
[**SearchAvailablePhoneNumbers**](PhoneControllerApi#SearchAvailablePhoneNumbers) | **Post** /phone/provisioning/search | Search available phone numbers
[**SendSmsFromPhoneNumber**](PhoneControllerApi#SendSmsFromPhoneNumber) | **Post** /phone/numbers/{phoneNumberId}/sms | Send TXT message from a phone number
[**SetConsentStatus**](PhoneControllerApi#SetConsentStatus) | **Post** /phone/consent | Set consent status
[**SetPhoneFavourited**](PhoneControllerApi#SetPhoneFavourited) | **Put** /phone/numbers/{phoneNumberId}/favourite | Set phone favourited state
[**TestPhoneNumberSendSms**](PhoneControllerApi#TestPhoneNumberSendSms) | **Post** /phone/numbers/{phoneNumberId}/test | Test sending an SMS to a number
[**UpdatePhoneNumber**](PhoneControllerApi#UpdatePhoneNumber) | **Put** /phone/numbers/{phoneNumberId} | Update a phone number
[**UpdatePhonePool**](PhoneControllerApi#UpdatePhonePool) | **Put** /phone/pools/{poolId} | Update phone pool
[**ValidatePhoneNumber**](PhoneControllerApi#ValidatePhoneNumber) | **Post** /phone/validate | Verify validity of a phone number



## AcquirePhonePoolLease

> PhonePoolLeaseDto AcquirePhonePoolLease(ctx, poolId, acquirePhonePoolLeaseOptions)

Acquire phone pool lease

Acquire an available phone number from the pool and mark it leased

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | [**string**]()|  | 
**acquirePhonePoolLeaseOptions** | [**AcquirePhonePoolLeaseOptions**](AcquirePhonePoolLeaseOptions)|  | 

### Return type

[**PhonePoolLeaseDto**](PhonePoolLeaseDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AddAllPhoneNumbersToPhonePool

> PhonePoolDetailDto AddAllPhoneNumbersToPhonePool(ctx, poolId)

Add all phone numbers to phone pool

Add all active owned phone numbers to a pool

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | [**string**]()|  | 

### Return type

[**PhonePoolDetailDto**](PhonePoolDetailDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AddPhoneNumbersToPhonePool

> PhonePoolDetailDto AddPhoneNumbersToPhonePool(ctx, poolId, addPhonePoolNumbersOptions)

Add phone numbers to phone pool

Add one or more owned phone numbers to a pool

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | [**string**]()|  | 
**addPhonePoolNumbersOptions** | [**AddPhonePoolNumbersOptions**](AddPhonePoolNumbersOptions)|  | 

### Return type

[**PhonePoolDetailDto**](PhonePoolDetailDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


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


## CreatePhonePool

> PhonePoolDetailDto CreatePhonePool(ctx, createPhonePoolOptions)

Create phone pool

Create a reusable pool of phone numbers for coordinated leasing

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createPhonePoolOptions** | [**CreatePhonePoolOptions**](CreatePhonePoolOptions)|  | 

### Return type

[**PhonePoolDetailDto**](PhonePoolDetailDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CreatePhoneProvisioningJob

> PhoneProvisioningJobDto CreatePhoneProvisioningJob(ctx, createPhoneProvisioningJobOptions)

Create a phone provisioning job

Create an advanced phone provisioning job from shortlisted numbers

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createPhoneProvisioningJobOptions** | [**CreatePhoneProvisioningJobOptions**](CreatePhoneProvisioningJobOptions)|  | 

### Return type

[**PhoneProvisioningJobDto**](PhoneProvisioningJobDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DeleteAllPhoneNumber

> DeleteAllPhoneNumber(ctx, )

Delete all phone numbers

Remove all phone number from account

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


## DeletePhoneMessageThreadItems

> EmptyResponseDto DeletePhoneMessageThreadItems(ctx, phoneNumberId, otherNumber)

Delete messages in a phone thread

Delete all messages in an SMS thread

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneNumberId** | [**string**]()|  | 
**otherNumber** | **string**|  | 

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


## DeletePhonePool

> DeletePhonePool(ctx, poolId)

Delete phone pool

Delete a phone pool and release any active leases from that pool

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | [**string**]()|  | 

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


## GetAllPhoneNumberReleases

> PagePhoneNumberReleaseProjection GetAllPhoneNumberReleases(ctx, optional)

Get all phone number releases

List released or deleted phone numbers

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllPhoneNumberReleasesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllPhoneNumberReleasesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**|  | [default to 0]
 **size** | **optional.Int32**|  | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to DESC]

### Return type

[**PagePhoneNumberReleaseProjection**](PagePhoneNumberReleaseProjection)

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


## GetOrCreatePhonePool

> PhonePoolDetailDto GetOrCreatePhonePool(ctx, getOrCreatePhonePoolOptions)

Get or create phone pool

Get a phone pool by name or create it if it does not exist

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**getOrCreatePhonePoolOptions** | [**GetOrCreatePhonePoolOptions**](GetOrCreatePhonePoolOptions)|  | 

### Return type

[**PhonePoolDetailDto**](PhonePoolDetailDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
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


## GetPhoneNumberByName

> PhoneNumberDto GetPhoneNumberByName(ctx, name)

Get a phone number by name

Get a phone number by name

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**|  | 

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


## GetPhoneNumberByPhoneNumber

> PhoneNumberDto GetPhoneNumberByPhoneNumber(ctx, phoneNumber)

Get a phone number by phone number

Get a phone number by phone number

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneNumber** | **string**|  | 

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


## GetPhoneNumberLineTypeIntelligence

> PhoneNumberLineTypeLookupDto GetPhoneNumberLineTypeIntelligence(ctx, validatePhoneNumberOptions)

Get line type intelligence for a phone number

Lookup line type intelligence for a phone number

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**validatePhoneNumberOptions** | [**ValidatePhoneNumberOptions**](ValidatePhoneNumberOptions)|  | 

### Return type

[**PhoneNumberLineTypeLookupDto**](PhoneNumberLineTypeLookupDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetPhoneNumberRelease

> PhoneNumberReleaseProjection GetPhoneNumberRelease(ctx, releaseId)

Get phone number release

Get a released or deleted phone numbers

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseId** | [**string**]()|  | 

### Return type

[**PhoneNumberReleaseProjection**](PhoneNumberReleaseProjection)

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
 **lineType** | **optional.String**| Optional line type filter | 
 **carrierName** | **optional.String**| Optional carrier name filter | 
 **mobileCountryCode** | **optional.String**| Optional mobile country code filter | 
 **mobileNetworkCode** | **optional.String**| Optional mobile network code filter | 
 **providerLabel** | **optional.String**| Optional provider label filter such as T1 or T2 | 
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


## GetPhonePool

> PhonePoolDetailDto GetPhonePool(ctx, poolId)

Get phone pool

Get phone pool details by ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | [**string**]()|  | 

### Return type

[**PhonePoolDetailDto**](PhonePoolDetailDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetPhonePoolByName

> PhonePoolDetailDto GetPhonePoolByName(ctx, name)

Get phone pool by name

Get phone pool details by name

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**|  | 

### Return type

[**PhonePoolDetailDto**](PhonePoolDetailDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetPhonePools

> []PhonePoolDto GetPhonePools(ctx, )

Get phone pools

List phone pools for the authenticated user

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]PhonePoolDto**](PhonePoolDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetPhoneProvisioningCapabilities

> PhoneProviderCapabilitiesResult GetPhoneProvisioningCapabilities(ctx, phoneCountry, optional)

Get phone provisioning capabilities

Get supported provider-country variant capabilities for advanced provisioning

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneCountry** | **string**|  | 
 **optional** | ***GetPhoneProvisioningCapabilitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPhoneProvisioningCapabilitiesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **providerLabel** | **optional.String**|  | 

### Return type

[**PhoneProviderCapabilitiesResult**](PhoneProviderCapabilitiesResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetPhoneProvisioningJob

> PhoneProvisioningJobDto GetPhoneProvisioningJob(ctx, jobId)

Get phone provisioning job

Get advanced phone provisioning job status

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | [**string**]()|  | 

### Return type

[**PhoneProvisioningJobDto**](PhoneProvisioningJobDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetPhoneSmsPrepaidCredit

> PhoneSmsPrepaidCreditDto GetPhoneSmsPrepaidCredit(ctx, creditId)

Get SMS prepaid credit

Get a specific SMS prepaid credit balance for the authenticated account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**creditId** | [**string**]()|  | 

### Return type

[**PhoneSmsPrepaidCreditDto**](PhoneSmsPrepaidCreditDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetPhoneSmsPrepaidCredits

> PhoneSmsPrepaidCreditsDto GetPhoneSmsPrepaidCredits(ctx, )

Get SMS prepaid credits

List SMS prepaid credits for the authenticated account

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**PhoneSmsPrepaidCreditsDto**](PhoneSmsPrepaidCreditsDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetPhoneSummary

> PhoneSummaryDto GetPhoneSummary(ctx, )

Get phone summary

Get overview of assigned phones

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**PhoneSummaryDto**](PhoneSummaryDto)

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


## ReassignPhoneNumberRelease

> PhoneNumberDto ReassignPhoneNumberRelease(ctx, releaseId)

Reassign phone number release

Reassign phone number that was released or deleted

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseId** | [**string**]()|  | 

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


## ReleasePhonePoolLease

> ReleasePhonePoolLease(ctx, poolId, leaseId)

Release phone pool lease

Release an active phone pool lease

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | [**string**]()|  | 
**leaseId** | [**string**]()|  | 

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


## RemovePhoneNumberFromPhonePool

> RemovePhoneNumberFromPhonePool(ctx, poolId, phoneNumberId)

Remove phone number from phone pool

Remove a phone number from a pool. If the number is leased from this pool the lease is released.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | [**string**]()|  | 
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


## SearchAvailablePhoneNumbers

> AvailablePhoneNumbersResult SearchAvailablePhoneNumbers(ctx, searchAvailablePhoneNumbersOptions)

Search available phone numbers

Search available numbers for advanced provisioning

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**searchAvailablePhoneNumbersOptions** | [**SearchAvailablePhoneNumbersOptions**](SearchAvailablePhoneNumbersOptions)|  | 

### Return type

[**AvailablePhoneNumbersResult**](AvailablePhoneNumbersResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
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

Test a phone number by sending an SMS to it. NOTE this is only for internal use to check that a phone number is working. For end-to-end phone testing see https://docs.mailslurp.com/txt-sms/

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


## UpdatePhonePool

> PhonePoolDetailDto UpdatePhonePool(ctx, poolId, updatePhonePoolOptions)

Update phone pool

Update phone pool metadata such as name or description

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | [**string**]()|  | 
**updatePhonePoolOptions** | [**UpdatePhonePoolOptions**](UpdatePhonePoolOptions)|  | 

### Return type

[**PhonePoolDetailDto**](PhonePoolDetailDto)

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

