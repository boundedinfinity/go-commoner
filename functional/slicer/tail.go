package slicer

import "github.com/boundedinfinity/go-commoner/functional/optioner"

func Tail[T any](elems ...T) optioner.Option[T] {
	l := len(elems)

	if l > 0 {
		return optioner.Some(elems[l-1])
	} else {
		return optioner.None[T]()
	}
}
