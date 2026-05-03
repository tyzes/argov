package argov

import (
	"fmt"
	"strings"
)

func (p *Parser) Parse(args []string, opts ...ParseOption) ([]string, error) {
	for _, f := range p.flags {
		if f.err != nil {
			return nil, f.err
		}
	}

	po := parsingOptions{}
	for _, opt := range opts {
		opt(&po)
	}

	p.isSet = make(map[string]struct{})

	var positionals = []string{}
	i := 0
	for ; i < len(args); i++ {
		if strings.HasPrefix(args[i], "--") {
			if len(args[i]) == 2 {
				if missing := p.checkRequired(); missing != "" {
					return nil, &MissingRequiredFlagError{missing}
				}
				if len(args) > i {
					positionals = append(positionals, args[i+1:]...)
				}
				return positionals, nil
			}

			consumed, err := p.parseDoubleDash(args, i)
			if err != nil {
				return nil, err
			}
			if consumed {
				i++
			}
		} else if strings.HasPrefix(args[i], "-") {
			consumed, err := p.parseSingleDash(args, i)
			if err != nil {
				return nil, err
			}
			if consumed {
				i++
			}
		} else {
			if po.noMixing {
				if missing := p.checkRequired(); missing != "" {
					return nil, &MissingRequiredFlagError{missing}
				}
				if len(args) > i {
					return args[i:], nil
				}
				return nil, nil
			}
			positionals = append(positionals, args[i])
		}
	}
	if missing := p.checkRequired(); missing != "" {
		return nil, &MissingRequiredFlagError{missing}
	}
	return positionals, nil
}

func (p *Parser) parseSingleDash(args []string, i int) (bool, error) {
	switch len(args[i]) {
	case 1:
		return false, &FlagSyntaxError{ErrMsg: fmt.Sprintf("missing flag: '%s'", args[i])}
	case 2:
		f, ok := p.lookup[args[i][1:]]
		if !ok {
			return false, &FlagUnknownError{Flag: args[i][1:]}
		}

		isBool := p.parseBool(f)
		if !isBool {
			err := p.parseValue(f, string(args[i][1]), args, i)
			if err != nil {
				return false, err
			}
			return true, nil
		}
	default:
		isEquals, err := p.parseEquals(args[i][1:])
		if err != nil {
			return false, err
		}
		if !isEquals {
			for _, arg := range args[i][1:] {
				f, ok := p.lookup[string(arg)]
				if !ok {
					return false, &FlagUnknownError{Flag: string(arg)}
				}

				isBool := p.parseBool(f)
				if !isBool {
					return false, &FlagSyntaxError{ErrMsg: fmt.Sprintf("non-boolean flag in flag group: '%s'", string(arg))}
				}
			}
			return false, nil
		}
	}
	return false, nil
}

func (p *Parser) parseDoubleDash(args []string, i int) (bool, error) {
	isEquals, err := p.parseEquals(args[i][2:])
	if err != nil {
		return false, err
	}
	if !isEquals {
		f, ok := p.lookup[string(args[i][2:])]
		if !ok {
			return false, &FlagUnknownError{Flag: string(args[i][2:])}
		}

		isBool := p.parseBool(f)
		if !isBool {
			err := p.parseValue(f, string(args[i][2:]), args, i)
			if err != nil {
				return false, err
			}
			return true, nil
		}
	}
	return false, nil
}

func (p *Parser) parseBool(f *flag) bool {
	if boolValue, ok := f.val.(*boolValue); ok {
		*boolValue.val = true
		p.isSet[f.names[0]] = struct{}{}
		return true
	}
	return false
}

func (p *Parser) parseEquals(flag string) (bool, error) {
	equals := strings.IndexRune(flag, '=')
	if equals < 0 {
		return false, nil
	}

	f, ok := p.lookup[flag[:equals]]
	if !ok {
		return true, &FlagUnknownError{Flag: flag[:equals]}
	}
	if len(flag)-1 <= equals {
		return false, &MissingValueError{Flag: flag[:equals]}
	}

	err := setValue(f, flag[equals+1:])
	if err != nil {
		return true, &InvalidValueError{Flag: flag[:equals], Value: flag[equals+1:], Err: err}
	}

	p.isSet[f.names[0]] = struct{}{}
	return true, nil
}

func (p *Parser) parseValue(f *flag, flag string, args []string, i int) error {
	if len(args) <= i+1 {
		return &MissingValueError{Flag: flag}
	}

	err := setValue(f, args[i+1])
	if err != nil {
		return &InvalidValueError{Flag: flag, Value: args[i+1], Err: err}
	}
	p.isSet[f.names[0]] = struct{}{}
	return nil
}
