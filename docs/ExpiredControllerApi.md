# MailSlurp\ExpiredControllerApi

All URIs are relative to *https://api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetExpirationDefaults**](ExpiredControllerApi#GetExpirationDefaults) | **Get** /expired/defaults | Get default expiration settings
[**GetExpiredInboxByInboxId**](ExpiredControllerApi#GetExpiredInboxByInboxId) | **Get** /expired/inbox/{inboxId} | Get expired inbox record for a previously existing inbox
[**GetExpiredInboxRecord**](ExpiredControllerApi#GetExpiredInboxRecord) | **Get** /expired/{expiredId} | Get an expired inbox record
[**GetExpiredInboxes**](ExpiredControllerApi#GetExpiredInboxes) | **Get** /expired | List records of expired inboxes



## GetExpirationDefaults

> ExpirationDefaults GetExpirationDefaults(ctx, )

Get default expiration settings

Return default times used for inbox expiration

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ExpirationDefaults**](ExpirationDefaults)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetExpiredInboxByInboxId

> ExpiredInboxDto GetExpiredInboxByInboxId(ctx, inboxId)

Get expired inbox record for a previously existing inbox

Use the inboxId to return an ExpiredInboxRecord if an inbox has expired. Inboxes expire and are disabled if an expiration date is set or plan requires. Returns 404 if no expired inbox is found for the inboxId

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()| ID of inbox you want to retrieve (not the inbox ID) | 

### Return type

[**ExpiredInboxDto**](ExpiredInboxDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetExpiredInboxRecord

> ExpiredInboxDto GetExpiredInboxRecord(ctx, iDOfTheExpiredInboxRecordYouWantToRetrieveThisIsDifferentFromTheIDOfTheInboxYouAreInterestedInSeeOtherMethodsForGettingExpiredInboxRecordForAnInboxInboxId)

Get an expired inbox record

Inboxes created with an expiration date will expire after the given date and be moved to an ExpiredInbox entity. You can still read emails in the inbox but it can no longer send or receive emails. Fetch the expired inboxes to view the old inboxes properties

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iDOfTheExpiredInboxRecordYouWantToRetrieveThisIsDifferentFromTheIDOfTheInboxYouAreInterestedInSeeOtherMethodsForGettingExpiredInboxRecordForAnInboxInboxId** | [**string**]()|  | 

### Return type

[**ExpiredInboxDto**](ExpiredInboxDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetExpiredInboxes

> PageExpiredInboxRecordProjection GetExpiredInboxes(ctx, optional)

List records of expired inboxes

Inboxes created with an expiration date will expire after the given date. An ExpiredInboxRecord is created that records the inboxes old ID and email address. You can still read emails in the inbox (using the inboxes old ID) but the email address associated with the inbox can no longer send or receive emails. Fetch expired inbox records to view the old inboxes properties

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetExpiredInboxesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetExpiredInboxesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Optional page index in inbox sent email list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in inbox sent email list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **since** | **optional.Time**| Filter by created at after the given timestamp | 
 **before** | **optional.Time**| Filter by created at before the given timestamp | 

### Return type

[**PageExpiredInboxRecordProjection**](PageExpiredInboxRecordProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

