package pipeline

import (
	"fmt"
)

type StepFn[T any] func(T) (T, error)

func New[T any]() *Pipeline[T] {
	return &Pipeline[T]{
		steps: []StepFn[T]{},
	}
}

type Pipeline[T any] struct {
	initial T
	steps   []StepFn[T]
}

func (t *Pipeline[T]) Step(fn StepFn[T]) *Pipeline[T] {
	t.steps = append(t.steps, fn)
	return t
}

func (t *Pipeline[T]) Run(in T) (T, error) {
	next := in

	for i, step := range t.steps {
		if step == nil {
			return next, fmt.Errorf("step %v is nil", i)
		}

		n, err := step(next)

		if err != nil {
			return next, fmt.Errorf("failed on step %v: %w", i, err)
		}

		next = n
	}

	return next, nil
}
