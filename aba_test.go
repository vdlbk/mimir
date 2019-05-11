package mimir

import (
	"fmt"
	"reflect"
	"testing"
)

// Tests IsABARTNValid with custom values
func TestIsABARTNValid(t *testing.T) {
	testCases := []struct {
		input       string
		expectedErr error
	}{
		{"111000025", nil},
		{"133563582", nil},
		{"021302567", nil},
		{"053902197", nil},
		{"036001808", nil},
		{"011600033", nil},
		{"765651956", nil},
		{"172193081", nil},
		{"510502124", nil},
		{"276612918", nil},
		{"663003088", nil},
		{"621243110", nil},
		{"383317096", nil},
		{"584513178", nil},
		{"326621435", nil},
		{"239932004", nil},
		{"157221130", nil},
		{"254831711", nil},
		{"925855136", nil},
		{"482770518", nil},
		{"789456124", nil},
		{"7 8 9 4 5 6 1 2 4", nil},
		{"13356358", ErrABARTNInvalidLength},
		{"1335635800", ErrABARTNInvalidLength},
		{"133563586", ErrABAInvalidChecksum},
	}

	for i, testCase := range testCases {
		err := IsABARTNValid(testCase.input)

		if err != testCase.expectedErr {
			t.Errorf("[%d] Output was incorrect, IsABARTNValid(%s)= %v, want: %v.", i, testCase.input, err, testCase.expectedErr)
		}
	}
}

func TestGetABARTNCheckDigit(t *testing.T) {
	testCases := []struct {
		input              string
		expectedCheckdigit string
		expectedABARTN     string
		expectedErr        error
	}{
		{"111000020", "5", "111000025", nil},
		{"133563580", "2", "133563582", nil},
		{"021302560", "7", "021302567", nil},
		{"053902190", "7", "053902197", nil},
		{"036001800", "8", "036001808", nil},
		{"011600030", "3", "011600033", nil},
		{"765651950", "6", "765651956", nil},
		{"172193080", "1", "172193081", nil},
		{"510502120", "4", "510502124", nil},
		{"276612910", "8", "276612918", nil},
		{"663003080", "8", "663003088", nil},
		{"621243110", "0", "621243110", nil},
		{"383317090", "6", "383317096", nil},
		{"584513170", "8", "584513178", nil},
		{"326621430", "5", "326621435", nil},
		{"239932000", "4", "239932004", nil},
		{"157221130", "0", "157221130", nil},
		{"254831710", "1", "254831711", nil},
		{"925855130", "6", "925855136", nil},
		{"482770510", "8", "482770518", nil},
		{"789456120", "4", "789456124", nil},
		{"7 8 9 4 5 6 1 2 0", "4", "789456124", nil},
		{"133563586", "2", "133563582", nil},
		{"13356358", "", "", ErrABARTNInvalidLength},
		{"1335635800", "", "", ErrABARTNInvalidLength},
	}

	for i, testCase := range testCases {
		c, verifiedABARTN, err := GetABARTNCheckDigit(testCase.input)
		if err != testCase.expectedErr {
			t.Errorf("[%d] Output was incorrect, IsABARTNValid(%s)= %v, want: %v.", i, testCase.input, err, testCase.expectedErr)
		}

		if c != testCase.expectedCheckdigit || verifiedABARTN != testCase.expectedABARTN {
			t.Errorf("[%d] Output was incorrect, IsABARTNValid(%s)= %v, %v, want: %v, %v.", i,
				testCase.input,
				c,
				verifiedABARTN,
				testCase.expectedCheckdigit,
				testCase.expectedABARTN,
			)
		}
	}
}

func TestSplitABARTN(t *testing.T) {
	testCases := []struct {
		input          string
		expectedKeys   []string
		expectedValues []string
		expectedErr    error
	}{
		{"111000025", []string{"f", "a", "k"}, []string{"1110", "0002", "5"}, nil},
		{"133563582", []string{"f", "a", "k"}, []string{"1335", "6358", "2"}, nil},
		{"789456124", []string{"f", "a", "k"}, []string{"7894", "5612", "4"}, nil},
		{"7 8 9 4 5 6 1 2 4", []string{"f", "a", "k"}, []string{"7894", "5612", "4"}, nil},
		{"13356358", nil, nil, ErrABARTNInvalidLength},
		{"1335635800", nil, nil, ErrABARTNInvalidLength},
	}

	for i, testCase := range testCases {
		keys, values, err := SplitABARTN(testCase.input)
		if err != testCase.expectedErr {
			t.Errorf("[%d] Output was incorrect, SplitABARTN(%s)= %v, want: %v.", i, testCase.input, err, testCase.expectedErr)
		}

		if !reflect.DeepEqual(keys, testCase.expectedKeys) || !reflect.DeepEqual(values, testCase.expectedValues) {
			t.Errorf("[%d] Output was incorrect, IsABARTNValid(%s)= %v, %v, want: %v, %v.", i,
				testCase.input,
				keys,
				values,
				testCase.expectedKeys,
				testCase.expectedValues,
			)
		}
	}
}

// Examples

func ExampleGetABARTNCheckDigit() {
	fmt.Println(GetABARTNCheckDigit("111000020"))
	fmt.Println(GetABARTNCheckDigit("111000025"))
	// Output:
	// 5 111000025 <nil>
	// 5 111000025 <nil>
}
