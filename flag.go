package pennant

// Flag represents a single flag in a FlagSet.
type Flag struct {
	Name                string              // name as it appears on command line
	Shorthand           string              // one-letter abbreviated flag
	Usage               string              // help message
	Value               Value               // value as set
	DefValue            string              // default value (as text)
	Changed             bool                // if the user set the value
	NoOptDefVal         string              // default value when flag is used without a value
	Deprecated          string              // deprecation message; empty means not deprecated
	Hidden              bool                // hidden from help/usage output
	ShorthandDeprecated string              // deprecation message for the shorthand
	Annotations         map[string][]string // used by cobra and other tools
}
