# MailSlurp\DevicePreviewsControllerApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelDevicePreviewRun**](DevicePreviewsControllerApi#CancelDevicePreviewRun) | **Post** /emails/device-previews/{runId}/cancel | Cancel a running device preview run
[**CreateDevicePreviewFeedback**](DevicePreviewsControllerApi#CreateDevicePreviewFeedback) | **Post** /emails/device-previews/feedback | Create device preview feedback
[**CreateDevicePreviewFolder**](DevicePreviewsControllerApi#CreateDevicePreviewFolder) | **Post** /emails/device-previews/folders | Create a device preview folder
[**CreateDevicePreviewImportAddress**](DevicePreviewsControllerApi#CreateDevicePreviewImportAddress) | **Post** /emails/device-previews/import-addresses | Create a temporary email address for device preview import
[**CreateDevicePreviewProfile**](DevicePreviewsControllerApi#CreateDevicePreviewProfile) | **Post** /emails/device-previews/profiles | Create a device preview profile
[**CreateDevicePreviewRun**](DevicePreviewsControllerApi#CreateDevicePreviewRun) | **Post** /emails/{emailId}/device-previews | Create a new device preview run for an email
[**CreateDevicePreviewRunFromFile**](DevicePreviewsControllerApi#CreateDevicePreviewRunFromFile) | **Post** /emails/device-previews/import/file | Create a device preview run from an HTML, EML, or ZIP file
[**CreateDevicePreviewRunFromHtmlImport**](DevicePreviewsControllerApi#CreateDevicePreviewRunFromHtmlImport) | **Post** /emails/device-previews/import/html | Create a device preview run from HTML
[**CreateDevicePreviewRunFromImport**](DevicePreviewsControllerApi#CreateDevicePreviewRunFromImport) | **Post** /emails/device-previews/import | Create a device preview run from imported raw MIME
[**CreateDevicePreviewRunFromMultipartImport**](DevicePreviewsControllerApi#CreateDevicePreviewRunFromMultipartImport) | **Post** /emails/device-previews/import/multipart | Create a device preview run from multipart EML upload
[**CreateDevicePreviewRunFromRawImport**](DevicePreviewsControllerApi#CreateDevicePreviewRunFromRawImport) | **Post** /emails/device-previews/import/raw | Create a device preview run from raw MIME bytes
[**CreateDevicePreviewShareLink**](DevicePreviewsControllerApi#CreateDevicePreviewShareLink) | **Post** /emails/device-previews/{runId}/share-links | Create a share link for a device preview run
[**DeleteDevicePreviewFolder**](DevicePreviewsControllerApi#DeleteDevicePreviewFolder) | **Delete** /emails/device-previews/folders/{folderId} | Delete a folder, preserving renders and child folders
[**DeleteDevicePreviewProfile**](DevicePreviewsControllerApi#DeleteDevicePreviewProfile) | **Delete** /emails/device-previews/profiles/{profileId} | Delete a device preview profile
[**DeleteDevicePreviewRun**](DevicePreviewsControllerApi#DeleteDevicePreviewRun) | **Delete** /emails/device-previews/{runId} | Delete local device preview run data
[**DeleteDevicePreviewShareLink**](DevicePreviewsControllerApi#DeleteDevicePreviewShareLink) | **Delete** /emails/device-previews/share-links/{shareLinkId} | Revoke a device preview share link
[**DownloadDevicePreviewEmailSubmissionSourceEmail**](DevicePreviewsControllerApi#DownloadDevicePreviewEmailSubmissionSourceEmail) | **Get** /emails/device-previews/email-submissions/{submissionId}/source-email/raw | Download the original RFC822 email for a device preview submission
[**DownloadDevicePreviewRunSourceEmail**](DevicePreviewsControllerApi#DownloadDevicePreviewRunSourceEmail) | **Get** /emails/device-previews/{runId}/source-email/raw | Download the original RFC822 email for a device preview run
[**EnsureDevicePreviewRun**](DevicePreviewsControllerApi#EnsureDevicePreviewRun) | **Put** /emails/{emailId}/device-previews/latest | Return active run for email or create one when none exists
[**GetDevicePreviewEmailDomainOrCreate**](DevicePreviewsControllerApi#GetDevicePreviewEmailDomainOrCreate) | **Put** /emails/device-previews/email-domain | Return account device preview email domain setup or create one
[**GetDevicePreviewEmailRuns**](DevicePreviewsControllerApi#GetDevicePreviewEmailRuns) | **Get** /emails/device-previews/email-runs | List device preview runs created from account email domains
[**GetDevicePreviewEmailSubmissionSourceEmail**](DevicePreviewsControllerApi#GetDevicePreviewEmailSubmissionSourceEmail) | **Get** /emails/device-previews/email-submissions/{submissionId}/source-email | Get the original email content for a device preview submission
[**GetDevicePreviewEmailSubmissions**](DevicePreviewsControllerApi#GetDevicePreviewEmailSubmissions) | **Get** /emails/device-previews/email-submissions | List received device preview email submissions
[**GetDevicePreviewFeedback**](DevicePreviewsControllerApi#GetDevicePreviewFeedback) | **Get** /emails/device-previews/feedback/{feedbackId} | Get a single device preview feedback item
[**GetDevicePreviewFeedbackItems**](DevicePreviewsControllerApi#GetDevicePreviewFeedbackItems) | **Get** /emails/device-previews/feedback | List device preview feedback
[**GetDevicePreviewFolder**](DevicePreviewsControllerApi#GetDevicePreviewFolder) | **Get** /emails/device-previews/folders/{folderId} | Get a device preview folder
[**GetDevicePreviewFolders**](DevicePreviewsControllerApi#GetDevicePreviewFolders) | **Get** /emails/device-previews/folders | List device preview folders
[**GetDevicePreviewImportAddress**](DevicePreviewsControllerApi#GetDevicePreviewImportAddress) | **Get** /emails/device-previews/import-addresses/{sessionId} | Get temporary device preview import address status
[**GetDevicePreviewNativeTargetAvailability**](DevicePreviewsControllerApi#GetDevicePreviewNativeTargetAvailability) | **Get** /emails/device-previews/native-targets/availability | Get native device preview target availability
[**GetDevicePreviewProfile**](DevicePreviewsControllerApi#GetDevicePreviewProfile) | **Get** /emails/device-previews/profiles/{profileId} | Get a device preview profile
[**GetDevicePreviewProfiles**](DevicePreviewsControllerApi#GetDevicePreviewProfiles) | **Get** /emails/device-previews/profiles | List device preview profiles
[**GetDevicePreviewRun**](DevicePreviewsControllerApi#GetDevicePreviewRun) | **Get** /emails/device-previews/{runId} | Get device preview run status
[**GetDevicePreviewRunProviderProgress**](DevicePreviewsControllerApi#GetDevicePreviewRunProviderProgress) | **Get** /emails/device-previews/{runId}/providers/{provider} | Get provider-level progress for a device preview run
[**GetDevicePreviewRunResults**](DevicePreviewsControllerApi#GetDevicePreviewRunResults) | **Get** /emails/device-previews/{runId}/results | Get device preview run results
[**GetDevicePreviewRunScreenshot**](DevicePreviewsControllerApi#GetDevicePreviewRunScreenshot) | **Get** /emails/device-previews/{runId}/screenshots/{screenshotId}/image | Get a device preview screenshot image
[**GetDevicePreviewRunSourceEmail**](DevicePreviewsControllerApi#GetDevicePreviewRunSourceEmail) | **Get** /emails/device-previews/{runId}/source-email | Get the original email content for a device preview run
[**GetDevicePreviewRuns**](DevicePreviewsControllerApi#GetDevicePreviewRuns) | **Get** /emails/{emailId}/device-previews | List previous device preview runs for an email
[**GetDevicePreviewRunsForAccount**](DevicePreviewsControllerApi#GetDevicePreviewRunsForAccount) | **Get** /emails/device-previews | List previous device preview runs for account
[**GetDevicePreviewRunsForAccountOffsetPaginated**](DevicePreviewsControllerApi#GetDevicePreviewRunsForAccountOffsetPaginated) | **Get** /emails/device-previews/offset-paginated | List previous device preview runs for account in paginated form
[**GetDevicePreviewRunsOffsetPaginated**](DevicePreviewsControllerApi#GetDevicePreviewRunsOffsetPaginated) | **Get** /emails/{emailId}/device-previews/offset-paginated | List previous device preview runs for an email in paginated form
[**GetDevicePreviewShareLinks**](DevicePreviewsControllerApi#GetDevicePreviewShareLinks) | **Get** /emails/device-previews/{runId}/share-links | List share links for a device preview run
[**GetDevicePreviewSharedResult**](DevicePreviewsControllerApi#GetDevicePreviewSharedResult) | **Get** /emails/device-previews/share/{shareToken} | Get a public shared device preview result
[**GetDevicePreviewSharedResultScreenshot**](DevicePreviewsControllerApi#GetDevicePreviewSharedResultScreenshot) | **Get** /emails/device-previews/share/{shareToken}/screenshots/{screenshotId}/image | Get a public shared device preview screenshot image
[**GetDevicePreviewTargets**](DevicePreviewsControllerApi#GetDevicePreviewTargets) | **Get** /device-preview-targets | List device preview targets
[**MoveDevicePreviewRuns**](DevicePreviewsControllerApi#MoveDevicePreviewRuns) | **Put** /emails/device-previews/folders/runs | Move device preview runs to a folder or Unfiled
[**RerunDevicePreviewTargets**](DevicePreviewsControllerApi#RerunDevicePreviewTargets) | **Post** /emails/device-previews/{runId}/targets/rerun | Rerun selected device preview targets in the same run
[**UpdateDevicePreviewEmailDomainSubdomain**](DevicePreviewsControllerApi#UpdateDevicePreviewEmailDomainSubdomain) | **Put** /emails/device-previews/email-domain/subdomain | Update account device preview email subdomain
[**UpdateDevicePreviewFeedback**](DevicePreviewsControllerApi#UpdateDevicePreviewFeedback) | **Put** /emails/device-previews/feedback/{feedbackId} | Update device preview feedback
[**UpdateDevicePreviewFolder**](DevicePreviewsControllerApi#UpdateDevicePreviewFolder) | **Put** /emails/device-previews/folders/{folderId} | Rename or move a device preview folder
[**UpdateDevicePreviewProfile**](DevicePreviewsControllerApi#UpdateDevicePreviewProfile) | **Put** /emails/device-previews/profiles/{profileId} | Update a device preview profile
[**WaitForDevicePreviewEmailRuns**](DevicePreviewsControllerApi#WaitForDevicePreviewEmailRuns) | **Get** /emails/device-previews/email-runs/wait | Wait for device preview runs created from account email domains
[**WaitForDevicePreviewEmailSubmissions**](DevicePreviewsControllerApi#WaitForDevicePreviewEmailSubmissions) | **Get** /emails/device-previews/email-submissions/wait | Wait for received device preview email submissions
[**WaitForDevicePreviewRun**](DevicePreviewsControllerApi#WaitForDevicePreviewRun) | **Get** /emails/device-previews/{runId}/wait | Wait for device preview run to complete



## CancelDevicePreviewRun

> CancelDevicePreviewRunResult CancelDevicePreviewRun(ctx, runId, optional)

Cancel a running device preview run

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | [**string**]()|  | 
 **optional** | ***CancelDevicePreviewRunOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CancelDevicePreviewRunOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cancelDevicePreviewRunOptions** | [**optional.Interface of CancelDevicePreviewRunOptions**](CancelDevicePreviewRunOptions)|  | 

### Return type

[**CancelDevicePreviewRunResult**](CancelDevicePreviewRunResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CreateDevicePreviewFeedback

> DevicePreviewFeedbackDto CreateDevicePreviewFeedback(ctx, createDevicePreviewFeedbackOptions)

Create device preview feedback

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createDevicePreviewFeedbackOptions** | [**CreateDevicePreviewFeedbackOptions**](CreateDevicePreviewFeedbackOptions)|  | 

### Return type

[**DevicePreviewFeedbackDto**](DevicePreviewFeedbackDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CreateDevicePreviewFolder

> DevicePreviewFolderDto CreateDevicePreviewFolder(ctx, createDevicePreviewFolderOptions)

Create a device preview folder

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createDevicePreviewFolderOptions** | [**CreateDevicePreviewFolderOptions**](CreateDevicePreviewFolderOptions)|  | 

### Return type

[**DevicePreviewFolderDto**](DevicePreviewFolderDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CreateDevicePreviewImportAddress

> DevicePreviewIngestAddressDto CreateDevicePreviewImportAddress(ctx, optional)

Create a temporary email address for device preview import

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateDevicePreviewImportAddressOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateDevicePreviewImportAddressOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDevicePreviewIngestAddressOptions** | [**optional.Interface of CreateDevicePreviewIngestAddressOptions**](CreateDevicePreviewIngestAddressOptions)|  | 

### Return type

[**DevicePreviewIngestAddressDto**](DevicePreviewIngestAddressDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CreateDevicePreviewProfile

> DevicePreviewProfileDto CreateDevicePreviewProfile(ctx, createDevicePreviewProfileOptions)

Create a device preview profile

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createDevicePreviewProfileOptions** | [**CreateDevicePreviewProfileOptions**](CreateDevicePreviewProfileOptions)|  | 

### Return type

[**DevicePreviewProfileDto**](DevicePreviewProfileDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CreateDevicePreviewRun

> CreateDevicePreviewRunResult CreateDevicePreviewRun(ctx, emailId, optional)

Create a new device preview run for an email

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()|  | 
 **optional** | ***CreateDevicePreviewRunOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateDevicePreviewRunOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDevicePreviewOptions** | [**optional.Interface of CreateDevicePreviewOptions**](CreateDevicePreviewOptions)|  | 

### Return type

[**CreateDevicePreviewRunResult**](CreateDevicePreviewRunResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CreateDevicePreviewRunFromFile

> CreateDevicePreviewRunResult CreateDevicePreviewRunFromFile(ctx, file, optional)

Create a device preview run from an HTML, EML, or ZIP file

Uploads an HTML or RFC822 EML file, or a ZIP archive containing exactly one HTML or EML file. ZIP uploads are bounded by entry count, expanded size, and compression ratio.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**file** | ***os.File*****os.File**| An .html, .htm, or .eml file, or a .zip archive containing exactly one supported file. | 
 **optional** | ***CreateDevicePreviewRunFromFileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateDevicePreviewRunFromFileOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subject** | **optional.String**|  | 
 **from** | **optional.String**|  | 
 **to** | **optional.String**|  | 
 **sourceAlias** | **optional.String**|  | 
 **externalId** | **optional.String**|  | 
 **providers** | [**optional.Interface of []string**](string)| Legacy web-provider keys to render. | 
 **allTargets** | **optional.Bool**| Render every currently configured target. | 
 **targets** | [**optional.Interface of []string**](string)| Device preview target keys to render. | 
 **emailAudit** | **optional.Bool**|  | 
 **domainMonitor** | **optional.Bool**|  | 
 **senderDomain** | **optional.String**|  | 
 **folderId** | [**optional.Interface of string**]()|  | 

### Return type

[**CreateDevicePreviewRunResult**](CreateDevicePreviewRunResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CreateDevicePreviewRunFromHtmlImport

> CreateDevicePreviewRunResult CreateDevicePreviewRunFromHtmlImport(ctx, createDevicePreviewHtmlImportOptions)

Create a device preview run from HTML

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createDevicePreviewHtmlImportOptions** | [**CreateDevicePreviewHtmlImportOptions**](CreateDevicePreviewHtmlImportOptions)|  | 

### Return type

[**CreateDevicePreviewRunResult**](CreateDevicePreviewRunResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CreateDevicePreviewRunFromImport

> CreateDevicePreviewRunResult CreateDevicePreviewRunFromImport(ctx, createDevicePreviewImportOptions)

Create a device preview run from imported raw MIME

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createDevicePreviewImportOptions** | [**CreateDevicePreviewImportOptions**](CreateDevicePreviewImportOptions)|  | 

### Return type

[**CreateDevicePreviewRunResult**](CreateDevicePreviewRunResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CreateDevicePreviewRunFromMultipartImport

> CreateDevicePreviewRunResult CreateDevicePreviewRunFromMultipartImport(ctx, file, optional)

Create a device preview run from multipart EML upload

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**file** | ***os.File*****os.File**|  | 
 **optional** | ***CreateDevicePreviewRunFromMultipartImportOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateDevicePreviewRunFromMultipartImportOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sourceAlias** | **optional.String**|  | 
 **externalId** | **optional.String**|  | 
 **allTargets** | **optional.Bool**| Render every currently configured target. | 
 **targets** | [**optional.Interface of []string**](string)| Device preview target keys to render. | 
 **folderId** | [**optional.Interface of string**]()|  | 

### Return type

[**CreateDevicePreviewRunResult**](CreateDevicePreviewRunResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CreateDevicePreviewRunFromRawImport

> CreateDevicePreviewRunResult CreateDevicePreviewRunFromRawImport(ctx, optional)

Create a device preview run from raw MIME bytes

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateDevicePreviewRunFromRawImportOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateDevicePreviewRunFromRawImportOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceAlias** | **optional.String**|  | 
 **externalId** | **optional.String**|  | 
 **allTargets** | **optional.Bool**| Render every currently configured target. | 
 **targets** | [**optional.Interface of []string**](string)| Device preview target keys to render. | 
 **folderId** | [**optional.Interface of string**]()|  | 

### Return type

[**CreateDevicePreviewRunResult**](CreateDevicePreviewRunResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CreateDevicePreviewShareLink

> DevicePreviewShareLinkDto CreateDevicePreviewShareLink(ctx, runId, optional)

Create a share link for a device preview run

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | [**string**]()|  | 
 **optional** | ***CreateDevicePreviewShareLinkOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateDevicePreviewShareLinkOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDevicePreviewShareLinkOptions** | [**optional.Interface of CreateDevicePreviewShareLinkOptions**](CreateDevicePreviewShareLinkOptions)|  | 

### Return type

[**DevicePreviewShareLinkDto**](DevicePreviewShareLinkDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DeleteDevicePreviewFolder

> DeleteDevicePreviewFolderResult DeleteDevicePreviewFolder(ctx, folderId)

Delete a folder, preserving renders and child folders

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | [**string**]()|  | 

### Return type

[**DeleteDevicePreviewFolderResult**](DeleteDevicePreviewFolderResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DeleteDevicePreviewProfile

> DeleteDevicePreviewProfileResult DeleteDevicePreviewProfile(ctx, profileId)

Delete a device preview profile

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | [**string**]()|  | 

### Return type

[**DeleteDevicePreviewProfileResult**](DeleteDevicePreviewProfileResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DeleteDevicePreviewRun

> DeleteDevicePreviewRunResult DeleteDevicePreviewRun(ctx, runId)

Delete local device preview run data

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | [**string**]()|  | 

### Return type

[**DeleteDevicePreviewRunResult**](DeleteDevicePreviewRunResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DeleteDevicePreviewShareLink

> DeleteDevicePreviewShareLinkResult DeleteDevicePreviewShareLink(ctx, shareLinkId)

Revoke a device preview share link

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shareLinkId** | [**string**]()|  | 

### Return type

[**DeleteDevicePreviewShareLinkResult**](DeleteDevicePreviewShareLinkResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DownloadDevicePreviewEmailSubmissionSourceEmail

> *os.File DownloadDevicePreviewEmailSubmissionSourceEmail(ctx, submissionId)

Download the original RFC822 email for a device preview submission

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**submissionId** | [**string**]()|  | 

### Return type

[***os.File**](*os.File)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: message/rfc822

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DownloadDevicePreviewRunSourceEmail

> *os.File DownloadDevicePreviewRunSourceEmail(ctx, runId)

Download the original RFC822 email for a device preview run

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | [**string**]()|  | 

### Return type

[***os.File**](*os.File)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: message/rfc822

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## EnsureDevicePreviewRun

> CreateDevicePreviewRunResult EnsureDevicePreviewRun(ctx, emailId, optional)

Return active run for email or create one when none exists

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()|  | 
 **optional** | ***EnsureDevicePreviewRunOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EnsureDevicePreviewRunOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDevicePreviewOptions** | [**optional.Interface of CreateDevicePreviewOptions**](CreateDevicePreviewOptions)|  | 

### Return type

[**CreateDevicePreviewRunResult**](CreateDevicePreviewRunResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDevicePreviewEmailDomainOrCreate

> DevicePreviewAccountSettingsDto GetDevicePreviewEmailDomainOrCreate(ctx, )

Return account device preview email domain setup or create one

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**DevicePreviewAccountSettingsDto**](DevicePreviewAccountSettingsDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDevicePreviewEmailRuns

> []DevicePreviewRunDto GetDevicePreviewEmailRuns(ctx, since, optional)

List device preview runs created from account email domains

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**since** | **time.Time**|  | 
 **optional** | ***GetDevicePreviewEmailRunsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDevicePreviewEmailRunsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **emailDomainId** | [**optional.Interface of string**]()|  | 
 **profileId** | [**optional.Interface of string**]()|  | 
 **localPart** | **optional.String**|  | 
 **limit** | **optional.Int32**|  | [default to 20]
 **hydrateThumbnail** | **optional.Bool**| Hydrate each returned run with a nullable primary screenshot thumbnail URL. | [default to false]

### Return type

[**[]DevicePreviewRunDto**](DevicePreviewRunDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDevicePreviewEmailSubmissionSourceEmail

> DevicePreviewSourceEmailDto GetDevicePreviewEmailSubmissionSourceEmail(ctx, submissionId)

Get the original email content for a device preview submission

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**submissionId** | [**string**]()|  | 

### Return type

[**DevicePreviewSourceEmailDto**](DevicePreviewSourceEmailDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDevicePreviewEmailSubmissions

> []DevicePreviewEmailSubmissionDto GetDevicePreviewEmailSubmissions(ctx, since, optional)

List received device preview email submissions

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**since** | **time.Time**|  | 
 **optional** | ***GetDevicePreviewEmailSubmissionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDevicePreviewEmailSubmissionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **emailDomainId** | [**optional.Interface of string**]()|  | 
 **profileId** | [**optional.Interface of string**]()|  | 
 **localPart** | **optional.String**|  | 
 **limit** | **optional.Int32**|  | [default to 20]
 **recipientEmailAddress** | **optional.String**|  | 

### Return type

[**[]DevicePreviewEmailSubmissionDto**](DevicePreviewEmailSubmissionDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDevicePreviewFeedback

> DevicePreviewFeedbackDto GetDevicePreviewFeedback(ctx, feedbackId)

Get a single device preview feedback item

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**feedbackId** | [**string**]()|  | 

### Return type

[**DevicePreviewFeedbackDto**](DevicePreviewFeedbackDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDevicePreviewFeedbackItems

> DevicePreviewFeedbackListDto GetDevicePreviewFeedbackItems(ctx, optional)

List device preview feedback

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetDevicePreviewFeedbackItemsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDevicePreviewFeedbackItemsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**|  | 
 **size** | **optional.Int32**|  | 
 **source** | **optional.String**|  | 
 **runId** | [**optional.Interface of string**]()|  | 
 **status** | **optional.String**|  | 
 **provider** | **optional.String**|  | 
 **category** | **optional.String**|  | 
 **search** | **optional.String**|  | 

### Return type

[**DevicePreviewFeedbackListDto**](DevicePreviewFeedbackListDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDevicePreviewFolder

> DevicePreviewFolderDto GetDevicePreviewFolder(ctx, folderId)

Get a device preview folder

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | [**string**]()|  | 

### Return type

[**DevicePreviewFolderDto**](DevicePreviewFolderDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDevicePreviewFolders

> []DevicePreviewFolderDto GetDevicePreviewFolders(ctx, )

List device preview folders

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]DevicePreviewFolderDto**](DevicePreviewFolderDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDevicePreviewImportAddress

> DevicePreviewIngestAddressDto GetDevicePreviewImportAddress(ctx, sessionId)

Get temporary device preview import address status

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | [**string**]()|  | 

### Return type

[**DevicePreviewIngestAddressDto**](DevicePreviewIngestAddressDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDevicePreviewNativeTargetAvailability

> DevicePreviewNativeTargetAvailabilityListDto GetDevicePreviewNativeTargetAvailability(ctx, )

Get native device preview target availability

Returns live native device preview target availability for authenticated API clients, including worker health, queue, timeout, and status fields. Use `/device-preview-targets` when you need the stable catalog without transient availability data.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**DevicePreviewNativeTargetAvailabilityListDto**](DevicePreviewNativeTargetAvailabilityListDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDevicePreviewProfile

> DevicePreviewProfileDto GetDevicePreviewProfile(ctx, profileId)

Get a device preview profile

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | [**string**]()|  | 

### Return type

[**DevicePreviewProfileDto**](DevicePreviewProfileDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDevicePreviewProfiles

> []DevicePreviewProfileDto GetDevicePreviewProfiles(ctx, )

List device preview profiles

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]DevicePreviewProfileDto**](DevicePreviewProfileDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDevicePreviewRun

> DevicePreviewRunDto GetDevicePreviewRun(ctx, runId)

Get device preview run status

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | [**string**]()|  | 

### Return type

[**DevicePreviewRunDto**](DevicePreviewRunDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDevicePreviewRunProviderProgress

> DevicePreviewProviderProgressDto GetDevicePreviewRunProviderProgress(ctx, runId, provider)

Get provider-level progress for a device preview run

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | [**string**]()|  | 
**provider** | **string**|  | 

### Return type

[**DevicePreviewProviderProgressDto**](DevicePreviewProviderProgressDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDevicePreviewRunResults

> DevicePreviewRunResultsDto GetDevicePreviewRunResults(ctx, runId)

Get device preview run results

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | [**string**]()|  | 

### Return type

[**DevicePreviewRunResultsDto**](DevicePreviewRunResultsDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDevicePreviewRunScreenshot

> string GetDevicePreviewRunScreenshot(ctx, runId, screenshotId)

Get a device preview screenshot image

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | [**string**]()|  | 
**screenshotId** | [**string**]()|  | 

### Return type

**string**

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDevicePreviewRunSourceEmail

> DevicePreviewSourceEmailDto GetDevicePreviewRunSourceEmail(ctx, runId)

Get the original email content for a device preview run

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | [**string**]()|  | 

### Return type

[**DevicePreviewSourceEmailDto**](DevicePreviewSourceEmailDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDevicePreviewRuns

> []DevicePreviewRunDto GetDevicePreviewRuns(ctx, emailId, optional)

List previous device preview runs for an email

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()|  | 
 **optional** | ***GetDevicePreviewRunsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDevicePreviewRunsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 50]
 **hydrateThumbnail** | **optional.Bool**| Hydrate each returned run with a nullable primary screenshot thumbnail URL. | [default to false]

### Return type

[**[]DevicePreviewRunDto**](DevicePreviewRunDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDevicePreviewRunsForAccount

> []DevicePreviewRunDto GetDevicePreviewRunsForAccount(ctx, optional)

List previous device preview runs for account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetDevicePreviewRunsForAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDevicePreviewRunsForAccountOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**|  | [default to 50]
 **hydrateThumbnail** | **optional.Bool**| Hydrate each returned run with a nullable primary screenshot thumbnail URL. | [default to false]

### Return type

[**[]DevicePreviewRunDto**](DevicePreviewRunDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDevicePreviewRunsForAccountOffsetPaginated

> PageDevicePreviewRunProjection GetDevicePreviewRunsForAccountOffsetPaginated(ctx, optional)

List previous device preview runs for account in paginated form

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetDevicePreviewRunsForAccountOffsetPaginatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDevicePreviewRunsForAccountOffsetPaginatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Optional page index in list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size for paginated result list. | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to DESC]
 **hydrateThumbnail** | **optional.Bool**| Hydrate each returned run with a nullable primary screenshot thumbnail URL. | [default to false]
 **folderId** | [**optional.Interface of string**]()|  | 
 **unfiledOnly** | **optional.Bool**|  | [default to false]
 **emailId** | [**optional.Interface of string**]()|  | 

### Return type

[**PageDevicePreviewRunProjection**](PageDevicePreviewRunProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDevicePreviewRunsOffsetPaginated

> PageDevicePreviewRunProjection GetDevicePreviewRunsOffsetPaginated(ctx, emailId, optional)

List previous device preview runs for an email in paginated form

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailId** | [**string**]()|  | 
 **optional** | ***GetDevicePreviewRunsOffsetPaginatedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDevicePreviewRunsOffsetPaginatedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Optional page index in list pagination | [default to 0]
 **size** | **optional.Int32**| Optional page size for paginated result list. | [default to 20]
 **sort** | **optional.String**| Optional createdAt sort direction ASC or DESC | [default to DESC]
 **hydrateThumbnail** | **optional.Bool**| Hydrate each returned run with a nullable primary screenshot thumbnail URL. | [default to false]

### Return type

[**PageDevicePreviewRunProjection**](PageDevicePreviewRunProjection)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDevicePreviewShareLinks

> []DevicePreviewShareLinkDto GetDevicePreviewShareLinks(ctx, runId)

List share links for a device preview run

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | [**string**]()|  | 

### Return type

[**[]DevicePreviewShareLinkDto**](DevicePreviewShareLinkDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDevicePreviewSharedResult

> DevicePreviewSharedResultDto GetDevicePreviewSharedResult(ctx, shareToken)

Get a public shared device preview result

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shareToken** | **string**|  | 

### Return type

[**DevicePreviewSharedResultDto**](DevicePreviewSharedResultDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDevicePreviewSharedResultScreenshot

> string GetDevicePreviewSharedResultScreenshot(ctx, shareToken, screenshotId)

Get a public shared device preview screenshot image

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shareToken** | **string**|  | 
**screenshotId** | [**string**]()|  | 

### Return type

**string**

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetDevicePreviewTargets

> DevicePreviewNativeTargetListDto GetDevicePreviewTargets(ctx, )

List device preview targets

Returns visible native device preview target metadata without worker availability, queue, or health fields.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**DevicePreviewNativeTargetListDto**](DevicePreviewNativeTargetListDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## MoveDevicePreviewRuns

> MoveDevicePreviewRunsResult MoveDevicePreviewRuns(ctx, moveDevicePreviewRunsOptions)

Move device preview runs to a folder or Unfiled

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moveDevicePreviewRunsOptions** | [**MoveDevicePreviewRunsOptions**](MoveDevicePreviewRunsOptions)|  | 

### Return type

[**MoveDevicePreviewRunsResult**](MoveDevicePreviewRunsResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## RerunDevicePreviewTargets

> RerunDevicePreviewTargetsResult RerunDevicePreviewTargets(ctx, runId, rerunDevicePreviewTargetsOptions)

Rerun selected device preview targets in the same run

Set failedOnly=false to rerun a successful target. Successful targets consume their configured credit cost again.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | [**string**]()|  | 
**rerunDevicePreviewTargetsOptions** | [**RerunDevicePreviewTargetsOptions**](RerunDevicePreviewTargetsOptions)|  | 

### Return type

[**RerunDevicePreviewTargetsResult**](RerunDevicePreviewTargetsResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## UpdateDevicePreviewEmailDomainSubdomain

> DevicePreviewAccountSettingsDto UpdateDevicePreviewEmailDomainSubdomain(ctx, updateDevicePreviewIngestDomainOptions)

Update account device preview email subdomain

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**updateDevicePreviewIngestDomainOptions** | [**UpdateDevicePreviewIngestDomainOptions**](UpdateDevicePreviewIngestDomainOptions)|  | 

### Return type

[**DevicePreviewAccountSettingsDto**](DevicePreviewAccountSettingsDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## UpdateDevicePreviewFeedback

> DevicePreviewFeedbackDto UpdateDevicePreviewFeedback(ctx, feedbackId, updateDevicePreviewFeedbackOptions)

Update device preview feedback

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**feedbackId** | [**string**]()|  | 
**updateDevicePreviewFeedbackOptions** | [**UpdateDevicePreviewFeedbackOptions**](UpdateDevicePreviewFeedbackOptions)|  | 

### Return type

[**DevicePreviewFeedbackDto**](DevicePreviewFeedbackDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## UpdateDevicePreviewFolder

> DevicePreviewFolderDto UpdateDevicePreviewFolder(ctx, folderId, updateDevicePreviewFolderOptions)

Rename or move a device preview folder

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | [**string**]()|  | 
**updateDevicePreviewFolderOptions** | [**UpdateDevicePreviewFolderOptions**](UpdateDevicePreviewFolderOptions)|  | 

### Return type

[**DevicePreviewFolderDto**](DevicePreviewFolderDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## UpdateDevicePreviewProfile

> DevicePreviewProfileDto UpdateDevicePreviewProfile(ctx, profileId, updateDevicePreviewProfileOptions)

Update a device preview profile

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | [**string**]()|  | 
**updateDevicePreviewProfileOptions** | [**UpdateDevicePreviewProfileOptions**](UpdateDevicePreviewProfileOptions)|  | 

### Return type

[**DevicePreviewProfileDto**](DevicePreviewProfileDto)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## WaitForDevicePreviewEmailRuns

> DevicePreviewIngestRunWaitResult WaitForDevicePreviewEmailRuns(ctx, since, optional)

Wait for device preview runs created from account email domains

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**since** | **time.Time**|  | 
 **optional** | ***WaitForDevicePreviewEmailRunsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WaitForDevicePreviewEmailRunsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **emailDomainId** | [**optional.Interface of string**]()|  | 
 **profileId** | [**optional.Interface of string**]()|  | 
 **localPart** | **optional.String**|  | 
 **limit** | **optional.Int32**|  | [default to 20]
 **timeoutMillis** | **optional.Int64**|  | [default to 25000]
 **pollIntervalMillis** | **optional.Int64**|  | [default to 1000]
 **hydrateThumbnail** | **optional.Bool**| Hydrate each returned run with a nullable primary screenshot thumbnail URL. | [default to false]

### Return type

[**DevicePreviewIngestRunWaitResult**](DevicePreviewIngestRunWaitResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## WaitForDevicePreviewEmailSubmissions

> DevicePreviewEmailSubmissionWaitResult WaitForDevicePreviewEmailSubmissions(ctx, since, optional)

Wait for received device preview email submissions

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**since** | **time.Time**|  | 
 **optional** | ***WaitForDevicePreviewEmailSubmissionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WaitForDevicePreviewEmailSubmissionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **emailDomainId** | [**optional.Interface of string**]()|  | 
 **profileId** | [**optional.Interface of string**]()|  | 
 **localPart** | **optional.String**|  | 
 **limit** | **optional.Int32**|  | [default to 20]
 **timeoutMillis** | **optional.Int64**|  | [default to 25000]
 **pollIntervalMillis** | **optional.Int64**|  | [default to 1000]
 **recipientEmailAddress** | **optional.String**|  | 

### Return type

[**DevicePreviewEmailSubmissionWaitResult**](DevicePreviewEmailSubmissionWaitResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## WaitForDevicePreviewRun

> DevicePreviewRunWaitResult WaitForDevicePreviewRun(ctx, runId, optional)

Wait for device preview run to complete

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | [**string**]()|  | 
 **optional** | ***WaitForDevicePreviewRunOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WaitForDevicePreviewRunOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timeoutMillis** | **optional.Int64**|  | [default to 25000]
 **pollIntervalMillis** | **optional.Int64**|  | [default to 1000]

### Return type

[**DevicePreviewRunWaitResult**](DevicePreviewRunWaitResult)

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

