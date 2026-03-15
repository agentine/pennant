package pennant

import "fmt"

// MarkDeprecated marks a flag as deprecated, with a usage message.
// When the deprecated flag is used, a warning is printed to the output.
func (f *FlagSet) MarkDeprecated(name string, usageMessage string) error {
	flag := f.Lookup(name)
	if flag == nil {
		return fmt.Errorf("flag %q does not exist", name)
	}
	if usageMessage == "" {
		return fmt.Errorf("deprecated message for flag %q must be set", name)
	}
	flag.Deprecated = usageMessage
	flag.Hidden = true
	return nil
}

// MarkShorthandDeprecated marks the shorthand of a flag as deprecated,
// with a usage message. The flag itself remains available.
func (f *FlagSet) MarkShorthandDeprecated(name string, usageMessage string) error {
	flag := f.Lookup(name)
	if flag == nil {
		return fmt.Errorf("flag %q does not exist", name)
	}
	if usageMessage == "" {
		return fmt.Errorf("deprecated message for flag %q shorthand must be set", name)
	}
	flag.ShorthandDeprecated = usageMessage
	return nil
}

// MarkHidden marks a flag as hidden. Hidden flags are still functional
// but are not shown in help/usage output.
func (f *FlagSet) MarkHidden(name string) error {
	flag := f.Lookup(name)
	if flag == nil {
		return fmt.Errorf("flag %q does not exist", name)
	}
	flag.Hidden = true
	return nil
}
