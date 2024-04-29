# DomainDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**UserId** | **string** |  | 
**Domain** | **string** | Custom domain name | 
**VerificationToken** | **string** | Verification tokens | 
**DkimTokens** | **[]string** | Unique token DKIM tokens | 
**MissingRecordsMessage** | Pointer to **string** | If the domain is missing records then show which pairs are missing. | [optional] 
**HasMissingRecords** | **bool** | Whether the domain has missing required records. If true then see the domain in the dashboard app. | 
**IsVerified** | **bool** | Whether domain has been verified or not. If the domain is not verified after 72 hours there is most likely an issue with the domains DNS records. | 
**DomainNameRecords** | [**[]DomainNameRecord**](DomainNameRecord) | List of DNS domain name records (C, MX, TXT) etc that you must add to the DNS server associated with your domain provider. | 
**CatchAllInboxId** | Pointer to **string** | The optional catch all inbox that will receive emails sent to the domain that cannot be matched. | [optional] 
**CreatedAt** | [**time.Time**](time.Time) |  | 
**UpdatedAt** | [**time.Time**](time.Time) |  | 
**DomainType** | **string** | Type of domain. Dictates type of inbox that can be created with domain. HTTP means inboxes are processed using SES while SMTP inboxes use a custom SMTP mail server. SMTP does not support sending so use HTTP for sending emails. | 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


