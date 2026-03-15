package pennant

import (
	"bytes"
	"strings"
	"testing"
)

func FuzzParse(f *testing.F) {
	// Seed corpus
	f.Add("--name=hello")
	f.Add("-v")
	f.Add("-abc")
	f.Add("--verbose=true")
	f.Add("--count=42")
	f.Add("--")
	f.Add("-n hello")
	f.Add("--name hello")
	f.Add("")
	f.Add("-")
	f.Add("---bad")
	f.Add("--=bad")
	f.Add("-v -n test --count=5 -- extra1 extra2")
	f.Add("--name=with spaces")
	f.Add("--name=with=equals")
	f.Add("-vntest")

	f.Fuzz(func(t *testing.T, input string) {
		args := strings.Fields(input)

		fs := NewFlagSet("fuzz", ContinueOnError)
		fs.SetOutput(&bytes.Buffer{})

		fs.StringP("name", "n", "", "name")
		fs.BoolP("verbose", "v", false, "verbose")
		fs.IntP("count", "c", 0, "count")
		fs.BoolP("alpha", "a", false, "alpha")
		fs.BoolP("beta", "b", false, "beta")

		// The parser should not panic regardless of input
		_ = fs.Parse(args)
	})
}
