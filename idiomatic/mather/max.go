package mather

func Max[V Number](numbers ...V) V {
	max, _ := MaxOk(numbers...)
	return max
}

func MaxOk[V Number](numbers ...V) (V, bool) {
	if len(numbers) == 0 {
		var zero V
		return zero, false
	}

	max := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] > max {
			max = numbers[i]
		}
	}

	return max, true
}

func MaxFunc[T any, V Number](fn func(T) V, numbers ...T) V {
	max, _ := MaxFuncOk(fn, numbers...)
	return max
}

func MaxFuncOk[T any, V Number](fn func(T) V, numbers ...T) (V, bool) {
	if len(numbers) == 0 {
		var zero V
		return zero, false
	}

	max := fn(numbers[0])

	for i := 1; i < len(numbers); i++ {
		if fn(numbers[i]) > max {
			max = fn(numbers[i])
		}
	}

	return max, true
}
