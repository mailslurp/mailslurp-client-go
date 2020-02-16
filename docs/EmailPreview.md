# EmailPreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attachments** | **[]string** | List of IDs of attachments found in the email. Use these IDs with the Inbox and Email Controllers to download attachments and attachment meta data such as filesize, name, extension. | [optional] 
**Bcc** | **[]string** | List of &#x60;BCC&#x60; recipients email was addressed to | [optional] 
**Cc** | **[]string** | List of &#x60;CC&#x60; recipients email was addressed to | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) | When was the email received by MailSlurp | [optional] 
**Id** | **string** | ID of the email | [optional] 
**Read** | **bool** | Has the email been viewed ever | [optional] 
**Subject** | **string** | The subject line of the email message | [optional] 
**To** | **[]string** | List of &#x60;To&#x60; recipients email was addressed to | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


