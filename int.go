package argov

import "strconv"

// Int is a wrapper for defaultParser.Int().
func Int(names []string, defaultValue int, opts ...Option) *int {
	return defaultParser.Int(names, defaultValue, opts...)
}

// Int registers a new int-flag.
func (p *Parser) Int(names []string, defaultValue int, opts ...Option) *int {
	return Custom[int](p, names, defaultValue, func(s string) (int, error) {
		v, err := strconv.ParseInt(s, 0, strconv.IntSize)
		return int(v), err
	}, opts...)
}
