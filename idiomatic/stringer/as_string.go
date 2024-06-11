package stringer

import (
	"fmt"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func AsString(elem fmt.Stringer) string {
	return elem.String()
}

func AsStrings(elems ...fmt.Stringer) []string {
	return slicer.Map(
		func(_ int, elem fmt.Stringer) string { return elem.String() },
		elems...,
	)
}
