package slicer

import "github.com/boundedinfinity/go-commoner/idiomatic/reflecter"

func NotZero[T any](elems ...T) (T, bool) {
	for _, elem := range elems {
		if !reflecter.IsZero[T](elem) {
			return elem, true
		}
	}

	var zero T
	return zero, false
}
