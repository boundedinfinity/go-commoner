package mapper

import (
	"github.com/boundedinfinity/go-commoner/functional/optioner"
	"github.com/boundedinfinity/go-commoner/idiomatic/mapper"
)

type Item[K comparable, V any] mapper.Item[K, V]

func Items[K comparable, V any](m map[K]V) optioner.Option[[]Item[K, V]] {
	fn := func(_ K, _ V) bool { return true }
	return ItemsFiltered(m, fn)
}

func ItemsFiltered[K comparable, V any](m map[K]V, fn func(K, V) bool) optioner.Option[[]Item[K, V]] {
	var elems []Item[K, V]

	for k, v := range m {
		if fn(k, v) {
			elems = append(elems, Item[K, V]{K: k, V: v})
		}
	}

	return optioner.OfSlice(elems)
}
