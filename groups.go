package pennant

// flagGroup defines a constraint between a set of flags.
type flagGroup struct {
	flags []string
}

// mutuallyExclusiveGroups tracks groups of flags where only one may be set.
// requiredTogetherGroups tracks groups where if any is set, all must be set.
// These are stored on FlagSet and validated after parsing.
