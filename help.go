package argov

import (
	"bytes"
)

type FlagInformation struct {
	Names       []string
	Description string
	Placeholder string
	Required    bool
	IsSliceVal  bool
	SplitRunes  []rune
}

func GetFlags() []FlagInformation {
	return parser.GetFlags()
}

func (p *Parser) GetFlags() []FlagInformation {
	var flags []FlagInformation
	for _, f := range p.flags {
		flags = append(flags, FlagInformation{f.names, f.description, f.placeholder, f.required, f.val.IsSliceValue(), f.splitRunes})
	}
	return flags
}

func GenerateHelp() string {
	return parser.GenerateHelp()
}

func (p *Parser) GenerateHelp() string {
	var buf bytes.Buffer

	maxShortsLen := 0
	maxLongsLen := 0

	for _, f := range p.flags {
		shortsLen := 0
		longsLen := 0
		for _, name := range f.names {
			if len(name) == 1 {
				if shortsLen > 0 {
					shortsLen += 2
				}
				shortsLen += 2
			} else {
				if longsLen > 0 {
					longsLen += 2
				}
				longsLen += 2 + len(name)
			}
		}
		if f.placeholder != "" {
			longsLen += len(f.placeholder) + 3
		}
		if shortsLen > maxShortsLen {
			maxShortsLen = shortsLen
		}
		if longsLen > maxLongsLen {
			maxLongsLen = longsLen
		}
	}

	for i, f := range p.flags {
		buf.WriteByte(' ')

		shortsLen := 0
		for _, name := range f.names {
			if len(name) == 1 {
				if shortsLen > 0 {
					buf.WriteString(", ")
					shortsLen += 2
				}
				buf.WriteByte('-')
				buf.WriteString(name)
				shortsLen += 2
			}
		}

		writeSpaces(&buf, maxShortsLen-shortsLen+2)

		longsLen := 0
		for _, name := range f.names {
			if len(name) > 1 {
				if longsLen > 0 {
					buf.WriteString(", ")
					longsLen += 2
				}
				buf.WriteString("--")
				buf.WriteString(name)
				longsLen += 2 + len(name)
			}
		}

		if f.placeholder != "" {
			buf.WriteString(" <")
			buf.WriteString(f.placeholder)
			buf.WriteRune('>')
			longsLen += len(f.placeholder) + 3
		}

		writeSpaces(&buf, maxLongsLen-longsLen+3)

		buf.WriteString(f.description)

		hintsLen := 0
		if f.required {
			buf.WriteString(" (Required")
			hintsLen++
		}
		if f.val.IsSliceValue() {
			if hintsLen > 0 {
				buf.WriteString(", ")
			} else {
				buf.WriteString(" (")
			}
			buf.WriteString("Accepts repeated calls")
			if len(f.splitRunes) > 0 {
				buf.WriteString(" or multiple values separated by ")
				buf.WriteString("'")
				for j, r := range f.splitRunes {
					buf.WriteRune(r)
					if j < len(f.splitRunes)-1 {
						buf.WriteString("' or '")
					}
				}
				buf.WriteRune('\'')
			}
			hintsLen++
		}
		if hintsLen > 0 {
			buf.WriteRune(')')
		}

		if i < len(p.flags)-1 {
			buf.WriteRune('\n')
		}
	}

	return buf.String()
}

func writeSpaces(buf *bytes.Buffer, n int) {
	for range n {
		buf.WriteByte(' ')
	}
}
