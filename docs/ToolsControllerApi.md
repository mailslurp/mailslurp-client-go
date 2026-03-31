# MailSlurp\ToolsControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalyzeDmarcReport**](ToolsControllerApi#AnalyzeDmarcReport) | **Post** /tools/analyze-dmarc-report | Parse and summarize a DMARC aggregate XML report
[**AnalyzeEmailHeaders**](ToolsControllerApi#AnalyzeEmailHeaders) | **Post** /tools/analyze-email-headers | Analyze email headers for auth results and delivery path
[**CheckCampaignProbe**](ToolsControllerApi#CheckCampaignProbe) | **Post** /tools/check-campaign-probe | Run a one-shot free campaign probe preflight check
[**CheckDnsPropagation**](ToolsControllerApi#CheckDnsPropagation) | **Post** /tools/check-dns-propagation | Check DNS propagation for a host and record type across configured resolvers
[**CheckDomainMonitor**](ToolsControllerApi#CheckDomainMonitor) | **Post** /tools/check-domain-monitor | Run a one-shot free domain monitor posture check
[**CheckEmailAudit**](ToolsControllerApi#CheckEmailAudit) | **Post** /tools/check-email-audit | Run a one-shot free email audit across links, images, HTML, and client support
[**CheckEmailAuthStack**](ToolsControllerApi#CheckEmailAuthStack) | **Post** /tools/check-email-auth-stack | Run a one-shot combined SPF, DKIM, DMARC, BIMI, MX, MTA-STS, and TLS-RPT check
[**CheckEmailBlacklist**](ToolsControllerApi#CheckEmailBlacklist) | **Post** /tools/check-email-blacklists | Check whether a domain or IP appears on configured DNS blacklists
[**CheckEmailFeaturesClientSupport**](ToolsControllerApi#CheckEmailFeaturesClientSupport) | **Post** /tools/check-email-features-client-support | Check email client support for email HTML and CSS features
[**CreateNewFakeEmailAddress**](ToolsControllerApi#CreateNewFakeEmailAddress) | **Post** /tools/fake-email | Create a new email address using the fake email domains
[**DeleteNewFakeEmailAddress**](ToolsControllerApi#DeleteNewFakeEmailAddress) | **Delete** /tools/fake-email | Delete a fake email address using the fake email domains
[**GenerateBimiRecord**](ToolsControllerApi#GenerateBimiRecord) | **Post** /tools/generate-bimi-record | Create a BIMI record policy
[**GenerateDmarcRecord**](ToolsControllerApi#GenerateDmarcRecord) | **Post** /tools/generate-dmarc-record | Create a DMARC record policy
[**GenerateMtaStsRecord**](ToolsControllerApi#GenerateMtaStsRecord) | **Post** /tools/generate-mta-sts-record | Create a TLS reporting record policy
[**GenerateSpfRecord**](ToolsControllerApi#GenerateSpfRecord) | **Post** /tools/generate-spf-record | Create an SPF record
[**GenerateTlsReportingRecord**](ToolsControllerApi#GenerateTlsReportingRecord) | **Post** /tools/generate-tls-reporting-record | Create a TLS reporting record policy
[**GetFakeEmailByEmailAddress**](ToolsControllerApi#GetFakeEmailByEmailAddress) | **Get** /tools/fake-email/byEmailAddress | 
[**GetFakeEmailById**](ToolsControllerApi#GetFakeEmailById) | **Get** /tools/fake-email | Get a fake email by its ID
[**GetFakeEmailRaw**](ToolsControllerApi#GetFakeEmailRaw) | **Get** /tools/fake-email/html | Get raw fake email content
[**GetFakeEmailsForAddress**](ToolsControllerApi#GetFakeEmailsForAddress) | **Get** /tools/fake-emails | Get fake emails for an address
[**LookupBimiDomain**](ToolsControllerApi#LookupBimiDomain) | **Post** /tools/lookup-bimi-domain | Lookup a BIMI record policy
[**LookupDkimDomain**](ToolsControllerApi#LookupDkimDomain) | **Post** /tools/lookup-dkim-domain | Lookup and validate a DKIM record
[**LookupDmarcDomain**](ToolsControllerApi#LookupDmarcDomain) | **Post** /tools/lookup-dmarc-domain | Lookup a DMARC record policy
[**LookupMtaStsDomain**](ToolsControllerApi#LookupMtaStsDomain) | **Post** /tools/lookup-mta-sts-domain | Lookup a MTA-STS domain policy
[**LookupMxRecord**](ToolsControllerApi#LookupMxRecord) | **Post** /tools/lookup-mx-records | Lookup a MX records for a domain
[**LookupPtr**](ToolsControllerApi#LookupPtr) | **Post** /tools/lookup-ptr | Lookup PTR records for an IP address
[**LookupSpfDomain**](ToolsControllerApi#LookupSpfDomain) | **Post** /tools/lookup-spf-domain | Lookup and validate an SPF record
[**LookupTlsReportingDomain**](ToolsControllerApi#LookupTlsReportingDomain) | **Post** /tools/lookup-tls-reporting-domain | Lookup a TLS reporting domain policy
[**TestSmtpServer**](ToolsControllerApi#TestSmtpServer) | **Post** /tools/test-smtp-server | Run a conservative SMTP connectivity, TLS, and AUTH diagnostic



## AnalyzeDmarcReport

> AnalyzeDmarcReportResults AnalyzeDmarcReport(ctx, analyzeDmarcReportOptions)

Parse and summarize a DMARC aggregate XML report

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analyzeDmarcReportOptions** | [**AnalyzeDmarcReportOptions**](AnalyzeDmarcReportOptions)|  | 

### Return type

[**AnalyzeDmarcReportResults**](AnalyzeDmarcReportResults)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## AnalyzeEmailHeaders

> AnalyzeEmailHeadersResults AnalyzeEmailHeaders(ctx, analyzeEmailHeadersOptions)

Analyze email headers for auth results and delivery path

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analyzeEmailHeadersOptions** | [**AnalyzeEmailHeadersOptions**](AnalyzeEmailHeadersOptions)|  | 

### Return type

[**AnalyzeEmailHeadersResults**](AnalyzeEmailHeadersResults)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CheckCampaignProbe

> CheckCampaignProbeResults CheckCampaignProbe(ctx, checkCampaignProbeOptions)

Run a one-shot free campaign probe preflight check

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkCampaignProbeOptions** | [**CheckCampaignProbeOptions**](CheckCampaignProbeOptions)|  | 

### Return type

[**CheckCampaignProbeResults**](CheckCampaignProbeResults)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CheckDnsPropagation

> CheckDnsPropagationResults CheckDnsPropagation(ctx, checkDnsPropagationOptions)

Check DNS propagation for a host and record type across configured resolvers

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkDnsPropagationOptions** | [**CheckDnsPropagationOptions**](CheckDnsPropagationOptions)|  | 

### Return type

[**CheckDnsPropagationResults**](CheckDnsPropagationResults)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CheckDomainMonitor

> CheckDomainMonitorResults CheckDomainMonitor(ctx, checkDomainMonitorOptions)

Run a one-shot free domain monitor posture check

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkDomainMonitorOptions** | [**CheckDomainMonitorOptions**](CheckDomainMonitorOptions)|  | 

### Return type

[**CheckDomainMonitorResults**](CheckDomainMonitorResults)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CheckEmailAudit

> EmailAuditAnalysisResult CheckEmailAudit(ctx, checkEmailAuditOptions)

Run a one-shot free email audit across links, images, HTML, and client support

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkEmailAuditOptions** | [**CheckEmailAuditOptions**](CheckEmailAuditOptions)|  | 

### Return type

[**EmailAuditAnalysisResult**](EmailAuditAnalysisResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CheckEmailAuthStack

> CheckEmailAuthStackResults CheckEmailAuthStack(ctx, checkEmailAuthStackOptions)

Run a one-shot combined SPF, DKIM, DMARC, BIMI, MX, MTA-STS, and TLS-RPT check

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkEmailAuthStackOptions** | [**CheckEmailAuthStackOptions**](CheckEmailAuthStackOptions)|  | 

### Return type

[**CheckEmailAuthStackResults**](CheckEmailAuthStackResults)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CheckEmailBlacklist

> CheckEmailBlacklistResults CheckEmailBlacklist(ctx, checkEmailBlacklistOptions)

Check whether a domain or IP appears on configured DNS blacklists

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkEmailBlacklistOptions** | [**CheckEmailBlacklistOptions**](CheckEmailBlacklistOptions)|  | 

### Return type

[**CheckEmailBlacklistResults**](CheckEmailBlacklistResults)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


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


## GenerateSpfRecord

> GenerateSpfRecordResults GenerateSpfRecord(ctx, generateSpfRecordOptions)

Create an SPF record

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**generateSpfRecordOptions** | [**GenerateSpfRecordOptions**](GenerateSpfRecordOptions)|  | 

### Return type

[**GenerateSpfRecordResults**](GenerateSpfRecordResults)

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


## LookupDkimDomain

> LookupDkimDomainResults LookupDkimDomain(ctx, lookupDkimDomainOptions)

Lookup and validate a DKIM record

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lookupDkimDomainOptions** | [**LookupDkimDomainOptions**](LookupDkimDomainOptions)|  | 

### Return type

[**LookupDkimDomainResults**](LookupDkimDomainResults)

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


## LookupMxRecord

> LookupMxRecordsResults LookupMxRecord(ctx, lookupMxRecordsOptions)

Lookup a MX records for a domain

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lookupMxRecordsOptions** | [**LookupMxRecordsOptions**](LookupMxRecordsOptions)|  | 

### Return type

[**LookupMxRecordsResults**](LookupMxRecordsResults)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## LookupPtr

> LookupPtrResults LookupPtr(ctx, lookupPtrOptions)

Lookup PTR records for an IP address

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lookupPtrOptions** | [**LookupPtrOptions**](LookupPtrOptions)|  | 

### Return type

[**LookupPtrResults**](LookupPtrResults)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## LookupSpfDomain

> LookupSpfDomainResults LookupSpfDomain(ctx, lookupSpfDomainOptions)

Lookup and validate an SPF record

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lookupSpfDomainOptions** | [**LookupSpfDomainOptions**](LookupSpfDomainOptions)|  | 

### Return type

[**LookupSpfDomainResults**](LookupSpfDomainResults)

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


## TestSmtpServer

> TestSmtpServerResults TestSmtpServer(ctx, testSmtpServerOptions)

Run a conservative SMTP connectivity, TLS, and AUTH diagnostic

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testSmtpServerOptions** | [**TestSmtpServerOptions**](TestSmtpServerOptions)|  | 

### Return type

[**TestSmtpServerResults**](TestSmtpServerResults)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

