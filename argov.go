package argov

type flag struct {
	names       []string
	description string
	placeholder string
	val         value
	required    bool
	splitRunes  []rune
	err         error
}

type value interface {
	set(string) error
	isSliceValue() bool
	string() string
}

// Parser is the central struct and holds all flags registered to it.
// New parsers should only be created using the NewParser() function.
type Parser struct {
	flags  []*flag
	lookup map[string]*flag
	isSet  map[string]struct{}
}

type parsingOptions struct {
	noMixing bool
}

// NewParser returns a new parser.
func NewParser() *Parser {
	return &Parser{lookup: make(map[string]*flag), isSet: make(map[string]struct{})}
}

var defaultParser = NewParser()

// Parse is a wrapper for defaultParser.Parse().
func Parse(args []string, opts ...ParseOption) ([]string, error) {
	return defaultParser.Parse(args, opts...)
}

// IsSet is a wrapper for defaultParser.IsSet().
func IsSet(name string) bool {
	return defaultParser.IsSet(name)
}

// IsSet returns if the given flag has been passed into Parse().
// If a flag has multiple names, any name can be passed to IsSet.
func (p *Parser) IsSet(name string) bool {
	f, ok := p.lookup[name]
	if !ok {
		return false
	}
	_, set := p.isSet[f.names[0]]
	return set
}
