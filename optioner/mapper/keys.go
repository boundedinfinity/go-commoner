package mapper

import (
	"github.com/boundedinfinity/go-commoner/mapper"
	"github.com/boundedinfinity/go-commoner/optioner"
)

func KeysFiltered[K comparable, V any](m map[K]V, fn func(K) bool) optioner.Option[[]K] {
	if m == nil || len(m) <= 0 {
		return optioner.None[[]K]()
	}

	res := mapper.KeysFiltered(m, fn)

	if res == nil || len(res) <= 0 {
		return optioner.None[[]K]()
	}

	return optioner.Some(res)
}

func Keys[K comparable, V any](m map[K]V) optioner.Option[[]K] {
	return KeysFiltered(m, func(k K) bool {
		return true
	})
}
