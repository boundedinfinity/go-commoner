package slicer

type DupCountResult[T any] struct {
	Count int
	Item  T
}

func DupCount[T comparable](items ...T) []DupCountResult[T] {
	fn := func(item T) T {
		return item
	}

	return DupCountFn(fn, items...)
}

func DupCountFn[T any, K comparable](fn func(T) K, items ...T) []DupCountResult[T] {
	fn2 := func(item T) (K, error) {
		k := fn(item)
		return k, nil
	}

	results, _ := DupCountFnErr(fn2, items...)

	return results
}

func DupCountFnErr[T any, K comparable](fn func(T) (K, error), items ...T) ([]DupCountResult[T], error) {
	m := make(map[K]*DupCountResult[T])
	var err error
	var key K
	var results []DupCountResult[T]

	for _, item := range items {
		if key, err = fn(item); err != nil {
			break
		}

		if _, ok := m[key]; !ok {
			m[key] = &DupCountResult[T]{
				Item:  item,
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
