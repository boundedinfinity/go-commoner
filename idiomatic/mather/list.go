package mather

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/mather/types"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func Sum[V types.Numbers](numbers ...V) V {
	fn := func(a V, v V) V {
		return a + v
	}

	return slicer.Reduce(0, numbers, fn)
}

func Mean[V types.Numbers](numbers ...V) V {
	l := len(numbers)

	if l == 0 {
		return 0
	}

	sum := Sum(numbers...)
	mean := sum / V(l)

	return mean
}

func Median[V types.Numbers](numbers ...V) V {
	sorted := slicer.Sort(numbers...)
	i := Ceil(len(sorted) / 2)
	return numbers[i]
}

func MinOf[V types.Numbers](numbers ...V) V {
	fn := func(a V, v V) V {
		if v < a {
			return v
		} else {
			return a
		}
	}

	head, _ := slicer.Head(numbers...)
	return slicer.Reduce(head, numbers, fn)
}

func MaxOf[V types.Numbers](numbers ...V) V {
	fn := func(a V, v V) V {
		if v > a {
			return v
		} else {
			return a
		}
	}

	head, _ := slicer.Head(numbers...)
	return slicer.Reduce(head, numbers, fn)
}
