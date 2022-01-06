# DomainDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] 
**UserId** | **string** |  | [optional] 
**Domain** | **string** | Custom domain name | [optional] 
**VerificationToken** | **string** | Verification tokens | [optional] 
**DkimTokens** | **[]string** | Unique token DKIM tokens | [optional] 
**DomainNameRecords** | [**[]DomainNameRecord**](DomainNameRecord) | List of DNS domain name records (C, MX, TXT) etc that you must add to the DNS server associated with your domain provider. | [optional] 
**CatchAllInboxId** | **string** | The optional catch all inbox that will receive emails sent to the domain that cannot be matched. | [optional] 
**CreatedAt** | [**time.Time**](time.Time) |  | [optional] 
**UpdatedAt** | [**time.Time**](time.Time) |  | [optional] 
**DomainType** | **string** | Type of domain. Dictates type of inbox that can be created with domain. HTTP means inboxes are processed using SES while SMTP inboxes use a custom SMTP mail server. SMTP does not support sending so use HTTP for sending emails. | [optional] 
**Verified** | **bool** |  | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


