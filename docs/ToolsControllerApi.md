# MailSlurp\ToolsControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckEmailFeaturesClientSupport**](ToolsControllerApi#CheckEmailFeaturesClientSupport) | **Post** /tools/check-email-features-client-support | Check email client support for email HTML and CSS features
[**CreateNewFakeEmailAddress**](ToolsControllerApi#CreateNewFakeEmailAddress) | **Post** /tools/fake-email | Create a new email address using the fake email domains
[**DeleteNewFakeEmailAddress**](ToolsControllerApi#DeleteNewFakeEmailAddress) | **Delete** /tools/fake-email | Delete a fake email address using the fake email domains
[**GenerateBimiRecord**](ToolsControllerApi#GenerateBimiRecord) | **Post** /tools/generate-bimi-record | Create a BIMI record policy
[**GenerateDmarcRecord**](ToolsControllerApi#GenerateDmarcRecord) | **Post** /tools/generate-dmarc-record | Create a DMARC record policy
[**GenerateMtaStsRecord**](ToolsControllerApi#GenerateMtaStsRecord) | **Post** /tools/generate-mta-sts-record | Create a TLS reporting record policy
[**GenerateTlsReportingRecord**](ToolsControllerApi#GenerateTlsReportingRecord) | **Post** /tools/generate-tls-reporting-record | Create a TLS reporting record policy
[**GetFakeEmailByEmailAddress**](ToolsControllerApi#GetFakeEmailByEmailAddress) | **Get** /tools/fake-email/byEmailAddress | 
[**GetFakeEmailById**](ToolsControllerApi#GetFakeEmailById) | **Get** /tools/fake-email | Get a fake email by its ID
[**GetFakeEmailRaw**](ToolsControllerApi#GetFakeEmailRaw) | **Get** /tools/fake-email/html | Get raw fake email content
[**GetFakeEmailsForAddress**](ToolsControllerApi#GetFakeEmailsForAddress) | **Get** /tools/fake-emails | Get fake emails for an address
[**LookupBimiDomain**](ToolsControllerApi#LookupBimiDomain) | **Post** /tools/lookup-bimi-domain | Lookup a BIMI record policy
[**LookupDmarcDomain**](ToolsControllerApi#LookupDmarcDomain) | **Post** /tools/lookup-dmarc-domain | Lookup a DMARC record policy
[**LookupMtaStsDomain**](ToolsControllerApi#LookupMtaStsDomain) | **Post** /tools/lookup-mta-sts-domain | Lookup a MTA-STS domain policy
[**LookupTlsReportingDomain**](ToolsControllerApi#LookupTlsReportingDomain) | **Post** /tools/lookup-tls-reporting-domain | Lookup a TLS reporting domain policy



## CheckEmailFeaturesClientSupport

> CheckEmailFeaturesClientSupportResults CheckEmailFeaturesClientSupport(ctx, checkEmailFeaturesClientSupportOptions)

Check email client support for email HTML and CSS features

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkEmailFeaturesClientSupportOptions** | [**CheckEmailFeaturesClientSupportOptions**](CheckEmailFeaturesClientSupportOptions)|  | 

### Return type

[**CheckEmailFeaturesClientSupportResults**](CheckEmailFeaturesClientSupportResults)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CreateNewFakeEmailAddress

> NewFakeEmailAddressResult CreateNewFakeEmailAddress(ctx, )

Create a new email address using the fake email domains

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**NewFakeEmailAddressResult**](NewFakeEmailAddressResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DeleteNewFakeEmailAddress

> DeleteNewFakeEmailAddress(ctx, emailAddress)

Delete a fake email address using the fake email domains

Delete a fake email address using the fake email domains

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailAddress** | **string**|  | 

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


## GenerateBimiRecord

> GenerateBimiRecordResults GenerateBimiRecord(ctx, generateBimiRecordOptions)

Create a BIMI record policy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**generateBimiRecordOptions** | [**GenerateBimiRecordOptions**](GenerateBimiRecordOptions)|  | 

### Return type

[**GenerateBimiRecordResults**](GenerateBimiRecordResults)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GenerateDmarcRecord

> GenerateDmarcRecordResults GenerateDmarcRecord(ctx, generateDmarcRecordOptions)

Create a DMARC record policy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**generateDmarcRecordOptions** | [**GenerateDmarcRecordOptions**](GenerateDmarcRecordOptions)|  | 

### Return type

[**GenerateDmarcRecordResults**](GenerateDmarcRecordResults)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GenerateMtaStsRecord

> GenerateMtaStsRecordResults GenerateMtaStsRecord(ctx, generateMtaStsRecordOptions)

Create a TLS reporting record policy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**generateMtaStsRecordOptions** | [**GenerateMtaStsRecordOptions**](GenerateMtaStsRecordOptions)|  | 

### Return type

[**GenerateMtaStsRecordResults**](GenerateMtaStsRecordResults)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GenerateTlsReportingRecord

> GenerateTlsReportingRecordResults GenerateTlsReportingRecord(ctx, generateTlsReportingRecordOptions)

Create a TLS reporting record policy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**generateTlsReportingRecordOptions** | [**GenerateTlsReportingRecordOptions**](GenerateTlsReportingRecordOptions)|  | 

### Return type

[**GenerateTlsReportingRecordResults**](GenerateTlsReportingRecordResults)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetFakeEmailByEmailAddress

> FakeEmailResult GetFakeEmailByEmailAddress(ctx, emailAddress)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailAddress** | **string**|  | 

### Return type

[**FakeEmailResult**](FakeEmailResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetFakeEmailById

> FakeEmailResult GetFakeEmailById(ctx, id)

Get a fake email by its ID

Get a fake email by its ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()|  | 

### Return type

[**FakeEmailResult**](FakeEmailResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetFakeEmailRaw

> string GetFakeEmailRaw(ctx, id)

Get raw fake email content

Retrieve the raw content of a fake email by its ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**]()|  | 

### Return type

**string**

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain;charset=utf-8, text/html;charset=utf-8, text/plain; charset=utf-8, text/html; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetFakeEmailsForAddress

> []FakeEmailPreview GetFakeEmailsForAddress(ctx, emailAddress, optional)

Get fake emails for an address

Get fake emails for an address

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailAddress** | **string**|  | 
 **optional** | ***GetFakeEmailsForAddressOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetFakeEmailsForAddressOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**|  | 

### Return type

[**[]FakeEmailPreview**](FakeEmailPreview)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## LookupBimiDomain

> LookupBimiDomainResults LookupBimiDomain(ctx, lookupBimiDomainOptions)

Lookup a BIMI record policy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lookupBimiDomainOptions** | [**LookupBimiDomainOptions**](LookupBimiDomainOptions)|  | 

### Return type

[**LookupBimiDomainResults**](LookupBimiDomainResults)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## LookupDmarcDomain

> LookupDmarcDomainResults LookupDmarcDomain(ctx, lookupDmarcDomainOptions)

Lookup a DMARC record policy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lookupDmarcDomainOptions** | [**LookupDmarcDomainOptions**](LookupDmarcDomainOptions)|  | 

### Return type

[**LookupDmarcDomainResults**](LookupDmarcDomainResults)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## LookupMtaStsDomain

> LookupMtaStsDomainResults LookupMtaStsDomain(ctx, lookupMtaStsDomainOptions)

Lookup a MTA-STS domain policy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lookupMtaStsDomainOptions** | [**LookupMtaStsDomainOptions**](LookupMtaStsDomainOptions)|  | 

### Return type

[**LookupMtaStsDomainResults**](LookupMtaStsDomainResults)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## LookupTlsReportingDomain

> LookupTlsReportingDomainResults LookupTlsReportingDomain(ctx, lookupTlsReportingDomainOptions)

Lookup a TLS reporting domain policy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lookupTlsReportingDomainOptions** | [**LookupTlsReportingDomainOptions**](LookupTlsReportingDomainOptions)|  | 

### Return type

[**LookupTlsReportingDomainResults**](LookupTlsReportingDomainResults)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

