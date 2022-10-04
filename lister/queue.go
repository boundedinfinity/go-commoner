package lister

import "github.com/boundedinfinity/go-commoner/optioner"

type queue[T any] struct {
	list *list[T]
}

func NewQueue[T any]() *stack[T] {
	return &stack[T]{
		list: NewList[T](),
	}
}

func (t *queue[T]) Empty() bool {
	return t.list.Empty()
}

func (t *queue[T]) Len() int {
	return t.list.Len()
}

func (t *queue[T]) Push(items ...T) {
	t.list.PushB(items...)
}

func (t *queue[T]) Pop() optioner.Option[T] {
	return t.list.PopF()
}

func (t *queue[T]) Peek() optioner.Option[T] {
	return t.list.PeekF()
}
