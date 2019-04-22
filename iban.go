package mimir

import (
	"math/big"
	"strings"
)

const (
	minimumIBANLength       = 2
	numberOfDigitRearranged = 4
)

// IsIBANValid checks if the given IBAN is valid based on the country its belong to
func IsIBANValid(iban string) (bool, error) {
	// the given IBAN's length cannot be less than `minimumIBANLength`
	if len(iban) < minimumIBANLength {
		return false, ErrIBANTooshort
	}

	iban = strings.ToUpper(iban)

	var conf configuration
	var ok bool

	// Extracts country code from the first 2 characters of the IBAN and if exists, get the country configuration associated with it.
	if conf, ok = countriesConfiguration[iban[:2]]; !ok {
		return false, ErrCountryCodeDoesNotExist
	}

	// Checks the length of the IBAN
	if len(iban) != conf.IBANDefinition.Length {
		return false, ErrIBANIncorrectLength
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
