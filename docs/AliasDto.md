# AliasDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**EmailAddress** | **string** | The alias&#39;s email address for receiving email | 
**MaskedEmailAddress** | Pointer to **string** | The underlying email address that is hidden and will received forwarded email | [optional] 
**UserId** | **string** |  | 
**InboxId** | **string** | Inbox that is associated with the alias | 
**Name** | Pointer to **string** |  | [optional] 
**UseThreads** | Pointer to **bool** | If alias will generate response threads or not when email are received by it | [optional] 
**IsVerified** | **bool** | Has the alias been verified. You must verify an alias if the masked email address has not yet been verified by your account | 
**CreatedAt** | Pointer to [**time.Time**](time.Time) |  | [optional] 
**UpdatedAt** | Pointer to [**time.Time**](time.Time) |  | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


