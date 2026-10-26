# AccountDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | [**UserDto**](UserDto) |  | 
**Subscription** | Pointer to [**SubscriptionDto**](SubscriptionDto) |  | [optional] 
**Addons** | Pointer to **[]string** |  | [optional] 
**IntercomHash** | Pointer to **string** | disabled | [optional] 
**ChurnkeyHash** | Pointer to **string** | disabled | [optional] 
**CurrentPeriod** | [**CurrentPeriodDto**](CurrentPeriodDto) |  | 
**RequiresVerification** | **bool** |  | 
**RequiredVerificationTypes** | Pointer to **[]string** |  | [optional] 
**DefaultCurrency** | **string** |  | 
**DefaultSchedule** | **string** |  | 
**HasBadSubscriptionItemStatuses** | Pointer to **bool** |  | [optional] 
**TrialsEnabled** | **bool** |  | 
**HasAnyActiveTrialSubscription** | **bool** |  | 
**CanEndTrial** | **bool** |  | 
**FreeAccountExpiresInDays** | Pointer to **int32** |  | [optional] 
**FreeAccountExpired** | **bool** |  | 
**FreeAccountEvaluationDenied** | **bool** |  | 
**FreeAccountEvaluationDeniedReason** | Pointer to **string** |  | [optional] 
**FreeAccountEvaluationDeniedAllowTrial** | **bool** |  | 
**FreeAccountMode** | **string** |  | 
**RequiresTermsAndConditionsAgreement** | Pointer to **bool** |  | [optional] 
**EnforceSendingAgreement** | **bool** |  | 
**AdminSendPause** | Pointer to [**AdminSendPauseDto**](AdminSendPauseDto) |  | [optional] 
**OrganizationName** | Pointer to **string** |  | [optional] 
**OrganizationSlug** | Pointer to **string** |  | [optional] 
**OrganizationBrandColor** | Pointer to **string** |  | [optional] 
**OrganizationLogoUrl** | Pointer to **string** |  | [optional] 
**UserPermissionRole** | Pointer to **string** |  | [optional] 
**EnvironmentAccounts** | Pointer to [**EnvironmentAccountsDto**](EnvironmentAccountsDto) |  | [optional] 
**DomainPlanQuantity** | **int32** |  | 
**AssignedDomainCount** | **int32** |  | 
**EffectiveDomainLimit** | Pointer to **int32** |  | [optional] 
**RequiresDashboardOnboardingV2** | **bool** |  | 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


