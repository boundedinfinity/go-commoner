package pipeline

import "github.com/boundedinfinity/go-commoner/slices"

type StepFn[T any] func(T) T

func New[T any]() *Pipeline[T] {
	return &Pipeline[T]{
		steps: []StepFn[T]{},
	}
}

type Pipeline[T any] struct {
	steps []StepFn[T]
}

func (t *Pipeline[T]) Append(fn StepFn[T]) *Pipeline[T] {
	t.steps = append(t.steps, fn)
	return t
}

func (t *Pipeline[T]) Prepend(fn StepFn[T]) *Pipeline[T] {
	t.steps = append([]StepFn[T]{fn}, t.steps...)
	return t
}

func (t *Pipeline[T]) RunSingle(item T) T {
	items := []T{item}
	results := t.RunList(items)
	return results[0]
}

func (t *Pipeline[T]) RunList(items []T) []T {
	results := items

	for _, step := range t.steps {
		results = slices.Map(results, step)
	}

	return results
}
