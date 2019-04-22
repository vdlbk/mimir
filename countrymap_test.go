package mimir

import "testing"

func Test_CountryMap(t *testing.T) {
	for _, configuration := range countriesConfiguration {

		if configuration.IBANDefinition.Length != len(configuration.IBANDefinition.Example) {
			t.Errorf("[%s] IBAN Incorrect length, got: %v, want: %v.",
				configuration.CountryCode.String(),
				len(configuration.IBANDefinition.Example),
				configuration.IBANDefinition.Length,
			)
		}

		if configuration.BBANDefinition.Length != len(configuration.BBANDefinition.Example) {
			t.Errorf("[%s] BBAN Incorrect length, got: %v, want: %v.",
				configuration.CountryCode.String(),
				len(configuration.BBANDefinition.Example),
				configuration.BBANDefinition.Length,
			)
		}
	}
}

func Test_SEPACountry(t *testing.T) {
	const expectedSEPACountry = 35
	ctrSEPACountry := 0

	for _, configuration := range countriesConfiguration {
		if configuration.IsSEPAAvailable {
			ctrSEPACountry++
		}
	}

	if ctrSEPACountry != expectedSEPACountry {
		t.Errorf("The number of Country with SEPA available is incorrect, got: %v, want: %v.", ctrSEPACountry, expectedSEPACountry)
	}
}
