# OpenRTB

[![Build Status](https://travis-ci.org/UnityTech/openrtb.svg?branch=master)](https://travis-ci.org/UnityTech/openrtb)

OpenRTB structs and validations for Go.

## Requirements

Requires Go 1.8+ for proper `json.RawMessage` marshaling.

## Installation

To install, use `go get`:

```shell
go get github.com/UnityTech/openrtb
```

## Usage

```go
package main

import (
  "log"

  "github.com/UnityTech/openrtb"
)

func main() {
  file, err := os.Open("stored.json")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  var req *openrtb.BidRequest
  if err := json.NewDecoder(file).Decode(&req); err != nil {
    log.Fatal(err)
  }

  log.Printf("%+v\n", req)
}
```
