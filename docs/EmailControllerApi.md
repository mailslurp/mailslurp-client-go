# MailSlurp\EmailControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyImapFlagOperation**](EmailControllerApi#ApplyImapFlagOperation) | **Post** /emails/{emailId}/imap-flag-operation | Set IMAP flags associated with a message. Only supports &#39;\\Seen&#39; flag.
[**CanSend**](EmailControllerApi#CanSend) | **Post** /emails/can-send | Check whether an email send would be accepted
[**CheckEmailAudit1**](EmailControllerApi#CheckEmailAudit1) | **Post** /emails/{emailId}/check-email-audit | Run aggregate email audit for a stored email
[**CheckEmailBody**](EmailControllerApi#CheckEmailBody) | **Post** /emails/{emailId}/check-email-body | Check email body for broken links, images, and spelling issues
[**CheckEmailBodyFeatureSupport**](EmailControllerApi#CheckEmailBodyFeatureSupport) | **Post** /emails/{emailId}/check-email-body-feature-support | Check client support for features used in a stored email body
[**CheckEmailClientSupport**](EmailControllerApi#CheckEmailClientSupport) | **Post** /emails/check-email-client-support | Check email-client support for a provided HTML body
[**CreateEmailAuditForEmail**](EmailControllerApi#CreateEmailAuditForEmail) | **Post** /emails/{emailId}/audit | Persist aggregate email audit for a stored email
[**DeleteAllEmails**](EmailControllerApi#DeleteAllEmails) | **Delete** /emails | Delete all emails in all inboxes.
[**DeleteEmail**](EmailControllerApi#DeleteEmail) | **Delete** /emails/{emailId} | Delete an email
[**DownloadAttachment**](EmailControllerApi#DownloadAttachment) | **Get** /emails/{emailId}/attachments/{attachmentId} | Get email attachment bytes. Returned as &#x60;octet-stream&#x60; with content type header. If you have trouble with byte responses try the &#x60;downloadAttachmentBase64&#x60; response endpoints and convert the base 64 encoded content to a file or string.
[**DownloadAttachmentBase64**](EmailControllerApi#DownloadAttachmentBase64) | **Get** /emails/{emailId}/attachments/{attachmentId}/base64 | Get email attachment as base64 encoded string as an alternative to binary responses. Decode the &#x60;base64FileContents&#x60; as a &#x60;utf-8&#x60; encoded string or array of bytes depending on the &#x60;contentType&#x60;.
[**DownloadBody**](EmailControllerApi#DownloadBody) | **Get** /emails/{emailId}/body | Get email body as string. Returned as &#x60;plain/text&#x60; with content type header.
[**DownloadBodyBytes**](EmailControllerApi#DownloadBodyBytes) | **Get** /emails/{emailId}/body-bytes | Get email body in bytes. Returned as &#x60;octet-stream&#x60; with content type header.
[**ForwardEmail**](EmailControllerApi#ForwardEmail) | **Post** /emails/{emailId}/forward | Forward email to recipients
[**GetAttachmentMetaData**](EmailControllerApi#GetAttachmentMetaData) | **Get** /emails/{emailId}/attachments/{attachmentId}/metadata | Get email attachment metadata. This is the &#x60;contentType&#x60; and &#x60;contentLength&#x60; of an attachment. To get the individual attachments  use the &#x60;downloadAttachment&#x60; methods.
[**GetEmail**](EmailControllerApi#GetEmail) | **Get** /emails/{emailId} | Get hydrated email (headers and body)
[**GetEmailAttachments**](EmailControllerApi#GetEmailAttachments) | **Get** /emails/{emailId}/attachments | List attachment metadata for an email
[**GetEmailCodes**](EmailControllerApi#GetEmailCodes) | **Post** /emails/{emailId}/codes | Extract verification codes from an email
[**GetEmailContentMatch**](EmailControllerApi#GetEmailContentMatch) | **Post** /emails/{emailId}/contentMatch | Run regex against hydrated email body and return matches
[**GetEmailContentPart**](EmailControllerApi#GetEmailContentPart) | **Get** /emails/{emailId}/contentPart | Get email content part by content type
[**GetEmailContentPartContent**](EmailControllerApi#GetEmailContentPartContent) | **Get** /emails/{emailId}/contentPart/raw | Get multipart content part as raw response
[**GetEmailCount**](EmailControllerApi#GetEmailCount) | **Get** /emails/emails/count | Get email count
[**GetEmailHTML**](EmailControllerApi#GetEmailHTML) | **Get** /emails/{emailId}/html | Get hydrated email HTML for browser rendering
[**GetEmailHTMLJson**](EmailControllerApi#GetEmailHTMLJson) | **Get** /emails/{emailId}/html/json | Get hydrated email HTML wrapped in JSON
[**GetEmailHTMLQuery**](EmailControllerApi#GetEmailHTMLQuery) | **Get** /emails/{emailId}/htmlQuery | Query hydrated HTML body and return matching text lines
[**GetEmailLinks**](EmailControllerApi#GetEmailLinks) | **Get** /emails/{emailId}/links | Extract links from an email HTML body
[**GetEmailPreviewURLs**](EmailControllerApi#GetEmailPreviewURLs) | **Get** /emails/{emailId}/urls | Get email URLs for viewing in browser or downloading
[**GetEmailScreenshotAsBase64**](EmailControllerApi#GetEmailScreenshotAsBase64) | **Post** /emails/{emailId}/screenshot/base64 | Take a screenshot of an email in a browser and return base64 encoded string
[**GetEmailScreenshotAsBinary**](EmailControllerApi#GetEmailScreenshotAsBinary) | **Post** /emails/{emailId}/screenshot/binary | Take a screenshot of an email in a browser
[**GetEmailSignature**](EmailControllerApi#GetEmailSignature) | **Get** /emails/{emailId}/signature | Extract signature from an inbound email
[**GetEmailSummary**](EmailControllerApi#GetEmailSummary) | **Get** /emails/{emailId}/summary | Get email summary (headers/metadata only)
[**GetEmailTextLines**](EmailControllerApi#GetEmailTextLines) | **Get** /emails/{emailId}/textLines | Extract normalized text lines from email body
[**GetEmailThread**](EmailControllerApi#GetEmailThread) | **Get** /emails/threads/{threadId} | Get email thread metadata by thread ID
[**GetEmailThreadItems**](EmailControllerApi#GetEmailThreadItems) | **Get** /emails/threads/{threadId}/items | Get messages in a specific email thread
[**GetEmailThreads**](EmailControllerApi#GetEmailThreads) | **Get** /emails/threads | List email threads in paginated form
[**GetEmailsOffsetPaginated**](EmailControllerApi#GetEmailsOffsetPaginated) | **Get** /emails/offset-paginated | Get all emails in all inboxes in paginated form. Email API list all.
[**GetEmailsPaginated**](EmailControllerApi#GetEmailsPaginated) | **Get** /emails | Get all emails in all inboxes in paginated form. Email API list all.
[**GetGravatarUrlForEmailAddress**](EmailControllerApi#GetGravatarUrlForEmailAddress) | **Get** /emails/gravatarFor | Get Gravatar URL for an email address
[**GetLatestEmail**](EmailControllerApi#GetLatestEmail) | **Get** /emails/latest | Get latest email in all inboxes. Most recently received.
[**GetLatestEmailInInbox1**](EmailControllerApi#GetLatestEmailInInbox1) | **Get** /emails/latestIn | Get latest email in an inbox. Use &#x60;WaitForController&#x60; to get emails that may not have arrived yet.
[**GetOrganizationEmailsPaginated**](EmailControllerApi#GetOrganizationEmailsPaginated) | **Get** /emails/organization | List organization-visible emails
[**GetRawEmailContents**](EmailControllerApi#GetRawEmailContents) | **Get** /emails/{emailId}/raw | Get raw email string. Returns unparsed raw SMTP message with headers and body.
[**GetRawEmailJson**](EmailControllerApi#GetRawEmailJson) | **Get** /emails/{emailId}/raw/json | Get raw email in JSON. Unparsed SMTP message in JSON wrapper format.
[**GetUnreadEmailCount**](EmailControllerApi#GetUnreadEmailCount) | **Get** /emails/unreadCount | Get unread email count
[**MarkAllAsRead**](EmailControllerApi#MarkAllAsRead) | **Patch** /emails/read | Mark all emails as read or unread
[**MarkAsRead**](EmailControllerApi#MarkAsRead) | **Patch** /emails/{emailId}/read | Mark an email as read or unread
[**ReplyToEmail**](EmailControllerApi#ReplyToEmail) | **Put** /emails/{emailId} | Reply to an email
[**SearchEmails**](EmailControllerApi#SearchEmails) | **Post** /emails/search | Get all emails by search criteria. Return in paginated form.
[**SendEmailSourceOptional**](EmailControllerApi#SendEmailSourceOptional) | **Post** /emails | Send email
[**SetEmailFavourited**](EmailControllerApi#SetEmailFavourited) | **Put** /emails/{emailId}/favourite | Set email favourited state
[**ValidateEmail**](EmailControllerApi#ValidateEmail) | **Post** /emails/{emailId}/validate | Validate email HTML contents



## ApplyImapFlagOperation

> EmailPreview ApplyImapFlagOperation(ctx, emailId, imapFlagOperationOptions)

Set IMAP flags associated with a message. Only supports '\\Seen' flag.

Applies RFC3501 IMAP flag operations on a message. Current implementation supports read/unread semantics via the `\\\\Seen` flag only.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()|  | 
**imapFlagOperationOptions** | [**ImapFlagOperationOptions**](ImapFlagOperationOptions)|  | 

### Return type

[**EmailPreview**](EmailPreview)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CanSend

> CanSendEmailResults CanSend(ctx, inboxId, sendEmailOptions)

Check whether an email send would be accepted

Validates sender/inbox context and recipient eligibility before attempting a send. Useful for preflight checks in UI or test workflows.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()| ID of the inbox you want to send the email from | 
**sendEmailOptions** | [**SendEmailOptions**](SendEmailOptions)|  | 

### Return type

[**CanSendEmailResults**](CanSendEmailResults)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CheckEmailAudit1

> EmailAuditAnalysisResult CheckEmailAudit1(ctx, emailId)

Run aggregate email audit for a stored email

Runs the same message-level audit bundle used by the email audit dashboard in one request. Combines content checks, HTML validation, compatibility analysis, and reputation verdict rollup when available.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()|  | 

### Return type

[**EmailAuditAnalysisResult**](EmailAuditAnalysisResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CheckEmailBody

> CheckEmailBodyResults CheckEmailBody(ctx, emailId)

Check email body for broken links, images, and spelling issues

Runs content quality checks against hydrated email body content. This endpoint performs outbound HTTP checks for linked resources, so avoid use with sensitive or stateful URLs.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()|  | 

### Return type

[**CheckEmailBodyResults**](CheckEmailBodyResults)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CheckEmailBodyFeatureSupport

> CheckEmailBodyFeatureSupportResults CheckEmailBodyFeatureSupport(ctx, emailId)

Check client support for features used in a stored email body

Detects HTML/CSS features in the target email body and reports compatibility across major email clients and devices.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()|  | 

### Return type

[**CheckEmailBodyFeatureSupportResults**](CheckEmailBodyFeatureSupportResults)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CheckEmailClientSupport

> CheckEmailClientSupportResults CheckEmailClientSupport(ctx, checkEmailClientSupportOptions)

Check email-client support for a provided HTML body

Evaluates HTML/CSS features in the supplied body and reports support coverage across major email clients and platforms.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkEmailClientSupportOptions** | [**CheckEmailClientSupportOptions**](CheckEmailClientSupportOptions)|  | 

### Return type

[**CheckEmailClientSupportResults**](CheckEmailClientSupportResults)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CreateEmailAuditForEmail

> EmailAuditDto CreateEmailAuditForEmail(ctx, emailId)

Persist aggregate email audit for a stored email

Runs the aggregate audit bundle for the target email and stores the resulting audit record for later review and history tracking.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()|  | 

### Return type

[**EmailAuditDto**](EmailAuditDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DeleteAllEmails

> DeleteAllEmails(ctx, )

Delete all emails in all inboxes.

Deletes all emails for the authenticated account context. This operation is destructive and cannot be undone.

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


## DeleteEmail

> DeleteEmail(ctx, emailId)

Delete an email

Deletes a single email from account scope. Operation is destructive and not reversible.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()| ID of email to delete | 

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


## DownloadAttachment

> string DownloadAttachment(ctx, emailId, attachmentId, optional)

Get email attachment bytes. Returned as `octet-stream` with content type header. If you have trouble with byte responses try the `downloadAttachmentBase64` response endpoints and convert the base 64 encoded content to a file or string.

Returns attachment bytes by attachment ID. Use attachment IDs from email payloads or attachment listing endpoints.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()| ID of email | 
**attachmentId** | **string**| ID of attachment | 
 **optional** | ***DownloadAttachmentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DownloadAttachmentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiKey** | **optional.String**| Can pass apiKey in url for this request if you wish to download the file in a browser. Content type will be set to original content type of the attachment file. This is so that browsers can download the file correctly. | 

### Return type

**string**

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DownloadAttachmentBase64

> DownloadAttachmentDto DownloadAttachmentBase64(ctx, emailId, attachmentId)

Get email attachment as base64 encoded string as an alternative to binary responses. Decode the `base64FileContents` as a `utf-8` encoded string or array of bytes depending on the `contentType`.

Returns attachment payload as base64 in JSON. Useful for clients that cannot reliably consume binary streaming responses.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()| ID of email | 
**attachmentId** | **string**| ID of attachment | 

### Return type

[**DownloadAttachmentDto**](DownloadAttachmentDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DownloadBody

> string DownloadBody(ctx, emailId)

Get email body as string. Returned as `plain/text` with content type header.

Returns hydrated email body text as a string with content type aligned to the underlying body format.

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
- **Accept**: text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DownloadBodyBytes

> string DownloadBodyBytes(ctx, emailId)

Get email body in bytes. Returned as `octet-stream` with content type header.

Streams hydrated email body bytes with content type derived from the message body format.

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
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## ForwardEmail

> SentEmailDto ForwardEmail(ctx, emailId, forwardEmailOptions)

Forward email to recipients

Forwards an existing email to new recipients. Uses the owning inbox context unless overridden by allowed sender options.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()| ID of email | 
**forwardEmailOptions** | [**ForwardEmailOptions**](ForwardEmailOptions)|  | 

### Return type

[**SentEmailDto**](SentEmailDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetAttachmentMetaData

> AttachmentMetaData GetAttachmentMetaData(ctx, emailId, attachmentId)

Get email attachment metadata. This is the `contentType` and `contentLength` of an attachment. To get the individual attachments  use the `downloadAttachment` methods.

Returns metadata for a specific attachment ID (name, content type, and size fields).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()| ID of email | 
**attachmentId** | **string**| ID of attachment | 

### Return type

[**AttachmentMetaData**](AttachmentMetaData)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetEmail

> Email GetEmail(ctx, emailId)

Get hydrated email (headers and body)

Returns parsed email content including headers and body fields. Accessing hydrated content may mark the email as read depending on read-state rules.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()|  | 

### Return type

[**Email**](Email)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/xml;charset=UTF-8, application/json; charset=UTF-8, application/xml; charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetEmailAttachments

> []AttachmentMetaData GetEmailAttachments(ctx, emailId)

List attachment metadata for an email

Returns metadata for all attachment IDs associated with the email (name, content type, size, and IDs).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()| ID of email | 

### Return type

[**[]AttachmentMetaData**](AttachmentMetaData)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetEmailCodes

> ExtractCodesResult GetEmailCodes(ctx, emailId, optional)

Extract verification codes from an email

Extracts one-time passcodes and similar tokens from email content. Supports deterministic extraction now with method/fallback flags (`AUTO`, `PATTERN`, `LLM`, `OCR`, `OCR_THEN_LLM`) for QA and future advanced pipelines.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()| ID of email to extract codes from | 
 **optional** | ***GetEmailCodesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEmailCodesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **extractCodesOptions** | [**optional.Interface of ExtractCodesOptions**](ExtractCodesOptions)|  | 

### Return type

[**ExtractCodesResult**](ExtractCodesResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetEmailContentMatch

> EmailContentMatchResult GetEmailContentMatch(ctx, emailId, contentMatchOptions)

Run regex against hydrated email body and return matches

Executes a Java regex pattern over hydrated email body text and returns the full match plus capture groups. Pattern syntax follows Java `Pattern` rules.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()| ID of email to match against | 
**contentMatchOptions** | [**ContentMatchOptions**](ContentMatchOptions)|  | 

### Return type

[**EmailContentMatchResult**](EmailContentMatchResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetEmailContentPart

> EmailContentPartResult GetEmailContentPart(ctx, emailId, contentType, optional)

Get email content part by content type

Extracts one MIME body part by `contentType` and optional `index`, returned in a structured DTO with metadata.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()| ID of email to match against | 
**contentType** | **string**| Content type | 
 **optional** | ***GetEmailContentPartOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEmailContentPartOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **strict** | **optional.Bool**| Strict content type match | 
 **index** | **optional.Int32**| Index of content type part if multiple | 

### Return type

[**EmailContentPartResult**](EmailContentPartResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/xml;charset=UTF-8, application/json; charset=UTF-8, application/xml; charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetEmailContentPartContent

> string GetEmailContentPartContent(ctx, emailId, contentType, optional)

Get multipart content part as raw response

Extracts one MIME body part by `contentType` and optional `index`, and returns raw content with matching response content type when valid.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()| ID of email to match against | 
**contentType** | **string**| Content type | 
 **optional** | ***GetEmailContentPartContentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEmailContentPartContentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **strict** | **optional.Bool**| Strict content type match | 
 **index** | **optional.Int32**| Index of content type part if multiple. Starts from 0 and applies to the result list after selecting for your content type. Content type parts are sorted by order found in original MIME message. | 

### Return type

**string**

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetEmailCount

> CountDto GetEmailCount(ctx, optional)

Get email count

Returns total email count for the authenticated user, or count scoped to a specific inbox when `inboxId` is provided.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetEmailCountOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEmailCountOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inboxId** | [**optional.Interface of string**]()|  | 

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


## GetEmailHTML

> string GetEmailHTML(ctx, emailId, optional)

Get hydrated email HTML for browser rendering

Returns hydrated HTML body directly with `text/html` content type. Supports temporary access/browser usage and optional CID replacement for inline asset rendering.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()|  | 
 **optional** | ***GetEmailHTMLOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEmailHTMLOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **replaceCidImages** | **optional.Bool**|  | [default to false]

### Return type

**string**

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html;charset=utf-8, text/html

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetEmailHTMLJson

> EmailHtmlDto GetEmailHTMLJson(ctx, emailId, optional)

Get hydrated email HTML wrapped in JSON

Returns hydrated HTML body and subject in a JSON DTO. Useful for API clients that need structured response payloads instead of raw HTML responses.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()|  | 
 **optional** | ***GetEmailHTMLJsonOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEmailHTMLJsonOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **replaceCidImages** | **optional.Bool**|  | [default to false]

### Return type

[**EmailHtmlDto**](EmailHtmlDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetEmailHTMLQuery

> EmailTextLinesResult GetEmailHTMLQuery(ctx, emailId, htmlSelector)

Query hydrated HTML body and return matching text lines

Applies a JSoup/CSS selector to hydrated HTML email body and returns matching text lines.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()| ID of email to perform HTML query on | 
**htmlSelector** | **string**| HTML selector to search for. Uses JQuery/JSoup/CSS style selector like &#39;.my-div&#39; to match content. See https://jsoup.org/apidocs/org/jsoup/select/Selector.html for more information. | 

### Return type

[**EmailTextLinesResult**](EmailTextLinesResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetEmailLinks

> EmailLinksResult GetEmailLinks(ctx, emailId, optional)

Extract links from an email HTML body

Parses HTML content and extracts link URLs (`href`). For non-HTML emails this endpoint returns a validation error.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()| ID of email to fetch text for | 
 **optional** | ***GetEmailLinksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEmailLinksOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **selector** | **optional.String**| Optional HTML query selector for links | 

### Return type

[**EmailLinksResult**](EmailLinksResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetEmailPreviewURLs

> EmailPreviewUrls GetEmailPreviewURLs(ctx, emailId)

Get email URLs for viewing in browser or downloading

Returns precomputed URLs for preview and raw message access for the specified email.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()|  | 

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


## GetEmailScreenshotAsBase64

> EmailScreenshotResult GetEmailScreenshotAsBase64(ctx, emailId, getEmailScreenshotOptions)

Take a screenshot of an email in a browser and return base64 encoded string

Renders the email in a browser engine and returns PNG data as base64. Useful for APIs and dashboards that cannot easily stream binary responses.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()|  | 
**getEmailScreenshotOptions** | [**GetEmailScreenshotOptions**](GetEmailScreenshotOptions)|  | 

### Return type

[**EmailScreenshotResult**](EmailScreenshotResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetEmailScreenshotAsBinary

> GetEmailScreenshotAsBinary(ctx, emailId, getEmailScreenshotOptions)

Take a screenshot of an email in a browser

Renders the email in a browser engine and returns PNG bytes. Intended for visual QA and rendering regression checks.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()|  | 
**getEmailScreenshotOptions** | [**GetEmailScreenshotOptions**](GetEmailScreenshotOptions)|  | 

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


## GetEmailSignature

> EmailSignatureParseResult GetEmailSignature(ctx, emailId)

Extract signature from an inbound email

Attempts to parse a sender signature block from an email body. Uses raw MIME content parts when possible and falls back to hydrated body parsing. This is heuristic and may not be accurate for all email clients or formats.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()| ID of email to extract signature from | 

### Return type

[**EmailSignatureParseResult**](EmailSignatureParseResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetEmailSummary

> EmailPreview GetEmailSummary(ctx, emailId, optional)

Get email summary (headers/metadata only)

Returns lightweight metadata and headers for an email. Use `getEmail` for hydrated body content or `getRawEmail` for original SMTP content.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()|  | 
 **optional** | ***GetEmailSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEmailSummaryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **decode** | **optional.Bool**| Deprecated and ignored. Retained for backwards compatibility. | 

### Return type

[**EmailPreview**](EmailPreview)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetEmailTextLines

> EmailTextLinesResult GetEmailTextLines(ctx, emailId, optional)

Extract normalized text lines from email body

Converts email body content to line-based plain text with optional HTML entity decoding and custom line separator.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()| ID of email to fetch text for | 
 **optional** | ***GetEmailTextLinesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEmailTextLinesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **decodeHtmlEntities** | **optional.Bool**| Decode HTML entities | 
 **lineSeparator** | **optional.String**| Line separator character | 

### Return type

[**EmailTextLinesResult**](EmailTextLinesResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetEmailThread

> EmailThreadDto GetEmailThread(ctx, threadId)

Get email thread metadata by thread ID

Returns thread metadata built from RFC 5322 `Message-ID`, `In-Reply-To`, and `References` headers. Use `getEmailThreadItems` to fetch the thread messages.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | [**string**]()|  | 

### Return type

[**EmailThreadDto**](EmailThreadDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetEmailThreadItems

> EmailThreadItemsDto GetEmailThreadItems(ctx, threadId, optional)

Get messages in a specific email thread

Returns all messages in a thread ordered by `createdAt` using the requested sort direction.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | [**string**]()|  | 
 **optional** | ***GetEmailThreadItemsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEmailThreadItemsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]

### Return type

[**EmailThreadItemsDto**](EmailThreadItemsDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetEmailThreads

> PageEmailThreadProjection GetEmailThreads(ctx, optional)

List email threads in paginated form

Lists conversation threads inferred from `Message-ID`, `In-Reply-To`, and `References`. Supports filtering by inbox, search text, and time range.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetEmailThreadsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEmailThreadsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **htmlSelector** | [**optional.Interface of string**]()| Optional inbox filter. Query parameter name is &#x60;htmlSelector&#x60; for legacy compatibility. | 
 **page** | **optional.Int32**| Optional page index in email thread pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in email thread pagination. Maximum size is 100. Use page index and sort to page through larger results | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **searchFilter** | **optional.String**| Optional search filter search filter for email threads. | 
 **since** | **optional.Time**| Optional filter email threads created since time | 
 **before** | **optional.Time**| Optional filter emails threads created before given date time | 

### Return type

[**PageEmailThreadProjection**](PageEmailThreadProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetEmailsOffsetPaginated

> PageEmailProjection GetEmailsOffsetPaginated(ctx, optional)

Get all emails in all inboxes in paginated form. Email API list all.

Offset-style pagination endpoint for listing emails across inboxes. Supports inbox filters, unread-only, search, date boundaries, favourites, connector sync, plus-address filtering, and explicit include IDs.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetEmailsOffsetPaginatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEmailsOffsetPaginatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inboxId** | [**optional.Interface of []string**](string)| Optional inbox ids to filter by. Can be repeated. By default will use all inboxes belonging to your account. | 
 **page** | **optional.Int32**| Optional page index in email list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in email list pagination. Maximum size is 100. Use page index and sort to page through larger results | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **unreadOnly** | **optional.Bool**| Optional filter for unread emails only. All emails are considered unread until they are viewed in the dashboard or requested directly | [default to false]
 **searchFilter** | **optional.String**| Optional search filter. Full email addresses match sender and receiving inbox email or receiving plus-address full address exactly. Address-like fragments containing @ (for example +alias@) also match sender, receiving inbox email, and receiving plus-address full address by contains. General text search matches sender, subject, and ID. | 
 **since** | **optional.Time**| Optional filter emails received after given date time | 
 **before** | **optional.Time**| Optional filter emails received before given date time | 
 **favourited** | **optional.Bool**| Optional filter emails that are favourited | 
 **syncConnectors** | **optional.Bool**| Sync connectors | 
 **plusAddressId** | [**optional.Interface of string**]()| Optional plus address ID filter | 
 **include** | [**optional.Interface of []string**](string)| Optional list of IDs to include in result | 

### Return type

[**PageEmailProjection**](PageEmailProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetEmailsPaginated

> PageEmailProjection GetEmailsPaginated(ctx, optional)

Get all emails in all inboxes in paginated form. Email API list all.

Primary paginated email listing endpoint. Returns emails across inboxes with support for inbox filters, unread-only, search, date boundaries, favourites, connector sync, and plus-address filtering.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetEmailsPaginatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEmailsPaginatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inboxId** | [**optional.Interface of []string**](string)| Optional inbox ids to filter by. Can be repeated. By default will use all inboxes belonging to your account. | 
 **page** | **optional.Int32**| Optional page index in email list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in email list pagination. Maximum size is 100. Use page index and sort to page through larger results | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **unreadOnly** | **optional.Bool**| Optional filter for unread emails only. All emails are considered unread until they are viewed in the dashboard or requested directly | [default to false]
 **searchFilter** | **optional.String**| Optional search filter. Full email addresses match sender and receiving inbox email or receiving plus-address full address exactly. Address-like fragments containing @ (for example +alias@) also match sender, receiving inbox email, and receiving plus-address full address by contains. General text search matches sender, subject, and ID. | 
 **since** | **optional.Time**| Optional filter emails received after given date time. If unset will use time 24hours prior to now. | 
 **before** | **optional.Time**| Optional filter emails received before given date time | 
 **syncConnectors** | **optional.Bool**| Sync connectors | 
 **plusAddressId** | [**optional.Interface of string**]()| Optional plus address ID filter | 
 **favourited** | **optional.Bool**| Optional filter emails that are favourited | 

### Return type

[**PageEmailProjection**](PageEmailProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetGravatarUrlForEmailAddress

> GravatarUrl GetGravatarUrlForEmailAddress(ctx, emailAddress, optional)

Get Gravatar URL for an email address

Builds a Gravatar image URL from the provided email address and optional size. This endpoint does not fetch image bytes; it only returns the computed URL.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailAddress** | **string**|  | 
 **optional** | ***GetGravatarUrlForEmailAddressOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetGravatarUrlForEmailAddressOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | **optional.String**|  | 

### Return type

[**GravatarUrl**](GravatarUrl)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetLatestEmail

> Email GetLatestEmail(ctx, optional)

Get latest email in all inboxes. Most recently received.

Returns the most recently received email across all inboxes or an optional subset of inbox IDs.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetLatestEmailOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetLatestEmailOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inboxIds** | [**optional.Interface of []string**](string)| Optional set of inboxes to filter by. Only get the latest email from these inbox IDs. If not provided will search across all inboxes | 

### Return type

[**Email**](Email)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetLatestEmailInInbox1

> Email GetLatestEmailInInbox1(ctx, inboxId)

Get latest email in an inbox. Use `WaitForController` to get emails that may not have arrived yet.

Returns the newest email for the specified inbox ID. For polling/wait semantics use wait endpoints.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()| ID of the inbox you want to get the latest email from | 

### Return type

[**Email**](Email)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetOrganizationEmailsPaginated

> PageEmailProjection GetOrganizationEmailsPaginated(ctx, optional)

List organization-visible emails

Returns paginated emails visible through organization/team access. Supports inbox filtering, unread-only filtering, search, favourites, plus-address filtering, and optional connector sync.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetOrganizationEmailsPaginatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetOrganizationEmailsPaginatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inboxId** | [**optional.Interface of []string**](string)| Optional inbox ids to filter by. Can be repeated. By default will use all inboxes belonging to your account. | 
 **page** | **optional.Int32**| Optional page index in email list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in email list pagination. Maximum size is 100. Use page index and sort to page through larger results | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **unreadOnly** | **optional.Bool**| Optional filter for unread emails only. All emails are considered unread until they are viewed in the dashboard or requested directly | [default to false]
 **searchFilter** | **optional.String**| Optional search filter search filter for emails. | 
 **since** | **optional.Time**| Optional filter emails received after given date time. If unset will use time 24hours prior to now. | 
 **before** | **optional.Time**| Optional filter emails received before given date time | 
 **syncConnectors** | **optional.Bool**| Sync connectors | 
 **favourited** | **optional.Bool**| Search only favorited emails | 
 **plusAddressId** | [**optional.Interface of string**]()| Optional plus address ID filter | 

### Return type

[**PageEmailProjection**](PageEmailProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetRawEmailContents

> GetRawEmailContents(ctx, emailId)

Get raw email string. Returns unparsed raw SMTP message with headers and body.

Returns the original unparsed SMTP/MIME message as `text/plain`. Use JSON variant if your client expects JSON transport.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()| ID of email | 

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


## GetRawEmailJson

> RawEmailJson GetRawEmailJson(ctx, emailId)

Get raw email in JSON. Unparsed SMTP message in JSON wrapper format.

Returns the original unparsed SMTP/MIME message wrapped in a JSON DTO for API clients that avoid plain-text stream responses.

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


## GetUnreadEmailCount

> UnreadCount GetUnreadEmailCount(ctx, optional)

Get unread email count

Returns unread email count. An email is considered read after dashboard/API retrieval depending on your read workflow.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetUnreadEmailCountOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetUnreadEmailCountOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inboxId** | [**optional.Interface of string**]()| Inbox ID filter for unread count scope | 

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


## MarkAllAsRead

> MarkAllAsRead(ctx, optional)

Mark all emails as read or unread

Sets read state for all emails, optionally scoped to one inbox. Use `read=false` to reset unread state in bulk.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MarkAllAsReadOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MarkAllAsReadOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **read** | **optional.Bool**| What value to assign to email read property. Default true. | [default to true]
 **inboxId** | [**optional.Interface of string**]()| Optional inbox ID filter | 

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


## MarkAsRead

> EmailPreview MarkAsRead(ctx, emailId, optional)

Mark an email as read or unread

Sets read state for one email. Useful when implementing custom mailbox workflows that treat viewed messages as unread.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()|  | 
 **optional** | ***MarkAsReadOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MarkAsReadOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **read** | **optional.Bool**| What value to assign to email read property. Default true. | [default to true]

### Return type

[**EmailPreview**](EmailPreview)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## ReplyToEmail

> SentEmailDto ReplyToEmail(ctx, emailId, replyToEmailOptions)

Reply to an email

Sends a reply using the original conversation context (subject/thread headers). Reply target resolution honors sender/reply-to semantics.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()| ID of the email that should be replied to | 
**replyToEmailOptions** | [**ReplyToEmailOptions**](ReplyToEmailOptions)|  | 

### Return type

[**SentEmailDto**](SentEmailDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## SearchEmails

> PageEmailProjection SearchEmails(ctx, searchEmailsOptions, optional)

Get all emails by search criteria. Return in paginated form.

Searches emails by sender/recipient/address/subject/id fields and returns paginated matches. Does not perform full-text body search.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**searchEmailsOptions** | [**SearchEmailsOptions**](SearchEmailsOptions)|  | 
 **optional** | ***SearchEmailsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchEmailsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **syncConnectors** | **optional.Bool**| Sync connectors | 
 **favourited** | **optional.Bool**| Search only favourited emails | 
 **plusAddressId** | [**optional.Interface of string**]()| Optional plus address ID filter | 

### Return type

[**PageEmailProjection**](PageEmailProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## SendEmailSourceOptional

> SendEmailSourceOptional(ctx, sendEmailOptions, optional)

Send email

Sends an email from an existing inbox, or creates a temporary inbox when `inboxId` is not provided. Supports `useDomainPool` and `virtualSend` inbox creation behavior for convenience sends.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sendEmailOptions** | [**SendEmailOptions**](SendEmailOptions)|  | 
 **optional** | ***SendEmailSourceOptionalOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SendEmailSourceOptionalOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inboxId** | [**optional.Interface of string**]()| ID of the inbox you want to send the email from | 
 **useDomainPool** | **optional.Bool**| Use domain pool. Optionally create inbox to send from using the mailslurp domain pool. | 
 **virtualSend** | **optional.Bool**| Optionally create inbox to send from that is a virtual inbox and won&#39;t send to external addresses | 

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


## SetEmailFavourited

> SetEmailFavourited(ctx, emailId, favourited)

Set email favourited state

Sets favourite state for an email for dashboard/search workflows.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()| ID of email to set favourite state | 
**favourited** | **bool**|  | 

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


## ValidateEmail

> ValidationDto ValidateEmail(ctx, emailId)

Validate email HTML contents

Runs HTML validation on the email body when HTML is present. Non-HTML emails are treated as valid for this check.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()| ID of email | 

### Return type

[**ValidationDto**](ValidationDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

