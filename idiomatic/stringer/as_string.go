package stringer

import (
	"fmt"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func AsString(item fmt.Stringer) string {
	return item.String()
}

func AsStrings(items ...fmt.Stringer) []string {
	return slicer.Map(
		func(item fmt.Stringer) string { return item.String() },
		items...,
	)
}
