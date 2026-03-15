# pennant — POSIX/GNU-style Flag Parsing for Go

**Replaces:** `spf13/pflag` (54,528+ importers, 2.7k stars)
**Language:** Go
**Package:** `github.com/agentine/pennant`

## Problem

`spf13/pflag` is the de facto standard for POSIX/GNU-style command-line flag parsing in Go, used by cobra, viper, Kubernetes, and ~30,000 open-source repositories. It is effectively unmaintained: 66 open issues, 73 open PRs, last release Sept 2024, and the original maintainer (spf13) has moved on. A GitHub issue (#385) about maintenance status shows community concern but no resolution.

Go's stdlib `flag` package only supports `-flag` style; it does not support `--long-flag`, `-s` shorthand, or other POSIX/GNU conventions that are standard in modern CLIs.

## Goals

1. **API-compatible drop-in replacement** for `spf13/pflag` — same `FlagSet` API, same `Var`/`StringVar`/`BoolVar` patterns
2. **Fix known pflag issues** — address the most common open issues (deprecated flag handling, flag grouping, better error messages)
3. **Modern Go practices** — Go 1.21+ minimum, generics where appropriate, `fs.FS`-style interfaces
4. **Zero dependencies** — no external imports beyond the Go standard library
5. **Comprehensive test coverage** — 90%+ coverage, fuzz testing for parser edge cases

## Architecture

### Core Package: `pennant`

```
pennant/
  flag.go          — FlagSet type, core parsing logic
  flagset.go       — FlagSet methods (AddFlag, Parse, etc.)
  value.go         — Value interface + built-in value types
  types.go         — Typed flag registration (String, Int, Bool, etc.)
  shorthand.go     — Shorthand flag resolution (-v, -abc combined)
  deprecated.go    — Deprecated/hidden flag support
  groups.go        — Flag groups (mutually exclusive, required together)
  normalize.go     — Flag name normalization
  errors.go        — Structured error types
  help.go          — Usage/help formatting
  commandline.go   — Default CommandLine flagset (mirrors stdlib)
```

### Key Design Decisions

- **Value interface** identical to `pflag.Value` (and `flag.Value`): `String() string`, `Set(string) error`, `Type() string`
- **ShorthandLookup** and **Lookup** methods preserved
- **Flag struct** fields match pflag: Name, Shorthand, Usage, Value, DefValue, Changed, Hidden, Deprecated, Annotations
- **Generics**: Add `TypedFlag[T]` for type-safe flag access alongside the traditional API
- **Normalization**: Support `SetNormalizeFunc` for flag name transforms (e.g., underscores to hyphens)
- **POSIX compliance**: `--`, `-`, combined shorthands (`-abc`), `--flag=value`, `--flag value`

## Deliverables

1. Core `pennant` package with full pflag-compatible API
2. Generics-based typed flag helpers
3. Flag groups (mutually exclusive, co-required)
4. Comprehensive test suite with fuzz tests
5. Migration guide from `spf13/pflag`
6. README with examples and comparison

## Compatibility Matrix

| pflag Feature | pennant Status |
|---|---|
| Long flags (`--name`) | Supported |
| Short flags (`-n`) | Supported |
| Combined shorts (`-abc`) | Supported |
| `--flag=value` | Supported |
| `--flag value` | Supported |
| `--` terminator | Supported |
| Bool flag without value | Supported |
| Deprecated flags | Supported |
| Hidden flags | Supported |
| Flag annotations | Supported |
| ShorthandLookup | Supported |
| SetNormalizeFunc | Supported |
| ip/ipnet/ipmask types | Supported |
| Slice types | Supported |
| Count flags | Supported |
