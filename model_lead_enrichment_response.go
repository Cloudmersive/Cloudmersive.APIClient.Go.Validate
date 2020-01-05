/*
 * validateapi
 *
 * The validation APIs help you validate data. Check if an E-mail address is real. Check if a domain is real. Check up on an IP address, and even where it is located. All this and much more is available in the validation API.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cloudmersive_validate_api_client

// Result of the lead enrichment process
type LeadEnrichmentResponse struct {
	// True if the operation was successful, false otherwise
	Successful bool `json:"Successful,omitempty"`
	// The type of the lead; possible types are Junk (a single individual using a disposable/throwaway email address); Individual (a single individual, typically a consumer, not purchasing on behalf of a business); SmallBusiness (a small business, typically with fewer than 100 employees); MediumBusiness (a medium business, larger than 100 employees but fewer than 1000 employees); Enterprise (a large business with greater than 1000 employees); Business (a business customer of unknown size)
	LeadType string `json:"LeadType,omitempty"`
	// The person's business email address for the lead
	ContactBusinessEmail string `json:"ContactBusinessEmail,omitempty"`
	// The person's first name for the lead
	ContactFirstName string `json:"ContactFirstName,omitempty"`
	// The person's last name for the lead
	ContactLastName string `json:"ContactLastName,omitempty"`
	// Gender for contact name; possible values are Male, Female, and Neutral (can be applied to Male or Female).  Requires ContactFirstName.
	ContactGender string `json:"ContactGender,omitempty"`
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
	// Country Name of the address of the company for the lead
	CompanyCountry string `json:"CompanyCountry,omitempty"`
	// Country Code (2-letter ISO 3166-1) of the address of the company for the lead
	CompanyCountryCode string `json:"CompanyCountryCode,omitempty"`
	// Telephone of the company office for the lead
	CompanyTelephone string `json:"CompanyTelephone,omitempty"`
	// VAT number of the company for the lead
	CompanyVATNumber string `json:"CompanyVATNumber,omitempty"`
	// Count of employees at the company (estimated), if available
	EmployeeCount int32 `json:"EmployeeCount,omitempty"`
}
