package argov

import "strconv"

func Int(names []string, defaultValue int, opts ...Option) *int {
	return parser.Int(names, defaultValue, opts...)
}

func (p *Parser) Int(names []string, defaultValue int, opts ...Option) *int {
	return Custom[int](p, names, defaultValue, func(s string) (int, error) {
		v, err := strconv.ParseInt(s, 0, strconv.IntSize)
		return int(v), err
	}, opts...)
}
