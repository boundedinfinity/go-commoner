package chain

import "fmt"

type chainer[T any] struct {
	steps []func(t T) (T, error)
}

func new[T any]() *chainer[T] {
	return &chainer[T]{
		steps: make([]func(t T) (T, error), 0),
	}
}

func CanErr[T any](fn func(T) (T, error)) *chainer[T] {
	return new[T]().CanErr(fn)
}

func Then[T any](fn func(T) T) *chainer[T] {
	return new[T]().Then(fn)
}

func (t *chainer[T]) CanErr(fn func(t T) (T, error)) *chainer[T] {
	t.steps = append(t.steps, fn)
	return t
}

func (t *chainer[T]) Then(fn func(T) T) *chainer[T] {
	t.steps = append(t.steps, func(t T) (T, error) {
		return fn(t), nil
	})

	return t
}

func (t *chainer[T]) Run(item T) (T, error) {
	out := item

	for i, step := range t.steps {
		if step != nil {
			r, err := step(out)

			if err != nil {
				return out, fmt.Errorf("step[%v] %w", i, err)
			}

			out = r
		}
	}

	return out, nil
}

func (t *chainer[T]) All(items []T) ([]T, error) {
	out := make([]T, 0)

	for i, item := range items {
		r, err := t.Run(item)

		if err != nil {
			return out, fmt.Errorf("items[%v] %w", i, err)
		}

		out = append(out, r)
	}

	return out, nil
}

type ChainItem[T any] struct {
	Result T
	Err    error
}

func (t *chainer[T]) Each(items []T) []ChainItem[T] {
	out := make([](ChainItem[T]), 0)

	for _, item := range items {
		r, err := t.Run(item)

		out = append(out, ChainItem[T]{
			Result: r,
			Err:    err,
		})
	}

	return out
}
