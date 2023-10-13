package mapper

import "github.com/boundedinfinity/go-commoner/functional/optioner"

func Get[K comparable, V any](m map[K]V, k K) optioner.Option[V] {
	v, ok := m[k]
	return optioner.OfOk(v, ok)
}
