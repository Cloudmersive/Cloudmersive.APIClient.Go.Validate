# \DateTimeApi

All URIs are relative to *https://api.cloudmersive.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DateTimeGetNowSimple**](DateTimeApi.md#DateTimeGetNowSimple) | **Get** /validate/date-time/get/now | Get current date and time as of now
[**DateTimeGetPublicHolidays**](DateTimeApi.md#DateTimeGetPublicHolidays) | **Post** /validate/date-time/get/holidays | Get public holidays in the specified country and year
[**DateTimeParseNaturalLanguageDateTime**](DateTimeApi.md#DateTimeParseNaturalLanguageDateTime) | **Post** /validate/date-time/parse/date-time/natural-language | Parses a free-form natural language date and time string into a date and time
[**DateTimeParseStandardDateTime**](DateTimeApi.md#DateTimeParseStandardDateTime) | **Post** /validate/date-time/parse/date-time/structured | Parses a standardized date and time string into a date and time


# **DateTimeGetNowSimple**
> DateTimeNowResult DateTimeGetNowSimple(ctx, )
Get current date and time as of now

Gets the current date and time.  Response time is syncronized with atomic clocks, and represents a monotonic, centrally available, consistent clock.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DateTimeNowResult**](DateTimeNowResult.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DateTimeGetPublicHolidays**
> PublicHolidaysResponse DateTimeGetPublicHolidays(ctx, input)
Get public holidays in the specified country and year

Enumerates all public holidays in a given country for a given year.  Supports over 100 countries.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **input** | [**GetPublicHolidaysRequest**](GetPublicHolidaysRequest.md)| Input request | 

### Return type

[**PublicHolidaysResponse**](PublicHolidaysResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DateTimeParseNaturalLanguageDateTime**
> DateTimeStandardizedParseResponse DateTimeParseNaturalLanguageDateTime(ctx, input)
Parses a free-form natural language date and time string into a date and time

Parses an unstructured, free-form, natural language date and time string into a date time object.  This is intended for lightweight human-entered input, such as \"tomorrow at 3pm\" or \"tuesday\".

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **input** | [**DateTimeNaturalLanguageParseRequest**](DateTimeNaturalLanguageParseRequest.md)| Input request | 

### Return type

[**DateTimeStandardizedParseResponse**](DateTimeStandardizedParseResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DateTimeParseStandardDateTime**
> DateTimeStandardizedParseResponse DateTimeParseStandardDateTime(ctx, input)
Parses a standardized date and time string into a date and time

Parses a structured date and time string into a date time object.  This is intended for standardized date strings that adhere to formatting conventions, rather than natural language input.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **input** | [**DateTimeStandardizedParseRequest**](DateTimeStandardizedParseRequest.md)| Input request | 

### Return type

[**DateTimeStandardizedParseResponse**](DateTimeStandardizedParseResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

