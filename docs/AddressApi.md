# \AddressApi

All URIs are relative to *https://api.cloudmersive.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddressCheckEUMembership**](AddressApi.md#AddressCheckEUMembership) | **Post** /validate/address/country/check-eu-membership | Check if a country is a member of the European Union (EU)
[**AddressCountry**](AddressApi.md#AddressCountry) | **Post** /validate/address/country | Validate and normalize country information, return ISO 3166-1 country codes and country name
[**AddressCountryList**](AddressApi.md#AddressCountryList) | **Post** /validate/address/country/list | Get a list of ISO 3166-1 countries
[**AddressGetTimezone**](AddressApi.md#AddressGetTimezone) | **Post** /validate/address/country/get-timezones | Gets IANA/Olsen time zones for a country
[**AddressParseString**](AddressApi.md#AddressParseString) | **Post** /validate/address/parse | Parse an unstructured input text string into an international, formatted address
[**AddressValidateAddress**](AddressApi.md#AddressValidateAddress) | **Post** /validate/address/street-address | Validate a street address
[**AddressValidateCity**](AddressApi.md#AddressValidateCity) | **Post** /validate/address/city | Validate a City and State/Province combination, get location information about it
[**AddressValidatePostalCode**](AddressApi.md#AddressValidatePostalCode) | **Post** /validate/address/postal-code | Validate a postal code, get location information about it
[**AddressValidateState**](AddressApi.md#AddressValidateState) | **Post** /validate/address/state | Validate a state or province, name or abbreviation, get location information about it


# **AddressCheckEUMembership**
> ValidateCountryResponse AddressCheckEUMembership(ctx, input)
Check if a country is a member of the European Union (EU)

Checks if the input country is a member of the European Union or not.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **input** | [**ValidateCountryRequest**](ValidateCountryRequest.md)| Input request | 

### Return type

[**ValidateCountryResponse**](ValidateCountryResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddressCountry**
> ValidateCountryResponse AddressCountry(ctx, input)
Validate and normalize country information, return ISO 3166-1 country codes and country name

Validates and normalizes country information, and returns key information about a country, as well as whether it is a member of the European Union.  Also returns distinct time zones in the country.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **input** | [**ValidateCountryRequest**](ValidateCountryRequest.md)| Input request | 

### Return type

[**ValidateCountryResponse**](ValidateCountryResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddressCountryList**
> CountryListResult AddressCountryList(ctx, )
Get a list of ISO 3166-1 countries

Enumerates the list of ISO 3166-1 countries, including name, country codes, and more.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CountryListResult**](CountryListResult.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddressGetTimezone**
> GetTimezonesResponse AddressGetTimezone(ctx, input)
Gets IANA/Olsen time zones for a country

Gets the IANA/Olsen time zones for a country.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **input** | [**GetTimezonesRequest**](GetTimezonesRequest.md)| Input request | 

### Return type

[**GetTimezonesResponse**](GetTimezonesResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddressParseString**
> ParseAddressResponse AddressParseString(ctx, input)
Parse an unstructured input text string into an international, formatted address

Uses machine learning and Natural Language Processing (NLP) to handle a wide array of cases, including non-standard and unstructured address strings across a wide array of countries and address formatting norms.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **input** | [**ParseAddressRequest**](ParseAddressRequest.md)| Input parse request | 

### Return type

[**ParseAddressResponse**](ParseAddressResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddressValidateAddress**
> ValidateAddressResponse AddressValidateAddress(ctx, input)
Validate a street address

Determines if an input structured street address is valid or invalid.  If the address is valid, also returns the latitude and longitude of the address.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **input** | [**ValidateAddressRequest**](ValidateAddressRequest.md)| Input parse request | 

### Return type

[**ValidateAddressResponse**](ValidateAddressResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddressValidateCity**
> ValidateCityResponse AddressValidateCity(ctx, input)
Validate a City and State/Province combination, get location information about it

Checks if the input city and state name or code is valid, and returns information about it such as normalized City name, State name and more.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **input** | [**ValidateCityRequest**](ValidateCityRequest.md)| Input parse request | 

### Return type

[**ValidateCityResponse**](ValidateCityResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddressValidatePostalCode**
> ValidatePostalCodeResponse AddressValidatePostalCode(ctx, input)
Validate a postal code, get location information about it

Checks if the input postal code is valid, and returns information about it such as City, State and more.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **input** | [**ValidatePostalCodeRequest**](ValidatePostalCodeRequest.md)| Input parse request | 

### Return type

[**ValidatePostalCodeResponse**](ValidatePostalCodeResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddressValidateState**
> ValidateStateResponse AddressValidateState(ctx, input)
Validate a state or province, name or abbreviation, get location information about it

Checks if the input state name or code is valid, and returns information about it such as normalized State name and more.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **input** | [**ValidateStateRequest**](ValidateStateRequest.md)| Input parse request | 

### Return type

[**ValidateStateResponse**](ValidateStateResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

