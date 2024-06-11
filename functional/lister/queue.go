package lister

import (
	"github.com/boundedinfinity/go-commoner/functional/optioner"
	"github.com/boundedinfinity/go-commoner/idiomatic/lister"
)

type Queue[T any] struct {
	idiomatic *lister.Queue[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		idiomatic: lister.NewQueue[T](),
	}
}

func (t *Queue[T]) Empty() bool {
	return t.idiomatic.Empty()
}

func (t *Queue[T]) Len() int {
	return t.idiomatic.Len()
}

func (t *Queue[T]) Push(elems ...T) {
	t.idiomatic.Push(elems...)
}

func (t *Queue[T]) Pop() optioner.Option[T] {
	return optioner.OfOk(t.idiomatic.Pop())
}

func (t *Queue[T]) Peek() optioner.Option[T] {
	return optioner.OfOk(t.idiomatic.Peek())
}
