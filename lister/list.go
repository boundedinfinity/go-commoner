package lister

import (
	"github.com/boundedinfinity/go-commoner/optioner"
)

type list[T any] struct {
	items []T
}

func NewList[T any]() *list[T] {
	s := &list[T]{
		items: make([]T, 0),
	}

	return s
}

func (t *list[T]) Len() int {
	return len(t.items)
}

func (t *list[T]) Empty() bool {
	return len(t.items) > 0
}

func (t *list[T]) PushF(items ...T) {
	t.items = append(items, t.items...)
}

func (t *list[T]) PushB(items ...T) {
	t.items = append(t.items, items...)
}

func (t *list[T]) PeekB() optioner.Option[T] {
	l := len(t.items) - 1

	if l > -1 {
		return optioner.Some(t.items[l])
	}

	return optioner.None[T]()
}

func (t *list[T]) PopB() optioner.Option[T] {
	l := len(t.items) - 1

	if l > -1 {
		item := t.items[l]
		t.items = t.items[:l]
		return optioner.Some(item)
	}

	return optioner.None[T]()
}

func (t *list[T]) PeekF() optioner.Option[T] {
	if len(t.items) > 0 {
		return optioner.Some(t.items[0])
	}

	return optioner.None[T]()
}

func (t *list[T]) PopF() optioner.Option[T] {
	l := len(t.items) - 1

	if l > -1 {
		item := t.items[0]
		t.items = t.items[1:]
		return optioner.Some(item)
	}

	return optioner.None[T]()
}
