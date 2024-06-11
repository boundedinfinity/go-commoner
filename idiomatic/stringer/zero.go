package stringer

import "github.com/boundedinfinity/go-commoner/idiomatic/slicer"

func IsZero[T ~string](s T) bool {
	return s == ""
}

func FindNonZero[T ~string](elems ...T) (T, bool) {
	return slicer.FindFn(func(_ int, elem T) bool {
		return !IsZero(elem)
	}, elems...)
}
