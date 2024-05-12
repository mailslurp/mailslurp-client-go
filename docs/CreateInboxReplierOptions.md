# CreateInboxReplierOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InboxId** | **string** | Inbox ID to attach replier to | 
**Name** | Pointer to **string** | Name for replier | [optional] 
**Field** | **string** | Field to match against to trigger inbox replier for inbound email | 
**Match** | **string** | String or wildcard style match for field specified when evaluating reply rules. Use &#x60;*&#x60; to match anything. | 
**ReplyTo** | Pointer to **string** | Reply-to email address when sending replying | [optional] 
**Subject** | Pointer to **string** | Subject override when replying to email | [optional] 
**From** | Pointer to **string** | Send email from address | [optional] 
**Charset** | Pointer to **string** | Email reply charset | [optional] 
**IgnoreReplyTo** | Pointer to **bool** | Ignore sender replyTo when responding. Send directly to the sender if enabled. | [optional] 
**IsHTML** | Pointer to **bool** | Send HTML email | [optional] 
**Body** | Pointer to **string** | Email body for reply | [optional] 
**TemplateId** | Pointer to **string** | ID of template to use when sending a reply | [optional] 
**TemplateVariables** | Pointer to **map[string]map[string]interface{}** | Template variable values | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


