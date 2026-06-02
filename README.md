# go-spdx

[![Go Reference](https://pkg.go.dev/badge/github.com/shyim/go-spdx.svg)](https://pkg.go.dev/github.com/shyim/go-spdx)
[![CI](https://github.com/shyim/go-spdx/actions/workflows/ci.yml/badge.svg)](https://github.com/shyim/go-spdx/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/shyim/go-spdx)](https://goreportcard.com/report/github.com/shyim/go-spdx)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Go library for validating [SPDX license expressions](https://spdx.dev/). Uses embedded SPDX license list data with zero dependencies — pure Go standard library.

## Installation

```bash
go get github.com/shyim/go-spdx
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/shyim/go-spdx"
)

func main() {
    s, err := spdx.NewSpdxLicenses()
    if err != nil {
        panic(err)
    }

    // Validate a single license
    valid, _ := s.Validate("MIT")
    fmt.Println("MIT valid:", valid) // true

    // Validate a license expression
    valid, _ = s.Validate("(LGPL-2.1-only OR GPL-3.0-or-later)")
    fmt.Println("Expression valid:", valid) // true

    // Validate a slice of licenses (combined with OR)
    valid, _ = s.Validate([]string{"MIT", "Apache-2.0"})
    fmt.Println("Slice valid:", valid) // true
}
```

## Features

- Validate SPDX license identifiers (case-insensitive)
- Validate compound expressions with `AND`, `OR`, and `WITH`
- Supports `LicenseRef-*` and `DocumentRef-*:LicenseRef-*` identifiers
- Handles deprecated license identifiers (`+` suffix)
- Zero dependencies — pure Go standard library
- Thread-safe after initialization

## License

MIT — see [LICENSE](LICENSE) for details.
