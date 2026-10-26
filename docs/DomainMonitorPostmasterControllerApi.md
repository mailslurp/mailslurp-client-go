# MailSlurp\DomainMonitorPostmasterControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDomainMonitorPostmasterConnection**](DomainMonitorPostmasterControllerApi#CreateDomainMonitorPostmasterConnection) | **Post** /domain-monitor/postmaster/connect | Create Postmaster OAuth connection URL
[**DeleteDomainMonitorPostmasterConnection**](DomainMonitorPostmasterControllerApi#DeleteDomainMonitorPostmasterConnection) | **Delete** /domain-monitor/postmaster/connection | Disconnect Postmaster OAuth
[**ExchangeDomainMonitorPostmasterCode**](DomainMonitorPostmasterControllerApi#ExchangeDomainMonitorPostmasterCode) | **Post** /domain-monitor/postmaster/oauth-exchange | Exchange Postmaster OAuth authorization code
[**GetDomainMonitorPostmasterConnection**](DomainMonitorPostmasterControllerApi#GetDomainMonitorPostmasterConnection) | **Get** /domain-monitor/postmaster/connection | Get Postmaster connection status
[**GetDomainMonitorPostmasterDomains**](DomainMonitorPostmasterControllerApi#GetDomainMonitorPostmasterDomains) | **Get** /domain-monitor/postmaster/domains | List local Postmaster domain snapshots
[**GetDomainMonitorPostmasterTrafficStats**](DomainMonitorPostmasterControllerApi#GetDomainMonitorPostmasterTrafficStats) | **Get** /domain-monitor/postmaster/domains/{domain}/traffic-stats | Get local Postmaster traffic stats
[**SyncDomainMonitorPostmasterDomains**](DomainMonitorPostmasterControllerApi#SyncDomainMonitorPostmasterDomains) | **Post** /domain-monitor/postmaster/domains/sync | Sync Postmaster domains
[**SyncDomainMonitorPostmasterTrafficStats**](DomainMonitorPostmasterControllerApi#SyncDomainMonitorPostmasterTrafficStats) | **Post** /domain-monitor/postmaster/domains/{domain}/traffic-stats/sync | Sync Postmaster traffic stats



## CreateDomainMonitorPostmasterConnection

> CreatePostmasterConnectionResult CreateDomainMonitorPostmasterConnection(ctx, createPostmasterConnectionOptions)

Create Postmaster OAuth connection URL

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createPostmasterConnectionOptions** | [**CreatePostmasterConnectionOptions**](CreatePostmasterConnectionOptions)|  | 

### Return type

[**CreatePostmasterConnectionResult**](CreatePostmasterConnectionResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DeleteDomainMonitorPostmasterConnection

> DeleteDomainMonitorPostmasterConnection(ctx, )

Disconnect Postmaster OAuth

### Required Parameters

This endpoint does not need any parameter.

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


## ExchangeDomainMonitorPostmasterCode

> PostmasterConnectionDto ExchangeDomainMonitorPostmasterCode(ctx, exchangePostmasterCodeOptions)

Exchange Postmaster OAuth authorization code

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exchangePostmasterCodeOptions** | [**ExchangePostmasterCodeOptions**](ExchangePostmasterCodeOptions)|  | 

### Return type

[**PostmasterConnectionDto**](PostmasterConnectionDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDomainMonitorPostmasterConnection

> PostmasterConnectionStatusDto GetDomainMonitorPostmasterConnection(ctx, )

Get Postmaster connection status

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**PostmasterConnectionStatusDto**](PostmasterConnectionStatusDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDomainMonitorPostmasterDomains

> []PostmasterDomainDto GetDomainMonitorPostmasterDomains(ctx, )

List local Postmaster domain snapshots

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]PostmasterDomainDto**](PostmasterDomainDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDomainMonitorPostmasterTrafficStats

> PostmasterTrafficStatsListDto GetDomainMonitorPostmasterTrafficStats(ctx, domain, optional)

Get local Postmaster traffic stats

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string**|  | 
 **optional** | ***GetDomainMonitorPostmasterTrafficStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDomainMonitorPostmasterTrafficStatsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **optional.String**|  | 
 **endDate** | **optional.String**|  | 

### Return type

[**PostmasterTrafficStatsListDto**](PostmasterTrafficStatsListDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## SyncDomainMonitorPostmasterDomains

> []PostmasterDomainDto SyncDomainMonitorPostmasterDomains(ctx, )

Sync Postmaster domains

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]PostmasterDomainDto**](PostmasterDomainDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## SyncDomainMonitorPostmasterTrafficStats

> PostmasterTrafficStatsListDto SyncDomainMonitorPostmasterTrafficStats(ctx, domain, optional)

Sync Postmaster traffic stats

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string**|  | 
 **optional** | ***SyncDomainMonitorPostmasterTrafficStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SyncDomainMonitorPostmasterTrafficStatsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **optional.String**|  | 
 **endDate** | **optional.String**|  | 

### Return type

[**PostmasterTrafficStatsListDto**](PostmasterTrafficStatsListDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

