package slicer

func Intersection[T comparable](as []T, bs []T) []T {
	fn2 := func(elem T) T { return elem }
	return IntersectionFn(fn2, as, bs)
}

func IntersectionFn[T any, C comparable](fn func(T) C, as []T, bs []T) []T {
	fn2 := func(elem T) (C, error) {
		return fn(elem), nil
	}

	result, _ := IntersectionFnErr(fn2, as, bs)
	return result
}

func IntersectionFnErr[T any, C comparable](fn func(T) (C, error), as []T, bs []T) ([]T, error) {
	fn2 := func(i int, elem T) (C, error) {
		return fn(elem)
	}

	results := []T{}

	uas, err := UniqByErrI(fn2, as...)
	if err != nil {
		return results, err
	}

	am, err := GroupErr(fn2, uas...)
	if err != nil {
		return results, err
	}

	ubs, err := UniqByErrI(fn2, bs...)
	if err != nil {
		return results, err
	}

	bm, err := GroupErr(fn2, ubs...)
	if err != nil {
		return results, err
	}

	for k, as := range am {
		if _, ok := bm[k]; ok {
			results = append(results, as[0])
		}
	}

	return results, nil
}

func HasIntersection[T comparable](as []T, bs []T) bool {
	return len(Intersection(as, bs)) > 0
}

func HasIntersectionFn[T any, C comparable](fn func(T) C, as []T, bs []T) bool {
	return len(IntersectionFn(fn, as, bs)) > 0
}

func HasIntersectionFnErr[T any, C comparable](fn func(T) (C, error), as []T, bs []T) (bool, error) {
	results, err := IntersectionFnErr(fn, as, bs)
	return len(results) > 0, err
}
