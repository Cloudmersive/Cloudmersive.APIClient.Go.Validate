# XxeDetectionRequestItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputText** | **string** | Individual input text item to protect from XXE | [optional] [default to null]
**AllowInternetUrls** | **bool** | Optional: Set to true to allow Internet-based dependency URLs for DTDs and other XML External Entitites, set to false to block.  Default is false. | [optional] [default to null]
**KnownSafeUrls** | **[]string** | Optional: Comma separated list of fully-qualified URLs that will automatically be considered safe. | [optional] [default to null]
**KnownUnsafeUrls** | **[]string** | Optional: Comma separated list of fully-qualified URLs that will automatically be considered unsafe. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


