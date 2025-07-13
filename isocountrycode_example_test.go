package e164_test

import (
	"fmt"

	"github.com/reiver/go-e164"
)

func ExampleISOCountryCode_ca() {

	e164CountryCode := "1"
	e164NationalDestinationCode := "604" // Also called an area-code

	isoCountryCode := e164.ISOCountryCode(e164CountryCode, e164NationalDestinationCode)

	fmt.Printf("ISO 3166-1 alpha-2 country-code: %s\n", isoCountryCode)

	// Output:
	// ISO 3166-1 alpha-2 country-code: CA
}

func ExampleISOCountryCode_kr() {

	e164CountryCode := "82"
	e164NationalDestinationCode := "2" // Also called an area-code

	isoCountryCode := e164.ISOCountryCode(e164CountryCode, e164NationalDestinationCode)

	fmt.Printf("ISO 3166-1 alpha-2 country-code: %s\n", isoCountryCode)

	// Output:
	// ISO 3166-1 alpha-2 country-code: KR
}

func ExampleISOCountryCode_us() {

	e164CountryCode := "1"
	e164NationalDestinationCode := "206" // Also called an area-code

	isoCountryCode := e164.ISOCountryCode(e164CountryCode, e164NationalDestinationCode)

	fmt.Printf("ISO 3166-1 alpha-2 country-code: %s\n", isoCountryCode)

	// Output:
	// ISO 3166-1 alpha-2 country-code: US
}
