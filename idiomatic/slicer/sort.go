package slicer

import (
	"sort"

	"golang.org/x/exp/constraints"
)

type builtInSort[T constraints.Ordered] struct {
	elems []T
}

func (t builtInSort[T]) Len() int {
	return len(t.elems)
}

func (t builtInSort[T]) Less(i, j int) bool {
	return t.elems[i] < t.elems[j]
}

func (t builtInSort[T]) Swap(i, j int) {
	t.elems[i], t.elems[j] = t.elems[j], t.elems[i]
}

func Sort[T constraints.Ordered](elems ...T) []T {
	var copied []T
	copy(copied, elems)

	sort.Sort(builtInSort[T]{elems: copied})
	return copied
}

type builtInFnSort[T any, O constraints.Ordered] struct {
	elems []T
	fn    func(T) O
}

func (t builtInFnSort[T, O]) Len() int {
	return len(t.elems)
}

func (t builtInFnSort[T, O]) Less(i, j int) bool {
	return t.fn(t.elems[i]) < t.fn(t.elems[j])
}

func (t builtInFnSort[T, O]) Swap(i, j int) {
	t.elems[i], t.elems[j] = t.elems[j], t.elems[i]
}

func SortFn[T any, O constraints.Ordered](fn func(T) O, elems ...T) []T {
	var copied []T
	copy(copied, elems)

	sort.Sort(builtInFnSort[T, O]{
		elems: copied,
		fn:    fn,
	})
	return copied
}
