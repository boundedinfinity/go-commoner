package errorer

import (
	"errors"
	"fmt"
)

// https://tip.golang.org/src/errors/wrap_test.go

type Errorer struct {
	wrapped error
}

func (e Errorer) String() string {
	if e.wrapped == nil {
		return ""
	}

	return e.wrapped.Error()
}

func (e Errorer) Error() string {
	return e.String()
}

func (e Errorer) Unwrap() error {
	return e.wrapped
}

func Errorf(format string, a ...any) error {
	if format == "" {
		return nil
	}

	return New(fmt.Sprintf(format, a...))
}

func New(message string) error {
	return &Errorer{
		wrapped: errors.New(message),
	}
}

func Wrap(err error) error {
	if err == nil {
		return nil
	}

	return &Errorer{
		wrapped: err,
	}
}

func None() error {
	return &Errorer{
		wrapped: nil,
	}
}

func WithParent(parent error) func(error) error {
	return func(child error) error {
		return fmt.Errorf("%w %w", parent, child)
	}
}

func WithParentf(parent error) func(string, ...any) error {
	return func(format string, a ...any) error {
		child := fmt.Errorf(format, a...)
		return fmt.Errorf("%w %w", parent, child)
	}
}

func WithChild(child error) func(error) error {
	return func(parent error) error {
		return fmt.Errorf("%w %w", parent, child)
	}
}

func WithChildf(child error) func(string, ...any) error {
	return func(format string, a ...any) error {
		parent := fmt.Errorf(format, a...)
		return fmt.Errorf("%w %w", parent, child)
	}
}
