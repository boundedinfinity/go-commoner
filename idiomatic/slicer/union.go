package slicer

func Union[T comparable](as []T, bs []T) ([]T, error) {
	wrap := func(t T) T {
		return t
	}

	return UnionFn(wrap, as, bs)
}

func UnionFn[T any, C comparable](fn func(T) C, as []T, bs []T) ([]T, error) {
	wrap := func(t T) (C, error) {
		return fn(t), nil
	}

	return UnionFnErr(wrap, as, bs)
}

func UnionFnErr[T any, C comparable](fn func(T) (C, error), as []T, bs []T) ([]T, error) {
	cs := append([]T{}, as...)
	cs = append(cs, bs...)

	return UniqFnErr(fn, cs...)
}
