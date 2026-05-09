package argov

// Option configures a flag and is passed upon flag registration.
type Option func(*flag)

// ParseOption modifies the behaviour of the parser and is passed along with Parse().
type ParseOption func(*parsingOptions)

// Description adds a description text to a flag.
// This is only used for the help menu generation.
func Description(description string) Option {
	return func(f *flag) {
		f.description = description
	}
}

// Placeholder adds a placeholder text to a flag (e.g. FILE for -f FILE).
// This is only used for the help menu generation.
func Placeholder(placeholder string) Option {
	return func(f *flag) {
		f.placeholder = placeholder
	}
}

// Required makes the parser throw an error if a flag registered with it is not passed.
func Required() Option {
	return func(f *flag) {
		f.required = true
	}
}

// SplitOn sets runes on which a slice flag's value is split into separate entries.
// Slice-flags can also be passed multiples times.
func SplitOn(splitRunes ...rune) Option {
	return func(f *flag) {
		if f.val.isSliceValue() {
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

// NoMixing tells the parser to not allow mixing of flags and positional arguments.
// With this option, the parser will stop parsing if it encounters anything that is not a flag or a value and will return the remaining as positional arguments.
func NoMixing() ParseOption {
	return func(po *parsingOptions) {
		po.noMixing = true
	}
}
