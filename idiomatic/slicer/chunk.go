package slicer

func Chunk[T any](size int, items ...T) [][]T {
	var buffer []T

	fn := func(current [][]T, item T) [][]T {
		if len(buffer) >= size {
			current = append(current, buffer)
			buffer = []T{}
		}

		buffer = append(buffer, item)
		return current
	}

	return append(FoldLeft([][]T{}, fn, items...), buffer)
}
