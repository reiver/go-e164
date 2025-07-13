package e164_test

import (
	"testing"

	"github.com/reiver/go-e164"
)

func TestParseTolerantly(t *testing.T) {

	tests := []struct{
		String string
		ExpectedCountryCode string
		ExpectedNationalDestinationCode string
		ExpectedSubscriberNumber string
	}{
		{
			String:                         "16045551234",
			ExpectedCountryCode:            "1",
			ExpectedNationalDestinationCode: "604",
			ExpectedSubscriberNumber:           "5551234",
		},
		{
			String:                        "+16045551234",
			ExpectedCountryCode:            "1",
			ExpectedNationalDestinationCode: "604",
			ExpectedSubscriberNumber:           "5551234",
		},



		{
			String:                         "1 6045551234",
			ExpectedCountryCode:            "1",
			ExpectedNationalDestinationCode : "604",
			ExpectedSubscriberNumber:            "5551234",
		},
		{
			String:                        "+1 6045551234",
			ExpectedCountryCode:            "1",
			ExpectedNationalDestinationCode : "604",
			ExpectedSubscriberNumber:            "5551234",
		},



		{
			String:                         "1604 5551234",
			ExpectedCountryCode:            "1",
			ExpectedNationalDestinationCode :"604",
			ExpectedSubscriberNumber:            "5551234",
		},
		{
			String:                        "+1604 5551234",
			ExpectedCountryCode:            "1",
			ExpectedNationalDestinationCode :"604",
			ExpectedSubscriberNumber:            "5551234",
		},



		{
			String:                         "1 604 5551234",
			ExpectedCountryCode:            "1",
			ExpectedNationalDestinationCode : "604",
			ExpectedSubscriberNumber:             "5551234",
		},
		{
			String:                        "+1 604 5551234",
			ExpectedCountryCode:            "1",
			ExpectedNationalDestinationCode : "604",
			ExpectedSubscriberNumber:             "5551234",
		},
	}

	for testNumber, test := range tests {

		actualCountryCode, actualNationalDestinationCode, actualSubscriberNumber, err := e164.ParseTolerantly(test.String)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("STRING: %q", test.String)
			continue
		}

		{
			actual := actualCountryCode
			expected := test.ExpectedCountryCode

			if expected != actual {
				t.Errorf("For test #%d, the actual 'country-code' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("STRING: %q", test.String)
				continue
			}
		}

		{
			actual := actualNationalDestinationCode
			expected := test.ExpectedNationalDestinationCode

			if expected != actual {
				t.Errorf("For test #%d, the actual 'national-destination-code' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("STRING: %q", test.String)
				continue
			}
		}

		{
			actual := actualSubscriberNumber
			expected := test.ExpectedSubscriberNumber

			if expected != actual {
				t.Errorf("For test #%d, the actual 'subscriber-number' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("STRING: %q", test.String)
				continue
			}
		}
	}
}
