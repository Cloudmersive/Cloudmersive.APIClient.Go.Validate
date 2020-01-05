/*
 * validateapi
 *
 * The validation APIs help you validate data. Check if an E-mail address is real. Check if a domain is real. Check up on an IP address, and even where it is located. All this and much more is available in the validation API.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package Go-CloudmersiveValidateApiClient

// Request to parse an address formatted as a string/free text into a structured address
type ParseAddressRequest struct {
	// A mailing address or street address formatted as a single text string; this will be parsed into its components
	AddressString string `json:"AddressString,omitempty"`
	// Optional; indicates how the parsed output should be capitalized; default is Title Case; possible values are: \"Uppercase\" will set the capitalization to UPPER CASE; \"Lowercase\" will set the capitalization to lower case; \"Titlecase\" will set the capitalization to Title Case; and \"Originalcase\" will preserve the original casing as much as possible
	CapitalizationMode string `json:"CapitalizationMode,omitempty"`
}
