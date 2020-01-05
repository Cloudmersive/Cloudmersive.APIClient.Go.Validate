# VatLookupResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | **string** | Two-letter country code | [optional] [default to null]
**VatNumber** | **string** | VAT number | [optional] [default to null]
**IsValid** | **bool** | True if the VAT code is valid, false otherwise | [optional] [default to null]
**BusinessName** | **string** | Name of the business | [optional] [default to null]
**BusinessAddress** | **string** | Business address as a single string | [optional] [default to null]
**BusinessBuilding** | **string** | For the business address, the name of the building, house or structure if applicable, such as \&quot;Cloudmersive Building 2\&quot;.  This will often by null. | [optional] [default to null]
**BusinessStreetNumber** | **string** | For the business address, the street number or house number of the address.  For example, in the address \&quot;1600 Pennsylvania Avenue NW\&quot; the street number would be \&quot;1600\&quot;.  This value will typically be populated for most addresses. | [optional] [default to null]
**BusinessStreet** | **string** | For the business address, the name of the street or road of the address.  For example, in the address \&quot;1600 Pennsylvania Avenue NW\&quot; the street number would be \&quot;Pennsylvania Avenue NW\&quot;. | [optional] [default to null]
**BusinessCity** | **string** | For the business address, the city of the address. | [optional] [default to null]
**BusinessStateOrProvince** | **string** | For the business address, the state or province of the address. | [optional] [default to null]
**BusinessPostalCode** | **string** | For the business address, the postal code or zip code of the address. | [optional] [default to null]
**BusinessCountry** | **string** | For the business address, country of the address, if present in the address.  If not included in the address it will be null. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


