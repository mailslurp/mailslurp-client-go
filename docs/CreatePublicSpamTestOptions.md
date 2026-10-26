# CreatePublicSpamTestOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromEmail** | Pointer to **string** | Optional sender email address. If provided, MailSlurp will also derive senderDomain for compatibility with inbox placement service versions that still expect it. | [optional] 
**SenderDomain** | Pointer to **string** | Optional sender domain. If omitted and fromEmail is provided, MailSlurp derives it automatically. | [optional] 
**RequestedSegment** | **string** | Requested seed cohort for the placement test. | 
**AddressFormat** | **string** | Address export format for the returned seed addresses. | 
**CaptchaToken** | Pointer to **string** | Optional captcha token when captcha protection is enabled. | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


