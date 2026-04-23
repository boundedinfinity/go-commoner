package slicer

func Unique[C comparable](elems ...C) []C {
	m := map[C]int{}

	for i, elem := range elems {
		m[elem] = i
	}

	var results []C

	for elem := range m {
		results = append(results, elem)
	}

	return results
}

func UniqueFunc[T any, C comparable](fn func(T) C, elems ...T) []T {
	m := map[C]T{}

	for _, elem := range elems {
		m[fn(elem)] = elem
	}

	var results []T

	for _, elem := range m {
		results = append(results, elem)
	}

	return results
}

func UniqueByI[T any, C comparable](fn func(int, T) C, elems ...T) []T {
	m := map[C]T{}

	for i, elem := range elems {
		m[fn(i, elem)] = elem
	}

	var results []T

	for _, elem := range m {
		results = append(results, elem)
	}

	return results
}

func UniqueByErr[T any, C comparable](fn func(T) (C, error), elems ...T) ([]T, error) {
	results := []T{}
	m := map[C]T{}

	for _, elem := range elems {
		c, err := fn(elem)

		if err != nil {
			return results, err
		}

		if _, ok := m[c]; !ok {
			m[c] = elem
		}
	}

	for _, elem := range m {
		results = append(results, elem)
	}

	return results, nil
}

func UniqueByErrI[T any, C comparable](fn func(int, T) (C, error), elems ...T) ([]T, error) {
	results := []T{}
	m := map[C]T{}

	for i, elem := range elems {
		c, err := fn(i, elem)

		if err != nil {
			return results, err
		}

		if _, ok := m[c]; !ok {
			m[c] = elem
		}
	}

	for _, elem := range m {
		results = append(results, elem)
	}

	return results, nil
}
