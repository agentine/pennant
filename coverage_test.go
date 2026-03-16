package pennant

import (
	"bytes"
	"net"
	"testing"
	"time"
)

// Tests for remaining uncovered XxxVar and TypedXxx functions

func TestInt16Var(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var v int16
	f.Int16Var(&v, "val", 10, "")
	if v != 10 {
		t.Errorf("expected 10, got %d", v)
	}
}

func TestInt32Var(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var v int32
	f.Int32Var(&v, "val", 20, "")
	if v != 20 {
		t.Errorf("expected 20, got %d", v)
	}
}

func TestInt64Var(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var v int64
	f.Int64Var(&v, "val", 30, "")
	if v != 30 {
		t.Errorf("expected 30, got %d", v)
	}
}

func TestUint8Var(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var v uint8
	f.Uint8Var(&v, "val", 100, "")
	if v != 100 {
		t.Errorf("expected 100, got %d", v)
	}
}

func TestUint16Var(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var v uint16
	f.Uint16Var(&v, "val", 1000, "")
	if v != 1000 {
		t.Errorf("expected 1000, got %d", v)
	}
}

func TestUint32Var(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var v uint32
	f.Uint32Var(&v, "val", 10000, "")
	if v != 10000 {
		t.Errorf("expected 10000, got %d", v)
	}
}

func TestUint64Var(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var v uint64
	f.Uint64Var(&v, "val", 100000, "")
	if v != 100000 {
		t.Errorf("expected 100000, got %d", v)
	}
}

func TestFloat32Var(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var v float32
	f.Float32Var(&v, "val", 1.5, "")
	if v != 1.5 {
		t.Errorf("expected 1.5, got %f", v)
	}
}

func TestFloat64Var(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var v float64
	f.Float64Var(&v, "val", 2.5, "")
	if v != 2.5 {
		t.Errorf("expected 2.5, got %f", v)
	}
}

func TestDurationVar(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var v time.Duration
	f.DurationVar(&v, "val", 3*time.Second, "")
	if v != 3*time.Second {
		t.Errorf("expected 3s, got %v", v)
	}
}

func TestIPVar(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var v net.IP
	f.IPVar(&v, "addr", net.IPv4(1, 2, 3, 4), "")
	if !v.Equal(net.IPv4(1, 2, 3, 4)) {
		t.Errorf("expected 1.2.3.4, got %v", v)
	}
}

func TestIPNetVar(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var v net.IPNet
	_, expected, _ := net.ParseCIDR("10.0.0.0/8")
	f.IPNetVar(&v, "net", *expected, "")
	if v.String() != expected.String() {
		t.Errorf("expected %v, got %v", expected, v)
	}
}

func TestIPMaskVar(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var v net.IPMask
	f.IPMaskVar(&v, "mask", net.IPv4Mask(255, 255, 0, 0), "")
	if v.String() != "ffff0000" {
		t.Errorf("expected ffff0000, got %v", v)
	}
}

func TestBytesHexVar(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var v []byte
	f.BytesHexVar(&v, "data", []byte{0xab}, "")
	if len(v) != 1 || v[0] != 0xab {
		t.Errorf("expected [ab], got %v", v)
	}
}

func TestBytesBase64Var(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var v []byte
	f.BytesBase64Var(&v, "data", []byte("hi"), "")
	if string(v) != "hi" {
		t.Errorf("expected 'hi', got '%s'", v)
	}
}

func TestStringSliceVar(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var v []string
	f.StringSliceVar(&v, "tags", []string{"a"}, "")
	if len(v) != 1 {
		t.Errorf("expected 1, got %d", len(v))
	}
}

func TestStringArrayVar(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var v []string
	f.StringArrayVar(&v, "arr", []string{"x"}, "")
	if len(v) != 1 {
		t.Errorf("expected 1, got %d", len(v))
	}
}

func TestIntSliceVar(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var v []int
	f.IntSliceVar(&v, "nums", []int{1, 2}, "")
	if len(v) != 2 {
		t.Errorf("expected 2, got %d", len(v))
	}
}

func TestFloat64SliceVar(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var v []float64
	f.Float64SliceVar(&v, "nums", []float64{1.0}, "")
	if len(v) != 1 {
		t.Errorf("expected 1, got %d", len(v))
	}
}

func TestBoolSliceVar(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var v []bool
	f.BoolSliceVar(&v, "flags", []bool{true}, "")
	if len(v) != 1 {
		t.Errorf("expected 1, got %d", len(v))
	}
}

func TestDurationSliceVar(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var v []time.Duration
	f.DurationSliceVar(&v, "durs", []time.Duration{time.Second}, "")
	if len(v) != 1 {
		t.Errorf("expected 1, got %d", len(v))
	}
}

func TestIPSliceVar(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var v []net.IP
	f.IPSliceVar(&v, "addrs", []net.IP{net.IPv4(1, 2, 3, 4)}, "")
	if len(v) != 1 {
		t.Errorf("expected 1, got %d", len(v))
	}
}

func TestStringToStringVar(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var v map[string]string
	f.StringToStringVar(&v, "labels", map[string]string{"a": "1"}, "")
	if v["a"] != "1" {
		t.Errorf("expected a=1, got %v", v)
	}
}

func TestStringToIntVar(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var v map[string]int
	f.StringToIntVar(&v, "counts", map[string]int{"x": 5}, "")
	if v["x"] != 5 {
		t.Errorf("expected x=5, got %v", v)
	}
}

func TestStringToInt64Var(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var v map[string]int64
	f.StringToInt64Var(&v, "sizes", map[string]int64{"a": 99}, "")
	if v["a"] != 99 {
		t.Errorf("expected a=99, got %v", v)
	}
}

// -- Typed generics coverage

func TestTypedInt8(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	v := TypedInt8(f, "val", "", 10, "")
	if v.Get() != 10 {
		t.Errorf("expected 10, got %d", v.Get())
	}
}

func TestTypedInt16(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	v := TypedInt16(f, "val", "", 100, "")
	if v.Get() != 100 {
		t.Errorf("expected 100, got %d", v.Get())
	}
}

func TestTypedInt32(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	v := TypedInt32(f, "val", "", 1000, "")
	if v.Get() != 1000 {
		t.Errorf("expected 1000, got %d", v.Get())
	}
}

func TestTypedInt64(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	v := TypedInt64(f, "val", "", 10000, "")
	if v.Get() != 10000 {
		t.Errorf("expected 10000, got %d", v.Get())
	}
}

func TestTypedUint(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	v := TypedUint(f, "val", "", 42, "")
	if v.Get() != 42 {
		t.Errorf("expected 42, got %d", v.Get())
	}
}

func TestTypedUint8(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	v := TypedUint8(f, "val", "", 255, "")
	if v.Get() != 255 {
		t.Errorf("expected 255, got %d", v.Get())
	}
}

func TestTypedUint16(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	v := TypedUint16(f, "val", "", 1000, "")
	if v.Get() != 1000 {
		t.Errorf("expected 1000, got %d", v.Get())
	}
}

func TestTypedUint32(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	v := TypedUint32(f, "val", "", 10000, "")
	if v.Get() != 10000 {
		t.Errorf("expected 10000, got %d", v.Get())
	}
}

func TestTypedUint64(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	v := TypedUint64(f, "val", "", 100000, "")
	if v.Get() != 100000 {
		t.Errorf("expected 100000, got %d", v.Get())
	}
}

func TestTypedFloat32(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	v := TypedFloat32(f, "val", "", 1.5, "")
	if v.Get() != 1.5 {
		t.Errorf("expected 1.5, got %f", v.Get())
	}
}

func TestTypedIP(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	v := TypedIP(f, "addr", "", net.IPv4(1, 2, 3, 4), "")
	if !v.Get().Equal(net.IPv4(1, 2, 3, 4)) {
		t.Errorf("expected 1.2.3.4, got %v", v.Get())
	}
}

func TestTypedIntSlice(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	v := TypedIntSlice(f, "nums", "", []int{1, 2}, "")
	if len(v.Get()) != 2 {
		t.Errorf("expected 2, got %d", len(v.Get()))
	}
}

func TestCountVar(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var v int
	f.CountVar(&v, "verbose", "")
	if err := f.Parse([]string{"--verbose", "--verbose"}); err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if v != 2 {
		t.Errorf("expected 2, got %d", v)
	}
}

func TestCountDirect(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	v := f.Count("verbose", "")
	if err := f.Parse([]string{"--verbose"}); err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if *v != 1 {
		t.Errorf("expected 1, got %d", *v)
	}
}

// -- Additional coverage for error branches and edge cases

func TestGetOutputDefault(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	// Without SetOutput, GetOutput returns os.Stderr
	w := f.GetOutput()
	if w == nil {
		t.Error("expected non-nil writer")
	}
}

func TestSetParseErrors(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.Int8P(   "i8",  "", 0, "")
	f.Int16P(  "i16", "", 0, "")
	f.Int32P(  "i32", "", 0, "")
	f.Int64P(  "i64", "", 0, "")
	f.UintP(   "u",   "", 0, "")
	f.Uint8P(  "u8",  "", 0, "")
	f.Uint16P( "u16", "", 0, "")
	f.Uint32P( "u32", "", 0, "")
	f.Uint64P( "u64", "", 0, "")
	f.Float32P("f32", "", 0, "")
	f.Float64P("f64", "", 0, "")
	f.DurationP("dur", "", 0, "")
	f.BoolP(   "b",   "", false, "")
	f.BytesHexP("bh", "", nil, "")
	f.BytesBase64P("bb", "", nil, "")

	// All should fail with invalid values
	badSets := []struct{ name, val string }{
		{"i8", "notnum"}, {"i16", "notnum"}, {"i32", "notnum"}, {"i64", "notnum"},
		{"u", "notnum"}, {"u8", "notnum"}, {"u16", "notnum"}, {"u32", "notnum"}, {"u64", "notnum"},
		{"f32", "notnum"}, {"f64", "notnum"}, {"dur", "notdur"},
		{"bh", "notahex!"}, {"bb", "not!base64!!!"},
	}
	for _, tc := range badSets {
		if err := f.Set(tc.name, tc.val); err == nil {
			t.Errorf("expected error for Set(%s, %s)", tc.name, tc.val)
		}
	}
}

func TestUnquoteUsageEdgeCases(t *testing.T) {
	// Unterminated backtick
	f := NewFlagSet("test", ContinueOnError)
	f.StringP("s", "", "", "use `name")
	flag := f.Lookup("s")
	name, usage := UnquoteUsage(flag)
	if name != "string" {
		t.Errorf("expected 'string' for unterminated backtick, got %q", name)
	}
	if usage != "use `name" {
		t.Errorf("unexpected usage: %q", usage)
	}

	// Type-based names for various types
	tests := []struct {
		typeName string
		expected string
	}{
		{"float64", "float"},
		{"int64", "int"},
		{"uint64", "uint"},
		{"stringSlice", "strings"},
		{"stringArray", "strings"},
		{"intSlice", "ints"},
		{"float64Slice", "floats"},
		{"boolSlice", "bools"},
		{"durationSlice", "durations"},
		{"ipSlice", "ips"},
	}
	for _, tc := range tests {
		flag := &Flag{Value: &fakeValue{typeName: tc.typeName}, Usage: "test"}
		name, _ := UnquoteUsage(flag)
		if name != tc.expected {
			t.Errorf("UnquoteUsage type=%s: expected %q, got %q", tc.typeName, tc.expected, name)
		}
	}
}

type fakeValue struct {
	typeName string
}

func (f *fakeValue) String() string   { return "" }
func (f *fakeValue) Set(string) error { return nil }
func (f *fakeValue) Type() string     { return f.typeName }

func TestParsePanicOnErrorCoverage(t *testing.T) {
	f := NewFlagSet("test", PanicOnError)
	f.SetOutput(&bytes.Buffer{})
	f.StringP("name", "", "", "")
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic")
		}
	}()
	_ = f.Parse([]string{"--unknown"})
}

func TestParseShortArgErrors(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.SetOutput(&bytes.Buffer{})
	f.StringP("name", "n", "", "")
	// Short flag that needs a value but has none
	err := f.Parse([]string{"-n"})
	if err == nil {
		t.Error("expected error for -n without value")
	}
}

func TestParseLongArgBadSyntax(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.SetOutput(&bytes.Buffer{})
	// --=value is bad syntax
	err := f.Parse([]string{"--=value"})
	if err == nil {
		t.Error("expected error for --=value")
	}
}

func TestChangedNonexistent(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	if f.Changed("nonexistent") {
		t.Error("expected false for nonexistent flag")
	}
}

func TestShorthandLookupEmpty(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	flag := f.ShorthandLookup("")
	if flag != nil {
		t.Error("expected nil for empty shorthand")
	}
}

func TestSliceSetErrors(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.IntSliceP("is", "", nil, "")
	f.Float64SliceP("fs", "", nil, "")
	f.BoolSliceP("bs", "", nil, "")
	f.DurationSliceP("ds", "", nil, "")
	f.IPSliceP("ips", "", nil, "")

	badSets := []struct{ name, val string }{
		{"is", "notint"},
		{"fs", "notfloat"},
		{"bs", "notbool"},
		{"ds", "notdur"},
		{"ips", "notanip"},
	}
	for _, tc := range badSets {
		if err := f.Set(tc.name, tc.val); err == nil {
			t.Errorf("expected error for Set(%s, %s)", tc.name, tc.val)
		}
	}
}

func TestMapSetErrors(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.StringToIntP("si", "", nil, "")
	f.StringToInt64P("si64", "", nil, "")

	if err := f.Set("si", "a=notint"); err == nil {
		t.Error("expected error for StringToInt bad value")
	}
	if err := f.Set("si64", "a=notint"); err == nil {
		t.Error("expected error for StringToInt64 bad value")
	}
}

func TestIPParseErrors(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.IPP("addr", "", net.IPv4(0, 0, 0, 0), "")
	f.IPNetP("net", "", net.IPNet{}, "")
	f.IPMaskP("mask", "", net.IPv4Mask(0, 0, 0, 0), "")

	if err := f.Set("addr", "notanip"); err == nil {
		t.Error("expected error for bad IP")
	}
	if err := f.Set("net", "notacidr"); err == nil {
		t.Error("expected error for bad IPNet")
	}
	if err := f.Set("mask", "notamask"); err == nil {
		t.Error("expected error for bad IPMask")
	}
}

func TestSliceAppendErrors(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.IntSliceP("is", "", []int{1}, "")
	f.Float64SliceP("fs", "", []float64{1.0}, "")
	f.BoolSliceP("bs", "", []bool{true}, "")
	f.DurationSliceP("ds", "", []time.Duration{time.Second}, "")
	f.IPSliceP("ips", "", []net.IP{net.IPv4(1, 2, 3, 4)}, "")

	// Test Append with bad values via the SliceValue interface
	appendTests := []string{"is", "fs", "bs", "ds", "ips"}
	for _, name := range appendTests {
		flag := f.Lookup(name)
		sv, ok := flag.Value.(SliceValue)
		if !ok {
			t.Errorf("%s does not implement SliceValue", name)
			continue
		}
		if err := sv.Append("!!!bad!!!"); err == nil {
			t.Errorf("expected Append error for %s", name)
		}
	}
}

func TestSliceReplaceErrors(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.IntSliceP("is", "", []int{1}, "")
	f.Float64SliceP("fs", "", []float64{1.0}, "")
	f.BoolSliceP("bs", "", []bool{true}, "")
	f.DurationSliceP("ds", "", []time.Duration{time.Second}, "")
	f.IPSliceP("ips", "", []net.IP{net.IPv4(1, 2, 3, 4)}, "")

	replaceTests := []string{"is", "fs", "bs", "ds", "ips"}
	for _, name := range replaceTests {
		flag := f.Lookup(name)
		sv := flag.Value.(SliceValue)
		if err := sv.Replace([]string{"!!!bad!!!"}); err == nil {
			t.Errorf("expected Replace error for %s", name)
		}
	}
}

func TestWrapUsageNarrowCols(t *testing.T) {
	var buf bytes.Buffer
	// cols <= indent should not wrap
	wrapUsage(&buf, "some usage text here", 20, 10)
	if buf.Len() == 0 {
		t.Error("expected output from wrapUsage with narrow cols")
	}
}

func TestDefaultUsageNoNameCoverage(t *testing.T) {
	f := NewFlagSet("", ContinueOnError)
	var buf bytes.Buffer
	f.SetOutput(&buf)
	f.defaultUsage()
	if !bytes.Contains(buf.Bytes(), []byte("Usage:")) {
		t.Error("expected 'Usage:' in output")
	}
}

func TestGetSliceFromSliceTypes(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.Float64SliceP("fs", "", []float64{1.0, 2.0}, "")
	f.BoolSliceP("bs", "", []bool{true, false}, "")
	f.DurationSliceP("ds", "", []time.Duration{time.Second}, "")
	f.IPSliceP("ips", "", []net.IP{net.IPv4(1, 2, 3, 4)}, "")

	// GetSlice returns string representation
	for _, name := range []string{"fs", "bs", "ds", "ips"} {
		flag := f.Lookup(name)
		sv := flag.Value.(SliceValue)
		sl := sv.GetSlice()
		if len(sl) == 0 {
			t.Errorf("expected non-empty slice for %s", name)
		}
	}
}

func TestStringSliceSetCSV(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.StringSliceP("ss", "", nil, "")
	// Set with CSV value
	if err := f.Set("ss", "a,b,c"); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	val, _ := f.GetStringSlice("ss")
	if len(val) != 3 {
		t.Errorf("expected 3 items, got %d", len(val))
	}
}
