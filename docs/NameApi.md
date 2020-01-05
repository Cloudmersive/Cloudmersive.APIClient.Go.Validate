# \NameApi

All URIs are relative to *https://api.cloudmersive.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NameGetGender**](NameApi.md#NameGetGender) | **Post** /validate/name/get-gender | Get the gender of a first name
[**NameIdentifier**](NameApi.md#NameIdentifier) | **Post** /validate/name/identifier | Validate a code identifier
[**NameValidateFirstName**](NameApi.md#NameValidateFirstName) | **Post** /validate/name/first | Validate a first name
[**NameValidateFullName**](NameApi.md#NameValidateFullName) | **Post** /validate/name/full-name | Parse and validate a full name
[**NameValidateLastName**](NameApi.md#NameValidateLastName) | **Post** /validate/name/last | Validate a last name


# **NameGetGender**
> GetGenderResponse NameGetGender(ctx, input)
Get the gender of a first name

Determines the gender of a first name (given name)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **input** | [**GetGenderRequest**](GetGenderRequest.md)| Gender request information | 

### Return type

[**GetGenderResponse**](GetGenderResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NameIdentifier**
> ValidateIdentifierResponse NameIdentifier(ctx, input)
Validate a code identifier

Determines if the input name is a valid technical / code identifier.  Configure input rules such as whether whitespace, hyphens, underscores, etc. are allowed.  For example, a valid identifier might be \"helloWorld\" but not \"hello*World\".

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **input** | [**ValidateIdentifierRequest**](ValidateIdentifierRequest.md)| Identifier validation request information | 

### Return type

[**ValidateIdentifierResponse**](ValidateIdentifierResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NameValidateFirstName**
> FirstNameValidationResponse NameValidateFirstName(ctx, input)
Validate a first name

Determines if a string is a valid first name (given name)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **input** | [**FirstNameValidationRequest**](FirstNameValidationRequest.md)| Validation request information | 

### Return type

[**FirstNameValidationResponse**](FirstNameValidationResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NameValidateFullName**
> FullNameValidationResponse NameValidateFullName(ctx, input)
Parse and validate a full name

Parses a full name string (e.g. \"Mr. Jon van der Waal Jr.\") into its component parts (and returns these component parts), and then validates whether it is a valid name string or not

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **input** | [**FullNameValidationRequest**](FullNameValidationRequest.md)| Validation request information | 

### Return type

[**FullNameValidationResponse**](FullNameValidationResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NameValidateLastName**
> LastNameValidationResponse NameValidateLastName(ctx, input)
Validate a last name

Determines if a string is a valid last name (surname)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **input** | [**LastNameValidationRequest**](LastNameValidationRequest.md)| Validation request information | 

### Return type

[**LastNameValidationResponse**](LastNameValidationResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

