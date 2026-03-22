package slicer

func Duplicate[T any](count int, elems ...T) []T {
	results := []T{}

	for range count {
		results = append(results, elems...)
	}

	return results
}
