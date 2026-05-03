package argov

import (
	"strconv"
)

type boolValue struct {
	val *bool
}

func (v *boolValue) IsSliceValue() bool {
	return false
}

func (v *boolValue) Set(s string) error {
	b, err := strconv.ParseBool(s)
	if err != nil {
		return err
	}
	*v.val = b
	return nil
}

func (v *boolValue) String() string {
	return strconv.FormatBool(*v.val)
}

func Bool(names []string, description string, defaultValue bool, opts ...Option) *bool {
	return parser.Bool(names, description, defaultValue, opts...)
}

func (p *Parser) Bool(names []string, description string, defaultValue bool, opts ...Option) *bool {
	val := new(bool)
	*val = defaultValue

	f := &flag{
		names:       names,
		description: description,
		val:         &boolValue{val},
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
