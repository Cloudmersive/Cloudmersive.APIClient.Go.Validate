/*
 * validateapi
 *
 * The validation APIs help you validate data. Check if an E-mail address is real. Check if a domain is real. Check up on an IP address, and even where it is located. All this and much more is available in the validation API.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package Go-CloudmersiveValidateApiClient

import (
	"time"
)

// Public holiday occurrence
type PublicHolidayOccurrence struct {
	// Name of the holiday in English
	EnglishName string `json:"EnglishName,omitempty"`
	// Local name of the holiday
	LocalName string `json:"LocalName,omitempty"`
	// Date of the holiday (start time)
	OccurrenceDate time.Time `json:"OccurrenceDate,omitempty"`
	// Type of the holiday; possible values are: Public, Bank, School, Authorities, Optional, Observance
	HolidayType string `json:"HolidayType,omitempty"`
	// True if the holiday is celebrated in all locales in the country, false otherwise
	Nationwaide bool `json:"Nationwaide,omitempty"`
}
