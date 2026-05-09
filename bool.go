package argov

import (
	"strconv"
)

type boolValue struct {
	val *bool
}

func (v *boolValue) isSliceValue() bool {
	return false
}

func (v *boolValue) set(s string) error {
	b, err := strconv.ParseBool(s)
	if err != nil {
		return err
	}
	*v.val = b
	return nil
}

func (v *boolValue) string() string {
	return strconv.FormatBool(*v.val)
}

// Bool is a wrapper for defaultParser.Bool().
func Bool(names []string, defaultValue bool, opts ...Option) *bool {
	return defaultParser.Bool(names, defaultValue, opts...)
}

// Bool registers a new boolean-flag.
func (p *Parser) Bool(names []string, defaultValue bool, opts ...Option) *bool {
	val := new(bool)
	*val = defaultValue

	f := &flag{
		names: names,
		val:   &boolValue{val},
	}

	for _, opt := range opts {
		opt(f)
	}

	p.flags = append(p.flags, f)
	for _, name := range names {
		p.lookup[name] = f
	}

	return val
}
