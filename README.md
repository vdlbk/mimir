# Mimir &nbsp;[![Build Status](https://cloud.drone.io/api/badges/vdlbk/mimir/status.svg)](https://cloud.drone.io/vdlbk/mimir) [![GoDoc](https://godoc.org/github.com/vdlbk/mimir?status.svg)](https://godoc.org/github.com/vdlbk/mimir)

## What is Mimir ?

**Mimir** is a lightweight library that exposes some functions to validate some banking informations.

## Validation

As the project just started, you can just validate an IBAN.    
Stay tuned for updates !

## Use

### Install
```bash
$ go get -u github.com/vdlbk/mimir
```

### Example : Validate an IBAN
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
	
	badIBAN := "FR1420041010050500013M02606"
    	
    ok, _ = mimir.IsIBANValid(badIBAN)
    fmt.Println(ok) // false
}
```