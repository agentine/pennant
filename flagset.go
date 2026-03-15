package pennant

import (
	"io"
	"os"
)

// ErrorHandling defines how FlagSet.Parse behaves on errors.
type ErrorHandling int

const (
	// ContinueOnError means Parse returns an error.
	ContinueOnError ErrorHandling = iota
	// ExitOnError means Parse calls os.Exit(2).
	ExitOnError
	// PanicOnError means Parse panics.
	PanicOnError
)

// FlagSet is a set of defined flags. It implements POSIX/GNU-style flag parsing.
type FlagSet struct {
	// Usage is the function called when an error occurs while parsing flags.
	// It defaults to printing a simple header and calling PrintDefaults.
	Usage func()

	// SortFlags controls whether flags are sorted by name in help output.
	SortFlags bool

	name          string
	parsed        bool
	errorHandling ErrorHandling
	output        io.Writer
	interspersed  bool
	args          []string // arguments after flags
	argsLenAtDash int      // length of args when -- was found, or -1
	formal        map[NormalizedName]*Flag
	orderedFormal []*Flag
	shorthands    map[byte]*Flag
	normalizeFunc NormalizeFunc
	addedGoFlags  bool
}

// NewFlagSet creates a new FlagSet with the given name and error handling behavior.
func NewFlagSet(name string, errorHandling ErrorHandling) *FlagSet {
	f := &FlagSet{
		name:          name,
		errorHandling: errorHandling,
		SortFlags:     true,
		interspersed:  true,
		argsLenAtDash: -1,
	}
	return f
}

// Init reinitializes a FlagSet with the given name and error handling.
// A FlagSet can be reused after calling Init.
func (f *FlagSet) Init(name string, errorHandling ErrorHandling) {
	f.name = name
	f.errorHandling = errorHandling
	f.argsLenAtDash = -1
}

// SetInterspersed sets whether non-flag arguments can be interspersed with flags.
// If false, parsing stops at the first non-flag argument.
func (f *FlagSet) SetInterspersed(interspersed bool) {
	f.interspersed = interspersed
}

// GetOutput returns the output writer, defaulting to os.Stderr.
func (f *FlagSet) GetOutput() io.Writer {
	if f.output == nil {
		return os.Stderr
	}
	return f.output
}

// SetOutput sets the destination for usage and error messages.
func (f *FlagSet) SetOutput(output io.Writer) {
	f.output = output
}

// Name returns the name of the FlagSet.
func (f *FlagSet) Name() string {
	return f.name
}

// Parsed returns true if Parse has been called.
func (f *FlagSet) Parsed() bool {
	return f.parsed
}

// Args returns the non-flag arguments.
func (f *FlagSet) Args() []string {
	return f.args
}

// NArg returns the number of non-flag arguments.
func (f *FlagSet) NArg() int {
	return len(f.args)
}

// Arg returns the i'th non-flag argument (0-indexed).
func (f *FlagSet) Arg(i int) string {
	if i < 0 || i >= len(f.args) {
		return ""
	}
	return f.args[i]
}

// NFlag returns the number of flags that have been set.
func (f *FlagSet) NFlag() int {
	count := 0
	for _, flag := range f.formal {
		if flag.Changed {
			count++
		}
	}
	return count
}

// ArgsLenAtDash returns the number of args before a "--" was found,
// or -1 if no "--" was found.
func (f *FlagSet) ArgsLenAtDash() int {
	return f.argsLenAtDash
}

// HasFlags returns true if any flags have been defined.
func (f *FlagSet) HasFlags() bool {
	return len(f.formal) > 0
}

// HasAvailableFlags returns true if any flags that are not hidden have been defined.
func (f *FlagSet) HasAvailableFlags() bool {
	for _, flag := range f.formal {
		if !flag.Hidden {
			return true
		}
	}
	return false
}

// Lookup returns the Flag for the given name, or nil if not defined.
func (f *FlagSet) Lookup(name string) *Flag {
	if f.formal == nil {
		return nil
	}
	return f.formal[f.normalizeName(name)]
}

// AddFlag adds a flag to the FlagSet.
func (f *FlagSet) AddFlag(flag *Flag) {
	normalizedName := f.normalizeName(flag.Name)

	if f.formal == nil {
		f.formal = make(map[NormalizedName]*Flag)
	}
	f.formal[normalizedName] = flag
	f.orderedFormal = append(f.orderedFormal, flag)

	if flag.Shorthand != "" && len(flag.Shorthand) == 1 {
		if f.shorthands == nil {
			f.shorthands = make(map[byte]*Flag)
		}
		f.shorthands[flag.Shorthand[0]] = flag
	}
}

// Set sets the value of the named flag.
func (f *FlagSet) Set(name, value string) error {
	flag := f.Lookup(name)
	if flag == nil {
		return &ErrUnknownFlag{FlagName: name}
	}
	if err := flag.Value.Set(value); err != nil {
		return err
	}
	flag.Changed = true
	return nil
}

// Visit visits the flags that have been set, in order of definition.
func (f *FlagSet) Visit(fn func(*Flag)) {
	for _, flag := range f.orderedFormal {
		if flag.Changed {
			fn(flag)
		}
	}
}

// VisitAll visits all flags in the FlagSet, in order of definition.
func (f *FlagSet) VisitAll(fn func(*Flag)) {
	for _, flag := range f.orderedFormal {
		fn(flag)
	}
}

// GetBool returns the bool value of the named flag, or an error.
func (f *FlagSet) GetBool(name string) (bool, error) {
	flag := f.Lookup(name)
	if flag == nil {
		return false, &ErrUnknownFlag{FlagName: name}
	}
	return flag.Value.String() == "true", nil
}

// GetString returns the string value of the named flag, or an error.
func (f *FlagSet) GetString(name string) (string, error) {
	flag := f.Lookup(name)
	if flag == nil {
		return "", &ErrUnknownFlag{FlagName: name}
	}
	return flag.Value.String(), nil
}
