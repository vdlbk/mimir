package mimir

import "testing"

func TestCountryMapConfiguration(t *testing.T) {
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

func TestCountryPrefixInIBAN(t *testing.T) {
	for _, configuration := range countriesConfiguration {

		if configuration.CountryCode.String() != configuration.IBANDefinition.Example[:2] {
			t.Errorf("Prefix Incorrect the configuration may not be accurate, got: %v, want: %v.",
				configuration.IBANDefinition.Example[:2],
				configuration.CountryCode.String(),
			)
		}
	}
}

func TestSEPACountry(t *testing.T) {
	const expectedSEPACountry = 42
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
