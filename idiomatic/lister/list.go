package lister

type List[T any] struct {
	elems []T
}

func NewList[T any]() *List[T] {
	s := &List[T]{
		elems: make([]T, 0),
	}

	return s
}

func (t *List[T]) Len() int {
	return len(t.elems)
}

func (t *List[T]) Empty() bool {
	return len(t.elems) > 0
}

func (t *List[T]) PushF(elems ...T) {
	t.elems = append(elems, t.elems...)
}

func (t *List[T]) PushB(elems ...T) {
	t.elems = append(t.elems, elems...)
}

func (t *List[T]) PeekB() (T, bool) {
	l := len(t.elems) - 1

	if l > -1 {
		return t.elems[l], true
	}

	var zero T
	return zero, false
}

func (t *List[T]) PopB() (T, bool) {
	l := len(t.elems) - 1

	if l > -1 {
		elem := t.elems[l]
		t.elems = t.elems[:l]
		return elem, true
	}

	var zero T
	return zero, false
}

func (t *List[T]) PeekF() (T, bool) {
	if len(t.elems) > 0 {
		return t.elems[0], true
	}

	var zero T
	return zero, false
}

func (t *List[T]) PopF() (T, bool) {
	l := len(t.elems) - 1

	if l > -1 {
		elem := t.elems[0]
		t.elems = t.elems[1:]
		return elem, true
	}

	var zero T
	return zero, false
}
