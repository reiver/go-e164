package e164

import (
	"strings"

	"github.com/reiver/go-erorr"
)

const (
	ErrBadCountryCode             = erorr.Error("bad country-code")
	ErrBadNationalDestinationCode = erorr.Error("bad national-destination-code")
	ErrBadSubscriberNumber        = erorr.Error("bad subscriber-number")
	ErrEmptyPhoneNumber           = erorr.Error("empty phone-number")
)

// ParseWithIsoCountryCodeTolerantly parses `value` for a E.164 phone-number (without the country-code) tolerantly,
// such that it allows spaces between the country-code, the national-destination-code and the subscriber-number.
// The country-code implied by the ISO 3166-1 alpha-2 country-code,
//
// The basic format is:
//
//	[area code][local phone number]
//
// For example:
//
//	var countryCode string = "CA"
//	var phoneNumber string = "6045551234"
//	
// 	countryCode, nationalDestinationCode, subscriberNumber, err := e164.ParseTolerantly(phoneNumber)
//	// countryCode == "1"
//	// nationalDestinationCode == "604"
//	// subscriberNumber == "5551234"
func ParseWithIsoCountryCodeTolerantly(isoCountryCode string, value string) (countryCode string, nationalDestinationCode string, subscriberNumber string, err error) {
	if len(value) <= 0 {
		err = ErrEmptyPhoneNumber
		return
	}

	countryCode = isoToCountryCode(isoCountryCode)
	if "" == countryCode {
		err = ErrBadCountryCode
		return
	}

	return ParseTolerantly(countryCode + value)
}

// ParseTolerantly parses `value` for a E.164 phone-number tolerantly,
// such that it allows spaces between the country-code, the national-destination-code and the subscriber-number.
//
// The basic format is:
//
//	[+][country code][area code][local phone number]
//
// For example:
//
//	var phoneNumber string = "+16045551234"
//	
// 	countryCode, nationalDestinationCode, subscriberNumber, err := e164.ParseTolerantly(phoneNumber)
//	// countryCode == "1"
//	// nationalDestinationCode == "604"
//	// subscriberNumber == "5551234"
func ParseTolerantly(value string) (countryCode string, nationalDestinationCode string, subscriberNumber string, err error) {

	if len(value) <= 0 {
		err = ErrEmptyPhoneNumber
		return
	}

	{
		const plus byte = '+'
		const plusString string = string(plus)

		val0 := value[0]

		if plus == val0 {
			value = value[len(plusString):]
		}
	}

	{
		n, found := peekCountryCode(value)
		if !found {
			err = ErrBadCountryCode
			return
		}

		countryCode = value[:n]
		value = value[n:]
	}

	{
		value = strings.TrimLeft(value, " ")
	}

	{
		n, found := peekNationalDestinationCode(countryCode, value)
		if !found {
			err = ErrBadNationalDestinationCode
			return
		}

		nationalDestinationCode = value[:n]
		value = value[n:]
	}

	{
		value = strings.TrimLeft(value, " ")
	}

	subscriberNumber = value
	for _, r := range subscriberNumber {
		switch {
		case '0' <= r && r <= '9':
			// nothing here
		default:
			err = ErrBadSubscriberNumber
			return
		}
	}

	return
}
