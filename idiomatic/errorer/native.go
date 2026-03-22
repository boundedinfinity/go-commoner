package errorer

import (
	"errors"
	"fmt"
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
