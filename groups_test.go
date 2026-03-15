package pennant

import (
	"bytes"
	"strings"
	"testing"
)

func TestMutuallyExclusive(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.SetOutput(&bytes.Buffer{})
	f.String("json", "", "output json")
	f.String("yaml", "", "output yaml")
	f.MarkFlagsMutuallyExclusive("json", "yaml")

	err := f.Parse([]string{"--json=true", "--yaml=true"})
	if err == nil {
		t.Fatal("expected error for mutually exclusive flags")
	}
	if _, ok := err.(*ErrMutuallyExclusive); !ok {
		t.Errorf("expected ErrMutuallyExclusive, got %T: %v", err, err)
	}
}

func TestMutuallyExclusiveOneSet(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.String("json", "", "output json")
	f.String("yaml", "", "output yaml")
	f.MarkFlagsMutuallyExclusive("json", "yaml")

	err := f.Parse([]string{"--json=true"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestMutuallyExclusiveNoneSet(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.String("json", "", "output json")
	f.String("yaml", "", "output yaml")
	f.MarkFlagsMutuallyExclusive("json", "yaml")

	err := f.Parse([]string{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestRequiredTogether(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.SetOutput(&bytes.Buffer{})
	f.String("user", "", "username")
	f.String("pass", "", "password")
	f.MarkFlagsRequiredTogether("user", "pass")

	err := f.Parse([]string{"--user=admin"})
	if err == nil {
		t.Fatal("expected error when only one of required-together flags is set")
	}
	if _, ok := err.(*ErrRequiredTogether); !ok {
		t.Errorf("expected ErrRequiredTogether, got %T: %v", err, err)
	}
}

func TestRequiredTogetherBothSet(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.String("user", "", "username")
	f.String("pass", "", "password")
	f.MarkFlagsRequiredTogether("user", "pass")

	err := f.Parse([]string{"--user=admin", "--pass=secret"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestRequiredTogetherNoneSet(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.String("user", "", "username")
	f.String("pass", "", "password")
	f.MarkFlagsRequiredTogether("user", "pass")

	err := f.Parse([]string{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestFlagUsages(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.StringP("name", "n", "", "your name")
	f.BoolP("verbose", "v", false, "verbose output")
	f.Int("count", 0, "number of items")

	usage := f.FlagUsages()
	if !strings.Contains(usage, "--name") {
		t.Error("expected --name in usage")
	}
	if !strings.Contains(usage, "-n") {
		t.Error("expected -n shorthand in usage")
	}
	if !strings.Contains(usage, "--verbose") {
		t.Error("expected --verbose in usage")
	}
	if !strings.Contains(usage, "--count") {
		t.Error("expected --count in usage")
	}
}

func TestFlagUsagesHidesHidden(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.String("visible", "", "shown")
	f.String("hidden", "", "not shown")
	f.MarkHidden("hidden")

	usage := f.FlagUsages()
	if !strings.Contains(usage, "visible") {
		t.Error("expected visible flag in usage")
	}
	if strings.Contains(usage, "hidden") {
		t.Error("hidden flag should not appear in usage")
	}
}

func TestFlagUsagesSorted(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.SortFlags = true
	f.String("zebra", "", "")
	f.String("alpha", "", "")
	f.String("middle", "", "")

	usage := f.FlagUsages()
	alphaIdx := strings.Index(usage, "alpha")
	middleIdx := strings.Index(usage, "middle")
	zebraIdx := strings.Index(usage, "zebra")
	if alphaIdx > middleIdx || middleIdx > zebraIdx {
		t.Error("expected flags sorted alphabetically")
	}
}

func TestFlagUsagesUnsorted(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.SortFlags = false
	f.String("zebra", "", "")
	f.String("alpha", "", "")

	usage := f.FlagUsages()
	zebraIdx := strings.Index(usage, "zebra")
	alphaIdx := strings.Index(usage, "alpha")
	if zebraIdx > alphaIdx {
		t.Error("expected flags in definition order when SortFlags=false")
	}
}

func TestPrintDefaults(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	var buf bytes.Buffer
	f.SetOutput(&buf)
	f.String("name", "world", "your name")
	f.PrintDefaults()

	output := buf.String()
	if !strings.Contains(output, "--name") {
		t.Error("expected --name in PrintDefaults output")
	}
	if !strings.Contains(output, `"world"`) {
		t.Error("expected quoted default value for string flag")
	}
}

func TestUnquoteUsage(t *testing.T) {
	var s string
	flag := &Flag{
		Name:  "output",
		Usage: "write to `file`",
		Value: newStringValue("", &s),
	}
	name, usage := UnquoteUsage(flag)
	if name != "file" {
		t.Errorf("expected name 'file', got '%s'", name)
	}
	if usage != "write to file" {
		t.Errorf("expected 'write to file', got '%s'", usage)
	}
}

func TestFlagUsagesWrapped(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.String("output", "", "the output file path for writing results to disk")

	usage := f.FlagUsagesWrapped(40)
	if !strings.Contains(usage, "--output") {
		t.Error("expected --output in wrapped usage")
	}
}
