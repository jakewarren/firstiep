# firstiep
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/jakewarren/firstiep)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](https://github.com/jakewarren/firstiep/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/jakewarren/firstiep)](https://goreportcard.com/report/github.com/jakewarren/firstiep)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=shields)](http://makeapullrequest.com)

> Tools for working with FIRST IEP (Information Exchange Policy) objects - https://www.first.org/iep/

## Demo
![](demo.gif)

## Install

```
go get github.com/jakewarren/firstiep/...
```

## Usage
### As a library

```go
package main

import (
	"fmt"

	"github.com/jakewarren/firstiep"
)

func main() {

	// create a new IEP object
	i := firstiep.New()
	
	// fill in fields you want represented
	i.TLP = "RED"

	// output the object for inspection
	fmt.Println(i)

	// perform validation on the object
	if err := i.Validate(); err == nil {
		fmt.Println("Valid!")
	} else {
		fmt.Println("Not valid:", err)
	}

}
```

### As a cmd line utility

On the command line, run `firstiep` and follow the wizard to select the options you want, the tool will then output the JSON representation.
