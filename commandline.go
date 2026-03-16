package pennant

import (
	"io"
	"net"
	"os"
	"time"
)

// CommandLine is the default FlagSet, used by the package-level functions.
// It parses flags from os.Args.
var CommandLine = NewFlagSet(os.Args[0], ExitOnError)

// Parse parses the command-line flags from os.Args[1:].
func Parse() {
	// CommandLine uses ExitOnError, so Parse will os.Exit on error
	// and never actually return a non-nil error.
	_ = CommandLine.Parse(os.Args[1:])
}

// Parsed returns true if the command-line flags have been parsed.
func Parsed() bool {
	return CommandLine.Parsed()
}

// -- Package-level flag query functions

// Args returns the non-flag command-line arguments.
func Args() []string { return CommandLine.Args() }

// NArg returns the number of non-flag command-line arguments.
func NArg() int { return CommandLine.NArg() }

// Arg returns the i'th non-flag command-line argument.
func Arg(i int) string { return CommandLine.Arg(i) }

// NFlag returns the number of command-line flags that have been set.
func NFlag() int { return CommandLine.NFlag() }

// Lookup returns the Flag for the given name on the CommandLine FlagSet.
func Lookup(name string) *Flag { return CommandLine.Lookup(name) }

// ShorthandLookup returns the Flag for the given one-letter shorthand.
func ShorthandLookup(name string) *Flag { return CommandLine.ShorthandLookup(name) }

// Visit visits the command-line flags that have been set.
func Visit(fn func(*Flag)) { CommandLine.Visit(fn) }

// VisitAll visits all command-line flags.
func VisitAll(fn func(*Flag)) { CommandLine.VisitAll(fn) }

// Set sets the value of the named command-line flag.
func Set(name, value string) error { return CommandLine.Set(name, value) }

// SetOutput sets the destination for usage and error messages.
func SetOutput(output io.Writer) { CommandLine.SetOutput(output) }

// SetInterspersed sets whether non-flag arguments can be interspersed with flags.
func SetInterspersed(interspersed bool) { CommandLine.SetInterspersed(interspersed) }

// SetNormalizeFunc sets a function to transform flag names on the CommandLine.
func SetNormalizeFunc(n NormalizeFunc) { CommandLine.SetNormalizeFunc(n) }

// PrintDefaults prints the default values of all command-line flags.
func PrintDefaults() { CommandLine.PrintDefaults() }

// FlagUsages returns the usage information for all command-line flags.
func FlagUsages() string { return CommandLine.FlagUsages() }

// FlagUsagesWrapped returns the usage information wrapped at the given column.
func FlagUsagesWrapped(cols int) string { return CommandLine.FlagUsagesWrapped(cols) }

// HasFlags returns true if any command-line flags have been defined.
func HasFlags() bool { return CommandLine.HasFlags() }

// HasAvailableFlags returns true if any non-hidden command-line flags have been defined.
func HasAvailableFlags() bool { return CommandLine.HasAvailableFlags() }

// Var registers a custom Value on the CommandLine.
func Var(value Value, name, usage string) { CommandLine.Var(value, name, usage) }

// VarP registers a custom Value with a shorthand on the CommandLine.
func VarP(value Value, name, shorthand, usage string) { CommandLine.VarP(value, name, shorthand, usage) }

// AddFlag adds a flag to the CommandLine FlagSet.
func AddFlag(flag *Flag) { CommandLine.AddFlag(flag) }

// MarkDeprecated marks a flag as deprecated on the CommandLine.
func MarkDeprecated(name, usageMessage string) error {
	return CommandLine.MarkDeprecated(name, usageMessage)
}

// MarkHidden marks a flag as hidden on the CommandLine.
func MarkHidden(name string) error { return CommandLine.MarkHidden(name) }

// MarkShorthandDeprecated marks a shorthand as deprecated on the CommandLine.
func MarkShorthandDeprecated(name, usageMessage string) error {
	return CommandLine.MarkShorthandDeprecated(name, usageMessage)
}

// MarkFlagsMutuallyExclusive marks flags as mutually exclusive on the CommandLine.
func MarkFlagsMutuallyExclusive(flagNames ...string) {
	CommandLine.MarkFlagsMutuallyExclusive(flagNames...)
}

// MarkFlagsRequiredTogether marks flags as required together on the CommandLine.
func MarkFlagsRequiredTogether(flagNames ...string) {
	CommandLine.MarkFlagsRequiredTogether(flagNames...)
}

// SetAnnotation sets an annotation on the CommandLine.
func SetAnnotation(name, key string, values []string) error {
	return CommandLine.SetAnnotation(name, key, values)
}

// -- Package-level typed flag registration (delegates to CommandLine)

// StringVar defines a string flag on the CommandLine.
func StringVar(p *string, name string, value string, usage string) {
	CommandLine.StringVar(p, name, value, usage)
}

// StringVarP defines a string flag with shorthand on the CommandLine.
func StringVarP(p *string, name, shorthand string, value string, usage string) {
	CommandLine.StringVarP(p, name, shorthand, value, usage)
}

// String defines a string flag on the CommandLine and returns a pointer.
func String(name string, value string, usage string) *string {
	return CommandLine.String(name, value, usage)
}

// StringP defines a string flag with shorthand on the CommandLine.
func StringP(name, shorthand string, value string, usage string) *string {
	return CommandLine.StringP(name, shorthand, value, usage)
}

// GetString returns the string value of the named CommandLine flag.
func GetString(name string) (string, error) { return CommandLine.GetString(name) }

// BoolVar defines a bool flag on the CommandLine.
func BoolVar(p *bool, name string, value bool, usage string) {
	CommandLine.BoolVar(p, name, value, usage)
}

// BoolVarP defines a bool flag with shorthand on the CommandLine.
func BoolVarP(p *bool, name, shorthand string, value bool, usage string) {
	CommandLine.BoolVarP(p, name, shorthand, value, usage)
}

// Bool defines a bool flag on the CommandLine and returns a pointer.
func Bool(name string, value bool, usage string) *bool { return CommandLine.Bool(name, value, usage) }

// BoolP defines a bool flag with shorthand on the CommandLine.
func BoolP(name, shorthand string, value bool, usage string) *bool {
	return CommandLine.BoolP(name, shorthand, value, usage)
}

// GetBool returns the bool value of the named CommandLine flag.
func GetBool(name string) (bool, error) { return CommandLine.GetBool(name) }

// IntVar defines an int flag on the CommandLine.
func IntVar(p *int, name string, value int, usage string) { CommandLine.IntVar(p, name, value, usage) }

// IntVarP defines an int flag with shorthand on the CommandLine.
func IntVarP(p *int, name, shorthand string, value int, usage string) {
	CommandLine.IntVarP(p, name, shorthand, value, usage)
}

// Int defines an int flag on the CommandLine and returns a pointer.
func Int(name string, value int, usage string) *int { return CommandLine.Int(name, value, usage) }

// IntP defines an int flag with shorthand on the CommandLine.
func IntP(name, shorthand string, value int, usage string) *int {
	return CommandLine.IntP(name, shorthand, value, usage)
}

// GetInt returns the int value of the named CommandLine flag.
func GetInt(name string) (int, error) { return CommandLine.GetInt(name) }

// Int64Var defines an int64 flag on the CommandLine.
func Int64Var(p *int64, name string, value int64, usage string) {
	CommandLine.Int64Var(p, name, value, usage)
}

// Int64VarP defines an int64 flag with shorthand on the CommandLine.
func Int64VarP(p *int64, name, shorthand string, value int64, usage string) {
	CommandLine.Int64VarP(p, name, shorthand, value, usage)
}

// Int64 defines an int64 flag on the CommandLine and returns a pointer.
func Int64(name string, value int64, usage string) *int64 {
	return CommandLine.Int64(name, value, usage)
}

// Int64P defines an int64 flag with shorthand on the CommandLine.
func Int64P(name, shorthand string, value int64, usage string) *int64 {
	return CommandLine.Int64P(name, shorthand, value, usage)
}

// Float64Var defines a float64 flag on the CommandLine.
func Float64Var(p *float64, name string, value float64, usage string) {
	CommandLine.Float64Var(p, name, value, usage)
}

// Float64VarP defines a float64 flag with shorthand on the CommandLine.
func Float64VarP(p *float64, name, shorthand string, value float64, usage string) {
	CommandLine.Float64VarP(p, name, shorthand, value, usage)
}

// Float64 defines a float64 flag on the CommandLine and returns a pointer.
func Float64(name string, value float64, usage string) *float64 {
	return CommandLine.Float64(name, value, usage)
}

// Float64P defines a float64 flag with shorthand on the CommandLine.
func Float64P(name, shorthand string, value float64, usage string) *float64 {
	return CommandLine.Float64P(name, shorthand, value, usage)
}

// DurationVar defines a duration flag on the CommandLine.
func DurationVar(p *time.Duration, name string, value time.Duration, usage string) {
	CommandLine.DurationVar(p, name, value, usage)
}

// DurationVarP defines a duration flag with shorthand on the CommandLine.
func DurationVarP(p *time.Duration, name, shorthand string, value time.Duration, usage string) {
	CommandLine.DurationVarP(p, name, shorthand, value, usage)
}

// Duration defines a duration flag on the CommandLine and returns a pointer.
func Duration(name string, value time.Duration, usage string) *time.Duration {
	return CommandLine.Duration(name, value, usage)
}

// DurationP defines a duration flag with shorthand on the CommandLine.
func DurationP(name, shorthand string, value time.Duration, usage string) *time.Duration {
	return CommandLine.DurationP(name, shorthand, value, usage)
}

// UintVar defines a uint flag on the CommandLine.
func UintVar(p *uint, name string, value uint, usage string) {
	CommandLine.UintVar(p, name, value, usage)
}

// Uint defines a uint flag on the CommandLine and returns a pointer.
func Uint(name string, value uint, usage string) *uint { return CommandLine.Uint(name, value, usage) }

// Uint64Var defines a uint64 flag on the CommandLine.
func Uint64Var(p *uint64, name string, value uint64, usage string) {
	CommandLine.Uint64Var(p, name, value, usage)
}

// Uint64 defines a uint64 flag on the CommandLine and returns a pointer.
func Uint64(name string, value uint64, usage string) *uint64 {
	return CommandLine.Uint64(name, value, usage)
}

// CountVar defines a count flag on the CommandLine.
func CountVar(p *int, name string, usage string) { CommandLine.CountVar(p, name, usage) }

// CountVarP defines a count flag with shorthand on the CommandLine.
func CountVarP(p *int, name, shorthand string, usage string) {
	CommandLine.CountVarP(p, name, shorthand, usage)
}

// Count defines a count flag on the CommandLine and returns a pointer.
func Count(name string, usage string) *int { return CommandLine.Count(name, usage) }

// CountP defines a count flag with shorthand on the CommandLine.
func CountP(name, shorthand string, usage string) *int {
	return CommandLine.CountP(name, shorthand, usage)
}

// IPVar defines an IP flag on the CommandLine.
func IPVar(p *net.IP, name string, value net.IP, usage string) {
	CommandLine.IPVar(p, name, value, usage)
}

// IP defines an IP flag on the CommandLine and returns a pointer.
func IP(name string, value net.IP, usage string) *net.IP {
	return CommandLine.IP(name, value, usage)
}

// StringSliceVar defines a string slice flag on the CommandLine.
func StringSliceVar(p *[]string, name string, value []string, usage string) {
	CommandLine.StringSliceVar(p, name, value, usage)
}

// StringSlice defines a string slice flag on the CommandLine and returns a pointer.
func StringSlice(name string, value []string, usage string) *[]string {
	return CommandLine.StringSlice(name, value, usage)
}

// IntSliceVar defines an int slice flag on the CommandLine.
func IntSliceVar(p *[]int, name string, value []int, usage string) {
	CommandLine.IntSliceVar(p, name, value, usage)
}

// IntSlice defines an int slice flag on the CommandLine and returns a pointer.
func IntSlice(name string, value []int, usage string) *[]int {
	return CommandLine.IntSlice(name, value, usage)
}

// StringToStringVar defines a string-to-string map flag on the CommandLine.
func StringToStringVar(p *map[string]string, name string, value map[string]string, usage string) {
	CommandLine.StringToStringVar(p, name, value, usage)
}

// StringToString defines a string-to-string map flag on the CommandLine and returns a pointer.
func StringToString(name string, value map[string]string, usage string) *map[string]string {
	return CommandLine.StringToString(name, value, usage)
}

// StringToIntVar defines a string-to-int map flag on the CommandLine.
func StringToIntVar(p *map[string]int, name string, value map[string]int, usage string) {
	CommandLine.StringToIntVar(p, name, value, usage)
}

// StringToInt defines a string-to-int map flag on the CommandLine and returns a pointer.
func StringToInt(name string, value map[string]int, usage string) *map[string]int {
	return CommandLine.StringToInt(name, value, usage)
}

// BytesHexVar defines a bytes-hex flag on the CommandLine.
func BytesHexVar(p *[]byte, name string, value []byte, usage string) {
	CommandLine.BytesHexVar(p, name, value, usage)
}

// BytesHex defines a bytes-hex flag on the CommandLine and returns a pointer.
func BytesHex(name string, value []byte, usage string) *[]byte {
	return CommandLine.BytesHex(name, value, usage)
}

// BytesBase64Var defines a bytes-base64 flag on the CommandLine.
func BytesBase64Var(p *[]byte, name string, value []byte, usage string) {
	CommandLine.BytesBase64Var(p, name, value, usage)
}

// BytesBase64 defines a bytes-base64 flag on the CommandLine and returns a pointer.
func BytesBase64(name string, value []byte, usage string) *[]byte {
	return CommandLine.BytesBase64(name, value, usage)
}
