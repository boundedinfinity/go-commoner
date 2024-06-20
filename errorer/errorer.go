package errorer

import (
	"encoding/json"
	"errors"
	"fmt"
)

// ----------------------------------------------------------------------------------------------------
// Constructors
// ----------------------------------------------------------------------------------------------------

func New(message string) *Errorer {
	return Errorf(message)
}

func Errorf(format string, a ...any) *Errorer {
	return FromList(fmt.Errorf(format, a...))
}

func FromList(errs ...error) *Errorer {
	return &Errorer{
		Errors: errs,
	}
}

type Errorer struct {
	Errors []error
}

func FormatFn(format string, errs ...error) func(a ...any) error {
	return FromList(errs...).FormatFn(format)
}

func ValueFn(errs ...error) func(v any) error {
	return FromList(errs...).ValueFn()
}

// ----------------------------------------------------------------------------------------------------
// Error interface
// ----------------------------------------------------------------------------------------------------

func (e Errorer) Error() string {
	switch len(e.Errors) {
	case 0:
		return ""
	case 1:
		return e.Errors[0].Error()
	default:
		combined := e.Errors[0]

		for _, err := range e.Errors[1:] {
			combined = fmt.Errorf("%w : %w", err, combined)
		}

		return combined.Error()
	}
}

func (e *Errorer) Is(err error) bool {
	for _, possible := range e.Errors {
		if errors.Is(err, possible) {
			return true
		}
	}

	return false
}

func (e *Errorer) Unwrap() []error {
	return e.Errors
}

// ----------------------------------------------------------------------------------------------------
// Stringer interface
// ----------------------------------------------------------------------------------------------------

func (e Errorer) String() string {
	return e.Error()
}

// ----------------------------------------------------------------------------------------------------
// Utility functions
// ----------------------------------------------------------------------------------------------------

// List returns the extracted error messages from the internal list of errors.
func (e Errorer) List() []string {
	var errs []string

	for _, err := range e.Errors {
		if err != nil && err.Error() != "" {
			errs = append(errs, err.Error())
		}
	}

	return errs
}

// Add appends an error to the internal ist of errors.
func (e *Errorer) Add(errs ...error) *Errorer {
	e.Errors = append(e.Errors, errs...)
	return e
}

// Sub returns a new instance with the errs included in it's internal list.
func (e Errorer) Sub(errs ...error) *Errorer {
	return FromList(append(e.Errors, errs...)...)
}

// Sub returns a new instance with an error created using fmt.Errorf and appended internal list of
// errors.
func (e Errorer) Subf(format string, a ...any) *Errorer {
	return FromList(append(e.Errors, fmt.Errorf(format, a...))...)
}

// FormatFn takes a format string as the input and returns a function that accepts a list of
// input arguments which will be supplied to the format string to create a new instance
// with the created error appended to it's internal list.
func (e Errorer) FormatFn(format string) func(a ...any) error {
	return func(a ...any) error { return e.Subf(format, a...) }
}

// ValueFn returns a function that accepts an input argument which will be supplied to the
// format string to create a new instance with the created error appended to it's internal list.
//
// This function basically perform the following:
//
//	FormatFn("%v")
func (e Errorer) ValueFn() func(v any) error {
	return func(v any) error { return e.Subf("%v", v) }
}

// WithValue returns an error the supplied value appened as and error appended to it's internal
// list.
//
// This function basically perform the following:
//
//	ValueFn()(v)
func (e Errorer) WithValue(v any) error {
	return e.ValueFn()(v)
}

// ----------------------------------------------------------------------------------------------------
// Marshal functions
// ----------------------------------------------------------------------------------------------------

func (e *Errorer) MarshalJSON() ([]byte, error) {
	if e.Errors == nil {
		return json.Marshal(nil)
	}

	raw := struct {
		Errors []string `json:"errors,omitempty"`
	}{}

	for _, err := range e.Errors {
		if err != nil {
			raw.Errors = append(raw.Errors, err.Error())
		}
	}

	return json.Marshal(raw)
}

func (e *Errorer) UnmarshalJSON(data []byte) error {
	if data == nil {
		return nil
	}

	raw := struct {
		Errors []string `json:"errors,omitempty"`
	}{}

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	for _, s := range raw.Errors {
		e.Errors = append(e.Errors, errors.New(s))
	}

	return nil
}
