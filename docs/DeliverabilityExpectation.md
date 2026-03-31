# DeliverabilityExpectation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Optional label for this expectation | [optional] 
**MinCount** | **int64** | Minimum number of matching messages required for this expectation to pass | 
**From** | Pointer to **string** | Optional sender filter. Matching is case-insensitive contains against inbound sender/from values. | [optional] 
**To** | Pointer to **string** | Optional recipient filter. Matching is case-insensitive contains against recipient/to values. | [optional] 
**Subject** | Pointer to **string** | Optional subject filter for INBOX scope tests. Ignored for PHONE scope tests. | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


