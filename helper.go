package argov

import (
	"slices"
	"strings"
)

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

func splitOnRunes(str string, runes []rune) []string {
	return strings.FieldsFunc(str, func(currentRune rune) bool {
		return slices.Contains(runes, currentRune)
	})
}

func setValue(f *flag, str string) error {
	if f.val.IsSliceValue() && len(f.splitRunes) > 0 {
		for _, s := range splitOnRunes(str, f.splitRunes) {
			err := f.val.Set(s)
			if err != nil {
				return err
			}
		}
	} else {
		return f.val.Set(str)
	}
	return nil
}
