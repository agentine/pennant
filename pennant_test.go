package pennant

import "testing"

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
