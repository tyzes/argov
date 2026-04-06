package argov

import (
	"fmt"
	"strings"
)

func (p *Parser) Parse(args []string) ([]string, error) {
	i := 0
	for ; i < len(args); i++ {
		if strings.HasPrefix(args[i], "--") {
			if len(args[i]) == 2 {
				if len(args) > i {
					return args[i+1:], nil
				}
				return []string{}, nil
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
			missing := p.checkRequired()
			if missing != "" {
				return nil, &MissingRequiredFlagError{missing}
			}
			return args[i:], nil
		}
	}
	missing := p.checkRequired()
	if missing != "" {
		return nil, &MissingRequiredFlagError{missing}
	}
	return []string{}, nil
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
			err := p.parseValue(string(args[i][1]), args, i)
			if err != nil {
				return false, err
			}
			return true, nil
		}
	default:
		isEquals, err := p.parseEquals(args[i][1:], args, i)
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
	isEquals, err := p.parseEquals(args[i][2:], args, i)
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
			err := p.parseValue(string(args[i][2:]), args, i)
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

func (p *Parser) parseEquals(flag string, args []string, i int) (bool, error) {
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

	err := f.val.Set(flag[equals+1:])
	if err != nil {
		return true, &InvalidValueError{Flag: flag[:equals], Value: flag[equals+1:], Err: err}
	}

	p.isSet[f.names[0]] = struct{}{}
	return true, nil
}

func (p *Parser) parseValue(flag string, args []string, i int) error {
	if len(args) <= i+1 {
		return &MissingValueError{Flag: flag}
	}

	f, ok := p.lookup[flag]
	if !ok {
		return &FlagUnknownError{Flag: flag}
	}

	err := f.val.Set(args[i+1])
	if err != nil {
		return &InvalidValueError{Flag: flag, Value: args[i+1], Err: err}
	}
	p.isSet[f.names[0]] = struct{}{}
	return nil
}
