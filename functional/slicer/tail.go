package slicer

import "github.com/boundedinfinity/go-commoner/functional/optioner"

func Tail[T any](items ...T) optioner.Option[T] {
	l := len(items)

	if l > 0 {
		return optioner.Some(items[l-1])
	} else {
		return optioner.None[T]()
	}
}
