/*
 * validateapi
 *
 * The validation APIs help you validate data. Check if an E-mail address is real. Check if a domain is real. Check up on an IP address, and even where it is located. All this and much more is available in the validation API.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package Go-CloudmersiveValidateApiClient

// Result from validating a phone number
type PhoneNumberValidationResponse struct {
	// True if the phone number is valid, false otherwise
	IsValid bool `json:"IsValid,omitempty"`
	// True if the operation was successful, false if there was an error during validation.  See IsValid for validation result.
	Successful bool `json:"Successful,omitempty"`
	// Type of phone number; possible values are: FixedLine, Mobile, FixedLineOrMobile, TollFree, PremiumRate,   SharedCost, Voip, PersonalNumber, Pager, Uan, Voicemail, Unknown
	PhoneNumberType string `json:"PhoneNumberType,omitempty"`
	// E.164 format of the phone number
	E164Format string `json:"E164Format,omitempty"`
	// Internaltional format of the phone number
	InternationalFormat string `json:"InternationalFormat,omitempty"`
	// National format of the phone number
	NationalFormat string `json:"NationalFormat,omitempty"`
	// Two digit country code of the phone number
	CountryCode string `json:"CountryCode,omitempty"`
	// User-friendly long name of the country for the phone number
	CountryName string `json:"CountryName,omitempty"`
}
