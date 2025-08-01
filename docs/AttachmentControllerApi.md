# MailSlurp\AttachmentControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAllAttachments**](AttachmentControllerApi#DeleteAllAttachments) | **Delete** /attachments | Delete all attachments
[**DeleteAttachment**](AttachmentControllerApi#DeleteAttachment) | **Delete** /attachments/{attachmentId} | Delete an attachment
[**DownloadAttachmentAsBase64Encoded**](AttachmentControllerApi#DownloadAttachmentAsBase64Encoded) | **Get** /attachments/{attachmentId}/base64 | Get email attachment as base64 encoded string as alternative to binary responses. To read the content decode the Base64 encoded contents.
[**DownloadAttachmentAsBytes**](AttachmentControllerApi#DownloadAttachmentAsBytes) | **Get** /attachments/{attachmentId}/bytes | Download attachments. Get email attachment bytes. If you have trouble with byte responses try the &#x60;downloadAttachmentBase64&#x60; response endpoints.
[**GetAttachment**](AttachmentControllerApi#GetAttachment) | **Get** /attachments/{attachmentId} | Get an attachment entity
[**GetAttachmentInfo**](AttachmentControllerApi#GetAttachmentInfo) | **Get** /attachments/{attachmentId}/metadata | Get email attachment metadata information
[**GetAttachments**](AttachmentControllerApi#GetAttachments) | **Get** /attachments | Get email attachments
[**UploadAttachment**](AttachmentControllerApi#UploadAttachment) | **Post** /attachments | Upload an attachment for sending using base64 file encoding. Returns an array whose first element is the ID of the uploaded attachment.
[**UploadAttachmentBytes**](AttachmentControllerApi#UploadAttachmentBytes) | **Post** /attachments/bytes | Upload an attachment for sending using file byte stream input octet stream. Returns an array whose first element is the ID of the uploaded attachment.
[**UploadMultipartForm**](AttachmentControllerApi#UploadMultipartForm) | **Post** /attachments/multipart | Upload an attachment for sending using a Multipart Form request. Returns an array whose first element is the ID of the uploaded attachment.



## DeleteAllAttachments

> DeleteAllAttachments(ctx, )

Delete all attachments

Delete all attachments

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


## DeleteAttachment

> DeleteAttachment(ctx, attachmentId)

Delete an attachment

Delete an attachment

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attachmentId** | **string**| ID of attachment | 

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


## DownloadAttachmentAsBase64Encoded

> DownloadAttachmentDto DownloadAttachmentAsBase64Encoded(ctx, attachmentId)

Get email attachment as base64 encoded string as alternative to binary responses. To read the content decode the Base64 encoded contents.

Returns the specified attachment for a given email as a base 64 encoded string. The response type is application/json. This method is similar to the `downloadAttachment` method but allows some clients to get around issues with binary responses.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attachmentId** | **string**| ID of attachment | 

### Return type

[**DownloadAttachmentDto**](DownloadAttachmentDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DownloadAttachmentAsBytes

> string DownloadAttachmentAsBytes(ctx, attachmentId)

Download attachments. Get email attachment bytes. If you have trouble with byte responses try the `downloadAttachmentBase64` response endpoints.

Returns the specified attachment for a given email as a stream / array of bytes. You can find attachment ids in email responses endpoint responses. The response type is application/octet-stream.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attachmentId** | **string**| ID of attachment | 

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


## GetAttachment

> AttachmentEntityDto GetAttachment(ctx, attachmentId)

Get an attachment entity

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attachmentId** | **string**| ID of attachment | 

### Return type

[**AttachmentEntityDto**](AttachmentEntityDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetAttachmentInfo

> AttachmentMetaData GetAttachmentInfo(ctx, attachmentId)

Get email attachment metadata information

Returns the metadata for an attachment. It is saved separately to the content of the attachment. Contains properties `name` and `content-type` and `content-length` in bytes for a given attachment.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attachmentId** | **string**| ID of attachment | 

### Return type

[**AttachmentMetaData**](AttachmentMetaData)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetAttachments

> PageAttachmentEntity GetAttachments(ctx, optional)

Get email attachments

Get all attachments in paginated response. Each entity contains meta data for the attachment such as `name` and `content-type`. Use the `attachmentId` and the download endpoints to get the file contents.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAttachmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAttachmentsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Optional page index for list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size for list pagination | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to ASC]
 **fileNameFilter** | **optional.String**| Optional file name and content type search filter | 
 **since** | **optional.Time**| Filter by created at after the given timestamp | 
 **before** | **optional.Time**| Filter by created at before the given timestamp | 
 **inboxId** | [**optional.Interface of string**]()| Optional inboxId to filter attachments by | 
 **emailId** | [**optional.Interface of string**]()| Optional emailId to filter attachments by | 
 **sentEmailId** | [**optional.Interface of string**]()| Optional sentEmailId to filter attachments by | 
 **include** | [**optional.Interface of []string**](string)| Optional list of IDs to include in result | 

### Return type

[**PageAttachmentEntity**](PageAttachmentEntity)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## UploadAttachment

> []string UploadAttachment(ctx, uploadAttachmentOptions)

Upload an attachment for sending using base64 file encoding. Returns an array whose first element is the ID of the uploaded attachment.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uploadAttachmentOptions** | [**UploadAttachmentOptions**](UploadAttachmentOptions)|  | 

### Return type

**[]string**

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## UploadAttachmentBytes

> []string UploadAttachmentBytes(ctx, optional)

Upload an attachment for sending using file byte stream input octet stream. Returns an array whose first element is the ID of the uploaded attachment.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UploadAttachmentBytesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UploadAttachmentBytesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **optional.String**|  | 
 **contentType2** | **optional.String**| Optional contentType for file. For instance &#x60;application/pdf&#x60; | 
 **contentId** | **optional.String**| Optional content ID (CID) to save upload with | 
 **filename** | **optional.String**| Optional filename to save upload with | 
 **fileSize** | **optional.Int64**| Optional byte length to save upload with | 
 **filename2** | **optional.String**|  | 

### Return type

**[]string**

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## UploadMultipartForm

> []string UploadMultipartForm(ctx, optional)

Upload an attachment for sending using a Multipart Form request. Returns an array whose first element is the ID of the uploaded attachment.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UploadMultipartFormOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UploadMultipartFormOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentId** | **optional.String**| Optional content ID of attachment | 
 **contentType** | **optional.String**| Optional content type of attachment | 
 **filename** | **optional.String**| Optional name of file | 
 **contentTypeHeader** | **optional.String**| Optional content type header of attachment | 
 **xFilename** | **optional.String**| Optional filename header of attachment | 
 **xFilenameRaw** | **optional.String**| Optional raw filename header of attachment that will be converted to punycode | 
 **xFilesize** | **optional.Int64**| Optional content size of attachment | 
 **inlineObject1** | [**optional.Interface of InlineObject1**](InlineObject1)|  | 

### Return type

**[]string**

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

