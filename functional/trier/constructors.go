package trier

import (
	"fmt"

	"github.com/boundedinfinity/go-commoner/errorer"
)

func Success[T any](result T) Try[T] {
	return Try[T]{
		Result: result,
	}
}

func Failure[T any](err error) Try[T] {
	var zero T
	return Complete(zero, err)
}

func Failuref[T any](format string, a ...any) Try[T] {
	return Failure[T](fmt.Errorf(format, a...))
}

func Complete[T any](result T, err error) Try[T] {
	var werr error

	if err != nil {
		werr = errorer.Err(err)
	}

	return Try[T]{
		Result: result,
		Error:  werr,
	}
}