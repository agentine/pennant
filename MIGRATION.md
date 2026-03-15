# Migrating from spf13/pflag to pennant

## Quick Start

Replace your import:

```go
// Before
import "github.com/spf13/pflag"

// After
import "github.com/agentine/pennant"
```

Replace all `pflag.` references with `pennant.`:

```go
// Before
pflag.StringP("name", "n", "", "your name")
pflag.Parse()

// After
pennant.StringP("name", "n", "", "your name")
pennant.Parse()
```

## API Compatibility

pennant implements the full pflag API surface:

| pflag API | pennant API | Notes |
|---|---|---|
| `pflag.FlagSet` | `pennant.FlagSet` | Same struct fields and methods |
| `pflag.Value` | `pennant.Value` | Same interface: `String()`, `Set()`, `Type()` |
| `pflag.Flag` | `pennant.Flag` | Same fields: Name, Shorthand, Usage, Value, DefValue, Changed, Hidden, Deprecated, Annotations |
| `pflag.NewFlagSet` | `pennant.NewFlagSet` | Same signature |
| `pflag.CommandLine` | `pennant.CommandLine` | Same global default |
| `pflag.Parse()` | `pennant.Parse()` | Same behavior |
| `pflag.ContinueOnError` | `pennant.ContinueOnError` | Same constants |
| `pflag.ExitOnError` | `pennant.ExitOnError` | |
| `pflag.PanicOnError` | `pennant.PanicOnError` | |

## Supported Types

All pflag types are supported:

- **Scalar**: String, Bool, Int, Int8, Int16, Int32, Int64, Uint, Uint8, Uint16, Uint32, Uint64, Float32, Float64, Duration, Count
- **Network**: IP, IPNet, IPMask
- **Bytes**: BytesHex, BytesBase64
- **Slice**: StringSlice, StringArray, IntSlice, Float64Slice, BoolSlice, DurationSlice, IPSlice
- **Map**: StringToString, StringToInt, StringToInt64

Each type has the full set of registration variants: `Xxx`, `XxxVar`, `XxxP`, `XxxVarP`, and `GetXxx`.

## New Features

### TypedFlag Generics

pennant adds type-safe flag access using Go generics:

```go
f := pennant.NewFlagSet("app", pennant.ExitOnError)
name := pennant.TypedString(f, "name", "n", "world", "your name")
count := pennant.TypedInt(f, "count", "c", 1, "repeat count")
f.Parse(os.Args[1:])

fmt.Println(name.Get()) // string, no error or type assertion needed
fmt.Println(count.Get()) // int, no error or type assertion needed
```

### Flag Groups

```go
f.MarkFlagsMutuallyExclusive("json", "yaml", "table")
f.MarkFlagsRequiredTogether("username", "password")
```

## Zero Dependencies

pennant uses only the Go standard library. No transitive dependencies.
