# EmailPreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attachments** | **[]string** | List of IDs of attachments found in the email. Use these IDs with the Inbox and Email Controllers to download attachments and attachment meta data such as filesize, name, extension. | [optional] 
**Bcc** | **[]string** | List of &#x60;BCC&#x60; recipients email was addressed to | [optional] 
**Cc** | **[]string** | List of &#x60;CC&#x60; recipients email was addressed to | [optional] 
**CreatedAt** | [**time.Time**](time.Time) | When was the email received by MailSlurp | [optional] 
**From** | **string** | Who the email was sent from | [optional] 
**Id** | **string** | ID of the email entity | [optional] 
**Read** | **bool** | Read flag. Has the email ever been viewed in the dashboard or fetched via the API? If so the email is marked as read. | [optional] 
**Subject** | **string** | The subject line of the email message | [optional] 
**To** | **[]string** | List of &#x60;To&#x60; recipients that email was addressed to | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


