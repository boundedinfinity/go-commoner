package trier

import (
	"fmt"

	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

func Complete[T any](result T, err error) Try[T] {
	return Try[T]{
		Value: result,
		Err:   err,
	}
}

func Success[T any](result T) Try[T] {
	return Complete(result, nil)
}

func Error[T any](err error) Try[T] {
	var zero T
	return Complete(zero, err)
}

func Errorf[T any](format string, a ...any) Try[T] {
	var zero T
	err := fmt.Errorf(format, a...)
	return Complete(zero, err)
}

func Zero[T comparable](result T, err error) Try[optioner.Option[T]] {
	return Complete(optioner.OfZero(result), err)
}

func Ptr[T any](result *T, err error) Try[optioner.Option[T]] {
	return Complete(optioner.OfPtr(result), err)
}

func Func[T comparable](result T, fn func(v T) bool, err error) Try[optioner.Option[T]] {
	return Complete(optioner.OfFn(result, fn), err)
}

func Slice[T any](result []T, err error) Try[optioner.Option[[]T]] {
	return Complete(optioner.OfSlice(result), err)
}
