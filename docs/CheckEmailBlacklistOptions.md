# CheckEmailBlacklistOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** | Domain to check against configured domain blocklists. Domain A and MX addresses are not treated as outbound sending IPs unless includeDomainInfrastructure is explicitly enabled. | [optional] 
**IpAddress** | Pointer to **string** | Specific observed or configured outbound IPv4 address to check directly | [optional] 
**MxHost** | Pointer to **string** | Specific MX host to resolve and check as inbound infrastructure. MX results do not prove outbound sender reputation. | [optional] 
**IncludeDomainInfrastructure** | Pointer to **bool** | Explicitly include the domain A address and discovered MX addresses as infrastructure diagnostics. Defaults to false because these addresses are normally not outbound sending IPs. | [optional] 
**CaptchaToken** | Pointer to **string** |  | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


