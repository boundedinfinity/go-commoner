package lister

import (
	"github.com/boundedinfinity/go-commoner/functional/optioner"
	"github.com/boundedinfinity/go-commoner/idiomatic/lister"
)

type List[T any] struct {
	idiomatic *lister.List[T]
}

func NewList[T any]() *List[T] {
	s := &List[T]{
		idiomatic: lister.NewList[T](),
	}

	return s
}

func (t *List[T]) Len() int {
	return t.idiomatic.Len()
}

func (t *List[T]) Empty() bool {
	return t.idiomatic.Empty()
}

func (t *List[T]) PushF(elems ...T) {
	t.idiomatic.PushF(elems...)
}

func (t *List[T]) PushB(elems ...T) {
	t.idiomatic.PushB(elems...)
}

func (t *List[T]) PeekB() optioner.Option[T] {
	return optioner.OfOk(t.idiomatic.PeekB())
}

func (t *List[T]) PopB() optioner.Option[T] {
	return optioner.OfOk(t.idiomatic.PopB())
}

func (t *List[T]) PeekF() optioner.Option[T] {
	return optioner.OfOk(t.idiomatic.PeekF())
}

func (t *List[T]) PopF() optioner.Option[T] {
	return optioner.OfOk(t.idiomatic.PopF())
}
