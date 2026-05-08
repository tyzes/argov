package argov

func String(names []string, defaultValue string, opts ...Option) *string {
	return parser.String(names, defaultValue, opts...)
}

func (p *Parser) String(names []string, defaultValue string, opts ...Option) *string {
	return Custom[string](p, names, defaultValue, func(s string) (string, error) {
		return s, nil
	}, opts...)
}

func StringSlice(names []string, opts ...Option) *[]string {
	return parser.StringSlice(names, opts...)
}

func (p *Parser) StringSlice(names []string, opts ...Option) *[]string {
	return CustomSlice(p, names, func(s string) (string, error) {
		return s, nil
	}, opts...)
}
