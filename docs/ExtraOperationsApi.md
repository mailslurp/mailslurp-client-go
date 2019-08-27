# \ExtraOperationsApi

All URIs are relative to *https://api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkCreateInboxesUsingPOST**](ExtraOperationsApi.md#BulkCreateInboxesUsingPOST) | **Post** /bulk/inboxes | Bulk create Inboxes (email addresses)
[**BulkDeleteInboxesUsingDELETE**](ExtraOperationsApi.md#BulkDeleteInboxesUsingDELETE) | **Delete** /bulk/inboxes | Bulk Delete Inboxes
[**BulkSendEmailsUsingPOST**](ExtraOperationsApi.md#BulkSendEmailsUsingPOST) | **Post** /bulk/send | Bulk Send Emails
[**CreateInboxUsingPOST**](ExtraOperationsApi.md#CreateInboxUsingPOST) | **Post** /inboxes | Create an Inbox (email address)
[**CreateInboxWebhookUsingPOST**](ExtraOperationsApi.md#CreateInboxWebhookUsingPOST) | **Post** /inboxes/{inboxId}/webhooks | Attach a webhook URL to an inbox
[**DeleteEmailUsingDELETE**](ExtraOperationsApi.md#DeleteEmailUsingDELETE) | **Delete** /emails/{emailId} | Delete Email
[**DeleteInboxUsingDELETE**](ExtraOperationsApi.md#DeleteInboxUsingDELETE) | **Delete** /inboxes/{inboxId} | Delete Inbox
[**GetEmailAttachmentUsingGET**](ExtraOperationsApi.md#GetEmailAttachmentUsingGET) | **Get** /emails/{emailId}/attachments/{attachmentId} | Get email attachment
[**GetEmailUsingGET**](ExtraOperationsApi.md#GetEmailUsingGET) | **Get** /emails/{emailId} | Get Email Content
[**GetEmailsUsingGET**](ExtraOperationsApi.md#GetEmailsUsingGET) | **Get** /inboxes/{inboxId}/emails | List an Inbox&#39;s Emails
[**GetInboxUsingGET**](ExtraOperationsApi.md#GetInboxUsingGET) | **Get** /inboxes/{inboxId} | Get Inbox
[**GetInboxWebhooksUsingDELETE**](ExtraOperationsApi.md#GetInboxWebhooksUsingDELETE) | **Delete** /inboxes/{inboxId}/webhooks/{webhookId} | Delete and disable a webhook for an inbox
[**GetInboxWebhooksUsingGET**](ExtraOperationsApi.md#GetInboxWebhooksUsingGET) | **Get** /inboxes/{inboxId}/webhooks | Get all webhooks for an inbox
[**GetInboxesUsingGET**](ExtraOperationsApi.md#GetInboxesUsingGET) | **Get** /inboxes | List Inboxes
[**GetRawEmailUsingGET**](ExtraOperationsApi.md#GetRawEmailUsingGET) | **Get** /emails/{emailId}/raw | Get Raw Email Content
[**SendEmailUsingPOST**](ExtraOperationsApi.md#SendEmailUsingPOST) | **Post** /inboxes/{inboxId} | Send Email


# **BulkCreateInboxesUsingPOST**
> []Inbox BulkCreateInboxesUsingPOST(ctx, count)
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

# **BulkDeleteInboxesUsingDELETE**
> BulkDeleteInboxesUsingDELETE(ctx, requestBody)
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

# **BulkSendEmailsUsingPOST**
> BulkSendEmailsUsingPOST(ctx, bulkSendEmailOptions)
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

# **CreateInboxUsingPOST**
> Inbox CreateInboxUsingPOST(ctx, )
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

# **CreateInboxWebhookUsingPOST**
> Webhook CreateInboxWebhookUsingPOST(ctx, inboxId, createWebhookOptions)
Attach a webhook URL to an inbox

Get notified whenever an inbox receives an email via a webhook URL. An emailID will be posted to this URL every time an email is received for this inbox. The URL must be publicly reachable by the MailSlurp server. You can provide basicAuth values if you wish to secure this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inboxId** | [**string**](.md)| inboxId | 
  **createWebhookOptions** | [**CreateWebhookOptions**](CreateWebhookOptions.md)| options | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEmailUsingDELETE**
> DeleteEmailUsingDELETE(ctx, emailId)
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

# **DeleteInboxUsingDELETE**
> DeleteInboxUsingDELETE(ctx, inboxId)
Delete Inbox

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

# **GetEmailAttachmentUsingGET**
> GetEmailAttachmentUsingGET(ctx, attachmentId, emailId)
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

# **GetEmailUsingGET**
> Email GetEmailUsingGET(ctx, emailId)
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

# **GetEmailsUsingGET**
> []EmailPreview GetEmailsUsingGET(ctx, inboxId, optional)
List an Inbox's Emails

List emails that an inbox has received. Only emails that are sent to the inbox's email address will appear in the inbox. It may take several seconds for any email you send to an inbox's email address to appear in the inbox. To make this endpoint wait for a minimum number of emails use the `minCount` parameter. The server will retry the inbox database until the `minCount` is satisfied or the `retryTimeout` is reached

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inboxId** | [**string**](.md)| Id of inbox that emails belongs to | 
 **optional** | ***GetEmailsUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetEmailsUsingGETOpts struct

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

# **GetInboxUsingGET**
> Inbox GetInboxUsingGET(ctx, inboxId)
Get Inbox

Returns an inbox's properties, including its email address and ID

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

# **GetInboxWebhooksUsingDELETE**
> GetInboxWebhooksUsingDELETE(ctx, inboxId, webhookId)
Delete and disable a webhook for an inbox

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

# **GetInboxWebhooksUsingGET**
> []Webhook GetInboxWebhooksUsingGET(ctx, inboxId)
Get all webhooks for an inbox

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

# **GetInboxesUsingGET**
> []Inbox GetInboxesUsingGET(ctx, )
List Inboxes

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

# **GetRawEmailUsingGET**
> string GetRawEmailUsingGET(ctx, emailId)
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

# **SendEmailUsingPOST**
> SendEmailUsingPOST(ctx, inboxId, sendEmailOptions)
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

