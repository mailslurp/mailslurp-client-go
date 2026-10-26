# ExtractCodesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Found** | **bool** | True if at least one code candidate was found | 
**Code** | Pointer to **string** | Best candidate code when found | [optional] 
**MethodUsed** | Pointer to **string** | Extraction strategy for verification codes. Unsupported strategies may fall back when allowFallback is true. | [optional] 
**Candidates** | [**[]CodeCandidate**](CodeCandidate) | Ranked code candidates | 
**Warnings** | **[]string** | Warnings or fallback notes explaining extraction behavior for debugging and QA. | 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


