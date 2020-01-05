/*
 * validateapi
 *
 * The validation APIs help you validate data. Check if an E-mail address is real. Check if a domain is real. Check up on an IP address, and even where it is located. All this and much more is available in the validation API.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package Go-CloudmersiveValidateApiClient

// Input lead contact; fill in known fields to extend them with matched field values
type LeadEnrichmentRequest struct {
	// The person's business email address for the lead
	ContactBusinessEmail string `json:"ContactBusinessEmail,omitempty"`
	// The person's first name for the lead
	ContactFirstName string `json:"ContactFirstName,omitempty"`
	// The person's last name for the lead
	ContactLastName string `json:"ContactLastName,omitempty"`
	// Name of the company for the lead
	CompanyName string `json:"CompanyName,omitempty"`
	// Domain name / website for the lead
	CompanyDomainName string `json:"CompanyDomainName,omitempty"`
	// House number of the address of the company for the lead
	CompanyHouseNumber string `json:"CompanyHouseNumber,omitempty"`
	// Street name of the address of the company for the lead
	CompanyStreet string `json:"CompanyStreet,omitempty"`
	// City of the address of the company for the lead
	CompanyCity string `json:"CompanyCity,omitempty"`
	// State or Province of the address of the company for the lead
	CompanyStateOrProvince string `json:"CompanyStateOrProvince,omitempty"`
	// Postal Code of the address of the company for the lead
	CompanyPostalCode string `json:"CompanyPostalCode,omitempty"`
	// Country of the address of the company for the lead
	CompanyCountry string `json:"CompanyCountry,omitempty"`
	// Country Code (2-letter ISO 3166-1) of the address of the company for the lead
	CompanyCountryCode string `json:"CompanyCountryCode,omitempty"`
	// Telephone of the company office for the lead
	CompanyTelephone string `json:"CompanyTelephone,omitempty"`
	// VAT number of the company for the lead
	CompanyVATNumber string `json:"CompanyVATNumber,omitempty"`
}
