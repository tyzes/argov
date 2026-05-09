package argov

import "fmt"

type genericValue[T any] struct {
	val   *T
	parse func(string) (T, error)
}

func (g *genericValue[T]) isSliceValue() bool {
	return false
}

func (g *genericValue[T]) set(s string) error {
	v, err := g.parse(s)
	if err != nil {
		return err
	}
	*g.val = v
	return nil
}

func (g *genericValue[T]) string() string {
	return fmt.Sprintf("%v", *g.val)
}

// Custom registers a new flag with a custom type and a custom parsing function.
func Custom[T any](p *Parser, names []string, defaultValue T, parse func(string) (T, error), opts ...Option) *T {
	val := new(T)
	*val = defaultValue

	f := &flag{
		names: names,
		val:   &genericValue[T]{val, parse},
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

func (g *genericSlice[T]) isSliceValue() bool {
	return true
}

func (g *genericSlice[T]) set(s string) error {
	v, err := g.parse(s)
	if err != nil {
		return err
	}
	*g.val = append(*g.val, v)
	return nil
}

func (g *genericSlice[T]) string() string {
	return fmt.Sprintf("%v", *g.val)
}

// CustomSlice registers a new slice-flag with a custom type and a custom parsing function.
func CustomSlice[T any](p *Parser, names []string, parse func(string) (T, error), opts ...Option) *[]T {
	val := new([]T)

	f := &flag{
		names: names,
		val:   &genericSlice[T]{val, parse},
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
