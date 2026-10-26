# UserDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ApiKey** | **string** | Deprecate api key use entity instead | 
**EmailAddress** | **string** |  | 
**EmailAddressMd5** | **string** |  | 
**Created** | Pointer to [**time.Time**](time.Time) |  | [optional] 
**UserType** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to **string** | Does user belong to an organization | [optional] 
**Verified** | Pointer to **string** | Has user accepted an organization invite | [optional] 
**HasPassword** | **bool** |  | 
**IsFrozen** | **bool** |  | 
**FrozenMessage** | Pointer to **string** |  | [optional] 
**AddNewContacts** | Pointer to **bool** |  | [optional] 
**SsoProvider** | Pointer to **string** |  | [optional] 
**CustomerId** | Pointer to **string** |  | [optional] 
**HasOnboarded** | Pointer to **bool** |  | [optional] 
**ImapUsername** | Pointer to **string** |  | [optional] 
**ImapPassword** | Pointer to **string** |  | [optional] 
**SmtpUsername** | Pointer to **string** |  | [optional] 
**SmtpPassword** | Pointer to **string** |  | [optional] 
**ChargePending** | Pointer to **bool** |  | [optional] 
**EnableSms** | Pointer to **bool** |  | [optional] 
**VendorEmailAddress** | Pointer to **string** |  | [optional] 
**Unsubscribed** | Pointer to **bool** |  | [optional] 
**ApiAccessToken** | Pointer to [**TemporaryAccessKeyDto**](TemporaryAccessKeyDto) |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**VerifiedPhone** | Pointer to **string** | Has user accepted an organization invite | [optional] 
**AccountRegion** | Pointer to **string** | Default account region | [optional] 
**Name** | Pointer to **string** | User display name | [optional] 
**AccessLinkSendResult** | Pointer to [**AccessLinkSendResultDto**](AccessLinkSendResultDto) |  | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


