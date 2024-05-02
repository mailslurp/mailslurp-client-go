# SearchEmailsOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InboxIds** | **[]string** | Optional inbox ids to filter by. Can be repeated. By default will use all inboxes belonging to your account. | [optional] 
**PageIndex** | **int32** | Optional page index in email list pagination | [optional] 
**PageSize** | **int32** | Optional page size in email list pagination. Maximum size is 100. Use page index and sort to page through larger results | [optional] 
**SortDirection** | **string** | Optional createdAt sort direction ASC or DESC | [optional] 
**UnreadOnly** | **bool** | Optional filter for unread emails only. All emails are considered unread until they are viewed in the dashboard or requested directly | [optional] 
**SearchFilter** | **string** | Optional search filter. Searches email recipients, sender, subject, email address and ID. Does not search email body | [optional] 
**Since** | [**time.Time**](time.Time) | Optional filter emails received after given date time | [optional] 
**Before** | [**time.Time**](time.Time) | Optional filter emails received before given date time | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


