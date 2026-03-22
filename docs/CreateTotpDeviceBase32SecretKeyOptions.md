# CreateTotpDeviceBase32SecretKeyOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Base32SecretKey** | **string** | Base32 secret key for TOTP device as alternative to OTP Auth URI or QR code. | 
**Name** | Pointer to **string** | Name for labeling the TOTP device | [optional] 
**Username** | Pointer to **string** | Optional username for the TOTP device | [optional] 
**Issuer** | Pointer to **string** | Optional issuer override for the TOTP device | [optional] 
**Digits** | Pointer to **int32** | Optional number of digits in TOTP code | [optional] [default to 6]
**Period** | Pointer to **int32** | Optional period in seconds for TOTP code expiration | [optional] [default to 30]
**Algorithm** | Pointer to **string** | Optional algorithm override | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


