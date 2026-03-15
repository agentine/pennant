package pennant

import (
	"fmt"
	"io"
	"os"
	"strings"
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
	normalizeFunc     NormalizeFunc
	addedGoFlags      bool
	mutuallyExclusive []flagGroup
	requiredTogether  []flagGroup
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

// GetInt returns the int value of the named flag, or an error.
func (f *FlagSet) GetInt(name string) (int, error) {
	flag := f.Lookup(name)
	if flag == nil {
		return 0, &ErrUnknownFlag{FlagName: name}
	}
	v, ok := flag.Value.(*intValue)
	if !ok {
		return 0, fmt.Errorf("trying to get int value of flag of type %s", flag.Value.Type())
	}
	return int(*v), nil
}

// -- Flag registration

// VarPF registers a flag with a Value, name, shorthand, and usage string.
// It returns the Flag so callers can set additional fields.
func (f *FlagSet) VarPF(value Value, name, shorthand, usage string) *Flag {
	flag := &Flag{
		Name:      name,
		Shorthand: shorthand,
		Usage:     usage,
		Value:     value,
		DefValue:  value.String(),
	}
	f.AddFlag(flag)
	return flag
}

// VarP is like VarPF but does not return the created Flag.
func (f *FlagSet) VarP(value Value, name, shorthand, usage string) {
	f.VarPF(value, name, shorthand, usage)
}

// Var is like VarP without a shorthand.
func (f *FlagSet) Var(value Value, name, usage string) {
	f.VarPF(value, name, "", usage)
}

// StringVarP defines a string flag with a shorthand.
func (f *FlagSet) StringVarP(p *string, name, shorthand string, value string, usage string) {
	f.VarPF(newStringValue(value, p), name, shorthand, usage)
}

// StringVar defines a string flag.
func (f *FlagSet) StringVar(p *string, name string, value string, usage string) {
	f.StringVarP(p, name, "", value, usage)
}

// StringP defines a string flag with a shorthand and returns the pointer.
func (f *FlagSet) StringP(name, shorthand string, value string, usage string) *string {
	p := new(string)
	f.StringVarP(p, name, shorthand, value, usage)
	return p
}

// String defines a string flag and returns the pointer.
func (f *FlagSet) String(name string, value string, usage string) *string {
	return f.StringP(name, "", value, usage)
}

// BoolVarP defines a bool flag with a shorthand.
func (f *FlagSet) BoolVarP(p *bool, name, shorthand string, value bool, usage string) {
	flag := f.VarPF(newBoolValue(value, p), name, shorthand, usage)
	flag.NoOptDefVal = "true"
}

// BoolVar defines a bool flag.
func (f *FlagSet) BoolVar(p *bool, name string, value bool, usage string) {
	f.BoolVarP(p, name, "", value, usage)
}

// BoolP defines a bool flag with a shorthand and returns the pointer.
func (f *FlagSet) BoolP(name, shorthand string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolVarP(p, name, shorthand, value, usage)
	return p
}

// Bool defines a bool flag and returns the pointer.
func (f *FlagSet) Bool(name string, value bool, usage string) *bool {
	return f.BoolP(name, "", value, usage)
}

// IntVarP defines an int flag with a shorthand.
func (f *FlagSet) IntVarP(p *int, name, shorthand string, value int, usage string) {
	f.VarPF(newIntValue(value, p), name, shorthand, usage)
}

// IntVar defines an int flag.
func (f *FlagSet) IntVar(p *int, name string, value int, usage string) {
	f.IntVarP(p, name, "", value, usage)
}

// IntP defines an int flag with a shorthand and returns the pointer.
func (f *FlagSet) IntP(name, shorthand string, value int, usage string) *int {
	p := new(int)
	f.IntVarP(p, name, shorthand, value, usage)
	return p
}

// Int defines an int flag and returns the pointer.
func (f *FlagSet) Int(name string, value int, usage string) *int {
	return f.IntP(name, "", value, usage)
}

// -- Parsing

// Parse parses the given arguments according to POSIX/GNU conventions.
func (f *FlagSet) Parse(arguments []string) error {
	f.parsed = true
	f.args = []string{}
	f.argsLenAtDash = -1

	err := f.parseArgs(arguments)
	if err == nil {
		err = f.validateFlagGroups()
	}
	if err != nil {
		switch f.errorHandling {
		case ContinueOnError:
			return err
		case ExitOnError:
			fmt.Fprintln(f.GetOutput(), err)
			os.Exit(2)
		case PanicOnError:
			panic(err)
		}
	}
	return nil
}

func (f *FlagSet) parseArgs(args []string) error {
	for len(args) > 0 {
		s := args[0]
		args = args[1:]

		if len(s) == 0 || s[0] != '-' || len(s) == 1 {
			// Not a flag (empty, no dash, or bare "-")
			if !f.interspersed {
				f.args = append(f.args, s)
				f.args = append(f.args, args...)
				return nil
			}
			f.args = append(f.args, s)
			continue
		}

		if s == "--" {
			f.argsLenAtDash = len(f.args)
			f.args = append(f.args, args...)
			return nil
		}

		if s[1] == '-' {
			var err error
			args, err = f.parseLongArg(s, args)
			if err != nil {
				return err
			}
		} else {
			var err error
			args, err = f.parseShortArg(s[1:], args)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (f *FlagSet) parseLongArg(s string, args []string) ([]string, error) {
	name := s[2:]
	if len(name) == 0 || name[0] == '-' || name[0] == '=' {
		return args, fmt.Errorf("bad flag syntax: %s", s)
	}

	// Check for --flag=value
	value := ""
	hasValue := false
	if idx := strings.IndexByte(name, '='); idx >= 0 {
		value = name[idx+1:]
		hasValue = true
		name = name[:idx]
	}

	flag := f.Lookup(name)
	if flag == nil {
		return args, &ErrUnknownFlag{FlagName: name}
	}

	if flag.Deprecated != "" {
		fmt.Fprintf(f.GetOutput(), "Flag --%s has been deprecated, %s\n", name, flag.Deprecated)
	}

	if bf, ok := flag.Value.(boolFlag); ok && bf.IsBoolFlag() {
		if hasValue {
			if err := flag.Value.Set(value); err != nil {
				return args, &ErrParseError{FlagName: name, Value: value, Type: flag.Value.Type(), Err: err}
			}
		} else {
			defVal := flag.NoOptDefVal
			if defVal == "" {
				defVal = "true"
			}
			if err := flag.Value.Set(defVal); err != nil {
				return args, &ErrParseError{FlagName: name, Value: defVal, Type: flag.Value.Type(), Err: err}
			}
		}
	} else {
		if !hasValue {
			if len(args) == 0 {
				return args, &ErrNoValue{FlagName: name}
			}
			value = args[0]
			args = args[1:]
		}
		if err := flag.Value.Set(value); err != nil {
			return args, &ErrParseError{FlagName: name, Value: value, Type: flag.Value.Type(), Err: err}
		}
	}

	flag.Changed = true
	return args, nil
}

func (f *FlagSet) parseShortArg(shorthands string, args []string) ([]string, error) {
	for i := 0; i < len(shorthands); i++ {
		c := shorthands[i]
		flag := f.shorthands[c]
		if flag == nil {
			return args, &ErrUnknownShorthand{Shorthand: c}
		}

		if flag.ShorthandDeprecated != "" {
			fmt.Fprintf(f.GetOutput(), "Flag shorthand -%c has been deprecated, %s\n", c, flag.ShorthandDeprecated)
		}
		if flag.Deprecated != "" {
			fmt.Fprintf(f.GetOutput(), "Flag --%s has been deprecated, %s\n", flag.Name, flag.Deprecated)
		}

		if bf, ok := flag.Value.(boolFlag); ok && bf.IsBoolFlag() {
			// Check for -b=value
			if i+1 < len(shorthands) && shorthands[i+1] == '=' {
				value := shorthands[i+2:]
				if err := flag.Value.Set(value); err != nil {
					return args, &ErrParseError{FlagName: flag.Name, Value: value, Type: flag.Value.Type(), Err: err}
				}
				flag.Changed = true
				return args, nil
			}
			defVal := flag.NoOptDefVal
			if defVal == "" {
				defVal = "true"
			}
			if err := flag.Value.Set(defVal); err != nil {
				return args, &ErrParseError{FlagName: flag.Name, Value: defVal, Type: flag.Value.Type(), Err: err}
			}
			flag.Changed = true
		} else {
			// Non-bool flag: rest of string or next arg is the value
			var value string
			if i+1 < len(shorthands) {
				value = shorthands[i+1:]
				if len(value) > 0 && value[0] == '=' {
					value = value[1:]
				}
			} else if len(args) > 0 {
				value = args[0]
				args = args[1:]
			} else {
				return args, &ErrNoValue{FlagName: flag.Name}
			}
			if err := flag.Value.Set(value); err != nil {
				return args, &ErrParseError{FlagName: flag.Name, Value: value, Type: flag.Value.Type(), Err: err}
			}
			flag.Changed = true
			return args, nil
		}
	}
	return args, nil
}

// Changed returns true if the named flag was explicitly set during parsing.
func (f *FlagSet) Changed(name string) bool {
	flag := f.Lookup(name)
	if flag == nil {
		return false
	}
	return flag.Changed
}

// SetAnnotation sets an annotation on the named flag.
func (f *FlagSet) SetAnnotation(name, key string, values []string) error {
	flag := f.Lookup(name)
	if flag == nil {
		return fmt.Errorf("no such flag -%s", name)
	}
	if flag.Annotations == nil {
		flag.Annotations = make(map[string][]string)
	}
	flag.Annotations[key] = values
	return nil
}
