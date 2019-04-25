# Mimir &nbsp;[![Build Status](https://cloud.drone.io/api/badges/vdlbk/mimir/status.svg)](https://cloud.drone.io/vdlbk/mimir) [![GoDoc](https://godoc.org/github.com/vdlbk/mimir?status.svg)](https://godoc.org/github.com/vdlbk/mimir)

## What is Mimir ?

**Mimir** is a lightweight library that exposes some functions to validate some banking informations.

### Disclaimer

This library is still in development, so it might be possible that some methods change over time. 

## Methods available

As the project just started, it is quite poor in method.    
Stay tuned for updates !

| Method | Short Description |
| ------ | ----------------- |
| `IsIBANValid(string) (bool, error)` | Validate an IBAN  |
| `GetCheckDigits(string) (string, string, error)` | Compute check digits from a valid IBAN  |
| `PrintFormatIBAN(string) (string, error)` | Format IBAN as it would be printed  |
| `GetCountryConfiguration(string) (*configuration, error)` | Get country configuration (ISO 13616-Compliant IBAN Formats)  |

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
	
	ok, _ := mimir.IsIBANValid(IBAN)
	fmt.Println(ok) // true
	
	badIBAN := "FR1420041010050500013M02605"
    	
    ok, _ = mimir.IsIBANValid(badIBAN)
    fmt.Println(ok) // false
}
```