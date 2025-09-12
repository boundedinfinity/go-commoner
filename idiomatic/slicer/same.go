package slicer

func Same[T comparable](as, bs []T) bool {
	if len(as) != len(bs) {
		return false
	}

	vc := make(map[T]int, len(as))

	for _, a := range as {
		vc[a]++
	}

	for _, b := range bs {
		vc[b]++
	}

	for _, c := range vc {
		if c%2 != 0 {
			return false
		}
	}

	return true
}

func SameBy[T any, C comparable](fn func(T) C, as, bs []T) bool {
	if len(as) != len(bs) {
		return false
	}

	vc := make(map[C]int, len(as))

	for _, a := range as {
		vc[fn(a)]++
	}

	for _, b := range bs {
		vc[fn(b)]++
	}

	for _, c := range vc {
		if c%2 != 0 {
			return false
		}
	}

	return true
}

func SameByI[T any, C comparable](fn func(int, T) C, as, bs []T) bool {
	if fn == nil {
		return false
	}

	if len(as) != len(bs) {
		return false
	}

	vc := make(map[C]int, len(as))

	for i, a := range as {
		vc[fn(i, a)]++
	}

	for i, b := range bs {
		vc[fn(i, b)]++
	}

	for _, c := range vc {
		if c%2 != 0 {
			return false
		}
	}

	return true
}

func SamByErr[T any, C comparable](fn func(T) (C, error), as, bs []T) (bool, error) {
	if len(as) != len(bs) {
		return false, nil
	}

	vc := make(map[C]int, len(as))

	for _, a := range as {
		if c, err := fn(a); err != nil {
			return false, err
		} else {
			vc[c]++
		}
	}

	for _, b := range bs {
		if c, err := fn(b); err != nil {
			return false, err
		} else {
			vc[c]++
		}
	}

	for _, c := range vc {
		if c%2 != 0 {
			return false, nil
		}
	}

	return true, nil
}

func SamByIErr[T any, C comparable](fn func(int, T) (C, error), as, bs []T) (bool, error) {
	if len(as) != len(bs) {
		return false, nil
	}

	vc := make(map[C]int, len(as))

	for i, a := range as {
		if c, err := fn(i, a); err != nil {
			return false, err
		} else {
			vc[c]++
		}
	}

	for i, b := range bs {
		if c, err := fn(i, b); err != nil {
			return false, err
		} else {
			vc[c]++
		}
	}

	for _, c := range vc {
		if c%2 != 0 {
			return false, nil
		}
	}

	return true, nil
}
