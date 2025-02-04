package pflag

import "strconv"

// -- uint16 value
type uint16Value uint16

func newUint16Value(val uint16, p *uint16) *uint16Value {
	*p = val
	return (*uint16Value)(p)
}

func (i *uint16Value) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, 16)
	*i = uint16Value(v)
	return err
}

func (i *uint16Value) Type() string {
	return "uint16"
}

func (i *uint16Value) String() string { return strconv.FormatUint(uint64(*i), 10) }

func uint16Conv(sval string) (interface{}, error) {
	v, err := strconv.ParseUint(sval, 0, 16)
	if err != nil {
		return 0, err
	}
	return uint16(v), nil
}

// GetUint16 return the uint16 value of a flag with the given name
func (f *FlagSet) GetUint16(name string) (uint16, error) {
	val, err := f.getFlagType(name, "uint16", uint16Conv)
	if err != nil {
		return 0, err
	}
	return val.(uint16), nil
}

// Uint16Var defines a uint flag with specified name, default value, and usage string.
// The argument p points to a uint variable in which to store the value of the flag.
func (f *FlagSet) Uint16Var(p *uint16, name string, value uint16, usage string) {
	f.Uint16VarP(p, name, "", value, usage)
}

// Uint16VarN is like Uint16VarP, but adds the name as shorthand (non-posix).
func (f *FlagSet) Uint16VarN(p *uint16, name, shorthand string, value uint16, usage string) {
	f.VarN(newUint16Value(value, p), name, shorthand, usage)
}

// Uint16VarP is like Uint16Var, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Uint16VarP(p *uint16, name, shorthand string, value uint16, usage string) {
	f.VarP(newUint16Value(value, p), name, shorthand, usage)
}

// Uint16VarS is like Uint16Var, but accepts a shorthand letter that can be used after a single dash, alone.
func (f *FlagSet) Uint16VarS(p *uint16, name, shorthand string, value uint16, usage string) {
	f.VarS(newUint16Value(value, p), name, shorthand, usage)
}

// Uint16Var defines a uint flag with specified name, default value, and usage string.
// The argument p points to a uint  variable in which to store the value of the flag.
func Uint16Var(p *uint16, name string, value uint16, usage string) {
	CommandLine.Uint16Var(p, name, value, usage)
}

// Uint16VarN is like Uint16VarP, but adds the name as shorthand (non-posix).
func Uint16VarN(p *uint16, name, shorthand string, value uint16, usage string) {
	CommandLine.Uint16VarN(p, name, shorthand, value, usage)
}

// Uint16VarP is like Uint16Var, but accepts a shorthand letter that can be used after a single dash.
func Uint16VarP(p *uint16, name, shorthand string, value uint16, usage string) {
	CommandLine.Uint16VarP(p, name, shorthand, value, usage)
}

// Uint16VarS is like Uint16Var, but accepts a shorthand letter that can be used after a single dash, alone.
func Uint16VarS(p *uint16, name, shorthand string, value uint16, usage string) {
	CommandLine.Uint16VarS(p, name, shorthand, value, usage)
}

// Uint16 defines a uint flag with specified name, default value, and usage string.
// The return value is the address of a uint  variable that stores the value of the flag.
func (f *FlagSet) Uint16(name string, value uint16, usage string) *uint16 {
	return f.Uint16P(name, "", value, usage)
}

// Uint16N is like Uint16P, but adds the name as shorthand (non-posix).
func (f *FlagSet) Uint16N(name, shorthand string, value uint16, usage string) *uint16 {
	p := new(uint16)
	f.Uint16VarN(p, name, shorthand, value, usage)
	return p
}

// Uint16P is like Uint16, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Uint16P(name, shorthand string, value uint16, usage string) *uint16 {
	p := new(uint16)
	f.Uint16VarP(p, name, shorthand, value, usage)
	return p
}

// Uint16S is like Uint16, but accepts a shorthand letter that can be used after a single dash, alone.
func (f *FlagSet) Uint16S(name, shorthand string, value uint16, usage string) *uint16 {
	p := new(uint16)
	f.Uint16VarS(p, name, shorthand, value, usage)
	return p
}

// Uint16 defines a uint flag with specified name, default value, and usage string.
// The return value is the address of a uint  variable that stores the value of the flag.
func Uint16(name string, value uint16, usage string) *uint16 {
	return CommandLine.Uint16(name, value, usage)
}

// Uint16N is like Uint16P, but adds the name as shorthand (non-posix).
func Uint16N(name, shorthand string, value uint16, usage string) *uint16 {
	return CommandLine.Uint16N(name, shorthand, value, usage)
}

// Uint16P is like Uint16, but accepts a shorthand letter that can be used after a single dash.
func Uint16P(name, shorthand string, value uint16, usage string) *uint16 {
	return CommandLine.Uint16P(name, shorthand, value, usage)
}

// Uint16S is like Uint16, but accepts a shorthand letter that can be used after a single dash, alone.
func Uint16S(name, shorthand string, value uint16, usage string) *uint16 {
	return CommandLine.Uint16S(name, shorthand, value, usage)
}
