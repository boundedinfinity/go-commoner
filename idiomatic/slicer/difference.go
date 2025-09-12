package slicer

func Difference[T comparable](as, bs []T) []T {
	var results []T
	bm := map[T]bool{}

	for _, b := range bs {
		bm[b] = true
	}

	for _, a := range as {
		if ok := bm[a]; !ok {
			results = append(results, a)
		}
	}

	return results
}

func DifferenceBy[T any, C comparable](fn func(T) C, as, bs []T) []T {
	var results []T

	if fn == nil {
		return results
	}

	bm := map[C]T{}

	for _, b := range bs {
		bm[fn(b)] = b
	}

	for _, a := range as {
		if _, ok := bm[fn(a)]; !ok {
			results = append(results, a)
		}
	}

	return results
}

func DifferenceByErr[T any, C comparable](fn func(T) (C, error), as, bs []T) ([]T, error) {
	var results []T
	var err error
	var c C

	if fn == nil {
		return results, err
	}

	bm := map[C]T{}

	for _, b := range bs {
		if c, err = fn(b); err != nil {
			break
		} else {
			bm[c] = b
		}

	}

	if err == nil {
		for _, a := range as {
			if _, err = fn(a); err != nil {
				break
			} else {
				results = append(results, a)
			}
		}
	}

	return results, err
}
