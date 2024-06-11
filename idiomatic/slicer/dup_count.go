package slicer

type DupCountResult[T any] struct {
	Count int
	Item  T
}

func DupCount[T comparable](elems ...T) []DupCountResult[T] {
	fn := func(elem T) T {
		return elem
	}

	return DupCountFn(fn, elems...)
}

func DupCountFn[T any, K comparable](fn func(T) K, elems ...T) []DupCountResult[T] {
	fn2 := func(elem T) (K, error) {
		k := fn(elem)
		return k, nil
	}

	results, _ := DupCountFnErr(fn2, elems...)

	return results
}

func DupCountFnErr[T any, K comparable](fn func(T) (K, error), elems ...T) ([]DupCountResult[T], error) {
	m := make(map[K]*DupCountResult[T])
	var err error
	var key K
	var results []DupCountResult[T]

	for _, elem := range elems {
		if key, err = fn(elem); err != nil {
			break
		}

		if _, ok := m[key]; !ok {
			m[key] = &DupCountResult[T]{
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
