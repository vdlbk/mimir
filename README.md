# Mimir &nbsp;[![Build Status](https://cloud.drone.io/api/badges/vdlbk/mimir/status.svg)](https://cloud.drone.io/vdlbk/mimir) [![GoDoc](https://godoc.org/github.com/vdlbk/mimir?status.svg)](https://godoc.org/github.com/vdlbk/mimir)

## What is Mimir ?

**Mimir** is a lightweight library that exposes some functions to validate some banking informations.

### Disclaimer

This library is still in development, so it might be possible that some methods change over time. 

## Methods available

As the project just started, it is quite poor in method.    
Stay tuned for updates !

+  Validate an IBAN (every country supported)
+  Compute check digits from a valid IBAN
+  Get country configuration (ISO 13616-Compliant IBAN Formats)

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