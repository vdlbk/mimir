package mimir

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

const (
	minimumIBANLength       = 2
	numberOfDigitRearranged = 4
)

// IsIBANValid checks if the given IBAN is valid based on the country its belong to
func IsIBANValid(iban string) (bool, error) {
	err := checkIBAN(iban)
	if err != nil {
		return false, err
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
		return false, ErrIBANInvalidChecksum
	}

	return true, nil
}

// GetCheckDigits compute the check digits from a given IBAN.
// Returns the computed digits, the IBAN set with the check digits or an error if something goes wrong
func GetCheckDigits(iban string) (string, string, error) {
	err := checkIBAN(iban)
	if err != nil {
		return "", "", err
	}

	// Replace the 2 checks digit by 00
	iban = iban[:2] + "00" + iban[4:]

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

	// Replace the 2 checks digit by the result
	iban = iban[:2] + fmt.Sprintf("%02d", check) + iban[4:]


	return strconv.Itoa(int(check)), iban, nil
}

func checkIBAN(iban string) error {
	// the given IBAN's length cannot be less than `minimumIBANLength`
	if len(iban) < minimumIBANLength {
		return ErrIBANTooshort
	}

	iban = strings.ToUpper(iban)

	var conf configuration
	var ok bool

	// Extracts country code from the first 2 characters of the IBAN and if exists, get the country configuration associated with it.
	if conf, ok = countriesConfiguration[iban[:2]]; !ok {
		return ErrCountryCodeDoesNotExist
	}

	// Checks the length of the IBAN
	if len(iban) != conf.IBANDefinition.Length {
		return ErrIBANIncorrectLength
	}
	return nil
}