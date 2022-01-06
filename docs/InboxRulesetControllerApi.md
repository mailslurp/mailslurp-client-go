# MailSlurp\InboxRulesetControllerApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewInboxRuleset**](InboxRulesetControllerApi#CreateNewInboxRuleset) | **Post** /rulesets | Create an inbox ruleset
[**DeleteInboxRuleset**](InboxRulesetControllerApi#DeleteInboxRuleset) | **Delete** /rulesets/{id} | Delete an inbox ruleset
[**DeleteInboxRulesets**](InboxRulesetControllerApi#DeleteInboxRulesets) | **Delete** /rulesets | Delete inbox rulesets
[**GetInboxRuleset**](InboxRulesetControllerApi#GetInboxRuleset) | **Get** /rulesets/{id} | Get an inbox ruleset
[**GetInboxRulesets**](InboxRulesetControllerApi#GetInboxRulesets) | **Get** /rulesets | List inbox rulesets
[**TestInboxRuleset**](InboxRulesetControllerApi#TestInboxRuleset) | **Post** /rulesets/{id}/test | Test an inbox ruleset
[**TestInboxRulesetsForInbox**](InboxRulesetControllerApi#TestInboxRulesetsForInbox) | **Put** /rulesets | Test inbox rulesets for inbox
[**TestNewInboxRuleset**](InboxRulesetControllerApi#TestNewInboxRuleset) | **Patch** /rulesets | Test new inbox ruleset



## CreateNewInboxRuleset

> InboxRulesetDto CreateNewInboxRuleset(ctx, inboxId, createInboxRulesetOptions)

Create an inbox ruleset

Create a new inbox rule for forwarding, blocking, and allowing emails when sending and receiving

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()| Inbox id to attach ruleset to | 
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


## DeleteInboxRuleset

> DeleteInboxRuleset(ctx, id)

Delete an inbox ruleset

Delete inbox ruleset

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()| ID of inbox ruleset | 

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


## DeleteInboxRulesets

> DeleteInboxRulesets(ctx, optional)

Delete inbox rulesets

Delete inbox rulesets. Accepts optional inboxId filter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteInboxRulesetsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteInboxRulesetsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inboxId** | [**optional.Interface of string**]()| Optional inbox id to attach ruleset to | 

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


## GetInboxRuleset

> InboxRulesetDto GetInboxRuleset(ctx, id)

Get an inbox ruleset

Get inbox ruleset

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()| ID of inbox ruleset | 

### Return type

[**InboxRulesetDto**](InboxRulesetDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetInboxRulesets

> PageInboxRulesetDto GetInboxRulesets(ctx, optional)

List inbox rulesets

List all rulesets attached to an inbox

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetInboxRulesetsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetInboxRulesetsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inboxId** | [**optional.Interface of string**]()| Optional inbox id to get rulesets from | 
 **page** | **optional.Int32**| Optional page index in inbox ruleset list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in inbox ruleset list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **searchFilter** | **optional.String**| Optional search filter | 
 **since** | **optional.Time**| Filter by created at after the given timestamp | 
 **before** | **optional.Time**| Filter by created at before the given timestamp | 

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


## TestInboxRuleset

> InboxRulesetTestResult TestInboxRuleset(ctx, id, inboxRulesetTestOptions)

Test an inbox ruleset

Test an inbox ruleset

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()| ID of inbox ruleset | 
**inboxRulesetTestOptions** | [**InboxRulesetTestOptions**](InboxRulesetTestOptions)|  | 

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


## TestInboxRulesetsForInbox

> InboxRulesetTestResult TestInboxRulesetsForInbox(ctx, inboxId, inboxRulesetTestOptions)

Test inbox rulesets for inbox

Test inbox rulesets for inbox

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inboxId** | [**string**]()| ID of inbox | 
**inboxRulesetTestOptions** | [**InboxRulesetTestOptions**](InboxRulesetTestOptions)|  | 

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


## TestNewInboxRuleset

> InboxRulesetTestResult TestNewInboxRuleset(ctx, testNewInboxRulesetOptions)

Test new inbox ruleset

Test new inbox ruleset

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

