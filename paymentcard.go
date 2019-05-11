package mimir

import (
	"regexp"
	"strconv"
	"strings"
)

// Major industry identifier name / Issuer keys
const (
	AmericanExpress         = "American Express"
	ChinaUnionPay           = "China UnionPay"
	DinnerClubInternational = "Dinner Club International"
	DinnerClubCarteBlanche  = "Dinner Club Carte Blanche"
	DinnerClub              = "Dinner Club"
	Discover                = "Discover"
	Rupay                   = "Rupay"
	InterPayment            = "InterPayment"
	InstaPayment            = "InstaPayment"
	JCB                     = "JCB"
	Maestro                 = "Maestro"
	Dankort                 = "Dankort"
	MIR                     = "MIR"
	Mastercard              = "Mastercard"
	Visa                    = "Visa"
	UATP                    = "UATP"
)

// MatchPaymentCard returns a list of the issuers that match the given number. If the issuers list has been filled with valid issuer keys,
// it will only check the number for those ones. So it can be useful when it comes to check if a payment card number belongs to a specific issuer.
// It can also return an error if something goes wrong or if the number does not match any known issuer.
func MatchPaymentCard(number string, issuers ...string) ([]string, error) {
	number = sanitizeCardNumber(number)

	configs := paymentCardConfigurations
	if len(issuers) > 0 {
		configs = make(map[string]paymentCardConfiguration, len(issuers))
		for _, i := range issuers {
			if conf, ok := paymentCardConfigurations[i]; ok {
				configs[i] = conf
			}
		}
	}

	var matches []string
	for k, v := range configs {
		rgx, err := regexp.Compile(v.Pattern)
		if err != nil {
			return nil, err
		}
		if rgx.MatchString(number) {
			if r := luhn(number); r != 0 {
				// The number does not match the current card pattern, so skip to the next one.
				continue
			}
			matches = append(matches, k)
		}
	}

	if len(matches) == 0 {
		return nil, ErrPaymentCardDoesNotMatchAnyIssuer
	}

	return matches, nil
}

// GetPaymentCardCheckDigits compute the check digits from a given payment card number.
// If you need to compute the check digit for 1234, please add an extra 0 at the end, otherwise, it's going to compute the
// check digit for 123. The last digit is ignored.
// Returns the computed digit, the payment card number set with the check digit or an error if something goes wrong.
func GetPaymentCardCheckDigits(number string) (string, string, error) {
	if len(number) == 0 {
		return "", "", ErrPaymentCardTooShort
	}
	r := luhn(number)
	if r == 0 {
		return "0", number[:len(number)-1] + "0", nil
	}
	checksum := 10 - r
	return strconv.FormatInt(checksum, 10), number[:len(number)-1] + strconv.FormatInt(checksum, 10), nil
}

func sanitizeCardNumber(number string) string {
	// Remove all spaces
	return strings.Replace(number, " ", "", -1)
}

// luhn algorithm implementation
func luhn(number string) int64 {
	var sum int64
	var alt bool

	for pos := len(number) - 1; pos > -1; pos-- {
		num := int64(number[pos] - 0x30)
		if alt {
			num *= 2
			if num > 9 {
				num = (num % 10) + 1
			}
		}
		alt = !alt
		sum += num
	}

	return sum % 10
}
