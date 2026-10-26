# PublicSpamTestRunDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Source** | **string** |  | 
**Status** | Pointer to **string** |  | [optional] 
**ReportStatus** | Pointer to **string** |  | [optional] 
**RequestedSegment** | Pointer to **string** |  | [optional] 
**SeedlistProfile** | Pointer to **string** |  | [optional] 
**SenderDomain** | Pointer to **string** |  | [optional] 
**FromEmail** | Pointer to **string** |  | [optional] 
**TrackingToken** | Pointer to **string** |  | [optional] 
**SeedAddresses** | **[]string** |  | 
**SeedAddressFormats** | **map[string]string** |  | 
**TotalTargets** | **int32** |  | 
**MatchedTargets** | **int32** |  | 
**InboxTargets** | **int32** |  | 
**SpamTargets** | **int32** |  | 
**NotReceivedTargets** | **int32** |  | 
**SubmittedAt** | Pointer to [**time.Time**](time.Time) |  | [optional] 
**StartedAt** | Pointer to [**time.Time**](time.Time) |  | [optional] 
**FirstMatchAt** | Pointer to [**time.Time**](time.Time) |  | [optional] 
**CompletedAt** | Pointer to [**time.Time**](time.Time) |  | [optional] 
**ExpiresAt** | Pointer to [**time.Time**](time.Time) |  | [optional] 
**LastEventAt** | Pointer to [**time.Time**](time.Time) |  | [optional] 
**LastSyncedAt** | Pointer to [**time.Time**](time.Time) |  | [optional] 
**FreeLimitReached** | **bool** |  | 
**ErrorCode** | Pointer to **string** |  | [optional] 
**CreatedAt** | [**time.Time**](time.Time) |  | 
**UpdatedAt** | [**time.Time**](time.Time) |  | 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


