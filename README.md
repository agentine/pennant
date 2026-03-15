# pennant

A POSIX/GNU-style command-line flag parsing library for Go. Drop-in replacement for [spf13/pflag](https://github.com/spf13/pflag) with zero dependencies.

## Features

- Full POSIX/GNU flag parsing: `--flag=value`, `--flag value`, `-s`, `-abc` combined shorthands, `--` terminator
- Drop-in pflag API compatibility — same `Value` interface, `FlagSet`, `Flag` struct, and convenience functions
- Type-safe generics with `TypedFlag[T]` — no more `GetString` error checking
- Flag groups: mutually exclusive and required-together constraints
- All pflag types: scalars, slices, maps, network types, bytes
- Zero external dependencies — only the Go standard library
- 92%+ test coverage with fuzz testing

## Installation

```
go get github.com/agentine/pennant
```

Requires Go 1.21+.

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/agentine/pennant"
)

func main() {
    name := pennant.StringP("name", "n", "world", "who to greet")
    count := pennant.IntP("count", "c", 1, "number of greetings")
    pennant.Parse()

    for i := 0; i < *count; i++ {
        fmt.Printf("Hello, %s!\n", *name)
    }
}
```

```
$ myapp --name=Go -c 3
Hello, Go!
Hello, Go!
Hello, Go!
```

## Usage

### FlagSet

```go
flags := pennant.NewFlagSet("app", pennant.ExitOnError)
flags.StringP("config", "c", "", "config file path")
flags.BoolP("verbose", "v", false, "enable verbose output")
flags.IntP("port", "p", 8080, "server port")
flags.Parse(os.Args[1:])
```

### Type-Safe Generics

```go
flags := pennant.NewFlagSet("app", pennant.ExitOnError)
name := pennant.TypedString(flags, "name", "n", "world", "who to greet")
port := pennant.TypedInt(flags, "port", "p", 8080, "server port")
flags.Parse(os.Args[1:])

fmt.Println(name.Get()) // string — no error handling needed
fmt.Println(port.Get()) // int — no type assertion needed
```

### All Supported Types

| Category | Types |
|---|---|
| Scalar | `String`, `Bool`, `Int`, `Int8`, `Int16`, `Int32`, `Int64`, `Uint`, `Uint8`, `Uint16`, `Uint32`, `Uint64`, `Float32`, `Float64`, `Duration`, `Count` |
| Network | `IP`, `IPNet`, `IPMask` |
| Bytes | `BytesHex`, `BytesBase64` |
| Slice | `StringSlice`, `StringArray`, `IntSlice`, `Float64Slice`, `BoolSlice`, `DurationSlice`, `IPSlice` |
| Map | `StringToString`, `StringToInt`, `StringToInt64` |

Each type has `Xxx`, `XxxVar`, `XxxP`, `XxxVarP`, and `GetXxx` variants.

### Flag Groups

```go
flags.MarkFlagsMutuallyExclusive("json", "yaml", "table")
flags.MarkFlagsRequiredTogether("username", "password")
```

### Custom Value Types

Implement the `Value` interface:

```go
type Value interface {
    String() string
    Set(string) error
    Type() string
}
```

Register with `Var` or `VarP`:

```go
flags.Var(&myValue, "flag", "usage")
flags.VarP(&myValue, "flag", "f", "usage")
```

## Migrating from pflag

Replace your import and all `pflag.` references:

```go
// Before
import "github.com/spf13/pflag"
pflag.StringP("name", "n", "", "your name")

// After
import "github.com/agentine/pennant"
pennant.StringP("name", "n", "", "your name")
```

See [MIGRATION.md](MIGRATION.md) for the full compatibility table.

## License

MIT
