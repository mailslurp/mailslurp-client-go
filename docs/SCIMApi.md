# MailSlurp\SCIMApi

All URIs are relative to *https://golang.api.mailslurp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGroup**](SCIMApi#CreateGroup) | **Post** /scim/v2/{slug}/Groups | 
[**CreateUser**](SCIMApi#CreateUser) | **Post** /scim/v2/{slug}/Users | 
[**DeleteGroup**](SCIMApi#DeleteGroup) | **Delete** /scim/v2/{slug}/Groups/{groupId} | 
[**DeleteUser**](SCIMApi#DeleteUser) | **Delete** /scim/v2/{slug}/Users/{userId} | 
[**GetGroup**](SCIMApi#GetGroup) | **Get** /scim/v2/{slug}/Groups/{groupId} | 
[**GetResourceType**](SCIMApi#GetResourceType) | **Get** /scim/v2/{slug}/ResourceTypes/{resourceType} | 
[**GetResourceTypes**](SCIMApi#GetResourceTypes) | **Get** /scim/v2/{slug}/ResourceTypes | 
[**GetSchema**](SCIMApi#GetSchema) | **Get** /scim/v2/{slug}/Schemas/{schemaId} | 
[**GetSchemas**](SCIMApi#GetSchemas) | **Get** /scim/v2/{slug}/Schemas | 
[**GetServiceProviderConfig**](SCIMApi#GetServiceProviderConfig) | **Get** /scim/v2/{slug}/ServiceProviderConfig | 
[**GetUser**](SCIMApi#GetUser) | **Get** /scim/v2/{slug}/Users/{userId} | 
[**ListGroups**](SCIMApi#ListGroups) | **Get** /scim/v2/{slug}/Groups | 
[**ListUsers**](SCIMApi#ListUsers) | **Get** /scim/v2/{slug}/Users | 
[**PatchGroup**](SCIMApi#PatchGroup) | **Patch** /scim/v2/{slug}/Groups/{groupId} | 
[**PatchUser**](SCIMApi#PatchUser) | **Patch** /scim/v2/{slug}/Users/{userId} | 
[**ReplaceGroup**](SCIMApi#ReplaceGroup) | **Put** /scim/v2/{slug}/Groups/{groupId} | 
[**ReplaceUser**](SCIMApi#ReplaceUser) | **Put** /scim/v2/{slug}/Users/{userId} | 



## CreateGroup

> map[string]interface{} CreateGroup(ctx, slug, scimGroupResource)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string**|  | 
**scimGroupResource** | [**ScimGroupResource**](ScimGroupResource)|  | 

### Return type

**map[string]interface{}**

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/scim+json, application/json
- **Accept**: application/scim+json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## CreateUser

> map[string]interface{} CreateUser(ctx, slug, scimUserResource)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string**|  | 
**scimUserResource** | [**ScimUserResource**](ScimUserResource)|  | 

### Return type

**map[string]interface{}**

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/scim+json, application/json
- **Accept**: application/scim+json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DeleteGroup

> map[string]interface{} DeleteGroup(ctx, slug, groupId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string**|  | 
**groupId** | **string**|  | 

### Return type

**map[string]interface{}**

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/scim+json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## DeleteUser

> map[string]interface{} DeleteUser(ctx, slug, userId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string**|  | 
**userId** | **string**|  | 

### Return type

**map[string]interface{}**

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/scim+json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetGroup

> map[string]interface{} GetGroup(ctx, slug, groupId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string**|  | 
**groupId** | **string**|  | 

### Return type

**map[string]interface{}**

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/scim+json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetResourceType

> map[string]interface{} GetResourceType(ctx, slug, resourceType)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string**|  | 
**resourceType** | **string**|  | 

### Return type

**map[string]interface{}**

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/scim+json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetResourceTypes

> map[string]interface{} GetResourceTypes(ctx, slug)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string**|  | 

### Return type

**map[string]interface{}**

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/scim+json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetSchema

> map[string]interface{} GetSchema(ctx, slug, schemaId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string**|  | 
**schemaId** | **string**|  | 

### Return type

**map[string]interface{}**

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/scim+json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetSchemas

> map[string]interface{} GetSchemas(ctx, slug)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string**|  | 

### Return type

**map[string]interface{}**

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/scim+json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetServiceProviderConfig

> map[string]interface{} GetServiceProviderConfig(ctx, slug)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string**|  | 

### Return type

**map[string]interface{}**

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/scim+json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## GetUser

> map[string]interface{} GetUser(ctx, slug, userId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string**|  | 
**userId** | **string**|  | 

### Return type

**map[string]interface{}**

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/scim+json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## ListGroups

> map[string]interface{} ListGroups(ctx, slug, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string**|  | 
 **optional** | ***ListGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListGroupsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**|  | 
 **startIndex** | **optional.Int32**|  | 
 **count** | **optional.Int32**|  | 

### Return type

**map[string]interface{}**

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/scim+json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## ListUsers

> map[string]interface{} ListUsers(ctx, slug, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string**|  | 
 **optional** | ***ListUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListUsersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**|  | 
 **startIndex** | **optional.Int32**|  | 
 **count** | **optional.Int32**|  | 

### Return type

**map[string]interface{}**

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/scim+json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## PatchGroup

> map[string]interface{} PatchGroup(ctx, slug, groupId, scimPatchRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string**|  | 
**groupId** | **string**|  | 
**scimPatchRequest** | [**ScimPatchRequest**](ScimPatchRequest)|  | 

### Return type

**map[string]interface{}**

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/scim+json, application/json
- **Accept**: application/scim+json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## PatchUser

> map[string]interface{} PatchUser(ctx, slug, userId, scimPatchRequest)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string**|  | 
**userId** | **string**|  | 
**scimPatchRequest** | [**ScimPatchRequest**](ScimPatchRequest)|  | 

### Return type

**map[string]interface{}**

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/scim+json, application/json
- **Accept**: application/scim+json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## ReplaceGroup

> map[string]interface{} ReplaceGroup(ctx, slug, groupId, scimGroupResource)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string**|  | 
**groupId** | **string**|  | 
**scimGroupResource** | [**ScimGroupResource**](ScimGroupResource)|  | 

### Return type

**map[string]interface{}**

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/scim+json, application/json
- **Accept**: application/scim+json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)


## ReplaceUser

> map[string]interface{} ReplaceUser(ctx, slug, userId, scimUserResource)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string**|  | 
**userId** | **string**|  | 
**scimUserResource** | [**ScimUserResource**](ScimUserResource)|  | 

### Return type

**map[string]interface{}**

### Authorization

[API_KEY](../README#API_KEY)

### HTTP request headers

- **Content-Type**: application/scim+json, application/json
- **Accept**: application/scim+json

[[Back to top]](#) [[Back to API list]](../README#documentation-for-api-endpoints)
[[Back to Model list]](../README#documentation-for-models)
[[Back to README]](../README)

