# StructuredOutputSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnyOf** | Pointer to [**[]StructuredOutputSchema**](StructuredOutputSchema) |  | [optional] 
**Default** | Pointer to [**map[string]interface{}**]() |  | [optional] 
**Description** | Pointer to **string** | Provide a description of the schema to help the AI understand the schema. | [optional] 
**EnumValues** | Pointer to **[]string** | When using type string and format enum pass a collection of enum values here. | [optional] 
**Example** | Pointer to [**map[string]interface{}**]() |  | [optional] 
**Format** | Pointer to **string** | Format for string types. Can be null, date-time or enum. | [optional] 
**Items** | Pointer to [**StructuredOutputSchema**](StructuredOutputSchema) |  | [optional] 
**MaxItems** | Pointer to **int64** |  | [optional] 
**MinItems** | Pointer to **int64** |  | [optional] 
**MaxLength** | Pointer to **int64** |  | [optional] 
**MinLength** | Pointer to **int64** |  | [optional] 
**Pattern** | Pointer to **string** | Regex pattern for STRING type | [optional] 
**Properties** | Pointer to [**map[string]StructuredOutputSchema**](StructuredOutputSchema) | Properties of an OBJECT schema. These are key value pairs where the key is the property name and the value is the schema for that property. | [optional] 
**PropertyOrdering** | Pointer to **[]string** | Pass an array of property names to specify the order of properties in the generated JSON object if required. | [optional] 
**Required** | Pointer to **[]string** | Is field required | [optional] 
**MaxProperties** | Pointer to **int64** |  | [optional] 
**MinProperties** | Pointer to **int64** |  | [optional] 
**Maximum** | Pointer to **float32** |  | [optional] 
**Minimum** | Pointer to **float32** |  | [optional] 
**Nullable** | Pointer to **bool** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | Primitive JSON schema types with a fallback CUSTOM for unknown values. | [optional] 

[[Back to Model list]](../README#documentation-for-models) [[Back to API list]](../README#documentation-for-api-endpoints) [[Back to README]](../README)


