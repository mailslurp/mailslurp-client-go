# MailSlurp\InboxControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelScheduledJob**](InboxControllerApi#CancelScheduledJob) | **Delete** /inboxes/scheduled-jobs/{jobId} | Cancel a scheduled email job
[**CreateInbox**](InboxControllerApi#CreateInbox) | **Post** /inboxes | Create an inbox email address. An inbox has a real email address and can send and receive emails. Inboxes can be either &#x60;SMTP&#x60; or &#x60;HTTP&#x60; inboxes.
[**CreateInboxRuleset**](InboxControllerApi#CreateInboxRuleset) | **Post** /inboxes/{inboxId}/rulesets | Create an inbox ruleset
[**CreateInboxWithDefaults**](InboxControllerApi#CreateInboxWithDefaults) | **Post** /inboxes/withDefaults | Create an inbox with default options. Uses MailSlurp domain pool address and is private.
[**CreateInboxWithOptions**](InboxControllerApi#CreateInboxWithOptions) | **Post** /inboxes/withOptions | Create an inbox with options. Extended options for inbox creation.
[**DeleteAllInboxEmails**](InboxControllerApi#DeleteAllInboxEmails) | **Delete** /inboxes/{inboxId}/deleteAllInboxEmails | Delete all emails in a given inboxes.
[**DeleteAllInboxes**](InboxControllerApi#DeleteAllInboxes) | **Delete** /inboxes | Delete all inboxes
[**DeleteInbox**](InboxControllerApi#DeleteInbox) | **Delete** /inboxes/{inboxId} | Delete inbox
[**DoesInboxExist**](InboxControllerApi#DoesInboxExist) | **Get** /inboxes/exists | Does inbox exist
[**FlushExpired**](InboxControllerApi#FlushExpired) | **Delete** /inboxes/expired | Remove expired inboxes
[**GetAllInboxes**](InboxControllerApi#GetAllInboxes) | **Get** /inboxes/paginated | List All Inboxes Paginated
[**GetAllScheduledJobs**](InboxControllerApi#GetAllScheduledJobs) | **Get** /inboxes/scheduled-jobs | Get all scheduled email sending jobs for account
[**GetDeliveryStatusesByInboxId**](InboxControllerApi#GetDeliveryStatusesByInboxId) | **Get** /inboxes/{inboxId}/delivery-status | 
[**GetEmails**](InboxControllerApi#GetEmails) | **Get** /inboxes/{inboxId}/emails | Get emails in an Inbox. This method is not idempotent as it allows retries and waits if you want certain conditions to be met before returning. For simple listing and sorting of known emails use the email controller instead.
[**GetImapSmtpAccess**](InboxControllerApi#GetImapSmtpAccess) | **Get** /inboxes/imap-smtp-access | 
[**GetInbox**](InboxControllerApi#GetInbox) | **Get** /inboxes/{inboxId} | Get Inbox. Returns properties of an inbox.
[**GetInboxByEmailAddress**](InboxControllerApi#GetInboxByEmailAddress) | **Get** /inboxes/byEmailAddress | Search for an inbox with the provided email address
[**GetInboxByName**](InboxControllerApi#GetInboxByName) | **Get** /inboxes/byName | Search for an inbox with the given name
[**GetInboxCount**](InboxControllerApi#GetInboxCount) | **Get** /inboxes/count | Get total inbox count
[**GetInboxEmailCount**](InboxControllerApi#GetInboxEmailCount) | **Get** /inboxes/{inboxId}/emails/count | Get email count in inbox
[**GetInboxEmailsPaginated**](InboxControllerApi#GetInboxEmailsPaginated) | **Get** /inboxes/{inboxId}/emails/paginated | Get inbox emails paginated
[**GetInboxIds**](InboxControllerApi#GetInboxIds) | **Get** /inboxes/ids | Get all inbox IDs
[**GetInboxSentEmails**](InboxControllerApi#GetInboxSentEmails) | **Get** /inboxes/{inboxId}/sent | Get Inbox Sent Emails
[**GetInboxTags**](InboxControllerApi#GetInboxTags) | **Get** /inboxes/tags | Get inbox tags
[**GetInboxes**](InboxControllerApi#GetInboxes) | **Get** /inboxes | List Inboxes and email addresses
[**GetLatestEmailInInbox**](InboxControllerApi#GetLatestEmailInInbox) | **Get** /inboxes/getLatestEmail | Get latest email in an inbox. Use &#x60;WaitForController&#x60; to get emails that may not have arrived yet.
[**GetOrganizationInboxes**](InboxControllerApi#GetOrganizationInboxes) | **Get** /inboxes/organization | List Organization Inboxes Paginated
[**GetScheduledJob**](InboxControllerApi#GetScheduledJob) | **Get** /inboxes/scheduled-jobs/{jobId} | Get a scheduled email job
[**GetScheduledJobsByInboxId**](InboxControllerApi#GetScheduledJobsByInboxId) | **Get** /inboxes/{inboxId}/scheduled-jobs | Get all scheduled email sending jobs for the inbox
[**ListInboxRulesets**](InboxControllerApi#ListInboxRulesets) | **Get** /inboxes/{inboxId}/rulesets | List inbox rulesets
[**ListInboxTrackingPixels**](InboxControllerApi#ListInboxTrackingPixels) | **Get** /inboxes/{inboxId}/tracking-pixels | List inbox tracking pixels
[**SendEmail**](InboxControllerApi#SendEmail) | **Post** /inboxes/{inboxId} | Send Email
[**SendEmailAndConfirm**](InboxControllerApi#SendEmailAndConfirm) | **Post** /inboxes/{inboxId}/confirm | Send email and return sent confirmation
[**SendEmailWithQueue**](InboxControllerApi#SendEmailWithQueue) | **Post** /inboxes/{inboxId}/with-queue | Send email with queue
[**SendSmtpEnvelope**](InboxControllerApi#SendSmtpEnvelope) | **Post** /inboxes/{inboxId}/smtp-envelope | Send email using an SMTP mail envelope and message body and return sent confirmation
[**SendTestEmail**](InboxControllerApi#SendTestEmail) | **Post** /inboxes/{inboxId}/send-test-email | Send a test email to inbox
[**SendWithSchedule**](InboxControllerApi#SendWithSchedule) | **Post** /inboxes/{inboxId}/with-schedule | Send email with with delay or schedule
[**SetInboxFavourited**](InboxControllerApi#SetInboxFavourited) | **Put** /inboxes/{inboxId}/favourite | Set inbox favourited state
[**UpdateInbox**](InboxControllerApi#UpdateInbox) | **Patch** /inboxes/{inboxId} | Update Inbox. Change name and description. Email address is not editable.



## CancelScheduledJob

> ScheduledJobDto CancelScheduledJob(ctx, jobId)

Cancel a scheduled email job

Get a scheduled email job and cancel it. Will fail if status of job is already cancelled, failed, or complete.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | [**string**]()|  | 

### Return type

[**ScheduledJobDto**](ScheduledJobDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CreateInbox

> InboxDto CreateInbox(ctx, optional)

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
 **emailAddress** | **optional.String**| A custom email address to use with the inbox. Defaults to null. When null MailSlurp will assign a random email address to the inbox such as &#x60;123@mailslurp.com&#x60;. If you use the &#x60;useDomainPool&#x60; option when the email address is null it will generate an email address with a more varied domain ending such as &#x60;123@mailslurp.info&#x60; or &#x60;123@mailslurp.biz&#x60;. When a custom email address is provided the address is split into a domain and the domain is queried against your user. If you have created the domain in the MailSlurp dashboard and verified it you can use any email address that ends with the domain. Note domain types must match the inbox type - so &#x60;SMTP&#x60; inboxes will only work with &#x60;SMTP&#x60; type domains. Avoid &#x60;SMTP&#x60; inboxes if you need to send emails as they can only receive. Send an email to this address and the inbox will receive and store it for you. To retrieve the email use the Inbox and Email Controller endpoints with the inbox ID. | 
 **tags** | [**optional.Interface of []string**](string)| Tags that inbox has been tagged with. Tags can be added to inboxes to group different inboxes within an account. You can also search for inboxes by tag in the dashboard UI. | 
 **name** | **optional.String**| Optional name of the inbox. Displayed in the dashboard for easier search and used as the sender name when sending emails. | 
 **description** | **optional.String**| Optional description of the inbox for labelling purposes. Is shown in the dashboard and can be used with | 
 **useDomainPool** | **optional.Bool**| Use the MailSlurp domain name pool with this inbox when creating the email address. Defaults to null. If enabled the inbox will be an email address with a domain randomly chosen from a list of the MailSlurp domains. This is useful when the default &#x60;@mailslurp.com&#x60; email addresses used with inboxes are blocked or considered spam by a provider or receiving service. When domain pool is enabled an email address will be generated ending in &#x60;@mailslurp.{world,info,xyz,...}&#x60; . This means a TLD is randomly selecting from a list of &#x60;.biz&#x60;, &#x60;.info&#x60;, &#x60;.xyz&#x60; etc to add variance to the generated email addresses. When null or false MailSlurp uses the default behavior of &#x60;@mailslurp.com&#x60; or custom email address provided by the emailAddress field. Note this feature is only available for &#x60;HTTP&#x60; inbox types. | 
 **favourite** | **optional.Bool**| Is the inbox a favorite. Marking an inbox as a favorite is typically done in the dashboard for quick access or filtering | 
 **expiresAt** | **optional.Time**| Optional inbox expiration date. If null then this inbox is permanent and the emails in it won&#39;t be deleted. If an expiration date is provided or is required by your plan the inbox will be closed when the expiration time is reached. Expired inboxes still contain their emails but can no longer send or receive emails. An ExpiredInboxRecord is created when an inbox and the email address and inbox ID are recorded. The expiresAt property is a timestamp string in ISO DateTime Format yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSXXX. | 
 **expiresIn** | **optional.Int64**| Number of milliseconds that inbox should exist for | 
 **allowTeamAccess** | **optional.Bool**| DEPRECATED (team access is always true). Grant team access to this inbox and the emails that belong to it for team members of your organization. | 
 **inboxType** | **optional.String**| HTTP (default) or SMTP inbox type. HTTP inboxes are default and best solution for most cases. SMTP inboxes are more reliable for public inbound email consumption (but do not support sending emails). When using custom domains the domain type must match the inbox type. HTTP inboxes are processed by AWS SES while SMTP inboxes use a custom mail server running at &#x60;mx.mailslurp.com&#x60;. | 
 **virtualInbox** | **optional.Bool**| Virtual inbox prevents any outbound emails from being sent. It creates sent email records but will never send real emails to recipients. Great for testing and faking email sending. | 
 **useShortAddress** | **optional.Bool**| Use a shorter email address under 31 characters | 

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


## CreateInboxRuleset

> InboxRulesetDto CreateInboxRuleset(ctx, inboxId, createInboxRulesetOptions)

Create an inbox ruleset

Create a new inbox rule for forwarding, blocking, and allowing emails when sending and receiving

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()| inboxId | 
**createInboxRulesetOptions** | [**CreateInboxRulesetOptions**](CreateInboxRulesetOptions)|  | 

### Return type

[**InboxRulesetDto**](InboxRulesetDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CreateInboxWithDefaults

> InboxDto CreateInboxWithDefaults(ctx, )

Create an inbox with default options. Uses MailSlurp domain pool address and is private.

### Required Parameters

This endpoint does not need any parameter.

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


## CreateInboxWithOptions

> InboxDto CreateInboxWithOptions(ctx, createInboxDto)

Create an inbox with options. Extended options for inbox creation.

Additional endpoint that allows inbox creation with request body options. Can be more flexible that other methods for some clients.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createInboxDto** | [**CreateInboxDto**](CreateInboxDto)|  | 

### Return type

[**InboxDto**](InboxDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DeleteAllInboxEmails

> DeleteAllInboxEmails(ctx, inboxId)

Delete all emails in a given inboxes.

Deletes all emails in an inbox. Be careful as emails cannot be recovered

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


## DoesInboxExist

> InboxExistsDto DoesInboxExist(ctx, emailAddress, optional)

Does inbox exist

Check if inboxes exist by email address. Useful if you are sending emails to mailslurp addresses

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailAddress** | **string**| Email address | 
 **optional** | ***DoesInboxExistOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DoesInboxExistOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **allowCatchAll** | **optional.Bool**|  | 

### Return type

[**InboxExistsDto**](InboxExistsDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## FlushExpired

> FlushExpiredInboxesResult FlushExpired(ctx, optional)

Remove expired inboxes

Remove any expired inboxes for your account (instead of waiting for scheduled removal on server)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FlushExpiredOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FlushExpiredOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **before** | **optional.Time**| Optional expired at before flag to flush expired inboxes that have expired before the given time | 

### Return type

[**FlushExpiredInboxesResult**](FlushExpiredInboxesResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetAllInboxes

> PageInboxProjection GetAllInboxes(ctx, optional)

List All Inboxes Paginated

List inboxes in paginated form. The results are available on the `content` property of the returned object. This method allows for page index (zero based), page size (how many results to return), and a sort direction (based on createdAt time). You Can also filter by whether an inbox is favorited or use email address pattern. This method is the recommended way to query inboxes. The alternative `getInboxes` method returns a full list of inboxes but is limited to 100 results.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllInboxesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllInboxesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Optional page index in list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **favourite** | **optional.Bool**| Optionally filter results for favourites only | [default to false]
 **search** | **optional.String**| Optionally filter by search words partial matching ID, tags, name, and email address | 
 **tag** | **optional.String**| Optionally filter by tags. Will return inboxes that include given tags | 
 **teamAccess** | **optional.Bool**| DEPRECATED. Optionally filter by team access. | 
 **since** | **optional.Time**| Optional filter by created after given date time | 
 **before** | **optional.Time**| Optional filter by created before given date time | 
 **inboxType** | **optional.String**| Optional filter by inbox type | 
 **domainId** | [**optional.Interface of string**]()| Optional domain ID filter | 

### Return type

[**PageInboxProjection**](PageInboxProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetAllScheduledJobs

> PageScheduledJobs GetAllScheduledJobs(ctx, optional)

Get all scheduled email sending jobs for account

Schedule sending of emails using scheduled jobs. These can be inbox or account level.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllScheduledJobsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllScheduledJobsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Optional page index in scheduled job list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in scheduled job list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **since** | **optional.Time**| Filter by created at after the given timestamp | 
 **before** | **optional.Time**| Filter by created at before the given timestamp | 

### Return type

[**PageScheduledJobs**](PageScheduledJobs)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDeliveryStatusesByInboxId

> PageDeliveryStatus GetDeliveryStatusesByInboxId(ctx, inboxId, optional)



Get all email delivery statuses for an inbox

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()|  | 
 **optional** | ***GetDeliveryStatusesByInboxIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDeliveryStatusesByInboxIdOpts struct


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

 **size** | **optional.Int32**| Alias for limit. Assessed first before assessing any passed limit. | 
 **limit** | **optional.Int32**| Limit the result set, ordered by received date time sort direction. Maximum 100. For more listing options see the email controller | 
 **sort** | **optional.String**| Sort the results by received date and direction ASC or DESC | 
 **retryTimeout** | **optional.Int64**| Maximum milliseconds to spend retrying inbox database until minCount emails are returned | 
 **delayTimeout** | **optional.Int64**|  | 
 **minCount** | **optional.Int64**| Minimum acceptable email count. Will cause request to hang (and retry) until minCount is satisfied or retryTimeout is reached. | 
 **unreadOnly** | **optional.Bool**|  | 
 **before** | **optional.Time**| Exclude emails received after this ISO 8601 date time | 
 **since** | **optional.Time**| Exclude emails received before this ISO 8601 date time | 

### Return type

[**[]EmailPreview**](EmailPreview)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetImapSmtpAccess

> ImapSmtpAccessDetails GetImapSmtpAccess(ctx, optional)



Get IMAP and SMTP access usernames and passwords

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetImapSmtpAccessOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetImapSmtpAccessOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inboxId** | [**optional.Interface of string**]()| Inbox ID | 

### Return type

[**ImapSmtpAccessDetails**](ImapSmtpAccessDetails)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetInbox

> InboxDto GetInbox(ctx, inboxId)

Get Inbox. Returns properties of an inbox.

Returns an inbox's properties, including its email address and ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()|  | 

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


## GetInboxByEmailAddress

> InboxByEmailAddressResult GetInboxByEmailAddress(ctx, emailAddress)

Search for an inbox with the provided email address

Get a inbox result by email address

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailAddress** | **string**|  | 

### Return type

[**InboxByEmailAddressResult**](InboxByEmailAddressResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetInboxByName

> InboxByNameResult GetInboxByName(ctx, name)

Search for an inbox with the given name

Get a inbox result by name

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**|  | 

### Return type

[**InboxByNameResult**](InboxByNameResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetInboxCount

> CountDto GetInboxCount(ctx, )

Get total inbox count

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


## GetInboxEmailCount

> CountDto GetInboxEmailCount(ctx, inboxId)

Get email count in inbox

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()| Id of inbox that emails belongs to | 

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
 **since** | **optional.Time**| Optional filter by received after given date time | 
 **before** | **optional.Time**| Optional filter by received before given date time | 

### Return type

[**PageEmailPreview**](PageEmailPreview)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetInboxIds

> InboxIdsResult GetInboxIds(ctx, )

Get all inbox IDs

Get list of inbox IDs

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InboxIdsResult**](InboxIdsResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

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
**inboxId** | [**string**]()|  | 
 **optional** | ***GetInboxSentEmailsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetInboxSentEmailsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Optional page index in inbox sent email list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in inbox sent email list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **searchFilter** | **optional.String**| Optional sent email search | 
 **since** | **optional.Time**| Optional filter by sent after given date time | 
 **before** | **optional.Time**| Optional filter by sent before given date time | 

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
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetInboxes

> []InboxDto GetInboxes(ctx, optional)

List Inboxes and email addresses

List the inboxes you have created. Note use of the more advanced `getAllInboxes` is recommended and allows paginated access using a limit and sort parameter.

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
 **since** | **optional.Time**| Optional filter by created after given date time | 
 **before** | **optional.Time**| Optional filter by created before given date time | 

### Return type

[**[]InboxDto**](InboxDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetLatestEmailInInbox

> Email GetLatestEmailInInbox(ctx, inboxId, timeoutMillis)

Get latest email in an inbox. Use `WaitForController` to get emails that may not have arrived yet.

Get the newest email in an inbox or wait for one to arrive

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()| ID of the inbox you want to get the latest email from | 
**timeoutMillis** | **int64**| Timeout milliseconds to wait for latest email | 

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
 **page** | **optional.Int32**| Optional page index in list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **searchFilter** | **optional.String**| Optional search filter | 
 **since** | **optional.Time**| Optional filter by created after given date time | 
 **before** | **optional.Time**| Optional filter by created before given date time | 

### Return type

[**PageOrganizationInboxProjection**](PageOrganizationInboxProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetScheduledJob

> ScheduledJobDto GetScheduledJob(ctx, jobId)

Get a scheduled email job

Get a scheduled email job details.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | [**string**]()|  | 

### Return type

[**ScheduledJobDto**](ScheduledJobDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetScheduledJobsByInboxId

> PageScheduledJobs GetScheduledJobsByInboxId(ctx, inboxId, optional)

Get all scheduled email sending jobs for the inbox

Schedule sending of emails using scheduled jobs.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()|  | 
 **optional** | ***GetScheduledJobsByInboxIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetScheduledJobsByInboxIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Optional page index in scheduled job list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in scheduled job list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **since** | **optional.Time**| Filter by created at after the given timestamp | 
 **before** | **optional.Time**| Filter by created at before the given timestamp | 

### Return type

[**PageScheduledJobs**](PageScheduledJobs)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## ListInboxRulesets

> PageInboxRulesetDto ListInboxRulesets(ctx, inboxId, optional)

List inbox rulesets

List all rulesets attached to an inbox

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()|  | 
 **optional** | ***ListInboxRulesetsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListInboxRulesetsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Optional page index in inbox ruleset list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in inbox ruleset list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **searchFilter** | **optional.String**| Optional search filter | 
 **since** | **optional.Time**| Optional filter by created after given date time | 
 **before** | **optional.Time**| Optional filter by created before given date time | 

### Return type

[**PageInboxRulesetDto**](PageInboxRulesetDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## ListInboxTrackingPixels

> PageTrackingPixelProjection ListInboxTrackingPixels(ctx, inboxId, optional)

List inbox tracking pixels

List all tracking pixels sent from an inbox

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()|  | 
 **optional** | ***ListInboxTrackingPixelsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListInboxTrackingPixelsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Optional page index in inbox tracking pixel list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in inbox tracking pixel list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **searchFilter** | **optional.String**| Optional search filter | 
 **since** | **optional.Time**| Optional filter by created after given date time | 
 **before** | **optional.Time**| Optional filter by created before given date time | 

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


## SendEmail

> SendEmail(ctx, inboxId, sendEmailOptions)

Send Email

Send an email from an inbox's email address.  The request body should contain the `SendEmailOptions` that include recipients, attachments, body etc. See `SendEmailOptions` for all available properties. Note the `inboxId` refers to the inbox's id not the inbox's email address. See https://www.mailslurp.com/guides/ for more information on how to send emails. This method does not return a sent email entity due to legacy reasons. To send and get a sent email as returned response use the sister method `sendEmailAndConfirm`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()| ID of the inbox you want to send the email from | 
**sendEmailOptions** | [**SendEmailOptions**](SendEmailOptions)|  | 

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

> SentEmailDto SendEmailAndConfirm(ctx, inboxId, sendEmailOptions)

Send email and return sent confirmation

Sister method for standard `sendEmail` method with the benefit of returning a `SentEmail` entity confirming the successful sending of the email with a link to the sent object created for it.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()| ID of the inbox you want to send the email from | 
**sendEmailOptions** | [**SendEmailOptions**](SendEmailOptions)|  | 

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


## SendEmailWithQueue

> SendEmailWithQueue(ctx, inboxId, validateBeforeEnqueue, sendEmailOptions)

Send email with queue

Send an email using a queue. Will place the email onto a queue that will then be processed and sent. Use this queue method to enable any failed email sending to be recovered. This will prevent lost emails when sending if your account encounters a block or payment issue.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()| ID of the inbox you want to send the email from | 
**validateBeforeEnqueue** | **bool**| Validate before adding to queue | 
**sendEmailOptions** | [**SendEmailOptions**](SendEmailOptions)|  | 

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


## SendSmtpEnvelope

> SentEmailDto SendSmtpEnvelope(ctx, inboxId, sendSmtpEnvelopeOptions)

Send email using an SMTP mail envelope and message body and return sent confirmation

Send email using an SMTP envelope containing RCPT TO, MAIL FROM, and a SMTP BODY.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()| ID of the inbox you want to send the email from | 
**sendSmtpEnvelopeOptions** | [**SendSmtpEnvelopeOptions**](SendSmtpEnvelopeOptions)|  | 

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


## SendTestEmail

> SendTestEmail(ctx, inboxId)

Send a test email to inbox

Send an inbox a test email to test email receiving is working

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


## SendWithSchedule

> ScheduledJobDto SendWithSchedule(ctx, inboxId, sendEmailOptions, optional)

Send email with with delay or schedule

Send an email using a delay. Will place the email onto a scheduler that will then be processed and sent. Use delays to schedule email sending.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()| ID of the inbox you want to send the email from | 
**sendEmailOptions** | [**SendEmailOptions**](SendEmailOptions)|  | 
 **optional** | ***SendWithScheduleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SendWithScheduleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sendAtTimestamp** | **optional.Time**| Sending timestamp | 
 **sendAtNowPlusSeconds** | **optional.Int64**| Send after n seconds | 
 **validateBeforeEnqueue** | **optional.Bool**| Validate before adding to queue | 

### Return type

[**ScheduledJobDto**](ScheduledJobDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## SetInboxFavourited

> InboxDto SetInboxFavourited(ctx, inboxId, setInboxFavouritedOptions)

Set inbox favourited state

Set and return new favourite state for an inbox

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()| ID of inbox to set favourite state | 
**setInboxFavouritedOptions** | [**SetInboxFavouritedOptions**](SetInboxFavouritedOptions)|  | 

### Return type

[**InboxDto**](InboxDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## UpdateInbox

> InboxDto UpdateInbox(ctx, inboxId, updateInboxOptions)

Update Inbox. Change name and description. Email address is not editable.

Update editable fields on an inbox

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()|  | 
**updateInboxOptions** | [**UpdateInboxOptions**](UpdateInboxOptions)|  | 

### Return type

[**InboxDto**](InboxDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

