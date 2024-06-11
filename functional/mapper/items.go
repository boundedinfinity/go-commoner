package mapper

import (
	"github.com/boundedinfinity/go-commoner/functional/optioner"
	"github.com/boundedinfinity/go-commoner/idiomatic/mapper"
)

type Item[K comparable, V any] mapper.Item[K, V]

func Items[K comparable, V any](m map[K]V) optioner.Option[[]Item[K, V]] {
	var elems []Item[K, V]

	for k, v := range m {
		elems = append(elems, Item[K, V]{K: k, V: v})
	}

	return optioner.OfSlice(elems)
}

// func ItemsFiltered[K comparable, V any](m map[K]V, fn func(K, V) bool) optioner.Option[[]Item[K, V]] {
// 	convert := func(_ int, elem Item[K, V]) mapper.Item[K, V] {
// 		return mapper.Item[K, V]{K: elem.K, V: elem.V}
// 	}

//     inputs := slicer.Map(convert, m)

// 	results := mapper.ItemsFiltered()

// 	return optioner.OfSlice(results)
// }
