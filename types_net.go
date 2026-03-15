package pennant

import (
	"fmt"
	"net"
	"strings"
)

// -- IP Value

type ipValue net.IP

func newIPValue(val net.IP, p *net.IP) *ipValue {
	*p = val
	return (*ipValue)(p)
}

func (ip *ipValue) Set(val string) error {
	parsed := net.ParseIP(strings.TrimSpace(val))
	if parsed == nil {
		return fmt.Errorf("failed to parse IP: %q", val)
	}
	*ip = ipValue(parsed)
	return nil
}

func (ip *ipValue) String() string { return net.IP(*ip).String() }
func (ip *ipValue) Type() string   { return "ip" }

func (f *FlagSet) IPVarP(p *net.IP, name, shorthand string, value net.IP, usage string) {
	f.VarPF(newIPValue(value, p), name, shorthand, usage)
}
func (f *FlagSet) IPVar(p *net.IP, name string, value net.IP, usage string) {
	f.IPVarP(p, name, "", value, usage)
}
func (f *FlagSet) IPP(name, shorthand string, value net.IP, usage string) *net.IP {
	p := new(net.IP)
	f.IPVarP(p, name, shorthand, value, usage)
	return p
}
func (f *FlagSet) IP(name string, value net.IP, usage string) *net.IP {
	return f.IPP(name, "", value, usage)
}
func (f *FlagSet) GetIP(name string) (net.IP, error) {
	flag := f.Lookup(name)
	if flag == nil {
		return nil, &ErrUnknownFlag{FlagName: name}
	}
	v, ok := flag.Value.(*ipValue)
	if !ok {
		return nil, fmt.Errorf("trying to get ip value of flag of type %s", flag.Value.Type())
	}
	return net.IP(*v), nil
}

// -- IPNet Value

type ipNetValue net.IPNet

func newIPNetValue(val net.IPNet, p *net.IPNet) *ipNetValue {
	*p = val
	return (*ipNetValue)(p)
}

func (ipn *ipNetValue) Set(val string) error {
	_, parsed, err := net.ParseCIDR(strings.TrimSpace(val))
	if err != nil {
		return err
	}
	*ipn = ipNetValue(*parsed)
	return nil
}

func (ipn *ipNetValue) String() string {
	n := net.IPNet(*ipn)
	return n.String()
}
func (ipn *ipNetValue) Type() string { return "ipNet" }

func (f *FlagSet) IPNetVarP(p *net.IPNet, name, shorthand string, value net.IPNet, usage string) {
	f.VarPF(newIPNetValue(value, p), name, shorthand, usage)
}
func (f *FlagSet) IPNetVar(p *net.IPNet, name string, value net.IPNet, usage string) {
	f.IPNetVarP(p, name, "", value, usage)
}
func (f *FlagSet) IPNetP(name, shorthand string, value net.IPNet, usage string) *net.IPNet {
	p := new(net.IPNet)
	f.IPNetVarP(p, name, shorthand, value, usage)
	return p
}
func (f *FlagSet) IPNet(name string, value net.IPNet, usage string) *net.IPNet {
	return f.IPNetP(name, "", value, usage)
}
func (f *FlagSet) GetIPNet(name string) (net.IPNet, error) {
	flag := f.Lookup(name)
	if flag == nil {
		return net.IPNet{}, &ErrUnknownFlag{FlagName: name}
	}
	v, ok := flag.Value.(*ipNetValue)
	if !ok {
		return net.IPNet{}, fmt.Errorf("trying to get ipNet value of flag of type %s", flag.Value.Type())
	}
	return net.IPNet(*v), nil
}

// -- IPMask Value

type ipMaskValue net.IPMask

func newIPMaskValue(val net.IPMask, p *net.IPMask) *ipMaskValue {
	*p = val
	return (*ipMaskValue)(p)
}

func (ipm *ipMaskValue) Set(val string) error {
	// Try parsing as dotted decimal (e.g., "255.255.255.0")
	ip := net.ParseIP(strings.TrimSpace(val))
	if ip != nil {
		ip4 := ip.To4()
		if ip4 != nil {
			*ipm = ipMaskValue(net.IPv4Mask(ip4[0], ip4[1], ip4[2], ip4[3]))
			return nil
		}
	}
	// Try parsing as hex
	mask, err := parseHexIPMask(val)
	if err != nil {
		return fmt.Errorf("failed to parse IPMask: %q", val)
	}
	*ipm = ipMaskValue(mask)
	return nil
}

func parseHexIPMask(val string) (net.IPMask, error) {
	val = strings.TrimSpace(val)
	if len(val) != 8 {
		return nil, fmt.Errorf("invalid mask length")
	}
	mask := make(net.IPMask, 4)
	for i := 0; i < 4; i++ {
		b, err := parseHexByte(val[i*2 : i*2+2])
		if err != nil {
			return nil, err
		}
		mask[i] = b
	}
	return mask, nil
}

func parseHexByte(s string) (byte, error) {
	var b byte
	for _, c := range s {
		b <<= 4
		switch {
		case c >= '0' && c <= '9':
			b |= byte(c - '0')
		case c >= 'a' && c <= 'f':
			b |= byte(c - 'a' + 10)
		case c >= 'A' && c <= 'F':
			b |= byte(c - 'A' + 10)
		default:
			return 0, fmt.Errorf("invalid hex character: %c", c)
		}
	}
	return b, nil
}

func (ipm *ipMaskValue) String() string {
	mask := net.IPMask(*ipm)
	if len(mask) != 4 {
		return "<nil>"
	}
	return fmt.Sprintf("%d.%d.%d.%d", mask[0], mask[1], mask[2], mask[3])
}
func (ipm *ipMaskValue) Type() string { return "ipMask" }

func (f *FlagSet) IPMaskVarP(p *net.IPMask, name, shorthand string, value net.IPMask, usage string) {
	f.VarPF(newIPMaskValue(value, p), name, shorthand, usage)
}
func (f *FlagSet) IPMaskVar(p *net.IPMask, name string, value net.IPMask, usage string) {
	f.IPMaskVarP(p, name, "", value, usage)
}
func (f *FlagSet) IPMaskP(name, shorthand string, value net.IPMask, usage string) *net.IPMask {
	p := new(net.IPMask)
	f.IPMaskVarP(p, name, shorthand, value, usage)
	return p
}
func (f *FlagSet) IPMask(name string, value net.IPMask, usage string) *net.IPMask {
	return f.IPMaskP(name, "", value, usage)
}
func (f *FlagSet) GetIPMask(name string) (net.IPMask, error) {
	flag := f.Lookup(name)
	if flag == nil {
		return nil, &ErrUnknownFlag{FlagName: name}
	}
	v, ok := flag.Value.(*ipMaskValue)
	if !ok {
		return nil, fmt.Errorf("trying to get ipMask value of flag of type %s", flag.Value.Type())
	}
	return net.IPMask(*v), nil
}
