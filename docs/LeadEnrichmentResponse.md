# LeadEnrichmentResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Successful** | **bool** | True if the operation was successful, false otherwise | [optional] [default to null]
**LeadType** | **string** | The type of the lead; possible types are Junk (a single individual using a disposable/throwaway email address); Individual (a single individual, typically a consumer, not purchasing on behalf of a business); SmallBusiness (a small business, typically with fewer than 100 employees); MediumBusiness (a medium business, larger than 100 employees but fewer than 1000 employees); Enterprise (a large business with greater than 1000 employees); Business (a business customer of unknown size) | [optional] [default to null]
**ContactBusinessEmail** | **string** | The person&#39;s business email address for the lead | [optional] [default to null]
**ContactFirstName** | **string** | The person&#39;s first name for the lead | [optional] [default to null]
**ContactLastName** | **string** | The person&#39;s last name for the lead | [optional] [default to null]
**ContactGender** | **string** | Gender for contact name; possible values are Male, Female, and Neutral (can be applied to Male or Female).  Requires ContactFirstName. | [optional] [default to null]
**CompanyName** | **string** | Name of the company for the lead | [optional] [default to null]
**CompanyDomainName** | **string** | Domain name / website for the lead | [optional] [default to null]
**CompanyHouseNumber** | **string** | House number of the address of the company for the lead | [optional] [default to null]
**CompanyStreet** | **string** | Street name of the address of the company for the lead | [optional] [default to null]
**CompanyCity** | **string** | City of the address of the company for the lead | [optional] [default to null]
**CompanyStateOrProvince** | **string** | State or Province of the address of the company for the lead | [optional] [default to null]
**CompanyPostalCode** | **string** | Postal Code of the address of the company for the lead | [optional] [default to null]
**CompanyCountry** | **string** | Country Name of the address of the company for the lead | [optional] [default to null]
**CompanyCountryCode** | **string** | Country Code (2-letter ISO 3166-1) of the address of the company for the lead | [optional] [default to null]
**CompanyTelephone** | **string** | Telephone of the company office for the lead | [optional] [default to null]
**CompanyVATNumber** | **string** | VAT number of the company for the lead | [optional] [default to null]
**EmployeeCount** | **int32** | Count of employees at the company (estimated), if available | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


