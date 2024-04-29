# SearchInboxesOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageIndex** | Pointer to **int32** | Optional page index in list pagination | [optional] 
**PageSize** | Pointer to **int32** | Optional page size in list pagination | [optional] 
**SortDirection** | Pointer to **string** | Optional createdAt sort direction ASC or DESC | [optional] 
**Favourite** | Pointer to **bool** | Optionally filter results for favourites only | [optional] 
**Search** | Pointer to **string** | Optionally filter by search words partial matching ID, tags, name, and email address | [optional] 
**Tag** | Pointer to **string** | Optionally filter by tags. Will return inboxes that include given tags | [optional] 
**Since** | Pointer to [**time.Time**](time.Time) | Optional filter by created after given date time | [optional] 
**Before** | Pointer to [**time.Time**](time.Time) | Optional filter by created before given date time | [optional] 
**InboxType** | Pointer to **string** | Type of inbox. HTTP inboxes are faster and better for most cases. SMTP inboxes are more suited for public facing inbound messages (but cannot send). | [optional] 
**InboxFunction** | Pointer to **string** | Optional filter by inbox function | [optional] 
**DomainId** | Pointer to **string** | Optional domain ID filter | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


