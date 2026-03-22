# EmailAuditAnalysisResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Health status for a one-shot email audit | 
**HealthScore** | **int32** |  | 
**TotalChecks** | **int32** |  | 
**PassingChecks** | **int32** |  | 
**FailingChecks** | **int32** |  | 
**DetectedLinks** | **int32** |  | 
**CheckedLinks** | **int32** |  | 
**DetectedImages** | **int32** |  | 
**CheckedImages** | **int32** |  | 
**LinkIssueCount** | **int32** |  | 
**ImageIssueCount** | **int32** |  | 
**SpellingIssueCount** | **int32** |  | 
**BrokenLinks** | [**[]EmailAuditUrlIssue**](EmailAuditUrlIssue) |  | 
**BrokenImages** | [**[]EmailAuditUrlIssue**](EmailAuditUrlIssue) |  | 
**SpellingIssues** | [**[]EmailAuditSpellingIssue**](EmailAuditSpellingIssue) |  | 
**CompatibilityWarningCount** | **int32** |  | 
**CompatibilityNotSupportedCount** | **int32** |  | 
**CompatibilityUnknownCount** | **int32** |  | 
**FeatureSupport** | [**EmailFeatureSupportResult**](EmailFeatureSupportResult) |  | [optional] 
**HtmlErrorCount** | **int32** |  | 
**HtmlWarningCount** | **int32** |  | 
**HtmlInfoCount** | **int32** |  | 
**HtmlValidation** | Pointer to [**HtmlValidationResult**](HTMLValidationResult) |  | [optional] 
**ReputationFailureCount** | **int32** |  | 
**AttachmentMentionIssueCount** | **int32** |  | 
**ExternalCheckSkippedCount** | **int32** |  | 
**Insights** | **[]string** |  | 
**ErrorMessage** | Pointer to **string** |  | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


