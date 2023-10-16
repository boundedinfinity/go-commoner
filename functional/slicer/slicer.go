package slicer

import "github.com/boundedinfinity/go-commoner/functional/optioner"

type Slicer[T any] struct {
	optioner.Option[[]T]
}

func (t Slicer[T]) Items() []T {
	return t.OrElse([]T{})
}

func (t *Slicer[T]) Len() int {
	return len(t.Get())
}

func (t *Slicer[T]) Append(item T) {
	t.Option = optioner.OfSlice(append(t.Get(), item))
}
