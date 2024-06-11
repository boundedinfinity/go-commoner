package slicer

import "fmt"

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// References
//      - https://lodash.com/docs/4.17.15#chunk
// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Chunk creates a slice where each element is a slice of elems of size length.
//
// If elems can't be split evenly, the final element will contain the remaining elems.
//
// Example:
//
//	elems := []int{1, 2, 3, 4, 5}
//	size := 2
//	chunks := Chunk(size, elems...)
//	fmt.Println(chunks)
//	// [[1 2] [3 4] [5]]
func Chunk[T any](size int, elems ...T) [][]T {
	var output [][]T
	var buffer []T

	for _, elem := range elems {
		buffer = append(buffer, elem)

		if len(buffer) >= size {
			output = append(output, buffer)
			buffer = []T{}
		}
	}

	if len(buffer) > 0 {
		output = append(output, buffer)
	}

	fmt.Println(output)

	return output
}
