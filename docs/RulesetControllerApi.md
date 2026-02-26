# MailSlurp\RulesetControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewRuleset**](RulesetControllerApi#CreateNewRuleset) | **Post** /rulesets | Create a ruleset
[**DeleteRuleset**](RulesetControllerApi#DeleteRuleset) | **Delete** /rulesets/{id} | Delete a ruleset
[**DeleteRulesets**](RulesetControllerApi#DeleteRulesets) | **Delete** /rulesets | Delete rulesets
[**GetRuleset**](RulesetControllerApi#GetRuleset) | **Get** /rulesets/{id} | Get a ruleset
[**GetRulesets**](RulesetControllerApi#GetRulesets) | **Get** /rulesets | List rulesets block and allow lists
[**TestInboxRulesetsForInbox**](RulesetControllerApi#TestInboxRulesetsForInbox) | **Put** /rulesets | Test inbox rulesets for inbox
[**TestNewRuleset**](RulesetControllerApi#TestNewRuleset) | **Patch** /rulesets | Test new ruleset
[**TestRuleset**](RulesetControllerApi#TestRuleset) | **Post** /rulesets/{id}/test | Test a ruleset
[**TestRulesetReceiving**](RulesetControllerApi#TestRulesetReceiving) | **Post** /rulesets/test-receiving | Test receiving with rulesets
[**TestRulesetSending**](RulesetControllerApi#TestRulesetSending) | **Post** /rulesets/test-sending | Test sending with rulesets



## CreateNewRuleset

> RulesetDto CreateNewRuleset(ctx, createRulesetOptions, optional)

Create a ruleset

Create a new inbox or phone number rule for forwarding, blocking, and allowing emails or SMS when sending and receiving

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createRulesetOptions** | [**CreateRulesetOptions**](CreateRulesetOptions)|  | 
 **optional** | ***CreateNewRulesetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateNewRulesetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inboxId** | [**optional.Interface of string**]()| Inbox id to attach ruleset to | 
 **phoneId** | [**optional.Interface of string**]()| Phone id to attach ruleset to | 

### Return type

[**RulesetDto**](RulesetDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DeleteRuleset

> DeleteRuleset(ctx, id)

Delete a ruleset

Delete ruleset

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()| ID of ruleset | 

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


## DeleteRulesets

> DeleteRulesets(ctx, optional)

Delete rulesets

Delete rulesets. Accepts optional inboxId or phoneId filters.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteRulesetsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteRulesetsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inboxId** | [**optional.Interface of string**]()| Optional inbox id to attach ruleset to | 
 **phoneId** | [**optional.Interface of string**]()|  | 

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


## GetRuleset

> RulesetDto GetRuleset(ctx, id)

Get a ruleset

Get ruleset

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()| ID of ruleset | 

### Return type

[**RulesetDto**](RulesetDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetRulesets

> PageRulesetDto GetRulesets(ctx, optional)

List rulesets block and allow lists

List all rulesets attached to an inbox or phone or account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetRulesetsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetRulesetsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inboxId** | [**optional.Interface of string**]()| Optional inbox id to get rulesets from | 
 **phoneId** | [**optional.Interface of string**]()| Optional phone id to get rulesets from | 
 **page** | **optional.Int32**| Optional page index in inbox ruleset list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in inbox ruleset list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **searchFilter** | **optional.String**| Optional search filter | 
 **since** | **optional.Time**| Filter by created at after the given timestamp | 
 **before** | **optional.Time**| Filter by created at before the given timestamp | 

### Return type

[**PageRulesetDto**](PageRulesetDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## TestInboxRulesetsForInbox

> InboxRulesetTestResult TestInboxRulesetsForInbox(ctx, inboxId, rulesetTestOptions)

Test inbox rulesets for inbox

Test inbox rulesets for inbox

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()| ID of inbox | 
**rulesetTestOptions** | [**RulesetTestOptions**](RulesetTestOptions)|  | 

### Return type

[**InboxRulesetTestResult**](InboxRulesetTestResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## TestNewRuleset

> InboxRulesetTestResult TestNewRuleset(ctx, testNewInboxRulesetOptions)

Test new ruleset

Test new ruleset

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testNewInboxRulesetOptions** | [**TestNewInboxRulesetOptions**](TestNewInboxRulesetOptions)|  | 

### Return type

[**InboxRulesetTestResult**](InboxRulesetTestResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## TestRuleset

> InboxRulesetTestResult TestRuleset(ctx, id, rulesetTestOptions)

Test a ruleset

Test an inbox or phone ruleset

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()| ID of ruleset | 
**rulesetTestOptions** | [**RulesetTestOptions**](RulesetTestOptions)|  | 

### Return type

[**InboxRulesetTestResult**](InboxRulesetTestResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## TestRulesetReceiving

> TestRulesetReceivingResult TestRulesetReceiving(ctx, testRulesetReceivingOptions)

Test receiving with rulesets

Test whether inbound emails from an email address would be blocked or allowed by inbox rulesets or test if phone number can receive SMS

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testRulesetReceivingOptions** | [**TestRulesetReceivingOptions**](TestRulesetReceivingOptions)|  | 

### Return type

[**TestRulesetReceivingResult**](TestRulesetReceivingResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## TestRulesetSending

> TestRulesetSendingResult TestRulesetSending(ctx, testInboxRulesetSendingOptions)

Test sending with rulesets

Test whether outbound emails to an email address would be blocked or allowed by inbox rulesets or whether a phone number can send SMS

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testInboxRulesetSendingOptions** | [**TestInboxRulesetSendingOptions**](TestInboxRulesetSendingOptions)|  | 

### Return type

[**TestRulesetSendingResult**](TestRulesetSendingResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

