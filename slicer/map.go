package slicer

func Map[T any, U any](items []T, fn func(T) U) []U {
	var us []U

	for _, t := range items {
		u := fn(t)
		us = append(us, u)
	}

	return us
}

func MapErr[T any, U any](items []T, fn func(T) (U, error)) ([]U, error) {
	var us []U

	for _, i := range items {
		u, err := fn(i)

		if err != nil {
			return us, err
		}

		us = append(us, u)
	}

	return us, nil
}
