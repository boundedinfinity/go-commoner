package mapper

import (
	"github.com/boundedinfinity/go-commoner/mapper"
	"github.com/boundedinfinity/go-commoner/optioner"
)

func ValuesFiltered[K comparable, V any](m map[K]V, fn func(V) bool) optioner.Option[[]V] {
	if m == nil || len(m) <= 0 {
		return optioner.None[[]V]()
	}

	res := mapper.ValuesFiltered(m, fn)

	if res == nil || len(res) <= 0 {
		return optioner.None[[]V]()
	}

	return optioner.Some(res)
}

func Values[K comparable, V any](m map[K]V) optioner.Option[[]V] {
	return ValuesFiltered(m, func(v V) bool {
		return true
	})
}
