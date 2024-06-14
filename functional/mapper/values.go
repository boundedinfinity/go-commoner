package mapper

import (
	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

func Values[K comparable, V any](m map[K]V) optioner.Option[[]V] {
	fn := func(_ V) bool { return true }
	return ValuesFiltered(m, fn)
}

func ValuesFiltered[K comparable, V any](m map[K]V, fn func(V) bool) optioner.Option[[]V] {
	var elems []V

	for _, v := range m {
		if fn(v) {
			elems = append(elems, v)
		}
	}

	return optioner.OfSlice(elems)
}
