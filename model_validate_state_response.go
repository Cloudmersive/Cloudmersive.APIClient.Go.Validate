/*
 * validateapi
 *
 * The validation APIs help you validate data. Check if an E-mail address is real. Check if a domain is real. Check up on an IP address, and even where it is located. All this and much more is available in the validation API.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package Go-CloudmersiveValidateApiClient

// Result of validating a state
type ValidateStateResponse struct {
	// True if the address is valid, false otherwise
	ValidState bool `json:"ValidState,omitempty"`
	// If valid; State or province corresponding to the input state name, such as 'CA' or 'California'
	StateOrProvince string `json:"StateOrProvince,omitempty"`
	// If the postal code is valid, the degrees latitude of the centroid of the state, null otherwise
	Latitude float64 `json:"Latitude,omitempty"`
	// If the postal code is valid, the degrees longitude of the centroid of the state, null otherwise
	Longitude float64 `json:"Longitude,omitempty"`
}
