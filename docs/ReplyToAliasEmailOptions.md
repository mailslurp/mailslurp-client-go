# ReplyToAliasEmailOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attachments** | **[]string** | List of uploaded attachments to send with the reply. Optional. | [optional] 
**Body** | **string** | Body of the reply email you want to send | [optional] 
**Charset** | **string** | The charset that your message should be sent with. Optional. Default is UTF-8 | [optional] 
**IsHTML** | **bool** | Is the reply HTML | [optional] 
**SendStrategy** | **string** | When to send the email. Typically immediately | [optional] 
**Template** | **string** | Template ID to use instead of body. Will use template variable map to fill defined variable slots. | [optional] 
**TemplateVariables** | [**map[string]interface{}**](.md) | Template variables if using a template | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


