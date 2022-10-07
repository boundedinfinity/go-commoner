package stringer

func Remove[T ~string, U ~string](s T, remove U) string {
	return Replace[T, U, string](s, remove, "")
}
