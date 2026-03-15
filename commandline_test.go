package pennant

import (
	"bytes"
	"net"
	"testing"
	"time"
)

// resetCommandLine creates a fresh CommandLine for testing.
func resetCommandLine() {
	CommandLine = NewFlagSet("test", ContinueOnError)
	CommandLine.SetOutput(&bytes.Buffer{})
}

func TestCommandLineStringVar(t *testing.T) {
	resetCommandLine()
	var s string
	StringVar(&s, "name", "default", "name")
	if s != "default" {
		t.Errorf("expected 'default', got '%s'", s)
	}
}

func TestCommandLineStringVarP(t *testing.T) {
	resetCommandLine()
	var s string
	StringVarP(&s, "name", "n", "val", "name")
	if s != "val" {
		t.Errorf("expected 'val', got '%s'", s)
	}
}

func TestCommandLineString(t *testing.T) {
	resetCommandLine()
	p := String("name", "hello", "name")
	if *p != "hello" {
		t.Errorf("expected 'hello', got '%s'", *p)
	}
}

func TestCommandLineStringP(t *testing.T) {
	resetCommandLine()
	p := StringP("name", "n", "world", "name")
	if *p != "world" {
		t.Errorf("expected 'world', got '%s'", *p)
	}
}

func TestCommandLineGetString(t *testing.T) {
	resetCommandLine()
	String("name", "test", "name")
	v, err := GetString("name")
	if err != nil || v != "test" {
		t.Errorf("expected 'test', got '%s', err=%v", v, err)
	}
}

func TestCommandLineBoolVar(t *testing.T) {
	resetCommandLine()
	var b bool
	BoolVar(&b, "flag", true, "")
	if !b {
		t.Error("expected true")
	}
}

func TestCommandLineBoolVarP(t *testing.T) {
	resetCommandLine()
	var b bool
	BoolVarP(&b, "flag", "f", false, "")
	if b {
		t.Error("expected false")
	}
}

func TestCommandLineBool(t *testing.T) {
	resetCommandLine()
	p := Bool("flag", true, "")
	if !*p {
		t.Error("expected true")
	}
}

func TestCommandLineBoolP(t *testing.T) {
	resetCommandLine()
	p := BoolP("flag", "f", false, "")
	if *p {
		t.Error("expected false")
	}
}

func TestCommandLineGetBool(t *testing.T) {
	resetCommandLine()
	Bool("flag", true, "")
	v, err := GetBool("flag")
	if err != nil || !v {
		t.Errorf("expected true, got %v, err=%v", v, err)
	}
}

func TestCommandLineIntVar(t *testing.T) {
	resetCommandLine()
	var i int
	IntVar(&i, "count", 42, "")
	if i != 42 {
		t.Errorf("expected 42, got %d", i)
	}
}

func TestCommandLineIntVarP(t *testing.T) {
	resetCommandLine()
	var i int
	IntVarP(&i, "count", "c", 10, "")
	if i != 10 {
		t.Errorf("expected 10, got %d", i)
	}
}

func TestCommandLineInt(t *testing.T) {
	resetCommandLine()
	p := Int("count", 7, "")
	if *p != 7 {
		t.Errorf("expected 7, got %d", *p)
	}
}

func TestCommandLineIntP(t *testing.T) {
	resetCommandLine()
	p := IntP("count", "c", 99, "")
	if *p != 99 {
		t.Errorf("expected 99, got %d", *p)
	}
}

func TestCommandLineGetInt(t *testing.T) {
	resetCommandLine()
	Int("count", 42, "")
	v, err := GetInt("count")
	if err != nil || v != 42 {
		t.Errorf("expected 42, got %d, err=%v", v, err)
	}
}

func TestCommandLineInt64(t *testing.T) {
	resetCommandLine()
	p := Int64("big", 999, "")
	if *p != 999 {
		t.Errorf("expected 999, got %d", *p)
	}
}

func TestCommandLineInt64Var(t *testing.T) {
	resetCommandLine()
	var v int64
	Int64Var(&v, "big", 100, "")
	if v != 100 {
		t.Errorf("expected 100, got %d", v)
	}
}

func TestCommandLineInt64VarP(t *testing.T) {
	resetCommandLine()
	var v int64
	Int64VarP(&v, "big", "b", 200, "")
	if v != 200 {
		t.Errorf("expected 200, got %d", v)
	}
}

func TestCommandLineInt64P(t *testing.T) {
	resetCommandLine()
	p := Int64P("big", "b", 300, "")
	if *p != 300 {
		t.Errorf("expected 300, got %d", *p)
	}
}

func TestCommandLineFloat64(t *testing.T) {
	resetCommandLine()
	p := Float64("rate", 1.5, "")
	if *p != 1.5 {
		t.Errorf("expected 1.5, got %f", *p)
	}
}

func TestCommandLineFloat64Var(t *testing.T) {
	resetCommandLine()
	var v float64
	Float64Var(&v, "rate", 2.5, "")
	if v != 2.5 {
		t.Errorf("expected 2.5, got %f", v)
	}
}

func TestCommandLineFloat64VarP(t *testing.T) {
	resetCommandLine()
	var v float64
	Float64VarP(&v, "rate", "r", 3.5, "")
	if v != 3.5 {
		t.Errorf("expected 3.5, got %f", v)
	}
}

func TestCommandLineFloat64P(t *testing.T) {
	resetCommandLine()
	p := Float64P("rate", "r", 4.5, "")
	if *p != 4.5 {
		t.Errorf("expected 4.5, got %f", *p)
	}
}

func TestCommandLineDuration(t *testing.T) {
	resetCommandLine()
	p := Duration("timeout", time.Second, "")
	if *p != time.Second {
		t.Errorf("expected 1s, got %v", *p)
	}
}

func TestCommandLineDurationVar(t *testing.T) {
	resetCommandLine()
	var v time.Duration
	DurationVar(&v, "timeout", 2*time.Second, "")
	if v != 2*time.Second {
		t.Errorf("expected 2s, got %v", v)
	}
}

func TestCommandLineDurationVarP(t *testing.T) {
	resetCommandLine()
	var v time.Duration
	DurationVarP(&v, "timeout", "t", 3*time.Second, "")
	if v != 3*time.Second {
		t.Errorf("expected 3s, got %v", v)
	}
}

func TestCommandLineDurationP(t *testing.T) {
	resetCommandLine()
	p := DurationP("timeout", "t", 4*time.Second, "")
	if *p != 4*time.Second {
		t.Errorf("expected 4s, got %v", *p)
	}
}

func TestCommandLineUintVar(t *testing.T) {
	resetCommandLine()
	var v uint
	UintVar(&v, "num", 5, "")
	if v != 5 {
		t.Errorf("expected 5, got %d", v)
	}
}

func TestCommandLineUint(t *testing.T) {
	resetCommandLine()
	p := Uint("num", 10, "")
	if *p != 10 {
		t.Errorf("expected 10, got %d", *p)
	}
}

func TestCommandLineUint64Var(t *testing.T) {
	resetCommandLine()
	var v uint64
	Uint64Var(&v, "num", 100, "")
	if v != 100 {
		t.Errorf("expected 100, got %d", v)
	}
}

func TestCommandLineUint64(t *testing.T) {
	resetCommandLine()
	p := Uint64("num", 200, "")
	if *p != 200 {
		t.Errorf("expected 200, got %d", *p)
	}
}

func TestCommandLineCountVar(t *testing.T) {
	resetCommandLine()
	var v int
	CountVar(&v, "verbose", "")
	if v != 0 {
		t.Errorf("expected 0, got %d", v)
	}
}

func TestCommandLineCountVarP(t *testing.T) {
	resetCommandLine()
	var v int
	CountVarP(&v, "verbose", "v", "")
	if v != 0 {
		t.Errorf("expected 0, got %d", v)
	}
}

func TestCommandLineCount(t *testing.T) {
	resetCommandLine()
	p := Count("verbose", "")
	if *p != 0 {
		t.Errorf("expected 0, got %d", *p)
	}
}

func TestCommandLineCountP(t *testing.T) {
	resetCommandLine()
	p := CountP("verbose", "v", "")
	if *p != 0 {
		t.Errorf("expected 0, got %d", *p)
	}
}

func TestCommandLineIPVar(t *testing.T) {
	resetCommandLine()
	var v net.IP
	IPVar(&v, "addr", net.IPv4(1, 2, 3, 4), "")
	if !v.Equal(net.IPv4(1, 2, 3, 4)) {
		t.Errorf("expected 1.2.3.4, got %v", v)
	}
}

func TestCommandLineIP(t *testing.T) {
	resetCommandLine()
	p := IP("addr", net.IPv4(5, 6, 7, 8), "")
	if !p.Equal(net.IPv4(5, 6, 7, 8)) {
		t.Errorf("expected 5.6.7.8, got %v", *p)
	}
}

func TestCommandLineStringSliceVar(t *testing.T) {
	resetCommandLine()
	var v []string
	StringSliceVar(&v, "tags", []string{"a"}, "")
	if len(v) != 1 {
		t.Errorf("expected 1, got %d", len(v))
	}
}

func TestCommandLineStringSlice(t *testing.T) {
	resetCommandLine()
	p := StringSlice("tags", []string{"a", "b"}, "")
	if len(*p) != 2 {
		t.Errorf("expected 2, got %d", len(*p))
	}
}

func TestCommandLineIntSliceVar(t *testing.T) {
	resetCommandLine()
	var v []int
	IntSliceVar(&v, "nums", []int{1}, "")
	if len(v) != 1 {
		t.Errorf("expected 1, got %d", len(v))
	}
}

func TestCommandLineIntSlice(t *testing.T) {
	resetCommandLine()
	p := IntSlice("nums", []int{1, 2}, "")
	if len(*p) != 2 {
		t.Errorf("expected 2, got %d", len(*p))
	}
}

func TestCommandLineStringToStringVar(t *testing.T) {
	resetCommandLine()
	var v map[string]string
	StringToStringVar(&v, "labels", map[string]string{"a": "1"}, "")
	if v["a"] != "1" {
		t.Errorf("expected {a:1}, got %v", v)
	}
}

func TestCommandLineStringToString(t *testing.T) {
	resetCommandLine()
	p := StringToString("labels", map[string]string{"b": "2"}, "")
	if (*p)["b"] != "2" {
		t.Errorf("expected {b:2}, got %v", *p)
	}
}

func TestCommandLineStringToIntVar(t *testing.T) {
	resetCommandLine()
	var v map[string]int
	StringToIntVar(&v, "counts", map[string]int{"x": 5}, "")
	if v["x"] != 5 {
		t.Errorf("expected {x:5}, got %v", v)
	}
}

func TestCommandLineStringToInt(t *testing.T) {
	resetCommandLine()
	p := StringToInt("counts", map[string]int{"y": 10}, "")
	if (*p)["y"] != 10 {
		t.Errorf("expected {y:10}, got %v", *p)
	}
}

func TestCommandLineBytesHexVar(t *testing.T) {
	resetCommandLine()
	var v []byte
	BytesHexVar(&v, "data", []byte{0xde, 0xad}, "")
	if len(v) != 2 {
		t.Errorf("expected 2 bytes, got %d", len(v))
	}
}

func TestCommandLineBytesHex(t *testing.T) {
	resetCommandLine()
	p := BytesHex("data", []byte{0xbe, 0xef}, "")
	if len(*p) != 2 {
		t.Errorf("expected 2 bytes, got %d", len(*p))
	}
}

func TestCommandLineBytesBase64Var(t *testing.T) {
	resetCommandLine()
	var v []byte
	BytesBase64Var(&v, "data", []byte("hello"), "")
	if string(v) != "hello" {
		t.Errorf("expected hello, got %s", string(v))
	}
}

func TestCommandLineBytesBase64(t *testing.T) {
	resetCommandLine()
	p := BytesBase64("data", []byte("world"), "")
	if string(*p) != "world" {
		t.Errorf("expected world, got %s", string(*p))
	}
}

func TestCommandLineUtilities(t *testing.T) {
	resetCommandLine()
	String("name", "", "name")

	if !HasFlags() {
		t.Error("expected HasFlags() true")
	}
	if !HasAvailableFlags() {
		t.Error("expected HasAvailableFlags() true")
	}

	flag := Lookup("name")
	if flag == nil {
		t.Error("expected non-nil flag")
	}

	SetNormalizeFunc(WordSepNormalizeFunc)
	SetInterspersed(false)
	SetOutput(&bytes.Buffer{})

	MarkHidden("name")
	MarkDeprecated("name", "use something else")
}

func TestCommandLineVisitVisitAll(t *testing.T) {
	resetCommandLine()
	String("a", "", "")
	String("b", "", "")
	Set("a", "val")

	visitCount := 0
	Visit(func(f *Flag) { visitCount++ })
	if visitCount != 1 {
		t.Errorf("expected Visit to hit 1 flag, got %d", visitCount)
	}

	allCount := 0
	VisitAll(func(f *Flag) { allCount++ })
	if allCount != 2 {
		t.Errorf("expected VisitAll to hit 2 flags, got %d", allCount)
	}
}

func TestCommandLinePrintDefaults(t *testing.T) {
	resetCommandLine()
	var buf bytes.Buffer
	SetOutput(&buf)
	String("name", "", "your name")
	PrintDefaults()
	if !bytes.Contains(buf.Bytes(), []byte("name")) {
		t.Error("expected name in output")
	}
}

func TestCommandLineFlagUsages(t *testing.T) {
	resetCommandLine()
	String("name", "", "your name")
	s := FlagUsages()
	if !bytes.Contains([]byte(s), []byte("name")) {
		t.Error("expected name in FlagUsages")
	}
}

func TestCommandLineFlagUsagesWrapped(t *testing.T) {
	resetCommandLine()
	String("name", "", "your name")
	s := FlagUsagesWrapped(80)
	if !bytes.Contains([]byte(s), []byte("name")) {
		t.Error("expected name in FlagUsagesWrapped")
	}
}

func TestCommandLineVarVarP(t *testing.T) {
	resetCommandLine()
	var s string
	Var(newStringValue("a", &s), "one", "")
	var s2 string
	VarP(newStringValue("b", &s2), "two", "t", "")
	if s != "a" || s2 != "b" {
		t.Errorf("expected a,b got %s,%s", s, s2)
	}
}

func TestCommandLineAddFlag(t *testing.T) {
	resetCommandLine()
	var s string
	flag := &Flag{
		Name:  "custom",
		Value: newStringValue("val", &s),
	}
	AddFlag(flag)
	if Lookup("custom") == nil {
		t.Error("expected to find custom flag")
	}
}

func TestCommandLineShorthandLookup(t *testing.T) {
	resetCommandLine()
	StringP("name", "n", "", "")
	if ShorthandLookup("n") == nil {
		t.Error("expected to find shorthand n")
	}
}

func TestCommandLineMarkShorthandDeprecated(t *testing.T) {
	resetCommandLine()
	StringP("name", "n", "", "")
	err := MarkShorthandDeprecated("name", "use --name")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestCommandLineMarkFlagsMutuallyExclusive(t *testing.T) {
	resetCommandLine()
	String("a", "", "")
	String("b", "", "")
	MarkFlagsMutuallyExclusive("a", "b")
}

func TestCommandLineMarkFlagsRequiredTogether(t *testing.T) {
	resetCommandLine()
	String("a", "", "")
	String("b", "", "")
	MarkFlagsRequiredTogether("a", "b")
}

func TestCommandLineSetAnnotation(t *testing.T) {
	resetCommandLine()
	String("name", "", "")
	err := SetAnnotation("name", "key", []string{"val"})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestCommandLineParsed(t *testing.T) {
	resetCommandLine()
	if Parsed() {
		t.Error("expected Parsed() false before parse")
	}
}

func TestCommandLineArgs(t *testing.T) {
	resetCommandLine()
	if len(Args()) != 0 {
		t.Error("expected empty args")
	}
	if NArg() != 0 {
		t.Error("expected NArg() == 0")
	}
	if Arg(0) != "" {
		t.Error("expected empty Arg(0)")
	}
	if NFlag() != 0 {
		t.Error("expected NFlag() == 0")
	}
}
