# MailSlurp\CommonActionsControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewEmailAddress**](CommonActionsControllerApi#CreateNewEmailAddress) | **Post** /newEmailAddress | Create new random inbox
[**CreateRandomInbox**](CommonActionsControllerApi#CreateRandomInbox) | **Post** /createInbox | Create new random inbox
[**DeleteEmailAddress**](CommonActionsControllerApi#DeleteEmailAddress) | **Delete** /deleteEmailAddress | Delete inbox email address by inbox id
[**EmptyInbox**](CommonActionsControllerApi#EmptyInbox) | **Delete** /emptyInbox | Delete all emails in an inbox
[**SendEmailQuery**](CommonActionsControllerApi#SendEmailQuery) | **Post** /sendEmailQuery | Send an email using query parameters
[**SendEmailSimple**](CommonActionsControllerApi#SendEmailSimple) | **Post** /sendEmail | Send an email



## CreateNewEmailAddress

> InboxDto CreateNewEmailAddress(ctx, optional)

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
 **allowTeamAccess** | **optional.Bool**|  | 
 **useDomainPool** | **optional.Bool**|  | 
 **expiresAt** | **optional.Time**|  | 
 **expiresIn** | **optional.Int64**|  | 
 **emailAddress** | **optional.String**|  | 
 **inboxType** | **optional.String**|  | 
 **description** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **tags** | [**optional.Interface of []string**](string)|  | 
 **favourite** | **optional.Bool**|  | 
 **virtualInbox** | **optional.Bool**|  | 
 **useShortAddress** | **optional.Bool**|  | 
 **domainName** | **optional.String**|  | 
 **domainId** | [**optional.Interface of string**]()|  | 
 **prefix** | **optional.String**|  | 

### Return type

[**InboxDto**](InboxDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CreateRandomInbox

> InboxDto CreateRandomInbox(ctx, optional)

Create new random inbox

Returns an Inbox with an `id` and an `emailAddress`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateRandomInboxOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateRandomInboxOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **allowTeamAccess** | **optional.Bool**|  | 
 **useDomainPool** | **optional.Bool**|  | 
 **expiresAt** | **optional.Time**|  | 
 **expiresIn** | **optional.Int64**|  | 
 **emailAddress** | **optional.String**|  | 
 **inboxType** | **optional.String**|  | 
 **description** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **tags** | [**optional.Interface of []string**](string)|  | 
 **favourite** | **optional.Bool**|  | 
 **virtualInbox** | **optional.Bool**|  | 
 **useShortAddress** | **optional.Bool**|  | 
 **domainName** | **optional.String**|  | 
 **domainId** | [**optional.Interface of string**]()|  | 
 **prefix** | **optional.String**|  | 

### Return type

[**InboxDto**](InboxDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DeleteEmailAddress

> DeleteEmailAddress(ctx, inboxId)

Delete inbox email address by inbox id

Deletes inbox email address

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()|  | 

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


## EmptyInbox

> EmptyInbox(ctx, inboxId)

Delete all emails in an inbox

Deletes all emails

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()|  | 

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


## SendEmailQuery

> SendEmailQuery(ctx, to, optional)

Send an email using query parameters

If no senderId or inboxId provided a random email address will be used to send from. Ensure your parameters are URL encoded.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**to** | **string**| Email address to send to | 
 **optional** | ***SendEmailQueryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SendEmailQueryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **senderId** | [**optional.Interface of string**]()| ID of inbox to send from. If null an inbox will be created for sending | 
 **body** | **optional.String**| Body of the email message. Supports HTML | 
 **subject** | **optional.String**| Subject line of the email | 

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


## SendEmailSimple

> SendEmailSimple(ctx, simpleSendEmailOptions)

Send an email

If no senderId or inboxId provided a random email address will be used to send from.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simpleSendEmailOptions** | [**SimpleSendEmailOptions**](SimpleSendEmailOptions)|  | 

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

