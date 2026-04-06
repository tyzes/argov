package argov

import "fmt"

type genericValue[T any] struct {
	val   *T
	parse func(string) (T, error)
}

func (g *genericValue[T]) Set(s string) error {
	v, err := g.parse(s)
	if err != nil {
		return err
	}
	*g.val = v
	return nil
}

func (g *genericValue[T]) String() string {
	return fmt.Sprintf("%v", *g.val)
}

func Custom[T any](p *Parser, names []string, description string, defaultValue T, parse func(string) (T, error), opts ...Option) *T {
	val := new(T)
	*val = defaultValue

	f := &flag{
		names:       names,
		description: description,
		val:         &genericValue[T]{val, parse},
	}

	for _, opt := range opts {
		opt(f)
	}

	p.flags = append(p.flags, f)
	for _, name := range names {
		p.lookup[name] = f
	}
	return val
}

type genericSlice[T any] struct {
	val   *[]T
	parse func(string) (T, error)
}

func (g *genericSlice[T]) Set(s string) error {
	v, err := g.parse(s)
	if err != nil {
		return err
	}
	*g.val = append(*g.val, v)
	return nil
}

func (g *genericSlice[T]) String() string {
	return fmt.Sprintf("%v", *g.val)
}

func Slice[T any](p *Parser, names []string, description string, parse func(string) (T, error), opts ...Option) *[]T {
	val := new([]T)

	f := &flag{
		names:       names,
		description: description,
		val:         &genericSlice[T]{val, parse},
	}

	for _, opt := range opts {
		opt(f)
	}

	p.flags = append(p.flags, f)
	for _, name := range names {
		p.lookup[name] = f
	}
	return val
}
