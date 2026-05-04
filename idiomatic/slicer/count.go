package slicer

type CountResult[T any] struct {
	Count int
	Item  T
}

func Count[T comparable](elems ...T) []CountResult[T] {
	m := make(map[T]*CountResult[T])

	for _, elem := range elems {
		if _, ok := m[elem]; !ok {
			m[elem] = &CountResult[T]{
				Count: 1,
				Item:  elem,
			}
		} else {
			m[elem].Count++
		}
	}

	results := make([]CountResult[T], 0)

	for _, elem := range m {
		results = append(results, *elem)
	}

	return results
}

func CountFunc[T any, C comparable](fn func(T) C, elems ...T) []CountResult[T] {
	results := make([]CountResult[T], 0)

	if fn == nil {
		return results
	}

	m := make(map[C]*CountResult[T])

	for _, elem := range elems {
		c := fn(elem)

		if _, ok := m[c]; !ok {
			m[c] = &CountResult[T]{
				Count: 1,
				Item:  elem,
			}
		} else {
			m[c].Count++
		}
	}

	for _, elem := range m {
		results = append(results, *elem)
	}

	return results
}

func CountFuncI[T any, C comparable](fn func(int, T) C, elems ...T) []CountResult[T] {
	results := make([]CountResult[T], 0)

	if fn == nil {
		return results
	}

	m := make(map[C]*CountResult[T])

	for i, elem := range elems {
		c := fn(i, elem)

		if _, ok := m[c]; !ok {
			m[c] = &CountResult[T]{
				Count: 1,
				Item:  elem,
			}
		} else {
			m[c].Count++
		}
	}

	for _, elem := range m {
		results = append(results, *elem)
	}

	return results
}

func CountFuncErr[T any, C comparable](fn func(T) (C, error), elems ...T) ([]CountResult[T], error) {
	m := make(map[C]*CountResult[T])
	var err error
	var key C
	var results []CountResult[T]

	if fn == nil {
		return results, err
	}

	for _, elem := range elems {
		if key, err = fn(elem); err != nil {
			break
		}

		if _, ok := m[key]; !ok {
			m[key] = &CountResult[T]{
				Item:  elem,
				Count: 0,
			}
		}

		m[key].Count++
	}

	if err != nil {
		return results, err
	}

	for _, result := range m {
		results = append(results, *result)
	}

	return results, nil
}

func CountFuncErrI[T any, C comparable](fn func(int, T) (C, error), elems ...T) ([]CountResult[T], error) {
	m := make(map[C]*CountResult[T])
	var err error
	var key C
	var results []CountResult[T]

	if fn == nil {
		return results, err
	}

	for i, elem := range elems {
		if key, err = fn(i, elem); err != nil {
			break
		}

		if _, ok := m[key]; !ok {
			m[key] = &CountResult[T]{
				Item:  elem,
				Count: 0,
			}
		}

		m[key].Count++
	}

	if err != nil {
		return results, err
	}

	for _, result := range m {
		results = append(results, *result)
	}

	return results, nil
}
