package argov

import "strconv"

// Int64 is a wrapper for defaultParser.Int64().
func Int64(names []string, defaultValue int64, opts ...Option) *int64 {
	return defaultParser.Int64(names, defaultValue, opts...)
}

// Int64 registers a new int64-flag.
func (p *Parser) Int64(names []string, defaultValue int64, opts ...Option) *int64 {
	return Custom[int64](p, names, defaultValue, func(s string) (int64, error) {
		return strconv.ParseInt(s, 0, 64)
	}, opts...)
}
