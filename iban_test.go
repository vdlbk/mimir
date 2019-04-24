package mimir

import (
	"testing"
)

func Test_IsIBANValid(t *testing.T) {
	for _, configuration := range countriesConfiguration {

		if ok, _ := IsIBANValid(configuration.IBANDefinition.Example); !ok {
			t.Errorf("Output was incorrect, IsIBANValid(%s)= false, want: true.", configuration.IBANDefinition.Example)
		}
	}
}

func Test_GetCheckDigits(t *testing.T) {
	for _, configuration := range countriesConfiguration {

		check, iban, err := GetCheckDigits(configuration.IBANDefinition.Example)
		if err != nil {
			t.Errorf("Unexpected error occured: %v", err)
		}

		if iban != configuration.IBANDefinition.Example {
			t.Errorf("Output was incorrect, GetCheckDigit(%s)= %v, %v; want: %v, %v.",
				configuration.IBANDefinition.Example,
				check,
				iban,
				configuration.IBANDefinition.Example[2:4],
				configuration.IBANDefinition.Example,
			)
		}
	}
}