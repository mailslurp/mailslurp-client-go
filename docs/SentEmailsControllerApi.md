# MailSlurp\SentEmailsControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAllSentEmails**](SentEmailsControllerApi#DeleteAllSentEmails) | **Delete** /sent | Delete all sent email receipts
[**DeleteSentEmail**](SentEmailsControllerApi#DeleteSentEmail) | **Delete** /sent/{id} | Delete sent email receipt
[**GetAllSentTrackingPixels**](SentEmailsControllerApi#GetAllSentTrackingPixels) | **Get** /sent/tracking-pixels | 
[**GetRawSentEmailContents**](SentEmailsControllerApi#GetRawSentEmailContents) | **Get** /sent/{emailId}/raw | Get raw sent email string. Returns unparsed raw SMTP message with headers and body.
[**GetRawSentEmailJson**](SentEmailsControllerApi#GetRawSentEmailJson) | **Get** /sent/{emailId}/raw/json | Get raw sent email in JSON. Unparsed SMTP message in JSON wrapper format.
[**GetSentDeliveryStatus**](SentEmailsControllerApi#GetSentDeliveryStatus) | **Get** /sent/delivery-status/{deliveryId} | 
[**GetSentDeliveryStatuses**](SentEmailsControllerApi#GetSentDeliveryStatuses) | **Get** /sent/delivery-status | 
[**GetSentDeliveryStatusesBySentId**](SentEmailsControllerApi#GetSentDeliveryStatusesBySentId) | **Get** /sent/{sentId}/delivery-status | 
[**GetSentEmail**](SentEmailsControllerApi#GetSentEmail) | **Get** /sent/{id} | Get sent email receipt
[**GetSentEmailHTMLContent**](SentEmailsControllerApi#GetSentEmailHTMLContent) | **Get** /sent/{id}/html | Get sent email HTML content
[**GetSentEmailPreviewURLs**](SentEmailsControllerApi#GetSentEmailPreviewURLs) | **Get** /sent/{id}/urls | Get sent email URL for viewing in browser or downloading
[**GetSentEmailTrackingPixels**](SentEmailsControllerApi#GetSentEmailTrackingPixels) | **Get** /sent/{id}/tracking-pixels | 
[**GetSentEmails**](SentEmailsControllerApi#GetSentEmails) | **Get** /sent | Get all sent emails in paginated form
[**GetSentEmailsWithQueueResults**](SentEmailsControllerApi#GetSentEmailsWithQueueResults) | **Get** /sent/queue-results | Get results of email sent with queues in paginated form
[**GetSentOrganizationEmails**](SentEmailsControllerApi#GetSentOrganizationEmails) | **Get** /sent/organization | 



## DeleteAllSentEmails

> DeleteAllSentEmails(ctx, )

Delete all sent email receipts

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


## DeleteSentEmail

> DeleteSentEmail(ctx, id)

Delete sent email receipt

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()|  | 

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


## GetAllSentTrackingPixels

> PageTrackingPixelProjection GetAllSentTrackingPixels(ctx, optional)



Get all sent email tracking pixels in paginated form

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllSentTrackingPixelsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllSentTrackingPixelsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Optional page index in sent email tracking pixel list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in sent email tracking pixel list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **searchFilter** | **optional.String**| Optional search filter | 
 **since** | **optional.Time**| Filter by created at after the given timestamp | 
 **before** | **optional.Time**| Filter by created at before the given timestamp | 

### Return type

[**PageTrackingPixelProjection**](PageTrackingPixelProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetRawSentEmailContents

> string GetRawSentEmailContents(ctx, emailId)

Get raw sent email string. Returns unparsed raw SMTP message with headers and body.

Returns a raw, unparsed, and unprocessed sent email. If your client has issues processing the response it is likely due to the response content-type which is text/plain. If you need a JSON response content-type use the getRawSentEmailJson endpoint

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()| ID of email | 

### Return type

**string**

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetRawSentEmailJson

> RawEmailJson GetRawSentEmailJson(ctx, emailId)

Get raw sent email in JSON. Unparsed SMTP message in JSON wrapper format.

Returns a raw, unparsed, and unprocessed sent email wrapped in a JSON response object for easier handling when compared with the getRawSentEmail text/plain response

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()| ID of email | 

### Return type

[**RawEmailJson**](RawEmailJson)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetSentDeliveryStatus

> DeliveryStatus GetSentDeliveryStatus(ctx, deliveryId)



Get a sent email delivery status

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deliveryId** | [**string**]()|  | 

### Return type

[**DeliveryStatus**](DeliveryStatus)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetSentDeliveryStatuses

> PageDeliveryStatus GetSentDeliveryStatuses(ctx, optional)



Get all sent email delivery statuses

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetSentDeliveryStatusesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSentDeliveryStatusesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Optional page index in delivery status list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in delivery status list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **since** | **optional.Time**| Filter by created at after the given timestamp | 
 **before** | **optional.Time**| Filter by created at before the given timestamp | 

### Return type

[**PageDeliveryStatus**](PageDeliveryStatus)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetSentDeliveryStatusesBySentId

> PageDeliveryStatus GetSentDeliveryStatusesBySentId(ctx, sentId, optional)



Get all sent email delivery statuses

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sentId** | [**string**]()|  | 
 **optional** | ***GetSentDeliveryStatusesBySentIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSentDeliveryStatusesBySentIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Optional page index in delivery status list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in delivery status list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **since** | **optional.Time**| Filter by created at after the given timestamp | 
 **before** | **optional.Time**| Filter by created at before the given timestamp | 

### Return type

[**PageDeliveryStatus**](PageDeliveryStatus)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetSentEmail

> SentEmailDto GetSentEmail(ctx, id)

Get sent email receipt

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()|  | 

### Return type

[**SentEmailDto**](SentEmailDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetSentEmailHTMLContent

> string GetSentEmailHTMLContent(ctx, id)

Get sent email HTML content

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()|  | 

### Return type

**string**

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetSentEmailPreviewURLs

> EmailPreviewUrls GetSentEmailPreviewURLs(ctx, id)

Get sent email URL for viewing in browser or downloading

Get a list of URLs for sent email content as text/html or raw SMTP message for viewing the message in a browser.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()|  | 

### Return type

[**EmailPreviewUrls**](EmailPreviewUrls)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetSentEmailTrackingPixels

> PageTrackingPixelProjection GetSentEmailTrackingPixels(ctx, id, optional)



Get all tracking pixels for a sent email in paginated form

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()|  | 
 **optional** | ***GetSentEmailTrackingPixelsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSentEmailTrackingPixelsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Optional page index in sent email tracking pixel list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in sent email tracking pixel list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **searchFilter** | **optional.String**| Optional search filter | 
 **since** | **optional.Time**| Filter by created at after the given timestamp | 
 **before** | **optional.Time**| Filter by created at before the given timestamp | 

### Return type

[**PageTrackingPixelProjection**](PageTrackingPixelProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetSentEmails

> PageSentEmailProjection GetSentEmails(ctx, optional)

Get all sent emails in paginated form

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetSentEmailsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSentEmailsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inboxId** | [**optional.Interface of string**]()| Optional inboxId to filter sender of sent emails by | 
 **page** | **optional.Int32**| Optional page index in inbox sent email list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in inbox sent email list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **searchFilter** | **optional.String**| Optional search filter | 
 **since** | **optional.Time**| Filter by created at after the given timestamp | 
 **before** | **optional.Time**| Filter by created at before the given timestamp | 

### Return type

[**PageSentEmailProjection**](PageSentEmailProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetSentEmailsWithQueueResults

> PageSentEmailWithQueueProjection GetSentEmailsWithQueueResults(ctx, optional)

Get results of email sent with queues in paginated form

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetSentEmailsWithQueueResultsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSentEmailsWithQueueResultsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Optional page index in inbox sent email list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in inbox sent email list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **since** | **optional.Time**| Filter by created at after the given timestamp | 
 **before** | **optional.Time**| Filter by created at before the given timestamp | 

### Return type

[**PageSentEmailWithQueueProjection**](PageSentEmailWithQueueProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetSentOrganizationEmails

> PageSentEmailProjection GetSentOrganizationEmails(ctx, optional)



Get all sent organization emails in paginated form

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetSentOrganizationEmailsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSentOrganizationEmailsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inboxId** | [**optional.Interface of string**]()| Optional inboxId to filter sender of sent emails by | 
 **page** | **optional.Int32**| Optional page index in sent email list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in sent email list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **searchFilter** | **optional.String**| Optional search filter | 
 **since** | **optional.Time**| Filter by created at after the given timestamp | 
 **before** | **optional.Time**| Filter by created at before the given timestamp | 

### Return type

[**PageSentEmailProjection**](PageSentEmailProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

