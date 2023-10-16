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

func JoinLift[T ~string](sep string, items ...T) string {
	return Join(sep, optioner.OfSlice(items))
}

func JoinOptions[T ~string](sep string, items ...optioner.Option[T]) string {
	var realItems []T

	for _, item := range items {
		if item.Defined() {
			realItems = append(realItems, item.Get())
		}
	}

	return Join(sep, optioner.OfSlice(realItems))
}
