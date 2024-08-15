package stringer

import "github.com/boundedinfinity/go-commoner/idiomatic/slicer"

func Duplicate[T ~string](count int, elems ...T) []T {
	return slicer.Duplicate(count, elems...)
}
