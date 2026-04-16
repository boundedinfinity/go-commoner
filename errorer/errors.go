package errorer

import (
	"errors"
	"fmt"
	"strings"
)

func Is(err, target error) bool {
	return errors.Is(err, target)
}

func As(err error, target any) bool {
	return errors.As(err, target)
}

func Join(errs ...error) error {
	return errors.Join(errs...)
}

func Errorf(format string, a ...any) error {
	return fmt.Errorf(format, a...)
}

func New(s string) error {
	return errors.New(s)
}

func Func(errs ...error) func(format string, a ...any) error {
	errFormat := strings.Repeat("%w : ", len(errs))
	errAny := make([]any, len(errs))

	for i, err := range errs {
		if err == nil {
			panic("err is null")
		}

		errAny[i] = err
	}

	return func(format string, a ...any) error {
		return fmt.Errorf(errFormat+format, append(errAny, a...)...)
	}
}

type errorHandler struct {
	errFormat string
	errAny    []any
}

func (this errorHandler) New(err error) error {
	return this.Errorf("%w", err)
}

func (this errorHandler) Errorf(format string, a ...any) error {
	return fmt.Errorf(this.errFormat+format, append(this.errAny, a...)...)
}

func Handler(errs ...error) errorHandler {
	handler := errorHandler{
		errFormat: strings.Repeat("%w : ", len(errs)),
		errAny:    make([]any, len(errs)),
	}

	for i, err := range errs {
		if err == nil {
			panic("err is null")
		}

		handler.errAny[i] = err
	}

	return handler
}
