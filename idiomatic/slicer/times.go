package slicer

func Times[T any](count int, elems ...T) []T {
	fn := func(_, _ int, t T) T { return t }
	return TimesFn(fn, count, elems...)
}

func TimesFn[T any, U any](fn func(i, j int, t T) U, count int, elems ...T) []U {
	fn2 := func(i, j int, t T) (U, error) {
		return fn(i, j, t), nil
	}

	results, _ := TimesFnErr(fn2, count, elems...)
	return results
}

func TimesFnErr[T any, U any](fn func(i, j int, t T) (U, error), count int, elems ...T) ([]U, error) {
	results := []U{}
	var result U
	var err error

	for i := 0; i < count; i++ {
		for j, elem := range elems {
			result, err = fn(i, j, elem)

			if err != nil {
				break
			}

			results = append(results, result)
		}

		if err != nil {
			break
		}
	}

	return results, err
}
