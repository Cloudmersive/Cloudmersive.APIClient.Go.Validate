# AddressVerifySyntaxOnlyResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValidAddress** | **bool** | True if the email address is syntactically valid, false if it is not | [optional] [default to null]
**Domain** | **string** | Domain name of the email address | [optional] [default to null]
**IsFreeEmailProvider** | **bool** | True if the email domain name is a free provider (typically a free to sign up web email provider for consumers / personal use), false otherwise. | [optional] [default to null]
**IsDisposable** | **bool** | True if the email address is a disposable email address, false otherwise; these disposable providers are not typically used to receive email and so will have a low likelihood of opening mail sent there. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


