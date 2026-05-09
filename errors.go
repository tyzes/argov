package argov

import (
	"fmt"
)

// FlagInvalidError occurs if a flag is passed which was not registered in the used parser.
type FlagInvalidError struct {
	Flag string
}

func (e *FlagInvalidError) Error() string {
	return fmt.Sprintf("invalid flag '%s'", e.Flag)
}

// FlagSyntaxError occurs if the given string to be parsed contains invalid syntax (e.g. a '-' without flag).
type FlagSyntaxError struct {
	ErrMsg string
}

func (e *FlagSyntaxError) Error() string {
	return fmt.Sprintf("invalid syntax: %s", e.ErrMsg)
}

// MissingValueError occurs if a flag requiring a value is passed without a value.
type MissingValueError struct {
	Flag string
}

func (e *MissingValueError) Error() string {
	return fmt.Sprintf("missing value for flag '%s'", e.Flag)
}

// MissingRequiredError occurs if a flag registered with the Required() option is not passed.
type MissingRequiredError struct {
	Flag string
}

func (e *MissingRequiredError) Error() string {
	return fmt.Sprintf("missing required flag '%s'", e.Flag)
}

// InvalidValueError occurs if a provided value cannot be parsed (e.g. an integer flag given a string).
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

// InvalidOptionError occurs if an option is incompatible with the flag it is applied to (e.g. SplitOn() being passed to a non-slice value).
//
// The error will be stored and only returned by Parse().
type InvalidOptionError struct {
	Flag   string
	ErrMsg string
}

func (e *InvalidOptionError) Error() string {
	return fmt.Sprintf("invalid option for flag '%s': %s", e.Flag, e.ErrMsg)
}
