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

func ContainsAny[T ~string, S ~string](s S, items ...T) bool {
	ns := string(s)

	for _, item := range items {
		if strings.Contains(ns, string(item)) {
			return true
		}
	}

	return false
}

func ContainsNone[T ~string, S ~string](s S, items ...T) bool {
	return !ContainsAny(s, items...)
}

func ContainsAnyIgnoreCase[T ~string, S ~string](s S, items ...T) bool {
	lowers := slicer.Map(
		func(item T) string { return ToLower(item) },
		items...,
	)
	return ContainsAny(ToLower(s), lowers...)
}

func ContainsNoneIgnoreCase[T ~string, S ~string](s S, items ...T) bool {
	return !ContainsAnyIgnoreCase(s, items...)
}
