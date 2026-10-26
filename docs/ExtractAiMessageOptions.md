# ExtractAiMessageOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | [**AiMessageReference**](AIMessageReference) |  | 
**OutputSchema** | Pointer to [**StructuredOutputSchema**](StructuredOutputSchema) |  | [optional] 
**TransformId** | Pointer to **string** |  | [optional] 
**Instructions** | Pointer to **string** |  | [optional] 
**Timeout** | **int64** |  | 
**AiOptions** | [**AiMessageAiOptions**](AIMessageAIOptions) |  | 
**ExtractionPreset** | Pointer to **string** | Built-in extraction schema. OTP_CODE returns data.code as a verbatim string, preserving leading zeros and letter case. Missing or equally eligible codes are INCONCLUSIVE. | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


