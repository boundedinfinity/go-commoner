package pipeline

import (
	"fmt"
)

type StepWithErrFn[T any] func(T) (T, error)

func SimpleStep[T any](fn func(T) T) StepWithErrFn[T] {
	return func(t T) (T, error) {
		return fn(t), nil
	}
}

func NewWithError[T any]() *PipelineWithErr[T] {
	return &PipelineWithErr[T]{
		steps: []StepWithErrFn[T]{},
	}
}

type PipelineWithErr[T any] struct {
	steps []StepWithErrFn[T]
}

func (t *PipelineWithErr[T]) Append(fn StepWithErrFn[T]) *PipelineWithErr[T] {
	t.steps = append(t.steps, fn)
	return t
}

func (t *PipelineWithErr[T]) Prepend(fn StepWithErrFn[T]) *PipelineWithErr[T] {
	t.steps = append([]StepWithErrFn[T]{fn}, t.steps...)
	return t
}

func (t *PipelineWithErr[T]) Run(in T) (T, error) {
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
