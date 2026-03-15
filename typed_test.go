package pennant

import (
	"testing"
	"time"
)

func TestTypedString(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	name := TypedString(f, "name", "n", "default", "your name")

	if name.Get() != "default" {
		t.Errorf("expected 'default', got '%s'", name.Get())
	}

	f.Parse([]string{"-n", "hello"})
	if name.Get() != "hello" {
		t.Errorf("expected 'hello', got '%s'", name.Get())
	}
}

func TestTypedBool(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	verbose := TypedBool(f, "verbose", "v", false, "verbose output")

	if verbose.Get() {
		t.Error("expected false")
	}

	f.Parse([]string{"-v"})
	if !verbose.Get() {
		t.Error("expected true")
	}
}

func TestTypedInt(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	count := TypedInt(f, "count", "c", 0, "count")

	f.Parse([]string{"--count=42"})
	if count.Get() != 42 {
		t.Errorf("expected 42, got %d", count.Get())
	}
}

func TestTypedFloat64(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	rate := TypedFloat64(f, "rate", "", 0.0, "rate")

	f.Parse([]string{"--rate=3.14"})
	if rate.Get() != 3.14 {
		t.Errorf("expected 3.14, got %f", rate.Get())
	}
}

func TestTypedDuration(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	timeout := TypedDuration(f, "timeout", "t", time.Second, "timeout")

	f.Parse([]string{"-t", "5m"})
	if timeout.Get() != 5*time.Minute {
		t.Errorf("expected 5m, got %v", timeout.Get())
	}
}

func TestTypedStringSlice(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	tags := TypedStringSlice(f, "tags", "", nil, "tags")

	f.Parse([]string{"--tags=a,b,c"})
	got := tags.Get()
	if len(got) != 3 || got[0] != "a" {
		t.Errorf("expected [a b c], got %v", got)
	}
}

func TestTypedPtr(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	name := TypedString(f, "name", "", "val", "name")

	p := name.Ptr()
	if *p != "val" {
		t.Errorf("expected 'val', got '%s'", *p)
	}
}
