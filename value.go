package pennant

import "strconv"

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

// boolFlag is an optional interface for boolean-like flag values.
// If a Value implements this and IsBoolFlag returns true, the flag
// can be used without an explicit value (--flag means --flag=true).
type boolFlag interface {
	Value
	IsBoolFlag() bool
}

// -- string Value

type stringValue string

func newStringValue(val string, p *string) *stringValue {
	*p = val
	return (*stringValue)(p)
}

func (s *stringValue) Set(val string) error {
	*s = stringValue(val)
	return nil
}

func (s *stringValue) String() string { return string(*s) }
func (s *stringValue) Type() string   { return "string" }

// -- bool Value

type boolValue bool

func newBoolValue(val bool, p *bool) *boolValue {
	*p = val
	return (*boolValue)(p)
}

func (b *boolValue) Set(val string) error {
	v, err := strconv.ParseBool(val)
	if err != nil {
		return err
	}
	*b = boolValue(v)
	return nil
}

func (b *boolValue) String() string  { return strconv.FormatBool(bool(*b)) }
func (b *boolValue) Type() string    { return "bool" }
func (b *boolValue) IsBoolFlag() bool { return true }

// -- int Value

type intValue int

func newIntValue(val int, p *int) *intValue {
	*p = val
	return (*intValue)(p)
}

func (i *intValue) Set(val string) error {
	v, err := strconv.ParseInt(val, 0, strconv.IntSize)
	if err != nil {
		return err
	}
	*i = intValue(v)
	return nil
}

func (i *intValue) String() string { return strconv.Itoa(int(*i)) }
func (i *intValue) Type() string   { return "int" }
