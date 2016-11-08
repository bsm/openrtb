# Go OpenRTB v2.x

[![Build Status](https://travis-ci.org/bsm/openrtb.svg?branch=master)](https://travis-ci.org/bsm/openrtb)

OpenRTB implementation for Go

## Installation

To install, use `go get`:

```shell
go get github.com/bsm/openrtb
```

## Usage

Import the package:

```go
package main

import (
  "log"
  "github.com/bsm/openrtb"
)

func main() {
  file, err := os.Open("stored.json")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  var req *openrtb.BidRequest
  err = json.NewDecoder(file).Decode(&req)
  if err != nil {
    log.Fatal(err)
  }

  log.Printf("%+v\n", req)
}
```

## Licence

    Copyright (c) 2015 Black Square Media Ltd. All rights reserved.
    (The MIT License)

    Permission is hereby granted, free of charge, to any person obtaining
    a copy of this software and associated documentation files (the
    'Software'), to deal in the Software without restriction, including
    without limitation the rights to use, copy, modify, merge, publish,
    distribute, sublicense, and/or sell copies of the Software, and to
    permit persons to whom the Software is furnished to do so, subject to
    the following conditions:

    The above copyright notice and this permission notice shall be
    included in all copies or substantial portions of the Software.

    THE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND,
    EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
    MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
    IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
    CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
    TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
    SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

    Some test examples were taken from:
    https://code.google.com/p/openrtb/wiki/OpenRTB_Examples
