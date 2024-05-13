# WaitForSmsConditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumberId** | **string** | ID of phone number to search within and apply conditions to. Essentially filtering the SMS found to give a count. | 
**Limit** | Pointer to **int32** | Limit results | [optional] 
**Count** | **int64** | Number of results that should match conditions. Either exactly or at least this amount based on the &#x60;countType&#x60;. If count condition is not met and the timeout has not been reached the &#x60;waitFor&#x60; method will retry the operation. | 
**DelayTimeout** | Pointer to **int64** | Max time in milliseconds to wait between retries if a &#x60;timeout&#x60; is specified. | [optional] 
**Timeout** | **int64** | Max time in milliseconds to retry the &#x60;waitFor&#x60; operation until conditions are met. | 
**UnreadOnly** | Pointer to **bool** | Apply conditions only to **unread** SMS. All SMS messages begin with &#x60;read&#x3D;false&#x60;. An SMS is marked &#x60;read&#x3D;true&#x60; when an &#x60;SMS&#x60; has been returned to the user at least once. For example you have called &#x60;getSms&#x60; or &#x60;waitForSms&#x60; etc., or you have viewed the SMS in the dashboard. | [optional] 
**CountType** | Pointer to **string** | How result size should be compared with the expected size. Exactly or at-least matching result? | [optional] 
**Matches** | Pointer to [**[]SmsMatchOption**](SmsMatchOption) | Conditions that should be matched for an SMS to qualify for results. Each condition will be applied in order to each SMS within a phone number to filter a result list of matching SMSs you are waiting for. | [optional] 
**SortDirection** | Pointer to **string** | Direction to sort matching SMSs by created time | [optional] 
**Since** | Pointer to [**time.Time**](time.Time) | ISO Date Time earliest time of SMS to consider. Filter for matching SMSs that were received after this date | [optional] 
**Before** | Pointer to [**time.Time**](time.Time) | ISO Date Time latest time of SMS to consider. Filter for matching SMSs that were received before this date | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


