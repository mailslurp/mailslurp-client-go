# EmailIntelligenceOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Targets** | **[]string** | Email addresses or domains to score. | 
**Page** | Pointer to **int32** | Zero-based page index for processing a subset of the target list. | [optional] 
**Size** | Pointer to **int32** | Page size for processing a subset of the target list. | [optional] 
**IgnoreCache** | Pointer to **bool** | Ignore cached intelligence values and force recomputation. | [optional] 
**IncludeEmailValidation** | Pointer to **bool** | Also run mailbox safety verification using the existing verification client for email inputs. | [optional] 
**Tests** | Pointer to [**EmailIntelligenceTestsOptions**](EmailIntelligenceTestsOptions) |  | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


