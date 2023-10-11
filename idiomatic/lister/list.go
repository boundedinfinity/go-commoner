package lister

type List[T any] struct {
	items []T
}

func NewList[T any]() *List[T] {
	s := &List[T]{
		items: make([]T, 0),
	}

	return s
}

func (t *List[T]) Len() int {
	return len(t.items)
}

func (t *List[T]) Empty() bool {
	return len(t.items) > 0
}

func (t *List[T]) PushF(items ...T) {
	t.items = append(items, t.items...)
}

func (t *List[T]) PushB(items ...T) {
	t.items = append(t.items, items...)
}

func (t *List[T]) PeekB() (T, bool) {
	l := len(t.items) - 1

	if l > -1 {
		return t.items[l], true
	}

	var zero T
	return zero, false
}

func (t *List[T]) PopB() (T, bool) {
	l := len(t.items) - 1

	if l > -1 {
		item := t.items[l]
		t.items = t.items[:l]
		return item, true
	}

	var zero T
	return zero, false
}

func (t *List[T]) PeekF() (T, bool) {
	if len(t.items) > 0 {
		return t.items[0], true
	}

	var zero T
	return zero, false
}

func (t *List[T]) PopF() (T, bool) {
	l := len(t.items) - 1

	if l > -1 {
		item := t.items[0]
		t.items = t.items[1:]
		return item, true
	}

	var zero T
	return zero, false
}
