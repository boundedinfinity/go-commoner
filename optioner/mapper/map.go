package mapper

import "github.com/boundedinfinity/go-commoner/optioner"

type Mapper[K comparable, V any] map[K]V

func (t Mapper[K, V]) Has(k K) bool {
	return Has(t, k)
}

func (t Mapper[K, V]) Get(k K) optioner.Option[V] {
	return Get(t, k)
}

func (t Mapper[K, V]) Keys(k K) optioner.Option[[]K] {
	return Keys(t)
}

func (t Mapper[K, V]) Values(k K) optioner.Option[[]V] {
	return Values(t)
}

func (t Mapper[K, V]) Each(fn func(K, V)) {
	Each(t, fn)
}

func (t Mapper[K, V]) Eacherr(fn func(K, V) error) error {
	return EachErr(t, fn)
}
