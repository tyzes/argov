package argov

import "strconv"

// Float64 is a wrapper for defaultParser.Float64().
func Float64(names []string, defaultValue float64, opts ...Option) *float64 {
	return defaultParser.Float64(names, defaultValue, opts...)
}

// Float64 registers a new float64-flag.
func (p *Parser) Float64(names []string, defaultValue float64, opts ...Option) *float64 {
	return Custom[float64](p, names, defaultValue, func(s string) (float64, error) {
		return strconv.ParseFloat(s, 64)
	}, opts...)
}
