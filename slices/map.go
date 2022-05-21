package slices

func MapErr[T any, U any](ts []T, fn func(T) (U, error)) ([]U, error) {
	var us []U

	for _, i := range ts {
		u, err := fn(i)

		if err != nil {
			return us, err
		}

		us = append(us, u)
	}

	return us, nil
}

func Map[T any, U any](ts []T, fn func(T) U) []U {
	var us []U

	for _, t := range ts {
		u := fn(t)
		us = append(us, u)
	}

	return us
}
