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
