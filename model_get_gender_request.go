/*
 * validateapi
 *
 * The validation APIs help you validate data. Check if an E-mail address is real. Check if a domain is real. Check up on an IP address, and even where it is located. All this and much more is available in the validation API.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package Go-CloudmersiveValidateApiClient

// Request to get the gender from a first name
type GetGenderRequest struct {
	// Input first name (given name) to get the gender of
	FirstName string `json:"FirstName,omitempty"`
	// Optional; the country for this name, possible values are \"US\", \"LY\", \"NI\", \"TT\", \"MK\", \"KZ\", \"BO\", \"UG\", \"TZ\", \"CL\", \"SI\", \"MA\", \"RW\", \"VN\", \"AW\", \"CY\", \"BH\", \"SG\", \"ZA\", \"MU\", \"BR\", \"TN\", \"KH\", \"US\", \"TH\", \"TW\", \"UY\", \"DO\", \"CO\", \"UA\", \"QA\", \"BY\", \"SN\", \"SD\", \"FJ\", \"LB\", \"BE\", \"ML\", \"LV\", \"FR\", \"TM\", \"NG\", \"EC\", \"NO\", \"SL\", \"CR\", \"PA\", \"GE\", \"CH\", \"KR\", \"RS\", \"ZM\", \"FI\", \"BF\", \"MC\", \"AU\", \"GA\", \"LS\", \"RU\", \"IN\", \"SE\", \"LK\", \"BZ\", \"MX\", \"GH\", \"AF\", \"TJ\", \"BN\", \"DZ\", \"CM\", \"GR\", \"MD\", \"HN\", \"AT\", \"NZ\", \"SV\", \"GW\", \"NA\", \"AR\", \"MZ\", \"PK\", \"MN\", \"IQ\", \"BW\", \"VE\", \"PT\", \"BS\", \"AL\", \"TG\", \"ID\", \"ET\", \"CF\", \"JP\", \"BB\", \"PH\", \"CU\", \"BD\", \"AO\", \"SM\", \"LC\", \"ME\", \"RO\",  DANIL\"O\"\", \"ES\", \"EE\", \"IL\", \"ZW\", \"SY\", \"MW\", \"LU\", \"IR\", \"SC\", \"NL\", \"JO\", \"AM\", \"DE\", \"GL\", \"OM\", \"DK\", \"HR\", \"LI\", \"TD\", \"KM\", \"BA\", \"GM\", \"GD\", \"CA\", \"CZ\", \"MR\", \"ST\", \"IS\", \"LR\", \"IE\", \"VC\", \"AE\", \"KG\", \"DJ\", \"TR\", \"KE\", \"NE\", \"UZ\", \"CN\", \"GQ\", \"SK\", \"BJ\", \"MG\", \"BT\", \"EG\", \"PL\", \"IT\", \"SA\", \"MY\", \"CI\", \"AG\", \"AD\", \"KS\", \"HU\", \"CG\", \"KP\", \"DM\", \"GN\", \"GT\", \"NP\", \"JM\", \"LA\", \"GB\", \"BG\", \"HT\", \"PE\", \"AZ\", \"LT\", \"SZ\", \"PY\", \"MT\", \"VA\"
	CountryCode string `json:"CountryCode,omitempty"`
}
