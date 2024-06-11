package slicer

import "github.com/boundedinfinity/go-commoner/functional/optioner"

func Head[T any](elems ...T) optioner.Option[T] {
	if len(elems) > 0 {
		return optioner.Some(elems[0])
	}

	return optioner.None[T]()
}
