package validater

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

var ErrStringIsInvalid = errors.New("string")

// Builder

func String[T ~string]() *StringValidatorBuilder[T] {
	return &StringValidatorBuilder[T]{
		fns: make([]func(T) error, 0),
	}
}

type StringValidatorBuilder[T ~string] struct {
	fns []func(T) error
}

func (t *StringValidatorBuilder[T]) Validate(item T, groups ...string) error {
	for _, fn := range t.fns {
		if err := fn(item); err != nil {
			return err
		}
	}

	return nil
}

func (t *StringValidatorBuilder[T]) Min(min int) *StringValidatorBuilder[T] {
	t.fns = append(t.fns, StringMin[T](min))
	return t
}

func (t *StringValidatorBuilder[T]) Max(max int) *StringValidatorBuilder[T] {
	t.fns = append(t.fns, StringMax[T](max))
	return t
}

func (t *StringValidatorBuilder[T]) NotZero() *StringValidatorBuilder[T] {
	t.fns = append(t.fns, StringNotEmpty[T]())
	return t
}

func (t *StringValidatorBuilder[T]) Matches(regex string) *StringValidatorBuilder[T] {
	t.fns = append(t.fns, StringMatches[T](regex))
	return t
}

// Stand alone

func StringMatches[T ~string](regex string) func(s T) error {
	re := regexp.MustCompile(regex)

	return func(s T) error {
		if !re.MatchString(string(s)) {
			return fmt.Errorf("%w doesn't match regular expression %v", ErrStringIsInvalid, s)
		}

		return nil
	}
}

func StringNotEmpty[T ~string]() func(s T) error {
	return func(s T) error {
		if stringer.IsEmpty(s) {
			return fmt.Errorf("%w is empty value", ErrStringIsInvalid)
		}

		return nil
	}
}

func StringMin[T ~string](min int) func(s T) error {
	return func(s T) error {
		if len(s) < min {
			return fmt.Errorf("%w less than %v", ErrStringIsInvalid, min)
		}

		return nil
	}
}

func StringMax[T ~string](max int) func(s T) error {
	return func(s T) error {
		if len(s) > max {
			return fmt.Errorf("%w greater than %v", ErrStringIsInvalid, max)
		}

		return nil
	}
}
