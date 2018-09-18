# envy

<p align="center">
  <img style="float: right;" src="assets/envy.png" alt="envy gopher"/>
</p>

A simple package to get environment variables, returns a string or an error. [Go Buffalo](https://github.com/gobuffalo/envy) already had a ENV package named envy it seems. Naming stuff is hard. Oh well.

## Installation

Use `go get github.com/shindakun/envy` to install into Go or use as a Go module.

## Usage

Import `github.com/shindakun/envy`, see full example below.

## Usage example

```go
package main

import (
  "fmt"
  "github.com/shindakun/envy"
)

func main() {
  gopath, err := envy.Get("GOPATH")
  if err != nil {
    panic(err)
  }
  fmt.Println(gopath)
}
```
