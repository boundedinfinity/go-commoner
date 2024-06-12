package slicer

func Times[T any](count int, elems ...T) []T {
	results := []T{}

	for i := 0; i < count; i++ {
		results = append(results, elems...)
	}

	return results
}
