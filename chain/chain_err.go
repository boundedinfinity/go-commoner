package chain

import (
	"github.com/boundedinfinity/go-commoner/slices"
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

func (t *ChainErr[T]) RunSingle(item T) (T, error) {
	results, err := t.RunList([]T{item})
	return results[0], err
}

func (t *ChainErr[T]) RunList(items []T) ([]T, error) {
	results := items[:]

	for _, step := range t.steps {
		_results, err := slices.MapErr(results, step)

		if err != nil {
			return _results, err
		}

		results = _results
	}

	return results, nil
}
