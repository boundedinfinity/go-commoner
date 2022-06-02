package pipeline

import (
	"github.com/boundedinfinity/commons/try"
)

type StepTryFn[T any] func(T) try.Try[T]

func NewTry[T any]() *PipelineTry[T] {
	return &PipelineTry[T]{
		steps: []StepTryFn[T]{},
	}
}

type PipelineTry[T any] struct {
	initial T
	steps   []StepTryFn[T]
}

func (t *PipelineTry[T]) Step(fn StepTryFn[T]) *PipelineTry[T] {
	t.steps = append(t.steps, fn)
	return t
}

func (t *PipelineTry[T]) Run(in T) try.Try[T] {
	next := in

	for i, step := range t.steps {
		if step == nil {
			return try.Failuref[T]("step %v is nil", i)
		}

		n := step(next)

		if n.Failure() {
			return n
		}

		next = n.Result
	}

	return try.Success(next)
}
