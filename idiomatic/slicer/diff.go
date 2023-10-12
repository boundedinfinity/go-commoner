package slicer

func Diff[T comparable](as, bs []T) []T {
	fn := func(a, b T) bool {
		return a == b
	}

	return DiffFn(as, bs, fn)
}

func DiffFn[T any](as, bs []T, fn func(T, T) bool) []T {
	var d []T

	for _, a := range as {
		var found bool

		for _, b := range bs {
			if found = fn(a, b); found {
				break
			}
		}

		if !found {
			d = append(d, a)
		}
	}

	return d
}

func DiffFnErr[T any](as, bs []T, fn func(T, T) (bool, error)) ([]T, error) {
	var d []T
	var err error

	for _, a := range as {
		var found bool

		for _, b := range bs {
			if found, err = fn(a, b); found || err != nil {
				break
			}
		}

		if !found {
			d = append(d, a)
		}
	}

	return d, err
}
