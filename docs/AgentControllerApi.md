# MailSlurp\AgentControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentAcquireConversationClaim**](AgentControllerApi#AgentAcquireConversationClaim) | **Post** /agent/conversations/{conversationId}/claims | Acquire a renewable conversation claim
[**AgentAddConversationNote**](AgentControllerApi#AgentAddConversationNote) | **Post** /agent/conversations/{conversationId}/notes | Add an internal conversation note
[**AgentAddSmsNote**](AgentControllerApi#AgentAddSmsNote) | **Post** /agent/sms-conversations/{conversationId}/notes | Add an internal SMS conversation note
[**AgentCheckInboxSendPolicy**](AgentControllerApi#AgentCheckInboxSendPolicy) | **Post** /agent/inboxes/{inboxId}/send-policy/check | Check whether sending rules allow recipients
[**AgentClaimSmsConversation**](AgentControllerApi#AgentClaimSmsConversation) | **Post** /agent/sms-conversations/{conversationId}/claims | Acquire an SMS conversation claim
[**AgentCompleteWorkItem**](AgentControllerApi#AgentCompleteWorkItem) | **Post** /agent/work-items/{conversationId}/complete | Complete a work item with an idempotent reply
[**AgentCreateEmailDraft**](AgentControllerApi#AgentCreateEmailDraft) | **Post** /agent/inboxes/{inboxId}/drafts | Create a scoped email draft for an agent
[**AgentCreateInbox**](AgentControllerApi#AgentCreateInbox) | **Post** /agent/inboxes | Create an inbox with an account-wide inbox-manager connection
[**AgentCreateReplyDraft**](AgentControllerApi#AgentCreateReplyDraft) | **Post** /agent/emails/{emailId}/reply-draft | Create a correctly threaded reply draft
[**AgentCreateTotpDeviceForBase32SecretKey**](AgentControllerApi#AgentCreateTotpDeviceForBase32SecretKey) | **Post** /agent/totp/devices/base32-secret-key | Create an agent-managed TOTP device from a base32 secret key
[**AgentCreateTotpDeviceForCustom**](AgentControllerApi#AgentCreateTotpDeviceForCustom) | **Post** /agent/totp/devices/custom | Create an agent-managed TOTP device from custom options
[**AgentCreateTotpDeviceForOtpAuthUrl**](AgentControllerApi#AgentCreateTotpDeviceForOtpAuthUrl) | **Post** /agent/totp/devices/otp-auth-url | Create an agent-managed TOTP device from an OTP Auth URL
[**AgentDeferWorkItem**](AgentControllerApi#AgentDeferWorkItem) | **Post** /agent/work-items/{conversationId}/defer | Defer a work item to human review
[**AgentDeleteEmailDraft**](AgentControllerApi#AgentDeleteEmailDraft) | **Delete** /agent/inboxes/{inboxId}/drafts/{draftId} | Delete a scoped email draft for an agent
[**AgentDeleteTotpDevice**](AgentControllerApi#AgentDeleteTotpDevice) | **Delete** /agent/totp/devices/{id} | Delete an agent-managed TOTP device
[**AgentExtractEmailAttachmentText**](AgentControllerApi#AgentExtractEmailAttachmentText) | **Post** /agent/emails/{emailId}/attachments/{attachmentId}/text | Extract text from a scoped email attachment for an agent
[**AgentGenerateTotpDeviceCode**](AgentControllerApi#AgentGenerateTotpDeviceCode) | **Get** /agent/totp/devices/{id}/code | Generate a code for an agent-managed TOTP device
[**AgentGetCapabilities**](AgentControllerApi#AgentGetCapabilities) | **Get** /agent/capabilities | Discover this agent connection&#39;s capabilities
[**AgentGetConversationV2**](AgentControllerApi#AgentGetConversationV2) | **Get** /agent/conversations/{conversationId} | Get one scoped conversation
[**AgentGetEmail**](AgentControllerApi#AgentGetEmail) | **Get** /agent/emails/{emailId} | Get scoped hydrated email for an agent
[**AgentGetEmailDraft**](AgentControllerApi#AgentGetEmailDraft) | **Get** /agent/inboxes/{inboxId}/drafts/{draftId} | Get a scoped email draft for an agent
[**AgentGetEmailSummary**](AgentControllerApi#AgentGetEmailSummary) | **Get** /agent/emails/{emailId}/summary | Get scoped email summary for an agent
[**AgentGetEmailThread**](AgentControllerApi#AgentGetEmailThread) | **Get** /agent/emails/{emailId}/thread | Get a scoped email thread for an agent
[**AgentGetLatestInboxEmail**](AgentControllerApi#AgentGetLatestInboxEmail) | **Get** /agent/inboxes/{inboxId}/emails/latest | Get latest scoped inbox email for an agent
[**AgentGetNextWorkItem**](AgentControllerApi#AgentGetNextWorkItem) | **Post** /agent/work-items/next | Claim the next conversation work item
[**AgentGetPhoneMessageThreadItems**](AgentControllerApi#AgentGetPhoneMessageThreadItems) | **Get** /agent/phone-numbers/{phoneNumberId}/message-threads/{otherPhoneNumber} | List SMS thread messages
[**AgentGetPhoneMessageThreads**](AgentControllerApi#AgentGetPhoneMessageThreads) | **Get** /agent/phone-numbers/{phoneNumberId}/message-threads | List SMS threads
[**AgentGetPhoneNumber**](AgentControllerApi#AgentGetPhoneNumber) | **Get** /agent/phone-numbers/{phoneNumberId} | Get an assigned phone number
[**AgentGetPhoneNumbers**](AgentControllerApi#AgentGetPhoneNumbers) | **Get** /agent/phone-numbers | List assigned phone numbers
[**AgentGetReceivedSmsMessage**](AgentControllerApi#AgentGetReceivedSmsMessage) | **Get** /agent/phone-numbers/{phoneNumberId}/sms/{smsId} | Get received SMS
[**AgentGetReceivedSmsMessages**](AgentControllerApi#AgentGetReceivedSmsMessages) | **Get** /agent/phone-numbers/{phoneNumberId}/sms | List received SMS for an assigned phone number
[**AgentGetSentSmsMessage**](AgentControllerApi#AgentGetSentSmsMessage) | **Get** /agent/phone-numbers/{phoneNumberId}/sms/sent/{sentSmsId} | Get sent SMS
[**AgentGetSentSmsMessages**](AgentControllerApi#AgentGetSentSmsMessages) | **Get** /agent/phone-numbers/{phoneNumberId}/sms/sent | List sent SMS for an assigned phone number
[**AgentGetSmsConversation**](AgentControllerApi#AgentGetSmsConversation) | **Get** /agent/sms-conversations/{conversationId} | Get one SMS conversation
[**AgentGetTotpDevice**](AgentControllerApi#AgentGetTotpDevice) | **Get** /agent/totp/devices/{id} | Get an agent-accessible TOTP device
[**AgentGetTotpDeviceBy**](AgentControllerApi#AgentGetTotpDeviceBy) | **Get** /agent/totp/devices/by | Find an agent-accessible TOTP device by name, issuer, or username
[**AgentGetTotpDevices**](AgentControllerApi#AgentGetTotpDevices) | **Get** /agent/totp/devices | List agent-accessible TOTP devices
[**AgentListAccessibleInboxes**](AgentControllerApi#AgentListAccessibleInboxes) | **Get** /agent/inboxes | List inboxes accessible to an agent
[**AgentListAutomationEvents**](AgentControllerApi#AgentListAutomationEvents) | **Get** /agent/events | Read durable scoped automation events
[**AgentListConversationsV2**](AgentControllerApi#AgentListConversationsV2) | **Get** /agent/conversations | List active conversations with an opaque keyset cursor
[**AgentListEmailAttachments**](AgentControllerApi#AgentListEmailAttachments) | **Get** /agent/emails/{emailId}/attachments | List attachments on a scoped email for an agent
[**AgentListEmailDrafts**](AgentControllerApi#AgentListEmailDrafts) | **Get** /agent/inboxes/{inboxId}/drafts | List scoped email drafts for an agent
[**AgentListInboxEmails**](AgentControllerApi#AgentListInboxEmails) | **Get** /agent/inboxes/{inboxId}/emails | List scoped inbox emails for an agent
[**AgentListInboxThreads**](AgentControllerApi#AgentListInboxThreads) | **Get** /agent/inboxes/{inboxId}/threads | List scoped inbox threads for an agent
[**AgentListSmsConversations**](AgentControllerApi#AgentListSmsConversations) | **Get** /agent/sms-conversations | List active SMS conversations
[**AgentPrepareWorkItemReply**](AgentControllerApi#AgentPrepareWorkItemReply) | **Post** /agent/work-items/{conversationId}/prepare-reply | Prepare a guarded reply draft
[**AgentReleaseConversationClaim**](AgentControllerApi#AgentReleaseConversationClaim) | **Delete** /agent/conversations/{conversationId}/claims/{claimId} | Release a conversation claim
[**AgentReleaseSmsClaim**](AgentControllerApi#AgentReleaseSmsClaim) | **Delete** /agent/sms-conversations/{conversationId}/claims/{claimId} | Release an SMS conversation claim
[**AgentRenewConversationClaim**](AgentControllerApi#AgentRenewConversationClaim) | **Put** /agent/conversations/{conversationId}/claims/{claimId} | Renew a conversation claim
[**AgentRenewSmsClaim**](AgentControllerApi#AgentRenewSmsClaim) | **Put** /agent/sms-conversations/{conversationId}/claims/{claimId} | Renew an SMS conversation claim
[**AgentReplyToEmail**](AgentControllerApi#AgentReplyToEmail) | **Post** /agent/emails/{emailId}/reply | Reply to a scoped email as an agent
[**AgentReplyToSms**](AgentControllerApi#AgentReplyToSms) | **Post** /agent/phone-numbers/{phoneNumberId}/sms/{smsId}/reply | Reply to received SMS
[**AgentReportWorkerHeartbeat**](AgentControllerApi#AgentReportWorkerHeartbeat) | **Post** /agent/runtime/heartbeat | Report a worker heartbeat
[**AgentRequestConversationHandoff**](AgentControllerApi#AgentRequestConversationHandoff) | **Post** /agent/conversations/{conversationId}/handoffs | Request human review for a conversation
[**AgentRequestSmsHandoff**](AgentControllerApi#AgentRequestSmsHandoff) | **Post** /agent/sms-conversations/{conversationId}/handoffs | Request human review for an SMS conversation
[**AgentSearchInboxEmails**](AgentControllerApi#AgentSearchInboxEmails) | **Post** /agent/inboxes/{inboxId}/emails/search | Search scoped inbox emails for an agent
[**AgentSendEmailDraft**](AgentControllerApi#AgentSendEmailDraft) | **Post** /agent/inboxes/{inboxId}/drafts/{draftId}/send | Send a scoped email draft for an agent
[**AgentSendSms**](AgentControllerApi#AgentSendSms) | **Post** /agent/phone-numbers/{phoneNumberId}/sms | Send SMS from an assigned phone number
[**AgentUpdateEmailDraft**](AgentControllerApi#AgentUpdateEmailDraft) | **Patch** /agent/inboxes/{inboxId}/drafts/{draftId} | Update a scoped email draft for an agent
[**AgentUpdateInbox**](AgentControllerApi#AgentUpdateInbox) | **Patch** /agent/inboxes/{inboxId} | Update an inbox visible to an inbox-manager connection
[**AgentUpdateTotpDevice**](AgentControllerApi#AgentUpdateTotpDevice) | **Put** /agent/totp/devices/{id} | Update an agent-managed TOTP device



## AgentAcquireConversationClaim

> AgentConversationClaimDto AgentAcquireConversationClaim(ctx, conversationId, acquireAgentConversationClaimOptions)

Acquire a renewable conversation claim

Atomically leases a conversation to one worker. A conflict means another unexpired worker owns it. Use expectedVersion to reject work selected from stale context.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | [**string**]()|  | 
**acquireAgentConversationClaimOptions** | [**AcquireAgentConversationClaimOptions**](AcquireAgentConversationClaimOptions)|  | 

### Return type

[**AgentConversationClaimDto**](AgentConversationClaimDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentAddConversationNote

> AgentConversationNoteDto AgentAddConversationNote(ctx, conversationId, createAgentConversationNoteOptions)

Add an internal conversation note

Adds agent/operator collaboration context without sending any email to the customer.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | [**string**]()|  | 
**createAgentConversationNoteOptions** | [**CreateAgentConversationNoteOptions**](CreateAgentConversationNoteOptions)|  | 

### Return type

[**AgentConversationNoteDto**](AgentConversationNoteDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentAddSmsNote

> AgentConversationNoteDto AgentAddSmsNote(ctx, conversationId, createAgentSmsConversationNoteOptions)

Add an internal SMS conversation note

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | [**string**]()|  | 
**createAgentSmsConversationNoteOptions** | [**CreateAgentSmsConversationNoteOptions**](CreateAgentSmsConversationNoteOptions)|  | 

### Return type

[**AgentConversationNoteDto**](AgentConversationNoteDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentCheckInboxSendPolicy

> AgentInboxSendPolicyDto AgentCheckInboxSendPolicy(ctx, inboxId, agentInboxSendPolicyCheckOptions)

Check whether sending rules allow recipients

Evaluates effective inbox and account sending rules before an agent creates a send action. A blocked result includes the operator recovery route.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()|  | 
**agentInboxSendPolicyCheckOptions** | [**AgentInboxSendPolicyCheckOptions**](AgentInboxSendPolicyCheckOptions)|  | 

### Return type

[**AgentInboxSendPolicyDto**](AgentInboxSendPolicyDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentClaimSmsConversation

> AgentConversationClaimDto AgentClaimSmsConversation(ctx, conversationId, acquireAgentConversationClaimOptions)

Acquire an SMS conversation claim

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | [**string**]()|  | 
**acquireAgentConversationClaimOptions** | [**AcquireAgentConversationClaimOptions**](AcquireAgentConversationClaimOptions)|  | 

### Return type

[**AgentConversationClaimDto**](AgentConversationClaimDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentCompleteWorkItem

> AgentCompletedWorkItemDto AgentCompleteWorkItem(ctx, conversationId, completeAgentWorkItemOptions)

Complete a work item with an idempotent reply

Revalidates conversation freshness, sends one threaded reply with an idempotency key, and releases the supplied claim only after the send succeeds.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | [**string**]()|  | 
**completeAgentWorkItemOptions** | [**CompleteAgentWorkItemOptions**](CompleteAgentWorkItemOptions)|  | 

### Return type

[**AgentCompletedWorkItemDto**](AgentCompletedWorkItemDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentCreateEmailDraft

> EmailDraftDto AgentCreateEmailDraft(ctx, inboxId, sendEmailOptions, optional)

Create a scoped email draft for an agent

Creates a saved draft in the path inbox using the same SendEmailOptions shape as sending, without sending the email.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()| ID of inbox to save the draft against | 
**sendEmailOptions** | [**SendEmailOptions**](SendEmailOptions)|  | 
 **optional** | ***AgentCreateEmailDraftOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AgentCreateEmailDraftOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sourceEmailId** | [**optional.Interface of string**]()| Optional source email this draft responds to. When supplied, MailSlurp links the draft and later draft actions to that conversation. | 

### Return type

[**EmailDraftDto**](EmailDraftDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentCreateInbox

> InboxDto AgentCreateInbox(ctx, createInboxDto)

Create an inbox with an account-wide inbox-manager connection

Creates an inbox without changing this agent connection's scopes. Account-wide scope is required so the new inbox is already covered.

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


## AgentCreateReplyDraft

> EmailDraftDto AgentCreateReplyDraft(ctx, emailId, replyToEmailOptions)

Create a correctly threaded reply draft

Builds the same recipients and RFC In-Reply-To/References headers as immediate reply, saves them as a draft, and preserves the required threading headers across later agent draft updates.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()|  | 
**replyToEmailOptions** | [**ReplyToEmailOptions**](ReplyToEmailOptions)|  | 

### Return type

[**EmailDraftDto**](EmailDraftDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentCreateTotpDeviceForBase32SecretKey

> TotpDeviceDto AgentCreateTotpDeviceForBase32SecretKey(ctx, createTotpDeviceBase32SecretKeyOptions)

Create an agent-managed TOTP device from a base32 secret key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createTotpDeviceBase32SecretKeyOptions** | [**CreateTotpDeviceBase32SecretKeyOptions**](CreateTotpDeviceBase32SecretKeyOptions)|  | 

### Return type

[**TotpDeviceDto**](TotpDeviceDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentCreateTotpDeviceForCustom

> TotpDeviceDto AgentCreateTotpDeviceForCustom(ctx, createTotpDeviceCustomOptions)

Create an agent-managed TOTP device from custom options

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createTotpDeviceCustomOptions** | [**CreateTotpDeviceCustomOptions**](CreateTotpDeviceCustomOptions)|  | 

### Return type

[**TotpDeviceDto**](TotpDeviceDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentCreateTotpDeviceForOtpAuthUrl

> TotpDeviceDto AgentCreateTotpDeviceForOtpAuthUrl(ctx, createTotpDeviceOtpAuthUrlOptions)

Create an agent-managed TOTP device from an OTP Auth URL

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createTotpDeviceOtpAuthUrlOptions** | [**CreateTotpDeviceOtpAuthUrlOptions**](CreateTotpDeviceOtpAuthUrlOptions)|  | 

### Return type

[**TotpDeviceDto**](TotpDeviceDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentDeferWorkItem

> AgentDeferredWorkItemDto AgentDeferWorkItem(ctx, conversationId, deferAgentWorkItemOptions)

Defer a work item to human review

Creates or returns the active review handoff and releases the supplied claim after the handoff is durable. Deferral is a normal workflow result rather than an operational failure.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | [**string**]()|  | 
**deferAgentWorkItemOptions** | [**DeferAgentWorkItemOptions**](DeferAgentWorkItemOptions)|  | 

### Return type

[**AgentDeferredWorkItemDto**](AgentDeferredWorkItemDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentDeleteEmailDraft

> AgentDeleteEmailDraft(ctx, inboxId, draftId)

Delete a scoped email draft for an agent

Deletes a saved draft without sending it.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()| ID of inbox the draft belongs to | 
**draftId** | [**string**]()| ID of draft | 

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


## AgentDeleteTotpDevice

> AgentDeleteTotpDevice(ctx, id)

Delete an agent-managed TOTP device

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()|  | 

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


## AgentExtractEmailAttachmentText

> ExtractAttachmentTextResult AgentExtractEmailAttachmentText(ctx, emailId, attachmentId, optional)

Extract text from a scoped email attachment for an agent

Extracts bounded text only after verifying both the email scope and that the attachment belongs to that email. List attachments first instead of guessing IDs.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()|  | 
**attachmentId** | **string**|  | 
 **optional** | ***AgentExtractEmailAttachmentTextOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AgentExtractEmailAttachmentTextOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **extractAttachmentTextOptions** | [**optional.Interface of ExtractAttachmentTextOptions**](ExtractAttachmentTextOptions)|  | 

### Return type

[**ExtractAttachmentTextResult**](ExtractAttachmentTextResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentGenerateTotpDeviceCode

> TotpDeviceCodeDto AgentGenerateTotpDeviceCode(ctx, id, optional)

Generate a code for an agent-managed TOTP device

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()|  | 
 **optional** | ***AgentGenerateTotpDeviceCodeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AgentGenerateTotpDeviceCodeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **at** | **optional.Time**|  | 
 **minSecondsUntilExpire** | **optional.Int32**| Minimum remaining code lifetime. The request may wait for the next period when needed. | [default to 5]

### Return type

[**TotpDeviceCodeDto**](TotpDeviceCodeDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentGetCapabilities

> AgentCapabilitiesDto AgentGetCapabilities(ctx, )

Discover this agent connection's capabilities

Returns the connection identity, role, effective permission names, declared resource scope, feature availability, and server-enforced limits. Call this once during worker startup instead of hard-coding role behavior.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**AgentCapabilitiesDto**](AgentCapabilitiesDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentGetConversationV2

> AgentConversationDetailV2Dto AgentGetConversationV2(ctx, conversationId)

Get one scoped conversation

Returns the projected conversation plus this connection's bounded activity timeline, internal notes, and active human-review handoff.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | [**string**]()|  | 

### Return type

[**AgentConversationDetailV2Dto**](AgentConversationDetailV2Dto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentGetEmail

> Email AgentGetEmail(ctx, emailId)

Get scoped hydrated email for an agent

Returns parsed email content after resolving the message inbox and checking agent scope.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()| ID of email to inspect | 

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


## AgentGetEmailDraft

> EmailDraftDto AgentGetEmailDraft(ctx, inboxId, draftId)

Get a scoped email draft for an agent

Returns a saved draft including the full SendEmailOptions payload.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()| ID of inbox the draft belongs to | 
**draftId** | [**string**]()| ID of draft | 

### Return type

[**EmailDraftDto**](EmailDraftDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentGetEmailSummary

> EmailPreview AgentGetEmailSummary(ctx, emailId)

Get scoped email summary for an agent

Returns lightweight metadata and headers for an email after resolving its inbox and checking the agent key scope.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()| ID of email to inspect | 

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


## AgentGetEmailThread

> ScopedEmailThread AgentGetEmailThread(ctx, emailId, optional)

Get a scoped email thread for an agent

Returns thread metadata and ordered messages using an email ID as the authorization anchor. Arbitrary thread-ID probing is not allowed.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()|  | 
 **optional** | ***AgentGetEmailThreadOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AgentGetEmailThreadOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | **optional.String**|  | [default to ASC]

### Return type

[**ScopedEmailThread**](ScopedEmailThread)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentGetLatestInboxEmail

> Email AgentGetLatestInboxEmail(ctx, inboxId, optional)

Get latest scoped inbox email for an agent

Waits for and returns the latest email in one scoped inbox. This is the agent-shaped equivalent of latest-message polling.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()| ID of inbox the agent may read | 
 **optional** | ***AgentGetLatestInboxEmailOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AgentGetLatestInboxEmailOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timeoutMillis** | **optional.Int64**| Timeout milliseconds to wait for latest email | [default to 30000]
 **unreadOnly** | **optional.Bool**| Only return unread emails | [default to false]
 **since** | **optional.Time**| Only return email received at or after this time | 
 **before** | **optional.Time**| Only return email received before this time | 

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


## AgentGetNextWorkItem

> AgentWorkItemDto AgentGetNextWorkItem(ctx, getNextAgentWorkItemOptions)

Claim the next conversation work item

Selects the oldest available scoped conversation and atomically claims it for one worker. This combines queue selection and claim acquisition so concurrent workers do not need to coordinate those steps themselves.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**getNextAgentWorkItemOptions** | [**GetNextAgentWorkItemOptions**](GetNextAgentWorkItemOptions)|  | 

### Return type

[**AgentWorkItemDto**](AgentWorkItemDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentGetPhoneMessageThreadItems

> PagePhoneMessageThreadItemProjection AgentGetPhoneMessageThreadItems(ctx, phoneNumberId, otherPhoneNumber, optional)

List SMS thread messages

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneNumberId** | [**string**]()|  | 
**otherPhoneNumber** | **string**| Other participant phone number in E.164 format | 
 **optional** | ***AgentGetPhoneMessageThreadItemsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AgentGetPhoneMessageThreadItemsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**|  | [default to 0]
 **size** | **optional.Int32**|  | [default to 20]

### Return type

[**PagePhoneMessageThreadItemProjection**](PagePhoneMessageThreadItemProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentGetPhoneMessageThreads

> PagePhoneMessageThreadProjection AgentGetPhoneMessageThreads(ctx, phoneNumberId, optional)

List SMS threads

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneNumberId** | [**string**]()|  | 
 **optional** | ***AgentGetPhoneMessageThreadsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AgentGetPhoneMessageThreadsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**|  | [default to 0]
 **size** | **optional.Int32**|  | [default to 20]

### Return type

[**PagePhoneMessageThreadProjection**](PagePhoneMessageThreadProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentGetPhoneNumber

> PhoneNumberDto AgentGetPhoneNumber(ctx, phoneNumberId)

Get an assigned phone number

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneNumberId** | [**string**]()|  | 

### Return type

[**PhoneNumberDto**](PhoneNumberDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentGetPhoneNumbers

> PageAgentPhoneNumber AgentGetPhoneNumbers(ctx, optional)

List assigned phone numbers

Returns only active phone numbers explicitly assigned to this agent connection.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AgentGetPhoneNumbersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AgentGetPhoneNumbersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**|  | [default to 0]
 **size** | **optional.Int32**|  | [default to 20]
 **sort** | **optional.String**|  | [default to DESC]
 **search** | **optional.String**|  | 

### Return type

[**PageAgentPhoneNumber**](PageAgentPhoneNumber)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentGetReceivedSmsMessage

> SmsDto AgentGetReceivedSmsMessage(ctx, phoneNumberId, smsId)

Get received SMS

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneNumberId** | [**string**]()|  | 
**smsId** | [**string**]()|  | 

### Return type

[**SmsDto**](SmsDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentGetReceivedSmsMessages

> PageSmsProjection AgentGetReceivedSmsMessages(ctx, phoneNumberId, optional)

List received SMS for an assigned phone number

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneNumberId** | [**string**]()|  | 
 **optional** | ***AgentGetReceivedSmsMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AgentGetReceivedSmsMessagesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**|  | [default to 0]
 **size** | **optional.Int32**|  | [default to 20]
 **sort** | **optional.String**|  | [default to DESC]
 **since** | **optional.Time**|  | 
 **before** | **optional.Time**|  | 
 **search** | **optional.String**|  | 
 **favourite** | **optional.Bool**|  | 

### Return type

[**PageSmsProjection**](PageSmsProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentGetSentSmsMessage

> SentSmsDto AgentGetSentSmsMessage(ctx, phoneNumberId, sentSmsId)

Get sent SMS

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneNumberId** | [**string**]()|  | 
**sentSmsId** | [**string**]()|  | 

### Return type

[**SentSmsDto**](SentSmsDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentGetSentSmsMessages

> PageSentSmsProjection AgentGetSentSmsMessages(ctx, phoneNumberId, optional)

List sent SMS for an assigned phone number

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneNumberId** | [**string**]()|  | 
 **optional** | ***AgentGetSentSmsMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AgentGetSentSmsMessagesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**|  | [default to 0]
 **size** | **optional.Int32**|  | [default to 20]
 **sort** | **optional.String**|  | [default to DESC]
 **since** | **optional.Time**|  | 
 **before** | **optional.Time**|  | 
 **search** | **optional.String**|  | 

### Return type

[**PageSentSmsProjection**](PageSentSmsProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentGetSmsConversation

> AgentSmsConversationDetailDto AgentGetSmsConversation(ctx, conversationId)

Get one SMS conversation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | [**string**]()|  | 

### Return type

[**AgentSmsConversationDetailDto**](AgentSmsConversationDetailDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentGetTotpDevice

> TotpDeviceDto AgentGetTotpDevice(ctx, id)

Get an agent-accessible TOTP device

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()|  | 

### Return type

[**TotpDeviceDto**](TotpDeviceDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentGetTotpDeviceBy

> TotpDeviceOptionalDto AgentGetTotpDeviceBy(ctx, optional)

Find an agent-accessible TOTP device by name, issuer, or username

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AgentGetTotpDeviceByOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AgentGetTotpDeviceByOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**|  | 
 **issuer** | **optional.String**|  | 
 **username** | **optional.String**|  | 

### Return type

[**TotpDeviceOptionalDto**](TotpDeviceOptionalDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentGetTotpDevices

> PageTotpDeviceProjection AgentGetTotpDevices(ctx, optional)

List agent-accessible TOTP devices

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AgentGetTotpDevicesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AgentGetTotpDevicesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**|  | [default to 0]
 **size** | **optional.Int32**|  | [default to 20]
 **sort** | **optional.String**|  | [default to DESC]
 **issuer** | **optional.String**|  | 
 **username** | **optional.String**|  | 

### Return type

[**PageTotpDeviceProjection**](PageTotpDeviceProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentListAccessibleInboxes

> AccessibleInboxList AgentListAccessibleInboxes(ctx, optional)

List inboxes accessible to an agent

Returns the inbox IDs, addresses, names, and tags visible to this resource-scoped agent key. Agents should use this endpoint instead of guessing inbox IDs.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AgentListAccessibleInboxesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AgentListAccessibleInboxesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **optional.Int32**|  | [default to 20]
 **sort** | **optional.String**|  | [default to ASC]

### Return type

[**AccessibleInboxList**](AccessibleInboxList)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentListAutomationEvents

> AgentAutomationEventPageDto AgentListAutomationEvents(ctx, optional)

Read durable scoped automation events

Returns an ordered at-least-once event batch. Persist nextCursor after processing and deduplicate by eventId. waitSeconds enables bounded long polling; production workers should not continuously call latest-email polling.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AgentListAutomationEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AgentListAutomationEventsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**|  | 
 **eventType** | [**optional.Interface of []string**](string)|  | 
 **size** | **optional.Int32**|  | [default to 20]
 **waitSeconds** | **optional.Int32**|  | [default to 0]

### Return type

[**AgentAutomationEventPageDto**](AgentAutomationEventPageDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentListConversationsV2

> AgentConversationCursorPageDto AgentListConversationsV2(ctx, optional)

List active conversations with an opaque keyset cursor

Reads the bounded mail-thread projection and applies current connection scope on every request. Persist nextCursor only after processing the returned page. Exact totals are intentionally omitted.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AgentListConversationsV2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AgentListConversationsV2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowState** | **optional.String**|  | 
 **outcome** | **optional.String**|  | 
 **reviewStatus** | **optional.String**|  | 
 **inboxId** | [**optional.Interface of string**]()|  | 
 **unread** | **optional.Bool**|  | 
 **hasAttachments** | **optional.Bool**|  | 
 **since** | **optional.Time**|  | 
 **before** | **optional.Time**|  | 
 **search** | **optional.String**|  | 
 **sort** | **optional.String**|  | [default to LAST_MESSAGE_DESC]
 **cursor** | **optional.String**|  | 
 **size** | **optional.Int32**|  | [default to 20]

### Return type

[**AgentConversationCursorPageDto**](AgentConversationCursorPageDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentListEmailAttachments

> []AttachmentMetaData AgentListEmailAttachments(ctx, emailId)

List attachments on a scoped email for an agent

Returns bounded attachment metadata only after verifying the email belongs to an inbox visible to the agent key.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()|  | 

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


## AgentListEmailDrafts

> PageEmailDraftProjection AgentListEmailDrafts(ctx, inboxId, optional)

List scoped email drafts for an agent

Lists draft previews for one scoped inbox.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()| ID of inbox to list drafts for | 
 **optional** | ***AgentListEmailDraftsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AgentListEmailDraftsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Optional page index in draft pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in draft pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]

### Return type

[**PageEmailDraftProjection**](PageEmailDraftProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentListInboxEmails

> PageEmailPreview AgentListInboxEmails(ctx, inboxId, optional)

List scoped inbox emails for an agent

Returns paginated email previews for one inbox after checking the agent key has AGENT_READ_EMAIL and is scoped to the inbox.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()| ID of inbox the agent may read | 
 **optional** | ***AgentListInboxEmailsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AgentListInboxEmailsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Optional page index in inbox emails list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in inbox emails list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **since** | **optional.Time**| Optional filter by received after given date time | 
 **before** | **optional.Time**| Optional filter by received before given date time | 
 **syncConnectors** | **optional.Bool**| Sync connectors before fetching emails | 

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


## AgentListInboxThreads

> PageEmailThreadProjection AgentListInboxThreads(ctx, inboxId, optional)

List scoped inbox threads for an agent

Lists conversation threads for one scoped inbox, including unread state, attachment presence, message count, and latest-message excerpts.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()|  | 
 **optional** | ***AgentListInboxThreadsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AgentListInboxThreadsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**|  | [default to 0]
 **size** | **optional.Int32**|  | [default to 20]
 **sort** | **optional.String**|  | [default to DESC]
 **search** | **optional.String**|  | 
 **since** | **optional.Time**|  | 
 **before** | **optional.Time**|  | 

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


## AgentListSmsConversations

> AgentSmsConversationCursorPageDto AgentListSmsConversations(ctx, optional)

List active SMS conversations

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AgentListSmsConversationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AgentListSmsConversationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowState** | **optional.String**|  | 
 **outcome** | **optional.String**|  | 
 **reviewStatus** | **optional.String**|  | 
 **phoneNumberId** | [**optional.Interface of string**]()|  | 
 **unread** | **optional.Bool**|  | 
 **hasMedia** | **optional.Bool**|  | 
 **since** | **optional.Time**|  | 
 **before** | **optional.Time**|  | 
 **search** | **optional.String**|  | 
 **sort** | **optional.String**|  | [default to LAST_MESSAGE_DESC]
 **cursor** | **optional.String**|  | 
 **size** | **optional.Int32**|  | [default to 20]

### Return type

[**AgentSmsConversationCursorPageDto**](AgentSmsConversationCursorPageDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentPrepareWorkItemReply

> AgentPreparedReplyDto AgentPrepareWorkItemReply(ctx, conversationId, prepareAgentReplyOptions)

Prepare a guarded reply draft

Checks that the selected conversation and latest inbound email are still current, then creates a correctly threaded reply draft. An optional internal note records context for human review without sending email.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | [**string**]()|  | 
**prepareAgentReplyOptions** | [**PrepareAgentReplyOptions**](PrepareAgentReplyOptions)|  | 

### Return type

[**AgentPreparedReplyDto**](AgentPreparedReplyDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentReleaseConversationClaim

> AgentReleaseConversationClaim(ctx, conversationId, claimId)

Release a conversation claim

Release is idempotent; repeating it does not fail.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | [**string**]()|  | 
**claimId** | [**string**]()|  | 

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


## AgentReleaseSmsClaim

> AgentReleaseSmsClaim(ctx, conversationId, claimId)

Release an SMS conversation claim

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | [**string**]()|  | 
**claimId** | [**string**]()|  | 

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


## AgentRenewConversationClaim

> AgentConversationClaimDto AgentRenewConversationClaim(ctx, conversationId, claimId, renewAgentConversationClaimOptions)

Renew a conversation claim

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | [**string**]()|  | 
**claimId** | [**string**]()|  | 
**renewAgentConversationClaimOptions** | [**RenewAgentConversationClaimOptions**](RenewAgentConversationClaimOptions)|  | 

### Return type

[**AgentConversationClaimDto**](AgentConversationClaimDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentRenewSmsClaim

> AgentConversationClaimDto AgentRenewSmsClaim(ctx, conversationId, claimId, renewAgentConversationClaimOptions)

Renew an SMS conversation claim

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | [**string**]()|  | 
**claimId** | [**string**]()|  | 
**renewAgentConversationClaimOptions** | [**RenewAgentConversationClaimOptions**](RenewAgentConversationClaimOptions)|  | 

### Return type

[**AgentConversationClaimDto**](AgentConversationClaimDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentReplyToEmail

> SentEmailDto AgentReplyToEmail(ctx, emailId, replyToEmailOptions, optional)

Reply to a scoped email as an agent

Sends a reply using the original conversation context after checking AGENT_REPLY_EMAIL and inbox scope.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()| ID of email to reply to | 
**replyToEmailOptions** | [**ReplyToEmailOptions**](ReplyToEmailOptions)|  | 
 **optional** | ***AgentReplyToEmailOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AgentReplyToEmailOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **idempotencyKey** | **optional.String**|  | 
 **expectedConversationVersion** | **optional.Int64**|  | 
 **expectedLatestEmailId** | [**optional.Interface of string**]()|  | 

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


## AgentReplyToSms

> SentSmsDto AgentReplyToSms(ctx, phoneNumberId, smsId, smsReplyOptions, optional)

Reply to received SMS

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneNumberId** | [**string**]()|  | 
**smsId** | [**string**]()|  | 
**smsReplyOptions** | [**SmsReplyOptions**](SmsReplyOptions)|  | 
 **optional** | ***AgentReplyToSmsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AgentReplyToSmsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **idempotencyKey** | **optional.String**|  | 
 **expectedConversationVersion** | **optional.Int64**|  | 
 **expectedLatestMessageId** | [**optional.Interface of string**]()|  | 

### Return type

[**SentSmsDto**](SentSmsDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentReportWorkerHeartbeat

> AgentWorkerHeartbeatDto AgentReportWorkerHeartbeat(ctx, agentWorkerHeartbeatOptions)

Report a worker heartbeat

Records one running worker instance and its deployed release for owner-side health monitoring. Send this during startup and periodically while the process can accept work; it does not claim or mutate a conversation.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentWorkerHeartbeatOptions** | [**AgentWorkerHeartbeatOptions**](AgentWorkerHeartbeatOptions)|  | 

### Return type

[**AgentWorkerHeartbeatDto**](AgentWorkerHeartbeatDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentRequestConversationHandoff

> AgentConversationHandoffDto AgentRequestConversationHandoff(ctx, conversationId, createAgentConversationHandoffOptions)

Request human review for a conversation

Creates or returns the active handoff and moves the shared workflow state into REVIEW without classifying the request as an operational failure.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | [**string**]()|  | 
**createAgentConversationHandoffOptions** | [**CreateAgentConversationHandoffOptions**](CreateAgentConversationHandoffOptions)|  | 

### Return type

[**AgentConversationHandoffDto**](AgentConversationHandoffDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentRequestSmsHandoff

> AgentConversationHandoffDto AgentRequestSmsHandoff(ctx, conversationId, createAgentSmsConversationHandoffOptions)

Request human review for an SMS conversation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | [**string**]()|  | 
**createAgentSmsConversationHandoffOptions** | [**CreateAgentSmsConversationHandoffOptions**](CreateAgentSmsConversationHandoffOptions)|  | 

### Return type

[**AgentConversationHandoffDto**](AgentConversationHandoffDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentSearchInboxEmails

> []EmailPreview AgentSearchInboxEmails(ctx, inboxId, matchOptions, optional)

Search scoped inbox emails for an agent

Searches one scoped inbox using explicit field matches and attachment conditions. Use this for deterministic automation recovery instead of repeatedly fetching the latest message.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()|  | 
**matchOptions** | [**MatchOptions**](MatchOptions)|  | 
 **optional** | ***AgentSearchInboxEmailsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AgentSearchInboxEmailsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **size** | **optional.Int32**|  | [default to 20]
 **sort** | **optional.String**|  | [default to DESC]
 **since** | **optional.Time**|  | 
 **before** | **optional.Time**|  | 
 **unreadOnly** | **optional.Bool**|  | 

### Return type

[**[]EmailPreview**](EmailPreview)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentSendEmailDraft

> SentEmailDto AgentSendEmailDraft(ctx, inboxId, draftId, optional)

Send a scoped email draft for an agent

Sends a saved draft through the existing inbox send pipeline and deletes it only after send success.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()| ID of inbox the draft belongs to | 
**draftId** | [**string**]()| ID of draft | 
 **optional** | ***AgentSendEmailDraftOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AgentSendEmailDraftOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **idempotencyKey** | **optional.String**|  | 

### Return type

[**SentEmailDto**](SentEmailDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentSendSms

> SentSmsDto AgentSendSms(ctx, phoneNumberId, smsSendOptions, optional)

Send SMS from an assigned phone number

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneNumberId** | [**string**]()|  | 
**smsSendOptions** | [**SmsSendOptions**](SmsSendOptions)|  | 
 **optional** | ***AgentSendSmsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AgentSendSmsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **idempotencyKey** | **optional.String**|  | 
 **expectedConversationVersion** | **optional.Int64**|  | 
 **expectedLatestMessageId** | [**optional.Interface of string**]()|  | 

### Return type

[**SentSmsDto**](SentSmsDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentUpdateEmailDraft

> EmailDraftDto AgentUpdateEmailDraft(ctx, inboxId, draftId, sendEmailOptions)

Update a scoped email draft for an agent

Replaces the SendEmailOptions stored for a draft. The draft remains attached to its original inbox.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()| ID of inbox the draft belongs to | 
**draftId** | [**string**]()| ID of draft | 
**sendEmailOptions** | [**SendEmailOptions**](SendEmailOptions)|  | 

### Return type

[**EmailDraftDto**](EmailDraftDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AgentUpdateInbox

> InboxDto AgentUpdateInbox(ctx, inboxId, updateInboxOptions)

Update an inbox visible to an inbox-manager connection

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


## AgentUpdateTotpDevice

> TotpDeviceDto AgentUpdateTotpDevice(ctx, id, updateTotpDeviceOptions)

Update an agent-managed TOTP device

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()|  | 
**updateTotpDeviceOptions** | [**UpdateTotpDeviceOptions**](UpdateTotpDeviceOptions)|  | 

### Return type

[**TotpDeviceDto**](TotpDeviceDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

