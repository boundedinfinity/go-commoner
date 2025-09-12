package stringer

import (
	"fmt"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func AsString[T any](v T) string {
	return fmt.Sprintf("%v", v)
}

func AsStrings[T any](elems ...T) []string {
	return slicer.Map(
		func(elem T) string { return fmt.Sprintf("%v", elem) },
		elems...,
	)
}
