package argov

type Option func(*flag)

type ParseOption func(*parsingOptions)

func Required() Option {
	return func(f *flag) {
		f.required = true
	}
}

func SplitOn(splitRunes ...rune) Option {
	return func(f *flag) {
		if f.val.IsSliceValue() {
			f.splitRunes = splitRunes
		} else {
			var name string
			if len(f.names) > 0 {
				name = f.names[0]
			}
			f.err = &InvalidOptionError{Flag: name, ErrMsg: "split runes provided on non-slice value"}
		}
	}
}

func NoMixing() ParseOption {
	return func(po *parsingOptions) {
		po.noMixing = true
	}
}
