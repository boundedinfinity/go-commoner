package slicer

import (
	"errors"
	"fmt"
)

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// References
//      - https://lodash.com/docs/4.17.15#groupBy
//      - https://www.scala-lang.org/api/3.4.2/scala/collection/ArrayOps.html#groupBy-f68
// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Group creates a map composed of keys generated from the fn function and the value is the value passed to the fn
// function which genereated they key.
func Group[T any, K comparable](fn func(int, T) K, elems ...T) map[K][]T {
	fn2 := func(i int, elem T) (K, error) {
		key := fn(i, elem)
		return key, nil
	}

	results, _ := GroupErr(fn2, elems...)

	return results
}

var (
	ErrGroup   = errors.New("parition error")
	errGroupFn = func(i int, err error) error {
		return fmt.Errorf("%w occured at index %v : %w", err, i, ErrGroup)
	}
)

// GroupErr creates a map composed of keys generated from the fn function and the value is the value passed to the fn
// function which genereated they key.
//
// If the fn function returns an error, processing through elements is stopped and the error is returned.  The values
// of the grouped elements will contain any elements processed up to but not including the element when the error
// occured.
func GroupErr[T any, K comparable](fn func(int, T) (K, error), elems ...T) (map[K][]T, error) {
	groups := make(map[K][]T)
	var err error
	var key K

	for i, elem := range elems {
		key, err = fn(i, elem)

		if err != nil {
			err = errGroupFn(i, err)
			break
		}

		if _, ok := groups[key]; !ok {
			groups[key] = []T{}
		}

		groups[key] = append(groups[key], elem)
	}

	return groups, err
}
