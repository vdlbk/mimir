package iban

import (
	"fmt"
	"github.com/vdlbk/mimir/models"
	"github.com/vdlbk/mimir/utils/consts"
	"math/big"
	"strings"
)

const (
	NumberOfDigitRearranged = 4
)

func IsIBANValid(iban string) (bool, error) {
	if len(iban) < 2 {
		return false, nil // TODO: iban too short
	}

	iban = strings.ToUpper(iban)

	countryCode := iban[:2]

	var conf models.Configuration
	var ok bool
	if conf, ok = consts.CountriesConfiguration[countryCode]; !ok {
		return false, nil // TODO: IBAN country code does not exists
	}

	if len(iban) != conf.IBANDefinition.Length {
		return false, nil // TODO: Length is not good
	}

	rearrangedIBAN := iban[NumberOfDigitRearranged:] + iban[:NumberOfDigitRearranged]

	sanitizedIBAN := ""
	for _, c := range rearrangedIBAN {
		sanitizedIBAN += consts.Alphabet[c]
	}

	x := new(big.Int)
	r := new(big.Int)
	x.SetString(sanitizedIBAN, 10)
	r.Mod(x, big.NewInt(97))

	fmt.Println(iban, sanitizedIBAN, r.Int64())

	ok = r.Int64() == 1
	var err error
	if !ok {
		err = nil // TODO: Invalid check sum
	}

	return ok, err
}
