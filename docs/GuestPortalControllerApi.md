# MailSlurp\GuestPortalControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGuestPortal**](GuestPortalControllerApi#CreateGuestPortal) | **Post** /guest-portal | Create a portal page for your customers or clients to log into email accounts and view emails.
[**CreateGuestPortalUser**](GuestPortalControllerApi#CreateGuestPortalUser) | **Post** /guest-portal/{portalId}/user | Create a portal guest user
[**GetAllGuestPortalUsers**](GuestPortalControllerApi#GetAllGuestPortalUsers) | **Get** /guest-portal/user | Get all guest users for portal
[**GetGuestPortal**](GuestPortalControllerApi#GetGuestPortal) | **Get** /guest-portal/{portalId} | Get a client email portal
[**GetGuestPortalUser**](GuestPortalControllerApi#GetGuestPortalUser) | **Get** /guest-portal/{portalId}/user/{guestId} | Get guest user for portal
[**GetGuestPortalUserById**](GuestPortalControllerApi#GetGuestPortalUserById) | **Get** /guest-portal/user/{guestId} | Get guest user
[**GetGuestPortalUsers**](GuestPortalControllerApi#GetGuestPortalUsers) | **Get** /guest-portal/{portalId}/user | Get all guest users for portal
[**GetGuestPortals**](GuestPortalControllerApi#GetGuestPortals) | **Get** /guest-portal | Get guest portals



## CreateGuestPortal

> GuestPortalDto CreateGuestPortal(ctx, createPortalOptions)

Create a portal page for your customers or clients to log into email accounts and view emails.

Create a guest login page for customers or clients to access assigned email addresses

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createPortalOptions** | [**CreatePortalOptions**](CreatePortalOptions)|  | 

### Return type

[**GuestPortalDto**](GuestPortalDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CreateGuestPortalUser

> GuestPortalUserCreateDto CreateGuestPortalUser(ctx, portalId, createPortalUserOptions)

Create a portal guest user

Add customer to portal

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portalId** | [**string**]()|  | 
**createPortalUserOptions** | [**CreatePortalUserOptions**](CreatePortalUserOptions)|  | 

### Return type

[**GuestPortalUserCreateDto**](GuestPortalUserCreateDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetAllGuestPortalUsers

> PageGuestPortalUsers GetAllGuestPortalUsers(ctx, optional)

Get all guest users for portal

Get all customers for a portal

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllGuestPortalUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllGuestPortalUsersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **portalId** | [**optional.Interface of string**]()| Optional portal ID | 
 **search** | **optional.String**| Optional search term | 
 **page** | **optional.Int32**| Optional page index in list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **since** | **optional.Time**| Filter by created at after the given timestamp | 
 **before** | **optional.Time**| Filter by created at before the given timestamp | 

### Return type

[**PageGuestPortalUsers**](PageGuestPortalUsers)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetGuestPortal

> GuestPortalDto GetGuestPortal(ctx, portalId)

Get a client email portal

Fetch a customer guest portal

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portalId** | [**string**]()|  | 

### Return type

[**GuestPortalDto**](GuestPortalDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetGuestPortalUser

> GuestPortalUserDto GetGuestPortalUser(ctx, portalId, guestId)

Get guest user for portal

Get customer for portal

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portalId** | [**string**]()|  | 
**guestId** | [**string**]()|  | 

### Return type

[**GuestPortalUserDto**](GuestPortalUserDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetGuestPortalUserById

> GuestPortalUserDto GetGuestPortalUserById(ctx, guestId)

Get guest user

Get customer by ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guestId** | [**string**]()|  | 

### Return type

[**GuestPortalUserDto**](GuestPortalUserDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetGuestPortalUsers

> PageGuestPortalUsers GetGuestPortalUsers(ctx, portalId, optional)

Get all guest users for portal

Get customers for a portal

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portalId** | [**string**]()|  | 
 **optional** | ***GetGuestPortalUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetGuestPortalUsersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **optional.String**| Optional search term | 
 **page** | **optional.Int32**| Optional page index in list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **since** | **optional.Time**| Filter by created at after the given timestamp | 
 **before** | **optional.Time**| Filter by created at before the given timestamp | 

### Return type

[**PageGuestPortalUsers**](PageGuestPortalUsers)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetGuestPortals

> []GuestPortalDto GetGuestPortals(ctx, )

Get guest portals

Get portals

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]GuestPortalDto**](GuestPortalDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

