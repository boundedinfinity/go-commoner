package slicer

func Flatten[T any](items ...[]T) []T {
	var o []T

	for _, item := range items {
		o = append(o, item...)
	}

	return o
}

func FlattenFn[T any, U any](fn func(vs []T) []U, items ...[]T) []U {
	var o []U

	for _, item := range items {
		us := fn(item)
		o = append(o, us...)
	}

	return o
}

func FlattenFnErr[T any, U any](fn func(vs []T) ([]U, error), items ...[]T) ([]U, error) {
	var o []U

	for _, item := range items {
		us, err := fn(item)

		if err != nil {
			return o, err
		}

		o = append(o, us...)
	}

	return o, nil
}
