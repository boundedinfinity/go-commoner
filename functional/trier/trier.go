package trier

import "errors"

var errTryForgotToSetError = errors.New("You forgot to assign and error here")

type Try[T any] struct {
	Value T     `json:"value"`
	Err   error `json:"error"`
}

func (t Try[T]) Failed() bool {
	return t.Err != nil
}

func (t Try[T]) Succeeded() bool {
	return !t.Failed()
}

func (t Try[T]) Error() string {
	if t.Err != nil {
		return t.Err.Error()
	}

	return errTryForgotToSetError.Error()
}
