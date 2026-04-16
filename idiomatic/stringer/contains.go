package stringer

import (
	"strings"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func Contains[T ~string, S ~string](s S, substr T) bool {
	return strings.Contains(string(s), string(substr))
}

func ContainsIgnoreCase[T ~string, S ~string](s S, substr T) bool {
	return strings.EqualFold(string(s), string(substr))
}

func ContainsAny[T ~string, S ~string](s S, elems ...T) bool {
	normal := string(s)

	for _, elem := range elems {
		if strings.Contains(normal, string(elem)) {
			return true
		}
	}

	return false
}

func ContainsNone[T ~string, S ~string](s S, elems ...T) bool {
	return !ContainsAny(s, elems...)
}

func ContainsAnyIgnoreCase[T ~string, S ~string](s S, elems ...T) bool {
	lowers := slicer.Map(
		func(elem T) string { return ToLower(elem) },
		elems...,
	)
	return ContainsAny(ToLower(s), lowers...)
}

func ContainsNoneIgnoreCase[T ~string, S ~string](s S, elems ...T) bool {
	return !ContainsAnyIgnoreCase(s, elems...)
}
