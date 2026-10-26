# SearchEmailsOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InboxIds** | Pointer to **[]string** | Optional inbox ids to filter by. Can be repeated. By default will use all inboxes belonging to your account. | [optional] 
**PageIndex** | Pointer to **int32** | Optional page index in email list pagination | [optional] 
**PageSize** | Pointer to **int32** | Optional page size in email list pagination. Maximum size is 100. Use page index and sort to page through larger results | [optional] 
**SortDirection** | Pointer to **string** | Optional createdAt sort direction ASC or DESC | [optional] 
**UnreadOnly** | Pointer to **bool** | Optional filter for unread emails only. All emails are considered unread until they are viewed in the dashboard or requested directly | [optional] 
**SearchFilter** | Pointer to **string** | Optional search filter. Matches email ID, inbox ID, sender address, subject, receiving inbox email address, receiving plus-address full address, or plus-address suffix. Pasted email address tokens are normalized before matching. | [optional] 
**Since** | Pointer to [**time.Time**](time.Time) | Optional filter emails received after given date time | [optional] 
**Before** | Pointer to [**time.Time**](time.Time) | Optional filter emails received before given date time | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


