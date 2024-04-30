# CreateContactOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Company** | Pointer to **string** |  | [optional] 
**EmailAddresses** | Pointer to **[]string** | Set of email addresses belonging to the contact | [optional] 
**Tags** | Pointer to **[]string** | Tags that can be used to search and group contacts | [optional] 
**MetaData** | Pointer to [**map[string]interface{}**]() |  | [optional] 
**OptOut** | Pointer to **bool** | Has the user explicitly or implicitly opted out of being contacted? If so MailSlurp will ignore them in all actions. | [optional] 
**GroupId** | Pointer to **string** | Group IDs that contact belongs to | [optional] 
**VerifyEmailAddresses** | Pointer to **bool** | Whether to validate contact email address exists | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


