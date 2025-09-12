package slicer

type CountDuplicatesResult[T any] struct {
	Count int
	Item  T
}

func CountDuplicates[T comparable](elems ...T) []CountDuplicatesResult[T] {
	m := make(map[T]*CountDuplicatesResult[T])

	for _, elem := range elems {
		if _, ok := m[elem]; !ok {
			m[elem] = &CountDuplicatesResult[T]{
				Count: 1,
				Item:  elem,
			}
		} else {
			m[elem].Count++
		}
	}

	results := make([]CountDuplicatesResult[T], 0)

	for _, elem := range m {
		results = append(results, *elem)
	}

	return results
}

func CountDuplicatesBy[T any, C comparable](by func(T) C, elems ...T) []CountDuplicatesResult[T] {
	results := make([]CountDuplicatesResult[T], 0)

	if by == nil {
		return results
	}

	m := make(map[C]*CountDuplicatesResult[T])

	for _, elem := range elems {
		c := by(elem)

		if _, ok := m[c]; !ok {
			m[c] = &CountDuplicatesResult[T]{
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

func CountDuplicatesByErr[T any, C comparable](by func(T) (C, error), elems ...T) ([]CountDuplicatesResult[T], error) {
	m := make(map[C]*CountDuplicatesResult[T])
	var err error
	var key C
	var results []CountDuplicatesResult[T]

	if by == nil {
		return results, err
	}

	for _, elem := range elems {
		if key, err = by(elem); err != nil {
			break
		}

		if _, ok := m[key]; !ok {
			m[key] = &CountDuplicatesResult[T]{
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
