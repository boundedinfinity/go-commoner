package slicer

import (
	"errors"
	"fmt"
)

// References
//      - https://lodash.com/docs/4.17.15#partition
//      - https://www.scala-lang.org/api/3.4.2/scala/collection/ArrayOps.html#partition-243

// PartitionErr splits the elements into two slices based on the result of the fn function.
//
// The fn function takes the index of the current index of the element of elems, and the current value of the element
// in elems.  It should return the computed compariable value extract from the element.
//
// The first slice contains the elements for which the fn function returns a true value.  The second slice contains the
// elements for which the fn function returns a false value.
//
// for which the fn function returns a true value.  The second slice contains the elems for which
// the fn fucntion returns a false value.
func Partition[T any](fn func(int, T) bool, elems ...T) ([]T, []T) {
	fn2 := func(i int, elem T) (bool, error) {
		k := fn(i, elem)
		return k, nil
	}

	p1, p2, _ := PartitionErr(fn2, elems...)

	return p1, p2
}

var (
	ErrPartition   = errors.New("parition error")
	errPartitionFn = func(i int, err error) error {
		return fmt.Errorf("%w occured at index %v : %w", err, i, ErrPartition)
	}
)

// PartitionErr splits the elements into two slices based on the result of the fn function.
//
// The fn function takes the index of the current index of the element of elems, and the current value of the element
// in elems.  It should return the computed compariable value extract from the element or an error if computed
// compariable value can't be computed.
//
// The first slice contains the elements for which the fn function returns a true value.  The second slice contains the
// elements for which the fn function returns a false value.
//
// If the fn function returns an error, processing through elements is stopped and the error is returned.  The values
// of the partitions will contain any partitioned data up to but not including the value when the error occured.
func PartitionErr[T any](fn func(int, T) (bool, error), elems ...T) ([]T, []T, error) {
	var p1, p2 []T
	var err error
	var ok bool

	for i, elem := range elems {
		ok, err = fn(i, elem)

		if err != nil {
			err = errPartitionFn(i, err)
			break
		}

		if ok {
			p1 = append(p1, elem)
		} else {
			p2 = append(p2, elem)
		}
	}

	return p1, p2, err
}
