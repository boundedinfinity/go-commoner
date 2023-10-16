package stringer

import "github.com/boundedinfinity/go-commoner/idiomatic/slicer"

func IsZero[T ~string](s T) bool {
	return s == ""
}

func IsEmpty[T ~string](s T) bool {
	return IsZero(s)
}

func FindNonZero[T ~string](items ...T) (T, bool) {
	return slicer.FindFn(func(item T) bool {
		return !IsZero(item)
	}, items...)
}
