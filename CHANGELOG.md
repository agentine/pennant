# Changelog

## v0.1.0 — 2026-03-16

Initial release. Drop-in replacement for `spf13/pflag` with modern Go practices.

### Features

- Full POSIX/GNU-style flag parsing (`--long`, `-s`, `--flag=value`, `--`, combined shorthands `-abc`)
- API-compatible with `spf13/pflag` — same `FlagSet`, `Var`/`StringVar`/`BoolVar` patterns
- All built-in value types: string, bool, int/int8-64, uint/uint8-64, float32/64, duration, count
- Slice types: string, int, float64, bool, duration, IP
- Map types: string-to-string, string-to-int
- Network types: IP, IPNet, IPMask
- Byte types: bytes (hex and base64)
- `TypedFlag[T]` generics for type-safe flag access
- Flag groups: mutually exclusive and required-together constraints
- Deprecated and hidden flag support
- Flag name normalization via `SetNormalizeFunc`
- Flag annotations
- Structured error types
- Comprehensive help/usage formatting
- Zero external dependencies
- 92%+ test coverage with fuzz tests
- Migration guide from `spf13/pflag`
