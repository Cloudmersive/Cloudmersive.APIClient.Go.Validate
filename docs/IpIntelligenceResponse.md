# IpIntelligenceResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsBot** | **bool** | True if the IP address is a known bot, otherwise false | [optional] [default to null]
**IsTorNode** | **bool** | True if the IP address is a known Tor exit node, otherwise false | [optional] [default to null]
**IsThreat** | **bool** | True if the IP address is a known threat IP, otherwise false | [optional] [default to null]
**IsEU** | **bool** | True if the IP address is in the European Union, otherwise false | [optional] [default to null]
**Location** | [***GeolocateResponse**](GeolocateResponse.md) | Returns the location of the IP address | [optional] [default to null]
**CurrencyCode** | **string** | ISO 4217 currency code for the IP address location | [optional] [default to null]
**CurrencyName** | **string** | Name of the currency in English | [optional] [default to null]
**RegionArea** | **string** | Region (continent) in which the country is located; possible values are None, Europe, Americas, Asia, Africa, Oceania | [optional] [default to null]
**SubregionArea** | **string** | Subregion in which the country is located; possible values are None, NorthernEurope, WesternEurope, SouthernEurope, EasternEurope, CentralAmerica, NorthernAmerica, SouthAmerica, EasternAfrica, MiddleAfrica, NorthernAfrica , SouthernAfrica , WesternAfrica , CentralAsia , EasternAsia , SouthernAsia , SouthEasternAsia , WesternAsia , Southern , Middle , AustraliaandNewZealand , Melanesia , Polynesia , Micronesia , Caribbean, | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


