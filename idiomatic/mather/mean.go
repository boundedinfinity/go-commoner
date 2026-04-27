package mather

func Mean[N Number](numbers ...N) N {
	l := len(numbers)

	if l == 0 {
		return 0
	}

	return Sum(numbers...) / N(l)
}

func MeanOk[N Number](numbers ...N) (N, bool) {
	l := len(numbers)

	if l == 0 {
		var zero N
		return zero, false
	}

	return Sum(numbers...) / N(l), true
}

func MeanFunc[E any, N Number](fn func(E) N, numbers ...E) N {
	l := len(numbers)

	if l == 0 {
		return 0
	}

	return SumFunc(fn, numbers...) / N(l)
}

func MeanFuncOk[E any, N Number](fn func(E) N, numbers ...E) (N, bool) {
	l := len(numbers)

	if l == 0 {
		var zero N
		return zero, false
	}

	return SumFunc(fn, numbers...) / N(l), true
}
