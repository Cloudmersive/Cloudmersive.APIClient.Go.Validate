/*
 * validateapi
 *
 * The validation APIs help you validate data. Check if an E-mail address is real. Check if a domain is real. Check up on an IP address, and even where it is located. All this and much more is available in the validation API.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package Go-CloudmersiveValidateApiClient

// IP address intelligence result
type IpIntelligenceResponse struct {
	// True if the IP address is a known bot, otherwise false
	IsBot bool `json:"IsBot,omitempty"`
	// True if the IP address is a known Tor exit node, otherwise false
	IsTorNode bool `json:"IsTorNode,omitempty"`
	// True if the IP address is a known threat IP, otherwise false
	IsThreat bool `json:"IsThreat,omitempty"`
	// True if the IP address is in the European Union, otherwise false
	IsEU bool `json:"IsEU,omitempty"`
	// Returns the location of the IP address
	Location *GeolocateResponse `json:"Location,omitempty"`
	// ISO 4217 currency code for the IP address location
	CurrencyCode string `json:"CurrencyCode,omitempty"`
	// Name of the currency in English
	CurrencyName string `json:"CurrencyName,omitempty"`
	// Region (continent) in which the country is located; possible values are None, Europe, Americas, Asia, Africa, Oceania
	RegionArea string `json:"RegionArea,omitempty"`
	// Subregion in which the country is located; possible values are None, NorthernEurope, WesternEurope, SouthernEurope, EasternEurope, CentralAmerica, NorthernAmerica, SouthAmerica, EasternAfrica, MiddleAfrica, NorthernAfrica , SouthernAfrica , WesternAfrica , CentralAsia , EasternAsia , SouthernAsia , SouthEasternAsia , WesternAsia , Southern , Middle , AustraliaandNewZealand , Melanesia , Polynesia , Micronesia , Caribbean,
	SubregionArea string `json:"SubregionArea,omitempty"`
}