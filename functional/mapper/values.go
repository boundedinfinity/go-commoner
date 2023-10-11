package mapper

import (
	"github.com/boundedinfinity/go-commoner/functional/optioner"
	"github.com/boundedinfinity/go-commoner/idiomatic/mapper"
)

func ValuesFiltered[K comparable, V any](m map[K]V, fn func(V) bool) optioner.Option[[]V] {
	res := mapper.ValuesFiltered(m, fn)

	if len(res) <= 0 {
		return optioner.None[[]V]()
	}

	return optioner.Some(res)
}

func Values[K comparable, V any](m map[K]V) optioner.Option[[]V] {
	return ValuesFiltered(m, func(v V) bool {
		return true
	})
}
