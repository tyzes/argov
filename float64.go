package argov

import "strconv"

func Float64(names []string, defaultValue float64, opts ...Option) *float64 {
	return parser.Float64(names, defaultValue, opts...)
}

func (p *Parser) Float64(names []string, defaultValue float64, opts ...Option) *float64 {
	return Custom[float64](p, names, defaultValue, func(s string) (float64, error) {
		return strconv.ParseFloat(s, 64)
	}, opts...)
}
