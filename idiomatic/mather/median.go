package mather

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func Median[N Number](numbers ...N) N {
	sorted := slicer.Sort(numbers...)
	i := len(sorted) / 2
	return sorted[i]
}

func MedianOk[N Number](numbers ...N) (N, bool) {
	if len(numbers) == 0 {
		var zero N
		return zero, false
	}

	return Median(numbers...), true
}

func MedianFunc[E any, N Number](fn func(E) N, numbers ...E) N {
	sorted := slicer.Map(fn, numbers...)
	i := len(sorted) / 2
	return sorted[i]
}

func MedianFuncOk[E any, N Number](fn func(E) N, numbers ...E) (N, bool) {
	if len(numbers) == 0 {
		var zero N
		return zero, false
	}

	return MedianFunc(fn, numbers...), true
}
