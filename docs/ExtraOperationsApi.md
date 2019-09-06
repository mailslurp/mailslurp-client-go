# \ExtraOperationsApi

All URIs are relative to *https://api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkCreateInboxes**](ExtraOperationsApi.md#BulkCreateInboxes) | **Post** /bulk/inboxes | Bulk create Inboxes (email addresses)
[**BulkDeleteInboxes**](ExtraOperationsApi.md#BulkDeleteInboxes) | **Delete** /bulk/inboxes | Bulk Delete Inboxes
[**BulkSendEmails**](ExtraOperationsApi.md#BulkSendEmails) | **Post** /bulk/send | Bulk Send Emails
[**CreateInbox**](ExtraOperationsApi.md#CreateInbox) | **Post** /inboxes | Create an Inbox (email address)
[**CreateWebhook**](ExtraOperationsApi.md#CreateWebhook) | **Post** /inboxes/{inboxId}/webhooks | Attach a WebHook URL to an inbox
[**DeleteEmail1**](ExtraOperationsApi.md#DeleteEmail1) | **Delete** /emails/{emailId} | Delete Email
[**DeleteInbox**](ExtraOperationsApi.md#DeleteInbox) | **Delete** /inboxes/{inboxId} | Delete Inbox / Email Address
[**DeleteWebhook**](ExtraOperationsApi.md#DeleteWebhook) | **Delete** /inboxes/{inboxId}/webhooks/{webhookId} | Delete and disable a WebHook for an Inbox
[**DownloadAttachment**](ExtraOperationsApi.md#DownloadAttachment) | **Get** /emails/{emailId}/attachments/{attachmentId} | Get email attachment
[**GetEmail**](ExtraOperationsApi.md#GetEmail) | **Get** /emails/{emailId} | Get Email Content
[**GetEmails**](ExtraOperationsApi.md#GetEmails) | **Get** /inboxes/{inboxId}/emails | List Emails in an Inbox / EmailAddress
[**GetInbox**](ExtraOperationsApi.md#GetInbox) | **Get** /inboxes/{inboxId} | Get Inbox / EmailAddress
[**GetInboxes**](ExtraOperationsApi.md#GetInboxes) | **Get** /inboxes | List Inboxes / Email Addresses
[**GetRawEmailContents**](ExtraOperationsApi.md#GetRawEmailContents) | **Get** /emails/{emailId}/raw | Get Raw Email Content
[**GetWebhooks**](ExtraOperationsApi.md#GetWebhooks) | **Get** /inboxes/{inboxId}/webhooks | Get all WebHooks for an Inbox
[**SendEmail**](ExtraOperationsApi.md#SendEmail) | **Post** /inboxes/{inboxId} | Send Email


# **BulkCreateInboxes**
> []Inbox BulkCreateInboxes(ctx, count)
Bulk create Inboxes (email addresses)

Enterprise Plan Required

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **count** | **int32**| Number of inboxes to be created in bulk | 

### Return type

[**[]Inbox**](Inbox.md)

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BulkDeleteInboxes**
> BulkDeleteInboxes(ctx, requestBody)
Bulk Delete Inboxes

Enterprise Plan Required

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestBody** | [**[]string**](array.md)| ids | 

### Return type

 (empty response body)

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BulkSendEmails**
> BulkSendEmails(ctx, bulkSendEmailOptions)
Bulk Send Emails

Enterprise Plan Required

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bulkSendEmailOptions** | [**BulkSendEmailOptions**](BulkSendEmailOptions.md)| bulkSendEmailOptions | 

### Return type

 (empty response body)

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateInbox**
> Inbox CreateInbox(ctx, )
Create an Inbox (email address)

Create a new inbox and ephemeral email address to send and receive from. This is a necessary step before sending or receiving emails. The response contains the inbox's ID and its associated email address. It is recommended that you create a new inbox during each test method so that it is unique and empty

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Inbox**](Inbox.md)

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateWebhook**
> Webhook CreateWebhook(ctx, inboxId, createWebhookOptions)
Attach a WebHook URL to an inbox

Get notified whenever an inbox receives an email via a WebHook URL. An emailID will be posted to this URL every time an email is received for this inbox. The URL must be publicly reachable by the MailSlurp server. You can provide basicAuth values if you wish to secure this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inboxId** | [**string**](.md)| inboxId | 
  **createWebhookOptions** | [**CreateWebhookOptions**](CreateWebhookOptions.md)| webhookOptions | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEmail1**
> DeleteEmail1(ctx, emailId)
Delete Email

Deletes an email and removes it from the inbox

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **emailId** | [**string**](.md)| emailId | 

### Return type

 (empty response body)

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteInbox**
> DeleteInbox(ctx, inboxId)
Delete Inbox / Email Address

Permanently delete an inbox and associated email address

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteWebhook**
> DeleteWebhook(ctx, inboxId, webhookId)
Delete and disable a WebHook for an Inbox

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inboxId** | [**string**](.md)| inboxId | 
  **webhookId** | [**string**](.md)| webhookId | 

### Return type

 (empty response body)

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DownloadAttachment**
> DownloadAttachment(ctx, attachmentId, emailId)
Get email attachment

Returns the specified attachment for a given email as a byte stream (file download). Get the attachmentId from the email response. Requires enterprise account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **attachmentId** | **string**| attachmentId | 
  **emailId** | [**string**](.md)| emailId | 

### Return type

 (empty response body)

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEmail**
> Email GetEmail(ctx, emailId)
Get Email Content

Returns a email summary object with headers and content. To retrieve the raw unparsed email use the getRawMessage endpoint

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **emailId** | [**string**](.md)| emailId | 

### Return type

[**Email**](Email.md)

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEmails**
> []EmailPreview GetEmails(ctx, inboxId, optional)
List Emails in an Inbox / EmailAddress

List emails that an inbox has received. Only emails that are sent to the inbox's email address will appear in the inbox. It may take several seconds for any email you send to an inbox's email address to appear in the inbox. To make this endpoint wait for a minimum number of emails use the `minCount` parameter. The server will retry the inbox database until the `minCount` is satisfied or the `retryTimeout` is reached

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inboxId** | [**string**](.md)| Id of inbox that emails belongs to | 
 **optional** | ***GetEmailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetEmailsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| Limit the result set, ordered by descending received date time | 
 **minCount** | **optional.Int64**| Minimum acceptable email count. Will cause request to hang (and retry) until minCount is satisfied or retryTimeout is reached. | 
 **retryTimeout** | **optional.Int64**| Maximum milliseconds to spend retrying inbox database until minCount emails are returned | 
 **since** | **optional.Time**| Exclude emails received before this ISO 8601 date time | 

### Return type

[**[]EmailPreview**](EmailPreview.md)

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInbox**
> Inbox GetInbox(ctx, inboxId)
Get Inbox / EmailAddress

Returns an inbox's properties, including its email address and ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inboxId** | [**string**](.md)| inboxId | 

### Return type

[**Inbox**](Inbox.md)

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInboxes**
> []Inbox GetInboxes(ctx, )
List Inboxes / Email Addresses

List the inboxes you have created

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Inbox**](Inbox.md)

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRawEmailContents**
> string GetRawEmailContents(ctx, emailId)
Get Raw Email Content

Returns a raw, unparsed and unprocessed email

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **emailId** | [**string**](.md)| emailId | 

### Return type

**string**

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWebhooks**
> []Webhook GetWebhooks(ctx, inboxId)
Get all WebHooks for an Inbox

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inboxId** | [**string**](.md)| inboxId | 

### Return type

[**[]Webhook**](Webhook.md)

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendEmail**
> SendEmail(ctx, inboxId, sendEmailOptions)
Send Email

Send an email from the inbox's email address. Specify the email recipients and contents in the request body. See the `SendEmailOptions` for more information. Note the `inboxId` refers to the inbox's id NOT its email address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inboxId** | [**string**](.md)| inboxId | 
  **sendEmailOptions** | [**SendEmailOptions**](SendEmailOptions.md)| sendEmailOptions | 

### Return type

 (empty response body)

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

