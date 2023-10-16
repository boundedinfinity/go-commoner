package mapper

import (
	"github.com/boundedinfinity/go-commoner/functional/optioner"
	"github.com/boundedinfinity/go-commoner/idiomatic/mapper"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

type Item[K comparable, V any] mapper.Item[K, V]

func Items[K comparable, V any](m map[K]V) optioner.Option[[]Item[K, V]] {
	fn := func(K, V) bool { return true }
	return ItemsFiltered(m, fn)
}

func ItemsFiltered[K comparable, V any](m map[K]V, fn func(K, V) bool) optioner.Option[[]Item[K, V]] {
	results := slicer.Map(func(item mapper.Item[K, V]) Item[K, V] {
		return Item[K, V]{K: item.K, V: item.V}
	}, mapper.ItemsFiltered(m, fn)...)

	return optioner.OfSlice(results)
}
