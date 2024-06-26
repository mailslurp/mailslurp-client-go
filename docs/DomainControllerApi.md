# MailSlurp\DomainControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDomainWildcardCatchAll**](DomainControllerApi#AddDomainWildcardCatchAll) | **Post** /domains/{id}/wildcard | Add catch all wild card inbox to domain
[**CreateDomain**](DomainControllerApi#CreateDomain) | **Post** /domains | Create Domain
[**DeleteDomain**](DomainControllerApi#DeleteDomain) | **Delete** /domains/{id} | Delete a domain
[**GetAvailableDomains**](DomainControllerApi#GetAvailableDomains) | **Get** /domains/available-domains | Get all usable domains
[**GetDomain**](DomainControllerApi#GetDomain) | **Get** /domains/{id} | Get a domain
[**GetDomainIssues**](DomainControllerApi#GetDomainIssues) | **Get** /domains/issues | Get domain issues
[**GetDomainWildcardCatchAllInbox**](DomainControllerApi#GetDomainWildcardCatchAllInbox) | **Get** /domains/{id}/wildcard | Get catch all wild card inbox for domain
[**GetDomains**](DomainControllerApi#GetDomains) | **Get** /domains | Get domains
[**GetMailSlurpDomains**](DomainControllerApi#GetMailSlurpDomains) | **Get** /domains/mailslurp-domains | Get MailSlurp domains
[**UpdateDomain**](DomainControllerApi#UpdateDomain) | **Put** /domains/{id} | Update a domain



## AddDomainWildcardCatchAll

> DomainDto AddDomainWildcardCatchAll(ctx, id)

Add catch all wild card inbox to domain

Add a catch all inbox to a domain so that any emails sent to it that cannot be matched will be sent to the catch all inbox generated

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()|  | 

### Return type

[**DomainDto**](DomainDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CreateDomain

> DomainDto CreateDomain(ctx, createDomainOptions)

Create Domain

Link a domain that you own with MailSlurp so you can create email addresses using it. Endpoint returns DNS records used for validation. You must add these verification records to your host provider's DNS setup to verify the domain.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createDomainOptions** | [**CreateDomainOptions**](CreateDomainOptions)|  | 

### Return type

[**DomainDto**](DomainDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DeleteDomain

> []string DeleteDomain(ctx, id)

Delete a domain

Delete a domain. This will disable any existing inboxes that use this domain.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()|  | 

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


## GetAvailableDomains

> DomainGroupsDto GetAvailableDomains(ctx, optional)

Get all usable domains

List all domains available for use with email address creation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAvailableDomainsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAvailableDomainsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inboxType** | **optional.String**|  | 

### Return type

[**DomainGroupsDto**](DomainGroupsDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDomain

> DomainDto GetDomain(ctx, id, optional)

Get a domain

Returns domain verification status and tokens for a given domain

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()|  | 
 **optional** | ***GetDomainOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDomainOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **checkForErrors** | **optional.Bool**|  | 

### Return type

[**DomainDto**](DomainDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDomainIssues

> DomainIssuesDto GetDomainIssues(ctx, )

Get domain issues

List domain issues for domains you have created

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**DomainIssuesDto**](DomainIssuesDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDomainWildcardCatchAllInbox

> InboxDto GetDomainWildcardCatchAllInbox(ctx, id)

Get catch all wild card inbox for domain

Get the catch all inbox for a domain for missed emails

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()|  | 

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


## GetDomains

> []DomainPreview GetDomains(ctx, )

Get domains

List all custom domains you have created

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]DomainPreview**](DomainPreview)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetMailSlurpDomains

> DomainGroupsDto GetMailSlurpDomains(ctx, optional)

Get MailSlurp domains

List all MailSlurp domains used with non-custom email addresses

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetMailSlurpDomainsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetMailSlurpDomainsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inboxType** | **optional.String**|  | 

### Return type

[**DomainGroupsDto**](DomainGroupsDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## UpdateDomain

> DomainDto UpdateDomain(ctx, id, updateDomainOptions)

Update a domain

Update values on a domain. Note you cannot change the domain name as it is immutable. Recreate the domain if you need to alter this.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()|  | 
**updateDomainOptions** | [**UpdateDomainOptions**](UpdateDomainOptions)|  | 

### Return type

[**DomainDto**](DomainDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

