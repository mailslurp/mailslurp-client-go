# ExtractAttachmentTextResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Present** | **bool** | True if extracted text is present | 
**Text** | Pointer to **string** | Extracted text when present | [optional] 
**MethodUsed** | Pointer to **string** | Method for extracting text from attachments. OCR/LLM methods are reserved for advanced extraction. | [optional] 
**ContentType** | Pointer to **string** | Detected attachment content type | [optional] 
**Warnings** | **[]string** | Warnings or fallback notes | 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


