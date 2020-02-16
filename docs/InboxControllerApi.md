# \InboxControllerApi

All URIs are relative to *https://api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInbox**](InboxControllerApi.md#CreateInbox) | **Post** /inboxes | Create an Inbox (email address)
[**DeleteAllInboxes**](InboxControllerApi.md#DeleteAllInboxes) | **Delete** /inboxes | Delete all inboxes
[**DeleteInbox**](InboxControllerApi.md#DeleteInbox) | **Delete** /inboxes/{inboxId} | Delete Inbox / Email Address
[**GetAllInboxes**](InboxControllerApi.md#GetAllInboxes) | **Get** /inboxes/paginated | List Inboxes Paginated
[**GetEmails**](InboxControllerApi.md#GetEmails) | **Get** /inboxes/{inboxId}/emails | Get emails in an Inbox
[**GetInbox**](InboxControllerApi.md#GetInbox) | **Get** /inboxes/{inboxId} | Get Inbox / EmailAddress
[**GetInboxEmailsPaginated**](InboxControllerApi.md#GetInboxEmailsPaginated) | **Get** /inboxes/{inboxId}/emails/paginated | Get inbox emails paginated
[**GetInboxes**](InboxControllerApi.md#GetInboxes) | **Get** /inboxes | List Inboxes / Email Addresses
[**SendEmail**](InboxControllerApi.md#SendEmail) | **Post** /inboxes/{inboxId} | Send Email
[**SetInboxFavourited**](InboxControllerApi.md#SetInboxFavourited) | **Put** /inboxes/{inboxId}/favourite | Set inbox favourited state



## CreateInbox

> Inbox CreateInbox(ctx, optional)

Create an Inbox (email address)

Create a new inbox and with a ranmdomized email address to send and receive from. Pass emailAddress parameter if you wish to use a specific email address. Creating an inbox is required before sending or receiving emails. If writing tests it is recommended that you create a new inbox during each test method so that it is unique and empty. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateInboxOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateInboxOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **description** | **optional.String**| Optional description for an inbox. | 
 **emailAddress** | **optional.String**| Optional email address including domain you wish inbox to use (eg: test123@mydomain.com). Only supports domains that you have registered and verified with MailSlurp using dashboard or &#x60;createDomain&#x60; method. | 
 **expiresAt** | **optional.Time**| Optional expires at timestamp. If your plan supports this feature you can specify when an inbox should expire. If left empty inbox will exist permanently or expire when your plan dictates | 
 **favourite** | **optional.Bool**| Is inbox favourited. | 
 **name** | **optional.String**| Optional name for an inbox. | 
 **tags** | [**optional.Interface of []string**](string.md)| Optional tags for an inbox. Can be used for searching and filtering inboxes. | 

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


## DeleteAllInboxes

> DeleteAllInboxes(ctx, )

Delete all inboxes

Permanently delete all inboxes and associated email addresses and all emails within the given inboxes

### Required Parameters

This endpoint does not need any parameter.

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


## DeleteInbox

> DeleteInbox(ctx, inboxId)

Delete Inbox / Email Address

Permanently delete an inbox and associated email address and all emails within the given inboxes

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


## GetAllInboxes

> PageInboxProjection GetAllInboxes(ctx, optional)

List Inboxes Paginated

List inboxes in paginated form. Allows for page index, page size, and sort direction. Can also filter by favourited or email address like pattern.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllInboxesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllInboxesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **favourite** | **optional.Bool**| Optionally filter results for favourites only | [default to false]
 **page** | **optional.Int32**| Optional page index in inbox list pagination | [default to 0]
 **search** | **optional.String**| Optionally filter by search words partial matching ID, tags, name, and email address | 
 **size** | **optional.Int32**| Optional page size in inbox list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]

### Return type

[**PageInboxProjection**](Page«InboxProjection».md)

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmails

> []EmailPreview GetEmails(ctx, inboxId, optional)

Get emails in an Inbox

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

 **limit** | **optional.Int32**| Limit the result set, ordered by received date time sort direction | 
 **minCount** | **optional.Int64**| Minimum acceptable email count. Will cause request to hang (and retry) until minCount is satisfied or retryTimeout is reached. | 
 **retryTimeout** | **optional.Int64**| Maximum milliseconds to spend retrying inbox database until minCount emails are returned | 
 **since** | **optional.Time**| Exclude emails received before this ISO 8601 date time | 
 **sort** | **optional.String**| Sort the results by received date and direction ASC or DESC | 

### Return type

[**[]EmailPreview**](EmailPreview.md)

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInbox

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInboxEmailsPaginated

> PageEmailPreview GetInboxEmailsPaginated(ctx, inboxId, optional)

Get inbox emails paginated

Get a paginated list of emails in an inbox. Does not hold connections open.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**](.md)| Id of inbox that emails belongs to | 
 **optional** | ***GetInboxEmailsPaginatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetInboxEmailsPaginatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Optional page index in inbox emails list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in inbox emails list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]

### Return type

[**PageEmailPreview**](Page«EmailPreview».md)

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInboxes

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendEmail

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetInboxFavourited

> Inbox SetInboxFavourited(ctx, inboxId, setInboxFavouritedOptions)

Set inbox favourited state

Set and return new favourite state for an inbox

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**](.md)| inboxId | 
**setInboxFavouritedOptions** | [**SetInboxFavouritedOptions**](SetInboxFavouritedOptions.md)| setInboxFavouritedOptions | 

### Return type

[**Inbox**](Inbox.md)

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

