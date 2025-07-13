package e164

import (
	"strings"
)

// peekCountryCode checks if `value` starts with a E.164 country-code,
// and if it does then it returns its length (in bytes).
//
// For example:
//
//	// A phone-number from Vancouver, BC, Canada
//	var value string = "16045551234"
//	
//	n, found := peekCountryCode(value)
//	// n == 1
//	// found == true
//
// Also, for example:
//
//	// A phone-number from Alexandria, Egypt
//	var value string = "20212345678"
//	
//	n, found := peekCountryCode(value)
//	// n == 2
//	// found == true
func peekCountryCode(value string) (n int, found bool) {
	for _, countyCode := range countryCodes {
		if strings.HasPrefix(value, countyCode) {
			return len(countyCode), true
		}
	}

	return 0, false
}
