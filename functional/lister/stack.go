package lister

import (
	"github.com/boundedinfinity/go-commoner/functional/optioner"
	"github.com/boundedinfinity/go-commoner/idiomatic/lister"
)

type Stack[T any] struct {
	idiomatic *lister.Stack[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		idiomatic: lister.NewStack[T](),
	}
}

func (t *Stack[T]) Empty() bool {
	return t.idiomatic.Empty()
}

func (t *Stack[T]) Len() int {
	return t.idiomatic.Len()
}

func (t *Stack[T]) Push(elems ...T) {
	t.idiomatic.Push(elems...)
}

func (t *Stack[T]) Pop() optioner.Option[T] {
	return optioner.OfOk(t.idiomatic.Pop())
}

func (t *Stack[T]) Peek() optioner.Option[T] {
	return optioner.OfOk(t.idiomatic.Peek())
}
