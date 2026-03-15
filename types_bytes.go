package pennant

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

// -- BytesHex Value

type bytesHexValue []byte

func newBytesHexValue(val []byte, p *[]byte) *bytesHexValue {
	*p = val
	return (*bytesHexValue)(p)
}

func (b *bytesHexValue) Set(val string) error {
	decoded, err := hex.DecodeString(val)
	if err != nil {
		return err
	}
	*b = decoded
	return nil
}

func (b *bytesHexValue) String() string { return hex.EncodeToString(*b) }
func (b *bytesHexValue) Type() string   { return "bytesHex" }

func (f *FlagSet) BytesHexVarP(p *[]byte, name, shorthand string, value []byte, usage string) {
	f.VarPF(newBytesHexValue(value, p), name, shorthand, usage)
}
func (f *FlagSet) BytesHexVar(p *[]byte, name string, value []byte, usage string) {
	f.BytesHexVarP(p, name, "", value, usage)
}
func (f *FlagSet) BytesHexP(name, shorthand string, value []byte, usage string) *[]byte {
	p := new([]byte)
	f.BytesHexVarP(p, name, shorthand, value, usage)
	return p
}
func (f *FlagSet) BytesHex(name string, value []byte, usage string) *[]byte {
	return f.BytesHexP(name, "", value, usage)
}
func (f *FlagSet) GetBytesHex(name string) ([]byte, error) {
	flag := f.Lookup(name)
	if flag == nil {
		return nil, &ErrUnknownFlag{FlagName: name}
	}
	v, ok := flag.Value.(*bytesHexValue)
	if !ok {
		return nil, fmt.Errorf("trying to get bytesHex value of flag of type %s", flag.Value.Type())
	}
	return []byte(*v), nil
}

// -- BytesBase64 Value

type bytesBase64Value []byte

func newBytesBase64Value(val []byte, p *[]byte) *bytesBase64Value {
	*p = val
	return (*bytesBase64Value)(p)
}

func (b *bytesBase64Value) Set(val string) error {
	decoded, err := base64.StdEncoding.DecodeString(val)
	if err != nil {
		return err
	}
	*b = decoded
	return nil
}

func (b *bytesBase64Value) String() string { return base64.StdEncoding.EncodeToString(*b) }
func (b *bytesBase64Value) Type() string   { return "bytesBase64" }

func (f *FlagSet) BytesBase64VarP(p *[]byte, name, shorthand string, value []byte, usage string) {
	f.VarPF(newBytesBase64Value(value, p), name, shorthand, usage)
}
func (f *FlagSet) BytesBase64Var(p *[]byte, name string, value []byte, usage string) {
	f.BytesBase64VarP(p, name, "", value, usage)
}
func (f *FlagSet) BytesBase64P(name, shorthand string, value []byte, usage string) *[]byte {
	p := new([]byte)
	f.BytesBase64VarP(p, name, shorthand, value, usage)
	return p
}
func (f *FlagSet) BytesBase64(name string, value []byte, usage string) *[]byte {
	return f.BytesBase64P(name, "", value, usage)
}
func (f *FlagSet) GetBytesBase64(name string) ([]byte, error) {
	flag := f.Lookup(name)
	if flag == nil {
		return nil, &ErrUnknownFlag{FlagName: name}
	}
	v, ok := flag.Value.(*bytesBase64Value)
	if !ok {
		return nil, fmt.Errorf("trying to get bytesBase64 value of flag of type %s", flag.Value.Type())
	}
	return []byte(*v), nil
}
