# WaitForConditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | Number of results that should match conditions | [optional] 
**CountType** | **string** | Should exactly count number of results be returned or at least that many. | [optional] 
**InboxId** | **string** | Inbox to search within | [optional] 
**Matches** | [**[]MatchOption**](MatchOption.md) | Conditions that should be matched | [optional] 
**SortDirection** | **string** | Direction to sort matching emails by created time | [optional] 
**Timeout** | **int64** | Max time in milliseconds to wait until conditions are met | [optional] 
**UnreadOnly** | **bool** | Apply only to unread emails | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


