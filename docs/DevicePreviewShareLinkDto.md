# DevicePreviewShareLinkDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**RunId** | **string** |  | 
**TokenPrefix** | **string** |  | 
**Token** | Pointer to **string** | Raw share token. Only returned immediately after creation. | [optional] 
**ShareUrl** | Pointer to **string** | Public share URL with social preview metadata. Only returned when the raw token is available. | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to [**time.Time**](time.Time) |  | [optional] 
**RevokedAt** | Pointer to [**time.Time**](time.Time) |  | [optional] 
**LastAccessedAt** | Pointer to [**time.Time**](time.Time) |  | [optional] 
**AccessCount** | **int64** |  | 
**CreatedAt** | [**time.Time**](time.Time) |  | 
**UpdatedAt** | [**time.Time**](time.Time) |  | 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


