package slicer

func Reduce[T any, R any](fn func(int, R, T) R, initial R, elems ...T) R {
	fn2 := func(i int, result R, elem T) (R, error) {
		return fn(i, result, elem), nil
	}

	result, _ := ReduceErr(fn2, initial, elems...)
	return result
}

func ReduceErr[T any, R any](fn func(int, R, T) (R, error), initial R, elems ...T) (R, error) {
	result := initial
	var err error

	for i, elem := range elems {
		result, err = fn(i, result, elem)

		if err != nil {
			break
		}
	}

	return result, err
}
