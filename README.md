# cld2 nlpt
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
See tests for full usage. This package consists of 5 public functions only.

The function `Detect` is the preferred way of using this package, but it requires many arguments that depend on user familiarity with the cld2 project. It uses the full set of options for cld2, including extended language detection. It will eventually support passing a struct of language hints.

If user has no familiarty with cld2 or doesn't want to be bothered with complex usage then `StaticDetect` is for you: it uses the extended language feature and pre-defines all the options to cld2; it requires 1 argument: text to be identified.

The 3 remaining public functions can be used to return language code, name, or displayed name; they do not use extended language features and use the most basic options in cld2.

In terms of accuracy and reliability `Detect` and `StaticDetect` are most reliable.

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
