package slicer

import "github.com/boundedinfinity/go-commoner/idiomatic/mapper"

func Distinct[T comparable](xs []T) []T {
	m := map[T]int{}

	for _, x := range xs {
		m[x]++
	}

	return mapper.Keys(m)
}
