package slicer

func Flatten[T any](elems ...[]T) []T {
	var o []T

	for _, elem := range elems {
		o = append(o, elem...)
	}

	return o
}

func FlattenFn[T any, U any](fn func(vs []T) []U, elems ...[]T) []U {
	var o []U

	for _, elem := range elems {
		us := fn(elem)
		o = append(o, us...)
	}

	return o
}

func FlattenFnErr[T any, U any](fn func(vs []T) ([]U, error), elems ...[]T) ([]U, error) {
	var o []U

	for _, elem := range elems {
		us, err := fn(elem)

		if err != nil {
			return o, err
		}

		o = append(o, us...)
	}

	return o, nil
}
