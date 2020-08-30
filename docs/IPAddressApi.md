# \IPAddressApi

All URIs are relative to *https://api.cloudmersive.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IPAddressIsTorNode**](IPAddressApi.md#IPAddressIsTorNode) | **Post** /validate/ip/is-tor-node | Check if IP address is a Tor node server
[**IPAddressPost**](IPAddressApi.md#IPAddressPost) | **Post** /validate/ip/geolocate | Geolocate an IP address


# **IPAddressIsTorNode**
> TorNodeResponse IPAddressIsTorNode(ctx, value)
Check if IP address is a Tor node server

Check if the input IP address is a Tor exit node server.  Tor servers are a type of privacy-preserving technology that can hide the original IP address who makes a request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **value** | **string**| IP address to check, e.g. \&quot;55.55.55.55\&quot;.  The input is a string so be sure to enclose it in double-quotes. | 

### Return type

[**TorNodeResponse**](TorNodeResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: text/javascript, application/json, text/json
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IPAddressPost**
> GeolocateResponse IPAddressPost(ctx, value)
Geolocate an IP address

Identify an IP address Country, State/Provence, City, Zip/Postal Code, etc.  Useful for security and UX applications.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **value** | **string**| IP address to geolocate, e.g. \&quot;55.55.55.55\&quot;.  The input is a string so be sure to enclose it in double-quotes. | 

### Return type

[**GeolocateResponse**](GeolocateResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: text/javascript, application/json, text/json
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

