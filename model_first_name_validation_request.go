/*
 * validateapi
 *
 * The validation APIs help you validate data. Check if an E-mail address is real. Check if a domain is real. Check up on an IP address, and even where it is located. All this and much more is available in the validation API.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cloudmersive_validate_api_client

// Request to validate a first name
type FirstNameValidationRequest struct {
	// First name to process
	FirstName string `json:"FirstName,omitempty"`
}
