package e164

// peekNationalDestinationCode checks if `value` starts with a E.164 national-destination-code,
// and if it does then it returns its length (in bytes).
//
// For example:
//
//	// A phone-number from Vancouver, BC, Canada
//	var countyCode string = "1"
//	var value string = "6045551234"
//	
//	n, found := peekCountryCode(countyCode, value)
//	// n == 3
//	// found == true
//
// Also, for example:
//
//	// A phone-number from Alexandria, Egypt
//	var countyCode string = "20"
//	var value string = "212345678"
//	
//	n, found := peekCountryCode(value)
//	// n == 1
//	// found == true
func peekNationalDestinationCode(countryCode string, value string) (n int, found bool) {

	switch countryCode {
	case "1":
		const length int = 3

		if len(value) < length {
			return 0, false
		}
		return 3, true
//@TODO
	default:
		return 0, false
	}

	return 0, false
}
