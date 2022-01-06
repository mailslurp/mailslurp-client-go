# MailSlurp\InboxControllerApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInbox**](InboxControllerApi#CreateInbox) | **Post** /inboxes | Create an inbox email address. An inbox has a real email address and can send and receive emails. Inboxes can be either &#x60;SMTP&#x60; or &#x60;HTTP&#x60; inboxes.
[**CreateInboxRuleset**](InboxControllerApi#CreateInboxRuleset) | **Post** /inboxes/{inboxId}/rulesets | Create an inbox ruleset
[**CreateInboxWithDefaults**](InboxControllerApi#CreateInboxWithDefaults) | **Post** /inboxes/withDefaults | Create an inbox with default options. Uses MailSlurp domain pool address and is private.
[**CreateInboxWithOptions**](InboxControllerApi#CreateInboxWithOptions) | **Post** /inboxes/withOptions | Create an inbox with options. Extended options for inbox creation.
[**DeleteAllInboxes**](InboxControllerApi#DeleteAllInboxes) | **Delete** /inboxes | Delete all inboxes
[**DeleteInbox**](InboxControllerApi#DeleteInbox) | **Delete** /inboxes/{inboxId} | Delete inbox
[**DoesInboxExist**](InboxControllerApi#DoesInboxExist) | **Get** /inboxes/exists | Does inbox exist
[**FlushExpired**](InboxControllerApi#FlushExpired) | **Delete** /inboxes/expired | Remove expired inboxes
[**GetAllInboxes**](InboxControllerApi#GetAllInboxes) | **Get** /inboxes/paginated | List All Inboxes Paginated
[**GetEmails**](InboxControllerApi#GetEmails) | **Get** /inboxes/{inboxId}/emails | Get emails in an Inbox. This method is not idempotent as it allows retries and waits if you want certain conditions to be met before returning. For simple listing and sorting of known emails use the email controller instead.
[**GetInbox**](InboxControllerApi#GetInbox) | **Get** /inboxes/{inboxId} | Get Inbox. Returns properties of an inbox.
[**GetInboxCount**](InboxControllerApi#GetInboxCount) | **Get** /inboxes/count | Get total inbox count
[**GetInboxEmailCount**](InboxControllerApi#GetInboxEmailCount) | **Get** /inboxes/{inboxId}/emails/count | Get email count in inbox
[**GetInboxEmailsPaginated**](InboxControllerApi#GetInboxEmailsPaginated) | **Get** /inboxes/{inboxId}/emails/paginated | Get inbox emails paginated
[**GetInboxSentEmails**](InboxControllerApi#GetInboxSentEmails) | **Get** /inboxes/{inboxId}/sent | Get Inbox Sent Emails
[**GetInboxTags**](InboxControllerApi#GetInboxTags) | **Get** /inboxes/tags | Get inbox tags
[**GetInboxes**](InboxControllerApi#GetInboxes) | **Get** /inboxes | List Inboxes and email addresses
[**GetOrganizationInboxes**](InboxControllerApi#GetOrganizationInboxes) | **Get** /inboxes/organization | List Organization Inboxes Paginated
[**ListInboxRulesets**](InboxControllerApi#ListInboxRulesets) | **Get** /inboxes/{inboxId}/rulesets | List inbox rulesets
[**ListInboxTrackingPixels**](InboxControllerApi#ListInboxTrackingPixels) | **Get** /inboxes/{inboxId}/tracking-pixels | List inbox tracking pixels
[**SendEmail**](InboxControllerApi#SendEmail) | **Post** /inboxes/{inboxId} | Send Email
[**SendEmailAndConfirm**](InboxControllerApi#SendEmailAndConfirm) | **Post** /inboxes/{inboxId}/confirm | Send email and return sent confirmation
[**SendTestEmail**](InboxControllerApi#SendTestEmail) | **Post** /inboxes/{inboxId}/send-test-email | Send a test email to inbox
[**SetInboxFavourited**](InboxControllerApi#SetInboxFavourited) | **Put** /inboxes/{inboxId}/favourite | Set inbox favourited state
[**UpdateInbox**](InboxControllerApi#UpdateInbox) | **Patch** /inboxes/{inboxId} | Update Inbox. Change name and description. Email address is not editable.



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
 **aCustomEmailAddressToUseWithTheInboxDefaultsToNullWhenNullMailSlurpWillAssignARandomEmailAddressToTheInboxSuchAs123mailslurpComIfYouUseTheUseDomainPoolOptionWhenTheEmailAddressIsNullItWillGenerateAnEmailAddressWithAMoreVariedDomainEndingSuchAs123mailslurpInfoOr123mailslurpBizWhenACustomEmailAddressIsProvidedTheAddressIsSplitIntoADomainAndTheDomainIsQueriedAgainstYourUserIfYouHaveCreatedTheDomainInTheMailSlurpDashboardAndVerifiedItYouCanUseAnyEmailAddressThatEndsWithTheDomainNoteDomainTypesMustMatchTheInboxTypeSoSMTPInboxesWillOnlyWorkWithSMTPTypeDomainsAvoidSMTPInboxesIfYouNeedToSendEmailsAsTheyCanOnlyReceiveSendAnEmailToThisAddressAndTheInboxWillReceiveAndStoreItForYouToRetrieveTheEmailUseTheInboxAndEmailControllerEndpointsWithTheInboxID** | **optional.String**|  | 
 **tagsThatInboxHasBeenTaggedWithTagsCanBeAddedToInboxesToGroupDifferentInboxesWithinAnAccountYouCanAlsoSearchForInboxesByTagInTheDashboardUI** | [**optional.Interface of []string**](string)|  | 
 **optionalNameOfTheInboxDisplayedInTheDashboardForEasierSearchAndUsedAsTheSenderNameWhenSendingEmails** | **optional.String**|  | 
 **optionalDescriptionOfTheInboxForLabellingPurposesIsShownInTheDashboardAndCanBeUsedWith** | **optional.String**|  | 
 **useTheMailSlurpDomainNamePoolWithThisInboxWhenCreatingTheEmailAddressDefaultsToNullIfEnabledTheInboxWillBeAnEmailAddressWithADomainRandomlyChosenFromAListOfTheMailSlurpDomainsThisIsUsefulWhenTheDefaultMailslurpComEmailAddressesUsedWithInboxesAreBlockedOrConsideredSpamByAProviderOrReceivingServiceWhenDomainPoolIsEnabledAnEmailAddressWillBeGeneratedEndingInMailslurpWorldinfoxyzThisMeansATLDIsRandomlySelectingFromAListOfBizInfoXyzEtcToAddVarianceToTheGeneratedEmailAddressesWhenNullOrFalseMailSlurpUsesTheDefaultBehaviorOfMailslurpComOrCustomEmailAddressProvidedByTheEmailAddressFieldNoteThisFeatureIsOnlyAvailableForHTTPInboxTypes** | **optional.Bool**|  | 
 **isTheInboxAFavoriteMarkingAnInboxAsAFavoriteIsTypicallyDoneInTheDashboardForQuickAccessOrFiltering** | **optional.Bool**|  | 
 **optionalInboxExpirationDateIfNullThenThisInboxIsPermanentAndTheEmailsInItWontBeDeletedIfAnExpirationDateIsProvidedOrIsRequiredByYourPlanTheInboxWillBeClosedWhenTheExpirationTimeIsReachedExpiredInboxesStillContainTheirEmailsButCanNoLongerSendOrReceiveEmailsAnExpiredInboxRecordIsCreatedWhenAnInboxAndTheEmailAddressAndInboxIDAreRecordedTheExpiresAtPropertyIsATimestampStringInISODateTimeFormatYyyyMMDdTHHmmssSSSXXX** | **optional.Time**|  | 
 **numberOfMillisecondsThatInboxShouldExistFor** | **optional.Int64**|  | 
 **dEPRECATEDTeamAccessIsAlwaysTrueGrantTeamAccessToThisInboxAndTheEmailsThatBelongToItForTeamMembersOfYourOrganization** | **optional.Bool**|  | 
 **hTTPDefaultOrSMTPInboxTypeHTTPInboxesAreDefaultAndBestSolutionForMostCasesSMTPInboxesAreMoreReliableForPublicInboundEmailConsumptionButDoNotSupportSendingEmailsWhenUsingCustomDomainsTheDomainTypeMustMatchTheInboxTypeHTTPInboxesAreProcessedByAWSSESWhileSMTPInboxesUseACustomMailServerRunningAtMxMailslurpCom** | **optional.String**|  | 

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
**inboxId** | [**string**]()|  | 
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

> InboxExistsDto DoesInboxExist(ctx, emailAddress)

Does inbox exist

Check if inboxes exist by email address. Useful if you are sending emails to mailslurp addresses

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailAddress** | **string**| Email address | 

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
 **dEPRECATEDOptionallyFilterByTeamAccess** | **optional.Bool**|  | 
 **since** | **optional.Time**| Optional filter by created after given date time | 
 **before** | **optional.Time**| Optional filter by created before given date time | 

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


## GetEmails

> []EmailPreview GetEmails(ctx, idOfInboxThatEmailsBelongsTo, optional)

Get emails in an Inbox. This method is not idempotent as it allows retries and waits if you want certain conditions to be met before returning. For simple listing and sorting of known emails use the email controller instead.

List emails that an inbox has received. Only emails that are sent to the inbox's email address will appear in the inbox. It may take several seconds for any email you send to an inbox's email address to appear in the inbox. To make this endpoint wait for a minimum number of emails use the `minCount` parameter. The server will retry the inbox database until the `minCount` is satisfied or the `retryTimeout` is reached

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOfInboxThatEmailsBelongsTo** | [**string**]()|  | 
 **optional** | ***GetEmailsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEmailsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aliasForLimitAssessedFirstBeforeAssessingAnyPassedLimit** | **optional.Int32**|  | 
 **limit** | **optional.Int32**| Limit the result set, ordered by received date time sort direction. Maximum 100. For more listing options see the email controller | 
 **sortTheResultsByReceivedDateAndDirectionASCOrDESC** | **optional.String**|  | 
 **retryTimeout** | **optional.Int64**| Maximum milliseconds to spend retrying inbox database until minCount emails are returned | 
 **delayTimeout** | **optional.Int64**|  | 
 **minCount** | **optional.Int64**| Minimum acceptable email count. Will cause request to hang (and retry) until minCount is satisfied or retryTimeout is reached. | 
 **unreadOnly** | **optional.Bool**|  | 
 **excludeEmailsReceivedAfterThisISO8601DateTime** | **optional.Time**|  | 
 **excludeEmailsReceivedBeforeThisISO8601DateTime** | **optional.Time**|  | 

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

> CountDto GetInboxEmailCount(ctx, idOfInboxThatEmailsBelongsTo)

Get email count in inbox

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOfInboxThatEmailsBelongsTo** | [**string**]()|  | 

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

> PageEmailPreview GetInboxEmailsPaginated(ctx, idOfInboxThatEmailsBelongsTo, optional)

Get inbox emails paginated

Get a paginated list of emails in an inbox. Does not hold connections open.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOfInboxThatEmailsBelongsTo** | [**string**]()|  | 
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
 **optionalPageSizeInInboxSentEmailListPagination** | **optional.Int32**|  | [default to 20]
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

List the inboxes you have created. Note use of the more advanced `getAllEmails` is recommended and allows paginated access using a limit and sort parameter.

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
 **optionalPageSizeInInboxTrackingPixelListPagination** | **optional.Int32**|  | [default to 20]
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

> SendEmail(ctx, iDOfTheInboxYouWantToSendTheEmailFrom, sendEmailOptions)

Send Email

Send an email from an inbox's email address.  The request body should contain the `SendEmailOptions` that include recipients, attachments, body etc. See `SendEmailOptions` for all available properties. Note the `inboxId` refers to the inbox's id not the inbox's email address. See https://www.mailslurp.com/guides/ for more information on how to send emails. This method does not return a sent email entity due to legacy reasons. To send and get a sent email as returned response use the sister method `sendEmailAndConfirm`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iDOfTheInboxYouWantToSendTheEmailFrom** | [**string**]()|  | 
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

> SentEmailDto SendEmailAndConfirm(ctx, iDOfTheInboxYouWantToSendTheEmailFrom, sendEmailOptions)

Send email and return sent confirmation

Sister method for standard `sendEmail` method with the benefit of returning a `SentEmail` entity confirming the successful sending of the email with a link to the sent object created for it.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iDOfTheInboxYouWantToSendTheEmailFrom** | [**string**]()|  | 
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


## SetInboxFavourited

> InboxDto SetInboxFavourited(ctx, inboxId, setInboxFavouritedOptions)

Set inbox favourited state

Set and return new favourite state for an inbox

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()|  | 
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

