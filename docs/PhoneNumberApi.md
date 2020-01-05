# \PhoneNumberApi

All URIs are relative to *https://api.cloudmersive.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PhoneNumberSyntaxOnly**](PhoneNumberApi.md#PhoneNumberSyntaxOnly) | **Post** /validate/phonenumber/basic | Validate phone number (basic)


# **PhoneNumberSyntaxOnly**
> PhoneNumberValidationResponse PhoneNumberSyntaxOnly(ctx, value)
Validate phone number (basic)

Validate a phone number by analyzing the syntax

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **value** | [**PhoneNumberValidateRequest**](PhoneNumberValidateRequest.md)| Phone number to validate in a PhoneNumberValidateRequest object.  Try a phone number such as \&quot;1.800.463.3339\&quot;, and either leave DefaultCountryCode blank or use \&quot;US\&quot;. | 

### Return type

[**PhoneNumberValidationResponse**](PhoneNumberValidationResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

