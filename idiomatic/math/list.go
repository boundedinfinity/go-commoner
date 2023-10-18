package math

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func Sum[V Numbers](numbers ...V) V {
	fn := func(a V, v V) V {
		return a + v
	}

	return slicer.Reduce(0, numbers, fn)
}

func Mean[V Numbers](numbers ...V) V {
	l := len(numbers)

	if l == 0 {
		return 0
	}

	sum := Sum(numbers...)
	mean := sum / V(l)

	return mean
}

func Median[V Numbers](numbers ...V) V {
	sorted := slicer.Sort(numbers...)
	i := Ceil(len(sorted) / 2)
	return numbers[i]
}

func MinOf[V Numbers](numbers ...V) V {
	fn := func(a V, v V) V {
		if v < a {
			return v
		} else {
			return a
		}
	}

	return slicer.Reduce(slicer.Head(numbers), numbers, fn)
}

func MaxOf[V Numbers](numbers ...V) V {
	fn := func(a V, v V) V {
		if v > a {
			return v
		} else {
			return a
		}
	}

	return slicer.Reduce(slicer.Head(numbers), numbers, fn)
}
