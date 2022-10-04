package slices

import "github.com/boundedinfinity/go-commoner/optioner"

func Tail[V any](vs []V) optioner.Option[V] {
	l := len(vs)

	if l > 0 {
		return optioner.Some(vs[l-1])
	} else {
		return optioner.None[V]()
	}
}

func TailVO[V any](vs ...V) optioner.Option[V] {
	return Tail(vs)
}
