package e164_test

import (
	"fmt"

	"github.com/reiver/go-e164"
)

func ExampleParseTolerantly_ca() {

	phoneNumber := "+16045551234"

	countryCode, nationalDestinationCode, subscriberNumber, err := e164.ParseTolerantly(phoneNumber)
	if nil != err {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	fmt.Printf("E.164 country-code: %s\n", countryCode)
	fmt.Printf("E.164 national-destination-code: %s\n", nationalDestinationCode)
	fmt.Printf("E.164 subscriber-number: %s\n", subscriberNumber)

	// Output:
	// E.164 country-code: 1
	// E.164 national-destination-code: 604
	// E.164 subscriber-number: 5551234
}

func ExampleParseTolerantly_us() {

	phoneNumber := "+12061234567"

	countryCode, nationalDestinationCode, subscriberNumber, err := e164.ParseTolerantly(phoneNumber)
	if nil != err {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	fmt.Printf("E.164 country-code: %s\n", countryCode)
	fmt.Printf("E.164 national-destination-code: %s\n", nationalDestinationCode)
	fmt.Printf("E.164 subscriber-number: %s\n", subscriberNumber)

	// Output:
	// E.164 country-code: 1
	// E.164 national-destination-code: 206
	// E.164 subscriber-number: 1234567
}
