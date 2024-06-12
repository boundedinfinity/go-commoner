package slicer

type CountDuplicatesResult[T any] struct {
	Count int
	Item  T
}

func CountDuplicates[T comparable](elems ...T) []CountDuplicatesResult[T] {
	fn := func(elem T) T {
		return elem
	}

	return CountDuplicatesFn(fn, elems...)
}

func CountDuplicatesFn[T any, K comparable](fn func(T) K, elems ...T) []CountDuplicatesResult[T] {
	if fn == nil {
		return []CountDuplicatesResult[T]{}
	}

	fn2 := func(elem T) (K, error) {
		k := fn(elem)
		return k, nil
	}

	results, _ := CountDuplicatesFnErr(fn2, elems...)

	return results
}

func CountDuplicatesFnErr[T any, K comparable](fn func(T) (K, error), elems ...T) ([]CountDuplicatesResult[T], error) {
	m := make(map[K]*CountDuplicatesResult[T])
	var err error
	var key K
	var results []CountDuplicatesResult[T]

	if fn == nil {
		return results, err
	}

	for _, elem := range elems {
		if key, err = fn(elem); err != nil {
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
