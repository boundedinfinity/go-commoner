package stringer

import "github.com/boundedinfinity/go-commoner/idiomatic/slicer"

func IsDefined[T ~string](s T) bool {
	return !IsEmpty[T](s)
}

func IsEmpty[T ~string](s T) bool {
	trimmed := TrimSpace(s)
	return IsZero(trimmed)
}

func FindNonEmpty[T ~string](items ...T) (T, bool) {
	return slicer.FindFn(func(item T) bool {
		return !IsEmpty(item)
	}, items...)
}
