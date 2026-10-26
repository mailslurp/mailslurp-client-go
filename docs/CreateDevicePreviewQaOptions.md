# CreateDevicePreviewQaOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAudit** | Pointer to **bool** | Run and persist an email content audit alongside the device preview run. | [optional] 
**DomainMonitor** | Pointer to **bool** | Create or reuse a manual domain monitor for the sender domain and run it alongside the device preview run. | [optional] 
**SenderDomain** | Pointer to **string** | Sender domain to use for domain monitor QA when the source email does not include a real From domain, such as HTML imports. | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


