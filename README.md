<!-- 
MIT License

Copyright (c) 2023 kerelape

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE. 
-->
# Convenient JSON

Convenient JSON (cjson) is a highly inspired by [eo-json](https://github.com/objectionary/eo-json), object-oriented JSON library for Go (golang).

[![EO principles respected here](https://www.elegantobjects.org/badge.svg)](https://www.elegantobjects.org)

[![Build](https://github.com/kerelape/cjson/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/kerelape/cjson/actions/workflows/build.yml)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](https://github.com/kerelape/cjson/blob/main/LICENSE.txt)

## How to use this library

### Add `cjson` to your Go module

```bash
$ go get -u github.com/kerelape/cjson
```

### Build JSON

```go
package main

import (
    "encoding/json"

    cjson "github.com/kerelape/cjson/pkg"
)

func main() {
    object := cjson.NewObject().With("key", cjson.NewString("value"))
    bytes, err := json.Marshal(object)
    if err != nil {
        panic(err)
    }
    println(string(bytes))
}
```

### Parse JSON

```go
package main

import (
    "encoding/json"

    cjson "github.com/kerelape/cjson/pkg"
)

func main() {
    document := cjson.NewDocument()
    err := json.Unmarshal([]byte(`{"a":"b"}`), document)
    if err != nil {
        panic(err)
    }
    println(document.Root().Object().Value().Found("a").Value().String().Value().Content())
}
```

## How to Contribute

Fork this repository, make changes, send us a pull request. We will review your changes and apply them to the `main` branch shortly, provided they don't violate our quality standards. To avoid frustration, before sending us your pull request please run the tests:

```bash
$ go test -v ./...
```

You will need Go 1.20+.
