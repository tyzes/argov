package argov

import "strconv"

// Float32 is a wrapper for defaultParser.Float32().
func Float32(names []string, defaultValue float32, opts ...Option) *float32 {
	return defaultParser.Float32(names, defaultValue, opts...)
}

// Float32 registers a new float32-flag.
func (p *Parser) Float32(names []string, defaultValue float32, opts ...Option) *float32 {
	return Custom[float32](p, names, defaultValue, func(s string) (float32, error) {
		v, err := strconv.ParseFloat(s, 32)
		return float32(v), err
	}, opts...)
}
