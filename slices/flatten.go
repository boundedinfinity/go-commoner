package slices

func Flatten[T any](vss ...[]T) []T {
	var o []T

	for _, vs := range vss {
		o = append(o, vs...)
	}

	return o
}

func FlattenFn[T any, U any](fn func(vs []T) []U, vss ...[]T) []U {
	var o []U

	for _, vs := range vss {
		us := fn(vs)
		o = append(o, us...)
	}

	return o
}
