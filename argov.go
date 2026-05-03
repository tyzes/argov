package argov

type flag struct {
	names       []string
	description string
	val         Value
	required    bool
	splitRunes  []rune
	err         error
}

type Value interface {
	Set(string) error
	String() string
	IsSliceValue() bool
}

type Parser struct {
	flags  []*flag
	lookup map[string]*flag
	isSet  map[string]struct{}
}

type parsingOptions struct {
	noMixing bool
}

func NewParser() *Parser {
	return &Parser{lookup: make(map[string]*flag), isSet: make(map[string]struct{})}
}

var parser = NewParser()

func Parse(args []string, opts ...ParseOption) ([]string, error) {
	return parser.Parse(args, opts...)
}

func IsSet(name string) bool {
	return parser.IsSet(name)
}

func (p *Parser) IsSet(name string) bool {
	f, ok := p.lookup[name]
	if !ok {
		return false
	}
	_, set := p.isSet[f.names[0]]
	return set
}
