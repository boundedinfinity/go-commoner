package slicer

import (
	"sort"

	"github.com/boundedinfinity/go-commoner/idiomatic"
)

func Sort[T idiomatic.Ordered](elems ...T) []T {
	fn := func(elem T) T { return elem }
	return SortBy(fn, elems...)
}

type builtInFnSort[T any, O idiomatic.Ordered] struct {
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

func SortBy[T any, O idiomatic.Ordered](fn func(T) O, elems ...T) []T {
	sorted := []T{}

	if fn == nil {
		return sorted
	}

	sorted = append(sorted, elems...)

	sort.Sort(builtInFnSort[T, O]{
		elems: sorted,
		fn:    fn,
	})

	return sorted
}
