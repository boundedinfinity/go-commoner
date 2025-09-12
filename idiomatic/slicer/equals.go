package slicer

func Equals[T comparable](as, bs []T) bool {
	if len(as) != len(bs) {
		return false
	}

	m := map[T]bool{}

	for _, a := range as {
		m[a] = false
	}

	for _, b := range bs {
		m[b] = true
	}

	for _, ok := range m {
		if !ok {
			return false
		}
	}

	return true
}

func EqualsBy[T any, C comparable](fn func(T) C, as, bs []T) bool {
	if len(as) != len(bs) {
		return false
	}

	if fn == nil {
		return false
	}

	m := map[C]bool{}

	for _, a := range as {
		m[fn(a)] = false
	}

	for _, b := range bs {
		m[fn(b)] = true
	}

	for _, ok := range m {
		if !ok {
			return false
		}
	}

	return true
}

func EqualsByI[T any, C comparable](fn func(int, T) C, as, bs []T) bool {
	if len(as) != len(bs) {
		return false
	}

	if fn == nil {
		return false
	}

	m := map[C]bool{}

	for i, a := range as {
		m[fn(i, a)] = false
	}

	for i, b := range bs {
		m[fn(i, b)] = true
	}

	for _, ok := range m {
		if !ok {
			return false
		}
	}

	return true
}

func EqualsFnErr[T any, C comparable](fn func(T) (C, error), as, bs []T) (bool, error) {
	if len(as) != len(bs) {
		return false, nil
	}

	if fn == nil {
		return false, nil
	}

	var err error
	var c C

	m := map[C]bool{}

	for _, a := range as {
		if c, err = fn(a); err != nil {
			break
		} else {
			m[c] = false
		}
	}

	if err == nil {
		for _, b := range bs {
			if c, err = fn(b); err != nil {
				break
			} else {
				m[c] = true
			}
		}
	}

	if err == nil {
		for _, ok := range m {
			if !ok {
				return false, nil
			}
		}
	}

	return true, nil
}
