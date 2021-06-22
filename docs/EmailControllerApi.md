# MailSlurp\EmailControllerApi

All URIs are relative to *https://api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAllEmails**](EmailControllerApi#DeleteAllEmails) | **Delete** /emails | Delete all emails in all inboxes.
[**DeleteEmail**](EmailControllerApi#DeleteEmail) | **Delete** /emails/{emailId} | Delete an email
[**DownloadAttachment**](EmailControllerApi#DownloadAttachment) | **Get** /emails/{emailId}/attachments/{attachmentId} | Get email attachment bytes. Returned as &#x60;octet-stream&#x60; with content type header. If you have trouble with byte responses try the &#x60;downloadAttachmentBase64&#x60; response endpoints and convert the base 64 encoded content to a file or string.
[**DownloadAttachmentBase64**](EmailControllerApi#DownloadAttachmentBase64) | **Get** /emails/{emailId}/attachments/{attachmentId}/base64 | Get email attachment as base64 encoded string as an alternative to binary responses. Decode the &#x60;base64FileContents&#x60; as a &#x60;utf-8&#x60; encoded string or array of bytes depending on the &#x60;contentType&#x60;.
[**DownloadBody**](EmailControllerApi#DownloadBody) | **Get** /emails/{emailId}/body | Get email body as string. Returned as &#x60;plain/text&#x60; with content type header.
[**DownloadBodyBytes**](EmailControllerApi#DownloadBodyBytes) | **Get** /emails/{emailId}/body-bytes | Get email body in bytes. Returned as &#x60;octet-stream&#x60; with content type header.
[**ForwardEmail**](EmailControllerApi#ForwardEmail) | **Post** /emails/{emailId}/forward | Forward email to recipients
[**GetAttachmentMetaData**](EmailControllerApi#GetAttachmentMetaData) | **Get** /emails/{emailId}/attachments/{attachmentId}/metadata | Get email attachment metadata. This is the &#x60;contentType&#x60; and &#x60;contentLength&#x60; of an attachment. To get the individual attachments  use the &#x60;downloadAttachment&#x60; methods.
[**GetAttachments1**](EmailControllerApi#GetAttachments1) | **Get** /emails/{emailId}/attachments | Get all email attachment metadata. Metadata includes name and size of attachments.
[**GetEmail**](EmailControllerApi#GetEmail) | **Get** /emails/{emailId} | Get email content including headers and body. Expects email to exist by ID. For emails that may not have arrived yet use the WaitForController.
[**GetEmailContentMatch**](EmailControllerApi#GetEmailContentMatch) | **Post** /emails/{emailId}/contentMatch | Get email content regex pattern match results. Runs regex against email body and returns match groups.
[**GetEmailHTML**](EmailControllerApi#GetEmailHTML) | **Get** /emails/{emailId}/html | Get email content as HTML. For displaying emails in browser context.
[**GetEmailHTMLQuery**](EmailControllerApi#GetEmailHTMLQuery) | **Get** /emails/{emailId}/htmlQuery | Parse and return text from an email, stripping HTML and decoding encoded characters
[**GetEmailTextLines**](EmailControllerApi#GetEmailTextLines) | **Get** /emails/{emailId}/textLines | Parse and return text from an email, stripping HTML and decoding encoded characters
[**GetEmailsPaginated**](EmailControllerApi#GetEmailsPaginated) | **Get** /emails | Get all emails in all inboxes in paginated form. Email API list all.
[**GetLatestEmail**](EmailControllerApi#GetLatestEmail) | **Get** /emails/latest | Get latest email in all inboxes. Most recently received.
[**GetLatestEmailInInbox**](EmailControllerApi#GetLatestEmailInInbox) | **Get** /emails/latestIn | Get latest email in an inbox. Use &#x60;WaitForController&#x60; to get emails that may not have arrived yet.
[**GetOrganizationEmailsPaginated**](EmailControllerApi#GetOrganizationEmailsPaginated) | **Get** /emails/organization | Get all organization emails. List team or shared test email accounts
[**GetRawEmailContents**](EmailControllerApi#GetRawEmailContents) | **Get** /emails/{emailId}/raw | Get raw email string. Returns unparsed raw SMTP message with headers and body.
[**GetRawEmailJson**](EmailControllerApi#GetRawEmailJson) | **Get** /emails/{emailId}/raw/json | Get raw email in JSON. Unparsed SMTP message in JSON wrapper format.
[**GetUnreadEmailCount**](EmailControllerApi#GetUnreadEmailCount) | **Get** /emails/unreadCount | Get unread email count
[**ReplyToEmail**](EmailControllerApi#ReplyToEmail) | **Put** /emails/{emailId} | Reply to an email
[**ValidateEmail**](EmailControllerApi#ValidateEmail) | **Post** /emails/{emailId}/validate | Validate email HTML contents



## DeleteAllEmails

> DeleteAllEmails(ctx, )

Delete all emails in all inboxes.

Deletes all emails in your account. Be careful as emails cannot be recovered

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

Deletes an email and removes it from the inbox. Deleted emails cannot be recovered.

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

> string DownloadAttachment(ctx, attachmentId, emailId, optional)

Get email attachment bytes. Returned as `octet-stream` with content type header. If you have trouble with byte responses try the `downloadAttachmentBase64` response endpoints and convert the base 64 encoded content to a file or string.

Returns the specified attachment for a given email as a stream / array of bytes. You can find attachment ids in email responses endpoint responses. The response type is application/octet-stream.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attachmentId** | **string**| ID of attachment | 
**emailId** | [**string**]()| ID of email | 
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

> DownloadAttachmentDto DownloadAttachmentBase64(ctx, attachmentId, emailId)

Get email attachment as base64 encoded string as an alternative to binary responses. Decode the `base64FileContents` as a `utf-8` encoded string or array of bytes depending on the `contentType`.

Returns the specified attachment for a given email as a base 64 encoded string. The response type is application/json. This method is similar to the `downloadAttachment` method but allows some clients to get around issues with binary responses.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attachmentId** | **string**| ID of attachment | 
**emailId** | [**string**]()| ID of email | 

### Return type

[**DownloadAttachmentDto**](DownloadAttachmentDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DownloadBody

> string DownloadBody(ctx, emailId)

Get email body as string. Returned as `plain/text` with content type header.

Returns the specified email body for a given email as a string

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
- **Accept**: text/html, text/plain

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DownloadBodyBytes

> string DownloadBodyBytes(ctx, emailId)

Get email body in bytes. Returned as `octet-stream` with content type header.

Returns the specified email body for a given email as a stream / array of bytes.

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

> ForwardEmail(ctx, emailId, forwardEmailOptions)

Forward email to recipients

Forward an existing email to new recipients. The sender of the email will be the inbox that received the email you are forwarding. You can override the sender with the `from` option. Note you must have access to the from address in MailSlurp to use the override. For more control consider fetching the email and sending it a new using the send email endpoints.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()| ID of email | 
**forwardEmailOptions** | [**ForwardEmailOptions**](ForwardEmailOptions)| forwardEmailOptions | 

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


## GetAttachmentMetaData

> AttachmentMetaData GetAttachmentMetaData(ctx, attachmentId, emailId)

Get email attachment metadata. This is the `contentType` and `contentLength` of an attachment. To get the individual attachments  use the `downloadAttachment` methods.

Returns the metadata such as name and content-type for a given attachment and email.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attachmentId** | **string**| ID of attachment | 
**emailId** | [**string**]()| ID of email | 

### Return type

[**AttachmentMetaData**](AttachmentMetaData)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetAttachments1

> []AttachmentMetaData GetAttachments1(ctx, emailId)

Get all email attachment metadata. Metadata includes name and size of attachments.

Returns an array of attachment metadata such as name and content-type for a given email if present.

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetEmail

> Email GetEmail(ctx, emailId, optional)

Get email content including headers and body. Expects email to exist by ID. For emails that may not have arrived yet use the WaitForController.

Returns a email summary object with headers and content. To retrieve the raw unparsed email use the getRawEmail endpoints

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()| emailId | 
 **optional** | ***GetEmailOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEmailOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **decode** | **optional.Bool**| Decode email body quoted-printable encoding to plain text. SMTP servers often encode text using quoted-printable format (for instance &#x60;&#x3D;D7&#x60;). This can be a pain for testing | [default to false]

### Return type

[**Email**](Email)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetEmailContentMatch

> EmailContentMatchResult GetEmailContentMatch(ctx, emailId, contentMatchOptions)

Get email content regex pattern match results. Runs regex against email body and returns match groups.

Return the matches for a given Java style regex pattern. Do not include the typical `/` at start or end of regex in some languages. Given an example `your code is: 12345` the pattern to extract match looks like `code is: (\\d{6})`. This will return an array of matches with the first matching the entire pattern and the subsequent matching the groups: `['code is: 123456', '123456']` See https://docs.oracle.com/javase/8/docs/api/java/util/regex/Pattern.html for more information of available patterns. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()| ID of email to match against | 
**contentMatchOptions** | [**ContentMatchOptions**](ContentMatchOptions)| contentMatchOptions | 

### Return type

[**EmailContentMatchResult**](EmailContentMatchResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetEmailHTML

> string GetEmailHTML(ctx, emailId, optional)

Get email content as HTML. For displaying emails in browser context.

Retrieve email content as HTML response for viewing in browsers. Decodes quoted-printable entities and converts charset to UTF-8. Pass your API KEY as a request parameter when viewing in a browser: `?apiKey=xxx`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()| emailId | 
 **optional** | ***GetEmailHTMLOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEmailHTMLOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **decode** | **optional.Bool**| decode | [default to false]

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


## GetEmailHTMLQuery

> EmailTextLinesResult GetEmailHTMLQuery(ctx, emailId, optional)

Parse and return text from an email, stripping HTML and decoding encoded characters

Parse an email body and return the content as an array of text. HTML parsing uses JSoup which supports JQuery/CSS style selectors

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()| ID of email to perform HTML query on | 
 **optional** | ***GetEmailHTMLQueryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEmailHTMLQueryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **htmlSelector** | **optional.String**| HTML selector to search for. Uses JQuery/JSoup/CSS style selector like &#39;.my-div&#39; to match content. See https://jsoup.org/apidocs/org/jsoup/select/Selector.html for more information. | 

### Return type

[**EmailTextLinesResult**](EmailTextLinesResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetEmailTextLines

> EmailTextLinesResult GetEmailTextLines(ctx, emailId, optional)

Parse and return text from an email, stripping HTML and decoding encoded characters

Parse an email body and return the content as an array of strings. HTML parsing uses JSoup and UNIX line separators.

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetEmailsPaginated

> PageEmailProjection GetEmailsPaginated(ctx, optional)

Get all emails in all inboxes in paginated form. Email API list all.

By default returns all emails across all inboxes sorted by ascending created at date. Responses are paginated. You can restrict results to a list of inbox IDs. You can also filter out read messages

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
 **searchFilter** | **optional.String**| Optional search filter. Searches email recipients, sender, subject, email address and ID. Does not search email body | [default to false]
 **size** | **optional.Int32**| Optional page size in email list pagination. Maximum size is 100. Use page index and sort to page through larger results | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **unreadOnly** | **optional.Bool**| Optional filter for unread emails only. All emails are considered unread until they are viewed in the dashboard or requested directly | [default to false]

### Return type

[**PageEmailProjection**](PageEmailProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetLatestEmail

> Email GetLatestEmail(ctx, optional)

Get latest email in all inboxes. Most recently received.

Get the newest email in all inboxes or in a passed set of inbox IDs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetLatestEmailOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetLatestEmailOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inboxIds** | [**optional.Interface of []string**](string)| Optional set of inboxes to filter by. Only get the latest email from these inbox IDs | 

### Return type

[**Email**](Email)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetLatestEmailInInbox

> Email GetLatestEmailInInbox(ctx, optional)

Get latest email in an inbox. Use `WaitForController` to get emails that may not have arrived yet.

Get the newest email in all inboxes or in a passed set of inbox IDs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetLatestEmailInInboxOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetLatestEmailInInboxOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inboxId** | [**optional.Interface of string**]()| ID of the inbox you want to get the latest email from | 

### Return type

[**Email**](Email)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetOrganizationEmailsPaginated

> PageEmailProjection GetOrganizationEmailsPaginated(ctx, optional)

Get all organization emails. List team or shared test email accounts

By default returns all emails across all team inboxes sorted by ascending created at date. Responses are paginated. You can restrict results to a list of inbox IDs. You can also filter out read messages

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
 **searchFilter** | **optional.String**| Optional search filter search filter for emails. | 
 **size** | **optional.Int32**| Optional page size in email list pagination. Maximum size is 100. Use page index and sort to page through larger results | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **unreadOnly** | **optional.Bool**| Optional filter for unread emails only. All emails are considered unread until they are viewed in the dashboard or requested directly | [default to false]

### Return type

[**PageEmailProjection**](PageEmailProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetRawEmailContents

> string GetRawEmailContents(ctx, emailId)

Get raw email string. Returns unparsed raw SMTP message with headers and body.

Returns a raw, unparsed, and unprocessed email. If your client has issues processing the response it is likely due to the response content-type which is text/plain. If you need a JSON response content-type use the getRawEmailJson endpoint

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


## GetRawEmailJson

> RawEmailJson GetRawEmailJson(ctx, emailId)

Get raw email in JSON. Unparsed SMTP message in JSON wrapper format.

Returns a raw, unparsed, and unprocessed email wrapped in a JSON response object for easier handling when compared with the getRawEmail text/plain response

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetUnreadEmailCount

> UnreadCount GetUnreadEmailCount(ctx, )

Get unread email count

Get number of emails unread. Unread means has not been viewed in dashboard or returned in an email API response

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**UnreadCount**](UnreadCount)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## ReplyToEmail

> SentEmailDto ReplyToEmail(ctx, emailId, replyToEmailOptions)

Reply to an email

Send the reply to the email sender or reply-to and include same subject cc bcc etc. Reply to an email and the contents will be sent with the existing subject to the emails `to`, `cc`, and `bcc`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()| ID of the email that should be replied to | 
**replyToEmailOptions** | [**ReplyToEmailOptions**](ReplyToEmailOptions)| replyToEmailOptions | 

### Return type

[**SentEmailDto**](SentEmailDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## ValidateEmail

> ValidationDto ValidateEmail(ctx, emailId)

Validate email HTML contents

Validate the HTML content of email if HTML is found. Considered valid if no HTML is present.

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

