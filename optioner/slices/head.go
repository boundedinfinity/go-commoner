package slices

import "github.com/boundedinfinity/go-commoner/optioner"

func Head[V any](vs []V) optioner.Option[V] {
	if len(vs) > 0 {
		return optioner.Some(vs[0])
	}

	return optioner.None[V]()
}

func HeadVO[V any](vs ...V) optioner.Option[V] {
	return Head(vs)
}
