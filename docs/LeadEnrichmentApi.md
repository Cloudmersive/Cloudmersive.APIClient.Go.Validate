# \LeadEnrichmentApi

All URIs are relative to *https://api.cloudmersive.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LeadEnrichmentEnrichLead**](LeadEnrichmentApi.md#LeadEnrichmentEnrichLead) | **Post** /validate/lead-enrichment/lead/enrich | Enrich an input lead with additional fields of data


# **LeadEnrichmentEnrichLead**
> LeadEnrichmentResponse LeadEnrichmentEnrichLead(ctx, request)
Enrich an input lead with additional fields of data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**LeadEnrichmentRequest**](LeadEnrichmentRequest.md)| Input lead with known fields set, and unknown fields left blank (null) | 

### Return type

[**LeadEnrichmentResponse**](LeadEnrichmentResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

