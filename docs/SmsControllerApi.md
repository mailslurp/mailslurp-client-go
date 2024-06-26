# MailSlurp\SmsControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSmsMessage**](SmsControllerApi#DeleteSmsMessage) | **Delete** /sms/{smsId} | Delete SMS message.
[**DeleteSmsMessages**](SmsControllerApi#DeleteSmsMessages) | **Delete** /sms | Delete all SMS messages
[**GetReplyForSmsMessage**](SmsControllerApi#GetReplyForSmsMessage) | **Get** /sms/{smsId}/reply | Get reply for an SMS message
[**GetSmsCount**](SmsControllerApi#GetSmsCount) | **Get** /sms/count | Get SMS count
[**GetSmsMessage**](SmsControllerApi#GetSmsMessage) | **Get** /sms/{smsId} | Get SMS content including body. Expects SMS to exist by ID. For SMS that may not have arrived yet use the WaitForController.
[**GetSmsMessagesPaginated**](SmsControllerApi#GetSmsMessagesPaginated) | **Get** /sms | Get all SMS messages in all phone numbers in paginated form. .
[**GetUnreadSmsCount**](SmsControllerApi#GetUnreadSmsCount) | **Get** /sms/unreadCount | Get unread SMS count
[**ReplyToSmsMessage**](SmsControllerApi#ReplyToSmsMessage) | **Post** /sms/{smsId}/reply | Send a reply to a received SMS message. Replies are sent from the receiving number.



## DeleteSmsMessage

> DeleteSmsMessage(ctx, smsId)

Delete SMS message.

Delete an SMS message

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**smsId** | [**string**]()|  | 

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


## DeleteSmsMessages

> DeleteSmsMessages(ctx, optional)

Delete all SMS messages

Delete all SMS messages or all messages for a given phone number

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteSmsMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteSmsMessagesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phoneNumberId** | [**optional.Interface of string**]()|  | 

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


## GetReplyForSmsMessage

> ReplyForSms GetReplyForSmsMessage(ctx, smsId)

Get reply for an SMS message

Get reply for an SMS message.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**smsId** | [**string**]()|  | 

### Return type

[**ReplyForSms**](ReplyForSms)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetSmsCount

> CountDto GetSmsCount(ctx, )

Get SMS count

Get number of SMS

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**CountDto**](CountDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetSmsMessage

> SmsDto GetSmsMessage(ctx, smsId)

Get SMS content including body. Expects SMS to exist by ID. For SMS that may not have arrived yet use the WaitForController.

Returns a SMS summary object with content.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**smsId** | [**string**]()|  | 

### Return type

[**SmsDto**](SmsDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetSmsMessagesPaginated

> PageSmsProjection GetSmsMessagesPaginated(ctx, optional)

Get all SMS messages in all phone numbers in paginated form. .

By default returns all SMS messages across all phone numbers sorted by ascending created at date. Responses are paginated. You can restrict results to a list of phone number IDs. You can also filter out read messages

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetSmsMessagesPaginatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSmsMessagesPaginatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phoneNumber** | [**optional.Interface of string**]()| Optional receiving phone number to filter SMS messages for | 
 **page** | **optional.Int32**| Optional page index in SMS list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in SMS list pagination. Maximum size is 100. Use page index and sort to page through larger results | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **unreadOnly** | **optional.Bool**| Optional filter for unread SMS only. All SMS are considered unread until they are viewed in the dashboard or requested directly | [default to false]
 **since** | **optional.Time**| Optional filter SMSs received after given date time | 
 **before** | **optional.Time**| Optional filter SMSs received before given date time | 

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


## GetUnreadSmsCount

> UnreadCount GetUnreadSmsCount(ctx, )

Get unread SMS count

Get number of SMS unread. Unread means has not been viewed in dashboard or returned in an email API response

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**UnreadCount**](UnreadCount)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## ReplyToSmsMessage

> SentSmsDto ReplyToSmsMessage(ctx, smsId, smsReplyOptions)

Send a reply to a received SMS message. Replies are sent from the receiving number.

Reply to an SMS message.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**smsId** | [**string**]()|  | 
**smsReplyOptions** | [**SmsReplyOptions**](SmsReplyOptions)|  | 

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

