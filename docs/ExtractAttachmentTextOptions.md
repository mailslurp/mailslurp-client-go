# ExtractAttachmentTextOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | Pointer to **string** | Method for extracting text from attachments. OCR/LLM methods are reserved for advanced extraction. | [optional] 
**AllowFallback** | **bool** | Allow fallback to native extraction when the selected method is unavailable. | [optional] [default to true]
**MaxBytes** | Pointer to **int64** | Maximum bytes to read from attachment for native text extraction. | [optional] [default to 1048576]

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


