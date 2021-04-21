# MailSlurp\BulkActionsControllerApi

All URIs are relative to *https://api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkCreateInboxes**](BulkActionsControllerApi#BulkCreateInboxes) | **Post** /bulk/inboxes | Bulk create Inboxes (email addresses)
[**BulkDeleteInboxes**](BulkActionsControllerApi#BulkDeleteInboxes) | **Delete** /bulk/inboxes | Bulk Delete Inboxes
[**BulkSendEmails**](BulkActionsControllerApi#BulkSendEmails) | **Post** /bulk/send | Bulk Send Emails



## BulkCreateInboxes

> []Inbox BulkCreateInboxes(ctx, count)

Bulk create Inboxes (email addresses)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**count** | **int32**| Number of inboxes to be created in bulk | 

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


## BulkDeleteInboxes

> BulkDeleteInboxes(ctx, ids)

Bulk Delete Inboxes

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ids** | [**[]string**](string)| ids | 

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


## BulkSendEmails

> BulkSendEmails(ctx, bulkSendEmailOptions)

Bulk Send Emails

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bulkSendEmailOptions** | [**BulkSendEmailOptions**](BulkSendEmailOptions)| bulkSendEmailOptions | 

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

