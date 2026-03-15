package pennant

import (
	"net"
	"time"
)

// TypedFlag provides type-safe access to a flag's value without type assertions.
type TypedFlag[T any] struct {
	ptr *T
}

// Get returns the current value of the flag.
func (tf *TypedFlag[T]) Get() T {
	return *tf.ptr
}

// Ptr returns a pointer to the flag's value.
func (tf *TypedFlag[T]) Ptr() *T {
	return tf.ptr
}

// TypedString creates a type-safe string flag on the given FlagSet.
func TypedString(f *FlagSet, name, shorthand string, value string, usage string) *TypedFlag[string] {
	p := new(string)
	f.StringVarP(p, name, shorthand, value, usage)
	return &TypedFlag[string]{ptr: p}
}

// TypedBool creates a type-safe bool flag on the given FlagSet.
func TypedBool(f *FlagSet, name, shorthand string, value bool, usage string) *TypedFlag[bool] {
	p := new(bool)
	f.BoolVarP(p, name, shorthand, value, usage)
	return &TypedFlag[bool]{ptr: p}
}

// TypedInt creates a type-safe int flag on the given FlagSet.
func TypedInt(f *FlagSet, name, shorthand string, value int, usage string) *TypedFlag[int] {
	p := new(int)
	f.IntVarP(p, name, shorthand, value, usage)
	return &TypedFlag[int]{ptr: p}
}

// TypedInt8 creates a type-safe int8 flag.
func TypedInt8(f *FlagSet, name, shorthand string, value int8, usage string) *TypedFlag[int8] {
	p := new(int8)
	f.Int8VarP(p, name, shorthand, value, usage)
	return &TypedFlag[int8]{ptr: p}
}

// TypedInt16 creates a type-safe int16 flag.
func TypedInt16(f *FlagSet, name, shorthand string, value int16, usage string) *TypedFlag[int16] {
	p := new(int16)
	f.Int16VarP(p, name, shorthand, value, usage)
	return &TypedFlag[int16]{ptr: p}
}

// TypedInt32 creates a type-safe int32 flag.
func TypedInt32(f *FlagSet, name, shorthand string, value int32, usage string) *TypedFlag[int32] {
	p := new(int32)
	f.Int32VarP(p, name, shorthand, value, usage)
	return &TypedFlag[int32]{ptr: p}
}

// TypedInt64 creates a type-safe int64 flag.
func TypedInt64(f *FlagSet, name, shorthand string, value int64, usage string) *TypedFlag[int64] {
	p := new(int64)
	f.Int64VarP(p, name, shorthand, value, usage)
	return &TypedFlag[int64]{ptr: p}
}

// TypedUint creates a type-safe uint flag.
func TypedUint(f *FlagSet, name, shorthand string, value uint, usage string) *TypedFlag[uint] {
	p := new(uint)
	f.UintVarP(p, name, shorthand, value, usage)
	return &TypedFlag[uint]{ptr: p}
}

// TypedUint8 creates a type-safe uint8 flag.
func TypedUint8(f *FlagSet, name, shorthand string, value uint8, usage string) *TypedFlag[uint8] {
	p := new(uint8)
	f.Uint8VarP(p, name, shorthand, value, usage)
	return &TypedFlag[uint8]{ptr: p}
}

// TypedUint16 creates a type-safe uint16 flag.
func TypedUint16(f *FlagSet, name, shorthand string, value uint16, usage string) *TypedFlag[uint16] {
	p := new(uint16)
	f.Uint16VarP(p, name, shorthand, value, usage)
	return &TypedFlag[uint16]{ptr: p}
}

// TypedUint32 creates a type-safe uint32 flag.
func TypedUint32(f *FlagSet, name, shorthand string, value uint32, usage string) *TypedFlag[uint32] {
	p := new(uint32)
	f.Uint32VarP(p, name, shorthand, value, usage)
	return &TypedFlag[uint32]{ptr: p}
}

// TypedUint64 creates a type-safe uint64 flag.
func TypedUint64(f *FlagSet, name, shorthand string, value uint64, usage string) *TypedFlag[uint64] {
	p := new(uint64)
	f.Uint64VarP(p, name, shorthand, value, usage)
	return &TypedFlag[uint64]{ptr: p}
}

// TypedFloat32 creates a type-safe float32 flag.
func TypedFloat32(f *FlagSet, name, shorthand string, value float32, usage string) *TypedFlag[float32] {
	p := new(float32)
	f.Float32VarP(p, name, shorthand, value, usage)
	return &TypedFlag[float32]{ptr: p}
}

// TypedFloat64 creates a type-safe float64 flag.
func TypedFloat64(f *FlagSet, name, shorthand string, value float64, usage string) *TypedFlag[float64] {
	p := new(float64)
	f.Float64VarP(p, name, shorthand, value, usage)
	return &TypedFlag[float64]{ptr: p}
}

// TypedDuration creates a type-safe duration flag.
func TypedDuration(f *FlagSet, name, shorthand string, value time.Duration, usage string) *TypedFlag[time.Duration] {
	p := new(time.Duration)
	f.DurationVarP(p, name, shorthand, value, usage)
	return &TypedFlag[time.Duration]{ptr: p}
}

// TypedIP creates a type-safe IP flag.
func TypedIP(f *FlagSet, name, shorthand string, value net.IP, usage string) *TypedFlag[net.IP] {
	p := new(net.IP)
	f.IPVarP(p, name, shorthand, value, usage)
	return &TypedFlag[net.IP]{ptr: p}
}

// TypedStringSlice creates a type-safe string slice flag.
func TypedStringSlice(f *FlagSet, name, shorthand string, value []string, usage string) *TypedFlag[[]string] {
	p := new([]string)
	f.StringSliceVarP(p, name, shorthand, value, usage)
	return &TypedFlag[[]string]{ptr: p}
}

// TypedIntSlice creates a type-safe int slice flag.
func TypedIntSlice(f *FlagSet, name, shorthand string, value []int, usage string) *TypedFlag[[]int] {
	p := new([]int)
	f.IntSliceVarP(p, name, shorthand, value, usage)
	return &TypedFlag[[]int]{ptr: p}
}
