package pennant

import (
	"bytes"
	"encoding/hex"
	"net"
	"strings"
	"testing"
	"time"
)

// -- GetXxx method tests

func TestGetInt8(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.Int8("val", 10, "")
	v, err := f.GetInt8("val")
	if err != nil || v != 10 {
		t.Errorf("expected 10, got %d, err=%v", v, err)
	}
	_, err = f.GetInt8("noexist")
	if err == nil {
		t.Error("expected error for nonexistent flag")
	}
}

func TestGetInt16(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.Int16("val", 1000, "")
	v, err := f.GetInt16("val")
	if err != nil || v != 1000 {
		t.Errorf("expected 1000, got %d, err=%v", v, err)
	}
}

func TestGetInt32(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.Int32("val", 100000, "")
	v, err := f.GetInt32("val")
	if err != nil || v != 100000 {
		t.Errorf("expected 100000, got %d, err=%v", v, err)
	}
}

func TestGetInt64(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.Int64("val", 9999999999, "")
	v, err := f.GetInt64("val")
	if err != nil || v != 9999999999 {
		t.Errorf("expected 9999999999, got %d, err=%v", v, err)
	}
}

func TestGetUint(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.Uint("val", 42, "")
	v, err := f.GetUint("val")
	if err != nil || v != 42 {
		t.Errorf("expected 42, got %d, err=%v", v, err)
	}
}

func TestGetUint8(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.Uint8("val", 255, "")
	v, err := f.GetUint8("val")
	if err != nil || v != 255 {
		t.Errorf("expected 255, got %d, err=%v", v, err)
	}
}

func TestGetUint16(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.Uint16("val", 65535, "")
	v, err := f.GetUint16("val")
	if err != nil || v != 65535 {
		t.Errorf("expected 65535, got %d, err=%v", v, err)
	}
}

func TestGetUint32(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.Uint32("val", 4294967295, "")
	v, err := f.GetUint32("val")
	if err != nil || v != 4294967295 {
		t.Errorf("expected 4294967295, got %d, err=%v", v, err)
	}
}

func TestGetUint64(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.Uint64("val", 18446744073709551615, "")
	v, err := f.GetUint64("val")
	if err != nil || v != 18446744073709551615 {
		t.Errorf("expected max uint64, got %d, err=%v", v, err)
	}
}

func TestGetFloat32(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.Float32("val", 3.14, "")
	v, err := f.GetFloat32("val")
	if err != nil || v < 3.13 || v > 3.15 {
		t.Errorf("expected ~3.14, got %f, err=%v", v, err)
	}
}

func TestGetFloat64(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.Float64("val", 2.718, "")
	v, err := f.GetFloat64("val")
	if err != nil || v != 2.718 {
		t.Errorf("expected 2.718, got %f, err=%v", v, err)
	}
}

func TestGetDuration(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.Duration("val", 5*time.Second, "")
	v, err := f.GetDuration("val")
	if err != nil || v != 5*time.Second {
		t.Errorf("expected 5s, got %v, err=%v", v, err)
	}
}

func TestGetCount(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.CountP("verbose", "v", "")
	f.Parse([]string{"-v", "-v"})
	v, err := f.GetCount("verbose")
	if err != nil || v != 2 {
		t.Errorf("expected 2, got %d, err=%v", v, err)
	}
}

func TestGetIP(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.IP("addr", net.IPv4(127, 0, 0, 1), "")
	v, err := f.GetIP("addr")
	if err != nil || !v.Equal(net.IPv4(127, 0, 0, 1)) {
		t.Errorf("expected 127.0.0.1, got %v, err=%v", v, err)
	}
}

func TestGetIPNet(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	_, expected, _ := net.ParseCIDR("10.0.0.0/8")
	f.IPNet("net", *expected, "")
	v, err := f.GetIPNet("net")
	if err != nil || v.String() != expected.String() {
		t.Errorf("expected %v, got %v, err=%v", expected, v, err)
	}
}

func TestGetIPMask(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	mask := net.IPv4Mask(255, 255, 255, 0)
	f.IPMask("mask", mask, "")
	v, err := f.GetIPMask("mask")
	if err != nil || v.String() != mask.String() {
		t.Errorf("expected %v, got %v, err=%v", mask, v, err)
	}
}

func TestGetBytesHex(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	data, _ := hex.DecodeString("deadbeef")
	f.BytesHex("data", data, "")
	v, err := f.GetBytesHex("data")
	if err != nil || hex.EncodeToString(v) != "deadbeef" {
		t.Errorf("expected deadbeef, got %s, err=%v", hex.EncodeToString(v), err)
	}
}

func TestGetBytesBase64(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.BytesBase64("data", []byte("hello"), "")
	v, err := f.GetBytesBase64("data")
	if err != nil || string(v) != "hello" {
		t.Errorf("expected hello, got %s, err=%v", string(v), err)
	}
}

func TestGetStringSlice(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.StringSlice("tags", []string{"a", "b"}, "")
	v, err := f.GetStringSlice("tags")
	if err != nil || len(v) != 2 {
		t.Errorf("expected 2 items, got %v, err=%v", v, err)
	}
}

func TestGetStringArray(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.StringArray("tags", []string{"x"}, "")
	v, err := f.GetStringArray("tags")
	if err != nil || len(v) != 1 {
		t.Errorf("expected 1 item, got %v, err=%v", v, err)
	}
}

func TestGetIntSlice(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.IntSlice("nums", []int{1, 2}, "")
	v, err := f.GetIntSlice("nums")
	if err != nil || len(v) != 2 {
		t.Errorf("expected 2 items, got %v, err=%v", v, err)
	}
}

func TestGetFloat64Slice(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.Float64Slice("nums", []float64{1.1}, "")
	v, err := f.GetFloat64Slice("nums")
	if err != nil || len(v) != 1 {
		t.Errorf("expected 1 item, got %v, err=%v", v, err)
	}
}

func TestGetBoolSlice(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.BoolSlice("flags", []bool{true}, "")
	v, err := f.GetBoolSlice("flags")
	if err != nil || len(v) != 1 || !v[0] {
		t.Errorf("expected [true], got %v, err=%v", v, err)
	}
}

func TestGetDurationSlice(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.DurationSlice("durs", []time.Duration{time.Second}, "")
	v, err := f.GetDurationSlice("durs")
	if err != nil || len(v) != 1 {
		t.Errorf("expected 1 item, got %v, err=%v", v, err)
	}
}

func TestGetIPSlice(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.IPSlice("addrs", []net.IP{net.IPv4(1, 2, 3, 4)}, "")
	v, err := f.GetIPSlice("addrs")
	if err != nil || len(v) != 1 {
		t.Errorf("expected 1 item, got %v, err=%v", v, err)
	}
}

func TestGetStringToString(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.StringToString("labels", map[string]string{"a": "1"}, "")
	v, err := f.GetStringToString("labels")
	if err != nil || v["a"] != "1" {
		t.Errorf("expected {a:1}, got %v, err=%v", v, err)
	}
}

func TestGetStringToInt(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.StringToInt("counts", map[string]int{"x": 5}, "")
	v, err := f.GetStringToInt("counts")
	if err != nil || v["x"] != 5 {
		t.Errorf("expected {x:5}, got %v, err=%v", v, err)
	}
}

func TestGetStringToInt64(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.StringToInt64("sizes", map[string]int64{"big": 99}, "")
	v, err := f.GetStringToInt64("sizes")
	if err != nil || v["big"] != 99 {
		t.Errorf("expected {big:99}, got %v, err=%v", v, err)
	}
}

// -- GetXxx type mismatch errors

func TestGetIntTypeMismatch(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.String("name", "", "")
	_, err := f.GetInt("name")
	if err == nil {
		t.Error("expected error for type mismatch")
	}
}

func TestGetBoolTypeMismatch(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.Int("val", 0, "")
	// GetBool checks string representation, so this actually works
	// but returns false since "0" != "true"
	v, err := f.GetBool("val")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if v {
		t.Error("expected false")
	}
}

// -- Error message tests

func TestErrUnknownFlagMessage(t *testing.T) {
	err := &ErrUnknownFlag{FlagName: "foo"}
	if !strings.Contains(err.Error(), "foo") {
		t.Errorf("expected 'foo' in error message: %s", err.Error())
	}
}

func TestErrUnknownShorthandMessage(t *testing.T) {
	err := &ErrUnknownShorthand{Shorthand: 'x'}
	if !strings.Contains(err.Error(), "x") {
		t.Errorf("expected 'x' in error message: %s", err.Error())
	}
}

func TestErrParseErrorMessage(t *testing.T) {
	err := &ErrParseError{FlagName: "count", Value: "abc", Type: "int", Err: nil}
	msg := err.Error()
	if !strings.Contains(msg, "count") || !strings.Contains(msg, "abc") {
		t.Errorf("expected flag info in error: %s", msg)
	}
}

func TestErrParseErrorUnwrap(t *testing.T) {
	inner := &ErrNoValue{FlagName: "x"}
	err := &ErrParseError{Err: inner}
	if err.Unwrap() != inner {
		t.Error("Unwrap should return inner error")
	}
}

func TestErrNoValueMessage(t *testing.T) {
	err := &ErrNoValue{FlagName: "name"}
	if !strings.Contains(err.Error(), "name") {
		t.Errorf("expected 'name' in error: %s", err.Error())
	}
}

func TestErrMutuallyExclusiveMessage(t *testing.T) {
	err := &ErrMutuallyExclusive{FlagNames: []string{"a", "b"}}
	if !strings.Contains(err.Error(), "a") {
		t.Errorf("expected flag names in error: %s", err.Error())
	}
}

func TestErrRequiredTogetherMessage(t *testing.T) {
	err := &ErrRequiredTogether{FlagNames: []string{"x", "y"}}
	if !strings.Contains(err.Error(), "x") {
		t.Errorf("expected flag names in error: %s", err.Error())
	}
}

// -- IPMask hex parsing

func TestIPMaskHexParsing(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.IPMask("mask", nil, "")
	f.Parse([]string{"--mask=ffffff00"})
	v, _ := f.GetIPMask("mask")
	if v.String() != "ffffff00" {
		expected := net.IPv4Mask(255, 255, 255, 0)
		if v.String() != expected.String() {
			t.Errorf("expected 255.255.255.0, got %v", v)
		}
	}
}

// -- Var and VarP

func TestVar(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var s string
	f.Var(newStringValue("hello", &s), "name", "a name")
	if s != "hello" {
		t.Errorf("expected 'hello', got '%s'", s)
	}
}

func TestVarP(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var s string
	f.VarP(newStringValue("world", &s), "name", "n", "a name")
	f.Parse([]string{"-n", "test"})
	if s != "test" {
		t.Errorf("expected 'test', got '%s'", s)
	}
}

// -- defaultUsage

func TestDefaultUsage(t *testing.T) {
	f := NewFlagSet("myapp", ContinueOnError)
	var buf bytes.Buffer
	f.SetOutput(&buf)
	f.String("name", "", "your name")
	f.defaultUsage()
	out := buf.String()
	if !strings.Contains(out, "myapp") {
		t.Error("expected app name in usage")
	}
	if !strings.Contains(out, "--name") {
		t.Error("expected --name in usage")
	}
}

func TestDefaultUsageNoName(t *testing.T) {
	f := NewFlagSet("", ContinueOnError)
	var buf bytes.Buffer
	f.SetOutput(&buf)
	f.defaultUsage()
	if !strings.Contains(buf.String(), "Usage:") {
		t.Error("expected 'Usage:' in output")
	}
}

// -- GetNormalizeFunc

func TestGetNormalizeFunc(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	if f.GetNormalizeFunc() != nil {
		t.Error("expected nil normalize func")
	}
	f.SetNormalizeFunc(WordSepNormalizeFunc)
	if f.GetNormalizeFunc() == nil {
		t.Error("expected non-nil normalize func")
	}
}

// -- Edge cases

func TestParseBadFlagSyntax(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.SetOutput(&bytes.Buffer{})
	err := f.Parse([]string{"---bad"})
	if err == nil {
		t.Error("expected error for bad flag syntax")
	}
}

func TestParseEqualsInLongFlag(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.SetOutput(&bytes.Buffer{})
	err := f.Parse([]string{"--=bad"})
	if err == nil {
		t.Error("expected error for flag starting with =")
	}
}

func TestCountWithExplicitValue(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	v := f.CountP("verbose", "v", "")
	f.Parse([]string{"--verbose=5"})
	if *v != 5 {
		t.Errorf("expected 5, got %d", *v)
	}
}

func TestParseIntFlagHex(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	v := f.Int("val", 0, "")
	f.Parse([]string{"--val=0xff"})
	if *v != 255 {
		t.Errorf("expected 255, got %d", *v)
	}
}

func TestSliceValueAppendReplace(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.IntSlice("nums", []int{1}, "")
	flag := f.Lookup("nums")
	sv := flag.Value.(SliceValue)

	sv.Append("2")
	got := sv.GetSlice()
	if len(got) != 2 || got[1] != "2" {
		t.Errorf("Append failed: %v", got)
	}

	sv.Replace([]string{"10", "20"})
	got = sv.GetSlice()
	if len(got) != 2 || got[0] != "10" {
		t.Errorf("Replace failed: %v", got)
	}
}

func TestFloat64SliceAppendReplace(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.Float64Slice("nums", []float64{1.0}, "")
	flag := f.Lookup("nums")
	sv := flag.Value.(SliceValue)
	sv.Append("2.5")
	sv.Replace([]string{"3.0"})
	got := sv.GetSlice()
	if len(got) != 1 || got[0] != "3" {
		t.Errorf("expected [3], got %v", got)
	}
}

func TestBoolSliceAppendReplace(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.BoolSlice("flags", nil, "")
	flag := f.Lookup("flags")
	sv := flag.Value.(SliceValue)
	sv.Append("true")
	sv.Replace([]string{"false", "true"})
	got := sv.GetSlice()
	if len(got) != 2 || got[0] != "false" {
		t.Errorf("expected [false true], got %v", got)
	}
}

func TestDurationSliceAppendReplace(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.DurationSlice("durs", nil, "")
	flag := f.Lookup("durs")
	sv := flag.Value.(SliceValue)
	sv.Append("1s")
	sv.Replace([]string{"2m", "3h"})
	got := sv.GetSlice()
	if len(got) != 2 {
		t.Errorf("expected 2, got %d", len(got))
	}
}

func TestIPSliceAppendReplace(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.IPSlice("addrs", nil, "")
	flag := f.Lookup("addrs")
	sv := flag.Value.(SliceValue)
	sv.Append("1.2.3.4")
	sv.Replace([]string{"5.6.7.8"})
	got := sv.GetSlice()
	if len(got) != 1 {
		t.Errorf("expected 1, got %d", len(got))
	}
}

func TestStringArrayAppendReplace(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.StringArray("arr", nil, "")
	flag := f.Lookup("arr")
	sv := flag.Value.(SliceValue)
	sv.Append("a")
	sv.Append("b")
	got := sv.GetSlice()
	if len(got) != 2 {
		t.Errorf("expected 2, got %d", len(got))
	}
	sv.Replace([]string{"x"})
	got = sv.GetSlice()
	if len(got) != 1 || got[0] != "x" {
		t.Errorf("expected [x], got %v", got)
	}
}

// -- Map type String() and edge cases

func TestStringToStringString(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.StringToString("labels", map[string]string{"a": "1"}, "")
	flag := f.Lookup("labels")
	s := flag.Value.String()
	if !strings.Contains(s, "a=1") {
		t.Errorf("expected 'a=1' in String(): %s", s)
	}
}

func TestStringToStringMerge(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	v := f.StringToString("labels", map[string]string{"a": "1"}, "")
	f.Parse([]string{"--labels=b=2", "--labels=c=3"})
	if (*v)["b"] != "2" || (*v)["c"] != "3" {
		t.Errorf("expected merge, got %v", *v)
	}
}

func TestStringToIntString(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.StringToInt("counts", map[string]int{"x": 1}, "")
	flag := f.Lookup("counts")
	s := flag.Value.String()
	if !strings.Contains(s, "x=1") {
		t.Errorf("expected 'x=1' in String(): %s", s)
	}
}

func TestStringToInt64String(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.StringToInt64("sizes", map[string]int64{"a": 99}, "")
	flag := f.Lookup("sizes")
	s := flag.Value.String()
	if !strings.Contains(s, "a=99") {
		t.Errorf("expected 'a=99' in String(): %s", s)
	}
}

// -- Type() method tests

func TestTypeStrings(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		value    Value
	}{
		{"int8", "int8", newInt8Value(0, new(int8))},
		{"int16", "int16", newInt16Value(0, new(int16))},
		{"int32", "int32", newInt32Value(0, new(int32))},
		{"int64", "int64", newInt64Value(0, new(int64))},
		{"uint", "uint", newUintValue(0, new(uint))},
		{"uint8", "uint8", newUint8Value(0, new(uint8))},
		{"uint16", "uint16", newUint16Value(0, new(uint16))},
		{"uint32", "uint32", newUint32Value(0, new(uint32))},
		{"uint64", "uint64", newUint64Value(0, new(uint64))},
		{"float32", "float32", newFloat32Value(0, new(float32))},
		{"float64", "float64", newFloat64Value(0, new(float64))},
		{"duration", "duration", newDurationValue(0, new(time.Duration))},
		{"count", "count", newCountValue(0, new(int))},
		{"ip", "ip", newIPValue(nil, new(net.IP))},
		{"ipNet", "ipNet", newIPNetValue(net.IPNet{}, new(net.IPNet))},
		{"ipMask", "ipMask", newIPMaskValue(nil, new(net.IPMask))},
		{"bytesHex", "bytesHex", newBytesHexValue(nil, new([]byte))},
		{"bytesBase64", "bytesBase64", newBytesBase64Value(nil, new([]byte))},
		{"stringSlice", "stringSlice", newStringSliceValue(nil, new([]string))},
		{"stringArray", "stringArray", newStringArrayValue(nil, new([]string))},
		{"intSlice", "intSlice", newIntSliceValue(nil, new([]int))},
		{"float64Slice", "float64Slice", newFloat64SliceValue(nil, new([]float64))},
		{"boolSlice", "boolSlice", newBoolSliceValue(nil, new([]bool))},
		{"durationSlice", "durationSlice", newDurationSliceValue(nil, new([]time.Duration))},
		{"ipSlice", "ipSlice", newIPSliceValue(nil, new([]net.IP))},
		{"stringToString", "stringToString", newStringToStringValue(nil, new(map[string]string))},
		{"stringToInt", "stringToInt", newStringToIntValue(nil, new(map[string]int))},
		{"stringToInt64", "stringToInt64", newStringToInt64Value(nil, new(map[string]int64))},
	}
	for _, tt := range tests {
		if tt.value.Type() != tt.expected {
			t.Errorf("%s: expected Type()=%q, got %q", tt.name, tt.expected, tt.value.Type())
		}
	}
}

// -- HasAvailableFlags when all hidden

func TestHasAvailableFlagsAllHidden(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.String("hidden", "", "")
	f.MarkHidden("hidden")
	if f.HasAvailableFlags() {
		t.Error("expected HasAvailableFlags() to be false when all flags are hidden")
	}
}

// -- MarkHidden error

func TestMarkHiddenError(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	err := f.MarkHidden("nonexistent")
	if err == nil {
		t.Error("expected error for nonexistent flag")
	}
}

// -- MarkShorthandDeprecated error

func TestMarkShorthandDeprecatedErrors(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	if err := f.MarkShorthandDeprecated("noexist", "msg"); err == nil {
		t.Error("expected error for nonexistent flag")
	}
	f.StringP("flag", "f", "", "")
	if err := f.MarkShorthandDeprecated("flag", ""); err == nil {
		t.Error("expected error for empty message")
	}
}

// -- Nil slice String()

func TestNilSliceString(t *testing.T) {
	v := &stringSliceValue{value: nil}
	if v.String() != "[]" {
		t.Errorf("expected '[]', got '%s'", v.String())
	}
}
