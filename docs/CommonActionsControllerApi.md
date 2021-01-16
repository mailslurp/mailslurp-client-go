# MailSlurp\CommonActionsControllerApi

All URIs are relative to *https://api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewEmailAddress**](CommonActionsControllerApi.md#CreateNewEmailAddress) | **Post** /createInbox | Create new random inbox
[**CreateNewEmailAddress1**](CommonActionsControllerApi.md#CreateNewEmailAddress1) | **Post** /newEmailAddress | Create new random inbox
[**EmptyInbox**](CommonActionsControllerApi.md#EmptyInbox) | **Delete** /emptyInbox | Delete all emails in an inbox
[**SendEmailSimple**](CommonActionsControllerApi.md#SendEmailSimple) | **Post** /sendEmail | Send an email



## CreateNewEmailAddress

> Inbox CreateNewEmailAddress(ctx, optional)

Create new random inbox

Returns an Inbox with an `id` and an `emailAddress`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateNewEmailAddressOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateNewEmailAddressOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expiresAt** | **optional.Time**| expiresAt | 
 **expiresIn** | **optional.Int64**| expiresIn | 
 **useDomainPool** | **optional.Bool**| useDomainPool | 

### Return type

[**Inbox**](Inbox.md)

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNewEmailAddress1

> Inbox CreateNewEmailAddress1(ctx, optional)

Create new random inbox

Returns an Inbox with an `id` and an `emailAddress`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateNewEmailAddress1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateNewEmailAddress1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expiresAt** | **optional.Time**| expiresAt | 
 **expiresIn** | **optional.Int64**| expiresIn | 
 **useDomainPool** | **optional.Bool**| useDomainPool | 

### Return type

[**Inbox**](Inbox.md)

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmptyInbox

> EmptyInbox(ctx, inboxId)

Delete all emails in an inbox

Deletes all emails

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**](.md)| inboxId | 

### Return type

 (empty response body)

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendEmailSimple

> SendEmailSimple(ctx, emailOptions)

Send an email

If no senderId or inboxId provided a random email address will be used to send from.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailOptions** | [**SimpleSendEmailOptions**](SimpleSendEmailOptions.md)| emailOptions | 

### Return type

 (empty response body)

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

