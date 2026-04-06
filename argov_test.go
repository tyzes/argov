package argov_test

import (
	"fmt"
	"slices"
	"strconv"
	"testing"

	"github.com/tyzes/argov"
)

var generalTests = []struct {
	flags       []string
	positionals []string
	s           string
	b, B        bool
	i           int
	I           int64
	f           float32
	F           float64
	S           []string
	C           []int64
	c           int64
}{
	{flags: []string{"-b"}, b: true},
	{flags: []string{"-bB"}, b: true, B: true},
	{flags: []string{"--bool"}, b: true},
	{flags: []string{"-b=true"}, b: true},
	{flags: []string{"-b=f"}},

	{flags: []string{"-s", "hello"}, s: "hello"},
	{flags: []string{"--string", "hello"}, s: "hello"},
	{flags: []string{"-s=hello"}, s: "hello"},
	{flags: []string{"--string=hello"}, s: "hello"},
	{flags: []string{"--string=hello world"}, s: "hello world"},

	{flags: []string{"-i", "42"}, i: 42},
	{flags: []string{"--int", "42"}, i: 42},
	{flags: []string{"-i=42"}, i: 42},
	{flags: []string{"-i=-42"}, i: -42},
	{flags: []string{"-i", "0"}},

	{flags: []string{"-I", "999999999999999999"}, I: 999999999999999999},
	{flags: []string{"--int64", "999999999999999999"}, I: 999999999999999999},
	{flags: []string{"--int64=999"}, I: 999},
	{flags: []string{"-I=-1"}, I: -1},

	{flags: []string{"-f", "3.14"}, f: 3.14},
	{flags: []string{"-f=3.14159"}, f: 3.14159},

	{flags: []string{"-F", "3.141592653589793"}, F: 3.141592653589793},

	{flags: []string{"-b", "--int", "1234"}, b: true, i: 1234},
	{flags: []string{"-s", "foo", "-i", "7", "-b"}, s: "foo", b: true, i: 7},

	{flags: []string{"-b", "file.txt"}, positionals: []string{"file.txt"}, b: true},
	{flags: []string{"-i", "1", "a.txt", "b.txt"}, positionals: []string{"a.txt", "b.txt"}, i: 1},

	{flags: []string{"--", "-b"}, positionals: []string{"-b"}},
	{flags: []string{"-i", "1", "--", "--string", "foo"}, positionals: []string{"--string", "foo"}, i: 1},
	{flags: []string{"-b", "--"}, positionals: []string{}, b: true},

	{flags: []string{}},
	{flags: []string{"file.txt"}, positionals: []string{"file.txt"}},

	{flags: []string{"-S=foo", "-S", "bar", "123"}, positionals: []string{"123"}, S: []string{"foo", "bar"}},
	{flags: []string{"--Slice=a", "-S", "b", "-S=c", "-S", "d", "--", "-S", "d"}, positionals: []string{"-S", "d"}, S: []string{"a", "b", "c", "d"}},
	{flags: []string{"-S", "foo", "-b", "--Slice", "bar", "--int", "123", "-S=456", "7890", "abc"}, positionals: []string{"7890", "abc"}, S: []string{"foo", "bar", "456"}, b: true, i: 123},

	{flags: []string{"-C=123", "--Slice", "a", "-b", "--Custom", "456", "-C", "789", "--Slice=b", "--Custom=0", "abc"}, positionals: []string{"abc"}, C: []int64{123, 456, 789, 0}, S: []string{"a", "b"}, b: true},
	{flags: []string{"-c", "123", "456"}, positionals: []string{"456"}, c: 123},
}

func TestParse(t *testing.T) {
	for _, tt := range generalTests {
		p := argov.NewParser()
		s := p.String([]string{"s", "string"}, "", "")
		b := p.Bool([]string{"b", "bool"}, "", false)
		B := p.Bool([]string{"B", "Bool"}, "", false)
		i := p.Int([]string{"i", "int"}, "", 0)
		I := p.Int64([]string{"I", "int64"}, "", 0)
		f := p.Float32([]string{"f", "float32"}, "", 0)
		F := p.Float64([]string{"F", "float64"}, "", 0)
		S := p.StringSlice([]string{"S", "Slice"}, "")
		argov.Custom(p, []string{"c", "custom"}, "", 0, func(s string) (int64, error) {
			return strconv.ParseInt(s, 0, 64)
		})
		C := argov.Slice(p, []string{"C", "Custom"}, "", func(s string) (int64, error) {
			return strconv.ParseInt(s, 0, 64)
		})

		res, err := p.Parse(tt.flags)
		if err != nil {
			t.Errorf("%v: unexpected error: %s", tt.flags, err.Error())
			continue
		}
		if *s != tt.s {
			t.Errorf("%v: string: expected %q, got %q", tt.flags, tt.s, *s)
		}
		if *b != tt.b {
			t.Errorf("%v: bool 1: expected %v, got %v", tt.flags, tt.b, *b)
		}
		if *B != tt.B {
			t.Errorf("%v: bool 2: expected %v, got %v", tt.flags, tt.B, *B)
		}
		if *i != tt.i {
			t.Errorf("%v: int: expected %v, got %v", tt.flags, tt.i, *i)
		}
		if *I != tt.I {
			t.Errorf("%v: int64: expected %v, got %v", tt.flags, tt.I, *I)
		}
		if *f != tt.f {
			t.Errorf("%v: float32: expected %v, got %v", tt.flags, tt.f, *f)
		}
		if *F != tt.F {
			t.Errorf("%v: float64: expected %v, got %v", tt.flags, tt.F, *F)
		}

		if !slices.Equal(*S, tt.S) {
			t.Errorf("%v: string: expected %v, got %v", tt.flags, tt.S, *S)
		}
		if !slices.Equal(*C, tt.C) {
			t.Errorf("%v: bool: expected %v, got %v", tt.flags, tt.C, *C)
		}
		if !slices.Equal(res, tt.positionals) {
			t.Errorf("%v: positionals: expected %v, got %v", tt.flags, tt.positionals, res)
		}
	}
}

var errorTests = []struct {
	flags   []string
	wantErr string
	wantMsg string
}{
	{[]string{"-"}, "*argov.FlagSyntaxError", "invalid syntax: missing flag: '-'"},
	{[]string{"-bs"}, "*argov.FlagSyntaxError", "invalid syntax: non-boolean flag in flag group: 's'"},
	{[]string{"-z"}, "*argov.FlagUnknownError", "unknown flag 'z'"},
	{[]string{"--unknown"}, "*argov.FlagUnknownError", "unknown flag 'unknown'"},
	{[]string{"-bz"}, "*argov.FlagUnknownError", "unknown flag 'z'"},
	{[]string{"-i", "abc"}, "*argov.InvalidValueError", "invalid value for flag 'i': 'abc'"},
	{[]string{"-i", "3.14"}, "*argov.InvalidValueError", "invalid value for flag 'i': '3.14'"},
	{[]string{"-b=maybe"}, "*argov.InvalidValueError", "invalid value for flag 'b': 'maybe'"},
	{[]string{"-s"}, "*argov.MissingValueError", "missing value for flag 's'"},
	{[]string{"-i"}, "*argov.MissingValueError", "missing value for flag 'i'"},
	{[]string{"--string"}, "*argov.MissingValueError", "missing value for flag 'string'"},
	{[]string{"--string="}, "*argov.MissingValueError", "missing value for flag 'string'"},
	{[]string{"--=value"}, "*argov.FlagUnknownError", "unknown flag ''"},
	{[]string{"-S="}, "*argov.MissingValueError", "missing value for flag 'S'"},
	{[]string{"-C=foo"}, "*argov.InvalidValueError", "invalid value for flag 'C': 'foo'"},
	{[]string{"-c=foo"}, "*argov.InvalidValueError", "invalid value for flag 'c': 'foo'"},
}

func TestParseErrors(t *testing.T) {
	for _, tt := range errorTests {
		p := argov.NewParser()
		p.Bool([]string{"b", "bool"}, "", false)
		p.String([]string{"s", "string"}, "", "")
		p.Int([]string{"i", "int"}, "", -2)
		p.Int64([]string{"I", "int64"}, "", -2)
		p.Float32([]string{"f", "float32"}, "", -2)
		p.Float64([]string{"F", "float64"}, "", -2)
		p.StringSlice([]string{"S", "Slice"}, "")
		argov.Custom(p, []string{"c", "custom"}, "", 0, func(s string) (int64, error) {
			return strconv.ParseInt(s, 0, 64)
		})
		argov.Slice(p, []string{"C", "Custom"}, "", func(s string) (int64, error) {
			return strconv.ParseInt(s, 0, 64)
		})

		_, err := p.Parse(tt.flags)
		if err == nil {
			t.Errorf("%v: expected error, got nil", tt.flags)
		} else if fmt.Sprintf("%T", err) != tt.wantErr {
			t.Errorf("%v: expected error type %v, got %T", tt.flags, tt.wantErr, err)
		} else if err.Error() != tt.wantMsg {
			t.Errorf("%v: expected message \"%v\", got \"%v\"", tt.flags, tt.wantMsg, err.Error())
		}
	}
}

var requiredTests = []struct {
	flags       []string
	positionals []string
	hasErr      bool
	wantMsg     string
	b, B        bool
	s, S        string
}{
	{flags: []string{"-b", "-B", "-s", "foo", "-S", "bar", "test", "TEST"}, positionals: []string{"test", "TEST"}, b: true, B: true, s: "foo", S: "bar"},
	{flags: []string{"-B", "-S", "foo", "--", "bar"}, positionals: []string{"bar"}, B: true, S: "foo"},
	{flags: []string{"-B", "-s", "foo", "-S", "bar"}, B: true, s: "foo", S: "bar"},
	{flags: []string{"-b", "-s", "foo", "-S", "bar"}, hasErr: true, wantMsg: "missing required flag 'B'"},
	{flags: []string{"-b", "-B", "-s", "foo"}, hasErr: true, wantMsg: "missing required flag 'S'"},
	{flags: []string{}, hasErr: true, wantMsg: "missing required flag 'B'"},
	{flags: []string{"-b", "-s", "foo"}, hasErr: true, wantMsg: "missing required flag 'B'"},
}

func TestRequired(t *testing.T) {
	for _, tt := range requiredTests {
		p := argov.NewParser()
		b := p.Bool([]string{"b", "bool"}, "", false)
		B := p.Bool([]string{"B", "Bool"}, "", false, argov.Required())
		s := p.String([]string{"s", "string"}, "", "")
		S := p.String([]string{"S", "String"}, "", "", argov.Required())

		res, err := p.Parse(tt.flags)
		if err != nil && !tt.hasErr {
			t.Errorf("%v: unexpected error: %v", tt.flags, err)
		}
		if err == nil && tt.hasErr {
			t.Errorf("%v: expected error, got nil", tt.flags)
			continue
		}
		if err != nil && err.Error() != tt.wantMsg {
			t.Errorf("%v: expected message \"%v\", got \"%v\"", tt.flags, tt.wantMsg, err.Error())
		}

		if err == nil {
			if *b != tt.b {
				t.Errorf("%v: bool: expected %v, got %v", tt.flags, tt.b, *b)
			}
			if *B != tt.B {
				t.Errorf("%v: required bool: expected %v, got %v", tt.flags, tt.B, *B)
			}
			if *s != tt.s {
				t.Errorf("%v: string: expected \"%v\", got \"%v\"", tt.flags, tt.s, *s)
			}
			if *S != tt.S {
				t.Errorf("%v: required string: expected \"%v\", got \"%v\"", tt.flags, tt.S, *S)
			}

			if !slices.Equal(res, tt.positionals) {
				t.Errorf("%v: positionals: expected %v, got %v", tt.flags, tt.positionals, res)
			}
		}
	}
}
