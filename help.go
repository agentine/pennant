package pennant

import (
	"bytes"
	"fmt"
	"io"
	"sort"
	"strings"
)

// UnquoteUsage extracts a back-quoted name from the usage string and returns
// it and the un-quoted usage. Given "a]name to `search` for" it returns
// ("search", "a name to search for").
func UnquoteUsage(flag *Flag) (name string, usage string) {
	usage = flag.Usage
	for i := 0; i < len(usage); i++ {
		if usage[i] == '`' {
			for j := i + 1; j < len(usage); j++ {
				if usage[j] == '`' {
					name = usage[i+1 : j]
					usage = usage[:i] + name + usage[j+1:]
					return name, usage
				}
			}
			break
		}
	}

	name = flag.Value.Type()
	switch name {
	case "bool":
		name = ""
	case "float64":
		name = "float"
	case "int64":
		name = "int"
	case "uint64":
		name = "uint"
	case "stringSlice", "stringArray":
		name = "strings"
	case "intSlice":
		name = "ints"
	case "float64Slice":
		name = "floats"
	case "boolSlice":
		name = "bools"
	case "durationSlice":
		name = "durations"
	case "ipSlice":
		name = "ips"
	}
	return
}

// FlagUsages returns a string containing the usage information for all flags
// in the FlagSet.
func (f *FlagSet) FlagUsages() string {
	return f.FlagUsagesWrapped(0)
}

// FlagUsagesWrapped returns a string containing the usage information for all
// flags in the FlagSet, wrapping lines at the specified column (0 = no wrap).
func (f *FlagSet) FlagUsagesWrapped(cols int) string {
	var buf bytes.Buffer

	var flags []*Flag
	f.VisitAll(func(flag *Flag) {
		if flag.Hidden {
			return
		}
		flags = append(flags, flag)
	})

	if f.SortFlags {
		sort.Slice(flags, func(i, j int) bool {
			return flags[i].Name < flags[j].Name
		})
	}

	// Compute max width of the flag column for alignment
	maxLen := 0
	lines := make([]struct {
		flagCol string
		usage   string
	}, len(flags))

	for i, flag := range flags {
		var line strings.Builder
		if flag.Shorthand != "" {
			fmt.Fprintf(&line, "  -%s, --%s", flag.Shorthand, flag.Name)
		} else {
			fmt.Fprintf(&line, "      --%s", flag.Name)
		}
		name, usage := UnquoteUsage(flag)
		if name != "" {
			fmt.Fprintf(&line, " %s", name)
		}
		flagCol := line.String()
		if len(flagCol) > maxLen {
			maxLen = len(flagCol)
		}

		defValue := flag.DefValue
		if flag.Value.Type() == "string" {
			defValue = fmt.Sprintf("%q", defValue)
		}
		if flag.DefValue != "" && flag.DefValue != "0" && flag.DefValue != "false" && flag.DefValue != "[]" {
			usage += fmt.Sprintf(" (default %s)", defValue)
		}

		lines[i].flagCol = flagCol
		lines[i].usage = usage
	}

	for _, l := range lines {
		// Pad flag column to maxLen
		padding := maxLen - len(l.flagCol) + 3
		if padding < 1 {
			padding = 1
		}

		if cols > 0 && len(l.flagCol)+padding+len(l.usage) > cols {
			// Wrap: flag column on one line, usage on next
			fmt.Fprintf(&buf, "%s\n", l.flagCol)
			wrapUsage(&buf, l.usage, maxLen+3, cols)
		} else {
			fmt.Fprintf(&buf, "%s%s%s\n", l.flagCol, strings.Repeat(" ", padding), l.usage)
		}
	}

	return buf.String()
}

// wrapUsage writes usage text with word wrapping, indented to the given offset.
func wrapUsage(w io.Writer, usage string, indent, cols int) {
	if cols <= indent {
		_, _ = fmt.Fprintf(w, "%s%s\n", strings.Repeat(" ", indent), usage)
		return
	}
	width := cols - indent
	prefix := strings.Repeat(" ", indent)
	words := strings.Fields(usage)
	lineLen := 0
	first := true
	for _, word := range words {
		if !first && lineLen+1+len(word) > width {
			_, _ = fmt.Fprintf(w, "\n%s", prefix)
			lineLen = 0
			first = true
		}
		if !first {
			_, _ = fmt.Fprint(w, " ")
			lineLen++
		} else {
			_, _ = fmt.Fprint(w, prefix)
		}
		_, _ = fmt.Fprint(w, word)
		lineLen += len(word)
		first = false
	}
	if !first {
		_, _ = fmt.Fprintln(w)
	}
}

// PrintDefaults prints, to the FlagSet's output writer, the default values
// of all defined flags in the set.
func (f *FlagSet) PrintDefaults() {
	usages := f.FlagUsages()
	_, _ = fmt.Fprint(f.GetOutput(), usages)
}

// defaultUsage is the default Usage function.
func (f *FlagSet) defaultUsage() {
	if f.name == "" {
		_, _ = fmt.Fprintf(f.GetOutput(), "Usage:\n")
	} else {
		_, _ = fmt.Fprintf(f.GetOutput(), "Usage of %s:\n", f.name)
	}
	f.PrintDefaults()
}
