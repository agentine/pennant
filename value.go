package pennant

// Value is the interface to the dynamic value stored in a flag.
// Compatible with pflag.Value and flag.Value.
type Value interface {
	String() string
	Set(string) error
	Type() string
}

// SliceValue is an optional interface for slice-type flag values.
type SliceValue interface {
	Value
	Append(string) error
	Replace([]string) error
	GetSlice() []string
}
