package pennant

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

// -- StringSlice Value

type stringSliceValue struct {
	value   *[]string
	changed bool
}

func newStringSliceValue(val []string, p *[]string) *stringSliceValue {
	ssv := &stringSliceValue{value: p}
	*p = val
	return ssv
}

func (s *stringSliceValue) Set(val string) error {
	v := parseCSV(val)
	if !s.changed {
		*s.value = v
	} else {
		*s.value = append(*s.value, v...)
	}
	s.changed = true
	return nil
}

func (s *stringSliceValue) Append(val string) error {
	*s.value = append(*s.value, val)
	return nil
}

func (s *stringSliceValue) Replace(val []string) error {
	*s.value = val
	return nil
}

func (s *stringSliceValue) GetSlice() []string {
	return *s.value
}

func (s *stringSliceValue) String() string {
	if s.value == nil {
		return "[]"
	}
	return "[" + strings.Join(*s.value, ",") + "]"
}
func (s *stringSliceValue) Type() string { return "stringSlice" }

func (f *FlagSet) StringSliceVarP(p *[]string, name, shorthand string, value []string, usage string) {
	f.VarPF(newStringSliceValue(value, p), name, shorthand, usage)
}
func (f *FlagSet) StringSliceVar(p *[]string, name string, value []string, usage string) {
	f.StringSliceVarP(p, name, "", value, usage)
}
func (f *FlagSet) StringSliceP(name, shorthand string, value []string, usage string) *[]string {
	p := new([]string)
	f.StringSliceVarP(p, name, shorthand, value, usage)
	return p
}
func (f *FlagSet) StringSlice(name string, value []string, usage string) *[]string {
	return f.StringSliceP(name, "", value, usage)
}
func (f *FlagSet) GetStringSlice(name string) ([]string, error) {
	flag := f.Lookup(name)
	if flag == nil {
		return nil, &ErrUnknownFlag{FlagName: name}
	}
	v, ok := flag.Value.(*stringSliceValue)
	if !ok {
		return nil, fmt.Errorf("trying to get stringSlice value of flag of type %s", flag.Value.Type())
	}
	return *v.value, nil
}

// -- StringArray Value (like StringSlice but does not split on comma)

type stringArrayValue struct {
	value   *[]string
	changed bool
}

func newStringArrayValue(val []string, p *[]string) *stringArrayValue {
	sav := &stringArrayValue{value: p}
	*p = val
	return sav
}

func (s *stringArrayValue) Set(val string) error {
	if !s.changed {
		*s.value = []string{val}
	} else {
		*s.value = append(*s.value, val)
	}
	s.changed = true
	return nil
}

func (s *stringArrayValue) Append(val string) error {
	*s.value = append(*s.value, val)
	return nil
}

func (s *stringArrayValue) Replace(val []string) error {
	*s.value = val
	return nil
}

func (s *stringArrayValue) GetSlice() []string {
	return *s.value
}

func (s *stringArrayValue) String() string {
	if s.value == nil {
		return "[]"
	}
	return "[" + strings.Join(*s.value, ",") + "]"
}
func (s *stringArrayValue) Type() string { return "stringArray" }

func (f *FlagSet) StringArrayVarP(p *[]string, name, shorthand string, value []string, usage string) {
	f.VarPF(newStringArrayValue(value, p), name, shorthand, usage)
}
func (f *FlagSet) StringArrayVar(p *[]string, name string, value []string, usage string) {
	f.StringArrayVarP(p, name, "", value, usage)
}
func (f *FlagSet) StringArrayP(name, shorthand string, value []string, usage string) *[]string {
	p := new([]string)
	f.StringArrayVarP(p, name, shorthand, value, usage)
	return p
}
func (f *FlagSet) StringArray(name string, value []string, usage string) *[]string {
	return f.StringArrayP(name, "", value, usage)
}
func (f *FlagSet) GetStringArray(name string) ([]string, error) {
	flag := f.Lookup(name)
	if flag == nil {
		return nil, &ErrUnknownFlag{FlagName: name}
	}
	v, ok := flag.Value.(*stringArrayValue)
	if !ok {
		return nil, fmt.Errorf("trying to get stringArray value of flag of type %s", flag.Value.Type())
	}
	return *v.value, nil
}

// -- IntSlice Value

type intSliceValue struct {
	value   *[]int
	changed bool
}

func newIntSliceValue(val []int, p *[]int) *intSliceValue {
	isv := &intSliceValue{value: p}
	*p = val
	return isv
}

func (s *intSliceValue) Set(val string) error {
	parts := parseCSV(val)
	parsed := make([]int, 0, len(parts))
	for _, part := range parts {
		v, err := strconv.ParseInt(strings.TrimSpace(part), 0, strconv.IntSize)
		if err != nil {
			return err
		}
		parsed = append(parsed, int(v))
	}
	if !s.changed {
		*s.value = parsed
	} else {
		*s.value = append(*s.value, parsed...)
	}
	s.changed = true
	return nil
}

func (s *intSliceValue) Append(val string) error {
	v, err := strconv.ParseInt(strings.TrimSpace(val), 0, strconv.IntSize)
	if err != nil {
		return err
	}
	*s.value = append(*s.value, int(v))
	return nil
}

func (s *intSliceValue) Replace(val []string) error {
	parsed := make([]int, 0, len(val))
	for _, v := range val {
		n, err := strconv.ParseInt(strings.TrimSpace(v), 0, strconv.IntSize)
		if err != nil {
			return err
		}
		parsed = append(parsed, int(n))
	}
	*s.value = parsed
	return nil
}

func (s *intSliceValue) GetSlice() []string {
	out := make([]string, len(*s.value))
	for i, v := range *s.value {
		out[i] = strconv.Itoa(v)
	}
	return out
}

func (s *intSliceValue) String() string {
	if s.value == nil {
		return "[]"
	}
	parts := make([]string, len(*s.value))
	for i, v := range *s.value {
		parts[i] = strconv.Itoa(v)
	}
	return "[" + strings.Join(parts, ",") + "]"
}
func (s *intSliceValue) Type() string { return "intSlice" }

func (f *FlagSet) IntSliceVarP(p *[]int, name, shorthand string, value []int, usage string) {
	f.VarPF(newIntSliceValue(value, p), name, shorthand, usage)
}
func (f *FlagSet) IntSliceVar(p *[]int, name string, value []int, usage string) {
	f.IntSliceVarP(p, name, "", value, usage)
}
func (f *FlagSet) IntSliceP(name, shorthand string, value []int, usage string) *[]int {
	p := new([]int)
	f.IntSliceVarP(p, name, shorthand, value, usage)
	return p
}
func (f *FlagSet) IntSlice(name string, value []int, usage string) *[]int {
	return f.IntSliceP(name, "", value, usage)
}
func (f *FlagSet) GetIntSlice(name string) ([]int, error) {
	flag := f.Lookup(name)
	if flag == nil {
		return nil, &ErrUnknownFlag{FlagName: name}
	}
	v, ok := flag.Value.(*intSliceValue)
	if !ok {
		return nil, fmt.Errorf("trying to get intSlice value of flag of type %s", flag.Value.Type())
	}
	return *v.value, nil
}

// -- Float64Slice Value

type float64SliceValue struct {
	value   *[]float64
	changed bool
}

func newFloat64SliceValue(val []float64, p *[]float64) *float64SliceValue {
	fsv := &float64SliceValue{value: p}
	*p = val
	return fsv
}

func (s *float64SliceValue) Set(val string) error {
	parts := parseCSV(val)
	parsed := make([]float64, 0, len(parts))
	for _, part := range parts {
		v, err := strconv.ParseFloat(strings.TrimSpace(part), 64)
		if err != nil {
			return err
		}
		parsed = append(parsed, v)
	}
	if !s.changed {
		*s.value = parsed
	} else {
		*s.value = append(*s.value, parsed...)
	}
	s.changed = true
	return nil
}

func (s *float64SliceValue) Append(val string) error {
	v, err := strconv.ParseFloat(strings.TrimSpace(val), 64)
	if err != nil {
		return err
	}
	*s.value = append(*s.value, v)
	return nil
}

func (s *float64SliceValue) Replace(val []string) error {
	parsed := make([]float64, 0, len(val))
	for _, v := range val {
		n, err := strconv.ParseFloat(strings.TrimSpace(v), 64)
		if err != nil {
			return err
		}
		parsed = append(parsed, n)
	}
	*s.value = parsed
	return nil
}

func (s *float64SliceValue) GetSlice() []string {
	out := make([]string, len(*s.value))
	for i, v := range *s.value {
		out[i] = strconv.FormatFloat(v, 'g', -1, 64)
	}
	return out
}

func (s *float64SliceValue) String() string {
	if s.value == nil {
		return "[]"
	}
	parts := make([]string, len(*s.value))
	for i, v := range *s.value {
		parts[i] = strconv.FormatFloat(v, 'g', -1, 64)
	}
	return "[" + strings.Join(parts, ",") + "]"
}
func (s *float64SliceValue) Type() string { return "float64Slice" }

func (f *FlagSet) Float64SliceVarP(p *[]float64, name, shorthand string, value []float64, usage string) {
	f.VarPF(newFloat64SliceValue(value, p), name, shorthand, usage)
}
func (f *FlagSet) Float64SliceVar(p *[]float64, name string, value []float64, usage string) {
	f.Float64SliceVarP(p, name, "", value, usage)
}
func (f *FlagSet) Float64SliceP(name, shorthand string, value []float64, usage string) *[]float64 {
	p := new([]float64)
	f.Float64SliceVarP(p, name, shorthand, value, usage)
	return p
}
func (f *FlagSet) Float64Slice(name string, value []float64, usage string) *[]float64 {
	return f.Float64SliceP(name, "", value, usage)
}
func (f *FlagSet) GetFloat64Slice(name string) ([]float64, error) {
	flag := f.Lookup(name)
	if flag == nil {
		return nil, &ErrUnknownFlag{FlagName: name}
	}
	v, ok := flag.Value.(*float64SliceValue)
	if !ok {
		return nil, fmt.Errorf("trying to get float64Slice value of flag of type %s", flag.Value.Type())
	}
	return *v.value, nil
}

// -- BoolSlice Value

type boolSliceValue struct {
	value   *[]bool
	changed bool
}

func newBoolSliceValue(val []bool, p *[]bool) *boolSliceValue {
	bsv := &boolSliceValue{value: p}
	*p = val
	return bsv
}

func (s *boolSliceValue) Set(val string) error {
	parts := parseCSV(val)
	parsed := make([]bool, 0, len(parts))
	for _, part := range parts {
		v, err := strconv.ParseBool(strings.TrimSpace(part))
		if err != nil {
			return err
		}
		parsed = append(parsed, v)
	}
	if !s.changed {
		*s.value = parsed
	} else {
		*s.value = append(*s.value, parsed...)
	}
	s.changed = true
	return nil
}

func (s *boolSliceValue) Append(val string) error {
	v, err := strconv.ParseBool(strings.TrimSpace(val))
	if err != nil {
		return err
	}
	*s.value = append(*s.value, v)
	return nil
}

func (s *boolSliceValue) Replace(val []string) error {
	parsed := make([]bool, 0, len(val))
	for _, v := range val {
		b, err := strconv.ParseBool(strings.TrimSpace(v))
		if err != nil {
			return err
		}
		parsed = append(parsed, b)
	}
	*s.value = parsed
	return nil
}

func (s *boolSliceValue) GetSlice() []string {
	out := make([]string, len(*s.value))
	for i, v := range *s.value {
		out[i] = strconv.FormatBool(v)
	}
	return out
}

func (s *boolSliceValue) String() string {
	if s.value == nil {
		return "[]"
	}
	parts := make([]string, len(*s.value))
	for i, v := range *s.value {
		parts[i] = strconv.FormatBool(v)
	}
	return "[" + strings.Join(parts, ",") + "]"
}
func (s *boolSliceValue) Type() string { return "boolSlice" }

func (f *FlagSet) BoolSliceVarP(p *[]bool, name, shorthand string, value []bool, usage string) {
	f.VarPF(newBoolSliceValue(value, p), name, shorthand, usage)
}
func (f *FlagSet) BoolSliceVar(p *[]bool, name string, value []bool, usage string) {
	f.BoolSliceVarP(p, name, "", value, usage)
}
func (f *FlagSet) BoolSliceP(name, shorthand string, value []bool, usage string) *[]bool {
	p := new([]bool)
	f.BoolSliceVarP(p, name, shorthand, value, usage)
	return p
}
func (f *FlagSet) BoolSlice(name string, value []bool, usage string) *[]bool {
	return f.BoolSliceP(name, "", value, usage)
}
func (f *FlagSet) GetBoolSlice(name string) ([]bool, error) {
	flag := f.Lookup(name)
	if flag == nil {
		return nil, &ErrUnknownFlag{FlagName: name}
	}
	v, ok := flag.Value.(*boolSliceValue)
	if !ok {
		return nil, fmt.Errorf("trying to get boolSlice value of flag of type %s", flag.Value.Type())
	}
	return *v.value, nil
}

// -- DurationSlice Value

type durationSliceValue struct {
	value   *[]time.Duration
	changed bool
}

func newDurationSliceValue(val []time.Duration, p *[]time.Duration) *durationSliceValue {
	dsv := &durationSliceValue{value: p}
	*p = val
	return dsv
}

func (s *durationSliceValue) Set(val string) error {
	parts := parseCSV(val)
	parsed := make([]time.Duration, 0, len(parts))
	for _, part := range parts {
		v, err := time.ParseDuration(strings.TrimSpace(part))
		if err != nil {
			return err
		}
		parsed = append(parsed, v)
	}
	if !s.changed {
		*s.value = parsed
	} else {
		*s.value = append(*s.value, parsed...)
	}
	s.changed = true
	return nil
}

func (s *durationSliceValue) Append(val string) error {
	v, err := time.ParseDuration(strings.TrimSpace(val))
	if err != nil {
		return err
	}
	*s.value = append(*s.value, v)
	return nil
}

func (s *durationSliceValue) Replace(val []string) error {
	parsed := make([]time.Duration, 0, len(val))
	for _, v := range val {
		d, err := time.ParseDuration(strings.TrimSpace(v))
		if err != nil {
			return err
		}
		parsed = append(parsed, d)
	}
	*s.value = parsed
	return nil
}

func (s *durationSliceValue) GetSlice() []string {
	out := make([]string, len(*s.value))
	for i, v := range *s.value {
		out[i] = v.String()
	}
	return out
}

func (s *durationSliceValue) String() string {
	if s.value == nil {
		return "[]"
	}
	parts := make([]string, len(*s.value))
	for i, v := range *s.value {
		parts[i] = v.String()
	}
	return "[" + strings.Join(parts, ",") + "]"
}
func (s *durationSliceValue) Type() string { return "durationSlice" }

func (f *FlagSet) DurationSliceVarP(p *[]time.Duration, name, shorthand string, value []time.Duration, usage string) {
	f.VarPF(newDurationSliceValue(value, p), name, shorthand, usage)
}
func (f *FlagSet) DurationSliceVar(p *[]time.Duration, name string, value []time.Duration, usage string) {
	f.DurationSliceVarP(p, name, "", value, usage)
}
func (f *FlagSet) DurationSliceP(name, shorthand string, value []time.Duration, usage string) *[]time.Duration {
	p := new([]time.Duration)
	f.DurationSliceVarP(p, name, shorthand, value, usage)
	return p
}
func (f *FlagSet) DurationSlice(name string, value []time.Duration, usage string) *[]time.Duration {
	return f.DurationSliceP(name, "", value, usage)
}
func (f *FlagSet) GetDurationSlice(name string) ([]time.Duration, error) {
	flag := f.Lookup(name)
	if flag == nil {
		return nil, &ErrUnknownFlag{FlagName: name}
	}
	v, ok := flag.Value.(*durationSliceValue)
	if !ok {
		return nil, fmt.Errorf("trying to get durationSlice value of flag of type %s", flag.Value.Type())
	}
	return *v.value, nil
}

// -- IPSlice Value

type ipSliceValue struct {
	value   *[]net.IP
	changed bool
}

func newIPSliceValue(val []net.IP, p *[]net.IP) *ipSliceValue {
	isv := &ipSliceValue{value: p}
	*p = val
	return isv
}

func (s *ipSliceValue) Set(val string) error {
	parts := parseCSV(val)
	parsed := make([]net.IP, 0, len(parts))
	for _, part := range parts {
		ip := net.ParseIP(strings.TrimSpace(part))
		if ip == nil {
			return fmt.Errorf("failed to parse IP: %q", part)
		}
		parsed = append(parsed, ip)
	}
	if !s.changed {
		*s.value = parsed
	} else {
		*s.value = append(*s.value, parsed...)
	}
	s.changed = true
	return nil
}

func (s *ipSliceValue) Append(val string) error {
	ip := net.ParseIP(strings.TrimSpace(val))
	if ip == nil {
		return fmt.Errorf("failed to parse IP: %q", val)
	}
	*s.value = append(*s.value, ip)
	return nil
}

func (s *ipSliceValue) Replace(val []string) error {
	parsed := make([]net.IP, 0, len(val))
	for _, v := range val {
		ip := net.ParseIP(strings.TrimSpace(v))
		if ip == nil {
			return fmt.Errorf("failed to parse IP: %q", v)
		}
		parsed = append(parsed, ip)
	}
	*s.value = parsed
	return nil
}

func (s *ipSliceValue) GetSlice() []string {
	out := make([]string, len(*s.value))
	for i, v := range *s.value {
		out[i] = v.String()
	}
	return out
}

func (s *ipSliceValue) String() string {
	if s.value == nil {
		return "[]"
	}
	parts := make([]string, len(*s.value))
	for i, v := range *s.value {
		parts[i] = v.String()
	}
	return "[" + strings.Join(parts, ",") + "]"
}
func (s *ipSliceValue) Type() string { return "ipSlice" }

func (f *FlagSet) IPSliceVarP(p *[]net.IP, name, shorthand string, value []net.IP, usage string) {
	f.VarPF(newIPSliceValue(value, p), name, shorthand, usage)
}
func (f *FlagSet) IPSliceVar(p *[]net.IP, name string, value []net.IP, usage string) {
	f.IPSliceVarP(p, name, "", value, usage)
}
func (f *FlagSet) IPSliceP(name, shorthand string, value []net.IP, usage string) *[]net.IP {
	p := new([]net.IP)
	f.IPSliceVarP(p, name, shorthand, value, usage)
	return p
}
func (f *FlagSet) IPSlice(name string, value []net.IP, usage string) *[]net.IP {
	return f.IPSliceP(name, "", value, usage)
}
func (f *FlagSet) GetIPSlice(name string) ([]net.IP, error) {
	flag := f.Lookup(name)
	if flag == nil {
		return nil, &ErrUnknownFlag{FlagName: name}
	}
	v, ok := flag.Value.(*ipSliceValue)
	if !ok {
		return nil, fmt.Errorf("trying to get ipSlice value of flag of type %s", flag.Value.Type())
	}
	return *v.value, nil
}

// parseCSV splits a string by commas, respecting quoted fields.
func parseCSV(val string) []string {
	if val == "" {
		return nil
	}
	return strings.Split(val, ",")
}
