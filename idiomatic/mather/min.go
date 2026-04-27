package mather

func Min[V Number](numbers ...V) V {
	min, _ := MinOk(numbers...)

	return min
}

func MinOk[V Number](numbers ...V) (V, bool) {
	if len(numbers) == 0 {
		var zero V
		return zero, false
	}

	min := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] < min {
			min = numbers[i]
		}
	}

	return min, true
}

func MinFunc[T any, V Number](fn func(T) V, numbers ...T) V {
	min, _ := MinFuncOk(fn, numbers...)
	return min
}

func MinFuncOk[T any, V Number](fn func(T) V, numbers ...T) (V, bool) {
	if len(numbers) == 0 {
		var zero V
		return zero, false
	}

	min := fn(numbers[0])

	for i := 1; i < len(numbers); i++ {
		if fn(numbers[i]) < min {
			min = fn(numbers[i])
		}
	}

	return min, true
}
