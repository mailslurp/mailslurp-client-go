# InboxForwarderDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**InboxId** | **string** |  | 
**Name** | Pointer to **string** | Name of inbox forwarder | [optional] 
**Field** | **string** | Which field to match against | 
**Match** | **string** | Wild-card type pattern to apply to field | 
**ForwardToRecipients** | **[]string** | Who to send forwarded email to | 
**CreatedAt** | [**time.Time**](time.Time) |  | 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


