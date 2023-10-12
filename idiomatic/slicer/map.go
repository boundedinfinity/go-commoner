package slicer

func Map[T any, U any](fn func(T) U, items ...T) []U {
	var us []U

	for _, t := range items {
		u := fn(t)
		us = append(us, u)
	}

	return us
}

func MapI[T any, U any](fn func(int, T) U, items ...T) []U {
	var us []U

	for i, item := range items {
		u := fn(i, item)
		us = append(us, u)
	}

	return us
}

func MapErr[T any, U any](fn func(T) (U, error), items ...T) ([]U, error) {
	var us []U
	var u U
	var err error

	for _, item := range items {
		if u, err = fn(item); err != nil {
			break
		} else {
			us = append(us, u)
		}
	}

	return us, err
}

func MapErrI[T any, U any](fn func(int, T) (U, error), items ...T) ([]U, error) {
	var us []U
	var u U
	var err error

	for i, item := range items {
		if u, err = fn(i, item); err != nil {
			break
		} else {
			us = append(us, u)
		}
	}

	return us, err
}
