/*
 * validateapi
 *
 * The validation APIs help you validate data. Check if an E-mail address is real. Check if a domain is real. Check up on an IP address, and even where it is located. All this and much more is available in the validation API.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package Go-CloudmersiveValidateApiClient

// Result of performing a batch XSS protection operation
type SqlInjectionCheckBatchResponse struct {
	// Results from performing a batch SQL Injection detection operation; order is preserved from input data
	ResultItems []SqlInjectionDetectionResult `json:"ResultItems,omitempty"`
}
