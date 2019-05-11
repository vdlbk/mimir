package mimir

import (
	"fmt"
	"testing"
)

// Tests MatchPaymentCard with payment card configurations examples
func TestMatchPaymentCard(t *testing.T) {
	for _, configuration := range paymentCardConfigurations {
		for _, ex := range configuration.Examples {

			t.Run("default", func(t *testing.T) {
				if m, err := MatchPaymentCard(ex); err != nil {
					t.Errorf("[%s] Unexpected error, MatchPaymentCard(%s)= %v, %v.", configuration.MajorIndustryIdentifierName, ex, m, err)
				} else if len(m) == 0 {
					t.Errorf("[%s] Output was incorrect, MatchPaymentCard(%s)= %v, %v; want un-empty match list.", configuration.MajorIndustryIdentifierName, ex, m, err)
				}
			})

			t.Run("match issuer", func(t *testing.T) {
				if m, err := MatchPaymentCard(ex, configuration.MajorIndustryIdentifierName); err != nil {
					t.Errorf("[%s] Unexpected error, MatchPaymentCard(%s, %s)= %v, %v.",
						configuration.MajorIndustryIdentifierName,
						ex,
						configuration.MajorIndustryIdentifierName,
						m,
						err,
					)
				}
			})
		}
	}
}

// Tests MatchPaymentCard with custom values
func TestMatchPaymentCardhWithCustomValues(t *testing.T) {
	testCases := []struct {
		input          string
		issuer         []string
		expectedResult []string
		expectedErr    error
	}{
		{"4012888888881881", []string{}, []string{Visa}, nil},
		{"4012888888881881", []string{}, []string{Visa}, nil},
		{"4012888888881881", []string{Visa}, []string{Visa}, nil},
		{"4012888888881881", []string{Mastercard}, []string{}, ErrPaymentCardDoesNotMatchAnyIssuer},
		{"foobar", []string{}, []string{}, ErrPaymentCardDoesNotMatchAnyIssuer},
		{"", []string{}, []string{}, ErrPaymentCardDoesNotMatchAnyIssuer},
	}

	for i, testCase := range testCases {
		r, err := MatchPaymentCard(testCase.input)

		if err != nil && err != testCase.expectedErr {
			t.Errorf("[%d] Unexpected error, MatchPaymentCard(%s)= %v, %v; want: %v, %v.",
				i,
				testCase.input,
				r,
				err,
				testCase.expectedResult,
				testCase.expectedErr,
			)
		} else if len(r) == 0 && len(testCase.expectedResult) != 0 {
			t.Errorf("[%d] Output was incorrect, MatchPaymentCard(%s)= %v, %v; want: %v, %v.",
				i,
				testCase.input,
				r,
				err,
				testCase.expectedResult,
				testCase.expectedErr,
			)
		}
	}
}

// Tests luhn with payment card configurations examples
func TestLuhn(t *testing.T) {
	for _, configuration := range paymentCardConfigurations {
		for _, ex := range configuration.Examples {
			t.Run("default", func(t *testing.T) {
				if r := luhn(ex); r != 0 {
					t.Errorf("[%s] Output was incorrect, luhn(%s)= %v, want: %v.", configuration.MajorIndustryIdentifierName, ex, r, 0)
				}
			})
		}
	}
}

// Tests luhn with custom values
func TestLuhnWithCustomValues(t *testing.T) {
	testCases := []struct {
		input          string
		expectedResult int64
	}{
		{"4012888888881881", 0},
		{"4012888888881882", 1},
	}

	for i, testCase := range testCases {
		r := luhn(testCase.input)

		if r != testCase.expectedResult {
			t.Errorf("[%d] Output was incorrect, luhn(%s)= %v, want: %v.", i, testCase.input, r, testCase.expectedResult)
		}
	}
}

// Tests GetPaymentCardCheckDigits with payment card configurations examples
func TestGetPaymentCardCheckDigits(t *testing.T) {
	for _, configuration := range paymentCardConfigurations {
		for _, ex := range configuration.Examples {
			t.Run("default", func(t *testing.T) {
				expected := ex[len(ex)-1]
				if checksum, number, err := GetPaymentCardCheckDigits(ex[:len(ex)-1] + "0"); number != ex || err != nil {
					t.Errorf("[%s] Output was incorrect, GetPaymentCardCheckDigits(%s)= %v, %v; want: %v, %v.",
						configuration.MajorIndustryIdentifierName,
						ex[:len(ex)-1]+"0",
						checksum,
						number,
						ex,
						expected,
					)
				}
			})
		}
	}
}

// Tests GetPaymentCardCheckDigits with custom values
func TestGetPaymentCardCheckDigitsWithCustomValues(t *testing.T) {
	testCases := []struct {
		input            string
		expectedChecksum string
		expectedResult   string
		expectedErr      error
	}{
		{"12340", "4", "12344", nil},
		{"12344", "0", "12340", nil},
		{"1234", "6", "1236", nil},
		{"1", "9", "9", nil},
		{"0", "0", "0", nil},
		{"", "", "", ErrPaymentCardTooShort},
	}

	for i, testCase := range testCases {
		checksum, number, err := GetPaymentCardCheckDigits(testCase.input)

		if err != nil && err != testCase.expectedErr {
			t.Errorf("[%d] Unexpected error, GetPaymentCardCheckDigits(%s)= %v; want: %v.", i, testCase.input, err, testCase.expectedErr)
		} else if number != testCase.expectedResult {
			t.Errorf("[%d] Output was incorrect, GetPaymentCardCheckDigits(%s)= %v, %v; want: %v, %v.",
				i,
				testCase.input,
				checksum,
				number,
				testCase.expectedChecksum,
				testCase.expectedResult,
			)
		}
	}
}

// Examples

func ExampleGetPaymentCardCheckDigits() {
	fmt.Println(GetPaymentCardCheckDigits("12340"))
	fmt.Println(GetPaymentCardCheckDigits("1234"))
	// Output:
	// 4 12344 <nil>
	// 6 1236 <nil>
}

func ExampleMatchPaymentCard() {
	fmt.Println(MatchPaymentCard("4012888888881881"))
	fmt.Println(MatchPaymentCard("4012888888881881", Visa))
	fmt.Println(MatchPaymentCard("4012888888881881", Mastercard))
	fmt.Println(MatchPaymentCard("30037174022893"))
	// Output:
	// [Visa] <nil>
	// [Visa] <nil>
	// [] Payment card does not match any issuer
	// [Dinner Club International Dinner Club Carte Blanche] <nil>
}
