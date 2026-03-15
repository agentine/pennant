package pennant

import (
	"fmt"
	"strconv"
	"time"
)

// -- int8 Value

type int8Value int8

func newInt8Value(val int8, p *int8) *int8Value {
	*p = val
	return (*int8Value)(p)
}

func (i *int8Value) Set(val string) error {
	v, err := strconv.ParseInt(val, 0, 8)
	if err != nil {
		return err
	}
	*i = int8Value(v)
	return nil
}

func (i *int8Value) String() string { return strconv.FormatInt(int64(*i), 10) }
func (i *int8Value) Type() string   { return "int8" }

func (f *FlagSet) Int8VarP(p *int8, name, shorthand string, value int8, usage string) {
	f.VarPF(newInt8Value(value, p), name, shorthand, usage)
}
func (f *FlagSet) Int8Var(p *int8, name string, value int8, usage string) {
	f.Int8VarP(p, name, "", value, usage)
}
func (f *FlagSet) Int8P(name, shorthand string, value int8, usage string) *int8 {
	p := new(int8)
	f.Int8VarP(p, name, shorthand, value, usage)
	return p
}
func (f *FlagSet) Int8(name string, value int8, usage string) *int8 {
	return f.Int8P(name, "", value, usage)
}
func (f *FlagSet) GetInt8(name string) (int8, error) {
	flag := f.Lookup(name)
	if flag == nil {
		return 0, &ErrUnknownFlag{FlagName: name}
	}
	v, ok := flag.Value.(*int8Value)
	if !ok {
		return 0, fmt.Errorf("trying to get int8 value of flag of type %s", flag.Value.Type())
	}
	return int8(*v), nil
}

// -- int16 Value

type int16Value int16

func newInt16Value(val int16, p *int16) *int16Value {
	*p = val
	return (*int16Value)(p)
}

func (i *int16Value) Set(val string) error {
	v, err := strconv.ParseInt(val, 0, 16)
	if err != nil {
		return err
	}
	*i = int16Value(v)
	return nil
}

func (i *int16Value) String() string { return strconv.FormatInt(int64(*i), 10) }
func (i *int16Value) Type() string   { return "int16" }

func (f *FlagSet) Int16VarP(p *int16, name, shorthand string, value int16, usage string) {
	f.VarPF(newInt16Value(value, p), name, shorthand, usage)
}
func (f *FlagSet) Int16Var(p *int16, name string, value int16, usage string) {
	f.Int16VarP(p, name, "", value, usage)
}
func (f *FlagSet) Int16P(name, shorthand string, value int16, usage string) *int16 {
	p := new(int16)
	f.Int16VarP(p, name, shorthand, value, usage)
	return p
}
func (f *FlagSet) Int16(name string, value int16, usage string) *int16 {
	return f.Int16P(name, "", value, usage)
}
func (f *FlagSet) GetInt16(name string) (int16, error) {
	flag := f.Lookup(name)
	if flag == nil {
		return 0, &ErrUnknownFlag{FlagName: name}
	}
	v, ok := flag.Value.(*int16Value)
	if !ok {
		return 0, fmt.Errorf("trying to get int16 value of flag of type %s", flag.Value.Type())
	}
	return int16(*v), nil
}

// -- int32 Value

type int32Value int32

func newInt32Value(val int32, p *int32) *int32Value {
	*p = val
	return (*int32Value)(p)
}

func (i *int32Value) Set(val string) error {
	v, err := strconv.ParseInt(val, 0, 32)
	if err != nil {
		return err
	}
	*i = int32Value(v)
	return nil
}

func (i *int32Value) String() string { return strconv.FormatInt(int64(*i), 10) }
func (i *int32Value) Type() string   { return "int32" }

func (f *FlagSet) Int32VarP(p *int32, name, shorthand string, value int32, usage string) {
	f.VarPF(newInt32Value(value, p), name, shorthand, usage)
}
func (f *FlagSet) Int32Var(p *int32, name string, value int32, usage string) {
	f.Int32VarP(p, name, "", value, usage)
}
func (f *FlagSet) Int32P(name, shorthand string, value int32, usage string) *int32 {
	p := new(int32)
	f.Int32VarP(p, name, shorthand, value, usage)
	return p
}
func (f *FlagSet) Int32(name string, value int32, usage string) *int32 {
	return f.Int32P(name, "", value, usage)
}
func (f *FlagSet) GetInt32(name string) (int32, error) {
	flag := f.Lookup(name)
	if flag == nil {
		return 0, &ErrUnknownFlag{FlagName: name}
	}
	v, ok := flag.Value.(*int32Value)
	if !ok {
		return 0, fmt.Errorf("trying to get int32 value of flag of type %s", flag.Value.Type())
	}
	return int32(*v), nil
}

// -- int64 Value

type int64Value int64

func newInt64Value(val int64, p *int64) *int64Value {
	*p = val
	return (*int64Value)(p)
}

func (i *int64Value) Set(val string) error {
	v, err := strconv.ParseInt(val, 0, 64)
	if err != nil {
		return err
	}
	*i = int64Value(v)
	return nil
}

func (i *int64Value) String() string { return strconv.FormatInt(int64(*i), 10) }
func (i *int64Value) Type() string   { return "int64" }

func (f *FlagSet) Int64VarP(p *int64, name, shorthand string, value int64, usage string) {
	f.VarPF(newInt64Value(value, p), name, shorthand, usage)
}
func (f *FlagSet) Int64Var(p *int64, name string, value int64, usage string) {
	f.Int64VarP(p, name, "", value, usage)
}
func (f *FlagSet) Int64P(name, shorthand string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64VarP(p, name, shorthand, value, usage)
	return p
}
func (f *FlagSet) Int64(name string, value int64, usage string) *int64 {
	return f.Int64P(name, "", value, usage)
}
func (f *FlagSet) GetInt64(name string) (int64, error) {
	flag := f.Lookup(name)
	if flag == nil {
		return 0, &ErrUnknownFlag{FlagName: name}
	}
	v, ok := flag.Value.(*int64Value)
	if !ok {
		return 0, fmt.Errorf("trying to get int64 value of flag of type %s", flag.Value.Type())
	}
	return int64(*v), nil
}

// -- uint Value

type uintValue uint

func newUintValue(val uint, p *uint) *uintValue {
	*p = val
	return (*uintValue)(p)
}

func (i *uintValue) Set(val string) error {
	v, err := strconv.ParseUint(val, 0, strconv.IntSize)
	if err != nil {
		return err
	}
	*i = uintValue(v)
	return nil
}

func (i *uintValue) String() string { return strconv.FormatUint(uint64(*i), 10) }
func (i *uintValue) Type() string   { return "uint" }

func (f *FlagSet) UintVarP(p *uint, name, shorthand string, value uint, usage string) {
	f.VarPF(newUintValue(value, p), name, shorthand, usage)
}
func (f *FlagSet) UintVar(p *uint, name string, value uint, usage string) {
	f.UintVarP(p, name, "", value, usage)
}
func (f *FlagSet) UintP(name, shorthand string, value uint, usage string) *uint {
	p := new(uint)
	f.UintVarP(p, name, shorthand, value, usage)
	return p
}
func (f *FlagSet) Uint(name string, value uint, usage string) *uint {
	return f.UintP(name, "", value, usage)
}
func (f *FlagSet) GetUint(name string) (uint, error) {
	flag := f.Lookup(name)
	if flag == nil {
		return 0, &ErrUnknownFlag{FlagName: name}
	}
	v, ok := flag.Value.(*uintValue)
	if !ok {
		return 0, fmt.Errorf("trying to get uint value of flag of type %s", flag.Value.Type())
	}
	return uint(*v), nil
}

// -- uint8 Value

type uint8Value uint8

func newUint8Value(val uint8, p *uint8) *uint8Value {
	*p = val
	return (*uint8Value)(p)
}

func (i *uint8Value) Set(val string) error {
	v, err := strconv.ParseUint(val, 0, 8)
	if err != nil {
		return err
	}
	*i = uint8Value(v)
	return nil
}

func (i *uint8Value) String() string { return strconv.FormatUint(uint64(*i), 10) }
func (i *uint8Value) Type() string   { return "uint8" }

func (f *FlagSet) Uint8VarP(p *uint8, name, shorthand string, value uint8, usage string) {
	f.VarPF(newUint8Value(value, p), name, shorthand, usage)
}
func (f *FlagSet) Uint8Var(p *uint8, name string, value uint8, usage string) {
	f.Uint8VarP(p, name, "", value, usage)
}
func (f *FlagSet) Uint8P(name, shorthand string, value uint8, usage string) *uint8 {
	p := new(uint8)
	f.Uint8VarP(p, name, shorthand, value, usage)
	return p
}
func (f *FlagSet) Uint8(name string, value uint8, usage string) *uint8 {
	return f.Uint8P(name, "", value, usage)
}
func (f *FlagSet) GetUint8(name string) (uint8, error) {
	flag := f.Lookup(name)
	if flag == nil {
		return 0, &ErrUnknownFlag{FlagName: name}
	}
	v, ok := flag.Value.(*uint8Value)
	if !ok {
		return 0, fmt.Errorf("trying to get uint8 value of flag of type %s", flag.Value.Type())
	}
	return uint8(*v), nil
}

// -- uint16 Value

type uint16Value uint16

func newUint16Value(val uint16, p *uint16) *uint16Value {
	*p = val
	return (*uint16Value)(p)
}

func (i *uint16Value) Set(val string) error {
	v, err := strconv.ParseUint(val, 0, 16)
	if err != nil {
		return err
	}
	*i = uint16Value(v)
	return nil
}

func (i *uint16Value) String() string { return strconv.FormatUint(uint64(*i), 10) }
func (i *uint16Value) Type() string   { return "uint16" }

func (f *FlagSet) Uint16VarP(p *uint16, name, shorthand string, value uint16, usage string) {
	f.VarPF(newUint16Value(value, p), name, shorthand, usage)
}
func (f *FlagSet) Uint16Var(p *uint16, name string, value uint16, usage string) {
	f.Uint16VarP(p, name, "", value, usage)
}
func (f *FlagSet) Uint16P(name, shorthand string, value uint16, usage string) *uint16 {
	p := new(uint16)
	f.Uint16VarP(p, name, shorthand, value, usage)
	return p
}
func (f *FlagSet) Uint16(name string, value uint16, usage string) *uint16 {
	return f.Uint16P(name, "", value, usage)
}
func (f *FlagSet) GetUint16(name string) (uint16, error) {
	flag := f.Lookup(name)
	if flag == nil {
		return 0, &ErrUnknownFlag{FlagName: name}
	}
	v, ok := flag.Value.(*uint16Value)
	if !ok {
		return 0, fmt.Errorf("trying to get uint16 value of flag of type %s", flag.Value.Type())
	}
	return uint16(*v), nil
}

// -- uint32 Value

type uint32Value uint32

func newUint32Value(val uint32, p *uint32) *uint32Value {
	*p = val
	return (*uint32Value)(p)
}

func (i *uint32Value) Set(val string) error {
	v, err := strconv.ParseUint(val, 0, 32)
	if err != nil {
		return err
	}
	*i = uint32Value(v)
	return nil
}

func (i *uint32Value) String() string { return strconv.FormatUint(uint64(*i), 10) }
func (i *uint32Value) Type() string   { return "uint32" }

func (f *FlagSet) Uint32VarP(p *uint32, name, shorthand string, value uint32, usage string) {
	f.VarPF(newUint32Value(value, p), name, shorthand, usage)
}
func (f *FlagSet) Uint32Var(p *uint32, name string, value uint32, usage string) {
	f.Uint32VarP(p, name, "", value, usage)
}
func (f *FlagSet) Uint32P(name, shorthand string, value uint32, usage string) *uint32 {
	p := new(uint32)
	f.Uint32VarP(p, name, shorthand, value, usage)
	return p
}
func (f *FlagSet) Uint32(name string, value uint32, usage string) *uint32 {
	return f.Uint32P(name, "", value, usage)
}
func (f *FlagSet) GetUint32(name string) (uint32, error) {
	flag := f.Lookup(name)
	if flag == nil {
		return 0, &ErrUnknownFlag{FlagName: name}
	}
	v, ok := flag.Value.(*uint32Value)
	if !ok {
		return 0, fmt.Errorf("trying to get uint32 value of flag of type %s", flag.Value.Type())
	}
	return uint32(*v), nil
}

// -- uint64 Value

type uint64Value uint64

func newUint64Value(val uint64, p *uint64) *uint64Value {
	*p = val
	return (*uint64Value)(p)
}

func (i *uint64Value) Set(val string) error {
	v, err := strconv.ParseUint(val, 0, 64)
	if err != nil {
		return err
	}
	*i = uint64Value(v)
	return nil
}

func (i *uint64Value) String() string { return strconv.FormatUint(uint64(*i), 10) }
func (i *uint64Value) Type() string   { return "uint64" }

func (f *FlagSet) Uint64VarP(p *uint64, name, shorthand string, value uint64, usage string) {
	f.VarPF(newUint64Value(value, p), name, shorthand, usage)
}
func (f *FlagSet) Uint64Var(p *uint64, name string, value uint64, usage string) {
	f.Uint64VarP(p, name, "", value, usage)
}
func (f *FlagSet) Uint64P(name, shorthand string, value uint64, usage string) *uint64 {
	p := new(uint64)
	f.Uint64VarP(p, name, shorthand, value, usage)
	return p
}
func (f *FlagSet) Uint64(name string, value uint64, usage string) *uint64 {
	return f.Uint64P(name, "", value, usage)
}
func (f *FlagSet) GetUint64(name string) (uint64, error) {
	flag := f.Lookup(name)
	if flag == nil {
		return 0, &ErrUnknownFlag{FlagName: name}
	}
	v, ok := flag.Value.(*uint64Value)
	if !ok {
		return 0, fmt.Errorf("trying to get uint64 value of flag of type %s", flag.Value.Type())
	}
	return uint64(*v), nil
}

// -- float32 Value

type float32Value float32

func newFloat32Value(val float32, p *float32) *float32Value {
	*p = val
	return (*float32Value)(p)
}

func (f *float32Value) Set(val string) error {
	v, err := strconv.ParseFloat(val, 32)
	if err != nil {
		return err
	}
	*f = float32Value(v)
	return nil
}

func (fv *float32Value) String() string { return strconv.FormatFloat(float64(*fv), 'g', -1, 32) }
func (fv *float32Value) Type() string   { return "float32" }

func (f *FlagSet) Float32VarP(p *float32, name, shorthand string, value float32, usage string) {
	f.VarPF(newFloat32Value(value, p), name, shorthand, usage)
}
func (f *FlagSet) Float32Var(p *float32, name string, value float32, usage string) {
	f.Float32VarP(p, name, "", value, usage)
}
func (f *FlagSet) Float32P(name, shorthand string, value float32, usage string) *float32 {
	p := new(float32)
	f.Float32VarP(p, name, shorthand, value, usage)
	return p
}
func (f *FlagSet) Float32(name string, value float32, usage string) *float32 {
	return f.Float32P(name, "", value, usage)
}
func (f *FlagSet) GetFloat32(name string) (float32, error) {
	flag := f.Lookup(name)
	if flag == nil {
		return 0, &ErrUnknownFlag{FlagName: name}
	}
	v, ok := flag.Value.(*float32Value)
	if !ok {
		return 0, fmt.Errorf("trying to get float32 value of flag of type %s", flag.Value.Type())
	}
	return float32(*v), nil
}

// -- float64 Value

type float64Value float64

func newFloat64Value(val float64, p *float64) *float64Value {
	*p = val
	return (*float64Value)(p)
}

func (fv *float64Value) Set(val string) error {
	v, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return err
	}
	*fv = float64Value(v)
	return nil
}

func (fv *float64Value) String() string { return strconv.FormatFloat(float64(*fv), 'g', -1, 64) }
func (fv *float64Value) Type() string   { return "float64" }

func (f *FlagSet) Float64VarP(p *float64, name, shorthand string, value float64, usage string) {
	f.VarPF(newFloat64Value(value, p), name, shorthand, usage)
}
func (f *FlagSet) Float64Var(p *float64, name string, value float64, usage string) {
	f.Float64VarP(p, name, "", value, usage)
}
func (f *FlagSet) Float64P(name, shorthand string, value float64, usage string) *float64 {
	p := new(float64)
	f.Float64VarP(p, name, shorthand, value, usage)
	return p
}
func (f *FlagSet) Float64(name string, value float64, usage string) *float64 {
	return f.Float64P(name, "", value, usage)
}
func (f *FlagSet) GetFloat64(name string) (float64, error) {
	flag := f.Lookup(name)
	if flag == nil {
		return 0, &ErrUnknownFlag{FlagName: name}
	}
	v, ok := flag.Value.(*float64Value)
	if !ok {
		return 0, fmt.Errorf("trying to get float64 value of flag of type %s", flag.Value.Type())
	}
	return float64(*v), nil
}

// -- duration Value

type durationValue time.Duration

func newDurationValue(val time.Duration, p *time.Duration) *durationValue {
	*p = val
	return (*durationValue)(p)
}

func (d *durationValue) Set(val string) error {
	v, err := time.ParseDuration(val)
	if err != nil {
		return err
	}
	*d = durationValue(v)
	return nil
}

func (d *durationValue) String() string { return time.Duration(*d).String() }
func (d *durationValue) Type() string   { return "duration" }

func (f *FlagSet) DurationVarP(p *time.Duration, name, shorthand string, value time.Duration, usage string) {
	f.VarPF(newDurationValue(value, p), name, shorthand, usage)
}
func (f *FlagSet) DurationVar(p *time.Duration, name string, value time.Duration, usage string) {
	f.DurationVarP(p, name, "", value, usage)
}
func (f *FlagSet) DurationP(name, shorthand string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.DurationVarP(p, name, shorthand, value, usage)
	return p
}
func (f *FlagSet) Duration(name string, value time.Duration, usage string) *time.Duration {
	return f.DurationP(name, "", value, usage)
}
func (f *FlagSet) GetDuration(name string) (time.Duration, error) {
	flag := f.Lookup(name)
	if flag == nil {
		return 0, &ErrUnknownFlag{FlagName: name}
	}
	v, ok := flag.Value.(*durationValue)
	if !ok {
		return 0, fmt.Errorf("trying to get duration value of flag of type %s", flag.Value.Type())
	}
	return time.Duration(*v), nil
}

// -- count Value (increments each time it's set; implements boolFlag)

type countValue int

func newCountValue(val int, p *int) *countValue {
	*p = val
	return (*countValue)(p)
}

func (c *countValue) Set(val string) error {
	if val == "true" {
		*c++
		return nil
	}
	v, err := strconv.ParseInt(val, 0, 0)
	if err != nil {
		return err
	}
	*c = countValue(v)
	return nil
}

func (c *countValue) String() string  { return strconv.Itoa(int(*c)) }
func (c *countValue) Type() string    { return "count" }
func (c *countValue) IsBoolFlag() bool { return true }

func (f *FlagSet) CountVarP(p *int, name, shorthand string, usage string) {
	flag := f.VarPF(newCountValue(0, p), name, shorthand, usage)
	flag.NoOptDefVal = "true"
}
func (f *FlagSet) CountVar(p *int, name string, usage string) {
	f.CountVarP(p, name, "", usage)
}
func (f *FlagSet) CountP(name, shorthand string, usage string) *int {
	p := new(int)
	f.CountVarP(p, name, shorthand, usage)
	return p
}
func (f *FlagSet) Count(name string, usage string) *int {
	return f.CountP(name, "", usage)
}
func (f *FlagSet) GetCount(name string) (int, error) {
	flag := f.Lookup(name)
	if flag == nil {
		return 0, &ErrUnknownFlag{FlagName: name}
	}
	v, ok := flag.Value.(*countValue)
	if !ok {
		return 0, fmt.Errorf("trying to get count value of flag of type %s", flag.Value.Type())
	}
	return int(*v), nil
}
