# CodeCandidate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Extracted code value | 
**Confidence** | **float64** | Relative confidence score from 0 to 1 | 
**Method** | **string** | Extraction strategy for verification codes. Unsupported strategies may fall back when allowFallback is true. | 
**Source** | **string** | Source fragment used for extraction, for example RAW_TEXT_PART or SMS_BODY | 
**Context** | Pointer to **string** | Nearby text context useful for debugging extraction decisions | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


