package pennant

import (
	"encoding/hex"
	"net"
	"testing"
	"time"
)

func TestInt8Flag(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var v int8
	f.Int8Var(&v, "val", 0, "")
	if err := f.Parse([]string{"--val=42"}); err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if v != 42 {
		t.Errorf("expected 42, got %d", v)
	}
	got, err := f.GetInt8("val")
	if err != nil || got != 42 {
		t.Errorf("GetInt8 failed: %v, %d", err, got)
	}
}

func TestInt16Flag(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	v := f.Int16("val", 0, "")
	if err := f.Parse([]string{"--val=1000"}); err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if *v != 1000 {
		t.Errorf("expected 1000, got %d", *v)
	}
}

func TestInt32Flag(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	v := f.Int32("val", 0, "")
	if err := f.Parse([]string{"--val=100000"}); err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if *v != 100000 {
		t.Errorf("expected 100000, got %d", *v)
	}
}

func TestInt64Flag(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	v := f.Int64("val", 0, "")
	if err := f.Parse([]string{"--val=9999999999"}); err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if *v != 9999999999 {
		t.Errorf("expected 9999999999, got %d", *v)
	}
}

func TestUintFlag(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	v := f.Uint("val", 0, "")
	if err := f.Parse([]string{"--val=42"}); err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if *v != 42 {
		t.Errorf("expected 42, got %d", *v)
	}
}

func TestUint8Flag(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	v := f.Uint8("val", 0, "")
	if err := f.Parse([]string{"--val=255"}); err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if *v != 255 {
		t.Errorf("expected 255, got %d", *v)
	}
}

func TestUint16Flag(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	v := f.Uint16("val", 0, "")
	if err := f.Parse([]string{"--val=65535"}); err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if *v != 65535 {
		t.Errorf("expected 65535, got %d", *v)
	}
}

func TestUint32Flag(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	v := f.Uint32("val", 0, "")
	if err := f.Parse([]string{"--val=4294967295"}); err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if *v != 4294967295 {
		t.Errorf("expected 4294967295, got %d", *v)
	}
}

func TestUint64Flag(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	v := f.Uint64("val", 0, "")
	if err := f.Parse([]string{"--val=18446744073709551615"}); err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if *v != 18446744073709551615 {
		t.Errorf("expected max uint64, got %d", *v)
	}
}

func TestFloat32Flag(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	v := f.Float32("val", 0, "")
	if err := f.Parse([]string{"--val=3.14"}); err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if *v < 3.13 || *v > 3.15 {
		t.Errorf("expected ~3.14, got %f", *v)
	}
}

func TestFloat64Flag(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	v := f.Float64("val", 0, "")
	if err := f.Parse([]string{"--val=3.141592653589793"}); err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if *v != 3.141592653589793 {
		t.Errorf("expected pi, got %f", *v)
	}
}

func TestDurationFlag(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	v := f.Duration("val", 0, "")
	if err := f.Parse([]string{"--val=5s"}); err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if *v != 5*time.Second {
		t.Errorf("expected 5s, got %v", *v)
	}
}

func TestCountFlag(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	v := f.CountP("verbose", "v", "verbosity")
	if err := f.Parse([]string{"-v", "-v", "-v"}); err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if *v != 3 {
		t.Errorf("expected 3, got %d", *v)
	}
}

func TestIPFlag(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	v := f.IP("addr", net.IPv4(127, 0, 0, 1), "")
	if err := f.Parse([]string{"--addr=192.168.1.1"}); err != nil {
		t.Fatalf("Parse: %v", err)
	}
	expected := net.ParseIP("192.168.1.1")
	if !v.Equal(expected) {
		t.Errorf("expected 192.168.1.1, got %v", *v)
	}
}

func TestIPNetFlag(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	v := f.IPNet("net", net.IPNet{}, "")
	if err := f.Parse([]string{"--net=10.0.0.0/8"}); err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if v.String() != "10.0.0.0/8" {
		t.Errorf("expected 10.0.0.0/8, got %v", v)
	}
}

func TestIPMaskFlag(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	v := f.IPMask("mask", net.IPv4Mask(255, 255, 255, 0), "")
	if err := f.Parse([]string{"--mask=255.255.0.0"}); err != nil {
		t.Fatalf("Parse: %v", err)
	}
	expected := net.IPv4Mask(255, 255, 0, 0)
	if v.String() != expected.String() {
		t.Errorf("expected %v, got %v", expected, *v)
	}
}

func TestBytesHexFlag(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	v := f.BytesHex("data", nil, "")
	if err := f.Parse([]string{"--data=deadbeef"}); err != nil {
		t.Fatalf("Parse: %v", err)
	}
	expected, _ := hex.DecodeString("deadbeef")
	if hex.EncodeToString(*v) != hex.EncodeToString(expected) {
		t.Errorf("expected deadbeef, got %s", hex.EncodeToString(*v))
	}
}

func TestBytesBase64Flag(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	v := f.BytesBase64("data", nil, "")
	if err := f.Parse([]string{"--data=aGVsbG8="}); err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if string(*v) != "hello" {
		t.Errorf("expected hello, got %s", string(*v))
	}
}

func TestStringSliceFlag(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	v := f.StringSlice("names", nil, "")
	if err := f.Parse([]string{"--names=a,b,c"}); err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if len(*v) != 3 || (*v)[0] != "a" || (*v)[1] != "b" || (*v)[2] != "c" {
		t.Errorf("expected [a b c], got %v", *v)
	}
}

func TestStringSliceMultiple(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	v := f.StringSlice("names", nil, "")
	if err := f.Parse([]string{"--names=a,b", "--names=c"}); err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if len(*v) != 3 {
		t.Errorf("expected 3, got %d: %v", len(*v), *v)
	}
}

func TestStringArrayFlag(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	v := f.StringArray("names", nil, "")
	if err := f.Parse([]string{"--names=hello,world", "--names=foo"}); err != nil {
		t.Fatalf("Parse: %v", err)
	}
	// StringArray does NOT split on comma
	if len(*v) != 2 || (*v)[0] != "hello,world" || (*v)[1] != "foo" {
		t.Errorf("expected [hello,world foo], got %v", *v)
	}
}

func TestIntSliceFlag(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	v := f.IntSlice("nums", nil, "")
	if err := f.Parse([]string{"--nums=1,2,3"}); err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if len(*v) != 3 || (*v)[0] != 1 || (*v)[1] != 2 || (*v)[2] != 3 {
		t.Errorf("expected [1 2 3], got %v", *v)
	}
}

func TestFloat64SliceFlag(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	v := f.Float64Slice("nums", nil, "")
	if err := f.Parse([]string{"--nums=1.1,2.2"}); err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if len(*v) != 2 {
		t.Errorf("expected 2, got %d", len(*v))
	}
}

func TestBoolSliceFlag(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	v := f.BoolSlice("flags", nil, "")
	if err := f.Parse([]string{"--flags=true,false,true"}); err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if len(*v) != 3 || !(*v)[0] || (*v)[1] || !(*v)[2] {
		t.Errorf("expected [true false true], got %v", *v)
	}
}

func TestDurationSliceFlag(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	v := f.DurationSlice("durs", nil, "")
	if err := f.Parse([]string{"--durs=1s,2m"}); err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if len(*v) != 2 || (*v)[0] != time.Second || (*v)[1] != 2*time.Minute {
		t.Errorf("expected [1s 2m0s], got %v", *v)
	}
}

func TestIPSliceFlag(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	v := f.IPSlice("addrs", nil, "")
	if err := f.Parse([]string{"--addrs=1.2.3.4,5.6.7.8"}); err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if len(*v) != 2 {
		t.Errorf("expected 2, got %d", len(*v))
	}
}

func TestStringToStringFlag(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	v := f.StringToString("labels", nil, "")
	if err := f.Parse([]string{"--labels=a=1,b=2"}); err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if (*v)["a"] != "1" || (*v)["b"] != "2" {
		t.Errorf("expected {a:1 b:2}, got %v", *v)
	}
}

func TestStringToIntFlag(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	v := f.StringToInt("counts", nil, "")
	if err := f.Parse([]string{"--counts=apples=5,oranges=3"}); err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if (*v)["apples"] != 5 || (*v)["oranges"] != 3 {
		t.Errorf("expected {apples:5 oranges:3}, got %v", *v)
	}
}

func TestStringToInt64Flag(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	v := f.StringToInt64("sizes", nil, "")
	if err := f.Parse([]string{"--sizes=big=9999999999,small=1"}); err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if (*v)["big"] != 9999999999 || (*v)["small"] != 1 {
		t.Errorf("expected {big:9999999999 small:1}, got %v", *v)
	}
}

func TestSliceValueInterface(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.StringSlice("names", []string{"a"}, "")
	flag := f.Lookup("names")

	sv, ok := flag.Value.(SliceValue)
	if !ok {
		t.Fatal("StringSlice should implement SliceValue")
	}
	if err := sv.Append("b"); err != nil {
		t.Fatalf("Append: %v", err)
	}
	got := sv.GetSlice()
	if len(got) != 2 || got[0] != "a" || got[1] != "b" {
		t.Errorf("expected [a b], got %v", got)
	}
	if err := sv.Replace([]string{"x", "y"}); err != nil {
		t.Fatalf("Replace: %v", err)
	}
	got = sv.GetSlice()
	if len(got) != 2 || got[0] != "x" || got[1] != "y" {
		t.Errorf("expected [x y], got %v", got)
	}
}
