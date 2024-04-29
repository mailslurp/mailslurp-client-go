# ReplyToEmailOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | **string** | Body of the reply email you want to send | 
**IsHTML** | **bool** | Is the reply HTML | 
**From** | Pointer to **string** | The from header that should be used. Optional | [optional] 
**ReplyTo** | Pointer to **string** | The replyTo header that should be used. Optional | [optional] 
**CustomHeaders** | Pointer to **map[string]string** | Optional custom headers | [optional] 
**Charset** | Pointer to **string** | The charset that your message should be sent with. Optional. Default is UTF-8 | [optional] 
**Attachments** | Pointer to **[]string** | List of uploaded attachments to send with the reply. Optional. | [optional] 
**TemplateVariables** | Pointer to **map[string]map[string]interface{}** | Template variables if using a template | [optional] 
**Template** | Pointer to **string** | Template ID to use instead of body. Will use template variable map to fill defined variable slots. | [optional] 
**SendStrategy** | Pointer to **string** | How an email should be sent based on its recipients | [optional] 
**UseInboxName** | Pointer to **bool** | Optionally use inbox name as display name for sender email address | [optional] 
**Html** | **bool** |  | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


