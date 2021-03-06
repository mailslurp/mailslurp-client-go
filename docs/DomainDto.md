# DomainDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CatchAllInboxId** | **string** | The optional catch all inbox that will receive emails sent to the domain that cannot be matched. | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) |  | 
**DkimTokens** | **[]string** | Unique token DKIM tokens | [optional] 
**Domain** | **string** | Custom domain name | [optional] 
**DomainNameRecords** | [**[]DomainNameRecord**](DomainNameRecord.md) | List of DNS domain name records (C, MX, TXT) etc that you must add to the DNS server associated with your domain provider. | [optional] 
**Id** | **string** |  | 
**IsVerified** | **bool** | Whether domain has been verified or not. If the domain is not verified after 72 hours there is most likely an issue with the domains DNS records. | [optional] 
**UpdatedAt** | [**time.Time**](time.Time.md) |  | 
**UserId** | **string** |  | 
**VerificationToken** | **string** | Verification tokens | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


