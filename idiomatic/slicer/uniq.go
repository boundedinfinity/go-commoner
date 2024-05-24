package slicer

func Uniq[T comparable](vs ...T) []T {
	o := []T{}
	m := map[T]bool{}

	for _, v := range vs {
		if _, ok := m[v]; !ok {
			m[v] = true
		}
	}

	for v := range m {
		o = append(o, v)
	}

	return o
}

func UniqFn[T any, U comparable](fn func(T) U, vs ...T) []T {
	o := []T{}
	m := map[U]T{}

	for _, v := range vs {
		c := fn(v)

		if _, ok := m[c]; !ok {
			m[c] = v
		}
	}

	for _, v := range m {
		o = append(o, v)
	}

	return o
}
