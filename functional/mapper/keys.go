package mapper

import (
	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

func Keys[K comparable, V any](m map[K]V) optioner.Option[[]K] {
	fn := func(_ K) bool { return true }
	return KeysFiltered(m, fn)
}

func KeysFiltered[K comparable, V any](m map[K]V, fn func(K) bool) optioner.Option[[]K] {
	var elems []K

	for k := range m {
		if fn(k) {
			elems = append(elems, k)
		}
	}

	return optioner.OfSlice(elems)
}
