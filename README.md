# Mimir &nbsp;[![Build Status](https://cloud.drone.io/api/badges/vdlbk/mimir/status.svg)](https://cloud.drone.io/vdlbk/mimir) [![Go Report Card](https://goreportcard.com/badge/github.com/vdlbk/mimir)](https://goreportcard.com/report/github.com/vdlbk/mimir)  [![GoDoc](https://godoc.org/github.com/vdlbk/mimir?status.svg)](https://godoc.org/github.com/vdlbk/mimir)

## What is Mimir ?

**Mimir** is a lightweight library that exposes some functions to validate banking informations.    
It does not store any of your personal data, you can use it freely.

### Disclaimer

This library is still in development, so it might be possible that some methods change over time. 

## Methods available

If you experienced troubles with one of our methods, don't hesitate telling us by submitting an issue.

### Bank account numbers

#### IBAN Methods 

| Method Name | Short Description |
| ----------- | ----------------- |
| `IsIBANValid` | Validate an IBAN  |
| `GetIBANCheckDigits` | Compute or fix check digits from an IBAN  |
| `FormatIBAN` | Format IBAN based on its structure  |
| `PrintFormatIBAN` | Format IBAN as it would be printed  |
| `GetCountryConfiguration` | Get country configuration and informations  |
| `SplitIBAN` | Split the IBAN based on its structure  |

#### ABA Routing Number Methods 

| Method Name | Short Description |
| ----------- | ----------------- |
| `IsABARTNValid` | Validate an ABA  |
| `GetABARTNCheckDigit` | Compute or fix check digits for an ABA  |
| `SplitABARTN` | Split the ABA based on its structure  |

#### Supported countries

Please visit [SUPPORTED_COUNTRIES.md](https://github.com/vdlbk/mimir/blob/master/SUPPORTED_COUNTRIES.md)

### Payment card numbers

All cards in the examples has been generated randomly

#### Methods
| Method Name | Short Description |
| ----------- | ----------------- |
| `MatchPaymentCard` | Validate a payment card number from a specific card issuer or not  |
| `GetPaymentCardCheckDigits` | Compute or fix check digits from a payment card  |
| `GetPaymentCardConfiguration` | Get payment cards information from an issuer  |

#### Supported cards

Please visit [SUPPORTED_CARDS.md](https://github.com/vdlbk/mimir/blob/master/SUPPORTED_CARDS.md)

## Use

### Get
```bash
$ go get -u github.com/vdlbk/mimir
```

### Examples
#### Validate an IBAN
```go
package main

import (
	"fmt"
	"github.com/vdlbk/mimir"
)

func main(){
	IBAN := "FR1420041010050500013M02606"
	
	err := mimir.IsIBANValid(IBAN)
	fmt.Println(err) // nil
	
	badIBAN := "FR1420041010050500013M02605"
	
	err = mimir.IsIBANValid(badIBAN)
	fmt.Println(err) // IBAN invalid checksum
}
```

#### Validate a payment card number
```go
package main

import (
	"fmt"
	"github.com/vdlbk/mimir"
)

func main(){
	number := "345752218692713"
	
	r, err := mimir.MatchPaymentCard(number, mimir.AmericanExpress)
	fmt.Println(r, err) // [American Express], nil
	
	badNumber := "foobar"
	
	r, err = mimir.MatchPaymentCard(badNumber, mimir.AmericanExpress)
	fmt.Println(r, err) // nil, Payment card does not match any issuer
}
```