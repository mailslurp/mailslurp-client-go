# AiMessageUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tokens** | **int64** | Consumed MailSlurp AI tokens after applying the model multiplier. These are not raw provider tokens. | 
**CandidatesEvaluated** | **int32** |  | 
**Complete** | **bool** | False when a timed-out provider invocation is still settling. The reservation remains held and actual usage is settled when it finishes. | 
**ReservedTokens** | **int64** | AI tokens held for provider attempts with unknown usage. Excluded from tokens; complete remains false until reconciliation. | 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


