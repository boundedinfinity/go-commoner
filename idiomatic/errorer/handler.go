package errorer

import (
	"fmt"
	"strings"
)

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
