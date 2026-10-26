# EmailIntelligenceTestsOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckRandomLocalPart** | Pointer to **bool** | Check local-part randomness for email inputs. | [optional] 
**CheckFreeProvider** | Pointer to **bool** | Check if domain is a known free email provider. | [optional] 
**CheckHttpsWebsite** | Pointer to **bool** | Check if the domain has a reachable HTTPS website. | [optional] 
**CheckDns** | Pointer to **bool** | Check DNS records (A, MX, SOA) for the domain. | [optional] 
**CheckDomainAgeHint** | Pointer to **bool** | Derive a domain age hint from DNS SOA serial when possible. | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


