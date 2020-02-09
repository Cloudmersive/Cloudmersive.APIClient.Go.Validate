# ValidateAddressRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StreetAddress** | **string** | Required; Street address to validate, such as &#39;2950 Buskirk Ave.&#39; | [optional] [default to null]
**City** | **string** | Required; City part of the addrerss to validate, such as &#39;Walnut Creek&#39; | [optional] [default to null]
**StateOrProvince** | **string** | Required; State or province of the address to validate, such as &#39;CA&#39; or &#39;California&#39; | [optional] [default to null]
**PostalCode** | **string** | Optional (recommended); Zip code or postal code of the address to validate, such as &#39;94597&#39; | [optional] [default to null]
**CountryFullName** | **string** | Optional (recommended); Name of the country, such as &#39;United States&#39;.  If left blank, and CountryCode is also left blank, will default to United States.  Global countries are supported. | [optional] [default to null]
**CountryCode** | **string** | Optional; two-letter country code (Two-letter ISO 3166-1 country code) of the country.  If left blank, and CountryFullName is also left blank, will default to United States.  Global countries are supported. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


