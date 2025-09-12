package slicer

func ignoreErr[T any](elem T, _ error) T {
	return elem
}

func wrapIndexFunc[T any, R any](fn func(T) R) func(int, T) R {
	return func(_ int, elem T) R { return fn(elem) }
}

func wrapIndexAndErr[T any, R any](fn func(T) (R, error)) func(int, T) (R, error) {
	return func(_ int, elem T) (R, error) { return fn(elem) }
}
