package argov

type Option func(*flag)

func Required() Option {
	return func(f *flag) {
		f.required = true
	}
}

func (p *Parser) checkRequired() string {
	for _, f := range p.flags {
		if f.required {
			if _, ok := p.isSet[f.names[0]]; !ok {
				return f.names[0]
			}
		}
	}
	return ""
}
