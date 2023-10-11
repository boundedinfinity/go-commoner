package lister

type Queue[T any] struct {
	list *List[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		list: NewList[T](),
	}
}

func (t *Queue[T]) Empty() bool {
	return t.list.Empty()
}

func (t *Queue[T]) Len() int {
	return t.list.Len()
}

func (t *Queue[T]) Push(items ...T) {
	t.list.PushB(items...)
}

func (t *Queue[T]) Pop() (T, bool) {
	return t.list.PopF()
}

func (t *Queue[T]) Peek() (T, bool) {
	return t.list.PeekF()
}
