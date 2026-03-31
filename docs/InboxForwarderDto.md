# InboxForwarderDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**InboxId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | Name of inbox forwarder | [optional] 
**Field** | Pointer to **string** | Which field to match against | [optional] 
**Match** | Pointer to **string** | Pattern to apply to field | [optional] 
**ForwardToRecipients** | **[]string** | Who to send forwarded email to | 
**CreatedAt** | [**time.Time**](time.Time) |  | 
**Should** | Pointer to **string** | Comparison mode for inbox automation matching. | [optional] 
**MatchOptions** | Pointer to [**InboxAutomationMatchOptions**](InboxAutomationMatchOptions) |  | [optional] 
**AttachmentTextExtractionMethod** | Pointer to **string** | Method for extracting text from attachments. | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


