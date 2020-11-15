/*
 * validateapi
 *
 * The validation APIs help you validate data. Check if an E-mail address is real. Check if a domain is real. Check up on an IP address, and even where it is located. All this and much more is available in the validation API.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package Go-CloudmersiveValidateApiClient

// Input parameter to a date time parsing request
type DateTimeNaturalLanguageParseRequest struct {
	// Raw string input of a natural language-formatted date and time for parsing
	RawDateTimeInput string `json:"RawDateTimeInput,omitempty"`
}
