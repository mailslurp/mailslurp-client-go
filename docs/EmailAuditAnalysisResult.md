# EmailAuditAnalysisResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentReview** | Pointer to [**[]EmailAuditCheck**](EmailAuditCheck) | Versioned static accessibility, unresolved template, and compliance review checks. Null on historical audits that did not run these checks. | [optional] 
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
**FeatureSupport** | Pointer to [**EmailFeatureSupportResult**](EmailFeatureSupportResult) |  | [optional] 
**HtmlErrorCount** | **int32** |  | 
**HtmlWarningCount** | **int32** |  | 
**HtmlInfoCount** | **int32** |  | 
**HtmlValidation** | Pointer to [**HtmlValidationResult**](HTMLValidationResult) |  | [optional] 
**ReputationFailureCount** | **int32** |  | 
**AttachmentMentionIssueCount** | **int32** |  | 
**ExternalCheckSkippedCount** | **int32** |  | 
**Insights** | **[]string** |  | 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**SpamVerdict** | Pointer to **string** | Spam verdict supplied by the source mail receiver | [optional] 
**VirusVerdict** | Pointer to **string** | Virus verdict supplied by the source mail receiver | [optional] 
**SpfVerdict** | Pointer to **string** | SPF verdict supplied by the source mail receiver | [optional] 
**DkimVerdict** | Pointer to **string** | DKIM verdict supplied by the source mail receiver | [optional] 
**DmarcVerdict** | Pointer to **string** | DMARC verdict supplied by the source mail receiver | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


