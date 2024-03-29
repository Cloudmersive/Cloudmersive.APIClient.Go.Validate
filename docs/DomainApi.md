# \DomainApi

All URIs are relative to *https://api.cloudmersive.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DomainCheck**](DomainApi.md#DomainCheck) | **Post** /validate/domain/check | Validate a domain name
[**DomainGetTopLevelDomainFromUrl**](DomainApi.md#DomainGetTopLevelDomainFromUrl) | **Post** /validate/domain/url/get-top-level-domain | Get top-level domain name from URL
[**DomainIsAdminPath**](DomainApi.md#DomainIsAdminPath) | **Post** /validate/domain/url/is-admin-path | Check if path is a high-risk or vulnerable server administration path
[**DomainPhishingCheck**](DomainApi.md#DomainPhishingCheck) | **Post** /validate/domain/url/phishing-threat-check | Check a URL for Phishing threats
[**DomainPost**](DomainApi.md#DomainPost) | **Post** /validate/domain/whois | Get WHOIS information for a domain
[**DomainQualityScore**](DomainApi.md#DomainQualityScore) | **Post** /validate/domain/quality-score | Validate a domain name&#39;s quality score
[**DomainSafetyCheck**](DomainApi.md#DomainSafetyCheck) | **Post** /validate/domain/url/safety-threat-check | Check a URL for safety threats
[**DomainSsrfCheck**](DomainApi.md#DomainSsrfCheck) | **Post** /validate/domain/url/ssrf-threat-check | Check a URL for SSRF threats
[**DomainSsrfCheckBatch**](DomainApi.md#DomainSsrfCheckBatch) | **Post** /validate/domain/url/ssrf-threat-check/batch | Check a URL for SSRF threats in batches
[**DomainUrlFull**](DomainApi.md#DomainUrlFull) | **Post** /validate/domain/url/full | Validate a URL fully
[**DomainUrlHtmlSsrfCheck**](DomainApi.md#DomainUrlHtmlSsrfCheck) | **Post** /validate/domain/url/ssrf-threat-check/html-embedded | Check a URL for HTML embedded SSRF threats
[**DomainUrlSyntaxOnly**](DomainApi.md#DomainUrlSyntaxOnly) | **Post** /validate/domain/url/syntax-only | Validate a URL syntactically


# **DomainCheck**
> CheckResponse DomainCheck(ctx, domain)
Validate a domain name

Check whether a domain name is valid or not.  API performs a live validation by contacting DNS services to validate the existence of the domain name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domain** | **string**| Domain name to check, for example \&quot;cloudmersive.com\&quot;.  The input is a string so be sure to enclose it in double-quotes. | 

### Return type

[**CheckResponse**](CheckResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomainGetTopLevelDomainFromUrl**
> ValidateUrlResponseSyntaxOnly DomainGetTopLevelDomainFromUrl(ctx, request)
Get top-level domain name from URL

Gets the top-level domain name from a URL, such as mydomain.com.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**ValidateUrlRequestSyntaxOnly**](ValidateUrlRequestSyntaxOnly.md)| Input URL information | 

### Return type

[**ValidateUrlResponseSyntaxOnly**](ValidateUrlResponseSyntaxOnly.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomainIsAdminPath**
> IsAdminPathResponse DomainIsAdminPath(ctx, value)
Check if path is a high-risk or vulnerable server administration path

Check if the input URL or relative path is a server Administration Path, and therefore a risk or vulnerability for remote access.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **value** | **string**| URL or relative path to check, e.g. \&quot;/admin/login\&quot;.  The input is a string so be sure to enclose it in double-quotes. | 

### Return type

[**IsAdminPathResponse**](IsAdminPathResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomainPhishingCheck**
> PhishingCheckResponse DomainPhishingCheck(ctx, request)
Check a URL for Phishing threats

Checks if an input URL is at risk of being an Phishing (fake login page, or other page designed to collect information via social engineering) threat or attack.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**PhishingCheckRequest**](PhishingCheckRequest.md)| Input URL request | 

### Return type

[**PhishingCheckResponse**](PhishingCheckResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomainPost**
> WhoisResponse DomainPost(ctx, domain)
Get WHOIS information for a domain

Validate whether a domain name exists, and also return the full WHOIS record for that domain name.  WHOIS records include all the registration details of the domain name, such as information about the domain's owners.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domain** | **string**| Domain name to check, for example \&quot;cloudmersive.com\&quot;.   The input is a string so be sure to enclose it in double-quotes. | 

### Return type

[**WhoisResponse**](WhoisResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomainQualityScore**
> DomainQualityResponse DomainQualityScore(ctx, domain)
Validate a domain name's quality score

Check the quality of a domain name.  Supports over 9 million domain names.  Higher quality scores indicate more trust and authority in the domain name, with values ranging from 0.0 (low quality) to 10.0 (maximum quality).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domain** | **string**| Domain name to check, for example \&quot;cloudmersive.com\&quot;. | 

### Return type

[**DomainQualityResponse**](DomainQualityResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomainSafetyCheck**
> UrlSafetyCheckResponseFull DomainSafetyCheck(ctx, request)
Check a URL for safety threats

Checks if an input URL is at risk of being a safety threat through malware, unwanted software, or social engineering threats.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**UrlSafetyCheckRequestFull**](UrlSafetyCheckRequestFull.md)| Input URL request | 

### Return type

[**UrlSafetyCheckResponseFull**](UrlSafetyCheckResponseFull.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomainSsrfCheck**
> UrlSsrfResponseFull DomainSsrfCheck(ctx, request)
Check a URL for SSRF threats

Checks if an input URL is at risk of being an SSRF (Server-side request forgery) threat or attack.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**UrlSsrfRequestFull**](UrlSsrfRequestFull.md)| Input URL request | 

### Return type

[**UrlSsrfResponseFull**](UrlSsrfResponseFull.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomainSsrfCheckBatch**
> UrlSsrfResponseBatch DomainSsrfCheckBatch(ctx, request)
Check a URL for SSRF threats in batches

Batch-checks if input URLs are at risk of being an SSRF (Server-side request forgery) threat or attack.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**UrlSsrfRequestBatch**](UrlSsrfRequestBatch.md)| Input URL request as a batch of multiple URLs | 

### Return type

[**UrlSsrfResponseBatch**](UrlSsrfResponseBatch.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomainUrlFull**
> ValidateUrlResponseFull DomainUrlFull(ctx, request)
Validate a URL fully

Validate whether a URL is syntactically valid (does not check endpoint for validity), whether it exists, and whether the endpoint is up and passes virus scan checks.  Accepts various types of input and produces a well-formed URL as output.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**ValidateUrlRequestFull**](ValidateUrlRequestFull.md)| Input URL request | 

### Return type

[**ValidateUrlResponseFull**](ValidateUrlResponseFull.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomainUrlHtmlSsrfCheck**
> UrlHtmlSsrfResponseFull DomainUrlHtmlSsrfCheck(ctx, request)
Check a URL for HTML embedded SSRF threats

Checks if an input URL HTML is at risk of containing one or more embedded SSRF (Server-side request forgery) threats or attacks.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**UrlHtmlSsrfRequestFull**](UrlHtmlSsrfRequestFull.md)| Input URL request | 

### Return type

[**UrlHtmlSsrfResponseFull**](UrlHtmlSsrfResponseFull.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomainUrlSyntaxOnly**
> ValidateUrlResponseSyntaxOnly DomainUrlSyntaxOnly(ctx, request)
Validate a URL syntactically

Validate whether a URL is syntactically valid (does not check endpoint for validity).  Accepts various types of input and produces a well-formed URL as output.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**ValidateUrlRequestSyntaxOnly**](ValidateUrlRequestSyntaxOnly.md)| Input URL information | 

### Return type

[**ValidateUrlResponseSyntaxOnly**](ValidateUrlResponseSyntaxOnly.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

