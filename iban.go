package mimir

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

const (
	checkDigitSize          = 2
	minimumIBANLength       = countryCodeSize
	numberOfDigitRearranged = 4
)

// IsIBANValid checks if the given IBAN is valid based on the country its belong to.
func IsIBANValid(iban string) error {
	iban = standardizeIBAN(iban)

	_, err := checkIBAN(iban)
	if err != nil {
		return err
	}

	// Move the `numberOfDigitRearranged` first characters to the end
	rearrangedIBAN := iban[numberOfDigitRearranged:] + iban[:numberOfDigitRearranged]

	// Replace all possible letters by their integer value
	sanitizedIBAN := ""
	for _, c := range rearrangedIBAN {
		sanitizedIBAN += alphabet[c]
	}

	iSanitizedIBAN := new(big.Int)
	mod := new(big.Int)
	iSanitizedIBAN.SetString(sanitizedIBAN, 10)
	mod.Mod(iSanitizedIBAN, big.NewInt(97))

	if mod.Int64() != 1 {
		return ErrIBANInvalidChecksum
	}

	return nil
}

// GetCheckDigits compute the check digits from a given IBAN.
// Returns the computed digits, the IBAN set with the check digits or an error if something goes wrong.
func GetCheckDigits(iban string) (string, string, error) {
	iban = standardizeIBAN(iban)

	_, err := checkIBAN(iban)
	if err != nil {
		return "", "", err
	}

	// Replace the checks digit by checkDigitSize x '0'
	iban = iban[:countryCodeSize] + strings.Repeat("0", checkDigitSize) + iban[countryCodeSize+2:]

	// Move the `numberOfDigitRearranged` first characters to the end
	rearrangedIBAN := iban[numberOfDigitRearranged:] + iban[:numberOfDigitRearranged]

	// Replace all possible letters by their integer value
	sanitizedIBAN := ""
	for _, c := range rearrangedIBAN {
		sanitizedIBAN += alphabet[c]
	}

	iSanitizedIBAN := new(big.Int)
	mod := new(big.Int)
	iSanitizedIBAN.SetString(sanitizedIBAN, 10)
	mod.Mod(iSanitizedIBAN, big.NewInt(97))

	check := 98 - mod.Int64()

	// Replace the checks digit by the result
	iban = iban[:countryCodeSize] + fmt.Sprintf("%02d", check) + iban[countryCodeSize+checkDigitSize:]

	return strconv.Itoa(int(check)), iban, nil
}

// checkIBAN runs some control check-ups like length or country codes and will returns the country configuration associated
// with the IBAN if everything's fine. Otherwise, it will returns an error.
func checkIBAN(iban string) (*configuration, error) {
	// the given IBAN's length cannot be less than `minimumIBANLength`
	if len(iban) < minimumIBANLength {
		return nil, ErrIBANTooshort
	}

	var conf configuration
	var ok bool

	// Extracts country code from the first `countryCodeSize` characters of the IBAN and if exists, get the country configuration associated with it.
	if conf, ok = countriesConfiguration[iban[:countryCodeSize]]; !ok {
		return nil, ErrCountryCodeDoesNotExist
	}

	// Checks the length of the IBAN
	if len(iban) != conf.IBANDefinition.Length {
		return nil, ErrIBANIncorrectLength
	}
	return &conf, nil
}

// standardizeIBAN removes useless spaces from IBAN and capitalizes it too.
func standardizeIBAN(iban string) string {
	// Remove all spaces
	iban = strings.Replace(iban, " ", "", -1)

	// Sets the IBAN in uppercase
	iban = strings.ToUpper(iban)
	return iban
}

// PrintFormatIBAN formats the given IBAN based on the regular way to print it depending the country.
// It will use the PrintFormat of the configuration to insure the result.
func PrintFormatIBAN(iban string) (string, error) {
	iban = standardizeIBAN(iban)

	conf, err := checkIBAN(iban)
	if err != nil {
		return "", err
	}

	// We use the PrintFormat to get the print pattern
	wantedParts := strings.Split(conf.IBANDefinition.PrintFormat, " ")

	parts := make([]string, 0, len(wantedParts))
	for _, p := range wantedParts {
		partLength := len(p)
		if partLength > len(iban) {
			partLength = len(iban)
		}
		prefix := iban[:partLength]
		parts = append(parts, prefix)
		iban = strings.TrimPrefix(iban, prefix)
	}
	return strings.Join(parts, " "), nil
}

// SplitIBAN splits a given IBAN depending the structure of its country. It returns a list of digit keys, a list
// that contains each part of the structure or an error if something went wrong.
func SplitIBAN(iban string) ([]string, []string, error) {
	iban = standardizeIBAN(iban)

	conf, err := checkIBAN(iban)
	if err != nil {
		return nil, nil, err
	}

	s := conf.IBANDefinition.Structure
	var keys []string
	var values []string
	lastRune := ""
	currentPart := ""
	for i, r := range iban {
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

// FormatIBAN formats the given IBAN with the different parts that are defined in its structure.
func FormatIBAN(iban string) (string, error) {
	iban = standardizeIBAN(iban)

	_, err := checkIBAN(iban)
	if err != nil {
		return "", err
	}

	_, v, err := SplitIBAN(iban)
	if err != nil {
		return "", err
	}

	return strings.Join(v, " "), nil
}