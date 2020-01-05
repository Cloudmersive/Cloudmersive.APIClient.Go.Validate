/*
 * validateapi
 *
 * The validation APIs help you validate data. Check if an E-mail address is real. Check if a domain is real. Check up on an IP address, and even where it is located. All this and much more is available in the validation API.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package Go-CloudmersiveValidateApiClient

// Input parameter to a country validation request
type ValidateCountryRequest struct {
	// Raw country input - can be a two-letter code (FIPS 10-4 or ISO 3166-1), three-letter code (ISO 3166-1) or country name
	RawCountryInput string `json:"RawCountryInput,omitempty"`
}
