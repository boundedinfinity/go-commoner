package slicer

// Dedup[T] creates a copy of elems with duplicates removed
func Dedup[T comparable](elems ...T) []T {
	fn := func(_ int, elem T) T {
		return elem
	}

	return DedupFn(fn, elems...)
}

// DedupFn[T, K] creates a copy of elems with duplicates removed based on the result of the
// fn(int, T) function
//
// The fn(int, T) K functions takes:
//   - int is the index of the current element
//   - T is the current element
//   - K the comparable type to deduplicate
func DedupFn[T any, K comparable](fn func(int, T) K, elems ...T) []T {
	if fn == nil {
		return []T{}
	}

	fn2 := func(i int, elem T) (K, error) {
		k := fn(i, elem)
		return k, nil
	}

	results, _ := DedupFnErr(fn2, elems...)

	return results
}

// DedupFn[T, K] creates a copy of elems with duplicates removed based on the result of the
// fn(int, T) function
//
// The fn(int, T) K functions takes:
//   - int is the index of the current element
//   - T is the current element
//   - K the comparable type to deduplicate
//
// If the fn function returns an error, processing through elems is stopped and the error is returned.
func DedupFnErr[T any, K comparable](fn func(int, T) (K, error), elems ...T) ([]T, error) {
	m := make(map[K]T)
	var err error
	var key K
	var results []T

	if fn == nil {
		return results, err
	}

	for i, elem := range elems {
		if key, err = fn(i, elem); err != nil {
			break
		} else {
			if _, ok := m[key]; !ok {
				m[key] = elem
				results = append(results, elem)
			}
		}
	}

	if err != nil {
		return results, err
	}

	return results, nil
}
