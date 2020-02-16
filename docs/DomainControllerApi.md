# \DomainControllerApi

All URIs are relative to *https://api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDomain**](DomainControllerApi.md#CreateDomain) | **Post** /domains | Create Domain
[**DeleteDomain**](DomainControllerApi.md#DeleteDomain) | **Delete** /domains/{id} | Delete a domain
[**GetDomain**](DomainControllerApi.md#GetDomain) | **Get** /domains/{id} | Get a domain
[**GetDomains**](DomainControllerApi.md#GetDomains) | **Get** /domains | Get domains



## CreateDomain

> DomainPlusVerificationRecordsAndStatus CreateDomain(ctx, domainOptions)

Create Domain

Link a domain that you own with MailSlurp so you can create email addresses using it. Endpoint returns DNS records used for validation. You must add these verification records to your host provider's DNS setup to verify the domain.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainOptions** | [**CreateDomainOptions**](CreateDomainOptions.md)| domainOptions | 

### Return type

[**DomainPlusVerificationRecordsAndStatus**](Domain plus verification records and status.md)

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDomain

> DeleteDomain(ctx, id)

Delete a domain

Delete a domain. This will disable any existing inboxes that use this domain.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| id | 

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


## GetDomain

> DomainPlusVerificationRecordsAndStatus GetDomain(ctx, id)

Get a domain

Returns domain verification status and tokens for a given domain

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| id | 

### Return type

[**DomainPlusVerificationRecordsAndStatus**](Domain plus verification records and status.md)

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDomains

> []DomainPreview GetDomains(ctx, )

Get domains

List all custom domains you have created

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]DomainPreview**](DomainPreview.md)

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

