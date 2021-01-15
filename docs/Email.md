# Email

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Analysis** | [**EmailAnalysis**](EmailAnalysis.md) |  | [optional] 
**Attachments** | **[]string** | List of IDs of attachments found in the email. Use these IDs with the Inbox and Email Controllers to download attachments and attachment meta data such as filesize, name, extension. | [optional] 
**Bcc** | **[]string** | List of &#x60;BCC&#x60; recipients email was addressed to | [optional] 
**Body** | **string** | The body of the email message | [optional] 
**BodyMD5Hash** | **string** | A hash signature of the email message | [optional] 
**Cc** | **[]string** | List of &#x60;CC&#x60; recipients email was addressed to | [optional] 
**Charset** | **string** | Detected character set of the email body such as UTF-8 | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) | When was the email received by MailSlurp | [optional] 
**From** | **string** | Who the email was sent from | [optional] 
**Headers** | **map[string]string** |  | [optional] 
**Id** | **string** | ID of the email | [optional] 
**InboxId** | **string** | ID of the inbox that received the email | [optional] 
**IsHTML** | **bool** | Was HTML sent in the email body | [optional] 
**Read** | **bool** | Has the email been viewed ever | [optional] 
**ReplyTo** | **string** | The replyTo field on the received email | [optional] 
**Subject** | **string** | The subject line of the email message | [optional] 
**To** | **[]string** | List of &#x60;To&#x60; recipients email was addressed to | [optional] 
**UpdatedAt** | [**time.Time**](time.Time.md) | When was the email last updated | [optional] 
**UserId** | **string** | ID of user that email belongs | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


