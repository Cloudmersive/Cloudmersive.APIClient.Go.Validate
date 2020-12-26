# Go API client for Go-CloudmersiveValidateApiClient

The validation APIs help you validate data. Check if an E-mail address is real. Check if a domain is real. Check up on an IP address, and even where it is located. All this and much more is available in the validation API.

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: 3.1.3
- Build package: io.swagger.codegen.languages.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./Go-CloudmersiveValidateApiClient"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.cloudmersive.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AddressApi* | [**AddressCheckEUMembership**](docs/AddressApi.md#addresscheckeumembership) | **Post** /validate/address/country/check-eu-membership | Check if a country is a member of the European Union (EU)
*AddressApi* | [**AddressCountry**](docs/AddressApi.md#addresscountry) | **Post** /validate/address/country | Validate and normalize country information, return ISO 3166-1 country codes and country name
*AddressApi* | [**AddressCountryList**](docs/AddressApi.md#addresscountrylist) | **Post** /validate/address/country/list | Get a list of ISO 3166-1 countries
*AddressApi* | [**AddressGeocode**](docs/AddressApi.md#addressgeocode) | **Post** /validate/address/geocode | Geocode a street address into latitude and longitude
*AddressApi* | [**AddressGetCountryCurrency**](docs/AddressApi.md#addressgetcountrycurrency) | **Post** /validate/address/country/get-currency | Get the currency of the input country
*AddressApi* | [**AddressGetCountryRegion**](docs/AddressApi.md#addressgetcountryregion) | **Post** /validate/address/country/get-region | Get the region, subregion and continent of the country
*AddressApi* | [**AddressGetTimezone**](docs/AddressApi.md#addressgettimezone) | **Post** /validate/address/country/get-timezones | Gets IANA/Olsen time zones for a country
*AddressApi* | [**AddressNormalizeAddress**](docs/AddressApi.md#addressnormalizeaddress) | **Post** /validate/address/street-address/normalize | Normalize a street address
*AddressApi* | [**AddressParseString**](docs/AddressApi.md#addressparsestring) | **Post** /validate/address/parse | Parse an unstructured input text string into an international, formatted address
*AddressApi* | [**AddressReverseGeocodeAddress**](docs/AddressApi.md#addressreversegeocodeaddress) | **Post** /validate/address/geocode/reverse | Reverse geocode a lattitude and longitude into an address
*AddressApi* | [**AddressValidateAddress**](docs/AddressApi.md#addressvalidateaddress) | **Post** /validate/address/street-address | Validate a street address
*AddressApi* | [**AddressValidateCity**](docs/AddressApi.md#addressvalidatecity) | **Post** /validate/address/city | Validate a City and State/Province combination, get location information about it
*AddressApi* | [**AddressValidatePostalCode**](docs/AddressApi.md#addressvalidatepostalcode) | **Post** /validate/address/postal-code | Validate a postal code, get location information about it
*AddressApi* | [**AddressValidateState**](docs/AddressApi.md#addressvalidatestate) | **Post** /validate/address/state | Validate a state or province, name or abbreviation, get location information about it
*DateTimeApi* | [**DateTimeGetNowSimple**](docs/DateTimeApi.md#datetimegetnowsimple) | **Get** /validate/date-time/get/now | Get current date and time as of now
*DateTimeApi* | [**DateTimeGetPublicHolidays**](docs/DateTimeApi.md#datetimegetpublicholidays) | **Post** /validate/date-time/get/holidays | Get public holidays in the specified country and year
*DateTimeApi* | [**DateTimeParseNaturalLanguageDateTime**](docs/DateTimeApi.md#datetimeparsenaturallanguagedatetime) | **Post** /validate/date-time/parse/date-time/natural-language | Parses a free-form natural language date and time string into a date and time
*DateTimeApi* | [**DateTimeParseStandardDateTime**](docs/DateTimeApi.md#datetimeparsestandarddatetime) | **Post** /validate/date-time/parse/date-time/structured | Parses a standardized date and time string into a date and time
*DomainApi* | [**DomainCheck**](docs/DomainApi.md#domaincheck) | **Post** /validate/domain/check | Validate a domain name
*DomainApi* | [**DomainPost**](docs/DomainApi.md#domainpost) | **Post** /validate/domain/whois | Get WHOIS information for a domain
*DomainApi* | [**DomainQualityScore**](docs/DomainApi.md#domainqualityscore) | **Post** /validate/domain/quality-score | Validate a domain name&#39;s quality score
*DomainApi* | [**DomainUrlFull**](docs/DomainApi.md#domainurlfull) | **Post** /validate/domain/url/full | Validate a URL fully
*DomainApi* | [**DomainUrlSyntaxOnly**](docs/DomainApi.md#domainurlsyntaxonly) | **Post** /validate/domain/url/syntax-only | Validate a URL syntactically
*EmailApi* | [**EmailAddressGetServers**](docs/EmailApi.md#emailaddressgetservers) | **Post** /validate/email/address/servers | Partially check whether an email address is valid
*EmailApi* | [**EmailFullValidation**](docs/EmailApi.md#emailfullvalidation) | **Post** /validate/email/address/full | Fully validate an email address
*EmailApi* | [**EmailPost**](docs/EmailApi.md#emailpost) | **Post** /validate/email/address/syntaxOnly | Validate email adddress for syntactic correctness only
*IPAddressApi* | [**IPAddressGeolocateStreetAddress**](docs/IPAddressApi.md#ipaddressgeolocatestreetaddress) | **Post** /validate/ip/geolocate/street-address | Geolocate an IP address to a street address
*IPAddressApi* | [**IPAddressIsThreat**](docs/IPAddressApi.md#ipaddressisthreat) | **Post** /validate/ip/is-threat | Check if IP address is a known threat
*IPAddressApi* | [**IPAddressIsTorNode**](docs/IPAddressApi.md#ipaddressistornode) | **Post** /validate/ip/is-tor-node | Check if IP address is a Tor node server
*IPAddressApi* | [**IPAddressPost**](docs/IPAddressApi.md#ipaddresspost) | **Post** /validate/ip/geolocate | Geolocate an IP address
*IPAddressApi* | [**IPAddressReverseDomainLookup**](docs/IPAddressApi.md#ipaddressreversedomainlookup) | **Post** /validate/ip/reverse-domain-lookup | Perform a reverse domain name (DNS) lookup on an IP address
*LeadEnrichmentApi* | [**LeadEnrichmentEnrichLead**](docs/LeadEnrichmentApi.md#leadenrichmentenrichlead) | **Post** /validate/lead-enrichment/lead/enrich | Enrich an input lead with additional fields of data
*NameApi* | [**NameGetGender**](docs/NameApi.md#namegetgender) | **Post** /validate/name/get-gender | Get the gender of a first name
*NameApi* | [**NameIdentifier**](docs/NameApi.md#nameidentifier) | **Post** /validate/name/identifier | Validate a code identifier
*NameApi* | [**NameValidateFirstName**](docs/NameApi.md#namevalidatefirstname) | **Post** /validate/name/first | Validate a first name
*NameApi* | [**NameValidateFullName**](docs/NameApi.md#namevalidatefullname) | **Post** /validate/name/full-name | Parse and validate a full name
*NameApi* | [**NameValidateLastName**](docs/NameApi.md#namevalidatelastname) | **Post** /validate/name/last | Validate a last name
*PhoneNumberApi* | [**PhoneNumberSyntaxOnly**](docs/PhoneNumberApi.md#phonenumbersyntaxonly) | **Post** /validate/phonenumber/basic | Validate phone number (basic)
*TextInputApi* | [**TextInputCheckXss**](docs/TextInputApi.md#textinputcheckxss) | **Post** /validate/text-input/check/xss | Check text input for Cross-Site-Scripting (XSS) attacks
*TextInputApi* | [**TextInputCheckXssBatch**](docs/TextInputApi.md#textinputcheckxssbatch) | **Post** /validate/text-input/check-and-protect/xss/batch | Check and protect multiple text inputs for Cross-Site-Scripting (XSS) attacks in batch
*TextInputApi* | [**TextInputProtectXss**](docs/TextInputApi.md#textinputprotectxss) | **Post** /validate/text-input/protect/xss | Protect text input from Cross-Site-Scripting (XSS) attacks through normalization
*UserAgentApi* | [**UserAgentParse**](docs/UserAgentApi.md#useragentparse) | **Post** /validate/useragent/parse | Parse an HTTP User-Agent string, identify robots
*VatApi* | [**VatVatLookup**](docs/VatApi.md#vatvatlookup) | **Post** /validate/vat/lookup | Validate a VAT number


## Documentation For Models

 - [AddressGetServersResponse](docs/AddressGetServersResponse.md)
 - [AddressVerifySyntaxOnlyResponse](docs/AddressVerifySyntaxOnlyResponse.md)
 - [CheckResponse](docs/CheckResponse.md)
 - [CountryDetails](docs/CountryDetails.md)
 - [CountryListResult](docs/CountryListResult.md)
 - [DateTimeNaturalLanguageParseRequest](docs/DateTimeNaturalLanguageParseRequest.md)
 - [DateTimeNowResult](docs/DateTimeNowResult.md)
 - [DateTimeStandardizedParseRequest](docs/DateTimeStandardizedParseRequest.md)
 - [DateTimeStandardizedParseResponse](docs/DateTimeStandardizedParseResponse.md)
 - [DomainQualityResponse](docs/DomainQualityResponse.md)
 - [FirstNameValidationRequest](docs/FirstNameValidationRequest.md)
 - [FirstNameValidationResponse](docs/FirstNameValidationResponse.md)
 - [FullEmailValidationResponse](docs/FullEmailValidationResponse.md)
 - [FullNameValidationRequest](docs/FullNameValidationRequest.md)
 - [FullNameValidationResponse](docs/FullNameValidationResponse.md)
 - [GeolocateResponse](docs/GeolocateResponse.md)
 - [GeolocateStreetAddressResponse](docs/GeolocateStreetAddressResponse.md)
 - [GetGenderRequest](docs/GetGenderRequest.md)
 - [GetGenderResponse](docs/GetGenderResponse.md)
 - [GetPublicHolidaysRequest](docs/GetPublicHolidaysRequest.md)
 - [GetTimezonesRequest](docs/GetTimezonesRequest.md)
 - [GetTimezonesResponse](docs/GetTimezonesResponse.md)
 - [IpReverseDnsLookupResponse](docs/IpReverseDnsLookupResponse.md)
 - [IpThreatResponse](docs/IpThreatResponse.md)
 - [LastNameValidationRequest](docs/LastNameValidationRequest.md)
 - [LastNameValidationResponse](docs/LastNameValidationResponse.md)
 - [LeadEnrichmentRequest](docs/LeadEnrichmentRequest.md)
 - [LeadEnrichmentResponse](docs/LeadEnrichmentResponse.md)
 - [NormalizeAddressResponse](docs/NormalizeAddressResponse.md)
 - [ParseAddressRequest](docs/ParseAddressRequest.md)
 - [ParseAddressResponse](docs/ParseAddressResponse.md)
 - [PhoneNumberValidateRequest](docs/PhoneNumberValidateRequest.md)
 - [PhoneNumberValidationResponse](docs/PhoneNumberValidationResponse.md)
 - [PublicHolidayOccurrence](docs/PublicHolidayOccurrence.md)
 - [PublicHolidaysResponse](docs/PublicHolidaysResponse.md)
 - [ReverseGeocodeAddressRequest](docs/ReverseGeocodeAddressRequest.md)
 - [ReverseGeocodeAddressResponse](docs/ReverseGeocodeAddressResponse.md)
 - [Timezone](docs/Timezone.md)
 - [TorNodeResponse](docs/TorNodeResponse.md)
 - [UserAgentValidateRequest](docs/UserAgentValidateRequest.md)
 - [UserAgentValidateResponse](docs/UserAgentValidateResponse.md)
 - [ValidateAddressRequest](docs/ValidateAddressRequest.md)
 - [ValidateAddressResponse](docs/ValidateAddressResponse.md)
 - [ValidateCityRequest](docs/ValidateCityRequest.md)
 - [ValidateCityResponse](docs/ValidateCityResponse.md)
 - [ValidateCountryRequest](docs/ValidateCountryRequest.md)
 - [ValidateCountryResponse](docs/ValidateCountryResponse.md)
 - [ValidateIdentifierRequest](docs/ValidateIdentifierRequest.md)
 - [ValidateIdentifierResponse](docs/ValidateIdentifierResponse.md)
 - [ValidatePostalCodeRequest](docs/ValidatePostalCodeRequest.md)
 - [ValidatePostalCodeResponse](docs/ValidatePostalCodeResponse.md)
 - [ValidateStateRequest](docs/ValidateStateRequest.md)
 - [ValidateStateResponse](docs/ValidateStateResponse.md)
 - [ValidateUrlRequestFull](docs/ValidateUrlRequestFull.md)
 - [ValidateUrlRequestSyntaxOnly](docs/ValidateUrlRequestSyntaxOnly.md)
 - [ValidateUrlResponseFull](docs/ValidateUrlResponseFull.md)
 - [ValidateUrlResponseSyntaxOnly](docs/ValidateUrlResponseSyntaxOnly.md)
 - [VatLookupRequest](docs/VatLookupRequest.md)
 - [VatLookupResponse](docs/VatLookupResponse.md)
 - [WhoisResponse](docs/WhoisResponse.md)
 - [XssProtectionBatchRequest](docs/XssProtectionBatchRequest.md)
 - [XssProtectionBatchResponse](docs/XssProtectionBatchResponse.md)
 - [XssProtectionRequestItem](docs/XssProtectionRequestItem.md)
 - [XssProtectionResult](docs/XssProtectionResult.md)


## Documentation For Authorization

## Apikey
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```

## Author



