package pennant

import "strings"

// NormalizedName is the result of applying a NormalizeFunc to a flag name.
type NormalizedName string

// NormalizeFunc transforms flag names for matching.
// For example, it can convert underscores to hyphens.
type NormalizeFunc func(f *FlagSet, name string) NormalizedName

// SetNormalizeFunc sets a function to transform flag names.
// All flag names are passed through this function before lookup.
func (f *FlagSet) SetNormalizeFunc(n NormalizeFunc) {
	f.normalizeFunc = n
}

// GetNormalizeFunc returns the normalization function, or nil if none is set.
func (f *FlagSet) GetNormalizeFunc() NormalizeFunc {
	return f.normalizeFunc
}

// normalizeName applies the normalization function if set, otherwise returns the name as-is.
func (f *FlagSet) normalizeName(name string) NormalizedName {
	if f.normalizeFunc != nil {
		return f.normalizeFunc(f, name)
	}
	return NormalizedName(name)
}

// WordSepNormalizeFunc is a NormalizeFunc that converts separators to hyphens.
func WordSepNormalizeFunc(_ *FlagSet, name string) NormalizedName {
	return NormalizedName(strings.ReplaceAll(name, "_", "-"))
}
