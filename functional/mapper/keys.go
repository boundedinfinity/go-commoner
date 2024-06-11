package mapper

import (
	"github.com/boundedinfinity/go-commoner/functional/optioner"
	"github.com/boundedinfinity/go-commoner/idiomatic/mapper"
)

func Keys[K comparable, V any](m map[K]V) optioner.Option[[]K] {
	return KeysFiltered(m, func(k K) bool {
		return true
	})
}

func KeysFiltered[K comparable, V any](m map[K]V, fn func(K) bool) optioner.Option[[]K] {
	return optioner.OfFn(mapper.KeysFiltered(m, fn), func(elems []K) bool {
		return len(elems) > 0
	})
}
