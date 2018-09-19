# cld2 nlpt

DEPRECATED for whatlanggo: [https://github.com/abadojack/whatlanggo](https://github.com/abadojack/whatlanggo)


This is a separate fork of [rainycape's wrapper of cld2](https://github.com/rainycape/cld2).

From the README of rainycape wrapper of cld2:

```sh
Go wrapper for the cld2 language detection library by Google Chrome.

Package cld2 implements language detection using the Compact Language Detector.

This package includes the relevant sources from the cld2 project, so it doesn't
require any external dependencies. For more information about CLD2, see
https://code.google.com/p/cld2/.
```

The `nlpt` part is a side project of mine for a Natural Language Processing Toolkit in go.

## External Sources
This wrapper owes its existence to 3 projects:

* cld2 project -- original code [cld2](https://code.google.com/p/cld2/)
* cld2 go wrapper -- [cld2](https://github.com/rainycape/cld2)
* rust-cld2 wrapper [rust-cld2](https://github.com/emk/rust-cld2)

I'm not very good at C/C++ so I leaved heavily on wrapper projects in Go and Rust mentioned.


## Get it

```sh
go get github.com/jbowles/nlpt-cld2
```

## Using it

```go
package main

import (
	"fmt"

	"github.com/jbowles/nlpt-cld2"
)

var s1 = "this sentence is in english dooode"

func this() {
}

func main() {
	res, err := cld2.SimpleDetect(s1)
	fmt.Printf("%v %v \n", res, err)
}
```
