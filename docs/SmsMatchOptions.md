# SmsMatchOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Matches** | Pointer to [**[]SmsMatchOption**](SmsMatchOption) | Zero or more match options such as &#x60;{ field: &#39;BODY&#39;, should: &#39;CONTAIN&#39;, value: &#39;Welcome&#39; }&#x60;. Options are additive so if one does not match the SMS is excluded from results | [optional] 
**Ai** | Pointer to [**AiMessageMatchOptions**](AIMessageMatchOptions) |  | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


