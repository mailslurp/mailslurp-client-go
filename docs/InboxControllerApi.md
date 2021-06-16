# MailSlurp\InboxControllerApi

All URIs are relative to *https://api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInbox**](InboxControllerApi#CreateInbox) | **Post** /inboxes | Create an inbox email address. An inbox has a real email address and can send and receive emails. Inboxes can be either &#x60;SMTP&#x60; or &#x60;HTTP&#x60; inboxes.
[**CreateInboxWithDefaults**](InboxControllerApi#CreateInboxWithDefaults) | **Post** /inboxes/withDefaults | Create an inbox with default options. Uses MailSlurp domain pool address and is private.
[**CreateInboxWithOptions**](InboxControllerApi#CreateInboxWithOptions) | **Post** /inboxes/withOptions | Create an inbox with options. Extended options for inbox creation.
[**DeleteAllInboxes**](InboxControllerApi#DeleteAllInboxes) | **Delete** /inboxes | Delete all inboxes
[**DeleteInbox**](InboxControllerApi#DeleteInbox) | **Delete** /inboxes/{inboxId} | Delete inbox
[**GetAllInboxes**](InboxControllerApi#GetAllInboxes) | **Get** /inboxes/paginated | List All Inboxes Paginated
[**GetEmails**](InboxControllerApi#GetEmails) | **Get** /inboxes/{inboxId}/emails | Get emails in an Inbox. This method is not idempotent as it allows retries and waits if you want certain conditions to be met before returning. For simple listing and sorting of known emails use the email controller instead.
[**GetInbox**](InboxControllerApi#GetInbox) | **Get** /inboxes/{inboxId} | Get Inbox. Returns properties of an inbox.
[**GetInboxEmailsPaginated**](InboxControllerApi#GetInboxEmailsPaginated) | **Get** /inboxes/{inboxId}/emails/paginated | Get inbox emails paginated
[**GetInboxSentEmails**](InboxControllerApi#GetInboxSentEmails) | **Get** /inboxes/{inboxId}/sent | Get Inbox Sent Emails
[**GetInboxTags**](InboxControllerApi#GetInboxTags) | **Get** /inboxes/tags | Get inbox tags
[**GetInboxes**](InboxControllerApi#GetInboxes) | **Get** /inboxes | List Inboxes and email eddresses
[**GetOrganizationInboxes**](InboxControllerApi#GetOrganizationInboxes) | **Get** /inboxes/organization | List Organization Inboxes Paginated
[**SendEmail**](InboxControllerApi#SendEmail) | **Post** /inboxes/{inboxId} | Send Email
[**SendEmailAndConfirm**](InboxControllerApi#SendEmailAndConfirm) | **Post** /inboxes/{inboxId}/confirm | Send email and return sent confirmation
[**SendTestEmail**](InboxControllerApi#SendTestEmail) | **Post** /inboxes/{inboxId}/send-test-email | Send a test email to inbox
[**SetInboxFavourited**](InboxControllerApi#SetInboxFavourited) | **Put** /inboxes/{inboxId}/favourite | Set inbox favourited state
[**UpdateInbox**](InboxControllerApi#UpdateInbox) | **Patch** /inboxes/{inboxId} | Update Inbox. Change name and description. Email address is not editable.



## CreateInbox

> Inbox CreateInbox(ctx, optional)

Create an inbox email address. An inbox has a real email address and can send and receive emails. Inboxes can be either `SMTP` or `HTTP` inboxes.

Create a new inbox and with a randomized email address to send and receive from. Pass emailAddress parameter if you wish to use a specific email address. Creating an inbox is required before sending or receiving emails. If writing tests it is recommended that you create a new inbox during each test method so that it is unique and empty. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateInboxOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateInboxOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **allowTeamAccess** | **optional.Bool**| Grant team access to this inbox and the emails that belong to it for team members of your organization. | 
 **description** | **optional.String**| Optional description of the inbox for labelling purposes. Is shown in the dashboard and can be used with | 
 **emailAddress** | **optional.String**| A custom email address to use with the inbox. Defaults to null. When null MailSlurp will assign a random email address to the inbox such as &#x60;123@mailslurp.com&#x60;. If you use the &#x60;useDomainPool&#x60; option when the email address is null it will generate an email address with a more varied domain ending such as &#x60;123@mailslurp.info&#x60; or &#x60;123@mailslurp.biz&#x60;. When a custom email address is provided the address is split into a domain and the domain is queried against your user. If you have created the domain in the MailSlurp dashboard and verified it you can use any email address that ends with the domain. Note domain types must match the inbox type - so &#x60;SMTP&#x60; inboxes will only work with &#x60;SMTP&#x60; type domains. Send an email to this address and the inbox will receive and store it for you. To retrieve the email use the Inbox and Email Controller endpoints with the inbox ID. | 
 **expiresAt** | **optional.Time**| Optional inbox expiration date. If null then this inbox is permanent and the emails in it won&#39;t be deleted. If an expiration date is provided or is required by your plan the inbox will be closed when the expiration time is reached. Expired inboxes still contain their emails but can no longer send or receive emails. An ExpiredInboxRecord is created when an inbox and the email address and inbox ID are recorded. The expiresAt property is a timestamp string in ISO DateTime Format yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSXXX. | 
 **expiresIn** | **optional.Int64**| Number of milliseconds that inbox should exist for | 
 **favourite** | **optional.Bool**| Is the inbox a favorite. Marking an inbox as a favorite is typically done in the dashboard for quick access or filtering | 
 **inboxType** | **optional.String**| HTTP (default) or SMTP inbox type. HTTP inboxes are best for testing while SMTP inboxes are more reliable for public inbound email consumption. When using custom domains the domain type must match the inbox type. HTTP inboxes are processed by AWS SES while SMTP inboxes use a custom mail server running at &#x60;mx.mailslurp.com&#x60;. | 
 **name** | **optional.String**| Optional name of the inbox. Displayed in the dashboard for easier search and used as the sender name when sending emails. | 
 **tags** | [**optional.Interface of []string**](string)| Tags that inbox has been tagged with. Tags can be added to inboxes to group different inboxes within an account. You can also search for inboxes by tag in the dashboard UI. | 
 **useDomainPool** | **optional.Bool**| Use the MailSlurp domain name pool with this inbox when creating the email address. Defaults to null. If enabled the inbox will be an email address with a domain randomly chosen from a list of the MailSlurp domains. This is useful when the default &#x60;@mailslurp.com&#x60; email addresses used with inboxes are blocked or considered spam by a provider or receiving service. When domain pool is enabled an email address will be generated ending in &#x60;@mailslurp.{world,info,xyz,...}&#x60; . This means a TLD is randomly selecting from a list of &#x60;.biz&#x60;, &#x60;.info&#x60;, &#x60;.xyz&#x60; etc to add variance to the generated email addresses. When null or false MailSlurp uses the default behavior of &#x60;@mailslurp.com&#x60; or custom email address provided by the emailAddress field. Note this feature is only available for &#x60;HTTP&#x60; inbox types. | 

### Return type

[**Inbox**](Inbox)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CreateInboxWithDefaults

> Inbox CreateInboxWithDefaults(ctx, )

Create an inbox with default options. Uses MailSlurp domain pool address and is private.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**Inbox**](Inbox)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CreateInboxWithOptions

> Inbox CreateInboxWithOptions(ctx, createInboxDto)

Create an inbox with options. Extended options for inbox creation.

Additional endpoint that allows inbox creation with request body options. Can be more flexible that other methods for some clients.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createInboxDto** | [**CreateInboxDto**](CreateInboxDto)| createInboxDto | 

### Return type

[**Inbox**](Inbox)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DeleteAllInboxes

> DeleteAllInboxes(ctx, )

Delete all inboxes

Permanently delete all inboxes and associated email addresses. This will also delete all emails within the inboxes. Be careful as inboxes cannot be recovered once deleted. Note: deleting inboxes will not impact your usage limits. Monthly inbox creation limits are based on how many inboxes were created in the last 30 days, not how many inboxes you currently have.

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


## DeleteInbox

> DeleteInbox(ctx, inboxId)

Delete inbox

Permanently delete an inbox and associated email address as well as all emails within the given inbox. This action cannot be undone. Note: deleting an inbox will not affect your account usage. Monthly inbox usage is based on how many inboxes you create within 30 days, not how many exist at time of request.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()| inboxId | 

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


## GetAllInboxes

> PageInboxProjection GetAllInboxes(ctx, optional)

List All Inboxes Paginated

List inboxes in paginated form. The results are available on the `content` property of the returned object. This method allows for page index (zero based), page size (how many results to return), and a sort direction (based on createdAt time). You Can also filter by whether an inbox is favorited or use email address pattern. This method is the recommended way to query inboxes. The alternative `getInboxes` method returns a full list of inboxes but is limited to 100 results. Results do not include team access inboxes by default. Use organization method to list team inboxes or set `teamAccess` to true.

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
 **tag** | **optional.String**| Optionally filter by tags. Will return inboxes that include given tags | 
 **teamAccess** | **optional.Bool**| Optionally filter by team access. Defaults to false so organization inboxes are not included | [default to false]

### Return type

[**PageInboxProjection**](PageInboxProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetEmails

> []EmailPreview GetEmails(ctx, inboxId, optional)

Get emails in an Inbox. This method is not idempotent as it allows retries and waits if you want certain conditions to be met before returning. For simple listing and sorting of known emails use the email controller instead.

List emails that an inbox has received. Only emails that are sent to the inbox's email address will appear in the inbox. It may take several seconds for any email you send to an inbox's email address to appear in the inbox. To make this endpoint wait for a minimum number of emails use the `minCount` parameter. The server will retry the inbox database until the `minCount` is satisfied or the `retryTimeout` is reached

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()| Id of inbox that emails belongs to | 
 **optional** | ***GetEmailsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEmailsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| Limit the result set, ordered by received date time sort direction. Maximum 100. For more listing options see the email controller | 
 **minCount** | **optional.Int64**| Minimum acceptable email count. Will cause request to hang (and retry) until minCount is satisfied or retryTimeout is reached. | 
 **retryTimeout** | **optional.Int64**| Maximum milliseconds to spend retrying inbox database until minCount emails are returned | 
 **since** | **optional.Time**| Exclude emails received before this ISO 8601 date time | 
 **size** | **optional.Int32**| Alias for limit. Assessed first before assessing any passed limit. | 
 **sort** | **optional.String**| Sort the results by received date and direction ASC or DESC | 

### Return type

[**[]EmailPreview**](EmailPreview)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetInbox

> Inbox GetInbox(ctx, inboxId)

Get Inbox. Returns properties of an inbox.

Returns an inbox's properties, including its email address and ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()| inboxId | 

### Return type

[**Inbox**](Inbox)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetInboxEmailsPaginated

> PageEmailPreview GetInboxEmailsPaginated(ctx, inboxId, optional)

Get inbox emails paginated

Get a paginated list of emails in an inbox. Does not hold connections open.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()| Id of inbox that emails belongs to | 
 **optional** | ***GetInboxEmailsPaginatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetInboxEmailsPaginatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Optional page index in inbox emails list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in inbox emails list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]

### Return type

[**PageEmailPreview**](PageEmailPreview)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetInboxSentEmails

> PageSentEmailProjection GetInboxSentEmails(ctx, inboxId, optional)

Get Inbox Sent Emails

Returns an inbox's sent email receipts. Call individual sent email endpoints for more details. Note for privacy reasons the full body of sent emails is never stored. An MD5 hash hex is available for comparison instead.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()| inboxId | 
 **optional** | ***GetInboxSentEmailsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetInboxSentEmailsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Optional page index in inbox sent email list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in inbox sent email list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]

### Return type

[**PageSentEmailProjection**](PageSentEmailProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetInboxTags

> []string GetInboxTags(ctx, )

Get inbox tags

Get all inbox tags

### Required Parameters

This endpoint does not need any parameter.

### Return type

**[]string**

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetInboxes

> []Inbox GetInboxes(ctx, optional)

List Inboxes and email eddresses

List the inboxes you have created. Note use of the more advanced `getAllEmails` is recommended. You can provide a limit and sort parameter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetInboxesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetInboxesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **optional.Int32**| Optional result size limit. Note an automatic limit of 100 results is applied. See the paginated &#x60;getAllEmails&#x60; for larger queries. | [default to 100]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]

### Return type

[**[]Inbox**](Inbox)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetOrganizationInboxes

> PageOrganizationInboxProjection GetOrganizationInboxes(ctx, optional)

List Organization Inboxes Paginated

List organization inboxes in paginated form. These are inboxes created with `allowTeamAccess` flag enabled. Organization inboxes are `readOnly` for non-admin users. The results are available on the `content` property of the returned object. This method allows for page index (zero based), page size (how many results to return), and a sort direction (based on createdAt time). 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetOrganizationInboxesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetOrganizationInboxesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Optional page index in inbox list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in inbox list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]

### Return type

[**PageOrganizationInboxProjection**](PageOrganizationInboxProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## SendEmail

> SendEmail(ctx, inboxId, optional)

Send Email

Send an email from an inbox's email address.  The request body should contain the `SendEmailOptions` that include recipients, attachments, body etc. See `SendEmailOptions` for all available properties. Note the `inboxId` refers to the inbox's id not the inbox's email address. See https://www.mailslurp.com/guides/ for more information on how to send emails. This method does not return a sent email entity due to legacy reasons. To send and get a sent email as returned response use the sister method `sendEmailAndConfirm`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()| ID of the inbox you want to send the email from | 
 **optional** | ***SendEmailOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SendEmailOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendEmailOptions** | [**optional.Interface of SendEmailOptions**](SendEmailOptions)| Options for the email | 

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


## SendEmailAndConfirm

> SentEmailDto SendEmailAndConfirm(ctx, inboxId, optional)

Send email and return sent confirmation

Sister method for standard `sendEmail` method with the benefit of returning a `SentEmail` entity confirming the successful sending of the email with a link to the sent object created for it.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()| ID of the inbox you want to send the email from | 
 **optional** | ***SendEmailAndConfirmOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SendEmailAndConfirmOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendEmailOptions** | [**optional.Interface of SendEmailOptions**](SendEmailOptions)| Options for the email | 

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


## SendTestEmail

> SendTestEmail(ctx, inboxId)

Send a test email to inbox

Send an inbox a test email to test email receiving is working

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()| inboxId | 

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


## SetInboxFavourited

> Inbox SetInboxFavourited(ctx, inboxId, setInboxFavouritedOptions)

Set inbox favourited state

Set and return new favourite state for an inbox

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()| inboxId | 
**setInboxFavouritedOptions** | [**SetInboxFavouritedOptions**](SetInboxFavouritedOptions)| setInboxFavouritedOptions | 

### Return type

[**Inbox**](Inbox)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## UpdateInbox

> Inbox UpdateInbox(ctx, inboxId, updateInboxOptions)

Update Inbox. Change name and description. Email address is not editable.

Update editable fields on an inbox

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()| inboxId | 
**updateInboxOptions** | [**UpdateInboxOptions**](UpdateInboxOptions)| updateInboxOptions | 

### Return type

[**Inbox**](Inbox)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

