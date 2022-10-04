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
		found := false

		for _, b := range bs {
			if fn(a, b) {
				found = true
				break
			}
		}

		if !found {
			d = append(d, a)
		}
	}

	return d
}
