package chain

import (
	"fmt"
)

type ChainErrFn[T any] func(T) (T, error)

func NewError[T any]() *ChainErr[T] {
	return &ChainErr[T]{
		steps: []ChainErrFn[T]{},
	}
}

type ChainErr[T any] struct {
	steps []ChainErrFn[T]
}

func (t *ChainErr[T]) Append(fn ChainErrFn[T]) *ChainErr[T] {
	t.steps = append(t.steps, fn)
	return t
}

func (t *ChainErr[T]) AppendNoErr(fn ChainFn[T]) *ChainErr[T] {
	return t.Append(func(t T) (T, error) {
		return fn(t), nil
	})
}

func (t *ChainErr[T]) Prepend(fn ChainErrFn[T]) *ChainErr[T] {
	t.steps = append([]ChainErrFn[T]{fn}, t.steps...)
	return t
}

func (t *ChainErr[T]) PrependNoErr(fn ChainFn[T]) *ChainErr[T] {
	return t.Prepend(func(t T) (T, error) {
		return fn(t), nil
	})
}

func (t *ChainErr[T]) Run(in T) (T, error) {
	res := in

	for i, step := range t.steps {
		if step == nil {
			return res, fmt.Errorf("step %v is nil", i)
		}

		n, err := step(res)

		if err != nil {
			return res, fmt.Errorf("failed on step %v: %w", i, err)
		}

		res = n
	}

	return res, nil
}
