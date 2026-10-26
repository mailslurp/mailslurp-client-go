# AgentConversationQueueItemDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConversationId** | **string** |  | 
**ThreadId** | **string** |  | 
**InboxId** | **string** |  | 
**Subject** | Pointer to **string** |  | [optional] 
**Customer** | Pointer to **string** |  | [optional] 
**LastSender** | Pointer to **string** |  | [optional] 
**LastExcerpt** | Pointer to **string** |  | [optional] 
**LatestMessageDirection** | **string** |  | 
**LatestEmailId** | Pointer to **string** |  | [optional] 
**LatestReceivedEmailId** | Pointer to **string** |  | [optional] 
**LatestSentEmailId** | Pointer to **string** |  | [optional] 
**MessageCount** | **int32** |  | 
**UnreadCount** | **int32** |  | 
**HasAttachments** | **bool** |  | 
**CreatedAt** | [**time.Time**](time.Time) |  | 
**LastMessageAt** | [**time.Time**](time.Time) |  | 
**ProjectedAt** | [**time.Time**](time.Time) |  | 
**Version** | **int64** |  | 
**WorkflowState** | **string** |  | 
**Outcome** | Pointer to **string** |  | [optional] 
**ReviewStatus** | **string** |  | 
**RequiresHumanReview** | **bool** |  | 
**NeedsOperationalAttention** | **bool** |  | 
**Claim** | Pointer to [**AgentConversationClaimDto**](AgentConversationClaimDto) |  | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


