package argov

import (
	"fmt"
)

type FlagUnknownError struct {
	Flag string
}

func (e *FlagUnknownError) Error() string {
	return fmt.Sprintf("unknown flag '%s'", e.Flag)
}

type FlagSyntaxError struct {
	ErrMsg string
}

func (e *FlagSyntaxError) Error() string {
	return fmt.Sprintf("invalid syntax: %s", e.ErrMsg)
}

type MissingValueError struct {
	Flag string
}

func (e *MissingValueError) Error() string {
	return fmt.Sprintf("missing value for flag '%s'", e.Flag)
}

type MissingRequiredFlagError struct {
	Flag string
}

func (e *MissingRequiredFlagError) Error() string {
	return fmt.Sprintf("missing required flag '%s'", e.Flag)
}

type InvalidValueError struct {
	Flag  string
	Value string
	Err   error
}

func (e *InvalidValueError) Error() string {
	return fmt.Sprintf("invalid value for flag '%s': '%s'", e.Flag, e.Value)
}

func (e *InvalidValueError) Unwrap() error {
	return e.Err
}

type InvalidOptionError struct {
	Flag   string
	ErrMsg string
}

func (e *InvalidOptionError) Error() string {
	return fmt.Sprintf("invalid option for flag '%s': %s", e.Flag, e.ErrMsg)
}
