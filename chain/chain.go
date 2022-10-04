package chain

import "github.com/boundedinfinity/go-commoner/slices"

type ChainFn[T any] func(T) T

func New[T any]() *Chain[T] {
	return &Chain[T]{
		steps: []ChainFn[T]{},
	}
}

type Chain[T any] struct {
	steps []ChainFn[T]
}

func (t *Chain[T]) Append(fn ChainFn[T]) *Chain[T] {
	t.steps = append(t.steps, fn)
	return t
}

func (t *Chain[T]) Prepend(fn ChainFn[T]) *Chain[T] {
	t.steps = append([]ChainFn[T]{fn}, t.steps...)
	return t
}

func (t *Chain[T]) RunSingle(item T) T {
	items := []T{item}
	results := t.RunList(items)
	return results[0]
}

func (t *Chain[T]) RunList(items []T) []T {
	results := items

	for _, step := range t.steps {
		results = slices.Map(results, step)
	}

	return results
}
