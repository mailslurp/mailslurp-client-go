# CreateDevicePreviewOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Targets** | Pointer to **[]string** | Optional device preview target keys to request for rendering. Null leaves target selection to service configuration. Values are forwarded to ESP so newly deployed targets can be used before this API&#39;s enum list is updated. | [optional] 
**AllTargets** | Pointer to **bool** | When true, ask the renderer to include every currently configured target. Cannot be combined with targets. | [optional] 
**Qa** | Pointer to [**CreateDevicePreviewQaOptions**](CreateDevicePreviewQaOptions) |  | [optional] 
**FolderId** | Pointer to **string** | Optional folder for organizing the render. Missing folders leave the render unfiled. | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


