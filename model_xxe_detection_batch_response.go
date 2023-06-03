/*
 * validateapi
 *
 * The validation APIs help you validate data. Check if an E-mail address is real. Check if a domain is real. Check up on an IP address, and even where it is located. All this and much more is available in the validation API.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package Go-CloudmersiveValidateApiClient

// Result of performing a batch XXE detection operation
type XxeDetectionBatchResponse struct {
	// Results from performing a batch XSS protection operation
	ResultItems []XxeDetectionResult `json:"ResultItems,omitempty"`
}