# AttachmentEntityDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier for this attachment. | 
**AttachmentId** | **string** | The identifier of the attachment file | 
**UserId** | **string** | The user identifier associated with this attachment. | 
**ContentType** | Pointer to **string** | The content type of the attachment. | [optional] 
**ContentLength** | Pointer to **int64** | The content length of the attachment in bytes. | [optional] 
**ContentId** | Pointer to **string** | The content identifier, which is a unique ID for the content part of the email. | [optional] 
**Name** | Pointer to **string** | The name of the attachment file. | [optional] 
**InboxId** | Pointer to **string** | The inbox identifier associated with this attachment. | [optional] 
**CreatedAt** | [**time.Time**](time.Time) | The timestamp when this attachment was created. | 
**UpdatedAt** | [**time.Time**](time.Time) | The timestamp when this attachment was last updated. | 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


