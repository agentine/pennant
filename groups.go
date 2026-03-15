package pennant

// flagGroup defines a constraint between a set of flags.
type flagGroup struct {
	flags []string
}

// MarkFlagsMutuallyExclusive marks a set of flags as mutually exclusive.
// If any flag in the group is set, none of the others may be set.
func (f *FlagSet) MarkFlagsMutuallyExclusive(flagNames ...string) {
	f.mutuallyExclusive = append(f.mutuallyExclusive, flagGroup{flags: flagNames})
}

// MarkFlagsRequiredTogether marks a set of flags as required together.
// If any flag in the group is set, all must be set.
func (f *FlagSet) MarkFlagsRequiredTogether(flagNames ...string) {
	f.requiredTogether = append(f.requiredTogether, flagGroup{flags: flagNames})
}

// validateFlagGroups checks all group constraints after parsing.
func (f *FlagSet) validateFlagGroups() error {
	for _, group := range f.mutuallyExclusive {
		var set []string
		for _, name := range group.flags {
			if f.Changed(name) {
				set = append(set, name)
			}
		}
		if len(set) > 1 {
			return &ErrMutuallyExclusive{FlagNames: group.flags}
		}
	}

	for _, group := range f.requiredTogether {
		anySet := false
		allSet := true
		for _, name := range group.flags {
			if f.Changed(name) {
				anySet = true
			} else {
				allSet = false
			}
		}
		if anySet && !allSet {
			return &ErrRequiredTogether{FlagNames: group.flags}
		}
	}

	return nil
}
