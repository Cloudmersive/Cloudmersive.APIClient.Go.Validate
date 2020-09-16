/*
 * validateapi
 *
 * The validation APIs help you validate data. Check if an E-mail address is real. Check if a domain is real. Check up on an IP address, and even where it is located. All this and much more is available in the validation API.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package Go-CloudmersiveValidateApiClient

// Result of performing a country validation operation
type ValidateCountryResponse struct {
	// True if successful, false otherwise
	Successful bool `json:"Successful,omitempty"`
	// Full name of the country
	CountryFullName string `json:"CountryFullName,omitempty"`
	// Two-letter ISO 3166-1 country code
	ISOTwoLetterCode string `json:"ISOTwoLetterCode,omitempty"`
	// Two-letter FIPS 10-4 country code
	FIPSTwoLetterCode string `json:"FIPSTwoLetterCode,omitempty"`
	// Three-letter ISO 3166-1 country code
	ThreeLetterCode string `json:"ThreeLetterCode,omitempty"`
	// True if this country is currently a member of the European Union (EU), false otherwise
	IsEuropeanUnionMember bool `json:"IsEuropeanUnionMember,omitempty"`
	// Time zones (IANA/Olsen) in the country
	Timezones []Timezone `json:"Timezones,omitempty"`
	// ISO 4217 currency three-letter code associated with the country
	ISOCurrencyCode string `json:"ISOCurrencyCode,omitempty"`
	// Symbol associated with the currency
	CurrencySymbol string `json:"CurrencySymbol,omitempty"`
	// Full name of the currency
	CurrencyEnglishName string `json:"CurrencyEnglishName,omitempty"`
	// Region (continent) in which the country is located; possible values are None, Europe, Americas, Asia, Africa, Oceania
	Region string `json:"Region,omitempty"`
	// Subregion in which the country is located; possible values are None, NorthernEurope, WesternEurope, SouthernEurope, EasternEurope, CentralAmerica, NorthernAmerica, SouthAmerica, EasternAfrica, MiddleAfrica, NorthernAfrica , SouthernAfrica , WesternAfrica , CentralAsia , EasternAsia , SouthernAsia , SouthEasternAsia , WesternAsia , Southern , Middle , AustraliaandNewZealand , Melanesia , Polynesia , Micronesia , Caribbean,
	Subregion string `json:"Subregion,omitempty"`
}
