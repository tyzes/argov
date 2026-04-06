package argov

import "strconv"

func Float32(names []string, description string, defaultValue float32, opts ...Option) *float32 {
	return parser.Float32(names, description, defaultValue, opts...)
}

func (p *Parser) Float32(names []string, description string, defaultValue float32, opts ...Option) *float32 {
	return Custom[float32](p, names, description, defaultValue, func(s string) (float32, error) {
		v, err := strconv.ParseFloat(s, 32)
		return float32(v), err
	}, opts...)
}
