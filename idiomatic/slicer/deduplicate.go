package slicer

// Deduplicate[T] creates a copy of elems with duplicates removed
func Deduplicate[T comparable](elems ...T) []T {
	m := make(map[T]bool)
	var results []T

	for _, elem := range elems {
		m[elem] = true
	}

	for elem := range m {
		results = append(results, elem)
	}

	return results
}

// DeduplicateFn[T, K] creates a copy of elems with duplicates removed based on the result of the
// fn(int, T) function
//
// The fn(int, T) K functions takes:
//   - int is the index of the current element
//   - T is the current element
//   - K the comparable type to deduplicate
func DeduplicateBy[T any, C comparable](fn func(T) C, elems ...T) []T {
	var results []T

	if fn == nil {
		return results
	}

	m := make(map[C]T)

	for _, elem := range elems {
		m[fn(elem)] = elem
	}

	for _, elem := range m {
		results = append(results, elem)
	}

	return results
}

// DeduplicateFn[T, K] creates a copy of elems with duplicates removed based on the result of the
// fn(int, T) function
//
// The fn(int, T) K functions takes:
//   - int is the index of the current element
//   - T is the current element
//   - K the comparable type to deduplicate
func DeduplicateByI[T any, C comparable](fn func(int, T) C, elems ...T) []T {
	var results []T

	if fn == nil {
		return results
	}

	m := make(map[C]T)

	for i, elem := range elems {
		m[fn(i, elem)] = elem
	}

	for _, elem := range m {
		results = append(results, elem)
	}

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
func DeduplicateByErr[T any, C comparable](fn func(T) (C, error), elems ...T) ([]T, error) {
	var results []T

	if fn == nil {
		return results, nil
	}

	m := make(map[C]T)

	var c C
	var err error

	for _, elem := range elems {
		if c, err = fn(elem); err != nil {
			break
		} else {
			m[c] = elem
		}
	}

	for _, elem := range m {
		results = append(results, elem)
	}

	return results, err
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
func DeduplicateByErrI[T any, C comparable](fn func(int, T) (C, error), elems ...T) ([]T, error) {
	var results []T

	if fn == nil {
		return results, nil
	}

	m := make(map[C]T)

	var c C
	var err error

	for i, elem := range elems {
		if c, err = fn(i, elem); err != nil {
			break
		} else {
			m[c] = elem
		}
	}

	for _, elem := range m {
		results = append(results, elem)
	}

	return results, err
}
