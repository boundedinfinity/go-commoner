package slicer

// Deduplicate[T] creates a copy of elems with duplicates removed
func Deduplicate[T comparable](elems ...T) []T {
	fn := func(_ int, elem T) T {
		return elem
	}

	return DeduplicateFn(fn, elems...)
}

// DeduplicateFn[T, K] creates a copy of elems with duplicates removed based on the result of the
// fn(int, T) function
//
// The fn(int, T) K functions takes:
//   - int is the index of the current element
//   - T is the current element
//   - K the comparable type to deduplicate
func DeduplicateFn[T any, C comparable](fn func(int, T) C, elems ...T) []T {
	if fn == nil {
		return []T{}
	}

	fn2 := func(i int, elem T) (C, error) {
		k := fn(i, elem)
		return k, nil
	}

	results, _ := DeduplicateFnErr(fn2, elems...)

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
func DeduplicateFnErr[T any, C comparable](fn func(int, T) (C, error), elems ...T) ([]T, error) {
	m := make(map[C]T)
	var err error
	var key C
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
