package slicer

func Flatten[T any](elems ...[]T) []T {
	var result []T

	for _, elem := range elems {
		result = append(result, elem...)
	}

	return result
}
