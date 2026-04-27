package mather

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func Sum[N Number](numbers ...N) N {
	fn := func(_ int, a N, v N) N { return a + v }
	return slicer.Reduce(fn, 0, numbers...)
}

func SumOk[N Number](numbers ...N) (N, bool) {
	if len(numbers) == 0 {
		var zero N
		return zero, false
	}

	return Sum(numbers...), true
}

func SumFunc[E any, N Number](fn func(E) N, numbers ...E) N {
	wrap := func(_ int, a N, e E) N { return a + fn(e) }
	return slicer.Reduce(wrap, 0, numbers...)
}

func SumFuncOk[E any, N Number](fn func(E) N, numbers ...E) (N, bool) {
	if len(numbers) == 0 {
		var zero N
		return zero, false
	}

	return SumFunc(fn, numbers...), true
}
