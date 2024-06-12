package slicer

import (
	"sort"

	"golang.org/x/exp/constraints"
)

func Sort[T constraints.Ordered](elems ...T) []T {
	fn := func(elem T) T { return elem }
	return SortFn(fn, elems...)
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
	copied := []T{}

	if fn == nil {
		return copied
	}

	copied = append(copied, elems...)

	sort.Sort(builtInFnSort[T, O]{
		elems: copied,
		fn:    fn,
	})
	return copied
}
