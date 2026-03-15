package pennant

import (
	"fmt"
	"strconv"
	"strings"
)

// -- StringToString Value

type stringToStringValue struct {
	value   *map[string]string
	changed bool
}

func newStringToStringValue(val map[string]string, p *map[string]string) *stringToStringValue {
	ssv := &stringToStringValue{value: p}
	*p = val
	return ssv
}

func (s *stringToStringValue) Set(val string) error {
	parsed, err := parseMapStringString(val)
	if err != nil {
		return err
	}
	if !s.changed {
		*s.value = parsed
	} else {
		for k, v := range parsed {
			(*s.value)[k] = v
		}
	}
	s.changed = true
	return nil
}

func parseMapStringString(val string) (map[string]string, error) {
	m := make(map[string]string)
	if val == "" {
		return m, nil
	}
	pairs := strings.Split(val, ",")
	for _, pair := range pairs {
		idx := strings.IndexByte(pair, '=')
		if idx < 0 {
			return nil, fmt.Errorf("%q must be formatted as key=value", pair)
		}
		m[pair[:idx]] = pair[idx+1:]
	}
	return m, nil
}

func (s *stringToStringValue) String() string {
	if s.value == nil {
		return ""
	}
	parts := make([]string, 0, len(*s.value))
	for k, v := range *s.value {
		parts = append(parts, k+"="+v)
	}
	return strings.Join(parts, ",")
}
func (s *stringToStringValue) Type() string { return "stringToString" }

func (f *FlagSet) StringToStringVarP(p *map[string]string, name, shorthand string, value map[string]string, usage string) {
	f.VarPF(newStringToStringValue(value, p), name, shorthand, usage)
}
func (f *FlagSet) StringToStringVar(p *map[string]string, name string, value map[string]string, usage string) {
	f.StringToStringVarP(p, name, "", value, usage)
}
func (f *FlagSet) StringToStringP(name, shorthand string, value map[string]string, usage string) *map[string]string {
	p := &map[string]string{}
	f.StringToStringVarP(p, name, shorthand, value, usage)
	return p
}
func (f *FlagSet) StringToString(name string, value map[string]string, usage string) *map[string]string {
	return f.StringToStringP(name, "", value, usage)
}
func (f *FlagSet) GetStringToString(name string) (map[string]string, error) {
	flag := f.Lookup(name)
	if flag == nil {
		return nil, &ErrUnknownFlag{FlagName: name}
	}
	v, ok := flag.Value.(*stringToStringValue)
	if !ok {
		return nil, fmt.Errorf("trying to get stringToString value of flag of type %s", flag.Value.Type())
	}
	return *v.value, nil
}

// -- StringToInt Value

type stringToIntValue struct {
	value   *map[string]int
	changed bool
}

func newStringToIntValue(val map[string]int, p *map[string]int) *stringToIntValue {
	siv := &stringToIntValue{value: p}
	*p = val
	return siv
}

func (s *stringToIntValue) Set(val string) error {
	parsed, err := parseMapStringInt(val)
	if err != nil {
		return err
	}
	if !s.changed {
		*s.value = parsed
	} else {
		for k, v := range parsed {
			(*s.value)[k] = v
		}
	}
	s.changed = true
	return nil
}

func parseMapStringInt(val string) (map[string]int, error) {
	m := make(map[string]int)
	if val == "" {
		return m, nil
	}
	pairs := strings.Split(val, ",")
	for _, pair := range pairs {
		idx := strings.IndexByte(pair, '=')
		if idx < 0 {
			return nil, fmt.Errorf("%q must be formatted as key=value", pair)
		}
		v, err := strconv.Atoi(pair[idx+1:])
		if err != nil {
			return nil, fmt.Errorf("invalid integer value %q for key %q: %v", pair[idx+1:], pair[:idx], err)
		}
		m[pair[:idx]] = v
	}
	return m, nil
}

func (s *stringToIntValue) String() string {
	if s.value == nil {
		return ""
	}
	parts := make([]string, 0, len(*s.value))
	for k, v := range *s.value {
		parts = append(parts, k+"="+strconv.Itoa(v))
	}
	return strings.Join(parts, ",")
}
func (s *stringToIntValue) Type() string { return "stringToInt" }

func (f *FlagSet) StringToIntVarP(p *map[string]int, name, shorthand string, value map[string]int, usage string) {
	f.VarPF(newStringToIntValue(value, p), name, shorthand, usage)
}
func (f *FlagSet) StringToIntVar(p *map[string]int, name string, value map[string]int, usage string) {
	f.StringToIntVarP(p, name, "", value, usage)
}
func (f *FlagSet) StringToIntP(name, shorthand string, value map[string]int, usage string) *map[string]int {
	p := &map[string]int{}
	f.StringToIntVarP(p, name, shorthand, value, usage)
	return p
}
func (f *FlagSet) StringToInt(name string, value map[string]int, usage string) *map[string]int {
	return f.StringToIntP(name, "", value, usage)
}
func (f *FlagSet) GetStringToInt(name string) (map[string]int, error) {
	flag := f.Lookup(name)
	if flag == nil {
		return nil, &ErrUnknownFlag{FlagName: name}
	}
	v, ok := flag.Value.(*stringToIntValue)
	if !ok {
		return nil, fmt.Errorf("trying to get stringToInt value of flag of type %s", flag.Value.Type())
	}
	return *v.value, nil
}

// -- StringToInt64 Value

type stringToInt64Value struct {
	value   *map[string]int64
	changed bool
}

func newStringToInt64Value(val map[string]int64, p *map[string]int64) *stringToInt64Value {
	siv := &stringToInt64Value{value: p}
	*p = val
	return siv
}

func (s *stringToInt64Value) Set(val string) error {
	parsed, err := parseMapStringInt64(val)
	if err != nil {
		return err
	}
	if !s.changed {
		*s.value = parsed
	} else {
		for k, v := range parsed {
			(*s.value)[k] = v
		}
	}
	s.changed = true
	return nil
}

func parseMapStringInt64(val string) (map[string]int64, error) {
	m := make(map[string]int64)
	if val == "" {
		return m, nil
	}
	pairs := strings.Split(val, ",")
	for _, pair := range pairs {
		idx := strings.IndexByte(pair, '=')
		if idx < 0 {
			return nil, fmt.Errorf("%q must be formatted as key=value", pair)
		}
		v, err := strconv.ParseInt(pair[idx+1:], 0, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid int64 value %q for key %q: %v", pair[idx+1:], pair[:idx], err)
		}
		m[pair[:idx]] = v
	}
	return m, nil
}

func (s *stringToInt64Value) String() string {
	if s.value == nil {
		return ""
	}
	parts := make([]string, 0, len(*s.value))
	for k, v := range *s.value {
		parts = append(parts, k+"="+strconv.FormatInt(v, 10))
	}
	return strings.Join(parts, ",")
}
func (s *stringToInt64Value) Type() string { return "stringToInt64" }

func (f *FlagSet) StringToInt64VarP(p *map[string]int64, name, shorthand string, value map[string]int64, usage string) {
	f.VarPF(newStringToInt64Value(value, p), name, shorthand, usage)
}
func (f *FlagSet) StringToInt64Var(p *map[string]int64, name string, value map[string]int64, usage string) {
	f.StringToInt64VarP(p, name, "", value, usage)
}
func (f *FlagSet) StringToInt64P(name, shorthand string, value map[string]int64, usage string) *map[string]int64 {
	p := &map[string]int64{}
	f.StringToInt64VarP(p, name, shorthand, value, usage)
	return p
}
func (f *FlagSet) StringToInt64(name string, value map[string]int64, usage string) *map[string]int64 {
	return f.StringToInt64P(name, "", value, usage)
}
func (f *FlagSet) GetStringToInt64(name string) (map[string]int64, error) {
	flag := f.Lookup(name)
	if flag == nil {
		return nil, &ErrUnknownFlag{FlagName: name}
	}
	v, ok := flag.Value.(*stringToInt64Value)
	if !ok {
		return nil, fmt.Errorf("trying to get stringToInt64 value of flag of type %s", flag.Value.Type())
	}
	return *v.value, nil
}
