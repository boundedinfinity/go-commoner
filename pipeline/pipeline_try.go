package pipeline

import (
	"github.com/boundedinfinity/go-commoner/trier"
)

func NewTry[T any]() *PipelineTry[T] {
	return &PipelineTry[T]{
		pipeline: NewWithError[T](),
	}
}

type PipelineTry[T any] struct {
	pipeline *PipelineWithErr[T]
}

func (t *PipelineTry[T]) Append(fn StepWithErrFn[T]) *PipelineTry[T] {
	t.pipeline.Append(fn)
	return t
}

func (t *PipelineTry[T]) Prepend(fn StepWithErrFn[T]) *PipelineTry[T] {
	t.pipeline.Prepend(fn)
	return t
}

func (t *PipelineTry[T]) Run(in T) trier.Try[T] {
	res, err := t.pipeline.Run(in)

	if err != nil {
		return trier.Complete(res, err)
	}

	return trier.Success(res)
}
