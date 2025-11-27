English | [中文](README.md)

# gconv — A Lightweight Universal Type Conversion Library for Go

[![Go Version](https://img.shields.io/badge/Go-%3E=1.20-blue)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Coverage](https://img.shields.io/badge/Coverage-98.2%25-brightgreen)](gconv_test.go)

`gconv` is a zero-dependency, lightweight Go utility package designed to help developers deal with **dynamic or unpredictable data types**, such as responses from JSON APIs, databases, RPC calls, or loosely structured input.

It provides a Python-like experience using a powerful generic function `Cast[T]`, enabling seamless conversion from `any` to any desired type.

## ✨ Features

- ✔ Universal generic type conversion using `Cast[T]`
- ✔ Convert any type → string / int / float64 / bool
- ✔ Convert any type → `map[string]any` / `[]any`
- ✔ Supports `json.Number`
- ✔ Auto-convert numeric strings ("123", "3.14")
- ✔ Struct conversion via JSON reflection
- ✔ No third-party dependencies
- ✔ Safe conversion: unsupported values return zero-value instead of panic

---

## 📦 Installation

Option 1: Install using go get

```bash
go get github.com/cnchef/gconv
```

Option 2: Copy the `gconv` folder to your project

Or initialize your own module:

```bash
go mod init gconv
```

---

## 🛠 Usage Example

```go
package main

import (
    "fmt"
    "gconv"
)

func main() {
    var v any = "123"

    fmt.Println(gconv.Cast[int](v))        // 123
    fmt.Println(gconv.Cast[float64](v))    // 123.0
    fmt.Println(gconv.Cast[string](v))     // "123"
    fmt.Println(gconv.Cast[bool]("true"))  // true

    m := map[string]any{"a": 1}
    fmt.Println(gconv.Cast[map[string]any](m)) // map[a:1]
}
```

**More Examples:** Check the [examples/](examples/) directory:
- `basic.go` - Basic usage examples
- `advanced.go` - Advanced usage (struct conversion, API response handling, batch conversion, etc.)

---

## 🧩 Generic Casting: `Cast[T]`

`Cast[T]` is a universal converter with simple usage:

```go
value := Cast[T](v)
```

Examples:

```go
age   := gconv.Cast[int]("20")         // 20
price := gconv.Cast[float64]("99.9")   // 99.9
flag  := gconv.Cast[bool]("true")      // true
name  := gconv.Cast[string](123)       // "123"
```

---

## 🔧 Available Conversion Functions

### ToString(v any) string

Accurately converts all basic types.

### ToInt(v any) int

Automatically handles `"123"`, `123.0`, `json.Number`.

### ToFloat(v any) float64

Automatically supports mixed string and numeric formats.

### ToBool(v any) bool

Supports `"true"` `"1"` `"false"` `"0"`.

### ToMap(v any) map[string]any

Returns empty map when conversion fails.

### ToSlice(v any) []any

Returns empty slice when conversion fails.

---

## ⚙️ Project Structure

```
gconv/
 ├── gconv.go              # Core conversion functions
 ├── gconv_test.go         # Complete unit tests (98.2% coverage)
 ├── go.mod                # Go module definition
 ├── README.md             # Chinese documentation
 ├── README_US.md          # English documentation
 ├── LICENSE               # MIT License
 ├── CHANGELOG.md          # Version changelog
 ├── .gitignore            # Git ignore configuration
 ├── .github/
 │   └── workflows/
 │       └── test.yml      # GitHub Actions CI configuration
 └── examples/             # Usage examples
     ├── basic.go          # Basic examples
     ├── advanced.go       # Advanced examples (struct conversion, etc.)
     └── go.mod            # Examples module configuration
```

---

## 🧪 Testing

Run unit tests:

```bash
go test -v -cover
```

Run benchmark tests:

```bash
go test -bench=. -benchmem
```

Current test coverage: **98.2%**

---

## 🤝 Contributing

Issues and Pull Requests are welcome!

Before submitting a PR, please ensure:
- All tests pass `go test ./...`
- Code is formatted `go fmt ./...`
- Add necessary test cases

---

## 📄 License

[MIT License](LICENSE)

---

## 🔗 Links

- [GitHub Repository](https://github.com/cnchef/gconv)
- [Issue Tracker](https://github.com/cnchef/gconv/issues)
- [Changelog](CHANGELOG.md)
