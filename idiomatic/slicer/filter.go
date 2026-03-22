package slicer

func Filter[T comparable](t T, elems ...T) []T {
	var result []T

	for _, elem := range elems {
		if elem == t {
			result = append(result, elem)
		}
	}

	return result
}

func FilterFn[T any](fn func(T) bool, elems ...T) []T {
	if fn == nil {
		return []T{}
	}

	var result []T

	for _, elem := range elems {
		if fn(elem) {
			result = append(result, elem)
		}
	}

	return result
}

func FilterFnI[T any](fn func(int, T) bool, elems ...T) []T {
	if fn == nil {
		return []T{}
	}

	var result []T

	for i, elem := range elems {
		if fn(i, elem) {
			result = append(result, elem)
		}
	}

	return result
}

func FilterFnErr[T any](fn func(T) (bool, error), elems ...T) ([]T, error) {
	var result []T

	if fn == nil {
		return result, nil
	}

	for _, elem := range elems {
		ok, err := fn(elem)

		if err != nil {
			return result, err
		}

		if ok {
			result = append(result, elem)
		}
	}

	return result, nil
}

func FilterFnErrI[T any](fn func(int, T) (bool, error), elems ...T) ([]T, error) {
	var result []T

	if fn == nil {
		return result, nil
	}

	for i, elem := range elems {
		ok, err := fn(i, elem)

		if err != nil {
			return result, err
		}

		if ok {
			result = append(result, elem)
		}
	}

	return result, nil
}
