package trier

import (
	"fmt"

	"github.com/boundedinfinity/go-commoner/errorer"
	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

func Success[T any](result T) Try[T] {
	return CompleteErr[T](result, nil)
}

func Failure[T any](err error) Try[T] {
	var zero T
	return CompleteErr(zero, err)
}

func Failuref[T any](format string, a ...any) Try[T] {
	var zero T
	err := fmt.Errorf(format, a...)
	return CompleteErr[T](zero, err)
}

func CompleteErr[T any](result T, err error) Try[T] {
	var werr error

	if err != nil {
		werr = errorer.Wrap(err)
	}

	return Try[T]{
		Value: result,
		Err:   werr,
	}
}

func CompleteZero[T comparable](result T, err error) Try[optioner.Option[T]] {
	return CompleteErr(optioner.OfZero(result), err)
}

func CompletePtr[T any](result *T, err error) Try[optioner.Option[T]] {
	return CompleteErr(optioner.OfPtr(result), err)
}

func CompleteFn[T comparable](result T, fn func(v T) bool, err error) Try[optioner.Option[T]] {
	return CompleteErr(optioner.OfFn(result, fn), err)
}

func CompleteSlice[T any](result []T, err error) Try[optioner.Option[[]T]] {
	return CompleteErr(optioner.OfSlice(result), err)
}
