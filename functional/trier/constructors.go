package trier

import (
	"fmt"

	"github.com/boundedinfinity/go-commoner/errorer"
	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

func Success[T any](result T) Try[T] {
	return Complete[T](result, nil)
}

func Failure[T any](err error) Try[T] {
	var zero T
	return Complete(zero, err)
}

func Failuref[T any](format string, a ...any) Try[T] {
	var zero T
	err := fmt.Errorf(format, a...)
	return Complete[T](zero, err)
}

func Complete[T any](result T, err error) Try[T] {
	var werr error

	if err != nil {
		werr = errorer.Err(err)
	}

	return Try[T]{
		Result: result,
		Err:    werr,
	}
}

func CompleteOfZero[T comparable](result T, err error) Try[optioner.Option[T]] {
	return Complete(optioner.OfZero(result), err)
}

func CompleteOfPtr[T any](result *T, err error) Try[optioner.Option[T]] {
	return Complete(optioner.OfPtr(result), err)
}

func CompleteOfFn[T comparable](result T, fn func(v T) bool, err error) Try[optioner.Option[T]] {
	return Complete(optioner.OfFn(result, fn), err)
}

func CompleteOfSlice[T any](result []T, err error) Try[optioner.Option[[]T]] {
	return Complete(optioner.OfSlice(result), err)
}
