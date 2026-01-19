package slicer

import (
	"errors"
	"fmt"
)

var (
	ErrReduceExitEarly = errors.New("reduce exit")
	ErrReduceFailure   = errors.New("reduce error")
	errReduceFailureFn = func(i int, format string, a ...any) error {
		message := fmt.Sprintf(format, a...)
		return fmt.Errorf("%w at index %d: %s", ErrReduceExitEarly, i, message)
	}
)

func Reduce[E any, A any](fn func(int, A, E) A, initial A, elems ...E) A {
	fn2 := func(i int, acc A, elem E) (A, error) {
		return fn(i, acc, elem), nil
	}

	result, _ := ReduceErr(fn2, initial, elems...)
	return result
}

func ReduceErr[E any, A any](fn func(int, A, E) (A, error), initial A, elems ...E) (A, error) {
	result := initial
	var err error

	for i, elem := range elems {
		result, err = fn(i, result, elem)

		if errors.Is(err, ErrReduceExitEarly) {
			err = nil
			break
		}

		if err != nil {
			err = errReduceFailureFn(i, err.Error())
			break
		}
	}

	return result, err
}
