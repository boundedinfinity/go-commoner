package lister

import "github.com/boundedinfinity/go-commoner/optioner"

type stack[T any] struct {
	list *list[T]
}

func NewStack[T any]() *stack[T] {
	return &stack[T]{
		list: NewList[T](),
	}
}

func (t *stack[T]) Empty() bool {
	return t.list.Empty()
}

func (t *stack[T]) Len() int {
	return t.list.Len()
}

func (t *stack[T]) Push(items ...T) {
	t.list.PushF(items...)
}

func (t *stack[T]) Pop() optioner.Option[T] {
	return t.list.PopF()
}

func (t *stack[T]) Peek() optioner.Option[T] {
	return t.list.PeekF()
}
