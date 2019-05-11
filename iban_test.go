package mimir

import (
	"fmt"
	"strings"
	"testing"
)

// Tests IsIBANValid with default values from countries configuration
func TestIsIBANValid(t *testing.T) {
	for _, configuration := range countriesConfiguration {

		if err := IsIBANValid(configuration.IBANDefinition.Example); err != nil {
			t.Errorf("Output was incorrect, IsIBANValid(%s)= false, %v; want: true.", configuration.IBANDefinition.Example, err)
		}

		if err := IsIBANValid(configuration.IBANDefinition.PrintFormat); err != nil {
			t.Errorf("Output was incorrect, IsIBANValid(%s)= false, %v; want: true.", configuration.IBANDefinition.PrintFormat, err)
		}
	}
}

// Tests IsIBANValid with custom values
func TestIsIBANValidWithCustomValues(t *testing.T) {
	testCases := []struct {
		input       string
		expectedErr error
	}{
		{"FR1420041010050500013M02606", nil},
		{"fr1420041010050500013M02606", nil},
		{"fR1420041010050500013M02606", nil},
		{"Fr1420041010050500013M02606", nil},
		{"FR1420041010050500013m02606", nil},
		{"FR1420041010050500013M02607", ErrIBANInvalidChecksum},
	}

	for i, testCase := range testCases {

		if err := IsIBANValid(testCase.input); err != testCase.expectedErr {
			t.Errorf("[%d] Output was incorrect, IsIBANValid(%s)= %v; want: %v.", i,
				testCase.input,
				err,
				testCase.expectedErr,
			)
		}
	}
}

// Tests checkIBAN with custom values
func TestCheckIBAN(t *testing.T) {
	testCases := []struct {
		input       string
		expectedErr error
	}{
		{"FR1420041010050500013M02606", nil},
		{"FR14 2004 1010 0505 0001 3M02 606", ErrIBANIncorrectLength},
		{"", ErrIBANTooshort},
		{"fr1420041010050500013M02606", ErrCountryCodeDoesNotExist},
		{"XX1420041010050500013M02606", ErrCountryCodeDoesNotExist},
		{"FR1420041010050500013M0260", ErrIBANIncorrectLength},
	}

	for i, testCase := range testCases {
		_, err := checkIBAN(testCase.input)

		if err != testCase.expectedErr {
			t.Errorf("[%d] Output was incorrect, checkIBAN(%s)= %v, want: %v.", i, testCase.input, err, testCase.expectedErr)
		}
	}
}

// Tests GetIBANCheckDigits with default values from countries configuration
func TestGetIBANCheckDigits(t *testing.T) {
	for _, configuration := range countriesConfiguration {

		check, iban, err := GetIBANCheckDigits(configuration.IBANDefinition.Example)
		if err != nil {
			t.Errorf("Unexpected error occurred: %v", err)
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

// Tests PrintFormatIBAN with default values from countries configuration
func TestPrintFormatIBAN(t *testing.T) {
	for _, configuration := range countriesConfiguration {

		formattedIBAN, err := PrintFormatIBAN(configuration.IBANDefinition.Example)
		if err != nil {
			t.Errorf("Unexpected error occurred: %v", err)
		}

		if formattedIBAN != configuration.IBANDefinition.PrintFormat {
			t.Errorf("Output was incorrect, PrintFormatIBAN(%s)= %v; want: %v.",
				configuration.IBANDefinition.Example,
				formattedIBAN,
				configuration.IBANDefinition.PrintFormat,
			)
		}

	}
}

// Tests SplitIBAN with default values from countries configuration
func TestSplitIBAN(t *testing.T) {
	for _, configuration := range countriesConfiguration {

		cc := configuration.CountryCode.String()

		m := map[string]int{}
		m[AccountNumberIBANDigitKey] = strings.Count(configuration.IBANDefinition.Structure, AccountNumberIBANDigitKey)
		m[NationalBankCodeIBANDigitKey] = strings.Count(configuration.IBANDefinition.Structure, NationalBankCodeIBANDigitKey)
		m[CountryCodeIBANDigitKey] = strings.Count(configuration.IBANDefinition.Structure, CountryCodeIBANDigitKey)
		m[CheckIBANDigitKey] = strings.Count(configuration.IBANDefinition.Structure, CheckIBANDigitKey)
		m[NationalIdentificationNumberIBANDigitKey] = strings.Count(configuration.IBANDefinition.Structure, NationalIdentificationNumberIBANDigitKey)
		m[CurrencyIBANDigitKey] = strings.Count(configuration.IBANDefinition.Structure, CurrencyIBANDigitKey)
		m[AccountHolderIBANDigitKey] = strings.Count(configuration.IBANDefinition.Structure, AccountHolderIBANDigitKey)
		m[ReserveNumberIBANDigitKey] = strings.Count(configuration.IBANDefinition.Structure, ReserveNumberIBANDigitKey)
		m[BranchCodeIBANDigitKey] = strings.Count(configuration.IBANDefinition.Structure, BranchCodeIBANDigitKey)
		m[AccountTypeIBANDigitKey] = strings.Count(configuration.IBANDefinition.Structure, AccountTypeIBANDigitKey)
		m[SWIFTBICCodeIBANDigitKey] = strings.Count(configuration.IBANDefinition.Structure, SWIFTBICCodeIBANDigitKey)
		m[NationalCheckIBANDigitKey] = strings.Count(configuration.IBANDefinition.Structure, NationalCheckIBANDigitKey)

		k, v, err := SplitIBAN(configuration.IBANDefinition.Example)
		if err != nil {
			t.Errorf("[%s] Unexpected error occurred: %v", cc, err)
		}

		if len(k) != len(v) {
			t.Errorf("[%s] Keys and values are not the same length. |keys|= %v; |values|= %v.", cc, len(k), len(v))
		}

		for i := range k {
			if n, ok := m[k[i]]; ok {
				if n != len(v[i]) {
					t.Errorf("[%s] Missing digit in part %v. want: %v; got: %v", cc, k[i], n, len(v[i]))
				}
			}
		}

	}
}

func BenchmarkSplitIBAN(b *testing.B) {
	const iban = "FR1420041010050500013M02606"

	b.Run("SplitIBAN", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_, _, err := SplitIBAN(iban)
			if err != nil {
				b.Errorf("Unexpected error occurred: %v", err)
			}
		}
	})
}

// Examples

func ExampleIsIBANValid() {
	fmt.Println(IsIBANValid("FR1420041010050500013M02606"))
	fmt.Println(IsIBANValid("FR1420041010050500013M02605"))
	// Output:
	// <nil>
	// IBAN invalid checksum
}

func ExampleGetIBANCheckDigits() {
	fmt.Println(GetIBANCheckDigits("FR1420041010050500013M02606"))
	fmt.Println(GetIBANCheckDigits("FR0020041010050500013M02606"))
	// Output:
	// 14 FR1420041010050500013M02606 <nil>
	// 14 FR1420041010050500013M02606 <nil>
}
