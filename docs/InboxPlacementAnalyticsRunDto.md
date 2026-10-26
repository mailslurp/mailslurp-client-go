# InboxPlacementAnalyticsRunDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**RemoteTestId** | **string** |  | 
**SenderDomain** | Pointer to **string** |  | [optional] 
**FromEmail** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ReportStatus** | Pointer to **string** |  | [optional] 
**CreatedAt** | [**time.Time**](time.Time) |  | 
**StartedAt** | Pointer to [**time.Time**](time.Time) |  | [optional] 
**CompletedAt** | Pointer to [**time.Time**](time.Time) |  | [optional] 
**TotalTargets** | **int32** |  | 
**MatchedTargets** | **int32** |  | 
**DeliveredTargets** | **int32** |  | 
**InboxTargets** | **int32** |  | 
**SpamTargets** | **int32** |  | 
**PromotionTargets** | **int32** |  | 
**SocialTargets** | **int32** |  | 
**UpdateTargets** | **int32** |  | 
**ForumTargets** | **int32** |  | 
**OtherFolderTargets** | **int32** |  | 
**NotReceivedTargets** | **int32** |  | 
**DeliveredRatePercentage** | **float64** |  | 
**InboxRatePercentage** | **float64** |  | 
**SpamRatePercentage** | **float64** |  | 
**PromotionRatePercentage** | **float64** |  | 
**NotReceivedRatePercentage** | **float64** |  | 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


