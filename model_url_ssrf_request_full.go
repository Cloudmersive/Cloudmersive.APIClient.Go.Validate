/*
 * validateapi
 *
 * The validation APIs help you validate data. Check if an E-mail address is real. Check if a domain is real. Check up on an IP address, and even where it is located. All this and much more is available in the validation API.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package Go-CloudmersiveValidateApiClient

// Request to determine if a URL is an SSRF threat check
type UrlSsrfRequestFull struct {
	// URL to validate
	URL string `json:"URL,omitempty"`
	// Top level domains that you do not want to allow access to, e.g. mydomain.com - will block all subdomains as well
	BlockedDomains []string `json:"BlockedDomains,omitempty"`
}
