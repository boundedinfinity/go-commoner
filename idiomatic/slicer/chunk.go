package slicer

import (
	"iter"
)

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// References
//      - https://lodash.com/docs/4.17.15#chunk
// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Chunk creates a slice where each element is a slice of elems of size length.
//
// If elems can't be split evenly, the final element will contain the remaining elems.
// If count < 1 the results are return as if count is 1
//
// Example:
//
//	elems := []int{1, 2, 3, 4, 5}
//	size := 2
//	chunks := Chunk(size, elems...)
//	fmt.Println(chunks)
//	// [[1 2] [3 4] [5]]
func Chunk[T any](size int, elems ...T) [][]T {
	var chunks [][]T

	for chunk := range chunkIter(size, elems...) {
		chunks = append(chunks, chunk)
	}

	return chunks
}

func chunkIter[T any](size int, elems ...T) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		var chunk []T

		for _, elem := range elems {
			if len(chunk) >= size {
				if !yield(chunk) {
					return
				}

				chunk = []T{}
			}

			chunk = append(chunk, elem)
		}
	}
}
