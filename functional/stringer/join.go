package stringer

import (
	"github.com/boundedinfinity/go-commoner/functional/optioner"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

func Join[T ~string](sep string, slice optioner.Option[[]T]) string {
	if slice.Defined() {
		return stringer.Join(sep, slice.Get()...)
	}

	return ""
}

func JoinLift[T ~string](sep string, elems ...T) string {
	return Join(sep, optioner.OfSlice(elems))
}

func JoinOptions[T ~string](sep string, elems ...optioner.Option[T]) string {
	var realItems []T

	for _, elem := range elems {
		if elem.Defined() {
			realItems = append(realItems, elem.Get())
		}
	}

	return Join(sep, optioner.OfSlice(realItems))
}
