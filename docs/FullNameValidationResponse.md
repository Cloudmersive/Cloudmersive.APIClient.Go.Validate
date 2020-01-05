# FullNameValidationResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Successful** | **bool** | True if the validation operation was successful, false otherwise | [optional] [default to null]
**ValidationResultFirstName** | **string** | Possible values are: ValidFirstName, ValidUnknownFirstName, InvalidSpamInput, InvalidCharacters, InvalidEmpty | [optional] [default to null]
**ValidationResultLastName** | **string** | Possible values are: ValidLastName, ValidUnknownLastName, InvalidSpamInput, InvalidCharacters, InvalidEmpty | [optional] [default to null]
**Title** | **string** | The person&#39;s title (if supplied), e.g. \&quot;Mr.\&quot; or \&quot;Ms.\&quot; | [optional] [default to null]
**FirstName** | **string** | The first name (given name) | [optional] [default to null]
**MiddleName** | **string** | The middle name(s); if there are multiple names they will be separated by spaces | [optional] [default to null]
**LastName** | **string** | The last name (surname) | [optional] [default to null]
**NickName** | **string** | Nickname (if supplied) | [optional] [default to null]
**Suffix** | **string** | Suffix to the name, e.g. \&quot;Jr.\&quot; or \&quot;Sr.\&quot; | [optional] [default to null]
**DisplayName** | **string** | The full display name of the name | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


