package slices

func Flatten[T any](items ...[]T) []T {
	var o []T

	for _, item := range items {
		o = append(o, item...)
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
