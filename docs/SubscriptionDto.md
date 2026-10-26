# SubscriptionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**UserId** | **string** |  | 
**CustomerId** | Pointer to **string** |  | [optional] 
**SubscriptionId** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Cost** | Pointer to **int32** |  | [optional] 
**IsDecimalLess** | **bool** |  | 
**SubscriptionType** | **string** |  | 
**SubscriptionSchedule** | **string** |  | 
**MonthlyInboxes** | **int32** |  | 
**MonthlyReceive** | **int32** |  | 
**OauthConnections** | **int32** |  | 
**StorageBytes** | **int64** |  | 
**TotalInboxes** | **int32** |  | 
**PhoneNumbers** | **int32** |  | 
**PhoneProvisionRequestLimit** | **int32** |  | 
**MonthlySend** | **int32** |  | 
**Domains** | **int32** |  | 
**Members** | **int32** |  | 
**Aliases** | **int32** |  | 
**Threads** | **int32** |  | 
**Recipients** | **int32** |  | 
**Attachments** | **bool** |  | 
**Bulk** | **bool** |  | 
**PermanentInbox** | **bool** |  | 
**Contacts** | **bool** |  | 
**Webhooks** | **bool** |  | 
**Password** | **bool** |  | 
**InboxRulesets** | **bool** |  | 
**Zapier** | **bool** |  | 
**TrackingPixels** | **bool** |  | 
**InboxForwarders** | **bool** |  | 
**InboxRepliers** | **bool** |  | 
**AttachmentAutomations** | **bool** |  | 
**CreatedAt** | [**time.Time**](time.Time) |  | 
**UpdatedAt** | [**time.Time**](time.Time) |  | 
**MonthlyGetMailServerRecords** | **int32** |  | 
**MonthlyVerifyEmailAddress** | **int32** |  | 
**Metered** | **bool** |  | 
**MaxUniqueRecipients** | **int32** |  | 
**ExternalPlanId** | Pointer to **string** |  | [optional] 
**ExternalValidationPlanId** | Pointer to **string** | Use validation entity | [optional] 
**ExternalValidationSubscriptionItemId** | Pointer to **string** | Use validation entity | [optional] 
**SendWithQueue** | Pointer to **bool** |  | [optional] 
**TrialDays** | Pointer to **int32** |  | [optional] 
**CancelAt** | Pointer to [**time.Time**](time.Time) |  | [optional] 
**TrialEnd** | Pointer to [**time.Time**](time.Time) |  | [optional] 
**TrialStart** | Pointer to [**time.Time**](time.Time) |  | [optional] 
**TrialAdjustmentReason** | Pointer to **string** |  | [optional] 
**TrialAdjustedAt** | Pointer to [**time.Time**](time.Time) |  | [optional] 
**SamlSSO** | Pointer to **bool** |  | [optional] 
**OidcSSO** | Pointer to **bool** |  | [optional] 
**ScimProvisioning** | Pointer to **bool** |  | [optional] 
**DemoDaysRemaining** | Pointer to **int32** |  | [optional] 
**Organizations** | Pointer to **int32** |  | [optional] 
**EmailFeatureCheck** | Pointer to **bool** |  | [optional] 
**Backup** | Pointer to **bool** |  | [optional] 
**Screenshots** | Pointer to **bool** |  | [optional] 
**DevicePreviews** | Pointer to **bool** |  | [optional] 
**UserPermissions** | Pointer to **bool** |  | [optional] 
**Environments** | Pointer to **bool** |  | [optional] 
**ImportExport** | Pointer to **bool** |  | [optional] 
**AuditLogs** | Pointer to **bool** |  | [optional] 
**DeliverabilityTest** | Pointer to **bool** |  | [optional] 
**EmailAudit** | Pointer to **bool** |  | [optional] 
**InboxWarmUp** | Pointer to **bool** |  | [optional] 
**WarmupPremiumSeedAccess** | Pointer to **bool** |  | [optional] 
**DomainMonitor** | Pointer to **bool** |  | [optional] 
**CampaignProbe** | Pointer to **bool** |  | [optional] 
**VirtualTotpDevices** | Pointer to **bool** |  | [optional] 
**CurrentPeriodEnd** | [**time.Time**](time.Time) |  | 
**CurrentPeriodStart** | [**time.Time**](time.Time) |  | 
**SubscriptionStatus** | Pointer to **string** |  | [optional] 
**SubscriptionStatusIsBad** | **bool** |  | 
**DailyReceive** | **int32** |  | 
**DailySend** | **int32** |  | 
**MaxEmailSizeMb** | **int32** |  | 
**AiTransformers** | **bool** |  | 
**GmailInboxConnection** | **bool** |  | 
**AiDestinationSheets** | **bool** |  | 
**GmailInboxConnectionLimit** | **int32** |  | 
**PermitsExternalSend** | Pointer to **bool** |  | [optional] 
**TotalPermanentInboxes** | Pointer to **int32** |  | [optional] 
**MonthlyTemporaryInboxes** | Pointer to **int32** |  | [optional] 
**Agents** | Pointer to **int32** | Maximum active API or MCP agent connections. Null identifies a legacy unlimited subscription. | [optional] 
**MaxInboxPlacementSeedAddressesPerTest** | **int32** | Maximum seed inbox addresses assigned to one inbox placement test. | 
**MaxAutomatedInboxPlacementSchedules** | **int32** | Maximum non-deleted automated inbox placement schedules. Paused schedules count; deleted schedules release capacity. | 
**IncludedWarmupSenderSeats** | **int32** | Warmup sender seats included by the base subscription before paid add-ons. | 
**IncludedMonthlyWarmupCreditsPerSender** | **int32** |  | 
**WarmupControlledRecipientCreditWeight** | **int32** |  | 
**WarmupPremiumRecipientCreditWeight** | **int32** |  | 
**LoadTest** | Pointer to **bool** |  | [optional] 
**CostMajor** | Pointer to **float32** | Subscription cost in major currency units for analytics and conversion value reporting. | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


