# CreateInboxForwarderOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | Pointer to **string** | Field to match against to trigger inbox forwarding for inbound email | [optional] 
**Match** | Pointer to **string** | String or wildcard style match for field specified when evaluating forwarding rules | [optional] 
**ForwardToRecipients** | **[]string** | Email addresses to forward an email to if it matches the field and match criteria of the forwarder | 
**Should** | Pointer to **string** | Comparison mode for inbox automation matching. | [optional] 
**MatchOptions** | Pointer to [**InboxAutomationMatchOptions**](InboxAutomationMatchOptions) |  | [optional] 
**AttachmentTextExtractionMethod** | Pointer to **string** | Method for extracting text from attachments. | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


