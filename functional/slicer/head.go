package slicer

import "github.com/boundedinfinity/go-commoner/functional/optioner"

func Head[T any](items ...T) optioner.Option[T] {
	if len(items) > 0 {
		return optioner.Some(items[0])
	}

	return optioner.None[T]()
}
