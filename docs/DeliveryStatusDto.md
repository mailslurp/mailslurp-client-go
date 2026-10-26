# DeliveryStatusDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**UserId** | **string** |  | 
**SentId** | Pointer to **string** |  | [optional] 
**RemoteMtaIp** | Pointer to **string** |  | [optional] 
**InboxId** | Pointer to **string** |  | [optional] 
**ReportingMta** | Pointer to **string** |  | [optional] 
**Recipients** | Pointer to **[]string** |  | [optional] 
**SmtpResponse** | Pointer to **string** |  | [optional] 
**SmtpStatusCode** | Pointer to **int32** |  | [optional] 
**ProcessingTimeMillis** | Pointer to **int64** |  | [optional] 
**Received** | Pointer to [**time.Time**](time.Time) |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**CreatedAt** | [**time.Time**](time.Time) |  | 
**UpdatedAt** | [**time.Time**](time.Time) |  | 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


