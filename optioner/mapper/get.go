package mapper

import "github.com/boundedinfinity/go-commoner/optioner"

func Get[K comparable, V any](m map[K]V, k K) optioner.Option[V] {
	v, ok := m[k]

	if ok {
		return optioner.Some(v)
	}

	return optioner.None[V]()
}
