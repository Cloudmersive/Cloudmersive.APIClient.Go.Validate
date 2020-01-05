# PhoneNumberValidationResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsValid** | **bool** | True if the phone number is valid, false otherwise | [optional] [default to null]
**Successful** | **bool** | True if the operation was successful, false if there was an error during validation.  See IsValid for validation result. | [optional] [default to null]
**PhoneNumberType** | **string** | Type of phone number; possible values are: FixedLine, Mobile, FixedLineOrMobile, TollFree, PremiumRate,   SharedCost, Voip, PersonalNumber, Pager, Uan, Voicemail, Unknown | [optional] [default to null]
**E164Format** | **string** | E.164 format of the phone number | [optional] [default to null]
**InternationalFormat** | **string** | Internaltional format of the phone number | [optional] [default to null]
**NationalFormat** | **string** | National format of the phone number | [optional] [default to null]
**CountryCode** | **string** | Two digit country code of the phone number | [optional] [default to null]
**CountryName** | **string** | User-friendly long name of the country for the phone number | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


