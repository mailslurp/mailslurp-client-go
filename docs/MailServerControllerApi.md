# MailSlurp\MailServerControllerApi

All URIs are relative to *https://api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DescribeMailServerDomain**](MailServerControllerApi.md#DescribeMailServerDomain) | **Post** /mail-server/describe/domain | Get DNS Mail Server records for a domain
[**VerifyEmailAddress**](MailServerControllerApi.md#VerifyEmailAddress) | **Post** /mail-server/verify/email-address | Verify the existence of an email address at a given mail server.



## DescribeMailServerDomain

> DescribeMailServerDomainResult DescribeMailServerDomain(ctx, describeOptions)

Get DNS Mail Server records for a domain

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**describeOptions** | [**DescribeDomainOptions**](DescribeDomainOptions.md)| describeOptions | 

### Return type

[**DescribeMailServerDomainResult**](DescribeMailServerDomainResult.md)

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyEmailAddress

> EmailVerificationResult VerifyEmailAddress(ctx, verifyOptions)

Verify the existence of an email address at a given mail server.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**verifyOptions** | [**VerifyEmailAddressOptions**](VerifyEmailAddressOptions.md)| verifyOptions | 

### Return type

[**EmailVerificationResult**](EmailVerificationResult.md)

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

