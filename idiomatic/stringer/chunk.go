package stringer

func Chunk[T ~string](size int, s T) []string {
	var x = string(s)
	var chunks []string
	var last = 0

	for i := size; i < len(s); i += size {
		chunks = append(chunks, x[last:i])
		last = i
	}

	chunks = append(chunks, x[last:])

	return chunks
}
