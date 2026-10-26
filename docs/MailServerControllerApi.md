# MailSlurp\MailServerControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DescribeMailServerDomain**](MailServerControllerApi#DescribeMailServerDomain) | **Post** /mail-server/describe/domain | Get DNS Mail Server records for a domain
[**GetDnsLookup**](MailServerControllerApi#GetDnsLookup) | **Post** /mail-server/describe/dns-lookup | Lookup DNS records for a domain
[**GetDnsLookups**](MailServerControllerApi#GetDnsLookups) | **Post** /mail-server/describe/dns-lookups | Lookup DNS records for multiple domains
[**GetIpAddress**](MailServerControllerApi#GetIpAddress) | **Post** /mail-server/describe/ip-address | Get IP address for a domain
[**VerifyEmailAddress**](MailServerControllerApi#VerifyEmailAddress) | **Post** /mail-server/verify/email-address | Deprecated. Use the EmailVerificationController methods for more accurate and reliable functionality. Verify the existence of an email address at a given mail server.



## DescribeMailServerDomain

> DescribeMailServerDomainResult DescribeMailServerDomain(ctx, describeDomainOptions)

Get DNS Mail Server records for a domain

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**describeDomainOptions** | [**DescribeDomainOptions**](DescribeDomainOptions)|  | 

### Return type

[**DescribeMailServerDomainResult**](DescribeMailServerDomainResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDnsLookup

> DnsLookupResults GetDnsLookup(ctx, dnsLookupOptions)

Lookup DNS records for a domain

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dnsLookupOptions** | [**DnsLookupOptions**](DnsLookupOptions)|  | 

### Return type

[**DnsLookupResults**](DNSLookupResults)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDnsLookups

> DnsLookupResults GetDnsLookups(ctx, dnsLookupsOptions)

Lookup DNS records for multiple domains

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dnsLookupsOptions** | [**DnsLookupsOptions**](DnsLookupsOptions)|  | 

### Return type

[**DnsLookupResults**](DNSLookupResults)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetIpAddress

> IpAddressResult GetIpAddress(ctx, name)

Get IP address for a domain

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**|  | 

### Return type

[**IpAddressResult**](IPAddressResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## VerifyEmailAddress

> EmailVerificationResult VerifyEmailAddress(ctx, verifyEmailAddressOptions)

Deprecated. Use the EmailVerificationController methods for more accurate and reliable functionality. Verify the existence of an email address at a given mail server.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**verifyEmailAddressOptions** | [**VerifyEmailAddressOptions**](VerifyEmailAddressOptions)|  | 

### Return type

[**EmailVerificationResult**](EmailVerificationResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

