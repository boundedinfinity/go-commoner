package slicer

func Uniq[C comparable](vs ...C) []C {
	wrap := func(t C) C {
		return t
	}

	return UniqFn(wrap, vs...)
}

func UniqFn[T any, C comparable](fn func(T) C, vs ...T) []T {
	wrap := func(t T) (C, error) {
		return fn(t), nil
	}

	o, _ := UniqFnErr(wrap, vs...)

	return o
}

func UniqFnErr[T any, C comparable](fn func(T) (C, error), vs ...T) ([]T, error) {
	o := []T{}
	m := map[C]T{}

	for _, v := range vs {
		c, err := fn(v)

		if err != nil {
			return o, err
		}

		if _, ok := m[c]; !ok {
			m[c] = v
		}
	}

	for _, v := range m {
		o = append(o, v)
	}

	return o, nil
}
