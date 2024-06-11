package slicer

func Map[T any, R any](fn func(int, T) R, elems ...T) []R {
	fn2 := func(i int, elem T) (R, error) {
		return fn(i, elem), nil
	}

	results, _ := MapErr(fn2, elems...)
	return results
}

func MapErr[T any, R any](fn func(int, T) (R, error), elems ...T) ([]R, error) {
	var results []R
	var result R
	var err error

	for i, elem := range elems {
		result, err = fn(i, elem)

		if err != nil {
			break
		}

		results = append(results, result)
	}

	return results, err
}

func MapPipe[T any, R any](fns []func(int, T) R, elems ...T) []R {
	var results []R

	for _, fn := range fns {
		results = Map(fn, elems...)
	}

	return results
}

func MapPipeErr[T any, R any](fns []func(int, T) (R, error), elems ...T) ([]R, error) {
	var results []R
	var err error

	for _, fn := range fns {
		results, err = MapErr(fn, elems...)
	}

	return results, err
}
