# CreateDeliverabilitySimulationJobOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SenderInboxId** | Pointer to **string** | Sender inbox ID for INBOX scope tests | [optional] 
**SenderPhoneId** | Pointer to **string** | Sender phone ID for PHONE scope tests | [optional] 
**Email** | Pointer to [**DeliverabilitySimulationEmailOptions**](DeliverabilitySimulationEmailOptions) |  | [optional] 
**Sms** | Pointer to [**DeliverabilitySimulationSmsOptions**](DeliverabilitySimulationSmsOptions) |  | [optional] 
**DelayMs** | Pointer to **int64** | Delay between individual sends in milliseconds | [optional] 
**BatchSize** | Pointer to **int32** | Maximum sends processed per scheduler batch | [optional] 
**SendsPerTarget** | Pointer to **int64** | Optional fixed sends per target. If omitted this is derived from test expectations. | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


