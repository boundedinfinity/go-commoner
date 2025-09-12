package slicer

func Exclude[T comparable](t T, elems ...T) []T {
	var result []T

	for _, elem := range elems {
		if elem != t {
			result = append(result, elem)
		}
	}

	return result
}

func ExcludeBy[T any](by func(T) bool, elems ...T) []T {
	if by == nil {
		return []T{}
	}

	var result []T

	for _, elem := range elems {
		if !by(elem) {
			result = append(result, elem)
		}
	}

	return result
}

func ExcludeByI[T any](by func(int, T) bool, elems ...T) []T {
	if by == nil {
		return []T{}
	}

	var result []T

	for i, elem := range elems {
		if !by(i, elem) {
			result = append(result, elem)
		}
	}

	return result
}

func ExcludeByErrI[T any](by func(int, T) (bool, error), elems ...T) ([]T, error) {
	var result []T

	if by == nil {
		return result, nil
	}

	for i, elem := range elems {
		ok, err := by(i, elem)

		if err != nil {
			return result, err
		}

		if !ok {
			result = append(result, elem)
		}
	}

	return result, nil
}
