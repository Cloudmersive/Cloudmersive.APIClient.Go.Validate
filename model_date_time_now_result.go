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

// Current date and time response
type DateTimeNowResult struct {
	// True if successful, false otherwise
	Successful bool `json:"Successful,omitempty"`
	// Current date, time, and time zone in standard JSON date format
	Now time.Time `json:"Now,omitempty"`
	// Current GMT-time-zone date, time, and time zone in standard JSON date format
	NowGmt time.Time `json:"NowGmt,omitempty"`
}
