# DevicePreviewRunDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunId** | **string** |  | 
**EmailId** | Pointer to **string** |  | [optional] 
**SourceType** | **string** |  | 
**SourceInboxId** | Pointer to **string** |  | [optional] 
**SourceAlias** | Pointer to **string** |  | [optional] 
**SourceExternalId** | Pointer to **string** |  | [optional] 
**SourceFrom** | Pointer to **string** |  | [optional] 
**SourceTo** | Pointer to **[]string** |  | [optional] 
**SourceSubject** | Pointer to **string** |  | [optional] 
**SourceMessageId** | Pointer to **string** |  | [optional] 
**SourceSizeBytes** | Pointer to **int64** |  | [optional] 
**SourceStorageProvider** | Pointer to **string** |  | [optional] 
**SourceStorageBucket** | Pointer to **string** |  | [optional] 
**SourceStorageObjectKey** | Pointer to **string** |  | [optional] 
**SourceStorageRegion** | Pointer to **string** |  | [optional] 
**SourceStorageEndpoint** | Pointer to **string** |  | [optional] 
**SourceStoragePublicBaseUrl** | Pointer to **string** |  | [optional] 
**SourceRecipientEmailAddress** | Pointer to **string** |  | [optional] 
**SourceRecipientLocalPart** | Pointer to **string** |  | [optional] 
**SourceIngestDomainId** | Pointer to **string** |  | [optional] 
**SourceProfileId** | Pointer to **string** |  | [optional] 
**SourceProfileLocalPart** | Pointer to **string** |  | [optional] 
**SourceProfileDisplayName** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 
**PrimaryScreenshotId** | Pointer to **string** |  | [optional] 
**RequestedProviders** | Pointer to **[]string** |  | [optional] 
**ImportedProviders** | Pointer to **[]string** |  | [optional] 
**Warnings** | Pointer to **[]string** |  | [optional] 
**ProviderMessageIds** | Pointer to **map[string]string** |  | [optional] 
**LatestError** | Pointer to **string** | Latest run-level error, if the preview run was blocked or failed after creation. | [optional] 
**TargetCount** | **int64** |  | 
**ScreenshotCount** | **int64** |  | 
**CreatedAt** | [**time.Time**](time.Time) |  | 
**UpdatedAt** | [**time.Time**](time.Time) |  | 
**Thumbnail** | Pointer to [**DevicePreviewRunThumbnailDto**](DevicePreviewRunThumbnailDto) |  | [optional] 
**QaBundle** | Pointer to [**DevicePreviewQaBundleDto**](DevicePreviewQaBundleDto) |  | [optional] 
**FolderId** | Pointer to **string** |  | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


