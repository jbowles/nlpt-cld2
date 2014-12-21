# cld2 nlpt
Go wrapper for the cld2 language detection library by Google Chrome.

## External Sources
This wrapper owes its existence to 3 projects:

* cld2 project -- original code
* cld2 go wrapper
* rust-cld2 wrapper

I'm not very good at C/C++ so I leaved heavily on wrapper projects in Go and Rust mentioned.


## Get it

```sh
go get github.com/jbowles/cld2_nlpt
```

## Using it

```go
package main

import "github.com/jbowles/cld2_nlpt"

func main() {

  string := "This is an english sentence"
  buffer_length := len(string)
  det := DetectLanguage(s, buffer_length)

```

## Misc.
The first version used the original go wrapper

```go
// Package cld2 implements language detection using the
// Compact Language Detector.
//
// This package includes the relevant sources from the cld2
// project, so it doesn't require any external dependencies.
// For more information about CLD2, see https://code.google.com/p/cld2/.

package cld2_nlpt

// #include <stdlib.h>
// #include "cld2_min/cld2.h"
import "C"
import "unsafe"

// Detect returns the language code for detected language
// in the given text.
func Detect(text string) (lang string) {
	cs := C.CString(text)
	res := C.DetectLang(cs, -1)
	defer C.free(unsafe.Pointer(cs))
	if res != nil {
		lang = C.GoString(res)
	}
	return
}
```
