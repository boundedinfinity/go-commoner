package stringer

import (
	"strings"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func Contains[T ~string, S ~string](s S, substr T) bool {
	return ContainsAny(s, substr)
}

func ContainsIgnoreCase[T ~string, S ~string](s S, substr T) bool {
	return ContainsAnyIgnoreCase(ToLower(s), ToLower(substr))
}

func ContainsAny[T ~string, S ~string](s S, elems ...T) bool {
	ns := string(s)

	for _, elem := range elems {
		if strings.Contains(ns, string(elem)) {
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
		func(_ int, elem T) string { return ToLower(elem) },
		elems...,
	)
	return ContainsAny(ToLower(s), lowers...)
}

func ContainsNoneIgnoreCase[T ~string, S ~string](s S, elems ...T) bool {
	return !ContainsAnyIgnoreCase(s, elems...)
}
