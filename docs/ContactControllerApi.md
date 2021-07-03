# MailSlurp\ContactControllerApi

All URIs are relative to *https://api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContact**](ContactControllerApi#CreateContact) | **Post** /contacts | Create a contact
[**DeleteContact**](ContactControllerApi#DeleteContact) | **Delete** /contacts/{contactId} | Delete contact
[**GetAllContacts**](ContactControllerApi#GetAllContacts) | **Get** /contacts/paginated | Get all contacts
[**GetContact**](ContactControllerApi#GetContact) | **Get** /contacts/{contactId} | Get contact
[**GetContactVCard**](ContactControllerApi#GetContactVCard) | **Get** /contacts/{contactId}/download | Get contact vCard vcf file
[**GetContacts**](ContactControllerApi#GetContacts) | **Get** /contacts | Get all contacts



## CreateContact

> ContactDto CreateContact(ctx, createContactOptions)

Create a contact

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createContactOptions** | [**CreateContactOptions**](CreateContactOptions)| createContactOptions | 

### Return type

[**ContactDto**](ContactDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DeleteContact

> DeleteContact(ctx, contactId)

Delete contact

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | [**string**]()| contactId | 

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


## GetAllContacts

> PageContactProjection GetAllContacts(ctx, optional)

Get all contacts

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllContactsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllContactsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Optional page index in list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size in list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]

### Return type

[**PageContactProjection**](PageContactProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetContact

> ContactDto GetContact(ctx, contactId)

Get contact

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | [**string**]()| contactId | 

### Return type

[**ContactDto**](ContactDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetContactVCard

> string GetContactVCard(ctx, contactId)

Get contact vCard vcf file

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | [**string**]()| contactId | 

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


## GetContacts

> []ContactProjection GetContacts(ctx, )

Get all contacts

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]ContactProjection**](ContactProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

