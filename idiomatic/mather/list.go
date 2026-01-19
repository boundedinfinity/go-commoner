package mather

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/mather/types"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func Sum[N types.Number](numbers ...N) N {
	fn := func(_ int, a N, v N) N { return a + v }
	return slicer.Reduce(fn, 0, numbers...)
}

func SumFn[E any, N types.Number](fn func(E) N, numbers ...E) N {
	wrap := func(_ int, a N, e E) N { return a + fn(e) }
	return slicer.Reduce(wrap, 0, numbers...)
}

func Mean[N types.Number](numbers ...N) N {
	l := len(numbers)

	if l == 0 {
		return 0
	}

	sum := Sum(numbers...)
	mean := sum / N(l)

	return mean
}

func MeanFn[E any, N types.Number](fn func(E) N, numbers ...E) N {
	l := len(numbers)

	if l == 0 {
		return 0
	}

	sum := SumFn(fn, numbers...)
	mean := sum / N(l)

	return mean
}

func Median[N types.Number](numbers ...N) N {
	sorted := slicer.Sort(numbers...)
	i := Ceil(len(sorted) / 2)
	return numbers[i]
}

func MedianFn[E any, O types.Ordered, N types.Number](fn func(E) O, numbers ...E) E {
	sorted := slicer.SortBy(fn, numbers...)
	i := Ceil(len(sorted) / 2)
	return numbers[i]
}

func MinOf[V types.Number](numbers ...V) V {
	fn := func(_ int, a V, v V) V {
		if v < a {
			return v
		} else {
			return a
		}
	}

	head, _ := slicer.Head(numbers...)
	return slicer.Reduce(fn, head, numbers...)
}

func MaxOf[V types.Number](numbers ...V) V {
	fn := func(_ int, a V, v V) V {
		if v > a {
			return v
		} else {
			return a
		}
	}

	head, _ := slicer.Head(numbers...)
	return slicer.Reduce(fn, head, numbers...)
}
