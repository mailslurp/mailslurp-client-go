# AiMessageResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EvaluationId** | **string** |  | 
**Status** | **string** |  | 
**Summary** | **string** |  | 
**Model** | [**AiMessageModelDto**](AIMessageModelDto) |  | 
**Usage** | [**AiMessageUsage**](AIMessageUsage) |  | 
**ElapsedMs** | **int64** |  | 
**Message** | Pointer to [**AiMessageReference**](AIMessageReference) |  | [optional] 
**Email** | Pointer to [**Email**](Email) |  | [optional] 
**Sms** | Pointer to [**SmsDto**](SmsDto) |  | [optional] 
**Match** | Pointer to [**AiMessageMatchResult**](AIMessageMatchResult) |  | [optional] 
**AssertionStatus** | Pointer to **string** |  | [optional] 
**Assertions** | [**[]AiMessageAssertionResult**](AIMessageAssertionResult) |  | 
**Data** | Pointer to [**map[string]interface{}**]() | Extracted JSON value conforming to the requested schema or extraction preset. Null when extraction is inconclusive or unavailable. | [optional] 
**ExtractionEvidence** | [**map[string][]AiMessageEvidence**](array) |  | 
**Candidates** | [**[]AiMessageCandidateResult**](AIMessageCandidateResult) |  | 
**ExtractionReason** | Pointer to **string** | Explanation when extraction is inconclusive. May contain a model-reported explanation of missing or ambiguous fields, or a server evidence-validation failure. Diagnostic text, not a stable machine-readable status. | [optional] 
**Successful** | **bool** | True only when all requested work succeeded. Assert this in tests and use summary as the assertion failure message. | [readonly] 
**FailedAssertionIds** | **[]string** |  | [readonly] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


