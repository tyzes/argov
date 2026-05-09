package argov

// String is a wrapper for defaultParser.String().
func String(names []string, defaultValue string, opts ...Option) *string {
	return defaultParser.String(names, defaultValue, opts...)
}

// String registers a new string-flag.
func (p *Parser) String(names []string, defaultValue string, opts ...Option) *string {
	return Custom[string](p, names, defaultValue, func(s string) (string, error) {
		return s, nil
	}, opts...)
}

// StringSlice is a wrapper for defaultParser.StringSlice().
func StringSlice(names []string, opts ...Option) *[]string {
	return defaultParser.StringSlice(names, opts...)
}

// StringSlice registers a new string-flag accepting multiple values.
func (p *Parser) StringSlice(names []string, opts ...Option) *[]string {
	return CustomSlice(p, names, func(s string) (string, error) {
		return s, nil
	}, opts...)
}
