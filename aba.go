package mimir

import (
	"strconv"
	"strings"
)

const (
	abaExpectedLength = 9
	abaStructure      = "ffffaaaak"
)

// IsABARTNValid checks if the given ABA Routing Number is valid.
func IsABARTNValid(aba string) error {
	aba = standardizeABARTN(aba)

	if len(aba) != abaExpectedLength {
		return ErrABARTNInvalidLength
	}

	n := computeABARTN(aba)

	if n == 0 || n%10 != 0 {
		return ErrABAInvalidChecksum
	}

	return nil
}

// GetABARTNCheckDigit compute the check digits from a given ABA Routing Number.
// Returns the computed digit and the ABA updated with the check digits or an error if something goes wrong.
// The ABA must contains 9 digits but the last one will ignored. It's possible to use this function by adding an extra 0
// at the end if the ABA is not a valid one.
func GetABARTNCheckDigit(aba string) (string, string, error) {
	aba = standardizeABARTN(aba)

	if len(aba) != abaExpectedLength {
		return "", "", ErrABARTNInvalidLength
	}

	n := computeABARTN(aba)

	// Removes the last digit from the computation
	n -= int(aba[len(aba)-1] - '0')

	checksum := 0
	if mod := n % 10; mod > checksum {
		checksum = 10 - mod
	}
	sChecksum := strconv.Itoa(checksum)

	aba = aba[:abaExpectedLength-1] + sChecksum
	return sChecksum, aba, nil
}

// SplitABARTN splits a given valid ABA Routing Number depending its structure. It returns a list of digit keys and a list
// that contains each part of the structure or an error if something went wrong.
func SplitABARTN(aba string) ([]string, []string, error) {
	aba = standardizeABARTN(aba)

	if len(aba) != abaExpectedLength {
		return nil, nil, ErrABARTNInvalidLength
	}

	s := abaStructure
	var keys []string
	var values []string
	lastRune := ""
	currentPart := ""
	for i, r := range aba {
		y := string(s[i])
		if lastRune != "" && lastRune != y {
			keys = append(keys, string(lastRune))
			values = append(values, currentPart)
			currentPart = ""
		}
		currentPart += string(r)

		lastRune = y
	}

	if lastRune != "" {
		keys = append(keys, string(lastRune))
		values = append(values, currentPart)
	}

	return keys, values, nil
}

// standardizeABARTN removes useless spaces from ABA Routing Number.
func standardizeABARTN(aba string) string {
	// Remove all spaces
	return strings.Replace(aba, " ", "", -1)
}

// computeABARTN expects an 9 length string that contains digit from 0-9.
// It going to sum them using the following formula according Wikipedia : https://en.wikipedia.org/wiki/ABA_routing_transit_number#Check_digit
// [3(d1 + d4 + d7) + 7(d2 + d5 + d8) + (d3 + d6 + d9)]
func computeABARTN(aba string) int {
	n := 0
	for i := 0; i < len(aba); i += 3 {
		n += int(aba[i]-'0') * 3
		n += int(aba[i+1]-'0') * 7
		n += int(aba[i+2] - '0')
	}
	return n
}
