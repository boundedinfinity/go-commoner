package lister

type Stack[T any] struct {
	list *List[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		list: NewList[T](),
	}
}

func (t *Stack[T]) Empty() bool {
	return t.list.Empty()
}

func (t *Stack[T]) Len() int {
	return t.list.Len()
}

func (t *Stack[T]) Push(items ...T) {
	t.list.PushF(items...)
}

func (t *Stack[T]) Pop() (T, bool) {
	return t.list.PopF()
}

func (t *Stack[T]) Peek() (T, bool) {
	return t.list.PeekF()
}
