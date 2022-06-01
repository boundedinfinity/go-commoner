package slices

import "github.com/boundedinfinity/commons/maps"

func Distinct[T comparable](xs []T) []T {
	m := map[T]int{}

	for _, x := range xs {
		m[x]++
	}

	return maps.Keys(m)
}
