package argov

func String(names []string, description string, defaultValue string, opts ...Option) *string {
	return parser.String(names, description, defaultValue, opts...)
}

func (p *Parser) String(names []string, description string, defaultValue string, opts ...Option) *string {
	return Custom[string](p, names, description, defaultValue, func(s string) (string, error) {
		return s, nil
	}, opts...)
}

func StringSlice(names []string, description string, opts ...Option) *[]string {
	return parser.StringSlice(names, description, opts...)
}

func (p *Parser) StringSlice(names []string, description string, opts ...Option) *[]string {
	return Slice(p, names, description, func(s string) (string, error) {
		return s, nil
	}, opts...)
}
