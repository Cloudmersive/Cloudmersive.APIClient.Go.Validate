/*
 * validateapi
 *
 * The validation APIs help you validate data. Check if an E-mail address is real. Check if a domain is real. Check up on an IP address, and even where it is located. All this and much more is available in the validation API.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package Go-CloudmersiveValidateApiClient

// Result of performing a IP threat check on an IP address
type IpThreatResponse struct {
	// True if the input IP address is a threat, false otherwise
	IsThreat bool `json:"IsThreat,omitempty"`
	// Specifies the type of IP threat; possible values include Blocklist, Botnet, WebBot
	ThreatType string `json:"ThreatType,omitempty"`
}
