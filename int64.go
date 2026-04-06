package argov

import "strconv"

func Int64(names []string, description string, defaultValue int64, opts ...Option) *int64 {
	return parser.Int64(names, description, defaultValue, opts...)
}

func (p *Parser) Int64(names []string, description string, defaultValue int64, opts ...Option) *int64 {
	return Custom[int64](p, names, description, defaultValue, func(s string) (int64, error) {
		return strconv.ParseInt(s, 0, 64)
	}, opts...)
}
