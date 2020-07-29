# ValidatePostalCodeResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValidPostalCode** | **bool** | True if the address is valid, false otherwise | [optional] [default to null]
**City** | **string** | If valid, City corresponding to the input postal code, such as &#39;Walnut Creek&#39; | [optional] [default to null]
**StateOrProvince** | **string** | If valid; State or province corresponding to the input postal code, such as &#39;CA&#39; or &#39;California&#39; | [optional] [default to null]
**Latitude** | **float64** | If the postal code is valid, the degrees latitude of the centroid of the postal code, null otherwise | [optional] [default to null]
**Longitude** | **float64** | If the postal code is valid, the degrees longitude of the centroid of the postal code, null otherwise | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


