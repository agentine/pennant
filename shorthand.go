package pennant

// Shorthand flag resolution is handled within the Parse method.
// This file contains shorthand-related lookup utilities.

// ShorthandLookup returns the Flag for a one-letter shorthand, or nil.
func (f *FlagSet) ShorthandLookup(name string) *Flag {
	if len(name) != 1 {
		return nil
	}
	if f.shorthands == nil {
		return nil
	}
	return f.shorthands[name[0]]
}
