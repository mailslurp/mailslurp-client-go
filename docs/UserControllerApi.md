# MailSlurp\UserControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInboxRetentionPolicy**](UserControllerApi#CreateInboxRetentionPolicy) | **Post** /user/inbox-retention-policies | 
[**CreateInboxRetentionPolicyRun**](UserControllerApi#CreateInboxRetentionPolicyRun) | **Post** /user/inbox-retention-policies/account/run | 
[**CreateOrUpdateInboxRetentionPolicyForAccount**](UserControllerApi#CreateOrUpdateInboxRetentionPolicyForAccount) | **Post** /user/inbox-retention-policies/account | 
[**DeleteInboxRetentionPolicy**](UserControllerApi#DeleteInboxRetentionPolicy) | **Delete** /user/inbox-retention-policies/{policyId} | 
[**DeleteInboxRetentionPolicyForAccount**](UserControllerApi#DeleteInboxRetentionPolicyForAccount) | **Delete** /user/inbox-retention-policies/account | 
[**GetEntityAutomations**](UserControllerApi#GetEntityAutomations) | **Get** /user/automations | 
[**GetEntityEvents**](UserControllerApi#GetEntityEvents) | **Get** /user/events | 
[**GetEntityFavorites**](UserControllerApi#GetEntityFavorites) | **Get** /user/favorites | 
[**GetInboxRetentionPolicy**](UserControllerApi#GetInboxRetentionPolicy) | **Get** /user/inbox-retention-policies/{policyId} | 
[**GetInboxRetentionPolicyForAccount**](UserControllerApi#GetInboxRetentionPolicyForAccount) | **Get** /user/inbox-retention-policies/account | 
[**GetInboxRetentionPolicyRun**](UserControllerApi#GetInboxRetentionPolicyRun) | **Get** /user/inbox-retention-policies/runs/{runId} | 
[**GetInboxRetentionPolicyRuns**](UserControllerApi#GetInboxRetentionPolicyRuns) | **Get** /user/inbox-retention-policies/runs | 
[**GetJsonPropertyAsString**](UserControllerApi#GetJsonPropertyAsString) | **Post** /user/json/pluck | 
[**GetUserInfo**](UserControllerApi#GetUserInfo) | **Get** /user/info | 
[**GetUserStorageInsights**](UserControllerApi#GetUserStorageInsights) | **Get** /user/storage/insights | 
[**ListInboxRetentionPolicies**](UserControllerApi#ListInboxRetentionPolicies) | **Get** /user/inbox-retention-policies | 
[**UpdateInboxRetentionPolicy**](UserControllerApi#UpdateInboxRetentionPolicy) | **Put** /user/inbox-retention-policies/{policyId} | 



## CreateInboxRetentionPolicy

> InboxRetentionPolicyDto CreateInboxRetentionPolicy(ctx, createInboxRetentionPolicyOptions)



Create an inbox retention policy for your account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createInboxRetentionPolicyOptions** | [**CreateInboxRetentionPolicyOptions**](CreateInboxRetentionPolicyOptions)|  | 

### Return type

[**InboxRetentionPolicyDto**](InboxRetentionPolicyDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CreateInboxRetentionPolicyRun

> InboxRetentionPolicyRunDto CreateInboxRetentionPolicyRun(ctx, )



Start an asynchronous inbox retention policy run for your account

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InboxRetentionPolicyRunDto**](InboxRetentionPolicyRunDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CreateOrUpdateInboxRetentionPolicyForAccount

> InboxRetentionPolicyDto CreateOrUpdateInboxRetentionPolicyForAccount(ctx, createInboxRetentionPolicyForAccountOptions)



Create inbox retention policy for your global account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createInboxRetentionPolicyForAccountOptions** | [**CreateInboxRetentionPolicyForAccountOptions**](CreateInboxRetentionPolicyForAccountOptions)|  | 

### Return type

[**InboxRetentionPolicyDto**](InboxRetentionPolicyDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DeleteInboxRetentionPolicy

> EmptyResponseDto DeleteInboxRetentionPolicy(ctx, policyId)



Delete an inbox retention policy for your account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | [**string**]()|  | 

### Return type

[**EmptyResponseDto**](EmptyResponseDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DeleteInboxRetentionPolicyForAccount

> EmptyResponseDto DeleteInboxRetentionPolicyForAccount(ctx, )



Delete inbox retention policy for your global account

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**EmptyResponseDto**](EmptyResponseDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetEntityAutomations

> PageEntityAutomationItems GetEntityAutomations(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetEntityAutomationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEntityAutomationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Optional page index | [default to 0]
 **size** | **optional.Int32**| Optional page size | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to DESC]
 **since** | **optional.Time**| Filter by created at after the given timestamp | 
 **before** | **optional.Time**| Filter by created at before the given timestamp | 
 **inboxId** | [**optional.Interface of string**]()| Optional inbox ID | 
 **phoneId** | [**optional.Interface of string**]()| Optional phone ID | 
 **filter** | **optional.String**| Optional automation type filter | 

### Return type

[**PageEntityAutomationItems**](PageEntityAutomationItems)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetEntityEvents

> PageEntityEventItems GetEntityEvents(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetEntityEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEntityEventsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Optional page index | [default to 0]
 **size** | **optional.Int32**| Optional page size | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to DESC]
 **since** | **optional.Time**| Filter by created at after the given timestamp | 
 **before** | **optional.Time**| Filter by created at before the given timestamp | 
 **inboxId** | [**optional.Interface of string**]()| Optional inbox ID | 
 **emailId** | [**optional.Interface of string**]()| Optional email ID | 
 **phoneId** | [**optional.Interface of string**]()| Optional phone ID | 
 **smsId** | [**optional.Interface of string**]()| Optional SMS ID | 
 **attachmentId** | [**optional.Interface of string**]()| Optional attachment ID | 
 **filter** | **optional.String**| Optional type filter | 

### Return type

[**PageEntityEventItems**](PageEntityEventItems)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetEntityFavorites

> PageEntityFavouriteItems GetEntityFavorites(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetEntityFavoritesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEntityFavoritesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Optional page index | [default to 0]
 **size** | **optional.Int32**| Optional page size | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to DESC]
 **since** | **optional.Time**| Filter by created at after the given timestamp | 
 **before** | **optional.Time**| Filter by created at before the given timestamp | 
 **filter** | **optional.String**| Optional type filter | 

### Return type

[**PageEntityFavouriteItems**](PageEntityFavouriteItems)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetInboxRetentionPolicy

> InboxRetentionPolicyDto GetInboxRetentionPolicy(ctx, policyId)



Get an inbox retention policy for your account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | [**string**]()|  | 

### Return type

[**InboxRetentionPolicyDto**](InboxRetentionPolicyDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetInboxRetentionPolicyForAccount

> InboxRetentionPolicyOptionalDto GetInboxRetentionPolicyForAccount(ctx, )



Get inbox retention policy for your global account

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InboxRetentionPolicyOptionalDto**](InboxRetentionPolicyOptionalDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetInboxRetentionPolicyRun

> InboxRetentionPolicyRunDto GetInboxRetentionPolicyRun(ctx, runId)



Get a previously submitted inbox retention policy run for your account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | [**string**]()|  | 

### Return type

[**InboxRetentionPolicyRunDto**](InboxRetentionPolicyRunDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetInboxRetentionPolicyRuns

> PageInboxRetentionPolicyRunProjection GetInboxRetentionPolicyRuns(ctx, optional)



List recent inbox retention policy runs for your account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetInboxRetentionPolicyRunsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetInboxRetentionPolicyRunsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Optional page index | [default to 0]
 **size** | **optional.Int32**| Optional page size | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to DESC]

### Return type

[**PageInboxRetentionPolicyRunProjection**](PageInboxRetentionPolicyRunProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetJsonPropertyAsString

> string GetJsonPropertyAsString(ctx, property, jsonNode)



Utility function to extract properties from JSON objects in language where this is cumbersome.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**property** | **string**| JSON property name or dot separated path selector such as &#x60;a.b.c&#x60; | 
**jsonNode** | [**JsonNode**](JsonNode)|  | 

### Return type

**string**

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetUserInfo

> UserInfoDto GetUserInfo(ctx, )



Get account information for your user

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**UserInfoDto**](UserInfoDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetUserStorageInsights

> UserStorageInsightsDto GetUserStorageInsights(ctx, optional)



Get retained storage usage insights for your account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetUserStorageInsightsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetUserStorageInsightsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Maximum number of inboxes and domains to return | [default to 10]

### Return type

[**UserStorageInsightsDto**](UserStorageInsightsDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## ListInboxRetentionPolicies

> InboxRetentionPoliciesDto ListInboxRetentionPolicies(ctx, )



List inbox retention policies for your account

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InboxRetentionPoliciesDto**](InboxRetentionPoliciesDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## UpdateInboxRetentionPolicy

> InboxRetentionPolicyDto UpdateInboxRetentionPolicy(ctx, policyId, updateInboxRetentionPolicyOptions)



Update an inbox retention policy for your account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | [**string**]()|  | 
**updateInboxRetentionPolicyOptions** | [**UpdateInboxRetentionPolicyOptions**](UpdateInboxRetentionPolicyOptions)|  | 

### Return type

[**InboxRetentionPolicyDto**](InboxRetentionPolicyDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

