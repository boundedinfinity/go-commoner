package errorer

import (
	"fmt"
	"strings"
)

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
