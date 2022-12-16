package chain

import (
	"fmt"

	"github.com/boundedinfinity/go-commoner/trier"
)

type chainer[T any] struct {
	steps []func(T) trier.Try[T]
}

func new[T any]() *chainer[T] {
	return &chainer[T]{
		steps: make([]func(T) trier.Try[T], 0),
	}
}

func Try[T any](fn func(T) trier.Try[T]) *chainer[T] {
	return new[T]().Try(fn)
}

func Then[T any](fn func(T) T) *chainer[T] {
	return new[T]().Then(fn)
}

func CanErr[T any](fn func(T) (T, error)) *chainer[T] {
	return new[T]().CanErr(fn)
}

func (t *chainer[T]) Try(fn func(T) trier.Try[T]) *chainer[T] {
	t.steps = append(t.steps, fn)
	return t
}

func (t *chainer[T]) CanErr(fn func(T) (T, error)) *chainer[T] {
	t.steps = append(t.steps, func(t T) trier.Try[T] {
		r, err := fn(t)
		return trier.Complete(r, err)
	})

	return t
}

func (t *chainer[T]) Then(fn func(T) T) *chainer[T] {
	t.steps = append(t.steps, func(t T) trier.Try[T] {
		r := fn(t)
		return trier.Success(r)
	})

	return t
}

func (t *chainer[T]) Run(item T) trier.Try[T] {
	x := item

	for i, step := range t.steps {
		r := step(item)

		if r.Failure() {
			return trier.Failuref[T]("failed at step %v: %w", i, r.Error)
		}

		x = r.Result
	}

	return trier.Success(x)
}

func (t *chainer[T]) Each(items []T) []trier.Try[T] {
	var xs []trier.Try[T]

	for i, item := range items {
		r := t.Run(item)

		if r.Failure() {
			xs = append(xs, trier.Failuref[T]("item %v %w", i, r.Error))
			continue
		}

		xs = append(xs, r)
	}

	return xs
}

func (t *chainer[T]) All(items []T) trier.Try[[]T] {
	var xs []T

	for i, item := range items {
		r := t.Run(item)

		if r.Failure() {
			err := fmt.Errorf("item %v %w", i, r.Error)
			return trier.Complete(xs, err)
		}

		xs = append(xs, r.Result)
	}

	return trier.Success(xs)
}
