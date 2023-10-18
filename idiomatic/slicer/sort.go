package slicer

import (
	"sort"

	"golang.org/x/exp/constraints"
)

type builtInSort[T constraints.Ordered] struct {
	items []T
}

func (t builtInSort[T]) Len() int {
	return len(t.items)
}

func (t builtInSort[T]) Less(i, j int) bool {
	return t.items[i] < t.items[j]
}

func (t builtInSort[T]) Swap(i, j int) {
	t.items[i], t.items[j] = t.items[j], t.items[i]
}

func Sort[T constraints.Ordered](items ...T) []T {
	copy := Copy(items...)
	sort.Sort(builtInSort[T]{items: copy})
	return copy
}
