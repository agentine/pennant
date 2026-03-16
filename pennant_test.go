package pennant

import (
	"bytes"
	"testing"
)

func TestNewFlagSet(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	if f.Name() != "test" {
		t.Errorf("expected name 'test', got '%s'", f.Name())
	}
	if f.Parsed() {
		t.Error("expected Parsed() to be false before parsing")
	}
	if f.NArg() != 0 {
		t.Errorf("expected NArg() == 0, got %d", f.NArg())
	}
	if f.NFlag() != 0 {
		t.Errorf("expected NFlag() == 0, got %d", f.NFlag())
	}
	if f.HasFlags() {
		t.Error("expected HasFlags() to be false")
	}
	if f.ArgsLenAtDash() != -1 {
		t.Errorf("expected ArgsLenAtDash() == -1, got %d", f.ArgsLenAtDash())
	}
}

func TestErrorHandlingConstants(t *testing.T) {
	if ContinueOnError != 0 {
		t.Error("ContinueOnError should be 0")
	}
	if ExitOnError != 1 {
		t.Error("ExitOnError should be 1")
	}
	if PanicOnError != 2 {
		t.Error("PanicOnError should be 2")
	}
}

func TestWordSepNormalizeFunc(t *testing.T) {
	result := WordSepNormalizeFunc(nil, "my_flag_name")
	if result != "my-flag-name" {
		t.Errorf("expected 'my-flag-name', got '%s'", result)
	}
}

// -- Long flag tests

func TestParseLongFlagEquals(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var name string
	f.StringVar(&name, "name", "", "a name")

	if err := f.Parse([]string{"--name=hello"}); err != nil {
		t.Fatal(err)
	}
	if name != "hello" {
		t.Errorf("expected 'hello', got '%s'", name)
	}
	if !f.Changed("name") {
		t.Error("expected Changed to be true")
	}
}

func TestParseLongFlagSpace(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var name string
	f.StringVar(&name, "name", "", "a name")

	if err := f.Parse([]string{"--name", "world"}); err != nil {
		t.Fatal(err)
	}
	if name != "world" {
		t.Errorf("expected 'world', got '%s'", name)
	}
}

func TestParseLongFlagEmptyValue(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var name string
	f.StringVar(&name, "name", "default", "a name")

	if err := f.Parse([]string{"--name="}); err != nil {
		t.Fatal(err)
	}
	if name != "" {
		t.Errorf("expected empty string, got '%s'", name)
	}
}

func TestParseBoolFlagNoValue(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var verbose bool
	f.BoolVar(&verbose, "verbose", false, "verbose output")

	if err := f.Parse([]string{"--verbose"}); err != nil {
		t.Fatal(err)
	}
	if !verbose {
		t.Error("expected verbose to be true")
	}
}

func TestParseBoolFlagExplicitTrue(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var verbose bool
	f.BoolVar(&verbose, "verbose", false, "verbose output")

	if err := f.Parse([]string{"--verbose=true"}); err != nil {
		t.Fatal(err)
	}
	if !verbose {
		t.Error("expected verbose to be true")
	}
}

func TestParseBoolFlagExplicitFalse(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var verbose bool
	f.BoolVar(&verbose, "verbose", true, "verbose output")

	if err := f.Parse([]string{"--verbose=false"}); err != nil {
		t.Fatal(err)
	}
	if verbose {
		t.Error("expected verbose to be false")
	}
}

// -- Shorthand tests

func TestParseShortFlag(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var name string
	f.StringVarP(&name, "name", "n", "", "a name")

	if err := f.Parse([]string{"-n", "hello"}); err != nil {
		t.Fatal(err)
	}
	if name != "hello" {
		t.Errorf("expected 'hello', got '%s'", name)
	}
}

func TestParseShortFlagAttached(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var name string
	f.StringVarP(&name, "name", "n", "", "a name")

	if err := f.Parse([]string{"-nhello"}); err != nil {
		t.Fatal(err)
	}
	if name != "hello" {
		t.Errorf("expected 'hello', got '%s'", name)
	}
}

func TestParseShortFlagEquals(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var name string
	f.StringVarP(&name, "name", "n", "", "a name")

	if err := f.Parse([]string{"-n=hello"}); err != nil {
		t.Fatal(err)
	}
	if name != "hello" {
		t.Errorf("expected 'hello', got '%s'", name)
	}
}

func TestParseCombinedShortBools(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var a, b, c bool
	f.BoolVarP(&a, "alpha", "a", false, "alpha flag")
	f.BoolVarP(&b, "beta", "b", false, "beta flag")
	f.BoolVarP(&c, "charlie", "c", false, "charlie flag")

	if err := f.Parse([]string{"-abc"}); err != nil {
		t.Fatal(err)
	}
	if !a || !b || !c {
		t.Errorf("expected all true, got a=%v b=%v c=%v", a, b, c)
	}
}

func TestParseCombinedShortWithValue(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var verbose bool
	var name string
	f.BoolVarP(&verbose, "verbose", "v", false, "verbose")
	f.StringVarP(&name, "name", "n", "", "name")

	if err := f.Parse([]string{"-vn", "hello"}); err != nil {
		t.Fatal(err)
	}
	if !verbose {
		t.Error("expected verbose to be true")
	}
	if name != "hello" {
		t.Errorf("expected 'hello', got '%s'", name)
	}
}

func TestParseCombinedShortWithAttachedValue(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var verbose bool
	var name string
	f.BoolVarP(&verbose, "verbose", "v", false, "verbose")
	f.StringVarP(&name, "name", "n", "", "name")

	if err := f.Parse([]string{"-vnhello"}); err != nil {
		t.Fatal(err)
	}
	if !verbose {
		t.Error("expected verbose to be true")
	}
	if name != "hello" {
		t.Errorf("expected 'hello', got '%s'", name)
	}
}

// -- Terminator tests

func TestParseTerminator(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var name string
	f.StringVar(&name, "name", "", "a name")

	if err := f.Parse([]string{"--", "--name=hello"}); err != nil {
		t.Fatal(err)
	}
	if name != "" {
		t.Errorf("expected empty, got '%s'", name)
	}
	if f.NArg() != 1 {
		t.Errorf("expected 1 arg, got %d", f.NArg())
	}
	if f.Arg(0) != "--name=hello" {
		t.Errorf("expected '--name=hello', got '%s'", f.Arg(0))
	}
	if f.ArgsLenAtDash() != 0 {
		t.Errorf("expected ArgsLenAtDash() == 0, got %d", f.ArgsLenAtDash())
	}
}

func TestParseTerminatorAfterArgs(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var name string
	f.StringVar(&name, "name", "def", "a name")

	if err := f.Parse([]string{"arg1", "--", "arg2"}); err != nil {
		t.Fatal(err)
	}
	if f.ArgsLenAtDash() != 1 {
		t.Errorf("expected ArgsLenAtDash() == 1, got %d", f.ArgsLenAtDash())
	}
	args := f.Args()
	if len(args) != 2 || args[0] != "arg1" || args[1] != "arg2" {
		t.Errorf("expected [arg1 arg2], got %v", args)
	}
}

// -- Interspersed args tests

func TestParseInterspersedArgs(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var name string
	f.StringVar(&name, "name", "", "a name")

	if err := f.Parse([]string{"arg1", "--name=hello", "arg2"}); err != nil {
		t.Fatal(err)
	}
	if name != "hello" {
		t.Errorf("expected 'hello', got '%s'", name)
	}
	args := f.Args()
	if len(args) != 2 || args[0] != "arg1" || args[1] != "arg2" {
		t.Errorf("expected [arg1 arg2], got %v", args)
	}
}

func TestParseNonInterspersed(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.SetInterspersed(false)
	var name string
	f.StringVar(&name, "name", "", "a name")

	if err := f.Parse([]string{"arg1", "--name=hello"}); err != nil {
		t.Fatal(err)
	}
	if name != "" {
		t.Errorf("expected empty (parsing should stop at arg1), got '%s'", name)
	}
	args := f.Args()
	if len(args) != 2 || args[0] != "arg1" || args[1] != "--name=hello" {
		t.Errorf("expected [arg1 --name=hello], got %v", args)
	}
}

// -- Error tests

func TestParseUnknownFlag(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.SetOutput(&bytes.Buffer{})

	err := f.Parse([]string{"--unknown"})
	if err == nil {
		t.Error("expected error for unknown flag")
	}
	if _, ok := err.(*ErrUnknownFlag); !ok {
		t.Errorf("expected ErrUnknownFlag, got %T: %v", err, err)
	}
}

func TestParseUnknownShorthand(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.SetOutput(&bytes.Buffer{})

	err := f.Parse([]string{"-x"})
	if err == nil {
		t.Error("expected error for unknown shorthand")
	}
	if _, ok := err.(*ErrUnknownShorthand); !ok {
		t.Errorf("expected ErrUnknownShorthand, got %T: %v", err, err)
	}
}

func TestParseMissingValue(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.SetOutput(&bytes.Buffer{})
	var name string
	f.StringVar(&name, "name", "", "a name")

	err := f.Parse([]string{"--name"})
	if err == nil {
		t.Error("expected error for missing value")
	}
	if _, ok := err.(*ErrNoValue); !ok {
		t.Errorf("expected ErrNoValue, got %T: %v", err, err)
	}
}

func TestParseMissingShorthandValue(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.SetOutput(&bytes.Buffer{})
	var name string
	f.StringVarP(&name, "name", "n", "", "a name")

	err := f.Parse([]string{"-n"})
	if err == nil {
		t.Error("expected error for missing shorthand value")
	}
	if _, ok := err.(*ErrNoValue); !ok {
		t.Errorf("expected ErrNoValue, got %T: %v", err, err)
	}
}

func TestParsePanicOnError(t *testing.T) {
	f := NewFlagSet("test", PanicOnError)
	f.SetOutput(&bytes.Buffer{})

	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic for unknown flag with PanicOnError")
		}
	}()
	_ = f.Parse([]string{"--unknown"})
}

// -- Multiple flags

func TestParseMultipleFlags(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var name string
	var verbose bool
	var count int
	f.StringVarP(&name, "name", "n", "", "a name")
	f.BoolVarP(&verbose, "verbose", "v", false, "verbose")
	f.IntVarP(&count, "count", "c", 0, "count")

	if err := f.Parse([]string{"--name=hello", "-v", "-c", "42"}); err != nil {
		t.Fatal(err)
	}
	if name != "hello" {
		t.Errorf("expected 'hello', got '%s'", name)
	}
	if !verbose {
		t.Error("expected verbose to be true")
	}
	if count != 42 {
		t.Errorf("expected 42, got %d", count)
	}
	if f.NFlag() != 3 {
		t.Errorf("expected NFlag() == 3, got %d", f.NFlag())
	}
}

// -- Pointer variants

func TestStringP(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	name := f.StringP("name", "n", "default", "a name")

	if *name != "default" {
		t.Errorf("expected 'default', got '%s'", *name)
	}

	if err := f.Parse([]string{"-n", "hello"}); err != nil {
		t.Fatal(err)
	}
	if *name != "hello" {
		t.Errorf("expected 'hello', got '%s'", *name)
	}
}

func TestBoolP(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	verbose := f.BoolP("verbose", "v", false, "verbose")

	if err := f.Parse([]string{"-v"}); err != nil {
		t.Fatal(err)
	}
	if !*verbose {
		t.Error("expected verbose to be true")
	}
}

func TestIntP(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	count := f.IntP("count", "c", 0, "count")

	if err := f.Parse([]string{"--count=99"}); err != nil {
		t.Fatal(err)
	}
	if *count != 99 {
		t.Errorf("expected 99, got %d", *count)
	}
}

// -- Lookup / ShorthandLookup

func TestLookup(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.String("name", "", "a name")

	flag := f.Lookup("name")
	if flag == nil {
		t.Fatal("expected to find flag 'name'")
	}
	if flag.Name != "name" {
		t.Errorf("expected Name='name', got '%s'", flag.Name)
	}

	if f.Lookup("nonexistent") != nil {
		t.Error("expected nil for nonexistent flag")
	}
}

func TestShorthandLookup(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.StringP("name", "n", "", "a name")

	flag := f.ShorthandLookup("n")
	if flag == nil {
		t.Fatal("expected to find shorthand 'n'")
	}
	if flag.Name != "name" {
		t.Errorf("expected Name='name', got '%s'", flag.Name)
	}

	if f.ShorthandLookup("x") != nil {
		t.Error("expected nil for nonexistent shorthand")
	}
	if f.ShorthandLookup("ab") != nil {
		t.Error("expected nil for multi-char shorthand lookup")
	}
}

// -- Set

func TestSetFlag(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var name string
	f.StringVar(&name, "name", "", "a name")

	if err := f.Set("name", "hello"); err != nil {
		t.Fatal(err)
	}
	if name != "hello" {
		t.Errorf("expected 'hello', got '%s'", name)
	}
	if !f.Changed("name") {
		t.Error("expected Changed to be true after Set")
	}
}

func TestSetUnknownFlag(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	err := f.Set("unknown", "val")
	if err == nil {
		t.Error("expected error for unknown flag")
	}
}

// -- Visit / VisitAll

func TestVisit(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.String("name", "", "name")
	f.String("other", "", "other")
	if err := f.Set("name", "hello"); err != nil {
		t.Fatalf("Set: %v", err)
	}

	var visited []string
	f.Visit(func(flag *Flag) {
		visited = append(visited, flag.Name)
	})
	if len(visited) != 1 || visited[0] != "name" {
		t.Errorf("expected [name], got %v", visited)
	}
}

func TestVisitAll(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.String("name", "", "name")
	f.String("other", "", "other")

	var visited []string
	f.VisitAll(func(flag *Flag) {
		visited = append(visited, flag.Name)
	})
	if len(visited) != 2 {
		t.Errorf("expected 2, got %d", len(visited))
	}
}

// -- NormalizeFunc

func TestNormalizeFuncLookup(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.SetNormalizeFunc(WordSepNormalizeFunc)
	f.String("my-flag", "", "a flag")

	flag := f.Lookup("my_flag")
	if flag == nil {
		t.Fatal("expected to find 'my-flag' via normalized 'my_flag'")
	}
	if flag.Name != "my-flag" {
		t.Errorf("expected Name='my-flag', got '%s'", flag.Name)
	}
}

func TestNormalizeFuncParse(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.SetNormalizeFunc(WordSepNormalizeFunc)
	var name string
	f.StringVar(&name, "my-flag", "", "a flag")

	if err := f.Parse([]string{"--my_flag=hello"}); err != nil {
		t.Fatal(err)
	}
	if name != "hello" {
		t.Errorf("expected 'hello', got '%s'", name)
	}
}

// -- Init

func TestInit(t *testing.T) {
	f := NewFlagSet("old", ContinueOnError)
	f.Init("new", PanicOnError)
	if f.Name() != "new" {
		t.Errorf("expected 'new', got '%s'", f.Name())
	}
}

// -- Deprecated

func TestDeprecatedFlagWarning(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var buf bytes.Buffer
	f.SetOutput(&buf)
	var old string
	f.StringVar(&old, "old-flag", "", "old flag")
	if err := f.MarkDeprecated("old-flag", "use --new-flag instead"); err != nil {
		t.Fatalf("MarkDeprecated: %v", err)
	}

	if err := f.Parse([]string{"--old-flag=val"}); err != nil {
		t.Fatal(err)
	}
	if old != "val" {
		t.Errorf("expected 'val', got '%s'", old)
	}
	if !bytes.Contains(buf.Bytes(), []byte("deprecated")) {
		t.Error("expected deprecation warning in output")
	}
}

func TestMarkDeprecatedErrors(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	if err := f.MarkDeprecated("nonexistent", "msg"); err == nil {
		t.Error("expected error for nonexistent flag")
	}

	f.String("flag", "", "usage")
	if err := f.MarkDeprecated("flag", ""); err == nil {
		t.Error("expected error for empty message")
	}
}

// -- HasFlags / HasAvailableFlags

func TestHasAvailableFlags(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.String("visible", "", "visible")
	f.String("hidden", "", "hidden")
	if err := f.MarkHidden("hidden"); err != nil {
		t.Fatalf("MarkHidden: %v", err)
	}

	if !f.HasAvailableFlags() {
		t.Error("expected HasAvailableFlags() to be true")
	}
}

// -- GetBool / GetString / GetInt

func TestGetString(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.String("name", "default", "name")

	val, err := f.GetString("name")
	if err != nil {
		t.Fatal(err)
	}
	if val != "default" {
		t.Errorf("expected 'default', got '%s'", val)
	}

	_, err = f.GetString("nonexistent")
	if err == nil {
		t.Error("expected error for nonexistent flag")
	}
}

func TestGetBool(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.Bool("verbose", false, "verbose")

	val, err := f.GetBool("verbose")
	if err != nil {
		t.Fatal(err)
	}
	if val {
		t.Error("expected false")
	}
}

func TestGetInt(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.Int("count", 42, "count")

	val, err := f.GetInt("count")
	if err != nil {
		t.Fatal(err)
	}
	if val != 42 {
		t.Errorf("expected 42, got %d", val)
	}
}

// -- Edge cases

func TestParseEmptyArgs(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.String("name", "default", "name")

	if err := f.Parse([]string{}); err != nil {
		t.Fatal(err)
	}
	if !f.Parsed() {
		t.Error("expected Parsed() to be true")
	}
	val, _ := f.GetString("name")
	if val != "default" {
		t.Errorf("expected 'default', got '%s'", val)
	}
}

func TestParseBareHyphen(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	if err := f.Parse([]string{"-"}); err != nil {
		t.Fatal(err)
	}
	if f.NArg() != 1 {
		t.Errorf("expected 1 arg, got %d", f.NArg())
	}
	if f.Arg(0) != "-" {
		t.Errorf("expected '-', got '%s'", f.Arg(0))
	}
}

func TestParseIntFlag(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var count int
	f.IntVar(&count, "count", 0, "count")

	if err := f.Parse([]string{"--count=42"}); err != nil {
		t.Fatal(err)
	}
	if count != 42 {
		t.Errorf("expected 42, got %d", count)
	}
}

func TestParseIntFlagInvalid(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.SetOutput(&bytes.Buffer{})
	f.Int("count", 0, "count")

	err := f.Parse([]string{"--count=notanumber"})
	if err == nil {
		t.Error("expected error for invalid int")
	}
}

func TestAnnotation(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.String("name", "", "name")

	if err := f.SetAnnotation("name", "key", []string{"val1", "val2"}); err != nil {
		t.Fatal(err)
	}
	flag := f.Lookup("name")
	vals := flag.Annotations["key"]
	if len(vals) != 2 || vals[0] != "val1" {
		t.Errorf("expected [val1 val2], got %v", vals)
	}

	if err := f.SetAnnotation("nonexistent", "key", nil); err == nil {
		t.Error("expected error for nonexistent flag")
	}
}

func TestArgOutOfBounds(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	if err := f.Parse([]string{"a"}); err != nil {
		t.Fatalf("Parse: %v", err)
	}

	if f.Arg(-1) != "" {
		t.Error("expected empty for negative index")
	}
	if f.Arg(5) != "" {
		t.Error("expected empty for out-of-bounds index")
	}
}

func TestShorthandBoolExplicitValue(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var verbose bool
	f.BoolVarP(&verbose, "verbose", "v", false, "verbose")

	if err := f.Parse([]string{"-v=false"}); err != nil {
		t.Fatal(err)
	}
	if verbose {
		t.Error("expected verbose to be false with -v=false")
	}
}

func TestParseShorthandDeprecatedWarning(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var buf bytes.Buffer
	f.SetOutput(&buf)
	var name string
	f.StringVarP(&name, "name", "n", "", "name")
	if err := f.MarkShorthandDeprecated("name", "use --name instead"); err != nil {
		t.Fatalf("MarkShorthandDeprecated: %v", err)
	}

	if err := f.Parse([]string{"-n", "hello"}); err != nil {
		t.Fatal(err)
	}
	if name != "hello" {
		t.Errorf("expected 'hello', got '%s'", name)
	}
	if !bytes.Contains(buf.Bytes(), []byte("deprecated")) {
		t.Error("expected deprecation warning for shorthand")
	}
}
