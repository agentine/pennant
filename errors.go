package pennant

import "fmt"

// ErrUnknownFlag is returned when an undefined flag is encountered during parsing.
type ErrUnknownFlag struct {
	FlagName string
}

func (e *ErrUnknownFlag) Error() string {
	return fmt.Sprintf("unknown flag: --%s", e.FlagName)
}

// ErrUnknownShorthand is returned when an undefined shorthand flag is encountered.
type ErrUnknownShorthand struct {
	Shorthand byte
}

func (e *ErrUnknownShorthand) Error() string {
	return fmt.Sprintf("unknown shorthand flag: %q in -%c", e.Shorthand, e.Shorthand)
}

// ErrParseError is returned when a flag value cannot be parsed.
type ErrParseError struct {
	FlagName string
	Value    string
	Type     string
	Err      error
}

func (e *ErrParseError) Error() string {
	return fmt.Sprintf("invalid argument %q for %q flag: %v", e.Value, "--"+e.FlagName, e.Err)
}

func (e *ErrParseError) Unwrap() error {
	return e.Err
}

// ErrNoValue is returned when a flag that requires a value is used without one.
type ErrNoValue struct {
	FlagName string
}

func (e *ErrNoValue) Error() string {
	return fmt.Sprintf("flag needs an argument: --%s", e.FlagName)
}

// ErrMutuallyExclusive is returned when mutually exclusive flags are both set.
type ErrMutuallyExclusive struct {
	FlagNames []string
}

func (e *ErrMutuallyExclusive) Error() string {
	return fmt.Sprintf("if any flags in the group %v are set none of the others can be", e.FlagNames)
}

// ErrRequiredTogether is returned when co-required flags are not all set.
type ErrRequiredTogether struct {
	FlagNames []string
}

func (e *ErrRequiredTogether) Error() string {
	return fmt.Sprintf("if any flags in the group %v are set they must all be set", e.FlagNames)
}
