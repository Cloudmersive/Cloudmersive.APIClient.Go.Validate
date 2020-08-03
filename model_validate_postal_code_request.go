/*
 * validateapi
 *
 * The validation APIs help you validate data. Check if an E-mail address is real. Check if a domain is real. Check up on an IP address, and even where it is located. All this and much more is available in the validation API.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package Go-CloudmersiveValidateApiClient

// Request to Validate a Postal Code
type ValidatePostalCodeRequest struct {
	// Required: Zip code or postal code of the address to validate, such as '94597'
	PostalCode string `json:"PostalCode,omitempty"`
	// Optional (recommended); Name of the country, such as 'United States'.  If left blank, and CountryCode is also left blank, will default to United States.  Global countries are supported.
	CountryFullName string `json:"CountryFullName,omitempty"`
	// Optional; two-letter country code (Two-letter ISO 3166-1 country code) of the country.  If left blank, and CountryFullName is also left blank, will default to United States.  Global countries are supported.
	CountryCode string `json:"CountryCode,omitempty"`
}
