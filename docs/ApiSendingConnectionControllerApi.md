# MailSlurp\ApiSendingConnectionControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSendingConnection**](ApiSendingConnectionControllerApi#CreateSendingConnection) | **Post** /sending-connections | Create a sending connection
[**DeleteDeliverySenderSeats**](ApiSendingConnectionControllerApi#DeleteDeliverySenderSeats) | **Delete** /sending-connections/sender-seats | Cancel purchased Delivery Sender seats
[**DeleteSendingConnection**](ApiSendingConnectionControllerApi#DeleteSendingConnection) | **Delete** /sending-connections/{connectionId} | Disconnect a sending connection
[**DisableSendingConnection**](ApiSendingConnectionControllerApi#DisableSendingConnection) | **Delete** /sending-connections/{connectionId}/enable | Disable Delivery Sender capacity
[**EnableSendingConnection**](ApiSendingConnectionControllerApi#EnableSendingConnection) | **Post** /sending-connections/{connectionId}/enable | Enable a verified Delivery Sender
[**GetDeliverySenderCapacity**](ApiSendingConnectionControllerApi#GetDeliverySenderCapacity) | **Get** /sending-connections/capacity | Get shared Delivery Sender capacity
[**GetSendingConnection**](ApiSendingConnectionControllerApi#GetSendingConnection) | **Get** /sending-connections/{connectionId} | Get a sending connection
[**GetSendingConnectionProviders**](ApiSendingConnectionControllerApi#GetSendingConnectionProviders) | **Get** /sending-connections/providers | List sending connection providers
[**GetSendingConnections**](ApiSendingConnectionControllerApi#GetSendingConnections) | **Get** /sending-connections | List sending connections
[**ProbeSendingConnectionDelivery**](ApiSendingConnectionControllerApi#ProbeSendingConnectionDelivery) | **Post** /sending-connections/{connectionId}/delivery-probe | Send an unbilled delivery test to a temporary MailSlurp inbox
[**RotateSendingConnectionCredentials**](ApiSendingConnectionControllerApi#RotateSendingConnectionCredentials) | **Post** /sending-connections/{connectionId}/credentials | Rotate sending connection credentials
[**UpdateDeliverySenderSeats**](ApiSendingConnectionControllerApi#UpdateDeliverySenderSeats) | **Put** /sending-connections/sender-seats | Update purchased Delivery Sender seats
[**UpdateSendingConnection**](ApiSendingConnectionControllerApi#UpdateSendingConnection) | **Patch** /sending-connections/{connectionId} | Update a sending connection
[**VerifySendingConnection**](ApiSendingConnectionControllerApi#VerifySendingConnection) | **Post** /sending-connections/{connectionId}/verify | Verify a sending connection



## CreateSendingConnection

> SendingConnectionDto CreateSendingConnection(ctx, createSendingConnectionOptions)

Create a sending connection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createSendingConnectionOptions** | [**CreateSendingConnectionOptions**](CreateSendingConnectionOptions)|  | 

### Return type

[**SendingConnectionDto**](SendingConnectionDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DeleteDeliverySenderSeats

> DeliverySenderCapacityDto DeleteDeliverySenderSeats(ctx, )

Cancel purchased Delivery Sender seats

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**DeliverySenderCapacityDto**](DeliverySenderCapacityDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DeleteSendingConnection

> DeleteSendingConnection(ctx, connectionId)

Disconnect a sending connection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | [**string**]()|  | 

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


## DisableSendingConnection

> SendingConnectionDto DisableSendingConnection(ctx, connectionId)

Disable Delivery Sender capacity

Releases this connection's shared sender allocation while retaining its verified credentials.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | [**string**]()|  | 

### Return type

[**SendingConnectionDto**](SendingConnectionDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## EnableSendingConnection

> SendingConnectionDto EnableSendingConnection(ctx, connectionId)

Enable a verified Delivery Sender

Allocates shared sender capacity. Included capacity is used first and supported additional capacity is added to the existing subscription when possible.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | [**string**]()|  | 

### Return type

[**SendingConnectionDto**](SendingConnectionDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDeliverySenderCapacity

> DeliverySenderCapacityDto GetDeliverySenderCapacity(ctx, )

Get shared Delivery Sender capacity

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**DeliverySenderCapacityDto**](DeliverySenderCapacityDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetSendingConnection

> SendingConnectionDto GetSendingConnection(ctx, connectionId)

Get a sending connection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | [**string**]()|  | 

### Return type

[**SendingConnectionDto**](SendingConnectionDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetSendingConnectionProviders

> []SendingConnectionProviderDto GetSendingConnectionProviders(ctx, )

List sending connection providers

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]SendingConnectionProviderDto**](SendingConnectionProviderDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetSendingConnections

> []SendingConnectionDto GetSendingConnections(ctx, )

List sending connections

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]SendingConnectionDto**](SendingConnectionDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## ProbeSendingConnectionDelivery

> SendingConnectionDeliveryProbeDto ProbeSendingConnectionDelivery(ctx, connectionId)

Send an unbilled delivery test to a temporary MailSlurp inbox

Uses only the selected customer-owned connection. A successful receipt proves provider acceptance and delivery to MailSlurp, not Gmail or Outlook placement.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | [**string**]()|  | 

### Return type

[**SendingConnectionDeliveryProbeDto**](SendingConnectionDeliveryProbeDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## RotateSendingConnectionCredentials

> SendingConnectionDto RotateSendingConnectionCredentials(ctx, connectionId, rotateSendingConnectionCredentialsOptions)

Rotate sending connection credentials

Verifies supported replacement credentials before atomically replacing the encrypted credential envelope.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | [**string**]()|  | 
**rotateSendingConnectionCredentialsOptions** | [**RotateSendingConnectionCredentialsOptions**](RotateSendingConnectionCredentialsOptions)|  | 

### Return type

[**SendingConnectionDto**](SendingConnectionDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## UpdateDeliverySenderSeats

> DeliverySenderCapacityDto UpdateDeliverySenderSeats(ctx, updateDeliverySenderSeatsOptions)

Update purchased Delivery Sender seats

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**updateDeliverySenderSeatsOptions** | [**UpdateDeliverySenderSeatsOptions**](UpdateDeliverySenderSeatsOptions)|  | 

### Return type

[**DeliverySenderCapacityDto**](DeliverySenderCapacityDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## UpdateSendingConnection

> SendingConnectionDto UpdateSendingConnection(ctx, connectionId, updateSendingConnectionOptions)

Update a sending connection

Updates non-secret connection metadata. Use the credentials endpoint to replace supported SMTP or API-key credentials.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | [**string**]()|  | 
**updateSendingConnectionOptions** | [**UpdateSendingConnectionOptions**](UpdateSendingConnectionOptions)|  | 

### Return type

[**SendingConnectionDto**](SendingConnectionDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## VerifySendingConnection

> SendingConnectionDto VerifySendingConnection(ctx, connectionId)

Verify a sending connection

Verifies the customer credential without enabling billable sender capacity. A consuming product enables the shared Delivery Sender seat when sending starts.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | [**string**]()|  | 

### Return type

[**SendingConnectionDto**](SendingConnectionDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

