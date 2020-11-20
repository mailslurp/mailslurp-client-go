# MailSlurp\AttachmentControllerApi

All URIs are relative to *https://api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UploadAttachment**](AttachmentControllerApi.md#UploadAttachment) | **Post** /attachments | Upload an attachment for sending using base64 file encoding
[**UploadAttachmentBytes**](AttachmentControllerApi.md#UploadAttachmentBytes) | **Post** /attachments/bytes | Upload an attachment for sending using file byte stream input octet stream
[**UploadMultipartForm**](AttachmentControllerApi.md#UploadMultipartForm) | **Post** /attachments/multipart | Upload an attachment for sending using Multipart Form



## UploadAttachment

> []string UploadAttachment(ctx, uploadOptions)

Upload an attachment for sending using base64 file encoding

When sending emails with attachments first upload each attachment with this endpoint. Record the returned attachment IDs. Then use these attachment IDs in the SendEmailOptions when sending an email. This means that attachments can easily be reused.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uploadOptions** | [**UploadAttachmentOptions**](UploadAttachmentOptions.md)| uploadOptions | 

### Return type

**[]string**

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadAttachmentBytes

> []string UploadAttachmentBytes(ctx, optional)

Upload an attachment for sending using file byte stream input octet stream

When sending emails with attachments first upload each attachment with this endpoint. Record the returned attachment IDs. Then use these attachment IDs in the SendEmailOptions when sending an email. This means that attachments can easily be reused.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UploadAttachmentBytesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UploadAttachmentBytesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filename** | **optional.String**| Optional filename to save upload with | 

### Return type

**[]string**

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadMultipartForm

> []string UploadMultipartForm(ctx, file, optional)

Upload an attachment for sending using Multipart Form

When sending emails with attachments first upload each attachment with this endpoint. Record the returned attachment IDs. Then use these attachment IDs in the SendEmailOptions when sending an email. This means that attachments can easily be reused.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**file** | ***os.File*****os.File**| file | 
 **optional** | ***UploadMultipartFormOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UploadMultipartFormOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **optional.String**| contentType | 
 **filename** | **optional.String**| filename | 
 **xFilename** | **optional.String**| x-filename | 

### Return type

**[]string**

### Authorization

[API_KEY](../README.md#API_KEY)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

