# InboxReplierDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**InboxId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Field** | Pointer to **string** |  | [optional] 
**Match** | Pointer to **string** |  | [optional] 
**ReplyTo** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**From** | Pointer to **string** |  | [optional] 
**Charset** | Pointer to **string** |  | [optional] 
**IsHTML** | **bool** |  | 
**TemplateId** | Pointer to **string** |  | [optional] 
**TemplateVariables** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**IgnoreReplyTo** | **bool** |  | 
**CreatedAt** | [**time.Time**](time.Time) |  | 
**Should** | Pointer to **string** | Comparison mode for inbox automation matching. | [optional] 
**MatchOptions** | Pointer to [**InboxAutomationMatchOptions**](InboxAutomationMatchOptions) |  | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


