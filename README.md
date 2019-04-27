# Mimir &nbsp;[![Build Status](https://cloud.drone.io/api/badges/vdlbk/mimir/status.svg)](https://cloud.drone.io/vdlbk/mimir) [![GoDoc](https://godoc.org/github.com/vdlbk/mimir?status.svg)](https://godoc.org/github.com/vdlbk/mimir)

## What is Mimir ?

**Mimir** is a lightweight library that exposes some functions to validate some banking informations.

### Disclaimer

This library is still in development, so it might be possible that some methods change over time. 

## Methods available

As the project just started, it is quite poor in methods.    
Stay tuned for updates !

| Method Name | Short Description |
| ----------- | ----------------- |
| `IsIBANValid` | Validate an IBAN  |
| `GetCheckDigits` | Compute check digits from a valid IBAN  |
| `FormatIBAN` | Format IBAN based on its structure  |
| `PrintFormatIBAN` | Format IBAN as it would be printed  |
| `GetCountryConfiguration` | Get country configuration and informations  |
| `SplitIBAN` | Split the IBAN based on its structure  |

### Supported countries

Please visit [SUPPORTED_COUNTRIES.md](https://github.com/vdlbk/mimir/blob/master/SUPPORTED_COUNTRIES.md)

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